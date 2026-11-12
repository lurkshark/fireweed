package evals_test

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
)

func requireDiagnosticError(t *testing.T, err error, code, cause string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected %s error containing %q", code, cause)
	}
	if got, _, ok := contracts.ErrorCode(err); !ok || got != code {
		t.Errorf("contracts.ErrorCode(%q) = %q, present %t; want %s", err, got, ok, code)
	}
	if !strings.Contains(err.Error(), cause) {
		t.Errorf("error = %q; want cause %q", err, cause)
	}
}

func requireFailFirstError(t *testing.T, err error, id string) {
	t.Helper()
	const prefix = "P2-01: "
	cause := "no fail-first case for " + id
	if err == nil || !strings.HasPrefix(err.Error(), prefix) || !strings.Contains(err.Error(), cause) ||
		strings.HasPrefix(strings.TrimPrefix(err.Error(), prefix), "P2-01:") {
		t.Errorf("error = %v; want P2-01 prefix and cause %q", err, cause)
	}
}

func checkSpecs() []contracts.Item {
	return []contracts.Item{{ID: "CHK-001"}, {ID: "CHK-002"}, {ID: "CHK-003"}}
}

func Test_CHK_001_CHK_029_CHK_030_NestedJUnitIDsResultsAndTime(t *testing.T) {
	data := []byte(`<testsuites><testsuite name="outer"><testsuite name="inner">
<testcase classname="suite.ChK_001" name="passes" time="1.25"/>
<testcase classname="suite" name="test_chk002_failed" time="0.5"><failure message="bad"/></testcase>
<testcase classname="suite" name="test_chk_003_error" time="0.75"><error message="broken"/></testcase>
<testcase classname="CHK-002" name="CHK-001_both" time="0.125"/>
</testsuite></testsuite></testsuites>`)
	report, err := checks.ParseJUnit(data, checkSpecs())
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]checks.Case{
		"passes":             {Name: "passes", Classname: "suite.ChK_001", IDs: []string{"CHK-001"}, Result: "pass", Seconds: 1.25},
		"test_chk002_failed": {Name: "test_chk002_failed", Classname: "suite", IDs: []string{"CHK-002"}, Result: "fail", Seconds: 0.5},
		"test_chk_003_error": {Name: "test_chk_003_error", Classname: "suite", IDs: []string{"CHK-003"}, Result: "fail", Seconds: 0.75},
		"CHK-001_both":       {Name: "CHK-001_both", Classname: "CHK-002", IDs: []string{"CHK-001", "CHK-002"}, Result: "pass", Seconds: 0.125},
	}
	if len(report.Cases) != len(want) {
		t.Fatalf("cases = %#v; want %d", report.Cases, len(want))
	}
	for _, got := range report.Cases {
		sort.Strings(got.IDs)
		got.Failure = ""
		if !reflect.DeepEqual(got, want[got.Name]) {
			t.Errorf("case = %#v; want %#v", got, want[got.Name])
		}
	}
}

func Test_CHK_001_CHK_029_CHK_030_SingleSuiteSkipped(t *testing.T) {
	report, err := checks.ParseJUnit([]byte(`<testsuite><testcase name="CHK-001" classname="unit" time="0.125"><skipped/></testcase></testsuite>`), checkSpecs())
	if err != nil {
		t.Fatal(err)
	}
	if len(report.Cases) != 1 || report.Cases[0].Result != "skip" || report.Cases[0].Seconds != 0.125 || !reflect.DeepEqual(report.Cases[0].IDs, []string{"CHK-001"}) {
		t.Fatalf("skipped case = %#v", report.Cases)
	}
}

func Test_CHK_079_ParseJUnitRetainsFailureAndErrorDetails(t *testing.T) {
	data := []byte(`<testsuite>
<testcase name="CHK-001_failure_combined"><failure message="want 2 &amp; got 3">assertion at example_test.go:17</failure></testcase>
<testcase name="CHK-001_failure_message"><failure message="expected true"/></testcase>
<testcase name="CHK-002_failure_body"><failure>actual value was false</failure></testcase>
<testcase name="CHK-002_error_combined"><error message="panic: nil pointer">stack at example_test.go:29</error></testcase>
<testcase name="CHK-003_error_message"><error message="setup failed"/></testcase>
<testcase name="CHK-003_error_body"><error>missing fixture</error></testcase>
<testcase name="CHK-001_empty_failure"><failure/></testcase>
<testcase name="CHK-002_pass"/>
<testcase name="CHK-003_skip"><skipped/></testcase>
</testsuite>`)
	report, err := checks.ParseJUnit(data, checkSpecs())
	if err != nil {
		t.Fatal(err)
	}
	want := map[string][]string{
		"CHK-001_failure_combined": {"want 2 & got 3", "assertion at example_test.go:17"},
		"CHK-001_failure_message":  {"expected true"},
		"CHK-002_failure_body":     {"actual value was false"},
		"CHK-002_error_combined":   {"panic: nil pointer", "stack at example_test.go:29"},
		"CHK-003_error_message":    {"setup failed"},
		"CHK-003_error_body":       {"missing fixture"},
		"CHK-001_empty_failure":    nil,
		"CHK-002_pass":             nil,
		"CHK-003_skip":             nil,
	}
	if len(report.Cases) != len(want) {
		t.Fatalf("cases = %#v; want %d", report.Cases, len(want))
	}
	for _, got := range report.Cases {
		parts, ok := want[got.Name]
		if !ok {
			t.Errorf("unexpected case %#v", got)
			continue
		}
		if len(parts) == 0 && got.Failure != "" {
			t.Errorf("%s: failure = %q; want empty", got.Name, got.Failure)
		}
		for _, part := range parts {
			if !strings.Contains(got.Failure, part) {
				t.Errorf("%s: failure = %q; missing %q", got.Name, got.Failure, part)
			}
		}
	}
}

func Test_CHK_002_SkippedJUnitCaseDoesNotCover(t *testing.T) {
	report, err := checks.ParseJUnit([]byte(`<testsuite><testcase name="CHK-001" classname="unit" time="0.125"><skipped/></testcase></testsuite>`), checkSpecs())
	if err != nil {
		t.Fatal(err)
	}
	requireDiagnosticError(t, checks.RequireCoverage(report, []contracts.Item{{ID: "CHK-001"}}), "C10", "no coverage for CHK-001")
}

func Test_CHK_027_RejectMalformedXML(t *testing.T) {
	for _, tc := range []struct{ xml, cause string }{
		{`<testsuite><testcase name="CHK-001"></testsuite>`, "parse JUnit:"},
		{`<document/>`, `invalid JUnit root "document"`},
	} {
		_, err := checks.ParseJUnit([]byte(tc.xml), checkSpecs())
		requireDiagnosticError(t, err, "C11", tc.cause)
	}
}

func Test_CHK_028_RejectMissingAndUnknownIDs(t *testing.T) {
	for _, xml := range []string{
		`<testsuite><testcase name="ordinary" classname="unit"/></testsuite>`,
		`<testsuite><testcase name="ordinary" classname="unit"><skipped/></testcase></testsuite>`,
		`<testsuite><testcase name="case_00" classname="unit"/></testsuite>`,
		`<testsuite><testcase name="CHK-999" classname="unit"/></testsuite>`,
		`<testsuite><testcase name="OCHK001" classname="unit"/></testsuite>`,
		`<testsuite><testcase name="CHK0010" classname="unit"/></testsuite>`,
	} {
		_, err := checks.ParseJUnit([]byte(xml), checkSpecs())
		cause := "has no known spec ID"
		if strings.Contains(xml, "CHK-999") {
			cause = `references unknown spec ID "CHK-999"`
		} else if strings.Contains(xml, "CHK0010") {
			cause = `references unknown spec ID "CHK0010"`
		}
		requireDiagnosticError(t, err, "C11", cause)
	}
}

func Test_CHK_028_ParseJUnitAcceptsNumberedLabels(t *testing.T) {
	const name = "ALP-001_case_00"
	const message = "expected status 200"
	const body = "got status 500\nassertion at example_test.go:17"
	specs := []contracts.Item{{ID: "ALP-001"}}
	for _, tc := range []struct {
		result, diagnostic string
	}{
		{result: "pass"},
		{result: "fail", diagnostic: `<failure message="` + message + `">` + body + `</failure>`},
	} {
		t.Run(tc.result, func(t *testing.T) {
			data := []byte(`<testsuite><testcase name="` + name + `">` + tc.diagnostic + `</testcase></testsuite>`)
			report, err := checks.ParseJUnit(data, specs)
			if err != nil {
				t.Fatal(err)
			}
			if len(report.Cases) != 1 {
				t.Fatalf("cases = %#v; want one", report.Cases)
			}
			got := report.Cases[0]
			if !reflect.DeepEqual(got.IDs, []string{"ALP-001"}) || got.Result != tc.result {
				t.Errorf("case = %#v; want ID ALP-001 and result %s", got, tc.result)
			}
			if tc.result == "fail" && (!strings.Contains(got.Failure, message) || !strings.Contains(got.Failure, body)) {
				t.Errorf("failure = %q; want complete message %q and body %q", got.Failure, message, body)
			}
		})
	}
}

func Test_CHK_028_ParseJUnitRejectsUnknownIDAlongsideKnownID(t *testing.T) {
	data := []byte(`<testsuite><testcase name="ALP-001_ALP-009"/></testsuite>`)
	_, err := checks.ParseJUnit(data, []contracts.Item{{ID: "ALP-001"}})
	requireDiagnosticError(t, err, "C11", `references unknown spec ID "ALP-009"`)
}

func Test_CHK_002_CoverageRequiresCurrentNonSkippedCases(t *testing.T) {
	specs := []contracts.Item{{ID: "CHK-001"}, {ID: "CHK-002"}}
	requireDiagnosticError(t, checks.RequireCoverage(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001"}, Result: "pass"}}}, specs), "C10", "no coverage for CHK-002")
	if err := checks.RequireCoverage(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001", "CHK-002"}, Result: "fail"}}}, specs); err != nil {
		t.Errorf("failed referencing case should cover both IDs: %v", err)
	}
}

func Test_CHK_002_CoverageExemptsReviewAndQuarantine(t *testing.T) {
	specs := []contracts.Item{{ID: "CHK-001"}, {ID: "CHK-002", Tag: "review"}, {ID: "CHK-003", Tag: "quarantine: flaky"}}
	if err := checks.RequireCoverage(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001"}, Result: "pass"}}}, specs); err != nil {
		t.Errorf("review/quarantine required coverage: %v", err)
	}
}

func Test_CHK_025_FailFirstRequiresEachRequestedFailure(t *testing.T) {
	ids := []string{"CHK-001", "CHK-002"}
	bad := []checks.Case{
		{IDs: []string{"CHK-001"}, Result: "pass"},
		{IDs: []string{"CHK-002"}, Result: "skip"},
		{IDs: []string{"CHK-003"}, Result: "fail"},
	}
	requireFailFirstError(t, checks.RequireFailFirst(checks.Report{Cases: bad}, ids, nil), "CHK-001")
	requireFailFirstError(t, checks.RequireFailFirst(checks.Report{}, ids, nil), "CHK-001")
	requireFailFirstError(t, checks.RequireFailFirst(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001"}, Result: "fail"}, {IDs: []string{"CHK-002"}, Result: "skip"}}}, ids, nil), "CHK-002")
	good := checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001"}, Result: "fail"}, {IDs: []string{"CHK-002"}, Result: "fail"}}}
	if err := checks.RequireFailFirst(good, ids, nil); err != nil {
		t.Errorf("both requested IDs failed: %v", err)
	}
}

func Test_CHK_025_FailFirstAlreadySatisfiedExemption(t *testing.T) {
	report := checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-002"}, Result: "fail"}}}
	if err := checks.RequireFailFirst(report, []string{"CHK-001", "CHK-002"}, []string{"CHK-001"}); err != nil {
		t.Errorf("already_satisfied ID incorrectly required: %v", err)
	}
}

func Test_CHK_026_PassRejectsNonQuarantinedFailures(t *testing.T) {
	specs := []contracts.Item{{ID: "CHK-001"}, {ID: "CHK-002", Tag: "review"}, {ID: "CHK-003", Tag: "quarantine: flaky"}}
	for _, id := range []string{"CHK-001", "CHK-002"} {
		requireDiagnosticError(t, checks.RequirePass(checks.Report{Cases: []checks.Case{{IDs: []string{id}, Result: "fail"}}}, specs), "E01", `failed case "" for `+id)
	}
	if err := checks.RequirePass(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-003"}, Result: "fail"}}}, specs); err != nil {
		t.Errorf("quarantined failure rejected: %v", err)
	}
}
