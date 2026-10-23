package evals_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
)

type baselineFixtureData struct {
	Root, Commit string
	Expected     p.Snapshot
	Files        map[string]string
}

func baselineInput(t *testing.T, root, input string, args ...string) (string, error) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = root
	cmd.Stdin = strings.NewReader(input)
	cmd.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_CONFIG_GLOBAL="+os.DevNull, "GIT_OPTIONAL_LOCKS=0", "GIT_AUTHOR_NAME=Baseline Fixture", "GIT_AUTHOR_EMAIL=baseline@example.invalid", "GIT_COMMITTER_NAME=Baseline Fixture", "GIT_COMMITTER_EMAIL=baseline@example.invalid")
	out, err := cmd.CombinedOutput()
	return string(out), err
}
func baselineGit(t *testing.T, root string, args ...string) string {
	t.Helper()
	out, err := baselineInput(t, root, "", args...)
	if err != nil {
		t.Fatalf("fixture Git %q: %v %s", args, err, out)
	}
	return strings.TrimSpace(out)
}
func baselineCommit(t *testing.T, root string) string {
	t.Helper()
	baselineGit(t, root, "add", ".")
	baselineGit(t, root, "commit", "-m", "baseline fixture")
	return baselineGit(t, root, "rev-parse", "HEAD")
}
func baselineCanonical(t *testing.T, path string) string {
	t.Helper()
	out, err := filepath.EvalSymlinks(path)
	if err != nil {
		t.Fatal(err)
	}
	return out
}
func baselineFixture(t *testing.T, format string) baselineFixtureData {
	t.Helper()
	root := filepath.Join(t.TempDir(), "repository ' λ with spaces")
	if err := os.MkdirAll(root, 0755); err != nil {
		t.Fatal(err)
	}
	out, err := baselineInput(t, root, "", "init", "-b", "main", "--object-format="+format)
	if err != nil {
		if format == "sha256" {
			t.Skipf("installed Git lacks SHA-256: %v %s", err, out)
		}
		t.Fatal(err, out)
	}
	root = baselineCanonical(t, root)
	cfg := strings.ReplaceAll(config, "touch MUST-NOT-RUN", "baseline-hook-trap") + "  setup: baseline-hook-trap\nagent:\n  command: [baseline-agent-trap]\nprotected: [private/**]\nworker_writable: [shared/**]\n"
	alpha := module("alpha", "ALP", "depends_on: [beta]\ninterface: []\nevals: []\ndisposable: []\npreserve: []\nappend_only: []\n")
	beta := module("beta", "BET", "hooks:\n  setup: 'opaque {module} command'\n")
	queued := change("beta", "set:\n  hooks:\n    eval: baseline-hook-trap\n") + "\n## Decisions\n\nPreserve queued decisions.\n\n## Notes\n\nRefer locally to new:fresh.\n"
	archive := strings.ReplaceAll(change("beta", "phase: applied\nrun: archived-run\n"), "new:fresh", "BET-002") + "\n## Notes\n\nHistorical prose may show new:undeclared-example.\n"
	files := map[string]string{"fireweed.yaml": cfg, "units/a/SPEC.md": alpha, "units/z/SPEC.md": beta, "changes/0002-zebra.md": queued, "changes/0002-alpha.md": queued, "changes/0000010-last.md": change("alpha", ""), "changes/archive/0009-old.md": archive, "changes/nested/0001-ignored.md": "private nested queue bytes must not be read", "changes/README.txt": "private nonmarkdown bytes must not be read", ".fireweed/hidden/SPEC.md": "private ignored module bytes must not be read", ".fireweed/playbooks/impl.md": "private playbook bytes must not be read", "units/z/internal/old.go": "private disposable oracle must not be read", "private/secret.txt": "private unrelated content must not be read", ".gitignore": "ignored\n", "ignored": "ignored outside snapshot"}
	for name, text := range files {
		put(t, root, name, text)
	}
	commit := baselineCommit(t, root)
	expected, err := p.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(expected.Diagnostics) != 0 || len(expected.Modules) != 2 || len(expected.Changes) != 3 || len(expected.Archive) != 1 {
		t.Fatalf("fixture discovery premise %+v", expected)
	}
	return baselineFixtureData{root, commit, expected, files}
}
func baselineFresh(t *testing.T, f *baselineFixtureData) {
	t.Helper()
	f.Commit = baselineCommit(t, f.Root)
	snapshot, err := p.Load(f.Root)
	if err != nil {
		t.Fatal(err)
	}
	f.Expected = snapshot
}
func baselineCorrective(message string) bool {
	if playbookCorrective(message) {
		return true
	}
	lower := strings.ToLower(message)
	for _, word := range []string{"fix ", "correct ", "make ", "move ", "update ", "declare ", "remove ", "resolve "} {
		if strings.Contains(lower, word) {
			return true
		}
	}
	return false
}
func baselineData(t *testing.T, got p.Snapshot) {
	t.Helper()
	if got.Modules == nil || got.Changes == nil || got.Archive == nil || got.Diagnostics == nil {
		t.Fatal("nil successful collections")
	}
	if !sort.SliceIsSorted(got.Diagnostics, func(i, j int) bool {
		a, b := got.Diagnostics[i], got.Diagnostics[j]
		return a.Path+"\x00"+a.Code+"\x00"+a.Severity+"\x00"+a.Message < b.Path+"\x00"+b.Code+"\x00"+b.Severity+"\x00"+b.Message
	}) {
		t.Fatal("unsorted diagnostics")
	}
	for _, d := range got.Diagnostics {
		if d.Path == "" || d.Message == "" || !baselineCorrective(d.Message) || (d.Severity != "error" && d.Severity != "warning") {
			t.Fatalf("unattributed/noncorrective diagnostic %+v", d)
		}
	}
}
func baselineAssert(t *testing.T, f baselineFixtureData, start, revision string) p.Snapshot {
	t.Helper()
	got, err := baselineCall(t, start, revision)
	if err != nil {
		t.Fatalf("LoadAt: %v", err)
	}
	baselineData(t, got)
	if len(f.Expected.Diagnostics) > 0 {
		baselineCategories(t, got, f.Expected)
		return got
	}
	if !reflect.DeepEqual(got, f.Expected) {
		t.Fatalf("selected committed data differs\ngot %+v\nwant %+v", got, f.Expected)
	}
	return got
}
func baselineCall(t *testing.T, start, revision string) (p.Snapshot, error) {
	t.Helper()
	type result struct {
		Snapshot p.Snapshot
		Err      error
	}
	ready := make(chan result, 1)
	go func() { value, err := p.LoadAt(start, revision); ready <- result{value, err} }()
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()
	select {
	case got := <-ready:
		return got.Snapshot, got.Err
	case <-timer.C:
		t.Fatal("LoadAt did not finish its local metadata read within the fixture deadline")
		return p.Snapshot{}, fmt.Errorf("fixture deadline")
	}
}
func baselineError(t *testing.T, start, revision, path string) {
	t.Helper()
	got, err := baselineCall(t, start, revision)
	if err == nil || !reflect.DeepEqual(got, p.Snapshot{}) || !strings.Contains(err.Error(), "C01") || !baselineCorrective(err.Error()) || (path != "" && !strings.Contains(err.Error(), path)) {
		t.Fatalf("want zero Snapshot + corrective C01 at %q: %+v %v", path, got, err)
	}
}
func baselineCategories(t *testing.T, got, want p.Snapshot) {
	t.Helper()
	baselineData(t, got)
	categories := func(ds []p.Diagnostic) []string {
		out := []string{}
		for _, d := range ds {
			out = append(out, d.Path+"\x00"+d.Code+"\x00"+d.Severity)
		}
		sort.Strings(out)
		return out
	}
	if !reflect.DeepEqual(categories(got.Diagnostics), categories(want.Diagnostics)) {
		t.Fatalf("categories got %+v want %+v", got.Diagnostics, want.Diagnostics)
	}
	got.Diagnostics = nil
	want.Diagnostics = nil
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("diagnosed data differs %+v want %+v", got, want)
	}
}
func Test_PRJ024_PRJ025CommittedRevisionSelection(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		t.Run(format, func(t *testing.T) {
			f := baselineFixture(t, format)
			baselineGit(t, f.Root, "branch", "approved", f.Commit)
			baselineGit(t, f.Root, "tag", "lightweight", f.Commit)
			baselineGit(t, f.Root, "tag", "-a", "annotated-λ", "-m", "snapshot", f.Commit)
			put(t, f.Root, "fireweed.yaml", "invalid newer config: [")
			newCommit := baselineCommit(t, f.Root)
			if newCommit == f.Commit {
				t.Fatal("two-commit premise")
			}
			for _, revision := range []string{f.Commit, "approved", "refs/heads/approved", "lightweight", "annotated-λ", "HEAD~1"} {
				t.Run(revision, func(t *testing.T) { baselineAssert(t, f, f.Root, revision) })
			}
		})
	}
}
func Test_PRJ024NestedLinkedAndLiteralReferences(t *testing.T) {
	f := baselineFixture(t, "sha1")
	name := "quoted'λ-tag"
	baselineGit(t, f.Root, "tag", name, f.Commit)
	nested := filepath.Join(f.Root, "nested ' λ")
	if err := os.Mkdir(nested, 0755); err != nil {
		t.Fatal(err)
	}
	baselineAssert(t, f, nested, name)
	peer := filepath.Join(t.TempDir(), "linked ' peer")
	baselineGit(t, f.Root, "worktree", "add", "-b", "peer", peer)
	peer = baselineCanonical(t, peer)
	put(t, peer, "fireweed.yaml", "invalid peer config: [")
	expected := f
	expected.Expected.Root = peer
	baselineAssert(t, expected, peer, f.Commit)
}
func Test_PRJ024InvalidSources(t *testing.T) {
	f := baselineFixture(t, "sha1")
	blob := baselineGit(t, f.Root, "rev-parse", "HEAD:fireweed.yaml")
	tree := baselineGit(t, f.Root, "rev-parse", "HEAD^{tree}")
	for _, revision := range []string{"", " ", "missing-ref", "--help", "--all", blob, tree, "HEAD:fireweed.yaml", "HEAD\x00bad"} {
		t.Run(revision, func(t *testing.T) { baselineError(t, f.Root, revision, "") })
	}
	t.Run("not-a-repository", func(t *testing.T) { baselineError(t, t.TempDir(), f.Commit, "") })
	t.Run("unborn", func(t *testing.T) {
		root := t.TempDir()
		baselineGit(t, root, "init", "-b", "main")
		baselineError(t, root, "HEAD", "")
	})
}
func Test_PRJ025ConfigurationFailuresAndFormats(t *testing.T) {
	for _, text := range []string{"", "version: [1]\n", "fireweed: '>=99'\n", "version: 99\n", "modules: ['../escape']\n", "module_defaults: {interface: ['[']}\n", "protected: ['/absolute']\n", "worker_writable: ['bad//suffix']\n"} {
		t.Run(fmt.Sprintf("%q", text), func(t *testing.T) {
			f := baselineFixture(t, "sha1")
			put(t, f.Root, "fireweed.yaml", text)
			commit := baselineCommit(t, f.Root)
			put(t, f.Root, "fireweed.yaml", f.Files["fireweed.yaml"])
			baselineError(t, f.Root, commit, "fireweed.yaml")
		})
	}
	t.Run("missing-config", func(t *testing.T) {
		f := baselineFixture(t, "sha1")
		if err := os.Remove(filepath.Join(f.Root, "fireweed.yaml")); err != nil {
			t.Fatal(err)
		}
		missing := baselineCommit(t, f.Root)
		put(t, f.Root, "fireweed.yaml", f.Files["fireweed.yaml"])
		baselineError(t, f.Root, missing, "fireweed.yaml")
	})
	for _, version := range []string{"0", "1"} {
		t.Run("format-"+version, func(t *testing.T) {
			f := baselineFixture(t, "sha1")
			put(t, f.Root, "fireweed.yaml", strings.Replace(f.Files["fireweed.yaml"], "version: 1", "version: "+version, 1)+"# selected format fixture "+version+"\n")
			baselineFresh(t, &f)
			baselineAssert(t, f, f.Root, f.Commit)
		})
	}
}
func Test_PRJ025EmptyAndRootSpecWarning(t *testing.T) {
	for _, kind := range []string{"empty", "root-spec"} {
		t.Run(kind, func(t *testing.T) {
			f := baselineFixture(t, "sha1")
			for _, dir := range []string{"units", "changes", ".fireweed"} {
				if err := os.RemoveAll(filepath.Join(f.Root, dir)); err != nil {
					t.Fatal(err)
				}
			}
			cfg := "version: 1\nbase: main\nmodules: []\n"
			if kind == "root-spec" {
				cfg = strings.Replace(config, "modules: [units/*]", "modules: ['**']", 1)
				put(t, f.Root, "SPEC.md", module("root", "ROT", ""))
			}
			put(t, f.Root, "fireweed.yaml", cfg)
			baselineFresh(t, &f)
			// Establish each fixture independently before the subject; a stub cannot mask a later premise.
			if len(f.Expected.Modules) != 0 || len(f.Expected.Changes) != 0 || len(f.Expected.Archive) != 0 {
				t.Fatalf("empty discovery premise %+v", f.Expected)
			}
			if kind == "root-spec" {
				requireDiag(t, f.Expected.Diagnostics, "C02", "warning", "SPEC.md")
			} else if len(f.Expected.Diagnostics) != 0 {
				t.Fatal("empty fixture diagnostics", f.Expected.Diagnostics)
			}
			baselineAssert(t, f, f.Root, f.Commit)
		})
	}
}
func Test_PRJ025MalformedDocumentsAndHandleRules(t *testing.T) {
	f := baselineFixture(t, "sha1")
	for name, text := range map[string]string{"outside/SPEC.md": "MUST NOT read out-of-glob bytes", "units/meta/SPEC.md": module("invalid name", "BAD", ""), "units/grammar/SPEC.md": strings.Replace(module("grammar", "GRA", ""), "MUST", "might", 1), "changes/0004-bad.md": "invalid", "changes/0005-unknown.md": change("missing", ""), "changes/0006-handles.md": strings.Replace(change("beta", ""), "new:fresh", "new:foreign", 1) + "\n## Notes\n\nOnly new:fresh is referenced here.\n", "changes/archive/0007-zebra.md": change("beta", ""), "changes/archive/0008-bad-live.md": strings.Replace(strings.ReplaceAll(change("beta", "phase: applied\nrun: archived-bad\n"), "new:fresh", "BET-004"), "The fixture MUST work.", "The fixture MUST refer to new:missing.", 1)} {
		put(t, f.Root, name, text)
	}
	baselineFresh(t, &f)
	for _, entry := range []struct{ code, severity, path string }{{"C02", "warning", "outside/SPEC.md"}, {"C03", "error", "units/meta/SPEC.md"}, {"C04", "error", "units/grammar/SPEC.md"}, {"C05", "error", "changes/0004-bad.md"}, {"C05", "error", "changes/0005-unknown.md"}, {"C09", "error", "changes/0006-handles.md"}, {"C05", "error", "changes/archive/0007-zebra.md"}, {"C09", "error", "changes/archive/0008-bad-live.md"}} {
		requireDiag(t, f.Expected.Diagnostics, entry.code, entry.severity, entry.path)
	}
	baselineData(t, f.Expected)
	got, err := baselineCall(t, f.Root, f.Commit)
	if err != nil {
		t.Fatal(err)
	}
	baselineCategories(t, got, f.Expected)
	for _, entry := range []struct{ code, severity, path string }{{"C02", "warning", "outside/SPEC.md"}, {"C03", "error", "units/meta/SPEC.md"}, {"C04", "error", "units/grammar/SPEC.md"}, {"C05", "error", "changes/0004-bad.md"}, {"C05", "error", "changes/0005-unknown.md"}, {"C09", "error", "changes/0006-handles.md"}, {"C05", "error", "changes/archive/0007-zebra.md"}, {"C09", "error", "changes/archive/0008-bad-live.md"}} {
		requireDiag(t, got.Diagnostics, entry.code, entry.severity, entry.path)
	}
	for _, ch := range got.Archive {
		if ch.Slug == "old" && !strings.Contains(ch.Document.Sections[len(ch.Document.Sections)-1].Body, "new:undeclared-example") {
			t.Fatal("historical prose lost")
		}
	}
}
func Test_PRJ025UnusualDirectoryPaths(t *testing.T) {
	f := baselineFixture(t, "sha1")
	old := "units/z"
	name := "units/beta ' λ\t\nname"
	if err := os.Rename(filepath.Join(f.Root, old), filepath.Join(f.Root, name)); err != nil {
		t.Fatal(err)
	}
	baselineFresh(t, &f)
	baselineAssert(t, f, f.Root, f.Commit)
}
func Test_PRJ026ContainedCommittedLinks(t *testing.T) {
	for _, kind := range []string{"relative-config", "absolute-config", "chain-config", "changed-current-target", "retargeted-current-link", "module", "module-parent-chain", "queue", "queue-directory", "archive-directory", "executable"} {
		t.Run(kind, func(t *testing.T) {
			f := baselineFixture(t, "sha1")
			selected := "fireweed.yaml"
			target := "metadata/config ' λ\n.yaml"
			text := f.Files[selected]
			switch kind {
			case "module", "module-parent-chain", "executable":
				selected = "units/a/SPEC.md"
				target = "metadata/alpha.md"
				text = f.Files[selected]
			case "queue":
				selected = "changes/0002-alpha.md"
				target = "metadata/change.md"
				text = f.Files[selected]
			}
			if kind == "queue-directory" || kind == "archive-directory" {
				selected = "changes"
				target = "metadata/queue"
				if kind == "archive-directory" {
					selected = "changes/archive"
					target = "metadata/archive"
				}
				if err := os.MkdirAll(filepath.Dir(filepath.Join(f.Root, target)), 0755); err != nil {
					t.Fatal(err)
				}
				if err := os.Rename(filepath.Join(f.Root, selected), filepath.Join(f.Root, target)); err != nil {
					t.Fatal(err)
				}
				relative, err := filepath.Rel(filepath.Dir(filepath.Join(f.Root, selected)), filepath.Join(f.Root, target))
				if err != nil {
					t.Fatal(err)
				}
				if err = os.Symlink(relative, filepath.Join(f.Root, selected)); err != nil {
					t.Fatal(err)
				}
			} else if kind == "executable" {
				if err := os.Chmod(filepath.Join(f.Root, selected), 0755); err != nil {
					t.Fatal(err)
				}
			} else {
				put(t, f.Root, target, text)
				if err := os.Remove(filepath.Join(f.Root, selected)); err != nil {
					t.Fatal(err)
				}
				destination := filepath.Join(f.Root, target)
				if kind == "chain-config" {
					link := filepath.Join(f.Root, "metadata/second-link")
					if err := os.Symlink(filepath.Base(target), link); err != nil {
						t.Fatal(err)
					}
					destination = link
				}
				if kind == "module-parent-chain" {
					link := filepath.Join(f.Root, "metadata/alias")
					if err := os.Symlink("actual", link); err != nil {
						t.Fatal(err)
					}
					put(t, f.Root, "metadata/actual/module.txt", text)
					destination = filepath.Join(link, "module.txt")
				}
				if kind != "absolute-config" {
					relative, err := filepath.Rel(filepath.Dir(filepath.Join(f.Root, selected)), destination)
					if err != nil {
						t.Fatal(err)
					}
					destination = relative
				}
				if err := os.Symlink(destination, filepath.Join(f.Root, selected)); err != nil {
					t.Fatal(err)
				}
			}
			baselineFresh(t, &f)
			for _, d := range f.Expected.Diagnostics {
				if d.Severity == "error" {
					t.Fatalf("contained link fixture %+v", d)
				}
			}
			if kind == "changed-current-target" {
				put(t, f.Root, target, "invalid current target: [")
			}
			if kind == "retargeted-current-link" {
				outside := baselineCanonical(t, t.TempDir())
				put(t, outside, "private", "invalid current outside target: [")
				if err := os.Remove(filepath.Join(f.Root, selected)); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(filepath.Join(outside, "private"), filepath.Join(f.Root, selected)); err != nil {
					t.Fatal(err)
				}
			}
			baselineAssert(t, f, f.Root, f.Commit)
		})
	}
}
func Test_PRJ026RejectedCommittedLinksAndGitlinks(t *testing.T) {
	for _, selected := range []string{"fireweed.yaml", "units/a/SPEC.md", "changes/0002-alpha.md"} {
		for _, kind := range []string{"escape", "absolute-outside", "prefix-sibling", "dangling", "loop", "directory", "nul", "backslash", "gitlink"} {
			t.Run(selected+"/"+kind, func(t *testing.T) {
				f := baselineFixture(t, "sha1")
				if kind == "gitlink" {
					baselineGit(t, f.Root, "update-index", "--add", "--cacheinfo", "160000,"+f.Commit+","+selected)
					baselineGit(t, f.Root, "commit", "-m", "gitlink metadata")
					f.Commit = baselineGit(t, f.Root, "rev-parse", "HEAD")
				} else {
					if err := os.Remove(filepath.Join(f.Root, selected)); err != nil {
						t.Fatal(err)
					}
					target := "../../../../outside/private"
					switch kind {
					case "absolute-outside":
						outside := baselineCanonical(t, t.TempDir())
						put(t, outside, "private", "outside unrelated bytes")
						target = filepath.Join(outside, "private")
					case "prefix-sibling":
						// Both potential mistaken interpretations contain valid metadata.
						// Rejection must follow containment, not a missing/invalid donor.
						outside := f.Root + "-sibling"
						put(t, outside, "private", f.Files[selected])
						put(t, f.Root, "-sibling/private", f.Files[selected])
						target = filepath.Join(outside, "private")
					case "dangling":
						target = "missing"
					case "loop":
						target = filepath.Base(selected)
					case "directory":
						target = "."
					case "backslash":
						target = "bad\\target"
					}
					if kind == "nul" {
						id, err := baselineInput(t, f.Root, "bad\x00target", "hash-object", "-w", "--stdin")
						if err != nil {
							t.Fatal(err, id)
						}
						baselineGit(t, f.Root, "update-index", "--add", "--cacheinfo", "120000,"+strings.TrimSpace(id)+","+selected)
						baselineGit(t, f.Root, "commit", "-m", "invalid link bytes")
						f.Commit = baselineGit(t, f.Root, "rev-parse", "HEAD")
					} else {
						if err := os.Symlink(target, filepath.Join(f.Root, selected)); err != nil {
							t.Fatal(err)
						}
						f.Commit = baselineCommit(t, f.Root)
					}
				}
				// Correct current metadata so no working-file read can explain the rejection.
				full := filepath.Join(f.Root, selected)
				if info, err := os.Lstat(full); err == nil {
					if info.IsDir() {
						_ = os.RemoveAll(full)
					} else {
						_ = os.Remove(full)
					}
				}
				put(t, f.Root, selected, f.Files[selected])
				if selected == "fireweed.yaml" {
					baselineError(t, f.Root, f.Commit, selected)
				} else {
					got, err := baselineCall(t, f.Root, f.Commit)
					if err != nil {
						t.Fatal(err)
					}
					code := "C03"
					if strings.HasPrefix(selected, "changes/") {
						code = "C05"
					}
					requireDiag(t, got.Diagnostics, code, "error", selected)
					baselineData(t, got)
				}
			})
		}
	}
}
func Test_PRJ026DirectoryAliasesNotDiscovered(t *testing.T) {
	f := baselineFixture(t, "sha1")
	if err := os.Symlink("a", filepath.Join(f.Root, "units/alias")); err != nil {
		t.Fatal(err)
	}
	baselineFresh(t, &f)
	baselineAssert(t, f, f.Root, f.Commit)
	if len(f.Expected.Modules) != 2 {
		t.Fatal("directory alias premise")
	}
}
func Test_PRJ027CurrentFilesNeverReplaceBaseline(t *testing.T) {
	for _, mode := range []string{"dirty", "staged", "deleted", "untracked", "unreadable", "outside-links"} {
		t.Run(mode, func(t *testing.T) {
			f := baselineFixture(t, "sha1")
			switch mode {
			case "dirty", "staged":
				put(t, f.Root, "fireweed.yaml", "invalid current config: [")
				put(t, f.Root, "units/a/SPEC.md", module("changed", "CHG", ""))
				put(t, f.Root, "changes/0002-alpha.md", "invalid changed document")
				if mode == "staged" {
					baselineGit(t, f.Root, "add", ".")
				}
			case "deleted":
				for _, name := range []string{"fireweed.yaml", "units/a/SPEC.md", "changes/0002-alpha.md"} {
					if err := os.Remove(filepath.Join(f.Root, name)); err != nil {
						t.Fatal(err)
					}
				}
			case "untracked":
				put(t, f.Root, "units/new/SPEC.md", module("new", "NEW", ""))
				put(t, f.Root, "changes/0001-untracked.md", change("alpha", ""))
			case "unreadable":
				for _, name := range []string{"fireweed.yaml", "units/a/SPEC.md", "changes/0002-alpha.md"} {
					full := filepath.Join(f.Root, name)
					if err := os.Chmod(full, 0000); err != nil {
						t.Fatal(err)
					}
					t.Cleanup(func() { _ = os.Chmod(full, 0600) })
				}
			case "outside-links":
				outside := baselineCanonical(t, t.TempDir())
				put(t, outside, "private", "private current data")
				for _, name := range []string{"fireweed.yaml", "units/a/SPEC.md", "changes/0002-alpha.md"} {
					full := filepath.Join(f.Root, name)
					if err := os.Remove(full); err != nil {
						t.Fatal(err)
					}
					if err := os.Symlink(filepath.Join(outside, "private"), full); err != nil {
						t.Fatal(err)
					}
				}
			}
			baselineAssert(t, f, f.Root, f.Commit)
		})
	}
}
func Test_PRJ027ExistingResultsAndEmptyStorage(t *testing.T) {
	f := baselineFixture(t, "sha1")
	one := baselineAssert(t, f, f.Root, f.Commit)
	two := baselineAssert(t, f, f.Root, f.Commit)
	before, _ := json.Marshal(two)
	input, _ := json.Marshal(f.Expected)
	one.Config.Modules[0] = "mutated"
	one.Config.ModuleDefaults.Interface[0] = "mutated"
	one.Modules[1].Document.Metadata["name"] = "mutated"
	one.Modules[1].Document.Metadata["hooks"].(map[string]any)["setup"] = "mutated"
	one.Modules[1].Specs[0].Body = "mutated"
	one.Changes[0].Document.Metadata["module"] = "mutated"
	one.Changes[0].Added[0].Body = "mutated"
	one.Changes[0].Set["hooks"].(map[string]any)["eval"] = "mutated"
	one.Archive[0].Document.Sections[0].Body = "mutated"
	now, _ := json.Marshal(two)
	expected, _ := json.Marshal(f.Expected)
	if string(now) != string(before) || string(expected) != string(input) {
		t.Fatal("result aliases an existing result or fixture expectation")
	}
	baselineAssert(t, f, f.Root, f.Commit)
	put(t, f.Root, "fireweed.yaml", "version: 1\nbase: main\nmodules: []\n")
	_ = os.RemoveAll(filepath.Join(f.Root, "units"))
	_ = os.RemoveAll(filepath.Join(f.Root, "changes"))
	_ = os.RemoveAll(filepath.Join(f.Root, ".fireweed"))
	baselineFresh(t, &f)
	a := baselineAssert(t, f, f.Root, f.Commit)
	b := baselineAssert(t, f, f.Root, f.Commit)
	a.Modules = append(a.Modules, c.Module{Name: "first"})
	b.Modules = append(b.Modules, c.Module{Name: "second"})
	a.Changes = append(a.Changes, c.Change{Slug: "first"})
	b.Changes = append(b.Changes, c.Change{Slug: "second"})
	a.Archive = append(a.Archive, c.Change{Slug: "first"})
	b.Archive = append(b.Archive, c.Change{Slug: "second"})
	a.Diagnostics = append(a.Diagnostics, p.Diagnostic{Code: "first"})
	b.Diagnostics = append(b.Diagnostics, p.Diagnostic{Code: "second"})
	if a.Modules[0].Name != "first" || a.Changes[0].Slug != "first" || a.Archive[0].Slug != "first" || a.Diagnostics[0].Code != "first" {
		t.Fatal("empty result backing storage aliases")
	}
}
func Test_PRJ024_PRJ025_PRJ026_PRJ027BaselineFixtureControls(t *testing.T) {
	f := baselineFixture(t, "sha1")
	if _, ok := f.Expected.Modules[1].Document.Metadata["hooks"].(map[string]any); !ok {
		t.Fatal("nested module metadata mutation premise")
	}
	if _, ok := f.Expected.Changes[0].Set["hooks"].(map[string]any); !ok {
		t.Fatal("nested change metadata mutation premise")
	}

	if f.Expected.Config.Version != 1 || len(f.Expected.Modules[0].Interface) != 0 || f.Expected.Modules[0].Interface == nil || !reflect.DeepEqual(f.Expected.Modules[1].Interface, []string{"api.go"}) || f.Expected.Changes[0].Slug != "alpha" || f.Expected.Changes[1].Slug != "zebra" || f.Expected.Changes[2].Slug != "last" {
		t.Fatal("fixture defaults/order premise")
	}
	raw, err := baselineInput(t, f.Root, "", "cat-file", "blob", f.Commit+":fireweed.yaml")
	if err != nil || raw != f.Files["fireweed.yaml"] {
		t.Fatal("Git blob fixture premise", err)
	}
	if len(f.Expected.Archive) != 1 || !strings.Contains(f.Expected.Archive[0].Document.Sections[len(f.Expected.Archive[0].Document.Sections)-1].Body, "new:undeclared-example") {
		t.Fatal("archive human-text premise")
	}
}

func Test_PRJ025ValidationRemainsExplicit(t *testing.T) {
	f := baselineFixture(t, "sha1")
	put(t, f.Root, "units/a/SPEC.md", module("alpha", "ALP", "depends_on: [absent]\n"))
	put(t, f.Root, "units/z/SPEC.md", module("beta", "BET", "evals: [internal/**]\nhooks:\n  eval: ''\n"))
	baselineFresh(t, &f)
	if len(f.Expected.Modules) != 2 || len(f.Expected.Diagnostics) != 0 {
		t.Fatal("static error fixture must remain discoverable")
	}
	ds := p.Check(f.Expected)
	requireDiag(t, ds, "C06", "error", "units/a/SPEC.md")
	requireDiag(t, ds, "C06", "error", "units/z/SPEC.md")
	requireDiag(t, ds, "C07", "error", "units/z/SPEC.md")
	baselineAssert(t, f, f.Root, f.Commit)
}

func baselineTree(t *testing.T, root, entries string) string {
	t.Helper()
	out, err := baselineInput(t, root, entries, "mktree", "-z")
	if err != nil {
		t.Fatalf("tree fixture: %v %s", err, out)
	}
	return strings.TrimSpace(out)
}
func baselineReplaceTreeEntry(t *testing.T, root, tree, name, mode, kind, object string) string {
	t.Helper()
	raw := baselineGit(t, root, "ls-tree", "-z", tree)
	entries := []string{}
	for _, entry := range strings.Split(raw, "\x00") {
		if entry == "" {
			continue
		}
		parts := strings.SplitN(entry, "\t", 2)
		if len(parts) != 2 {
			t.Fatal("tree entry premise")
		}
		if parts[1] != name {
			entries = append(entries, entry)
		}
	}
	entries = append(entries, mode+" "+kind+" "+object+"\t"+name)
	return baselineTree(t, root, strings.Join(entries, "\x00")+"\x00")
}
func baselineHiddenTree(t *testing.T, f *baselineFixtureData, mode string) {
	t.Helper()
	private, err := baselineInput(t, f.Root, "unique ignored historical SPEC content "+mode, "hash-object", "-w", "--stdin")
	if err != nil {
		t.Fatal(err)
	}
	hidden := baselineTree(t, f.Root, "100644 blob "+strings.TrimSpace(private)+"\tSPEC.md\x00")
	rootTree := baselineGit(t, f.Root, "rev-parse", f.Commit+"^{tree}")
	switch mode {
	case "root-dotgit":
		rootTree = baselineReplaceTreeEntry(t, f.Root, rootTree, ".git", "040000", "tree", hidden)
	default:
		units := baselineGit(t, f.Root, "rev-parse", f.Commit+":units")
		if mode == "gitlink-module" {
			units = baselineReplaceTreeEntry(t, f.Root, units, "foreign", "160000", "commit", f.Commit)
		} else {
			alpha := baselineGit(t, f.Root, "rev-parse", f.Commit+":units/a")
			name := ".git"
			if mode == "nested-fireweed" {
				name = ".fireweed"
			}
			alpha = baselineReplaceTreeEntry(t, f.Root, alpha, name, "040000", "tree", hidden)
			units = baselineReplaceTreeEntry(t, f.Root, units, "a", "040000", "tree", alpha)
		}
		rootTree = baselineReplaceTreeEntry(t, f.Root, rootTree, "units", "040000", "tree", units)
	}
	f.Commit = baselineGit(t, f.Root, "commit-tree", rootTree, "-p", f.Commit, "-m", "uncheckoutable ignored tree fixture")
}
func Test_PRJ025_PRJ026IgnoredGitTreesAndGitlinks(t *testing.T) {
	for _, mode := range []string{"root-dotgit", "nested-dotgit", "nested-fireweed", "gitlink-module"} {
		t.Run(mode, func(t *testing.T) {
			f := baselineFixture(t, "sha1")
			baselineHiddenTree(t, &f, mode)
			baselineAssert(t, f, f.Root, f.Commit)
		})
	}
}
func Test_PRJ025_PRJ026HiddenTreeFixtureControl(t *testing.T) {
	f := baselineFixture(t, "sha1")
	old := f.Commit
	baselineHiddenTree(t, &f, "root-dotgit")
	if f.Commit == old {
		t.Fatal("distinct hidden tree commit premise")
	}
	tree := baselineGit(t, f.Root, "ls-tree", "-r", "-z", f.Commit)
	if !strings.Contains(tree, "\t.git/SPEC.md\x00") {
		t.Fatal("ignored tree entry not present")
	}
	current, err := p.Load(f.Root)
	if err != nil || !reflect.DeepEqual(current, f.Expected) {
		t.Fatal("object-only tree fixture changed working discovery")
	}
}

func Test_PRJ025SelectedConfigResolvesLocalBase(t *testing.T) {
	for _, mode := range []string{"remote-default", "tracking-over-local", "local-fallback", "missing-local"} {
		t.Run(mode, func(t *testing.T) {
			f := baselineFixture(t, "sha1")
			cfg := strings.Replace(f.Files["fireweed.yaml"], "base: main\n", "remote: offline\n", 1)
			branch := "release/approved"
			want := "refs/heads/" + branch
			if mode != "remote-default" {
				cfg = "base: " + branch + "\n" + cfg
			}
			if mode == "remote-default" || mode == "tracking-over-local" {
				baselineGit(t, f.Root, "update-ref", "refs/remotes/offline/"+branch, f.Commit)
				want = "refs/remotes/offline/" + branch
			}
			if mode == "remote-default" {
				baselineGit(t, f.Root, "symbolic-ref", "refs/remotes/offline/HEAD", want)
			}
			if mode == "tracking-over-local" || mode == "local-fallback" {
				baselineGit(t, f.Root, "branch", branch, f.Commit)
			}
			put(t, f.Root, "fireweed.yaml", cfg)
			baselineFresh(t, &f)
			if f.Expected.BaseRef != want || f.Expected.Config.Remote != "offline" {
				t.Fatalf("selected BaseRef premise %+v", f.Expected)
			}
			// A newer committed and current config must not determine the selected baseline's base.
			put(t, f.Root, "fireweed.yaml", f.Files["fireweed.yaml"])
			baselineCommit(t, f.Root)
			current, err := p.Load(f.Root)
			if err != nil || current.BaseRef == want {
				t.Fatalf("distinguishing current base premise %+v %v", current, err)
			}
			baselineAssert(t, f, f.Root, f.Commit)
		})
	}
}
