package entry

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/contracts"
	"gopkg.in/yaml.v3"
)

const initUsage = "Use init --recipe go|none|./directory --agent codex|claude|<command containing {prompt}> [--audit-agent <agent>] [--no-app] [--json]."

const managedBlock = "<!-- fireweed:start -->\nThis project uses fireweed. Specs and change files are the source of truth.\nRun `fireweed playbook author-change` before authoring changes and\n`fireweed context <change>` before each step.\n<!-- fireweed:end -->"

const noneInstructions = "# Fireweed setup\n\nDefine the eval, boundaries, and setup hooks for this project. The eval hook must write a JUnit report for a module; the boundaries hook must check module imports and zones; the setup hook must prepare dependencies. Add module templates and canaries before running doctor.\n"

func isInitRequest(args []string) bool {
	for _, arg := range args {
		if arg != "--json" {
			return arg == "init"
		}
	}
	return false
}

func runInit(args []string, out, errOut io.Writer) int {
	jsonMode := contains(args, "--json")
	fail := func(code, message string) int {
		return cliFailure(2, jsonMode, out, errOut, code, message)
	}
	recipeName, agentName, auditName := "", "", ""
	noApp := false
	seen := map[string]bool{}
	commandSeen := false
	for i := 0; i < len(args); i++ {
		switch arg := args[i]; arg {
		case "init":
			if commandSeen {
				return fail("U01", "init takes no positional arguments. "+initUsage)
			}
			commandSeen = true
		case "--json", "--no-app":
			if seen[arg] {
				return fail("U01", "Duplicate init option "+arg+". "+initUsage)
			}
			seen[arg] = true
			if arg == "--no-app" {
				noApp = true
			}
		case "--recipe", "--agent", "--audit-agent":
			if seen[arg] || i+1 >= len(args) || strings.HasPrefix(args[i+1], "--") || args[i+1] == "" {
				return fail("U01", arg+" requires one value. "+initUsage)
			}
			seen[arg] = true
			i++
			switch arg {
			case "--recipe":
				recipeName = args[i]
			case "--agent":
				agentName = args[i]
			case "--audit-agent":
				auditName = args[i]
			}
		default:
			return fail("U01", fmt.Sprintf("Unexpected init argument %q. %s", arg, initUsage))
		}
	}
	if recipeName == "" || agentName == "" {
		return fail("U01", "init requires both --recipe and --agent. "+initUsage)
	}
	agent, ok := initAgent(agentName)
	if !ok {
		return fail("U01", "Invalid --agent value. "+initUsage)
	}
	var auditor any
	if auditName != "" {
		auditor, ok = initAgent(auditName)
		if !ok {
			return fail("U01", "Invalid --audit-agent value. "+initUsage)
		}
	}

	root, err := os.Getwd()
	if err != nil {
		return fail("P01", fmt.Sprintf("Cannot read current directory: %v. Choose a writable project directory.", err))
	}
	if _, err := os.Lstat(filepath.Join(root, "fireweed.yaml")); err == nil {
		return fail("P01", "fireweed.yaml already exists. Choose an uninitialized project directory.")
	} else if !os.IsNotExist(err) {
		return fail("P01", fmt.Sprintf("Cannot inspect fireweed.yaml: %v. Choose a writable project directory.", err))
	}
	recipe, code, err := selectInitRecipe(root, recipeName)
	if err != nil {
		return fail(code, err.Error())
	}
	files := make(map[string][]byte)
	for path, content := range recipe.Files {
		if path == "fireweed.yaml" || path == "SYSTEM.md" || path == "AGENTS.md" || path == ".gitignore" || path == "packages/app/SPEC.md" {
			return fail("P01", fmt.Sprintf("Recipe file %s conflicts with generated setup. Move that file in the recipe.", path))
		}
		files[path] = []byte(content)
	}
	if !noApp {
		if err := addAppModule(files); err != nil {
			return fail("C03", fmt.Sprintf("Invalid module template: %v. Correct the selected recipe.", err))
		}
	}
	config := map[string]any{
		"fireweed": "=" + assets.Version(), "version": 1,
		"base": "main", "remote": "origin", "modules": []string{"packages/*"},
		"module_defaults": zoneMap(recipe.ModuleDefaults), "hooks": recipe.Hooks,
		"protected": recipe.Protected, "worker_writable": recipe.WorkerWritable,
		"agent": map[string]any{"command": agent},
		"audit": map[string]any{"enabled": true},
	}
	if auditor != nil {
		config["audit"].(map[string]any)["command"] = auditor
	}
	data, err := yaml.Marshal(config)
	if err != nil {
		return fail("C12", fmt.Sprintf("Cannot encode recipe metadata: %v. Correct the selected recipe.", err))
	}
	files["fireweed.yaml"] = data
	if _, err := os.Lstat(filepath.Join(root, "SYSTEM.md")); os.IsNotExist(err) {
		files["SYSTEM.md"] = []byte("# Project system\n\nDescribe the project's goal, module boundaries, and agent instructions here.\n")
	} else if err != nil {
		return fail("P01", fmt.Sprintf("Cannot inspect SYSTEM.md: %v. Correct the project path.", err))
	}
	agents, err := updateManagedFile(root)
	if err != nil {
		return fail("P01", fmt.Sprintf("Cannot update AGENTS.md: %v. Correct the managed block.", err))
	}
	ignore, err := updateIgnoreFile(root)
	if err != nil {
		return fail("P01", fmt.Sprintf("Cannot update .gitignore: %v. Correct the project path.", err))
	}
	if err := validateInitDestinations(root, files); err != nil {
		return fail("P01", fmt.Sprintf("Init would overwrite a project file: %v. Remove or move the conflicting file.", err))
	}
	if !noApp && recipeName == "go" {
		for _, dir := range []string{"packages/app/evals", "packages/app/internal"} {
			info, err := os.Lstat(filepath.Join(root, filepath.FromSlash(dir)))
			if err == nil && !info.IsDir() {
				return fail("P01", fmt.Sprintf("Go module directory %s conflicts with a project file. Move the conflicting file.", dir))
			}
			if err != nil && !os.IsNotExist(err) {
				return fail("P01", fmt.Sprintf("Cannot inspect Go module directory %s: %v. Check the project path.", dir, err))
			}
		}
	}
	if err := writeInitFiles(root, files); err != nil {
		return fail("P01", fmt.Sprintf("Cannot write project setup: %v. Check directory permissions.", err))
	}
	if err := os.WriteFile(filepath.Join(root, "AGENTS.md"), agents, 0644); err != nil {
		return fail("P01", fmt.Sprintf("Cannot write AGENTS.md: %v. Check directory permissions.", err))
	}
	if err := os.WriteFile(filepath.Join(root, ".gitignore"), ignore, 0644); err != nil {
		return fail("P01", fmt.Sprintf("Cannot write .gitignore: %v. Check directory permissions.", err))
	}
	if !noApp && recipeName == "go" {
		for _, dir := range []string{"packages/app/evals", "packages/app/internal"} {
			if err := os.MkdirAll(filepath.Join(root, filepath.FromSlash(dir)), 0755); err != nil {
				return fail("P01", fmt.Sprintf("Cannot create Go module layout: %v. Check directory permissions.", err))
			}
		}
	}
	if jsonMode {
		_ = json.NewEncoder(out).Encode(envelope{Schema: 1, Data: recipe.Instructions, Diagnostics: []diagnostic{}})
	} else {
		_, _ = io.WriteString(out, recipe.Instructions)
	}
	return 0
}

func initAgent(name string) (any, bool) {
	switch name {
	case "codex":
		return []string{"codex", "exec", "{prompt}"}, true
	case "claude":
		return []string{"claude", "-p", "{prompt}", "--permission-mode", "acceptEdits"}, true
	default:
		return name, strings.Contains(name, "{prompt}")
	}
}

func zoneMap(z contracts.Zones) map[string]any {
	result := map[string]any{}
	for _, field := range []struct {
		name string
		list []string
	}{{"protected", z.Protected}, {"interface", z.Interface}, {"evals", z.Evals}, {"append_only", z.AppendOnly}, {"preserve", z.Preserve}, {"disposable", z.Disposable}} {
		if field.list != nil {
			result[field.name] = field.list
		}
	}
	return result
}

func selectInitRecipe(root, name string) (assets.Recipe, string, error) {
	switch name {
	case "go":
		r, err := assets.BuiltinRecipe("go")
		return r, "C03", err
	case "none":
		return assets.Recipe{Name: "none", Instructions: noneInstructions, Hooks: map[string]any{}, Files: map[string]string{}}, "", nil
	}
	if !strings.HasPrefix(name, "./") && !strings.HasPrefix(name, "../") && !filepath.IsAbs(name) {
		return assets.Recipe{}, "U02", fmt.Errorf("unknown recipe %q. Use go, none, or a local ./directory recipe", name)
	}
	path := filepath.Join(root, name)
	metadata, err := os.ReadFile(filepath.Join(path, "recipe.yaml"))
	if err != nil {
		return assets.Recipe{}, "C03", fmt.Errorf("cannot read recipe.yaml: %v. Provide a readable local recipe", err)
	}
	instructions, err := os.ReadFile(filepath.Join(path, "recipe.md"))
	if err != nil {
		return assets.Recipe{}, "C03", fmt.Errorf("cannot read recipe.md: %v. Provide a readable local recipe", err)
	}
	var parsed any
	if err := yaml.Unmarshal(metadata, &parsed); err != nil || parsed == nil {
		return assets.Recipe{}, "C03", fmt.Errorf("invalid recipe.yaml: %v. Correct the YAML document", err)
	}
	document, ok := parsed.(map[string]any)
	if !ok {
		return assets.Recipe{}, "C12", fmt.Errorf("recipe.yaml must be a metadata mapping. Use valid recipe metadata")
	}
	r := assets.Recipe{Name: "local", Instructions: string(instructions), Files: map[string]string{".fireweed/recipes/local.md": string(instructions)}}
	if raw, present := document["module_defaults"]; present {
		fields, ok := raw.(map[string]any)
		if !ok {
			return r, "C12", fmt.Errorf("module_defaults must be a mapping. Use valid recipe metadata")
		}
		for key, value := range fields {
			list, err := initStringList(value)
			if err != nil {
				return r, "C12", fmt.Errorf("module_defaults.%s must be a list of strings. Use valid recipe metadata", key)
			}
			switch key {
			case "protected":
				r.ModuleDefaults.Protected = list
			case "interface":
				r.ModuleDefaults.Interface = list
			case "evals":
				r.ModuleDefaults.Evals = list
			case "append_only":
				r.ModuleDefaults.AppendOnly = list
			case "preserve":
				r.ModuleDefaults.Preserve = list
			case "disposable":
				r.ModuleDefaults.Disposable = list
			default:
				return r, "C12", fmt.Errorf("unknown module_defaults field %q. Use valid recipe metadata", key)
			}
		}
	}
	if raw, present := document["hooks"]; present {
		hooks, ok := raw.(map[string]any)
		if !ok {
			return r, "C12", fmt.Errorf("hooks must be a mapping. Use valid recipe metadata")
		}
		for key, value := range hooks {
			if _, ok := value.(string); !ok {
				return r, "C12", fmt.Errorf("hooks.%s must be a string. Use valid recipe metadata", key)
			}
		}
		r.Hooks = hooks
	} else {
		r.Hooks = map[string]any{}
	}
	for _, field := range []struct {
		name   string
		target *[]string
	}{{"protected", &r.Protected}, {"worker_writable", &r.WorkerWritable}} {
		if value, present := document[field.name]; present {
			list, err := initStringList(value)
			if err != nil {
				return r, "C12", fmt.Errorf("%s must be a list of strings. Use valid recipe metadata", field.name)
			}
			*field.target = list
		}
	}
	err = filepath.WalkDir(filepath.Join(path, "files"), func(source string, entry fs.DirEntry, walkErr error) error {
		if os.IsNotExist(walkErr) && source == filepath.Join(path, "files") {
			return nil
		}
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(filepath.Join(path, "files"), source)
		if err != nil {
			return err
		}
		content, err := os.ReadFile(source)
		if err != nil {
			return err
		}
		r.Files[filepath.ToSlash(rel)] = string(content)
		return nil
	})
	if err != nil {
		return r, "C03", fmt.Errorf("cannot read recipe files: %v. Correct local recipe", err)
	}
	return r, "", nil
}

func initStringList(value any) ([]string, error) {
	items, ok := value.([]any)
	if !ok {
		return nil, fmt.Errorf("not a list")
	}
	result := make([]string, 0, len(items))
	for _, item := range items {
		str, ok := item.(string)
		if !ok {
			return nil, fmt.Errorf("not a string")
		}
		result = append(result, str)
	}
	return result, nil
}

func addAppModule(files map[string][]byte) error {
	const prefix = ".fireweed/templates/module/"
	foundSpec := false
	for path, content := range files {
		if !strings.HasPrefix(path, prefix) {
			continue
		}
		rel := strings.TrimPrefix(path, prefix)
		if strings.HasSuffix(rel, ".tmpl") {
			rel = strings.TrimSuffix(rel, ".tmpl")
		}
		target := "packages/app/" + rel
		if _, exists := files[target]; exists {
			return fmt.Errorf("duplicate template destination %s", target)
		}
		rendered := strings.NewReplacer("{{name}}", "app", "{{prefix}}", "APP", "{{package}}", "app").Replace(string(content))
		files[target] = []byte(rendered)
		if rel == "SPEC.md" {
			foundSpec = true
		}
	}
	if !foundSpec {
		files["packages/app/SPEC.md"] = []byte("---\nname: app\nid_prefix: APP\ndepends_on: []\n---\n\n# app\n\n## Specs\n\n## Decisions\n")
	}
	return nil
}

func updateManagedFile(root string) ([]byte, error) {
	data, err := os.ReadFile(filepath.Join(root, "AGENTS.md"))
	if os.IsNotExist(err) {
		return []byte(managedBlock + "\n"), nil
	}
	if err != nil {
		return nil, err
	}
	text := string(data)
	start := strings.Index(text, "<!-- fireweed:start -->")
	end := strings.Index(text, "<!-- fireweed:end -->")
	if start < 0 && end < 0 {
		if text != "" && !strings.HasSuffix(text, "\n") {
			text += "\n"
		}
		if text != "" {
			text += "\n"
		}
		return []byte(text + managedBlock + "\n"), nil
	}
	if start < 0 || end < start {
		return nil, fmt.Errorf("unpaired fireweed markers")
	}
	end += len("<!-- fireweed:end -->")
	return []byte(text[:start] + managedBlock + text[end:]), nil
}

func updateIgnoreFile(root string) ([]byte, error) {
	data, err := os.ReadFile(filepath.Join(root, ".gitignore"))
	if os.IsNotExist(err) {
		return []byte(".fireweed/runs/\n"), nil
	}
	if err != nil {
		return nil, err
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.TrimSpace(line) == ".fireweed/runs/" {
			return data, nil
		}
	}
	if len(data) > 0 && data[len(data)-1] != '\n' {
		data = append(data, '\n')
	}
	return append(data, []byte(".fireweed/runs/\n")...), nil
}

func validateInitDestinations(root string, files map[string][]byte) error {
	seen := map[string]string{}
	for path := range files {
		clean := filepath.Clean(filepath.FromSlash(path))
		if filepath.IsAbs(clean) || clean == "." || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
			return fmt.Errorf("invalid destination %q", path)
		}
		if earlier, exists := seen[clean]; exists {
			return fmt.Errorf("destinations %q and %q overlap", earlier, path)
		}
		seen[clean] = path
		target := filepath.Join(root, clean)
		if _, err := os.Lstat(target); err == nil {
			return fmt.Errorf("%s exists", path)
		} else if !os.IsNotExist(err) {
			return err
		}
		for parent := filepath.Dir(target); parent != root && parent != "."; parent = filepath.Dir(parent) {
			info, err := os.Lstat(parent)
			if os.IsNotExist(err) {
				continue
			}
			if err != nil {
				return err
			}
			if !info.IsDir() {
				return fmt.Errorf("%s has a non-directory parent", path)
			}
		}
	}
	for path := range seen {
		for parent := filepath.Dir(path); parent != "."; parent = filepath.Dir(parent) {
			if earlier, exists := seen[parent]; exists {
				return fmt.Errorf("destinations %q and %q overlap", earlier, seen[path])
			}
		}
	}
	return nil
}

func writeInitFiles(root string, files map[string][]byte) error {
	paths := make([]string, 0, len(files))
	for path := range files {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		target := filepath.Join(root, filepath.FromSlash(path))
		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			return err
		}
		if err := os.WriteFile(target, files[path], 0644); err != nil {
			return err
		}
	}
	return nil
}
