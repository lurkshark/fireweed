package local

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore/model"
)

const gitTimeout = 30 * time.Second

// Run keeps the command and its ordinary descendants in one process group.
func Run(ctx context.Context, request model.Request) (model.Result, error) {
	var result model.Result
	if len(request.Args) == 0 {
		return result, errors.New("empty command")
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if request.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, request.Timeout)
		defer cancel()
	}
	if err := ctx.Err(); err != nil {
		return result, err
	}
	cmd := exec.Command(request.Args[0], request.Args[1:]...)
	cmd.Dir = request.Dir
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Env = os.Environ()
	keys := make([]string, 0, len(request.Env))
	for key := range request.Env {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		cmd.Env = append(cmd.Env, key+"="+request.Env[key])
	}
	cmd.Stdin = bytes.NewReader(request.Input)
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	if err := cmd.Start(); err != nil {
		return result, err
	}
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	var err error
	select {
	case err = <-done:
	case <-ctx.Done():
		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		err = <-done
		result = model.Result{Stdout: stdout.String(), Stderr: stderr.String(), ExitCode: cmd.ProcessState.ExitCode()}
		return result, ctx.Err()
	}
	result = model.Result{Stdout: stdout.String(), Stderr: stderr.String(), ExitCode: cmd.ProcessState.ExitCode()}
	if err != nil {
		return result, err
	}
	return result, nil
}

func git(root string, input []byte, args ...string) ([]byte, error) {
	argv := append([]string{"git", "-C", root, "-c", "core.hooksPath=/dev/null"}, args...)
	result, err := Run(context.Background(), model.Request{
		Args: argv, Input: input, Timeout: gitTimeout,
		Env: map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0"},
	})
	if err != nil {
		return nil, fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, result.Stderr)
	}
	return []byte(result.Stdout), nil
}

func line(root string, args ...string) (string, error) {
	out, err := git(root, nil, args...)
	return strings.TrimSpace(string(out)), err
}

func canonical(path string) (string, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return filepath.EvalSymlinks(path)
}

func Inspect(dir string) (model.Repository, error) {
	var repo model.Repository
	root, err := line(dir, "rev-parse", "--show-toplevel")
	if err != nil {
		return repo, err
	}
	repo.Root, err = canonical(root)
	if err != nil {
		return repo, err
	}
	common, err := line(dir, "rev-parse", "--git-common-dir")
	if err != nil {
		return repo, err
	}
	if !filepath.IsAbs(common) {
		common = filepath.Join(dir, common)
	}
	repo.CommonDir, err = canonical(common)
	if err != nil {
		return repo, err
	}
	repo.Head, err = line(dir, "rev-parse", "--verify", "HEAD")
	if err != nil {
		return repo, err
	}
	// A detached HEAD makes symbolic-ref exit nonzero.
	repo.Branch, _ = line(dir, "symbolic-ref", "--quiet", "--short", "HEAD")
	return repo, nil
}

func refExists(root, name string) bool {
	_, err := git(root, nil, "show-ref", "--verify", "--quiet", name)
	return err == nil
}

func BaseRef(root, remote, base string) (string, error) {
	if remote == "" {
		remote = "origin"
	}
	if base == "" {
		prefix := "refs/remotes/" + remote + "/"
		if target, err := line(root, "symbolic-ref", "--quiet", prefix+"HEAD"); err == nil && strings.HasPrefix(target, prefix) {
			base = strings.TrimPrefix(target, prefix)
		}
		if base == "" {
			base = "main"
		}
	}
	remoteRef := "refs/remotes/" + remote + "/" + base
	if refExists(root, remoteRef) {
		return remoteRef, nil
	}
	localRef := "refs/heads/" + base
	if refExists(root, localRef) {
		return localRef, nil
	}
	return "", fmt.Errorf("base branch %q not found", base)
}

func Changes(root, base string) ([]model.Change, error) {
	changes := make([]model.Change, 0)
	seen := make(map[string]int)
	add := func(change model.Change) {
		if at, ok := seen[change.Path]; ok {
			// Keep the old path when the index identifies a rename.
			if change.OldPath != "" && changes[at].OldPath == "" {
				changes[at] = change
			}
			return
		}
		seen[change.Path] = len(changes)
		changes = append(changes, change)
	}
	for _, cached := range []bool{false, true} {
		args := []string{"--no-optional-locks", "diff", "--no-ext-diff", "--no-textconv", "--name-status", "-z", "-M"}
		if cached {
			args = append(args, "--cached")
		}
		args = append(args, base, "--")
		out, err := git(root, nil, args...)
		if err != nil {
			return nil, err
		}
		fields := bytes.Split(out, []byte{0})
		for i := 0; i < len(fields)-1; {
			status := string(fields[i])
			i++
			if i >= len(fields)-1 || status == "" {
				return nil, errors.New("malformed git diff output")
			}
			path := string(fields[i])
			i++
			change := model.Change{Path: path, Status: status[:1]}
			if status[0] == 'R' || status[0] == 'C' {
				if i >= len(fields)-1 {
					return nil, errors.New("malformed git rename output")
				}
				change.OldPath = path
				change.Path = string(fields[i])
				i++
			}
			add(change)
		}
	}
	untracked, err := git(root, nil, "ls-files", "--others", "--exclude-standard", "-z")
	if err != nil {
		return nil, err
	}
	for _, path := range bytes.Split(untracked, []byte{0}) {
		if len(path) != 0 {
			add(model.Change{Path: string(path), Status: "?"})
		}
	}
	sort.Slice(changes, func(i, j int) bool { return changes[i].Path < changes[j].Path })
	return changes, nil
}

func ReadAt(root, revision, path string) ([]byte, error) {
	return git(root, nil, "cat-file", "blob", revision+":"+path)
}

func Commit(root, subject string, paths []string) (string, error) {
	if len(paths) == 0 {
		return "", errors.New("no paths selected")
	}
	stage := append([]string{"--literal-pathspecs", "add", "-A", "--"}, paths...)
	if _, err := git(root, nil, stage...); err != nil {
		return "", err
	}
	args := []string{"--literal-pathspecs", "commit", "--only", "-m", subject, "--"}
	args = append(args, paths...)
	if _, err := git(root, nil, args...); err != nil {
		return "", err
	}
	return line(root, "rev-parse", "--verify", "HEAD")
}

func WriteBlob(root string, data []byte) (string, error) {
	out, err := git(root, data, "hash-object", "-w", "--stdin")
	return strings.TrimSpace(string(out)), err
}

func ReadBlob(root, oid string) ([]byte, error) {
	return git(root, nil, "cat-file", "blob", oid)
}

func Ref(root, name string) (string, bool, error) {
	if _, err := git(root, nil, "check-ref-format", name); err != nil {
		return "", false, err
	}
	if !refExists(root, name) {
		return "", false, nil
	}
	oid, err := line(root, "show-ref", "--verify", "--hash", name)
	return oid, err == nil, err
}

func CASRef(root, name, old, new string) error {
	if new == "" {
		if old == "" {
			return errors.New("cannot delete absent ref")
		}
		_, err := git(root, nil, "update-ref", "-d", name, old)
		return err
	}
	if old == "" {
		old = strings.Repeat("0", 40)
	}
	_, err := git(root, nil, "update-ref", name, new, old)
	return err
}
