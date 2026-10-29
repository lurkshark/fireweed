package evals_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/runner"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func usagePolicy(t *testing.T, root, policy string) {
	t.Helper()
	nextWrite(t, root, "fireweed.yaml", nextRead(t, root, "fireweed.yaml")+"limit_patterns: [RATE LIMIT]\n"+policy)
	nextCommit(t, root, "fixture usage policy")
}

// These additions run inside the existing fixture agent after it records its
// invocation and prompt. They keep all work in the fixture repository.
func nextUsageAgent(t *testing.T, f nextFixture, behavior string) {
	t.Helper()
	script := nextRead(t, f.logDir, "worker.sh")
	var addition string
	switch behavior {
	case "first":
		addition = `if [ "$phase" = evals ] && [ "$attempt" -eq 1 ]; then
  printf 'package evals\n// USAGE_PARTIAL_WORK\n' > packages/alpha/evals/partial_test.go
  printf 'limited evals'; printf 'RATE LIMIT' >&2; exit 9
fi
`
	case "twice":
		addition = `if [ "$phase" = evals ] && [ "$attempt" -le 2 ]; then
  printf 'package evals\n// USAGE_PARTIAL_WORK\n' > packages/alpha/evals/partial_test.go
  printf 'limited evals'; printf 'RATE LIMIT' >&2; exit 9
fi
`
	case "feedback":
		addition = `if [ "$phase" = evals ] && [ "$attempt" -eq 1 ]; then
  cat > packages/alpha/evals/next_test.go <<'GO'
package evals
import ("testing"; "example.invalid/next-fixture/packages/alpha")
func Test_ALP_002_NewValue(t *testing.T) { if alpha.Value() != "old" { t.Fatal("wrong") } }
GO
  printf 'package evals\n// RETAINED_EVAL_WORK\n' > packages/alpha/evals/retained_test.go
  printf 'first verification'; exit 0
fi
if [ "$phase" = evals ] && [ "$attempt" -eq 2 ]; then
  printf 'limited evals'; printf 'RATE LIMIT' >&2; exit 9
fi
`
	case "phases":
		addition = `if { [ "$phase" = evals ] || [ "$phase" = impl ]; } && [ "$attempt" -eq 1 ]; then
  if [ "$phase" = impl ]; then printf '// USAGE_IMPL_WORK\n' >> packages/alpha/internal/value.go; fi
  printf 'limited %s' "$phase"; printf 'RATE LIMIT' >&2; exit 9
fi
`
	default:
		t.Fatal("unknown fixture behavior")
	}
	anchor := "case \"$2/$phase\" in\n"
	if !strings.Contains(script, anchor) {
		t.Fatal("fixture agent insertion point missing")
	}
	nextWrite(t, f.logDir, "worker.sh", strings.Replace(script, anchor, addition+anchor, 1))
}

func runUsageAgent(t *testing.T, f runFixture, targets string) {
	t.Helper()
	script := nextRead(t, f.agents, "worker.sh")
	anchor := "if [ \"$2\" = fail-beta ]"
	addition := fmt.Sprintf(`case "$module/$phase" in
  %s)
    count_file="$1/usage-$module-$phase"
    if [ ! -f "$count_file" ]; then
      printf 1 > "$count_file"
      printf 'package evals\n// USAGE_PARTIAL_WORK\n' > "packages/$module/evals/usage_test.go"
      printf 'limited %%s %%s' "$module" "$phase"
      printf 'RATE LIMIT' >&2
      exit 9
    fi
    ;;
esac
`, targets)
	if !strings.Contains(script, anchor) {
		t.Fatal("run fixture agent insertion point missing")
	}
	nextWrite(t, f.agents, "worker.sh", strings.Replace(script, anchor, addition+anchor, 1))
}

func usageHeld(t *testing.T, root, slug, phase, run, head, claim string) workflow.State {
	t.Helper()
	state, err := workflow.Current(root, slug)
	if err != nil {
		t.Fatalf("held state: %v", err)
	}
	if state.Phase != phase || state.Run == "" || (run != "" && state.Run != run) || state.Head != head || state.ClaimOID == "" || (claim != "" && state.ClaimOID != claim) || state.Snapshot.Root != root || len(state.Change.Added) != 1 || state.Change.Added[0].ID == "" {
		t.Errorf("held identity changed: phase=%q run=%q head=%q claim=%q added=%#v", state.Phase, state.Run, state.Head, state.ClaimOID, state.Change.Added)
	}
	if got := nextGit(t, root, "rev-parse", "HEAD"); got != head {
		t.Errorf("HEAD = %s, want %s", got, head)
	}
	if got := nextGit(t, root, "for-each-ref", "--format=%(refname)", "refs/fireweed/claims"); got != "refs/fireweed/claims/"+slug {
		t.Errorf("held claim ref = %q", got)
	}
	return state
}

func Test_RUN_029_InvalidUsagePolicyPrecedesFreshStart(t *testing.T) {
	for _, tc := range []struct{ name, policy string }{
		{"bad wait", "limit_wait: nonsense\n"},
		{"bad budget", "limit_budget: nonsense\n"},
		{"zero wait", "limit_wait: 0s\n"},
		{"negative budget", "limit_budget: -1s\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := newNextFixture(t, true, true, "normal")
			usagePolicy(t, f.root, tc.policy)
			head := nextGit(t, f.root, "rev-parse", "HEAD")
			change := nextRead(t, f.root, nextChange)
			_, err := f.run(false)
			if err == nil || !strings.Contains(err.Error(), "C01") {
				t.Errorf("invalid policy error = %v, want C01", err)
			}
			if got, readErr := os.ReadFile(filepath.Join(f.root, nextChange)); readErr != nil || string(got) != change {
				t.Errorf("invalid policy changed queued document: %v", readErr)
			}
			if nextGit(t, f.root, "rev-parse", "HEAD") != head || nextGit(t, f.root, "for-each-ref", "--format=%(refname)", "refs/fireweed/claims") != "" {
				t.Error("invalid policy changed HEAD or claimed work")
			}
			f.noAgent(t)
		})
	}
}

func Test_RUN_029_RUN_031_DefaultUsageDurationsStopImmediately(t *testing.T) {
	for _, tc := range []struct{ name, policy, wait, budget string }{
		{"default wait", "limit_budget: 0s\n", "30m", "0s"},
		{"default budget", "limit_wait: 5h\n", "5h", "4h"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := newNextFixture(t, true, true, "normal")
			usagePolicy(t, f.root, tc.policy)
			nextUsageAgent(t, f, "first")
			head := nextGit(t, f.root, "rev-parse", "HEAD")
			r, err := f.run(false)
			if err == nil || !runner.IsUsageStop(err) || !strings.Contains(err.Error(), tc.wait) || !strings.Contains(err.Error(), tc.budget) || !strings.Contains(err.Error(), "0s") {
				t.Errorf("default policy stop = %v", err)
			}
			if len(r.Steps) != 1 || !r.Steps[0].Agent.UsageLimit || r.Steps[0].Agent.ExitCode != 9 || r.State.Phase != "evals" {
				t.Errorf("captured stop = %#v", r)
			}
			state := usageHeld(t, f.root, "next", "evals", r.State.Run, head, r.State.ClaimOID)
			if state.Change.Added[0].ID != "ALP-002" || !strings.Contains(nextRead(t, f.root, "packages/alpha/evals/partial_test.go"), "USAGE_PARTIAL_WORK") || len(nextInvocations(t, f)) != 1 {
				t.Error("stop lost assigned ID, partial work, or invocation")
			}
		})
	}
}

func Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep(t *testing.T) {
	f := newNextFixture(t, true, true, "normal")
	usagePolicy(t, f.root, "limit_wait: 5ms\nlimit_budget: 5ms\nattempts: 1\n")
	nextUsageAgent(t, f, "twice")
	head := nextGit(t, f.root, "rev-parse", "HEAD")
	first, err := f.run(false)
	if err == nil || !runner.IsUsageStop(err) || strings.Count(err.Error(), "5ms") < 3 || len(first.Steps) != 2 || !first.Steps[0].Agent.UsageLimit || !first.Steps[1].Agent.UsageLimit {
		t.Fatalf("short-budget stop = %#v, %v", first, err)
	}
	state := usageHeld(t, f.root, "next", "evals", first.State.Run, head, first.State.ClaimOID)
	partialWork := nextRead(t, f.root, "packages/alpha/evals/partial_test.go")
	if state.Change.Added[0].ID != "ALP-002" || !strings.Contains(partialWork, "USAGE_PARTIAL_WORK") {
		t.Error("held eval work or assigned ID lost")
	}
	approved := state.Snapshot
	if approved.Root != f.root || approved.Config["limit_wait"] != "5ms" || approved.Config["limit_budget"] != "5ms" || fmt.Sprint(approved.Config["attempts"]) != "1" {
		t.Errorf("held approved snapshot has wrong fixture policy: root=%q wait=%v budget=%v attempts=%v", approved.Root, approved.Config["limit_wait"], approved.Config["limit_budget"], approved.Config["attempts"])
	}
	approvedSpec := ""
	for _, module := range approved.Modules {
		if module.Name == "alpha" {
			approvedSpec = module.Document.Body
		}
	}
	if !strings.Contains(approvedSpec, "### ALP-002") || !strings.Contains(approvedSpec, "Alpha MUST return the new value.") {
		t.Errorf("held approved snapshot lacks projected alpha spec: %q", approvedSpec)
	}
	if got := nextGit(t, f.root, "rev-list", "--count", head+"..HEAD"); got != "0" {
		t.Errorf("stop committed %s times", got)
	}
	second, err := f.run(false)
	if err != nil || second.State.Phase != "applied" || second.State.Run != state.Run || len(second.Steps) != 3 || second.Steps[0].Phase != "evals" || len(nextInvocations(t, f)) != 5 {
		t.Errorf("owned resume = %#v, %v", second, err)
	}
	if !strings.Contains(nextRead(t, f.root, "changes/archive/0001-next.md"), "ALP-002") {
		t.Error("resume lost assigned ID")
	}
	if got := nextRead(t, f.root, "packages/alpha/evals/partial_test.go"); got != partialWork {
		t.Errorf("resume changed held partial eval work: got %q, want %q", got, partialWork)
	}
	resumedSpec := ""
	for _, module := range second.State.Snapshot.Modules {
		if module.Name == "alpha" {
			resumedSpec = module.Document.Body
		}
	}
	if second.State.Snapshot.Root != approved.Root || second.State.Snapshot.Config["limit_wait"] != approved.Config["limit_wait"] || second.State.Snapshot.Config["limit_budget"] != approved.Config["limit_budget"] || fmt.Sprint(second.State.Snapshot.Config["attempts"]) != fmt.Sprint(approved.Config["attempts"]) || resumedSpec != approvedSpec {
		t.Errorf("resumed State did not use held approved snapshot: root=%q wait=%v budget=%v attempts=%v projected alpha=%q", second.State.Snapshot.Root, second.State.Snapshot.Config["limit_wait"], second.State.Snapshot.Config["limit_budget"], second.State.Snapshot.Config["attempts"], resumedSpec)
	}
}

func Test_RUN_030_UsageRetryPreservesVerificationFeedback(t *testing.T) {
	f := newNextFixture(t, true, true, "normal")
	usagePolicy(t, f.root, "limit_wait: 5ms\nlimit_budget: 20ms\nattempts: 2\n")
	nextUsageAgent(t, f, "feedback")
	r, err := f.run(false)
	if err != nil || r.State.Phase != "applied" || len(r.Steps) != 5 {
		t.Fatalf("verification and usage retry = %#v, %v", r, err)
	}
	if r.Steps[0].Agent.Stdout != "first verification" || !r.Steps[1].Agent.UsageLimit || r.Steps[1].Phase != "evals" || r.Steps[2].Phase != "evals" || r.Steps[2].Agent.ExitCode != 0 {
		t.Errorf("evals attempt order = %#v", r.Steps)
	}
	if prompt := nextRead(t, f.logDir, "worker-evals-3.prompt"); !strings.Contains(prompt, "P2-01: no fail-first case for ALP-002") {
		t.Errorf("post-limit retry lost verification feedback: %q", prompt)
	}
	if !strings.Contains(nextRead(t, f.root, "packages/alpha/evals/retained_test.go"), "RETAINED_EVAL_WORK") {
		t.Error("retry lost earlier eval work")
	}
}

func Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases(t *testing.T) {
	f := newNextFixture(t, true, true, "normal")
	usagePolicy(t, f.root, "limit_wait: 5ms\nlimit_budget: 5ms\nattempts: 1\n")
	nextUsageAgent(t, f, "phases")
	r, err := f.run(false)
	if err == nil || !runner.IsUsageStop(err) || !strings.Contains(err.Error(), "5ms") || len(r.Steps) != 4 {
		t.Fatalf("cross-phase budget = %#v, %v", r, err)
	}
	for i, want := range []string{"evals", "evals", "audit", "impl"} {
		if r.Steps[i].Phase != want {
			t.Errorf("step %d phase = %q, want %q", i, r.Steps[i].Phase, want)
		}
	}
	if !r.Steps[0].Agent.UsageLimit || !r.Steps[3].Agent.UsageLimit || r.State.Phase != "impl" {
		t.Errorf("held phase/steps = %#v", r)
	}
	usageHeld(t, f.root, "next", "impl", r.State.Run, r.State.Head, r.State.ClaimOID)
	if !strings.Contains(nextRead(t, f.root, "packages/alpha/internal/value.go"), "USAGE_IMPL_WORK") {
		t.Error("impl partial work lost")
	}
}

func Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied(t *testing.T) {
	f := newRunFixture(t, 3, false, false)
	usagePolicy(t, f.root, "limit_wait: 5ms\nlimit_budget: 5ms\n")
	runUsageAgent(t, f, "alpha/evals|beta/evals")
	r, err := runner.Run(context.Background(), f.request(0, false))
	if err == nil || !runner.IsUsageStop(err) || !strings.Contains(err.Error(), "5ms") {
		t.Errorf("Run stop = %v", err)
	}
	assertRunApplied(t, f, r.Applied, "alpha")
	state := usageHeld(t, f.root, "beta", "evals", "", nextGit(t, f.root, "rev-parse", "HEAD"), "")
	if state.Change.Added[0].ID != "BET-002" || !strings.Contains(nextRead(t, f.root, "packages/beta/evals/usage_test.go"), "USAGE_PARTIAL_WORK") {
		t.Error("Run lost held beta ID or work")
	}
	if got := nextRead(t, f.root, "changes/0003-gamma.md"); strings.Contains(got, "phase:") || strings.Contains(got, "run:") {
		t.Errorf("later change started: %q", got)
	}
	if got := runInvocations(t, f); len(got) != 5 {
		t.Errorf("agent calls = %#v, want two alpha evals, audit, impl, beta evals", got)
	}
}

func Test_RUN_034_UsageStopClassifierWrapping(t *testing.T) {
	if runner.IsUsageStop(nil) || runner.IsUsageStop(errors.New("ordinary failure")) {
		t.Error("ordinary or nil error classified as clean stop")
	}
	f := newNextFixture(t, true, true, "normal")
	usagePolicy(t, f.root, "limit_wait: 5ms\nlimit_budget: 0s\n")
	nextUsageAgent(t, f, "first")
	_, err := f.run(false)
	if err == nil || !runner.IsUsageStop(fmt.Errorf("outer: %w", err)) {
		t.Errorf("wrapped clean stop classification = %v", err)
	}
}

func Test_RUN_035_EarlierBlockMakesLaterUsageStopFailure(t *testing.T) {
	f := runBlockedDependencyFixture(t)
	usagePolicy(t, f.root, "limit_wait: 5ms\nlimit_budget: 0s\n")
	runUsageAgent(t, f, "gamma/evals")
	r, err := runner.Run(context.Background(), f.request(0, false))
	if err == nil || !strings.Contains(err.Error(), "P2-01: no fail-first case for ALP-002") || !strings.Contains(err.Error(), "refs/fireweed/attempts/alpha/1") || !strings.Contains(err.Error(), "5ms") || runner.IsUsageStop(err) || len(r.Applied) != 0 {
		t.Errorf("combined block and usage failure = %#v, %v", r, err)
	}
	if saved := nextGit(t, f.root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/attempts/alpha/1"); saved == "" {
		t.Error("blocked alpha lost its saved attempt ref")
	}
	if blocked := nextRead(t, f.root, "changes/0001-alpha.md"); !strings.Contains(blocked, "blocked:") || !strings.Contains(blocked, "phase: evals") {
		t.Errorf("alpha block state = %q", blocked)
	}
	assertRunBetaWaiting(t, f)
	usageHeld(t, f.root, "gamma", "evals", "", nextGit(t, f.root, "rev-parse", "HEAD"), "")
}

func Test_RUN_031_CancelledUsageWaitRetainsHeldWork(t *testing.T) {
	f := newNextFixture(t, true, true, "normal")
	usagePolicy(t, f.root, "limit_wait: 5s\nlimit_budget: 5s\n")
	nextUsageAgent(t, f, "first")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	r, err := runner.Next(ctx, runner.NextRequest{Root: f.root, Bin: "/fixture/next-bin", Timeout: 20 * time.Second})
	if !errors.Is(err, context.DeadlineExceeded) || runner.IsUsageStop(err) || r.State.Phase != "evals" || len(r.Steps) != 1 {
		t.Errorf("cancelled wait = %#v, %v", r, err)
	}
	usageHeld(t, f.root, "next", "evals", r.State.Run, r.State.Head, r.State.ClaimOID)
}
