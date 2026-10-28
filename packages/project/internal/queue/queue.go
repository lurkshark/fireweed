package queue

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

var changeName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)

func Queue(root string) ([]model.Change, error) {
	snapshot, err := fixture.Load(root)
	if err != nil {
		return nil, err
	}
	return readQueue(root, snapshot)
}

func readQueue(root string, snapshot model.Snapshot) ([]model.Change, error) {
	known := map[string]bool{}
	for _, module := range snapshot.Modules {
		known[module.Name] = true
	}
	entries, err := os.ReadDir(filepath.Join(root, "changes"))
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, withCause("C03", fmt.Errorf("read changes: %w", err))
	}
	var changes []model.Change
	slugs := map[string]bool{}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		parts := changeName.FindStringSubmatch(entry.Name())
		if parts == nil {
			continue
		}
		number, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, withCause("C13", fmt.Errorf("invalid change number %s: %w", entry.Name(), err))
		}
		if slugs[parts[2]] {
			return nil, withCause("C13", fmt.Errorf("ambiguous change slug %q", parts[2]))
		}
		slugs[parts[2]] = true
		path := "changes/" + entry.Name()
		data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(path)))
		if err != nil {
			return nil, withCause("C03", fmt.Errorf("read %s: %w", path, err))
		}
		doc, err := contracts.ParseDocument(string(data))
		if err != nil {
			return nil, withCause("C03", fmt.Errorf("parse %s: %w", path, err))
		}
		module, ok := doc.Frontmatter["module"].(string)
		if !ok || !known[module] {
			return nil, withCause("C07", fmt.Errorf("%s names unknown module %q", path, module))
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
				return nil, withCause(itemErrorCode(err), fmt.Errorf("%s %s: %w", path, section.name, err))
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

func resolve(changes []model.Change, target string) (int, error) {
	for i, change := range changes {
		if target == change.Slug || target == filepath.Base(change.Path) || target == change.Path {
			return i, nil
		}
	}
	return -1, fmt.Errorf("pending change %q not found", target)
}
