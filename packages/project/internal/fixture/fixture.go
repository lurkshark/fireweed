package fixture

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/model"
	"gopkg.in/yaml.v3"
)

type loadFailure struct {
	code string
	path string
	err  error
}

func (e *loadFailure) Error() string { return e.err.Error() }
func (e *loadFailure) Unwrap() error { return e.err }

// LoadCode identifies the diagnostic cause retained by a Load error.
func LoadCode(err error) string {
	var failure *loadFailure
	if errors.As(err, &failure) {
		return failure.code
	}
	return "C01"
}

func loadCause(code string, err error) error { return &loadFailure{code: code, err: err} }

// LoadPath identifies the document responsible for a Load failure.
func LoadPath(err error) string {
	var failure *loadFailure
	if errors.As(err, &failure) && failure.path != "" {
		return failure.path
	}
	return "fireweed.yaml"
}

func moduleCause(rel string, err error) error {
	code, _, ok := contracts.ErrorCode(err)
	if !ok {
		code = "C03"
	}
	path := rel + "/SPEC.md"
	return &loadFailure{
		code: code,
		path: path,
		err:  fmt.Errorf("parse %s: %w. Use valid frontmatter and unique spec item headings, then retry", path, err),
	}
}

func Load(root string) (model.Snapshot, error) {
	data, err := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
	if err != nil {
		return model.Snapshot{}, fmt.Errorf("read fireweed.yaml: %w", err)
	}
	config, patterns, defaults, err := parseConfig(data)
	if err != nil {
		return model.Snapshot{}, err
	}

	snapshot := model.Snapshot{Root: root, Config: config}
	err = filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if !entry.IsDir() {
			return nil
		}
		if entry.Name() == ".git" {
			return filepath.SkipDir
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if rel == "." || !matchesModule(patterns, rel) {
			return nil
		}
		data, err := os.ReadFile(filepath.Join(path, "SPEC.md"))
		if err != nil {
			return &loadFailure{code: "C03", path: rel + "/SPEC.md", err: fmt.Errorf("read %s/SPEC.md: %w", rel, err)}
		}
		module, err := parseModule(data, defaults, rel)
		if err != nil {
			return err
		}
		snapshot.Modules = append(snapshot.Modules, module)
		return nil
	})
	if err != nil {
		return model.Snapshot{}, err
	}
	return snapshot, nil
}

func parseConfig(data []byte) (map[string]any, []string, map[string]any, error) {
	var config map[string]any
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, nil, nil, fmt.Errorf("parse fireweed.yaml: %w", err)
	}
	if config == nil {
		return nil, nil, nil, fmt.Errorf("fireweed.yaml must be a mapping")
	}
	patterns, err := stringList(config["modules"], "fireweed.yaml modules")
	if err != nil {
		return nil, nil, nil, err
	}
	defaults := map[string]any{}
	if value, ok := config["module_defaults"]; ok {
		defaults, ok = value.(map[string]any)
		if !ok {
			return nil, nil, nil, fmt.Errorf("fireweed.yaml module_defaults must be a mapping")
		}
	}
	return config, patterns, defaults, nil
}

func parseModule(data []byte, defaults map[string]any, rel string) (model.Module, error) {
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		return model.Module{}, moduleCause(rel, err)
	}
	section, _ := contracts.Section(doc.Body, "Specs")
	if _, err := contracts.ParseItems(section); err != nil {
		return model.Module{}, moduleCause(rel, err)
	}
	name, ok := doc.Frontmatter["name"].(string)
	if !ok {
		return model.Module{}, loadCause("C12", fmt.Errorf("%s/SPEC.md name must be a string", rel))
	}
	prefix, ok := doc.Frontmatter["id_prefix"].(string)
	if !ok {
		return model.Module{}, loadCause("C12", fmt.Errorf("%s/SPEC.md id_prefix must be a string", rel))
	}
	var dependencies []string
	if value, ok := doc.Frontmatter["depends_on"]; ok {
		dependencies, err = stringList(value, rel+"/SPEC.md depends_on")
		if err != nil {
			return model.Module{}, loadCause("C12", err)
		}
	}
	zones, err := moduleZones(doc.Frontmatter, defaults, rel)
	if err != nil {
		return model.Module{}, loadCause("C12", err)
	}
	return model.Module{
		Name: name, IDPrefix: prefix, Dir: rel, DependsOn: dependencies,
		Zones: zones, Document: doc,
	}, nil
}

func matchesModule(patterns []string, dir string) bool {
	for _, pattern := range patterns {
		if contracts.MatchGlob(pattern, dir) {
			return true
		}
	}
	return false
}

func stringList(value any, label string) ([]string, error) {
	items, ok := value.([]any)
	if !ok {
		return nil, fmt.Errorf("%s must be a list of strings", label)
	}
	result := make([]string, 0, len(items))
	for _, item := range items {
		text, ok := item.(string)
		if !ok {
			return nil, fmt.Errorf("%s must be a list of strings", label)
		}
		result = append(result, text)
	}
	return result, nil
}

func moduleZones(frontmatter, defaults map[string]any, dir string) (contracts.Zones, error) {
	var zones contracts.Zones
	fields := []struct {
		name string
		into *[]string
	}{
		{"protected", &zones.Protected},
		{"evals", &zones.Evals},
		{"interface", &zones.Interface},
		{"append_only", &zones.AppendOnly},
		{"preserve", &zones.Preserve},
		{"disposable", &zones.Disposable},
	}
	for _, field := range fields {
		value, exists := frontmatter[field.name]
		label := dir + "/SPEC.md " + field.name
		if !exists {
			value, exists = defaults[field.name]
			label = "fireweed.yaml module_defaults " + field.name
		}
		if exists {
			list, err := stringList(value, label)
			if err != nil {
				return contracts.Zones{}, err
			}
			*field.into = list
		}
	}
	return zones, nil
}

func ValidateGraph(modules []model.Module) error {
	byName := make(map[string]model.Module, len(modules))
	prefixes := make(map[string]string, len(modules))
	for _, module := range modules {
		if module.Name == "" {
			return fmt.Errorf("empty module name")
		}
		if module.IDPrefix == "" {
			return fmt.Errorf("empty ID prefix for module %s", module.Name)
		}
		if _, exists := byName[module.Name]; exists {
			return fmt.Errorf("duplicate module name %q", module.Name)
		}
		if other, exists := prefixes[module.IDPrefix]; exists {
			return fmt.Errorf("duplicate ID prefix %q for %s and %s", module.IDPrefix, other, module.Name)
		}
		byName[module.Name] = module
		prefixes[module.IDPrefix] = module.Name
	}
	state := make(map[string]uint8, len(modules))
	var visit func(string) error
	visit = func(name string) error {
		state[name] = 1
		for _, dependency := range byName[name].DependsOn {
			if dependency == name {
				return fmt.Errorf("module %s has a self-dependency", name)
			}
			if _, exists := byName[dependency]; !exists {
				return fmt.Errorf("module %s depends on missing module %s", name, dependency)
			}
			if state[dependency] == 1 {
				return fmt.Errorf("dependency cycle between %s and %s", name, dependency)
			}
			if state[dependency] == 0 {
				if err := visit(dependency); err != nil {
					return err
				}
			}
		}
		state[name] = 2
		return nil
	}
	for _, module := range modules {
		if state[module.Name] == 0 {
			if err := visit(module.Name); err != nil {
				return err
			}
		}
	}
	return nil
}

var slugPattern = regexp.MustCompile(`^[a-z0-9-]+$`)
var changeNamePattern = regexp.MustCompile(`^([0-9]+)-(.+)\.md$`)

func NewChange(root, module, slug string) (string, error) {
	if !slugPattern.MatchString(slug) {
		return "", fmt.Errorf("invalid change slug %q", slug)
	}
	snapshot, err := Load(root)
	if err != nil {
		return "", err
	}
	found := false
	for _, item := range snapshot.Modules {
		if item.Name == module {
			found = true
			break
		}
	}
	if !found {
		return "", fmt.Errorf("unknown module %q", module)
	}
	maxNumber := 0
	for _, directory := range []string{"changes", "changes/archive"} {
		entries, err := os.ReadDir(filepath.Join(root, directory))
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return "", fmt.Errorf("read %s: %w", directory, err)
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			parts := changeNamePattern.FindStringSubmatch(entry.Name())
			if parts == nil {
				continue
			}
			if parts[2] == slug {
				return "", fmt.Errorf("change slug %q already exists in %s", slug, directory)
			}
			number, err := strconv.Atoi(parts[1])
			if err != nil {
				return "", fmt.Errorf("invalid change number in %s/%s: %w", directory, entry.Name(), err)
			}
			if number > maxNumber {
				maxNumber = number
			}
		}
	}
	path := fmt.Sprintf("changes/%04d-%s.md", maxNumber+1, slug)
	body := "\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n\n"
	content, err := contracts.RenderDocument(contracts.Document{Frontmatter: map[string]any{"module": module}, Body: body})
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(filepath.Join(root, "changes"), 0755); err != nil {
		return "", fmt.Errorf("create changes directory: %w", err)
	}
	file, err := os.OpenFile(filepath.Join(root, filepath.FromSlash(path)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return "", fmt.Errorf("create %s: %w", path, err)
	}
	defer file.Close()
	if _, err := file.WriteString(content); err != nil {
		return "", fmt.Errorf("write %s: %w", path, err)
	}
	return path, nil
}
