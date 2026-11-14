package evals_test

import (
	"context"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner"
	"github.com/lurkshark/fireweed/packages/workflow"
)

const warningBase = "refs/remotes/origin/main"

func setNewerBaseTip(t *testing.T, f nextFixture) string {
	t.Helper()
	head := nextGit(t, f.root, "rev-parse", "HEAD")
	// Advance the tracking ref beyond main without changing which files it
	// contains or moving the checkout's HEAD.
	tip := nextGit(t, f.root, "commit-tree", "main^{tree}", "-p", "main", "-m", "newer base tip")
	nextGit(t, f.root, "update-ref", warningBase, tip)
	return head
}

func isBaseTipDiagnostic(d project.Diagnostic) bool {
	return strings.Contains(d.Message, warningBase)
}

func assertStartWarning(t *testing.T, diagnostics []project.Diagnostic, want project.Diagnostic) {
	t.Helper()
	var got []project.Diagnostic
	for _, d := range diagnostics {
		if isBaseTipDiagnostic(d) {
			got = append(got, d)
		}
	}
	if !slices.Equal(got, []project.Diagnostic{want}) {
		t.Errorf("base-tip diagnostics = %#v, want exactly one unchanged Start warning %#v; all diagnostics = %#v", got, want, diagnostics)
	}
}

func assertAppliedWarningResult(t *testing.T, f nextFixture, result runner.NextResult) {
	t.Helper()
	state := result.State
	if result.Empty || state.Phase != "applied" || state.Run == "" || state.Change.Slug != "next" || state.ClaimOID != "" || len(result.Steps) != 3 {
		t.Fatalf("result did not apply the change: phase=%q run=%q slug=%q claim=%q empty=%t steps=%d", state.Phase, state.Run, state.Change.Slug, state.ClaimOID, result.Empty, len(result.Steps))
	}
	archive := nextRead(t, f.root, "changes/archive/0001-next.md")
	if !strings.Contains(archive, "phase: applied\n") || !strings.Contains(archive, "run: "+state.Run+"\n") {
		t.Error("archive did not retain applied phase and run")
	}
}

func Test_RUN_072_RetainsStartWarning(t *testing.T) {
	reference := newNextFixture(t, true, false, "normal")
	setNewerBaseTip(t, reference)
	started, err := workflow.Start(reference.root, nextChange)
	if err != nil {
		t.Fatalf("Start in reference fixture: %v", err)
	}
	index := slices.IndexFunc(started.Diagnostics, isBaseTipDiagnostic)
	if index < 0 || started.Diagnostics[index].Code != "P05" || !started.Diagnostics[index].Warning {
		t.Fatalf("reference Start did not return the base-tip P05 warning: %#v", started.Diagnostics)
	}
	want := started.Diagnostics[index]
	isUnmerged := func(d project.Diagnostic) bool { return d.Code == "P04" }
	unmergedIndex := slices.IndexFunc(started.Diagnostics, isUnmerged)
	if unmergedIndex < 0 || !started.Diagnostics[unmergedIndex].Warning {
		t.Fatalf("reference Start did not warn about the unmerged change: %#v", started.Diagnostics)
	}
	wantUnmerged := started.Diagnostics[unmergedIndex]

	t.Run("applied with worker recovery", func(t *testing.T) {
		f := newNextFixture(t, true, true, "commit-success")
		setNewerBaseTip(t, f)
		result, err := f.run(false)
		if err != nil {
			t.Fatalf("Next: %v", err)
		}
		assertAppliedWarningResult(t, f, result)
		assertStartWarning(t, result.State.Diagnostics, want)
		recovery := slices.DeleteFunc(slices.Clone(result.State.Diagnostics), isBaseTipDiagnostic)
		assertRecoveryWarnings(t, recovery, "evals", "impl")
	})

	t.Run("verification exhaustion", func(t *testing.T) {
		f := newNextFixture(t, true, true, "exhaust-impl")
		head := setNewerBaseTip(t, f)
		result, err := f.run(false)
		if code, _, _ := contracts.ErrorCode(err); code != "E01" || !strings.Contains(err.Error(), nextValueCase) || len(result.Steps) != 4 {
			t.Fatalf("expected implementation verification exhaustion, got %d steps and error %v", len(result.Steps), err)
		}
		nextAssertExhaustedBlock(t, f, result, err, "impl", head, 2)
		assertStartWarning(t, result.State.Diagnostics, want)
	})

	for _, tc := range []struct {
		name, output string
		usageStop    bool
	}{
		{"invocation failure", "auditor failed", false},
		{"usage stop", "RATE LIMIT", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := newNextFixture(t, true, true, "normal")
			if tc.usageStop {
				usagePolicy(t, f.root, "limit_budget: 0s\n")
			}
			head := setNewerBaseTip(t, f)
			auditBeforeEdit(t, f, "printf '%s' "+shellQuote(tc.output)+" >&2\nexit 9\n")
			result, err := f.run(false)
			if err == nil || runner.IsUsageStop(err) != tc.usageStop || len(result.Steps) != 2 || result.Steps[0].Phase != "evals" {
				t.Fatalf("expected a stop in audit, got %d steps and error %v", len(result.Steps), err)
			}
			if step := result.Steps[1]; step.Phase != "audit" || step.Agent.ExitCode != 9 || step.Agent.Stderr != tc.output {
				t.Errorf("captured audit invocation = %#v", step)
			}
			state := result.State
			if state.Phase != "audit" || state.Run == "" || state.ClaimOID == "" || nextGit(t, f.root, "rev-list", "--count", head+"..HEAD") != "1" {
				t.Fatalf("stop did not retain the audit step after evals: phase=%q run=%q claim=%q", state.Phase, state.Run, state.ClaimOID)
			}
			usageHeld(t, f.root, "next", "audit", state.Run, state.Head, state.ClaimOID)
			assertStartWarning(t, state.Diagnostics, want)

			if tc.usageStop {
				resumed, err := f.run(false)
				if !runner.IsUsageStop(err) || resumed.State.Run != state.Run || resumed.State.ClaimOID != state.ClaimOID || resumed.State.Phase != "audit" || len(resumed.Steps) != 1 {
					t.Fatalf("expected to resume held audit and stop again, got phase=%q run=%q steps=%d err=%v", resumed.State.Phase, resumed.State.Run, len(resumed.Steps), err)
				}
				if slices.ContainsFunc(resumed.State.Diagnostics, isBaseTipDiagnostic) {
					t.Errorf("resumed invocation returned the earlier Start warning: %#v", resumed.State.Diagnostics)
				}
			}
		})
	}

	t.Run("Run applied result", func(t *testing.T) {
		f := newNextFixture(t, true, false, "normal")
		setNewerBaseTip(t, f)
		result, err := runner.Run(context.Background(), runner.RunRequest{
			Root: f.root, Bin: "/fixture/next-bin", AllowUnmerged: true, Max: 1, Timeout: 20 * time.Second,
		})
		if err != nil || !result.Empty || result.Limited || len(result.Applied) != 1 {
			t.Fatalf("expected Run to apply one queued change, got %d applied, empty=%t, limited=%t, err=%v", len(result.Applied), result.Empty, result.Limited, err)
		}
		assertAppliedWarningResult(t, f, result.Applied[0])
		diagnostics := result.Applied[0].State.Diagnostics
		assertStartWarning(t, diagnostics, want)
		var unmerged []project.Diagnostic
		for _, d := range diagnostics {
			if isUnmerged(d) {
				unmerged = append(unmerged, d)
			}
		}
		if !slices.Equal(unmerged, []project.Diagnostic{wantUnmerged}) {
			t.Errorf("unmerged diagnostics = %#v, want exactly one unchanged Start warning %#v", unmerged, wantUnmerged)
		}
	})
}
