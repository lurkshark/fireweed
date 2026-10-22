package evals

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"testing"

	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
)

func cliQueueChange(t *testing.T, root, path, target, prefix, metadata, edits string) {
	t.Helper()
	text := "---\nmodule: " + target + "\n" + metadata + "---\n\n## Motivation\n\nValidate authored intent.\n\n" + edits
	if _, err := c.ParseChange(text, path, prefix); err != nil {
		t.Fatalf("invalid valid-change fixture: %v", err)
	}
	projectPut(t, root, path, text)
}
func cliQueueExpected(t *testing.T, root string) []pathDiagnostic {
	t.Helper()
	s, err := p.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	ds := []pathDiagnostic{}
	for _, group := range [][]p.Diagnostic{p.Check(s), p.CheckQueue(s)} {
		for _, d := range group {
			ds = append(ds, pathDiagnostic{d.Code, d.Severity, d.Message, d.Path})
		}
	}
	sort.Slice(ds, func(i, j int) bool {
		a, b := ds[i], ds[j]
		return a.Path+"\x00"+a.Code+"\x00"+a.Severity+"\x00"+a.Message < b.Path+"\x00"+b.Code+"\x00"+b.Severity+"\x00"+b.Message
	})
	return ds
}
func cliQueueResult(t *testing.T, root string, valid bool, want []pathDiagnostic) {
	t.Helper()
	code := 0
	if !valid {
		code = 1
	}
	for _, args := range [][]string{{"check", "--json"}, {"--json", "check"}} {
		got := invoke(t, root, args...)
		x := decode(t, got, "check")
		if got.code != code || got.err != "" || !reflect.DeepEqual(decodeData(t, x.Data), map[string]any{"valid": valid}) || !reflect.DeepEqual(detailedDiagnostics(t, got), want) {
			t.Fatalf("queue JSON %+v want %+v", got, want)
		}
	}
	got := invoke(t, root, "check")
	if got.code != code || !reflect.DeepEqual(decodeData(t, json.RawMessage(got.out)), map[string]any{"valid": valid}) {
		t.Fatalf("queue plain %+v", got)
	}
	for _, d := range want {
		if !strings.Contains(got.err, d.Code) || !strings.Contains(got.err, d.Path) || !strings.Contains(got.err, d.Message) {
			t.Fatalf("missing plain diagnostic %+v", got)
		}
	}
	if len(want) == 0 && got.err != "" {
		t.Fatalf("unexpected stderr %q", got.err)
	}
}
func Test_APP008QueueComposition(t *testing.T) {
	root := projectFixture(t)
	projectModule(t, root, "units/a", "alpha", "ALP", "hooks: {eval: ''}\n")
	projectModule(t, root, "units/b", "beta", "BET", "")
	projectPut(t, root, "outside/SPEC.md", "warning outside")
	cliQueueChange(t, root, "changes/0001-stale.md", "alpha", "ALP", "bases: {ALP-001: deadbeef}\nset: {hooks: {boundaries: ''}}\n", "## Modified\n\n### ALP-001\nIt MUST retain this intended edit.\n")
	cliQueueChange(t, root, "changes/0002-ids.md", "beta", "BET", "fixes: [BET-999]\n", "## Added\n\n### BET-009\nIt MUST receive an assigned ID.\n")
	projectPut(t, root, "changes/0003-handle.md", "---\nmodule: beta\nblocked: waiting for new:foreign\n---\n\n## Motivation\n\nReject a foreign handle.\n")
	want := cliQueueExpected(t, root)
	counts := map[string]int{}
	for _, d := range want {
		counts[d.Code+"@"+d.Path]++
	}
	for _, key := range []string{"C02@outside/SPEC.md", "C07@units/a/SPEC.md", "C07@changes/0001-stale.md", "C10@changes/0001-stale.md", "C09@changes/0003-handle.md"} {
		if counts[key] != 1 {
			t.Fatalf("fixture or repetition %s %+v", key, want)
		}
	}
	if counts["C05@changes/0002-ids.md"] != 2 {
		t.Fatalf("ID/reference fixture %+v", want)
	}
	nested := filepath.Join(root, "units/a/nested")
	if err := os.MkdirAll(nested, 0755); err != nil {
		t.Fatal(err)
	}
	cliQueueResult(t, nested, false, want)
	currentSnapshot, current := projectExpected(t, root)
	got := invoke(t, nested, "modules", "--json")
	moduleJSON := decode(t, got, "modules")
	if !reflect.DeepEqual(decodeData(t, moduleJSON.Data), canonical(t, currentSnapshot.Modules)) {
		t.Fatal("modules data leaked projected queue state")
	}
	if got.code != 1 || !reflect.DeepEqual(detailedDiagnostics(t, got), current) {
		t.Fatalf("modules projected errors %+v", got)
	}
}
func Test_APP008QueueOrderAndWarnings(t *testing.T) {
	root := projectFixture(t)
	projectModule(t, root, "units/a", "alpha", "ALP", "")
	first := c.Spec{ID: "ALP-001", Tag: "[review]", Body: "It MUST retain this body."}
	for i, v := range []struct{ path, body string }{{"changes/0002-alpha.md", "It MUST be first."}, {"changes/00002-zulu.md", "It MUST be second."}, {"changes/00010-last.md", "It MUST be last."}} {
		meta := "bases: {ALP-001: " + c.SpecHash(first) + "}\n"
		if i == 2 {
			meta += "set: {preserve: [data/**]}\n"
		}
		cliQueueChange(t, root, v.path, "alpha", "ALP", meta, "## Modified\n\n### ALP-001 [review]\n"+v.body+"\n")
		first.Body = v.body
	}
	if want := cliQueueExpected(t, root); len(want) != 0 {
		t.Fatalf("ordered fixture %+v", want)
	}
	cliQueueResult(t, root, true, []pathDiagnostic{})
	cliQueueChange(t, root, "changes/0011-optout.md", "alpha", "ALP", "set: {hooks: {boundaries: ''}, boundaries_opt_out: 'visible queued reason'}\n", "")
	projectPut(t, root, "outside/SPEC.md", "warning outside")
	want := cliQueueExpected(t, root)
	if len(want) != 2 || want[0].Code != "D04" || want[0].Path != "changes/0011-optout.md" || !strings.Contains(want[0].Message, "visible queued reason") || want[1].Code != "C02" {
		t.Fatalf("warning fixture %+v", want)
	}
	cliQueueResult(t, root, true, want)
	got := invoke(t, root, "modules", "--json")
	moduleJSON := decode(t, got, "modules")
	currentSnapshot, _ := projectExpected(t, root)
	if !reflect.DeepEqual(decodeData(t, moduleJSON.Data), canonical(t, currentSnapshot.Modules)) {
		t.Fatal("modules data leaked queued hooks/specs")
	}
	if got.code != 0 || len(detailedDiagnostics(t, got)) != 1 || detailedDiagnostics(t, got)[0].Code != "C02" {
		t.Fatalf("modules included queued warning %+v", got)
	}
	// A structurally valid projected error affects only check.
	cliQueueChange(t, root, "changes/0012-dependency.md", "alpha", "ALP", "set: {depends_on: [missing]}\n", "")
	got = invoke(t, root, "modules", "--json")
	moduleJSON = decode(t, got, "modules")
	if !reflect.DeepEqual(decodeData(t, moduleJSON.Data), canonical(t, currentSnapshot.Modules)) {
		t.Fatal("modules data leaked queued dependencies")
	}
	if got.code != 0 || len(detailedDiagnostics(t, got)) != 1 {
		t.Fatalf("modules included projected errors %+v", got)
	}
	cliQueueResult(t, root, false, cliQueueExpected(t, root))
}
func Test_APP008MissingBasesAndQueueSafety(t *testing.T) {
	root := projectFixture(t)
	projectModule(t, root, "units/a", "alpha", "ALP", "")
	cliQueueChange(t, root, "changes/0001-missing.md", "alpha", "ALP", "", "## Modified\n\n### ALP-001\nIt MUST need a stamp.\n")
	projectPut(t, root, ".gitignore", "ignored\n")
	projectPut(t, root, "ignored", "KEEP BYTES")
	gitFixtureCommand(t, root, "add", ".")
	gitFixtureCommand(t, root, "-c", "user.name=Fixture", "-c", "user.email=fixture@example.invalid", "commit", "-m", "fixture")
	before := gitFixtureCommand(t, root, "show-ref") + gitFixtureCommand(t, root, "status", "--porcelain=v1", "--untracked-files=all")
	originals := map[string]string{}
	if err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			if entry.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		b, e := os.ReadFile(path)
		if e == nil {
			originals[path] = string(b)
		}
		return e
	}); err != nil {
		t.Fatal(err)
	}
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
		t.Cleanup(func() {
			if _, e := os.Stat(marker); !os.IsNotExist(e) {
				t.Error("queue check executed external trap")
			}
		})
	}
	want := cliQueueExpected(t, root)
	if len(want) != 1 || want[0].Code != "C10" {
		t.Fatalf("missing base fixture %+v", want)
	}
	cliQueueResult(t, root, false, want)
	after := gitFixtureCommand(t, root, "show-ref") + gitFixtureCommand(t, root, "status", "--porcelain=v1", "--untracked-files=all")
	if before != after {
		t.Fatal("queue check mutated refs/status")
	}
	for path, text := range originals {
		b, err := os.ReadFile(path)
		if err != nil || string(b) != text {
			t.Fatalf("mutated %s: %v", path, err)
		}
	}
	if _, err := os.Stat(filepath.Join(root, "MUST-NOT-RUN")); !os.IsNotExist(err) {
		t.Fatal(fmt.Sprintf("queue check ran hook: %v", err))
	}
}
