package evals_test

import (
	"context"
	"encoding/json"
	"fmt"
	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"
)

func git(t *testing.T, root string, args ...string) string {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_CONFIG_GLOBAL=/dev/null", "GIT_AUTHOR_NAME=Fixture", "GIT_AUTHOR_EMAIL=fixture@example.invalid", "GIT_COMMITTER_NAME=Fixture", "GIT_COMMITTER_EMAIL=fixture@example.invalid")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %s %v", args, out, err)
	}
	return string(out)
}
func put(t *testing.T, root, path, text string) {
	t.Helper()
	full := filepath.Join(root, filepath.FromSlash(path))
	if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(full, []byte(text), 0644); err != nil {
		t.Fatal(err)
	}
}

const config = "fireweed: '>=0.4, <0.5'\nversion: 1\nbase: main\nmodules: [units/*]\nmodule_defaults:\n  interface: [api.go]\n  evals: [evals/**]\n  disposable: [internal/**]\nhooks:\n  eval: touch MUST-NOT-RUN\n  boundaries: touch MUST-NOT-RUN\n"

func fixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	git(t, root, "init", "-b", "main")
	put(t, root, "fireweed.yaml", config)
	return root
}
func module(name, prefix, extra string) string {
	return fmt.Sprintf("---\nname: %s\nid_prefix: %s\n%s---\n\n# Human preamble\n\n## Specs\n\n### %s-001\nThe module MUST retain this behavior.\n\n## Decisions\n\nKeep this reason.\n", name, prefix, extra, prefix)
}
func change(target, extra string) string {
	return "---\nmodule: " + target + "\n" + extra + "---\n\n## Motivation\n\nA fixture needs this behavior.\n\n## Added\n\n### new:fresh\nThe fixture MUST work.\n\n## Modified\n\n## Removed\n\n## Provenance\n\n- authored fixture\n\n## Audit\n\nNo findings.\n"
}
func requireDiag(t *testing.T, ds []p.Diagnostic, code, severity, path string) {
	t.Helper()
	for _, d := range ds {
		if d.Code == code && d.Severity == severity && d.Path == path && strings.TrimSpace(d.Message) != "" {
			return
		}
	}
	t.Fatalf("missing %s %s %s: %+v", code, severity, path, ds)
}
func validModule(name, prefix, dir string) c.Module {
	return c.Module{Name: name, IDPrefix: prefix, Dir: dir, Hooks: c.Hooks{Eval: "opaque eval", Boundaries: "opaque boundaries"}}
}
func requireClean(t *testing.T, s p.Snapshot) {
	t.Helper()
	ds := p.Check(s)
	if ds == nil || len(ds) != 0 {
		t.Fatalf("expected nonnil empty diagnostics: %+v", ds)
	}
}
func Test_PRJ001Discovery(t *testing.T) {
	root := fixture(t)
	put(t, root, "units/z/SPEC.md", module("zeta", "ZED", "depends_on: [alpha]\n"))
	put(t, root, "units/a/SPEC.md", module("alpha", "ALP", "evals: []\ndisposable: []\nhooks:\n  setup: custom setup\n"))
	put(t, root, "units/a/nested/.keep", "")
	git(t, root, "add", ".")
	git(t, root, "commit", "-m", "fixture")
	before := git(t, root, "status", "--porcelain=v1", "--untracked-files=all") + git(t, root, "show-ref")
	s, err := p.Load(filepath.Join(root, "units/a/nested"))
	if err != nil {
		t.Fatal(err)
	}
	canonical, _ := filepath.EvalSymlinks(root)
	if s.Root != canonical || s.BaseRef != "refs/heads/main" || s.Config.Version != 1 || len(s.Modules) != 2 {
		t.Fatalf("snapshot %+v", s)
	}
	a, z := s.Modules[0], s.Modules[1]
	if a.Name != "alpha" || z.Name != "zeta" || a.Dir != "units/a" || z.IDPrefix != "ZED" || !reflect.DeepEqual(z.DependsOn, []string{"alpha"}) || !reflect.DeepEqual(z.Interface, []string{"api.go"}) || !reflect.DeepEqual(z.Disposable, []string{"internal/**"}) || len(a.Evals) != 0 || a.Evals == nil || len(a.Disposable) != 0 || a.Disposable == nil {
		t.Fatalf("resolved modules %+v", s.Modules)
	}
	if a.Hooks.Eval != "touch MUST-NOT-RUN" || a.Hooks.Boundaries != "touch MUST-NOT-RUN" || a.Hooks.Setup != "custom setup" || len(z.Specs) != 1 || z.Specs[0].ID != "ZED-001" || !strings.Contains(z.Decisions, "Keep this reason.") || !strings.Contains(z.Document.Preamble, "Human preamble") || z.Document.Metadata["name"] != "zeta" {
		t.Fatalf("lost module data %+v", s.Modules)
	}
	requireClean(t, s)
	after := git(t, root, "status", "--porcelain=v1", "--untracked-files=all") + git(t, root, "show-ref")
	if before != after {
		t.Fatalf("discovery mutated repository\n%s\n%s", before, after)
	}
	if _, err := os.Stat(filepath.Join(root, "MUST-NOT-RUN")); !os.IsNotExist(err) {
		t.Fatal("hook executed")
	}
}
func Test_PRJ001ConfigurationFailures(t *testing.T) {
	for _, body := range []string{"", "version: [1]\n", "fireweed: '>=99'\n", "version: 99\n", "modules: ['../escape']\n", "module_defaults: {interface: ['[']}\n", "protected: ['/absolute']\n", "worker_writable: ['bad/**suffix']\n"} {
		root := fixture(t)
		put(t, root, "fireweed.yaml", body)
		if _, err := p.Load(root); err == nil || !strings.Contains(err.Error(), "C01") {
			t.Fatalf("config %q: %v", body, err)
		}
	}
	root := fixture(t)
	if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
		t.Fatal(err)
	}
	if _, err := p.Load(root); err == nil || !strings.Contains(err.Error(), "C01") {
		t.Fatalf("missing config: %v", err)
	}
	if _, err := p.Load(t.TempDir()); err == nil {
		t.Fatal("accepted nonrepository")
	}
}
func Test_PRJ002DocumentsAndOrdering(t *testing.T) {
	root := fixture(t)
	put(t, root, "units/a/SPEC.md", module("alpha", "ALP", ""))
	put(t, root, "outside/SPEC.md", module("outside", "OUT", ""))
	put(t, root, "units/meta/SPEC.md", module("invalid name", "BAD", ""))
	put(t, root, "units/grammar/SPEC.md", strings.Replace(module("grammar", "GRA", ""), "MUST", "might", 1))
	put(t, root, ".fireweed/units/ignored/SPEC.md", "broken")
	put(t, root, ".git/ignored/SPEC.md", "broken")
	for _, name := range []string{"0010-last", "0002-zebra", "0002-alpha"} {
		put(t, root, "changes/"+name+".md", change("alpha", ""))
	}
	put(t, root, "changes/archive/0009-old.md", change("alpha", "phase: applied\nrun: fixture-run\n"))
	put(t, root, "changes/archive/0001-first.md", change("alpha", "phase: applied\n"))
	put(t, root, "changes/0004-bad.md", "invalid")
	put(t, root, "changes/0005-unknown.md", change("missing", ""))
	put(t, root, "changes/archive/0007-zebra.md", change("alpha", ""))
	put(t, root, "changes/nested/0001-hidden.md", "bad")
	put(t, root, "changes/README.txt", "bad")
	s, err := p.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(s.Modules) != 1 || s.Modules[0].Name != "alpha" {
		t.Fatalf("modules %+v", s.Modules)
	}
	for _, v := range []struct{ code, severity, path string }{{"C02", "warning", "outside/SPEC.md"}, {"C03", "error", "units/meta/SPEC.md"}, {"C04", "error", "units/grammar/SPEC.md"}, {"C05", "error", "changes/0004-bad.md"}, {"C05", "error", "changes/0005-unknown.md"}, {"C05", "error", "changes/archive/0007-zebra.md"}} {
		requireDiag(t, s.Diagnostics, v.code, v.severity, v.path)
	}
	if len(s.Diagnostics) != 6 {
		t.Fatalf("unexpected discovery diagnostics %+v", s.Diagnostics)
	}
	names := func(cs []c.Change) []string {
		out := []string{}
		for _, v := range cs {
			out = append(out, v.Slug)
		}
		return out
	}
	if !reflect.DeepEqual(names(s.Changes), []string{"alpha", "zebra", "last"}) || !reflect.DeepEqual(names(s.Archive), []string{"first", "zebra", "old"}) {
		t.Fatalf("queue/archive %+v %+v", names(s.Changes), names(s.Archive))
	}
	old := s.Archive[2]
	if old.Run != "fixture-run" || old.Phase != "applied" || old.Module != "alpha" || old.Added[0].ID != "new:fresh" || old.Path != "changes/archive/0009-old.md" {
		t.Fatalf("lost change %+v", old)
	}
	raw, _ := c.RenderMarkdown(old.Document)
	for _, text := range []string{"A fixture needs this behavior.", "authored fixture", "No findings."} {
		if !strings.Contains(raw, text) {
			t.Fatalf("lost %q: %s", text, raw)
		}
	}
}
func Test_PRJ001_PRJ002Symlinks(t *testing.T) {
	root := fixture(t)
	outside := t.TempDir()
	put(t, outside, "SPEC.md", module("secret", "SEC", ""))
	if err := os.Symlink(outside, filepath.Join(root, "units")); err != nil {
		t.Fatal(err)
	}
	s, err := p.Load(root)
	if err != nil || len(s.Modules) != 0 || len(s.Diagnostics) != 0 {
		t.Fatalf("followed directory symlink: %+v %v", s, err)
	}
	if err := os.Remove(filepath.Join(root, "units")); err != nil {
		t.Fatal(err)
	}
	put(t, root, "units/escape/.keep", "")
	if err := os.Symlink(filepath.Join(outside, "SPEC.md"), filepath.Join(root, "units/escape/SPEC.md")); err != nil {
		t.Fatal(err)
	}
	put(t, root, "changes/.keep", "")
	if err := os.Symlink(filepath.Join(outside, "SPEC.md"), filepath.Join(root, "changes/0001-escape.md")); err != nil {
		t.Fatal(err)
	}
	s, err = p.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	requireDiag(t, s.Diagnostics, "C03", "error", "units/escape/SPEC.md")
	requireDiag(t, s.Diagnostics, "C05", "error", "changes/0001-escape.md")
	if len(s.Modules) != 0 || len(s.Changes) != 0 {
		t.Fatalf("read escape %+v", s)
	}
}
func Test_PRJ003Graph(t *testing.T) {
	a := validModule("a", "AAA", "units/a")
	b := validModule("b", "BBB", "units/b")
	d := validModule("d", "DDD", "units/d")
	a.DependsOn = []string{"d"}
	b.DependsOn = []string{"d"}
	requireClean(t, p.Snapshot{Modules: []c.Module{b, d, a}})
	cases := []struct {
		name, code string
		mods       []c.Module
	}{{"name", "C03", []c.Module{a, validModule("a", "ZZZ", "units/z")}}, {"prefix", "C03", []c.Module{b, validModule("z", "BBB", "units/z")}}, {"nested", "C06", []c.Module{b, validModule("z", "ZZZ", "units/b/future")}}, {"same-dir", "C06", []c.Module{b, validModule("z", "ZZZ", "units/b")}}, {"root", "C06", []c.Module{b, validModule("z", "ZZZ", ".")}}}
	for _, kind := range []string{"missing", "self", "cycle"} {
		x, y := b, d
		if kind == "missing" {
			x.DependsOn = []string{"gone"}
		}
		if kind == "self" {
			x.DependsOn = []string{"b"}
		}
		if kind == "cycle" {
			x.DependsOn = []string{"d"}
			y.DependsOn = []string{"b"}
		}
		cases = append(cases, struct {
			name, code string
			mods       []c.Module
		}{kind, "C06", []c.Module{x, y}})
	}
	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			ds := p.Check(p.Snapshot{Modules: v.mods})
			found := false
			for _, d := range ds {
				if d.Code == v.code && d.Severity == "error" && strings.HasSuffix(d.Path, "SPEC.md") && d.Message != "" {
					found = true
				}
			}
			if !found {
				t.Fatalf("missing %s %+v", v.code, ds)
			}
		})
	}
}
func Test_PRJ004ZoneValidation(t *testing.T) {
	for _, zone := range []string{"interface", "evals", "disposable", "preserve", "append_only"} {
		m := validModule("a", "AAA", "future/a")
		bad := []string{"../bad"}
		switch zone {
		case "interface":
			m.Interface = bad
		case "evals":
			m.Evals = bad
		case "disposable":
			m.Disposable = bad
		case "preserve":
			m.Preserve = bad
		case "append_only":
			m.AppendOnly = bad
		}
		requireDiag(t, p.Check(p.Snapshot{Modules: []c.Module{m}}), "C03", "error", "future/a/SPEC.md")
	}
	m := validModule("a", "AAA", "future/a")
	m.Evals = []string{"src/**/check?.go"}
	m.Disposable = []string{"src/new/check[0-9].go"}
	requireDiag(t, p.Check(p.Snapshot{Modules: []c.Module{m}}), "C06", "error", "future/a/SPEC.md")
	m.Evals = []string{"evals/**"}
	m.Disposable = []string{"internal/**"}
	m.Interface = []string{"internal/api.go"}
	m.Preserve = []string{"internal/**"}
	m.AppendOnly = []string{"internal/log/*"}
	requireClean(t, p.Snapshot{Modules: []c.Module{m}})
}
func Test_PRJ005Hooks(t *testing.T) {
	for _, v := range []struct {
		eval, bound, reason string
		disposable          bool
		code, severity      string
	}{{" ", "ok", "", true, "C07", "error"}, {"ok", " ", "", true, "C07", "error"}, {"ok", "", "   ", true, "C07", "error"}, {"ok", "", "legacy reason", true, "D04", "warning"}, {"ok", "ok", "review exception", true, "D04", "warning"}, {"ok", "", "", false, "D04", "warning"}, {"", "", "reason", true, "C07", "error"}} {
		m := validModule("a", "AAA", "units/a")
		m.Hooks = c.Hooks{Eval: v.eval, Boundaries: v.bound}
		m.BoundariesOptOut = v.reason
		if v.disposable {
			m.Disposable = []string{"src/**"}
		}
		ds := p.Check(p.Snapshot{Modules: []c.Module{m}})
		requireDiag(t, ds, v.code, v.severity, "units/a/SPEC.md")
		if strings.TrimSpace(v.reason) != "" {
			requireDiag(t, ds, "D04", "warning", "units/a/SPEC.md")
			for _, d := range ds {
				if d.Code == "D04" && !strings.Contains(d.Message, v.reason) {
					t.Fatalf("reason hidden %+v", ds)
				}
			}
		}
		if v.eval == "ok" && strings.TrimSpace(v.reason) != "" {
			for _, d := range ds {
				if d.Code == "C07" {
					t.Fatalf("optout did not waive boundaries %+v", ds)
				}
			}
		}
	}
	m := validModule("a", "AAA", "units/a")
	m.Disposable = []string{"src/**"}
	requireClean(t, p.Snapshot{Modules: []c.Module{m}})
}
func Test_PRJ006ReservationsAndDeterminism(t *testing.T) {
	m := validModule("a", "AAA", "units/a")
	m.Specs = []c.Spec{{ID: "AAA-002", Body: "MUST survive."}}
	s := p.Snapshot{Modules: []c.Module{m}, Archive: []c.Change{{Path: "changes/archive/0001-old.md", Removed: []c.Spec{{ID: "AAA-002"}}}}, Diagnostics: []p.Diagnostic{{Code: "C05", Severity: "error", Message: "Fix fixture two", Path: "z.md"}, {Code: "C03", Severity: "error", Message: "Fix fixture one", Path: "a.md"}, {Code: "C03", Severity: "warning", Message: "Review fixture", Path: "a.md"}}}
	before, _ := json.Marshal(s)
	ds := p.Check(s)
	requireDiag(t, ds, "C08", "error", "units/a/SPEC.md")
	for _, v := range s.Diagnostics {
		found := false
		for _, d := range ds {
			if d == v {
				found = true
			}
		}
		if !found {
			t.Fatalf("lost load diagnostic %+v", v)
		}
	}
	if !sort.SliceIsSorted(ds, func(i, j int) bool {
		a, b := ds[i], ds[j]
		return a.Path+"\x00"+a.Code+"\x00"+a.Severity+"\x00"+a.Message < b.Path+"\x00"+b.Code+"\x00"+b.Severity+"\x00"+b.Message
	}) {
		t.Fatalf("unsorted %+v", ds)
	}
	if !reflect.DeepEqual(ds, p.Check(s)) {
		t.Fatal("unstable results")
	}
	after, _ := json.Marshal(s)
	if string(before) != string(after) {
		t.Fatal("mutated snapshot")
	}
	m.Specs[0].ID = "AAA-003"
	requireClean(t, p.Snapshot{Modules: []c.Module{m}, Archive: s.Archive})
}
