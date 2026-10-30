package history

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
	snapshot, err := fixture.Load(root)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fixture.LoadCode(err), err)
	}

	var target *model.Module
	for i := range snapshot.Modules {
		if snapshot.Modules[i].Name == module {
			target = &snapshot.Modules[i]
			break
		}
	}
	if target == nil {
		return nil, fmt.Errorf("U02: unknown module %q", module)
	}

	result := make([]model.SpecHistoryEntry, 0)
	archiveDir := filepath.Join(root, "changes", "archive")
	files, err := os.ReadDir(archiveDir)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("C03: read changes/archive: %w", err)
	}
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".md" {
			continue
		}
		path := "changes/archive/" + file.Name()
		entry, err := readArchive(root, path)
		if err != nil {
			return nil, err
		}
		if entry.Module != module {
			continue
		}
		if id != "" {
			if !contains(entry.IDs, id) {
				continue
			}
			entry.IDs = []string{id}
		} else if len(entry.IDs) == 0 {
			continue
		}
		result = append(result, entry)
	}

	if id != "" && len(result) == 0 {
		body, _ := contracts.Section(target.Document.Body, "Specs")
		items, err := contracts.ParseItems(body)
		if err != nil {
			return nil, fmt.Errorf("C03: %s/SPEC.md Specs: %w", target.Dir, err)
		}
		found := false
		for _, item := range items {
			if item.ID == id {
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("U02: unknown spec ID %q", id)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		a, b := result[i], result[j]
		if a.ImplementationDate.IsZero() != b.ImplementationDate.IsZero() {
			return !a.ImplementationDate.IsZero()
		}
		if !a.ImplementationDate.Equal(b.ImplementationDate) {
			return a.ImplementationDate.After(b.ImplementationDate)
		}
		return a.Path < b.Path
	})
	return result, nil
}

func readArchive(root, path string) (model.SpecHistoryEntry, error) {
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(path)))
	if err != nil {
		return model.SpecHistoryEntry{}, fmt.Errorf("C03: read %s: %w", path, err)
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		return model.SpecHistoryEntry{}, fmt.Errorf("C03: parse %s: %w", path, err)
	}
	module, ok := doc.Frontmatter["module"].(string)
	if !ok || module == "" {
		return model.SpecHistoryEntry{}, fmt.Errorf("C03: %s has no module", path)
	}
	ids := make(map[string]bool)
	for _, name := range []string{"Added", "Modified"} {
		body, _ := contracts.Section(doc.Body, name)
		items, err := contracts.ParseItems(body)
		if err != nil {
			return model.SpecHistoryEntry{}, fmt.Errorf("C03: %s %s: %w", path, name, err)
		}
		for _, item := range items {
			ids[item.ID] = true
		}
	}
	entry := model.SpecHistoryEntry{Path: path, Module: module, IDs: make([]string, 0, len(ids))}
	for itemID := range ids {
		entry.IDs = append(entry.IDs, itemID)
	}
	sort.Strings(entry.IDs)
	entry.Motivation, _ = contracts.Section(doc.Body, "Motivation")
	entry.Decisions, _ = contracts.Section(doc.Body, "Decisions")
	entry.RemovedDecisions, _ = contracts.Section(doc.Body, "Removed decisions")
	entry.Provenance, _ = contracts.Section(doc.Body, "Provenance")
	entry.ImplementationDate = latestImplementationDate(entry.Provenance)
	return entry, nil
}

func latestImplementationDate(provenance string) time.Time {
	var latest time.Time
	for _, line := range strings.Split(provenance, "\n") {
		fields := strings.Fields(line)
		if len(fields) < 3 || fields[0] != "-" || fields[1] != "impl" {
			continue
		}
		date, err := time.Parse(time.RFC3339Nano, fields[2])
		if err == nil && date.After(latest) {
			latest = date
		}
	}
	return latest
}

func contains(ids []string, id string) bool {
	for _, candidate := range ids {
		if candidate == id {
			return true
		}
	}
	return false
}
