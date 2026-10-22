package documents

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts/internal/config"
	"github.com/lurkshark/fireweed/packages/contracts/internal/specs"
	"github.com/lurkshark/fireweed/packages/contracts/model"
	"gopkg.in/yaml.v3"
)

// Fences shield example headings from the document structure.
type fence struct {
	character byte
	width     int
}

func (f *fence) live(line string) bool {
	trimmed := strings.TrimLeft(line, " ")
	if len(line)-len(trimmed) > 3 || len(trimmed) < 3 || (trimmed[0] != '`' && trimmed[0] != '~') {
		return f.width == 0
	}
	n := 0
	for n < len(trimmed) && trimmed[n] == trimmed[0] {
		n++
	}
	if n < 3 {
		return f.width == 0
	}
	if f.width == 0 {
		f.character, f.width = trimmed[0], n
		return false
	}
	if trimmed[0] == f.character && n >= f.width && strings.TrimSpace(trimmed[n:]) == "" {
		f.width = 0
	}
	return false
}
func outerBlankLines(text string) string {
	lines := strings.Split(text, "\n")
	start, end := 0, len(lines)
	for start < end && strings.TrimSpace(lines[start]) == "" {
		start++
	}
	for end > start && strings.TrimSpace(lines[end-1]) == "" {
		end--
	}
	return strings.Join(lines[start:end], "\n")
}
func Parse(text string) (model.Document, error) {
	result := model.Document{Metadata: map[string]any{}, Sections: []model.Section{}}
	lines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	if len(lines) == 0 || lines[0] != "---" {
		return result, fmt.Errorf("document must begin with YAML frontmatter")
	}
	end := 1
	for end < len(lines) && lines[end] != "---" {
		end++
	}
	if end == len(lines) {
		return result, fmt.Errorf("unclosed frontmatter")
	}
	if err := config.Decode(strings.Join(lines[1:end], "\n")+"\n", &result.Metadata); err != nil {
		return result, err
	}
	var scanner fence
	var body []string
	seen := map[string]bool{}
	finish := func() {
		text := outerBlankLines(strings.Join(body, "\n"))
		if len(result.Sections) == 0 {
			result.Preamble = text
		} else {
			result.Sections[len(result.Sections)-1].Body = text
		}
		body = nil
	}
	for _, line := range lines[end+1:] {
		if scanner.live(line) && strings.HasPrefix(line, "## ") {
			name := strings.TrimSpace(strings.TrimPrefix(line, "## "))
			if name == "" || seen[name] {
				return result, fmt.Errorf("empty or duplicate section %q", name)
			}
			seen[name] = true
			finish()
			result.Sections = append(result.Sections, model.Section{Name: name})
			continue
		}
		body = append(body, line)
	}
	finish()
	return result, nil
}
func Render(document model.Document) (string, error) {
	if document.Metadata == nil {
		return "", fmt.Errorf("metadata mapping is required")
	}
	meta, err := yaml.Marshal(document.Metadata)
	if err != nil {
		return "", err
	}
	text := "---\n" + string(meta) + "---\n"
	if document.Preamble != "" {
		text += "\n" + outerBlankLines(document.Preamble) + "\n"
	}
	seen := map[string]bool{}
	for _, section := range document.Sections {
		if section.Name == "" || seen[section.Name] {
			return "", fmt.Errorf("invalid section %q", section.Name)
		}
		seen[section.Name] = true
		text += "\n## " + section.Name + "\n"
		if section.Body != "" {
			text += "\n" + outerBlankLines(section.Body) + "\n"
		}
	}
	return text, nil
}
func section(document model.Document, name string) (string, bool) {
	for _, s := range document.Sections {
		if s.Name == name {
			return s.Body, true
		}
	}
	return "", false
}
func clone(list []string) []string {
	if list == nil {
		return nil
	}
	copy := make([]string, len(list))
	copy = append(copy[:0], list...)
	return copy
}
func cloneZones(z model.Zones) model.Zones {
	return model.Zones{Interface: clone(z.Interface), Evals: clone(z.Evals), Disposable: clone(z.Disposable), Preserve: clone(z.Preserve), AppendOnly: clone(z.AppendOnly)}
}
func decodeMetadata(document model.Document, target any) error {
	data, err := yaml.Marshal(document.Metadata)
	if err != nil {
		return err
	}
	return config.Decode(string(data), target)
}

var namePattern = regexp.MustCompile(`^[a-z0-9][a-z0-9-]*$`)
var prefixPattern = regexp.MustCompile(`^[A-Z][A-Z0-9]*$`)

func convert(items []specs.Item) []model.Spec {
	result := make([]model.Spec, len(items))
	for i, item := range items {
		result[i] = model.Spec{ID: item.ID, Tag: item.Tag, Body: item.Body}
	}
	return result
}
func Module(text, dir string, defaults model.Zones, hooks model.Hooks) (model.Module, error) {
	result := model.Module{Zones: cloneZones(defaults), Hooks: hooks, DependsOn: []string{}}
	document, err := Parse(text)
	if err != nil {
		return result, err
	}
	if err = decodeMetadata(document, &result); err != nil {
		return result, err
	}
	if !namePattern.MatchString(result.Name) || !prefixPattern.MatchString(result.IDPrefix) {
		return result, fmt.Errorf("valid name and id_prefix required")
	}
	body, ok := section(document, "Specs")
	if !ok {
		return result, fmt.Errorf("Specs section required")
	}
	items, err := specs.ParseItems(body, result.IDPrefix, false)
	if err != nil {
		return result, err
	}
	result.Dir = dir
	result.Specs = convert(items)
	result.Decisions, _ = section(document, "Decisions")
	result.Document = document
	return result, nil
}

var filenamePattern = regexp.MustCompile(`^changes/(?:archive/)?([0-9]{4,})-([a-z0-9][a-z0-9-]*)\.md$`)
var realID = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9]*-[0-9]{3,}$`)
var hashPattern = regexp.MustCompile(`^[0-9a-f]{8}$`)
var handleReference = regexp.MustCompile(`new:[a-z0-9-]+`)

func Change(text, path, prefix string) (model.Change, error) {
	result := model.Change{}
	document, err := Parse(text)
	if err != nil {
		return result, err
	}
	if err = decodeMetadata(document, &result); err != nil {
		return result, err
	}
	if !namePattern.MatchString(result.Module) {
		return result, fmt.Errorf("module required")
	}
	match := filenamePattern.FindStringSubmatch(path)
	if match == nil {
		return result, fmt.Errorf("invalid change filename %q", path)
	}
	number, err := strconv.Atoi(match[1])
	if err != nil {
		return result, err
	}
	result.Path, result.Slug, result.Number = path, match[2], number
	if result.Phase != "" && result.Phase != "evals" && result.Phase != "audit" && result.Phase != "impl" && result.Phase != "applied" {
		return result, fmt.Errorf("invalid phase %s", result.Phase)
	}
	motivation, _ := section(document, "Motivation")
	if strings.TrimSpace(motivation) == "" {
		return result, fmt.Errorf("Motivation required")
	}
	for _, entry := range []struct {
		name    string
		target  *[]model.Spec
		handles bool
	}{{"Added", &result.Added, true}, {"Modified", &result.Modified, false}} {
		body, _ := section(document, entry.name)
		items, err := specs.ParseItems(body, prefix, entry.handles)
		if err != nil {
			return result, fmt.Errorf("%s: %w", entry.name, err)
		}
		*entry.target = convert(items)
	}
	removed, _ := section(document, "Removed")
	result.Removed = []model.Spec{}
	for _, line := range strings.Split(removed, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		if !strings.HasPrefix(line, "### ") {
			return result, fmt.Errorf("Removed contains a body")
		}
		items, err := specs.ParseItems(line+"\nRemoval MUST be recorded.", prefix, false)
		if err != nil || len(items) != 1 {
			return result, fmt.Errorf("invalid Removed heading %q", line)
		}
		result.Removed = append(result.Removed, model.Spec{ID: items[0].ID, Tag: items[0].Tag})
	}
	declared := map[string]bool{}
	edited := map[string]bool{}
	for _, items := range [][]model.Spec{result.Added, result.Modified, result.Removed} {
		for _, item := range items {
			if edited[item.ID] {
				return result, fmt.Errorf("duplicate edit %s", item.ID)
			}
			edited[item.ID] = true
		}
	}
	for _, item := range result.Added {
		if strings.HasPrefix(item.ID, "new:") {
			declared[item.ID] = true
		}
	}
	for _, reference := range handleReference.FindAllString(text, -1) {
		if !declared[reference] {
			return result, fmt.Errorf("undeclared handle %s", reference)
		}
	}
	for id, hash := range result.Bases {
		if !realID.MatchString(id) || !hashPattern.MatchString(hash) {
			return result, fmt.Errorf("invalid base %s: %s", id, hash)
		}
	}
	for _, id := range result.Fixes {
		if !realID.MatchString(id) {
			return result, fmt.Errorf("invalid fixes spec ID %s", id)
		}
	}
	for _, id := range result.AlreadySatisfied {
		if !realID.MatchString(id) && !declared[id] {
			return result, fmt.Errorf("invalid already_satisfied spec ID %s", id)
		}
	}
	if err = validateSet(result.Set); err != nil {
		return result, err
	}
	switch {
	case result.Regenerate:
		result.Kind = "regenerate"
	case len(result.Added)+len(result.Modified)+len(result.Removed) > 0:
		result.Kind = "spec"
	case len(result.Fixes) > 0:
		result.Kind = "fix"
	default:
		result.Kind = "maintenance"
	}
	result.Document = document
	return result, nil
}
func validateSet(set map[string]any) error {
	allowed := map[string]bool{"depends_on": true, "interface": true, "evals": true, "disposable": true, "preserve": true, "append_only": true, "boundaries_opt_out": true, "hooks": true}
	for key := range set {
		if !allowed[key] {
			return fmt.Errorf("set cannot change %s", key)
		}
	}
	if set == nil {
		return nil
	}
	data, err := yaml.Marshal(set)
	if err != nil {
		return err
	}
	return config.Decode(string(data), &model.Module{})
}
