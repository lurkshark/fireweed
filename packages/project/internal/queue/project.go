package queue

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

// An owner distinguishes equal handle text declared by separate changes.
type ownedItem struct {
	item  contracts.Item
	owner string
}

type projection struct {
	snapshot model.Snapshot
	items    map[string][]ownedItem
}

func newProjection(snapshot model.Snapshot) (*projection, error) {
	p := &projection{snapshot: snapshot, items: map[string][]ownedItem{}}
	for _, module := range snapshot.Modules {
		section, _ := contracts.Section(module.Document.Body, "Specs")
		items, err := contracts.ParseItems(section)
		if err != nil {
			return nil, withCause(itemErrorCode(err), fmt.Errorf("%s/SPEC.md: %w", module.Dir, err))
		}
		for _, item := range items {
			p.items[module.Name] = append(p.items[module.Name], ownedItem{item: item})
		}
	}
	return p, nil
}

func (p *projection) find(module, id string) (contracts.Item, bool) {
	for _, entry := range p.items[module] {
		if entry.item.ID == id && entry.owner == "" {
			return entry.item, true
		}
	}
	return contracts.Item{}, false
}

func (p *projection) apply(change model.Change) error {
	index := -1
	for i := range p.snapshot.Modules {
		if p.snapshot.Modules[i].Name == change.Module {
			index = i
			break
		}
	}
	if index < 0 {
		return withCause("C07", fmt.Errorf("unknown module %q", change.Module))
	}
	if err := applySet(&p.snapshot.Modules[index], change.Document.Frontmatter["set"]); err != nil {
		return withCause("C12", fmt.Errorf("%s: %w", change.Path, err))
	}
	items := p.items[change.Module]
	for _, edit := range change.Modified {
		found := false
		for i := range items {
			if items[i].item.ID == edit.ID && items[i].owner == "" {
				items[i].item = edit
				found = true
				break
			}
		}
		if !found {
			return withCause("C08", fmt.Errorf("%s: Modified ID %s absent from preceding projection", change.Path, edit.ID))
		}
	}
	for _, edit := range change.Removed {
		found := false
		for i := range items {
			if items[i].item.ID == edit.ID && items[i].owner == "" {
				items = append(items[:i], items[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			return withCause("C08", fmt.Errorf("%s: Removed ID %s absent from preceding projection", change.Path, edit.ID))
		}
	}
	for _, edit := range change.Added {
		owner := ""
		if strings.HasPrefix(edit.ID, "new:") {
			owner = change.Path
		}
		items = append(items, ownedItem{item: edit, owner: owner})
	}
	p.items[change.Module] = items
	if err := fixture.ValidateGraph(p.snapshot.Modules); err != nil {
		return withCause("C02", err)
	}
	return nil
}

func applySet(module *model.Module, raw any) error {
	if raw == nil {
		return nil
	}
	settings, ok := raw.(map[string]any)
	if !ok {
		return fmt.Errorf("set must be a mapping")
	}
	front := make(map[string]any, len(module.Document.Frontmatter)+len(settings))
	for key, value := range module.Document.Frontmatter {
		front[key] = value
	}
	for key, value := range settings {
		switch key {
		case "depends_on", "interface", "evals", "disposable", "preserve", "append_only":
			list, err := asStrings(value)
			if err != nil {
				return fmt.Errorf("set %s: %w", key, err)
			}
			switch key {
			case "depends_on":
				module.DependsOn = list
			case "interface":
				module.Zones.Interface = list
			case "evals":
				module.Zones.Evals = list
			case "disposable":
				module.Zones.Disposable = list
			case "preserve":
				module.Zones.Preserve = list
			case "append_only":
				module.Zones.AppendOnly = list
			}
		case "boundaries_opt_out", "hooks":
		default:
			return fmt.Errorf("set cannot change %q", key)
		}
		front[key] = value
	}
	module.Document.Frontmatter = front
	return nil
}

func asStrings(value any) ([]string, error) {
	values, ok := value.([]any)
	if !ok {
		return nil, fmt.Errorf("expected a string list")
	}
	result := make([]string, 0, len(values))
	for _, value := range values {
		text, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("expected a string list")
		}
		result = append(result, text)
	}
	return result, nil
}

func (p *projection) result() model.Snapshot {
	for i := range p.snapshot.Modules {
		module := &p.snapshot.Modules[i]
		var body strings.Builder
		body.WriteByte('\n')
		for j, entry := range p.items[module.Name] {
			if j > 0 && !strings.HasSuffix(body.String(), "\n\n") {
				body.WriteByte('\n')
			}
			body.WriteString(contracts.RenderItem(entry.item))
			if !strings.HasSuffix(body.String(), "\n") {
				body.WriteByte('\n')
			}
		}
		module.Document.Body = contracts.ReplaceSection(module.Document.Body, "Specs", body.String())
	}
	return p.snapshot
}

func Project(root, through string) (model.Snapshot, error) {
	snapshot, err := fixture.Load(root)
	if err != nil {
		return model.Snapshot{}, err
	}
	changes, err := readQueue(root, snapshot)
	if err != nil {
		return model.Snapshot{}, err
	}
	limit := len(changes)
	if through != "" {
		index, err := resolve(changes, through)
		if err != nil {
			return model.Snapshot{}, err
		}
		limit = index + 1
	}
	p, err := newProjection(snapshot)
	if err != nil {
		return model.Snapshot{}, err
	}
	for _, change := range changes[:limit] {
		if err := p.apply(change); err != nil {
			return model.Snapshot{}, err
		}
	}
	if err := fixture.ValidateGraph(p.snapshot.Modules); err != nil {
		return model.Snapshot{}, err
	}
	return p.result(), nil
}

func Stamp(root, target string) error {
	snapshot, err := fixture.Load(root)
	if err != nil {
		return err
	}
	changes, err := readQueue(root, snapshot)
	if err != nil {
		return err
	}
	index, err := resolve(changes, target)
	if err != nil {
		return err
	}
	p, err := newProjection(snapshot)
	if err != nil {
		return err
	}
	for _, preceding := range changes[:index] {
		if err := p.apply(preceding); err != nil {
			return err
		}
	}
	change := changes[index]
	bases := map[string]any{}
	if raw, exists := change.Document.Frontmatter["bases"]; exists {
		current, ok := raw.(map[string]any)
		if !ok {
			return contracts.CodedError{Code: "C08", Subject: change.Path, Cause: fmt.Errorf("%s bases must be a mapping", change.Path)}
		}
		for key, value := range current {
			bases[key] = value
		}
	}
	for _, edit := range append(append([]contracts.Item{}, change.Modified...), change.Removed...) {
		base, ok := p.find(change.Module, edit.ID)
		if !ok {
			return fmt.Errorf("%s: ID %s absent from preceding projection", change.Path, edit.ID)
		}
		bases[edit.ID] = contracts.SpecHash(base)
	}
	change.Document.Frontmatter["bases"] = bases
	content, err := contracts.RenderDocument(change.Document)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(root, filepath.FromSlash(change.Path)), []byte(content), 0644)
}
