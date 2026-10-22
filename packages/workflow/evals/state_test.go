package evals

import (
	"encoding/json"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	k "github.com/lurkshark/fireweed/packages/checks"
	c "github.com/lurkshark/fireweed/packages/contracts"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

const statePath = "changes/0042-intent.md"

func stateModule() c.Module {
	return c.Module{Name: "alpha", IDPrefix: "ALP", Dir: "units/alpha", DependsOn: []string{"beta"}, Zones: c.Zones{Interface: []string{"api.go"}, Evals: []string{"evals/**"}, Disposable: []string{"internal/**"}, Preserve: []string{"data/**"}, AppendOnly: []string{"migrations/**"}}, Hooks: c.Hooks{Eval: "hook-trap", Boundaries: "hook-trap", Setup: "hook-trap"}, Specs: []c.Spec{{ID: "ALP-001", Body: "It MUST remain."}}, Decisions: "Retain this.", Document: c.Document{Metadata: map[string]any{"name": "alpha", "hooks": map[string]any{"eval": "hook-trap", "boundaries": "hook-trap", "setup": "hook-trap"}, "interface": []any{"api.go"}}, Preamble: "# Alpha", Sections: []c.Section{{Name: "Decisions", Body: "Retain this."}}}}
}
func stateFixture(t *testing.T, phase string) string {
	t.Helper()
	text := `---
module: alpha
set:
  depends_on: [beta]
  interface: [api.go]
  evals: [evals/**]
  disposable: [internal/**]
  preserve: [data/**]
  append_only: [migrations/**]
  boundaries_opt_out: legacy integration
  hooks: {eval: 'printf eval', boundaries: 'printf boundary', setup: 'printf setup'}
bases: {ALP-001: abcdef12, ALP-002: 1234abcd}
fixes: [ALP-001]
already_satisfied: [ALP-004]
regenerate: false
run: started-run
phase: PHASE
blocked: initial question
---

# Approved intent

## Motivation

Retain the actual requirement.

## Interface

Expose the approved interface.

## Added

### ALP-004
It MUST add exactly this behavior.

## Modified

### ALP-001
It MUST retain the approved correction.

## Removed

### ALP-002

## Decisions

Before reason.

` + "```md\n## Motivation\nExample heading in a fence.\n```" + `

## Removed decisions

> Remove this historical reason.

## Rejected alternatives

Keep the original tradeoff.

## Notes

Before notes.

## Questions

Before question.

## Audit

Before audit.

## Provenance

- evals | original record

## Unknown human section

Keep this too.
`
	text = strings.Replace(text, "PHASE", phase, 1)
	if _, err := c.ParseChange(text, statePath, "ALP"); err != nil {
		t.Fatalf("invalid state fixture %v", err)
	}
	return text
}
func stateDoc(t *testing.T, text string) c.Document {
	t.Helper()
	doc, err := c.ParseMarkdown(text)
	if err != nil {
		t.Fatal(err)
	}
	return doc
}
func stateRender(t *testing.T, doc c.Document) string {
	t.Helper()
	text, err := c.RenderMarkdown(doc)
	if err != nil {
		t.Fatal(err)
	}
	return text
}
func stateAlter(t *testing.T, text string, change func(*c.Document)) string {
	t.Helper()
	doc := stateDoc(t, text)
	change(&doc)
	return stateRender(t, doc)
}
func stateSection(doc *c.Document, name string, body string) {
	for i := range doc.Sections {
		if doc.Sections[i].Name == name {
			doc.Sections[i].Body = body
			return
		}
	}
	doc.Sections = append(doc.Sections, c.Section{Name: name, Body: body})
}
func stateRemove(doc *c.Document, name string) {
	sections := []c.Section{}
	for _, section := range doc.Sections {
		if section.Name != name {
			sections = append(sections, section)
		}
	}
	doc.Sections = sections
}
func stateBody(t *testing.T, text, name string) string {
	t.Helper()
	for _, s := range stateDoc(t, text).Sections {
		if s.Name == name {
			return s.Body
		}
	}
	t.Fatalf("missing section %s", name)
	return ""
}
func stateJSON(t *testing.T, value any) string {
	t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}
func stateAssert(t *testing.T, ds []k.Diagnostic, code, path string, count int) {
	t.Helper()
	if ds == nil || len(ds) != count {
		t.Fatalf("diagnostics %+v expected %d %s", ds, count, code)
	}
	keys := []string{}
	for _, d := range ds {
		lower := strings.ToLower(d.Message)
		if d.Code != code || d.Path != path || d.Severity != "error" || strings.TrimSpace(d.Message) == "" || !(strings.Contains(lower, "correct") || strings.Contains(lower, "restore") || strings.Contains(lower, "repair") || strings.Contains(lower, "resume")) {
			t.Fatalf("non-corrective/path error %+v", d)
		}
		keys = append(keys, d.Path+"\x00"+d.Code+"\x00"+d.Severity+"\x00"+d.Message)
	}
	if !sort.StringsAreSorted(keys) {
		t.Fatalf("unsorted diagnostics %+v", ds)
	}
}
func Test_FLOW001UnchangedStartedStates(t *testing.T) {
	for _, phase := range []string{"evals", "audit", "impl"} {
		for _, kind := range []string{"spec", "maintenance", "regenerate"} {
			t.Run(phase+"_"+kind, func(t *testing.T) {
				text := stateFixture(t, phase)
				if kind != "spec" {
					meta := ""
					if kind == "regenerate" {
						meta = "regenerate: true\n"
					}
					text = "---\nmodule: alpha\nrun: started\nphase: " + phase + "\n" + meta + "---\n\n## Motivation\n\nPerform approved maintenance.\n"
				}
				parsed, err := c.ParseChange(text, statePath, "ALP")
				if err != nil || parsed.Kind != kind {
					t.Fatalf("kind premise %s %+v %v", kind, parsed, err)
				}
				stateAssert(t, w.CheckState(stateModule(), text, text, statePath), "", statePath, 0)
			})
		}
	}
}
func Test_FLOW001InvalidApprovalAndCurrent(t *testing.T) {
	good := stateFixture(t, "evals")
	for _, mode := range []string{"empty", "malformed", "wrong-prefix", "unknown-metadata", "wrong-module", "missing-run", "empty-run", "blank-run", "missing-phase", "empty-phase", "applied-phase", "unknown-phase", "bad-path"} {
		t.Run("approved_"+mode, func(t *testing.T) {
			before := good
			module := stateModule()
			path := statePath
			switch mode {
			case "empty":
				before = ""
			case "malformed":
				before = "---\nmodule: [\n---\n## Motivation\nBroken"
			case "wrong-prefix":
				module.IDPrefix = "BET"
			case "unknown-metadata":
				before = strings.Replace(before, "module: alpha", "unknown: true\nmodule: alpha", 1)
			case "wrong-module":
				module.Name = "Alpha"
			case "missing-run":
				before = stateAlter(t, before, func(d *c.Document) { delete(d.Metadata, "run") })
			case "empty-run":
				before = stateAlter(t, before, func(d *c.Document) { d.Metadata["run"] = "" })
			case "blank-run":
				before = stateAlter(t, before, func(d *c.Document) { d.Metadata["run"] = " \t" })
			case "empty-phase":
				before = stateAlter(t, before, func(d *c.Document) { d.Metadata["phase"] = "" })
			case "missing-phase":
				before = stateAlter(t, before, func(d *c.Document) { delete(d.Metadata, "phase") })
			case "applied-phase":
				before = stateAlter(t, before, func(d *c.Document) { d.Metadata["phase"] = "applied" })
			case "unknown-phase":
				before = stateAlter(t, before, func(d *c.Document) { d.Metadata["phase"] = "queued" })
			case "bad-path":
				path = "changes/not-numbered.md"
			}
			ds := w.CheckState(module, before, good, path)
			stateAssert(t, ds, "P02", path, 1)
			if !strings.Contains(strings.ToLower(ds[0].Message), "approv") && !strings.Contains(strings.ToLower(ds[0].Message), "claim") {
				t.Fatalf("missing claim guidance %+v", ds)
			}
		})
	}
	for _, mode := range []string{"empty", "malformed", "unknown-metadata", "wrong-prefix", "missing-motivation", "nonnormative", "duplicate-id", "invalid-and-metadata"} {
		t.Run("current_"+mode, func(t *testing.T) {
			current := good
			switch mode {
			case "empty":
				current = ""
			case "malformed":
				current = "---\nmodule: [\n---\n## Motivation\nBroken"
			case "unknown-metadata":
				current = strings.Replace(current, "module: alpha", "unknown: true\nmodule: alpha", 1)
			case "wrong-prefix":
				current = strings.Replace(current, "### ALP-004", "### BET-004", 1)
			case "missing-motivation":
				current = stateAlter(t, current, func(d *c.Document) { stateRemove(d, "Motivation") })
			case "nonnormative":
				current = strings.Replace(current, "It MUST add exactly this behavior.", "Not a normative requirement.", 1)
			case "duplicate-id":
				current = stateAlter(t, current, func(d *c.Document) {
					stateSection(d, "Added", "### ALP-004\nIt MUST work.\n\n### ALP-004\nIt MUST duplicate.")
				})
			case "invalid-and-metadata":
				current = strings.Replace(current, "run: started-run", "run: worker-owned", 1)
				current = strings.Replace(current, "It MUST add exactly this behavior.", "Not normative.", 1)
			}
			if _, err := c.ParseChange(current, statePath, "ALP"); err == nil {
				t.Fatal("invalid-current premise")
			}
			ds := w.CheckState(stateModule(), good, current, statePath)
			stateAssert(t, ds, "P2-01", statePath, 1)
			if !strings.Contains(strings.ToLower(ds[0].Message), "worker") && !strings.Contains(strings.ToLower(ds[0].Message), "edit") {
				t.Fatalf("missing worker guidance %+v", ds)
			}
		})
	}
	stateAssert(t, w.CheckState(stateModule(), "bad", "bad", statePath), "P02", statePath, 1)
}
func Test_FLOW002MetadataAndExplicitKeys(t *testing.T) {
	before := stateFixture(t, "evals")
	mutations := map[string]func(*c.Document){
		"module":            func(d *c.Document) { d.Metadata["module"] = "beta" },
		"bases":             func(d *c.Document) { d.Metadata["bases"].(map[string]any)["ALP-001"] = "11111111" },
		"fixes":             func(d *c.Document) { d.Metadata["fixes"] = []any{"ALP-003"} },
		"already_satisfied": func(d *c.Document) { d.Metadata["already_satisfied"] = []any{"ALP-001"} },
		"regenerate":        func(d *c.Document) { d.Metadata["regenerate"] = true },
		"run":               func(d *c.Document) { d.Metadata["run"] = "worker-owned" },
		"phase":             func(d *c.Document) { d.Metadata["phase"] = "impl" },
	}
	for _, field := range []string{"depends_on", "interface", "evals", "disposable", "preserve", "append_only"} {
		field := field
		mutations["set_"+field] = func(d *c.Document) {
			value := "changed/**"
			if field == "depends_on" {
				value = "omega"
			}
			d.Metadata["set"].(map[string]any)[field] = []any{value}
		}
	}
	mutations["set_boundaries_opt_out"] = func(d *c.Document) { d.Metadata["set"].(map[string]any)["boundaries_opt_out"] = "worker exemption" }
	for _, field := range []string{"eval", "boundaries", "setup"} {
		field := field
		mutations["hook_"+field] = func(d *c.Document) {
			d.Metadata["set"].(map[string]any)["hooks"].(map[string]any)[field] = "worker command"
		}
	}
	for _, field := range []string{"set", "bases", "fixes", "already_satisfied", "regenerate", "run", "phase"} {
		field := field
		mutations["remove_"+field] = func(d *c.Document) { delete(d.Metadata, field) }
	}
	names := []string{}
	for name := range mutations {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			current := stateAlter(t, before, mutations[name])
			if _, err := c.ParseChange(current, statePath, "ALP"); err != nil {
				t.Fatalf("valid metadata edit premise %v", err)
			}
			stateAssert(t, w.CheckState(stateModule(), before, current, statePath), "P2-01", statePath, 1)
		})
	}
	for _, field := range []string{"fixes", "already_satisfied", "set", "regenerate"} {
		t.Run("add_explicit_"+field, func(t *testing.T) {
			omitted := stateAlter(t, before, func(d *c.Document) { delete(d.Metadata, field) })
			current := stateAlter(t, omitted, func(d *c.Document) {
				switch field {
				case "set":
					d.Metadata[field] = map[string]any{}
				case "regenerate":
					d.Metadata[field] = false
				default:
					d.Metadata[field] = []any{}
				}
			})
			if _, err := c.ParseChange(current, statePath, "ALP"); err != nil {
				t.Fatal(err)
			}
			stateAssert(t, w.CheckState(stateModule(), omitted, current, statePath), "P2-01", statePath, 1)
		})
	}
}
func Test_FLOW002BlockedAndPermittedSections(t *testing.T) {
	for _, phase := range []string{"evals", "audit", "impl"} {
		before := stateFixture(t, phase)
		for _, action := range []string{"add", "replace", "remove"} {
			t.Run(phase+"_blocked_"+action, func(t *testing.T) {
				approved := before
				if action == "add" {
					approved = stateAlter(t, approved, func(d *c.Document) { delete(d.Metadata, "blocked") })
				}
				current := stateAlter(t, approved, func(d *c.Document) {
					if action == "remove" {
						delete(d.Metadata, "blocked")
					} else {
						d.Metadata["blocked"] = "worker question needing an answer"
					}
				})
				stateAssert(t, w.CheckState(stateModule(), approved, current, statePath), "", statePath, 0)
			})
		}
		allowed := []string{"Decisions", "Questions", "Notes"}
		if phase == "audit" {
			allowed = []string{"Audit", "Questions"}
		}
		for _, name := range allowed {
			for _, action := range []string{"add", "change", "move", "remove"} {
				t.Run(phase+"_"+name+"_"+action, func(t *testing.T) {
					approved := before
					if action == "add" {
						approved = stateAlter(t, approved, func(d *c.Document) { stateRemove(d, name) })
					}
					current := stateAlter(t, approved, func(d *c.Document) {
						switch action {
						case "add", "change":
							stateSection(d, name, "Worker prose.\n\n```md\n## Motivation\nThis is a fenced example.\n```")
						case "remove":
							stateRemove(d, name)
						case "move":
							var selected c.Section
							for _, s := range d.Sections {
								if s.Name == name {
									selected = s
								}
							}
							stateRemove(d, name)
							d.Sections = append([]c.Section{selected}, d.Sections...)
						}
					})
					if _, err := c.ParseChange(current, statePath, "ALP"); err != nil {
						t.Fatalf("allowed-section fixture %v", err)
					}
					stateAssert(t, w.CheckState(stateModule(), approved, current, statePath), "", statePath, 0)
				})
			}
		}
	}
}
func Test_FLOW002ProtectedSectionsAndEscalation(t *testing.T) {
	for _, phase := range []string{"evals", "audit", "impl"} {
		before := stateFixture(t, phase)
		forbidden := []string{"Motivation", "Interface", "Added", "Modified", "Removed", "Removed decisions", "Rejected alternatives", "Provenance", "Unknown human section", "preamble", "protected-order", "new-unknown", "case-sensitive"}
		if phase == "audit" {
			forbidden = append(forbidden, "Decisions", "Notes")
		} else {
			forbidden = append(forbidden, "Audit")
		}
		for _, name := range forbidden {
			t.Run(phase+"_"+name, func(t *testing.T) {
				current := stateAlter(t, before, func(d *c.Document) {
					switch name {
					case "preamble":
						d.Preamble += "\nWorker changed the approval."
					case "protected-order":
						d.Sections[0], d.Sections[1] = d.Sections[1], d.Sections[0]
					case "new-unknown":
						d.Sections = append(d.Sections, c.Section{Name: "Worker invented section", Body: "Changed approval."})
					case "case-sensitive":
						d.Sections = append(d.Sections, c.Section{Name: "decisions", Body: "Lowercase is not permission."})
					case "Removed":
						stateSection(d, name, "### ALP-003")
					default:
						for i := range d.Sections {
							if d.Sections[i].Name == name {
								d.Sections[i].Body += "\nWorker changed this requirement."
								break
							}
						}
					}
				})
				if _, err := c.ParseChange(current, statePath, "ALP"); err != nil {
					t.Fatalf("protected mutation must remain parseable: %v", err)
				}
				stateAssert(t, w.CheckState(stateModule(), before, current, statePath), "P2-01", statePath, 1)
			})
		}
		t.Run(phase+"_phase_escalation", func(t *testing.T) {
			current := stateAlter(t, before, func(d *c.Document) {
				if phase == "audit" {
					d.Metadata["phase"] = "impl"
					stateSection(d, "Decisions", "Auditor self-granted implementation permission.")
				} else {
					d.Metadata["phase"] = "audit"
					stateSection(d, "Audit", "Worker self-granted audit permission.")
				}
			})
			stateAssert(t, w.CheckState(stateModule(), before, current, statePath), "P2-01", statePath, 2)
		})
	}
}
func Test_FLOW002FencesDuplicatesAndNormalization(t *testing.T) {
	before := stateFixture(t, "evals")
	// Fenced pseudo-headings are part of permitted prose, but a later actual protected section is not.
	permitted := stateAlter(t, before, func(d *c.Document) {
		stateSection(d, "Decisions", "```md\n## Notes\n## Motivation\nChanged example, not actual approval.\n```")
	})
	t.Run("fenced_allowed", func(t *testing.T) {
		stateAssert(t, w.CheckState(stateModule(), before, permitted, statePath), "", statePath, 0)
	})
	protected := stateAlter(t, permitted, func(d *c.Document) { stateSection(d, "Provenance", "Spoofed after fenced text.") })
	t.Run("actual_protected_after_fence", func(t *testing.T) {
		stateAssert(t, w.CheckState(stateModule(), before, protected, statePath), "P2-01", statePath, 1)
	})
	for _, name := range []string{"Decisions", "Unknown human section"} {
		t.Run("duplicate_"+name, func(t *testing.T) {
			approved := before + "\n## " + name + "\n\nSecond occurrence.\n"
			current := strings.Replace(approved, "Second occurrence.", "Worker changed the second occurrence.", 1)
			_, err := c.ParseChange(approved, statePath, "ALP")
			if err != nil {
				stateAssert(t, w.CheckState(stateModule(), approved, current, statePath), "P02", statePath, 1)
				return
			}
			count := 1
			if name == "Decisions" {
				count = 0
			}
			stateAssert(t, w.CheckState(stateModule(), approved, current, statePath), "P2-01", statePath, count)
		})
	}
	current := strings.Replace(before, "module: alpha\n", "module: 'alpha'\n", 1)
	current = strings.Replace(current, "run: started-run\nphase: evals\n", "phase: evals\nrun: started-run\n", 1)
	current = strings.ReplaceAll(current, "\n", "\r\n")
	if !reflect.DeepEqual(stateDoc(t, before), stateDoc(t, current)) {
		t.Fatal("public normalized-value premise")
	}
	t.Run("normalized_yaml_and_crlf", func(t *testing.T) {
		stateAssert(t, w.CheckState(stateModule(), before, current, statePath), "", statePath, 0)
	})
}
func Test_FLOW003BothViolationsAndResultOwnership(t *testing.T) {
	module := stateModule()
	before := stateFixture(t, "impl")
	current := stateAlter(t, before, func(d *c.Document) {
		d.Metadata["run"] = "worker"
		stateSection(d, "Motivation", "Worker changed approved motivation.")
	})
	if _, err := c.ParseChange(current, statePath, "ALP"); err != nil {
		t.Fatal(err)
	}
	savedModule, savedBefore, savedCurrent := stateJSON(t, module), before, current
	first, second := w.CheckState(module, before, current, statePath), w.CheckState(module, before, current, statePath)
	stateAssert(t, first, "P2-01", statePath, 2)
	stateAssert(t, second, "P2-01", statePath, 2)
	if !reflect.DeepEqual(first, second) || first[0].Message == first[1].Message {
		t.Fatal("lost distinct deterministic violations")
	}
	savedResult := stateJSON(t, second)
	first[0].Message = "mutate first result"
	first = append(first, k.Diagnostic{Message: "new result"})
	if savedResult != stateJSON(t, second) || savedModule != stateJSON(t, module) || before != savedBefore || current != savedCurrent {
		t.Fatal("result/input aliases or mutation")
	}
	emptyA, emptyB := w.CheckState(module, before, before, statePath), w.CheckState(module, before, before, statePath)
	stateAssert(t, emptyA, "", statePath, 0)
	stateAssert(t, emptyB, "", statePath, 0)
	emptyA = append(emptyA, k.Diagnostic{Message: "first empty"})
	emptyB = append(emptyB, k.Diagnostic{Message: "second empty"})
	if emptyA[0].Message != "first empty" || emptyB[0].Message != "second empty" {
		t.Fatal("empty results share writable storage")
	}
	invalidA, invalidB := w.CheckState(module, "bad", before, statePath), w.CheckState(module, "bad", before, statePath)
	stateAssert(t, invalidA, "P02", statePath, 1)
	stateAssert(t, invalidB, "P02", statePath, 1)
	savedInvalid := stateJSON(t, invalidB)
	invalidA[0].Path = "changed"
	if savedInvalid != stateJSON(t, invalidB) {
		t.Fatal("error results share storage")
	}
}
func stateFiles(t *testing.T, root string) map[string]string {
	t.Helper()
	files := map[string]string{}
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		relative, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		files[relative] = string(raw)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return files
}
func Test_FLOW003SuppliedTextAndSideEffectSafety(t *testing.T) {
	root := stateCanonicalDirectory(t, t.TempDir())
	trapRoot := t.TempDir()
	marker := filepath.Join(trapRoot, "unexpected")
	for path, text := range map[string]string{"sentinel": "user bytes", "ignored": "ignored bytes", ".git/refs/heads/sentinel": "not a real repository", ".gitignore": "ignored\n"} {
		full := filepath.Join(root, path)
		if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(text), 0644); err != nil {
			t.Fatal(err)
		}
	}
	script := "#!/bin/sh\nprintf attempted >> '" + strings.ReplaceAll(marker, "'", `'"'"'`) + "'\nexit 97\n"
	for _, name := range []string{"git", "hook-trap", "agent-trap", "codex", "claude", "fireweed", "fireweed-bootstrap", "curl", "wget", "ssh", "nc"} {
		if err := os.WriteFile(filepath.Join(trapRoot, name), []byte(script), 0755); err != nil {
			t.Fatal(err)
		}
	}
	t.Setenv("PATH", trapRoot+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("FIREWEED_BIN", filepath.Join(trapRoot, "fireweed-bootstrap"))
	control := exec.Command("git", "--version")
	if err := control.Run(); err == nil {
		t.Fatal("trap control unexpectedly succeeded")
	}
	if _, err := os.Stat(marker); err != nil {
		t.Fatal("trap control did not run")
	}
	if err := os.Remove(marker); err != nil {
		t.Fatal(err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(root); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(cwd)
	before := stateFixture(t, "audit")
	current := stateAlter(t, before, func(d *c.Document) { stateSection(d, "Decisions", "Forbidden auditor decision.") })
	module := stateModule()
	moduleBefore := stateJSON(t, module)
	filesBefore := stateFiles(t, root)
	envBefore := os.Environ()
	for _, tt := range []struct {
		approved, current, code string
		count                   int
	}{{before, before, "", 0}, {before, current, "P2-01", 1}, {"invalid", before, "P02", 1}, {before, "invalid", "P2-01", 1}} {
		t.Run(tt.code+"_"+tt.approved[:3]+"_"+tt.current[:3], func(t *testing.T) {
			stateAssert(t, w.CheckState(module, tt.approved, tt.current, statePath), tt.code, statePath, tt.count)
			afterCwd, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}
			if afterCwd != root || !reflect.DeepEqual(envBefore, os.Environ()) || moduleBefore != stateJSON(t, module) || !reflect.DeepEqual(filesBefore, stateFiles(t, root)) {
				t.Fatal("state guard mutated parent inputs or repository/user/ignored/ref files")
			}
			if _, err := os.Stat(marker); !os.IsNotExist(err) {
				t.Fatal("state guard executed Git/hooks/agents/bootstrap/network commands")
			}
		})
	}
}

func Test_FLOW001_FLOW002_FLOW003StateFixtureControl(t *testing.T) {
	for _, phase := range []string{"evals", "audit", "impl"} {
		text := stateFixture(t, phase)
		parsed, err := c.ParseChange(text, statePath, "ALP")
		if err != nil {
			t.Fatal(err)
		}
		if parsed.Module != "alpha" || parsed.Phase != phase || parsed.Run != "started-run" || parsed.Kind != "spec" || len(parsed.Added) != 1 || parsed.Added[0].ID != "ALP-004" || len(parsed.Modified) != 1 || parsed.Modified[0].ID != "ALP-001" || len(parsed.Removed) != 1 || parsed.Removed[0].ID != "ALP-002" {
			t.Fatalf("fixture semantics %+v", parsed)
		}
		doc := stateDoc(t, text)
		if len(doc.Sections) != 13 || stateBody(t, text, "Motivation") != "Retain the actual requirement." || !strings.Contains(stateBody(t, text, "Decisions"), "```md\n## Motivation\nExample heading in a fence.\n```") {
			t.Fatalf("fenced/section fixture %+v", doc)
		}
		if !reflect.DeepEqual(doc, stateDoc(t, stateRender(t, doc))) {
			t.Fatal("document helper changes public normalized values")
		}
	}
}

func Test_FLOW002NestedExplicitMetadataKeys(t *testing.T) {
	before := stateFixture(t, "evals")
	keys := []string{"depends_on", "interface", "evals", "disposable", "preserve", "append_only", "boundaries_opt_out", "hooks", "hooks.eval", "hooks.boundaries", "hooks.setup"}
	for _, key := range keys {
		for _, action := range []string{"add", "remove"} {
			t.Run(key+"_"+action, func(t *testing.T) {
				omitted := stateAlter(t, before, func(d *c.Document) {
					set := d.Metadata["set"].(map[string]any)
					if strings.HasPrefix(key, "hooks.") {
						delete(set["hooks"].(map[string]any), strings.TrimPrefix(key, "hooks."))
					} else {
						delete(set, key)
					}
				})
				explicit := stateAlter(t, omitted, func(d *c.Document) {
					set := d.Metadata["set"].(map[string]any)
					if strings.HasPrefix(key, "hooks.") {
						set["hooks"].(map[string]any)[strings.TrimPrefix(key, "hooks.")] = ""
					} else {
						switch key {
						case "hooks":
							set[key] = map[string]any{}
						case "boundaries_opt_out":
							set[key] = ""
						default:
							set[key] = []any{}
						}
					}
				})
				approved, current := omitted, explicit
				if action == "remove" {
					approved, current = current, approved
				}
				if _, err := c.ParseChange(approved, statePath, "ALP"); err != nil {
					t.Fatalf("nested approved premise %v", err)
				}
				if _, err := c.ParseChange(current, statePath, "ALP"); err != nil {
					t.Fatalf("nested current premise %v", err)
				}
				stateAssert(t, w.CheckState(stateModule(), approved, current, statePath), "P2-01", statePath, 1)
			})
		}
	}
}
func Test_FLOW002ParseableProtectedRemovals(t *testing.T) {
	for _, phase := range []string{"evals", "audit", "impl"} {
		before := stateFixture(t, phase)
		names := []string{"Interface", "Added", "Modified", "Removed", "Removed decisions", "Rejected alternatives", "Provenance", "Unknown human section"}
		if phase == "audit" {
			names = append(names, "Decisions", "Notes")
		} else {
			names = append(names, "Audit")
		}
		for _, name := range names {
			t.Run(phase+"_"+name, func(t *testing.T) {
				current := stateAlter(t, before, func(d *c.Document) { stateRemove(d, name) })
				if _, err := c.ParseChange(current, statePath, "ALP"); err != nil {
					t.Fatalf("parseable protected removal premise %v", err)
				}
				stateAssert(t, w.CheckState(stateModule(), before, current, statePath), "P2-01", statePath, 1)
			})
		}
	}
}

func Test_FLOW002CurrentDuplicateHeadingsCannotHideEdits(t *testing.T) {
	before := stateFixture(t, "evals")
	for _, name := range []string{"Decisions", "Unknown human section", "Provenance"} {
		t.Run(name, func(t *testing.T) {
			current := before + "\n## " + name + "\n\nWorker appended a second occurrence after protected content.\n"
			if _, err := c.ParseChange(current, statePath, "ALP"); err == nil {
				t.Fatal("public grammar must reject duplicate-section fixture")
			}
			stateAssert(t, w.CheckState(stateModule(), before, current, statePath), "P2-01", statePath, 1)
		})
	}
}

func stateCanonicalDirectory(t *testing.T, path string) string {
	t.Helper()
	canonical, err := filepath.EvalSymlinks(path)
	if err != nil {
		t.Fatal(err)
	}
	return canonical
}
func Test_FLOW003CanonicalDirectoryFixtureControl(t *testing.T) {
	root := stateCanonicalDirectory(t, t.TempDir())
	previous, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(root); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(previous)
	actual, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if actual != root {
		t.Fatalf("temporary root must be canonical before preservation checks: %s != %s", root, actual)
	}
}
