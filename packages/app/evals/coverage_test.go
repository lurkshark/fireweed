package evals

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	app "github.com/lurkshark/fireweed/packages/app"
	k "github.com/lurkshark/fireweed/packages/checks"
	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
)

const coverageDir = "units/alpha '雪 space"
const coveragePassing = `<testsuites><testsuite><testcase name="ALP001" classname="eval.alpha"/><testcase name="ALP002"><skipped message="quarantine"/></testcase><testcase name="ALP003"><skipped/></testcase></testsuite></testsuites>`

func coverageQuote(s string) string { return "'" + strings.ReplaceAll(s, "'", `'"'"'`) + "'" }
func coverageHook(label string) string {
	return "exec sh hooks/eval.sh {report} {module} {module_dir} " + coverageQuote(label)
}
func coverageFixture(t *testing.T) string {
	t.Helper()
	root := projectFixture(t)
	projectPut(t, root, "fireweed.yaml", projectConfig+"timeout: 1s\n")
	encoded, _ := json.Marshal(coverageHook("current"))
	projectModule(t, root, coverageDir, "alpha", "ALP", "hooks:\n  eval: "+string(encoded)+"\n")
	raw, err := os.ReadFile(filepath.Join(root, coverageDir, "SPEC.md"))
	if err != nil {
		t.Fatal(err)
	}
	text := strings.Replace(string(raw), "### ALP-001 [review]", "### ALP-001", 1)
	text = strings.Replace(text, "## Decisions", "### ALP-002 [quarantine: clock]\nIt MUST remain visible.\n\n### ALP-003 [review]\nIt MUST be reviewed.\n\n## Decisions", 1)
	projectPut(t, root, coverageDir+"/SPEC.md", text)
	projectModule(t, root, "units/beta", "beta", "BET", "")
	script := `#!/bin/sh
set -eu
mode=$(cat "$FIREWEED_MODULE_DIR/$4.mode")
case "$mode" in
 missing) ;;
 directory) mkdir "$1" ;;
 *) cat "$FIREWEED_MODULE_DIR/$4.xml" > "$1" ;;
esac
printf 'eval-%s-out\n' "$4"
printf 'eval-%s-err\n' "$4" >&2
printf '%s\n' "$1" "$FIREWEED_REPORT" "$2" "$FIREWEED_MODULE" "$3" "$FIREWEED_MODULE_DIR" "$FIREWEED_BIN" "$4" > "$FIREWEED_MODULE_DIR/trace"
case "$mode" in
 fail) exit 37 ;;
 signal) kill -TERM $$ ;;
 timeout) sleep 2; printf late > "$FIREWEED_MODULE_DIR/LATE" ;;
esac
`
	projectPut(t, root, "hooks/eval.sh", script)
	coverageSet(t, root, "current", "zero", coveragePassing)
	s := coverageSnapshot(t, root)
	if ds := p.Check(s); len(ds) != 0 {
		t.Fatalf("invalid coverage fixture %+v", ds)
	}
	return root
}
func coverageSet(t *testing.T, root, label, mode, xml string) {
	t.Helper()
	projectPut(t, root, coverageDir+"/"+label+".mode", mode)
	projectPut(t, root, coverageDir+"/"+label+".xml", xml)
}
func coverageSnapshot(t *testing.T, root string) p.Snapshot {
	t.Helper()
	s, err := p.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	return s
}
func coverageModule(t *testing.T, s p.Snapshot, name string) c.Module {
	t.Helper()
	for _, m := range s.Modules {
		if m.Name == name {
			return m
		}
	}
	t.Fatalf("missing %s", name)
	return c.Module{}
}
func coverageTrace(t *testing.T, root, label string, cleanup bool) string {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(root, coverageDir, "trace"))
	if err != nil {
		t.Fatal(err)
	}
	parts := strings.Split(strings.TrimSuffix(string(raw), "\n"), "\n")
	if len(parts) != 8 || !filepath.IsAbs(parts[0]) || parts[0] != parts[1] || parts[2] != "alpha" || parts[3] != "alpha" || parts[4] != coverageDir || parts[5] != coverageDir || parts[6] != subject || parts[7] != label {
		t.Fatalf("hook identity %+v expected executable %s", parts, subject)
	}
	if cleanup {
		if _, err := os.Stat(filepath.Dir(parts[0])); !os.IsNotExist(err) {
			t.Fatalf("report directory leaked %s", parts[0])
		}
	}
	return parts[0]
}
func coveragePretty(t *testing.T, output string) {
	t.Helper()
	var formatted bytes.Buffer
	if err := json.Indent(&formatted, []byte(strings.TrimSpace(output)), "", "  "); err != nil {
		t.Fatal(err)
	}
	if output != formatted.String()+"\n" || !strings.Contains(output, "\n  ") {
		t.Fatalf("coverage must be pretty JSON: %q", output)
	}
}
func coverageCorrective(t *testing.T, d k.Diagnostic) {
	t.Helper()
	lower := strings.ToLower(d.Message)
	useful := false
	for _, verb := range []string{"configure", "repair", "correct", "fix", "provide", "ensure"} {
		useful = useful || strings.Contains(lower, verb)
	}
	if d.Severity != "error" || strings.TrimSpace(d.Message) == "" || !useful {
		t.Fatalf("missing corrective error diagnostic %+v", d)
	}
}
func coverageExpected(t *testing.T, s p.Snapshot, m c.Module, xml, label string, exit int) k.Coverage {
	t.Helper()
	tests, err := k.ParseJUnit([]byte(xml))
	if err != nil {
		t.Fatal(err)
	}
	prefixes := []string{}
	for _, module := range s.Modules {
		prefixes = append(prefixes, module.IDPrefix)
	}
	cov := k.Analyze(m, prefixes, tests, exit, nil, nil)
	cov.Output = "eval-" + label + "-out\neval-" + label + "-err\n"
	return cov
}
func coverageAssert(t *testing.T, root string, args []string, want k.Coverage) result {
	t.Helper()
	got := invoke(t, root, args...)
	x := decode(t, got, "coverage")
	code := 0
	ds := []pathDiagnostic{}
	for _, d := range want.Diagnostics {
		ds = append(ds, pathDiagnostic{d.Code, d.Severity, d.Message, d.Path})
		if d.Severity == "error" {
			code = 1
		}
	}
	if got.code != code || !reflect.DeepEqual(decodeData(t, x.Data), canonical(t, want)) || !reflect.DeepEqual(detailedDiagnostics(t, got), ds) {
		t.Fatalf("coverage %+v expected %+v", got, want)
	}
	return got
}
func Test_APP013RealCurrentCoverage(t *testing.T) {
	root := coverageFixture(t)
	s := coverageSnapshot(t, root)
	m := coverageModule(t, s, "alpha")
	want := coverageExpected(t, s, m, coveragePassing, "current", 0)
	if len(want.Tests) != 3 || len(want.Diagnostics) != 0 || !reflect.DeepEqual(want.BySpec, map[string][]int{"ALP-001": {0}, "ALP-002": {1}, "ALP-003": {2}}) {
		t.Fatalf("coverage premise %+v", want)
	}
	nested := filepath.Join(root, coverageDir, "nested")
	if err := os.Mkdir(nested, 0755); err != nil {
		t.Fatal(err)
	}
	for _, args := range [][]string{{"coverage", "alpha", "--json"}, {"--json", "coverage", "alpha"}} {
		coverageAssert(t, nested, args, want)
		coverageTrace(t, root, "current", true)
	}
	plain := invoke(t, nested, "coverage", "alpha")
	if plain.code != 0 || plain.err != "" || !reflect.DeepEqual(decodeData(t, json.RawMessage(plain.out)), canonical(t, want)) {
		t.Fatalf("plain coverage %+v", plain)
	}
	coveragePretty(t, plain.out)
	// Current evidence remains available despite unrelated queued/static errors.
	cliQueueChange(t, root, "changes/0001-invalid.md", "alpha", "ALP", "bases: {ALP-999: deadbeef}\n", "## Modified\n\n### ALP-999\nIt MUST fail preview.\n")
	projectPut(t, root, "outside/SPEC.md", "outside warning")
	if ds := cliQueueExpected(t, root); len(ds) < 2 {
		t.Fatal("missing unrelated validation errors")
	}
	coverageAssert(t, nested, []string{"coverage", "alpha", "--json"}, want)
	for _, name := range []string{"missing", "Alpha", "ALPHA"} {
		os.Remove(filepath.Join(root, coverageDir, "trace"))
		got := invoke(t, root, "coverage", name, "--json")
		x := decode(t, got, "coverage")
		if got.code != 2 || string(x.Data) != "null" || len(x.Diagnostics) != 1 || x.Diagnostics[0].Code != "U01" || !strings.Contains(x.Diagnostics[0].Message, "alpha") || !strings.Contains(x.Diagnostics[0].Message, "beta") {
			t.Fatalf("exact module usage %+v", got)
		}
		if _, err := os.Stat(filepath.Join(root, coverageDir, "trace")); !os.IsNotExist(err) {
			t.Fatal("usage ran hook")
		}
	}
}
func Test_APP013FatalCoverageConfiguration(t *testing.T) {
	for _, mode := range []string{"nongit", "missing", "invalid", "incompatible"} {
		t.Run(mode, func(t *testing.T) {
			root := t.TempDir()
			if mode != "nongit" {
				gitFixtureCommand(t, root, "init", "--initial-branch=main")
			}
			if mode == "invalid" {
				projectPut(t, root, "fireweed.yaml", "version: [bad]\n")
			}
			if mode == "incompatible" {
				projectPut(t, root, "fireweed.yaml", "version: 1\nfireweed: '>=999'\n")
			}
			got := invoke(t, root, "coverage", "alpha", "--json")
			x := decode(t, got, "coverage")
			if got.code != 2 || string(x.Data) != "null" || len(x.Diagnostics) != 1 || x.Diagnostics[0].Code != "C01" || x.Diagnostics[0].Severity != "error" {
				t.Fatalf("configuration %+v", got)
			}
			plain := invoke(t, root, "coverage", "alpha")
			coveragePretty(t, plain.out)
			if plain.code != 2 || plain.out != "" || !strings.Contains(plain.err, "C01") || !strings.Contains(plain.err, "fireweed.yaml") {
				t.Fatalf("plain configuration %+v", plain)
			}
		})
	}
}
func coverageProjectedFixture(t *testing.T) string {
	t.Helper()
	root := coverageFixture(t)
	pre, _ := json.Marshal(coverageHook("pre"))
	projected, _ := json.Marshal(coverageHook("projected"))
	after, _ := json.Marshal(coverageHook("after"))
	assigned := "run: fixture\nphase: impl\n"
	cliQueueChange(t, root, "changes/0002-aa.md", "alpha", "ALP", assigned+"set: {hooks: {eval: "+string(pre)+"}}\n", "## Added\n\n### ALP-004\nIt MUST precede removal.\n")
	cliQueueChange(t, root, "changes/00002-zulu.md", "alpha", "ALP", assigned+"set: {hooks: {eval: "+string(projected)+"}}\n", "## Added\n\n### ALP-005\nIt MUST survive through the selected change.\n")
	cliQueueChange(t, root, "changes/0010-target.md", "alpha", "ALP", assigned+"bases: {ALP-001: deadbeef, ALP-003: deadbeef, ALP-004: deadbeef}\nfixes: [ALP-001]\nalready_satisfied: [ALP-006]\nset: {interface: [next.go]}\n", "## Added\n\n### ALP-006\nIt MUST be selected.\n\n## Modified\n\n### ALP-001\nIt MUST change inclusively.\n\n## Removed\n\n### ALP-003\n### ALP-004\n")
	cliQueueChange(t, root, "changes/0100-after.md", "alpha", "ALP", assigned+"set: {hooks: {eval: "+string(after)+"}}\n", "")
	cliQueueChange(t, root, "changes/archive/0001-archived.md", "alpha", "ALP", "run: old\nphase: applied\n", "## Added\n\n### ALP-007\nIt MUST stay historical.\n")
	coverageSet(t, root, "pre", "zero", `<testsuite><testcase name="ALP001"/><testcase name="ALP002"><skipped/></testcase><testcase name="ALP003"><skipped/></testcase><testcase name="ALP004"/></testsuite>`)
	coverageSet(t, root, "projected", "zero", `<testsuite><testcase name="ALP001"/><testcase name="ALP002"><skipped/></testcase><testcase name="ALP005"/><testcase name="ALP006"/></testsuite>`)
	coverageSet(t, root, "after", "zero", `<testsuite><testcase name="ALP001"/></testsuite>`)
	return root
}
func coverageSelected(t *testing.T, root, slug string) (p.Snapshot, c.Module) {
	t.Helper()
	s := coverageSnapshot(t, root)
	before, err := p.Project(s, slug)
	if err != nil {
		t.Fatal(err)
	}
	m := coverageModule(t, before, "alpha")
	for _, ch := range s.Changes {
		if ch.Slug == slug {
			m, err = p.ApplyChange(m, ch)
			if err != nil {
				t.Fatal(err)
			}
			return before, m
		}
	}
	t.Fatal("missing selected change")
	return p.Snapshot{}, c.Module{}
}
func Test_APP014InclusiveProjectedCoverage(t *testing.T) {
	root := coverageProjectedFixture(t)
	s, m := coverageSelected(t, root, "target")
	ids := []string{}
	for _, item := range m.Specs {
		ids = append(ids, item.ID)
	}
	if !reflect.DeepEqual(ids, []string{"ALP-001", "ALP-002", "ALP-005", "ALP-006"}) || m.Specs[0].Body != "It MUST change inclusively." || m.Hooks.Eval != coverageHook("projected") || !reflect.DeepEqual(m.Interface, []string{"next.go"}) {
		t.Fatalf("projection premise %+v", m)
	}
	raw, err := os.ReadFile(filepath.Join(root, coverageDir, "projected.xml"))
	if err != nil {
		t.Fatal(err)
	}
	want := coverageExpected(t, s, m, string(raw), "projected", 0)
	for _, args := range [][]string{{"coverage", "alpha", "--with-change", "target", "--json"}, {"--with-change=target", "--json", "coverage", "alpha"}, {"coverage", "--with-change", "aa", "alpha", "--with-change=target", "--json"}} {
		coverageAssert(t, root, args, want)
		coverageTrace(t, root, "projected", true)
	}
	// Neither fixes nor already_satisfied metadata turns normal coverage into fail-first.
	for _, id := range []string{"ALP001", "ALP006"} {
		failed := strings.Replace(string(raw), `name="`+id+`"/>`, `name="`+id+`"><failure/></testcase>`, 1)
		coverageSet(t, root, "projected", "fail", failed)
		expected := coverageExpected(t, s, m, failed, "projected", 37)
		if len(expected.Diagnostics) != 1 || expected.Diagnostics[0].Code != "C14" {
			t.Fatalf("ordinary failure premise %+v", expected)
		}
		coverageAssert(t, root, []string{"coverage", "alpha", "--with-change=target", "--json"}, expected)
	}
	missing := strings.Replace(string(raw), `<testcase name="ALP006"/>`, "", 1)
	coverageSet(t, root, "projected", "zero", missing)
	expected := coverageExpected(t, s, m, missing, "projected", 0)
	if len(expected.Diagnostics) != 1 || expected.Diagnostics[0].Code != "C13" {
		t.Fatalf("already_satisfied cannot waive coverage %+v", expected)
	}
	coverageAssert(t, root, []string{"coverage", "alpha", "--with-change=target", "--json"}, expected)
	current := coverageExpected(t, coverageSnapshot(t, root), coverageModule(t, coverageSnapshot(t, root), "alpha"), coveragePassing, "current", 0)
	coverageAssert(t, root, []string{"coverage", "alpha", "--json"}, current)
}
func Test_APP014ProjectionErrorsAndNoAllocation(t *testing.T) {
	root := coverageProjectedFixture(t)
	for _, args := range [][]string{{"coverage", "alpha", "--with-change=missing", "--json"}, {"coverage", "alpha", "--with-change=archived", "--json"}, {"coverage", "alpha", "--with-change=Target", "--json"}, {"coverage", "beta", "--with-change=target", "--json"}} {
		os.Remove(filepath.Join(root, coverageDir, "trace"))
		got := invoke(t, root, args...)
		x := decode(t, got, "coverage")
		if got.code != 2 || string(x.Data) != "null" || len(x.Diagnostics) != 1 || x.Diagnostics[0].Code != "U01" || !strings.Contains(x.Diagnostics[0].Message, "coverage") {
			t.Fatalf("selected usage %+v", got)
		}
		message := x.Diagnostics[0].Message
		if args[1] == "beta" {
			if !strings.Contains(message, "alpha") || !strings.Contains(message, "fireweed coverage alpha --with-change target") {
				t.Fatalf("missing target correction %s", message)
			}
		} else {
			for _, slug := range []string{"aa", "zulu", "target", "after"} {
				if !strings.Contains(message, slug) {
					t.Fatalf("missing queued choice %s: %s", slug, message)
				}
			}
			if !strings.Contains(message, "fireweed coverage alpha --with-change") {
				t.Fatalf("missing usable selection command %s", message)
			}
		}
		if _, err := os.Stat(filepath.Join(root, coverageDir, "trace")); !os.IsNotExist(err) {
			t.Fatal("bad selection ran alpha hook")
		}
		if _, err := os.Stat(filepath.Join(root, "MUST-NOT-RUN")); !os.IsNotExist(err) {
			t.Fatal("bad selection ran beta/default hook")
		}
	}
	cliQueueChange(t, root, "changes/0000-invalid-before.md", "alpha", "ALP", "bases: {ALP-999: deadbeef}\n", "## Modified\n\n### ALP-999\nIt MUST reject preview.\n")
	got := invoke(t, root, "coverage", "alpha", "--with-change=target", "--json")
	x := decode(t, got, "coverage")
	ds := detailedDiagnostics(t, got)
	if got.code != 1 || string(x.Data) != "null" || len(ds) != 1 || ds[0].Code != "C05" || ds[0].Path != "changes/0010-target.md" || !(strings.Contains(strings.ToLower(ds[0].Message), "correct") || strings.Contains(strings.ToLower(ds[0].Message), "repair")) {
		t.Fatalf("projection error %+v", got)
	}
	if _, err := os.Stat(filepath.Join(root, coverageDir, "trace")); !os.IsNotExist(err) {
		t.Fatal("invalid preview ran hook")
	}
	// The selected entry itself may fail even when the preceding projection succeeds.
	selectedRoot := coverageFixture(t)
	cliQueueChange(t, selectedRoot, "changes/0001-bad-selected.md", "alpha", "ALP", "bases: {ALP-999: deadbeef}\n", "## Modified\n\n### ALP-999\nIt MUST reject selection.\n")
	selectedSnapshot := coverageSnapshot(t, selectedRoot)
	projected, err := p.Project(selectedSnapshot, "bad-selected")
	if err != nil {
		t.Fatalf("preceding projection premise %v", err)
	}
	if _, err := p.ApplyChange(coverageModule(t, projected, "alpha"), selectedSnapshot.Changes[0]); err == nil {
		t.Fatal("selected application must fail")
	}
	selectedResult := invoke(t, selectedRoot, "coverage", "alpha", "--with-change=bad-selected", "--json")
	selectedEnvelope := decode(t, selectedResult, "coverage")
	selectedDS := detailedDiagnostics(t, selectedResult)
	if selectedResult.code != 1 || string(selectedEnvelope.Data) != "null" || len(selectedDS) != 1 || selectedDS[0].Code != "C05" || selectedDS[0].Path != "changes/0001-bad-selected.md" || !(strings.Contains(strings.ToLower(selectedDS[0].Message), "correct") || strings.Contains(strings.ToLower(selectedDS[0].Message), "repair")) {
		t.Fatalf("selected application error %+v", selectedResult)
	}
	if _, err := os.Stat(filepath.Join(selectedRoot, coverageDir, "trace")); !os.IsNotExist(err) {
		t.Fatal("failed selected application ran hook")
	}
	fresh := coverageFixture(t)
	cliQueueChange(t, fresh, "changes/0001-local.md", "alpha", "ALP", "", "## Added\n\n### new:fresh\nIt MUST wait for allocation.\n")
	before := showFiles(t, fresh)
	r := invoke(t, fresh, "coverage", "alpha", "--with-change=local", "--json")
	v := decode(t, r, "coverage")
	var cov k.Coverage
	if err := json.Unmarshal(v.Data, &cov); err != nil {
		t.Fatal(err)
	}
	if r.code != 1 || len(cov.Diagnostics) != 1 || cov.Diagnostics[0].Code != "C13" || !strings.Contains(cov.Diagnostics[0].Message, "new:fresh") {
		t.Fatalf("unallocated preview %+v", r)
	}
	after := showFiles(t, fresh)
	delete(after, coverageDir+"/trace")
	if !reflect.DeepEqual(before, after) {
		t.Fatal("coverage allocated or rewrote state")
	}
}
func Test_APP015CoverageRemovedIDs(t *testing.T) {
	root := coverageProjectedFixture(t)
	s, m := coverageSelected(t, root, "target")
	data, err := os.ReadFile(filepath.Join(root, coverageDir, "projected.xml"))
	if err != nil {
		t.Fatal(err)
	}
	xml := strings.Replace(string(data), "</testsuite>", `<testcase name="ALP003"/></testsuite>`, 1)
	coverageSet(t, root, "projected", "zero", xml)
	want := coverageExpected(t, s, m, xml, "projected", 0)
	if len(want.Diagnostics) != 1 || want.Diagnostics[0].Code != "C12" || !strings.Contains(want.Diagnostics[0].Message, "ALP-003") {
		t.Fatalf("removed-ID fixture %+v", want)
	}
	coverageAssert(t, root, []string{"coverage", "alpha", "--with-change=target", "--json"}, want)
	coverageTrace(t, root, "projected", true)
}
func Test_APP015CoverageDiagnostics(t *testing.T) {
	root := coverageFixture(t)
	s := coverageSnapshot(t, root)
	m := coverageModule(t, s, "alpha")
	for _, tt := range []struct {
		name, xml, mode string
		exit            int
		codes           []string
	}{
		{"ordinary failure", `<testsuite><testcase name="ALP001"><failure message="bad">trace</failure></testcase></testsuite>`, "fail", 37, []string{"C14"}},
		{"reporter inconsistency", `<testsuite><testcase name="ALP001"><error/></testcase></testsuite>`, "zero", 0, []string{"C11", "C14"}},
		{"unknown and missing", `<testsuite><testcase name="BET001"/><testcase name="ALP999"/><testcase name="without identifier"/></testsuite>`, "zero", 0, []string{"C12", "C12", "C12", "C13"}},
		{"skipped coverage", `<testsuite><testcase name="ALP001"><skipped/></testcase></testsuite>`, "zero", 0, []string{"C13"}},
		{"quarantined failure", `<testsuite><testcase name="ALP001"/><testcase name="ALP002"><failure/></testcase><testcase name="ALP003"><error/></testcase></testsuite>`, "fail", 37, []string{}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			coverageSet(t, root, "current", tt.mode, tt.xml)
			want := coverageExpected(t, s, m, tt.xml, "current", tt.exit)
			codes := []string{}
			for _, d := range want.Diagnostics {
				codes = append(codes, d.Code)
			}
			if !reflect.DeepEqual(codes, tt.codes) {
				t.Fatalf("fixture policy %+v", want)
			}
			coverageAssert(t, root, []string{"coverage", "alpha", "--json"}, want)
			coverageTrace(t, root, "current", true)
			plain := invoke(t, root, "coverage", "alpha")
			code := 0
			if len(want.Diagnostics) > 0 {
				code = 1
			}
			if plain.code != code || !reflect.DeepEqual(decodeData(t, json.RawMessage(plain.out)), canonical(t, want)) {
				t.Fatalf("plain evidence %+v", plain)
			}
			for _, d := range want.Diagnostics {
				if !strings.Contains(plain.err, d.Code) || !strings.Contains(plain.err, d.Path) || !strings.Contains(plain.err, d.Message) {
					t.Fatalf("plain diagnostic %+v", plain)
				}
			}
		})
	}
	for _, mode := range []string{"missing", "directory", "malformed", "timeout", "signal"} {
		t.Run(mode, func(t *testing.T) {
			xml := `<testsuite><testcase name="ALP001"><failure/></testcase></testsuite>`
			if mode == "malformed" {
				xml = `<testsuite><testcase name="ALP001"/><broken>`
			}
			coverageSet(t, root, "current", mode, xml)
			got := invoke(t, root, "coverage", "alpha", "--json")
			x := decode(t, got, "coverage")
			var cov k.Coverage
			if err := json.Unmarshal(x.Data, &cov); err != nil {
				t.Fatal(err)
			}
			if got.code != 1 || cov.Module != "alpha" || cov.Tests == nil || len(cov.Tests) != 0 || cov.BySpec == nil || len(cov.BySpec) != 0 || len(cov.Diagnostics) != 1 || cov.Diagnostics[0].Code != "C11" || cov.Diagnostics[0].Path != coverageDir || cov.Output != "eval-current-out\neval-current-err\n" {
				t.Fatalf("invalid report %+v", got)
			}
			coverageCorrective(t, cov.Diagnostics[0])
			if !reflect.DeepEqual(canonical(t, cov.Diagnostics), canonical(t, detailedDiagnostics(t, got))) {
				t.Fatal("envelope/data diagnostics differ")
			}
			coverageTrace(t, root, "current", true)
			if mode == "timeout" || mode == "signal" {
				if cov.HookExit != -1 {
					t.Fatalf("fatal actual exit %+v", cov)
				}
			} else if cov.HookExit != 0 {
				t.Fatalf("actual exit %+v", cov)
			}
			if mode == "timeout" {
				time.Sleep(2200 * time.Millisecond)
				if _, err := os.Stat(filepath.Join(root, coverageDir, "LATE")); !os.IsNotExist(err) {
					t.Fatal("timeout left hook running")
				}
			}
		})
	}
	raw, err := os.ReadFile(filepath.Join(root, coverageDir, "SPEC.md"))
	if err != nil {
		t.Fatal(err)
	}
	doc, err := c.ParseMarkdown(string(raw))
	if err != nil {
		t.Fatal(err)
	}
	doc.Metadata["hooks"] = map[string]any{"eval": ""}
	text, err := c.RenderMarkdown(doc)
	if err != nil {
		t.Fatal(err)
	}
	projectPut(t, root, coverageDir+"/SPEC.md", text)
	got := invoke(t, root, "coverage", "alpha", "--json")
	x := decode(t, got, "coverage")
	var cov k.Coverage
	if err := json.Unmarshal(x.Data, &cov); err != nil {
		t.Fatal(err)
	}
	if got.code != 1 || len(cov.Diagnostics) != 1 || cov.Diagnostics[0].Code != "C07" || cov.Diagnostics[0].Path != coverageDir || cov.Tests == nil || len(cov.Tests) != 0 || cov.BySpec == nil || len(cov.BySpec) != 0 {
		t.Fatalf("missing hook %+v", got)
	}
	coverageCorrective(t, cov.Diagnostics[0])
	if !reflect.DeepEqual(canonical(t, cov.Diagnostics), canonical(t, detailedDiagnostics(t, got))) {
		t.Fatal("missing hook diagnostics differ")
	}
}
func Test_APP016CoverageArgumentsAndWrites(t *testing.T) {
	for _, args := range [][]string{{"coverage"}, {"coverage", "alpha", "extra"}, {"coverage", "alpha", "--with-change"}, {"coverage", "alpha", "--with-change="}, {"coverage", "alpha", "--with-change", "--help"}, {"coverage", "alpha", "--with-change", "-h"}, {"coverage", "alpha", "--json=true"}, {"coverage", "alpha", "--unknown"}, {"coverage", "alpha", "--projected"}, {"coverage", "alpha", "--history"}, {"modules", "--with-change=x"}, {"check", "--with-change", "x"}, {"show", "alpha", "--with-change=x"}, {"version", "--with-change=x"}, {"help", "--with-change=x"}, {"explain", "C11", "--with-change=x"}, {"playbook", "evals", "--with-change=x"}} {
		got := invoke(t, t.TempDir(), append(args, "--json")...)
		x := decode(t, got, args[0])
		if got.code != 2 || string(x.Data) != "null" || len(x.Diagnostics) != 1 || x.Diagnostics[0].Code != "U01" || !strings.Contains(x.Diagnostics[0].Message, "help") {
			t.Fatalf("argument error %+v", got)
		}
		assertUsage(t, invoke(t, t.TempDir(), args...))
	}
	for _, args := range [][]string{{"coverage", "--help"}, {"coverage", "--with-change=x", "-h"}, {"--json", "coverage", "--help"}} {
		got := invoke(t, t.TempDir(), args...)
		if got.code != 0 || got.err != "" || !strings.Contains(got.out, "coverage") {
			t.Fatalf("help %+v", got)
		}
	}
	assertUsage(t, invoke(t, t.TempDir(), "coverage", "alpha", "--", "--json"))
	assertUsage(t, invoke(t, t.TempDir(), "coverage", "alpha", "--", "--with-change=x"))
	root := coverageFixture(t)
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(root); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(cwd)
	for _, args := range [][]string{{"coverage", "alpha"}, {"coverage", "alpha", "--json"}} {
		if code := app.Run(args, failedWriter{}, io.Discard); code != 2 {
			t.Fatalf("stdout failure%d", code)
		}
	}
	coverageSet(t, root, "current", "fail", `<testsuite><testcase name="ALP001"><failure/></testcase></testsuite>`)
	if code := app.Run([]string{"coverage", "alpha"}, io.Discard, failedWriter{}); code != 2 {
		t.Fatalf("stderr failure%d", code)
	}
}
func coverageControl(t *testing.T) string {
	t.Helper()
	root := coverageProjectedFixture(t)
	// Positive control proves the authored script's report protocol without product dispatch.
	control := filepath.Join(t.TempDir(), "control.xml")
	command := exec.Command("sh", "hooks/eval.sh", control, "alpha", coverageDir, "current")
	command.Dir = root
	command.Env = append(os.Environ(), "FIREWEED_MODULE=alpha", "FIREWEED_MODULE_DIR="+coverageDir, "FIREWEED_REPORT="+control, "FIREWEED_BIN="+subject)
	output, err := command.CombinedOutput()
	if err != nil || string(output) != "eval-current-out\neval-current-err\n" {
		t.Fatalf("helper control %v %s", err, output)
	}
	bytes, err := os.ReadFile(control)
	if err != nil || string(bytes) != coveragePassing {
		t.Fatalf("helper report %v %q", err, bytes)
	}
	coverageTrace(t, root, "current", false)
	os.Remove(filepath.Join(root, coverageDir, "trace"))
	return root
}
func Test_APP016CoverageHelperPositiveControl(t *testing.T) { coverageControl(t) }
func Test_APP016CoverageSafetyAndHelperControl(t *testing.T) {
	root := coverageControl(t)
	projectPut(t, root, ".gitignore", "ignored-sentinel\n"+coverageDir+"/trace\n")
	projectPut(t, root, "ignored-sentinel", "unchanged ignored bytes")
	realGit, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	trapRoot := t.TempDir()
	trapMarker := filepath.Join(trapRoot, "UNEXPECTED")
	script := "#!/bin/sh\nprintf trap >> " + coverageQuote(trapMarker) + "\nexit 97\n"
	for _, name := range []string{"codex", "claude", "agent-trap", "fireweed", "fireweed-bootstrap", "curl", "wget", "ssh", "nc"} {
		if err := os.WriteFile(filepath.Join(trapRoot, name), []byte(script), 0755); err != nil {
			t.Fatal(err)
		}
	}
	gitScript := "#!/bin/sh\nfor a in \"$@\"; do case \"$a\" in fetch|push|pull|clone|ls-remote|commit|update-ref|reset|checkout|switch|add|stash|clean|rm|mv) printf trap >> " + coverageQuote(trapMarker) + "; exit 97;; esac; done\nexec " + coverageQuote(realGit) + " \"$@\"\n"
	if err := os.WriteFile(filepath.Join(trapRoot, "git"), []byte(gitScript), 0755); err != nil {
		t.Fatal(err)
	}
	gitFixtureCommand(t, root, "add", ".")
	gitFixtureCommand(t, root, "-c", "user.name=Fixture", "-c", "user.email=fixture@example.invalid", "commit", "-m", "fixture")
	t.Setenv("PATH", trapRoot+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("FIREWEED_BIN", filepath.Join(trapRoot, "fireweed-bootstrap"))
	controlTrap := exec.Command("sh", "-c", "codex")
	if err := controlTrap.Run(); err == nil {
		t.Fatal("trap control did not fail")
	}
	if _, err := os.Stat(trapMarker); err != nil {
		t.Fatal("trap control did not run")
	}
	os.Remove(trapMarker)
	before := showFiles(t, root)
	refs := gitFixtureCommand(t, root, "show-ref")
	status := gitFixtureCommand(t, root, "status", "--porcelain=v1", "--untracked-files=all")
	for _, args := range [][]string{{"coverage", "alpha", "--json"}, {"coverage", "alpha", "--with-change=target", "--json"}} {
		got := invoke(t, root, args...)
		if got.code != 0 {
			t.Fatalf("valid safe coverage %+v", got)
		}
		label := "current"
		if len(args) > 3 {
			label = "projected"
		}
		coverageTrace(t, root, label, true)
	}
	afterSuccess := showFiles(t, root)
	delete(afterSuccess, coverageDir+"/trace")
	if !reflect.DeepEqual(before, afterSuccess) || gitFixtureCommand(t, root, "status", "--porcelain=v1", "--untracked-files=all") != status {
		t.Fatal("successful coverage mutated repository or ignored files")
	}
	for _, label := range []string{"current", "projected"} {
		xml := coveragePassing
		if label == "projected" {
			data, err := os.ReadFile(filepath.Join(root, coverageDir, "projected.xml"))
			if err != nil {
				t.Fatal(err)
			}
			xml = string(data)
		}
		coverageSet(t, root, label, "missing", xml)
		beforeFailure := showFiles(t, root)
		delete(beforeFailure, coverageDir+"/trace")
		failureStatus := gitFixtureCommand(t, root, "status", "--porcelain=v1", "--untracked-files=all")
		args := []string{"coverage", "alpha", "--json"}
		if label == "projected" {
			args = append(args, "--with-change=target")
		}
		got := invoke(t, root, args...)
		if got.code != 1 {
			t.Fatal(got)
		}
		coverageTrace(t, root, label, true)
		afterFailure := showFiles(t, root)
		delete(afterFailure, coverageDir+"/trace")
		if !reflect.DeepEqual(beforeFailure, afterFailure) || gitFixtureCommand(t, root, "status", "--porcelain=v1", "--untracked-files=all") != failureStatus {
			t.Fatal("failed coverage mutated repository or ignored files")
		}
		if gitFixtureCommand(t, root, "show-ref") != refs {
			t.Fatal("coverage mutated refs")
		}
	}
	if _, err := os.Stat(trapMarker); !os.IsNotExist(err) {
		t.Fatal("coverage invoked forbidden external command")
	}
	if _, err := os.Stat(filepath.Join(root, "MUST-NOT-RUN")); !os.IsNotExist(err) {
		t.Fatal("coverage ran setup/boundaries/default hook")
	}
}
