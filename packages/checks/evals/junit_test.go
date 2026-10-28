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
	if !strings.HasPrefix(err.Error(), code+": ") || !strings.Contains(err.Error(), cause) || strings.HasPrefix(strings.TrimPrefix(err.Error(), code+": "), code+":") {
		t.Errorf("error = %q; want %s prefix and cause %q", err, code, cause)
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
	requireDiagnosticError(t, checks.RequireFailFirst(checks.Report{Cases: bad}, ids, nil), "P2-01", "no fail-first case for CHK-001")
	requireDiagnosticError(t, checks.RequireFailFirst(checks.Report{}, ids, nil), "P2-01", "no fail-first case for CHK-001")
	requireDiagnosticError(t, checks.RequireFailFirst(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001"}, Result: "fail"}, {IDs: []string{"CHK-002"}, Result: "skip"}}}, ids, nil), "P2-01", "no fail-first case for CHK-002")
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
