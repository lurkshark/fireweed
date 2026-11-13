package fixture

import (
	"fmt"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func stringList(key string, raw any) ([]string, error) {
	values, ok := raw.([]any)
	if !ok {
		return nil, fmt.Errorf("%s must be a list of strings", key)
	}
	items := make([]string, 0, len(values))
	for _, value := range values {
		item, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("%s must be a list of strings", key)
		}
		items = append(items, item)
	}
	return items, nil
}

func readZones(fields map[string]any, zones contracts.Zones) (contracts.Zones, error) {
	for _, field := range []struct {
		key  string
		list *[]string
	}{
		{"protected", &zones.Protected}, {"evals", &zones.Evals}, {"interface", &zones.Interface},
		{"append_only", &zones.AppendOnly}, {"preserve", &zones.Preserve}, {"disposable", &zones.Disposable},
	} {
		if raw, present := fields[field.key]; present {
			list, err := stringList(field.key, raw)
			if err != nil {
				return contracts.Zones{}, err
			}
			*field.list = list
		}
	}
	return zones, nil
}

func ReadModule(dir string, doc contracts.Document, defaults contracts.Zones) (model.Module, error) {
	module := model.Module{Dir: dir, Document: doc}
	fields := doc.Frontmatter
	for _, field := range []struct {
		key   string
		value *string
	}{
		{"name", &module.Name}, {"id_prefix", &module.IDPrefix}, {"boundaries_opt_out", nil},
	} {
		if raw, present := fields[field.key]; present {
			value, ok := raw.(string)
			if !ok {
				return model.Module{}, fmt.Errorf("%s must be a string", field.key)
			}
			if field.value != nil {
				*field.value = value
			}
		}
	}
	if raw, present := fields["hooks"]; present {
		if _, ok := raw.(map[string]any); !ok {
			return model.Module{}, fmt.Errorf("hooks must be a mapping")
		}
	}
	if raw, present := fields["depends_on"]; present {
		deps, err := stringList("depends_on", raw)
		if err != nil {
			return model.Module{}, err
		}
		module.DependsOn = deps
	}
	zones, err := readZones(fields, defaults)
	if err != nil {
		return model.Module{}, err
	}
	module.Zones = zones
	return module, nil
}

func ValidateGraph(modules []model.Module) error {
	byName := make(map[string]model.Module, len(modules))
	prefixes := make(map[string]bool, len(modules))
	for _, module := range modules {
		if module.Name == "" {
			return fmt.Errorf("empty module name in %s", module.Dir)
		}
		if _, exists := byName[module.Name]; exists {
			return fmt.Errorf("duplicate module name %q", module.Name)
		}
		if module.IDPrefix == "" {
			return fmt.Errorf("empty ID prefix for module %q", module.Name)
		}
		if prefixes[module.IDPrefix] {
			return fmt.Errorf("duplicate module ID prefix %q", module.IDPrefix)
		}
		byName[module.Name] = module
		prefixes[module.IDPrefix] = true
	}

	visited := make(map[string]bool, len(modules))
	active := make(map[string]bool)
	var visit func(model.Module) error
	visit = func(module model.Module) error {
		if active[module.Name] {
			return fmt.Errorf("dependency cycle involving module %q", module.Name)
		}
		if visited[module.Name] {
			return nil
		}
		active[module.Name] = true
		for _, name := range module.DependsOn {
			dependency, exists := byName[name]
			if !exists {
				return fmt.Errorf("module %q depends on unknown module %q", module.Name, name)
			}
			if name == module.Name {
				return fmt.Errorf("module %q has a self dependency", name)
			}
			if err := visit(dependency); err != nil {
				return err
			}
		}
		delete(active, module.Name)
		visited[module.Name] = true
		return nil
	}
	for _, module := range modules {
		if err := visit(module); err != nil {
			return err
		}
	}
	return nil
}
