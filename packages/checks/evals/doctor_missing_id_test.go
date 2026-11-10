package evals_test

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/checks"
	"gopkg.in/yaml.v3"
)

func Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary(t *testing.T) {
	const canaryPath = "evals/canary_noid_eval_test.go"
	const canarySource = `package evals

import "testing"

func TestMissingIDCanary(t *testing.T) {
	t.Run("child", func(t *testing.T) {})
}
`
	for _, tc := range []struct {
		name, reported string
		detected       bool
	}{
		{name: "declared_test", detected: true},
		{name: "subtest", reported: "TestMissingIDCanary/child", detected: true},
		{name: "unrelated_test_with_shared_prefix", reported: "TestMissingIDCanaryExtra"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var f doctorFixture
			if tc.reported == "" {
				f = newDoctorFixture(t)
				f.bin = doctorBinary(t)
			} else {
				f = localDoctorFixture(t)
				// Isolate the testcase reported while the Go canary is planted.
				// The Go recipe normally reports the parent of a subtest too.
				hook := string(mustRead(t, filepath.Join(f.root, ".fireweed/eval.sh")))
				injection := fmt.Sprintf(`if test -f "$2/%s"; then
  printf '<testsuite><testcase name="%s"/></testsuite>' > "$3"
  exit 0
fi
`, canaryPath, tc.reported)
				f.put(".fireweed/eval.sh", strings.Replace(hook, "#!/bin/sh\n", "#!/bin/sh\n"+injection, 1))
			}
			const recipePath = ".fireweed/canaries/canaries.yaml"
			var recipe map[string]any
			if err := yaml.Unmarshal(mustRead(t, filepath.Join(f.root, recipePath)), &recipe); err != nil {
				t.Fatal(err)
			}
			recipe["missing_id"] = map[string]string{canaryPath: canarySource}
			data, err := yaml.Marshal(recipe)
			if err != nil {
				t.Fatal(err)
			}
			f.put(recipePath, string(data))
			f.commit()

			report := checks.Doctor(context.Background(), f.request(true))
			const probe = "missing_id global eval"
			if tc.detected {
				if doctorErrors(report) || !hasDoctorCheck(report, probe, true) {
					t.Errorf("Doctor missed the supplied Go canary: %s", doctorText(report))
				}
			} else if hasDoctorCheck(report, probe, true) || !hasDoctorDiagnostic(report, "C14", false, "missing_id", "eval") {
				t.Errorf("Doctor accepted unrelated testcase %q: %s", tc.reported, doctorText(report))
			}
			for _, check := range report.Checks {
				if check.Passed && strings.HasPrefix(check.Name, "recovery ") && strings.Contains(check.Name+" "+check.Detail, "missing_id") {
					return
				}
			}
			t.Errorf("Doctor did not report passing missing_id recovery: %s", doctorText(report))
		})
	}
}
