package evals

import (
	c "github.com/lurkshark/fireweed/packages/contracts"
	"reflect"
	"strings"
	"testing"
)

func changeDoc(meta, body string) string {
	return "---\nmodule: auth\n" + meta + "---\n\n## Motivation\nProtect accounts.\n\n" + body
}
func moduleDoc(meta, body string) string {
	return "---\nname: auth\nid_prefix: AUTH\n" + meta + "---\n# Auth\n## Specs\n" + body + "\n## Decisions\nKeep errors stable.\n"
}

func Test_CON012Markdown(t *testing.T) {
	text := "---\nname: auth\nnested: {enabled: true, list: [one, two]}\nnote: |\n  ## metadata text\n---\n\n# Auth\n\n## Notes\n  Keep indentation.\n\n```md\n## hidden\n```\n\n## Unknown human section\n~~~\n## also hidden\n~~~\n"
	got, err := c.ParseMarkdown(text)
	if err != nil {
		t.Fatal(err)
	}
	wantSections := []c.Section{{Name: "Notes", Body: "  Keep indentation.\n\n```md\n## hidden\n```"}, {Name: "Unknown human section", Body: "~~~\n## also hidden\n~~~"}}
	if got.Preamble != "# Auth" || !reflect.DeepEqual(got.Sections, wantSections) || got.Metadata["name"] != "auth" {
		t.Fatalf("parsed %+v", got)
	}
	rendered, err := c.RenderMarkdown(got)
	if err != nil {
		t.Fatal(err)
	}
	round, err := c.ParseMarkdown(rendered)
	if err != nil || !reflect.DeepEqual(round, got) {
		t.Errorf("round trip: %+v %v, want %+v", round, err, got)
	}
	crlf, err := c.ParseMarkdown(strings.ReplaceAll(text, "\n", "\r\n"))
	if err != nil || !reflect.DeepEqual(crlf, got) {
		t.Errorf("CRLF: %+v %v", crlf, err)
	}
	for _, bad := range []string{"# no frontmatter", "---\nname: auth", "---\n[one,two]\n---", "---\nname: [\n---", "---\nname: a\nname: b\n---", "---\n{}\n---\n## Notes\na\n## Notes\nb", "---\nnull\n---"} {
		if _, err := c.ParseMarkdown(bad); err == nil {
			t.Errorf("accepted bad Markdown %q", bad)
		}
	}
}

func Test_CON013ModuleInheritance(t *testing.T) {
	defaults := c.Zones{Interface: []string{"api.go"}, Evals: []string{"evals/**"}, Disposable: []string{"internal/**"}, Preserve: []string{"kept/**"}, AppendOnly: []string{"migration/**"}}
	hooks := c.Hooks{Eval: "test {report}", Boundaries: "bound", Setup: "setup"}
	got, err := c.ParseModule(moduleDoc("depends_on: [clock]\nboundaries_opt_out: 'legacy import'\n", "### AUTH-001\nIt MUST reject login."), "packages/auth", defaults, hooks)
	if err != nil {
		t.Fatal(err)
	}
	if got.Name != "auth" || got.IDPrefix != "AUTH" || got.Dir != "packages/auth" || !reflect.DeepEqual(got.Zones, defaults) || got.Hooks != hooks || !reflect.DeepEqual(got.DependsOn, []string{"clock"}) || got.BoundariesOptOut != "legacy import" || got.Decisions != "Keep errors stable." || len(got.Specs) != 1 || got.Specs[0].ID != "AUTH-001" {
		t.Fatalf("module %+v", got)
	}
	overrides := "interface: []\nevals: []\ndisposable: []\npreserve: []\nappend_only: []\nhooks: {eval: '', boundaries: '', setup: ''}\n"
	empty, err := c.ParseModule(moduleDoc(overrides, ""), "packages/auth", defaults, hooks)
	if err != nil {
		t.Fatal(err)
	}
	for key, list := range map[string][]string{"interface": empty.Interface, "evals": empty.Evals, "disposable": empty.Disposable, "preserve": empty.Preserve, "append_only": empty.AppendOnly} {
		if list == nil || len(list) != 0 {
			t.Errorf("explicit empty %s: %#v", key, list)
		}
	}
	if empty.Hooks != (c.Hooks{}) || len(empty.DependsOn) != 0 {
		t.Errorf("empty override/default deps %+v", empty)
	}
	partial, err := c.ParseModule(moduleDoc("hooks: {eval: custom}\n", ""), "packages/auth", defaults, hooks)
	if err != nil || partial.Hooks != (c.Hooks{Eval: "custom", Boundaries: "bound", Setup: "setup"}) {
		t.Errorf("partial hooks %+v %v", partial, err)
	}
	other, err := c.ParseModule(moduleDoc("", ""), "packages/auth", defaults, hooks)
	if err != nil {
		t.Fatal(err)
	}
	got.Interface[0] = "changed"
	got.Evals[0] = "changed"
	got.Disposable[0] = "changed"
	got.Preserve[0] = "changed"
	got.AppendOnly[0] = "changed"
	if defaults.Interface[0] != "api.go" || other.Interface[0] != "api.go" || defaults.Evals[0] != "evals/**" || other.Evals[0] != "evals/**" || defaults.Disposable[0] != "internal/**" || other.Disposable[0] != "internal/**" || defaults.Preserve[0] != "kept/**" || other.Preserve[0] != "kept/**" || defaults.AppendOnly[0] != "migration/**" || other.AppendOnly[0] != "migration/**" {
		t.Error("inherited slices shared")
	}
	for _, bad := range []string{
		strings.Replace(moduleDoc("", ""), "name: auth", "name: Bad", 1), strings.Replace(moduleDoc("", ""), "id_prefix: AUTH", "id_prefix: 123", 1),
		strings.Replace(moduleDoc("", ""), "name: auth\n", "", 1), strings.Replace(moduleDoc("", ""), "id_prefix: AUTH\n", "", 1),
		strings.Replace(moduleDoc("", ""), "## Specs", "## Other", 1), moduleDoc("unknown: true\n", ""), moduleDoc("depends_on: clock\n", ""),
		moduleDoc("hooks: {eval: 123}\n", ""), moduleDoc("interface: null\n", ""), moduleDoc("", "### OTHER-001\nIt MUST work."), moduleDoc("", "### AUTH-001\nNo normative text."),
	} {
		if _, err := c.ParseModule(bad, "packages/auth", defaults, hooks); err == nil {
			t.Errorf("invalid module accepted %q", bad)
		}
	}
}

func Test_CON014ChangeMetadata(t *testing.T) {
	text := changeDoc("run: abc\nphase: impl\nblocked: waiting\nbases: {AUTH-001: abcdef12}\nfixes: [AUTH-002]\nalready_satisfied: [AUTH-003]\nregenerate: true\n", "## Notes\nHuman notes.")
	got, err := c.ParseChange(text, "changes/0042-some-change.md", "AUTH")
	if err != nil {
		t.Fatal(err)
	}
	if got.Module != "auth" || got.Path != "changes/0042-some-change.md" || got.Slug != "some-change" || got.Number != 42 || got.Run != "abc" || got.Phase != "impl" || got.Blocked != "waiting" || got.Bases["AUTH-001"] != "abcdef12" || !reflect.DeepEqual(got.Fixes, []string{"AUTH-002"}) || !reflect.DeepEqual(got.AlreadySatisfied, []string{"AUTH-003"}) || !got.Regenerate || got.Kind != "regenerate" {
		t.Errorf("metadata %+v", got)
	}
	for _, tc := range []struct{ meta, body, kind string }{
		{"", "", "maintenance"}, {"fixes: [AUTH-001]\n", "", "fix"},
		{"", "## Added\n### new:one\nIt MUST work.", "spec"},
		{"", "## Modified\n### AUTH-001\nIt MUST work.", "spec"},
		{"", "## Removed\n### AUTH-001", "spec"},
		{"regenerate: true\n", "## Added\n### new:one\nIt MUST work.", "regenerate"},
	} {
		if got, err := c.ParseChange(changeDoc(tc.meta, tc.body), "changes/archive/0001-case.md", "AUTH"); err != nil || got.Kind != tc.kind {
			t.Errorf("kind %s: %+v %v", tc.kind, got, err)
		}
	}
	for _, phase := range []string{"evals", "audit", "impl", "applied"} {
		if _, err := c.ParseChange(changeDoc("phase: "+phase+"\n", ""), "changes/0001-case.md", "AUTH"); err != nil {
			t.Errorf("valid phase %s: %v", phase, err)
		}
	}
	for _, path := range []string{"0001-case.md", "changes/1-case.md", "changes/0001-Bad.md", "changes/0001-../case.md", "elsewhere/0001-case.md", "changes/0001-.md", "changes/0001-case.txt"} {
		if _, err := c.ParseChange(changeDoc("", ""), path, "AUTH"); err == nil {
			t.Errorf("bad path accepted %q", path)
		}
	}
	for _, text := range []string{
		changeDoc("phase: queued\n", ""), changeDoc("run: []\n", ""), changeDoc("fixes: invalid\n", ""), changeDoc("blocked: []\n", ""), changeDoc("regenerate: nonsense\n", ""), changeDoc("unknown: true\n", ""),
		strings.Replace(changeDoc("", ""), "module: auth\n", "", 1), strings.Replace(changeDoc("", ""), "Protect accounts.", "", 1),
	} {
		if _, err := c.ParseChange(text, "changes/0001-case.md", "AUTH"); err == nil {
			t.Errorf("invalid change accepted %q", text)
		}
	}
}

func Test_CON015ChangeSections(t *testing.T) {
	body := "## Interface\nExpose new:one.\n## Added\n### new:one\nIt MUST work.\n## Modified\n### AUTH-001\nIt SHOULD be clear.\n## Removed\n### AUTH-002\n## Decisions\nUse UTC.\n## Removed decisions\nOld behavior.\n## Questions\nWhy?\n## Audit\nNo findings.\n## Provenance\nmanual\n## Extra prose\nRetain this."
	got, err := c.ParseChange(changeDoc("", body), "changes/0001-case.md", "AUTH")
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Added) != 1 || got.Added[0].ID != "new:one" || len(got.Modified) != 1 || got.Modified[0].Body != "It SHOULD be clear." || len(got.Removed) != 1 || got.Removed[0].ID != "AUTH-002" || got.Removed[0].Body != "" {
		t.Errorf("edits %+v", got)
	}
	content := map[string]string{}
	for _, s := range got.Document.Sections {
		content[s.Name] = s.Body
	}
	for name, want := range map[string]string{"Motivation": "Protect accounts.", "Interface": "Expose new:one.", "Decisions": "Use UTC.", "Removed decisions": "Old behavior.", "Questions": "Why?", "Audit": "No findings.", "Provenance": "manual", "Extra prose": "Retain this."} {
		if content[name] != want {
			t.Errorf("section %s=%q, want %q", name, content[name], want)
		}
	}
	for _, bad := range []string{
		"## Added\n### new:one\nNot normative.", "## Modified\n### AUTH-001\n", "## Removed\n### AUTH-001\nForbidden body.",
		"## Added\n### new:one\nIt MUST work.\n### new:one\nIt MUST work.", "## Modified\n### new:one\nIt MUST work.", "## Removed\n### new:one",
		"## Interface\nExpose new:missing.", "## Added\n### AUTH-001\nIt MUST work.\n## Modified\n### AUTH-001\nIt MUST work.",
		"## Modified\n### AUTH-001\nIt MUST work.\n## Removed\n### AUTH-001",
	} {
		if _, err := c.ParseChange(changeDoc("", bad), "changes/0001-case.md", "AUTH"); err == nil {
			t.Errorf("invalid edits accepted %q", bad)
		}
	}
	if _, err := c.ParseChange(changeDoc("", "## Added\n## Modified\n## Removed\n"), "changes/0001-case.md", "AUTH"); err != nil {
		t.Errorf("empty sections: %v", err)
	}
}

func Test_CON016SetFields(t *testing.T) {
	meta := "set:\n  depends_on: [clock]\n  interface: []\n  evals: []\n  disposable: []\n  preserve: []\n  append_only: []\n  boundaries_opt_out: legacy\n  hooks: {eval: '', boundaries: custom, setup: ''}\n"
	got, err := c.ParseChange(changeDoc(meta, ""), "changes/0001-case.md", "AUTH")
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Set) != 8 || got.Set["boundaries_opt_out"] != "legacy" {
		t.Errorf("set %+v", got.Set)
	}
	for _, key := range []string{"interface", "evals", "disposable", "preserve", "append_only"} {
		if value, ok := got.Set[key].([]any); !ok || len(value) != 0 {
			t.Errorf("empty set %s: %#v", key, got.Set[key])
		}
	}
	if hooks, ok := got.Set["hooks"].(map[string]any); !ok || hooks["eval"] != "" || hooks["boundaries"] != "custom" || hooks["setup"] != "" {
		t.Errorf("hooks %+v", got.Set["hooks"])
	}
	for _, bad := range []string{"set: {name: other}\n", "set: {id_prefix: OTHER}\n", "set: {unknown: []}\n", "set: {depends_on: nope}\n", "set: {interface: false}\n", "set: {hooks: {eval: 123}}\n", "set: {hooks: {unknown: x}}\n", "set: {boundaries_opt_out: []}\n", "bases: {AUTH-001: ABCDEF12}\n", "bases: {AUTH-001: abc}\n", "bases: {new:one: abcdef12}\n", "fixes: [new:one]\n", "already_satisfied: [new:one]\n"} {
		if _, err := c.ParseChange(changeDoc(bad, ""), "changes/0001-case.md", "AUTH"); err == nil {
			t.Errorf("invalid metadata accepted %q", bad)
		}
	}
}

func Test_CON012IndependentRoundTrip(t *testing.T) {
	document := c.Document{Metadata: map[string]any{"name": "auth", "nested": map[string]any{"enabled": true, "list": []any{"one", "two"}}, "note": "## metadata text\n"}, Preamble: "# Auth", Sections: []c.Section{{Name: "Notes", Body: "Keep this."}}}
	text, err := c.RenderMarkdown(document)
	if err != nil {
		t.Fatal(err)
	}
	got, err := c.ParseMarkdown(text)
	if err != nil || !reflect.DeepEqual(got, document) {
		t.Errorf("independent round trip: %+v %v, want %+v", got, err, document)
	}
}
func Test_CON013TypedMetadata(t *testing.T) {
	for _, text := range []string{
		strings.Replace(moduleDoc("", ""), "id_prefix: AUTH", "id_prefix: auth", 1),
		moduleDoc("interface: [123]\n", ""), moduleDoc("depends_on: [123]\n", ""),
	} {
		if _, err := c.ParseModule(text, "packages/auth", c.Zones{}, c.Hooks{}); err == nil {
			t.Errorf("malformed module metadata accepted %q", text)
		}
	}
}
func Test_CON014ClassificationAndRequiredValues(t *testing.T) {
	for _, value := range []string{"''", "null"} {
		text := strings.Replace(changeDoc("", ""), "module: auth", "module: "+value, 1)
		if _, err := c.ParseChange(text, "changes/0001-case.md", "AUTH"); err == nil {
			t.Errorf("required module accepted %s", value)
		}
	}
	got, err := c.ParseChange(changeDoc("fixes: [AUTH-001]\n", "## Added\n### new:one\nIt MUST work."), "changes/10000-priority.md", "AUTH")
	if err != nil || got.Kind != "spec" || got.Number != 10000 {
		t.Errorf("spec beats fix: %+v %v", got, err)
	}
	got, err = c.ParseChange(changeDoc("regenerate: true\n", "## Added\n### new:one\nIt MUST work."), "changes/0001-case.md", "AUTH")
	if err != nil || got.Kind != "regenerate" || len(got.Added) != 1 || got.Added[0].ID != "new:one" {
		t.Errorf("regenerate retained edits: %+v %v", got, err)
	}
}
func Test_CON015ProseAndHandleReferences(t *testing.T) {
	for _, body := range []string{
		"## Modified\n### AUTH-001\nThis has no normative language.",
		"## Added\n### new:one\nIt MUST call new:missing.",
		"## Modified\n### AUTH-001\nIt MUST call new:missing.",
	} {
		if _, err := c.ParseChange(changeDoc("", body), "changes/0001-case.md", "AUTH"); err == nil {
			t.Errorf("invalid body accepted %q", body)
		}
	}
	text := strings.Replace(changeDoc("", "## Added\nIntroductory author prose.\n### new:one\nIt MUST work."), "## Motivation", "# Author title\n\n## Motivation", 1)
	got, err := c.ParseChange(text, "changes/0001-case.md", "AUTH")
	if err != nil {
		t.Fatal(err)
	}
	if got.Document.Preamble != "# Author title" {
		t.Errorf("preamble lost: %q", got.Document.Preamble)
	}
	found := false
	for _, section := range got.Document.Sections {
		if section.Name == "Added" {
			found = strings.HasPrefix(section.Body, "Introductory author prose.\n")
		}
	}
	if !found {
		t.Error("prose before spec item lost")
	}
}
func Test_CON016SetAndIDValidation(t *testing.T) {
	got, err := c.ParseChange(changeDoc("set: {depends_on: []}\n", ""), "changes/0001-case.md", "AUTH")
	if err != nil {
		t.Fatal(err)
	}
	if list, ok := got.Set["depends_on"].([]any); !ok || len(list) != 0 {
		t.Errorf("explicit empty dependencies lost: %#v", got.Set)
	}
	for _, meta := range []string{"bases: {AUTH-001: gggggggg}\n", "bases: {invalid: abcdef12}\n", "fixes: [invalid]\n", "already_satisfied: [invalid]\n"} {
		if _, err := c.ParseChange(changeDoc(meta, ""), "changes/0001-case.md", "AUTH"); err == nil {
			t.Errorf("invalid hash or ID accepted %q", meta)
		}
	}
}
