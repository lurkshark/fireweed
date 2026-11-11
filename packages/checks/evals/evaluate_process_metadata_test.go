package evals_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
)

func Test_CHK_034_EvaluateRetainsHookOutputOnRejectedReports(t *testing.T) {
	for _, tc := range []struct {
		name, reportCommand, cause string
	}{
		{
			name:          "unmapped PackageFailure",
			reportCommand: `printf '%s' '<testsuite><testcase name="CHK-003_valid"/><testcase name="PackageFailure"><failure message="compiler failed"/></testcase></testsuite>' > {report};`,
			cause:         `testcase "PackageFailure" has no known spec ID`,
		},
		{
			name:          "malformed XML",
			reportCommand: `printf '%s' '<testsuite><testcase name="CHK-003"></testsuite>' > {report};`,
			cause:         "parse JUnit:",
		},
		{
			name:  "missing report",
			cause: "read JUnit report:",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			stdout, stderr := "hook stdout marker\n", "hook stderr marker\n"
			command := `printf 'hook stdout marker\n'; printf 'hook stderr marker\n' >&2; ` + tc.reportCommand + `exit 7`
			snap := hookSnapshot(t, command, nil)
			report, err := checks.Evaluate(context.Background(), checks.HookRequest{
				Snapshot: snap, Module: "alpha", Timeout: time.Second,
			})
			requireDiagnosticError(t, err, "C11", tc.cause)
			if report.Stdout != stdout || report.Stderr != stderr || report.HookExit != 7 {
				t.Errorf("rejected report process metadata = %#v; want stdout %q, stderr %q, exit 7", report, stdout, stderr)
			}
			if len(report.Cases) != 0 || report.Seconds != 0 {
				t.Errorf("rejected JUnit supplied testcase evidence: %#v", report)
			}
			if tc.name == "unmapped PackageFailure" {
				specs := []contracts.Item{{ID: "CHK-003"}}
				requireDiagnosticError(t, checks.RequireCoverage(report, specs), "C10", "no coverage for CHK-003")
				requireFailFirstError(t, checks.RequireFailFirst(report, []string{"CHK-003"}, nil), "CHK-003")
			}
		})
	}
}

func Test_CHK_034_EvaluateRetainsOutputOnHookTimeout(t *testing.T) {
	snap := hookSnapshot(t, `printf 'timeout stdout marker\n'; printf 'timeout stderr marker\n' >&2; sleep 2`, nil)
	start := time.Now()
	report, err := checks.Evaluate(context.Background(), checks.HookRequest{
		Snapshot: snap, Module: "alpha", Timeout: 100 * time.Millisecond,
	})
	requireDiagnosticError(t, err, "C11", "run hook:")
	if report.Stdout != "timeout stdout marker\n" || report.Stderr != "timeout stderr marker\n" || report.HookExit == 0 {
		t.Errorf("timed-out report lost process metadata: %#v", report)
	}
	if len(report.Cases) != 0 || report.Seconds != 0 {
		t.Errorf("timed-out hook supplied testcase evidence: %#v", report)
	}
	if elapsed := time.Since(start); elapsed > 3*time.Second {
		t.Errorf("hook timeout took %s", elapsed)
	}
	if err != nil && !strings.Contains(err.Error(), "deadline exceeded") {
		t.Errorf("hook timeout cause lost: %v", err)
	}
}
