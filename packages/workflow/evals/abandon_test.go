package evals_test

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func abandonGitIdentity(t *testing.T, root string) {
	t.Helper()
	git(t, root, "config", "user.name", "Fixture")
	git(t, root, "config", "user.email", "fixture@example.invalid")
}

func assertAbandonCommit(t *testing.T, root, parent, oid string) {
	t.Helper()
	if !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(oid) || oid != head(t, root) {
		t.Errorf("Abandon returned %q; HEAD is %s", oid, head(t, root))
	}
	if got := git(t, root, "rev-parse", "HEAD^"); got != parent {
		t.Errorf("Abandon parent = %s, want %s (one new commit)", got, parent)
	}
	if got := git(t, root, "status", "--porcelain=v1", "--untracked-files=all"); got != "" {
		t.Errorf("Abandon left dirty worktree: %s", got)
	}
	for _, path := range []string{pendingCommitChange, archivedCommitChange} {
		if _, err := os.Stat(filepath.Join(root, path)); !os.IsNotExist(err) {
			t.Errorf("change remains at %s: %v", path, err)
		}
	}
	// FLOW-094: the only change-document edit is deletion; no new phase is assigned.
	if diff := git(t, root, "diff", "--unified=0", parent, oid, "--", "changes"); strings.Contains(diff, "+phase:") {
		t.Errorf("Abandon assigned a phase in its commit: %s", diff)
	}
}

func assertAbandonRefusal(t *testing.T, root, selector, code string) {
	t.Helper()
	beforeHead := head(t, root)
	beforeTree := git(t, root, "rev-parse", "HEAD^{tree}")
	beforeStatus := git(t, root, "status", "--porcelain=v1", "--untracked-files=all")
	beforeRefs := releaseRefs(t, root)
	if oid, err := workflow.Abandon(root, selector); err == nil || !strings.HasPrefix(err.Error(), code+":") || oid != "" {
		t.Errorf("Abandon(%q) = %q, %v; want %s refusal without OID", selector, oid, err, code)
	}
	if got := head(t, root); got != beforeHead {
		t.Errorf("rejected Abandon moved HEAD: %s -> %s", beforeHead, got)
	}
	if got := git(t, root, "rev-parse", "HEAD^{tree}"); got != beforeTree {
		t.Errorf("rejected Abandon changed committed tree: %s -> %s", beforeTree, got)
	}
	if got := git(t, root, "status", "--porcelain=v1", "--untracked-files=all"); got != beforeStatus {
		t.Errorf("rejected Abandon changed worktree status: %q -> %q", beforeStatus, got)
	}
	if got := releaseRefs(t, root); got != beforeRefs {
		t.Errorf("rejected Abandon changed refs: %q -> %q", beforeRefs, got)
	}
}

func Test_FLOW_089_FLOW_091_FLOW_094_AbandonUnclaimedQueuedChange(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	abandonGitIdentity(t, root)
	parent := head(t, root)
	beforeSpec := readFile(t, root, "packages/alpha/SPEC.md")
	if commits, err := workflow.Commits(root, "manual"); err != nil || len(commits) != 0 {
		t.Fatalf("queued fixture has phase commits: %v, %v", commits, err)
	}
	oid, err := workflow.Abandon(root, "manual")
	if err != nil {
		t.Fatalf("Abandon unclaimed queued change: %v", err)
	}
	assertAbandonCommit(t, root, parent, oid)
	if got := git(t, root, "diff-tree", "--no-commit-id", "--name-status", "-r", oid); got != "D\t"+pendingCommitChange {
		t.Errorf("queued Abandon effects = %q, want change-file deletion only", got)
	}
	if got := readFile(t, root, "packages/alpha/SPEC.md"); got != beforeSpec {
		t.Errorf("queued Abandon changed applied spec: %q", got)
	}
}

func Test_FLOW_089_FLOW_091_FLOW_094_AbandonCleanOwnedSteppedChange(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	abandonGitIdentity(t, root)
	initialSpec := readFile(t, root, "packages/alpha/SPEC.md")
	started := start(t, root)
	write(t, root, "packages/alpha/evals/step_test.go", "committed eval effect\n")
	evals, err := done(root)
	if err != nil || evals.State.Phase != "audit" {
		t.Fatalf("fixture evals Done = %#v, %v", evals, err)
	}
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	audited, err := done(root)
	if err != nil || audited.State.Phase != "impl" {
		t.Fatalf("fixture audit Done = %#v, %v", audited, err)
	}
	phaseCommits, err := workflow.Commits(root, "manual")
	if err != nil || len(phaseCommits) != 2 {
		t.Fatalf("fixture phase commits = %v, %v; want evals and audit", phaseCommits, err)
	}
	if got := readFile(t, root, "packages/alpha/evals/step_test.go"); got != "committed eval effect\n" {
		t.Fatalf("fixture lacks committed phase effect: %q", got)
	}
	before := head(t, root)
	current, err := workflow.Current(root, "manual")
	if err != nil || current.ClaimOID == "" || current.Run != started.Run {
		t.Fatalf("fixture lost owned claim: %#v, %v", current, err)
	}
	oid, err := workflow.Abandon(root, "manual")
	if err != nil {
		t.Fatalf("Abandon owned stepped change: %v", err)
	}
	assertAbandonCommit(t, root, before, oid)
	if got := git(t, root, "diff-tree", "--no-commit-id", "--name-status", "-r", oid); got != "D\t"+pendingCommitChange+"\nD\tpackages/alpha/evals/step_test.go" {
		t.Errorf("stepped Abandon effects = %q; want pending change and committed eval removed", got)
	}
	if got := readFile(t, root, "packages/alpha/SPEC.md"); got != initialSpec {
		t.Errorf("stepped Abandon changed applied spec: %q", got)
	}
	if claim := git(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/claims/manual"); claim != "" {
		t.Errorf("owned claim remains: %s", claim)
	}
	if _, err := workflow.Current(root, "manual"); err == nil || !strings.HasPrefix(err.Error(), "U02:") {
		t.Errorf("released claim is still current: %v", err)
	}
}

func Test_FLOW_090_AbandonRejectsDirtyWorkingTree(t *testing.T) {
	for _, kind := range []string{"tracked", "untracked"} {
		t.Run(kind, func(t *testing.T) {
			root := fixture(t, "", "", "", "", "")
			if kind == "tracked" {
				write(t, root, "SYSTEM.md", "dirty tracked content\n")
			} else {
				write(t, root, "stray.txt", "dirty untracked content\n")
			}
			assertAbandonRefusal(t, root, "manual", "P01")
			if kind == "tracked" && readFile(t, root, "SYSTEM.md") != "dirty tracked content\n" {
				t.Error("dirty tracked content changed")
			}
			if kind == "untracked" && readFile(t, root, "stray.txt") != "dirty untracked content\n" {
				t.Error("dirty untracked content changed")
			}
		})
	}
}

func Test_FLOW_091_AbandonRejectsForeignClaim(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	abandonGitIdentity(t, root)
	other := filepath.Join(t.TempDir(), "linked")
	git(t, root, "worktree", "add", "-b", "linked", other, "HEAD")
	started := start(t, root)
	// The linked worktree is clean and sees the same pending change at HEAD.
	if status := git(t, other, "status", "--porcelain=v1", "--untracked-files=all"); status != "" {
		t.Fatalf("foreign worktree is dirty: %s", status)
	}
	ownerHead := head(t, root)
	ownerChange := readFile(t, root, pendingCommitChange)
	assertAbandonRefusal(t, other, "manual", "P02")
	if head(t, root) != ownerHead || readFile(t, root, pendingCommitChange) != ownerChange {
		t.Error("foreign Abandon changed owner's worktree")
	}
	if current, err := workflow.Current(root, "manual"); err != nil || current.ClaimOID != started.ClaimOID {
		t.Errorf("foreign Abandon replaced owned claim: %#v, %v", current, err)
	}
}

func Test_FLOW_092_AbandonAppliedRefusalThenGitRevertToPending(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	abandonGitIdentity(t, root)
	initialSpec := readFile(t, root, "packages/alpha/SPEC.md")
	start(t, root)
	evals, err := done(root)
	if err != nil || evals.State.Phase != "audit" {
		t.Fatalf("fixture evals Done = %#v, %v", evals, err)
	}
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	audited, err := done(root)
	if err != nil || audited.State.Phase != "impl" {
		t.Fatalf("fixture audit Done = %#v, %v", audited, err)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	applied, err := done(root)
	if err != nil || applied.State.Phase != "applied" {
		t.Fatalf("fixture impl Done = %#v, %v", applied, err)
	}
	if _, err := os.Stat(filepath.Join(root, archivedCommitChange)); err != nil {
		t.Fatalf("fixture lacks applied archive: %v", err)
	}
	assertAbandonRefusal(t, root, "manual", "P01")
	git(t, root, "revert", "--no-edit", "HEAD")
	if _, err := os.Stat(filepath.Join(root, pendingCommitChange)); err != nil {
		t.Fatalf("git revert did not restore pending change: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, archivedCommitChange)); !os.IsNotExist(err) {
		t.Fatalf("git revert retained applied archive: %v", err)
	}
	if got := readFile(t, root, "packages/alpha/SPEC.md"); got != initialSpec {
		t.Fatalf("git revert failed to restore original spec: %q", got)
	}
	before := head(t, root)
	oid, err := workflow.Abandon(root, "manual")
	if err != nil {
		t.Fatalf("Abandon after git revert: %v", err)
	}
	assertAbandonCommit(t, root, before, oid)
	if got := readFile(t, root, "packages/alpha/SPEC.md"); got != initialSpec {
		t.Errorf("Abandon after revert changed original spec: %q", got)
	}
}

func Test_FLOW_093_FLOW_095_AbandonRetainsRealAttemptAndFreesAssignedID(t *testing.T) {
	root, started, saved := savedEvalsAttempt(t)
	if len(started.Change.Added) != 1 || started.Change.Added[0].ID != "ALP-003" {
		t.Fatalf("fixture did not assign ALP-003: %#v", started.Change.Added)
	}
	beforeAttempt := git(t, root, "rev-parse", saved.AttemptRef)
	if beforeAttempt != saved.AttemptOID {
		t.Fatalf("Block attempt ref %s points to %s, want %s", saved.AttemptRef, beforeAttempt, saved.AttemptOID)
	}
	parent := head(t, root)
	oid, err := workflow.Abandon(root, "manual")
	if err != nil {
		t.Fatalf("Abandon blocked unclaimed change: %v", err)
	}
	assertAbandonCommit(t, root, parent, oid)
	if got := git(t, root, "rev-parse", saved.AttemptRef); got != beforeAttempt {
		t.Errorf("saved attempt ref changed: %s -> %s", beforeAttempt, got)
	}
	if got := git(t, root, "cat-file", "-t", saved.AttemptOID); got != "commit" {
		t.Errorf("saved attempt object disappeared: %s", got)
	}
	path, err := project.NewChange(root, "alpha", "replacement")
	if err != nil {
		t.Fatalf("NewChange after Abandon: %v", err)
	}
	doc, err := contracts.ParseDocument(readFile(t, root, path))
	if err != nil {
		t.Fatal(err)
	}
	doc.Body = contracts.ReplaceSection(doc.Body, "Motivation", "\nUsers need a replacement alpha behavior.\n")
	doc.Body = contracts.ReplaceSection(doc.Body, "Added", "\n### new:replacement\nReplacement MUST be allocated.\n")
	rendered, err := contracts.RenderDocument(doc)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, path, rendered)
	commit(t, root, "queue replacement change")
	replacement, err := workflow.Start(root, "replacement")
	if err != nil {
		t.Fatalf("Start replacement: %v", err)
	}
	if len(replacement.Change.Added) != 1 || replacement.Change.Added[0].ID != "ALP-003" {
		t.Errorf("replacement assigned IDs = %#v, want reused ALP-003", replacement.Change.Added)
	}
	if got := git(t, root, "rev-parse", saved.AttemptRef); got != beforeAttempt {
		t.Errorf("subsequent Start changed saved attempt ref: %s -> %s", beforeAttempt, got)
	}
}

func Test_FLOW_096_AbandonSelectorsAndMissingAmbiguousRefusal(t *testing.T) {
	for _, selector := range []string{pendingCommitChange, "0001-manual.md", "manual"} {
		t.Run(selector, func(t *testing.T) {
			root := fixture(t, "", "", "", "", "")
			abandonGitIdentity(t, root)
			if commits, err := workflow.Commits(root, selector); err != nil || len(commits) != 0 {
				t.Fatalf("Commits selector %q = %v, %v", selector, commits, err)
			}
			parent := head(t, root)
			oid, err := workflow.Abandon(root, selector)
			if err != nil {
				t.Fatalf("Abandon selector %q: %v", selector, err)
			}
			assertAbandonCommit(t, root, parent, oid)
		})
	}
	root := fixture(t, "", "", "", "", "")
	for _, selector := range []string{"absent", "changes/absent.md"} {
		assertAbandonRefusal(t, root, selector, "U02")
	}
	write(t, root, "changes/0002-manual.md", readFile(t, root, pendingCommitChange))
	commit(t, root, "queue second manual change")
	if _, err := workflow.Commits(root, "manual"); err == nil || !strings.HasPrefix(err.Error(), "U02:") {
		t.Fatalf("fixture slug is not ambiguous to Commits: %v", err)
	}
	assertAbandonRefusal(t, root, "manual", "U02")
}
