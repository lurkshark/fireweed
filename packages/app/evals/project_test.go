package evals

import (
	"encoding/json"
	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

const projectConfig = "version: 1\nbase: main\nmodules: [units/*]\nmodule_defaults:\n  interface: [api.go]\n  evals: [evals/**]\n  disposable: [internal/**]\nhooks:\n  eval: touch MUST-NOT-RUN\n  boundaries: touch MUST-NOT-RUN\n  setup: touch MUST-NOT-RUN\nagent: {command: [agent-trap]}\naudit: {command: [agent-trap]}\n"

func projectPut(t *testing.T, root, path, text string) {
	t.Helper()
	full := filepath.Join(root, filepath.FromSlash(path))
	if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(full, []byte(text), 0644); err != nil {
		t.Fatal(err)
	}
}
func projectFixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	gitFixtureCommand(t, root, "init", "--initial-branch=main")
	projectPut(t, root, "fireweed.yaml", projectConfig)
	return root
}
func projectModule(t *testing.T, root, dir, name, prefix, extra string) {
	t.Helper()
	text := "---\nname: " + name + "\nid_prefix: " + prefix + "\n" + extra + "---\n\n# Human preamble\n\n## Specs\n\n### " + prefix + "-001 [review]\nIt MUST retain this body.\n\n## Decisions\n\nRetain this reason.\n"
	cfg, err := c.ParseConfig(projectConfig, "0.4.0")
	if err != nil {
		t.Fatal(err)
	}
	if _, err = c.ParseModule(text, dir, cfg.ModuleDefaults, cfg.Hooks); err != nil {
		t.Fatalf("invalid valid-module fixture: %v", err)
	}
	projectPut(t, root, dir+"/SPEC.md", text)
}
func canonical(t *testing.T, v any) any {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	var out any
	if err = json.Unmarshal(b, &out); err != nil {
		t.Fatal(err)
	}
	return out
}
func decodeData(t *testing.T, raw json.RawMessage) any {
	t.Helper()
	var out any
	if err := json.Unmarshal(raw, &out); err != nil {
		t.Fatal(err)
	}
	return out
}

type pathDiagnostic struct {
	Code     string `json:"code"`
	Severity string `json:"severity"`
	Message  string `json:"message"`
	Path     string `json:"path"`
}

func detailedDiagnostics(t *testing.T, got result) []pathDiagnostic {
	t.Helper()
	var x struct {
		Diagnostics []pathDiagnostic `json:"diagnostics"`
	}
	if err := json.Unmarshal([]byte(got.out), &x); err != nil {
		t.Fatal(err)
	}
	if x.Diagnostics == nil {
		t.Fatal("null diagnostics")
	}
	return x.Diagnostics
}
func projectExpected(t *testing.T, root string) (p.Snapshot, []pathDiagnostic) {
	t.Helper()
	s, err := p.Load(root)
	if err != nil {
		t.Fatalf("invalid project fixture: %v", err)
	}
	ds := []pathDiagnostic{}
	for _, v := range p.Check(s) {
		ds = append(ds, pathDiagnostic{v.Code, v.Severity, v.Message, v.Path})
	}
	return s, ds
}
func Test_APP005Modules(t *testing.T) {
	root := projectFixture(t)
	projectModule(t, root, "units/first", "zeta", "ZED", "depends_on: [alpha]\n")
	projectModule(t, root, "units/z", "alpha", "ALP", "evals: []\ndisposable: []\npreserve: [data/**]\nappend_only: [events/**]\nhooks: {setup: own setup}\n")
	projectPut(t, root, "units/z/nested/.keep", "")
	s, ds := projectExpected(t, root)
	if len(ds) != 0 || len(s.Modules) != 2 {
		t.Fatalf("fixture %+v %+v", s, ds)
	}
	for _, args := range [][]string{{"modules", "--json"}, {"--json", "modules"}} {
		got := invoke(t, filepath.Join(root, "units/z/nested"), args...)
		x := successJSON(t, got, "modules")
		if !reflect.DeepEqual(decodeData(t, x.Data), canonical(t, s.Modules)) {
			t.Fatalf("module fields missing %s", x.Data)
		}
	}
	got := invoke(t, root, "modules")
	if got.code != 0 || got.err != "" {
		t.Fatalf("plain modules %+v", got)
	}
	if !reflect.DeepEqual(decodeData(t, json.RawMessage(got.out)), canonical(t, s.Modules)) {
		t.Fatal("plain module values differ")
	}
	empty := projectFixture(t)
	for _, args := range [][]string{{"modules"}, {"modules", "--json"}} {
		got := invoke(t, empty, args...)
		raw := json.RawMessage(got.out)
		if len(args) > 1 {
			raw = successJSON(t, got, "modules").Data
		}
		if got.code != 0 || string(raw) != "[]" && strings.TrimSpace(string(raw)) != "[]" {
			t.Fatalf("empty module array %+v", got)
		}
	}
}
func Test_APP006ValidationDiagnostics(t *testing.T) {
	root := projectFixture(t)
	projectModule(t, root, "units/a", "alpha", "ALP", "hooks: {eval: ''}\nboundaries_opt_out: review reason\n")
	projectModule(t, root, "units/z", "zeta", "ZED", "depends_on: [missing]\nevals: [internal/**]\n")
	projectPut(t, root, "outside/SPEC.md", "outside warning")
	projectPut(t, root, "units/bad/SPEC.md", "---\nname: bad\nid_prefix: BAD\n---\n\n## Specs\n\n### BAD-001\nNo normative content.\n")
	projectPut(t, root, "changes/archive/0001-removed.md", "---\nmodule: zeta\nphase: applied\n---\n\n## Motivation\n\nReserve an ID.\n\n## Removed\n\n### ZED-001\n")
	s, ds := projectExpected(t, root)
	for _, code := range []string{"C02", "C04", "C06", "C07", "C08", "D04"} {
		found := false
		for _, d := range ds {
			if d.Code == code {
				found = true
			}
		}
		if !found {
			t.Fatalf("fixture lacks %s: %+v", code, ds)
		}
	}
	for _, command := range []string{"modules", "check"} {
		got := invoke(t, root, command, "--json")
		x := decode(t, got, command)
		if got.code != 1 {
			t.Fatalf("invalid exit %+v", got)
		}
		if !reflect.DeepEqual(detailedDiagnostics(t, got), ds) {
			t.Fatalf("lost/path/order diagnostics %s", got.out)
		}
		want := canonical(t, s.Modules)
		if command == "check" {
			want = map[string]any{"valid": false}
		}
		if !reflect.DeepEqual(decodeData(t, x.Data), want) {
			t.Fatalf("error lost data %s", x.Data)
		}
		got = invoke(t, root, command)
		if got.code != 1 || got.out == "" {
			t.Fatalf("plain failure %+v", got)
		}
		for _, d := range ds {
			if !strings.Contains(got.err, d.Code) || !strings.Contains(got.err, d.Path) {
				t.Fatalf("plain missing diagnostic path %+v", got)
			}
		}
	}
	// Warnings alone preserve success and data.
	warning := projectFixture(t)
	projectModule(t, warning, "units/a", "alpha", "ALP", "boundaries_opt_out: visible reason\n")
	_, warningDS := projectExpected(t, warning)
	if len(warningDS) != 1 || warningDS[0].Code != "D04" {
		t.Fatalf("warning fixture %+v", warningDS)
	}
	for _, command := range []string{"check", "modules"} {
		got := invoke(t, warning, "--json", command)
		x := decode(t, got, command)
		if got.code != 0 || !reflect.DeepEqual(detailedDiagnostics(t, got), warningDS) {
			t.Fatalf("warning result %+v", got)
		}
		if command == "check" && !reflect.DeepEqual(decodeData(t, x.Data), map[string]any{"valid": true}) {
			t.Fatal("warning made check invalid")
		}
	}
	valid := projectFixture(t)
	projectModule(t, valid, "units/a", "alpha", "ALP", "")
	for _, args := range [][]string{{"check", "--json"}, {"--json", "check"}} {
		x := successJSON(t, invoke(t, valid, args...), "check")
		if !reflect.DeepEqual(decodeData(t, x.Data), map[string]any{"valid": true}) {
			t.Fatalf("valid data %s", x.Data)
		}
	}
}
func Test_APP006ConfigurationErrors(t *testing.T) {
	for _, command := range []string{"check", "modules"} {
		for _, kind := range []string{"nongit", "missing", "invalid", "incompatible"} {
			root := t.TempDir()
			if kind != "nongit" {
				root = projectFixture(t)
			}
			switch kind {
			case "missing":
				if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
					t.Fatal(err)
				}
			case "invalid":
				projectPut(t, root, "fireweed.yaml", "version: [wrong]\n")
			case "incompatible":
				projectPut(t, root, "fireweed.yaml", "fireweed: '>99'\n")
			}
			got := invoke(t, root, command, "--json")
			x := decode(t, got, command)
			ds := detailedDiagnostics(t, got)
			if got.code != 2 || string(x.Data) != "null" || len(ds) != 1 || ds[0].Code != "C01" || ds[0].Severity != "error" || ds[0].Path != "fireweed.yaml" || ds[0].Message == "" {
				t.Fatalf("%s %s %+v", command, kind, got)
			}
			plain := invoke(t, root, command)
			if plain.code != 2 || plain.out != "" || !strings.Contains(plain.err, "C01") || !strings.Contains(plain.err, "fireweed.yaml") || strings.Contains(plain.err, "U01") {
				t.Fatalf("plain config %+v", plain)
			}
		}
	}
}
func Test_APP007ProjectArgumentsAndSafety(t *testing.T) {
	for _, command := range []string{"check", "modules"} {
		for _, args := range [][]string{{command, "extra"}, {command, "--unknown"}, {command, "--json=true"}, {command, "--", "--json"}} {
			assertUsage(t, invoke(t, t.TempDir(), args...))
		}
		for _, flag := range []string{"--help", "-h"} {
			successJSON(t, invoke(t, t.TempDir(), command, flag, "--json"), "help")
		}
		got := invoke(t, t.TempDir(), command, "extra", "--json")
		x := decode(t, got, command)
		if got.code != 2 || string(x.Data) != "null" || len(x.Diagnostics) != 1 || x.Diagnostics[0].Code != "U01" {
			t.Fatalf("operand loading order %+v", got)
		}
	}
	root := projectFixture(t)
	projectModule(t, root, "units/a", "alpha", "ALP", "")
	projectPut(t, root, ".gitignore", "ignored\n")
	projectPut(t, root, "ignored", "KEEP BYTES")
	gitFixtureCommand(t, root, "add", ".")
	gitFixtureCommand(t, root, "-c", "user.name=Fixture", "-c", "user.email=fixture@example.invalid", "commit", "-m", "fixture")
	before := gitFixtureCommand(t, root, "show-ref") + gitFixtureCommand(t, root, "status", "--porcelain=v1", "--untracked-files=all")
	if runtime.GOOS != "windows" {
		trap := t.TempDir()
		marker := filepath.Join(trap, "EXECUTED")
		realGit, err := exec.LookPath("git")
		if err != nil {
			t.Fatal(err)
		}
		quote := func(s string) string { return "'" + strings.ReplaceAll(s, "'", "'\"'\"'") + "'" }
		for _, name := range []string{"agent-trap", "codex", "claude", "fireweed", "fireweed-bootstrap", "curl", "wget", "ssh", "nc"} {
			projectPut(t, trap, name, "#!/bin/sh\necho invoked >> "+quote(marker)+"\nexit 99\n")
			if err = os.Chmod(filepath.Join(trap, name), 0755); err != nil {
				t.Fatal(err)
			}
		}
		projectPut(t, trap, "git", "#!/bin/sh\ncase \" $* \" in *' fetch '*|*' push '*|*' ls-remote '*|*' clone '*) echo network >> "+quote(marker)+";exit 99;;esac\nexec "+quote(realGit)+" \"$@\"\n")
		if err = os.Chmod(filepath.Join(trap, "git"), 0755); err != nil {
			t.Fatal(err)
		}
		t.Setenv("PATH", trap+string(os.PathListSeparator)+os.Getenv("PATH"))
		t.Setenv("FIREWEED_BIN", filepath.Join(trap, "fireweed-bootstrap"))
		for _, command := range []string{"modules", "check"} {
			successJSON(t, invoke(t, root, command, "--json"), command)
		}
		if _, err = os.Stat(marker); !os.IsNotExist(err) {
			t.Fatal("inspection executed external trap")
		}
	}
	after := gitFixtureCommand(t, root, "show-ref") + gitFixtureCommand(t, root, "status", "--porcelain=v1", "--untracked-files=all")
	ignored, err := os.ReadFile(filepath.Join(root, "ignored"))
	if before != after || err != nil || string(ignored) != "KEEP BYTES" {
		t.Fatal("inspection mutated files/ignored content/refs")
	}
	if _, err = os.Stat(filepath.Join(root, "MUST-NOT-RUN")); !os.IsNotExist(err) {
		t.Fatal("inspection ran hook")
	}
}
