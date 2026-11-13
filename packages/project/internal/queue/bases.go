package queue

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/internal/source"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func preceding(root, selector string) (model.Change, map[string]contracts.Item, error) {
	snapshot, err := fixture.Load(root)
	if err != nil {
		return model.Change{}, nil, err
	}
	changes, err := workingQueue(snapshot)
	if err != nil {
		return model.Change{}, nil, err
	}
	index, err := selectChange(changes, selector)
	if err != nil {
		return model.Change{}, nil, err
	}
	change := changes[index]
	p, err := projectChanges(snapshot, changes[:index])
	if err != nil {
		return model.Change{}, nil, err
	}
	return change, assignedItems(p.items[p.modules[change.Module]]), nil
}

func bases(change model.Change) (map[string]any, error) {
	value, exists := change.Document.Frontmatter["bases"]
	if !exists {
		return make(map[string]any), nil
	}
	if mapping, ok := value.(map[string]any); ok {
		return mapping, nil
	}
	return nil, source.FileError("C08", change.Path, fmt.Errorf("bases must be a mapping"))
}

func recordedHash(value any) string {
	switch value := value.(type) {
	case string:
		return value
	case int:
		// YAML decodes a digits-only unquoted hash as an integer.
		return strconv.Itoa(value)
	default:
		return ""
	}
}

func Stamp(root, selector string) error {
	return stamp(root, selector, false)
}

func StampMissing(root, selector string) error {
	return stamp(root, selector, true)
}

func stamp(root, selector string, missingOnly bool) error {
	change, previous, err := preceding(root, selector)
	if err != nil {
		return err
	}
	mapping, err := bases(change)
	if err != nil {
		return err
	}
	updated := false
	for _, edit := range baseEdits(change) {
		item, exists := previous[edit.ID]
		if !exists {
			return missingSpec(change, edit.ID)
		}
		recorded, exists := mapping[edit.ID]
		if missingOnly && exists {
			continue
		}
		hash := contracts.SpecHash(item)
		if recorded != hash {
			mapping[edit.ID] = hash
			updated = true
		}
	}
	if !updated {
		return nil
	}
	change.Document.Frontmatter["bases"] = mapping
	text, err := contracts.RenderDocument(change.Document)
	if err != nil {
		return err
	}
	if err := replaceChange(source.Path(root, change.Path), text); err != nil {
		return fmt.Errorf("stamp %s: %w", change.Path, err)
	}
	return nil
}

func replaceChange(file, text string) error {
	info, err := os.Stat(file)
	if err != nil {
		return err
	}
	temp, err := os.CreateTemp(path.Dir(file), ".fireweed-stamp-*")
	if err != nil {
		return err
	}
	defer os.Remove(temp.Name())
	err = temp.Chmod(info.Mode().Perm())
	if err == nil {
		_, err = temp.WriteString(text)
	}
	closeErr := temp.Close()
	if err != nil {
		return err
	}
	if closeErr != nil {
		return closeErr
	}
	return os.Rename(temp.Name(), file)
}

func StaleSpecs(root, selector string) ([]model.StaleSpec, error) {
	change, previous, err := preceding(root, selector)
	if err != nil {
		return nil, err
	}
	mapping, err := bases(change)
	if err != nil {
		return nil, err
	}

	stale := make([]model.StaleSpec, 0)
	for _, edit := range baseEdits(change) {
		value, recorded := mapping[edit.ID]
		item, exists := previous[edit.ID]
		hash := recordedHash(value)
		if !recorded || !exists || hash == contracts.SpecHash(item) {
			continue
		}
		stale = append(stale, model.StaleSpec{
			ID: edit.ID, RecordedHash: hash, Current: contracts.RenderItem(item),
		})
	}
	if len(stale) != 0 {
		historical := historicalBases(root, change.Path)
		for i := range stale {
			entry := &stale[i]
			if item, exists := historical[entry.ID]; exists && contracts.SpecHash(item) == entry.RecordedHash {
				entry.Before = contracts.RenderItem(item)
				entry.Recovered = true
			}
		}
	}
	return stale, nil
}

func historicalBases(root, file string) map[string]contracts.Item {
	revision := lastBasesRevision(root, file)
	if revision == "" {
		return nil
	}
	snapshot, err := fixture.LoadAt(root, revision)
	if err != nil {
		return nil
	}
	listing, err := source.Git(snapshot.Root, "ls-tree", "--name-only", "-z", revision+":changes")
	if err != nil {
		return nil
	}
	changes, err := readQueue(snapshot, strings.Split(listing, "\x00"), func(file string) ([]byte, error) {
		return gitstore.ReadAt(snapshot.Root, revision, file)
	})
	if err != nil {
		return nil
	}
	index, err := selectChange(changes, file)
	if err != nil {
		return nil
	}
	p, err := projectChanges(snapshot, changes[:index])
	if err != nil {
		return nil
	}
	return assignedItems(p.items[p.modules[changes[index].Module]])
}

func lastBasesRevision(root, file string) string {
	// Compare successive recorded versions on HEAD's first-parent history.
	// Commits changing only the document's prose leave the mapping equal.
	listing, err := source.Git(root, "log", "--first-parent", "--full-history", "--format=%H", "HEAD", "--", file)
	if err != nil {
		return ""
	}
	var latest any
	var revision string
	for _, commit := range strings.Fields(listing) {
		data, err := gitstore.ReadAt(root, commit, file)
		if err != nil {
			return revision
		}
		doc, err := contracts.ParseDocument(string(data))
		if err != nil {
			return ""
		}
		mapping := doc.Frontmatter["bases"]
		if revision == "" {
			if _, ok := mapping.(map[string]any); !ok {
				return ""
			}
			latest = mapping
		} else if !reflect.DeepEqual(latest, mapping) {
			break
		}
		revision = commit
	}
	return revision
}
