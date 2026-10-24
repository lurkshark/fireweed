package repository

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// readView gives Git its native comparison behavior without a configured
// promisor remote. Old Git versions may ignore GIT_NO_LAZY_FETCH, so the
// original partially populated repository cannot safely perform object reads.
// All mutable metadata, including the index, belongs to this temporary view.
type readView struct {
	root, dir string
	env       []string
}

func newReadView(root, selector string) (*readView, error) {
	root, gitDir, commonDir, err := Open(root)
	if err != nil {
		return nil, err
	}
	dir, err := os.MkdirTemp("", "fireweed-git-read-")
	if err != nil {
		return nil, err
	}
	v := &readView{root: root, dir: dir}
	success := false
	defer func() {
		if !success {
			os.RemoveAll(dir)
		}
	}()
	for _, name := range []string{"objects", "refs", "info"} {
		if err := os.Mkdir(filepath.Join(dir, name), 0700); err != nil {
			return nil, err
		}
	}
	config, err := git(root, "config", "--null", "--list", "--includes")
	if err != nil {
		return nil, err
	}
	var flattened strings.Builder
	for _, entry := range bytes.Split(config, []byte{0}) {
		if len(entry) == 0 {
			continue
		}
		key, value, valued := strings.Cut(string(entry), "\n")
		lower := strings.ToLower(key)
		// Included values are already expanded for the original checkout. Do not
		// re-evaluate relative includes or reintroduce any lazy-fetch configuration.
		if strings.HasPrefix(lower, "include.") || strings.HasPrefix(lower, "includeif.") || lower == "extensions.partialclone" || lower == "extensions.worktreeconfig" || (strings.HasPrefix(lower, "remote.") && (strings.HasSuffix(lower, ".promisor") || strings.HasSuffix(lower, ".partialclonefilter"))) {
			continue
		}
		// Inspection never needs external hook, monitor or conversion commands.
		// Built-in normalization and all other native comparison settings survive.
		if strings.HasPrefix(lower, "filter.") {
			continue
		}
		if err := writeConfigEntry(&flattened, key, value, valued); err != nil {
			return nil, err
		}
	}
	for _, entry := range [][2]string{{"core.bare", "false"}, {"core.worktree", root}, {"core.hooksPath", os.DevNull}, {"core.fsmonitor", ""}} {
		if err := writeConfigEntry(&flattened, entry[0], entry[1], true); err != nil {
			return nil, err
		}
	}
	if err := os.WriteFile(filepath.Join(dir, "config"), []byte(flattened.String()), 0600); err != nil {
		return nil, err
	}
	if err := copyReadFile(filepath.Join(gitDir, "HEAD"), filepath.Join(dir, "HEAD")); err != nil {
		return nil, err
	}
	refs, err := git(root, "for-each-ref", "--format=%(refname)%00%(objectname)%00%(symref)")
	if err != nil {
		return nil, err
	}
	// Keep existing ref names in Git's case-sensitive packed representation.
	// Materializing loose refs can merge distinct names on the host filesystem.
	var names []string
	objectsByRef := map[string]string{}
	for _, line := range bytes.Split(refs, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}
		fields := strings.Split(string(line), "\x00")
		if len(fields) != 3 {
			return nil, fmt.Errorf("invalid local ref inventory")
		}
		name, oid, target := fields[0], fields[1], fields[2]
		if !strings.HasPrefix(name, "refs/") || strings.Contains(name, "..") || strings.ContainsAny(name, "\\\x00") || filepath.IsAbs(name) {
			return nil, fmt.Errorf("invalid local ref inventory")
		}
		// Keep symbolic aliases so Git can follow their target's reflog.
		if target != "" {
			path := filepath.Join(dir, filepath.FromSlash(name))
			if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
				return nil, err
			}
			if err := os.WriteFile(path, []byte("ref: "+target+"\n"), 0600); err != nil {
				return nil, err
			}
			continue
		}
		names = append(names, name)
		objectsByRef[name] = oid
	}
	sort.Strings(names)
	var packed strings.Builder
	packed.WriteString("# pack-refs with: sorted\n")
	for _, name := range names {
		packed.WriteString(objectsByRef[name])
		packed.WriteByte(' ')
		packed.WriteString(name)
		packed.WriteByte('\n')
	}
	if err := os.WriteFile(filepath.Join(dir, "packed-refs"), []byte(packed.String()), 0600); err != nil {
		return nil, err
	}
	// Revision expressions can name per-worktree pseudorefs and reflog entries.
	// Preserve only their local metadata; resolving them still uses the view's
	// fetch-free object context.
	for _, name := range []string{"ORIG_HEAD", "FETCH_HEAD", "MERGE_HEAD", "CHERRY_PICK_HEAD", "REVERT_HEAD", "REBASE_HEAD", "AUTO_MERGE", "BISECT_HEAD"} {
		if err := copyReadOptional(filepath.Join(gitDir, name), filepath.Join(dir, name)); err != nil {
			return nil, err
		}
	}
	if strings.Contains(selector, "@{") {
		if err := copyReadTree(filepath.Join(commonDir, "logs", "refs"), filepath.Join(dir, "logs", "refs")); err != nil {
			return nil, err
		}
		if commonDir != gitDir {
			if err := copyReadTree(filepath.Join(gitDir, "logs", "refs"), filepath.Join(dir, "logs", "refs")); err != nil {
				return nil, err
			}
		}
		if err := copyReadOptional(filepath.Join(gitDir, "logs", "HEAD"), filepath.Join(dir, "logs", "HEAD")); err != nil {
			return nil, err
		}
	}
	index, err := readGitPath(root, "index")
	if err != nil {
		return nil, err
	}
	if err := copyReadOptional(index, filepath.Join(dir, "index")); err != nil {
		return nil, err
	}
	// Ask Git for this index's shared object instead of reading unrelated
	// shared indexes belonging to other checkouts.
	shared, err := git(root, "-c", "core.fsmonitor=", "rev-parse", "--shared-index-path")
	if err != nil {
		return nil, err
	}
	if path := strings.TrimSuffix(string(shared), "\n"); path != "" {
		if !filepath.IsAbs(path) {
			path = filepath.Join(root, path)
		}
		if err := copyReadFile(path, filepath.Join(dir, filepath.Base(path))); err != nil {
			return nil, err
		}
	}
	for _, name := range []string{"info/attributes", "info/exclude", "info/sparse-checkout", "info/grafts", "shallow"} {
		path, err := readGitPath(root, name)
		if err != nil {
			return nil, err
		}
		if err := copyReadOptional(path, filepath.Join(dir, filepath.FromSlash(name))); err != nil {
			return nil, err
		}
	}
	objects, err := readGitPath(root, "objects")
	if err != nil {
		return nil, err
	}
	for _, entry := range os.Environ() {
		key, _, _ := strings.Cut(entry, "=")
		key = strings.ToUpper(key)
		switch key {
		case "GIT_DIR", "GIT_COMMON_DIR", "GIT_WORK_TREE", "GIT_INDEX_FILE", "GIT_OBJECT_DIRECTORY":
			continue
		}
		if strings.HasPrefix(key, "GIT_CONFIG") {
			continue
		}
		v.env = append(v.env, entry)
	}
	v.env = append(v.env, "GIT_DIR="+dir, "GIT_WORK_TREE="+root, "GIT_INDEX_FILE="+filepath.Join(dir, "index"), "GIT_OBJECT_DIRECTORY="+objects, "GIT_CONFIG_NOSYSTEM=1", "GIT_CONFIG_GLOBAL="+os.DevNull, "GIT_CONFIG_SYSTEM="+os.DevNull, "GIT_TERMINAL_PROMPT=0", "GIT_OPTIONAL_LOCKS=0", "GIT_NO_LAZY_FETCH=1")
	success = true
	return v, nil
}
func readGitPath(root, name string) (string, error) {
	data, err := git(root, "rev-parse", "--git-path", name)
	if err != nil {
		return "", err
	}
	path := strings.TrimSuffix(string(data), "\n")
	if !filepath.IsAbs(path) {
		path = filepath.Join(root, path)
	}
	return path, nil
}
func writeConfigEntry(out *strings.Builder, key, value string, valued bool) error {
	first, last := strings.IndexByte(key, '.'), strings.LastIndexByte(key, '.')
	if first < 1 || last == len(key)-1 {
		return fmt.Errorf("invalid Git configuration key %q", key)
	}
	section, name := key[:first], key[last+1:]
	out.WriteByte('[')
	out.WriteString(section)
	if last > first {
		out.WriteByte(' ')
		out.WriteString(quoteGitConfig(key[first+1 : last]))
	}
	out.WriteString("]\n\t")
	out.WriteString(name)
	if valued {
		out.WriteString(" = ")
		out.WriteString(quoteGitConfig(value))
	}
	out.WriteByte('\n')
	return nil
}
func quoteGitConfig(value string) string {
	value = strings.NewReplacer("\\", "\\\\", "\"", "\\\"", "\n", "\\n", "\t", "\\t", "\b", "\\b").Replace(value)
	return "\"" + value + "\""
}
func copyReadFile(source, destination string) error {
	info, err := os.Stat(source)
	if err != nil {
		return err
	}
	data, err := os.ReadFile(source)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(destination), 0700); err != nil {
		return err
	}
	if err := os.WriteFile(destination, data, 0600); err != nil {
		return err
	}
	// Git compares cached file stats against the index timestamp to detect
	// racy entries. A freshly timestamped copy can hide same-size rewrites.
	return os.Chtimes(destination, info.ModTime(), info.ModTime())
}
func copyReadOptional(source, destination string) error {
	err := copyReadFile(source, destination)
	if os.IsNotExist(err) {
		return nil
	}
	return err
}
func copyReadTree(source, destination string) error {
	err := filepath.WalkDir(source, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		relative, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		return copyReadFile(path, filepath.Join(destination, relative))
	})
	if os.IsNotExist(err) {
		return nil
	}
	return err
}
func (v *readView) close() { os.RemoveAll(v.dir) }
func (v *readView) git(args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = v.root
	cmd.Env = v.env
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	data, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("local Git inspection %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(stderr.String()))
	}
	return data, nil
}
