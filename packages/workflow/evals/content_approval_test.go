package evals

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

const contentApprovalPath = "changes/0001-current.md"

func contentApprovalText() string {
	return `---
module: alpha
set:
  depends_on: [bravo, delta]
  interface: [api.go, other.go]
  evals: [evals/**, tests/**]
  disposable: [internal/**, src/**]
  preserve: [kept/**, saved/**]
  append_only: [log/**, history/**]
  boundaries_opt_out: authored exception
  hooks:
    eval: check new:first new:first-long
    boundaries: boundary-trap
    setup: setup-trap
bases: {ALP-001: abcdef12, ALP-009: "12345678"}
fixes: [ALP-001, ALP-004]
already_satisfied: [new:first, new:first-long]
---

Author preamble new:first and new:first-long.

## Motivation

Keep approval tied to the complete intent.

## Interface

Expose new:first and new:first-long.

## Added

### new:first
The target MUST retain new:first-long exactly.

### new:first-long [review]
The target MUST keep new:first and new:first-long.

## Modified

### ALP-001
The module MUST refer to new:first.

## Removed

### ALP-009

## Removed decisions

> An earlier decision to remove.

## Rejected alternatives

Skipping intent checks loses authority.

## Author guidance

` + "```markdown\n## Audit\nnew:first stays an author example.\n```\n" + `
## Decisions

Initial worker decision.

## Questions

Initial question.

## Audit

Initial audit.

## Provenance

Initial provenance.

## Notes

Initial note.
`
}
func contentApprovalParse(t *testing.T, text string) c.Change {
	t.Helper()
	ch, err := c.ParseChange(text, contentApprovalPath, "ALP")
	if err != nil {
		t.Fatalf("valid change fixture: %v\n%s", err, text)
	}
	return ch
}
func contentApprovalEdit(t *testing.T, text string, edit func(*c.Document)) string {
	t.Helper()
	doc, err := c.ParseMarkdown(text)
	if err != nil {
		t.Fatal(err)
	}
	edit(&doc)
	out, err := c.RenderMarkdown(doc)
	if err != nil {
		t.Fatal(err)
	}
	return out
}
func contentApprovalRuntime(t *testing.T, text string) string {
	return contentApprovalEdit(t, text, func(d *c.Document) {
		d.Metadata["run"] = "started-run"
		d.Metadata["phase"] = "audit"
		d.Metadata["blocked"] = "waiting for feedback"
	})
}
func contentApprovalPair(t *testing.T) (c.Module, string, string) {
	t.Helper()
	module := eligibleModule(t, "alpha")
	base := contentApprovalText()
	ch := contentApprovalParse(t, base)
	s := p.Snapshot{Modules: []c.Module{module}, Changes: []c.Change{ch}}
	assigned, bindings, err := p.Allocate(s, "current")
	if err != nil || bindings["new:first"] != "ALP-002" || bindings["new:first-long"] != "ALP-003" {
		t.Fatalf("public late allocation control: %+v %v", bindings, err)
	}
	text, err := c.RenderMarkdown(assigned.Document)
	if err != nil {
		t.Fatal(err)
	}
	current := contentApprovalRuntime(t, text)
	parsed := contentApprovalParse(t, current)
	if len(parsed.Added) != 2 || parsed.Added[1].ID != "ALP-003" || !strings.Contains(current, "check ALP-002 ALP-003") || !strings.Contains(current, "Author preamble ALP-002 and ALP-003.") {
		t.Fatal("complete late binding premise")
	}
	return module, base, current
}
func contentApprovalAssert(t *testing.T, module c.Module, base, current string, want bool) {
	t.Helper()
	before := policyJSON(t, module)
	got, err := w.ApprovalMatches(module, base, current, contentApprovalPath)
	if got != want || err != nil {
		t.Errorf("ApprovalMatches=%v,%v want %v,nil", got, err, want)
	}
	if policyJSON(t, module) != before {
		t.Error("module input mutated")
	}
}
func contentApprovalError(t *testing.T, module c.Module, base, current, path, detail string) {
	t.Helper()
	before := policyJSON(t, module)
	got, err := w.ApprovalMatches(module, base, current, path)
	if got || err == nil || !strings.HasPrefix(err.Error(), "C05:") || !policyCorrective(err.Error()) {
		t.Errorf("invalid current result %v,%v", got, err)
	}
	if err != nil && detail != "" && !strings.Contains(err.Error(), detail) {
		t.Errorf("lost parser detail %q: %v", detail, err)
	}
	if policyJSON(t, module) != before {
		t.Error("rejected module input mutated")
	}
}
func Test_FLOW028ApprovalInputAndMismatchResults(t *testing.T) {
	m, base, current := contentApprovalPair(t)
	contentApprovalAssert(t, m, base, base, true)
	contentApprovalAssert(t, m, base, current, true)
	for _, badBase := range []string{"", "---\nmodule: [\n---\n", "not a change", strings.Replace(base, "module: alpha", "module: bravo", 1)} {
		contentApprovalAssert(t, m, badBase, current, false)
	}
	for _, row := range []struct{ name, text, path string }{
		{"syntax", "---\nmodule: [\n---\n", contentApprovalPath},
		{"unknown-metadata", strings.Replace(base, "module: alpha", "module: alpha\nunknown: true", 1), contentApprovalPath},
		{"invalid-tag", strings.Replace(base, "[review]", "[unknown]", 1), contentApprovalPath},
		{"invalid-path", base, "changes/1-current.md"},
		{"undeclared-handle", strings.Replace(base, "Expose new:first", "Expose new:absent", 1), contentApprovalPath},
	} {
		t.Run(row.name, func(t *testing.T) {
			_, err := c.ParseChange(row.text, row.path, "ALP")
			if err == nil {
				t.Fatal("public invalid current premise")
			}
			for _, b := range []string{base, "", "bad base"} {
				contentApprovalError(t, m, b, row.text, row.path, err.Error())
			}
		})
	}
	wrong := strings.Replace(base, "module: alpha", "module: bravo", 1)
	contentApprovalParse(t, wrong)
	contentApprovalError(t, m, base, wrong, contentApprovalPath, "")
	wrongPrefix := m
	wrongPrefix.IDPrefix = "BET"
	_, prefixError := c.ParseChange(current, contentApprovalPath, wrongPrefix.IDPrefix)
	if prefixError == nil {
		t.Fatal("module prefix error premise")
	}
	contentApprovalError(t, wrongPrefix, base, current, contentApprovalPath, prefixError.Error())
}
func Test_FLOW029ApprovalAuthoredMetadata(t *testing.T) {
	m := eligibleModule(t, "alpha")
	base := contentApprovalText()
	contentApprovalParse(t, base)
	edits := []struct {
		name string
		edit func(*c.Document)
	}{
		{"bases", func(d *c.Document) { d.Metadata["bases"].(map[string]any)["ALP-001"] = "00000000" }},
		{"fix-order", func(d *c.Document) { d.Metadata["fixes"] = []string{"ALP-004", "ALP-001"} }},
		{"already-order", func(d *c.Document) { d.Metadata["already_satisfied"] = []string{"new:first-long", "new:first"} }},
		{"regenerate", func(d *c.Document) { d.Metadata["regenerate"] = true }},
	}
	for _, field := range []string{"depends_on", "interface", "evals", "disposable", "preserve", "append_only"} {
		field := field
		edits = append(edits, struct {
			name string
			edit func(*c.Document)
		}{field, func(d *c.Document) { d.Metadata["set"].(map[string]any)[field] = []string{} }})
	}
	edits = append(edits, struct {
		name string
		edit func(*c.Document)
	}{"dependency-order", func(d *c.Document) { d.Metadata["set"].(map[string]any)["depends_on"] = []string{"delta", "bravo"} }})
	edits = append(edits, struct {
		name string
		edit func(*c.Document)
	}{"opt-out", func(d *c.Document) { d.Metadata["set"].(map[string]any)["boundaries_opt_out"] = "different reason" }})
	for _, field := range []string{"eval", "boundaries", "setup"} {
		field := field
		edits = append(edits, struct {
			name string
			edit func(*c.Document)
		}{"hook-" + field, func(d *c.Document) {
			d.Metadata["set"].(map[string]any)["hooks"].(map[string]any)[field] = "different-hook"
		}})
	}
	for _, field := range []string{"depends_on", "interface", "evals", "disposable", "preserve", "append_only"} {
		field := field
		edits = append(edits, struct {
			name string
			edit func(*c.Document)
		}{"ordered-" + field, func(d *c.Document) {
			values := d.Metadata["set"].(map[string]any)[field].([]any)
			d.Metadata["set"].(map[string]any)[field] = []any{values[1], values[0]}
		}})
	}
	for _, row := range edits {
		t.Run(row.name, func(t *testing.T) {
			changed := contentApprovalEdit(t, base, row.edit)
			contentApprovalParse(t, changed)
			contentApprovalAssert(t, m, base, changed, false)
		})
	}
	for _, runtimeField := range []string{"run", "phase", "blocked"} {
		t.Run("runtime-"+runtimeField, func(t *testing.T) {
			changed := contentApprovalEdit(t, base, func(d *c.Document) {
				d.Metadata[runtimeField] = map[string]string{"run": "other-run", "phase": "impl", "blocked": "   "}[runtimeField]
			})
			contentApprovalParse(t, changed)
			contentApprovalAssert(t, m, base, changed, true)
		})
	}
	simple := "---\nmodule: alpha\n---\n\n## Motivation\n\nKeep the same intent.\n"
	explicit := contentApprovalEdit(t, simple, func(d *c.Document) {
		d.Metadata["set"] = map[string]any{}
		d.Metadata["bases"] = map[string]any{}
		d.Metadata["fixes"] = []string{}
		d.Metadata["already_satisfied"] = []string{}
		d.Metadata["regenerate"] = false
	})
	contentApprovalParse(t, simple)
	contentApprovalParse(t, explicit)
	contentApprovalAssert(t, m, simple, explicit, true)
	emptySet := contentApprovalEdit(t, simple, func(d *c.Document) { d.Metadata["set"] = map[string]any{"interface": []string{}} })
	contentApprovalParse(t, emptySet)
	contentApprovalAssert(t, m, simple, emptySet, false)
	for _, row := range []struct{ name, meta string }{
		{"empty-hook", "set: {hooks: {eval: ''}}\n"},
		{"empty-opt-out", "set: {boundaries_opt_out: ''}\n"},
		{"empty-hooks-map", "set: {hooks: {}}\n"},
	} {
		t.Run(row.name, func(t *testing.T) {
			changed := strings.Replace(simple, "module: alpha\n", "module: alpha\n"+row.meta, 1)
			contentApprovalParse(t, changed)
			contentApprovalAssert(t, m, simple, changed, false)
		})
	}
}
func Test_FLOW029ApprovalDocumentStructureAndProse(t *testing.T) {
	m := eligibleModule(t, "alpha")
	base := contentApprovalText()
	for _, name := range []string{"Motivation", "Interface", "Added", "Modified", "Removed", "Removed decisions", "Rejected alternatives", "Author guidance"} {
		name := name
		t.Run(name, func(t *testing.T) {
			changed := contentApprovalEdit(t, base, func(d *c.Document) {
				for i := range d.Sections {
					if d.Sections[i].Name == name {
						if name == "Removed" {
							d.Sections[i].Body = "### ALP-008"
						} else {
							d.Sections[i].Body = strings.Replace(d.Sections[i].Body, "The target MUST retain", "The target MUST additionally retain", 1)
							if name != "Added" {
								d.Sections[i].Body += "\n\nAdditional authored content."
							}
						}
					}
				}
			})
			contentApprovalParse(t, changed)
			contentApprovalAssert(t, m, base, changed, false)
		})
	}
	for _, mode := range []string{"preamble", "tag", "name", "order", "new-empty", "remove-empty", "interior-space", "same-hash-author-text", "fenced-exclusion", "case-exclusion"} {
		t.Run(mode, func(t *testing.T) {
			left := base
			changed := contentApprovalEdit(t, base, func(d *c.Document) {
				switch mode {
				case "preamble":
					d.Preamble += "Changed author preamble."
				case "tag":
					for i := range d.Sections {
						if d.Sections[i].Name == "Added" {
							d.Sections[i].Body = strings.Replace(d.Sections[i].Body, "### new:first\n", "### new:first [review]\n", 1)
						}
					}
				case "name":
					for i := range d.Sections {
						if d.Sections[i].Name == "Author guidance" {
							d.Sections[i].Name = "Other author guidance"
						}
					}
				case "order":
					d.Sections[0], d.Sections[1] = d.Sections[1], d.Sections[0]
				case "new-empty":
					d.Sections = append(d.Sections, c.Section{Name: "Empty authored section"})
				case "remove-empty":
					d.Sections = append(d.Sections, c.Section{Name: "Empty authored section"})
				case "interior-space":
					for i := range d.Sections {
						if d.Sections[i].Name == "Added" {
							d.Sections[i].Body = strings.Replace(d.Sections[i].Body, "The target MUST retain", "The  target MUST retain", 1)
						}
					}
				case "same-hash-author-text":
					for i := range d.Sections {
						if d.Sections[i].Name == "Added" {
							d.Sections[i].Body = strings.Replace(d.Sections[i].Body, "exactly.\n", "exactly. \t\n", 1)
						}
					}
				case "fenced-exclusion":
					for i := range d.Sections {
						if d.Sections[i].Name == "Author guidance" {
							d.Sections[i].Body = strings.Replace(d.Sections[i].Body, "stays an author example", "changes the author example", 1)
						}
					}
				case "case-exclusion":
					d.Sections = append(d.Sections, c.Section{Name: "notes", Body: "Different casing remains authored."})
				}
			})
			contentApprovalParse(t, changed)
			if mode == "remove-empty" {
				left, changed = changed, base
			}
			if mode == "same-hash-author-text" {
				a, b := contentApprovalParse(t, base), contentApprovalParse(t, changed)
				if c.SpecHash(a.Added[0]) != c.SpecHash(b.Added[0]) {
					t.Fatal("same normalized spec-hash premise")
				}
			}
			contentApprovalAssert(t, m, left, changed, false)
		})
	}
	for _, name := range []string{"Decisions", "Questions", "Audit", "Provenance", "Notes"} {
		name := name
		t.Run("excluded-"+name, func(t *testing.T) {
			changed := contentApprovalEdit(t, base, func(d *c.Document) {
				for i := range d.Sections {
					if d.Sections[i].Name == name {
						d.Sections[i].Body = "Different worker text.\n\n```markdown\n## Motivation\nFenced ignored text.\n```"
						d.Sections = append(d.Sections[:i], d.Sections[i+1:]...)
						d.Sections = append([]c.Section{{Name: name, Body: "Different moved worker text."}}, d.Sections...)
						break
					}
				}
			})
			contentApprovalParse(t, changed)
			contentApprovalAssert(t, m, base, changed, true)
			removed := contentApprovalEdit(t, base, func(d *c.Document) {
				for i := range d.Sections {
					if d.Sections[i].Name == name {
						d.Sections = append(d.Sections[:i], d.Sections[i+1:]...)
						break
					}
				}
			})
			contentApprovalParse(t, removed)
			contentApprovalAssert(t, m, base, removed, true)
		})
	}
}
func Test_FLOW029ApprovalExactExcludedNames(t *testing.T) {
	m := eligibleModule(t, "alpha")
	base := contentApprovalText()
	for _, name := range []string{"Decisions", "Questions", "Audit", "Provenance", "Notes"} {
		t.Run(name, func(t *testing.T) {
			changed := contentApprovalEdit(t, base, func(d *c.Document) {
				d.Sections = append(d.Sections, c.Section{Name: strings.ToLower(name), Body: "Author text remains significant."})
			})
			contentApprovalParse(t, changed)
			contentApprovalAssert(t, m, base, changed, false)
		})
	}
}

func Test_FLOW029ApprovalPresentationNormalization(t *testing.T) {
	m := eligibleModule(t, "alpha")
	base := "---\nmodule: alpha\nfixes: [ALP-001, ALP-002]\nregenerate: false\n---\n\n## Motivation\n\nKeep exact authored text.\n"
	current := "---\n# Presentation only\nregenerate: false\nfixes:\n - 'ALP-001'\n - \"ALP-002\"\nmodule: 'alpha'\n---\n\n\n## Motivation\n\n\nKeep exact authored text.\n\n\n"
	contentApprovalParse(t, base)
	contentApprovalParse(t, current)
	contentApprovalAssert(t, m, base, strings.ReplaceAll(current, "\n", "\r\n"), true)
}
func Test_FLOW030ApprovalAllLateBindingLocations(t *testing.T) {
	m, base, current := contentApprovalPair(t)
	contentApprovalAssert(t, m, base, current, true)
	reordered := contentApprovalEdit(t, current, func(d *c.Document) {
		parsed := contentApprovalParse(t, current)
		a, b := parsed.Added[0], parsed.Added[1]
		for i := range d.Sections {
			if d.Sections[i].Name == "Added" {
				d.Sections[i].Body = "### " + b.ID + " [review]\n" + b.Body + "\n\n### " + a.ID + "\n" + a.Body
			}
		}
	})
	contentApprovalParse(t, reordered)
	contentApprovalAssert(t, m, base, reordered, false)
	sameIDsDifferentIntent := strings.Replace(current, "Keep approval tied", "Different approval tied", 1)
	contentApprovalParse(t, sameIDsDifferentIntent)
	contentApprovalAssert(t, m, current, sameIDsDifferentIntent, false)
	for _, marker := range []string{"check ALP-002 ALP-003", "Author preamble ALP-002 and ALP-003.", "Expose ALP-002 and ALP-003.", "ALP-002 stays an author example."} {
		if !strings.Contains(current, marker) {
			t.Fatal("missing binding premise", marker)
		}
	}
	changed := strings.Replace(current, "Expose ALP-002 and ALP-003.", "Expose ALP-003 and ALP-002.", 1)
	contentApprovalParse(t, changed)
	contentApprovalAssert(t, m, base, changed, false)
	changed = contentApprovalEdit(t, current, func(d *c.Document) {
		for i := range d.Sections {
			if d.Sections[i].Name == "Added" {
				d.Sections[i].Body += "\n\n### ALP-020\nThe extra item MUST differ."
			}
		}
	})
	contentApprovalParse(t, changed)
	contentApprovalAssert(t, m, base, changed, false)
	fixed := strings.ReplaceAll(current, "ALP-003", "ALP-004")
	contentApprovalParse(t, fixed)
	contentApprovalAssert(t, m, current, fixed, false)
	// The first binding's result is itself the second binding's old handle.
	cascadeBase := "---\nmodule: alpha\nalready_satisfied: [new:a, new:b]\n---\n\nPreamble new:a and new:b.\n\n## Motivation\n\nnew:a MUST differ from new:b.\n\n## Added\n\n### new:a\nThis MUST keep new:b.\n\n### new:b\nThis MUST keep new:a.\n\n## new:a\n\nAuthor section new:b.\n"
	cascadeCurrent := "---\nmodule: alpha\nalready_satisfied: [new:b, ALP-020]\n---\n\nPreamble new:b and ALP-020.\n\n## Motivation\n\nnew:b MUST differ from ALP-020.\n\n## Added\n\n### new:b\nThis MUST keep ALP-020.\n\n### ALP-020\nThis MUST keep new:b.\n\n## new:b\n\nAuthor section ALP-020.\n"
	contentApprovalParse(t, cascadeBase)
	contentApprovalParse(t, cascadeCurrent)
	contentApprovalAssert(t, m, cascadeBase, cascadeCurrent, true)
}
func Test_FLOW031ApprovalOwnershipAndAmbientIndependence(t *testing.T) {
	m, base, current := contentApprovalPair(t)
	_, pristineError := w.ApprovalMatches(m, base, "invalid", contentApprovalPath)
	if pristineError == nil {
		t.Fatal("invalid current unexpectedly accepted")
	}
	backing := []string{"private-spare-one", "private-spare-two"}
	m.DependsOn = backing[:0]
	m.Document.Metadata["depends_on"] = backing[:0]
	m.Hooks = c.Hooks{Eval: "hook-trap", Boundaries: "hook-trap", Setup: "hook-trap"}
	m.Specs = []c.Spec{{ID: "invalid", Tag: "invalid", Body: "irrelevant"}}
	m.Zones = c.Zones{Interface: []string{"bad["}, Evals: []string{"overlap/**"}, Disposable: []string{"overlap/**"}}
	m.Dir = "/private/not-a-repository"
	m.Decisions = "irrelevant"
	before := policyJSON(t, m)
	expected := append([]string(nil), backing...)
	different := strings.Replace(current, "Keep approval tied", "Different approval tied", 1)
	contentApprovalParse(t, different)
	_, parseError := c.ParseChange("invalid", contentApprovalPath, "ALP")
	if parseError == nil {
		t.Fatal("error fixture control")
	}
	var initialError string
	originalCwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	originalEnv := os.Environ()
	for i := 0; i < 2; i++ {
		contentApprovalAssert(t, m, base, current, true)
		contentApprovalAssert(t, m, base, different, false)
		contentApprovalError(t, m, base, "invalid", contentApprovalPath, parseError.Error())
		_, err := w.ApprovalMatches(m, base, "invalid", contentApprovalPath)
		if err == nil {
			t.Fatal("missing invalid-current error")
		}
		if err.Error() != pristineError.Error() {
			t.Errorf("irrelevant module fields changed error: %v/%v", err, pristineError)
		}
		if i == 0 {
			initialError = err.Error()
		} else if err.Error() != initialError {
			t.Fatal("nondeterministic error")
		}
	}
	if before != policyJSON(t, m) || !reflect.DeepEqual(backing, expected) {
		t.Fatal("nested/backing source mutation")
	}
	sameCwd, err := os.Getwd()
	if err != nil || sameCwd != originalCwd || !reflect.DeepEqual(originalEnv, os.Environ()) {
		t.Fatal("initial in-process cwd/environment mutation")
	}
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	dir, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	if err = os.WriteFile(filepath.Join(dir, "fireweed.yaml"), []byte("malformed: ["), 0644); err != nil {
		t.Fatal(err)
	}
	if err = os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(cwd)
	t.Setenv("FIREWEED_BIN", "must-not-run")
	t.Setenv("FIREWEED_APPROVAL_DECOY", "changed")
	expectedEnv := os.Environ()
	contentApprovalAssert(t, m, base, current, true)
	contentApprovalAssert(t, m, base, different, false)
	contentApprovalError(t, m, base, "invalid", contentApprovalPath, parseError.Error())
	_, err = w.ApprovalMatches(m, base, "invalid", contentApprovalPath)
	if err == nil || err.Error() != initialError {
		t.Fatal("ambient changed error", err)
	}
	after, err := os.Getwd()
	if err != nil || after != dir || !reflect.DeepEqual(expectedEnv, os.Environ()) || before != policyJSON(t, m) || !reflect.DeepEqual(backing, expected) {
		t.Fatal("same-process or input mutation")
	}
	if raw, err := os.ReadFile(filepath.Join(dir, "fireweed.yaml")); err != nil || string(raw) != "malformed: [" {
		t.Fatal("ambient file mutation", err)
	}
}
func Test_FLOW028_FLOW030ApprovalFixtureControls(t *testing.T) {
	m, base, current := contentApprovalPair(t)
	a, b := contentApprovalParse(t, base), contentApprovalParse(t, current)
	if m.Name != "alpha" || len(a.Added) != 2 || a.Added[0].ID != "new:first" || b.Added[0].ID != "ALP-002" || b.Phase != "audit" {
		t.Fatal("public approval fixture premise")
	}
	if _, err := c.RenderMarkdown(a.Document); err != nil {
		t.Fatal(fmt.Errorf("public render control: %w", err))
	}
}
