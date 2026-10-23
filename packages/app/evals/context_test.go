package evals

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	app "github.com/lurkshark/fireweed/packages/app"
	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

const contextChangePath = "changes/0002-current.md"

func cliContextFixture(t *testing.T, format string) string {
	t.Helper()
	root, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	gitFixtureCommand(t, root, "init", "--initial-branch=main", "--object-format="+format)
	gitFixtureCommand(t, root, "config", "user.name", "Context Fixture")
	gitFixtureCommand(t, root, "config", "user.email", "context@example.invalid")
	projectPut(t, root, "fireweed.yaml", projectConfig)
	projectPut(t, root, ".gitignore", "ignored/**\n")
	projectPut(t, root, "SYSTEM.md", "Worker system bytes.\r\n")
	projectModule(t, root, "units/a", "alpha", "ALP", "depends_on: [beta]\npreserve: [retained/**]\nappend_only: [events/**]\n")
	projectModule(t, root, "units/b", "beta", "BET", "depends_on: [gamma]\n")
	projectModule(t, root, "units/g", "gamma", "GAM", "")
	for _, item := range []struct{ path, text string }{
		{"units/a/api.go", "Old unselected target interface.\n"},
		{"units/a/public/new.txt", "Projected target interface.\r\n"},
		{"units/a/evals/case.txt", "Target eval input.\n"},
		{"units/a/retained/data.txt", "Preserved target input.\n"},
		{"units/a/events/log.txt", "Append-only target input.\n"},
		{"units/a/internal/private.txt", "Disposable secret.\n"},
		{"units/a/nested/SPEC.md", "Nested spec excluded from file inputs.\n"},
		{"units/b/api.go", "Old dependency interface.\n"},
		{"units/b/public/new.txt", "Projected direct dependency interface.\n"},
		{"units/g/api.go", "Transitive private contract.\n"},
		{"private/secret.txt", "Unrelated private bytes.\n"},
		{"ignored/secret.txt", "Ignored bytes.\n"},
	} {
		projectPut(t, root, item.path, item.text)
	}
	cliQueueChange(t, root, "changes/0001-before.md", "beta", "BET", "set: {interface: ['public/**']}\n", "## Added\n\n### BET-002 [review]\nThe dependency MUST expose this preceding addition.\n")
	cliQueueChange(t, root, contextChangePath, "alpha", "ALP", "set: {interface: ['public/**']}\n", "## Added\n\n### ALP-002 [review]\nThe target MUST expose this selected addition.\n\n## Decisions\n\nSelected reason.\n")
	cliQueueChange(t, root, "changes/0003-later.md", "alpha", "ALP", "", "## Added\n\n### ALP-003 [review]\nLater behavior MUST remain later.\n")
	projectPut(t, root, "changes/archive/0009-archived.md", "---\nmodule: alpha\nphase: applied\n---\n\n## Motivation\n\nHistorical selector.\n")
	gitFixtureCommand(t, root, "add", ".")
	gitFixtureCommand(t, root, "commit", "-qm", "Context source baseline")
	// Discovery warnings are legitimate feedback, not a static approval gate.
	s, err := p.Load(root)
	if err != nil || len(s.Modules) != 3 || len(s.Changes) != 3 {
		t.Fatalf("fixture source: %v %+v", err, s)
	}
	return root
}

func cliContextClaim(t *testing.T, root, phase string, regenerate bool) w.Claim {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(root, contextChangePath))
	if err != nil {
		t.Fatal(err)
	}
	doc, err := c.ParseMarkdown(string(raw))
	if err != nil {
		t.Fatal(err)
	}
	doc.Metadata["phase"], doc.Metadata["run"] = phase, "context-run"
	if regenerate {
		doc.Metadata["regenerate"] = true
	}
	text, err := c.RenderMarkdown(doc)
	if err != nil {
		t.Fatal(err)
	}
	projectPut(t, root, contextChangePath, text)
	base := strings.TrimSpace(gitFixtureCommand(t, root, "rev-parse", "HEAD"))
	claim, err := w.CreateClaim(root, "current", "context-run", base, text)
	if err != nil {
		t.Fatal(err)
	}
	_, loaded, err := w.LoadApproval(root, "current")
	if err != nil || !reflect.DeepEqual(loaded, claim) {
		t.Fatalf("owned source control: %v %+v", err, loaded)
	}
	return claim
}

func cliContextExpected(t *testing.T, root string) w.WorkerContext {
	t.Helper()
	s, _, err := w.LoadApproval(root, "current")
	if err != nil {
		t.Fatal(err)
	}
	want, err := w.BuildContext(s, "current")
	if err != nil {
		t.Fatal(err)
	}
	if want.Change.Slug != "current" || len(want.Module.Specs) != 2 || want.Module.Specs[1].ID != "ALP-002" || len(want.Dependencies) != 1 || len(want.Dependencies[0].Specs) != 2 || want.Dependencies[0].Specs[1].ID != "BET-002" {
		t.Fatalf("projection control: %+v", want)
	}
	wantFiles := []string{"units/a/evals/case.txt", "units/a/public/new.txt"}
	if want.Change.Phase != "audit" {
		wantFiles = []string{"units/a/evals/case.txt", "units/a/events/log.txt", "units/a/public/new.txt", "units/a/retained/data.txt", "units/b/public/new.txt"}
	}
	paths := []string{}
	for _, f := range want.Files {
		paths = append(paths, f.Path)
	}
	if !reflect.DeepEqual(paths, wantFiles) {
		t.Fatalf("selection control %v want %v", paths, wantFiles)
	}
	if want.PlaybookHash == "" || want.GitCommonDir == "" {
		t.Fatal("incomplete public context control")
	}
	return want
}

func cliContextAssert(t *testing.T, where string, want w.WorkerContext) {
	t.Helper()
	for _, args := range [][]string{{"context", "current", "--json"}, {"--json", "context", "current"}, {"context", "--json", "current"}, {"context", "current"}} {
		got := invoke(t, where, args...)
		if got.code != 0 || got.err != "" {
			t.Errorf("context %v: %+v", args, got)
			continue
		}
		raw := json.RawMessage(got.out)
		if len(args) > 2 {
			raw = successJSON(t, got, "context").Data
		} else {
			coveragePretty(t, got.out)
		}
		if !reflect.DeepEqual(decodeData(t, raw), canonical(t, want)) {
			t.Errorf("complete context mismatch %s", raw)
		}
	}
}

func Test_APP019ContextSourceComposition(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		for _, phase := range []string{"", "evals", "audit", "impl", "regenerate"} {
			t.Run(format+"/"+phase, func(t *testing.T) {
				root := cliContextFixture(t, format)
				if phase != "" {
					actual := phase
					if phase == "regenerate" {
						actual = "impl"
					}
					cliContextClaim(t, root, actual, phase == "regenerate")
				}
				book := "author-change"
				if phase != "" {
					book = phase
				}
				projectPut(t, root, ".fireweed/playbooks/"+book+".md", "---\r\nbased_on: stale\r\n---\r\nContext override.\r\n")
				projectPut(t, root, "units/a/nested/.keep", "")
				want := cliContextExpected(t, root)
				if want.Playbook != "---\r\nbased_on: stale\r\n---\r\nContext override.\r\n" {
					t.Fatal("override control")
				}
				cliContextAssert(t, filepath.Join(root, "units/a/nested"), want)
			})
		}
	}
}

func Test_APP019ContextApprovedMetadataAndFeedback(t *testing.T) {
	for _, mode := range []string{"invalid-current", "deleted-config", "deleted-change", "empty-change", "moved-head", "linked"} {
		t.Run(mode, func(t *testing.T) {
			root := cliContextFixture(t, "sha1")
			if mode == "linked" {
				linked := filepath.Join(t.TempDir(), "linked checkout")
				gitFixtureCommand(t, root, "worktree", "add", "--detach", linked, "HEAD")
				var err error
				root, err = filepath.EvalSymlinks(linked)
				if err != nil {
					t.Fatal(err)
				}
			}
			claim := cliContextClaim(t, root, "impl", false)
			source, _, err := w.LoadApproval(root, "current")
			if err != nil {
				t.Fatal(err)
			}
			switch mode {
			case "invalid-current", "deleted-config":
				projectPut(t, root, "fireweed.yaml", "invalid: [\n")
				if mode == "deleted-config" {
					if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
						t.Fatal(err)
					}
				}
				projectPut(t, root, "units/a/SPEC.md", "Worker-invalid SPEC bytes.\r\n")
				projectPut(t, root, "changes/0001-before.md", "Broken current predecessor.")
				projectPut(t, root, contextChangePath, "Worker-invalid change bytes.\r\n")
			case "deleted-change":
				if err := os.Remove(filepath.Join(root, contextChangePath)); err != nil {
					t.Fatal(err)
				}
			case "empty-change":
				projectPut(t, root, contextChangePath, "")
			case "moved-head":
				gitFixtureCommand(t, root, "commit", "--allow-empty", "-qm", "Moved HEAD feedback")
			}
			want, err := w.BuildContext(source, "current")
			if err != nil {
				t.Fatal(err)
			}
			if want.Change.Run != claim.Run || want.Change.Phase != "impl" || want.Change.Path != contextChangePath {
				t.Fatal("lost approved state control")
			}
			if mode == "deleted-change" && want.ChangeText != claim.State {
				t.Fatal("missing recovery control")
			}
			if mode == "empty-change" && want.ChangeText != "" {
				t.Fatal("empty feedback control")
			}
			if (mode == "invalid-current" || mode == "deleted-config") && (want.SpecText != "Worker-invalid SPEC bytes.\r\n" || want.ChangeText != "Worker-invalid change bytes.\r\n") {
				t.Fatal("feedback bytes control")
			}
			cliContextAssert(t, root, want)
		})
	}
}

func cliContextError(t *testing.T, where, slug string, want error, code int) {
	t.Helper()
	if want == nil {
		t.Fatal("missing public error premise")
	}
	text := want.Error()
	prefix, _, ok := strings.Cut(text, ":")
	if !ok {
		t.Fatalf("uncoded public error %v", want)
	}
	for _, jsonMode := range []bool{true, false} {
		args := []string{"context", slug}
		if jsonMode {
			args = append(args, "--json")
		}
		got := invoke(t, where, args...)
		if got.code != code {
			t.Errorf("operational exit want%d got%+v", code, got)
			continue
		}
		if jsonMode {
			x := decode(t, got, "context")
			ds := detailedDiagnostics(t, got)
			if string(x.Data) != "null" || len(ds) != 1 || ds[0].Code != prefix || ds[0].Severity != "error" || ds[0].Message != text || ds[0].Path != "" {
				t.Errorf("public error lost %v %+v", want, got)
			}
			var raw struct{ Diagnostics []map[string]any }
			if err := json.Unmarshal([]byte(got.out), &raw); err != nil {
				t.Fatal(err)
			}
			if len(raw.Diagnostics) == 1 {
				if _, exists := raw.Diagnostics[0]["path"]; exists {
					t.Error("invented error path")
				}
			}
		} else if got.out != "" || got.err != "error "+prefix+": "+text+"\n" {
			t.Errorf("plain error lost %+v", got)
		}
	}
}

func cliContextPublicError(root, slug string) error {
	s, _, err := w.LoadApproval(root, slug)
	if err == nil {
		_, err = w.BuildContext(s, slug)
	}
	return err
}

func Test_APP020ContextOperationalErrors(t *testing.T) {
	for _, mode := range []string{"missing", "prefix", "archived", "ambiguous", "missing-target", "malformed-claim", "foreign", "reservation", "baseline", "state", "projection", "missing-spec", "required-directory", "missing-change", "missing-config", "bad-config", "override-directory", "escape"} {
		t.Run(mode, func(t *testing.T) {
			root := cliContextFixture(t, "sha1")
			slug := "current"
			code := 1
			base := strings.TrimSpace(gitFixtureCommand(t, root, "rev-parse", "HEAD"))
			switch mode {
			case "missing":
				slug = "absent"
			case "prefix":
				slug = "cur"
			case "archived":
				slug = "archived"
			case "ambiguous":
				raw, e := os.ReadFile(filepath.Join(root, contextChangePath))
				if e != nil {
					t.Fatal(e)
				}
				projectPut(t, root, "changes/0004-current.md", string(raw))
			case "missing-target":
				raw, e := os.ReadFile(filepath.Join(root, contextChangePath))
				if e != nil {
					t.Fatal(e)
				}
				projectPut(t, root, contextChangePath, strings.Replace(string(raw), "module: alpha", "module: absent", 1))
			case "malformed-claim":
				projectPut(t, root, "payload", "not JSON")
				oid := strings.TrimSpace(gitFixtureCommand(t, root, "hash-object", "-w", "payload"))
				gitFixtureCommand(t, root, "update-ref", "refs/fireweed/claims/current", oid)
			case "foreign", "reservation", "baseline", "state":
				owner := root
				state := ""
				revision := base
				if mode != "reservation" {
					raw, e := os.ReadFile(filepath.Join(root, contextChangePath))
					if e != nil {
						t.Fatal(e)
					}
					state = string(raw)
				}
				if mode == "foreign" {
					owner = filepath.Join(t.TempDir(), "foreign")
				}
				if mode == "baseline" {
					revision = strings.Repeat("1", 40)
				}
				payload, e := json.Marshal(w.Claim{Worktree: owner, Change: "current", Run: "context-run", Base: revision, State: state})
				if e != nil {
					t.Fatal(e)
				}
				projectPut(t, root, "payload", string(payload))
				oid := strings.TrimSpace(gitFixtureCommand(t, root, "hash-object", "-w", "payload"))
				gitFixtureCommand(t, root, "update-ref", "refs/fireweed/claims/current", oid)
			case "projection":
				cliQueueChange(t, root, contextChangePath, "alpha", "ALP", "", "## Added\n\n### ALP-001 [review]\nDuplicate MUST fail projection.\n")
			case "missing-spec", "required-directory", "escape":
				cliContextClaim(t, root, "impl", false)
				code = 2
				if mode == "missing-spec" {
					if e := os.Remove(filepath.Join(root, "units/a/SPEC.md")); e != nil {
						t.Fatal(e)
					}
				} else {
					file := filepath.Join(root, "units/a/public/new.txt")
					if e := os.Remove(file); e != nil {
						t.Fatal(e)
					}
					if mode == "required-directory" {
						if e := os.Mkdir(file, 0755); e != nil {
							t.Fatal(e)
						}
					} else {
						outside := filepath.Join(t.TempDir(), "secret")
						projectPut(t, filepath.Dir(outside), "secret", "Outside bytes")
						if e := os.Symlink(outside, file); e != nil {
							t.Fatal(e)
						}
					}
				}
			case "missing-change":
				if e := os.Remove(filepath.Join(root, contextChangePath)); e != nil {
					t.Fatal(e)
				}
				slug = "current" // No selected source remains: C05.
			case "missing-config":
				if e := os.Remove(filepath.Join(root, "fireweed.yaml")); e != nil {
					t.Fatal(e)
				}
				code = 2
			case "bad-config":
				projectPut(t, root, "fireweed.yaml", "invalid: [")
				code = 2
			case "override-directory":
				if e := os.MkdirAll(filepath.Join(root, ".fireweed/playbooks/author-change.md"), 0755); e != nil {
					t.Fatal(e)
				}
				code = 2
			}
			err := cliContextPublicError(root, slug)
			wantPrefix := "C05:"
			if strings.Contains(mode, "claim") || mode == "foreign" || mode == "reservation" || mode == "baseline" || mode == "state" {
				wantPrefix = "P02:"
			}
			if code == 2 {
				wantPrefix = "C01:"
			}
			if mode == "override-directory" {
				wantPrefix = "D05:"
			}
			if err == nil || !strings.HasPrefix(err.Error(), wantPrefix) {
				t.Fatalf("error fixture %s: %v", mode, err)
			}
			cliContextError(t, root, slug, err, code)
		})
	}
	outside := t.TempDir()
	cliContextError(t, outside, "current", cliContextPublicError(outside, "current"), 2)
	cliContextError(t, outside, "../escape", cliContextPublicError(outside, "../escape"), 1)
}

func Test_APP019_APP020ContextOptionalAndUnrelatedProblems(t *testing.T) {
	root := cliContextFixture(t, "sha1")
	if err := os.Remove(filepath.Join(root, "SYSTEM.md")); err != nil {
		t.Fatal(err)
	}
	projectPut(t, root, "units/g/SPEC.md", "Malformed unrelated current module.\n")
	// Remove the transitive relationship, retaining the direct selected dependency.
	raw, err := os.ReadFile(filepath.Join(root, "units/b/SPEC.md"))
	if err != nil {
		t.Fatal(err)
	}
	projectPut(t, root, "units/b/SPEC.md", strings.Replace(string(raw), "depends_on: [gamma]\n", "", 1))
	s, err := p.Load(root)
	if err != nil || len(s.Diagnostics) == 0 {
		t.Fatalf("discovery diagnostic premise %v %+v", err, s.Diagnostics)
	}
	want := cliContextExpected(t, root)
	if want.System != "" {
		t.Fatal("optional system control")
	}
	cliContextAssert(t, root, want)
}

func Test_APP020_APP021ContextArgumentsAndWriters(t *testing.T) {
	outside := t.TempDir()
	for _, args := range [][]string{{"context"}, {"context", "current", "extra"}, {"context", "current", "--unknown"}, {"context", "current", "--json=true"}, {"context", "current", "--help=true"}, {"context", "current", "--projected"}, {"context", "current", "--history"}, {"context", "current", "--with-change", "other"}, {"context", "--", "--json", "current"}} {
		assertUsage(t, invoke(t, outside, args...))
		jsonArgs := append([]string{"--json"}, args...)
		got := invoke(t, outside, jsonArgs...)
		x := decode(t, got, "context")
		if got.code != 2 || string(x.Data) != "null" || len(x.Diagnostics) != 1 || x.Diagnostics[0].Code != "U01" || x.Diagnostics[0].Severity != "error" || !strings.Contains(x.Diagnostics[0].Message, "help") {
			t.Errorf("JSON usage %+v", got)
		}
	}
	for _, args := range [][]string{{"context", "--help"}, {"context", "-h"}, {"--json", "context", "--help"}} {
		got := invoke(t, outside, args...)
		if got.code != 0 {
			t.Errorf("help %+v", got)
		}
		if args[0] == "--json" {
			successJSON(t, got, "help")
		} else if !strings.Contains(got.out, "context") {
			t.Error("missing help")
		}
	}
	got := invoke(t, outside, "context", "--", "--json")
	if got.code != 1 || got.out != "" || !strings.Contains(got.err, "C05") {
		t.Errorf("stopped JSON %+v", got)
	}
	root := cliContextFixture(t, "sha1")
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err = os.Chdir(root); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(cwd)
	for _, args := range [][]string{{"context", "current"}, {"context", "current", "--json"}, {"context", "missing"}, {"context", "missing", "--json"}, {"context"}, {"context", "--json"}} {
		out, errOut := &bytes.Buffer{}, &bytes.Buffer{}
		jsonMode := args[len(args)-1] == "--json"
		plainFailure := !jsonMode && (len(args) == 1 || args[1] == "missing")
		if plainFailure {
			if code := app.Run(args, out, failedWriter{}); code != 2 {
				t.Errorf("stderr failure %v exit%d", args, code)
			}
		} else {
			if code := app.Run(args, failedWriter{}, errOut); code != 2 {
				t.Errorf("stdout failure %v exit%d", args, code)
			}
		}
	}
}

func Test_APP021ContextCompleteStatePreservation(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("POSIX process traps; portable composition/error cases run separately")
	}
	root := cliContextFixture(t, "sha1")
	cliContextClaim(t, root, "impl", false)
	// An unrelated corrupt claim must not prevent an explicitly selected context.
	projectPut(t, root, "unrelated-payload", "not JSON")
	oid := strings.TrimSpace(gitFixtureCommand(t, root, "hash-object", "-w", "unrelated-payload"))
	gitFixtureCommand(t, root, "update-ref", "refs/fireweed/claims/unrelated", oid)
	gitFixtureCommand(t, root, "update-ref", "refs/custom/context", oid)
	projectPut(t, root, "units/a/public/dirty.txt", "Staged retained input.")
	gitFixtureCommand(t, root, "add", "units/a/public/dirty.txt")
	projectPut(t, root, "units/a/public/new.txt", "Dirty feedback.")
	projectPut(t, root, "units/a/evals/untracked.txt", "Untracked selected input.")
	linked := filepath.Join(t.TempDir(), "peer")
	gitFixtureCommand(t, root, "worktree", "add", "--detach", linked, "HEAD")
	s, _, err := w.LoadApproval(root, "current")
	if err != nil {
		t.Fatal(err)
	}
	want, err := w.BuildContext(s, "current")
	if err != nil {
		t.Fatal(err)
	}
	monitorRoot := t.TempDir()
	monitorMarker := filepath.Join(monitorRoot, "invoked")
	projectPut(t, monitorRoot, "monitor", "#!/bin/sh\nprintf invoked >> "+coverageQuote(monitorMarker)+"\nprintf '\\000'\n")
	if err := os.Chmod(filepath.Join(monitorRoot, "monitor"), 0755); err != nil {
		t.Fatal(err)
	}
	gitFixtureCommand(t, root, "config", "core.fsmonitor", coverageQuote(filepath.Join(monitorRoot, "monitor")))
	gitFixtureCommand(t, root, "status", "--porcelain")
	if _, err := os.Stat(monitorMarker); err != nil {
		t.Fatalf("native fsmonitor positive control: %v", err)
	}
	if err := os.Remove(monitorMarker); err != nil {
		t.Fatal(err)
	}
	marker := playbookCLITraps(t)
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	env := os.Environ()
	temp := t.TempDir()
	t.Setenv("TMPDIR", temp)
	env = os.Environ()
	before, peer, tempBefore := playbookCLIState(t, root), playbookCLIState(t, linked), playbookCLIState(t, temp)
	cliContextAssert(t, root, want)
	for _, slug := range []string{"unrelated", "missing", "../bad"} {
		cliContextError(t, root, slug, cliContextPublicError(root, slug), 1)
	}
	if !reflect.DeepEqual(before, playbookCLIState(t, root)) || !reflect.DeepEqual(peer, playbookCLIState(t, linked)) || !reflect.DeepEqual(tempBefore, playbookCLIState(t, temp)) {
		t.Fatal("context changed file/index/ref/common Git/temp state")
	}
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Fatalf("forbidden command attempt %v", err)
	}
	if _, err := os.Stat(monitorMarker); !os.IsNotExist(err) {
		t.Fatalf("native Git hook executed %v", err)
	}
	after, err := os.Getwd()
	if err != nil || after != cwd || !reflect.DeepEqual(env, os.Environ()) {
		t.Fatal("parent cwd/environment changed")
	}
}

func Test_APP019ContextFixtureControls(t *testing.T) {
	for _, phase := range []string{"", "evals", "audit", "impl", "regenerate"} {
		t.Run(fmt.Sprintf("phase_%s", phase), func(t *testing.T) {
			root := cliContextFixture(t, "sha1")
			if phase != "" {
				actual := phase
				if actual == "regenerate" {
					actual = "impl"
				}
				cliContextClaim(t, root, actual, phase == "regenerate")
			}
			cliContextExpected(t, root)
		})
	}
}

func Test_APP020UnreadableContextRequiredInput(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("POSIX permission denial control")
	}
	root := cliContextFixture(t, "sha1")
	cliContextClaim(t, root, "impl", false)
	file := filepath.Join(root, "units/a/public/new.txt")
	if err := os.Chmod(file, 0000); err != nil {
		t.Fatal(err)
	}
	defer os.Chmod(file, 0644)
	if _, err := os.ReadFile(file); err == nil {
		t.Fatal("unreadable fixture control unexpectedly readable")
	}
	err := cliContextPublicError(root, "current")
	if err == nil || !strings.HasPrefix(err.Error(), "C01:") {
		t.Fatalf("required read error premise: %v", err)
	}
	cliContextError(t, root, "current", err, 2)
}

func Test_APP020ContextAllOperationalWriterFailures(t *testing.T) {
	root := cliContextFixture(t, "sha1")
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(root); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(cwd)
	for _, mode := range []string{"C05", "P02", "D05", "C01"} {
		slug := "current"
		switch mode {
		case "C05":
			slug = "missing"
		case "P02":
			if _, err := w.CreateClaim(root, "current", "reserved", strings.TrimSpace(gitFixtureCommand(t, root, "rev-parse", "HEAD")), ""); err != nil {
				t.Fatal(err)
			}
		case "D05":
			claim, err := w.LoadClaim(root, "current")
			if err != nil {
				t.Fatal(err)
			}
			if err = w.ReleaseClaim(root, "current", claim.Object); err != nil {
				t.Fatal(err)
			}
			if err = os.MkdirAll(filepath.Join(root, ".fireweed/playbooks/author-change.md"), 0755); err != nil {
				t.Fatal(err)
			}
		case "C01":
			projectPut(t, root, "fireweed.yaml", "invalid: [")
		}
		premise := cliContextPublicError(root, slug)
		if premise == nil || !strings.HasPrefix(premise.Error(), mode+":") {
			t.Fatalf("writer error premise %s %v", mode, premise)
		}
		if code := app.Run([]string{"context", slug, "--json"}, failedWriter{}, &bytes.Buffer{}); code != 2 {
			t.Errorf("%s JSON writer exit%d", mode, code)
		}
		if code := app.Run([]string{"context", slug}, &bytes.Buffer{}, failedWriter{}); code != 2 {
			t.Errorf("%s plain writer exit%d", mode, code)
		}
	}
}

func Test_APP019_APP021ContextRetainsUnallocatedHandles(t *testing.T) {
	root := cliContextFixture(t, "sha1")
	for _, row := range []struct{ path, id, handle string }{{"changes/0001-before.md", "BET-002", "new:preceding"}, {contextChangePath, "ALP-002", "new:selected"}, {"changes/0003-later.md", "ALP-003", "new:later"}} {
		raw, err := os.ReadFile(filepath.Join(root, row.path))
		if err != nil {
			t.Fatal(err)
		}
		text := strings.ReplaceAll(string(raw), row.id, row.handle)
		if row.path == contextChangePath {
			text += "\n## Interface\n\nReferences new:selected without allocating it.\n"
		}
		if _, err := c.ParseChange(text, row.path, map[string]string{"BET-002": "BET", "ALP-002": "ALP", "ALP-003": "ALP"}[row.id]); err != nil {
			t.Fatal(err)
		}
		projectPut(t, root, row.path, text)
	}
	s, claim, err := w.LoadApproval(root, "current")
	if err != nil || claim != (w.Claim{}) {
		t.Fatalf("unclaimed handle source %v %+v", err, claim)
	}
	want, err := w.BuildContext(s, "current")
	if err != nil {
		t.Fatal(err)
	}
	if want.Module.Specs[1].ID != "new:selected" || want.Dependencies[0].Specs[1].ID != "new:preceding" || want.Change.Added[0].ID != "new:selected" || !strings.Contains(want.ChangeText, "References new:selected") {
		t.Fatalf("handle fixture control %+v", want)
	}
	before := playbookCLIState(t, root)
	cliContextAssert(t, root, want)
	if !reflect.DeepEqual(before, playbookCLIState(t, root)) {
		t.Fatal("context allocated or rewrote handle-bearing changes/specs/refs")
	}
}
