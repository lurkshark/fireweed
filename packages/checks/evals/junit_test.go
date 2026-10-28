package evals_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
)

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

func Test_CHK_001_CHK_002_CHK_029_CHK_030_SingleSuiteSkippedDoesNotCover(t *testing.T) {
	report, err := checks.ParseJUnit([]byte(`<testsuite><testcase name="CHK-001" classname="unit" time="0.125"><skipped/></testcase></testsuite>`), checkSpecs())
	if err != nil {
		t.Fatal(err)
	}
	if len(report.Cases) != 1 || report.Cases[0].Result != "skip" || report.Cases[0].Seconds != 0.125 || !reflect.DeepEqual(report.Cases[0].IDs, []string{"CHK-001"}) {
		t.Fatalf("skipped case = %#v", report.Cases)
	}
	if err := checks.RequireCoverage(report, []contracts.Item{{ID: "CHK-001"}}); err == nil {
		t.Error("skipped testcase counted as coverage")
	}
}

func Test_CHK_027_RejectMalformedXML(t *testing.T) {
	if _, err := checks.ParseJUnit([]byte(`<testsuite><testcase name="CHK-001"></testsuite>`), checkSpecs()); err == nil {
		t.Error("malformed XML accepted")
	}
}

func Test_CHK_028_CHK_029_RejectMissingAndUnknownIDs(t *testing.T) {
	for _, xml := range []string{
		`<testsuite><testcase name="ordinary" classname="unit"/></testsuite>`,
		`<testsuite><testcase name="ordinary" classname="unit"><skipped/></testcase></testsuite>`,
		`<testsuite><testcase name="CHK-999" classname="unit"/></testsuite>`,
		`<testsuite><testcase name="OCHK001" classname="unit"/></testsuite>`,
		`<testsuite><testcase name="CHK0010" classname="unit"/></testsuite>`,
	} {
		if _, err := checks.ParseJUnit([]byte(xml), checkSpecs()); err == nil {
			t.Errorf("ID-less or unknown testcase accepted: %s", xml)
		}
	}
}

func Test_CHK_002_CoverageRequiresCurrentNonSkippedCases(t *testing.T) {
	specs := []contracts.Item{{ID: "CHK-001"}, {ID: "CHK-002"}}
	if err := checks.RequireCoverage(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001"}, Result: "pass"}}}, specs); err == nil {
		t.Error("missing current ID accepted")
	}
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
	if err := checks.RequireFailFirst(checks.Report{Cases: bad}, ids, nil); err == nil {
		t.Error("passing, skipped, and unrelated failure accepted")
	}
	if err := checks.RequireFailFirst(checks.Report{}, ids, nil); err == nil {
		t.Error("empty report accepted")
	}
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
		if err := checks.RequirePass(checks.Report{Cases: []checks.Case{{IDs: []string{id}, Result: "fail"}}}, specs); err == nil {
			t.Errorf("failed non-quarantined %s accepted", id)
		}
	}
	if err := checks.RequirePass(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-003"}, Result: "fail"}}}, specs); err != nil {
		t.Errorf("quarantined failure rejected: %v", err)
	}
}
