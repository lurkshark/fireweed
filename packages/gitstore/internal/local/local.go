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

// Run starts each command in its own process group so cancellation also stops
// children that remain in that group.
func Run(ctx context.Context, request model.Request) (model.Result, error) {
	var result model.Result
	if len(request.Args) == 0 {
		return result, errors.New("empty command")
	}
	if err := ctx.Err(); err != nil {
		return result, err
	}
	if request.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, request.Timeout)
		defer cancel()
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

func git(root string, input []byte, env map[string]string, args ...string) ([]byte, error) {
	if env == nil {
		env = make(map[string]string)
	}
	env["GIT_NO_LAZY_FETCH"] = "1"
	env["GIT_TERMINAL_PROMPT"] = "0"
	argv := append([]string{"git", "-C", root, "-c", "core.hooksPath=/dev/null"}, args...)
	result, err := Run(context.Background(), model.Request{Args: argv, Env: env, Input: input, Timeout: gitTimeout})
	if err != nil {
		return nil, fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(result.Stderr))
	}
	return []byte(result.Stdout), nil
}

func gitText(root string, args ...string) (string, error) {
	out, err := git(root, nil, nil, args...)
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
	root, err := gitText(dir, "rev-parse", "--show-toplevel")
	if err != nil {
		return repo, err
	}
	repo.Root, err = canonical(root)
	if err != nil {
		return repo, err
	}
	common, err := gitText(dir, "rev-parse", "--git-common-dir")
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
	repo.Head, err = gitText(dir, "rev-parse", "--verify", "HEAD")
	if err != nil {
		return repo, err
	}
	repo.Branch, _ = gitText(dir, "symbolic-ref", "--quiet", "--short", "HEAD")
	return repo, nil
}

func BaseRef(root, remote, base string) (string, error) {
	if remote == "" {
		remote = "origin"
	}
	if base == "" {
		remoteHead := "refs/remotes/" + remote + "/HEAD"
		resolved, err := gitText(root, "symbolic-ref", "--quiet", remoteHead)
		if err == nil && strings.HasPrefix(resolved, "refs/remotes/"+remote+"/") {
			base = strings.TrimPrefix(resolved, "refs/remotes/"+remote+"/")
		} else {
			base = "main"
		}
	}
	for _, name := range []string{"refs/remotes/" + remote + "/" + base, "refs/heads/" + base} {
		if _, exists, err := Ref(root, name); err != nil {
			return "", err
		} else if exists {
			return name, nil
		}
	}
	return "", fmt.Errorf("base %q not found", base)
}

func Changes(root, base string) ([]model.Change, error) {
	byPath := make(map[string]model.Change)
	for _, args := range [][]string{
		{"diff", "--cached", "--name-status", "-z", "-M", base, "--"},
		{"diff", "--name-status", "-z", "-M", "--"},
	} {
		out, err := git(root, nil, nil, args...)
		if err != nil {
			return nil, err
		}
		fields := bytes.Split(out, []byte{0})
		for i := 0; i < len(fields)-1; {
			status := string(fields[i])
			i++
			if status == "" || i >= len(fields)-1 {
				return nil, errors.New("invalid Git name-status output")
			}
			change := model.Change{Path: string(fields[i]), Status: status[:1]}
			i++
			if status[0] == 'R' || status[0] == 'C' {
				if i >= len(fields)-1 {
					return nil, errors.New("invalid Git rename output")
				}
				change.OldPath, change.Path = change.Path, string(fields[i])
				i++
			}
			if old, ok := byPath[change.Path]; !ok || (old.OldPath == "" && change.OldPath != "") {
				byPath[change.Path] = change
			}
		}
	}
	out, err := git(root, nil, nil, "ls-files", "--others", "--exclude-standard", "-z")
	if err != nil {
		return nil, err
	}
	for _, path := range bytes.Split(out, []byte{0}) {
		if len(path) > 0 {
			name := string(path)
			if _, exists := byPath[name]; !exists {
				byPath[name] = model.Change{Path: name, Status: "?"}
			}
		}
	}
	paths := make([]string, 0, len(byPath))
	for path := range byPath {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	changes := make([]model.Change, 0, len(paths))
	for _, path := range paths {
		changes = append(changes, byPath[path])
	}
	return changes, nil
}

func ReadAt(root, revision, path string) ([]byte, error) {
	return git(root, nil, nil, "show", revision+":"+path)
}

func Commit(root, subject string, paths []string) (string, error) {
	if len(paths) == 0 {
		return "", errors.New("no paths selected")
	}
	old, err := gitText(root, "rev-parse", "HEAD")
	if err != nil {
		return "", err
	}
	file, err := os.CreateTemp("", "fireweed-index-*")
	if err != nil {
		return "", err
	}
	index := file.Name()
	_ = file.Close()
	defer os.Remove(index)
	env := map[string]string{"GIT_INDEX_FILE": index}
	if _, err := git(root, nil, env, "read-tree", "HEAD"); err != nil {
		return "", err
	}
	args := append([]string{"add", "-A", "--"}, paths...)
	if _, err := git(root, nil, env, args...); err != nil {
		return "", err
	}
	tree, err := git(root, nil, env, "write-tree")
	if err != nil {
		return "", err
	}
	commit, err := git(root, nil, nil, "commit-tree", strings.TrimSpace(string(tree)), "-p", old, "-m", subject)
	if err != nil {
		return "", err
	}
	newHead := strings.TrimSpace(string(commit))
	if _, err := git(root, nil, nil, "update-ref", "HEAD", newHead, old); err != nil {
		return "", err
	}
	args = append([]string{"reset", "-q", "HEAD", "--"}, paths...)
	if _, err := git(root, nil, nil, args...); err != nil {
		return newHead, err
	}
	return newHead, nil
}

func WriteBlob(root string, data []byte) (string, error) {
	out, err := git(root, data, nil, "hash-object", "-w", "--stdin")
	return strings.TrimSpace(string(out)), err
}

func ReadBlob(root, oid string) ([]byte, error) {
	return git(root, nil, nil, "cat-file", "blob", oid)
}

func Ref(root, name string) (string, bool, error) {
	if _, err := git(root, nil, nil, "check-ref-format", name); err != nil {
		return "", false, err
	}
	out, err := git(root, nil, nil, "show-ref", "--verify", "--hash", name)
	if err == nil {
		return strings.TrimSpace(string(out)), true, nil
	}
	// Git versions differ on the status returned for a missing valid ref.
	var exit *exec.ExitError
	if errors.As(err, &exit) {
		return "", false, nil
	}
	return "", false, err
}

func CASRef(root, name, old, new string) error {
	if old == "" {
		old = strings.Repeat("0", 40)
	}
	var args []string
	if new == "" {
		args = []string{"update-ref", "-d", name, old}
	} else {
		args = []string{"update-ref", name, new, old}
	}
	_, err := git(root, nil, nil, args...)
	return err
}
