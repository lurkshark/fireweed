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

const workerBlockReason = "The approved rule is ambiguous."
const workerBlockQuestions = "- Which value takes precedence?\n- What should happen when neither value is set?\n"
const workerBlockNotes = "Unfinished worker notes."

func holdWorkerPhase(t *testing.T, phase string) (string, workflow.State) {
	t.Helper()
	root := blockFixture(t, "", "", "### new:review [review]\nThe extra behavior MUST be reviewed.\n", "", "")
	held := start(t, root)
	if phase == "impl" {
		if result, err := done(root); err != nil || result.State.Phase != "audit" {
			t.Fatalf("fixture evals Done = phase %q, error %v", result.State.Phase, err)
		}
		audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
		result, err := done(root)
		if err != nil {
			t.Fatalf("fixture audit Done: %v", err)
		}
		held = result.State
	}
	if held.Phase != phase || held.Run == "" || held.ClaimOID == "" {
		t.Fatalf("fixture state = phase %q, run %q, claim %q; want held %s", held.Phase, held.Run, held.ClaimOID, phase)
	}
	return root, held
}

func declareWorkerBlock(t *testing.T, root string) {
	t.Helper()
	doc, err := contracts.ParseDocument(readFile(t, root, pendingCommitChange))
	if err != nil {
		t.Fatal(err)
	}
	doc.Frontmatter["blocked"] = workerBlockReason
	doc.Body = contracts.ReplaceSection(doc.Body, "Notes", "\n"+workerBlockNotes+"\n\n")
	doc.Body = contracts.ReplaceSection(doc.Body, "Questions", "\n"+workerBlockQuestions)
	rendered, err := contracts.RenderDocument(doc)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, pendingCommitChange, rendered)
}

func Test_FLOW_137_DoneSavesWorkerDeclaredBlock(t *testing.T) {
	for _, phase := range []string{"evals", "impl"} {
		t.Run(phase, func(t *testing.T) {
			root, held := holdWorkerPhase(t, phase)
			base := head(t, root)
			path, edit := "packages/alpha/evals/unfinished.txt", "unfinished eval work\n"
			if phase == "impl" {
				path, edit = "packages/alpha/internal/ready", "unfinished implementation work\n"
			}
			write(t, root, path, edit)
			if _, err := verify(root, false); err != nil {
				t.Fatalf("fixture should pass full %s Verify before the declaration: %v", phase, err)
			}
			declareWorkerBlock(t, root)

			result, err := done(root)
			assertErrorCode(t, "worker-declared Done", err, "P06")
			if result.State.Phase != phase || result.State.Run != held.Run {
				t.Errorf("Done returned phase %q and run %q, want %q and %q",
					result.State.Phase, result.State.Run, phase, held.Run)
			}
			_, err = workflow.Current(root, "manual")
			assertErrorCode(t, "Current after worker-declared Done", err, "P02")
			if claim := git(t, root, "for-each-ref", "--format=%(refname)", "refs/fireweed/claims/manual"); claim != "" {
				t.Errorf("blocked change still holds %s", claim)
			}
			if _, err := os.Stat(filepath.Join(root, archivedCommitChange)); !os.IsNotExist(err) {
				t.Errorf("archive should be absent after blocking: %v", err)
			}
			if got := git(t, root, "rev-parse", "HEAD^"); got != base {
				t.Errorf("block commit parent = %s, want %s", got, base)
			}
			if got := git(t, root, "diff-tree", "--no-commit-id", "--name-only", "-r", "HEAD"); got != pendingCommitChange {
				t.Fatalf("block commit changed %q, want only %s", got, pendingCommitChange)
			}
			if _, err := os.Stat(filepath.Join(root, path)); !os.IsNotExist(err) {
				t.Errorf("saved step edit remains in the checkout: %v", err)
			}

			committed, err := contracts.ParseDocument(git(t, root, "show", "HEAD:"+pendingCommitChange))
			if err != nil {
				t.Fatal(err)
			}
			approved := held.Change.Document
			wantFrontmatter := make(map[string]any, len(approved.Frontmatter)+1)
			for key, value := range approved.Frontmatter {
				wantFrontmatter[key] = value
			}
			wantFrontmatter["blocked"] = workerBlockReason
			if !reflect.DeepEqual(committed.Frontmatter, wantFrontmatter) {
				t.Errorf("committed frontmatter = %#v, want %#v", committed.Frontmatter, wantFrontmatter)
			}
			questions, _ := contracts.Section(committed.Body, "Questions")
			if strings.TrimSpace(questions) != strings.TrimSpace(workerBlockQuestions) {
				t.Errorf("committed Questions = %q, want %q", questions, workerBlockQuestions)
			}
			for _, name := range []string{"Added", "Modified", "Audit", "Provenance", "Notes"} {
				want, _ := contracts.Section(approved.Body, name)
				got, _ := contracts.Section(committed.Body, name)
				if strings.TrimSpace(got) != strings.TrimSpace(want) {
					t.Errorf("committed %s = %q, want approved %q", name, got, want)
				}
			}

			attempt := git(t, root, "for-each-ref", "--format=%(refname)", "refs/fireweed/attempts/manual")
			if attempt == "" || strings.Contains(attempt, "\n") {
				t.Fatalf("saved attempts = %q, want one", attempt)
			}
			if got := git(t, root, "show", attempt+":"+path); got != strings.TrimSpace(edit) {
				t.Errorf("saved %s = %q, want %q", path, got, edit)
			}
			if got := git(t, root, "show", attempt+":"+pendingCommitChange); !strings.Contains(got, workerBlockNotes) {
				t.Errorf("saved change lacks worker notes: %s", got)
			}
		})
	}
}

func Test_FLOW_137_DoneValidatesApprovalBeforeRecordingWorkerBlock(t *testing.T) {
	root, held := holdWorkerPhase(t, "evals")
	declareWorkerBlock(t, root)
	declared := readFile(t, root, pendingCommitChange)
	edited := strings.Replace(declared, "Users need a revised alpha value.", "Unapproved motivation.", 1)
	if edited == declared {
		t.Fatal("fixture could not edit approved motivation")
	}
	write(t, root, pendingCommitChange, edited)

	_, err := done(root)
	assertErrorCode(t, "worker-declared Done with unapproved edit", err, "P03")
	if head(t, root) != held.PhaseBase || git(t, root, "for-each-ref", "--format=%(refname)", "refs/fireweed/attempts/manual") != "" {
		t.Error("rejected declaration committed a block or saved an attempt")
	}
	if current, err := workflow.Current(root, "manual"); err != nil || current.ClaimOID != held.ClaimOID {
		t.Errorf("rejected declaration changed the held claim: %q, %v", current.ClaimOID, err)
	}
}

func Test_FLOW_138_ContextExplainsHowToDeclareWorkerBlock(t *testing.T) {
	for _, phase := range []string{"evals", "impl"} {
		t.Run(phase, func(t *testing.T) {
			root, _ := holdWorkerPhase(t, phase)
			context, err := workflow.Context(root, "manual", false)
			if err != nil {
				t.Fatal(err)
			}
			context = strings.ToLower(context)
			for _, term := range []string{"blocked", "reason", "## questions"} {
				if !strings.Contains(context, term) {
					t.Errorf("%s Context lacks worker-block guidance for %q", phase, term)
				}
			}
		})
	}
}
