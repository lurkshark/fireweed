package entry

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"maps"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/contracts"
	"gopkg.in/yaml.v3"
)

const (
	templateDir  = ".fireweed/templates/module"
	startMarker  = "<!-- fireweed:start -->"
	endMarker    = "<!-- fireweed:end -->"
	managedBlock = startMarker + "\nThis project uses fireweed. Specs and change files are the source of truth.\n" +
		"Run `fireweed playbook author-change` before authoring changes and\n" +
		"`fireweed context <change>` before each step.\n" + endMarker
)

func readOptional(file string) (string, error) {
	body, err := os.ReadFile(file)
	if os.IsNotExist(err) {
		return "", nil
	}
	return string(body), err
}

func withManagedBlock(body string) (string, error) {
	start, end := strings.Index(body, startMarker), strings.Index(body, endMarker)
	if start < 0 && end < 0 {
		if body != "" && !strings.HasSuffix(body, "\n") {
			body += "\n"
		}
		return body + managedBlock + "\n", nil
	}
	if start < 0 || end < start || strings.Count(body, startMarker) != 1 || strings.Count(body, endMarker) != 1 {
		return "", errors.New("unpaired Fireweed managed block markers in AGENTS.md")
	}
	return body[:start] + managedBlock + body[end+len(endMarker):], nil
}

func refreshAgents(root string) error {
	file := filepath.Join(root, "AGENTS.md")
	body, err := readOptional(file)
	if err != nil {
		return err
	}
	updated, err := withManagedBlock(body)
	if err != nil || body == updated {
		return err
	}
	return os.WriteFile(file, []byte(updated), 0644)
}

func agentCommand(value string) (any, error) {
	switch value {
	case "codex":
		return []string{"codex", "exec", "{prompt}"}, nil
	case "claude":
		return []string{"claude", "-p", "{prompt}", "--permission-mode", "acceptEdits"}, nil
	default:
		if strings.Contains(value, "{prompt}") || strings.Contains(value, "{prompt_file}") {
			return value, nil
		}
		return nil, fmt.Errorf("invalid agent command %q", value)
	}
}

func readTree(dir string) (map[string]string, error) {
	files := map[string]string{}
	err := filepath.WalkDir(dir, func(file string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		if !entry.Type().IsRegular() {
			return fmt.Errorf("%s must be a regular file", file)
		}
		name, err := filepath.Rel(dir, file)
		if err != nil {
			return err
		}
		body, err := os.ReadFile(file)
		if err == nil {
			files[filepath.ToSlash(name)] = string(body)
		}
		return err
	})
	return files, err
}

func readRecipe(root, selection string) (assets.Recipe, error) {
	if selection == "none" {
		return assets.Recipe{Instructions: "Configure an eval hook that accepts the module and report path and writes JUnit XML, a boundaries hook that reports violations, and a setup hook to prepare scratch worktrees. See fireweed help doctor to check the hooks.\n"}, nil
	}
	if !strings.Contains(selection, "/") && selection != "." && selection != ".." {
		recipe, err := assets.BuiltinRecipe(selection)
		if err != nil {
			err = coded("U02", selection, err)
		}
		return recipe, err
	}
	dir := selection
	if !filepath.IsAbs(dir) {
		dir = filepath.Join(root, dir)
	}
	metaFile := filepath.Join(dir, "recipe.yaml")
	metadata, err := os.ReadFile(metaFile)
	if err != nil {
		return assets.Recipe{}, coded("C03", metaFile, err)
	}
	instructionsFile := filepath.Join(dir, "recipe.md")
	instructions, err := os.ReadFile(instructionsFile)
	if err != nil {
		return assets.Recipe{}, coded("C03", instructionsFile, err)
	}
	var raw any
	dec := yaml.NewDecoder(strings.NewReader(string(metadata)))
	if err := dec.Decode(&raw); err != nil {
		return assets.Recipe{}, coded("C03", metaFile, err)
	}
	var extra any
	if err := dec.Decode(&extra); !errors.Is(err, io.EOF) {
		if err == nil {
			err = errors.New("recipe.yaml must contain one document")
		}
		return assets.Recipe{}, coded("C03", metaFile, err)
	}
	recipe, err := recipeMetadata(raw)
	if err != nil {
		return recipe, coded("C12", metaFile, err)
	}
	recipe.Name, recipe.Instructions = "local", string(instructions)
	filesDir := filepath.Join(dir, "files")
	recipe.Files, err = readTree(filesDir)
	if err != nil && !os.IsNotExist(err) {
		return recipe, coded("C03", filesDir, err)
	}
	const localInstructions = ".fireweed/recipes/local.md"
	if _, exists := recipe.Files[localInstructions]; exists {
		return recipe, coded("P01", localInstructions, errors.New("local recipe payload conflicts with copied instructions"))
	}
	recipe.Files[localInstructions] = recipe.Instructions
	return recipe, nil
}

func zoneFields(z *contracts.Zones) map[string]*[]string {
	return map[string]*[]string{
		"protected": &z.Protected, "interface": &z.Interface, "evals": &z.Evals,
		"append_only": &z.AppendOnly, "preserve": &z.Preserve, "disposable": &z.Disposable,
	}
}

func stringList(raw any) ([]string, bool) {
	entries, ok := raw.([]any)
	if !ok {
		return nil, false
	}
	values := make([]string, 0, len(entries))
	for _, entry := range entries {
		value, ok := entry.(string)
		if !ok {
			return nil, false
		}
		values = append(values, value)
	}
	return values, true
}

func recipeMetadata(raw any) (assets.Recipe, error) {
	var recipe assets.Recipe
	fields, ok := raw.(map[string]any)
	if !ok {
		return recipe, errors.New("recipe metadata must be a mapping")
	}
	if value, present := fields["module_defaults"]; present {
		defaults, ok := value.(map[string]any)
		if !ok {
			return recipe, errors.New("module_defaults must be a mapping")
		}
		zones := zoneFields(&recipe.ModuleDefaults)
		for _, key := range slices.Sorted(maps.Keys(zones)) {
			if value, present := defaults[key]; present {
				values, ok := stringList(value)
				if !ok {
					return recipe, fmt.Errorf("module_defaults.%s must be a list of strings", key)
				}
				*zones[key] = values
			}
		}
	}
	if value, present := fields["hooks"]; present {
		recipe.Hooks, ok = value.(map[string]any)
		if !ok {
			return recipe, errors.New("hooks must be a mapping")
		}
		for _, hook := range slices.Sorted(maps.Keys(recipe.Hooks)) {
			value := recipe.Hooks[hook]
			_, isShell := value.(string)
			_, isArgv := stringList(value)
			if !isShell && !isArgv {
				return recipe, fmt.Errorf("hooks.%s must be a command string or list of strings", hook)
			}
		}
	}
	for _, field := range []struct {
		key    string
		target *[]string
	}{{"protected", &recipe.Protected}, {"worker_writable", &recipe.WorkerWritable}} {
		if value, present := fields[field.key]; present {
			values, ok := stringList(value)
			if !ok {
				return recipe, fmt.Errorf("%s must be a list of strings", field.key)
			}
			*field.target = values
		}
	}
	return recipe, nil
}

func initConfig(recipe assets.Recipe, agent, auditor any) ([]byte, error) {
	parts := strings.Split(assets.Version(), ".")
	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	audit := map[string]any{"enabled": true}
	if auditor != nil {
		audit["command"] = auditor
	}
	config := map[string]any{
		"fireweed": fmt.Sprintf(">=%s.%d, <%s.%d", parts[0], minor, parts[0], minor+1),
		"version":  1, "base": "main", "remote": "origin", "modules": []string{"packages/*"},
		"agent": map[string]any{"command": agent}, "audit": audit,
	}
	defaults := map[string][]string{}
	for key, list := range zoneFields(&recipe.ModuleDefaults) {
		if *list != nil {
			defaults[key] = *list
		}
	}
	if len(defaults) > 0 {
		config["module_defaults"] = defaults
	}
	if recipe.Hooks != nil {
		config["hooks"] = recipe.Hooks
	}
	if recipe.Protected != nil {
		config["protected"] = recipe.Protected
	}
	if recipe.WorkerWritable != nil {
		config["worker_writable"] = recipe.WorkerWritable
	}
	return yaml.Marshal(config)
}

func (c command) initProject(root string) response {
	agent, err := agentCommand(c.options["--agent"])
	if err != nil {
		return failed(2, err, "U01", initUsage)
	}
	var auditor any
	if c.has("--audit-agent") {
		auditor, err = agentCommand(c.options["--audit-agent"])
		if err != nil {
			return failed(2, err, "U01", initUsage)
		}
	}
	refused := func(err error) response {
		return failed(2, err, "P01", "Use a directory without conflicting setup files and retry.")
	}
	present, err := exists(filepath.Join(root, "fireweed.yaml"))
	if err != nil {
		return refused(err)
	}
	if present {
		return refused(errors.New("fireweed.yaml already exists"))
	}
	recipe, err := readRecipe(root, c.options["--recipe"])
	if err != nil {
		return failed(2, err, "C03", initUsage)
	}
	config, err := initConfig(recipe, agent, auditor)
	if err != nil {
		return refused(err)
	}

	files := map[string]string{}
	add := func(name, body string) error {
		if _, duplicate := files[name]; duplicate {
			return fmt.Errorf("setup files conflict at %s", name)
		}
		files[name] = body
		return nil
	}
	for _, name := range slices.Sorted(maps.Keys(recipe.Files)) {
		if err := add(name, recipe.Files[name]); err != nil {
			return refused(err)
		}
	}
	if err := add("fireweed.yaml", string(config)); err != nil {
		return refused(err)
	}
	system, err := exists(filepath.Join(root, "SYSTEM.md"))
	if err != nil {
		return refused(err)
	}
	if !system {
		if err := add("SYSTEM.md", "# Project\n\nDescribe the project, its architecture, and worker conventions here.\n"); err != nil {
			return refused(err)
		}
	}

	var moduleDirs []string
	if !c.has("--no-app") {
		templates := map[string]string{}
		for name, body := range recipe.Files {
			if rel, ok := strings.CutPrefix(name, templateDir+"/"); ok {
				templates[rel] = body
			}
		}
		if len(templates) == 0 {
			templates["SPEC.md"] = "---\nname: {{name}}\nid_prefix: {{prefix}}\ndepends_on: []\n---\n\n# {{name}}\n\n## Specs\n\n## Decisions\n"
		}
		rendered, err := renderModule(root, "packages/app", "app", "APP", templates)
		if err != nil {
			return failed(2, err, "C03", "Correct the recipe module templates and retry.")
		}
		for _, name := range slices.Sorted(maps.Keys(rendered)) {
			if err := add(path.Join("packages/app", name), rendered[name]); err != nil {
				return refused(err)
			}
		}
		moduleDirs = scaffoldDirs("packages/app", recipe.ModuleDefaults)
	}

	updates := map[string]string{}
	for _, name := range []string{"AGENTS.md", ".gitignore"} {
		if _, supplied := files[name]; supplied {
			return refused(fmt.Errorf("recipe payload conflicts with managed setup file %s", name))
		}
		body, err := readOptional(filepath.Join(root, name))
		if err != nil {
			return refused(err)
		}
		switch name {
		case "AGENTS.md":
			updates[name], err = withManagedBlock(body)
			if err != nil {
				return refused(err)
			}
		case ".gitignore":
			if !slices.Contains(strings.Split(body, "\n"), ".fireweed/runs/") {
				if body != "" && !strings.HasSuffix(body, "\n") {
					body += "\n"
				}
				updates[name] = body + ".fireweed/runs/\n"
			}
		}
	}
	if err := preflightFiles(root, files, moduleDirs); err != nil {
		return refused(err)
	}
	configBody := files["fireweed.yaml"]
	delete(files, "fireweed.yaml")
	if err := writeFiles(root, files, moduleDirs); err != nil {
		return refused(err)
	}
	for _, name := range slices.Sorted(maps.Keys(updates)) {
		if err := os.WriteFile(filepath.Join(root, name), []byte(updates[name]), 0644); err != nil {
			return refused(err)
		}
	}
	if err := writeExclusive(filepath.Join(root, "fireweed.yaml"), configBody); err != nil {
		return refused(err)
	}
	return verbatim(recipe.Instructions)
}

func scaffoldDirs(dir string, zones contracts.Zones) []string {
	var directories []string
	for _, glob := range slices.Concat(zones.Evals, zones.Disposable) {
		parent, tail := path.Split(glob)
		if tail == "**" && parent != "" && !strings.Contains(parent, "*") && fs.ValidPath(strings.TrimSuffix(parent, "/")) {
			directories = append(directories, path.Join(dir, parent))
		}
	}
	return directories
}

func preflightFiles(root string, files map[string]string, dirs []string) error {
	for _, file := range slices.Sorted(maps.Keys(files)) {
		full := filepath.Join(root, filepath.FromSlash(file))
		if present, err := exists(full); err != nil || present {
			if err != nil {
				return err
			}
			return fmt.Errorf("%s already exists", file)
		}
		dirs = append(dirs, path.Dir(file))
	}
	for _, dir := range dirs {
		for parent := dir; parent != "."; parent = path.Dir(parent) {
			if _, found := files[parent]; found {
				return fmt.Errorf("%s is also required as a directory", parent)
			}
			full := filepath.Join(root, filepath.FromSlash(parent))
			info, err := os.Lstat(full)
			if os.IsNotExist(err) {
				continue
			}
			if err != nil {
				return err
			}
			if !info.IsDir() {
				return fmt.Errorf("%s must be a directory", parent)
			}
		}
	}
	return nil
}

func writeExclusive(file, body string) error {
	w, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return err
	}
	_, err = w.WriteString(body)
	err = errors.Join(err, w.Close())
	if err != nil {
		_ = os.Remove(file)
	}
	return err
}

func writeFiles(root string, files map[string]string, dirs []string) error {
	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(root, filepath.FromSlash(dir)), 0755); err != nil {
			return err
		}
	}
	var created []string
	for _, name := range slices.Sorted(maps.Keys(files)) {
		file := filepath.Join(root, filepath.FromSlash(name))
		err := os.MkdirAll(filepath.Dir(file), 0755)
		if err == nil {
			err = writeExclusive(file, files[name])
		}
		if err != nil {
			for _, file := range created {
				_ = os.Remove(file)
			}
			return err
		}
		created = append(created, file)
	}
	return nil
}
