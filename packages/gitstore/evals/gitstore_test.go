package evals_test

import (
	"bytes"
	"context"
	"errors"
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
	put(t, root, "staged.txt", []byte("staged worktree\n"))
	put(t, root, "unstaged.txt", []byte("unstaged\n"))
	git(t, root, "mv", "old-name.txt", "new-name.txt")
	put(t, root, "new-name.txt", []byte("renamed worktree\n"))
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
	if len(changes) != len(byPath) {
		t.Errorf("duplicate reported paths: %#v", changes)
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
	if got := git(t, root, "show", ":staged.txt"); got != "staged" {
		t.Errorf("unrelated staged index content changed: %q", got)
	}
	if got := git(t, root, "diff", "--cached", "--name-only"); !strings.Contains(got, "staged.txt") || !strings.Contains(got, "new-name.txt") {
		t.Errorf("unrelated index changes lost: %q", got)
	}
	if data, err := os.ReadFile(filepath.Join(root, "unstaged.txt")); err != nil || string(data) != "unstaged\n" {
		t.Errorf("unrelated worktree change lost: %q, %v", data, err)
	}
	put(t, root, "selected.txt", []byte("selected index\n"))
	git(t, root, "add", "selected.txt")
	put(t, root, "selected.txt", []byte("selected worktree\n"))
	put(t, root, "new-selected.txt", []byte("new selected\n"))
	gotHead, err = gitstore.Commit(root, "selected worktree and new file", []string{"selected.txt", "new-selected.txt"})
	if err != nil || gotHead == "" || gotHead != head(t, root) {
		t.Fatalf("Commit selected worktree and new file = %q, %v", gotHead, err)
	}
	if got := git(t, root, "diff-tree", "--no-commit-id", "--name-only", "-r", gotHead); got != "new-selected.txt\nselected.txt" {
		t.Errorf("committed selected paths = %q, want new-selected.txt and selected.txt", got)
	}
	if got := git(t, root, "log", "-1", "--format=%s"); got != "selected worktree and new file" {
		t.Errorf("selected commit subject = %q", got)
	}
	if got := git(t, root, "show", "HEAD:selected.txt"); got != "selected worktree" {
		t.Errorf("selected worktree HEAD content = %q", got)
	}
	if got := git(t, root, "show", "HEAD:new-selected.txt"); got != "new selected" {
		t.Errorf("new selected HEAD content = %q", got)
	}
	if got := git(t, root, "show", "HEAD:staged.txt"); got != "original" {
		t.Errorf("unrelated staged file was committed: %q", got)
	}
	if got := git(t, root, "show", ":staged.txt"); got != "staged" {
		t.Errorf("unrelated staged index content changed: %q", got)
	}
	if got := git(t, root, "diff", "--cached", "--name-only"); !strings.Contains(got, "staged.txt") || !strings.Contains(got, "new-name.txt") {
		t.Errorf("unrelated index changes lost: %q", got)
	}
	if data, err := os.ReadFile(filepath.Join(root, "unstaged.txt")); err != nil || string(data) != "unstaged\n" {
		t.Errorf("unrelated worktree change lost: %q, %v", data, err)
	}
}

func Test_GIT_003_UnstagedTrackedRename(t *testing.T) {
	root := repo(t)
	base := head(t, root)
	if err := os.Rename(filepath.Join(root, "seed.txt"), filepath.Join(root, "renamed-seed.txt")); err != nil {
		t.Fatal(err)
	}
	git(t, root, "add", "-N", "renamed-seed.txt")
	changes, err := gitstore.Changes(root, base)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for _, change := range changes {
		if change.Path == "renamed-seed.txt" {
			count++
			if change.OldPath != "seed.txt" {
				t.Errorf("unstaged rename = %#v, want old path seed.txt", change)
			}
		}
	}
	if count != 1 {
		t.Errorf("renamed-seed.txt reported %d times: %#v", count, changes)
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

func Test_GIT_010_DeleteWithAbsentExpectedRef(t *testing.T) {
	for _, tc := range []struct {
		name string
		old  string
	}{
		{name: "Empty", old: ""},
		{name: "ZeroSHA1", old: strings.Repeat("0", 40)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := repo(t)
			name := "refs/fireweed/eval-claim"

			if err := gitstore.CASRef(root, name, tc.old, ""); err != nil {
				t.Errorf("delete of absent ref: %v", err)
			}
			if got, exists, err := gitstore.Ref(root, name); err != nil || exists || got != "" {
				t.Errorf("Ref after absent delete = %q, %v, %v, want absent", got, exists, err)
			}

			oid := head(t, root)
			git(t, root, "update-ref", name, oid)
			if err := gitstore.CASRef(root, name, tc.old, ""); err == nil {
				t.Error("delete of existing ref succeeded despite expected absence")
			}
			if got, exists, err := gitstore.Ref(root, name); err != nil || !exists || got != oid {
				t.Errorf("Ref after mismatched delete = %q, %v, %v, want %q, true, nil", got, exists, err, oid)
			}
		})
	}
}

func Test_GIT_005_RunRequiresExecutableForGitSubcommand(t *testing.T) {
	root := repo(t)
	gitPath, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	bin := t.TempDir()
	if err := os.Symlink(gitPath, filepath.Join(bin, "git")); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", bin)

	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Args: []string{"ls-tree", "HEAD"}, Dir: root, Timeout: 3 * time.Second,
	})
	if err == nil {
		t.Errorf("Run launched ls-tree without an executable on PATH: %#v", result)
	}
}

func Test_GIT_011_RunLaunchFailure(t *testing.T) {
	dir := t.TempDir()
	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Args: []string{filepath.Join(dir, "missing-executable")}, Dir: dir, Timeout: 3 * time.Second,
	})
	if err == nil {
		t.Fatalf("Run with missing executable succeeded: %#v", result)
	}
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		t.Errorf("Run with missing executable returned an *exec.ExitError: %v", err)
	}
}

func Test_GIT_005_GIT_011_RunCaptureExitTimeoutAndGroupCancel(t *testing.T) {
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
	const (
		stdoutMarker = "gitstore-stdout-marker"
		stderrMarker = "gitstore-stderr-marker"
		emitOutput   = `printf '%s' "$FIREWEED_EVAL_STDOUT"; printf '%s' "$FIREWEED_EVAL_STDERR" >&2; `
	)
	env := map[string]string{"FIREWEED_EVAL_STDOUT": stdoutMarker, "FIREWEED_EVAL_STDERR": stderrMarker}
	checkFailedOutput := func(t *testing.T, result gitstore.Result, err error) {
		t.Helper()
		if result.Stdout != stdoutMarker || result.Stderr != stderrMarker {
			t.Errorf("failed Run captured stdout %q and stderr %q, want %q and %q", result.Stdout, result.Stderr, stdoutMarker, stderrMarker)
		}
		if err != nil && (strings.Contains(err.Error(), stdoutMarker) || strings.Contains(err.Error(), stderrMarker)) {
			t.Errorf("Run error includes captured output: %v", err)
		}
	}
	result, err = gitstore.Run(context.Background(), gitstore.Request{
		Args: []string{"/bin/sh", "-c", emitOutput + `exit 7`}, Dir: dir, Env: env, Timeout: 3 * time.Second,
	})
	if result.ExitCode != 7 {
		t.Errorf("failed Run = %#v, %v", result, err)
	}
	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) {
		t.Errorf("failed Run error = %v, want *exec.ExitError", err)
	} else if code := exitErr.ExitCode(); code != result.ExitCode {
		t.Errorf("process error exit code = %d, want Result.ExitCode %d", code, result.ExitCode)
	}
	checkFailedOutput(t, result, err)
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
			type outcome struct {
				result gitstore.Result
				err    error
			}
			done := make(chan outcome, 1)
			start := time.Now()
			go func() {
				// Keep the shell running while the test observes readiness and cancels Run.
				result, runErr := gitstore.Run(ctx, gitstore.Request{
					Args: []string{"/bin/sh", "-c", emitOutput + `(sleep 1; printf child > "$2") & printf ready > "$1"; wait; sleep 5`, "sh", ready, marker},
					Dir:  dir, Env: env, Timeout: mode.timeout,
				})
				done <- outcome{result: result, err: runErr}
			}()
			deadline := time.After(2 * time.Second)
			for {
				if _, statErr := os.Stat(ready); statErr == nil {
					break
				}
				select {
				case got := <-done:
					t.Fatalf("Run exited before child was ready: %v", got.err)
				case <-deadline:
					t.Fatal("child did not become ready")
				case <-time.After(10 * time.Millisecond):
				}
			}
			wantErr := context.DeadlineExceeded
			if mode.cancel {
				wantErr = context.Canceled
				cancel()
			}
			select {
			case got := <-done:
				if !errors.Is(got.err, wantErr) {
					t.Errorf("terminated Run error = %v, want %v", got.err, wantErr)
				}
				checkFailedOutput(t, got.result, got.err)
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

func Test_GIT_012_RunEnvironmentOverrides(t *testing.T) {
	t.Setenv("FIREWEED_EVAL_ENV_OVERRIDE", "parent-value")
	t.Setenv("FIREWEED_EVAL_ENV_INHERITED", "inherited-value")

	const script = `#!/bin/sh
printf '%s\n' "$FIREWEED_EVAL_ENV_OVERRIDE" "$FIREWEED_EVAL_ENV_INHERITED"
`
	dir := t.TempDir()
	fixture := filepath.Join(dir, "print-environment.sh")
	if err := os.WriteFile(fixture, []byte(script), 0755); err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		name string
		env  map[string]string
		want string
	}{
		{name: "Override", env: map[string]string{"FIREWEED_EVAL_ENV_OVERRIDE": "request-value"}, want: "request-value\ninherited-value\n"},
		{name: "NilMap", want: "parent-value\ninherited-value\n"},
		{name: "EmptyMap", env: map[string]string{}, want: "parent-value\ninherited-value\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			result, err := gitstore.Run(context.Background(), gitstore.Request{
				Args: []string{fixture}, Dir: dir, Env: tc.env, Timeout: 3 * time.Second,
			})
			if err != nil {
				t.Fatalf("Run: %v, stderr = %q", err, result.Stderr)
			}
			if result.Stdout != tc.want {
				t.Errorf("Run stdout = %q, want %q", result.Stdout, tc.want)
			}
		})
	}
}

func Test_GIT_005_RunCancelWhileShellLaunchesChild(t *testing.T) {
	const script = `#!/bin/bash
printf 'captured stdout\n'
printf 'captured stderr\n' >&2
printf ready > ready
sleep 5
`
	fixture := filepath.Join(t.TempDir(), "cancel.sh")
	if err := os.WriteFile(fixture, []byte(script), 0755); err != nil {
		t.Fatal(err)
	}
	// Cancellation races the child launch, so repeat the same fixture.
	for attempt := 1; attempt <= 60 && !t.Failed(); attempt++ {
		checkCancelWhileShellLaunchesChild(t, fixture, attempt)
	}
}

func checkCancelWhileShellLaunchesChild(t *testing.T, fixture string, attempt int) {
	t.Helper()
	dir := t.TempDir()
	ready := filepath.Join(dir, "ready")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	type outcome struct {
		result gitstore.Result
		err    error
	}
	done := make(chan outcome, 1)
	go func() {
		result, err := gitstore.Run(ctx, gitstore.Request{
			Args: []string{fixture}, Dir: dir, Timeout: 8 * time.Second,
		})
		done <- outcome{result: result, err: err}
	}()

	deadline := time.NewTimer(3 * time.Second)
	defer deadline.Stop()
	for {
		if _, err := os.Stat(ready); err == nil {
			break
		} else if !os.IsNotExist(err) {
			t.Fatal(err)
		}
		select {
		case got := <-done:
			t.Fatalf("attempt %d: Run exited before shell readiness: %#v, %v", attempt, got.result, got.err)
		case <-deadline.C:
			t.Fatalf("attempt %d: shell did not signal readiness", attempt)
		case <-time.After(time.Millisecond):
		}
	}

	canceledAt := time.Now()
	cancel()
	select {
	case got := <-done:
		if elapsed := time.Since(canceledAt); elapsed >= 3*time.Second {
			t.Errorf("attempt %d: Run returned %s after cancellation, want under 3s for a 5s child", attempt, elapsed)
		}
		if !errors.Is(got.err, context.Canceled) {
			t.Errorf("attempt %d: Run error = %v, want context.Canceled", attempt, got.err)
		}
		if got.result.Stdout != "captured stdout\n" || got.result.Stderr != "captured stderr\n" {
			t.Errorf("attempt %d: Run captured stdout %q and stderr %q after cancellation", attempt, got.result.Stdout, got.result.Stderr)
		}
	case <-time.After(7 * time.Second):
		t.Fatalf("attempt %d: Run did not return within 7s after cancellation", attempt)
	}
}
