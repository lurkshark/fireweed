package evals_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/workflow"
)

func releaseAppRefs(t *testing.T, root string) string {
	t.Helper()
	return cliGit(t, root, "for-each-ref", "--format=%(refname) %(objectname)")
}

func releaseAppFailure(t *testing.T, status, want int, out, errOut string, jsonMode bool, code string) {
	t.Helper()
	message := commandFailure(t, status, want, out, errOut, jsonMode)
	if jsonMode {
		message = assertFailedJSON(t, out, code)
	} else if !strings.HasPrefix(strings.TrimSpace(message), code+":") {
		t.Errorf("diagnostic=%q; want %s cause", message, code)
	}
	if !corrective(message) {
		t.Errorf("diagnostic %q must suggest a correction", message)
	}
}

func Test_APP_091_APP_092_APP_093_ReleaseOwnedClaimRetainsDirtyWork(t *testing.T) {
	for _, tc := range []struct {
		name, selector string
		json           bool
	}{
		{"plain-path", "changes/0001-manual.md", false},
		{"json-filename", "0001-manual.md", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := lifecycleCLIProject(t, "fixture-start-does-not-run-hooks", false)
			started, err := workflow.Start(root, "manual")
			if err != nil || started.ClaimOID == "" || started.Run == "" || started.Phase != "evals" || len(started.Change.Added) != 1 {
				t.Fatalf("fixture Start: state=%#v err=%v", started, err)
			}
			changePath := "changes/0001-manual.md"
			change, err := os.ReadFile(filepath.Join(root, changePath))
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(string(change), "run: "+started.Run) || !strings.Contains(string(change), "phase: evals") || !strings.Contains(string(change), "### "+started.Change.Added[0].ID) {
				t.Fatalf("Start did not record runtime state and assigned ID: %q", change)
			}
			wantChange := string(change) + "\nWorker draft remains.\n"
			appPut(t, root, changePath, wantChange)
			appPut(t, root, "packages/alpha/internal/worker.txt", "unfinished worker work\n")
			appPut(t, root, "worker-untracked.txt", "untracked worker work\n")
			cliGit(t, root, "update-ref", "refs/fireweed/fixture-unrelated", cliGit(t, root, "rev-parse", "HEAD"))
			beforeHead := cliGit(t, root, "rev-parse", "HEAD")
			beforeStatus := cliGit(t, root, "status", "--porcelain=v1", "--untracked-files=all")
			beforeRefs := releaseAppRefs(t, root)
			claimLine := "refs/fireweed/claims/manual " + started.ClaimOID
			if !strings.Contains(beforeRefs, claimLine) {
				t.Fatalf("fixture lacks owned claim in refs: %q", beforeRefs)
			}
			args := []string{"release", tc.selector}
			if tc.json {
				args = append(args, "--json")
			}
			status, out, errOut := runSavedAttemptApp(t, appNestedDir(t, root), args...)
			if status != 0 || errOut != "" {
				t.Fatalf("release: exit=%d stdout=%q stderr=%q", status, out, errOut)
			}
			confirmation := out
			if tc.json {
				confirmation = commandStringData(t, out)
			}
			if !strings.Contains(strings.ToLower(confirmation), "releas") || !strings.Contains(confirmation, "manual") {
				t.Errorf("missing release confirmation: %q", confirmation)
			}
			if got := cliGit(t, root, "rev-parse", "HEAD"); got != beforeHead {
				t.Errorf("release moved HEAD: %s -> %s", beforeHead, got)
			}
			if got := cliGit(t, root, "status", "--porcelain=v1", "--untracked-files=all"); got != beforeStatus {
				t.Errorf("release changed worktree status: %q -> %q", beforeStatus, got)
			}
			for path, want := range map[string]string{
				changePath:                           wantChange,
				"packages/alpha/internal/worker.txt": "unfinished worker work\n",
				"worker-untracked.txt":               "untracked worker work\n",
			} {
				body, err := os.ReadFile(filepath.Join(root, path))
				if err != nil || string(body) != want {
					t.Errorf("release changed %s: %q, %v", path, body, err)
				}
			}
			wantRefs := strings.TrimSpace(strings.Replace(beforeRefs, claimLine, "", 1))
			if got := releaseAppRefs(t, root); got != wantRefs {
				t.Errorf("release changed refs beyond owned claim: got %q, want %q", got, wantRefs)
			}
			if _, err := workflow.Current(root, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P02:") {
				t.Errorf("released claim remains current: %v", err)
			}
			status, out, errOut = runSavedAttemptApp(t, root, args...)
			releaseAppFailure(t, status, 1, out, errOut, tc.json, "P02")
			if got := releaseAppRefs(t, root); got != wantRefs {
				t.Errorf("repeated release changed refs: %q -> %q", wantRefs, got)
			}
		})
	}
}

func Test_APP_092_APP_093_ReleaseForeignClaimKeepsOwner(t *testing.T) {
	root := lifecycleCLIProject(t, "fixture-start-does-not-run-hooks", false)
	other := filepath.Join(t.TempDir(), "linked")
	cliGit(t, root, "worktree", "add", "-b", "linked", other, "HEAD")
	owner, err := workflow.Start(root, "manual")
	if err != nil || owner.ClaimOID == "" {
		t.Fatalf("fixture Start: state=%#v err=%v", owner, err)
	}
	appPut(t, root, "owner-work.txt", "owner's unfinished work\n")
	appPut(t, other, "foreign-work.txt", "other worktree's unfinished work\n")
	ownerHead, foreignHead := cliGit(t, root, "rev-parse", "HEAD"), cliGit(t, other, "rev-parse", "HEAD")
	ownerStatus := cliGit(t, root, "status", "--porcelain=v1", "--untracked-files=all")
	foreignStatus := cliGit(t, other, "status", "--porcelain=v1", "--untracked-files=all")
	refs := releaseAppRefs(t, root)
	ownerChange, err := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if err != nil {
		t.Fatal(err)
	}
	foreignChange, err := os.ReadFile(filepath.Join(other, "changes/0001-manual.md"))
	if err != nil {
		t.Fatal(err)
	}
	status, out, errOut := runSavedAttemptApp(t, appNestedDir(t, other), "release", "manual", "--json")
	releaseAppFailure(t, status, 1, out, errOut, true, "P02")
	if releaseAppRefs(t, root) != refs || cliGit(t, root, "rev-parse", "HEAD") != ownerHead || cliGit(t, other, "rev-parse", "HEAD") != foreignHead ||
		cliGit(t, root, "status", "--porcelain=v1", "--untracked-files=all") != ownerStatus ||
		cliGit(t, other, "status", "--porcelain=v1", "--untracked-files=all") != foreignStatus {
		t.Error("foreign release changed refs, HEAD or worktree status")
	}
	for path, want := range map[string]string{
		filepath.Join(root, "changes/0001-manual.md"):  string(ownerChange),
		filepath.Join(other, "changes/0001-manual.md"): string(foreignChange),
		filepath.Join(root, "owner-work.txt"):          "owner's unfinished work\n",
		filepath.Join(other, "foreign-work.txt"):       "other worktree's unfinished work\n",
	} {
		body, err := os.ReadFile(path)
		if err != nil || string(body) != want {
			t.Errorf("foreign release changed %s: %q, %v", path, body, err)
		}
	}
	current, err := workflow.Current(root, "manual")
	if err != nil || current.ClaimOID != owner.ClaimOID || current.Run != owner.Run || current.Phase != owner.Phase {
		t.Errorf("foreign release changed owner's claim: state=%#v err=%v", current, err)
	}
}

func Test_APP_092_APP_093_ReleaseUsageAndConfiguration(t *testing.T) {
	root := lifecycleCLIProject(t, "fixture-start-does-not-run-hooks", false)
	beforeHead := cliGit(t, root, "rev-parse", "HEAD")
	beforeRefs := releaseAppRefs(t, root)
	for _, tc := range []struct {
		name string
		args []string
		json bool
	}{
		{"missing-argument", []string{"release"}, false},
		{"extra-argument-json", []string{"release", "manual", "extra", "--json"}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			status, out, errOut := runSavedAttemptApp(t, root, tc.args...)
			releaseAppFailure(t, status, 2, out, errOut, tc.json, "U01")
		})
	}
	if got := cliGit(t, root, "rev-parse", "HEAD"); got != beforeHead || releaseAppRefs(t, root) != beforeRefs {
		t.Errorf("usage rejection changed HEAD or refs: %s -> %s", beforeHead, got)
	}
	appPut(t, root, "fireweed.yaml", "modules: [broken\n")
	status, out, errOut := runSavedAttemptApp(t, root, "release", "manual", "--json")
	releaseAppFailure(t, status, 2, out, errOut, true, "C01")
}
