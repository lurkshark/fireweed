package evals_test

import (
	"reflect"
	"strings"
	"testing"

	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
)

func finalSection(d *c.Document, name, body string) {
	for i, s := range d.Sections {
		if s.Name == name {
			d.Sections[i].Body = body
			return
		}
	}
	d.Sections = append(d.Sections, c.Section{Name: name, Body: body})
}
func finalBody(d c.Document, name string) (string, bool) {
	for _, s := range d.Sections {
		if s.Name == name {
			return s.Body, true
		}
	}
	return "", false
}
func finalAgreement(t *testing.T, m c.Module) {
	t.Helper()
	text, err := c.RenderMarkdown(m.Document)
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := c.ParseModule(text, m.Dir, c.Zones{}, c.Hooks{})
	if err != nil {
		t.Fatalf("not a valid current module: %v\n%s", err, text)
	}
	if !reflect.DeepEqual(parsed, m) {
		t.Fatalf("canonical module disagreement\n%+v\n%+v", parsed, m)
	}
}
func finalFixture(t *testing.T) (c.Module, c.Change) {
	t.Helper()
	m := queueModule(t, "alpha", "ALP", "units/a")
	ch := projectionChange(t, m, 1, "final", "run: held\nphase: impl\nset:\n  depends_on: [beta]\n  interface: [api2.go]\n  evals: [tests/**]\n  disposable: [impl/**]\n  preserve: [keep/**]\n  append_only: [events/**]\n  boundaries_opt_out: visible reason\n  hooks: {eval: next eval, boundaries: next boundaries, setup: next setup}\n", []c.Spec{{ID: "ALP-009", Tag: "[review]", Body: "It MUST append a permanent spec."}}, []c.Spec{{ID: "ALP-001", Tag: "[quarantine: test clock]", Body: "It MUST replace the first behavior."}}, []c.Spec{{ID: "ALP-002"}})
	finalSection(&ch.Document, "Decisions", "Keep the new choice.")
	return m, ch
}
func Test_PRJ018FinalSpecsMetadataAndAuthorSeparation(t *testing.T) {
	m, ch := finalFixture(t)
	ch.Number = 88
	ch.Modified[0].Body = "\r\n\r\nIt MUST replace the first behavior.  \r\n\r\n"
	normalizedItems, err := c.ParseItems("### ALP-001 [quarantine: test clock]\n\n"+ch.Modified[0].Body, m.IDPrefix, false)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"Motivation", "Interface", "Notes", "Audit", "Questions", "Provenance"} {
		finalSection(&ch.Document, name, "PRIVATE "+name+" context.")
	}
	before := encoded(t, []any{m, ch})
	got, err := p.Finalize(m, ch)
	if err != nil {
		t.Fatal(err)
	}
	want := []c.Spec{normalizedItems[0], m.Specs[2], ch.Added[0]}
	if !reflect.DeepEqual(got.Specs, want) || got.Name != m.Name || got.IDPrefix != m.IDPrefix || got.Dir != m.Dir || got.Document.Preamble != m.Document.Preamble {
		t.Fatalf("final item/identity/preamble values %+v", got)
	}
	if !reflect.DeepEqual(got.DependsOn, []string{"beta"}) || !reflect.DeepEqual(got.Interface, []string{"api2.go"}) || !reflect.DeepEqual(got.Evals, []string{"tests/**"}) || !reflect.DeepEqual(got.Disposable, []string{"impl/**"}) || !reflect.DeepEqual(got.Preserve, []string{"keep/**"}) || !reflect.DeepEqual(got.AppendOnly, []string{"events/**"}) || got.BoundariesOptOut != "visible reason" || got.Hooks != (c.Hooks{Eval: "next eval", Boundaries: "next boundaries", Setup: "next setup"}) {
		t.Fatalf("lost final set metadata %+v", got)
	}
	body, ok := finalBody(got.Document, "Human notes")
	if !ok || body != "Keep unrelated notes." {
		t.Fatal("lost unrelated module section")
	}
	text, err := c.RenderMarkdown(got.Document)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(text, "PRIVATE ") || strings.Contains(text, "held") || strings.Contains(text, "phase:") {
		t.Fatalf("leaked author/state context %s", text)
	}
	if !strings.Contains(got.Decisions, "[final]\nKeep the new choice.") {
		t.Fatalf("missing tagged Decisions %q", got.Decisions)
	}
	finalAgreement(t, got)
	if encoded(t, []any{m, ch}) != before || ch.Run != "held" || ch.Phase != "impl" {
		t.Fatal("final application mutated input/state")
	}
}
func Test_PRJ019QuotedRemovalsAndMultilineAppend(t *testing.T) {
	m, ch := finalFixture(t)
	m.Decisions = "[old]\nKeep important. Drop repeated.\n\nDrop repeated.\n\nKeep lower drop repeated.\n\n>Quoted marker.\n\nKeep final choice."
	finalSection(&m.Document, "Decisions", "Stale retained choice MUST disappear.")
	finalSection(&ch.Document, "Removed decisions", " \t> Drop repeated. \t\n\n  > >Quoted marker.  \n> absent quote\n")
	addition := "Choose a reproducible path.\n\n```text\n## fenced decision example\n```\n\nKeep another new paragraph."
	finalSection(&ch.Document, "Decisions", addition)
	got, err := p.Finalize(m, ch)
	if err != nil {
		t.Fatal(err)
	}
	expected := "[old]\nKeep important. \n\n\n\nKeep lower drop repeated.\n\n\n\nKeep final choice.\n\n[final]\n" + addition
	// The public document contract defines retained internal whitespace and CRLF.
	d := c.Document{Metadata: map[string]any{}, Sections: []c.Section{{Name: "Decisions", Body: expected}}}
	normalized, err := c.RenderMarkdown(d)
	if err != nil {
		t.Fatal(err)
	}
	parsed, err := c.ParseMarkdown(normalized)
	if err != nil {
		t.Fatal(err)
	}
	expected, _ = finalBody(parsed, "Decisions")
	if got.Decisions != expected {
		t.Fatalf("decision transform %q want %q", got.Decisions, expected)
	}
	body, ok := finalBody(got.Document, "Decisions")
	if !ok || body != got.Decisions {
		t.Fatal("typed Decisions disagree")
	}
	finalAgreement(t, got)
}
func Test_PRJ019QuoteModesAndAppendSpacing(t *testing.T) {
	for _, quote := range []string{"Erase this.", "> Erase this.", ">Erase this.", "  > Erase this.  "} {
		m, ch := finalFixture(t)
		m.Decisions = "Erase this.\nKeep this."
		finalSection(&ch.Document, "Decisions", "")
		finalSection(&ch.Document, "Removed decisions", quote)
		got, err := p.Finalize(m, ch)
		if err != nil || got.Decisions != "Keep this." {
			t.Fatalf("quote %q: %q %v", quote, got.Decisions, err)
		}
		finalAgreement(t, got)
	}
	m, ch := finalFixture(t)
	m.Decisions = "\r\nKeep existing.\r\n\r\n"
	finalSection(&ch.Document, "Decisions", "\r\nAdded choice.\r\n")
	got, err := p.Finalize(m, ch)
	if err != nil || got.Decisions != "Keep existing.\n\n[final]\nAdded choice." {
		t.Fatalf("append spacing %q %v", got.Decisions, err)
	}
	finalAgreement(t, got)
}

func Test_PRJ019DecisionSectionPresence(t *testing.T) {
	for _, kind := range []string{"absent", "existing-empty", "typed-empty-stale", "remove-last", "unmatched-removal", "add-first", "blank-addition"} {
		m := queueModule(t, "alpha", "ALP", "units/a")
		m.Decisions = ""
		sections := []c.Section{}
		for _, s := range m.Document.Sections {
			if s.Name != "Decisions" {
				sections = append(sections, s)
			}
		}
		m.Document.Sections = sections
		ch := projectionChange(t, m, 1, "section", "", nil, nil, nil)
		present := false
		want := ""
		switch kind {
		case "existing-empty":
			finalSection(&m.Document, "Decisions", "")
			present = true
		case "typed-empty-stale":
			finalSection(&m.Document, "Decisions", "STALE content must disappear.")
			present = true
		case "remove-last":
			m.Decisions = "Choice."
			finalSection(&m.Document, "Decisions", "stale")
			finalSection(&ch.Document, "Removed decisions", "> Choice.")
			present = true
		case "unmatched-removal":
			finalSection(&ch.Document, "Removed decisions", "> not present")
			present = true
		case "add-first":
			finalSection(&ch.Document, "Decisions", "First choice.")
			present = true
			want = "[section]\nFirst choice."
		case "blank-addition":
			finalSection(&ch.Document, "Decisions", " \r\n \r\n")
		}
		got, err := p.Finalize(m, ch)
		if err != nil {
			t.Fatalf("%s %v", kind, err)
		}
		body, exists := finalBody(got.Document, "Decisions")
		if got.Decisions != want || body != want || exists != present {
			t.Fatalf("%s section %q %v typed %q", kind, body, exists, got.Decisions)
		}
		finalAgreement(t, got)
	}
}
func Test_PRJ018FinalErrorsAreAtomic(t *testing.T) {
	for _, kind := range []string{"target", "body", "current-body", "current-prefix", "missing-modified", "missing-removed", "duplicate-edit", "set", "wrong-prefix", "current-handle", "result-handle", "empty-slug", "malformed-slug", "slug-path", "bad-path"} {
		m, ch := finalFixture(t)
		switch kind {
		case "target":
			ch.Module = "missing"
		case "body":
			ch.Added[0].Body = "No normative text."
		case "current-body":
			m.Specs[2].Body = "An untouched item without normative text."
		case "current-prefix":
			m.Specs[2].ID = "BAD-003"
		case "missing-modified":
			ch.Modified[0].ID = "ALP-999"
		case "missing-removed":
			ch.Removed[0].ID = "ALP-999"
		case "duplicate-edit":
			ch.Removed = []c.Spec{{ID: "ALP-001"}}
		case "set":
			ch.Set = map[string]any{"id_prefix": "BAD"}
		case "wrong-prefix":
			ch.Added[0].ID = "BAD-009"
		case "current-handle":
			m.Specs[2].ID = "new:preview"
		case "result-handle":
			ch.Added[0].ID = "new:pending"
		case "empty-slug":
			ch.Slug = ""
		case "malformed-slug":
			ch.Slug = "INVALID"
			ch.Path = "changes/0001-INVALID.md"
		case "slug-path":
			ch.Slug = "different"
		case "bad-path":
			ch.Path = "elsewhere/change.md"
		}
		before := encoded(t, []any{m, ch})
		got, err := p.Finalize(m, ch)
		if err == nil || !reflect.DeepEqual(got, c.Module{}) || encoded(t, []any{m, ch}) != before {
			t.Fatalf("%s non-atomic or mutating error %+v %v", kind, got, err)
		}
		again, e := p.Finalize(m, ch)
		if e == nil || e.Error() != err.Error() || !reflect.DeepEqual(again, got) {
			t.Fatalf("%s nondeterministic error", kind)
		}
	}
}
func Test_PRJ020FinalResultIndependenceAndPreview(t *testing.T) {
	for _, maintenance := range []bool{false, true} {
		m, ch := finalFixture(t)
		if maintenance {
			ch = projectionChange(t, m, 1, "maintenance", "set: {preserve: []}\n", nil, nil, nil)
			finalSection(&ch.Document, "Decisions", "Preserve a maintenance choice.")
		}
		finalSection(&ch.Document, "Removed decisions", "> Keep this reason.")
		input := encoded(t, []any{m, ch})
		first, err := p.Finalize(m, ch)
		if err != nil {
			t.Fatal(err)
		}
		second, err := p.Finalize(m, ch)
		if err != nil || !reflect.DeepEqual(first, second) {
			t.Fatalf("unstable final result %v", err)
		}
		if maintenance {
			if second.Preserve == nil || len(second.Preserve) != 0 || !strings.Contains(second.Decisions, "[maintenance]\nPreserve a maintenance choice.") {
				t.Fatalf("maintenance transform omitted: %+v", second)
			}
			finalAgreement(t, second)
		}
		secondBefore := encoded(t, second)
		poison(reflect.ValueOf(&first).Elem())
		if encoded(t, []any{m, ch}) != input || encoded(t, second) != secondBefore {
			t.Fatal("final result aliases input/another result")
		}
		third, err := p.Finalize(m, ch)
		if err != nil || !reflect.DeepEqual(second, third) {
			t.Fatal("mutation poisoned later finalization")
		}
		preview, err := p.ApplyChange(m, ch)
		if err != nil {
			t.Fatal(err)
		}
		body, _ := finalBody(preview.Document, "Decisions")
		if preview.Decisions != m.Decisions || body != m.Decisions || strings.Contains(body, "["+ch.Slug+"]") {
			t.Fatalf("preview applied Decisions early %+v", preview)
		}
		if encoded(t, []any{m, ch}) != input {
			t.Fatal("preview mutated finalization input")
		}
	}
}

func Test_PRJ020EmptySliceStorageIndependence(t *testing.T) {
	m := queueModule(t, "alpha", "ALP", "units/a")
	typedBacking := []string{"current backing"}
	m.Preserve = typedBacking[:0]
	docBacking := []any{"document backing"}
	m.Document.Metadata["preserve"] = docBacking[:0]
	ch := projectionChange(t, m, 1, "empty-storage", "set: {preserve: []}\n", nil, nil, nil)
	changeBacking := []any{"change backing"}
	ch.Set["preserve"] = changeBacking[:0]
	before := encoded(t, []any{m, ch})
	a, err := p.Finalize(m, ch)
	if err != nil {
		t.Fatal(err)
	}
	b, err := p.Finalize(m, ch)
	if err != nil {
		t.Fatal(err)
	}
	finalAgreement(t, a)
	finalAgreement(t, b)
	if a.Preserve == nil || b.Preserve == nil || len(a.Preserve) != 0 || len(b.Preserve) != 0 {
		t.Fatal("explicit empty typed preserve lost")
	}
	a.Preserve = append(a.Preserve, "first typed append")
	b.Preserve = append(b.Preserve, "second typed append")
	if a.Preserve[0] != "first typed append" {
		t.Fatal("returned empty typed slices share backing")
	}
	aMeta, aOK := a.Document.Metadata["preserve"].([]any)
	bMeta, bOK := b.Document.Metadata["preserve"].([]any)
	if !aOK || !bOK || aMeta == nil || bMeta == nil || len(aMeta) != 0 || len(bMeta) != 0 {
		t.Fatal("explicit empty metadata preserve lost")
	}
	aMeta = append(aMeta, "first metadata append")
	bMeta = append(bMeta, "second metadata append")
	if aMeta[0] != "first metadata append" {
		t.Fatal("returned empty metadata slices share backing")
	}
	if typedBacking[0] != "current backing" || docBacking[0] != "document backing" || changeBacking[0] != "change backing" || encoded(t, []any{m, ch}) != before {
		t.Fatal("empty result append mutated input backing")
	}
	fresh, err := p.Finalize(m, ch)
	if err != nil || len(fresh.Preserve) != 0 {
		t.Fatal("empty append poisoned later call")
	}
	finalAgreement(t, fresh)
}
