package evals_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/runner"
	"github.com/lurkshark/fireweed/packages/workflow"
)

// The ordinary fixture hook supplies fixed phase evidence so these cases measure
// Run scheduling rather than Go compilation time. The agent still writes work.
func timeFastHook(t *testing.T, f runFixture) {
	t.Helper()
	nextWrite(t, f.root, ".fireweed/eval.sh", `#!/bin/sh
case "$1" in alpha) prefix=ALP;; beta) prefix=BET;; gamma) prefix=GAM;; *) exit 35;; esac
phase=''; active=''
for file in changes/*.md; do
  [ -f "$file" ] || continue
  [ -z "$(sed -n 's/^blocked: //p' "$file" | head -1)" ] || continue
  case "$(sed -n 's/^phase: //p' "$file" | head -1)" in
    evals|impl) phase=$(sed -n 's/^phase: //p' "$file" | head -1); active=$(sed -n 's/^module: //p' "$file" | head -1); break;;
  esac
done
[ -n "$phase" ] || exit 36
if [ "$1" != "$active" ]; then
  printf '<testsuite><testcase name="%s-001" classname="fixture"/></testsuite>' "$prefix" > "$2"
  exit 0
fi
if [ "$phase" = evals ] && [ ! -f "packages/$1/evals/added_test.go" ]; then
  printf '<testsuite><testcase name="%s-001" classname="fixture"/><testcase name="%s-002" classname="fixture"/></testsuite>' "$prefix" "$prefix" > "$2"
  exit 0
fi
if [ "$phase" = evals ]; then result='<failure message="new value not implemented"/>'; code=1
else result=''; code=0; fi
printf '<testsuite><testcase name="%s-001" classname="fixture"/><testcase name="%s-002" classname="fixture">%s</testcase></testsuite>' "$prefix" "$prefix" "$result" > "$2"
exit "$code"
`)
	config := nextRead(t, f.root, "fireweed.yaml")
	nextWrite(t, f.root, "fireweed.yaml", strings.Replace(config, "timeout: 50ms\n", "timeout: 15s\n", 1))
	nextCommit(t, f.root, "bounded local time hook")
}

func timeAgentAfterWork(t *testing.T, f runFixture, body string) {
	t.Helper()
	script := nextRead(t, f.agents, "worker.sh")
	anchor := "printf '%s-%s' \"$module\" \"$phase\"\n"
	if !strings.Contains(script, anchor) {
		t.Fatal("fixture agent insertion point missing")
	}
	nextWrite(t, f.agents, "worker.sh", strings.Replace(script, anchor, body+anchor, 1))
}

func timeStop(t *testing.T, err error, limit string) {
	t.Helper()
	if err == nil || !runner.IsTimeStop(err) || !strings.Contains(err.Error(), limit) {
		t.Errorf("clean time stop = %v, want classified nonnil error naming %s", err, limit)
	}
}

func timeHeldWork(t *testing.T, f runFixture, slug, phase, marker, want, head string) workflow.State {
	t.Helper()
	state := usageHeld(t, f.root, slug, phase, "", head, "")
	if !strings.Contains(nextRead(t, f.root, marker), want) {
		t.Errorf("interrupted %s work missing", slug)
	}
	if got := nextGit(t, f.root, "rev-list", "--count", head+"..HEAD"); got != "0" {
		t.Errorf("interrupted step committed %s times", got)
	}
	if state.Change.Added[0].ID != runModules[indexRunModule(slug)].prefix+"-002" {
		t.Errorf("held %s ID = %q", slug, state.Change.Added[0].ID)
	}
	return state
}

func Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied(t *testing.T) {
	f := newRunFixture(t, 2, false, false)
	timeFastHook(t, f)
	// Alpha consumes more than the elapsed-time margin, so a fresh budget for
	// beta would exceed the bound even when the rest of the fixture runs quickly.
	const margin = 4 * time.Second
	timeAgentAfterWork(t, f, `if [ "$module" = alpha ] && [ "$phase" = impl ]; then sleep 6; fi
if [ "$module" = beta ] && [ "$phase" = evals ]; then
  printf 'TIME_HELD_WORK\n' > packages/beta/evals/time-work.txt
  sleep 120
fi
`)
	req := f.request(0, false)
	req.For = 30 * time.Second
	req.Timeout = time.Minute
	start := time.Now()
	r, err := runner.Run(context.Background(), req)
	elapsed := time.Since(start)
	timeStop(t, err, req.For.String())
	if elapsed > req.For+margin {
		t.Errorf("total deadline waited %s; budget may have reset after alpha", elapsed)
	}
	if len(r.Applied) == 1 {
		assertRunApplied(t, f, r.Applied, "alpha")
		timeHeldWork(t, f, "beta", "evals", "packages/beta/evals/time-work.txt", "TIME_HELD_WORK", nextGit(t, f.root, "rev-parse", "HEAD"))
	} else {
		t.Errorf("Applied count = %d, want one retained alpha", len(r.Applied))
	}
	if got := runInvocations(t, f); len(got) != 4 || !strings.HasPrefix(got[3], "beta|evals|") {
		t.Errorf("invocations through held beta = %#v", got)
	}
}

func Test_RUN_036_ZeroForKeepsOrdinaryRunUnbounded(t *testing.T) {
	f := newRunFixture(t, 1, false, false)
	timeFastHook(t, f)
	req := f.request(0, false)
	req.For = 0
	r, err := runner.Run(context.Background(), req)
	if err != nil || !r.Empty || runner.IsTimeStop(err) {
		t.Fatalf("zero For run = %#v, %v", r, err)
	}
	assertRunApplied(t, f, r.Applied, "alpha")
}

func Test_RUN_037_RUN_039_AgentDeadlineRetainsHeldWork(t *testing.T) {
	f := newRunFixture(t, 1, false, false)
	timeFastHook(t, f)
	timeAgentAfterWork(t, f, `if [ "$phase" = evals ]; then
  printf 'TIME_HELD_WORK\n' > packages/alpha/evals/time-work.txt
  sleep 120
fi
`)
	head := nextGit(t, f.root, "rev-parse", "HEAD")
	original, startErr := workflow.Start(f.root, "alpha")
	if startErr != nil {
		t.Fatalf("start interrupted agent step: %v", startErr)
	}
	req := f.request(0, false)
	req.For = 10 * time.Second
	r, err := runner.Run(context.Background(), req)
	timeStop(t, err, req.For.String())
	if err != nil && !runner.IsTimeStop(fmt.Errorf("outer: %w", err)) {
		t.Errorf("wrapped time stop classification = %v", err)
	}
	if len(r.Applied) != 0 {
		t.Errorf("interrupted agent applied %d changes", len(r.Applied))
	} else {
		held := timeHeldWork(t, f, "alpha", "evals", "packages/alpha/evals/time-work.txt", "TIME_HELD_WORK", head)
		if held.Run != original.Run || held.ClaimOID != original.ClaimOID || held.Phase != original.Phase || held.Change.Added[0].ID != original.Change.Added[0].ID || held.Head != original.Head {
			t.Errorf("interrupted agent identity changed: held=%#v, original=%#v", held, original)
		}
	}
}

func Test_RUN_037_VerificationDeadlineRetainsHeldWork(t *testing.T) {
	f := newRunFixture(t, 1, false, false)
	timeFastHook(t, f)
	hook := nextRead(t, f.root, ".fireweed/eval.sh")
	anchor := "if [ \"$phase\" = evals ]; then result="
	addition := `if [ "$phase" = evals ]; then
  printf 'TIME_HELD_WORK\n' > packages/alpha/evals/time-work.txt
  sleep 120
fi
`
	if !strings.Contains(hook, anchor) {
		t.Fatal("fixture hook insertion point missing")
	}
	nextWrite(t, f.root, ".fireweed/eval.sh", strings.Replace(hook, anchor, addition+anchor, 1))
	nextCommit(t, f.root, "bounded verification interruption")
	head := nextGit(t, f.root, "rev-parse", "HEAD")
	original, startErr := workflow.Start(f.root, "alpha")
	if startErr != nil {
		t.Fatalf("start interrupted verification step: %v", startErr)
	}
	req := f.request(0, false)
	req.For = 10 * time.Second
	r, err := runner.Run(context.Background(), req)
	invocations, invocationErr := os.ReadFile(filepath.Join(f.agents, "invocations"))
	workPath := filepath.Join(f.root, "packages/alpha/evals/added_test.go")
	workInfo, workStatErr := os.Stat(workPath)
	work, workReadErr := os.ReadFile(workPath)
	markerPath := filepath.Join(f.root, "packages/alpha/evals/time-work.txt")
	markerInfo, markerStatErr := os.Stat(markerPath)
	marker, markerReadErr := os.ReadFile(markerPath)
	reportPath := filepath.Join(f.root, ".fireweed", "runs", original.Run, "evals", "alpha.json")
	reportInfo, reportStatErr := os.Stat(reportPath)
	report, reportReadErr := os.ReadFile(reportPath)
	held, heldErr := workflow.Current(f.root, "alpha")
	readGit := func(args ...string) (string, error) {
		command := exec.Command("git", args...)
		command.Dir = f.root
		output, commandErr := command.Output()
		return strings.TrimSpace(string(output)), commandErr
	}
	actualHead, headErr := readGit("rev-parse", "HEAD")
	claimRef, claimErr := readGit("for-each-ref", "--format=%(refname)", "refs/fireweed/claims")
	commits, commitsErr := readGit("rev-list", "--count", head+"..HEAD")
	t.Logf("RUN-037 verification stop: result=%#v error=%v clean_time_stop=%t", r, err, runner.IsTimeStop(err))
	t.Logf("agent invocations: read_error=%v content=%q", invocationErr, invocations)
	t.Logf("agent work %s: exists=%t stat_error=%v read_error=%v content=%q", workPath, workStatErr == nil && workInfo != nil, workStatErr, workReadErr, work)
	t.Logf("hook marker %s: exists=%t stat_error=%v read_error=%v content=%q", markerPath, markerStatErr == nil && markerInfo != nil, markerStatErr, markerReadErr, marker)
	t.Logf("verification report %s: exists=%t stat_error=%v read_error=%v content=%q", reportPath, reportStatErr == nil && reportInfo != nil, reportStatErr, reportReadErr, report)
	t.Logf("held identity: state=%#v read_error=%v original phase=%q run=%q claim=%q assigned=%#v HEAD=%q; actual HEAD=%q error=%v; claim refs=%q error=%v; commits=%q error=%v", held, heldErr, original.Phase, original.Run, original.ClaimOID, original.Change.Added, original.Head, actualHead, headErr, claimRef, claimErr, commits, commitsErr)
	timeStop(t, err, req.For.String())
	if len(r.Applied) != 0 {
		t.Errorf("interrupted verification applied %d changes", len(r.Applied))
	}
	if markerReadErr != nil || !strings.Contains(string(marker), "TIME_HELD_WORK") {
		t.Errorf("interrupted alpha work missing: marker=%q read_error=%v", marker, markerReadErr)
	}
	if workStatErr != nil {
		t.Errorf("agent work lost before interrupted verification: %v", workStatErr)
	}
	if commitsErr != nil || commits != "0" {
		t.Errorf("interrupted step committed %s times: %v", commits, commitsErr)
	}
	if heldErr != nil {
		t.Errorf("held state: %v", heldErr)
	} else {
		if held.Phase != "evals" || held.Run == "" || held.Head != head || held.ClaimOID == "" || held.Snapshot.Root != f.root || len(held.Change.Added) != 1 || (len(held.Change.Added) == 1 && held.Change.Added[0].ID != "ALP-002") {
			t.Errorf("held identity changed: phase=%q run=%q head=%q claim=%q added=%#v snapshot=%q", held.Phase, held.Run, held.Head, held.ClaimOID, held.Change.Added, held.Snapshot.Root)
		}
		if held.Run != original.Run || held.ClaimOID != original.ClaimOID || held.Phase != original.Phase || len(held.Change.Added) != len(original.Change.Added) || (len(held.Change.Added) == 1 && len(original.Change.Added) == 1 && held.Change.Added[0].ID != original.Change.Added[0].ID) || held.Head != original.Head {
			t.Errorf("interrupted verification identity changed: held=%#v, original=%#v", held, original)
		}
	}
	if headErr != nil || actualHead != head {
		t.Errorf("HEAD = %s, want %s: %v", actualHead, head, headErr)
	}
	if claimErr != nil || claimRef != "refs/fireweed/claims/alpha" {
		t.Errorf("held claim ref = %q: %v", claimRef, claimErr)
	}
}

func Test_RUN_038_UsageWaitPreflightStopsBeforeWaiting(t *testing.T) {
	f := newRunFixture(t, 1, false, false)
	timeFastHook(t, f)
	usagePolicy(t, f.root, "limit_wait: 1m\nlimit_budget: 2m\n")
	runUsageAgent(t, f, "alpha/evals")
	head := nextGit(t, f.root, "rev-parse", "HEAD")
	req := f.request(0, false)
	req.For = 20 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), req.For+5*time.Second)
	defer cancel()
	start := time.Now()
	r, err := runner.Run(ctx, req)
	elapsed := time.Since(start)
	timeStop(t, err, req.For.String())
	if elapsed >= req.For/2 {
		t.Errorf("usage wait preflight took %s", elapsed)
	}
	if len(r.Applied) != 0 {
		t.Errorf("preflight applied %d changes", len(r.Applied))
	} else {
		timeHeldWork(t, f, "alpha", "evals", "packages/alpha/evals/usage_test.go", "USAGE_PARTIAL_WORK", head)
	}
}

func Test_RUN_039_ClassifierAndParentCancellation(t *testing.T) {
	if runner.IsTimeStop(nil) || runner.IsTimeStop(errors.New("ordinary")) || runner.IsTimeStop(context.Canceled) {
		t.Error("ordinary, nil or parent cancellation classified as time stop")
	}
	f := newRunFixture(t, 1, false, false)
	timeFastHook(t, f)
	usagePolicy(t, f.root, "limit_wait: 1m\nlimit_budget: 0s\n")
	runUsageAgent(t, f, "alpha/evals")
	req := f.request(0, false)
	_, usageErr := runner.Run(context.Background(), req)
	if !runner.IsUsageStop(usageErr) || runner.IsTimeStop(usageErr) {
		t.Errorf("usage stop classification = %v", usageErr)
	}

	other := newRunFixture(t, 1, false, false)
	timeFastHook(t, other)
	timeAgentAfterWork(t, other, `if [ "$phase" = evals ]; then sleep 8; fi
`)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	parentReq := other.request(0, false)
	parentReq.For = 10 * time.Second
	_, parentErr := runner.Run(ctx, parentReq)
	if !errors.Is(parentErr, context.DeadlineExceeded) || runner.IsTimeStop(parentErr) {
		t.Errorf("parent cancellation = %v", parentErr)
	}
}

func Test_RUN_040_EarlierBlockRetainsCauseBeforeTimeStop(t *testing.T) {
	f := runBlockedDependencyFixture(t)
	timeFastHook(t, f)
	timeAgentAfterWork(t, f, `if [ "$module" = gamma ] && [ "$phase" = evals ]; then
  printf 'TIME_HELD_WORK\n' > packages/gamma/evals/time-work.txt
  sleep 120
fi
`)
	req := f.request(0, false)
	req.For = 20 * time.Second
	req.Timeout = time.Minute
	r, err := runner.Run(context.Background(), req)
	if err == nil || !strings.Contains(err.Error(), "P2-01: no fail-first case for ALP-002") || !strings.Contains(err.Error(), "refs/fireweed/attempts/alpha/1") || !strings.Contains(err.Error(), req.For.String()) || runner.IsTimeStop(err) {
		t.Errorf("combined block and time failure: Applied=%d, err=%v", len(r.Applied), err)
	}
	if saved := nextGit(t, f.root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/attempts/alpha/1"); saved == "" {
		t.Error("blocked alpha lost saved attempt ref")
	}
	if len(r.Applied) != 0 {
		t.Errorf("blocked run applied %d changes", len(r.Applied))
	} else {
		timeHeldWork(t, f, "gamma", "evals", "packages/gamma/evals/time-work.txt", "TIME_HELD_WORK", nextGit(t, f.root, "rev-parse", "HEAD"))
	}
	assertRunBetaWaiting(t, f)
}
