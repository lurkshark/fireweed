package evals

import (
	"bytes"
	"encoding/json"
	"fmt"
	app "github.com/lurkshark/fireweed/packages/app"
	a "github.com/lurkshark/fireweed/packages/assets"
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

func boundPlaybookRoot(t *testing.T) string {
	t.Helper()
	root := projectFixture(t)
	root, err := filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatal(err)
	}
	projectPut(t, root, ".gitignore", "ignored\n")
	projectPut(t, root, "ignored", "ignored retained bytes")
	gitFixtureCommand(t, root, "add", ".")
	gitFixtureCommand(t, root, "-c", "user.name=Playbook Fixture", "-c", "user.email=fixture@example.invalid", "commit", "-m", "playbook fixture")
	return root
}
func boundPlaybookData(v p.Playbook) map[string]any {
	return map[string]any{"name": v.Name, "text": v.Text, "hash": v.Hash, "embedded_hash": v.EmbeddedHash, "overridden": v.Overridden}
}
func boundPlaybookAssert(t *testing.T, where, name string, want p.Playbook) {
	t.Helper()
	plain := invoke(t, where, "playbook", name)
	if plain.code != 0 || plain.out != want.Text || plain.err != "" {
		t.Fatalf("plain playbook %+v; want exact %q", plain, want.Text)
	}
	for _, args := range [][]string{{"playbook", name, "--json"}, {"--json", "playbook", name}, {"playbook", "--json", name}} {
		got := invoke(t, where, args...)
		env := successJSON(t, got, "playbook")
		if !reflect.DeepEqual(decodeData(t, env.Data), boundPlaybookData(want)) {
			t.Fatalf("JSON playbook %s; want %+v", env.Data, boundPlaybookData(want))
		}
	}
}
func boundPlaybookExpected(t *testing.T, root, name string) p.Playbook {
	t.Helper()
	v, err := p.ResolvePlaybook(root, name)
	if err != nil {
		t.Fatal(err)
	}
	return v
}
func boundPlaybookEmbedded(t *testing.T, name string) p.Playbook {
	t.Helper()
	v, err := a.Playbook(name)
	if err != nil {
		t.Fatal(err)
	}
	return p.Playbook{Name: name, Text: v.Text, Hash: v.Hash, EmbeddedHash: v.Hash}
}
func Test_APP002_APP017ProjectOverrideComposition(t *testing.T) {
	root := boundPlaybookRoot(t)
	nested := filepath.Join(root, "nested 'quoted 雪 space'")
	projectPut(t, nested, "fireweed.yaml", "invalid nested decoy [")
	for _, name := range a.Playbooks() {
		asset, _ := a.Playbook(name)
		for _, row := range []struct{ name, text string }{{"current", "---\nbased_on: '" + asset.Hash + "'\n---\nCurrent project instructions without final newline"}, {"stale", "---\nbased_on: wrong\n---\nStale project instructions.\n"}, {"malformed", "---\nbased_on: [\n---\nMalformed retained instructions.\n"}, {"empty", ""}, {"crlf", "---\r\nbased_on: '" + asset.Hash + "'\r\n---\r\nCRLF instructions.\r\n"}} {
			t.Run(name+"_"+row.name, func(t *testing.T) {
				projectPut(t, root, ".fireweed/playbooks/"+name+".md", row.text)
				projectPut(t, nested, ".fireweed/playbooks/"+name+".md", "nested override decoy")
				want := boundPlaybookExpected(t, root, name)
				if !want.Overridden || want.Text != row.text {
					t.Fatal("public resolver premise")
				}
				boundPlaybookAssert(t, nested, name, want)
			})
		}
	}
}
func Test_APP017LinkedAndNestedCheckoutBoundaries(t *testing.T) {
	root := boundPlaybookRoot(t)
	projectPut(t, root, ".fireweed/playbooks/impl.md", "primary checkout instructions")
	peer := filepath.Join(t.TempDir(), "linked")
	gitFixtureCommand(t, root, "worktree", "add", "--detach", peer, "HEAD")
	projectPut(t, peer, ".fireweed/playbooks/impl.md", "linked checkout instructions")
	nested := filepath.Join(peer, "nested")
	projectPut(t, nested, "fireweed.yaml", "invalid nested decoy [")
	projectPut(t, nested, ".fireweed/playbooks/impl.md", "nested decoy")
	boundPlaybookAssert(t, nested, "impl", boundPlaybookExpected(t, peer, "impl"))
	child := filepath.Join(root, "nested-checkout")
	if err := os.Mkdir(child, 0755); err != nil {
		t.Fatal(err)
	}
	gitFixtureCommand(t, child, "init", "--initial-branch=main")
	projectPut(t, child, ".fireweed/playbooks/impl.md", "unbound child override")
	boundPlaybookAssert(t, child, "impl", boundPlaybookEmbedded(t, "impl"))
	projectPut(t, child, "fireweed.yaml", projectConfig)
	boundPlaybookAssert(t, child, "impl", boundPlaybookExpected(t, child, "impl"))
}
func Test_APP017EmbeddedFallbackAndUnrelatedInvalidDocuments(t *testing.T) {
	root := boundPlaybookRoot(t)
	name := "evals"
	boundPlaybookAssert(t, root, name, boundPlaybookExpected(t, root, name))
	projectPut(t, root, "units/a/SPEC.md", "invalid unrelated module [")
	projectPut(t, root, "changes/0001-broken.md", "invalid unrelated change [")
	projectPut(t, root, ".fireweed/playbooks/evals.md", "Selected instructions despite unrelated invalid documents.")
	boundPlaybookAssert(t, root, name, boundPlaybookExpected(t, root, name))
}
func boundPlaybookFailure(t *testing.T, where, name, code, path string) {
	t.Helper()
	plain := invoke(t, where, "playbook", name)
	if plain.code != 2 || plain.out != "" || !strings.Contains(plain.err, code) || !strings.Contains(plain.err, path) {
		t.Fatalf("plain operational failure %+v", plain)
	}
	got := invoke(t, where, "--json", "playbook", name)
	env := decode(t, got, "playbook")
	ds := detailedDiagnostics(t, got)
	if got.code != 2 || string(env.Data) != "null" || len(ds) != 1 || ds[0].Code != code || ds[0].Severity != "error" || ds[0].Path != path || !playbookAction(ds[0].Message) {
		t.Fatalf("operational error %+v %+v", got, ds)
	}
	if code == "D05" {
		_, err := p.ResolvePlaybook(where, name)
		if err == nil || !strings.Contains(ds[0].Message, err.Error()) {
			t.Fatalf("lost exact public resolver detail %+v %v", ds, err)
		}
	}
}
func playbookAction(text string) bool {
	text = strings.ToLower(text)
	for _, word := range []string{"restore", "correct", "run ", "open", "use ", "repair", "choose", "check"} {
		if strings.Contains(text, word) {
			return true
		}
	}
	return false
}
func Test_APP003_APP018OverrideOperationalErrors(t *testing.T) {
	for _, mode := range []string{"directory", "unreadable", "outside", "dangling", "outside_directory", "missing_outside_descendant"} {
		t.Run(mode, func(t *testing.T) {
			root := boundPlaybookRoot(t)
			path := ".fireweed/playbooks/impl.md"
			projectPut(t, root, path, "old override")
			full := filepath.Join(root, path)
			outside := t.TempDir()
			outside, err := filepath.EvalSymlinks(outside)
			if err != nil {
				t.Fatal(err)
			}
			projectPut(t, outside, "private", "PRIVATE RETAINED")
			switch mode {
			case "directory":
				if err := os.Remove(full); err != nil {
					t.Fatal(err)
				}
				if err := os.Mkdir(full, 0755); err != nil {
					t.Fatal(err)
				}
			case "unreadable":
				if err := os.Chmod(full, 0000); err != nil {
					t.Fatal(err)
				}
				defer os.Chmod(full, 0644)
				if _, err := os.ReadFile(full); !os.IsPermission(err) {
					t.Skipf("host does not establish unreadability: %v", err)
				}
			case "outside", "dangling":
				if err := os.Remove(full); err != nil {
					t.Fatal(err)
				}
				target := filepath.Join(outside, "private")
				if mode == "dangling" {
					target = filepath.Join(outside, "missing")
				}
				if err := os.Symlink(target, full); err != nil {
					t.Skipf("symlink unavailable: %v", err)
				}
			case "outside_directory", "missing_outside_descendant":
				if mode == "outside_directory" {
					projectPut(t, outside, "impl.md", "private override")
				}
				if err := os.RemoveAll(filepath.Dir(full)); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(outside, filepath.Dir(full)); err != nil {
					t.Skipf("symlink unavailable: %v", err)
				}
			}
			boundPlaybookFailure(t, root, "impl", "D05", path)
		})
	}
}
func Test_APP003_APP018ConfigAndGitOperationalErrors(t *testing.T) {
	for _, mode := range []string{"malformed", "pin", "format", "unknown_field", "wrong_type", "directory", "unreadable", "outside", "dangling", "missing_git"} {
		t.Run(mode, func(t *testing.T) {
			root := boundPlaybookRoot(t)
			full := filepath.Join(root, "fireweed.yaml")
			switch mode {
			case "malformed":
				projectPut(t, root, "fireweed.yaml", "version: [\n")
			case "pin":
				projectPut(t, root, "fireweed.yaml", projectConfig+"fireweed: '>999.0'\n")
			case "format":
				projectPut(t, root, "fireweed.yaml", strings.Replace(projectConfig, "version: 1", "version: 99", 1))
			case "unknown_field":
				projectPut(t, root, "fireweed.yaml", projectConfig+"unexpected_setting: true\n")
			case "wrong_type":
				projectPut(t, root, "fireweed.yaml", strings.Replace(projectConfig, "modules: [units/*]", "modules: units/*", 1))
			case "directory":
				if err := os.Remove(full); err != nil {
					t.Fatal(err)
				}
				if err := os.Mkdir(full, 0755); err != nil {
					t.Fatal(err)
				}
			case "unreadable":
				if err := os.Chmod(full, 0000); err != nil {
					t.Fatal(err)
				}
				defer os.Chmod(full, 0644)
				if _, err := os.ReadFile(full); !os.IsPermission(err) {
					t.Skipf("host does not establish unreadability: %v", err)
				}
			case "outside", "dangling":
				if err := os.Remove(full); err != nil {
					t.Fatal(err)
				}
				outside := filepath.Join(t.TempDir(), "config")
				if mode == "outside" {
					projectPut(t, filepath.Dir(outside), "config", projectConfig)
				}
				if err := os.Symlink(outside, full); err != nil {
					t.Skipf("symlink unavailable: %v", err)
				}
			case "missing_git":
				if runtime.GOOS == "windows" {
					t.Skip("POSIX Git discovery trap")
				}
				traps := t.TempDir()
				marker := filepath.Join(traps, "git-attempted")
				projectPut(t, traps, "git", "#!/bin/sh\nprintf attempted > "+coverageQuote(marker)+"\nexit 97\n")
				if err := os.Chmod(filepath.Join(traps, "git"), 0755); err != nil {
					t.Fatal(err)
				}
				t.Setenv("PATH", traps)
				boundPlaybookFailure(t, root, "impl", "C01", "fireweed.yaml")
				if _, err := os.Stat(marker); err != nil {
					t.Fatal("bound project never attempted Git discovery")
				}
				return
			}
			boundPlaybookFailure(t, root, "impl", "C01", "fireweed.yaml")
		})
	}
}
func Test_APP017StandaloneNoGitAndNoConfigBinding(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("POSIX Git/bootstrap traps")
	}
	root := boundPlaybookRoot(t)
	if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
		t.Fatal(err)
	}
	projectPut(t, root, ".fireweed/playbooks/impl.md", "must remain unselected without config")
	outside := t.TempDir()
	projectPut(t, outside, "fireweed.yaml", "invalid unbound config [")
	projectPut(t, outside, ".fireweed/playbooks/impl.md", "must remain unselected without checkout")
	traps := t.TempDir()
	marker := filepath.Join(traps, "attempted")
	for _, name := range []string{"git", "fireweed", "fireweed-bootstrap", "codex", "claude"} {
		projectPut(t, traps, name, "#!/bin/sh\nprintf attempted >> "+coverageQuote(marker)+"\nexit 97\n")
		if err := os.Chmod(filepath.Join(traps, name), 0755); err != nil {
			t.Fatal(err)
		}
	}
	t.Setenv("PATH", traps)
	t.Setenv("FIREWEED_BIN", filepath.Join(traps, "fireweed-bootstrap"))
	if err := exec.Command("git", "--version").Run(); err == nil {
		t.Fatal("trap control passed")
	}
	if _, err := os.Stat(marker); err != nil {
		t.Fatal("trap control not reached")
	}
	if err := os.Remove(marker); err != nil {
		t.Fatal(err)
	}
	for _, where := range []string{root, outside} {
		boundPlaybookAssert(t, where, "impl", boundPlaybookEmbedded(t, "impl"))
	}
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Fatal("standalone lookup executed Git/agent/bootstrap")
	}
}
func Test_APP018ArgumentsAndWriterFailures(t *testing.T) {
	root := boundPlaybookRoot(t)
	projectPut(t, root, "fireweed.yaml", "invalid config [")
	for _, args := range [][]string{{"playbook"}, {"playbook", "unknown", "--json"}, {"playbook", "impl", "extra", "--json"}, {"playbook", "impl", "--unknown", "--json"}, {"playbook", "impl", "--json=true"}, {"playbook", "--", "--json"}} {
		got := invoke(t, root, args...)
		if got.code != 2 || !strings.Contains(got.out+got.err, "U01") {
			t.Fatalf("usage must precede binding: %v %+v", args, got)
		}
	}
	for _, args := range [][]string{{"playbook", "--help"}, {"playbook", "-h", "--json"}, {"playbook", "impl", "--help"}} {
		if got := invoke(t, root, args...); got.code != 0 {
			t.Fatalf("help loaded project: %+v", got)
		}
	}
	original, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(root); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(original)
	for _, mode := range []string{"embedded", "override", "operational_error"} {
		switch mode {
		case "embedded":
			if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
				t.Fatal(err)
			}
		case "override":
			projectPut(t, root, "fireweed.yaml", projectConfig)
			projectPut(t, root, ".fireweed/playbooks/impl.md", "project override")
		case "operational_error":
			projectPut(t, root, "fireweed.yaml", "invalid config [")
		}
		for _, args := range [][]string{{"playbook", "impl"}, {"playbook", "impl", "--json"}} {
			for _, writers := range [][2]bool{{true, false}, {false, true}} {
				var out, errOut bytes.Buffer
				var stdout, stderr interface{ Write([]byte) (int, error) } = &out, &errOut
				if writers[0] {
					stdout = failedWriter{}
				}
				if writers[1] {
					stderr = failedWriter{}
				}
				code := app.Run(args, stdout, stderr)
				if mode != "operational_error" && writers[1] {
					if code != 0 {
						t.Fatalf("unused stderr failure should not fail success: %s %v %d", mode, args, code)
					}
				} else if code != 2 {
					t.Fatalf("writer failure status %s %v %d", mode, args, code)
				}
			}
		}
	}
}
func Test_APP017ConfigFixtureControl(t *testing.T) {
	if _, err := c.ParseConfig(projectConfig, a.Info().Version); err != nil {
		t.Fatal(err)
	}
	var zero p.Playbook
	raw, err := json.Marshal(boundPlaybookData(zero))
	if err != nil || !strings.Contains(string(raw), `"overridden":false`) {
		t.Fatal("JSON fixture invalid")
	}
}

func playbookCLIState(t *testing.T, root string) map[string]string {
	t.Helper()
	result := map[string]string{}
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		info, err := os.Lstat(path)
		if err != nil {
			return err
		}
		value := info.Mode().String()
		if info.Mode()&os.ModeSymlink != 0 {
			target, err := os.Readlink(path)
			if err != nil {
				return err
			}
			value += " link:" + target
		} else if info.Mode().IsRegular() {
			raw, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			value += " bytes:" + string(raw)
		}
		result[filepath.ToSlash(rel)] = value
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return result
}
func playbookCLITraps(t *testing.T) string {
	t.Helper()
	traps := t.TempDir()
	marker := filepath.Join(traps, "attempted")
	realGit, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	script := "#!/bin/sh\nfor a in \"$@\";do case \"$a\" in fetch|push|pull|clone|ls-remote|commit|branch|checkout|switch|reset|restore|add|worktree|update-ref|hash-object|stash|clean|rm|mv) printf attempted >> " + coverageQuote(marker) + ";exit 97;;esac;done\nexec " + coverageQuote(realGit) + " \"$@\"\n"
	projectPut(t, traps, "git", script)
	for _, name := range []string{"git", "hook-trap", "touch", "agent-trap", "codex", "claude", "fireweed", "fireweed-bootstrap", "curl", "wget", "ssh", "nc", "sh", "bash", "zsh"} {
		if name != "git" {
			projectPut(t, traps, name, "#!/bin/sh\nprintf attempted >> "+coverageQuote(marker)+"\nexit 97\n")
		}
		if err := os.Chmod(filepath.Join(traps, name), 0755); err != nil {
			t.Fatal(err)
		}
	}
	t.Setenv("PATH", traps+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("FIREWEED_BIN", filepath.Join(traps, "fireweed-bootstrap"))
	for _, argv := range [][]string{{"git", "fetch", "forbidden.invalid"}, {"hook-trap"}} {
		if err := exec.Command(argv[0], argv[1:]...).Run(); err == nil {
			t.Fatal("trap positive control passed")
		}
		if _, err := os.Stat(marker); err != nil {
			t.Fatal("trap control absent")
		}
		if err := os.Remove(marker); err != nil {
			t.Fatal(err)
		}
	}
	return marker
}
func Test_APP018PlaybookCompleteStatePreservation(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("POSIX execution traps; ordinary command cases remain portable")
	}
	for _, linked := range []bool{false, true} {
		for _, mode := range []string{"embedded", "override", "stale", "config_error", "override_error"} {
			t.Run(fmt.Sprintf("%s_linked_%t", mode, linked), func(t *testing.T) {
				root := boundPlaybookRoot(t)
				base := strings.TrimSpace(gitFixtureCommand(t, root, "rev-parse", "HEAD"))
				projectPut(t, root, ".git/claim-fixture.json", fmt.Sprintf(`{"worktree":%q,"change":"fixture","run":"guard","base":%q,"state":""}`, root, base))
				blob := strings.TrimSpace(gitFixtureCommand(t, root, "hash-object", "-w", ".git/claim-fixture.json"))
				gitFixtureCommand(t, root, "update-ref", "refs/fireweed/claims/fixture", blob)
				gitFixtureCommand(t, root, "update-ref", "refs/custom/fixture", base)
				if linked {
					peer := filepath.Join(t.TempDir(), "peer")
					gitFixtureCommand(t, root, "worktree", "add", "--detach", peer, "HEAD")
					root = peer
				}
				root, err := filepath.EvalSymlinks(root)
				if err != nil {
					t.Fatal(err)
				}
				common := strings.TrimSpace(gitFixtureCommand(t, root, "rev-parse", "--git-common-dir"))
				if !filepath.IsAbs(common) {
					common = filepath.Join(root, common)
				}
				common, err = filepath.EvalSymlinks(common)
				if err != nil {
					t.Fatal(err)
				}
				outside, err := filepath.EvalSymlinks(t.TempDir())
				if err != nil {
					t.Fatal(err)
				}
				projectPut(t, outside, "private", "PRIVATE RETAINED")
				path := ".fireweed/playbooks/impl.md"
				asset, _ := a.Playbook("impl")
				switch mode {
				case "override":
					projectPut(t, root, path, "---\nbased_on: '"+asset.Hash+"'\n---\nPreserved override instructions.\n")
				case "stale":
					projectPut(t, root, path, "Preserved stale instructions.\n")
				case "config_error":
					projectPut(t, root, "fireweed.yaml", "invalid config [")
				case "override_error":
					projectPut(t, root, ".fireweed/playbooks/.keep", "")
					if err := os.Symlink(filepath.Join(outside, "private"), filepath.Join(root, path)); err != nil {
						t.Skipf("symlink unavailable: %v", err)
					}
				}
				marker := playbookCLITraps(t)
				temporary, err := filepath.EvalSymlinks(t.TempDir())
				if err != nil {
					t.Fatal(err)
				}
				t.Setenv("TMPDIR", temporary)
				t.Setenv("TMP", temporary)
				t.Setenv("TEMP", temporary)
				gitFixtureCommand(t, root, "--version")
				original, err := os.Getwd()
				if err != nil {
					t.Fatal(err)
				}
				if err := os.Chdir(root); err != nil {
					t.Fatal(err)
				}
				t.Cleanup(func() { os.Chdir(original) })
				beforeRoot, beforeCommon, beforeOutside, beforeTemp := playbookCLIState(t, root), playbookCLIState(t, common), playbookCLIState(t, outside), playbookCLIState(t, temporary)
				environment := os.Environ()
				cwd, err := os.Getwd()
				if err != nil {
					t.Fatal(err)
				}
				if mode == "config_error" {
					boundPlaybookFailure(t, root, "impl", "C01", "fireweed.yaml")
				} else if mode == "override_error" {
					boundPlaybookFailure(t, root, "impl", "D05", path)
				} else {
					boundPlaybookAssert(t, root, "impl", boundPlaybookExpected(t, root, "impl"))
				}
				for _, args := range [][]string{{"playbook", "impl"}, {"playbook", "impl", "--json"}} {
					var out, errOut bytes.Buffer
					code := app.Run(args, &out, &errOut)
					want := 0
					if strings.HasSuffix(mode, "error") {
						want = 2
					}
					if code != want {
						t.Fatalf("public entry status %s %d; expected %d", mode, code, want)
					}
				}
				after, err := os.Getwd()
				if err != nil || after != cwd || !reflect.DeepEqual(environment, os.Environ()) || !reflect.DeepEqual(beforeRoot, playbookCLIState(t, root)) || !reflect.DeepEqual(beforeCommon, playbookCLIState(t, common)) || !reflect.DeepEqual(beforeOutside, playbookCLIState(t, outside)) || !reflect.DeepEqual(beforeTemp, playbookCLIState(t, temporary)) {
					t.Fatal("playbook inspection changed files, ignored/private data, Git/index/refs/common storage, cwd/environment or temporary storage")
				}
				if _, err := os.Stat(marker); !os.IsNotExist(err) {
					t.Fatal("playbook inspection executed forbidden commands")
				}
			})
		}
	}
}

func Test_APP017PreviousFormatAndContainedConfigLinks(t *testing.T) {
	for _, mode := range []string{"previous", "inside_links"} {
		t.Run(mode, func(t *testing.T) {
			root := boundPlaybookRoot(t)
			config := projectConfig
			if mode == "previous" {
				config = strings.Replace(config, "version: 1", "version: 0", 1)
			}
			if _, err := c.ParseConfig(config, a.Info().Version); err != nil {
				t.Fatalf("compatible config premise: %v", err)
			}
			projectPut(t, root, "fireweed.yaml", config)
			projectPut(t, root, ".fireweed/playbooks/impl.md", "Contained override instructions.")
			if mode == "inside_links" {
				projectPut(t, root, "config/current.yaml", config)
				projectPut(t, root, "config/instructions.md", "Contained linked instructions.")
				if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(filepath.Join(root, "config/current.yaml"), filepath.Join(root, "fireweed.yaml")); err != nil {
					t.Skipf("symlink unavailable: %v", err)
				}
				if err := os.Remove(filepath.Join(root, ".fireweed/playbooks/impl.md")); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(filepath.Join(root, "config/instructions.md"), filepath.Join(root, ".fireweed/playbooks/impl.md")); err != nil {
					t.Fatal(err)
				}
			}
			boundPlaybookAssert(t, root, "impl", boundPlaybookExpected(t, root, "impl"))
		})
	}
}
