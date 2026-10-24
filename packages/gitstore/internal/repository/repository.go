package repository

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Edit struct{ Status, Path, OldPath string }

func git(root string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GIT_TERMINAL_PROMPT=0")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	data, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(stderr.String()))
	}
	return data, nil
}
func canonical(path string) (string, error) {
	absolute, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return filepath.EvalSymlinks(absolute)
}
func Open(start string) (string, string, string, error) {
	rootBytes, err := git(start, "rev-parse", "--show-toplevel")
	if err != nil {
		return "", "", "", err
	}
	root, err := canonical(strings.TrimSuffix(string(rootBytes), "\n"))
	if err != nil {
		return "", "", "", err
	}
	gitBytes, err := git(root, "rev-parse", "--absolute-git-dir")
	if err != nil {
		return "", "", "", err
	}
	gitDir, err := canonical(strings.TrimSuffix(string(gitBytes), "\n"))
	if err != nil {
		return "", "", "", err
	}
	commonBytes, err := git(root, "rev-parse", "--git-common-dir")
	if err != nil {
		return "", "", "", err
	}
	commonPath := strings.TrimSuffix(string(commonBytes), "\n")
	if !filepath.IsAbs(commonPath) {
		commonPath = filepath.Join(root, commonPath)
	}
	commonDir, err := canonical(commonPath)
	if err != nil {
		return "", "", "", err
	}
	return root, gitDir, commonDir, nil
}
func Base(root, base, remote string) (string, error) {
	if _, err := git(root, "rev-parse", "--git-dir"); err != nil {
		return "", err
	}
	if remote == "" {
		remote = "origin"
	}
	if strings.HasPrefix(remote, "-") {
		return "", fmt.Errorf("invalid remote %q", remote)
	}
	if _, err := git(root, "check-ref-format", "refs/remotes/"+remote+"/HEAD"); err != nil {
		return "", err
	}
	if base == "" {
		if data, err := git(root, "symbolic-ref", "-q", "refs/remotes/"+remote+"/HEAD"); err == nil {
			target := strings.TrimSpace(string(data))
			prefix := "refs/remotes/" + remote + "/"
			if strings.HasPrefix(target, prefix) {
				base = strings.TrimPrefix(target, prefix)
			}
		}
		if base == "" {
			base = "main"
		}
	}
	if strings.HasPrefix(base, "-") {
		return "", fmt.Errorf("invalid base %q", base)
	}
	local := "refs/heads/" + base
	if _, err := git(root, "check-ref-format", local); err != nil {
		return "", err
	}
	tracking := "refs/remotes/" + remote + "/" + base
	if _, err := git(root, "rev-parse", "--verify", "--quiet", "--end-of-options", tracking); err == nil {
		return tracking, nil
	}
	return local, nil
}
func Clean(root string) (bool, error) {
	data, err := git(root, "status", "--porcelain=v1", "-z", "--untracked-files=all")
	if err != nil {
		return false, err
	}
	return len(data) == 0, nil
}
func Edits(root, base string) ([]Edit, error) {
	if base == "" {
		return nil, fmt.Errorf("explicit base commit required")
	}
	view, err := newReadView(root, base)
	if err != nil {
		return nil, err
	}
	defer view.close()
	oid, err := view.git("rev-parse", "--verify", "--end-of-options", base+"^{commit}")
	if err != nil {
		return nil, err
	}
	data, err := view.git("diff", "--name-status", "-z", "--find-renames", "--no-ext-diff", "--no-textconv", strings.TrimSpace(string(oid)), "--")
	if err != nil {
		return nil, err
	}
	tokens := strings.Split(string(data), "\x00")
	result := []Edit{}
	for i := 0; i < len(tokens) && tokens[i] != ""; {
		status := tokens[i]
		i++
		if i >= len(tokens) || tokens[i] == "" {
			return nil, fmt.Errorf("incomplete Git diff record")
		}
		entry := Edit{Status: status[:1], Path: tokens[i]}
		i++
		if entry.Status == "R" || entry.Status == "C" {
			entry.OldPath = entry.Path
			if i >= len(tokens) || tokens[i] == "" {
				return nil, fmt.Errorf("incomplete Git rename")
			}
			entry.Path = tokens[i]
			i++
		}
		result = append(result, entry)
	}
	untracked, err := view.git("ls-files", "--others", "--exclude-standard", "-z")
	if err != nil {
		return nil, err
	}
	for _, path := range strings.Split(string(untracked), "\x00") {
		if path != "" {
			result = append(result, Edit{Status: "A", Path: path})
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Path != result[j].Path {
			return result[i].Path < result[j].Path
		}
		if result[i].OldPath != result[j].OldPath {
			return result[i].OldPath < result[j].OldPath
		}
		return result[i].Status < result[j].Status
	})
	deduplicated := result[:0]
	for _, entry := range result {
		if len(deduplicated) == 0 || entry != deduplicated[len(deduplicated)-1] {
			deduplicated = append(deduplicated, entry)
		}
	}
	return deduplicated, nil
}
func relativePath(path string) error {
	if path == "" || strings.HasPrefix(path, "/") || strings.ContainsAny(path, "\\\x00") {
		return fmt.Errorf("invalid relative path %q", path)
	}
	if len(path) > 1 && path[1] == ':' && ((path[0] >= 'a' && path[0] <= 'z') || (path[0] >= 'A' && path[0] <= 'Z')) {
		return fmt.Errorf("absolute drive path %q", path)
	}
	for _, part := range strings.Split(path, "/") {
		if part == "" || part == "." || part == ".." {
			return fmt.Errorf("invalid component in %q", path)
		}
	}
	return nil
}
func contained(root, path string) bool {
	relative, err := filepath.Rel(root, path)
	return err == nil && !filepath.IsAbs(relative) && relative != ".." && !strings.HasPrefix(relative, ".."+string(filepath.Separator))
}
func SafePath(root, relative string) (string, error) {
	if err := relativePath(relative); err != nil {
		return "", err
	}
	canonicalRoot, err := canonical(root)
	if err != nil {
		return "", err
	}
	info, err := os.Stat(canonicalRoot)
	if err != nil {
		return "", err
	}
	if !info.IsDir() {
		return "", fmt.Errorf("root is not a directory")
	}
	candidate := filepath.Join(canonicalRoot, filepath.FromSlash(relative))
	ancestor := candidate
	for {
		_, err := os.Lstat(ancestor)
		if err == nil {
			break
		}
		if !os.IsNotExist(err) {
			return "", err
		}
		parent := filepath.Dir(ancestor)
		if parent == ancestor {
			return "", fmt.Errorf("no existing ancestor")
		}
		ancestor = parent
	}
	resolved, err := filepath.EvalSymlinks(ancestor)
	if err != nil {
		return "", err
	}
	if !contained(canonicalRoot, resolved) {
		return "", fmt.Errorf("path escapes root through symlink: %s", relative)
	}
	return candidate, nil
}
