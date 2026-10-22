package evals

import (
	g "github.com/lurkshark/fireweed/packages/gitstore"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func git(t *testing.T, root string, args ...string) string {
	t.Helper()
	argv := append([]string{"-c", "user.name=Fireweed Eval", "-c", "user.email=eval@example.invalid", "-c", "core.autocrlf=false"}, args...)
	cmd := exec.Command("git", argv...)
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_CONFIG_GLOBAL="+os.DevNull)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v: %s", args, err, output)
	}
	return strings.TrimSpace(string(output))
}
func fixture(t *testing.T) string {
	t.Helper()
	root := filepath.Join(t.TempDir(), "repo with spaces")
	if err := os.Mkdir(root, 0755); err != nil {
		t.Fatal(err)
	}
	root, err := filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatal(err)
	}
	git(t, root, "init", "--initial-branch=main")
	write(t, root, "tracked.txt", "original\n")
	git(t, root, "add", ".")
	git(t, root, "commit", "-m", "fixture baseline")
	return root
}
func write(t *testing.T, root, path, text string) {
	t.Helper()
	absolute := filepath.Join(root, path)
	if err := os.MkdirAll(filepath.Dir(absolute), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(absolute, []byte(text), 0644); err != nil {
		t.Fatal(err)
	}
}
func Test_GIT001DiscoveryAndWorktrees(t *testing.T) {
	root := fixture(t)
	nested := filepath.Join(root, "a", "b")
	if err := os.MkdirAll(nested, 0755); err != nil {
		t.Fatal(err)
	}
	refs := git(t, root, "show-ref")
	head := git(t, root, "rev-parse", "HEAD")
	git(t, root, "remote", "add", "origin", "file:///fireweed-must-not-contact-remote")
	got, err := g.Open(nested)
	if err != nil {
		t.Fatal(err)
	}
	if got.Root != root || got.GitDir != filepath.Join(root, ".git") || got.CommonDir != got.GitDir {
		t.Fatalf("repository %+v", got)
	}
	fromRoot, err := g.Open(root)
	if err != nil || fromRoot != got {
		t.Errorf("root discovery %+v %v", fromRoot, err)
	}
	if git(t, root, "show-ref") != refs || git(t, root, "rev-parse", "HEAD") != head {
		t.Error("discovery changed refs")
	}
	scratch := filepath.Join(t.TempDir(), "linked checkout")
	git(t, root, "worktree", "add", "-b", "fixture-linked", scratch)
	defer git(t, root, "worktree", "remove", "--force", scratch)
	linked, err := g.Open(scratch)
	if err != nil {
		t.Fatal(err)
	}
	canonical, err := filepath.EvalSymlinks(scratch)
	if err != nil {
		t.Fatal(err)
	}
	if linked.Root != canonical || linked.Root == got.Root || linked.GitDir == got.GitDir || linked.CommonDir != got.CommonDir || !filepath.IsAbs(linked.GitDir) {
		t.Errorf("linked %+v vs %+v", linked, got)
	}
	if _, err := g.Open(t.TempDir()); err == nil {
		t.Error("non-repository accepted")
	}
}
func Test_GIT002LocalBaseResolution(t *testing.T) {
	root := fixture(t)
	head := git(t, root, "rev-parse", "HEAD")
	git(t, root, "remote", "add", "origin", "file:///fireweed-unreachable")
	if got, err := g.BaseRef(root, "", ""); err != nil || got != "refs/heads/main" {
		t.Errorf("fallback %q %v", got, err)
	}
	git(t, root, "update-ref", "refs/remotes/origin/main", head)
	if got, err := g.BaseRef(root, "main", ""); err != nil || got != "refs/remotes/origin/main" {
		t.Errorf("remote preferred %q %v", got, err)
	}
	git(t, root, "update-ref", "refs/remotes/upstream/team/trunk", head)
	git(t, root, "symbolic-ref", "refs/remotes/upstream/HEAD", "refs/remotes/upstream/team/trunk")
	refs := git(t, root, "show-ref")
	if got, err := g.BaseRef(root, "", "upstream"); err != nil || got != "refs/remotes/upstream/team/trunk" {
		t.Errorf("remote default %q %v", got, err)
	}
	if got, err := g.BaseRef(root, "team/local", "upstream"); err != nil || got != "refs/heads/team/local" {
		t.Errorf("explicit local fallback %q %v", got, err)
	}
	if git(t, root, "show-ref") != refs {
		t.Error("base inspection changed refs")
	}
	for _, tc := range [][2]string{{"--help", "origin"}, {"bad..branch", "origin"}, {"main", "../remote"}} {
		if _, err := g.BaseRef(root, tc[0], tc[1]); err == nil {
			t.Errorf("invalid base/remote accepted: %v", tc)
		}
	}
}
func Test_GIT003CleanAndDiff(t *testing.T) {
	root := fixture(t)
	for _, path := range []string{"deleted.txt", "rename old.txt", "staged.txt", "unstaged.txt"} {
		write(t, root, path, "original "+path+"\n")
	}
	write(t, root, ".gitignore", "ignored*\n")
	git(t, root, "add", ".")
	git(t, root, "commit", "-m", "inspection fixture")
	base := git(t, root, "rev-parse", "HEAD")
	if clean, err := g.Clean(root); err != nil || !clean {
		t.Fatalf("initial clean=%v %v", clean, err)
	}
	write(t, root, "ignored-output", "ignored")
	if clean, err := g.Clean(root); err != nil || !clean {
		t.Errorf("ignored file dirtied checkout: %v %v", clean, err)
	}
	write(t, root, "untracked only.txt", "new")
	if clean, err := g.Clean(root); err != nil || clean {
		t.Errorf("untracked not detected: %v %v", clean, err)
	}
	if err := os.Remove(filepath.Join(root, "untracked only.txt")); err != nil {
		t.Fatal(err)
	}
	write(t, root, "staged.txt", "staged edit\n")
	git(t, root, "add", "staged.txt")
	if clean, err := g.Clean(root); err != nil || clean {
		t.Errorf("staged not detected: %v %v", clean, err)
	}
	write(t, root, "unstaged.txt", "unstaged edit\n")
	if err := os.Remove(filepath.Join(root, "deleted.txt")); err != nil {
		t.Fatal(err)
	}
	git(t, root, "mv", "rename old.txt", "rename new.txt")
	write(t, root, "staged new.txt", "staged addition\n")
	git(t, root, "add", "staged new.txt")
	strange := "untracked\twith\nnewline.txt"
	write(t, root, strange, "new\n")
	want := []g.Edit{{Status: "D", Path: "deleted.txt"}, {Status: "R", Path: "rename new.txt", OldPath: "rename old.txt"}, {Status: "A", Path: "staged new.txt"}, {Status: "M", Path: "staged.txt"}, {Status: "M", Path: "unstaged.txt"}, {Status: "A", Path: strange}}
	sort.Slice(want, func(i, j int) bool { return want[i].Path < want[j].Path })
	got, err := g.Edits(root, base)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("edits %+v, want %+v", got, want)
	}
	again, err := g.Edits(root, base)
	if err != nil || !reflect.DeepEqual(got, again) {
		t.Errorf("nondeterministic %+v %v", again, err)
	}
	if clean, err := g.Clean(root); err != nil || clean {
		t.Errorf("dirty checkout: %v %v", clean, err)
	}
	for _, revision := range []string{"missing-ref", "--help", ""} {
		if _, err := g.Edits(root, revision); err == nil {
			t.Errorf("invalid revision accepted %q", revision)
		}
	}
	if _, err := g.Edits(t.TempDir(), base); err == nil {
		t.Error("invalid diff root accepted")
	}
	if _, err := g.Clean(t.TempDir()); err == nil {
		t.Error("invalid status root accepted")
	}
}
func Test_GIT004SafePaths(t *testing.T) {
	root := fixture(t)
	outside := root + "-sibling"
	if err := os.MkdirAll(outside, 0755); err != nil {
		t.Fatal(err)
	}
	write(t, root, "inside/file.txt", "inside")
	write(t, outside, "secret.txt", "secret")
	for link, target := range map[string]string{"internal-link": filepath.Join(root, "inside"), "escape": outside, "dangling": filepath.Join(outside, "does-not-exist")} {
		if err := os.Symlink(target, filepath.Join(root, link)); err != nil {
			t.Fatal(err)
		}
	}
	for _, relative := range []string{"inside/file.txt", "new/deep/file.txt", "internal-link/file.txt", "internal-link/new/file.txt"} {
		got, err := g.SafePath(root, relative)
		want := filepath.Join(root, filepath.FromSlash(relative))
		if err != nil || got != want {
			t.Errorf("safe path %q=%q %v, want %q", relative, got, err, want)
		}
	}
	for _, relative := range []string{"", "/absolute", "C:/absolute", "../escape", "inside/../file.txt", "./file.txt", "inside//file.txt", "inside/", "inside\\file.txt", "nul\x00file", "escape/secret.txt", "escape/new/deep.txt", "escape", "dangling/file.txt"} {
		if _, err := g.SafePath(root, relative); err == nil {
			t.Errorf("unsafe path accepted %q", relative)
		}
	}
	if _, err := g.SafePath(filepath.Join(root, "missing-root"), "file"); err == nil {
		t.Error("missing root accepted")
	}
	if _, err := os.Stat(filepath.Join(root, "new")); !os.IsNotExist(err) {
		t.Errorf("safe path created files: %v", err)
	}
	if data, err := os.ReadFile(filepath.Join(outside, "secret.txt")); err != nil || string(data) != "secret" {
		t.Errorf("outside file modified: %q %v", data, err)
	}
}
