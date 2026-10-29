package evals_test

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/workflow"
)

// These fixtures use the public lifecycle to create a saved commit and release
// its claim. An author then commits the unblock before any resume request.
func unblockAttempt(t *testing.T, root string) {
	t.Helper()
	path := "changes/0001-manual.md"
	doc, err := contracts.ParseDocument(readFile(t, root, path))
	if err != nil {
		t.Fatal(err)
	}
	if _, blocked := doc.Frontmatter["blocked"]; !blocked {
		t.Fatal("fixture has no committed block to remove")
	}
	delete(doc.Frontmatter, "blocked")
	doc.Body = contracts.ReplaceSection(doc.Body, "Motivation", "\nAuthor approved a revised purpose.\n")
	doc.Body = contracts.ReplaceSection(doc.Body, "Interface", "\nAuthor approved the current interface.\n")
	rendered, err := contracts.RenderDocument(doc)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, path, rendered)
	commit(t, root, "author unblocks saved attempt")
}

func savedEvalsAttempt(t *testing.T) (string, workflow.State, workflow.BlockResult) {
	t.Helper()
	root := blockFixture(t, "", "", "### new:extra\nExtra MUST be checked.\n", "", "")
	write(t, root, "packages/alpha/evals/current_test.go", "original eval\n")
	write(t, root, "packages/alpha/evals/retired_test.go", "retired eval\n")
	commit(t, root, "fixture eval files")
	started := start(t, root)
	if started.Phase != "evals" || started.Run == "" || len(started.Change.Added) != 1 || started.Change.Added[0].ID != "ALP-003" {
		t.Fatalf("invalid assigned evals fixture: %#v", started)
	}
	write(t, root, "packages/alpha/evals/current_test.go", "saved modified eval\n")
	if err := os.Remove(filepath.Join(root, "packages/alpha/evals/retired_test.go")); err != nil {
		t.Fatal(err)
	}
	write(t, root, "packages/alpha/evals/new_test.go", "saved new eval\n")
	path := "changes/0001-manual.md"
	doc, err := contracts.ParseDocument(readFile(t, root, path))
	if err != nil {
		t.Fatal(err)
	}
	doc.Body = contracts.ReplaceSection(doc.Body, "Decisions", "\nSaved worker decision.\n")
	doc.Body = contracts.ReplaceSection(doc.Body, "Questions", "\nSaved worker question.\n")
	rendered, err := contracts.RenderDocument(doc)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, path, rendered)
	saved := block(t, root, "eval attempts exhausted")
	if saved.AttemptRef != "refs/fireweed/attempts/manual/1" || saved.AttemptOID == "" {
		t.Fatalf("Block did not save attempt 1: %#v", saved)
	}
	if got := git(t, root, "status", "--porcelain", "-uall"); got != "" {
		t.Fatalf("Block left checkout dirty: %s", got)
	}
	return root, started, saved
}

func assertNoResumeMutation(t *testing.T, root, beforeHead, beforeChange string) {
	t.Helper()
	if head(t, root) != beforeHead || readFile(t, root, "changes/0001-manual.md") != beforeChange {
		t.Error("rejected resume changed HEAD or change document")
	}
	if got := git(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/claims/manual"); got != "" {
		t.Errorf("rejected resume claimed change: %s", got)
	}
	if got := git(t, root, "status", "--porcelain", "-uall"); got != "" {
		t.Errorf("rejected resume dirtied checkout: %s", got)
	}
}

func Test_FLOW_073_FLOW_074_FLOW_075_StartFromAttemptRestoresEvalsStepAndRetainsCurrentApproval(t *testing.T) {
	root, started, saved := savedEvalsAttempt(t)
	write(t, root, "later-independent.txt", "later independent commit\n")
	commit(t, root, "later independent work")
	unblockAttempt(t, root)
	beforeHead := head(t, root)
	before := readFile(t, root, "changes/0001-manual.md")
	if strings.Contains(before, "Saved worker decision.") || strings.Contains(before, "Saved worker question.") {
		t.Fatal("blocked fixture leaked saved worker body")
	}
	got, err := workflow.StartFromAttempt(root, "manual", 1)
	if err != nil {
		t.Fatalf("StartFromAttempt(%s): %v", saved.AttemptRef, err)
	}
	if got.Run != started.Run || got.Phase != started.Phase || len(got.Change.Added) != 1 || got.Change.Added[0].ID != started.Change.Added[0].ID {
		t.Errorf("resume changed run, phase or assigned ID: %#v", got)
	}
	if got.Head != beforeHead || got.PhaseBase != beforeHead || head(t, root) != beforeHead {
		t.Errorf("resume changed HEAD or phase base: state=%#v, HEAD=%s", got, head(t, root))
	}
	current, currentErr := workflow.Current(root, "manual")
	if got.ClaimOID == "" || currentErr != nil || current.ClaimOID != got.ClaimOID || got.Snapshot.Root != root {
		t.Errorf("resume did not return an owned held claim: state=%#v, Current=%#v, %v", got, current, currentErr)
	}
	if text := readFile(t, root, "packages/alpha/evals/current_test.go"); text != "saved modified eval\n" {
		t.Errorf("modified eval = %q", text)
	}
	if _, err := os.Stat(filepath.Join(root, "packages/alpha/evals/retired_test.go")); !os.IsNotExist(err) {
		t.Errorf("saved deletion was not restored: %v", err)
	}
	if text := readFile(t, root, "packages/alpha/evals/new_test.go"); text != "saved new eval\n" {
		t.Errorf("new eval = %q", text)
	}
	if text := readFile(t, root, "later-independent.txt"); text != "later independent commit\n" {
		t.Errorf("later independent work = %q", text)
	}
	resumed, err := contracts.ParseDocument(readFile(t, root, "changes/0001-manual.md"))
	if err != nil {
		t.Fatal(err)
	}
	approved, err := contracts.ParseDocument(before)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"Motivation", "Interface", "Added", "Modified", "Removed", "Removed decisions"} {
		want, _ := contracts.Section(approved.Body, name)
		actual, _ := contracts.Section(resumed.Body, name)
		if actual != want {
			t.Errorf("current approved %s changed: got %q, want %q", name, actual, want)
		}
	}
	if !reflect.DeepEqual(resumed.Frontmatter, approved.Frontmatter) {
		t.Errorf("current runtime/author frontmatter changed: got %#v, want %#v", resumed.Frontmatter, approved.Frontmatter)
	}
	for name, want := range map[string]string{"Decisions": "Saved worker decision.", "Questions": "Saved worker question."} {
		body, ok := contracts.Section(resumed.Body, name)
		if !ok || !strings.Contains(body, want) {
			t.Errorf("saved worker %s absent: %q", name, body)
		}
	}
}

func Test_FLOW_073_FLOW_074_FLOW_075_StartFromAttemptImplRetainsCommittedEvalsAuditAndProvenance(t *testing.T) {
	root := blockFixture(t, "", "", "", "", "")
	start(t, root)
	write(t, root, "packages/alpha/evals/committed_test.go", "committed eval\n")
	evals, err := done(root)
	if err != nil || evals.State.Phase != "audit" {
		t.Fatalf("evals Done: %#v, %v", evals, err)
	}
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	impl, err := done(root)
	if err != nil || impl.State.Phase != "impl" {
		t.Fatalf("audit Done: %#v, %v", impl, err)
	}
	write(t, root, "packages/alpha/internal/attempt.go", "saved implementation\n")
	path := "changes/0001-manual.md"
	doc, err := contracts.ParseDocument(readFile(t, root, path))
	if err != nil {
		t.Fatal(err)
	}
	doc.Body = contracts.ReplaceSection(doc.Body, "Decisions", "\nSaved impl decision.\n")
	rendered, err := contracts.RenderDocument(doc)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, path, rendered)
	saved := block(t, root, "impl attempts exhausted")
	if saved.AttemptRef != "refs/fireweed/attempts/manual/1" {
		t.Fatalf("wrong impl attempt ref: %s", saved.AttemptRef)
	}
	unblockAttempt(t, root)
	beforeHead := head(t, root)
	before, err := contracts.ParseDocument(readFile(t, root, path))
	if err != nil {
		t.Fatal(err)
	}
	provenance, hasProvenance := contracts.Section(before.Body, "Provenance")
	auditBody, hasAudit := contracts.Section(before.Body, "Audit")
	if !hasProvenance || strings.TrimSpace(provenance) == "" || !hasAudit || !strings.Contains(auditBody, "No findings; checked ALP-001") {
		t.Fatalf("fixture lacks committed controller provenance or audit: %#v", before)
	}
	got, err := workflow.StartFromAttempt(root, "manual", 1)
	if err != nil {
		t.Fatalf("impl StartFromAttempt: %v", err)
	}
	if got.Run != impl.State.Run || got.Phase != "impl" || got.Head != beforeHead || head(t, root) != beforeHead || got.ClaimOID == "" {
		t.Errorf("impl resume lost run/phase/claim or changed HEAD: %#v", got)
	}
	if text := readFile(t, root, "packages/alpha/evals/committed_test.go"); text != "committed eval\n" {
		t.Errorf("committed eval changed: %q", text)
	}
	if text := readFile(t, root, "packages/alpha/internal/attempt.go"); text != "saved implementation\n" {
		t.Errorf("impl step work = %q", text)
	}
	resumed, err := contracts.ParseDocument(readFile(t, root, path))
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"Audit", "Provenance", "Motivation", "Interface"} {
		want, _ := contracts.Section(before.Body, name)
		actual, _ := contracts.Section(resumed.Body, name)
		if actual != want {
			t.Errorf("committed/current %s changed: got %q, want %q", name, actual, want)
		}
	}
	if !reflect.DeepEqual(resumed.Frontmatter, before.Frontmatter) {
		t.Errorf("impl resume changed runtime/author frontmatter: %#v, want %#v", resumed.Frontmatter, before.Frontmatter)
	}
	if body, _ := contracts.Section(resumed.Body, "Decisions"); !strings.Contains(body, "Saved impl decision.") {
		t.Errorf("saved impl decision missing: %q", body)
	}
}

func Test_FLOW_073_StartFromAttemptRequiresCleanEligibleUnblockedAndUnclaimedChange(t *testing.T) {
	t.Run("dirty", func(t *testing.T) {
		root, _, _ := savedEvalsAttempt(t)
		unblockAttempt(t, root)
		write(t, root, "uncommitted.txt", "dirty\n")
		if _, err := workflow.StartFromAttempt(root, "manual", 1); err == nil || !strings.HasPrefix(err.Error(), "P01") {
			t.Errorf("dirty resume = %v, want P01", err)
		}
		if got := git(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/claims/manual"); got != "" {
			t.Errorf("dirty resume created claim: %s", got)
		}
	})
	t.Run("still blocked", func(t *testing.T) {
		root, _, _ := savedEvalsAttempt(t)
		beforeHead, before := head(t, root), readFile(t, root, "changes/0001-manual.md")
		if _, err := workflow.StartFromAttempt(root, "manual", 1); err == nil || !strings.HasPrefix(err.Error(), "P01") {
			t.Errorf("blocked resume = %v, want P01", err)
		}
		assertNoResumeMutation(t, root, beforeHead, before)
	})
	t.Run("not first eligible", func(t *testing.T) {
		root, _, _ := savedEvalsAttempt(t)
		unblockAttempt(t, root)
		write(t, root, "changes/0000-first.md", "---\nmodule: alpha\n---\n## Motivation\nEarlier approved work.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n")
		commit(t, root, "earlier eligible change")
		beforeHead, before := head(t, root), readFile(t, root, "changes/0001-manual.md")
		if _, err := workflow.StartFromAttempt(root, "manual", 1); err == nil || !strings.HasPrefix(err.Error(), "P01") {
			t.Errorf("waiting resume = %v, want P01", err)
		}
		assertNoResumeMutation(t, root, beforeHead, before)
	})
	t.Run("foreign claim", func(t *testing.T) {
		root, _, _ := savedEvalsAttempt(t)
		unblockAttempt(t, root)
		other := filepath.Join(t.TempDir(), "linked")
		git(t, root, "worktree", "add", "-b", "linked", other, "HEAD")
		owner := start(t, root)
		beforeHead, before := head(t, other), readFile(t, other, "changes/0001-manual.md")
		if _, err := workflow.StartFromAttempt(other, "manual", 1); err == nil || !strings.HasPrefix(err.Error(), "P02") {
			t.Errorf("foreign claim resume = %v, want P02", err)
		}
		if head(t, other) != beforeHead || readFile(t, other, "changes/0001-manual.md") != before {
			t.Error("foreign resume changed other worktree")
		}
		if current, err := workflow.Current(root, "manual"); err != nil || current.ClaimOID != owner.ClaimOID {
			t.Errorf("foreign resume changed owner's claim: %#v, %v", current, err)
		}
	})
}

func Test_FLOW_076_StartFromAttemptRejectsInvalidOrMissingAttemptBeforeMutation(t *testing.T) {
	root, _, _ := savedEvalsAttempt(t)
	unblockAttempt(t, root)
	beforeHead, before := head(t, root), readFile(t, root, "changes/0001-manual.md")
	for _, attempt := range []int{0, -1, 2} {
		if _, err := workflow.StartFromAttempt(root, "manual", attempt); err == nil || !strings.HasPrefix(err.Error(), "P01") {
			t.Errorf("attempt %d = %v, want P01", attempt, err)
		}
		assertNoResumeMutation(t, root, beforeHead, before)
	}
}
