package history

import (
	"cmp"
	"fmt"
	"maps"
	"os"
	"path"
	"slices"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/internal/source"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
	snapshot, err := fixture.Load(root)
	if err != nil {
		return nil, err
	}
	index := slices.IndexFunc(snapshot.Modules, func(candidate model.Module) bool { return candidate.Name == module })
	if index < 0 {
		return nil, source.Failure("U02", module, fmt.Errorf("unknown module %q", module))
	}
	const archive = "changes/archive"
	files, err := source.Files(root, archive)
	if err != nil {
		return nil, err
	}
	entries := make([]model.SpecHistoryEntry, 0)
	for _, file := range files {
		if !strings.HasSuffix(file, ".md") {
			continue
		}
		entry, err := readEntry(root, path.Join(archive, file), module, id)
		if err != nil {
			return nil, err
		}
		if len(entry.IDs) != 0 {
			entries = append(entries, entry)
		}
	}
	if id != "" && len(entries) == 0 {
		current, err := source.Items(snapshot.Modules[index].Document, "Specs")
		if err != nil {
			return nil, err
		}
		if !slices.ContainsFunc(current, func(item contracts.Item) bool { return item.ID == id }) {
			return nil, source.Failure("U02", id, fmt.Errorf("unknown spec %q in module %q", id, module))
		}
	}
	slices.SortFunc(entries, func(a, b model.SpecHistoryEntry) int {
		if a.ImplementationDate.IsZero() != b.ImplementationDate.IsZero() {
			if a.ImplementationDate.IsZero() {
				return 1
			}
			return -1
		}
		if byDate := b.ImplementationDate.Compare(a.ImplementationDate); byDate != 0 {
			return byDate
		}
		return cmp.Compare(a.Path, b.Path)
	})
	return entries, nil
}

func readEntry(root, file, module, id string) (model.SpecHistoryEntry, error) {
	var entry model.SpecHistoryEntry
	data, err := os.ReadFile(source.Path(root, file))
	if err != nil {
		return entry, source.FileError("C03", file, err)
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		return entry, source.FileError("C03", file, err)
	}
	archiveModule, _ := doc.Frontmatter["module"].(string)
	ids := make(map[string]bool)
	for _, section := range []string{"Added", "Modified"} {
		items, err := source.Items(doc, section)
		if err != nil {
			return entry, source.FileError("C03", file, err)
		}
		for _, item := range items {
			if id == "" || item.ID == id {
				ids[item.ID] = true
			}
		}
	}
	if archiveModule != module || len(ids) == 0 {
		return entry, nil
	}
	entry.Path = file
	entry.Module = module
	entry.IDs = slices.Sorted(maps.Keys(ids))
	entry.Motivation, _ = contracts.Section(doc.Body, "Motivation")
	entry.Decisions, _ = contracts.Section(doc.Body, "Decisions")
	entry.RemovedDecisions, _ = contracts.Section(doc.Body, "Removed decisions")
	entry.Provenance, _ = contracts.Section(doc.Body, "Provenance")
	for _, line := range strings.Split(entry.Provenance, "\n") {
		fields := strings.Fields(line)
		if len(fields) < 3 || fields[0] != "-" || fields[1] != "impl" {
			continue
		}
		date, err := time.Parse(time.RFC3339, fields[2])
		if err == nil && (entry.ImplementationDate.IsZero() || date.After(entry.ImplementationDate)) {
			entry.ImplementationDate = date
		}
	}
	return entry, nil
}
