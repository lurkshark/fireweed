// Package projection transforms typed snapshots without repository access.
package projection

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/changeview"
	"github.com/lurkshark/fireweed/packages/project/internal/inspection"
)

func Apply(module contracts.Module, change contracts.Change) (contracts.Module, error) {
	fail := func(err error) (contracts.Module, error) { return contracts.Module{}, err }
	if module.Name != change.Module {
		return fail(fmt.Errorf("change targets %q, not module %q", change.Module, module.Name))
	}
	m, ch := clone(module), clone(change)
	checked, err := validateChange(m, ch)
	if err != nil {
		return fail(err)
	}
	// Parsing is validation only: typed item bodies remain the caller's exact values.
	current := map[string]bool{}
	for _, item := range m.Specs {
		if _, err := contracts.ParseItems(items([]contracts.Spec{item}), m.IDPrefix, true); err != nil {
			return fail(fmt.Errorf("invalid existing item %s: %w", item.ID, err))
		}
		if !strings.HasPrefix(item.ID, "new:") {
			if current[item.ID] {
				return fail(fmt.Errorf("duplicate current ID %s", item.ID))
			}
			current[item.ID] = true
		}
	}
	modified := map[string]contracts.Spec{}
	removed := map[string]bool{}
	for _, item := range ch.Modified {
		if !current[item.ID] {
			return fail(fmt.Errorf("modified target %s is missing", item.ID))
		}
		modified[item.ID] = item
	}
	for _, item := range ch.Removed {
		if !current[item.ID] {
			return fail(fmt.Errorf("removed target %s is missing", item.ID))
		}
		removed[item.ID] = true
	}
	for _, item := range ch.Added {
		if !strings.HasPrefix(item.ID, "new:") && current[item.ID] {
			return fail(fmt.Errorf("added ID %s already exists", item.ID))
		}
	}
	doc := moduleDocument(m)
	for key, value := range checked.Set {
		if key == "hooks" {
			hooks := doc.Metadata["hooks"].(map[string]any)
			for name, command := range value.(map[string]any) {
				hooks[name] = command
			}
		} else {
			doc.Metadata[key] = clone(value)
		}
	}
	metaDoc := clone(doc)
	section(&metaDoc, "Specs", "")
	rendered, err := contracts.RenderMarkdown(metaDoc)
	if err != nil {
		return fail(err)
	}
	resolved, err := contracts.ParseModule(rendered, m.Dir, contracts.Zones{}, contracts.Hooks{})
	if err != nil {
		return fail(fmt.Errorf("invalid projected metadata: %w", err))
	}
	// Only named overrides replace values. No project defaults are consulted.
	for key := range ch.Set {
		switch key {
		case "depends_on":
			m.DependsOn = resolved.DependsOn
		case "interface":
			m.Interface = resolved.Interface
		case "evals":
			m.Evals = resolved.Evals
		case "disposable":
			m.Disposable = resolved.Disposable
		case "preserve":
			m.Preserve = resolved.Preserve
		case "append_only":
			m.AppendOnly = resolved.AppendOnly
		case "boundaries_opt_out":
			m.BoundariesOptOut = resolved.BoundariesOptOut
		case "hooks":
			m.Hooks = resolved.Hooks
		}
	}
	specs := make([]contracts.Spec, 0, len(m.Specs)+len(ch.Added))
	for _, item := range m.Specs {
		if removed[item.ID] {
			continue
		}
		if replacement, ok := modified[item.ID]; ok {
			item = replacement
		}
		specs = append(specs, item)
	}
	m.Specs = append(specs, ch.Added...)
	// Rebuild known metadata from the resolved typed values, including partial hooks.
	m.Document = doc
	m.Document = moduleDocument(m)
	return m, nil
}

func Project(snapshot inspection.Data, before string) (inspection.Data, error) {
	s := clone(snapshot)
	if before != "" {
		if _, err := selected(s.Changes, before); err != nil {
			return inspection.Data{}, err
		}
	}
	queue := append([]contracts.Change{}, s.Changes...)
	sort.SliceStable(queue, func(i, j int) bool {
		if queue[i].Number != queue[j].Number {
			return queue[i].Number < queue[j].Number
		}
		return queue[i].Slug < queue[j].Slug
	})
	for _, ch := range queue {
		if before != "" && ch.Slug == before {
			break
		}
		i, err := target(s.Modules, ch.Module)
		if err != nil {
			return inspection.Data{}, fmt.Errorf("change %s: %w", ch.Slug, err)
		}
		m, err := Apply(s.Modules[i], ch)
		if err != nil {
			return inspection.Data{}, fmt.Errorf("change %s: %w", ch.Slug, err)
		}
		s.Modules[i] = m
	}
	return s, nil
}

func Bases(snapshot inspection.Data, slug string) (map[string]string, error) {
	i, err := selected(snapshot.Changes, slug)
	if err != nil {
		return nil, err
	}
	ch := snapshot.Changes[i]
	preceding, err := Project(snapshot, slug)
	if err != nil {
		return nil, err
	}
	m, err := target(preceding.Modules, ch.Module)
	if err != nil {
		return nil, err
	}
	specs := map[string]contracts.Spec{}
	for _, item := range preceding.Modules[m].Specs {
		if !strings.HasPrefix(item.ID, "new:") {
			specs[item.ID] = item
		}
	}
	result := map[string]string{}
	for _, group := range [][]contracts.Spec{ch.Modified, ch.Removed} {
		for _, item := range group {
			base, ok := specs[item.ID]
			if !ok {
				return nil, fmt.Errorf("base target %s is missing before %s", item.ID, slug)
			}
			result[item.ID] = contracts.SpecHash(base)
		}
	}
	return result, nil
}

func selected(changes []contracts.Change, slug string) (int, error) {
	found := -1
	for i, ch := range changes {
		if ch.Slug == slug && slug != "" {
			if found >= 0 {
				return -1, fmt.Errorf("change selector %q is ambiguous", slug)
			}
			found = i
		}
	}
	if found < 0 {
		return -1, fmt.Errorf("change selector %q is unknown", slug)
	}
	return found, nil
}
func target(modules []contracts.Module, name string) (int, error) {
	found := -1
	for i, m := range modules {
		if m.Name == name {
			if found >= 0 {
				return -1, fmt.Errorf("module %q is ambiguous", name)
			}
			found = i
		}
	}
	if found < 0 {
		return -1, fmt.Errorf("target module %q is missing", name)
	}
	return found, nil
}

// Build a canonical change for public contract validation. Arbitrary author
// sections are retained so all handle references must be declared locally.
func validateChange(m contracts.Module, ch contracts.Change) (contracts.Change, error) {
	text, err := changeview.Text(ch)
	if err != nil {
		return contracts.Change{}, err
	}
	parsed, err := contracts.ParseChange(text, changeview.Path(ch), m.IDPrefix)
	if err != nil {
		return parsed, fmt.Errorf("invalid typed change: %w", err)
	}
	return parsed, nil
}
func moduleDocument(m contracts.Module) contracts.Document {
	d := clone(m.Document)
	if d.Metadata == nil {
		d.Metadata = map[string]any{}
	}
	d.Metadata["name"] = m.Name
	d.Metadata["id_prefix"] = m.IDPrefix
	lists := map[string][]string{"depends_on": m.DependsOn, "interface": m.Interface, "evals": m.Evals, "disposable": m.Disposable, "preserve": m.Preserve, "append_only": m.AppendOnly}
	for key, list := range lists {
		if list == nil {
			delete(d.Metadata, key)
		} else {
			d.Metadata[key] = clone(list)
		}
	}
	d.Metadata["hooks"] = map[string]any{"eval": m.Hooks.Eval, "boundaries": m.Hooks.Boundaries, "setup": m.Hooks.Setup}
	d.Metadata["boundaries_opt_out"] = m.BoundariesOptOut
	section(&d, "Specs", items(m.Specs))
	if m.Decisions != "" || hasSection(d, "Decisions") {
		section(&d, "Decisions", m.Decisions)
	}
	return d
}
func hasSection(d contracts.Document, name string) bool {
	for _, s := range d.Sections {
		if s.Name == name {
			return true
		}
	}
	return false
}
func section(d *contracts.Document, name, body string) {
	for i, s := range d.Sections {
		if s.Name == name {
			d.Sections[i].Body = body
			return
		}
	}
	d.Sections = append(d.Sections, contracts.Section{Name: name, Body: body})
}
func items(specs []contracts.Spec) string {
	var out strings.Builder
	for i, item := range specs {
		if i > 0 {
			out.WriteString("\n\n")
		}
		out.WriteString("### " + item.ID)
		if item.Tag != "" {
			out.WriteString(" " + item.Tag)
		}
		if item.Body != "" {
			out.WriteString("\n\n" + item.Body)
		}
	}
	return out.String()
}

// Snapshot types consist of exported records, interfaces, maps and slices.
// Preserve nil/empty distinctions and scalar types while severing all aliases.
func clone[T any](value T) T { return copyValue(reflect.ValueOf(value)).Interface().(T) }
func copyValue(v reflect.Value) reflect.Value {
	if !v.IsValid() {
		return v
	}
	switch v.Kind() {
	case reflect.Interface:
		out := reflect.New(v.Type()).Elem()
		if !v.IsNil() {
			out.Set(copyValue(v.Elem()))
		}
		return out
	case reflect.Ptr:
		if v.IsNil() {
			return reflect.Zero(v.Type())
		}
		out := reflect.New(v.Type().Elem())
		out.Elem().Set(copyValue(v.Elem()))
		return out
	case reflect.Map:
		if v.IsNil() {
			return reflect.Zero(v.Type())
		}
		out := reflect.MakeMapWithSize(v.Type(), v.Len())
		for _, k := range v.MapKeys() {
			out.SetMapIndex(copyValue(k), copyValue(v.MapIndex(k)))
		}
		return out
	case reflect.Slice:
		if v.IsNil() {
			return reflect.Zero(v.Type())
		}
		out := reflect.MakeSlice(v.Type(), v.Len(), v.Len())
		for i := 0; i < v.Len(); i++ {
			out.Index(i).Set(copyValue(v.Index(i)))
		}
		return out
	case reflect.Struct, reflect.Array:
		out := reflect.New(v.Type()).Elem()
		if v.Kind() == reflect.Array {
			for i := 0; i < v.Len(); i++ {
				out.Index(i).Set(copyValue(v.Index(i)))
			}
		} else {
			for i := 0; i < v.NumField(); i++ {
				out.Field(i).Set(copyValue(v.Field(i)))
			}
		}
		return out
	default:
		return v
	}
}
