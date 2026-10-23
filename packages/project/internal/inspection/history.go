package inspection

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project/internal/changeview"
)

type committedNode struct{ mode, kind, oid string }
type committedTree struct {
	root    string
	entries map[string]committedNode
	bodies  map[string]string
	objects map[string]bool
}

// LoadAt restores metadata from one immutable source without exporting its files.
func LoadAt(start, revision string) (Data, error) {
	fail := func(action string, err error) (Data, error) {
		return Data{}, fmt.Errorf("C01: %s; restore a valid local baseline: %w", action, err)
	}
	if strings.TrimSpace(revision) == "" || strings.HasPrefix(revision, "-") || strings.ContainsRune(revision, 0) {
		return fail("select an explicit local commit", fmt.Errorf("invalid revision %q", revision))
	}
	repo, err := gitstore.Open(start)
	if err != nil {
		return fail("open the calling Git checkout", err)
	}
	objects, err := baselineLocalObjects(repo.Root)
	if err != nil {
		return fail("inspect locally available metadata objects", err)
	}
	oid, err := baselineCommit(repo.Root, revision, objects)
	if err != nil {
		return fail("resolve the explicit local commit", err)
	}
	raw, err := baselineTreeMetadata(repo.Root, oid, objects)
	if err != nil {
		return fail("enumerate the selected commit tree", err)
	}
	tree, err := parseCommittedTree(repo.Root, raw)
	if err != nil {
		return fail("read the selected commit tree", err)
	}
	tree.objects = objects
	text, err := tree.file("fireweed.yaml")
	if err != nil {
		return fail("read selected fireweed.yaml", err)
	}
	cfg, err := contracts.ParseConfig(text, assets.Info().Version)
	if err != nil {
		return fail("fix selected fireweed.yaml or use a compatible binary", err)
	}
	patterns := append([]string{}, cfg.Modules...)
	patterns = append(patterns, zonePatterns(cfg.ModuleDefaults)...)
	patterns = append(patterns, cfg.Protected...)
	patterns = append(patterns, cfg.WorkerWritable...)
	for _, pattern := range patterns {
		if _, err := contracts.MatchGlob(pattern, "future-path"); err != nil {
			return fail("fix the configured glob in selected fireweed.yaml", err)
		}
	}
	base, err := gitstore.BaseRef(repo.Root, cfg.Base, cfg.Remote)
	if err != nil {
		return fail("resolve the selected fireweed.yaml local base", err)
	}
	d := Data{Root: repo.Root, BaseRef: base, Config: cfg, Modules: []contracts.Module{}, Changes: []contracts.Change{}, Archive: []contracts.Change{}, Diagnostics: []Diagnostic{}}
	candidates := []string{}
	for relative := range tree.entries {
		if path.Base(relative) != "SPEC.md" || hiddenTreePath(relative) {
			continue
		}
		candidates = append(candidates, relative)
	}
	sort.Strings(candidates)
	for _, relative := range candidates {
		dir := path.Dir(relative)
		matched := false
		for _, pattern := range cfg.Modules {
			ok, _ := contracts.MatchGlob(pattern, dir)
			if ok {
				matched = true
				break
			}
		}
		if !matched {
			d.add("C02", "warning", relative, "Move SPEC.md under a configured modules glob or update fireweed.yaml.")
			continue
		}
		text, err := tree.file(relative)
		if err != nil {
			d.add("C03", "error", relative, "Restore a safe readable committed module: "+err.Error())
			continue
		}
		m, code, err := parseModule(text, dir, cfg)
		if err != nil {
			d.add(code, "error", relative, "Correct the module document: "+err.Error())
			continue
		}
		d.Modules = append(d.Modules, m)
	}
	sort.Slice(d.Modules, func(i, j int) bool {
		a, b := d.Modules[i], d.Modules[j]
		if a.Name != b.Name {
			return a.Name < b.Name
		}
		return a.Dir < b.Dir
	})
	prefixes := map[string]string{}
	for _, m := range d.Modules {
		prefixes[m.Name] = m.IDPrefix
	}
	d.Changes = tree.changes(&d, "changes", prefixes)
	d.Archive = tree.changes(&d, "changes/archive", prefixes)
	seen := map[string]string{}
	for _, list := range [][]contracts.Change{d.Changes, d.Archive} {
		for _, ch := range list {
			if earlier, ok := seen[ch.Slug]; ok {
				d.add("C05", "error", ch.Path, fmt.Sprintf("Slug %q also appears in %s; choose unique change slugs.", ch.Slug, earlier))
			} else {
				seen[ch.Slug] = ch.Path
			}
		}
	}
	sortDiagnostics(d.Diagnostics)
	return d, nil
}

func baselineGit(root string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	argv := append([]string{"git", "-c", "core.fsmonitor=false", "-c", "core.hooksPath=" + os.DevNull}, args...)
	result := gitstore.Capture(ctx, root, argv, []string{"GIT_TERMINAL_PROMPT=0", "GIT_OPTIONAL_LOCKS=0", "GIT_NO_REPLACE_OBJECTS=1", "GIT_NO_LAZY_FETCH=1"})
	if result.Err != nil {
		return "", fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), result.Err, strings.TrimSpace(result.Stderr))
	}
	return result.Stdout, nil
}
func fullObjectID(s string) bool {
	if len(s) != 40 && len(s) != 64 {
		return false
	}
	for _, r := range s {
		if !(r >= '0' && r <= '9' || r >= 'a' && r <= 'f') {
			return false
		}
	}
	return true
}
func parseCommittedTree(root, raw string) (*committedTree, error) {
	t := &committedTree{root: root, entries: map[string]committedNode{".": {mode: "040000", kind: "tree"}}, bodies: map[string]string{}}
	if raw != "" && !strings.HasSuffix(raw, "\x00") {
		return nil, fmt.Errorf("incomplete NUL-delimited tree")
	}
	for _, record := range strings.Split(raw, "\x00") {
		if record == "" {
			continue
		}
		header, relative, ok := strings.Cut(record, "\t")
		fields := strings.Fields(header)
		if !ok || relative == "" || len(fields) != 3 || !fullObjectID(fields[2]) || path.IsAbs(relative) {
			return nil, fmt.Errorf("invalid tree metadata record")
		}
		for _, component := range strings.Split(relative, "/") {
			if component == "" || component == "." || component == ".." {
				return nil, fmt.Errorf("invalid tree path %q", relative)
			}
		}
		if _, ok := t.entries[relative]; ok {
			return nil, fmt.Errorf("duplicate tree path %q", relative)
		}
		t.entries[relative] = committedNode{mode: fields[0], kind: fields[1], oid: fields[2]}
	}
	return t, nil
}
func hiddenTreePath(relative string) bool {
	for _, part := range strings.Split(relative, "/") {
		if part == ".git" || part == ".fireweed" {
			return true
		}
	}
	return false
}
func (t *committedTree) blob(oid string) (string, error) {
	if body, ok := t.bodies[oid]; ok {
		return body, nil
	}
	if !t.objects[oid] {
		return "", fmt.Errorf("selected metadata blob %s is unavailable locally; restore local metadata objects", oid)
	}
	body, err := baselineGit(t.root, "cat-file", "blob", oid)
	if err == nil {
		t.bodies[oid] = body
	}
	return body, err
}

// Resolve link targets as virtual tree components. Expanding a link is recursive;
// its remaining caller components are processed only after its target resolves.
// Thus a/../b also respects a directory link at a, without filesystem access.
func (t *committedTree) resolve(relative string, active map[string]bool) (string, committedNode, error) {
	if strings.ContainsAny(relative, "\x00\\") {
		return "", committedNode{}, fmt.Errorf("invalid committed path %q", relative)
	}
	if path.IsAbs(relative) || filepath.IsAbs(relative) {
		root := filepath.ToSlash(t.root)
		prefix := root
		if !strings.HasSuffix(prefix, "/") {
			prefix += "/"
		}
		if relative == root {
			relative = "."
		} else if strings.HasPrefix(relative, prefix) {
			relative = strings.TrimPrefix(relative, prefix)
		} else {
			return "", committedNode{}, fmt.Errorf("committed link escapes the canonical checkout")
		}
	}
	current := "."
	for _, part := range strings.Split(relative, "/") {
		if part == "" {
			continue
		}
		if t.entries[current].kind != "tree" {
			return "", committedNode{}, fmt.Errorf("committed path parent %q is not a directory", current)
		}
		if part == "." {
			continue
		}
		if part == ".." {
			if current == "." {
				return "", committedNode{}, fmt.Errorf("committed link escapes the selected tree")
			}
			current = path.Dir(current)
			continue
		}
		next := path.Join(current, part)
		node, ok := t.entries[next]
		if !ok {
			return "", committedNode{}, fmt.Errorf("missing committed path %q: %w", next, os.ErrNotExist)
		}
		if node.mode != "120000" {
			current = next
			continue
		}
		if active[next] {
			return "", committedNode{}, fmt.Errorf("committed symlink loop at %q", next)
		}
		target, err := t.blob(node.oid)
		if err != nil {
			return "", committedNode{}, err
		}
		if target == "" {
			return "", committedNode{}, fmt.Errorf("empty committed symlink at %q", next)
		}
		active[next] = true
		if !path.IsAbs(target) && !filepath.IsAbs(target) {
			target = path.Dir(next) + "/" + target
		}
		resolved, _, err := t.resolve(target, active)
		delete(active, next)
		if err != nil {
			// A broken link is an error, not an ordinarily absent optional directory.
			return "", committedNode{}, fmt.Errorf("repair committed symlink %q: %v", next, err)
		}
		current = resolved
	}
	return current, t.entries[current], nil
}
func (t *committedTree) file(relative string) (string, error) {
	_, node, err := t.resolve(relative, map[string]bool{})
	if err != nil {
		return "", err
	}
	if node.kind != "blob" || (node.mode != "100644" && node.mode != "100755") {
		return "", fmt.Errorf("selected metadata %q is not a regular committed blob", relative)
	}
	return t.blob(node.oid)
}
func (t *committedTree) changes(d *Data, logical string, prefixes map[string]string) []contracts.Change {
	out := []contracts.Change{}
	physical, node, err := t.resolve(logical, map[string]bool{})
	if errors.Is(err, os.ErrNotExist) {
		return out
	}
	if err != nil {
		d.add("C05", "error", logical, "Restore a safe committed change directory: "+err.Error())
		return out
	}
	if node.kind != "tree" {
		d.add("C05", "error", logical, "Restore a committed directory for changes.")
		return out
	}
	names := []string{}
	for relative, entry := range t.entries {
		if path.Dir(relative) == physical && entry.kind != "tree" && path.Ext(relative) == ".md" {
			names = append(names, path.Base(relative))
		}
	}
	sort.Strings(names)
	for _, name := range names {
		relative := logical + "/" + name
		text, err := t.file(relative)
		if err != nil {
			d.add("C05", "error", relative, "Restore a safe readable committed change: "+err.Error())
			continue
		}
		doc, err := contracts.ParseMarkdown(text)
		if err != nil {
			d.add("C05", "error", relative, "Correct the change document: "+err.Error())
			continue
		}
		target, _ := doc.Metadata["module"].(string)
		prefix, ok := prefixes[target]
		if !ok {
			d.add("C05", "error", relative, fmt.Sprintf("Target module %q is unknown; select an existing valid module.", target))
			continue
		}
		validation := text
		if logical == "changes/archive" {
			validation, err = changeview.ArchiveText(doc)
			if err != nil {
				d.add("C05", "error", relative, "Correct the archived change document: "+err.Error())
				continue
			}
		}
		ch, err := contracts.ParseChange(validation, relative, prefix)
		if err != nil {
			d.add(changeview.ErrorCode(err), "error", relative, "Correct the change document for its target module; declare each live handle locally: "+err.Error())
			continue
		}
		ch.Document = doc
		out = append(out, ch)
	}
	sort.Slice(out, func(i, j int) bool {
		a, b := out[i], out[j]
		if a.Number != b.Number {
			return a.Number < b.Number
		}
		return a.Slug < b.Slug
	})
	return out
}
