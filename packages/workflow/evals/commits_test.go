package evals_test

import (
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/workflow"
)

const pendingCommitChange = "changes/0001-manual.md"
const archivedCommitChange = "changes/archive/0001-manual.md"

func commitPhase(t *testing.T, root, path, old, next, message string) string {
	t.Helper()
	before := readFile(t, root, path)
	var after string
	if old == "" {
		after = strings.Replace(before, "module: alpha\n", "module: alpha\nphase: "+next+"\n", 1)
	} else {
		after = strings.Replace(before, "phase: "+old+"\n", "phase: "+next+"\n", 1)
	}
	if after == before {
		t.Fatalf("fixture phase %q not found in %s", old, path)
	}
	write(t, root, path, after)
	commit(t, root, message)
	return head(t, root)
}

func assertCommits(t *testing.T, root, selector string, want []string) {
	t.Helper()
	got, err := workflow.Commits(root, selector)
	if err != nil {
		t.Fatalf("Commits(%q): %v", selector, err)
	}
	for _, oid := range got {
		if !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(oid) {
			t.Errorf("Commits(%q) returned non-full SHA-1 OID %q", selector, oid)
		}
	}
	if !(len(got) == 0 && len(want) == 0) && !reflect.DeepEqual(got, want) {
		t.Errorf("Commits(%q) = %v, want newest-first %v", selector, got, want)
	}
}

func Test_FLOW_087_FLOW_088_CommitsFollowPendingAndArchivePhaseContents(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	author := head(t, root)
	write(t, root, "unrelated.txt", "unrelated work\n")
	commit(t, root, "phase: evals")
	unrelated := head(t, root)
	evals := commitPhase(t, root, pendingCommitChange, "", "evals", "same message")
	change := readFile(t, root, pendingCommitChange)
	write(t, root, pendingCommitChange, strings.Replace(change, "Users need a revised alpha value.", "Users need a revised alpha value with detail.", 1))
	commit(t, root, "phase: audit")
	noPhaseChange := head(t, root)
	audit := commitPhase(t, root, pendingCommitChange, "evals", "audit", "same message")
	impl := commitPhase(t, root, pendingCommitChange, "audit", "impl", "same message")
	git(t, root, "checkout", "-b", "side")
	unreachable := commitPhase(t, root, pendingCommitChange, "impl", "applied", "same message")
	git(t, root, "checkout", "main")
	got, err := workflow.Commits(root, pendingCommitChange)
	if err != nil {
		t.Fatalf("Commits after side branch: %v", err)
	}
	for _, oid := range got {
		if oid == unreachable {
			t.Errorf("Commits included unreachable phase change %s", unreachable)
		}
	}
	for _, selector := range []string{pendingCommitChange, "0001-manual.md", "manual"} {
		assertCommits(t, root, selector, []string{impl, audit, evals})
	}

	if err := os.MkdirAll(filepath.Join(root, "changes", "archive"), 0o755); err != nil {
		t.Fatalf("create archive directory: %v", err)
	}
	git(t, root, "mv", pendingCommitChange, archivedCommitChange)
	archived := readFile(t, root, archivedCommitChange)
	write(t, root, archivedCommitChange, strings.Replace(archived, "phase: impl\n", "phase: applied\n", 1))
	commit(t, root, "same message")
	applied := head(t, root)
	for _, selector := range []string{archivedCommitChange, "0001-manual.md", "manual"} {
		assertCommits(t, root, selector, []string{applied, impl, audit, evals})
	}
	for _, excluded := range []string{author, unrelated, noPhaseChange} {
		if excluded == evals || excluded == audit || excluded == impl || excluded == applied {
			t.Fatalf("fixture OID collision: %s", excluded)
		}
	}
}

func Test_FLOW_087_CommitsUseFirstParentOfMerge(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	evals := commitPhase(t, root, pendingCommitChange, "", "evals", "same message")
	git(t, root, "branch", "alternate", evals)
	mainPhase := commitPhase(t, root, pendingCommitChange, "evals", "impl", "same message")
	git(t, root, "checkout", "alternate")
	branchPhase := commitPhase(t, root, pendingCommitChange, "evals", "audit", "same message")
	git(t, root, "checkout", "main")
	git(t, root, "merge", "--no-ff", "-s", "ours", "-m", "phase: applied", "alternate")
	merge := head(t, root)
	got, err := workflow.Commits(root, "manual")
	if err != nil {
		t.Fatalf("Commits after merge: %v", err)
	}
	want := map[string]bool{evals: true, mainPhase: true, branchPhase: true}
	if len(got) != len(want) {
		t.Errorf("Commits after merge = %v, want three phase changes", got)
	}
	for _, oid := range got {
		if !want[oid] {
			t.Errorf("Commits included %s; merge %s has no phase difference from its first parent", oid, merge)
		}
		delete(want, oid)
	}
	for oid := range want {
		t.Errorf("Commits omitted reachable phase change %s", oid)
	}
}

func Test_FLOW_088_CommitsRejectMissingAndAmbiguousSelectors(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	for _, selector := range []string{"absent", "changes/absent.md"} {
		if _, err := workflow.Commits(root, selector); err == nil || !strings.HasPrefix(err.Error(), "U02:") {
			t.Errorf("Commits(%q) error = %v, want U02", selector, err)
		}
	}
	write(t, root, "changes/0002-manual.md", readFile(t, root, pendingCommitChange))
	commit(t, root, "same message")
	if _, err := workflow.Commits(root, "manual"); err == nil || !strings.HasPrefix(err.Error(), "U02:") {
		t.Errorf("ambiguous slug error = %v, want U02", err)
	}
	assertCommits(t, root, pendingCommitChange, []string{})
}
