package evals_test

import (
	"context"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/workflow"
)

func abandonAppFailure(t *testing.T, root string, args []string, wantStatus int, code string) {
	t.Helper()
	jsonMode := len(args) > 0 && args[len(args)-1] == "--json"
	status, out, errOut := runSavedAttemptApp(t, appNestedDir(t, root), args...)
	message := commandFailure(t, status, wantStatus, out, errOut, jsonMode)
	if jsonMode {
		message = assertFailedJSON(t, out, code)
	} else if !strings.HasPrefix(strings.TrimSpace(message), code+":") {
		t.Errorf("diagnostic=%q; want %s cause", message, code)
	}
	if !corrective(message) {
		t.Errorf("diagnostic %q must suggest a correction", message)
	}
}

func Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON(t *testing.T) {
	for _, tc := range []struct {
		name, selector string
		json           bool
		owned          bool
	}{
		{"plain-queued", "changes/0001-manual.md", false, false},
		{"json-owned", "manual", true, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			bin := "fixture-start-does-not-run-hooks"
			if tc.owned {
				bin = buildAppBinary(t)
			}
			root := lifecycleCLIProject(t, bin, false)
			cliGit(t, root, "config", "user.name", "Fixture")
			cliGit(t, root, "config", "user.email", "fixture@example.invalid")
			if tc.owned {
				appPut(t, root, ".fireweed/emit-new", "yes\n")
				cliCommit(t, root, "fixture projected eval")
				started, err := workflow.Start(root, "manual")
				if err != nil || started.ClaimOID == "" || started.Phase != "evals" {
					t.Fatalf("fixture Start: state=%#v err=%v", started, err)
				}
				appPut(t, root, "packages/alpha/extra.go", "package alpha\n")
				advanced, err := workflow.Done(context.Background(), workflow.Request{Root: root, Change: "manual", Bin: bin, Timeout: 2 * time.Second})
				if err != nil || advanced.State.Phase != "impl" {
					t.Fatalf("fixture evals Done: phase=%q err=%v", advanced.State.Phase, err)
				}
				if current, err := workflow.Current(root, "manual"); err != nil || current.ClaimOID == "" || current.Phase != "impl" {
					t.Fatalf("fixture lost owned claim: phase=%q claim=%q err=%v", current.Phase, current.ClaimOID, err)
				}
				if dirty := cliGit(t, root, "status", "--porcelain=v1", "--untracked-files=all"); dirty != "" {
					t.Fatalf("fixture is dirty before abandon: %s", dirty)
				}
			}
			parent := cliGit(t, root, "rev-parse", "HEAD")
			args := []string{"abandon", tc.selector}
			if tc.json {
				args = append(args, "--json")
			}
			status, out, errOut := runSavedAttemptApp(t, appNestedDir(t, root), args...)
			if status != 0 || errOut != "" {
				t.Fatalf("abandon: exit=%d stdout=%q stderr=%q", status, out, errOut)
			}
			confirmation := out
			if tc.json {
				confirmation = commandStringData(t, out)
			}
			oid := cliGit(t, root, "rev-parse", "HEAD")
			if !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(oid) || oid == parent {
				t.Errorf("abandon HEAD=%q; want a new full object ID after %q", oid, parent)
			}
			if !strings.Contains(strings.ToLower(confirmation), "abandon") || !strings.Contains(confirmation, "manual") || !strings.Contains(confirmation, oid) {
				t.Errorf("confirmation %q must name abandon, selected change, and exact full HEAD %s", confirmation, oid)
			}
			if got := cliGit(t, root, "rev-parse", "HEAD^"); got != parent {
				t.Errorf("abandon commit parent=%s; want %s", got, parent)
			}
			if _, err := os.Stat(filepath.Join(root, "changes/0001-manual.md")); !os.IsNotExist(err) {
				t.Errorf("pending change remains after abandon: %v", err)
			}
			if dirty := cliGit(t, root, "status", "--porcelain=v1", "--untracked-files=all"); dirty != "" {
				t.Errorf("abandon left dirty worktree: %s", dirty)
			}
			if claim := cliGit(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/claims/manual"); claim != "" {
				t.Errorf("abandon retained owned claim: %s", claim)
			}
		})
	}
}

func Test_APP_098_APP_099_AbandonFailures(t *testing.T) {
	root := lifecycleCLIProject(t, "fixture-start-does-not-run-hooks", false)
	for _, tc := range []struct {
		name string
		args []string
		code string
	}{
		{"missing", []string{"abandon"}, "U01"},
		{"extra-json", []string{"abandon", "manual", "extra", "--json"}, "U01"},
		{"unsupported-option", []string{"abandon", "manual", "--bogus"}, "U01"},
		{"unknown-selector", []string{"abandon", "missing", "--json"}, "U02"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			abandonAppFailure(t, root, tc.args, 2, tc.code)
		})
	}

	configured := lifecycleCLIProject(t, "fixture-start-does-not-run-hooks", false)
	appPut(t, configured, "fireweed.yaml", "modules: [broken\n")
	abandonAppFailure(t, configured, []string{"abandon", "manual", "--json"}, 2, "C01")

	dirty := lifecycleCLIProject(t, "fixture-start-does-not-run-hooks", false)
	appPut(t, dirty, "stray.txt", "worker work\n")
	before := cliGit(t, dirty, "rev-parse", "HEAD")
	abandonAppFailure(t, dirty, []string{"abandon", "manual", "--json"}, 1, "P01")
	if got := cliGit(t, dirty, "rev-parse", "HEAD"); got != before {
		t.Errorf("dirty refusal moved HEAD: %s -> %s", before, got)
	}
	if body, err := os.ReadFile(filepath.Join(dirty, "stray.txt")); err != nil || string(body) != "worker work\n" {
		t.Errorf("dirty refusal changed worker file: %q, %v", body, err)
	}

	foreign := lifecycleCLIProject(t, "fixture-start-does-not-run-hooks", false)
	linked := filepath.Join(t.TempDir(), "linked")
	cliGit(t, foreign, "worktree", "add", "-b", "linked", linked, "HEAD")
	owner, err := workflow.Start(foreign, "manual")
	if err != nil || owner.ClaimOID == "" {
		t.Fatalf("fixture Start: state=%#v err=%v", owner, err)
	}
	abandonAppFailure(t, linked, []string{"abandon", "manual", "--json"}, 1, "P02")
	if current, err := workflow.Current(foreign, "manual"); err != nil || current.ClaimOID != owner.ClaimOID {
		t.Errorf("foreign refusal changed owner claim: state=%#v err=%v", current, err)
	}

	applied := lifecycleCLIProject(t, "fixture-start-does-not-run-hooks", false)
	if err := os.MkdirAll(filepath.Join(applied, "changes/archive"), 0755); err != nil {
		t.Fatal(err)
	}
	cliGit(t, applied, "mv", "changes/0001-manual.md", "changes/archive/0001-manual.md")
	commitsAppPhase(t, applied, "changes/archive/0001-manual.md", "", "applied")
	abandonAppFailure(t, applied, []string{"abandon", "manual"}, 1, "P01")
}
