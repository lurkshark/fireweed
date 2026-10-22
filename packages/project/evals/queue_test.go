package evals_test

import (
	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func queueModule(t *testing.T, name, prefix, dir string) c.Module {
	t.Helper()
	m := projectionModule(t, name, prefix, dir)
	m.DependsOn = []string{}
	return m
}
func queueBases(ch c.Change, m c.Module) c.Change {
	ch.Bases = map[string]string{}
	for _, group := range [][]c.Spec{ch.Modified, ch.Removed} {
		for _, edited := range group {
			for _, old := range m.Specs {
				if edited.ID == old.ID {
					ch.Bases[edited.ID] = c.SpecHash(old)
				}
			}
		}
	}
	return ch
}
func queueFixture(t *testing.T) p.Snapshot {
	t.Helper()
	m := queueModule(t, "alpha", "ALP", "units/a")
	one := queueBases(projectionChange(t, m, 2, "alpha", "", nil, []c.Spec{{ID: "ALP-001", Tag: "[review]", Body: "It MUST become first."}}, nil), m)
	two := projectionChange(t, m, 2, "zulu", "", nil, []c.Spec{{ID: "ALP-001", Body: "It MUST become second."}}, []c.Spec{{ID: "ALP-002"}})
	two.Bases = map[string]string{"ALP-001": c.SpecHash(one.Modified[0]), "ALP-002": c.SpecHash(m.Specs[1])}
	three := projectionChange(t, m, 10, "last", "", nil, []c.Spec{{ID: "ALP-001", Body: "It MUST become third."}}, nil)
	three.Bases = map[string]string{"ALP-001": c.SpecHash(two.Modified[0])}
	return p.Snapshot{Root: "/nonexistent/queue", BaseRef: "refs/heads/not-present", Modules: []c.Module{m}, Changes: []c.Change{three, two, one}}
}
func queueClean(t *testing.T, s p.Snapshot) {
	t.Helper()
	ds := p.CheckQueue(s)
	if ds == nil || len(ds) != 0 {
		t.Fatalf("unexpected queue diagnostics %+v", ds)
	}
}
func Test_PRJ011QueueBasesAndIntent(t *testing.T) {
	s := queueFixture(t)
	queueClean(t, s)
	for _, kind := range []string{"missing-first", "stale-first", "missing-removed", "missing-modified-target", "missing-removed-target"} {
		v := s
		v.Changes = append([]c.Change{}, s.Changes...)
		path := v.Changes[2].Path
		switch kind {
		case "missing-first":
			v.Changes[2].Bases = nil
		case "stale-first":
			v.Changes[2].Bases = map[string]string{"ALP-001": "deadbeef"}
		case "missing-removed":
			v.Changes[1].Bases = map[string]string{"ALP-001": s.Changes[1].Bases["ALP-001"]}
			path = v.Changes[1].Path
		case "missing-modified-target":
			v.Changes[1].Modified = []c.Spec{{ID: "ALP-999", Body: "It MUST be absent."}}
			path = v.Changes[1].Path
		case "missing-removed-target":
			v.Changes[1].Removed = []c.Spec{{ID: "ALP-999"}}
			path = v.Changes[1].Path
		}
		ds := p.CheckQueue(v)
		requireDiag(t, ds, "C10", "error", path)
		if kind == "missing-first" || kind == "stale-first" || kind == "missing-removed" {
			if len(ds) != 1 {
				t.Fatalf("lost stale but valid intent %s %+v", kind, ds)
			}
		}
	}
}
func Test_PRJ011TagsNormalizationAndTypedOrder(t *testing.T) {
	s := queueFixture(t)
	s.Modules = append([]c.Module{}, s.Modules...)
	s.Modules[0].Specs = append([]c.Spec{}, s.Modules[0].Specs...)
	s.Modules[0].Specs[0].Body = "\r\n" + s.Modules[0].Specs[0].Body + " \t\r\n\r\n"
	queueClean(t, s)
	s.Modules[0].Specs[0].Tag = "[review]"
	requireDiag(t, p.CheckQueue(s), "C10", "error", s.Changes[2].Path)
	s = queueFixture(t)
	s.Changes = append([]c.Change{}, s.Changes...)
	s.Changes[2].Path = "changes/0010-first-path.md"
	s.Changes[2].Number = 1
	s.Changes[1].Path = "changes/0001-second-path.md"
	s.Changes[1].Number = 2
	s.Changes[0].Number = 3
	queueClean(t, s)
}
func Test_PRJ012ProjectedEnforcement(t *testing.T) {
	m := queueModule(t, "alpha", "ALP", "units/a")
	cases := []struct {
		set                    map[string]any
		code, severity, detail string
	}{{map[string]any{"interface": []any{"["}}, "C03", "error", ""}, {map[string]any{"disposable": []any{"evals/**"}}, "C06", "error", ""}, {map[string]any{"depends_on": []any{"missing"}}, "C06", "error", ""}, {map[string]any{"hooks": map[string]any{"eval": ""}}, "C07", "error", ""}, {map[string]any{"hooks": map[string]any{"boundaries": ""}}, "C07", "error", ""}, {map[string]any{"hooks": map[string]any{"boundaries": ""}, "boundaries_opt_out": "visible queue reason"}, "D04", "warning", "visible queue reason"}}
	for _, v := range cases {
		ch := projectionChange(t, m, 1, "enforcement", "", nil, nil, nil)
		ch.Set = v.set
		s := p.Snapshot{Modules: []c.Module{m}, Changes: []c.Change{ch}}
		ds := p.CheckQueue(s)
		requireDiag(t, ds, v.code, v.severity, ch.Path)
		if v.detail != "" {
			for _, d := range ds {
				if d.Code == v.code && !strings.Contains(d.Message, v.detail) {
					t.Fatalf("hidden optout %+v", ds)
				}
				if d.Severity == "error" {
					t.Fatalf("optout rejected %+v", ds)
				}
			}
		}
	}
	b := queueModule(t, "beta", "BET", "units/b")
	b.DependsOn = []string{"alpha"}
	ch := projectionChange(t, m, 1, "cycle", "set: {depends_on: [beta]}\n", nil, nil, nil)
	requireDiag(t, p.CheckQueue(p.Snapshot{Modules: []c.Module{m, b}, Changes: []c.Change{ch}}), "C06", "error", ch.Path)
}
func Test_PRJ012IntermediateStateAndPreexisting(t *testing.T) {
	m := queueModule(t, "alpha", "ALP", "units/a")
	bad := projectionChange(t, m, 1, "bad", "set: {evals: [internal/**]}\n", nil, nil, nil)
	repair := projectionChange(t, m, 2, "repair", "set: {evals: [evals/**]}\n", nil, nil, nil)
	ds := p.CheckQueue(p.Snapshot{Modules: []c.Module{m}, Changes: []c.Change{repair, bad}})
	requireDiag(t, ds, "C06", "error", bad.Path)
	if len(ds) != 1 {
		t.Fatalf("repair repeated existing error %+v", ds)
	}
	m.Hooks.Eval = ""
	b := queueModule(t, "beta", "BET", "units/b")
	maintenance := projectionChange(t, b, 1, "other", "", nil, nil, nil)
	queueClean(t, p.Snapshot{Modules: []c.Module{m, b}, Changes: []c.Change{maintenance}, Diagnostics: []p.Diagnostic{{Code: "C02", Severity: "warning", Message: "Preexisting discovery issue", Path: "outside/SPEC.md"}}})
}
func Test_PRJ013AssignedIDsAndReservations(t *testing.T) {
	m := queueModule(t, "alpha", "ALP", "units/a")
	ch := projectionChange(t, m, 1, "ids", "", []c.Spec{{ID: "ALP-009", Body: "It MUST get an assigned ID."}}, nil, nil)
	for _, state := range []struct{ run, phase string }{{"", ""}, {"run", ""}, {"", "evals"}} {
		v := ch
		v.Run = state.run
		v.Phase = state.phase
		requireDiag(t, p.CheckQueue(p.Snapshot{Modules: []c.Module{m}, Changes: []c.Change{v}}), "C05", "error", v.Path)
	}
	ch.Run = "run"
	ch.Phase = "evals"
	queueClean(t, p.Snapshot{Modules: []c.Module{m}, Changes: []c.Change{ch}})
	archived := projectionChange(t, m, 0, "reservation", "", nil, nil, []c.Spec{{ID: "ALP-009"}})
	archived.Path = "changes/archive/0000-reservation.md"
	for _, assigned := range []bool{false, true} {
		v := ch
		if !assigned {
			v.Run = ""
			v.Phase = ""
		}
		ds := p.CheckQueue(p.Snapshot{Modules: []c.Module{m}, Changes: []c.Change{v}, Archive: []c.Change{archived}})
		requireDiag(t, ds, "C08", "error", v.Path)
		if !assigned {
			requireDiag(t, ds, "C05", "error", v.Path)
		}
	}
	local := projectionChange(t, m, 1, "local", "", []c.Spec{{ID: "new:local", Body: "It MUST remain local."}}, nil, nil)
	queueClean(t, p.Snapshot{Modules: []c.Module{m}, Changes: []c.Change{local}})
}
func Test_PRJ013ReferencesAndIndependentErrors(t *testing.T) {
	m := queueModule(t, "alpha", "ALP", "units/a")
	valid := projectionChange(t, m, 1, "references", "", nil, nil, nil)
	valid.Fixes = []string{"ALP-001"}
	valid.AlreadySatisfied = []string{"ALP-001"}
	queueClean(t, p.Snapshot{Modules: []c.Module{m}, Changes: []c.Change{valid}})
	for _, kind := range []string{"missing-fix", "missing-satisfied", "removed-fix", "removed-satisfied", "new-fix"} {
		ch := valid
		ch.Removed = nil
		ch.Fixes = append([]string{}, valid.Fixes...)
		ch.AlreadySatisfied = append([]string{}, valid.AlreadySatisfied...)
		switch kind {
		case "missing-fix":
			ch.Fixes = []string{"ALP-999"}
		case "missing-satisfied":
			ch.AlreadySatisfied = []string{"ALP-999"}
		case "removed-fix", "removed-satisfied":
			ch.Removed = []c.Spec{{ID: "ALP-001"}}
			ch = queueBases(ch, m)
			if kind == "removed-fix" {
				ch.AlreadySatisfied = nil
			} else {
				ch.Fixes = nil
			}
		case "new-fix":
			ch.Added = []c.Spec{{ID: "ALP-009", Body: "It MUST be newly added."}}
			ch.Run = "run"
			ch.Phase = "evals"
			ch.Fixes = []string{"ALP-009"}
		}
		requireDiag(t, p.CheckQueue(p.Snapshot{Modules: []c.Module{m}, Changes: []c.Change{ch}}), "C05", "error", ch.Path)
	}
	b := queueModule(t, "beta", "BET", "units/b")
	for _, badKind := range []string{"target", "application"} {
		bad := projectionChange(t, m, 1, "broken", "", nil, nil, nil)
		if badKind == "target" {
			bad.Module = "unknown"
		} else {
			bad.Added = []c.Spec{{ID: "new:bad", Body: "No normative words."}}
		}
		independent := projectionChange(t, b, 2, "independent", "", nil, []c.Spec{{ID: "BET-001", Body: "It MUST still be checked."}}, nil)
		ds := p.CheckQueue(p.Snapshot{Modules: []c.Module{m, b}, Changes: []c.Change{independent, bad}})
		requireDiag(t, ds, "C05", "error", bad.Path)
		requireDiag(t, ds, "C10", "error", independent.Path)
	}
}
func Test_PRJ013HandleLocalityEverySection(t *testing.T) {
	m := queueModule(t, "alpha", "ALP", "units/a")
	one := projectionChange(t, m, 1, "one", "", []c.Spec{{ID: "new:local", Body: "It MUST use new:local."}}, nil, nil)
	two := projectionChange(t, m, 2, "two", "", []c.Spec{{ID: "new:local", Body: "It MUST use new:local."}}, nil, nil)
	queueClean(t, p.Snapshot{Modules: []c.Module{m}, Changes: []c.Change{two, one}})
	for _, kind := range []string{"body", "metadata", "notes"} {
		bad := two
		bad.Document.Sections = append([]c.Section{}, two.Document.Sections...)
		bad.Added = []c.Spec{{ID: "new:other", Body: "It MUST work locally."}}
		switch kind {
		case "body":
			bad.Added[0].Body = "It MUST NOT borrow new:local."
		case "metadata":
			bad.Blocked = "Waiting for new:local"
		case "notes":
			bad.Document.Sections = append(bad.Document.Sections, c.Section{Name: "Extra notes", Body: "Do not borrow new:local."})
		}
		requireDiag(t, p.CheckQueue(p.Snapshot{Modules: []c.Module{m}, Changes: []c.Change{one, bad}}), "C09", "error", bad.Path)
	}
}
func Test_PRJ002UndeclaredHandleDiagnostic(t *testing.T) {
	root := fixture(t)
	put(t, root, "units/a/SPEC.md", module("alpha", "ALP", ""))
	put(t, root, "changes/0001-first.md", change("alpha", ""))
	bad := strings.Replace(change("alpha", ""), "### new:fresh", "### new:other", 1)
	bad = strings.Replace(bad, "The fixture MUST work.", "It MUST use new:fresh.", 1)
	put(t, root, "changes/0002-cross.md", bad)
	put(t, root, "changes/0003-local.md", change("alpha", ""))
	put(t, root, "changes/0004-notes.md", change("alpha", "")+"\n## Extra notes\n\nUse new:foreign.\n")
	put(t, root, "changes/0005-meta.md", change("alpha", "blocked: waiting for new:foreign\n"))
	s, err := p.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{"changes/0002-cross.md", "changes/0004-notes.md", "changes/0005-meta.md"} {
		requireDiag(t, s.Diagnostics, "C09", "error", path)
	}
	if len(s.Changes) != 2 || s.Changes[0].Slug != "first" || s.Changes[1].Slug != "local" {
		t.Fatalf("lost valid local declarations %+v", s.Changes)
	}
}
func Test_PRJ014QueueDiagnosticsAndPurity(t *testing.T) {
	s := queueFixture(t)
	s.Changes = append([]c.Change{}, s.Changes...)
	s.Changes[2].Bases = nil
	s.Changes[1].Bases = nil
	s.Changes[0].Bases = nil
	s.Changes[2].Set = map[string]any{"interface": []any{"["}, "depends_on": []any{"zzz", "aaa"}, "hooks": map[string]any{"eval": "", "boundaries": ""}, "boundaries_opt_out": "sorting warning reason"}
	s.Diagnostics = []p.Diagnostic{{Code: "C02", Severity: "warning", Message: "Unrelated", Path: "outside/SPEC.md"}}
	before := encoded(t, s)
	ds := p.CheckQueue(s)
	if ds == nil || len(ds) < 3 {
		t.Fatalf("expected queued errors %+v", ds)
	}
	for _, d := range ds {
		if d.Message == "" || d.Severity != "error" && d.Severity != "warning" || d.Path == "outside/SPEC.md" {
			t.Fatalf("bad or repeated diagnostic %+v", d)
		}
	}
	if !sort.SliceIsSorted(ds, func(i, j int) bool {
		a, b := ds[i], ds[j]
		return a.Path+"\x00"+a.Code+"\x00"+a.Severity+"\x00"+a.Message < b.Path+"\x00"+b.Code+"\x00"+b.Severity+"\x00"+b.Message
	}) {
		t.Fatalf("unsorted %+v", ds)
	}
	if !reflect.DeepEqual(ds, p.CheckQueue(s)) {
		t.Fatal("unstable diagnostics")
	}
	poison(reflect.ValueOf(&ds).Elem())
	if encoded(t, s) != before {
		t.Fatal("queue mutated or aliased input")
	}
	empty := s
	empty.Changes = nil
	queueClean(t, empty)
	list := p.CheckQueue(empty)
	list = append(list, p.Diagnostic{Code: "poison"})
	queueClean(t, empty)
}

func Test_PRJ002ArchiveProseAndLiveReferences(t *testing.T) {
	root := fixture(t)
	put(t, root, "units/a/SPEC.md", module("alpha", "ALP", ""))
	text := strings.ReplaceAll(change("alpha", "phase: applied\nrun: archive-run\n"), "new:fresh", "ALP-009") + "\n## Decisions\n\nExplain the literal new:historical example.\n\n## Extra author context\n\nPast notes mention new:another-example.\n"
	put(t, root, "changes/archive/0007-prose.md", text)
	original, err := c.ParseMarkdown(text)
	if err != nil {
		t.Fatal(err)
	}
	badEdit := strings.ReplaceAll(change("alpha", "phase: applied\nrun: other-run\n"), "new:fresh", "ALP-010")
	badEdit = strings.Replace(badEdit, "The fixture MUST work.", "It MUST use new:unresolved.", 1)
	put(t, root, "changes/archive/0008-edit.md", badEdit)
	put(t, root, "changes/archive/0009-meta.md", change("alpha", "phase: applied\nblocked: waiting for new:unknown\n"))
	s, err := p.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(s.Archive) != 1 || s.Archive[0].Slug != "prose" || s.Archive[0].Added[0].ID != "ALP-009" || s.Archive[0].Run != "archive-run" || s.Archive[0].Phase != "applied" || !reflect.DeepEqual(s.Archive[0].Document, original) {
		t.Fatalf("lost historical archive %+v", s.Archive)
	}
	requireDiag(t, s.Diagnostics, "C09", "error", "changes/archive/0008-edit.md")
	requireDiag(t, s.Diagnostics, "C09", "error", "changes/archive/0009-meta.md")
	if len(s.Diagnostics) != 2 {
		t.Fatalf("historical examples treated as live references %+v", s.Diagnostics)
	}
}
