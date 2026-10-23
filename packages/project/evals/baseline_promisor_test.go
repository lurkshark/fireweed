package evals_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	p "github.com/lurkshark/fireweed/packages/project"
)

// This helper is a local transport endpoint: it writes a marker and exits.
// No network service or bootstrap executable participates in these fixtures.
func baselinePromisorHelper(t *testing.T) {
	t.Helper()
	dir := baselineCanonical(t, t.TempDir())
	source := filepath.Join(dir, "transport.go")
	program := `package main
import "os"
func main(){f,e:=os.OpenFile(os.Getenv("FW_BASELINE_PROMISOR_MARKER"),os.O_CREATE|os.O_WRONLY|os.O_APPEND,0600);if e!=nil{os.Exit(92)};f.WriteString("attempted lazy transport\n");f.Close();os.Exit(91)}
`
	if err := os.WriteFile(source, []byte(program), 0600); err != nil {
		t.Fatal(err)
	}
	name := "git-remote-baselinefetch"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "build", "-o", filepath.Join(dir, name), source)
	cmd.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("local helper build: %v %s", err, out)
	}
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
	// Subjects must impose their own local-read policy, independent of this value.
	t.Setenv("GIT_NO_LAZY_FETCH", "0")
	t.Setenv("FW_BASELINE_TRACE_DIR", dir)
	temporary := baselineCanonical(t, t.TempDir())
	t.Setenv("TMPDIR", temporary)
	t.Setenv("TMP", temporary)
	t.Setenv("TEMP", temporary)
	t.Setenv("FW_BASELINE_OWNED_TEMP", temporary)
}

type baselinePromisorCase struct {
	Fixture                                      baselineFixtureData
	Common, Object, Marker, Revision, Path, Code string
	RealGit, Trace                               string
	Payload                                      []byte
	Fatal, Success                               bool
}

func baselinePromisorFixture(t *testing.T, format, kind string, linked bool) baselinePromisorCase {
	t.Helper()
	f := baselineFixture(t, format)
	if kind == "link-entry" || kind == "link-target" {
		put(t, f.Root, "selected-config.yaml", f.Files["fireweed.yaml"])
		// Construct a committed link directly so this premise needs no OS symlink privilege.
		baselineGit(t, f.Root, "add", "selected-config.yaml")
		raw, err := baselineInput(t, f.Root, "selected-config.yaml", "hash-object", "-w", "--stdin")
		if err != nil {
			t.Fatal(err, raw)
		}
		baselineGit(t, f.Root, "update-index", "--cacheinfo", "120000,"+strings.TrimSpace(raw)+",fireweed.yaml")
		baselineGit(t, f.Root, "commit", "-m", "selected committed configuration link")
		f.Commit = baselineGit(t, f.Root, "rev-parse", "HEAD")
	}
	revision := f.Commit
	target := "fireweed.yaml"
	code := "C01"
	fatal := true
	success := false
	switch kind {
	case "module":
		target = "units/a/SPEC.md"
		code = "C03"
		fatal = false
	case "queue":
		target = "changes/0002-zebra.md"
		code = "C05"
		fatal = false
	case "archive":
		target = "changes/archive/0009-old.md"
		code = "C05"
		fatal = false
	case "link-target":
		target = "selected-config.yaml"
	case "private":
		target = "private/secret.txt"
		success = true
		fatal = false
	case "disposable":
		target = "units/z/internal/old.go"
		success = true
		fatal = false
	case "hidden":
		target = ".fireweed/hidden/SPEC.md"
		success = true
		fatal = false
	}
	oid := ""
	switch kind {
	case "commit", "ancestor":
		oid = f.Commit
	case "tag", "tag-target":
		baselineGit(t, f.Root, "tag", "-a", "promisor-selected", "-m", "selected local tag", f.Commit)
		oid = baselineGit(t, f.Root, "rev-parse", "promisor-selected")
		if kind == "tag-target" {
			oid = f.Commit
		}
		revision = "promisor-selected"
	case "root-tree":
		oid = baselineGit(t, f.Root, "rev-parse", f.Commit+"^{tree}")
	case "nested-tree":
		oid = baselineGit(t, f.Root, "rev-parse", f.Commit+":units/z")
	default:
		oid = baselineGit(t, f.Root, "rev-parse", f.Commit+":"+target)
	}
	// Current files are deliberately different, so they cannot supply missing selected bytes.
	put(t, f.Root, "fireweed.yaml", "invalid current metadata: [")
	baselineCommit(t, f.Root)
	if kind == "ancestor" {
		revision = "HEAD~1"
	}
	if linked {
		peer := filepath.Join(t.TempDir(), "linked ' promisor")
		baselineGit(t, f.Root, "worktree", "add", "-b", "promisor-peer", peer)
		f.Root = baselineCanonical(t, peer)
		f.Expected.Root = f.Root
	}
	common := baselineGit(t, f.Root, "rev-parse", "--git-common-dir")
	if !filepath.IsAbs(common) {
		common = filepath.Join(f.Root, common)
	}
	common = baselineCanonical(t, common)
	for key, value := range map[string]string{"extensions.partialClone": "origin", "remote.origin.promisor": "true", "remote.origin.partialclonefilter": "blob:none", "remote.origin.url": "baselinefetch::marker-only", "protocol.baselinefetch.allow": "always"} {
		baselineGit(t, f.Root, "config", key, value)
	}
	marker := filepath.Join(baselineCanonical(t, t.TempDir()), "transport")
	t.Setenv("FW_BASELINE_PROMISOR_MARKER", marker)
	object := filepath.Join(common, "objects", oid[:2], oid[2:])
	payload, err := os.ReadFile(object)
	if err != nil {
		t.Fatal("loose object premise", err)
	}
	real, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	traceFile, err := os.CreateTemp(os.Getenv("FW_BASELINE_TRACE_DIR"), "trace-")
	if err != nil {
		t.Fatal(err)
	}
	trace := traceFile.Name()
	traceFile.Close()
	t.Setenv("GIT_TRACE2_EVENT", trace)
	return baselinePromisorCase{Fixture: f, Common: common, Object: object, Marker: marker, Revision: revision, Path: target, Code: code, Payload: payload, Fatal: fatal, Success: success, RealGit: real, Trace: trace}
}

// Trace2 records actual Git command identity even when protocol policy rejects a
// fetch before the executable marker transport can run.
func baselineFetchTrace(t *testing.T, path string) bool {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	for _, line := range strings.Split(strings.TrimSpace(string(raw)), "\n") {
		if line == "" {
			continue
		}
		var e struct {
			Event, Name string
			Argv        []string
		}
		if err := json.Unmarshal([]byte(line), &e); err != nil {
			t.Fatal("invalid independent Git trace", err)
		}
		if e.Event == "cmd_name" && (e.Name == "fetch" || e.Name == "fetch-pack") {
			return true
		}
		if e.Event == "child_start" && len(e.Argv) > 1 && (filepath.Base(e.Argv[0]) == "git" || filepath.Base(e.Argv[0]) == "git.exe") {
			for i := 1; i < len(e.Argv); i++ {
				a := e.Argv[i]
				if a == "-c" || a == "-C" || a == "--git-dir" || a == "--work-tree" || a == "--exec-path" {
					i++
					continue
				}
				if strings.HasPrefix(a, "-") {
					continue
				}
				if a == "fetch" || a == "fetch-pack" {
					return true
				}
				break
			}
		}
	}
	return false
}
func baselineResetTrace(t *testing.T, path string) {
	t.Helper()
	if err := os.WriteFile(path, nil, 0600); err != nil {
		t.Fatal(err)
	}
}
func baselinePromisorRemove(t *testing.T, f baselinePromisorCase) {
	t.Helper()
	if err := os.Remove(f.Object); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.MkdirAll(filepath.Dir(f.Object), 0755); os.WriteFile(f.Object, f.Payload, 0444) })
	oid := filepath.Base(filepath.Dir(f.Object)) + filepath.Base(f.Object)
	run := func(deny bool) {
		baselineResetTrace(t, f.Trace)
		args := []string{}
		if deny {
			args = append(args, "-c", "protocol.baselinefetch.allow=never")
		}
		args = append(args, "cat-file", "-p", oid)
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		cmd := exec.CommandContext(ctx, f.RealGit, args...)
		cmd.Dir = f.Fixture.Root
		out, err := cmd.CombinedOutput()
		if err == nil {
			t.Fatal("missing object unexpectedly readable", string(out))
		}
		if !baselineFetchTrace(t, f.Trace) {
			t.Fatal("missing-object positive control did not observe actual fetch command", string(out))
		}
		raw, e := os.ReadFile(f.Marker)
		if deny {
			if !os.IsNotExist(e) {
				t.Fatal("denied transport unexpectedly invoked marker", e, string(raw))
			}
		} else {
			if e != nil || !strings.Contains(string(raw), "attempted lazy transport\n") {
				t.Fatal("promisor transport positive control", string(raw), e)
			}
			if err := os.Remove(f.Marker); err != nil {
				t.Fatal(err)
			}
		}
	}
	run(false)
	run(true)
	baselineResetTrace(t, f.Trace)
}
func baselinePromisorResult(t *testing.T, f baselinePromisorCase) {
	t.Helper()
	got, err := baselineCall(t, f.Fixture.Root, f.Revision)
	if f.Success {
		if err != nil || !reflect.DeepEqual(got, f.Fixture.Expected) {
			t.Errorf("unrelated missing blob changed valid selected snapshot: %+v %v", got, err)
		}
	} else if f.Fatal {
		if err == nil || !reflect.DeepEqual(got, p.Snapshot{}) || !strings.Contains(err.Error(), "C01") || !baselineCorrective(err.Error()) {
			t.Errorf("missing selected object must fail atomically with corrective C01: %+v %v", got, err)
		}
	} else {
		if err != nil {
			t.Errorf("selected document error must remain diagnosed: %v", err)
		}
		found := false
		for _, d := range got.Diagnostics {
			if d.Path == f.Path && d.Code == f.Code && d.Severity == "error" && baselineCorrective(d.Message) {
				found = true
			}
		}
		if !found {
			t.Errorf("missing selected document lacks corrective %s at %s: %+v", f.Code, f.Path, got.Diagnostics)
		}
	}
	if baselineFetchTrace(t, f.Trace) {
		t.Error("LoadAt attempted a proven forbidden fetch command")
	}
	if raw, err := os.ReadFile(f.Marker); !os.IsNotExist(err) {
		t.Errorf("LoadAt invoked proven lazy transport: %q %v", raw, err)
	}
}
func Test_PRJ027PromisorFixtureControls(t *testing.T) {
	baselinePromisorHelper(t)
	for _, kind := range []string{"config", "module", "queue", "archive", "commit", "ancestor", "tag", "tag-target", "root-tree", "nested-tree", "link-entry", "link-target", "private", "disposable", "hidden"} {
		t.Run(kind, func(t *testing.T) {
			f := baselinePromisorFixture(t, "sha1", kind, false)
			// Available selected metadata remains loadable despite promisor configuration.
			baselineAssert(t, f.Fixture, f.Fixture.Root, f.Revision)
			baselinePromisorRemove(t, f)
			// A metadata-only inventory must not read missing objects or invoke transport.
			baselineGit(t, f.Fixture.Root, "cat-file", "--batch-all-objects", "--batch-check=%(objectname)")
			baselineNoMarker(t, f.Marker)
		})
	}
}
func Test_PRJ027BaselineMissingPromisorObjects(t *testing.T) {
	baselinePromisorHelper(t)
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/linked-%v", format, linked), func(t *testing.T) {
				for _, kind := range []string{"config", "module", "queue", "archive", "commit", "ancestor", "tag", "tag-target", "root-tree", "nested-tree", "link-entry", "link-target", "private", "disposable", "hidden"} {
					t.Run(kind, func(t *testing.T) {
						f := baselinePromisorFixture(t, format, kind, linked)
						baselinePromisorRemove(t, f)
						files := playbookTree(t, f.Fixture.Root)
						metadata := playbookTree(t, f.Common)
						temporary := os.Getenv("FW_BASELINE_OWNED_TEMP")
						temps := playbookTree(t, temporary)
						baselineResetTrace(t, f.Trace)
						env := os.Environ()
						cwd, err := os.Getwd()
						if err != nil {
							t.Fatal(err)
						}
						baselinePromisorResult(t, f)
						if !reflect.DeepEqual(files, playbookTree(t, f.Fixture.Root)) || !reflect.DeepEqual(metadata, playbookTree(t, f.Common)) {
							t.Error("LoadAt changed files, config, refs, index or common Git")
						}
						if !reflect.DeepEqual(temps, playbookTree(t, temporary)) {
							t.Error("LoadAt changed temporary storage")
						}
						after, err := os.Getwd()
						if err != nil || after != cwd || !reflect.DeepEqual(env, os.Environ()) {
							t.Error("LoadAt changed parent cwd/environment")
						}
					})
				}
			})
		}
	}
}

func Test_PRJ027PromisorFetchObserverControls(t *testing.T) {
	baselinePromisorHelper(t)
	f := baselinePromisorFixture(t, "sha1", "config", false)
	baselineGit(t, f.Fixture.Root, "update-ref", "refs/heads/fetch", f.Fixture.Commit)
	baselineResetTrace(t, f.Trace)
	baselineGit(t, f.Fixture.Root, "rev-parse", "--verify", "fetch")
	if baselineFetchTrace(t, f.Trace) {
		t.Fatal("observer mistook a literal ref argument for a fetch command")
	}
	baselinePromisorRemove(t, f)
	baselineNoMarker(t, f.Marker)
	if baselineFetchTrace(t, f.Trace) {
		t.Fatal("reset trace retained positive-control fetch")
	}
}

func Test_PRJ027PromisorGuardOperationControls(t *testing.T) {
	for _, layer := range []string{"body", "outer", "composed"} {
		t.Run(layer, func(t *testing.T) {
			f := baselineFixture(t, "sha1")
			baselineGit(t, f.Root, "branch", "fetch", f.Commit)
			markers := []string{}
			if layer != "outer" {
				m, _ := baselineGuard(t, f, "")
				markers = append(markers, m)
			}
			if layer != "body" {
				markers = append(markers, playbookTraps(t))
			}
			queries := [][]string{{"cat-file", "commit", f.Commit}, {"rev-parse", "--verify", "fetch"}, {"-c", "probe.value=fetch", "rev-parse", "fetch"}, {"branch"}, {"branch", "--list"}, {"branch", "--list", "fetch"}, {"branch", "fetch", "--list"}, {"branch", "--format", "fetch", "--list"}, {"hash-object", "--stdin"}, {"branch", "--show-current"}, {"branch", "--list", "--abbrev=12"}, {"branch", "-i", "--list", "fetch"}, {"branch", "-l"}, {"branch", "--sort", "refname"}, {"branch", "--format", "%(refname)"}, {"worktree", "list"}, {"worktree", "list", "--porcelain"}}
			for _, args := range queries {
				if out, err := baselineInput(t, f.Root, "", args...); err != nil {
					t.Fatalf("read-only operation rejected %q: %v %s", args, err, out)
				}
				for _, m := range markers {
					baselineNoMarker(t, m)
				}
			}
			for _, args := range [][]string{{"hash-object", "-w", "--stdin"}, {"branch", "must-not-create"}, {"branch", "-D", "fetch"}, {"worktree", "add", filepath.Join(t.TempDir(), "must-not-create")}, {"fetch", "marker-only.invalid"}} {
				if out, err := baselineInput(t, f.Root, "", args...); err == nil {
					t.Fatalf("forbidden operation accepted %q %s", args, out)
				}
				found := false
				for _, m := range markers {
					if _, e := os.Stat(m); e == nil {
						found = true
						os.Remove(m)
					} else if !os.IsNotExist(e) {
						t.Fatal(e)
					}
				}
				if !found {
					t.Fatalf("forbidden-operation trap did not observe %q", args)
				}
			}
		})
	}
}
