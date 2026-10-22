package evals_test

import (
	"encoding/json"
	"fmt"
	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
	"reflect"
	"strings"
	"testing"
)

func itemBody(items []c.Spec) string {
	var b strings.Builder
	for _, v := range items {
		fmt.Fprintf(&b, "### %s", v.ID)
		if v.Tag != "" {
			fmt.Fprintf(&b, " %s", v.Tag)
		}
		fmt.Fprintf(&b, "\n\n%s\n\n", v.Body)
	}
	return strings.TrimSpace(b.String())
}
func projectionModule(t *testing.T, name, prefix, dir string) c.Module {
	t.Helper()
	text := module(name, prefix, "depends_on: [helper]\n")
	text = strings.Replace(text, "## Decisions", "### "+prefix+"-002 [quarantine: old clock]\nThe second MUST retain its body.\n\n### "+prefix+"-003\nThe third SHOULD remain.\n\n## Decisions", 1) + "\n## Human notes\n\nKeep unrelated notes.\n"
	m, err := c.ParseModule(text, dir, c.Zones{Interface: []string{"api.go"}, Evals: []string{"evals/**"}, Disposable: []string{"internal/**"}, Preserve: []string{"data/**"}, AppendOnly: []string{"history/**"}}, c.Hooks{Eval: "default eval", Boundaries: "default boundaries", Setup: "default setup"})
	if err != nil {
		t.Fatal(err)
	}
	return m
}
func projectionChange(t *testing.T, m c.Module, number int, slug, meta string, added, modified, removed []c.Spec) c.Change {
	t.Helper()
	text := "---\nmodule: " + m.Name + "\n" + meta + "---\n\n## Motivation\n\nA projection fixture needs this change.\n\n## Interface\n\nRetain this author interface prose.\n\n## Added\n\n" + itemBody(added) + "\n\n## Modified\n\n" + itemBody(modified) + "\n\n## Removed\n\n" + itemBody(removed) + "\n\n## Notes\n\nKeep change notes.\n"
	ch, err := c.ParseChange(text, fmt.Sprintf("changes/%04d-%s.md", number, slug), m.IDPrefix)
	if err != nil {
		t.Fatalf("invalid fixture: %v\n%s", err, text)
	}
	return ch
}
func encoded(t *testing.T, v any) string {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}
func Test_PRJ007ApplySpecs(t *testing.T) {
	m := projectionModule(t, "alpha", "ALP", "units/a")
	added := []c.Spec{{ID: "new:first", Tag: "[quarantine: clock]", Body: "The first MUST refer to new:second.\n\n```text\n## Example heading\n```"}, {ID: "new:second", Body: "The second SHOULD refer to new:first."}}
	modified := []c.Spec{{ID: "ALP-002", Tag: "[review]", Body: "The modified MUST use new:first."}}
	ch := projectionChange(t, m, 1, "apply", "", added, modified, []c.Spec{{ID: "ALP-001"}})
	before := encoded(t, []any{m, ch})
	out, err := p.ApplyChange(m, ch)
	if err != nil {
		t.Fatal(err)
	}
	want := []c.Spec{modified[0], m.Specs[2], added[0], added[1]}
	if !reflect.DeepEqual(out.Specs, want) {
		t.Fatalf("spec application\n%+v\n%+v", out.Specs, want)
	}
	if out.Document.Preamble != m.Document.Preamble || out.Decisions != m.Decisions {
		t.Fatal("lost preamble/Decisions")
	}
	for _, s := range m.Document.Sections {
		if s.Name == "Specs" {
			continue
		}
		found := false
		for _, v := range out.Document.Sections {
			if v == s {
				found = true
			}
		}
		if !found {
			t.Fatalf("lost unrelated section %+v", s)
		}
	}
	specsBody := ""
	for _, s := range out.Document.Sections {
		if s.Name == "Specs" {
			specsBody = s.Body
		}
	}
	items, err := c.ParseItems(specsBody, "ALP", true)
	if err != nil || !reflect.DeepEqual(items, want) {
		t.Fatalf("document disagrees %+v %v", items, err)
	}
	if encoded(t, []any{m, ch}) != before {
		t.Fatal("mutated Apply inputs")
	}
}
func Test_PRJ007ApplyRejectsInvalid(t *testing.T) {
	m := projectionModule(t, "alpha", "ALP", "units/a")
	valid := projectionChange(t, m, 1, "edit", "", nil, []c.Spec{{ID: "ALP-002", Body: "It MUST change."}}, nil)
	cases := []c.Change{}
	add := func(edit func(*c.Change)) {
		v := valid
		v.Added = append([]c.Spec{}, valid.Added...)
		v.Modified = append([]c.Spec{}, valid.Modified...)
		v.Removed = append([]c.Spec{}, valid.Removed...)
		edit(&v)
		cases = append(cases, v)
	}
	add(func(v *c.Change) { v.Module = "other" })
	add(func(v *c.Change) { v.Modified[0].ID = "ALP-099" })
	add(func(v *c.Change) { v.Removed = []c.Spec{{ID: "ALP-099"}} })
	add(func(v *c.Change) { v.Added = []c.Spec{{ID: "ALP-001", Body: "It MUST collide."}} })
	add(func(v *c.Change) { v.Added = []c.Spec{{ID: "WRONG-001", Body: "It MUST fail prefix."}} })
	add(func(v *c.Change) { v.Modified[0].Body = "No normative content." })
	add(func(v *c.Change) { v.Modified[0].Tag = "[unknown]" })
	add(func(v *c.Change) { v.Removed = []c.Spec{{ID: "ALP-001", Body: "Forbidden removal body."}} })
	add(func(v *c.Change) { v.Modified = append(v.Modified, v.Modified[0]) })
	add(func(v *c.Change) { v.Removed = []c.Spec{{ID: "ALP-002"}} })
	add(func(v *c.Change) {
		v.Added = []c.Spec{{ID: "new:one", Body: "It MUST add."}, {ID: "new:one", Body: "It MUST duplicate."}}
	})
	add(func(v *c.Change) { v.Added = []c.Spec{{ID: "new:one", Body: "It MUST refer to new:elsewhere."}} })
	add(func(v *c.Change) { v.Modified[0].ID = "new:bad" })
	for i, ch := range cases {
		before := encoded(t, []any{m, ch})
		out, err := p.ApplyChange(m, ch)
		if err == nil || err.Error() == "" || !reflect.DeepEqual(out, c.Module{}) {
			t.Fatalf("case %d partial result/error %+v %v", i, out, err)
		}
		if before != encoded(t, []any{m, ch}) {
			t.Fatalf("invalid case %d mutated inputs", i)
		}
	}
	out, err := p.ApplyChange(c.Module{Name: m.Name, IDPrefix: m.IDPrefix, Dir: m.Dir, Specs: m.Specs}, projectionChange(t, m, 2, "minimal", "", nil, nil, nil))
	if err != nil || len(out.Specs) != 3 || out.Document.Metadata["name"] != m.Name {
		t.Fatalf("minimal synthesized document %+v %v", out, err)
	}
}
func Test_PRJ008ProjectedMetadata(t *testing.T) {
	m := projectionModule(t, "alpha", "ALP", "units/a")
	meta := "set:\n  depends_on: [other, helper]\n  interface: []\n  evals: []\n  disposable: []\n  preserve: []\n  append_only: []\n  boundaries_opt_out: documented reason\n  hooks: {eval: '', setup: custom setup}\n"
	ch := projectionChange(t, m, 1, "zones", meta, nil, nil, nil)
	out, err := p.ApplyChange(m, ch)
	if err != nil {
		t.Fatal(err)
	}
	if out.Name != m.Name || out.IDPrefix != m.IDPrefix || out.Dir != m.Dir || !reflect.DeepEqual(out.DependsOn, []string{"other", "helper"}) || out.BoundariesOptOut != "documented reason" || out.Hooks != (c.Hooks{Eval: "", Boundaries: "default boundaries", Setup: "custom setup"}) {
		t.Fatalf("metadata %+v", out)
	}
	for _, zs := range [][]string{out.Interface, out.Evals, out.Disposable, out.Preserve, out.AppendOnly} {
		if zs == nil || len(zs) != 0 {
			t.Fatalf("explicit empty lost %+v", out.Zones)
		}
	}
	raw, err := c.RenderMarkdown(out.Document)
	if err != nil {
		t.Fatal(err)
	}
	again, err := c.ParseModule(raw, out.Dir, c.Zones{}, c.Hooks{})
	if err != nil || !reflect.DeepEqual(again.Zones, out.Zones) || !reflect.DeepEqual(again.DependsOn, out.DependsOn) || again.Hooks != out.Hooks || again.BoundariesOptOut != out.BoundariesOptOut {
		t.Fatalf("document metadata disagrees\n%s\n%+v %v", raw, again, err)
	}
	partial := projectionChange(t, m, 2, "partial", "set: {hooks: {boundaries: ''}}\n", nil, nil, nil)
	same, err := p.ApplyChange(m, partial)
	if err != nil || same.Hooks != (c.Hooks{Eval: m.Hooks.Eval, Boundaries: "", Setup: m.Hooks.Setup}) || !reflect.DeepEqual(same.Zones, m.Zones) || !reflect.DeepEqual(same.DependsOn, m.DependsOn) {
		t.Fatalf("partial omitted values %+v %v", same, err)
	}
}
func Test_PRJ008RejectSetTypes(t *testing.T) {
	m := projectionModule(t, "alpha", "ALP", "units/a")
	ch := projectionChange(t, m, 1, "bad-set", "", nil, nil, nil)
	for _, set := range []map[string]any{{"name": "renamed"}, {"id_prefix": "NEW"}, {"dir": "elsewhere"}, {"unknown": []string{}}, {"depends_on": "other"}, {"interface": false}, {"evals": []any{1}}, {"disposable": "src/**"}, {"preserve": nil}, {"append_only": 3}, {"boundaries_opt_out": false}, {"hooks": "shell"}, {"hooks": map[string]any{"unknown": "command"}}, {"hooks": map[string]any{"eval": true}}} {
		ch.Set = set
		before := encoded(t, []any{m, ch})
		out, err := p.ApplyChange(m, ch)
		if err == nil || !reflect.DeepEqual(out, c.Module{}) {
			t.Fatalf("accepted set %+v: %+v %v", set, out, err)
		}
		if encoded(t, []any{m, ch}) != before {
			t.Fatal("invalid set mutated input")
		}
	}
}
func projectionFixture(t *testing.T) p.Snapshot {
	t.Helper()
	a := projectionModule(t, "alpha", "ALP", "units/a")
	b := projectionModule(t, "beta", "BET", "units/b")
	first := projectionChange(t, a, 2, "alpha", "set: {preserve: [first/**]}\n", nil, []c.Spec{{ID: "ALP-001", Body: "It MUST become first."}}, nil)
	second := projectionChange(t, a, 2, "zulu", "set: {preserve: [second/**]}\n", nil, []c.Spec{{ID: "ALP-001", Tag: "[review]", Body: "It MUST become second."}}, nil)
	last := projectionChange(t, b, 10, "last", "", []c.Spec{{ID: "new:last", Body: "It MUST be last."}}, nil, nil)
	archived := projectionChange(t, a, 1, "archived", "", nil, []c.Spec{{ID: "ALP-001", Body: "It MUST NOT replay archive."}}, nil)
	archived.Path = "changes/archive/0001-archived.md"
	return p.Snapshot{Root: "/does-not-exist/projection", BaseRef: "refs/heads/nonexistent", Config: c.Config{Modules: []string{"units/*"}, Protected: []string{"keep/**"}}, Modules: []c.Module{b, a}, Changes: []c.Change{last, second, first}, Archive: []c.Change{archived}, Diagnostics: []p.Diagnostic{{Code: "C02", Severity: "warning", Path: "outside/SPEC.md", Message: "Keep original diagnostic"}}}
}
func Test_PRJ009QueueOrderAndStops(t *testing.T) {
	s := projectionFixture(t)
	before := encoded(t, s)
	for _, v := range []struct {
		stop, body, zone string
		bet              int
	}{{"alpha", s.Modules[1].Specs[0].Body, "data/**", 3}, {"zulu", "It MUST become first.", "first/**", 3}, {"last", "It MUST become second.", "second/**", 3}, {"", "It MUST become second.", "second/**", 4}} {
		out, err := p.Project(s, v.stop)
		if err != nil {
			t.Fatal(err)
		}
		if out.Modules[1].Specs[0].Body != v.body || !reflect.DeepEqual(out.Modules[1].Preserve, []string{v.zone}) || len(out.Modules[0].Specs) != v.bet {
			t.Fatalf("stop %q: %+v", v.stop, out.Modules)
		}
		if out.Modules[0].Name != "beta" || out.Modules[1].Name != "alpha" || out.Root != s.Root || out.BaseRef != s.BaseRef || !reflect.DeepEqual(out.Config, s.Config) || !reflect.DeepEqual(out.Changes, s.Changes) || !reflect.DeepEqual(out.Archive, s.Archive) || !reflect.DeepEqual(out.Diagnostics, s.Diagnostics) {
			t.Fatal("changed unrelated snapshot data/order")
		}
	}
	if before != encoded(t, s) {
		t.Fatal("Project mutated input")
	}
}
func poison(v reflect.Value) {
	if !v.IsValid() {
		return
	}
	switch v.Kind() {
	case reflect.Interface:
		if !v.IsNil() {
			x := reflect.New(v.Elem().Type()).Elem()
			x.Set(v.Elem())
			poison(x)
			if v.CanSet() {
				v.Set(x)
			}
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			poison(v.Field(i))
		}
	case reflect.Map:
		if !v.IsNil() {
			for _, k := range v.MapKeys() {
				x := reflect.New(v.Type().Elem()).Elem()
				x.Set(v.MapIndex(k))
				poison(x)
				v.SetMapIndex(k, x)
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			poison(v.Index(i))
		}
	case reflect.String:
		if v.CanSet() {
			v.SetString("poisoned")
		}
	case reflect.Int:
		if v.CanSet() {
			v.SetInt(42)
		}
	case reflect.Bool:
		if v.CanSet() {
			v.SetBool(!v.Bool())
		}
	}
}
func Test_PRJ007_PRJ008_PRJ009Independence(t *testing.T) {
	s := projectionFixture(t)
	before := encoded(t, s)
	out, err := p.Project(s, "")
	if err != nil {
		t.Fatal(err)
	}
	poison(reflect.ValueOf(&out).Elem())
	if before != encoded(t, s) {
		t.Fatal("returned projected state aliases input")
	}
	m, ch := s.Modules[1], s.Changes[1]
	before = encoded(t, []any{m, ch})
	applied, err := p.ApplyChange(m, ch)
	if err != nil {
		t.Fatal(err)
	}
	poison(reflect.ValueOf(&applied).Elem())
	if before != encoded(t, []any{m, ch}) {
		t.Fatal("applied module aliases inputs")
	}
}
func Test_PRJ009RejectInvalidProjection(t *testing.T) {
	s := projectionFixture(t)
	for _, kind := range []string{"unknown-selector", "ambiguous-selector", "unknown-module", "invalid-application"} {
		v := s
		v.Changes = append([]c.Change{}, s.Changes...)
		selector := ""
		switch kind {
		case "unknown-selector":
			selector = "missing"
		case "ambiguous-selector":
			v.Changes[1].Slug = v.Changes[2].Slug
			selector = v.Changes[2].Slug
		case "unknown-module":
			v.Changes[2].Module = "missing"
		case "invalid-application":
			v.Changes[2].Modified = []c.Spec{{ID: "ALP-999", Body: "It MUST be absent."}}
		}
		before := encoded(t, v)
		out, err := p.Project(v, selector)
		if err == nil || err.Error() == "" {
			t.Fatalf("accepted %s %+v", kind, out)
		}
		poison(reflect.ValueOf(&out).Elem())
		if encoded(t, v) != before {
			t.Fatalf("failed %s aliases/mutates inputs", kind)
		}
	}
}
func Test_PRJ010ProjectedBases(t *testing.T) {
	s := projectionFixture(t)
	target := s.Changes[1]
	target.Modified = []c.Spec{{ID: "ALP-001", Body: "It MUST become final."}}
	target.Removed = []c.Spec{{ID: "ALP-002"}}
	target.Bases = map[string]string{"ALP-001": "deadbeef", "ALP-002": "deadbeef", "ALP-999": "deadbeef"}
	s.Changes[1] = target
	before := encoded(t, s)
	got, err := p.BaseHashes(s, target.Slug)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{"ALP-001": c.SpecHash(c.Spec{ID: "ALP-001", Body: "It MUST become first."}), "ALP-002": c.SpecHash(s.Modules[1].Specs[1])}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("projected bases %+v want %+v", got, want)
	}
	got["ALP-001"] = "poisoned"
	again, err := p.BaseHashes(s, target.Slug)
	if err != nil || !reflect.DeepEqual(again, want) || encoded(t, s) != before {
		t.Fatal("base map aliases/mutates inputs")
	}
	for _, slug := range []string{"last"} {
		empty, err := p.BaseHashes(s, slug)
		if err != nil || empty == nil || len(empty) != 0 {
			t.Fatalf("added-only bases %+v %v", empty, err)
		}
	}
	maintenance := projectionChange(t, s.Modules[1], 20, "maintenance", "", nil, nil, nil)
	s.Changes = append(s.Changes, maintenance)
	empty, err := p.BaseHashes(s, "maintenance")
	if err != nil || empty == nil || len(empty) != 0 {
		t.Fatalf("maintenance %+v %v", empty, err)
	}
	tagged := s
	tagged.Modules = append([]c.Module{}, s.Modules...)
	tagged.Modules[1].Specs = append([]c.Spec{}, s.Modules[1].Specs...)
	tagged.Modules[1].Specs[1].Tag = "[review]"
	changed, err := p.BaseHashes(tagged, target.Slug)
	if err != nil || changed["ALP-002"] == want["ALP-002"] {
		t.Fatal("tag excluded from base hash")
	}
}
func Test_PRJ010BaseErrors(t *testing.T) {
	s := projectionFixture(t)
	for _, kind := range []string{"missing-change", "ambiguous-change", "missing-module", "missing-edited"} {
		v := s
		v.Changes = append([]c.Change{}, s.Changes...)
		slug := "zulu"
		switch kind {
		case "missing-change":
			slug = "missing"
		case "ambiguous-change":
			v.Changes[2].Slug = slug
		case "missing-module":
			v.Changes[1].Module = "missing"
		case "missing-edited":
			v.Changes[1].Removed = []c.Spec{{ID: "ALP-999"}}
		}
		before := encoded(t, v)
		got, err := p.BaseHashes(v, slug)
		if err == nil || err.Error() == "" {
			t.Fatalf("accepted %s %+v", kind, got)
		}
		if encoded(t, v) != before {
			t.Fatalf("error %s mutated inputs", kind)
		}
	}
}

func Test_PRJ007PositionsAndAssignedAdditions(t *testing.T) {
	m := projectionModule(t, "alpha", "ALP", "units/a")
	m.Specs = append(m.Specs, c.Spec{ID: "ALP-004", Body: "The fourth MUST stay."}, c.Spec{ID: "ALP-005", Body: "The fifth MUST stay."})
	mods := []c.Spec{{ID: "ALP-004", Tag: "[review]", Body: "Fourth MUST change in place."}, {ID: "ALP-002", Body: "Second MUST change in place."}}
	adds := []c.Spec{{ID: "ALP-009", Body: "Assigned MUST append first."}, {ID: "new:local", Body: "Local MUST append second."}}
	ch := projectionChange(t, m, 1, "assigned", "run: assigned-run\nphase: impl\n", adds, mods, nil)
	out, err := p.ApplyChange(m, ch)
	if err != nil {
		t.Fatal(err)
	}
	want := []c.Spec{m.Specs[0], mods[1], m.Specs[2], mods[0], m.Specs[4], adds[0], adds[1]}
	if !reflect.DeepEqual(out.Specs, want) {
		t.Fatalf("lost positions/assigned additions %+v", out.Specs)
	}
	for _, adds := range [][]c.Spec{{{ID: "new:valid", Body: "No normative words."}}, {{ID: "new:valid", Tag: "[unknown]", Body: "It MUST fail tag."}}, {{ID: "new:Bad", Body: "It MUST fail handle."}}, {{ID: "ALP-009", Body: "It MUST add."}, {ID: "ALP-009", Body: "It MUST duplicate."}}} {
		bad := ch
		bad.Added = adds
		got, err := p.ApplyChange(m, bad)
		if err == nil || !reflect.DeepEqual(got, c.Module{}) {
			t.Fatalf("accepted Added %+v %v", adds, err)
		}
	}
}
func Test_PRJ007_PRJ009LocalHandleScopes(t *testing.T) {
	m := projectionModule(t, "alpha", "ALP", "units/a")
	one := projectionChange(t, m, 1, "one", "", []c.Spec{{ID: "new:local", Body: "First MUST use new:local."}}, nil, nil)
	two := projectionChange(t, m, 2, "two", "", []c.Spec{{ID: "new:local", Body: "Second MUST use new:local."}}, nil, nil)
	s := p.Snapshot{Modules: []c.Module{m}, Changes: []c.Change{two, one}}
	out, err := p.Project(s, "")
	if err != nil {
		t.Fatal(err)
	}
	want := append(append([]c.Spec{}, m.Specs...), one.Added[0], two.Added[0])
	if !reflect.DeepEqual(out.Modules[0].Specs, want) {
		t.Fatalf("collapsed local scopes %+v", out.Modules[0].Specs)
	}
	bad := two
	bad.Added = []c.Spec{{ID: "new:other", Body: "Second MUST NOT borrow new:local from another change."}}
	s.Changes = []c.Change{one, bad}
	if _, err = p.Project(s, ""); err == nil {
		t.Fatal("accepted cross-change-only handle reference")
	}
}
func moduleFields(m c.Module) c.Module { m.Document = c.Document{}; return m }
func assertProjectedDocument(t *testing.T, m c.Module) {
	t.Helper()
	raw, err := c.RenderMarkdown(m.Document)
	if err != nil {
		t.Fatal(err)
	}
	again, err := c.ParseModule(raw, m.Dir, c.Zones{}, c.Hooks{})
	if err != nil || !reflect.DeepEqual(moduleFields(again), moduleFields(m)) {
		t.Fatalf("document/fields disagree\n%s\n%+v\n%+v %v", raw, again, m, err)
	}
}
func Test_PRJ008EverySetFieldAndDocument(t *testing.T) {
	m := projectionModule(t, "alpha", "ALP", "units/a")
	m.BoundariesOptOut = "existing reason"
	ch := projectionChange(t, m, 1, "matrix", "", nil, nil, nil)
	cases := []struct {
		key   string
		value any
		edit  func(*c.Module)
	}{{"depends_on", []string{}, func(v *c.Module) { v.DependsOn = []string{} }}, {"interface", []string{"pub/**/*.go"}, func(v *c.Module) { v.Interface = []string{"pub/**/*.go"} }}, {"evals", []string{"tests/**"}, func(v *c.Module) { v.Evals = []string{"tests/**"} }}, {"disposable", []string{"src/**"}, func(v *c.Module) { v.Disposable = []string{"src/**"} }}, {"preserve", []string{"keep/**"}, func(v *c.Module) { v.Preserve = []string{"keep/**"} }}, {"append_only", []string{"events/**"}, func(v *c.Module) { v.AppendOnly = []string{"events/**"} }}, {"boundaries_opt_out", "", func(v *c.Module) { v.BoundariesOptOut = "" }}, {"hooks", map[string]any{"eval": "custom eval", "boundaries": "custom boundaries", "setup": ""}, func(v *c.Module) { v.Hooks = c.Hooks{Eval: "custom eval", Boundaries: "custom boundaries", Setup: ""} }}}
	for _, v := range cases {
		ch.Set = map[string]any{v.key: v.value}
		out, err := p.ApplyChange(m, ch)
		if err != nil {
			t.Fatalf("set %s: %v", v.key, err)
		}
		want := m
		v.edit(&want)
		if !reflect.DeepEqual(moduleFields(out), moduleFields(want)) {
			t.Fatalf("set %s %+v want %+v", v.key, out, want)
		}
		assertProjectedDocument(t, out)
		if out.Document.Preamble != m.Document.Preamble {
			t.Fatal("set-only lost preamble")
		}
		for _, s := range m.Document.Sections {
			if s.Name == "Specs" {
				continue
			}
			found := false
			for _, v := range out.Document.Sections {
				if v == s {
					found = true
				}
			}
			if !found {
				t.Fatalf("set-only lost section %+v", s)
			}
		}
	}
	minimal := m
	minimal.Document = c.Document{}
	out, err := p.ApplyChange(minimal, projectionChange(t, m, 1, "synthesize", "", nil, nil, nil))
	if err != nil {
		t.Fatal(err)
	}
	assertProjectedDocument(t, out)
	if out.Name != m.Name || out.IDPrefix != m.IDPrefix || out.Dir != m.Dir || !reflect.DeepEqual(out.Specs, m.Specs) || !reflect.DeepEqual(out.Zones, m.Zones) || out.Hooks != m.Hooks {
		t.Fatal("synthesis changed typed values")
	}
}
func Test_PRJ008_PRJ009NoReinheritanceOrEmptyAlias(t *testing.T) {
	m := projectionModule(t, "alpha", "ALP", "units/a")
	m.Interface = []string{}
	m.Hooks.Eval = ""
	m.Preserve = make([]string, 2, 8)
	copy(m.Preserve, []string{"keep-a", "keep-b"})
	m.AppendOnly = make([]string, 2, 8)
	copy(m.AppendOnly, []string{"history-a", "history-b"})
	ch := projectionChange(t, m, 1, "empty", "set: {preserve: [], append_only: [], hooks: {setup: ''}}\n", nil, nil, nil)
	s := p.Snapshot{Modules: []c.Module{m}, Changes: []c.Change{ch}, Config: c.Config{ModuleDefaults: c.Zones{Interface: []string{"WRONG/**"}, Evals: []string{"WRONG/**"}, Disposable: []string{"WRONG/**"}, Preserve: []string{"WRONG/**"}, AppendOnly: []string{"WRONG/**"}}, Hooks: c.Hooks{Eval: "WRONG", Boundaries: "WRONG", Setup: "WRONG"}}}
	before := encoded(t, s)
	out, err := p.Project(s, "")
	if err != nil {
		t.Fatal(err)
	}
	v := out.Modules[0]
	if v.Interface == nil || len(v.Interface) != 0 || v.Preserve == nil || len(v.Preserve) != 0 || v.AppendOnly == nil || len(v.AppendOnly) != 0 || v.Hooks.Eval != "" || v.Hooks.Setup != "" || v.Hooks.Boundaries != m.Hooks.Boundaries || !reflect.DeepEqual(v.Evals, m.Evals) || !reflect.DeepEqual(v.Disposable, m.Disposable) {
		t.Fatalf("re-inherited defaults %+v", v)
	}
	assertProjectedDocument(t, v)
	v.Preserve = append(v.Preserve, "MUTATED")
	v.AppendOnly = append(v.AppendOnly, "MUTATED")
	v.Interface = append(v.Interface, "MUTATED")
	if before != encoded(t, s) {
		t.Fatal("cleared slice aliases original capacity")
	}
}
func Test_PRJ009TypedNumbersAndExcludedTail(t *testing.T) {
	s := projectionFixture(t)
	s.Changes = append([]c.Change{}, s.Changes...)
	first := s.Changes[2]
	last := s.Changes[1]
	first.Number = 20
	last.Number = 1
	s.Changes = []c.Change{first, last}
	out, err := p.Project(s, "")
	if err != nil || out.Modules[1].Specs[0].Body != "It MUST become first." {
		t.Fatalf("sorted Path instead of typed Number %+v %v", out, err)
	}
	s = projectionFixture(t)
	s.Changes = append([]c.Change{}, s.Changes...)
	s.Changes[1].Modified = []c.Spec{{ID: "ALP-999", Body: "It MUST NOT apply at selector."}}
	s.Changes[0].Module = "missing-tail"
	out, err = p.Project(s, "zulu")
	if err != nil || out.Modules[1].Specs[0].Body != "It MUST become first." {
		t.Fatalf("evaluated excluded target/tail %+v %v", out, err)
	}
}
func Test_PRJ010NormalizedAndFreshEmptyBases(t *testing.T) {
	s := projectionFixture(t)
	s.Modules = append([]c.Module{}, s.Modules...)
	s.Modules[1].Specs = append([]c.Spec{}, s.Modules[1].Specs...)
	raw := s.Modules[1].Specs[1]
	raw.Body = "\r\nThe second MUST retain its body. \t\r\n\r\n"
	s.Modules[1].Specs[1] = raw
	s.Changes = append([]c.Change{}, s.Changes...)
	s.Changes[1].Removed = []c.Spec{{ID: "ALP-002"}}
	got, err := p.BaseHashes(s, "zulu")
	normalized := c.Spec{ID: raw.ID, Tag: raw.Tag, Body: "The second MUST retain its body."}
	if err != nil || got[raw.ID] != c.SpecHash(raw) || got[raw.ID] != c.SpecHash(normalized) || len(got[raw.ID]) != 8 {
		t.Fatalf("normalization %+v %v", got, err)
	}
	maintenance := projectionChange(t, s.Modules[1], 20, "maintenance", "", nil, nil, nil)
	s.Changes = append(s.Changes, maintenance)
	for _, slug := range []string{"last", "maintenance"} {
		empty, err := p.BaseHashes(s, slug)
		if err != nil || empty == nil || len(empty) != 0 {
			t.Fatalf("empty %s %+v %v", slug, empty, err)
		}
		empty["poison"] = "poison"
		again, err := p.BaseHashes(s, slug)
		if err != nil || again == nil || len(again) != 0 {
			t.Fatalf("shared empty map %s %+v %v", slug, again, err)
		}
	}
	s.Changes[1].Modified = []c.Spec{{ID: "ALP-999", Body: "It MUST be missing."}}
	if _, err = p.BaseHashes(s, "zulu"); err == nil {
		t.Fatal("accepted missing Modified base")
	}
}
