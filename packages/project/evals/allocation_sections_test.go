package evals_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
)

func allocationSectionModule(t *testing.T) c.Module {
	t.Helper()
	m, err := c.ParseModule("---\nname: alpha\nid_prefix: ALP\n---\n\n## Specs\n\n### ALP-001 [review]\nThe module MUST retain its contract.\n", "units/a", c.Zones{}, c.Hooks{})
	if err != nil {
		t.Fatal(err)
	}
	return m
}
func allocationSectionChange(t *testing.T, phase, edits string) c.Change {
	t.Helper()
	meta := ""
	if phase != "" {
		meta = "run: existing-run\nphase: " + phase + "\n"
	}
	raw := "---\nmodule: alpha\n" + meta + "---\n\nRetained preamble.\n\n## Motivation\n\nPreserve authored structure.\n\n## Empty author section\n\n## Notes\n\n```markdown\n## Added\n## Modified\n## Removed\n### ALP-777\nThis MUST remain a fenced example.\n```\n\n## Questions\n\nRetained question.\n" + edits
	ch, err := c.ParseChange(raw, "changes/0001-shape.md", "ALP")
	if err != nil {
		t.Fatal(err)
	}
	return ch
}
func allocationSectionNames(d c.Document) []string {
	names := []string{}
	for _, s := range d.Sections {
		names = append(names, s.Name)
	}
	return names
}
func allocationSectionHumanAgreement(t *testing.T, before, after c.Change) {
	t.Helper()
	if before.Document.Preamble != after.Document.Preamble {
		t.Error("changed preamble")
	}
	for _, section := range before.Document.Sections {
		switch section.Name {
		case "Added", "Modified", "Removed":
			continue
		}
		found := false
		for _, actual := range after.Document.Sections {
			if actual.Name == section.Name {
				found = true
				if !reflect.DeepEqual(section, actual) {
					t.Errorf("changed human section %s", section.Name)
				}
			}
		}
		if !found {
			t.Errorf("lost human section %s", section.Name)
		}
	}
}
func allocationSectionSnapshot(t *testing.T, ch c.Change) p.Snapshot {
	return p.Snapshot{Root: "/missing/allocation-section-root", Modules: []c.Module{allocationSectionModule(t)}, Changes: []c.Change{ch}}
}
func allocationSectionChecked(t *testing.T, s p.Snapshot) (c.Change, map[string]string) {
	t.Helper()
	before := encoded(t, s)
	got, bindings, err := p.Allocate(s, "shape")
	if err != nil {
		t.Fatalf("valid selected typed input: %v", err)
	}
	if bindings == nil {
		t.Error("nil mapping")
	}
	if encoded(t, s) != before {
		t.Error("allocation mutated supplied source")
	}
	allocationDocumentAgreement(t, got)
	allocationSectionHumanAgreement(t, s.Changes[0], got)
	if got.Run != s.Changes[0].Run || got.Phase != s.Changes[0].Phase || got.Path != s.Changes[0].Path || got.Slug != s.Changes[0].Slug || got.Number != s.Changes[0].Number {
		t.Error("changed lifecycle or typed identity")
	}
	return got, bindings
}
func Test_PRJ016AllocationEmptySectionPresence(t *testing.T) {
	for _, phase := range []string{"", "evals", "audit", "impl"} {
		for mask := 0; mask < 8; mask++ {
			t.Run(fmt.Sprintf("phase-%s-mask-%d", phase, mask), func(t *testing.T) {
				edits := ""
				for i, name := range []string{"Modified", "Added", "Removed"} {
					if mask&(1<<i) != 0 {
						edits += "\n## " + name + "\n"
					}
				}
				source := allocationSectionChange(t, phase, edits)
				expected := []string{"Motivation", "Empty author section", "Notes", "Questions"}
				for i, name := range []string{"Modified", "Added", "Removed"} {
					if mask&(1<<i) != 0 {
						expected = append(expected, name)
					}
				}
				if !reflect.DeepEqual(allocationSectionNames(source.Document), expected) {
					t.Fatal("source/fenced heading premise")
				}
				for _, empty := range []bool{false, true} {
					ch := source
					if empty {
						ch.Added = []c.Spec{}
						ch.Modified = []c.Spec{}
						ch.Removed = []c.Spec{}
					} else {
						ch.Added = nil
						ch.Modified = nil
						ch.Removed = nil
					}
					got, bindings := allocationSectionChecked(t, allocationSectionSnapshot(t, ch))
					if !reflect.DeepEqual(allocationSectionNames(got.Document), expected) {
						t.Errorf("empty=%v names %v want %v", empty, allocationSectionNames(got.Document), expected)
					}
					if len(got.Added) != 0 || len(got.Modified) != 0 || len(got.Removed) != 0 || len(bindings) != 0 {
						t.Errorf("empty lists invented edits or bindings: %+v %+v", got, bindings)
					}
					for _, section := range got.Document.Sections {
						switch section.Name {
						case "Added", "Modified", "Removed":
							if section.Body != "" {
								t.Errorf("empty %s acquired body %q", section.Name, section.Body)
							}
						}
					}
				}
			})
		}
	}
}
func Test_PRJ016AllocationTypedSectionsAreCreatedOrUpdated(t *testing.T) {
	for _, phase := range []string{"", "impl"} {
		for _, name := range []string{"Added", "Modified", "Removed"} {
			for _, present := range []bool{false, true} {
				t.Run(fmt.Sprintf("phase-%s-%s-present-%v", phase, name, present), func(t *testing.T) {
					edits := ""
					if present {
						edits = "\n## " + name + "\n"
					}
					ch := allocationSectionChange(t, phase, edits)
					item := c.Spec{ID: "ALP-001", Body: "The authoritative typed body MUST remain intact."}
					switch name {
					case "Added":
						if phase == "" {
							item.ID = "new:generated"
							item.Body = "The typed item MUST keep new:generated."
						} else {
							item.ID = "ALP-010"
						}
						ch.Added = []c.Spec{item}
					case "Modified":
						ch.Modified = []c.Spec{item}
					case "Removed":
						item.Body = ""
						ch.Removed = []c.Spec{item}
					}
					got, bindings := allocationSectionChecked(t, allocationSectionSnapshot(t, ch))
					expected := allocationSectionNames(ch.Document)
					if !present {
						expected = append(expected, name)
					}
					if !reflect.DeepEqual(allocationSectionNames(got.Document), expected) {
						t.Errorf("unexpected section creation/order %v want %v", allocationSectionNames(got.Document), expected)
					}
					var items []c.Spec
					switch name {
					case "Added":
						items = got.Added
					case "Modified":
						items = got.Modified
					case "Removed":
						items = got.Removed
					}
					want := item
					if name == "Added" && phase == "" {
						want.ID = "ALP-002"
						want.Body = strings.ReplaceAll(item.Body, "new:generated", "ALP-002")
						if !reflect.DeepEqual(bindings, map[string]string{"new:generated": "ALP-002"}) {
							t.Errorf("new binding %+v", bindings)
						}
					} else if len(bindings) != 0 {
						t.Errorf("unexpected binding %+v", bindings)
					}
					if !reflect.DeepEqual(items, []c.Spec{want}) {
						t.Errorf("typed items %+v want %+v", items, want)
					}
					for _, other := range []string{"Added", "Modified", "Removed"} {
						if other == name {
							continue
						}
						for _, section := range got.Document.Sections {
							if section.Name == other {
								t.Errorf("invented unrelated empty %s", other)
							}
						}
					}
				})
			}
		}
	}
}
func Test_PRJ016AllocationStaleBodiesBecomeRetainedEmptyHeadings(t *testing.T) {
	for _, name := range []string{"Added", "Modified", "Removed"} {
		t.Run(name, func(t *testing.T) {
			item := "### ALP-050\n\nThe stale retained item MUST be replaced by typed intent.\n"
			if name == "Removed" {
				item = "### ALP-050\n"
			}
			source := allocationSectionChange(t, "", "\n## "+name+"\n\n"+item)
			for _, empty := range []bool{false, true} {
				ch := source
				switch name {
				case "Added":
					if empty {
						ch.Added = []c.Spec{}
					} else {
						ch.Added = nil
					}
				case "Modified":
					if empty {
						ch.Modified = []c.Spec{}
					} else {
						ch.Modified = nil
					}
				case "Removed":
					if empty {
						ch.Removed = []c.Spec{}
					} else {
						ch.Removed = nil
					}
				}
				got, bindings := allocationSectionChecked(t, allocationSectionSnapshot(t, ch))
				if !reflect.DeepEqual(allocationSectionNames(got.Document), allocationSectionNames(ch.Document)) {
					t.Error("empty typed replacement moved or invented headings")
				}
				found := false
				for _, section := range got.Document.Sections {
					if section.Name == name {
						found = true
						if section.Body != "" {
							t.Errorf("stale retained body survives: %q", section.Body)
						}
					}
				}
				if !found {
					t.Error("explicit edit heading removed")
				}
				if len(bindings) != 0 || len(got.Added) != 0 || len(got.Modified) != 0 || len(got.Removed) != 0 {
					t.Error("stale items affected allocation")
				}
			}
		})
	}
}
func Test_PRJ016AllocationSectionFixtureControls(t *testing.T) {
	for _, phase := range []string{"", "evals", "audit", "impl"} {
		ch := allocationSectionChange(t, phase, "")
		if len(ch.Document.Sections) != 4 || ch.Kind != "maintenance" || ch.Module != "alpha" || ch.Slug != "shape" {
			t.Fatalf("public fixture %+v", ch)
		}
		for _, name := range allocationSectionNames(ch.Document) {
			if name == "Added" || name == "Modified" || name == "Removed" {
				t.Fatal("fenced headings became real sections")
			}
		}
		text, err := c.RenderMarkdown(ch.Document)
		if err != nil {
			t.Fatal(err)
		}
		again, err := c.ParseChange(text, ch.Path, "ALP")
		if err != nil || !reflect.DeepEqual(again, ch) {
			t.Fatal("public fixture round trip", err)
		}
	}
}
