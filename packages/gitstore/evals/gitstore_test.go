package evals_test

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
)

func git(t *testing.T, dir string, args ...string) string {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", append([]string{"-C", dir}, args...)...)
	cmd.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_CONFIG_GLOBAL=/dev/null", "GIT_NO_LAZY_FETCH=1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v: %s", args, err, out)
	}
	return strings.TrimSpace(string(out))
}

func put(t *testing.T, root, name string, data []byte) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		t.Fatal(err)
	}
}

func repo(t *testing.T) string {
	t.Helper()
	root := filepath.Join(t.TempDir(), "repo")
	if err := os.Mkdir(root, 0755); err != nil {
		t.Fatal(err)
	}
	var err error
	root, err = filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatal(err)
	}
	git(t, root, "init", "-q", "-b", "main", "--object-format=sha1")
	git(t, root, "config", "user.name", "Fixture")
	git(t, root, "config", "user.email", "fixture@example.invalid")
	put(t, root, "seed.txt", []byte("seed\n"))
	git(t, root, "add", ".")
	git(t, root, "commit", "-qm", "seed")
	return root
}

func head(t *testing.T, root string) string { return git(t, root, "rev-parse", "HEAD") }

func Test_GIT_001_InspectWorktreesAndDetachedHead(t *testing.T) {
	root := repo(t)
	inside := filepath.Join(root, "nested")
	if err := os.Mkdir(inside, 0755); err != nil {
		t.Fatal(err)
	}
	primary, err := gitstore.Inspect(inside)
	if err != nil {
		t.Fatal(err)
	}
	common, err := filepath.EvalSymlinks(filepath.Join(root, ".git"))
	if err != nil {
		t.Fatal(err)
	}
	if primary.Root != root || primary.CommonDir != common || primary.Head != head(t, root) || primary.Branch != "main" {
		t.Errorf("primary = %#v", primary)
	}
	linkedRoot := filepath.Join(t.TempDir(), "linked")
	git(t, root, "worktree", "add", "-qb", "linked", linkedRoot)
	linkedRoot, err = filepath.EvalSymlinks(linkedRoot)
	if err != nil {
		t.Fatal(err)
	}
	linked, err := gitstore.Inspect(linkedRoot)
	if err != nil {
		t.Fatal(err)
	}
	if linked.Root != linkedRoot || linked.CommonDir != common || linked.Head != head(t, root) || linked.Branch != "linked" {
		t.Errorf("linked = %#v", linked)
	}
	git(t, linkedRoot, "checkout", "--detach", "-q")
	detached, err := gitstore.Inspect(linkedRoot)
	if err != nil {
		t.Fatal(err)
	}
	if detached.Root != linkedRoot || detached.CommonDir != common || detached.Head != head(t, root) || detached.Branch != "" {
		t.Errorf("detached = %#v", detached)
	}
	if _, err := gitstore.Inspect(t.TempDir()); err == nil {
		t.Error("Inspect(non-repository) succeeded")
	}
}

func Test_GIT_002_BaseRefPreferenceDefaultsAndMissing(t *testing.T) {
	root := repo(t)
	if got, err := gitstore.BaseRef(root, "", ""); err != nil || got != "refs/heads/main" {
		t.Errorf("default main = %q, %v", got, err)
	}
	git(t, root, "branch", "release")
	if got, err := gitstore.BaseRef(root, "upstream", "release"); err != nil || got != "refs/heads/release" {
		t.Errorf("local fallback = %q, %v", got, err)
	}
	if _, err := gitstore.BaseRef(root, "upstream", "missing"); err == nil {
		t.Error("missing base succeeded")
	}
	oid := head(t, root)
	git(t, root, "update-ref", "refs/remotes/origin/main", oid)
	if got, err := gitstore.BaseRef(root, "", "main"); err != nil || got != "refs/remotes/origin/main" {
		t.Errorf("origin preference = %q, %v", got, err)
	}
	git(t, root, "update-ref", "refs/remotes/upstream/release", oid)
	if got, err := gitstore.BaseRef(root, "upstream", "release"); err != nil || got != "refs/remotes/upstream/release" {
		t.Errorf("named remote preference = %q, %v", got, err)
	}
	git(t, root, "update-ref", "refs/remotes/origin/develop", oid)
	git(t, root, "symbolic-ref", "refs/remotes/origin/HEAD", "refs/remotes/origin/develop")
	if got, err := gitstore.BaseRef(root, "", ""); err != nil || got != "refs/remotes/origin/develop" {
		t.Errorf("remote HEAD default = %q, %v", got, err)
	}
	git(t, root, "symbolic-ref", "refs/remotes/upstream/HEAD", "refs/remotes/upstream/release")
	if got, err := gitstore.BaseRef(root, "upstream", ""); err != nil || got != "refs/remotes/upstream/release" {
		t.Errorf("named remote HEAD default = %q, %v", got, err)
	}
}

func Test_GIT_003_GIT_006_GIT_007_ChangesReadAtAndSelectedCommit(t *testing.T) {
	root := repo(t)
	for _, name := range []string{"selected.txt", "staged.txt", "unstaged.txt", "old-name.txt"} {
		put(t, root, name, []byte("original\n"))
	}
	git(t, root, "add", ".")
	git(t, root, "commit", "-qm", "files")
	base := head(t, root)
	put(t, root, "committed.txt", []byte("committed after base\n"))
	git(t, root, "add", "committed.txt")
	git(t, root, "commit", "-qm", "after base")
	put(t, root, "selected.txt", []byte("selected\n"))
	put(t, root, "staged.txt", []byte("staged\n"))
	git(t, root, "add", "selected.txt", "staged.txt")
	put(t, root, "unstaged.txt", []byte("unstaged\n"))
	git(t, root, "mv", "old-name.txt", "new-name.txt")
	put(t, root, "untracked.txt", []byte("new\n"))
	changes, err := gitstore.Changes(root, base)
	if err != nil {
		t.Fatal(err)
	}
	byPath := make(map[string]gitstore.Change)
	for _, change := range changes {
		if change.Status == "" || filepath.IsAbs(change.Path) {
			t.Errorf("invalid change: %#v", change)
		}
		byPath[change.Path] = change
	}
	for _, path := range []string{"committed.txt", "selected.txt", "staged.txt", "unstaged.txt", "untracked.txt", "new-name.txt"} {
		if _, ok := byPath[path]; !ok {
			t.Errorf("missing change %q: %#v", path, changes)
		}
	}
	if got := byPath["new-name.txt"].OldPath; got != "old-name.txt" {
		t.Errorf("rename OldPath = %q", got)
	}
	data, err := gitstore.ReadAt(root, base, "selected.txt")
	if err != nil || !bytes.Equal(data, []byte("original\n")) {
		t.Errorf("ReadAt(base) = %q, %v", data, err)
	}
	if _, err := gitstore.ReadAt(root, base, "untracked.txt"); err == nil {
		t.Error("ReadAt(absent path) succeeded")
	}
	gotHead, err := gitstore.Commit(root, "selected only", []string{"selected.txt"})
	if err != nil || gotHead == "" || gotHead != head(t, root) {
		t.Fatalf("Commit = %q, %v", gotHead, err)
	}
	if got := git(t, root, "diff-tree", "--no-commit-id", "--name-only", "-r", gotHead); got != "selected.txt" {
		t.Errorf("committed paths = %q, want selected.txt", got)
	}
	if got := git(t, root, "log", "-1", "--format=%s"); got != "selected only" {
		t.Errorf("subject = %q", got)
	}
	if got := git(t, root, "show", "HEAD:selected.txt"); got != "selected" {
		t.Errorf("selected HEAD content = %q", got)
	}
	if got := git(t, root, "show", "HEAD:staged.txt"); got != "original" {
		t.Errorf("unrelated staged file was committed: %q", got)
	}
	if got := git(t, root, "diff", "--cached", "--name-only"); !strings.Contains(got, "staged.txt") || !strings.Contains(got, "new-name.txt") {
		t.Errorf("unrelated index changes lost: %q", got)
	}
	if data, err := os.ReadFile(filepath.Join(root, "unstaged.txt")); err != nil || string(data) != "unstaged\n" {
		t.Errorf("unrelated worktree change lost: %q, %v", data, err)
	}
}

func Test_GIT_004_GIT_008_GIT_009_GIT_010_BlobsAndRefCAS(t *testing.T) {
	root := repo(t)
	payload := []byte{0, 'b', '\n', 255}
	oid, err := gitstore.WriteBlob(root, payload)
	if err != nil || !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(oid) {
		t.Fatalf("WriteBlob = %q, %v", oid, err)
	}
	if got, err := gitstore.ReadBlob(root, oid); err != nil || !bytes.Equal(got, payload) {
		t.Errorf("ReadBlob = %v, %v", got, err)
	}
	name := "refs/fireweed/eval-claim"
	if got, exists, err := gitstore.Ref(root, name); err != nil || exists || got != "" {
		t.Errorf("absent Ref = %q, %v, %v", got, exists, err)
	}
	first := head(t, root)
	git(t, root, "commit", "-q", "--allow-empty", "-m", "second")
	second := head(t, root)
	git(t, root, "commit", "-q", "--allow-empty", "-m", "third")
	third := head(t, root)
	if err := gitstore.CASRef(root, name, "", first); err != nil {
		t.Fatal(err)
	}
	if got, exists, err := gitstore.Ref(root, name); err != nil || !exists || got != first {
		t.Errorf("created Ref = %q, %v, %v", got, exists, err)
	}
	if err := gitstore.CASRef(root, name, "", second); err == nil {
		t.Error("create-only overwrote existing ref")
	}
	if err := gitstore.CASRef(root, name, first, second); err != nil {
		t.Fatal(err)
	}
	t.Run("StaleUpdatePreservesNewerRef", func(t *testing.T) {
		if err := gitstore.CASRef(root, name, first, third); err == nil {
			t.Error("stale update succeeded")
		}
		if got, exists, err := gitstore.Ref(root, name); err != nil || !exists || got != second {
			t.Errorf("stale update changed Ref = %q, %v, %v", got, exists, err)
		}
	})
	if err := gitstore.CASRef(root, name, first, ""); err == nil {
		t.Error("stale delete succeeded")
	}
	if got, exists, err := gitstore.Ref(root, name); err != nil || !exists || got != second {
		t.Errorf("stale operation changed Ref = %q, %v, %v", got, exists, err)
	}
	if err := gitstore.CASRef(root, name, second, ""); err != nil {
		t.Fatal(err)
	}
	if got, exists, err := gitstore.Ref(root, name); err != nil || exists || got != "" {
		t.Errorf("deleted Ref = %q, %v, %v", got, exists, err)
	}
}

func Test_GIT_005_RunCaptureExitTimeoutAndGroupCancel(t *testing.T) {
	dir := t.TempDir()
	var err error
	dir, err = filepath.EvalSymlinks(dir)
	if err != nil {
		t.Fatal(err)
	}
	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Args: []string{"/bin/sh", "-c", `printf '%s|%s|' "$PWD" "$FIREWEED_EVAL"; cat; printf 'warn' >&2`},
		Dir:  dir, Env: map[string]string{"FIREWEED_EVAL": "supplied"}, Input: []byte("input"), Timeout: 3 * time.Second,
	})
	if err != nil || result.Stdout != dir+"|supplied|input" || result.Stderr != "warn" || result.ExitCode != 0 {
		t.Errorf("successful Run = %#v, %v", result, err)
	}
	result, err = gitstore.Run(context.Background(), gitstore.Request{
		Args: []string{"/bin/sh", "-c", `printf 'out'; printf 'err' >&2; exit 7`}, Dir: dir, Timeout: 3 * time.Second,
	})
	if err == nil || result.Stdout != "out" || result.Stderr != "err" || result.ExitCode != 7 {
		t.Errorf("failed Run = %#v, %v", result, err)
	}
	for _, mode := range []struct {
		name    string
		timeout time.Duration
		cancel  bool
	}{
		{name: "TimeoutTerminatesGroup", timeout: 500 * time.Millisecond},
		{name: "ContextCancellationTerminatesGroup", timeout: 3 * time.Second, cancel: true},
	} {
		t.Run(mode.name, func(t *testing.T) {
			ready := filepath.Join(dir, mode.name+"-ready")
			marker := filepath.Join(dir, mode.name+"-child-marker")
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			done := make(chan error, 1)
			start := time.Now()
			go func() {
				_, runErr := gitstore.Run(ctx, gitstore.Request{
					Args: []string{"/bin/sh", "-c", `(sleep 1; printf child > "$2") & printf ready > "$1"; wait`, "sh", ready, marker},
					Dir:  dir, Timeout: mode.timeout,
				})
				done <- runErr
			}()
			deadline := time.After(2 * time.Second)
			for {
				if _, statErr := os.Stat(ready); statErr == nil {
					break
				}
				select {
				case runErr := <-done:
					t.Fatalf("Run exited before child was ready: %v", runErr)
				case <-deadline:
					t.Fatal("child did not become ready")
				case <-time.After(10 * time.Millisecond):
				}
			}
			if mode.cancel {
				cancel()
			}
			select {
			case runErr := <-done:
				if runErr == nil {
					t.Error("terminated Run succeeded")
				}
			case <-time.After(2 * time.Second):
				t.Error("terminated Run did not return promptly")
			}
			if elapsed := time.Since(start); elapsed > 3*time.Second {
				t.Errorf("terminated Run took %s", elapsed)
			}
			time.Sleep(1200 * time.Millisecond)
			if _, statErr := os.Stat(marker); !os.IsNotExist(statErr) {
				t.Errorf("child survived group termination: marker stat = %v", statErr)
			}
		})
	}
}
