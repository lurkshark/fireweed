package evals

import (
	"encoding/json"
	"reflect"
	"sort"
	"strings"
	"testing"

	k "github.com/lurkshark/fireweed/packages/checks"
	c "github.com/lurkshark/fireweed/packages/contracts"
)

func Test_CHK001JUnitCasesAndMetadata(t *testing.T) {
	one := `<testcase name="Test_ALP001 &amp; more" classname="pkg.alpha" spec_ids="ALP-999"><failure message="first &amp; reason">line1
line2</failure><failure/><error message="panic"><![CDATA[trace <stack>]]></error><error/><skipped message="reason">skip text</skipped><properties><property name="spec_ids" value="ALP-999"/></properties></testcase>`
	want := []k.TestCase{{Name: "Test_ALP001 & more", Classname: "pkg.alpha", Failures: []k.Message{{Message: "first & reason", Text: "line1\nline2"}, {}}, Errors: []k.Message{{Message: "panic", Text: "trace <stack>"}, {}}, Skipped: &k.Message{Message: "reason", Text: "skip text"}, IDs: []string{}}, {Name: "second", Classname: "ALP002", IDs: []string{}}, {Name: "third", IDs: []string{}}}
	for _, xml := range []string{
		`<testsuite tests="999" failures="0"><properties/><system-out>ALP999</system-out>` + one + `<testsuite><testcase name="second" classname="ALP002"/></testsuite><testcase name="third"/></testsuite>`,
		`<?xml version="1.0"?><!-- report --><testsuites tests="0"><testsuite>` + one + `</testsuite><testsuites><testsuite><testcase name="second" classname="ALP002"/></testsuite></testsuites><testsuite><testcase name="third"/></testsuite></testsuites>`,
	} {
		raw := []byte(xml)
		before := append([]byte(nil), raw...)
		got, err := k.ParseJUnit(raw)
		if err != nil || !reflect.DeepEqual(got, want) {
			t.Fatalf("case evidence %+v %v want %+v", got, err, want)
		}
		if !reflect.DeepEqual(raw, before) {
			t.Fatal("parser changed bytes")
		}
	}
	for _, xml := range []string{`<testsuite/>`, `<testsuites tests="400" failures="99"><testsuite/></testsuites>`, " \n<testsuite></testsuite>\n "} {
		got, err := k.ParseJUnit([]byte(xml))
		if err != nil || got == nil || len(got) != 0 {
			t.Fatalf("valid empty %+v %v", got, err)
		}
	}
}
func Test_CHK001InvalidReportsAreAtomic(t *testing.T) {
	for _, xml := range []string{"", " \n", `<wrong/>`, `<testcase name="ALP001"/>`, `<testsuite>`, `<testsuite><testcase name="ALP001"/><broken></testsuite>`, `<testsuite/><testsuite/>`, `before<testsuite/>`, `<testsuite/>after`, `<testsuite><testcase name="ALP001"/></testsuite><`, `<testsuite>&undefined;</testsuite>`, `<testsuite><testcase name="ALP001"></testsuite>`, `<testsuites><testsuite/></testsuites>garbage`} {
		got, err := k.ParseJUnit([]byte(xml))
		if err == nil || len(got) != 0 {
			t.Errorf("invalid report accepted or partial: %q %+v %v", xml, got, err)
		}
	}
}
func evidenceModule(ids ...string) c.Module {
	m := c.Module{Name: "alpha", IDPrefix: "ALP", Dir: "units/alpha", Specs: []c.Spec{}}
	for _, id := range ids {
		m.Specs = append(m.Specs, c.Spec{ID: id, Body: "It MUST hold."})
	}
	return m
}
func evidenceCodes(t *testing.T, cov k.Coverage) []string {
	t.Helper()
	codes := []string{}
	for _, d := range cov.Diagnostics {
		if d.Severity != "error" || d.Path != "units/alpha" || strings.TrimSpace(d.Message) == "" {
			t.Fatalf("malformed diagnostic %+v", d)
		}
		lower := strings.ToLower(d.Message)
		corrective := false
		for _, word := range []string{"add", "remove", "name", "reference", "fix", "repair", "must", "run", "require", "produce", "report", "eval"} {
			corrective = corrective || strings.Contains(lower, word)
		}
		if !corrective {
			t.Fatalf("diagnostic lacks useful correction %+v", d)
		}
		codes = append(codes, d.Code)
	}
	sort.Strings(codes)
	return codes
}
func requireCodes(t *testing.T, cov k.Coverage, want ...string) {
	t.Helper()
	sort.Strings(want)
	if want == nil {
		want = []string{}
	}
	if got := evidenceCodes(t, cov); !reflect.DeepEqual(got, want) {
		t.Fatalf("diagnostics %+v want %v", cov.Diagnostics, want)
	}
}
func Test_CHK002IDJoiningAndForeignReferences(t *testing.T) {
	m := evidenceModule("ALP-001", "ALP-002")
	cases := []k.TestCase{
		{Name: "Test_alp_002_then_ALP001_and_alp-001", Classname: "suite/ALP002", IDs: []string{"ALP-999", "BET-001"}},
		{Name: "ALP001After", Classname: "pkg_ALP_002", Skipped: &k.Message{}},
		{Name: "ALP-999", Classname: "BET001 ALP-999"},
		{Name: "no reference", Classname: "unrelated"},
	}
	cov := k.Analyze(m, []string{"BET", "ALP", "BET"}, cases, 0, nil, nil)
	wantIDs := [][]string{{"ALP-001", "ALP-002"}, {"ALP-001", "ALP-002"}, {"ALP-999", "BET-001"}, {}}
	if len(cov.Tests) != len(cases) {
		t.Fatalf("cases %+v", cov)
	}
	for i, want := range wantIDs {
		if !reflect.DeepEqual(cov.Tests[i].IDs, want) {
			t.Fatalf("case%d IDs %v want %v", i, cov.Tests[i].IDs, want)
		}
	}
	wantMap := map[string][]int{"ALP-001": {0, 1}, "ALP-002": {0, 1}, "ALP-999": {2}, "BET-001": {2}}
	if !reflect.DeepEqual(cov.BySpec, wantMap) {
		t.Fatalf("mapping %+v", cov.BySpec)
	}
	requireCodes(t, cov, "C12", "C12", "C12")
	for _, needle := range []string{"ALP-999", "BET-001", "no reference"} {
		found := false
		for _, d := range cov.Diagnostics {
			found = found || strings.Contains(d.Message, needle)
		}
		if !found {
			t.Fatalf("missing %q attribution %+v", needle, cov.Diagnostics)
		}
	}
	for _, prefixes := range [][]string{nil, {}, {"BET"}} {
		one := k.Analyze(evidenceModule("ALP-001"), prefixes, []k.TestCase{{Name: "alp001"}}, 0, nil, nil)
		requireCodes(t, one)
		if !reflect.DeepEqual(one.BySpec, map[string][]int{"ALP-001": {0}}) {
			t.Fatalf("own prefix fallback %+v", one)
		}
	}
}
func Test_CHK002ExactPublicMatching(t *testing.T) {
	m := evidenceModule("ALP-001", "ALP-0015", "ALP-002")
	for i := range m.Specs {
		m.Specs[i].Tag = "[review]"
	}
	for _, text := range []string{"ALP-001", "alp_001", "ALP001", "x/AlP001Body", "ALP001Z", "(ALP-001)", "ALP0015 ALP001", "ALP0015", "OALP001", "1ALP001", "éALP001", "ALP001２", "ALP_01", "ALP--001", "ALP 001", "ALP-999", "BET001", "ALP001 ALP002", ""} {
		for _, classname := range []bool{false, true} {
			tc := k.TestCase{Name: "unrelated"}
			if classname {
				tc.Classname = text
			} else {
				tc.Name = text
			}
			cov := k.Analyze(m, nil, []k.TestCase{tc}, 0, nil, nil)
			if len(cov.Tests) != 1 {
				t.Fatal(cov)
			}
			want := []string{}
			for _, id := range []string{"ALP-001", "ALP-0015", "ALP-002", "ALP-999"} {
				if c.MatchesID(tc.Name, tc.Classname, id) {
					want = append(want, id)
				}
			}
			sort.Strings(want)
			if !reflect.DeepEqual(cov.Tests[0].IDs, want) {
				t.Fatalf("%q class%v IDs%v want%v", text, classname, cov.Tests[0].IDs, want)
			}
			wantCodes := []string{}
			if len(want) == 0 || strings.Contains(text, "999") {
				wantCodes = append(wantCodes, "C12")
			}
			requireCodes(t, cov, wantCodes...)
		}
	}
}
func Test_CHK003CoverageAndFailFirst(t *testing.T) {
	fail := k.TestCase{Name: "ALP001", Failures: []k.Message{{}}}
	errCase := k.TestCase{Name: "ALP001", Errors: []k.Message{{}}}
	pass := k.TestCase{Name: "ALP001"}
	skip := k.TestCase{Name: "ALP001", Skipped: &k.Message{}}
	mixedSkip := k.TestCase{Name: "ALP001", Failures: []k.Message{{}}, Errors: []k.Message{{}}, Skipped: &k.Message{}}
	type policy struct {
		name           string
		tests          []k.TestCase
		exit           int
		first, already []string
		codes          []string
	}
	for _, tt := range []policy{
		{name: "missing", codes: []string{"C13"}},
		{name: "pass", tests: []k.TestCase{pass}},
		{name: "skipped", tests: []k.TestCase{skip}, codes: []string{"C13"}},
		{name: "ordinary failure", tests: []k.TestCase{fail}, exit: 1, codes: []string{"C14"}},
		{name: "ordinary error", tests: []k.TestCase{errCase}, exit: 1, codes: []string{"C14"}},
		{name: "failfirst failure", tests: []k.TestCase{fail}, exit: 1, first: []string{"ALP-001"}},
		{name: "failfirst error", tests: []k.TestCase{errCase}, exit: 1, first: []string{"ALP-001"}},
		{name: "failfirst passing", tests: []k.TestCase{pass}, first: []string{"ALP-001"}, codes: []string{"P1-02"}},
		{name: "missing failfirst", first: []string{"ALP-001"}, codes: []string{"C13", "P1-02"}},
		{name: "already with genuine failfirst", tests: []k.TestCase{fail}, exit: 1, first: []string{"ALP-001"}, already: []string{"ALP-001"}},
		{name: "allowed passing", tests: []k.TestCase{pass}, first: []string{"ALP-001"}, already: []string{"ALP-001"}},
		{name: "allowed missing", already: []string{"ALP-001"}, first: []string{"ALP-001"}, codes: []string{"C13"}},
		{name: "allowed skipped", tests: []k.TestCase{skip}, already: []string{"ALP-001"}, first: []string{"ALP-001"}, codes: []string{"C13"}},
		{name: "allowed ordinary failure", tests: []k.TestCase{fail}, exit: 1, already: []string{"ALP-001"}, codes: []string{"C14"}},
		{name: "fake hook failure", tests: []k.TestCase{pass}, exit: 42, first: []string{"ALP-001"}, codes: []string{"C11", "P1-02"}},
		{name: "fake hook success", tests: []k.TestCase{fail}, first: []string{"ALP-001"}, codes: []string{"C11"}},
		{name: "skip beats failure", tests: []k.TestCase{mixedSkip}, exit: 1, first: []string{"ALP-001"}, codes: []string{"C11", "C13", "P1-02"}},
		{name: "skip fake success", tests: []k.TestCase{mixedSkip}, first: []string{"ALP-001"}, codes: []string{"C13", "P1-02"}},
		{name: "pass and failure", tests: []k.TestCase{pass, fail, skip}, exit: 1, first: []string{"ALP-001"}},
		{name: "duplicates", tests: []k.TestCase{pass}, first: []string{"ALP-001", "ALP-001"}, codes: []string{"P1-02"}},
		{name: "unknown wanted", tests: []k.TestCase{pass}, first: []string{"ALP-999", "ALP-999"}, codes: []string{"C12"}},
		{name: "unknown already", tests: []k.TestCase{pass}, already: []string{"BET-001", "BET-001"}, codes: []string{"C12"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			cov := k.Analyze(evidenceModule("ALP-001"), []string{"BET"}, tt.tests, tt.exit, tt.first, tt.already)
			requireCodes(t, cov, tt.codes...)
		})
	}
	m := evidenceModule("ALP-001", "ALP-002")
	unrelated := k.Analyze(m, nil, []k.TestCase{pass, {Name: "ALP002", Errors: []k.Message{{Text: "unrelated"}}}}, 1, []string{"ALP-001"}, nil)
	requireCodes(t, unrelated, "C14", "P1-02")
	for _, d := range unrelated.Diagnostics {
		if d.Code == "C14" && !strings.Contains(d.Message, "ALP-002") || d.Code == "P1-02" && !strings.Contains(d.Message, "ALP-001") {
			t.Fatalf("misattributed policy %+v", d)
		}
	}
}
func Test_CHK003TagsAndReporterConsistency(t *testing.T) {
	for _, tag := range []string{"[review]", "[quarantine: flaky clock]", "[review] [quarantine: temporary]"} {
		m := evidenceModule("ALP-001")
		m.Specs[0].Tag = tag
		cov := k.Analyze(m, nil, nil, 0, []string{"ALP-001"}, nil)
		requireCodes(t, cov)
		requireCodes(t, k.Analyze(m, nil, []k.TestCase{{Name: "ALP001"}}, 0, []string{"ALP-001"}, nil))
		requireCodes(t, k.Analyze(m, nil, []k.TestCase{{Name: "ALP001", Failures: []k.Message{{}}}}, 1, nil, nil))
		cases := []k.TestCase{{Name: "ALP001", Errors: []k.Message{{}}}, {Name: "ALP001", Skipped: &k.Message{}}}
		cov = k.Analyze(m, nil, cases, 1, []string{"ALP-001"}, nil)
		requireCodes(t, cov)
		if len(cov.Tests) != 2 || !reflect.DeepEqual(cov.BySpec, map[string][]int{"ALP-001": {0, 1}}) || !reflect.DeepEqual(cov.Tests[0].IDs, []string{"ALP-001"}) {
			t.Fatalf("tag visibility %+v", cov)
		}
		requireCodes(t, k.Analyze(m, nil, cases, 0, []string{"ALP-001"}, nil), "C11")
		requireCodes(t, k.Analyze(m, nil, []k.TestCase{{Name: "ALP001"}}, 1, nil, nil), "C11")
	}
	// Summary counts cannot create executed evidence or erase testcase errors.
	for _, xml := range []string{`<testsuite tests="10" failures="10"/>`, `<testsuite failures="7"><testcase name="ALP001"><skipped/></testcase></testsuite>`} {
		cases, err := k.ParseJUnit([]byte(xml))
		if err != nil {
			t.Fatal(err)
		}
		requireCodes(t, k.Analyze(evidenceModule("ALP-001"), nil, cases, 1, []string{"ALP-001"}, nil), "C11", "C13", "P1-02")
	}
}
func jsonSnapshot(t *testing.T, v any) string {
	t.Helper()
	b, e := json.Marshal(v)
	if e != nil {
		t.Fatal(e)
	}
	return string(b)
}
func Test_CHK004OwnershipDeterminismAndEmptyValues(t *testing.T) {
	m := evidenceModule("ALP-002", "ALP-001")
	m.DependsOn = []string{"beta"}
	m.Document = c.Document{Metadata: map[string]any{"nested": []any{"retained"}}, Sections: []c.Section{{Name: "Specs", Body: "original"}}}
	prefixes := []string{"BET", "ALP", "BET"}
	first := []string{"ALP-002", "ALP-001", "ALP-002"}
	already := []string{"ALP-002", "ALP-001", "ALP-002"}
	tests := []k.TestCase{{Name: "ALP001", Failures: []k.Message{{Message: "failure", Text: "trace"}}, Errors: []k.Message{{Message: "error", Text: "stack"}}, IDs: []string{"untrusted"}}, {Name: "ALP002", Skipped: &k.Message{Message: "skip", Text: "why"}, IDs: []string{"wrong"}}, {Name: "no ID"}}
	before := jsonSnapshot(t, []any{m, prefixes, tests, first, already})
	a := k.Analyze(m, prefixes, tests, 1, first, already)
	b := k.Analyze(m, prefixes, tests, 1, first, already)
	if !reflect.DeepEqual(a, b) || a.Module != "alpha" || a.HookExit != 1 || len(a.Tests) != 3 || a.Tests == nil || a.BySpec == nil || a.Diagnostics == nil {
		t.Fatalf("deterministic owned result %+v %+v", a, b)
	}
	if a.Tests[0].Failures[0] != tests[0].Failures[0] || a.Tests[0].Errors[0] != tests[0].Errors[0] || *a.Tests[1].Skipped != *tests[1].Skipped || a.Tests[0].Name != tests[0].Name || a.Tests[0].Classname != tests[0].Classname {
		t.Fatal("lost testcase evidence")
	}
	if before != jsonSnapshot(t, []any{m, prefixes, tests, first, already}) {
		t.Fatal("inputs changed during analysis")
	}
	if !sort.SliceIsSorted(a.Diagnostics, func(i, j int) bool {
		x, y := a.Diagnostics[i], a.Diagnostics[j]
		if x.Path != y.Path {
			return x.Path < y.Path
		}
		if x.Code != y.Code {
			return x.Code < y.Code
		}
		if x.Severity != y.Severity {
			return x.Severity < y.Severity
		}
		return x.Message < y.Message
	}) {
		t.Fatalf("diagnostics unsorted %+v", a.Diagnostics)
	}
	saved := jsonSnapshot(t, b)
	a.Tests[0].Name = "changed"
	a.Tests[0].Failures[0].Message = "changed"
	a.Tests[0].Failures[0].Text = "changed"
	a.Tests[0].Errors[0].Text = "changed"
	a.Tests[0].Errors[0].Message = "changed"
	a.Tests[0].IDs[0] = "changed"
	a.Tests[1].Skipped.Text = "changed"
	a.Tests[1].Skipped.Message = "changed"
	a.BySpec["ALP-001"][0] = 99
	a.Diagnostics[0].Message = "changed"
	if before != jsonSnapshot(t, []any{m, prefixes, tests, first, already}) || saved != jsonSnapshot(t, b) {
		t.Fatal("returned result aliases input or another result")
	}
	// Input mutation must not retroactively rewrite an already produced result.
	tests[0].Failures[0].Text = "input changed"
	tests[0].Errors[0].Message = "input changed"
	tests[0].IDs[0] = "input changed"
	tests[1].Skipped.Message = "input changed"
	m.Specs[0].ID = "changed"
	prefixes[0] = "changed"
	first[0] = "changed"
	already[0] = "changed"
	if saved != jsonSnapshot(t, b) {
		t.Fatal("input mutation changed result")
	}
	for _, tests := range [][]k.TestCase{nil, {}} {
		empty := k.Analyze(evidenceModule(), nil, tests, 0, nil, nil)
		if empty.Tests == nil || empty.BySpec == nil || empty.Diagnostics == nil || len(empty.Tests)+len(empty.BySpec)+len(empty.Diagnostics) != 0 || empty.Module != "alpha" || empty.HookExit != 0 {
			t.Fatalf("empty result %+v", empty)
		}
	}
	// Empty slices with hidden backing data must not be exposed through a result.
	backingMessages := []k.Message{{Text: "hidden"}}
	backingIDs := []string{"hidden"}
	backingCases := []k.TestCase{{Name: "hidden"}}
	tc := k.TestCase{Name: "ALP001", Failures: backingMessages[:0], Errors: backingMessages[:0], IDs: backingIDs[:0]}
	owned := k.Analyze(evidenceModule("ALP-001"), nil, []k.TestCase{tc}, 0, nil, nil)
	owned.Tests[0].Failures = append(owned.Tests[0].Failures, k.Message{Text: "result"})
	owned.Tests[0].Errors = append(owned.Tests[0].Errors, k.Message{Text: "result"})
	if backingMessages[0].Text != "hidden" || backingIDs[0] != "hidden" {
		t.Fatal("empty slice backing aliased")
	}
	empty := k.Analyze(evidenceModule(), nil, backingCases[:0], 0, nil, nil)
	empty.Tests = append(empty.Tests, k.TestCase{Name: "new"})
	if backingCases[0].Name != "hidden" {
		t.Fatal("empty case backing aliased")
	}
}
func Test_CHK004ParserCallsAreIndependent(t *testing.T) {
	xml := []byte(`<testsuite><testcase name="ALP001" classname="pkg"><failure message="first">text</failure><error message="error">trace</error><skipped message="skip">why</skipped></testcase></testsuite>`)
	a, err := k.ParseJUnit(xml)
	if err != nil {
		t.Fatal(err)
	}
	b, err := k.ParseJUnit(xml)
	if err != nil {
		t.Fatal(err)
	}
	if len(a) != 1 || len(b) != 1 || len(a[0].Failures) != 1 || len(a[0].Errors) != 1 || a[0].Skipped == nil || a[0].IDs == nil {
		t.Fatalf("parse %+v %+v", a, b)
	}
	saved := jsonSnapshot(t, b)
	raw := string(xml)
	a[0].Name = "changed"
	a[0].Classname = "changed"
	a[0].Failures[0].Message = "changed"
	a[0].Failures[0].Text = "changed"
	a[0].Errors[0].Message = "changed"
	a[0].Errors[0].Text = "changed"
	a[0].Skipped.Message = "changed"
	a[0].Skipped.Text = "changed"
	a[0].IDs = append(a[0].IDs, "changed")
	b[0].IDs = append(b[0].IDs, "other result")
	if a[0].IDs[0] != "changed" {
		t.Fatal("empty parser IDs shared between results")
	}
	b[0].IDs = b[0].IDs[:0]
	if saved != jsonSnapshot(t, b) || raw != string(xml) {
		t.Fatal("parser results alias")
	}
	for i := range xml {
		xml[i] = ' '
	}
	if saved != jsonSnapshot(t, b) {
		t.Fatal("parsed strings alias bytes")
	}
}

func Test_CHK003MultipleCoverageObligations(t *testing.T) {
	cov := k.Analyze(evidenceModule("ALP-003", "ALP-001", "ALP-002"), nil, []k.TestCase{{Name: "ALP001", Skipped: &k.Message{}}}, 0, []string{"ALP-003", "ALP-001"}, nil)
	requireCodes(t, cov, "C13", "C13", "C13", "P1-02", "P1-02")
	for _, id := range []string{"ALP-001", "ALP-002", "ALP-003"} {
		count := 0
		for _, d := range cov.Diagnostics {
			if d.Code == "C13" && strings.Contains(d.Message, id) {
				count++
			}
		}
		if count != 1 {
			t.Fatalf("coverage for %s count%d: %+v", id, count, cov.Diagnostics)
		}
	}
}
func Test_CHK004SameCodeOrderingAndEmptyStorage(t *testing.T) {
	// Distinct messages under one code require the last comparison key.
	cases := []k.TestCase{}
	for _, name := range []string{"zeta", "omega", "nu", "mu", "lambda", "kappa", "iota", "gamma", "eta", "epsilon", "delta", "alpha"} {
		cases = append(cases, k.TestCase{Name: name})
	}
	cov := k.Analyze(evidenceModule(), nil, cases, 0, nil, nil)
	if len(cov.Diagnostics) != len(cases) {
		t.Fatalf("missing testcase diagnostics %+v", cov)
	}
	messages := []string{}
	seen := map[string]bool{}
	for _, d := range cov.Diagnostics {
		if d.Code != "C12" || seen[d.Message] {
			t.Fatalf("diagnostic fixture %+v", cov.Diagnostics)
		}
		seen[d.Message] = true
		messages = append(messages, d.Message)
	}
	evidenceCodes(t, cov)
	if !sort.StringsAreSorted(messages) {
		t.Fatalf("same-code messages unsorted %v", messages)
	}
	// Fallback must not append into the caller's empty prefix slice.
	backing := []string{"sentinel prefix"}
	prefixes := backing[:0]
	own := k.Analyze(evidenceModule("ALP-001"), prefixes, []k.TestCase{{Name: "ALP001"}}, 0, nil, nil)
	requireCodes(t, own)
	if backing[0] != "sentinel prefix" || len(prefixes) != 0 {
		t.Fatal("target-prefix fallback changed caller storage")
	}
	// Independently empty results must not share maps, slices or hidden storage.
	a := k.Analyze(evidenceModule(), nil, nil, 0, nil, nil)
	b := k.Analyze(evidenceModule(), nil, nil, 0, nil, nil)
	if a.Tests == nil || b.Tests == nil || a.BySpec == nil || b.BySpec == nil || a.Diagnostics == nil || b.Diagnostics == nil {
		t.Fatal("nil empty result")
	}
	a.Tests = append(a.Tests, k.TestCase{Name: "first"})
	b.Tests = append(b.Tests, k.TestCase{Name: "second"})
	a.Diagnostics = append(a.Diagnostics, k.Diagnostic{Message: "first"})
	b.Diagnostics = append(b.Diagnostics, k.Diagnostic{Message: "second"})
	a.BySpec["first"] = []int{1}
	b.BySpec["second"] = []int{2}
	if a.Tests[0].Name != "first" || a.Diagnostics[0].Message != "first" || len(a.BySpec) != 1 || len(b.BySpec) != 1 {
		t.Fatal("empty results shared storage")
	}
	inputMessages := []k.Message{{Text: "hidden"}}
	inputIDs := []string{"hidden"}
	tc := k.TestCase{Name: "no ID", Failures: inputMessages[:0], Errors: inputMessages[:0], IDs: inputIDs[:0]}
	a = k.Analyze(evidenceModule(), nil, []k.TestCase{tc}, 0, nil, nil)
	b = k.Analyze(evidenceModule(), nil, []k.TestCase{tc}, 0, nil, nil)
	if len(a.Tests) != 1 || len(b.Tests) != 1 || a.Tests[0].IDs == nil || b.Tests[0].IDs == nil {
		t.Fatal("missing empty-ID cases")
	}
	a.Tests[0].Failures = append(a.Tests[0].Failures, k.Message{Text: "failure A"})
	b.Tests[0].Failures = append(b.Tests[0].Failures, k.Message{Text: "failure B"})
	a.Tests[0].Errors = append(a.Tests[0].Errors, k.Message{Text: "error A"})
	b.Tests[0].Errors = append(b.Tests[0].Errors, k.Message{Text: "error B"})
	a.Tests[0].IDs = append(a.Tests[0].IDs, "ID A")
	b.Tests[0].IDs = append(b.Tests[0].IDs, "ID B")
	if a.Tests[0].Failures[0].Text != "failure A" || a.Tests[0].Errors[0].Text != "error A" || a.Tests[0].IDs[0] != "ID A" || inputMessages[0].Text != "hidden" || inputIDs[0] != "hidden" {
		t.Fatal("empty nested result storage shared")
	}
}
