package queue

import (
	"fmt"
	"maps"
	"path"
	"slices"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/internal/source"
	"github.com/lurkshark/fireweed/packages/project/model"
)

type projection struct {
	snapshot model.Snapshot
	items    [][]contracts.Item
	modules  map[string]int
}

func newProjection(snapshot model.Snapshot) (*projection, error) {
	if err := fixture.ValidateGraph(snapshot.Modules); err != nil {
		return nil, source.Failure("C02", "", err)
	}
	p := &projection{
		snapshot: snapshot,
		items:    make([][]contracts.Item, len(snapshot.Modules)),
		modules:  make(map[string]int, len(snapshot.Modules)),
	}
	for i, module := range snapshot.Modules {
		items, err := source.Items(module.Document, "Specs")
		if err != nil {
			return nil, source.DocumentError(path.Join(module.Dir, "SPEC.md"), err)
		}
		p.modules[module.Name] = i
		p.items[i] = items
	}
	return p, nil
}

func Project(root, through string) (model.Snapshot, error) {
	snapshot, err := fixture.Load(root)
	if err != nil {
		return model.Snapshot{}, err
	}
	changes, err := workingQueue(snapshot)
	if err != nil {
		return model.Snapshot{}, err
	}
	if through != "" {
		index, err := selectChange(changes, through)
		if err != nil {
			return model.Snapshot{}, err
		}
		changes = changes[:index+1]
	}
	p, err := projectChanges(snapshot, changes)
	if err != nil {
		return model.Snapshot{}, err
	}
	return p.snapshot, nil
}

func projectChanges(snapshot model.Snapshot, changes []model.Change) (*projection, error) {
	p, err := newProjection(snapshot)
	if err != nil {
		return nil, err
	}
	for _, change := range changes {
		if err := p.apply(change); err != nil {
			return nil, err
		}
	}
	return p, nil
}

func (p *projection) apply(change model.Change) error {
	index := p.modules[change.Module]
	module := p.snapshot.Modules[index]
	if raw, present := change.Document.Frontmatter["set"]; present {
		updated, err := applySet(module, raw)
		if err != nil {
			return source.FileError("C12", change.Path, err)
		}
		p.snapshot.Modules[index] = updated
		if err := fixture.ValidateGraph(p.snapshot.Modules); err != nil {
			p.snapshot.Modules[index] = module
			return source.FileError("C02", change.Path, err)
		}
		module = updated
	}
	if len(change.Added)+len(change.Modified)+len(change.Removed) == 0 {
		return nil
	}
	items, err := editSpecs(p.items[index], change)
	if err != nil {
		return err
	}
	p.items[index] = items
	var section strings.Builder
	section.WriteByte('\n')
	for i, item := range items {
		text := contracts.RenderItem(item)
		section.WriteString(text)
		if i < len(items)-1 && !strings.HasSuffix(text, "\n") {
			section.WriteByte('\n')
		}
	}
	body, _ := contracts.Section(module.Document.Body, "Specs")
	updated := section.String()
	if !strings.HasSuffix(updated, "\n") && (body == "" || strings.HasSuffix(body, "\n")) {
		updated += "\n"
	}
	module.Document.Body = contracts.ReplaceSection(module.Document.Body, "Specs", updated)
	p.snapshot.Modules[index] = module
	return nil
}

func applySet(module model.Module, raw any) (model.Module, error) {
	fields, ok := raw.(map[string]any)
	if !ok {
		return model.Module{}, fmt.Errorf("set must be a mapping")
	}
	doc := module.Document
	doc.Frontmatter = maps.Clone(doc.Frontmatter)
	for _, key := range slices.Sorted(maps.Keys(fields)) {
		switch key {
		case "depends_on", "interface", "evals", "disposable", "preserve", "append_only", "boundaries_opt_out", "hooks":
			doc.Frontmatter[key] = fields[key]
		default:
			return model.Module{}, fmt.Errorf("invalid set field %q", key)
		}
	}
	return fixture.ReadModule(module.Dir, doc, module.Zones)
}

func baseEdits(change model.Change) []contracts.Item {
	return slices.Concat(change.Modified, change.Removed)
}

func assignedItems(items []contracts.Item) map[string]contracts.Item {
	assigned := make(map[string]contracts.Item, len(items))
	for _, item := range items {
		if !strings.HasPrefix(item.ID, "new:") {
			assigned[item.ID] = item
		}
	}
	return assigned
}

func missingSpec(change model.Change, id string) error {
	return source.FileError("C08", change.Path, fmt.Errorf("spec %s is absent from the preceding projection", id))
}

func editSpecs(previous []contracts.Item, change model.Change) ([]contracts.Item, error) {
	assigned := assignedItems(previous)
	for _, edit := range baseEdits(change) {
		if _, exists := assigned[edit.ID]; !exists {
			return nil, missingSpec(change, edit.ID)
		}
	}
	replaced := make(map[string]contracts.Item, len(change.Modified))
	removed := make(map[string]bool, len(change.Removed))
	for _, item := range change.Modified {
		replaced[item.ID] = item
	}
	for _, item := range change.Removed {
		removed[item.ID] = true
	}
	items := make([]contracts.Item, 0, len(previous)+len(change.Added))
	for _, item := range previous {
		if removed[item.ID] {
			continue
		}
		if update, exists := replaced[item.ID]; exists {
			item = update
		}
		items = append(items, item)
	}
	return append(items, change.Added...), nil
}
