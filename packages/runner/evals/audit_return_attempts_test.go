package evals_test

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/runner"
)

const auditReturnClause = "Alpha MUST return the new value on subsequent calls."

// Preserve all other sections when the fixture auditor submits new findings.
const replaceFixtureAudit = `SUBMISSION="$submission" awk '
  BEGIN { submission=ENVIRON["SUBMISSION"] }
  /^## Audit$/ { print; print submission; print ""; found=1; replacing=1; next }
  /^## / { replacing=0 }
  !replacing { print }
  END { if (!found) { print "## Audit"; print submission } }
' "$held" > "$1/audit-submission"
cp "$1/audit-submission" "$held"
`

func auditReturnChange(t *testing.T, root, path string) {
	t.Helper()
	change := nextRead(t, root, path)
	const firstClause = " MUST return the new value.\n"
	if !strings.Contains(change, firstClause) {
		t.Fatal("fixture value clause missing")
	}
	nextWrite(t, root, path, strings.Replace(change, firstClause, firstClause+auditReturnClause+"\n", 1))
	nextCommit(t, root, "add fixture clause for audit")
}

func auditReturnFixture(t *testing.T, workerMode string, correct bool) nextFixture {
	t.Helper()
	f := newNextFixture(t, true, true, workerMode)
	auditReturnChange(t, f.root, nextChange)
	if correct {
		worker := nextRead(t, f.logDir, "worker.sh")
		const output = "printf 'worker-%s' \"$phase\"\n"
		if !strings.Contains(worker, output) {
			t.Fatal("fixture worker output missing")
		}
		correction := `if [ "$phase" = evals ] && grep -q '^## Audit$' changes/0001-next.md; then
  cat >> packages/alpha/evals/next_test.go <<'GO'
func Test_ALP_002_SubsequentCalls(t *testing.T) {
    _ = alpha.Value()
    if got := alpha.Value(); got != "new" { t.Fatalf("subsequent call = %q", got) }
}
GO
fi
`
		nextWrite(t, f.logDir, "worker.sh", strings.Replace(worker, output, correction+output, 1))
	}
	// The auditor changes only its copy of the Audit section. Its observations
	// record which approved phase and run preceded each submission.
	auditBeforeEdit(t, f, `set -e
held=changes/0001-next.md
rounds=$(sed -n 's/^audit_rounds: //p' "$held")
round=$((${rounds:-0} + 1))
test "$round" -le 3
cp "$held" "$1/audit-input-$round"
if grep -q 'func Test_ALP_002_SubsequentCalls' packages/alpha/evals/next_test.go; then
  submission=$'No findings\nChecked IDs: ALP-002'
else
  submission='- ALP-002: "`+auditReturnClause+`" — missing an assertion for subsequent calls in audit '"$round"
fi
`+replaceFixtureAudit+`printf auditor
exit 0
`)
	return f
}

func assertAuditReturnSteps(t *testing.T, f nextFixture, result runner.NextResult, phases []string) {
	t.Helper()
	var captured, invoked []string
	for _, step := range result.Steps {
		captured = append(captured, step.Phase)
		role := "worker"
		if step.Phase == "audit" {
			role = "audit"
		}
		if step.Role != role || step.Agent.ExitCode != 0 && !step.Agent.UsageLimit {
			t.Errorf("unexpected captured invocation for %s: %#v", step.Phase, step)
		}
	}
	for _, line := range nextInvocations(t, f) {
		phase, _, _ := strings.Cut(line, "|")
		invoked = append(invoked, phase)
	}
	if !slices.Equal(captured, phases) || !slices.Equal(invoked, phases) {
		t.Errorf("phases captured=%v invoked=%v, want %v", captured, invoked, phases)
	}
	auditWorkspaceGone(t, f)
}

func Test_RUN_021_AuditCorrectionWithinAllowance(t *testing.T) {
	for _, tc := range []struct {
		name           string
		usage          bool
		phases         []string
		correctionCall int
	}{
		{"correction", false, []string{"evals", "audit", "evals", "audit", "impl", "impl"}, 2},
		{"usage retry", true, []string{"evals", "evals", "audit", "evals", "audit", "impl", "impl"}, 3},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := auditReturnFixture(t, "retry-impl", true)
			if tc.usage {
				usagePolicy(t, f.root, "limit_wait: 5ms\nlimit_budget: 5ms\n")
				nextUsageAgent(t, f, "first")
			}
			r, err := f.run(false)
			if err != nil || r.State.Phase != "applied" {
				if n := len(r.Steps); n > 0 && r.Steps[n-1].Agent.Stderr != "" {
					t.Logf("last agent stderr: %s", r.Steps[n-1].Agent.Stderr)
				}
				t.Fatalf("audit correction ended in %q: %v", r.State.Phase, err)
			}
			assertAuditReturnSteps(t, f, r, tc.phases)
			if tc.usage && (len(r.Steps) < 2 || !r.Steps[0].Agent.UsageLimit || r.Steps[0].Agent.LogPath != r.Steps[1].Agent.LogPath) {
				t.Error("usage retry did not reuse its evals attempt")
			}
			prompt := nextRead(t, f.logDir, fmt.Sprintf("worker-evals-%d.prompt", tc.correctionCall))
			if !strings.Contains(prompt, "missing an assertion for subsequent calls in audit 1") {
				t.Error("correction prompt omitted the new audit finding")
			}
			archive := nextRead(t, f.root, "changes/archive/0001-next.md")
			if !strings.Contains(archive, "audit_rounds: 1") || !strings.Contains(archive, "No findings; checked ALP-002") {
				t.Errorf("archive did not record the resolved audit return: %q", archive)
			}
		})
	}
}

func Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals(t *testing.T) {
	for _, tc := range []struct {
		name, workerMode string
		auditRound       int
		phases           []string
	}{
		{"later return", "normal", 2, []string{"evals", "audit", "evals", "audit"}},
		{"earlier verification retry", "retry-both", 1, []string{"evals", "evals", "audit"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := auditReturnFixture(t, tc.workerMode, false)
			base := nextGit(t, f.root, "rev-parse", "HEAD")
			r, err := f.run(false)
			assertAuditReturnSteps(t, f, r, tc.phases)
			assertEvalsAllowanceError(t, err)
			if r.State.Phase != "evals" || r.State.ClaimOID != "" {
				t.Fatalf("audit return did not block the held evals step: phase=%q claim=%q", r.State.Phase, r.State.ClaimOID)
			}
			prior, parseErr := contracts.ParseDocument(nextRead(t, f.logDir, fmt.Sprintf("audit-input-%d", tc.auditRound)))
			if parseErr != nil {
				t.Fatal(parseErr)
			}
			// The last phase commit accepts the finding and returns to evals.
			phaseHead := nextGit(t, f.root, "rev-parse", "HEAD^")
			if count := nextGit(t, f.root, "rev-list", "--count", base+".."+phaseHead); count != fmt.Sprint(2*tc.auditRound) {
				t.Errorf("phase commits before block = %s, want %d", count, 2*tc.auditRound)
			}
			ref := assertSavedBlock(t, f.root, "next", nextChange, phaseHead, err)
			if r.Empty || r.State.Run == "" || r.State.Run != prior.Frontmatter["run"] ||
				len(r.State.Change.Added) != 1 || r.State.Change.Added[0].ID != "ALP-002" {
				t.Errorf("blocked result lost identity: phase=%q run=%q added=%#v", r.State.Phase, r.State.Run, r.State.Change.Added)
			}
			finding := fmt.Sprintf("- ALP-002: %q — missing an assertion for subsequent calls in audit %d", auditReturnClause, tc.auditRound)
			assertFinding := func(label, source string) {
				t.Helper()
				doc, err := contracts.ParseDocument(source)
				if err != nil {
					t.Fatalf("%s: %v", label, err)
				}
				audit, _ := contracts.Section(doc.Body, "Audit")
				if doc.Frontmatter["run"] != prior.Frontmatter["run"] || doc.Frontmatter["phase"] != "evals" ||
					doc.Frontmatter["audit_rounds"] != tc.auditRound || strings.TrimSpace(audit) != finding ||
					!strings.Contains(doc.Body, "### ALP-002\n") {
					t.Errorf("%s did not retain the held evals identity and latest audit finding: %q", label, source)
				}
			}
			assertFinding("phase base", nextGit(t, f.root, "show", phaseHead+":"+nextChange))
			blocked := nextRead(t, f.root, nextChange)
			assertFinding("blocked change", blocked)
			if !strings.Contains(blocked, "blocked:") {
				t.Error("blocked change has no blocking reason")
			}
			if audit, _ := contracts.Section(r.State.Change.Document.Body, "Audit"); strings.TrimSpace(audit) != finding {
				t.Errorf("returned state lost the latest audit finding: %q", audit)
			}
			if ref != "" {
				assertFinding("saved attempt", nextGit(t, f.root, "show", ref+":"+nextChange))
				if parent := nextGit(t, f.root, "rev-parse", ref+"^"); parent != phaseHead {
					t.Errorf("saved attempt parent = %s, want evals phase base %s", parent, phaseHead)
				}
			}
		})
	}
}

func assertEvalsAllowanceError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("evals attempt exhaustion returned no error")
		return
	}
	message := strings.ToLower(err.Error())
	if !strings.Contains(message, "evals") || !strings.Contains(message, "attempt") && !strings.Contains(message, "allowance") ||
		!strings.Contains(message, "exhaust") && !strings.Contains(message, "limit") {
		t.Errorf("error does not identify exhausted evals allowance: %v", err)
	}
}

func Test_RUN_021_AuditRoundLimitCanStopBeforeEvalsAllowance(t *testing.T) {
	f := auditReturnFixture(t, "normal", false)
	config := nextRead(t, f.root, "fireweed.yaml")
	nextWrite(t, f.root, "fireweed.yaml", strings.Replace(config, "audit:\n", "audit:\n  max_rounds: 1\n", 1))
	nextAttempts(t, f, 3)
	r, err := f.run(false)
	assertAuditReturnSteps(t, f, r, []string{"evals", "audit", "evals", "audit"})
	if code, _, ok := contracts.ErrorCode(err); !ok || code != "P06" {
		t.Errorf("audit round limit returned %v, want P06", err)
	}
	questions, _ := contracts.Section(r.State.Change.Document.Body, "Questions")
	if r.State.Phase != "audit" || r.State.ClaimOID != "" || r.State.Change.Document.Frontmatter["audit_rounds"] != 1 ||
		!strings.Contains(questions, auditReturnClause) {
		t.Errorf("audit did not block at the persisted round limit: phase=%q claim=%q questions=%q", r.State.Phase, r.State.ClaimOID, questions)
	}
}

func Test_RUN_070_RunContinuesIndependentWorkAfterAuditReturnBlock(t *testing.T) {
	f := newRunFixture(t, 3, false, false)
	beta := nextRead(t, f.root, "packages/beta/SPEC.md")
	nextWrite(t, f.root, "packages/beta/SPEC.md", strings.Replace(beta, "depends_on: []", "depends_on: [alpha]", 1))
	auditReturnChange(t, f.root, "changes/0001-alpha.md")
	timeFastHook(t, f)
	nextAttempts(t, nextFixture{root: f.root}, 1)

	auditor := nextRead(t, f.agents, "auditor.sh")
	const noFindings = "printf '\\n## Audit\\nNo findings"
	if !strings.Contains(auditor, noFindings) {
		t.Fatal("fixture auditor output missing")
	}
	finding := `if [ "$module" = alpha ]; then
  set -e
  submission='- ALP-002: "` + auditReturnClause + `" — missing an assertion for subsequent calls'
` + replaceFixtureAudit + `
  printf alpha-audit
  exit 0
fi
`
	nextWrite(t, f.agents, "auditor.sh", strings.Replace(auditor, noFindings, finding+noFindings, 1))
	r, err := runner.Run(context.Background(), f.request(0, false))
	assertEvalsAllowanceError(t, err)
	assertSavedAttempt(t, f.root, "alpha", err)
	var calls []string
	for _, line := range runInvocations(t, f) {
		parts := strings.SplitN(line, "|", 3)
		calls = append(calls, strings.Join(parts[:2], "/"))
	}
	want := []string{"alpha/evals", "alpha/audit", "gamma/evals", "gamma/audit", "gamma/impl"}
	if !slices.Equal(calls, want) {
		t.Errorf("Run invocations = %v, want %v", calls, want)
	}
	assertRunApplied(t, f, r.Applied, "gamma")
	assertRunBetaWaiting(t, f)
}
