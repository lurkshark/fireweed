package queue

import (
	"cmp"
	"fmt"
	"os"
	"path"
	"slices"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/internal/source"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func Queue(root string) ([]model.Change, error) {
	snapshot, err := fixture.Load(root)
	if err != nil {
		return nil, err
	}
	return workingQueue(snapshot)
}

func workingQueue(snapshot model.Snapshot) ([]model.Change, error) {
	files, err := source.Files(snapshot.Root, "changes")
	if err != nil {
		return nil, err
	}
	return readQueue(snapshot, files, func(file string) ([]byte, error) {
		return os.ReadFile(source.Path(snapshot.Root, file))
	})
}

func readQueue(snapshot model.Snapshot, files []string, read func(string) ([]byte, error)) ([]model.Change, error) {
	changes := make([]model.Change, 0, len(files))
	slugs := make(map[string]bool)
	for _, file := range files {
		number, slug, ok := source.ChangeName(file)
		if !ok {
			continue
		}
		relative := path.Join("changes", file)
		if slugs[slug] {
			return nil, source.FileError("C13", relative, fmt.Errorf("duplicate pending change slug %q", slug))
		}
		slugs[slug] = true
		changes = append(changes, model.Change{Number: number, Slug: slug, Path: relative})
	}
	slices.SortFunc(changes, func(a, b model.Change) int {
		if byNumber := cmp.Compare(a.Number, b.Number); byNumber != 0 {
			return byNumber
		}
		return cmp.Compare(a.Slug, b.Slug)
	})

	modules := make(map[string]bool, len(snapshot.Modules))
	for _, module := range snapshot.Modules {
		modules[module.Name] = true
	}
	for i := range changes {
		change := &changes[i]
		data, err := read(change.Path)
		if err != nil {
			return nil, source.FileError("C03", change.Path, err)
		}
		change.Document, err = contracts.ParseDocument(string(data))
		if err != nil {
			return nil, source.DocumentError(change.Path, err)
		}
		change.Module, _ = change.Document.Frontmatter["module"].(string)
		if change.Module == "" || !modules[change.Module] {
			return nil, source.FileError("C07", change.Path, fmt.Errorf("unknown module %q", change.Module))
		}
		for _, section := range []struct {
			name  string
			items *[]contracts.Item
		}{{"Added", &change.Added}, {"Modified", &change.Modified}, {"Removed", &change.Removed}} {
			*section.items, err = source.Items(change.Document, section.name)
			if err != nil {
				return nil, source.DocumentError(change.Path, err)
			}
		}
	}
	return changes, nil
}

func selectChange(changes []model.Change, selector string) (int, error) {
	index := -1
	slugs := make([]string, 0, len(changes))
	for i, change := range changes {
		slugs = append(slugs, change.Slug)
		if selector == change.Slug || selector == path.Base(change.Path) || selector == change.Path {
			if index >= 0 {
				return -1, source.Failure("C13", selector, fmt.Errorf("ambiguous pending change %q", selector))
			}
			index = i
		}
	}
	if index >= 0 {
		return index, nil
	}
	return -1, source.Failure("U02", selector,
		fmt.Errorf("pending change %q not found; valid pending slugs: [%s]", selector, strings.Join(slugs, ", ")))
}
