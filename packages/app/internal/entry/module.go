package entry

import (
	"fmt"
	"maps"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/lurkshark/fireweed/packages/project"
	"gopkg.in/yaml.v3"
)

var (
	moduleName   = regexp.MustCompile(`^[a-z][a-z0-9]*(?:-[a-z0-9]+)*$`)
	modulePrefix = regexp.MustCompile(`^[A-Z][A-Z0-9]*$`)
)

func (c command) newModule(root string) response {
	name, prefix := c.arg(0), c.options["--prefix"]
	refused := func(err error) response {
		return failed(2, err, "P01", "Use an unused name, prefix and destination.")
	}
	configFile := filepath.Join(root, "fireweed.yaml")
	config, err := os.ReadFile(configFile)
	if err != nil {
		return configFailure(err)
	}
	var fields map[string]any
	if err := yaml.Unmarshal(config, &fields); err != nil {
		return configFailure(fmt.Errorf("fireweed.yaml: %w", err))
	}
	if err := project.CheckCompatibility(fields); err != nil {
		return configFailure(err)
	}
	dir := moduleDestination(fields, name)
	if dir == "" {
		return configFailure(fmt.Errorf("fireweed.yaml must configure a module glob ending in a single * under a literal parent"))
	}
	full := filepath.Join(root, filepath.FromSlash(dir))
	if present, err := exists(full); err != nil || present {
		if err == nil {
			err = fmt.Errorf("module destination %s already exists", dir)
		}
		return refused(err)
	}
	snapshot, err := loadGraph(root)
	if err != nil {
		return configFailure(err)
	}
	for _, module := range snapshot.Modules {
		if module.Name == name || module.IDPrefix == prefix {
			return refused(fmt.Errorf("module name %q or ID prefix %q conflicts with module %q (%s)", name, prefix, module.Name, module.IDPrefix))
		}
	}

	templateError := func(err error) response {
		return failed(2, err, "C03", "Correct the project templates in "+templateDir+" and retry.")
	}
	templates, err := readTree(filepath.Join(root, filepath.FromSlash(templateDir)))
	if err != nil {
		return templateError(err)
	}
	files, err := renderModule(root, dir, name, prefix, templates)
	if err != nil {
		return templateError(err)
	}
	module, err := validateModuleTemplate(files["SPEC.md"], snapshot, dir, name, prefix)
	if err != nil {
		return templateError(err)
	}

	paths := make(map[string]string, len(files))
	for file, body := range files {
		paths[path.Join(dir, file)] = body
	}
	dirs := scaffoldDirs(dir, module.Zones)
	if err := preflightFiles(root, paths, dirs); err != nil {
		return refused(err)
	}
	if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
		return refused(err)
	}
	if err := os.Mkdir(full, 0755); err != nil {
		return refused(err)
	}
	if err := writeFiles(root, paths, dirs); err != nil {
		_ = os.RemoveAll(full)
		return refused(err)
	}
	return message(dir)
}

func moduleDestination(config map[string]any, name string) string {
	patterns, _ := stringList(config["modules"])
	for _, pattern := range patterns {
		cleaned := path.Clean(pattern)
		if path.IsAbs(cleaned) || strings.HasPrefix(cleaned, "../") {
			continue
		}
		parent, tail := path.Split(cleaned)
		if tail == "*" && !strings.ContainsAny(parent, "*?[]") {
			return path.Join(parent, name)
		}
	}
	return ""
}

func renderModule(root, dir, name, prefix string, templates map[string]string) (map[string]string, error) {
	goMod, err := readOptional(filepath.Join(root, "go.mod"))
	if err != nil {
		return nil, err
	}
	var modulePath string
	for _, line := range strings.Split(goMod, "\n") {
		parts := strings.Fields(line)
		if len(parts) >= 2 && parts[0] == "module" {
			modulePath = strings.Trim(parts[1], "\"")
			break
		}
	}
	replace := strings.NewReplacer("{{name}}", name, "{{prefix}}", prefix, "{{package}}", strings.ReplaceAll(name, "-", "_"),
		"{{module_dir}}", dir, "{{go_module}}", modulePath)
	rendered := make(map[string]string, len(templates))
	for _, file := range slices.Sorted(maps.Keys(templates)) {
		name := strings.TrimSuffix(file, ".tmpl")
		if name == "" || path.Base(name) == "." || path.Base(name) == ".." || strings.HasSuffix(name, "/") {
			return nil, fmt.Errorf("invalid rendered template path %q", name)
		}
		if _, found := rendered[name]; found {
			return nil, fmt.Errorf("templates collide at %s", name)
		}
		rendered[name] = replace.Replace(templates[file])
	}
	return rendered, nil
}

// Load the rendered SPEC with the same project parser and defaults before
// reserving the destination directory.
func validateModuleTemplate(spec string, snapshot project.Snapshot, dir, name, prefix string) (project.Module, error) {
	var module project.Module
	temp, err := os.MkdirTemp("", "fireweed-module-")
	if err != nil {
		return module, err
	}
	defer os.RemoveAll(temp)
	config := map[string]any{"modules": []string{"."}}
	if defaults, ok := snapshot.Config["module_defaults"]; ok {
		config["module_defaults"] = defaults
	}
	data, err := yaml.Marshal(config)
	if err != nil {
		return module, err
	}
	if err := os.WriteFile(filepath.Join(temp, "fireweed.yaml"), data, 0600); err != nil {
		return module, err
	}
	if err := os.WriteFile(filepath.Join(temp, "SPEC.md"), []byte(spec), 0600); err != nil {
		return module, err
	}
	candidate, err := project.Load(temp)
	if err != nil {
		return module, fmt.Errorf("invalid rendered SPEC.md: %v", err)
	}
	module = candidate.Modules[0]
	module.Dir = dir
	if module.Name != name || module.IDPrefix != prefix {
		return module, fmt.Errorf("rendered SPEC.md must declare name %q and id_prefix %q", name, prefix)
	}
	if err := project.ValidateGraph(append(snapshot.Modules[:len(snapshot.Modules):len(snapshot.Modules)], module)); err != nil {
		return module, fmt.Errorf("invalid rendered SPEC.md: %v", err)
	}
	return module, nil
}
