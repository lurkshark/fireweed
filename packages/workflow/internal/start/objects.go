package start

import (
	"fmt"
	"github.com/lurkshark/fireweed/packages/contracts"
	"strings"
)

type commitData struct {
	tree    string
	parents []string
}
type treeEntry struct{ mode, kind, oid, name string }

// Older Git versions ignore GIT_NO_LAZY_FETCH. Enumerating local object names
// first lets every content read prove availability without requesting a remote.
func (g *localGit) prepare() error {
	raw, e := g.raw("cat-file", "--batch-all-objects", "--batch-check=%(objectname)")
	if e != nil {
		return e
	}
	g.objects = map[string]bool{}
	for _, id := range strings.Fields(raw) {
		g.objects[id] = true
	}
	g.commits = map[string]commitData{}
	g.trees = map[string][]treeEntry{}
	g.bodies = map[string]string{}
	return nil
}
func (g localGit) available(oid string) error {
	if !g.objects[oid] {
		return problem("local object %s is unavailable; restore the local repository objects", oid)
	}
	return nil
}
func (g localGit) commit(oid string) (commitData, error) {
	if v, ok := g.commits[oid]; ok {
		return v, nil
	}
	if e := g.available(oid); e != nil {
		return commitData{}, e
	}
	kind, e := g.text("cat-file", "-t", oid)
	if e != nil {
		return commitData{}, e
	}
	if kind != "commit" {
		return commitData{}, problem("local object %s must be a commit", oid)
	}
	raw, e := g.raw("cat-file", "-p", oid)
	if e != nil {
		return commitData{}, e
	}
	v := commitData{}
	for _, line := range strings.Split(strings.SplitN(raw, "\n\n", 2)[0], "\n") {
		if strings.HasPrefix(line, "tree ") {
			v.tree = strings.TrimPrefix(line, "tree ")
		}
		if strings.HasPrefix(line, "parent ") {
			v.parents = append(v.parents, strings.TrimPrefix(line, "parent "))
		}
	}
	if v.tree == "" {
		return v, problem("local commit %s has no tree", oid)
	}
	g.commits[oid] = v
	return v, nil
}
func (g localGit) resolveCommit(ref string) (string, error) {
	id, e := g.text("rev-parse", "--verify", "--end-of-options", ref)
	if e != nil {
		return "", e
	}
	if _, e = g.commit(id); e != nil {
		return "", e
	}
	return id, nil
}
func (g localGit) tree(oid string) ([]treeEntry, error) {
	if v, ok := g.trees[oid]; ok {
		return v, nil
	}
	if e := g.available(oid); e != nil {
		return nil, e
	}
	raw, e := g.raw("ls-tree", "-z", oid)
	if e != nil {
		return nil, e
	}
	if raw != "" && !strings.HasSuffix(raw, "\x00") {
		return nil, problem("incomplete local tree %s", oid)
	}
	entries := make([]treeEntry, 0)
	for _, record := range strings.Split(raw, "\x00") {
		if record == "" {
			continue
		}
		header, name, ok := strings.Cut(record, "\t")
		fields := strings.Fields(header)
		if !ok || len(fields) != 3 || name == "" || strings.Contains(name, "/") || name == "." || name == ".." {
			return nil, problem("invalid local tree record in %s", oid)
		}
		entries = append(entries, treeEntry{fields[0], fields[1], fields[2], name})
	}
	g.trees[oid] = entries
	return entries, nil
}
func (g localGit) blob(oid string) (string, error) {
	if text, ok := g.bodies[oid]; ok {
		return text, nil
	}
	if e := g.available(oid); e != nil {
		return "", e
	}
	text, e := g.raw("cat-file", "blob", oid)
	if e == nil {
		g.bodies[oid] = text
	}
	return text, e
}
func (g localGit) entry(commit, path string) (treeEntry, bool, error) {
	c, e := g.commit(commit)
	if e != nil {
		return treeEntry{}, false, e
	}
	current := treeEntry{kind: "tree", oid: c.tree}
	for _, part := range strings.Split(path, "/") {
		if current.kind != "tree" {
			return treeEntry{}, false, fmt.Errorf("%s is not a directory", part)
		}
		entries, e := g.tree(current.oid)
		if e != nil {
			return treeEntry{}, false, e
		}
		found := false
		for _, entry := range entries {
			if entry.name == part {
				current = entry
				found = true
				break
			}
		}
		if !found {
			return treeEntry{}, false, nil
		}
	}
	return current, true, nil
}
func (g localGit) file(commit, path string) (string, error) {
	entry, found, e := g.entry(commit, path)
	if e != nil {
		return "", e
	}
	if !found || entry.kind != "blob" {
		return "", problem("selected file %s is absent from local commit %s", path, commit)
	}
	return g.blob(entry.oid)
}
func (g localGit) completeTree(oid string, seen map[string]bool) error {
	if seen[oid] {
		return nil
	}
	seen[oid] = true
	entries, e := g.tree(oid)
	if e != nil {
		return e
	}
	for _, entry := range entries {
		if entry.kind == "commit" {
			continue
		}
		if e = g.available(entry.oid); e != nil {
			return e
		}
		if entry.kind == "tree" {
			if e = g.completeTree(entry.oid, seen); e != nil {
				return e
			}
		}
	}
	return nil
}
func (g localGit) ancestor(base, head string) bool {
	pending := []string{head}
	seen := map[string]bool{}
	for len(pending) > 0 {
		id := pending[len(pending)-1]
		pending = pending[:len(pending)-1]
		if id == base {
			return true
		}
		if seen[id] {
			continue
		}
		seen[id] = true
		c, e := g.commit(id)
		if e == nil {
			pending = append(pending, c.parents...)
		}
	}
	return false
}
func (g localGit) reachable(head string) (map[string]commitData, error) {
	result := map[string]commitData{}
	pending := []string{head}
	for len(pending) > 0 {
		id := pending[len(pending)-1]
		pending = pending[:len(pending)-1]
		if _, ok := result[id]; ok {
			continue
		}
		c, e := g.commit(id)
		if e != nil {
			return nil, e
		}
		result[id] = c
		pending = append(pending, c.parents...)
	}
	return result, nil
}
func (g localGit) later(base, head string) (map[string]commitData, error) {
	if base == head {
		return map[string]commitData{}, nil
	}
	old, e := g.reachable(base)
	if e != nil {
		return nil, e
	}
	all, e := g.reachable(head)
	if e != nil {
		return nil, e
	}
	for id := range old {
		delete(all, id)
	}
	return all, nil
}
func (g localGit) applies(commit, module string) (map[string]bool, error) {
	result := map[string]bool{}
	dir, found, e := g.entry(commit, "changes/archive")
	if e != nil {
		return nil, e
	}
	if !found {
		return result, nil
	}
	if dir.kind != "tree" {
		return nil, problem("restore archive directory in %s", commit)
	}
	entries, e := g.tree(dir.oid)
	if e != nil {
		return nil, e
	}
	for _, entry := range entries {
		if !strings.HasSuffix(entry.name, ".md") {
			continue
		}
		if entry.kind != "blob" {
			return nil, problem("restore archive document %s", entry.name)
		}
		raw, e := g.blob(entry.oid)
		if e != nil {
			return nil, e
		}
		doc, e := contracts.ParseMarkdown(raw)
		if e != nil {
			return nil, e
		}
		if doc.Metadata["module"] == module && doc.Metadata["phase"] == "applied" {
			run, ok := doc.Metadata["run"].(string)
			if !ok || run == "" {
				return nil, problem("restore applied archive run %s", entry.name)
			}
			result[run] = true
		}
	}
	return result, nil
}
