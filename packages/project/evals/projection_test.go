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
