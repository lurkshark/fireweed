package evals_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/workflow"
)

func releaseRefs(t *testing.T, root string) string {
	t.Helper()
	return git(t, root, "for-each-ref", "--format=%(refname) %(objectname)")
}

func Test_FLOW_084_ReleaseOwnedClaimPreservesDirtyWorkAndRecordedState(t *testing.T) {
	root := fixture(t, "", "", "### new:release-work\nRelease work MUST remain.\n", "", "")
	started := start(t, root)
	if started.Run == "" || started.Phase != "evals" || started.ClaimOID == "" || len(started.Change.Added) != 1 || started.Change.Added[0].ID != "ALP-003" {
		t.Fatalf("Start did not assign runtime state and ID: %#v", started)
	}
	changePath := "changes/0001-manual.md"
	change := readFile(t, root, changePath)
	if !strings.Contains(change, "run: "+started.Run) || !strings.Contains(change, "phase: "+started.Phase) || !strings.Contains(change, "### ALP-003") {
		t.Fatalf("Start did not record runtime state and assigned ID: %q", change)
	}
	write(t, root, changePath, strings.Replace(change, "Use the revised policy.", "Use the revised policy.\nWorker draft remains.", 1))
	write(t, root, "packages/alpha/internal/worker.txt", "unfinished worker work\n")
	write(t, root, "worker-untracked.txt", "untracked worker work\n")
	git(t, root, "update-ref", "refs/fireweed/fixture-unrelated", head(t, root))
	beforeHead := head(t, root)
	beforeStatus := git(t, root, "status", "--porcelain=v1", "--untracked-files=all")
	beforeRefs := releaseRefs(t, root)
	beforeChange := readFile(t, root, changePath)
	beforeTracked := readFile(t, root, "packages/alpha/internal/worker.txt")
	beforeUntracked := readFile(t, root, "worker-untracked.txt")
	if err := workflow.Release(root, "manual"); err != nil {
		t.Fatalf("Release owned claim: %v", err)
	}
	if got := head(t, root); got != beforeHead {
		t.Errorf("HEAD moved: %s -> %s", beforeHead, got)
	}
	if got := git(t, root, "status", "--porcelain=v1", "--untracked-files=all"); got != beforeStatus {
		t.Errorf("worktree status changed: %q -> %q", beforeStatus, got)
	}
	for path, want := range map[string]string{
		changePath:                           beforeChange,
		"packages/alpha/internal/worker.txt": beforeTracked,
		"worker-untracked.txt":               beforeUntracked,
	} {
		if got := readFile(t, root, path); got != want {
			t.Errorf("%s changed: %q -> %q", path, want, got)
		}
	}
	if _, err := workflow.Current(root, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P02:") {
		t.Errorf("released claim remains current: %v", err)
	}
	wantRefs := strings.Replace(beforeRefs, "refs/fireweed/claims/manual "+started.ClaimOID+"\n", "", 1)
	if wantRefs == beforeRefs {
		t.Fatalf("Start claim absent from refs: %q", beforeRefs)
	}
	if got := releaseRefs(t, root); got != wantRefs {
		t.Errorf("Release changed refs beyond owned claim: got %q, want %q", got, wantRefs)
	}
}

func Test_FLOW_085_ReleaseMissingAndForeignClaimsPreserveRefs(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	before := releaseRefs(t, root)
	if err := workflow.Release(root, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P02:") {
		t.Errorf("missing claim Release = %v, want P02", err)
	}
	if got := releaseRefs(t, root); got != before {
		t.Errorf("missing claim changed refs: %q -> %q", before, got)
	}

	other := filepath.Join(t.TempDir(), "linked")
	git(t, root, "worktree", "add", "-b", "linked", other, "HEAD")
	owner := start(t, root)
	write(t, root, "owner-work.txt", "owner's unfinished work\n")
	write(t, other, "foreign-work.txt", "other worktree's unfinished work\n")
	ownerHead, foreignHead := head(t, root), head(t, other)
	ownerStatus := git(t, root, "status", "--porcelain=v1", "--untracked-files=all")
	foreignStatus := git(t, other, "status", "--porcelain=v1", "--untracked-files=all")
	ownerChange := readFile(t, root, "changes/0001-manual.md")
	foreignChange := readFile(t, other, "changes/0001-manual.md")
	before = releaseRefs(t, root)
	if err := workflow.Release(other, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P02:") {
		t.Errorf("foreign claim Release = %v, want P02", err)
	}
	if got := releaseRefs(t, root); got != before {
		t.Errorf("foreign claim changed refs: %q -> %q", before, got)
	}
	if head(t, root) != ownerHead || head(t, other) != foreignHead ||
		git(t, root, "status", "--porcelain=v1", "--untracked-files=all") != ownerStatus ||
		git(t, other, "status", "--porcelain=v1", "--untracked-files=all") != foreignStatus ||
		readFile(t, root, "changes/0001-manual.md") != ownerChange ||
		readFile(t, other, "changes/0001-manual.md") != foreignChange ||
		readFile(t, root, "owner-work.txt") != "owner's unfinished work\n" ||
		readFile(t, other, "foreign-work.txt") != "other worktree's unfinished work\n" {
		t.Error("foreign Release changed HEAD, status or working files")
	}
	current, err := workflow.Current(root, "manual")
	if err != nil || current.ClaimOID != owner.ClaimOID || current.Run != owner.Run {
		t.Errorf("owner's exact claim changed: current = %#v, error = %v", current, err)
	}
}

func Test_FLOW_086_ReleaseAcceptsCurrentChangeSelectors(t *testing.T) {
	for _, selector := range []struct {
		name, value string
	}{
		{"path", "changes/0001-manual.md"},
		{"filename", "0001-manual.md"},
		{"slug", "manual"},
	} {
		t.Run(selector.name, func(t *testing.T) {
			root := fixture(t, "", "", "", "", "")
			started := start(t, root)
			selected := selector.value
			current, err := workflow.Current(root, selected)
			if err != nil || current.ClaimOID != started.ClaimOID {
				t.Fatalf("Current selector %q = %#v, %v", selected, current, err)
			}
			if err := workflow.Release(root, selected); err != nil {
				t.Fatalf("Release selector %q: %v", selected, err)
			}
			if _, err := workflow.Current(root, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P02:") {
				t.Errorf("selector %q left claim current: %v", selected, err)
			}
		})
	}
}
