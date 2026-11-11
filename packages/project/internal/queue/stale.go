package queue

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func preceding(root, target string) (model.Change, *projection, error) {
	snapshot, err := fixture.Load(root)
	if err != nil {
		return model.Change{}, nil, err
	}
	changes, err := readQueue(root, snapshot)
	if err != nil {
		if errorCode(err, "") == "C13" {
			return model.Change{}, nil, contracts.CodedError{Code: "C13", Subject: target, Cause: err}
		}
		return model.Change{}, nil, err
	}
	index, err := resolve(changes, target)
	if err != nil {
		return model.Change{}, nil, err
	}
	p, err := newProjection(snapshot)
	if err != nil {
		return model.Change{}, nil, err
	}
	for _, change := range changes[:index] {
		if err := p.apply(change); err != nil {
			return model.Change{}, nil, err
		}
	}
	return changes[index], p, nil
}

func changeBases(change model.Change) (map[string]any, error) {
	raw, exists := change.Document.Frontmatter["bases"]
	if !exists {
		return map[string]any{}, nil
	}
	bases, ok := raw.(map[string]any)
	if !ok {
		return nil, contracts.CodedError{Code: "C08", Subject: change.Path, Cause: fmt.Errorf("%s bases must be a mapping", change.Path)}
	}
	return bases, nil
}

func edited(change model.Change) []contracts.Item {
	items := make([]contracts.Item, 0, len(change.Modified)+len(change.Removed))
	items = append(items, change.Modified...)
	return append(items, change.Removed...)
}

func absentEditError(path, id string) error {
	return contracts.CodedError{Code: "C08", Subject: path, Cause: fmt.Errorf("%s: ID %s absent from preceding projection", path, id)}
}

// StaleSpecs reports recorded bases that differ from the preceding projection.
func StaleSpecs(root, target string) ([]model.StaleSpec, error) {
	change, p, err := preceding(root, target)
	if err != nil {
		return nil, err
	}
	bases, err := changeBases(change)
	if err != nil {
		return nil, err
	}
	stale := make([]model.StaleSpec, 0)
	for _, edit := range edited(change) {
		raw, exists := bases[edit.ID]
		if !exists {
			continue
		}
		recorded := fmt.Sprint(raw)
		current, ok := p.find(change.Module, edit.ID)
		if !ok {
			return nil, absentEditError(change.Path, edit.ID)
		}
		if recorded != contracts.SpecHash(current) {
			stale = append(stale, model.StaleSpec{ID: edit.ID, RecordedHash: recorded, Current: contracts.RenderItem(current)})
		}
	}
	if len(stale) == 0 {
		return stale, nil
	}
	if before, ok := stampedProjection(root, change, bases); ok {
		for i := range stale {
			item, found := before.find(change.Module, stale[i].ID)
			if found && contracts.SpecHash(item) == stale[i].RecordedHash {
				stale[i].Before = contracts.RenderItem(item)
				stale[i].Recovered = true
			}
		}
	}
	return stale, nil
}

// StampMissing fills absent bases without changing recorded values.
func StampMissing(root, target string) error {
	change, p, err := preceding(root, target)
	if err != nil {
		return err
	}
	old, err := changeBases(change)
	if err != nil {
		return err
	}
	bases := make(map[string]any, len(old)+len(edited(change)))
	for id, hash := range old {
		bases[id] = hash
	}
	missing := false
	for _, edit := range edited(change) {
		base, ok := p.find(change.Module, edit.ID)
		if !ok {
			return absentEditError(change.Path, edit.ID)
		}
		if _, exists := bases[edit.ID]; !exists {
			bases[edit.ID] = contracts.SpecHash(base)
			missing = true
		}
	}
	if !missing {
		return nil
	}
	change.Document.Frontmatter["bases"] = bases
	content, err := contracts.RenderDocument(change.Document)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(root, filepath.FromSlash(change.Path)), []byte(content), 0644)
}

// stampedProjection uses the commit that introduced the current bases mapping.
func stampedProjection(root string, change model.Change, bases map[string]any) (*projection, bool) {
	out, err := staleGit(root, "log", "--format=%H", "HEAD", "--", change.Path)
	if err != nil {
		return nil, false
	}
	commits := strings.Fields(string(out))
	if len(commits) == 0 {
		return nil, false
	}
	var stamped string
	for i, commit := range commits {
		data, err := gitstore.ReadAt(root, commit, change.Path)
		if err != nil {
			return nil, false
		}
		doc, err := contracts.ParseDocument(string(data))
		if err != nil {
			return nil, false
		}
		recorded, exists := doc.Frontmatter["bases"]
		if !exists {
			recorded = map[string]any{}
		}
		if i == 0 && !reflect.DeepEqual(recorded, bases) {
			return nil, false
		}
		if !reflect.DeepEqual(recorded, bases) {
			break
		}
		stamped = commit
	}
	if stamped == "" {
		return nil, false
	}
	snapshot, err := fixture.LoadAt(root, stamped)
	if err != nil {
		return nil, false
	}
	changes, err := queueAt(root, stamped)
	if err != nil {
		return nil, false
	}
	p, err := newProjection(snapshot)
	if err != nil {
		return nil, false
	}
	for _, historical := range changes {
		if historical.Path == change.Path {
			return p, true
		}
		if err := p.apply(historical); err != nil {
			return nil, false
		}
	}
	return nil, false
}

func queueAt(root, commit string) ([]model.Change, error) {
	out, err := staleGit(root, "ls-tree", "-r", "--name-only", "-z", commit, "--", "changes")
	if err != nil {
		return nil, err
	}
	var changes []model.Change
	for _, name := range bytes.Split(out, []byte{0}) {
		path := string(name)
		if !strings.HasPrefix(path, "changes/") || strings.Contains(strings.TrimPrefix(path, "changes/"), "/") {
			continue
		}
		parts := changeName.FindStringSubmatch(filepath.Base(path))
		if parts == nil {
			continue
		}
		number, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		data, err := gitstore.ReadAt(root, commit, path)
		if err != nil {
			return nil, err
		}
		doc, err := contracts.ParseDocument(string(data))
		if err != nil {
			return nil, err
		}
		module, ok := doc.Frontmatter["module"].(string)
		if !ok {
			return nil, fmt.Errorf("%s has no module", path)
		}
		change := model.Change{Number: number, Slug: parts[2], Path: path, Module: module, Document: doc}
		for _, section := range []struct {
			name string
			into *[]contracts.Item
		}{
			{"Added", &change.Added}, {"Modified", &change.Modified}, {"Removed", &change.Removed},
		} {
			body, _ := contracts.Section(doc.Body, section.name)
			items, err := contracts.ParseItems(body)
			if err != nil {
				return nil, err
			}
			*section.into = items
		}
		changes = append(changes, change)
	}
	sort.Slice(changes, func(i, j int) bool {
		if changes[i].Number != changes[j].Number {
			return changes[i].Number < changes[j].Number
		}
		return changes[i].Slug < changes[j].Slug
	})
	return changes, nil
}

func staleGit(root string, args ...string) ([]byte, error) {
	argv := append([]string{"git", "-C", root, "-c", "core.hooksPath=/dev/null", "-c", "core.fsmonitor=false"}, args...)
	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Args: argv, Timeout: 30 * time.Second,
		Env: map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0"},
	})
	if err != nil {
		return nil, fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, result.Stderr)
	}
	return []byte(result.Stdout), nil
}
