package evals_test

import (
	"reflect"
	"strings"
	"testing"

	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
)

func allocationFixture(t *testing.T) p.Snapshot {
	t.Helper()
	m := queueModule(t, "alpha", "ALP", "units/a")
	m.Specs = append(m.Specs, c.Spec{ID: "ALP-008", Body: "It MUST reserve current IDs."})
	ch := projectionChange(t, m, 1, "allocate", "already_satisfied: [new:a]\nbases: {ALP-9000000: abcdef12}\nfixes: [ALP-8000000]\nset: {hooks: {setup: 'show new:aa and new:a'}}\n", []c.Spec{{ID: "new:a", Tag: "[review]", Body: "It MUST refer to new:aa."}, {ID: "new:aa", Body: "It MUST refer to new:a."}}, nil, nil)
	archive := projectionChange(t, m, 0, "past", "phase: applied\nrun: history\n", []c.Spec{{ID: "ALP-010", Body: "It MUST reserve earlier additions."}}, []c.Spec{{ID: "ALP-011", Body: "It MUST reserve earlier modifications."}}, []c.Spec{{ID: "ALP-1000"}})
	archive.Path = "changes/archive/0000-past.md"
	assigned := projectionChange(t, m, 2, "assigned", "run: assigned\nphase: evals\n", []c.Spec{{ID: "ALP-1005", Body: "It MUST reserve assigned IDs."}}, nil, nil)
	unstarted := projectionChange(t, m, 3, "unstarted", "", []c.Spec{{ID: "new:other", Body: "It MUST keep its own handle."}}, nil, nil)
	// The selected grammar alone is validated; an excluded invalid author-chosen
	// real ID must not reserve numbers, even at a higher number.
	invalidTail := unstarted
	invalidTail.Slug = "excluded"
	invalidTail.Number = 4
	invalidTail.Path = "changes/0004-excluded.md"
	invalidTail.Added = []c.Spec{{ID: "ALP-9000", Body: "Invalid excluded body."}}
	other := queueModule(t, "beta", "BET", "units/b")
	other.Specs = append(other.Specs, c.Spec{ID: "BET-999999", Body: "It MUST belong to beta."})
	return p.Snapshot{Root: "/nonexistent/allocation-root", BaseRef: "nonexistent-base", Modules: []c.Module{m, other}, Changes: []c.Change{invalidTail, assigned, ch, unstarted}, Archive: []c.Change{archive}, Diagnostics: []p.Diagnostic{{Code: "C02", Severity: "warning", Message: "Opaque discovery", Path: "elsewhere/SPEC.md"}}}
}
func Test_PRJ015ReservationsAndDeclarationOrder(t *testing.T) {
	s := allocationFixture(t)
	before := encoded(t, s)
	ch, mapping, err := p.Allocate(s, "allocate")
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{"new:a": "ALP-1006", "new:aa": "ALP-1007"}
	if !reflect.DeepEqual(mapping, want) || len(ch.Added) != 2 || ch.Added[0].ID != "ALP-1006" || ch.Added[1].ID != "ALP-1007" {
		t.Fatalf("allocation %+v %+v", ch, mapping)
	}
	if ch.Run != "" || ch.Phase != "" {
		t.Fatal("allocation started workflow")
	}
	if encoded(t, s) != before {
		t.Fatal("allocation mutated input")
	}
	reversed := allocationFixture(t)
	reversed.Changes[2].Added = []c.Spec{reversed.Changes[2].Added[1], reversed.Changes[2].Added[0]}
	reverseResult, reverseMap, e := p.Allocate(reversed, "allocate")
	if e != nil || reverseMap["new:aa"] != "ALP-1006" || reverseMap["new:a"] != "ALP-1007" || reverseResult.Added[0].ID != "ALP-1006" {
		t.Fatalf("declaration order sorted lexically: %+v %+v %v", reverseResult, reverseMap, e)
	}
	for _, state := range []struct{ run, phase string }{{"", ""}, {"held", ""}, {"", "evals"}} {
		v := allocationFixture(t)
		v.Changes[0].Run = state.run
		v.Changes[0].Phase = state.phase
		v.Archive = append(v.Archive, c.Change{Module: "beta", Removed: []c.Spec{{ID: "BET-99999999"}}})
		got, mp, e := p.Allocate(v, "allocate")
		if e != nil || got.Added[0].ID != "ALP-1006" || mp["new:aa"] != "ALP-1007" {
			t.Fatalf("excluded partial state reserved numbers: %+v %+v %v", got, mp, e)
		}
	}
	// Isolate every reservation source, so one larger reservation cannot mask
	// an implementation that ignores any smaller source.
	for _, source := range []string{"current", "archive-added", "archive-modified", "archive-removed", "queued-assigned"} {
		v := allocationFixture(t)
		m := v.Modules[0]
		m.Specs = []c.Spec{}
		v.Modules = []c.Module{m}
		v.Archive = nil
		v.Changes = []c.Change{s.Changes[2]}
		id := "ALP-023"
		spec := c.Spec{ID: id, Body: "It MUST reserve this ID."}
		switch source {
		case "current":
			v.Modules[0].Specs = []c.Spec{spec}
		case "archive-added":
			v.Archive = []c.Change{{Module: m.Name, Added: []c.Spec{spec}}}
		case "archive-modified":
			v.Archive = []c.Change{{Module: m.Name, Modified: []c.Spec{spec}}}
		case "archive-removed":
			v.Archive = []c.Change{{Module: m.Name, Removed: []c.Spec{{ID: id}}}}
		case "queued-assigned":
			v.Changes = append(v.Changes, c.Change{Module: m.Name, Run: "held", Phase: "impl", Added: []c.Spec{spec}})
		}
		got, mp, e := p.Allocate(v, "allocate")
		if e != nil || got.Added[0].ID != "ALP-024" || mp["new:aa"] != "ALP-025" {
			t.Fatalf("%s reservation %+v %+v %v", source, got, mp, e)
		}
	}
}

func Test_PRJ015AllNumericReservationSourcesAndScope(t *testing.T) {
	for _, source := range []string{"current", "archive-added", "archive-modified", "archive-removed", "queued-assigned"} {
		for _, width := range []struct{ high, want string }{{"ALP-0999", "ALP-1000"}, {"ALP-999999999999999999999999999999", "ALP-1000000000000000000000000000000"}} {
			s := allocationFixture(t)
			s.Modules = s.Modules[:1]
			s.Modules[0].Specs = []c.Spec{}
			s.Archive = nil
			s.Changes = []c.Change{s.Changes[2]}
			items := []c.Spec{{ID: width.high, Body: "It MUST reserve the numeric maximum."}, {ID: "ALP-100", Body: "It MUST distinguish variable widths."}}
			switch source {
			case "current":
				s.Modules[0].Specs = items
			case "archive-added":
				s.Archive = []c.Change{{Module: "alpha", Added: items}}
			case "archive-modified":
				s.Archive = []c.Change{{Module: "alpha", Modified: items}}
			case "archive-removed":
				s.Archive = []c.Change{{Module: "alpha", Removed: []c.Spec{{ID: width.high}, {ID: "ALP-100"}}}}
			case "queued-assigned":
				s.Changes = append(s.Changes, c.Change{Module: "alpha", Run: "held", Phase: "impl", Added: items})
			}
			got, mp, err := p.Allocate(s, "allocate")
			if err != nil || got.Added[0].ID != width.want || mp["new:a"] != width.want {
				t.Fatalf("%s %s: %+v %+v %v", source, width.high, got, mp, err)
			}
		}
	}
	s := allocationFixture(t)
	// Prefix uniqueness is a separate validator's responsibility. The selected
	// allocator must filter module ownership even when another record is invalid.
	s.Modules[1].IDPrefix = "ALP"
	s.Modules[1].Specs = []c.Spec{{ID: "ALP-999999999", Body: "It MUST belong to the other module."}}
	s.Archive = append(s.Archive, c.Change{Module: "beta", Added: []c.Spec{{ID: "ALP-999999998"}}, Modified: []c.Spec{{ID: "ALP-999999997"}}, Removed: []c.Spec{{ID: "ALP-999999996"}}})
	s.Changes = append(s.Changes, c.Change{Module: "beta", Run: "other", Phase: "evals", Added: []c.Spec{{ID: "ALP-999999995"}}})
	s.Changes[1].Modified = []c.Spec{{ID: "ALP-7777777", Body: "It MUST only reference an existing ID."}}
	s.Changes[1].Removed = []c.Spec{{ID: "ALP-6666666"}}
	// A malformed predecessor must not be replayed to allocate the selected change.
	predecessor := s.Changes[0]
	predecessor.Number = 0
	predecessor.Slug = "invalid-before"
	predecessor.Path = "changes/0000-invalid-before.md"
	predecessor.Set = map[string]any{"name": "forbidden"}
	s.Changes = append(s.Changes, predecessor)
	got, mp, err := p.Allocate(s, "allocate")
	if err != nil || got.Added[0].ID != "ALP-1006" || mp["new:aa"] != "ALP-1007" {
		t.Fatalf("excluded owners/references/predecessor affected allocation: %+v %+v %v", got, mp, err)
	}
}

func allocationDocumentAgreement(t *testing.T, ch c.Change) {
	t.Helper()
	text, err := c.RenderMarkdown(ch.Document)
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := c.ParseChange(text, ch.Path, "ALP")
	if err != nil {
		t.Fatal(err)
	}
	parsed.Number = ch.Number
	if !reflect.DeepEqual(parsed, ch) {
		t.Fatalf("document/typed disagreement: %+v %+v", parsed, ch)
	}
}

func Test_PRJ015WidthsAndResume(t *testing.T) {
	for _, v := range []struct{ reserved, want string }{{"", "ALP-001"}, {"ALP-009", "ALP-010"}, {"ALP-0999", "ALP-1000"}, {"ALP-999999999999999999999999999999", "ALP-1000000000000000000000000000000"}} {
		s := allocationFixture(t)
		s.Modules = s.Modules[:1]
		s.Modules[0].Specs = []c.Spec{}
		if v.reserved != "" {
			s.Modules[0].Specs = []c.Spec{{ID: v.reserved, Body: "It MUST reserve this width."}}
		}
		s.Archive = nil
		s.Changes = []c.Change{s.Changes[2]}
		got, mp, err := p.Allocate(s, "allocate")
		if err != nil || got.Added[0].ID != v.want || mp["new:a"] != v.want {
			t.Fatalf("width %s %+v %+v %v", v.reserved, got, mp, err)
		}
	}
	for _, phase := range []string{"evals", "audit", "impl"} {
		s := allocationFixture(t)
		ch := s.Changes[1]
		ch.Phase = phase
		ch.Number = 83
		ch.Run = "typed-run"
		s.Changes = []c.Change{ch}
		before := encoded(t, s)
		got, mp, err := p.Allocate(s, ch.Slug)
		if err != nil || mp == nil || len(mp) != 0 || got.Added[0].ID != "ALP-1005" || got.Run != "typed-run" || got.Phase != phase || got.Number != 83 {
			t.Fatalf("resume %+v %+v %v", got, mp, err)
		}
		allocationDocumentAgreement(t, got)
		second, mp2, err := p.Allocate(s, ch.Slug)
		if err != nil {
			t.Fatal(err)
		}
		secondBefore, mapBefore := encoded(t, second), encoded(t, mp2)
		poison(reflect.ValueOf(&got).Elem())
		mp["poison"] = "poison"
		if encoded(t, s) != before || encoded(t, second) != secondBefore || encoded(t, mp2) != mapBefore {
			t.Fatal("real-ID resume aliases input or another result")
		}
		third, _, err := p.Allocate(s, ch.Slug)
		if err != nil || !reflect.DeepEqual(second, third) {
			t.Fatal("real-ID resume mutation poisoned future results")
		}
	}
}
func Test_PRJ016RewriteEveryLocationAndTypedIdentity(t *testing.T) {
	s := allocationFixture(t)
	ch := s.Changes[2]
	ch.Number = 77
	ch.Added = append([]c.Spec{}, ch.Added...)
	ch.Added[0].Body = "It MUST use a typed reference to new:aa."
	ch.Added[0].Tag = "[quarantine: authored reason]"
	ch.Document.Preamble = "# Human new:a and new:aa ALP-1234567"
	ch.Document.Sections = append([]c.Section{}, ch.Document.Sections...)
	for i, section := range ch.Document.Sections {
		if section.Name == "Motivation" {
			ch.Document.Sections[i].Body = "Motivate new:a and new:aa."
		}
		if section.Name == "Interface" {
			ch.Document.Sections[i].Body = "Use new:aa before new:a."
		}
		if section.Name == "Notes" {
			ch.Document.Sections[i].Body = "Keep punctuation: (new:a), new:aa.\n\n```text\nnew:aa new:a\n```"
		}
	}
	ch.Document.Sections = append(ch.Document.Sections, c.Section{Name: "Decisions", Body: "Choose new:a and new:aa."}, c.Section{Name: "Extra author text", Body: "Preserve this unrelated text exactly."})
	ch.Modified = []c.Spec{{ID: "ALP-001", Tag: "[review]", Body: "It MUST refer to new:a and new:aa."}}
	ch.Removed = []c.Spec{{ID: "ALP-002", Tag: "[review]"}}
	ch.Regenerate = true
	ch.Blocked = "Wait for new:aa"
	ch.Set["preserve"] = []any{}
	ch.Set["depends_on"] = []any{"beta"}
	for i := range ch.Document.Sections {
		if ch.Document.Sections[i].Name == "Extra author text" {
			ch.Document.Sections[i].Body = "\r\nPreserve this unrelated text exactly.  \r\nSecond line.\r\n\r\n"
		}
	}
	// All of these retained fields are stale; typed values define the transform.
	ch.Document.Metadata["already_satisfied"] = []any{"ALP-999"}
	ch.Document.Metadata["set"] = map[string]any{"hooks": map[string]any{"setup": "stale"}}
	ch.Document.Metadata["module"] = "stale"
	ch.Document.Metadata["blocked"] = "stale"
	ch.Document.Metadata["regenerate"] = false
	ch.Document.Metadata["run"] = "stale-run"
	ch.Document.Metadata["phase"] = "impl"
	normalizedText, err := c.RenderMarkdown(ch.Document)
	if err != nil {
		t.Fatal(err)
	}
	normalizedDoc, err := c.ParseMarkdown(normalizedText)
	if err != nil {
		t.Fatal(err)
	}
	var expectedHumanBody string
	for _, section := range normalizedDoc.Sections {
		if section.Name == "Extra author text" {
			expectedHumanBody = section.Body
		}
	}
	s.Changes = []c.Change{ch}
	s.Archive = nil
	s.Modules[0].Specs = s.Modules[0].Specs[:3]
	got, mp, err := p.Allocate(s, ch.Slug)
	if err != nil {
		t.Fatal(err)
	}
	if mp["new:a"] != "ALP-004" || mp["new:aa"] != "ALP-005" || got.Path != ch.Path || got.Slug != ch.Slug || got.Number != 77 || got.Module != "alpha" || got.Blocked != "Wait for ALP-005" || got.Added[0].Tag != "[quarantine: authored reason]" {
		t.Fatalf("typed identity %+v %+v", got, mp)
	}
	if !reflect.DeepEqual(got.AlreadySatisfied, []string{"ALP-004"}) || got.Set["hooks"].(map[string]any)["setup"] != "show ALP-005 and ALP-004" || got.Modified[0].Body != "It MUST refer to ALP-004 and ALP-005." || got.Modified[0].Tag != "[review]" {
		t.Fatalf("typed rewrite %+v", got)
	}
	if !reflect.DeepEqual(got.Bases, ch.Bases) || !reflect.DeepEqual(got.Fixes, ch.Fixes) || !reflect.DeepEqual(got.Removed, ch.Removed) || got.Modified[0].ID != "ALP-001" || !got.Regenerate || got.Kind != "regenerate" || got.Run != "" || got.Phase != "" || !reflect.DeepEqual(got.Set["preserve"], []any{}) || !reflect.DeepEqual(got.Set["depends_on"], []any{"beta"}) {
		t.Fatalf("lost unrelated real IDs/metadata: %+v", got)
	}
	for _, section := range got.Document.Sections {
		if section.Name == "Extra author text" && section.Body != expectedHumanBody {
			t.Fatalf("human normalization changed: %q want %q", section.Body, expectedHumanBody)
		}
	}
	text, err := c.RenderMarkdown(got.Document)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(text, "new:") {
		t.Fatalf("unrewritten handle %s", text)
	}
	for _, want := range []string{"# Human ALP-004 and ALP-005 ALP-1234567", "Motivate ALP-004 and ALP-005.", "Use ALP-005 before ALP-004.", "Choose ALP-004 and ALP-005.", "(ALP-004), ALP-005.", "```text\nALP-005 ALP-004\n```", "Preserve this unrelated text exactly.", "It MUST use a typed reference to ALP-005."} {
		if !strings.Contains(text, want) {
			t.Fatalf("lost rewrite %q in %s", want, text)
		}
	}
	parsed, err := c.ParseChange(text, ch.Path, "ALP")
	if err != nil {
		t.Fatal(err)
	}
	parsed.Number = got.Number
	if !reflect.DeepEqual(parsed, got) {
		t.Fatalf("document disagrees with typed result\n%+v\n%+v", parsed, got)
	}
	if len(got.Document.Sections) != len(ch.Document.Sections) || got.Document.Preamble == "" {
		t.Fatal("lost section order or preamble")
	}
	for i, section := range ch.Document.Sections {
		if got.Document.Sections[i].Name != section.Name {
			t.Fatal("section order changed")
		}
	}
}
func Test_PRJ017AllocationErrorsAreAtomic(t *testing.T) {
	for _, kind := range []string{"empty", "missing", "duplicate", "target", "duplicate-target", "body", "set", "handle-body", "handle-meta", "cross-only", "malformed-handle", "duplicate-handle", "phase", "hash", "author-id", "run-only", "phase-only", "applied", "started-handles"} {
		s := allocationFixture(t)
		s.Changes = []c.Change{s.Changes[2]}
		slug := "allocate"
		switch kind {
		case "empty":
			slug = ""
		case "missing":
			slug = "missing"
		case "duplicate":
			s.Changes = append(s.Changes, s.Changes[0])
		case "target":
			s.Changes[0].Module = "missing"
		case "duplicate-target":
			s.Modules = append(s.Modules, s.Modules[0])
		case "body":
			s.Changes[0].Added[0].Body = "No normative text."
		case "set":
			s.Changes[0].Set = map[string]any{"name": "renamed"}
		case "handle-body":
			s.Changes[0].Added[0].Body = "It MUST reject new:foreign."
		case "handle-meta":
			s.Changes[0].Blocked = "Waiting for new:foreign"
		case "cross-only":
			s.Changes[0].Added[0].Body = "It MUST reject borrowing new:foreign."
			other := projectionChange(t, s.Modules[1], 2, "foreign", "", []c.Spec{{ID: "new:foreign", Body: "It MUST remain local to another change."}}, nil, nil)
			s.Changes = append(s.Changes, other)
		case "malformed-handle":
			s.Changes[0].Added = []c.Spec{{ID: "new:UPPER", Body: "It MUST reject malformed declarations."}}
			s.Changes[0].AlreadySatisfied = nil
			s.Changes[0].Set = nil
		case "duplicate-handle":
			s.Changes[0].Added = []c.Spec{{ID: "new:a", Body: "It MUST remain unique."}, {ID: "new:a", Body: "It MUST remain unique."}}
			s.Changes[0].Set = nil
		case "phase":
			s.Changes[0].Phase = "invalid"
		case "hash":
			s.Changes[0].Bases = map[string]string{"ALP-001": "invalid"}
		case "author-id":
			s.Changes[0].Added = []c.Spec{{ID: "ALP-099", Body: "It MUST wait for assignment."}}
			s.Changes[0].AlreadySatisfied = nil
			s.Changes[0].Set = nil
		case "run-only":
			s.Changes[0].Added = []c.Spec{{ID: "ALP-099", Body: "It MUST keep its assigned ID."}}
			s.Changes[0].Set = nil
			s.Changes[0].AlreadySatisfied = nil
			s.Changes[0].Run = "run"
		case "phase-only":
			s.Changes[0].Added = []c.Spec{{ID: "ALP-099", Body: "It MUST keep its assigned ID."}}
			s.Changes[0].Set = nil
			s.Changes[0].AlreadySatisfied = nil
			s.Changes[0].Phase = "evals"
		case "applied":
			s.Changes[0].Added = []c.Spec{{ID: "ALP-099", Body: "It MUST keep its assigned ID."}}
			s.Changes[0].Set = nil
			s.Changes[0].AlreadySatisfied = nil
			s.Changes[0].Run = "run"
			s.Changes[0].Phase = "applied"
		case "started-handles":
			s.Changes[0].Run = "run"
			s.Changes[0].Phase = "impl"
		}
		before := encoded(t, s)
		got, mp, err := p.Allocate(s, slug)
		if err == nil || mp != nil || !reflect.DeepEqual(got, c.Change{}) {
			t.Fatalf("%s non-atomic error %+v %+v %v", kind, got, mp, err)
		}
		repeated, repeatedMap, repeatedError := p.Allocate(s, slug)
		if repeatedError == nil || repeatedError.Error() != err.Error() || !reflect.DeepEqual(repeated, got) || !reflect.DeepEqual(repeatedMap, mp) {
			t.Fatalf("%s error is nondeterministic", kind)
		}
		if encoded(t, s) != before {
			t.Fatalf("%s error mutated input", kind)
		}
	}
}
func Test_PRJ017AllocationIndependenceAndEmptyMaps(t *testing.T) {
	s := allocationFixture(t)
	s.Config = c.Config{Agent: c.Agent{Command: []string{"opaque"}}, ModuleDefaults: c.Zones{Preserve: []string{"keep/**"}}}
	before := encoded(t, s)
	first, mp, err := p.Allocate(s, "allocate")
	if err != nil {
		t.Fatal(err)
	}
	second, mp2, err := p.Allocate(s, "allocate")
	if err != nil || !reflect.DeepEqual(first, second) || !reflect.DeepEqual(mp, mp2) {
		t.Fatalf("nondeterministic %v", err)
	}
	secondBefore := encoded(t, second)
	mappingBefore := encoded(t, mp2)
	poison(reflect.ValueOf(&first).Elem())
	mp["new:a"] = "poison"
	if encoded(t, s) != before || encoded(t, second) != secondBefore || encoded(t, mp2) != mappingBefore {
		t.Fatal("allocation outputs alias each other or inputs")
	}
	third, mp3, err := p.Allocate(s, "allocate")
	if err != nil || !reflect.DeepEqual(second, third) || !reflect.DeepEqual(mp2, mp3) {
		t.Fatal("output mutation poisoned later allocation")
	}
	for _, started := range []bool{false, true} {
		empty := projectionChange(t, s.Modules[0], 1, "empty", "set: {preserve: [keep/**]}\n", nil, nil, nil)
		if started {
			empty.Run = "held"
			empty.Phase = "impl"
		}
		s.Changes = []c.Change{empty}
		input := encoded(t, s)
		got, a, e := p.Allocate(s, "empty")
		if e != nil || a == nil || len(a) != 0 || len(got.Added) != 0 {
			t.Fatalf("empty allocation %+v %+v %v", got, a, e)
		}
		a["poison"] = "poison"
		poison(reflect.ValueOf(&got).Elem())
		_, b, e := p.Allocate(s, "empty")
		if e != nil || b == nil || len(b) != 0 || encoded(t, s) != input {
			t.Fatal("empty map/result aliases caller or future results")
		}
	}
}
