package evals_test

import (
	"context"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/checks"
)

func Test_CHK_084_CanaryNamesBeginWithProbeAndHookScope(t *testing.T) {
	f := localDoctorFixture(t)
	spec := string(mustRead(t, filepath.Join(f.root, "packages/alpha/SPEC.md")))
	f.put("packages/alpha/SPEC.md", strings.Replace(spec, "depends_on: []", `depends_on: []
hooks:
  eval: sh .fireweed/eval.sh {module} {module_dir} {report}
  boundaries: sh .fireweed/boundary.sh`, 1))
	f.commit()

	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) {
		t.Fatalf("canary fixture failed: %s", doctorText(report))
	}
	want := map[string]bool{}
	for _, hook := range []struct {
		kind     string
		canaries []string
	}{
		{"eval", []string{"failing_eval", "missing_id"}},
		{"boundaries", []string{"undeclared", "internal_eval"}},
	} {
		for _, scope := range []string{"global " + hook.kind, "alpha " + hook.kind} {
			for _, probe := range append([]string{"baseline", "recovery"}, hook.canaries...) {
				want[probe+" "+scope] = false
			}
		}
	}
	for _, check := range report.Checks {
		matched := false
		for prefix := range want {
			if check.Name == prefix || strings.HasPrefix(check.Name, prefix+" ") {
				want[prefix] = true
				matched = true
				break
			}
		}
		if !matched {
			t.Errorf("canary name %q does not begin with its probe and hook scope", check.Name)
		}
	}
	for prefix, seen := range want {
		if !seen {
			t.Errorf("missing canary name beginning with %q", prefix)
		}
	}
}

func Test_CHK_084_FailedRecoveryNameBeginsWithProbeAndHookScope(t *testing.T) {
	f := localDoctorFixture(t)
	spec := string(mustRead(t, filepath.Join(f.root, "packages/alpha/SPEC.md")))
	f.put("packages/alpha/SPEC.md", strings.Replace(spec, "depends_on: []", `depends_on: []
hooks:
  eval: sh .fireweed/override-eval.sh {module} {module_dir} {report}`, 1))
	f.put(".fireweed/override-eval.sh", `#!/bin/sh
if test -f "$2/evals/failing.txt"; then
  touch .fireweed/recovery-pending
elif test -f .fireweed/recovery-pending; then
  rm .fireweed/recovery-pending
  printf '<testsuite/>' > "$3"
  exit 0
fi
exec sh .fireweed/eval.sh "$@"
`)
	f.commit()

	report := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(report, "C14", false, "recovery", "failing_eval", "eval") {
		t.Fatalf("fixture did not report the induced recovery failure: %s", doctorText(report))
	}
	const want = "recovery alpha eval"
	failed := false
	for _, check := range report.Checks {
		if check.Passed {
			continue
		}
		failed = true
		if check.Name != want && !strings.HasPrefix(check.Name, want+" ") {
			t.Errorf("failed recovery name = %q; want it to begin with %q", check.Name, want)
		}
	}
	if !failed {
		t.Errorf("missing failed recovery check: %s", doctorText(report))
	}
}
