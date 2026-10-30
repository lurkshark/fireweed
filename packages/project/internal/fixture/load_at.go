package fixture

import (
	"bytes"
	"context"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project/model"
)

// LoadAt reads project documents from one resolved commit, independent of the worktree.
func LoadAt(root, revision string) (model.Snapshot, error) {
	top, err := readGit(root, "rev-parse", "--show-toplevel")
	if err != nil {
		return model.Snapshot{}, fmt.Errorf("find repository root: %w", err)
	}
	canonical, err := filepath.EvalSymlinks(strings.TrimSpace(string(top)))
	if err != nil {
		return model.Snapshot{}, fmt.Errorf("canonical repository root: %w", err)
	}
	resolved, err := readGit(canonical, "rev-parse", "--verify", "--end-of-options", revision+"^{commit}")
	if err != nil {
		return model.Snapshot{}, fmt.Errorf("resolve revision %q: %w", revision, err)
	}
	commit := strings.TrimSpace(string(resolved))

	data, err := gitstore.ReadAt(canonical, commit, "fireweed.yaml")
	if err != nil {
		return model.Snapshot{}, fmt.Errorf("read fireweed.yaml at %s: %w", commit, err)
	}
	config, patterns, defaults, err := parseConfig(data)
	if err != nil {
		return model.Snapshot{}, err
	}

	trees, err := readGit(canonical, "ls-tree", "-r", "-d", "--name-only", "-z", commit)
	if err != nil {
		return model.Snapshot{}, fmt.Errorf("list module directories at %s: %w", commit, err)
	}
	snapshot := model.Snapshot{Root: canonical, Config: config}
	var dirs []string
	for _, entry := range bytes.Split(trees, []byte{0}) {
		if len(entry) == 0 {
			continue
		}
		dir := string(entry)
		if dir == ".git" || strings.Contains(dir, "/.git/") || strings.HasSuffix(dir, "/.git") {
			continue
		}
		if matchesModule(patterns, dir) {
			dirs = append(dirs, dir)
		}
	}
	sort.Strings(dirs)
	for _, dir := range dirs {
		path := dir + "/SPEC.md"
		data, err := gitstore.ReadAt(canonical, commit, path)
		if err != nil {
			return model.Snapshot{}, fmt.Errorf("read %s at %s: %w", path, commit, err)
		}
		module, err := parseModule(data, defaults, dir)
		if err != nil {
			return model.Snapshot{}, err
		}
		snapshot.Modules = append(snapshot.Modules, module)
	}
	return snapshot, nil
}

func readGit(root string, args ...string) ([]byte, error) {
	argv := append([]string{"git", "-C", root, "-c", "core.hooksPath=/dev/null"}, args...)
	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Args: argv, Timeout: 30 * time.Second,
		Env: map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0"},
	})
	if err != nil {
		return nil, fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, result.Stderr)
	}
	return []byte(result.Stdout), nil
}
