package evals

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	k "github.com/lurkshark/fireweed/packages/checks"
	c "github.com/lurkshark/fireweed/packages/contracts"
	g "github.com/lurkshark/fireweed/packages/gitstore"
	p "github.com/lurkshark/fireweed/packages/project"
)

const boundaryBaseline = ".fireweed/boundaries-baseline"

type boundaryHarness struct {
	Snapshot                           p.Snapshot
	Script, Output, Trace, Ready, Late string
}

func boundaryPut(t *testing.T, path, text string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(text), 0644); err != nil {
		t.Fatal(err)
	}
}
func boundaryFixture(t *testing.T) boundaryHarness {
	t.Helper()
	s, m := hookFixture(t)
	m.Disposable = []string{"internal/**"}
	aux := filepath.Join(t.TempDir(), "helper '雪 space")
	h := boundaryHarness{Snapshot: s, Script: filepath.Join(aux, "checker.sh"), Output: filepath.Join(aux, "output"), Trace: filepath.Join(aux, "trace"), Ready: filepath.Join(aux, "ready"), Late: filepath.Join(aux, "late")}
	script := `#!/bin/sh
set -eu
if IFS= read -r unexpected; then exit 93; fi
if [ "$4" = stderr ]; then cat "$2" >&2; else cat "$2"; fi
printf '%s\n' "$1" "$PWD" "$FIREWEED_MODULE" "$FIREWEED_MODULE_DIR" "$FIREWEED_REPORT" "$FIREWEED_BIN" "$FIREWEED_INHERITED_HOOK" "$8" "$9" "${10}" "${11}" EOF >> "$3"
if [ "$4" = hold ]; then
 ( printf ready > "$7.child.ready"; sleep 2; printf late > "$7.child" ) &
fi
printf ready > "$6"
case "$4" in
 signal) kill -TERM $$ ;;
 hold) sleep 2; printf late > "$7"; wait ;;
 gated) while [ ! -f "$7.release" ]; do sleep 0.01; done ;;
esac
exit "$5"
`
	boundaryPut(t, h.Script, script)
	boundaryPut(t, h.Output, "")
	m.Hooks.Boundaries = h.command("A", "zero", 0)
	h.Snapshot.Modules = []c.Module{m, {Name: "beta", IDPrefix: "BET", Dir: "units/beta", Zones: c.Zones{Disposable: []string{}}, Hooks: c.Hooks{}}, {Name: "gamma", IDPrefix: "GAM", Dir: "units/gamma", Zones: c.Zones{Disposable: []string{"internal/**"}}, BoundariesOptOut: "legacy interoperability", Hooks: c.Hooks{Boundaries: hookCommand(t, "marker", filepath.Join(s.Root, "MUST-NOT-RUN"))}}}
	return h
}
func (h boundaryHarness) args(label, mode string, code int) []string {
	return []string{label, h.Output, h.Trace, mode, fmt.Sprint(code), h.Ready, h.Late}
}
func (h boundaryHarness) command(label, mode string, code int) string {
	command := "exec sh " + shellLiteral(h.Script)
	for _, arg := range h.args(label, mode, code) {
		command += " " + shellLiteral(arg)
	}
	return command + " {module} {module_dir} {report} {git_common_dir}"
}
func boundaryExpected(t *testing.T, ds []k.Diagnostic, code string, paths, lines []string) {
	t.Helper()
	if ds == nil || len(ds) != len(paths) {
		t.Fatalf("diagnostics %+v expected paths %v", ds, paths)
	}
	keys := []string{}
	for i, d := range ds {
		lower := strings.ToLower(d.Message)
		if d.Code != code || d.Severity != "error" || d.Path != paths[i] || strings.TrimSpace(d.Message) == "" || !(strings.Contains(lower, "correct") || strings.Contains(lower, "repair") || strings.Contains(lower, "configure") || strings.Contains(lower, "fix")) {
			t.Fatalf("corrective attributed diagnostic %+v", d)
		}
		if i < len(lines) && !strings.Contains(d.Message, lines[i]) {
			t.Fatalf("lost violation %q: %+v", lines[i], d)
		}
		keys = append(keys, d.Path+"\x00"+d.Code+"\x00"+d.Severity+"\x00"+d.Message)
	}
	if !sort.StringsAreSorted(keys) {
		t.Fatalf("unsorted diagnostics %+v", ds)
	}
}
func boundaryTrace(t *testing.T, h boundaryHarness, labels []string) {
	t.Helper()
	data, err := os.ReadFile(h.Trace)
	if err != nil {
		t.Fatal(err)
	}
	fields := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	if len(fields) != 12*len(labels) {
		t.Fatalf("trace %q", data)
	}
	repo, err := g.Open(h.Snapshot.Root)
	if err != nil {
		t.Fatal(err)
	}
	exe, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	for i, label := range labels {
		want := []string{label, repo.Root, "", "", "", exe, os.Getenv("FIREWEED_INHERITED_HOOK"), "", "", "", repo.CommonDir, "EOF"}
		if !reflect.DeepEqual(fields[i*12:(i+1)*12], want) {
			t.Fatalf("trace %v expected %v", fields[i*12:(i+1)*12], want)
		}
	}
}
func boundaryEnv(t *testing.T) {
	t.Helper()
	t.Setenv("FIREWEED_INHERITED_HOOK", "parent '雪")
	for _, name := range []string{"FIREWEED_MODULE", "FIREWEED_MODULE_DIR", "FIREWEED_REPORT", "FIREWEED_BIN"} {
		t.Setenv(name, "stale-controller-value")
	}
}
func boundaryNoStartup(t *testing.T, h boundaryHarness) {
	t.Helper()
	for _, path := range []string{h.Trace, h.Ready, filepath.Join(h.Snapshot.Root, "MUST-NOT-RUN")} {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Fatalf("unexpected startup %s: %v", path, err)
		}
	}
}
func Test_CHK009_CHK010_CHK011_CHK012BoundaryHelperControl(t *testing.T) {
	boundaryEnv(t)
	h := boundaryFixture(t)
	boundaryPut(t, h.Output, "alpha: actual helper evidence\n")
	repo, err := g.Open(h.Snapshot.Root)
	if err != nil {
		t.Fatal(err)
	}
	exe, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	args := append([]string{h.Script}, h.args("control", "stderr", 17)...)
	args = append(args, "", "", "", repo.CommonDir)
	command := exec.Command("sh", args...)
	command.Dir = repo.Root
	command.Env = append(os.Environ(), "FIREWEED_MODULE=", "FIREWEED_MODULE_DIR=", "FIREWEED_REPORT=", "FIREWEED_BIN="+exe)
	output, err := command.CombinedOutput()
	if err == nil || command.ProcessState.ExitCode() != 17 || string(output) != "alpha: actual helper evidence\n" {
		t.Fatalf("helper control %v %q", err, output)
	}
	boundaryTrace(t, h, []string{"control"})
	if data, err := os.ReadFile(h.Ready); err != nil || string(data) != "ready" {
		t.Fatalf("helper readiness %v %q", err, data)
	}
}
func Test_CHK009ResolvedSelectionAndGlobalIdentity(t *testing.T) {
	boundaryEnv(t)
	h := boundaryFixture(t)
	h.Snapshot.Root = filepath.Join(h.Snapshot.Root, "nested")
	h.Snapshot.Modules[0].Hooks.Boundaries = h.command("B", "zero", 0)
	h.Snapshot.Modules[1].Hooks.Boundaries = h.command("A", "zero", 0)
	h.Snapshot.Modules = append(h.Snapshot.Modules, c.Module{Name: "optional", IDPrefix: "OPT", Dir: "units/optional", Hooks: c.Hooks{Boundaries: h.command("A", "zero", 0)}})
	boundaryPut(t, h.Output, "informational progress\n")
	before := jsonSnapshot(t, h.Snapshot)
	ds := k.Boundaries(context.Background(), h.Snapshot)
	boundaryExpected(t, ds, "", []string{}, nil)
	boundaryTrace(t, h, []string{"A", "B"})
	if jsonSnapshot(t, h.Snapshot) != before {
		t.Fatal("selection mutated snapshot")
	}
	if _, err := os.Stat(filepath.Join(filepath.Dir(h.Snapshot.Root), "MUST-NOT-RUN")); !os.IsNotExist(err) {
		t.Fatal("resolved override replaced by default or opted-out hook")
	}
}
func Test_CHK009RequiredAndOptionalCommands(t *testing.T) {
	boundaryEnv(t)
	for _, mode := range []string{"empty", "optional-empty", "opted-out", "missing", "blank", "two-missing"} {
		t.Run(mode, func(t *testing.T) {
			h := boundaryFixture(t)
			alpha := h.Snapshot.Modules[0]
			paths := []string{}
			switch mode {
			case "empty":
				h.Snapshot.Modules = []c.Module{}
			case "optional-empty":
				alpha.Disposable = []string{}
				alpha.Hooks.Boundaries = ""
				h.Snapshot.Modules = []c.Module{alpha}
			case "opted-out":
				alpha.BoundariesOptOut = "explicit reason"
				h.Snapshot.Modules = []c.Module{alpha}
			default:
				alpha.Hooks.Boundaries = ""
				if mode == "blank" {
					alpha.Hooks.Boundaries = " \t\n"
				}
				h.Snapshot.Modules = []c.Module{alpha, {Name: "zeta", IDPrefix: "ZET", Dir: "units/zeta", Hooks: c.Hooks{Boundaries: h.command("Z", "zero", 0)}}}
				paths = []string{alpha.Dir}
				if mode == "two-missing" {
					h.Snapshot.Modules = append(h.Snapshot.Modules, c.Module{Name: "aardvark", IDPrefix: "ARD", Dir: "units/aardvark", Zones: c.Zones{Disposable: []string{"internal/**"}}})
					paths = []string{"units/aardvark", alpha.Dir}
					sort.Strings(paths)
				}
			}
			ds := k.Boundaries(context.Background(), h.Snapshot)
			boundaryExpected(t, ds, "C07", paths, nil)
			boundaryNoStartup(t, h)
		})
	}
}
func Test_CHK010AttributionAndExactBaseline(t *testing.T) {
	boundaryEnv(t)
	for _, tt := range []struct {
		name, baseline, output string
		paths, lines           []string
	}{
		{"no baseline", "", " \talpha: first violation \r\n beta second violation\t\nAlpha: informational\nalpha informational: detail\nalpha-extra: informational\ngamma: intentionally ignored\n", []string{coverageBoundaryAlphaDir(), "units/beta"}, []string{"alpha: first violation", "beta second violation"}},
		{"exact trimmed", " \talpha: existing \r\n", " alpha: existing \r\ngamma ignored\n", []string{}, nil},
		{"prefix", "alpha:\n", "alpha: different violation\n", []string{coverageBoundaryAlphaDir()}, []string{"alpha: different violation"}},
		{"substring", "existing\n", "alpha: existing\n", []string{coverageBoundaryAlphaDir()}, []string{"alpha: existing"}},
		{"case folding", "Alpha: existing\n", "alpha: existing\n", []string{coverageBoundaryAlphaDir()}, []string{"alpha: existing"}},
		{"different message", "alpha: existing\n", "alpha: existing\nalpha: changed\n", []string{coverageBoundaryAlphaDir()}, []string{"alpha: changed"}},
		{"colon spacing", "", "alpha : spaced violation\n", []string{coverageBoundaryAlphaDir()}, []string{"alpha : spaced violation"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			h := boundaryFixture(t)
			boundaryPut(t, h.Output, tt.output)
			if tt.baseline != "" {
				boundaryPut(t, filepath.Join(h.Snapshot.Root, boundaryBaseline), tt.baseline)
			}
			ds := k.Boundaries(context.Background(), h.Snapshot)
			boundaryExpected(t, ds, "C15", tt.paths, tt.lines)
			boundaryTrace(t, h, []string{"A"})
		})
	}
}
func coverageBoundaryAlphaDir() string { return "units/a '雪 {report} $HOME ; $(touch INJECTED)" }
func Test_CHK010BaselineReadFailures(t *testing.T) {
	boundaryEnv(t)
	for _, mode := range []string{"directory", "unreadable", "escaping-symlink"} {
		t.Run(mode, func(t *testing.T) {
			h := boundaryFixture(t)
			path := filepath.Join(h.Snapshot.Root, boundaryBaseline)
			if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
				t.Fatal(err)
			}
			switch mode {
			case "directory":
				if err := os.Mkdir(path, 0755); err != nil {
					t.Fatal(err)
				}
			case "unreadable":
				boundaryPut(t, path, "alpha: hidden\n")
				if err := os.Chmod(path, 0000); err != nil {
					t.Fatal(err)
				}
				defer os.Chmod(path, 0644)
				if _, err := os.ReadFile(path); err == nil {
					t.Skip("current user bypasses file read permissions")
				}
			case "escaping-symlink":
				outside := filepath.Join(t.TempDir(), "outside")
				boundaryPut(t, outside, "alpha: hidden\n")
				if err := os.Symlink(outside, path); err != nil {
					t.Skipf("symlink unavailable: %v", err)
				}
			}
			ds := k.Boundaries(context.Background(), h.Snapshot)
			boundaryExpected(t, ds, "C15", []string{boundaryBaseline}, nil)
			boundaryNoStartup(t, h)
		})
	}
}
func Test_CHK011OrdinaryExitEvidence(t *testing.T) {
	boundaryEnv(t)
	for _, tt := range []struct {
		name, baseline, output string
		code                   int
		fail                   bool
	}{
		{"empty failure", "", "", 17, true},
		{"unknown failure", "", "compiler failed\n", 23, true},
		{"unknown baseline cannot waive", "compiler failed\n", "compiler failed\n", 23, true},
		{"mixed attribution", "alpha: existing\n", "alpha: existing\ncompiler failed\n", 17, true},
		{"all baseline", "alpha: existing\n", "alpha: existing\n", 17, false},
		{"all opted out", "", "gamma: legacy import\n", 17, false},
		{"baseline and opt out", "alpha: existing\n", "alpha: existing\ngamma: legacy import\n", 17, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			h := boundaryFixture(t)
			h.Snapshot.Modules[0].Hooks.Boundaries = h.command("A", "zero", tt.code)
			boundaryPut(t, h.Output, tt.output)
			if tt.baseline != "" {
				boundaryPut(t, filepath.Join(h.Snapshot.Root, boundaryBaseline), tt.baseline)
			}
			ds := k.Boundaries(context.Background(), h.Snapshot)
			if !tt.fail {
				boundaryExpected(t, ds, "", []string{}, nil)
			} else {
				if ds == nil || len(ds) == 0 {
					t.Fatal("nonzero checker accepted")
				}
				found := false
				for _, d := range ds {
					if d.Code != "C15" || d.Severity != "error" || !strings.Contains(d.Message, fmt.Sprint(tt.code)) {
						t.Fatalf("missing actual exit %+v", ds)
					}
					if strings.Contains(d.Message, tt.output) {
						found = true
					}
				}
				if !found {
					t.Fatalf("lost actual output %+v", ds)
				}
				boundaryExpected(t, ds, "C15", make([]string, len(ds)), nil)
			}
			boundaryTrace(t, h, []string{"A"})
		})
	}
}
func boundaryWaitReady(t *testing.T, h boundaryHarness, done <-chan []k.Diagnostic, child bool) {
	t.Helper()
	deadline := time.NewTimer(2 * time.Second)
	defer deadline.Stop()
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()
	for {
		_, err := os.Stat(h.Ready)
		_, childErr := os.Stat(h.Late + ".child.ready")
		if err == nil && (!child || childErr == nil) {
			return
		}
		select {
		case ds := <-done:
			t.Fatalf("checker ended before readiness %+v", ds)
		case <-deadline.C:
			t.Fatal("checker did not become ready")
		case <-ticker.C:
		}
	}
}
func boundaryFatal(t *testing.T, ds []k.Diagnostic, exit int, output string) {
	t.Helper()
	if ds == nil || len(ds) == 0 {
		t.Fatal("fatal checker accepted")
	}
	found := false
	for _, d := range ds {
		if d.Code != "C15" || d.Severity != "error" {
			t.Fatalf("fatal diagnostic %+v", d)
		}
		lower := strings.ToLower(d.Message)
		if strings.Contains(d.Message, fmt.Sprint(exit)) && strings.Contains(d.Message, strings.TrimSpace(output)) && (strings.Contains(lower, "repair") || strings.Contains(lower, "correct") || strings.Contains(lower, "configure")) {
			found = true
		}
	}
	if !found {
		t.Fatalf("fatal output/exit/remedy missing %+v", ds)
	}
}
func boundaryNoLate(t *testing.T, h boundaryHarness) {
	t.Helper()
	time.Sleep(2200 * time.Millisecond)
	for _, path := range []string{h.Late, h.Late + ".child"} {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Fatalf("cancelled checker survived: %s", path)
		}
	}
}
func Test_CHK011FatalExecutionAndPreconditions(t *testing.T) {
	boundaryEnv(t)
	for _, mode := range []string{"nil-context", "pre-cancelled", "invalid-timeout", "invalid-repository", "missing-shell", "signal-baselined", "timeout-baselined"} {
		t.Run(mode, func(t *testing.T) {
			h := boundaryFixture(t)
			ctx := context.Background()
			exit := 2
			output := ""
			switch mode {
			case "nil-context":
				ctx = nil
			case "pre-cancelled":
				cancelled, cancel := context.WithCancel(ctx)
				cancel()
				ctx = cancelled
				exit = -1
			case "invalid-timeout":
				h.Snapshot.Config.Timeout = "not-a-duration"
			case "invalid-repository":
				h.Snapshot.Root = t.TempDir()
			case "missing-shell":
				realGit, err := exec.LookPath("git")
				if err != nil {
					t.Fatal(err)
				}
				bin := t.TempDir()
				if err := os.Symlink(realGit, filepath.Join(bin, "git")); err != nil {
					t.Skipf("symlink unavailable: %v", err)
				}
				t.Setenv("PATH", bin)
				if _, err := g.Open(h.Snapshot.Root); err != nil {
					t.Fatalf("git-only path premise %v", err)
				}
			default:
				output = "alpha: existing\n"
				exit = -1
				boundaryPut(t, h.Output, output)
				boundaryPut(t, filepath.Join(h.Snapshot.Root, boundaryBaseline), output)
				hookMode := "signal"
				if mode == "timeout-baselined" {
					hookMode = "hold"
					h.Snapshot.Config.Timeout = "120ms"
				}
				h.Snapshot.Modules[0].Hooks.Boundaries = h.command("A", hookMode, 0)
			}
			start := time.Now()
			ds := k.Boundaries(ctx, h.Snapshot)
			boundaryFatal(t, ds, exit, output)
			if time.Since(start) > 1500*time.Millisecond {
				t.Fatal("fatal boundary exceeded execution bound")
			}
			if strings.HasSuffix(mode, "baselined") {
				boundaryTrace(t, h, []string{"A"})
				if mode == "timeout-baselined" {
					if _, err := os.Stat(h.Late + ".child.ready"); err != nil {
						t.Fatal("child cancellation premise", err)
					}
					boundaryNoLate(t, h)
				}
			} else {
				boundaryNoStartup(t, h)
			}
		})
	}
}
func Test_CHK011ParentCancellationAfterOptedOutOutput(t *testing.T) {
	boundaryEnv(t)
	h := boundaryFixture(t)
	boundaryPut(t, h.Output, "gamma: intentionally ignored\n")
	h.Snapshot.Modules[0].Hooks.Boundaries = h.command("A", "hold", 0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	done := make(chan []k.Diagnostic, 1)
	go func() { done <- k.Boundaries(ctx, h.Snapshot) }()
	boundaryWaitReady(t, h, done, true)
	start := time.Now()
	cancel()
	select {
	case ds := <-done:
		boundaryFatal(t, ds, -1, "gamma: intentionally ignored")
	case <-time.After(1500 * time.Millisecond):
		t.Fatal("parent cancellation did not terminate checker")
	}
	if time.Since(start) > 1500*time.Millisecond {
		t.Fatal("slow cancellation")
	}
	boundaryNoLate(t, h)
}
func boundaryFiles(t *testing.T, root string) map[string]string {
	t.Helper()
	out := map[string]string{}
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			if entry.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		relative, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		out[relative] = string(data)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return out
}
func Test_CHK012OwnershipAndSafetyAcrossResults(t *testing.T) {
	boundaryEnv(t)
	for _, mode := range []string{"success", "violation", "fatal"} {
		t.Run(mode, func(t *testing.T) {
			h := boundaryFixture(t)
			root := h.Snapshot.Root
			boundaryPut(t, filepath.Join(root, ".gitignore"), "ignored-sentinel\n")
			boundaryPut(t, filepath.Join(root, "ignored-sentinel"), "ignored user bytes")
			output := "informational only\n"
			if mode != "success" {
				output = "alpha: current violation\n"
			}
			boundaryPut(t, h.Output, output)
			baseline := "alpha: existing\n"
			if mode == "fatal" {
				baseline = output
				h.Snapshot.Config.Timeout = "120ms"
				h.Snapshot.Modules[0].Hooks.Boundaries = h.command("A", "hold", 0)
			}
			boundaryPut(t, filepath.Join(root, boundaryBaseline), baseline)
			hookGit(t, root, "add", ".")
			hookGit(t, root, "commit", "-qm", "boundary fixture")
			alphaDir := h.Snapshot.Modules[0].Dir
			h.Snapshot.Modules = []c.Module{h.Snapshot.Modules[2], h.Snapshot.Modules[0], h.Snapshot.Modules[1]}
			marker := installHookTraps(t, root)
			snapshotBefore := jsonSnapshot(t, h.Snapshot)
			filesBefore := boundaryFiles(t, root)
			refsBefore := hookGit(t, root, "show-ref")
			statusBefore := hookGit(t, root, "status", "--porcelain=v1", "--untracked-files=all")
			envBefore := os.Environ()
			cwdBefore, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}
			first := k.Boundaries(context.Background(), h.Snapshot)
			second := k.Boundaries(context.Background(), h.Snapshot)
			if mode == "success" {
				boundaryExpected(t, first, "", []string{}, nil)
				boundaryExpected(t, second, "", []string{}, nil)
			} else if mode == "violation" {
				boundaryExpected(t, first, "C15", []string{alphaDir}, []string{strings.TrimSpace(output)})
				if !reflect.DeepEqual(first, second) {
					t.Fatalf("nondeterministic diagnostics %+v %+v", first, second)
				}
			} else {
				boundaryFatal(t, first, -1, output)
				boundaryFatal(t, second, -1, output)
				boundaryNoLate(t, h)
			}
			saved := jsonSnapshot(t, second)
			if len(first) == 0 && len(second) == 0 {
				first = append(first, k.Diagnostic{Message: "first empty result"})
				second = append(second, k.Diagnostic{Message: "second empty result"})
				if first[0].Message != "first empty result" || second[0].Message != "second empty result" {
					t.Fatal("empty results share writable backing storage")
				}
				first, second = first[:0], second[:0]
			}
			if len(first) > 0 {
				first[0].Message = "mutated first result"
			}
			first = append(first, k.Diagnostic{Message: "new result"})
			_ = first
			cwdAfter, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}
			if saved != jsonSnapshot(t, second) || snapshotBefore != jsonSnapshot(t, h.Snapshot) || !reflect.DeepEqual(filesBefore, boundaryFiles(t, root)) || refsBefore != hookGit(t, root, "show-ref") || statusBefore != hookGit(t, root, "status", "--porcelain=v1", "--untracked-files=all") || !reflect.DeepEqual(envBefore, os.Environ()) || cwdBefore != cwdAfter {
				t.Fatal("boundary calls shared results or mutated caller/repository state")
			}
			if _, err := os.Stat(marker); !os.IsNotExist(err) {
				t.Fatal("boundary ran forbidden external command")
			}
			if _, err := os.Stat(filepath.Join(root, "MUST-NOT-RUN")); !os.IsNotExist(err) {
				t.Fatal("boundary ran eval/setup/default/opted-out hook")
			}
			boundaryTrace(t, h, []string{"A", "A"})
		})
	}
}
func Test_CHK012SortDeduplicateAndOverlap(t *testing.T) {
	boundaryEnv(t)
	h := boundaryFixture(t)
	boundaryPut(t, h.Output, "beta: later\nalpha: z violation\nalpha: a violation\n")
	h.Snapshot.Modules[1].Hooks.Boundaries = h.command("B", "stderr", 0)
	first := k.Boundaries(context.Background(), h.Snapshot)
	boundaryExpected(t, first, "C15", []string{h.Snapshot.Modules[0].Dir, h.Snapshot.Modules[0].Dir, "units/beta"}, []string{"alpha: a violation", "alpha: z violation", "beta: later"})
	boundaryTrace(t, h, []string{"A", "B"})
	long := boundaryFixture(t)
	brief := boundaryFixture(t)
	boundaryPut(t, long.Output, "alpha: existing\n")
	boundaryPut(t, filepath.Join(long.Snapshot.Root, boundaryBaseline), "alpha: existing\n")
	long.Snapshot.Modules[0].Hooks.Boundaries = long.command("long", "hold", 0)
	brief.Snapshot.Modules[0].Hooks.Boundaries = brief.command("brief", "gated", 0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	longDone := make(chan []k.Diagnostic, 1)
	briefDone := make(chan []k.Diagnostic, 1)
	long.Snapshot.Modules = []c.Module{long.Snapshot.Modules[2], long.Snapshot.Modules[0], long.Snapshot.Modules[1]}
	brief.Snapshot.Modules = []c.Module{brief.Snapshot.Modules[2], brief.Snapshot.Modules[0], brief.Snapshot.Modules[1]}
	envBefore := os.Environ()
	cwdBefore, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	longFiles, briefFiles := boundaryFiles(t, long.Snapshot.Root), boundaryFiles(t, brief.Snapshot.Root)
	before := jsonSnapshot(t, []p.Snapshot{long.Snapshot, brief.Snapshot})
	briefCtx, briefCancel := context.WithCancel(context.Background())
	defer briefCancel()
	t.Cleanup(func() { boundaryPut(t, brief.Late+".release", "release") })
	go func() { longDone <- k.Boundaries(ctx, long.Snapshot) }()
	boundaryWaitReady(t, long, longDone, true)
	go func() { briefDone <- k.Boundaries(briefCtx, brief.Snapshot) }()
	boundaryWaitReady(t, brief, briefDone, false)
	cancel()
	select {
	case ds := <-longDone:
		boundaryFatal(t, ds, -1, "alpha: existing")
	case <-time.After(1500 * time.Millisecond):
		t.Fatal("overlap cancellation stuck")
	}
	select {
	case ds := <-briefDone:
		t.Fatalf("independent checker was cancelled early %+v", ds)
	default:
	}
	boundaryPut(t, brief.Late+".release", "release")
	select {
	case ds := <-briefDone:
		boundaryExpected(t, ds, "", []string{}, nil)
	case <-time.After(1500 * time.Millisecond):
		t.Fatal("independent checker did not finish")
	}
	cwdAfter, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if before != jsonSnapshot(t, []p.Snapshot{long.Snapshot, brief.Snapshot}) || cwdBefore != cwdAfter || !reflect.DeepEqual(envBefore, os.Environ()) || !reflect.DeepEqual(longFiles, boundaryFiles(t, long.Snapshot.Root)) || !reflect.DeepEqual(briefFiles, boundaryFiles(t, brief.Snapshot.Root)) {
		t.Fatal("overlap mutated inputs, parent state or baseline/repository files")
	}
	boundaryTrace(t, long, []string{"long"})
	boundaryTrace(t, brief, []string{"brief"})
	boundaryNoLate(t, long)
}

func Test_CHK010NestedCanonicalBaseline(t *testing.T) {
	boundaryEnv(t)
	h := boundaryFixture(t)
	boundaryPut(t, h.Output, "alpha: existing\n")
	boundaryPut(t, filepath.Join(h.Snapshot.Root, boundaryBaseline), "alpha: existing\n")
	h.Snapshot.Root = filepath.Join(h.Snapshot.Root, "nested")
	boundaryPut(t, filepath.Join(h.Snapshot.Root, boundaryBaseline), "alpha: unrelated nested policy\n")
	before := boundaryFiles(t, filepath.Dir(h.Snapshot.Root))
	ds := k.Boundaries(context.Background(), h.Snapshot)
	boundaryExpected(t, ds, "", []string{}, nil)
	boundaryTrace(t, h, []string{"A"})
	if !reflect.DeepEqual(before, boundaryFiles(t, filepath.Dir(h.Snapshot.Root))) {
		t.Fatal("canonical baseline lookup mutated files")
	}
}
func Test_CHK011NonzeroAttributedViolations(t *testing.T) {
	boundaryEnv(t)
	h := boundaryFixture(t)
	h.Snapshot.Modules[0].Hooks.Boundaries = h.command("A", "stderr", 17)
	boundaryPut(t, h.Output, "beta: unexpected dependency\nalpha: unexpected interface reach\ngamma: ignored\n")
	boundaryPut(t, filepath.Join(h.Snapshot.Root, boundaryBaseline), "alpha: existing\n")
	ds := k.Boundaries(context.Background(), h.Snapshot)
	boundaryExpected(t, ds, "C15", []string{h.Snapshot.Modules[0].Dir, "units/beta"}, []string{"alpha: unexpected interface reach", "beta: unexpected dependency"})
	boundaryTrace(t, h, []string{"A"})
}
func Test_CHK012BoundaryGateHelperControl(t *testing.T) {
	boundaryEnv(t)
	h := boundaryFixture(t)
	repo, err := g.Open(h.Snapshot.Root)
	if err != nil {
		t.Fatal(err)
	}
	exe, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	args := append([]string{h.Script}, h.args("gate-control", "gated", 0)...)
	args = append(args, "", "", "", repo.CommonDir)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	command := exec.CommandContext(ctx, "sh", args...)
	command.Dir = repo.Root
	command.Env = append(os.Environ(), "FIREWEED_MODULE=", "FIREWEED_MODULE_DIR=", "FIREWEED_REPORT=", "FIREWEED_BIN="+exe)
	type completion struct {
		output string
		err    error
	}
	done := make(chan completion, 1)
	go func() { output, err := command.CombinedOutput(); done <- completion{string(output), err} }()
	ready := time.NewTimer(2 * time.Second)
	defer ready.Stop()
	tick := time.NewTicker(5 * time.Millisecond)
	defer tick.Stop()
	waiting := true
	for waiting {
		if _, err := os.Stat(h.Ready); err == nil {
			break
		}
		select {
		case c := <-done:
			t.Fatalf("gate helper ended before ready: %+v", c)
		case <-ready.C:
			t.Fatal("gate helper not ready")
		case <-tick.C:
		}
	}
	select {
	case c := <-done:
		t.Fatalf("gate helper completed before release: %+v", c)
	case <-time.After(50 * time.Millisecond):
	}
	boundaryPut(t, h.Late+".release", "release")
	select {
	case c := <-done:
		if c.err != nil || c.output != "" {
			t.Fatalf("gate helper failed after release: %+v", c)
		}
	case <-time.After(time.Second):
		t.Fatal("gate helper ignored release")
	}
	boundaryTrace(t, h, []string{"gate-control"})
}
