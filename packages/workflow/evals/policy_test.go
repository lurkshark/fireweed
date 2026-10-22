package evals

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	k "github.com/lurkshark/fireweed/packages/checks"
	c "github.com/lurkshark/fireweed/packages/contracts"
	g "github.com/lurkshark/fireweed/packages/gitstore"
	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

const policyChangePath = "changes/0024-current.md"

func policyFixture(phase string) (p.Snapshot, c.Module, c.Change) {
	m := c.Module{Name: "alpha", IDPrefix: "ALP", Dir: "units/a", Zones: c.Zones{
		Evals: []string{"tests/**", "overlap/eval/**"}, Interface: []string{"contract.go", "overlap/api/**"},
		Disposable: []string{"impl/**", "keep/**"}, Preserve: []string{"keep/**"}, AppendOnly: []string{"migrations/**", "overlap/**"}},
		DependsOn: []string{"beta"}, Hooks: c.Hooks{Eval: "hook-trap", Boundaries: "hook-trap", Setup: "hook-trap"}, Specs: []c.Spec{{ID: "ALP-001", Body: "A MUST retain approved ownership."}},
		Document: c.Document{Metadata: map[string]any{"nested": map[string]any{"list": []any{"target"}}}, Sections: []c.Section{{Name: "Specs", Body: "Approved content."}}}}
	stale := m
	stale.Zones = c.Zones{Evals: []string{"impl/**"}, Disposable: []string{"tests/**"}, AppendOnly: []string{"old-append/**"}}
	foreign := c.Module{Name: "beta", IDPrefix: "BET", Dir: "units/b", Zones: c.Zones{AppendOnly: []string{"migrations/**"}}, Document: c.Document{Metadata: map[string]any{"nested": []any{"foreign"}}}}
	s := p.Snapshot{Root: "/not-a-repository", BaseRef: "never-read", Config: c.Config{Protected: []string{".github/**", "units/a/protected/**", "shared/protected/**"}, WorkerWritable: []string{"shared/**", "go.sum"}, Agent: c.Agent{Command: []string{"agent-trap"}}, Hooks: c.Hooks{Eval: "hook-trap", Boundaries: "hook-trap"}}, Modules: []c.Module{foreign, stale}, Changes: []c.Change{{Module: "beta", Path: "changes/0025-other.md", Set: map[string]any{"preserve": []any{"keep/**"}}}}}
	ch := c.Change{Module: m.Name, Run: "approved-run", Phase: phase, Path: policyChangePath, Slug: "current", Set: map[string]any{"interface": []any{"contract.go"}}, Bases: map[string]string{"ALP-001": "hash"}, Document: c.Document{Metadata: map[string]any{"nested": []any{"approved"}}, Sections: []c.Section{{Name: "Interface", Body: "Change the retained contract."}, {Name: "Motivation", Body: "Preserve approved intent."}}}}
	return s, m, ch
}
func policyJSON(t *testing.T, v any) string {
	t.Helper()
	b, e := json.Marshal(v)
	if e != nil {
		t.Fatal(e)
	}
	return string(b)
}
func policyLess(a, b k.Diagnostic) bool {
	if a.Path != b.Path {
		return a.Path < b.Path
	}
	if a.Code != b.Code {
		return a.Code < b.Code
	}
	if a.Severity != b.Severity {
		return a.Severity < b.Severity
	}
	return a.Message < b.Message
}
func policyCorrective(message string) bool {
	lower := strings.ToLower(message)
	for _, action := range []string{"restore ", "use ", "remove ", "keep ", "provide ", "supply ", "edit only ", "revert ", "resolve ", "replace ", "check ", "retry ", "resume ", "fix "} {
		if strings.Contains(lower, action) {
			return true
		}
	}
	return false
}
func policyAssert(t *testing.T, got []k.Diagnostic, expected map[string][]string) {
	t.Helper()
	if got == nil {
		t.Fatal("nil diagnostics")
	}
	if !sort.SliceIsSorted(got, func(i, j int) bool { return policyLess(got[i], got[j]) }) {
		t.Fatalf("unsorted diagnostics: %+v", got)
	}
	actual := map[string][]string{}
	seen := map[k.Diagnostic]bool{}
	for _, d := range got {
		if d.Severity != "error" || strings.TrimSpace(d.Message) == "" || !policyCorrective(d.Message) {
			t.Fatalf("missing corrective diagnostic: %+v", d)
		}
		if seen[d] {
			t.Fatalf("duplicate diagnostic: %+v", d)
		}
		seen[d] = true
		actual[d.Path] = append(actual[d.Path], d.Code)
	}
	if expected == nil {
		expected = map[string][]string{}
	}
	for _, codes := range expected {
		sort.Strings(codes)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("diagnostics %+v; want codes by path %+v", got, expected)
	}
}
func policyWant(path string, codes ...string) map[string][]string {
	if len(codes) == 0 {
		return nil
	}
	return map[string][]string{path: codes}
}
func Test_FLOW004OwnershipAndProtectedPriority(t *testing.T) {
	denied := []string{"fireweed.yaml", "SYSTEM.md", ".fireweed/hooks.go", "AGENTS.md", "nested/AGENTS.md", "nested/CLAUDE.md", ".github/workflow.yml", "units/a/protected/code.go", "shared/protected/file", "units/a/SPEC.md", "units/a/nested/SPEC.md", "SPEC.md", "outside/SPEC.md", "units/b/code.go", "units/a-other/code.go", "human.txt", "changes/0025-other.md", "changes/archive/0024-current.md", "changes/0024-current.md.extra"}
	for _, phase := range []string{"evals", "audit", "impl"} {
		for _, path := range denied {
			t.Run(phase+"_"+path, func(t *testing.T) {
				s, m, ch := policyFixture(phase)
				policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{{Status: "M", Path: path}}), policyWant(path, "P2-01"))
			})
		}
	}
	for _, path := range []string{"fireweed.yaml", "units/b/code.go", "changes/0025-other.md", "outside/SPEC.md", "units/a/protected/code.go"} {
		t.Run("broad_worker_writable_"+path, func(t *testing.T) {
			s, m, ch := policyFixture("impl")
			s.Config.WorkerWritable = []string{"**"}
			policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{{Status: "M", Path: path}}), policyWant(path, "P2-01"))
		})
	}
	for _, phase := range []string{"evals", "audit", "impl"} {
		t.Run("exact_current_"+phase, func(t *testing.T) {
			s, m, ch := policyFixture(phase)
			policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{{Status: "M", Path: ch.Path}}), nil)
		})
	}
	for _, status := range []string{"A", "D", "T", "R", "C"} {
		t.Run("current_not_ordinary_M_"+status, func(t *testing.T) {
			s, m, ch := policyFixture("impl")
			e := g.Edit{Status: status, Path: ch.Path}
			if status == "R" || status == "C" {
				e.OldPath = "units/a/manifest.old"
			}
			policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{e}), policyWant(ch.Path, "P2-01"))
		})
	}
}
func Test_FLOW004ProjectedZonesReplaceStaleSnapshot(t *testing.T) {
	s, m, ch := policyFixture("impl")
	policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{{Status: "M", Path: "units/a/impl/existing.go"}}), nil)
	policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{{Status: "M", Path: "units/a/tests/frozen.go"}}), policyWant("units/a/tests/frozen.go", "P2-01"))
	policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{{Status: "M", Path: "units/a/old-append/ordinary.go"}}), nil)
	ch.Phase = "evals"
	policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{{Status: "M", Path: "units/a/impl/existing.go"}}), policyWant("units/a/impl/existing.go", "P1-01"))
	policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{{Status: "M", Path: "units/a/tests/frozen.go"}}), nil)
	// Projected metadata must come from the supplied module, not from this stale change's raw set.
	ch.Set = map[string]any{"evals": []any{"impl/**"}}
	policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{{Status: "M", Path: "units/a/impl/existing.go"}}), policyWant("units/a/impl/existing.go", "P1-01"))
}
func Test_FLOW005EvalFileMatrix(t *testing.T) {
	for _, row := range []struct {
		path              string
		ordinary, newFile []string
	}{
		{"units/a/tests/assert.go", nil, nil}, {"units/a/contract.go", nil, nil}, {"units/a/package.json", nil, nil}, {"shared/workspace.lock", nil, nil}, {"go.sum", nil, nil},
		{"units/a/impl/code.go", []string{"P1-01"}, nil}, {"units/a/keep/code.go", []string{"P1-01"}, []string{"P1-01"}},
		{"units/a/migrations/001.sql", []string{"H03", "P1-01"}, []string{"P1-01"}},
	} {
		for _, status := range []string{"A", "M", "D", "T", "R", "C"} {
			t.Run(row.path+"_"+status, func(t *testing.T) {
				s, m, ch := policyFixture("evals")
				e := g.Edit{Status: status, Path: row.path}
				codes := row.ordinary
				if status == "A" || status == "C" {
					codes = row.newFile
				}
				if status == "R" || status == "C" {
					e.OldPath = "units/a/manifests/old"
				}
				policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{e}), policyWant(row.path, append([]string{}, codes...)...))
			})
		}
	}
}
func Test_FLOW005InterfaceRequiresApprovedDescription(t *testing.T) {
	for _, row := range []struct {
		name, section, body string
		permit              bool
	}{
		{"absent", "", "", false}, {"empty", "Interface", "", false}, {"blank", "Interface", " \t\n ", false}, {"lowercase", "interface", "Changes.", false}, {"substring", "Interface proposal", "Changes.", false}, {"elsewhere", "Decisions", "## Interface\nFenced example.", false}, {"described", "Interface", "Describe approved edits.", true},
	} {
		for _, status := range []string{"A", "M", "D", "T", "R", "C"} {
			t.Run(row.name+"_"+status, func(t *testing.T) {
				s, m, ch := policyFixture("evals")
				ch.Document.Sections = nil
				if row.section != "" {
					ch.Document.Sections = []c.Section{{Name: row.section, Body: row.body}}
				}
				e := g.Edit{Status: status, Path: "units/a/contract.go"}
				if status == "R" || status == "C" {
					e.OldPath = "units/a/package.old"
				}
				codes := []string{"P1-01"}
				if row.permit {
					codes = nil
				}
				policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{e}), policyWant(e.Path, codes...))
			})
		}
	}
}
func Test_FLOW006ImplAndAuditFileMatrix(t *testing.T) {
	for _, row := range []struct {
		path              string
		ordinary, newFile []string
	}{
		{"units/a/impl/code.go", nil, nil}, {"units/a/keep/code.go", nil, nil}, {"units/a/package.json", nil, nil}, {"shared/workspace.lock", nil, nil}, {"go.sum", nil, nil},
		{"units/a/tests/assert.go", []string{"P2-01"}, []string{"P2-01"}}, {"units/a/contract.go", []string{"P2-01"}, []string{"P2-01"}},
		{"units/a/migrations/001.sql", []string{"H03"}, nil},
	} {
		for _, status := range []string{"A", "M", "D", "T", "R", "C"} {
			for _, phase := range []string{"impl", "audit"} {
				t.Run(phase+"_"+row.path+"_"+status, func(t *testing.T) {
					s, m, ch := policyFixture(phase)
					e := g.Edit{Status: status, Path: row.path}
					codes := append([]string{}, row.ordinary...)
					if status == "A" || status == "C" {
						codes = append([]string{}, row.newFile...)
					}
					if status == "R" || status == "C" {
						e.OldPath = "units/a/package.old"
					}
					expected := policyWant(row.path, codes...)
					if phase == "audit" {
						codes = []string{"P2-01"}
						if row.path == "units/a/migrations/001.sql" && status != "A" && status != "C" {
							codes = append(codes, "H03")
						}
						expected = policyWant(row.path, codes...)
						if status == "R" {
							expected[e.OldPath] = []string{"P2-01"}
						}
					}
					policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{e}), expected)
				})
			}
		}
	}
}
func Test_FLOW005_FLOW006RenameSourcesAndDestinations(t *testing.T) {
	rows := []struct {
		name, phase, from, to string
		want                  map[string][]string
	}{
		{"frozen_eval_source", "impl", "units/a/tests/frozen.go", "units/a/impl/new.go", policyWant("units/a/tests/frozen.go", "P2-01")},
		{"frozen_interface_source", "impl", "units/a/contract.go", "shared/new", policyWant("units/a/contract.go", "P2-01")},
		{"frozen_eval_destination", "impl", "units/a/impl/existing.go", "units/a/tests/new.go", policyWant("units/a/tests/new.go", "P2-01")},
		{"protected_source", "impl", ".fireweed/hook", "units/a/impl/new.go", policyWant(".fireweed/hook", "P2-01")},
		{"foreign_source", "impl", "units/b/code.go", "units/a/impl/new.go", policyWant("units/b/code.go", "P2-01")},
		{"preserved_source", "evals", "units/a/keep/code.go", "units/a/tests/new.go", policyWant("units/a/keep/code.go", "P1-01")},
		{"disposable_source", "evals", "units/a/impl/code.go", "units/a/tests/new.go", policyWant("units/a/impl/code.go", "P1-01")},
		{"disposable_destination", "evals", "units/a/package.json", "units/a/impl/new.go", policyWant("units/a/impl/new.go", "P1-01")},
		{"current_source", "impl", policyChangePath, "units/a/impl/new.go", policyWant(policyChangePath, "P2-01")},
		{"other_change_destination", "impl", "units/a/impl/code.go", "changes/0025-other.md", policyWant("changes/0025-other.md", "P2-01")},
		{"two_distinct_rejections", "impl", "units/b/code.go", ".fireweed/hook", map[string][]string{"units/b/code.go": {"P2-01"}, ".fireweed/hook": {"P2-01"}}},
		{"allowed_eval_rename", "evals", "units/a/tests/old.go", "units/a/tests/new.go", nil},
		{"interface_source_without_description", "evals", "units/a/contract.go", "units/a/tests/new.go", policyWant("units/a/contract.go", "P1-01")},
		{"allowed_impl_rename", "impl", "units/a/impl/old.go", "units/a/keep/new.go", nil},
	}
	for _, row := range rows {
		t.Run(row.name, func(t *testing.T) {
			s, m, ch := policyFixture(row.phase)
			if row.name == "interface_source_without_description" {
				ch.Document.Sections = nil
			}
			policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{{Status: "R", OldPath: row.from, Path: row.to}}), row.want)
		})
	}
	for _, phase := range []string{"evals", "impl"} {
		for _, source := range []string{".fireweed/hook", "units/b/code.go", policyChangePath, "units/a/migrations/old.sql"} {
			t.Run("unchanged_copy_source_"+phase+"_"+source, func(t *testing.T) {
				s, m, ch := policyFixture(phase)
				policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{{Status: "C", OldPath: source, Path: "units/a/impl/new.go"}}), nil)
			})
		}
	}
}
func Test_FLOW006AppendOnlyGlobalAndOverlap(t *testing.T) {
	for _, phase := range []string{"evals", "audit", "impl"} {
		for _, status := range []string{"A", "M", "D", "T", "R", "C"} {
			for _, path := range []string{"units/a/overlap/eval/case.go", "units/a/overlap/api/case.go", "units/a/overlap/protected/file", "units/b/migrations/001.sql"} {
				t.Run(phase+"_"+status+"_"+path, func(t *testing.T) {
					s, m, ch := policyFixture(phase)
					s.Config.Protected = append(s.Config.Protected, "units/a/overlap/protected/**")
					e := g.Edit{Status: status, Path: path}
					if status == "R" || status == "C" {
						e.OldPath = "units/a/package.old"
					}
					codes := []string{}
					if status != "A" && status != "C" {
						codes = append(codes, "H03")
					}
					if phase == "evals" && (status == "A" || status == "C") {
						codes = append(codes, "P1-01")
					}
					switch {
					case strings.HasPrefix(path, "units/b/") || strings.Contains(path, "/protected/") || phase != "evals":
						codes = append(codes, "P2-01")
					}
					expected := policyWant(path, codes...)
					if phase == "audit" && status == "R" {
						expected[e.OldPath] = []string{"P2-01"}
					}
					policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{e}), expected)
				})
			}
		}
	}
	for _, phase := range []string{"evals", "audit", "impl"} {
		t.Run("both_rename_append_endpoints_"+phase, func(t *testing.T) {
			s, m, ch := policyFixture(phase)
			codes := []string{"H03"}
			if phase == "evals" {
				codes = append(codes, "P1-01")
			}
			if phase == "audit" {
				codes = append(codes, "P2-01")
			}
			policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{{Status: "R", OldPath: "units/a/migrations/old.sql", Path: "units/a/migrations/new.sql"}}), map[string][]string{"units/a/migrations/old.sql": append([]string{}, codes...), "units/a/migrations/new.sql": append([]string{}, codes...)})
		})
	}
}
func Test_FLOW007InvalidApprovalStopsComparison(t *testing.T) {
	for _, row := range []struct {
		name string
		edit func(*c.Change)
	}{
		{"no_run", func(ch *c.Change) { ch.Run = "" }}, {"blank_run", func(ch *c.Change) { ch.Run = " \t\n" }}, {"wrong_target", func(ch *c.Change) { ch.Module = "beta" }},
		{"queued", func(ch *c.Change) { ch.Phase = "" }}, {"applied", func(ch *c.Change) { ch.Phase = "applied" }}, {"unknown", func(ch *c.Change) { ch.Phase = "EVALS" }},
	} {
		t.Run(row.name, func(t *testing.T) {
			s, m, ch := policyFixture("evals")
			row.edit(&ch)
			policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{{Status: "M", Path: "fireweed.yaml"}}), policyWant(ch.Path, "P02"))
			malformed := []g.Edit{{Status: "M", Path: "../escape"}, {Status: "R100", Path: "units/a/new", OldPath: "/absolute"}, {Status: "C", Path: "units/a/new"}}
			policyAssert(t, w.CheckEdits(s, m, ch, malformed), policyWant(ch.Path, "P02"))
		})
	}
}
func Test_FLOW007MalformedEdits(t *testing.T) {
	for _, path := range []string{"", ".", "..", "/absolute", "./units/a/file", "units//a/file", "units/../file", "units/./file", "units/a/", "units\\a\\file", "units/a/nu\x00ll"} {
		for _, status := range []string{"M", "R", "C"} {
			t.Run(fmt.Sprintf("%q_%s", path, status), func(t *testing.T) {
				s, m, ch := policyFixture("impl")
				e := g.Edit{Status: status, Path: path}
				if status == "R" || status == "C" {
					e.OldPath = "units/a/from"
				}
				policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{e}), policyWant(path, "P2-01"))
			})
		}
	}
	for _, e := range []g.Edit{{Status: "", Path: "units/a/file"}, {Status: "U", Path: "units/a/file"}, {Status: "R100", Path: "units/a/file", OldPath: "units/a/from"}, {Status: "R", Path: "units/a/file"}, {Status: "C", Path: "units/a/file"}, {Status: "R", Path: "units/a/file", OldPath: "units/a/file"}, {Status: "C", Path: "units/a/file", OldPath: "units/a/file"}, {Status: "M", Path: "units/a/file", OldPath: "units/a/from"}, {Status: "A", Path: "units/a/file", OldPath: "units/a/from"}, {Status: "D", Path: "units/a/file", OldPath: "units/a/from"}, {Status: "T", Path: "units/a/file", OldPath: "units/a/from"}} {
		t.Run(fmt.Sprintf("shape_%+v", e), func(t *testing.T) {
			s, m, ch := policyFixture("impl")
			policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{e}), policyWant(e.Path, "P2-01"))
		})
	}
	for _, path := range []string{"../source", "/source", "units//source", "units/./source", "units\\source", "source\x00bad"} {
		for _, status := range []string{"R", "C"} {
			t.Run("bad_source_"+status+"_"+fmt.Sprintf("%q", path), func(t *testing.T) {
				s, m, ch := policyFixture("impl")
				policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{{Status: status, Path: "units/a/new.go", OldPath: path}}), policyWant(path, "P2-01"))
			})
		}
	}
}
func Test_FLOW007DiagnosticAndInputOwnership(t *testing.T) {
	s, m, ch := policyFixture("evals")
	edits := []g.Edit{{Status: "R", OldPath: "units/b/migrations/old.sql", Path: "units/a/migrations/new.sql"}, {Status: "M", Path: "fireweed.yaml"}, {Status: "M", Path: "fireweed.yaml"}}
	want := map[string][]string{"units/b/migrations/old.sql": {"H03", "P2-01"}, "units/a/migrations/new.sql": {"H03", "P1-01"}, "fireweed.yaml": {"P2-01"}}
	before := policyJSON(t, []any{s, m, ch, edits})
	first := w.CheckEdits(s, m, ch, edits)
	policyAssert(t, first, want)
	second := w.CheckEdits(s, m, ch, edits)
	policyAssert(t, second, want)
	if !reflect.DeepEqual(first, second) {
		t.Fatal("nondeterministic results")
	}
	saved := policyJSON(t, second)
	first[0].Message = "mutated"
	first = append(first, k.Diagnostic{Code: "new"})
	if before != policyJSON(t, []any{s, m, ch, edits}) || saved != policyJSON(t, second) {
		t.Fatal("immediate result/input aliases")
	}
	// Input ordering, including duplicates, does not control diagnostic ordering.
	reversed := append([]g.Edit{}, edits...)
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	if !reflect.DeepEqual(second, w.CheckEdits(s, m, ch, reversed)) {
		t.Fatal("input order changed diagnostics")
	}
	for _, call := range []struct {
		name  string
		edits []g.Edit
	}{{"empty", nil}, {"allowed", []g.Edit{{Status: "A", Path: "units/a/impl/stub.go"}}}} {
		t.Run(call.name, func(t *testing.T) {
			a, b := w.CheckEdits(s, m, ch, call.edits), w.CheckEdits(s, m, ch, call.edits)
			policyAssert(t, a, nil)
			policyAssert(t, b, nil)
			a = append(a, k.Diagnostic{Message: "a"})
			b = append(b, k.Diagnostic{Message: "b"})
			if a[0].Message != "a" || b[0].Message != "b" {
				t.Fatal("empty writable results alias")
			}
		})
	}
	ch.Run = ""
	savedInputs := policyJSON(t, []any{s, m, ch, edits})
	a, b := w.CheckEdits(s, m, ch, edits), w.CheckEdits(s, m, ch, edits)
	policyAssert(t, a, policyWant(ch.Path, "P02"))
	policyAssert(t, b, policyWant(ch.Path, "P02"))
	a[0].Code = "mutated"
	if b[0].Code != "P02" || savedInputs != policyJSON(t, []any{s, m, ch, edits}) {
		t.Fatal("invalid input/result aliases")
	}
}
func Test_FLOW007OutsideRepositoryAndCommandSafety(t *testing.T) {
	root := stateCanonicalDirectory(t, t.TempDir())
	traps := t.TempDir()
	marker := filepath.Join(traps, "called")
	script := "#!/bin/sh\nprintf called >> '" + strings.ReplaceAll(marker, "'", `'"'"'`) + "'\nexit 97\n"
	for _, name := range []string{"git", "hook-trap", "agent-trap", "codex", "claude", "fireweed", "fireweed-bootstrap", "curl", "wget", "ssh", "nc"} {
		if err := os.WriteFile(filepath.Join(traps, name), []byte(script), 0755); err != nil {
			t.Fatal(err)
		}
	}
	t.Setenv("PATH", traps+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("FIREWEED_BIN", filepath.Join(traps, "fireweed-bootstrap"))
	if err := exec.Command("git", "--version").Run(); err == nil {
		t.Fatal("command trap control")
	}
	if _, err := os.Stat(marker); err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(marker); err != nil {
		t.Fatal(err)
	}
	for path, text := range map[string]string{"user": "retain", "ignored": "ignored bytes", ".git/refs/heads/retained": "fake ref"} {
		full := filepath.Join(root, path)
		if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(text), 0644); err != nil {
			t.Fatal(err)
		}
	}
	previous, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err = os.Chdir(root); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(previous)
	env := os.Environ()
	files := stateFiles(t, root)
	for _, mode := range []string{"allowed", "denied", "invalid"} {
		t.Run(mode, func(t *testing.T) {
			s, m, ch := policyFixture("impl")
			s.Root = root
			edits := []g.Edit{{Status: "M", Path: "units/a/impl/code.go"}}
			expected := map[string][]string(nil)
			if mode == "denied" {
				edits[0].Path = "fireweed.yaml"
				expected = policyWant("fireweed.yaml", "P2-01")
			}
			if mode == "invalid" {
				ch.Phase = "applied"
				expected = policyWant(ch.Path, "P02")
			}
			inputs := policyJSON(t, []any{s, m, ch, edits})
			policyAssert(t, w.CheckEdits(s, m, ch, edits), expected)
			cwd, err := os.Getwd()
			if err != nil || cwd != root || !reflect.DeepEqual(env, os.Environ()) || !reflect.DeepEqual(files, stateFiles(t, root)) || inputs != policyJSON(t, []any{s, m, ch, edits}) {
				t.Fatal("mutated files/refs/cwd/environment/inputs")
			}
			if _, err := os.Stat(marker); !os.IsNotExist(err) {
				t.Fatal("executed a trapped command")
			}
		})
	}
}

func Test_FLOW004_FLOW005_FLOW006WorkerWritableCannotWidenTarget(t *testing.T) {
	for _, row := range []struct {
		phase, path string
		codes       []string
	}{
		{"impl", "units/a/tests/frozen.go", []string{"P2-01"}},
		{"impl", "units/a/contract.go", []string{"P2-01"}},
		{"evals", "units/a/impl/existing.go", []string{"P1-01"}},
		{"evals", "units/a/keep/existing.go", []string{"P1-01"}},
		{"evals", "units/a/migrations/existing.sql", []string{"H03", "P1-01"}},
		{"audit", "units/a/package.json", []string{"P2-01"}},
	} {
		t.Run(row.phase+"_"+row.path, func(t *testing.T) {
			s, m, ch := policyFixture(row.phase)
			s.Config.WorkerWritable = []string{"**"}
			policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{{Status: "M", Path: row.path}}), policyWant(row.path, row.codes...))
		})
	}
	s, m, ch := policyFixture("evals")
	s.Config.WorkerWritable = []string{"units/a/**"}
	ch.Document.Sections = nil
	policyAssert(t, w.CheckEdits(s, m, ch, []g.Edit{{Status: "M", Path: "units/a/contract.go"}}), policyWant("units/a/contract.go", "P1-01"))
}
func Test_FLOW007CorrectiveMessageFixtureControl(t *testing.T) {
	for _, text := range []string{"", "Invalid edit.", "P02", "An error occurred."} {
		if policyCorrective(text) {
			t.Fatalf("non-guidance accepted: %q", text)
		}
	}
	for _, text := range []string{"Restore the approved claim before continuing.", "Remove the forbidden edit.", "Use a canonical path.", "Keep worker edits in the permitted zones."} {
		if !policyCorrective(text) {
			t.Fatalf("corrective action rejected: %q", text)
		}
	}
}
