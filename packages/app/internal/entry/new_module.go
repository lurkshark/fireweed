package entry

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

const newModuleUsage = "Use new-module <name> --prefix <P> [--json]."

var moduleNamePattern = regexp.MustCompile(`^[a-z][a-z0-9]*(?:-[a-z0-9]+)*$`)
var modulePrefixPattern = regexp.MustCompile(`^[A-Z][A-Z0-9]*$`)

func isNewModuleRequest(args []string) bool {
	for _, arg := range args {
		if arg != "--json" {
			return arg == "new-module"
		}
	}
	return false
}

func runNewModule(args []string, out, errOut io.Writer) int {
	jsonMode := contains(args, "--json")
	fail := func(code, message string) int {
		return cliFailure(2, jsonMode, out, errOut, code, message)
	}
	name, prefix := "", ""
	seenCommand, seenJSON, seenPrefix := false, false, false
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "new-module":
			if seenCommand {
				return fail("U01", "new-module takes one name. "+newModuleUsage)
			}
			seenCommand = true
		case "--json":
			if seenJSON {
				return fail("U01", "--json may be specified once. "+newModuleUsage)
			}
			seenJSON = true
		case "--prefix":
			if seenPrefix || i+1 >= len(args) || strings.HasPrefix(args[i+1], "-") {
				return fail("U01", "--prefix requires one ID prefix. "+newModuleUsage)
			}
			seenPrefix = true
			i++
			prefix = args[i]
		default:
			if strings.HasPrefix(args[i], "-") || name != "" || !seenCommand {
				return fail("U01", fmt.Sprintf("Unexpected new-module argument %q. %s", args[i], newModuleUsage))
			}
			name = args[i]
		}
	}
	if !moduleNamePattern.MatchString(name) || !modulePrefixPattern.MatchString(prefix) {
		return fail("U01", "Provide a lowercase module name and uppercase ID prefix. "+newModuleUsage)
	}

	snapshot, err := loadValidatedProject()
	if err != nil {
		return fail(projectFailureCode(err, "C01"), fmt.Sprintf("%v. Use a valid project configuration and module declarations, then retry.", err))
	}
	for _, module := range snapshot.Modules {
		if module.Name == name || module.IDPrefix == prefix {
			return fail("P01", fmt.Sprintf("Module name %q or ID prefix %q already exists. Choose an unused name and prefix.", name, prefix))
		}
	}
	dir, ok := moduleDestination(snapshot.Config, name)
	if !ok {
		return fail("C01", "No configured module glob has a literal parent and one terminal wildcard. Use a modules pattern such as packages/*, then retry.")
	}
	target := filepath.Join(snapshot.Root, filepath.FromSlash(dir))
	if _, err := os.Lstat(target); err == nil {
		return fail("P01", fmt.Sprintf("Module destination %s already exists. Choose an unused module name or move the existing directory.", dir))
	} else if !os.IsNotExist(err) {
		return fail("P01", fmt.Sprintf("Cannot inspect module destination %s: %v. Use a writable directory and retry.", dir, err))
	}
	files, err := renderModuleTemplates(snapshot.Root, dir, name, prefix)
	if err != nil {
		return fail("C03", fmt.Sprintf("Invalid module templates: %v. Provide valid files in .fireweed/templates/module and retry.", err))
	}
	if err := validateModuleTemplate(snapshot, dir, name, prefix, files); err != nil {
		return fail("C03", fmt.Sprintf("Invalid module SPEC template: %v. Provide a valid .fireweed/templates/module/SPEC.md and retry.", err))
	}
	if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
		return fail("P01", fmt.Sprintf("Cannot create module parent: %v. Use a writable destination directory and retry.", err))
	}
	if err := os.Mkdir(target, 0755); err != nil {
		return fail("P01", fmt.Sprintf("Cannot create module destination: %v. Choose an unused writable directory and retry.", err))
	}
	if err := writeInitFiles(snapshot.Root, files); err != nil {
		return fail("P01", fmt.Sprintf("Cannot write module files: %v. Use a writable destination directory and retry.", err))
	}
	return cliResponse(0, jsonMode, out, errOut, dir, dir+"\n", nil)
}

func moduleDestination(config map[string]any, name string) (string, bool) {
	patterns, ok := config["modules"].([]any)
	if !ok {
		return "", false
	}
	for _, value := range patterns {
		pattern, ok := value.(string)
		if !ok || !strings.HasSuffix(pattern, "/*") || strings.Count(pattern, "*") != 1 {
			continue
		}
		parent := strings.TrimSuffix(pattern, "/*")
		if parent == "" || parent == "." || path.IsAbs(parent) || path.Clean(parent) != parent || strings.ContainsAny(parent, "*?[]\\") || parent == ".." || strings.HasPrefix(parent, "../") {
			continue
		}
		dir := parent + "/" + name
		if contracts.MatchGlob(pattern, dir) {
			return dir, true
		}
	}
	return "", false
}

func renderModuleTemplates(root, dir, name, prefix string) (map[string][]byte, error) {
	templateRoot := filepath.Join(root, ".fireweed", "templates", "module")
	info, err := os.Stat(templateRoot)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("template path is not a directory")
	}
	goModule := ""
	if data, err := os.ReadFile(filepath.Join(root, "go.mod")); err == nil {
		for _, line := range strings.Split(string(data), "\n") {
			fields := strings.Fields(line)
			if len(fields) == 2 && fields[0] == "module" {
				goModule = fields[1]
				break
			}
		}
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("read go.mod: %w", err)
	}
	replacer := strings.NewReplacer(
		"{{name}}", name, "{{prefix}}", prefix,
		"{{package}}", strings.ReplaceAll(name, "-", "_"),
		"{{module_dir}}", dir, "{{go_module}}", goModule,
	)
	files := map[string][]byte{}
	err = filepath.WalkDir(templateRoot, func(source string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if source == templateRoot || entry.IsDir() {
			return nil
		}
		if !entry.Type().IsRegular() {
			return fmt.Errorf("%s is not a regular template file", source)
		}
		rel, err := filepath.Rel(templateRoot, source)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if strings.HasSuffix(rel, ".tmpl") {
			rel = strings.TrimSuffix(rel, ".tmpl")
		}
		if rel == "" || rel == "." {
			return fmt.Errorf("empty template output path")
		}
		output := path.Join(dir, rel)
		if _, exists := files[output]; exists {
			return fmt.Errorf("duplicate template output %s", output)
		}
		content, err := os.ReadFile(source)
		if err != nil {
			return err
		}
		files[output] = []byte(replacer.Replace(string(content)))
		return nil
	})
	if err != nil {
		return nil, err
	}
	if _, ok := files[path.Join(dir, "SPEC.md")]; !ok {
		return nil, fmt.Errorf("SPEC.md is missing")
	}
	return files, nil
}

func validateModuleTemplate(snapshot project.Snapshot, dir, name, prefix string, files map[string][]byte) error {
	doc, err := contracts.ParseDocument(string(files[path.Join(dir, "SPEC.md")]))
	if err != nil {
		return err
	}
	if doc.Frontmatter["name"] != name || doc.Frontmatter["id_prefix"] != prefix {
		return fmt.Errorf("SPEC.md must declare name %q and id_prefix %q", name, prefix)
	}
	module := project.Module{Name: name, IDPrefix: prefix, Dir: dir, Document: doc}
	if value, exists := doc.Frontmatter["depends_on"]; exists {
		items, ok := value.([]any)
		if !ok {
			return fmt.Errorf("depends_on must be a list of strings")
		}
		for _, item := range items {
			dependency, ok := item.(string)
			if !ok {
				return fmt.Errorf("depends_on must be a list of strings")
			}
			module.DependsOn = append(module.DependsOn, dependency)
		}
	}
	defaults, _ := snapshot.Config["module_defaults"].(map[string]any)
	for _, field := range []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"} {
		value, exists := doc.Frontmatter[field]
		if !exists {
			value, exists = defaults[field]
		}
		if !exists {
			continue
		}
		items, ok := value.([]any)
		if !ok {
			return fmt.Errorf("%s must be a list of strings", field)
		}
		for _, item := range items {
			if _, ok := item.(string); !ok {
				return fmt.Errorf("%s must be a list of strings", field)
			}
		}
	}
	modules := append(append([]project.Module(nil), snapshot.Modules...), module)
	return project.ValidateGraph(modules)
}
