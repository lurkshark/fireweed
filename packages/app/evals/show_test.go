package evals

import (
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	app "github.com/lurkshark/fireweed/packages/app"
	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
)

func showFixture(t *testing.T) string {
	t.Helper()
	root := projectFixture(t)
	projectModule(t, root, "units/a", "alpha", "ALP", "depends_on: [beta]\npreserve: [keep/**]\nappend_only: [events/**]\nhooks: {setup: own setup}\n")
	projectModule(t, root, "units/b", "beta", "BET", "")
	path := filepath.Join(root, "units/a/SPEC.md")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	text := strings.Replace(string(raw), "## Decisions", "### ALP-002 [quarantine: past clock]\nSecond SHOULD stay quarantined.\n\n### ALP-003\nThird MUST remain current.\n\n## Decisions", 1) + "\n## Extra module context\n\nRetain this module context.\n"
	cfg, err := c.ParseConfig(projectConfig, "0.4.0")
	if err != nil {
		t.Fatal(err)
	}
	if _, err = c.ParseModule(text, "units/a", cfg.ModuleDefaults, cfg.Hooks); err != nil {
		t.Fatal(err)
	}
	projectPut(t, root, "units/a/SPEC.md", text)
	return root
}
func showModule(t *testing.T, s p.Snapshot, name string) c.Module {
	t.Helper()
	for _, m := range s.Modules {
		if m.Name == name {
			return m
		}
	}
	t.Fatalf("missing fixture module %s", name)
	return c.Module{}
}
func showSnapshot(t *testing.T, root string) p.Snapshot {
	t.Helper()
	s, err := p.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	return s
}
func showDiagCode(ds []pathDiagnostic) int {
	for _, d := range ds {
		if d.Severity == "error" {
			return 1
		}
	}
	return 0
}
func assertShow(t *testing.T, root string, args []string, data any, ds []pathDiagnostic, code int) result {
	t.Helper()
	got := invoke(t, root, args...)
	x := decode(t, got, "show")
	if got.code != code || got.err != "" || !reflect.DeepEqual(decodeData(t, x.Data), canonical(t, data)) || !reflect.DeepEqual(detailedDiagnostics(t, got), ds) {
		t.Fatalf("show %+v expected data %+v diagnostics %+v exit%d", got, data, ds, code)
	}
	return got
}
func showUsageJSON(t *testing.T, root string, args ...string) {
	t.Helper()
	got := invoke(t, root, args...)
	command := args[0]
	x := decode(t, got, command)
	if got.code != 2 || got.err != "" || string(x.Data) != "null" || len(x.Diagnostics) != 1 || x.Diagnostics[0].Code != "U01" || x.Diagnostics[0].Severity != "error" || x.Diagnostics[0].Message == "" {
		t.Fatalf("show usage %+v", got)
	}
}

const currentShowText = "### ALP-001 [review]\nIt MUST retain this body.\n\n### ALP-002 [quarantine: past clock]\nSecond SHOULD stay quarantined.\n\n### ALP-003\nThird MUST remain current.\n"

func Test_APP009CurrentSpecReads(t *testing.T) {
	root := showFixture(t)
	s := showSnapshot(t, root)
	m := showModule(t, s, "alpha")
	if ds := p.Check(s); len(ds) != 0 {
		t.Fatalf("invalid fixture %+v", ds)
	}
	nested := filepath.Join(root, "units/a/nested/deeper")
	if err := os.MkdirAll(nested, 0755); err != nil {
		t.Fatal(err)
	}
	for _, args := range [][]string{{"show", "alpha", "--json"}, {"--json", "show", "alpha"}} {
		assertShow(t, nested, args, m, []pathDiagnostic{}, 0)
	}
	for _, args := range [][]string{{"show", "alpha"}, {"show", "--", "alpha"}} {
		got := invoke(t, nested, args...)
		if got.code != 0 || got.err != "" || got.out != currentShowText {
			t.Fatalf("current plain %+v", got)
		}
	}
	for _, item := range m.Specs {
		assertShow(t, nested, []string{"show", "alpha", item.ID, "--json"}, item, []pathDiagnostic{}, 0)
		want := "### " + item.ID + " " + item.Tag + "\n" + item.Body + "\n"
		if item.Tag == "" {
			want = "### " + item.ID + "\n" + item.Body + "\n"
		}
		got := invoke(t, nested, "show", "alpha", item.ID)
		if got.code != 0 || got.err != "" || got.out != want {
			t.Fatalf("single spec plain %+v", got)
		}
	}
	for _, id := range []string{"alp-001", "ALP001", "ALP-999"} {
		showUsageJSON(t, nested, "show", "alpha", id, "--json")
	}
	showUsageJSON(t, nested, "show", "missing", "--json")
	got := invoke(t, nested, "show", "missing", "--json")
	x := decode(t, got, "show")
	if !strings.Contains(x.Diagnostics[0].Message, "alpha") || !strings.Contains(x.Diagnostics[0].Message, "beta") {
		t.Fatal("unknown module lacks choices")
	}
	empty := showFixture(t)
	em := showModule(t, showSnapshot(t, empty), "alpha")
	for i := range em.Document.Sections {
		if em.Document.Sections[i].Name == "Specs" {
			em.Document.Sections[i].Body = ""
		}
	}
	text, err := c.RenderMarkdown(em.Document)
	if err != nil {
		t.Fatal(err)
	}
	projectPut(t, empty, "units/a/SPEC.md", text)
	em = showModule(t, showSnapshot(t, empty), "alpha")
	if em.Specs == nil || len(em.Specs) != 0 {
		t.Fatal("empty module fixture")
	}
	assertShow(t, empty, []string{"show", "alpha", "--json"}, em, []pathDiagnostic{}, 0)
	plain := invoke(t, empty, "show", "alpha")
	if plain.code != 0 || plain.out != "" || plain.err != "" {
		t.Fatalf("empty plain %+v", plain)
	}
}
func Test_APP009CurrentDiagnosticsAndConfig(t *testing.T) {
	for _, errors := range []bool{false, true} {
		root := showFixture(t)
		projectPut(t, root, "outside/SPEC.md", "outside warning")
		if errors {
			projectPut(t, root, "units/bad/SPEC.md", "---\nname: bad\nid_prefix: BAD\n---\n\n## Specs\n\n### BAD-001\nNo normative body.\n")
		}
		s, ds := projectExpected(t, root)
		if len(ds) < 1 || ds[0].Code != "C02" {
			t.Fatalf("diagnostic fixture %+v", ds)
		}
		code := 0
		if errors {
			code = 1
			found := false
			for _, d := range ds {
				if d.Code == "C04" {
					found = true
				}
			}
			if !found {
				t.Fatal("missing C04 fixture")
			}
		}
		assertShow(t, root, []string{"show", "alpha", "--json"}, showModule(t, s, "alpha"), ds, code)
		plain := invoke(t, root, "show", "alpha")
		if plain.code != code || plain.out != currentShowText {
			t.Fatalf("plain diagnostics %+v", plain)
		}
		for _, d := range ds {
			if !strings.Contains(plain.err, d.Code) || !strings.Contains(plain.err, d.Path) || !strings.Contains(plain.err, d.Message) {
				t.Fatalf("lost plain diagnostic %+v", plain)
			}
		}
	}
	for _, kind := range []string{"nongit", "missing", "invalid", "incompatible"} {
		root := t.TempDir()
		if kind != "nongit" {
			root = projectFixture(t)
		}
		switch kind {
		case "missing":
			if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
				t.Fatal(err)
			}
		case "invalid":
			projectPut(t, root, "fireweed.yaml", "version: []\n")
		case "incompatible":
			projectPut(t, root, "fireweed.yaml", "fireweed: '>99'\n")
		}
		for _, mode := range []string{"", "--projected", "--history"} {
			args := []string{"show", "alpha"}
			if mode != "" {
				args = append(args, mode)
			}
			jsonArgs := append(append([]string{}, args...), "--json")
			got := invoke(t, root, jsonArgs...)
			x := decode(t, got, "show")
			if got.code != 2 || got.err != "" || string(x.Data) != "null" || len(x.Diagnostics) != 1 || x.Diagnostics[0].Code != "C01" {
				t.Fatalf("config %s %s %+v", kind, mode, got)
			}
			plain := invoke(t, root, args...)
			if plain.code != 2 || plain.out != "" || !strings.Contains(plain.err, "C01") || strings.Contains(plain.err, "U01") {
				t.Fatalf("plain config %+v", plain)
			}
		}
	}
}
func projectedShowFixture(t *testing.T) string {
	t.Helper()
	root := showFixture(t)
	m := showModule(t, showSnapshot(t, root), "alpha")
	first := c.Spec{ID: "ALP-001", Tag: "[review]", Body: "It MUST become first."}
	second := c.Spec{ID: "ALP-001", Tag: "[review]", Body: "It MUST become second."}
	cliQueueChange(t, root, "changes/0002-alpha.md", "alpha", "ALP", "bases: {ALP-001: '"+c.SpecHash(m.Specs[0])+"'}\n", "## Modified\n\n### ALP-001 [review]\n"+first.Body+"\n")
	cliQueueChange(t, root, "changes/00002-zulu.md", "alpha", "ALP", "bases: {ALP-001: '"+c.SpecHash(first)+"'}\n", "## Modified\n\n### ALP-001 [review]\n"+second.Body+"\n")
	cliQueueChange(t, root, "changes/00010-last.md", "alpha", "ALP", "bases: {ALP-002: '"+c.SpecHash(m.Specs[1])+"'}\nset: {interface: [next.go], preserve: [], hooks: {setup: projected setup}}\n", "## Added\n\n### new:fresh\nIt MUST echo new:fresh.\n\n## Removed\n\n### ALP-002\n")
	cliQueueChange(t, root, "changes/archive/0001-past.md", "alpha", "ALP", "run: archived\nphase: applied\nbases: {ALP-003: abcdef12}\n", "## Modified\n\n### ALP-003\nIt MUST remain historical only.\n")
	if ds := cliQueueExpected(t, root); len(ds) != 0 {
		t.Fatalf("invalid projection fixture %+v", ds)
	}
	return root
}
func Test_APP010ProjectedSpecReads(t *testing.T) {
	root := projectedShowFixture(t)
	s := showSnapshot(t, root)
	projected, err := p.Project(s, "")
	if err != nil {
		t.Fatal(err)
	}
	m := showModule(t, projected, "alpha")
	if len(m.Specs) != 3 || m.Specs[0].Body != "It MUST become second." || m.Specs[1].ID != "ALP-003" || m.Specs[1].Body != "Third MUST remain current." || m.Specs[2].ID != "new:fresh" || m.Preserve == nil || len(m.Preserve) != 0 || !reflect.DeepEqual(m.Interface, []string{"next.go"}) || m.Hooks.Setup != "projected setup" {
		t.Fatalf("projection premise %+v", m)
	}
	for _, args := range [][]string{{"show", "alpha", "--projected", "--json"}, {"--json", "show", "--projected", "alpha"}} {
		assertShow(t, root, args, m, []pathDiagnostic{}, 0)
	}
	assertShow(t, root, []string{"show", "alpha", "ALP-001", "--projected", "--json"}, m.Specs[0], []pathDiagnostic{}, 0)
	assertShow(t, root, []string{"show", "alpha", "new:fresh", "--projected", "--json"}, m.Specs[2], []pathDiagnostic{}, 0)
	plain := invoke(t, root, "show", "alpha", "--projected")
	want := "### ALP-001 [review]\nIt MUST become second.\n\n### ALP-003\nThird MUST remain current.\n\n### new:fresh\nIt MUST echo new:fresh.\n"
	if plain.code != 0 || plain.err != "" || plain.out != want {
		t.Fatalf("projected plain %+v", plain)
	}
	assertShow(t, root, []string{"show", "alpha", "--json"}, showModule(t, s, "alpha"), []pathDiagnostic{}, 0)
	cliQueueChange(t, root, "changes/0011-optout.md", "alpha", "ALP", "set: {hooks: {boundaries: ''}, boundaries_opt_out: visible projected reason}\n", "")
	projectPut(t, root, "outside/SPEC.md", "outside warning")
	s = showSnapshot(t, root)
	projected, err = p.Project(s, "")
	if err != nil {
		t.Fatal(err)
	}
	ds := cliQueueExpected(t, root)
	if len(ds) != 2 || ds[0].Code != "D04" || ds[0].Path != "changes/0011-optout.md" || ds[1].Code != "C02" {
		t.Fatalf("projected warning fixture %+v", ds)
	}
	assertShow(t, root, []string{"show", "alpha", "--projected", "--json"}, showModule(t, projected, "alpha"), ds, 0)
	plain = invoke(t, root, "show", "alpha", "--projected")
	if plain.code != 0 || plain.out != want {
		t.Fatalf("plain projected warning %+v", plain)
	}
	for _, d := range ds {
		if !strings.Contains(plain.err, d.Code) || !strings.Contains(plain.err, d.Path) || !strings.Contains(plain.err, d.Message) {
			t.Fatalf("plain projected diagnostic %+v", plain)
		}
	}
}
func Test_APP010ProjectionFailuresAndAmbiguity(t *testing.T) {
	root := projectedShowFixture(t)
	path := filepath.Join(root, "changes/0002-alpha.md")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(string(raw), "\n")
	kept := []string{}
	for _, line := range lines {
		if !strings.HasPrefix(line, "bases:") {
			kept = append(kept, line)
		}
	}
	projectPut(t, root, "changes/0002-alpha.md", strings.Join(kept, "\n"))
	s := showSnapshot(t, root)
	projected, err := p.Project(s, "")
	if err != nil {
		t.Fatal(err)
	}
	ds := cliQueueExpected(t, root)
	if len(ds) != 1 || ds[0].Code != "C10" {
		t.Fatalf("stale fixture %+v", ds)
	}
	assertShow(t, root, []string{"show", "alpha", "--projected", "--json"}, showModule(t, projected, "alpha"), ds, 1)
	cliQueueChange(t, root, "changes/0000-invalid-before.md", "alpha", "ALP", "bases: {ALP-999: deadbeef}\n", "## Modified\n\n### ALP-999\nIt MUST fail application.\n")
	ds = cliQueueExpected(t, root)
	found := false
	for _, d := range ds {
		if d.Code == "C05" && d.Path == "changes/0000-invalid-before.md" {
			found = true
		}
	}
	if !found {
		t.Fatal("application fixture lacks C05")
	}
	assertShow(t, root, []string{"show", "alpha", "--projected", "--json"}, nil, ds, 1)
	plain := invoke(t, root, "show", "alpha", "--projected")
	if plain.code != 1 || plain.out != "" || !strings.Contains(plain.err, "C05") || !strings.Contains(plain.err, "changes/0000-invalid-before.md") {
		t.Fatalf("partial plain projection %+v", plain)
	}
	duplicate := projectedShowFixture(t)
	cliQueueChange(t, duplicate, "changes/0011-another.md", "alpha", "ALP", "", "## Added\n\n### new:fresh\nIt MUST remain a distinct declaration.\n")
	s = showSnapshot(t, duplicate)
	projected, err = p.Project(s, "")
	if err != nil {
		t.Fatal(err)
	}
	if ds := cliQueueExpected(t, duplicate); len(ds) != 0 {
		t.Fatalf("duplicate local namespace fixture %+v", ds)
	}
	assertShow(t, duplicate, []string{"show", "alpha", "--projected", "--json"}, showModule(t, projected, "alpha"), []pathDiagnostic{}, 0)
	showUsageJSON(t, duplicate, "show", "alpha", "new:fresh", "--projected", "--json")
	got := invoke(t, duplicate, "show", "alpha", "new:fresh", "--projected", "--json")
	x := decode(t, got, "show")
	if !strings.Contains(strings.ToLower(x.Diagnostics[0].Message), "ambig") {
		t.Fatal("missing handle ambiguity guidance")
	}
	showUsageJSON(t, duplicate, "show", "alpha", "new:fresh", "--json")
}

func historyBody(d c.Document, name string) string {
	for _, s := range d.Sections {
		if s.Name == name {
			return s.Body
		}
	}
	return ""
}
func archiveShow(t *testing.T, root, file, module, metadata, edits, motivation, decisions, provenance string) {
	t.Helper()
	text := "---\nmodule: " + module + "\nrun: history\nphase: applied\n" + metadata + "---\n\n## Motivation\n\n" + motivation + "\n\n" + edits + "\n\n## Decisions\n\n" + decisions + "\n\n## Provenance\n\n" + provenance + "\n"
	if _, err := c.ParseMarkdown(text); err != nil {
		t.Fatal(err)
	}
	projectPut(t, root, "changes/archive/"+file, text)
}
func expectedHistory(t *testing.T, s p.Snapshot, paths []string) any {
	t.Helper()
	rows := []any{}
	for _, path := range paths {
		found := false
		for _, ch := range s.Archive {
			if ch.Path == "changes/archive/"+path {
				rows = append(rows, map[string]any{"change": ch, "motivation": historyBody(ch.Document, "Motivation"), "decisions": historyBody(ch.Document, "Decisions"), "provenance": historyBody(ch.Document, "Provenance")})
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("archive fixture missing %s", path)
		}
	}
	return canonical(t, rows)
}
func historyShowFixture(t *testing.T) string {
	t.Helper()
	root := showFixture(t)
	archiveShow(t, root, "0001-origin.md", "alpha", "", "## Added\n\n### ALP-001\nInitial behavior MUST exist.", "Why origin.", "Original choice.", "- impl | 2026-10-22T10:00:00Z | agent: fixture")
	archiveShow(t, root, "0010-older-high.md", "alpha", "bases: {ALP-001: abcdef12}\n", "## Modified\n\n### ALP-001\nEarlier behavior MUST exist.", "Why older.", "Earlier choice.", "- impl | 2026-10-22T11:00:00Z | agent: fixture")
	archiveShow(t, root, "0005-offset-time.md", "alpha", "bases: {ALP-001: '12345678'}\n", "## Modified\n\n### ALP-001\nIntermediate behavior MUST exist.", "Why offset.", "Offset choice.", "- impl | 2026-10-22T13:30:00+02:00 | agent: fixture")
	// Earlier impl and later audit/unparseable entries must not replace the final
	// parseable impl time. Historical human handle examples are retained verbatim.
	provenance := "- impl | 2026-10-22T23:00:00Z | agent: earlier attempt\n- impl | 2026-10-22T12:00:00Z | agent: final\n- audit | 2026-10-22T23:30:00Z | agent: other\n- impl | invalid timestamp | agent: note"
	archiveShow(t, root, "0002-newer-low.md", "alpha", "bases: {ALP-001: '87654321'}\n", "## Modified\n\n### ALP-001 [review]\nIt MUST retain this body.", "Why newest, including historical new:example.", "Keep literal new:history in historical choices.\n\n```text\nnew:fenced-example\n```", provenance)
	archiveShow(t, root, "0013-highest-date.md", "alpha", "", "", "Why latest maintenance.", "Keep latest maintenance.", "- impl | 2026-10-22T14:00:00+01:00 | agent: fixture")
	for _, slug := range []string{"aa", "zz"} {
		archiveShow(t, root, "0012-"+slug+".md", "alpha", "", "", "Why tie "+slug+".", "Tie choice "+slug+".", "- impl | 2026-10-22T12:00:00Z | agent: fixture")
	}
	archiveShow(t, root, "0004-maintenance.md", "alpha", "", "", "Why legacy maintenance.", "Keep legacy maintenance.", "Legacy unstructured provenance")
	for _, slug := range []string{"fix-a", "fix-z"} {
		archiveShow(t, root, "0003-"+slug+".md", "alpha", "fixes: [ALP-001]\n", "", "Why "+slug+".", "Fix context.", "- audit | 2026-10-22T23:00:00Z | agent: audit only")
	}
	archiveShow(t, root, "0000-removed.md", "alpha", "bases: {ALP-099: abcdef12}\n", "## Removed\n\n### ALP-099", "Why removal.", "Removal context.", "")
	archiveShow(t, root, "9999-other.md", "beta", "", "## Added\n\n### BET-001\nIt MUST retain this body.", "Other module.", "Other choice.", "- impl | 2026-10-22T23:59:00Z | agent: other")
	current := showModule(t, showSnapshot(t, root), "alpha")
	cliQueueChange(t, root, "changes/9999-pending.md", "alpha", "ALP", "bases: {ALP-001: '"+c.SpecHash(current.Specs[0])+"'}\n", "## Modified\n\n### ALP-001 [review]\nPending behavior MUST stay outside history.\n")
	s := showSnapshot(t, root)
	if ds := p.Check(s); len(ds) != 0 || len(s.Archive) != 12 {
		t.Fatalf("history fixture %d %+v", len(s.Archive), ds)
	}
	return root
}
func Test_APP011ArchiveHistoryChronologyAndSelection(t *testing.T) {
	root := historyShowFixture(t)
	s := showSnapshot(t, root)
	all := []string{"0013-highest-date.md", "0012-zz.md", "0012-aa.md", "0002-newer-low.md", "0005-offset-time.md", "0010-older-high.md", "0001-origin.md", "0004-maintenance.md", "0003-fix-z.md", "0003-fix-a.md", "0000-removed.md"}
	spec := []string{"0002-newer-low.md", "0005-offset-time.md", "0010-older-high.md", "0001-origin.md"}
	for _, v := range []struct {
		args  []string
		paths []string
	}{{[]string{"show", "alpha", "--history", "--json"}, all}, {[]string{"--json", "show", "alpha", "ALP-001", "--history"}, spec}, {[]string{"show", "alpha", "ALP-099", "--history", "--json"}, []string{"0000-removed.md"}}, {[]string{"show", "alpha", "ALP-999", "--history", "--json"}, []string{}}, {[]string{"show", "alpha", "alp-001", "--history", "--json"}, []string{}}} {
		assertShow(t, root, v.args, expectedHistory(t, s, v.paths), []pathDiagnostic{}, 0)
		plainArgs := []string{}
		for _, a := range v.args {
			if a != "--json" {
				plainArgs = append(plainArgs, a)
			}
		}
		got := invoke(t, root, plainArgs...)
		if got.code != 0 || got.err != "" || !reflect.DeepEqual(decodeData(t, json.RawMessage(got.out)), expectedHistory(t, s, v.paths)) {
			t.Fatalf("plain history %+v", got)
		}
	}
	// Global current diagnostics are preserved while history ignores valid or
	// invalid pending applications and keeps the same archived reasons.
	cliQueueChange(t, root, "changes/0000-bad-pending.md", "alpha", "ALP", "bases: {ALP-999: abcdef12}\n", "## Modified\n\n### ALP-999\nIt MUST remain pending.\n")
	projectPut(t, root, "outside/SPEC.md", "warning outside")
	s, ds := projectExpected(t, root)
	if len(ds) != 1 || ds[0].Code != "C02" {
		t.Fatalf("current history warnings %+v", ds)
	}
	assertShow(t, root, []string{"show", "alpha", "--history", "--json"}, expectedHistory(t, s, all), ds, 0)
	for _, file := range []string{"0012-tie.md", "00012-tie.md"} {
		archiveShow(t, root, file, "alpha", "", "", "Tie path.", "Tie path choice.", "- impl | 2026-10-22T12:00:00Z | agent: fixture")
	}
	s, ds = projectExpected(t, root)
	found := false
	for _, d := range ds {
		if d.Code == "C05" {
			found = true
		}
	}
	if !found {
		t.Fatal("duplicate slug fixture lacks C05")
	}
	withPaths := []string{"0013-highest-date.md", "0012-zz.md", "0012-tie.md", "00012-tie.md", "0012-aa.md", "0002-newer-low.md", "0005-offset-time.md", "0010-older-high.md", "0001-origin.md", "0004-maintenance.md", "0003-fix-z.md", "0003-fix-a.md", "0000-removed.md"}
	assertShow(t, root, []string{"show", "alpha", "--history", "--json"}, expectedHistory(t, s, withPaths), ds, 1)
	plain := invoke(t, root, "show", "alpha", "--history")
	if plain.code != 1 || !reflect.DeepEqual(decodeData(t, json.RawMessage(plain.out)), expectedHistory(t, s, withPaths)) {
		t.Fatalf("plain history diagnostics %+v", plain)
	}
	for _, d := range ds {
		if !strings.Contains(plain.err, d.Code) || !strings.Contains(plain.err, d.Path) || !strings.Contains(plain.err, d.Message) {
			t.Fatalf("plain history stream %+v", plain)
		}
	}
	empty := showFixture(t)
	assertShow(t, empty, []string{"show", "alpha", "--history", "--json"}, []any{}, []pathDiagnostic{}, 0)
}

func showFiles(t *testing.T, root string) map[string]string {
	t.Helper()
	out := map[string]string{}
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			if entry.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		data, e := os.ReadFile(path)
		if e != nil {
			return e
		}
		relative, e := filepath.Rel(root, path)
		if e != nil {
			return e
		}
		out[relative] = string(data)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return out
}
func Test_APP012ShowArgumentsAndSafety(t *testing.T) {
	for _, args := range [][]string{{"show"}, {"show", "alpha", "ALP-001", "extra"}, {"show", "alpha", "--unknown"}, {"show", "alpha", "--json=true"}, {"show", "alpha", "--projected=true"}, {"show", "alpha", "--history=false"}, {"show", "alpha", "--projected", "--history"}, {"modules", "--history"}, {"check", "--projected"}, {"version", "--history"}, {"help", "--projected"}, {"explain", "C01", "--history"}, {"playbook", "author-change", "--projected"}} {
		assertUsage(t, invoke(t, t.TempDir(), args...))
		jsonArgs := append(append([]string{}, args...), "--json")
		showUsageJSON(t, t.TempDir(), jsonArgs...)
	}
	for _, args := range [][]string{{"show", "--help", "--json"}, {"show", "--projected", "-h", "--json"}, {"show", "--history", "--help", "--json"}} {
		successJSON(t, invoke(t, t.TempDir(), args...), "help")
	}
	root := showFixture(t)
	assertUsage(t, invoke(t, root, "show", "alpha", "--", "--json"))
	projectPut(t, root, ".gitignore", "ignored\n")
	projectPut(t, root, "ignored", "KEEP IGNORED BYTES")
	cliQueueChange(t, root, "changes/0001-safe.md", "alpha", "ALP", "", "## Added\n\n### new:guard\nIt MUST remain an inspection-only handle.\n")
	gitFixtureCommand(t, root, "add", ".")
	gitFixtureCommand(t, root, "-c", "user.name=Fixture", "-c", "user.email=fixture@example.invalid", "commit", "-m", "fixture")
	marker := ""
	if runtime.GOOS != "windows" {
		trap := t.TempDir()
		marker = filepath.Join(trap, "EXECUTED")
		realGit, err := exec.LookPath("git")
		if err != nil {
			t.Fatal(err)
		}
		quote := func(s string) string { return "'" + strings.ReplaceAll(s, "'", "'\"'\"'") + "'" }
		for _, name := range []string{"agent-trap", "codex", "claude", "fireweed", "fireweed-bootstrap", "curl", "wget", "ssh", "nc"} {
			projectPut(t, trap, name, "#!/bin/sh\necho invoked >> "+quote(marker)+"\nexit 99\n")
			if err = os.Chmod(filepath.Join(trap, name), 0755); err != nil {
				t.Fatal(err)
			}
		}
		projectPut(t, trap, "git", "#!/bin/sh\ncase \" $* \" in *' fetch '*|*' push '*|*' clone '*|*' ls-remote '*|*' update-ref '*|*' reset '*|*' checkout '*|*' switch '*|*' commit '*|*' add '*|*' stash '*|*' clean '*|*' rm '*|*' mv '*) echo mutation-or-network >> "+quote(marker)+";exit 99;;esac\nexec "+quote(realGit)+" \"$@\"\n")
		if err = os.Chmod(filepath.Join(trap, "git"), 0755); err != nil {
			t.Fatal(err)
		}
		t.Setenv("PATH", trap+string(os.PathListSeparator)+os.Getenv("PATH"))
		t.Setenv("FIREWEED_BIN", filepath.Join(trap, "fireweed-bootstrap"))
	}
	nested := filepath.Join(root, "units/a/nested")
	if err := os.MkdirAll(nested, 0755); err != nil {
		t.Fatal(err)
	}
	for _, invalid := range []bool{false, true} {
		if invalid {
			cliQueueChange(t, root, "changes/0001-safe.md", "alpha", "ALP", "bases: {ALP-999: deadbeef}\n", "## Modified\n\n### ALP-999\nIt MUST fail application only.\n")
		}
		before := showFiles(t, root)
		refs := gitFixtureCommand(t, root, "show-ref") + gitFixtureCommand(t, root, "status", "--porcelain=v1", "--untracked-files=all")
		s := showSnapshot(t, root)
		current, ds := projectExpected(t, root)
		m := showModule(t, current, "alpha")
		for _, mode := range []string{"", "--projected", "--history"} {
			args := []string{"show", "alpha"}
			if mode != "" {
				args = append(args, mode)
			}
			data := any(m)
			expectedDS := ds
			code := 0
			if mode == "--history" {
				data = []any{}
			}
			if mode == "--projected" {
				expectedDS = cliQueueExpected(t, root)
				code = showDiagCode(expectedDS)
				if invalid {
					data = nil
				} else {
					projected, err := p.Project(s, "")
					if err != nil {
						t.Fatal(err)
					}
					data = showModule(t, projected, "alpha")
				}
			}
			assertShow(t, nested, append(append([]string{}, args...), "--json"), data, expectedDS, code)
			plain := invoke(t, nested, args...)
			if plain.code != code {
				t.Fatalf("safe plain %+v", plain)
			}
		}
		if !reflect.DeepEqual(showFiles(t, root), before) || refs != gitFixtureCommand(t, root, "show-ref")+gitFixtureCommand(t, root, "status", "--porcelain=v1", "--untracked-files=all") {
			t.Fatal("show mutated files, ignored contents or refs")
		}
		if marker != "" {
			if _, err := os.Stat(marker); !os.IsNotExist(err) {
				t.Fatal("show executed agent/bootstrap/network/mutation command")
			}
		}
		if _, err := os.Stat(filepath.Join(root, "MUST-NOT-RUN")); !os.IsNotExist(err) {
			t.Fatal("show ran a project hook")
		}
	}
}
func Test_APP012ShowWriteFailures(t *testing.T) {
	root := projectedShowFixture(t)
	projectPut(t, root, "outside/SPEC.md", "outside warning")
	old, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err = os.Chdir(root); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(old); err != nil {
			t.Error(err)
		}
	}()
	for _, mode := range []string{"", "--projected", "--history"} {
		args := []string{"show", "alpha"}
		if mode != "" {
			args = append(args, mode)
		}
		for _, jsonMode := range []bool{false, true} {
			actual := append([]string{}, args...)
			if jsonMode {
				actual = append(actual, "--json")
			}
			if code := app.Run(actual, failedWriter{}, io.Discard); code != 2 {
				t.Fatalf("stdout failure %v returned %d", actual, code)
			}
		}
		if code := app.Run(args, io.Discard, failedWriter{}); code != 2 {
			t.Fatalf("diagnostic writer failure %v returned%d", args, code)
		}
	}
}
