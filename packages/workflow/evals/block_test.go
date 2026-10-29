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

func block(t *testing.T, root, reason string) workflow.BlockResult {
	t.Helper()
	r, err := workflow.Block(workflow.BlockRequest{Root: root, Change: "manual", Reason: reason})
	if err != nil {
		t.Fatalf("Block: %v", err)
	}
	return r
}

func blockFixture(t *testing.T, configExtra, changeFront, added, modified, removed string) string {
	t.Helper()
	root := fixture(t, configExtra, changeFront, added, modified, removed)
	git(t, root, "config", "user.name", "Fixture")
	git(t, root, "config", "user.email", "fixture@example.invalid")
	return root
}

func readFile(t *testing.T, root, name string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(root, name))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func assertBlockDocumentOnlyAddsReason(t *testing.T, approved, blocked, reason string) {
	t.Helper()
	before, err := contracts.ParseDocument(approved)
	if err != nil {
		t.Fatalf("parse approved change: %v", err)
	}
	after, err := contracts.ParseDocument(blocked)
	if err != nil {
		t.Fatalf("parse blocked change: %v", err)
	}
	if _, exists := before.Frontmatter["blocked"]; exists {
		t.Fatal("fixture change was already blocked")
	}
	want := make(map[string]any, len(before.Frontmatter)+1)
	for key, value := range before.Frontmatter {
		want[key] = value
	}
	want["blocked"] = reason
	if !reflect.DeepEqual(after.Frontmatter, want) {
		t.Errorf("block changed approved frontmatter: got %#v, want %#v", after.Frontmatter, want)
	}
	if strings.TrimSpace(after.Body) != strings.TrimSpace(before.Body) {
		t.Errorf("block changed approved Markdown body: got %q, want %q", after.Body, before.Body)
	}
}

func assertBlockReleased(t *testing.T, root string, r workflow.BlockResult, reason string) {
	t.Helper()
	if r.State.Phase == "" || r.State.ClaimOID != "" || r.State.Head != head(t, root) || r.State.PhaseBase != head(t, root) {
		t.Errorf("returned blocked state = %#v", r.State)
	}
	if got := r.State.Change.Document.Frontmatter["blocked"]; got != reason {
		t.Errorf("returned blocked reason = %#v, want %q", got, reason)
	}
	if got := git(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/claims/manual"); got != "" {
		t.Errorf("claim still exists: %s", got)
	}
	if _, err := workflow.Current(root, "manual"); err == nil {
		t.Error("Current still found a held claim")
	}
	if got := git(t, root, "status", "--porcelain", "-uall"); got != "" {
		t.Errorf("checkout is not clean: %s", got)
	}
	if index, tree := git(t, root, "write-tree"), git(t, root, "rev-parse", "HEAD^{tree}"); index != tree {
		t.Errorf("index tree %s differs from blocked HEAD tree %s", index, tree)
	}
}

func Test_FLOW_063_FLOW_064_FLOW_065_FLOW_066_FLOW_067_BlockSavesHeldEvalsWorkAndCommitsOnlyState(t *testing.T) {
	root := blockFixture(t, "", "", "### new:extra\nExtra MUST be checked.\n", "", "")
	write(t, root, "packages/alpha/evals/current_test.go", "original eval\n")
	write(t, root, "packages/alpha/evals/retired_test.go", "retired eval\n")
	commit(t, root, "fixture eval files")
	started := start(t, root)
	if started.Phase != "evals" || started.Run == "" || len(started.Change.Added) != 1 || started.Change.Added[0].ID != "ALP-003" {
		t.Fatalf("invalid held evals fixture: %#v", started)
	}
	approved := readFile(t, root, "changes/0001-manual.md")
	if !strings.Contains(approved, "### ALP-003") || !strings.Contains(approved, started.Run) {
		t.Fatalf("assigned IDs/run absent from fixture: %s", approved)
	}
	base := head(t, root)
	prior := "refs/fireweed/attempts/manual/1"
	git(t, root, "update-ref", prior, base)
	write(t, root, "packages/alpha/evals/current_test.go", "staged modified eval\n")
	git(t, root, "add", "packages/alpha/evals/current_test.go")
	if err := os.Remove(filepath.Join(root, "packages/alpha/evals/retired_test.go")); err != nil {
		t.Fatal(err)
	}
	write(t, root, "packages/alpha/evals/staged_test.go", "staged new eval\n")
	git(t, root, "add", "packages/alpha/evals/staged_test.go")
	write(t, root, "packages/alpha/evals/untracked_test.go", "untracked eval\n")
	if got := git(t, root, "status", "--porcelain", "-uall"); !strings.Contains(got, "M  packages/alpha/evals/current_test.go") || !strings.Contains(got, " D packages/alpha/evals/retired_test.go") || !strings.Contains(got, "A  packages/alpha/evals/staged_test.go") || !strings.Contains(got, "?? packages/alpha/evals/untracked_test.go") {
		t.Fatalf("fixture lacks staged, deleted or untracked work: %s", got)
	}
	if got := git(t, root, "show", ":packages/alpha/evals/current_test.go"); got != "staged modified eval" {
		t.Fatalf("fixture index was not staged: %q", got)
	}
	reason := "verification attempts exhausted"
	r := block(t, root, reason)
	if r.State.Phase != "evals" || r.State.Run != started.Run || len(r.State.Change.Added) != 1 || r.State.Change.Added[0].ID != "ALP-003" {
		t.Errorf("returned state lost held evals phase, run or assigned ID: %#v", r.State)
	}
	wantRef := "refs/fireweed/attempts/manual/2"
	if r.AttemptRef != wantRef || r.AttemptOID == "" || git(t, root, "rev-parse", wantRef) != r.AttemptOID || git(t, root, "cat-file", "-t", r.AttemptOID) != "commit" {
		t.Errorf("attempt ref/OID = %q %q, want %q at commit", r.AttemptRef, r.AttemptOID, wantRef)
	}
	if got := git(t, root, "rev-parse", prior); got != base {
		t.Errorf("prior attempt ref changed: %s, want %s", got, base)
	}
	for path, want := range map[string]string{
		"packages/alpha/evals/current_test.go":   "staged modified eval",
		"packages/alpha/evals/staged_test.go":    "staged new eval",
		"packages/alpha/evals/untracked_test.go": "untracked eval",
	} {
		if got := git(t, root, "show", r.AttemptOID+":"+path); got != want {
			t.Errorf("saved %s = %q, want %q", path, got, want)
		}
	}
	if got := git(t, root, "ls-tree", "--name-only", r.AttemptOID, "packages/alpha/evals/retired_test.go"); got != "" {
		t.Errorf("saved attempt retained deleted file: %s", got)
	}
	if got := git(t, root, "rev-parse", "HEAD^"); got != base {
		t.Errorf("block parent = %s, want phase base %s", got, base)
	}
	if got := git(t, root, "diff-tree", "--no-commit-id", "--name-only", "-r", "HEAD"); got != "changes/0001-manual.md" {
		t.Errorf("block commit contains step work: %s", got)
	}
	blocked := readFile(t, root, "changes/0001-manual.md")
	if !strings.Contains(blocked, "### ALP-003") || !strings.Contains(blocked, started.Run) || !strings.Contains(blocked, "phase: evals") || !strings.Contains(blocked, "blocked: "+reason) || !strings.Contains(blocked, "Users need a revised alpha value.") {
		t.Errorf("block commit lost assigned/author state: %s", blocked)
	}
	assertBlockDocumentOnlyAddsReason(t, approved, blocked, reason)
	if got := readFile(t, root, "packages/alpha/evals/current_test.go"); got != "original eval\n" {
		t.Errorf("tracked checkout not restored: %q", got)
	}
	if got := readFile(t, root, "packages/alpha/evals/retired_test.go"); got != "retired eval\n" {
		t.Errorf("deleted checkout not restored: %q", got)
	}
	for _, path := range []string{"packages/alpha/evals/staged_test.go", "packages/alpha/evals/untracked_test.go"} {
		if _, err := os.Stat(filepath.Join(root, path)); !os.IsNotExist(err) {
			t.Errorf("step file %s remains after block: %v", path, err)
		}
	}
	assertBlockReleased(t, root, r, reason)
}

func Test_FLOW_066_BlockImplPreservesCommittedEvalsAndAudit(t *testing.T) {
	root := blockFixture(t, "", "", "", "", "")
	start(t, root)
	write(t, root, "packages/alpha/evals/committed_test.go", "committed eval\n")
	evals, err := done(root)
	if err != nil || evals.State.Phase != "audit" {
		t.Fatalf("evals Done: %#v, %v", evals, err)
	}
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	audited, err := done(root)
	if err != nil || audited.State.Phase != "impl" {
		t.Fatalf("audit Done: %#v, %v", audited, err)
	}
	base := head(t, root)
	committedChange := readFile(t, root, "changes/0001-manual.md")
	if !strings.Contains(committedChange, "No findings; checked ALP-001") || readFile(t, root, "packages/alpha/evals/committed_test.go") != "committed eval\n" {
		t.Fatal("fixture lacks committed eval/audit content")
	}
	write(t, root, "packages/alpha/internal/attempt.go", "unfinished implementation\n")
	r := block(t, root, "implementation attempts exhausted")
	if r.State.Phase != "impl" || git(t, root, "rev-parse", "HEAD^") != base {
		t.Errorf("impl block base/phase = %s, %#v", base, r.State)
	}
	if got := git(t, root, "show", r.AttemptOID+":packages/alpha/internal/attempt.go"); got != "unfinished implementation" {
		t.Errorf("impl attempt work = %q", got)
	}
	if got := git(t, root, "diff-tree", "--no-commit-id", "--name-only", "-r", "HEAD"); got != "changes/0001-manual.md" {
		t.Errorf("impl block contains work: %s", got)
	}
	if got := readFile(t, root, "changes/0001-manual.md"); !strings.Contains(got, "No findings; checked ALP-001") || !strings.Contains(got, "blocked: implementation attempts exhausted") {
		t.Errorf("impl block lost committed audit state: %s", got)
	}
	assertBlockDocumentOnlyAddsReason(t, committedChange, readFile(t, root, "changes/0001-manual.md"), "implementation attempts exhausted")
	if got := readFile(t, root, "packages/alpha/evals/committed_test.go"); got != "committed eval\n" {
		t.Errorf("impl block lost committed eval: %q", got)
	}
	if _, err := os.Stat(filepath.Join(root, "packages/alpha/internal/attempt.go")); !os.IsNotExist(err) {
		t.Errorf("impl work remains: %v", err)
	}
	assertBlockReleased(t, root, r, "implementation attempts exhausted")
}

func Test_FLOW_063_BlockRejectsForeignClaimApprovedEditsAndChangedHead(t *testing.T) {
	t.Run("no claim", func(t *testing.T) {
		root := blockFixture(t, "", "", "", "", "")
		before := head(t, root)
		if _, err := workflow.Block(workflow.BlockRequest{Root: root, Change: "manual", Reason: "exhausted"}); err == nil {
			t.Error("Block accepted change without a claim")
		}
		if head(t, root) != before || git(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/attempts/manual") != "" {
			t.Error("unclaimed Block changed HEAD or attempts")
		}
	})
	t.Run("foreign claim", func(t *testing.T) {
		root := blockFixture(t, "", "", "", "", "")
		other := filepath.Join(t.TempDir(), "linked")
		git(t, root, "worktree", "add", "-b", "linked", other, "HEAD")
		started := start(t, root)
		before := head(t, other)
		if _, err := workflow.Block(workflow.BlockRequest{Root: other, Change: "manual", Reason: "exhausted"}); err == nil {
			t.Error("foreign worktree blocked owner's claim")
		}
		if head(t, other) != before {
			t.Error("foreign Block moved HEAD")
		}
		if current, err := workflow.Current(root, "manual"); err != nil || current.ClaimOID != started.ClaimOID {
			t.Errorf("owner's claim changed: %#v, %v", current, err)
		}
	})
	for _, edit := range []string{"author", "phase", "audit"} {
		t.Run(edit, func(t *testing.T) {
			root := blockFixture(t, "", "", "", "", "")
			started := start(t, root)
			path := "changes/0001-manual.md"
			before := readFile(t, root, path)
			switch edit {
			case "author":
				write(t, root, path, strings.Replace(before, "Users need a revised alpha value.", "Unapproved purpose.", 1))
			case "phase":
				write(t, root, path, strings.Replace(before, "phase: evals", "phase: impl", 1))
			case "audit":
				write(t, root, path, strings.Replace(before, "phase: evals", "audit_rounds: 1\nphase: evals", 1))
			}
			if before == readFile(t, root, path) {
				t.Fatal("fixture did not alter approved state")
			}
			base := head(t, root)
			if _, err := workflow.Block(workflow.BlockRequest{Root: root, Change: "manual", Reason: "exhausted"}); err == nil {
				t.Errorf("Block accepted changed %s", edit)
			}
			if head(t, root) != base || git(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/attempts/manual") != "" {
				t.Errorf("rejected Block committed %s work", edit)
			}
			if current, err := workflow.Current(root, "manual"); err != nil || current.ClaimOID != started.ClaimOID {
				t.Errorf("rejected Block released claim: %#v, %v", current, err)
			}
		})
	}
	t.Run("changed HEAD", func(t *testing.T) {
		root := blockFixture(t, "", "", "", "", "")
		started := start(t, root)
		write(t, root, "packages/alpha/evals/committed_test.go", "premature commit\n")
		git(t, root, "add", "packages/alpha/evals/committed_test.go")
		git(t, root, "commit", "-m", "fixture head advance")
		advanced := head(t, root)
		if advanced == started.PhaseBase {
			t.Fatal("fixture did not advance HEAD")
		}
		if _, err := workflow.Block(workflow.BlockRequest{Root: root, Change: "manual", Reason: "exhausted"}); err == nil {
			t.Error("Block accepted changed HEAD")
		}
		if head(t, root) != advanced || git(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/attempts/manual") != "" {
			t.Error("stale-HEAD Block changed refs or HEAD")
		}
		if current, err := workflow.Current(root, "manual"); err != nil || current.ClaimOID != started.ClaimOID {
			t.Errorf("stale-HEAD Block released claim: %#v, %v", current, err)
		}
	})
}
