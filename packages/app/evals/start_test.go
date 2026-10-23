package evals

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	app "github.com/lurkshark/fireweed/packages/app"
	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

const cliStartPath = "changes/0001-current.md"

type cliStartCase struct {
	Root, Where, Kind, Phase, Head, Slug, Path string
	Warning                                    bool
	FailedObjects                              int
}

func cliStartFixture(t *testing.T, format, kind, phase string, linked, nested, warning bool) cliStartCase {
	t.Helper()
	root, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	gitFixtureCommand(t, root, "init", "--initial-branch=main", "--object-format="+format)
	gitFixtureCommand(t, root, "config", "user.name", "Start CLI Fixture")
	gitFixtureCommand(t, root, "config", "user.email", "start@example.invalid")
	t.Setenv("GIT_AUTHOR_DATE", "2001-01-01T00:00:00Z")
	t.Setenv("GIT_COMMITTER_DATE", "2001-01-01T00:00:00Z")
	cfg := "fireweed: '>=0.4, <0.5'\n" + strings.Replace(projectConfig, "base: main", "base: approval-base", 1)
	projectPut(t, root, "fireweed.yaml", cfg)
	projectPut(t, root, ".gitignore", "ignored/**\n")
	projectPut(t, root, "SYSTEM.md", "Start system bytes.\r\n")
	projectModule(t, root, "units/a", "alpha", "ALP", "depends_on: [beta]\npreserve: [retained/**]\nappend_only: [events/**]\n")
	projectModule(t, root, "units/b", "beta", "BET", "")
	for name, text := range map[string]string{"units/a/api.go": "Current API.\n", "units/a/public/new.txt": "Projected API.\r\n", "units/a/evals/case.txt": "Retained eval.\n", "units/a/internal/private.txt": "Disposable implementation.\n", "units/a/retained/data.txt": "Preserved bytes.\n", "units/a/events/log.txt": "Append-only bytes.\n", "units/b/api.go": "Direct dependency API.\n", "private/secret.txt": "Unrelated private bytes.\n", "ignored/secret.txt": "Ignored bytes.\n"} {
		projectPut(t, root, name, text)
	}
	if warning {
		projectModule(t, root, "outside/a", "outside", "OUT", "")
	}
	gitFixtureCommand(t, root, "add", ".")
	gitFixtureCommand(t, root, "commit", "-qm", "Start project baseline")
	base := strings.TrimSpace(gitFixtureCommand(t, root, "rev-parse", "HEAD"))
	gitFixtureCommand(t, root, "update-ref", "refs/heads/approval-base", base)
	metadata := ""
	edits := ""
	switch kind {
	case "spec":
		metadata = "set: {interface: ['public/**']}\n"
		id := "new:visible"
		if phase != "" {
			id = "ALP-002"
		}
		edits = "## Added\n\n### " + id + "\nThe new behavior MUST retain " + id + ".\n"
	case "fix":
		metadata = "fixes: [ALP-001]\n"
	case "regenerate":
		metadata = "regenerate: true\n"
	case "maintenance":
	default:
		t.Fatal("unknown fixture kind", kind)
	}
	if phase != "" {
		metadata += "run: retained-cli-run\nphase: " + phase + "\n"
	}
	cliQueueChange(t, root, cliStartPath, "alpha", "ALP", metadata, edits)
	gitFixtureCommand(t, root, "add", ".")
	gitFixtureCommand(t, root, "commit", "-qm", "Authored start intent")
	head := strings.TrimSpace(gitFixtureCommand(t, root, "rev-parse", "HEAD"))
	if !warning {
		gitFixtureCommand(t, root, "update-ref", "refs/heads/approval-base", head)
	}
	if linked {
		peer := filepath.Join(t.TempDir(), "linked ' start")
		gitFixtureCommand(t, root, "worktree", "add", "--detach", peer)
		root, err = filepath.EvalSymlinks(peer)
		if err != nil {
			t.Fatal(err)
		}
	}
	projectPut(t, root, "ignored/secret.txt", "Ignored bytes.\n")
	where := root
	if nested {
		where = filepath.Join(root, "units/a")
	}
	s, err := p.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range append(p.Check(s), p.CheckQueue(s)...) {
		if d.Severity == "error" {
			t.Fatal("invalid start premise", d)
		}
	}
	if raw := gitFixtureCommand(t, root, "status", "--porcelain"); raw != "" {
		t.Fatal("dirty fixture", raw)
	}
	other := filepath.Join(t.TempDir(), "other worktree")
	gitFixtureCommand(t, root, "worktree", "add", "--detach", other)
	if _, e := w.CreateClaim(other, "unrelated", "other-run", head, "opaque unrelated claim state"); e != nil {
		t.Fatal(e)
	}
	return cliStartCase{root, where, kind, phase, head, "current", cliStartPath, warning, 0}
}
func cliStartPhase(f cliStartCase) string {
	if f.Phase != "" {
		return f.Phase
	}
	if f.Kind == "spec" || f.Kind == "fix" {
		return "evals"
	}
	return "impl"
}
func cliStartOwned(t *testing.T, f cliStartCase) w.WorkerContext {
	t.Helper()
	claim, err := w.LoadClaim(f.Root, f.Slug)
	if err != nil {
		t.Fatalf("start did not retain a valid owned claim: %v", err)
	}
	raw, err := os.ReadFile(filepath.Join(f.Root, f.Path))
	if err != nil {
		t.Fatal(err)
	}
	if claim.Worktree != f.Root || claim.Base != f.Head || claim.State != string(raw) || claim.Run == "" || claim.Object == "" {
		t.Errorf("incomplete exact owned state: %+v", claim)
	}
	if f.Phase != "" && claim.Run != "retained-cli-run" {
		t.Error("resume changed run")
	}
	s, loaded, err := w.LoadApproval(f.Root, f.Slug)
	if err != nil || !reflect.DeepEqual(loaded, claim) {
		t.Fatal("approved source mismatch", err)
	}
	ctx, err := w.BuildContext(s, f.Slug)
	if err != nil {
		t.Fatal(err)
	}
	if ctx.Change.Phase != cliStartPhase(f) || ctx.Change.Run != claim.Run || ctx.Change.Path != f.Path || ctx.ChangeText != string(raw) || ctx.Module.Name != "alpha" || len(ctx.Dependencies) != 1 || ctx.Dependencies[0].Name != "beta" || ctx.PlaybookHash == "" || ctx.GitCommonDir == "" {
		t.Errorf("incomplete context %+v", ctx)
	}
	if f.Kind == "spec" && (len(ctx.Change.Added) != 1 || ctx.Change.Added[0].ID != "ALP-002" || strings.Contains(string(raw), "new:visible")) {
		t.Error("late IDs/bindings lost")
	}
	exists := true
	if _, e := os.Lstat(filepath.Join(f.Root, "units/a/internal/private.txt")); os.IsNotExist(e) {
		exists = false
	} else if e != nil {
		t.Fatal(e)
	}
	if exists == (f.Kind == "regenerate") {
		t.Error("wrong disposable deletion")
	}
	if strings.TrimSpace(gitFixtureCommand(t, f.Root, "rev-parse", "HEAD")) != f.Head {
		t.Error("start committed or moved HEAD")
	}
	return ctx
}
func cliStartDiagnostics(t *testing.T, root string, ds []pathDiagnostic, expected w.StartResult, from string) {
	t.Helper()
	want := []pathDiagnostic{}
	for _, d := range expected.Diagnostics {
		want = append(want, pathDiagnostic{d.Code, d.Severity, strings.ReplaceAll(d.Message, from, root), d.Path})
	}
	if !reflect.DeepEqual(ds, want) {
		t.Errorf("warnings differ got%+v want%+v", ds, want)
	}
}
func Test_APP022_APP023_APP024StartFixtureControls(t *testing.T) {
	for _, item := range []struct{ kind, phase string }{{"spec", ""}, {"fix", ""}, {"maintenance", ""}, {"regenerate", ""}, {"spec", "evals"}, {"spec", "audit"}, {"spec", "impl"}, {"regenerate", "impl"}} {
		t.Run(item.kind+"/"+item.phase, func(t *testing.T) {
			f := cliStartFixture(t, "sha1", item.kind, item.phase, false, true, false)
			before := cliStartSnapshot(t, f)
			r, e := w.Start(f.Where, "current")
			cliStartPreserved(t, f, before, e == nil)
			if e != nil {
				t.Fatal("public Start fixture", e)
			}
			ctx := cliStartOwned(t, f)
			if !reflect.DeepEqual(ctx, r.Context) {
				t.Fatal("public result not exact context")
			}
		})
	}
}
func Test_APP022_APP023StartExecutableStateAndOutput(t *testing.T) {
	for _, item := range []struct{ kind, phase string }{{"spec", ""}, {"fix", ""}, {"maintenance", ""}, {"regenerate", ""}, {"spec", "evals"}, {"spec", "audit"}, {"spec", "impl"}, {"regenerate", "impl"}} {
		for _, plain := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/%s/plain-%v", item.kind, item.phase, plain), func(t *testing.T) {
				f := cliStartFixture(t, "sha1", item.kind, item.phase, false, true, false)
				args := []string{"start", "current"}
				if !plain {
					args = append(args, "--json")
				}
				before := cliStartSnapshot(t, f)
				got := invoke(t, f.Where, args...)
				cliStartPreserved(t, f, before, got.code == 0)
				if got.code != 0 {
					t.Errorf("executable start failed: %+v", got)
					return
				}
				ctx := cliStartOwned(t, f)
				raw := json.RawMessage(got.out)
				if !plain {
					raw = decode(t, got, "start").Data
				} else {
					coveragePretty(t, got.out)
				}
				if !reflect.DeepEqual(decodeData(t, raw), canonical(t, ctx)) {
					t.Error("dropped/added context fields")
				}
				if item.kind == "regenerate" {
					if plain {
						if !strings.Contains(got.err, "D01") {
							t.Error("missing drill warning")
						}
					} else {
						ds := detailedDiagnostics(t, got)
						if len(ds) != 1 || ds[0].Code != "D01" || ds[0].Severity != "warning" {
							t.Error("wrong warning", ds)
						}
					}
				} else if got.err != "" || (!plain && len(detailedDiagnostics(t, got)) != 0) {
					t.Error("unexpected diagnostics", got.err)
				}
			})
		}
	}
}
func Test_APP022_APP023StartLinkedWarningsAndSchema(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		for _, position := range []int{-1, 0, 1, 2} {
			t.Run(fmt.Sprintf("%s/json-%d", format, position), func(t *testing.T) {
				control := cliStartFixture(t, format, "regenerate", "", false, false, true)
				expected, e := w.Start(control.Where, "current")
				if e != nil {
					t.Fatal(e)
				}
				if len(expected.Diagnostics) < 3 {
					t.Fatal("multi-warning premise", expected.Diagnostics)
				}
				f := cliStartFixture(t, format, "regenerate", "", true, true, true)
				args := []string{"start", "current"}
				if position >= 0 {
					args = append(args[:position], append([]string{"--json"}, args[position:]...)...)
				}
				before := cliStartSnapshot(t, f)
				got := invoke(t, f.Where, args...)
				cliStartPreserved(t, f, before, got.code == 0)
				if got.code != 0 {
					t.Errorf("linked start failed %+v", got)
					return
				}
				ctx := cliStartOwned(t, f)
				var v envelope
				if position >= 0 {
					v = decode(t, got, "start")
				} else {
					coveragePretty(t, got.out)
					v.Data = json.RawMessage(got.out)
				}
				if !reflect.DeepEqual(decodeData(t, v.Data), canonical(t, ctx)) {
					t.Error("linked full context mismatch")
				}
				if position >= 0 {
					cliStartDiagnostics(t, f.Root, detailedDiagnostics(t, got), expected, control.Root)
				} else {
					lines := strings.Split(strings.TrimSpace(got.err), "\n")
					if len(lines) != len(expected.Diagnostics) {
						t.Errorf("extra/missing plain warning records %q", got.err)
					}
					for i, d := range expected.Diagnostics {
						if i >= len(lines) {
							break
						}
						message := strings.ReplaceAll(d.Message, control.Root, f.Root)
						if !strings.Contains(lines[i], message) || !strings.Contains(lines[i], d.Code) || (d.Path != "" && !strings.Contains(lines[i], d.Path)) {
							t.Errorf("plain warning code/path/message/order lost: %s", got.err)
						}
					}

				}
			})
		}
	}
}
func cliStartError(t *testing.T, mode string) (cliStartCase, string, error) {
	t.Helper()
	f := cliStartFixture(t, "sha1", "spec", "", false, false, false)
	slug := "current"
	commit := func() {
		gitFixtureCommand(t, f.Root, "add", ".")
		gitFixtureCommand(t, f.Root, "commit", "-qm", "Error fixture")
		f.Head = strings.TrimSpace(gitFixtureCommand(t, f.Root, "rev-parse", "HEAD"))
		gitFixtureCommand(t, f.Root, "update-ref", "refs/heads/approval-base", f.Head)
	}
	switch mode {
	case "dirty-config":
		projectPut(t, f.Root, "fireweed.yaml", "malformed: [")
	case "P01":
		projectPut(t, f.Root, "untracked", "retain me")
	case "P03":
		raw, e := os.ReadFile(filepath.Join(f.Root, f.Path))
		if e != nil {
			t.Fatal(e)
		}
		projectPut(t, f.Root, cliStartPath, strings.Replace(string(raw), "module: alpha", "module: alpha\nblocked: unresolved question", 1))
		commit()
	case "C01":
		if e := os.Remove(filepath.Join(f.Root, "SYSTEM.md")); e != nil {
			t.Fatal(e)
		}
		projectPut(t, f.Root, "SYSTEM.md/keep", "unreadable system directory")
		commit()
	case "D05":
		projectPut(t, f.Root, ".fireweed/playbooks/evals.md/keep", "unreadable override directory")
		commit()
	case "P02":
		peer := filepath.Join(t.TempDir(), "foreign owner")
		gitFixtureCommand(t, f.Root, "worktree", "add", "--detach", peer)
		if _, e := w.CreateClaim(peer, "current", "held", f.Head, ""); e != nil {
			t.Fatal(e)
		}
	case "uncoded":
		projectPut(t, f.Root, "fireweed.yaml", "invalid: [")
		commit()
	case "missing-selector":
		slug = "not-present"
	case "C05":
		slug = "../invalid"
	default:
		t.Fatal("unknown error mode")
	}
	before := cliStartSnapshot(t, f)
	got, err := w.Start(f.Where, slug)
	if err == nil || !reflect.DeepEqual(got, w.StartResult{}) {
		t.Fatal("public rejection premise", got, err)
	}

	f.FailedObjects = cliStartNewObjectCount(t, before)
	cliStartPreserved(t, f, before, false)
	return f, slug, err
}
func Test_APP023StartErrorFixtureControls(t *testing.T) {
	for _, mode := range []string{"P01", "P02", "P03", "C05", "C01", "D05", "uncoded", "dirty-config", "missing-selector"} {
		t.Run(mode, func(t *testing.T) {
			f, slug, expected := cliStartError(t, mode)
			before := cliStartSnapshot(t, f)
			got, err := w.Start(f.Where, slug)
			if err == nil || err.Error() != expected.Error() || !reflect.DeepEqual(got, w.StartResult{}) {
				t.Fatalf("repeated public rejection premise: %+v %v", got, err)
			}
			cliStartPreserved(t, f, before, false)
		})
	}
}
func Test_APP023StartOperationalErrors(t *testing.T) {
	for _, mode := range []string{"P01", "P02", "P03", "C05", "C01", "D05", "uncoded", "dirty-config", "missing-selector"} {
		for _, plain := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/plain-%v", mode, plain), func(t *testing.T) {
				f, slug, expected := cliStartError(t, mode)
				before := cliStartSnapshot(t, f)
				args := []string{"start", slug}
				if !plain {
					args = append(args, "--json")
				}
				got := invoke(t, f.Where, args...)
				diag, code := cliStartErrorCode(expected)
				if got.code != code {
					t.Errorf("wrong exit %+v want%d", got, code)
				}
				if plain {
					if got.out != "" || !strings.Contains(got.err, diag) || !strings.Contains(got.err, expected.Error()) {
						t.Errorf("wrong plain error %+v", got)
					}
				} else {
					v := decode(t, got, "start")
					ds := detailedDiagnostics(t, got)
					if string(v.Data) != "null" || len(ds) != 1 || ds[0].Code != diag || ds[0].Severity != "error" || ds[0].Message != expected.Error() {
						t.Errorf("wrong operational envelope %+v", got)
					}
				}
				cliStartPreserved(t, f, before, false)
			})
		}
	}
}
func Test_APP024StartArguments(t *testing.T) {
	root := t.TempDir()
	for _, args := range [][]string{{"start"}, {"start", "a", "b"}, {"start", "current", "--bad"}, {"start", "current", "--json=true"}, {"start", "current", "--help=true"}, {"start", "current", "--with-change", "other"}, {"start", "current", "--history"}, {"start", "current", "--projected"}, {"start", "current", "--with-change=other"}, {"start", "current", "--", "--json"}} {
		t.Run(strings.Join(args, "_"), func(t *testing.T) { got := invoke(t, root, args...); assertUsage(t, got) })
	}
	for _, args := range [][]string{{"start", "--help"}, {"--json", "start", "-h"}} {
		got := invoke(t, root, args...)
		if got.code != 0 || got.err != "" {
			t.Errorf("start help %+v", got)
		}
	}
	for _, args := range [][]string{{"start", "../invalid", "--json"}, {"--json", "start", "--", "-bad"}} {
		got := invoke(t, root, args...)
		v := decode(t, got, "start")
		if got.code != 1 || string(v.Data) != "null" || len(v.Diagnostics) != 1 || v.Diagnostics[0].Code != "C05" {
			t.Errorf("invalid slug did not retain C05 %+v", got)
		}
	}
	if entries, e := os.ReadDir(root); e != nil || len(entries) != 0 {
		t.Error("arguments created files", e)
	}
}
func Test_APP023StartOutputFailuresRetainOwnership(t *testing.T) {
	for _, mode := range []string{"json", "plain", "warning"} {
		t.Run(mode, func(t *testing.T) {
			warning := mode == "warning"
			f := cliStartFixture(t, "sha1", "spec", "", false, false, warning)
			cwd, e := os.Getwd()
			if e != nil {
				t.Fatal(e)
			}
			if e = os.Chdir(f.Root); e != nil {
				t.Fatal(e)
			}
			defer os.Chdir(cwd)
			out, errOut := &bytes.Buffer{}, &bytes.Buffer{}
			before := cliStartSnapshot(t, f)
			env := os.Environ()
			var code int
			if warning {
				code = app.Run([]string{"start", "current"}, out, failedWriter{})
			} else if mode == "plain" {
				code = app.Run([]string{"start", "current"}, failedWriter{}, errOut)
			} else {
				code = app.Run([]string{"start", "current", "--json"}, failedWriter{}, errOut)
			}
			if code != 2 {
				t.Error("writer failure exit", code)
			}
			cliStartOwned(t, f)
			cliStartPreserved(t, f, before, true)
			after, e := os.Getwd()
			if e != nil || after != f.Root || !reflect.DeepEqual(env, os.Environ()) {
				t.Error("CLI changed parent cwd/environment")
			}
		})
	}
}

// A file snapshot is fixture evidence; subjects continue to use public APIs.
func queueFiles(t *testing.T, root string) map[string]string {
	t.Helper()
	out := map[string]string{}
	err := filepath.WalkDir(root, func(full string, d os.DirEntry, e error) error {
		if e != nil {
			return e
		}
		rel, e := filepath.Rel(root, full)
		if e != nil {
			return e
		}
		if rel == ".git" {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		info, e := os.Lstat(full)
		if e != nil {
			return e
		}
		v := info.Mode().String()
		if info.Mode()&os.ModeSymlink != 0 {
			target, e := os.Readlink(full)
			if e != nil {
				return e
			}
			v += " link:" + target
		} else if info.Mode().IsRegular() {
			b, e := os.ReadFile(full)
			if e != nil {
				return e
			}
			v += " bytes:" + string(b)
		}
		out[filepath.ToSlash(rel)] = v
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func Test_APP022_APP023StartWarningFixtureControls(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/linked-%v", format, linked), func(t *testing.T) {
				f := cliStartFixture(t, format, "regenerate", "", linked, true, true)
				before := cliStartSnapshot(t, f)
				r, e := w.Start(f.Where, "current")
				cliStartPreserved(t, f, before, e == nil)
				if e != nil {
					t.Fatal(e)
				}
				cliStartOwned(t, f)
				if len(r.Diagnostics) < 3 {
					t.Fatal("multi-warning fixture did not exercise independent diagnostics", r.Diagnostics)
				}
			})
		}
	}
}

// Snapshots independently constrain the complete checkout, index, common Git
// metadata and a private temp directory; only the public transaction's files and
// new immutable objects / exact selected claim are permitted to differ.
type cliStartState struct {
	Files, Git, Temp map[string]string
	Common, Scratch  string
}

func cliStartSnapshot(t *testing.T, f cliStartCase) cliStartState {
	t.Helper()
	scratch, e := filepath.EvalSymlinks(t.TempDir())
	if e != nil {
		t.Fatal(e)
	}
	t.Setenv("TMPDIR", scratch)
	t.Setenv("TMP", scratch)
	t.Setenv("TEMP", scratch)
	if runtime.GOOS != "windows" {
		traps := t.TempDir()
		marker := filepath.Join(scratch, "forbidden-command")
		t.Setenv("FIREWEED_START_TRAP_MARKER", marker)
		for _, name := range []string{"agent-trap", "fireweed", "fireweed-bootstrap", "codex", "curl", "wget"} {
			path := filepath.Join(traps, name)
			if e := os.WriteFile(path, []byte("#!/bin/sh\nprintf denied > \"$FIREWEED_START_TRAP_MARKER\"\nexit 91\n"), 0755); e != nil {
				t.Fatal(e)
			}
			cmd := exec.Command(path)
			if e := cmd.Run(); e == nil {
				t.Fatal("command trap ineffective")
			}
			if raw, e := os.ReadFile(marker); e != nil || string(raw) != "denied" {
				t.Fatalf("unobservable command trap %s %v", name, e)
			}
			if e := os.Remove(marker); e != nil {
				t.Fatal(e)
			}
		}
		t.Setenv("PATH", traps+string(os.PathListSeparator)+os.Getenv("PATH"))
	}
	proxy := cliStartLocalGit(t, filepath.Join(scratch, "forbidden-git"))
	t.Setenv("PATH", proxy+string(os.PathListSeparator)+os.Getenv("PATH"))
	// Apple's Git launcher may initialize its own cache on first use of TMPDIR.
	gitFixtureCommand(t, f.Root, "--version")
	common := strings.TrimSpace(gitFixtureCommand(t, f.Root, "rev-parse", "--git-common-dir"))
	if !filepath.IsAbs(common) {
		common = filepath.Join(f.Root, common)
	}
	return cliStartState{queueFiles(t, f.Root), queueFiles(t, common), queueFiles(t, scratch), common, scratch}
}
func cliStartPreserved(t *testing.T, f cliStartCase, b cliStartState, started bool) {
	t.Helper()
	after := queueFiles(t, f.Root)
	want := map[string]string{}
	for k, v := range b.Files {
		want[k] = v
	}
	if started {
		if v, ok := after[f.Path]; ok {
			if strings.Split(v, " bytes:")[0] != strings.Split(want[f.Path], " bytes:")[0] {
				t.Error("change file mode changed")
			}
			want[f.Path] = v
		}
		if f.Kind == "regenerate" {
			delete(want, "units/a/internal/private.txt")
		}
	}
	if !reflect.DeepEqual(want, after) {
		t.Errorf("CLI added unexpected checkout changes: before%v after%v", want, after)
	}
	after = queueFiles(t, b.Common)
	allowed := func(k string) bool {
		return started && (k == "refs/fireweed" || k == "refs/fireweed/claims" || k == "refs/fireweed/claims/"+f.Slug || k == "logs/refs/fireweed" || k == "logs/refs/fireweed/claims" || k == "logs/refs/fireweed/claims/"+f.Slug)
	}
	for k, v := range b.Git {
		if !(started && (k == "refs/fireweed/claims/"+f.Slug || k == "logs/refs/fireweed/claims/"+f.Slug)) && after[k] != v {
			t.Errorf("Git metadata/object changed %s", k)
		}
	}
	newObjects := 0
	failedRun := ""
	for k, v := range after {
		if _, ok := b.Git[k]; ok || allowed(k) {
			continue
		}
		if (started || f.FailedObjects > 0) && strings.HasPrefix(k, "objects/") {
			rel := strings.TrimPrefix(k, "objects/")
			parts := strings.Split(rel, "/")
			if len(parts) == 1 && len(parts[0]) == 2 && strings.HasPrefix(v, "d") {
				populated := false
				for child := range after {
					if strings.HasPrefix(child, k+"/") {
						populated = true
						break
					}
				}
				if !populated {
					t.Errorf("empty leaked object directory %s", k)
				}
				continue
			}
			if len(parts) == 2 && len(parts[0]) == 2 && (len(parts[1]) == 38 || len(parts[1]) == 62) && !strings.Contains(v, " link:") {
				newObjects++
				oid := parts[0] + parts[1]
				if strings.TrimSpace(gitFixtureCommand(t, f.Root, "cat-file", "-t", oid)) != "blob" {
					t.Errorf("extra non-claim object %s", k)
					continue
				}
				raw := gitFixtureCommand(t, f.Root, "cat-file", "blob", oid)
				var c w.Claim
				if json.Unmarshal([]byte(raw), &c) != nil || c.Worktree != f.Root || c.Change != f.Slug || c.Base != f.Head || c.Run == "" {
					t.Errorf("extra non-transaction object %s", k)
					continue
				}
				if started {
					claim, e := w.LoadClaim(f.Root, f.Slug)
					// Reservation State is opaque; exact final State is asserted separately
					// by cliStartOwned. Every added blob must still be for this same claim/run.
					if e != nil || c.Run != claim.Run {
						t.Errorf("extra unrelated claim object %s", k)
					}
				} else {
					if failedRun == "" {
						failedRun = c.Run
					}
					if c.Run != failedRun || (f.Phase != "" && c.Run != "retained-cli-run") {
						t.Errorf("failed call added unrelated run object %s", k)
					}
				}
				continue
			}
		}
		t.Errorf("unexpected Git metadata %s", k)
	}

	if !started && newObjects > f.FailedObjects {
		t.Errorf("failed CLI created %d objects beyond public transaction allowance %d", newObjects, f.FailedObjects)
	}
	if !reflect.DeepEqual(b.Temp, queueFiles(t, b.Scratch)) {
		t.Error("CLI leaked or mutated temporary content")
	}
}
func Test_APP023StartRejectedOutputFailures(t *testing.T) {
	root := t.TempDir()
	cwd, e := os.Getwd()
	if e != nil {
		t.Fatal(e)
	}
	if e = os.Chdir(root); e != nil {
		t.Fatal(e)
	}
	defer os.Chdir(cwd)
	for _, args := range [][]string{{"start"}, {"start", "../invalid"}} {
		if code := app.Run(args, io.Discard, failedWriter{}); code != 2 {
			t.Errorf("plain failed diagnostic writer exit%d", code)
		}
		args = append(args, "--json")
		if code := app.Run(args, failedWriter{}, io.Discard); code != 2 {
			t.Errorf("JSON failed diagnostic writer exit%d", code)
		}
	}
	if code := app.Run([]string{"start", "--help"}, failedWriter{}, io.Discard); code != 2 {
		t.Errorf("help writer exit%d", code)
	}
}

func cliStartErrorCode(err error) (string, int) {
	for _, code := range []string{"P01", "P02", "P03", "C05"} {
		if strings.HasPrefix(err.Error(), code+":") {
			return code, 1
		}
	}
	for _, code := range []string{"C01", "D05"} {
		if strings.HasPrefix(err.Error(), code+":") {
			return code, 2
		}
	}
	return "C01", 2
}
func cliStartRenamed(t *testing.T, f cliStartCase, slug string) cliStartCase {
	t.Helper()
	next := "changes/0001-" + slug + ".md"
	if e := os.Rename(filepath.Join(f.Root, f.Path), filepath.Join(f.Root, next)); e != nil {
		t.Fatal(e)
	}
	gitFixtureCommand(t, f.Root, "add", ".")
	gitFixtureCommand(t, f.Root, "commit", "-qm", "Non-default selector")
	f.Head = strings.TrimSpace(gitFixtureCommand(t, f.Root, "rev-parse", "HEAD"))
	gitFixtureCommand(t, f.Root, "update-ref", "refs/heads/approval-base", f.Head)
	f.Slug = slug
	f.Path = next
	return f
}
func Test_APP022_APP024StartExactOperand(t *testing.T) {
	for _, control := range []bool{true, false} {
		t.Run(fmt.Sprint(control), func(t *testing.T) {
			f := cliStartRenamed(t, cliStartFixture(t, "sha1", "spec", "", true, true, false), "different-change")
			before := cliStartSnapshot(t, f)
			if control {
				if _, e := w.Start(f.Where, f.Slug); e != nil {
					t.Fatal("selector fixture", e)
				}
			} else {
				got := invoke(t, f.Where, "--json", "start", "--", f.Slug)
				if got.code != 0 {
					t.Fatalf("exact selector failed %+v", got)
				}
				v := decode(t, got, "start")
				if !reflect.DeepEqual(decodeData(t, v.Data), canonical(t, cliStartOwned(t, f))) {
					t.Error("selector data mismatch")
				}
			}
			cliStartPreserved(t, f, before, true)
			cliStartOwned(t, f)
		})
	}
}

// Record only concrete new loose objects created by the independent public
// failure control. Its full storage/file/ref validation happens immediately;
// later rejected CLI calls may reuse immutable objects but cannot add more.
func cliStartNewObjectCount(t *testing.T, b cliStartState) int {
	t.Helper()
	after := queueFiles(t, b.Common)
	count := 0
	for k, v := range after {
		if _, ok := b.Git[k]; ok {
			continue
		}
		parts := strings.Split(strings.TrimPrefix(k, "objects/"), "/")
		if strings.HasPrefix(k, "objects/") && len(parts) == 2 && len(parts[0]) == 2 && (len(parts[1]) == 38 || len(parts[1]) == 62) && !strings.HasPrefix(v, "d") {
			count++
		}
	}
	return count
}
