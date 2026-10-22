package evals

import (
	a "github.com/lurkshark/fireweed/packages/assets"
	"strings"
	"testing"
)

func Test_AST002PrescriptiveContent(t *testing.T) {
	instructions := map[string][]string{
		"author-change": {
			"Target one module and use unique new:<slug> handles for Added specs.",
			"State each behavior with MUST, MUST NOT or SHOULD; Modified carries the full item and Removed headings only.",
			"Run fireweed stamp for Modified or Removed bases, then fireweed check.",
			"A human's merge of the authored change into the base branch is its approval.",
		},
		"evals": {
			"Run fireweed context for the current change before editing.",
			"Write permanent black-box evals with a current spec ID in every testcase name or classname.",
			"For every Added, Modified or fixes ID, demonstrate a genuinely failing testcase before implementation; compile failures and skips do not count.",
			"Edit only evals, the described interface, module manifests, worker_writable paths and new disposable stub files; never modify existing implementation.",
			"Run fireweed verify to inspect fail-first evidence, then fireweed done; the controller owns commits.",
		},
		"audit": {
			"Use an independent context containing only approved change, specs, interfaces and evals.",
			"Each finding must cite a real spec ID and an exact quote from its clause.",
			"Write No findings and a checked IDs list naming every reviewed ID only when all required assertions are present.",
			"An untestable clause becomes a blocking Questions entry; other valid findings return to evals.",
			"Edit only Audit and Questions, then run fireweed done; never change tests or implementation.",
		},
		"impl": {
			"You MUST NOT edit evals, interfaces, any SPEC.md, protected paths, other modules or other changes.",
			"For append-only paths, add new files; never edit, delete or rename existing files.",
			"Pass the module's evals, all configured dependents and boundaries.",
			"Record behavioral choices missing from specs in Decisions.",
			"Run fireweed verify, then fireweed done to apply specs and archive the change; never commit yourself.",
		},
		"regenerate": {
			"Use fireweed context <change> for a regenerate change; fireweed drill <module> is the author command for a scratch rehearsal.",
			"Rebuild only deleted disposable files using retained specs, Decisions, interfaces, evals and allowed dependency interfaces.",
			"You MUST NOT consult the old implementation, Git history, saved attempts, old binaries or bootstrap source as an oracle.",
			"Preserve interfaces and evals; run own and transitive dependent evals plus boundaries.",
			"Report each contradicted Decision and gap with the nearest spec ID (or no spec) and target module.",
			"Do not author changes automatically; the human reviews the gaps.",
		},
	}
	for name, clauses := range instructions {
		asset, err := a.Playbook(name)
		if err != nil {
			t.Fatal(err)
		}
		for _, clause := range clauses {
			if !strings.Contains(asset.Text, clause) {
				t.Errorf("%s missing concrete instruction: %s", name, clause)
			}
		}
	}
}
func Test_AST003ActionableDiagnostics(t *testing.T) {
	messages := map[string]string{
		"U01":   "Usage is invalid. Run fireweed help and supply the required command arguments.",
		"C01":   "The configuration or version is incompatible. Fix fireweed.yaml or use a compatible binary before migrating.",
		"C02":   "A SPEC.md is outside the modules globs. Move it into an owned module or correct the configured globs.",
		"C03":   "Module metadata is invalid. Correct its name, unique prefix and relative zone paths.",
		"C04":   "Spec grammar or tags are invalid. Use one Specs section, normative items and a nonempty quarantine reason.",
		"C05":   "The change document is invalid. Target one module, supply Motivation and complete Added/Modified items with Removed headings only.",
		"C06":   "The module graph or zones conflict. Remove ownership/eval-disposable overlaps and resolve missing or cyclic dependencies.",
		"C07":   "An enforcement hook is missing. Configure eval and working boundaries, or document an explicit reasoned boundary opt-out.",
		"C08":   "A removed spec ID was reused. Allocate through local new: handles and retain archive reservations.",
		"C09":   "A handle is not declared in this change. Add it locally or combine changes that need the same handle.",
		"C10":   "A modified or removed base is stale or missing. Review projected specs and run fireweed stamp before approval.",
		"C11":   "JUnit evidence is malformed or inconsistent with the hook exit. Repair the reporter; compile failures cannot prove fail-first.",
		"C12":   "A testcase has a missing or unknown spec ID. Name a current ID and remove references to removed specs.",
		"C13":   "An enforced spec has no executed eval. Add a non-skipped black-box test for its behavior.",
		"C14":   "An enforced eval fails. Implement the approved behavior while keeping its evals unchanged.",
		"C15":   "A boundary hook or import rule failed. Respect declared dependencies/interfaces and fix new violations rather than broadening the baseline.",
		"C16":   "The managed AGENTS.md block is stale or malformed. Repair markers and run fireweed agents-md --write.",
		"P01":   "The working tree is dirty. Commit or stash tracked and untracked work before starting a change.",
		"P02":   "The claim or phase baseline conflicts. Inspect the owning worktree before using fireweed release to recover ownership.",
		"P03":   "The change is ineligible behind blocked work. Resolve the blocker or choose an independent eligible module.",
		"P04":   "The authored change is absent from the known base or HEAD is stale. Review approval and merge state; unattended bypass requires explicit --allow-unmerged.",
		"P05":   "The change is blocked. Answer Questions and remove blocked to resume the recorded phase.",
		"P06":   "An agent created commits. Soft-reset to the step baseline, retain edits and verify before letting the controller commit.",
		"P1-01": "Evals edited a forbidden zone. Limit edits to evals, described interfaces, manifests, shared paths and new disposable stubs.",
		"P1-02": "Fail-first proof is missing. Produce an actual failed testcase per changed/fixes ID; compile failures and skipped tests do not count.",
		"P1-03": "Audit evidence is incomplete. Cite real IDs with exact clauses, or write No findings with all checked IDs; block untestable Questions.",
		"P2-01": "Implementation edited a forbidden zone. Restore immutable evals/interfaces/specs/protected files and confine work to this module/change.",
		"D01":   "The module has no current passing drill. Run fireweed drill and review own/dependent results before regeneration.",
		"D02":   "Scratch setup failed. Repair the setup hook/templates and inspect worktrees under the common Git directory.",
		"D03":   "Audit reuses the worker command. Configure an independent audit command to reduce shared review blind spots.",
		"D04":   "Boundary enforcement is absent or opted out. Review the module's reason and keep the exception visible in status and PR bodies.",
		"D05":   "A playbook override is unsafe, unreadable or has stale based_on metadata. Restore a readable contained override; for stale metadata, compare the embedded hash, review the instructions and update based_on.",
		"D06":   "A doctor canary or recovery check failed. Repair protected hooks and all four canaries; an always-success hook is invalid.",
		"D07":   "Regeneration failed. Review the drill report/ref and author focused changes for accepted gaps.",
		"H01":   "A worker edited protected files. Restore them and move author configuration/glue changes to a separate author commit.",
		"H02":   "An implementation commit changed immutable evals or interfaces. Restore them and put approved contract edits in the evals phase.",
		"H03":   "An append-only file was edited, deleted or renamed by anyone, including merge resolutions. Add a new file instead.",
		"H04":   "An archived change is immutable; only deletion for git revert is allowed. Abort the merge, run fireweed commits <change>, git revert the incoming duplicate's commits, then merge again. Do not keep both implementations by picking one archive.",
	}
	for code, want := range messages {
		got, err := a.Explain(code)
		if err != nil || got != want {
			t.Errorf("%s problem/remedy: %q %v, want %q", code, got, err, want)
		}
	}
	if len(messages) != len(a.Codes()) {
		t.Errorf("not every diagnostic has an asserted problem and remedy")
	}
}

func Test_AST003OverrideProblemsAndRemedies(t *testing.T) {
	text, err := a.Explain("D05")
	if err != nil {
		t.Fatal(err)
	}
	for _, clause := range []struct {
		name  string
		terms []string
	}{
		{"unsafe-override", []string{"playbook override", "unsafe"}},
		{"unreadable-override", []string{"playbook override", "unreadable"}},
		{"stale-metadata", []string{"stale", "based_on"}},
		{"contained-readable-remedy", []string{"restore", "readable", "contained", "override"}},
		{"stale-remedy", []string{"compare", "embedded hash", "review", "instructions", "update based_on"}},
	} {
		t.Run(clause.name, func(t *testing.T) {
			for _, term := range clause.terms {
				if !strings.Contains(strings.ToLower(text), term) {
					t.Errorf("D05 %s guidance missing %q: %q", clause.name, term, text)
				}
			}
		})
	}
}
