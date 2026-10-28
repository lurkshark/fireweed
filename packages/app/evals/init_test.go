package evals_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
	"gopkg.in/yaml.v3"
)

type initConfig struct {
	Fireweed       string   `yaml:"fireweed"`
	Version        int      `yaml:"version"`
	Base           string   `yaml:"base"`
	Remote         string   `yaml:"remote"`
	Modules        []string `yaml:"modules"`
	ModuleDefaults struct {
		Protected  []string `yaml:"protected"`
		Interface  []string `yaml:"interface"`
		Evals      []string `yaml:"evals"`
		AppendOnly []string `yaml:"append_only"`
		Preserve   []string `yaml:"preserve"`
		Disposable []string `yaml:"disposable"`
	} `yaml:"module_defaults"`
	Hooks          map[string]any `yaml:"hooks"`
	Protected      []string       `yaml:"protected"`
	WorkerWritable []string       `yaml:"worker_writable"`
	Agent          struct {
		Command any `yaml:"command"`
	} `yaml:"agent"`
	Audit struct {
		Enabled bool `yaml:"enabled"`
		Command any  `yaml:"command"`
	} `yaml:"audit"`
}

func readInitFile(t *testing.T, root, name string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(name)))
	if err != nil {
		t.Errorf("read %s: %v", name, err)
		return ""
	}
	return string(data)
}

func readInitConfig(t *testing.T, root string) initConfig {
	t.Helper()
	var config initConfig
	if err := yaml.Unmarshal([]byte(readInitFile(t, root, "fireweed.yaml")), &config); err != nil {
		t.Errorf("parse fireweed.yaml: %v", err)
	}
	return config
}

func initCommand(value any) []string {
	items, ok := value.([]any)
	if !ok {
		return nil
	}
	result := make([]string, 0, len(items))
	for _, item := range items {
		value, ok := item.(string)
		if !ok {
			return nil
		}
		result = append(result, value)
	}
	return result
}

func assertInitOutput(t *testing.T, status int, stdout, stderr, instructions string) {
	t.Helper()
	if status != 0 || stderr != "" {
		t.Errorf("init status=%d stderr=%q; want 0 and empty", status, stderr)
	}
	if strings.TrimSuffix(stdout, "\n") != strings.TrimSuffix(instructions, "\n") {
		t.Errorf("init stdout=%q; want instructions %q", stdout, instructions)
	}
}

func assertNoInitConfig(t *testing.T, root string) {
	t.Helper()
	if _, err := os.Stat(filepath.Join(root, "fireweed.yaml")); !os.IsNotExist(err) {
		t.Errorf("rejected init created fireweed.yaml or returned unexpected stat error: %v", err)
	}
}

func Test_APP_041_APP_042_APP_043_APP_044_APP_045_APP_046_APP_050_APP_053_APP_054_APP_059_GoInit(t *testing.T) {
	bin := buildAppBinary(t)
	root := t.TempDir()
	appPut(t, root, "go.mod", "module example.test/initial\n\ngo 1.23\n")
	appPut(t, root, "SYSTEM.md", "# Existing project instructions\n")
	appPut(t, root, "AGENTS.md", "Before the block.\n\n<!-- fireweed:start -->\nold directions\n<!-- fireweed:end -->\n\nAfter the block.\n")
	appPut(t, root, ".gitignore", "custom-entry\n")
	recipe, err := assets.BuiltinRecipe("go")
	if err != nil {
		t.Fatal(err)
	}
	status, stdout, stderr := runAppBinary(t, bin, root, "init", "--recipe", "go", "--agent", "codex")
	assertInitOutput(t, status, stdout, stderr, recipe.Instructions) // APP-046
	config := readInitConfig(t, root)
	if config.Version != 1 || !strings.Contains(config.Fireweed, strings.TrimSuffix(assets.Version(), ".0")) || config.Base != "main" || config.Remote != "origin" || !reflect.DeepEqual(config.Modules, []string{"packages/*"}) {
		t.Errorf("APP-041 configuration header = %+v", config)
	}
	if !reflect.DeepEqual(config.ModuleDefaults.Interface, recipe.ModuleDefaults.Interface) || !reflect.DeepEqual(config.ModuleDefaults.Evals, recipe.ModuleDefaults.Evals) || !reflect.DeepEqual(config.ModuleDefaults.Disposable, recipe.ModuleDefaults.Disposable) || !reflect.DeepEqual(config.Hooks, recipe.Hooks) || !reflect.DeepEqual(config.Protected, recipe.Protected) || !reflect.DeepEqual(config.WorkerWritable, recipe.WorkerWritable) {
		t.Errorf("APP-042 recipe metadata differs: %+v", config)
	}
	if got := readInitFile(t, root, "SYSTEM.md"); got != "# Existing project instructions\n" {
		t.Errorf("APP-043 SYSTEM.md changed: %q", got)
	}
	agents := readInitFile(t, root, "AGENTS.md")
	for _, required := range []string{"Before the block.", "After the block.", "<!-- fireweed:start -->", "<!-- fireweed:end -->", "playbook author-change", "context"} {
		if !strings.Contains(agents, required) {
			t.Errorf("APP-044 AGENTS.md missing %q: %q", required, agents)
		}
	}
	if strings.Contains(agents, "old directions") || strings.Count(agents, "<!-- fireweed:start -->") != 1 {
		t.Errorf("APP-044 managed block was not replaced once: %q", agents)
	}
	if !strings.HasPrefix(agents, "Before the block.\n\n<!-- fireweed:start -->") || !strings.HasSuffix(agents, "<!-- fireweed:end -->\n\nAfter the block.\n") {
		t.Errorf("APP-044 text outside managed block changed: %q", agents)
	}
	for name, want := range recipe.Files {
		if got := readInitFile(t, root, name); got != want {
			t.Errorf("APP-045 %s differs from embedded recipe", name)
		}
	}
	spec := readInitFile(t, root, "packages/app/SPEC.md")
	if !strings.Contains(spec, "name: app") || !strings.Contains(spec, "id_prefix: APP") || !strings.Contains(spec, "depends_on: []") {
		t.Errorf("APP-050 rendered app module metadata: %q", spec)
	}
	if !strings.Contains(readInitFile(t, root, "packages/app/api.go"), "package app") {
		t.Error("APP-050 Go interface template was not rendered")
	}
	for _, directory := range []string{"packages/app/evals", "packages/app/internal"} {
		if info, err := os.Stat(filepath.Join(root, filepath.FromSlash(directory))); err != nil || !info.IsDir() {
			t.Errorf("APP-050 Go layout lacks directory %s: %v", directory, err)
		}
	}
	if _, err := os.Stat(filepath.Join(root, "packages/app/api.go.tmpl")); !os.IsNotExist(err) {
		t.Errorf("APP-050 left template filename: %v", err)
	}
	if got := initCommand(config.Agent.Command); !reflect.DeepEqual(got, []string{"codex", "exec", "{prompt}"}) {
		t.Errorf("APP-053 codex command = %q", got)
	}
	if !config.Audit.Enabled || config.Audit.Command != nil {
		t.Errorf("APP-054 audit = %+v; want enabled with agent fallback", config.Audit)
	}
	var rawConfig map[string]any
	if err := yaml.Unmarshal([]byte(readInitFile(t, root, "fireweed.yaml")), &rawConfig); err == nil {
		if audit, ok := rawConfig["audit"].(map[string]any); ok {
			if _, present := audit["command"]; present {
				t.Error("APP-054 audit.command must be unset for agent fallback")
			}
		}
	}
	if ignore := readInitFile(t, root, ".gitignore"); !strings.Contains(ignore, "custom-entry") || !strings.Contains(ignore, ".fireweed/runs/") {
		t.Errorf("APP-059 .gitignore = %q", ignore)
	}
	if status, stdout, stderr := runAppBinary(t, bin, root, "modules", "--json"); status != 0 || stderr != "" {
		t.Errorf("APP-050 generated Go graph modules status=%d stdout=%q stderr=%q", status, stdout, stderr)
	} else {
		envelope := decodeEnvelope(t, stdout)
		var modules []struct {
			Name      string   `json:"name"`
			IDPrefix  string   `json:"id_prefix"`
			Dir       string   `json:"dir"`
			DependsOn []string `json:"depends_on"`
			Zones     struct {
				Interface  []string `json:"interface"`
				Evals      []string `json:"evals"`
				Disposable []string `json:"disposable"`
			} `json:"zones"`
		}
		if err := json.Unmarshal(envelope["data"], &modules); err != nil || len(modules) != 1 {
			t.Errorf("APP-050 generated modules data=%s: %v", envelope["data"], err)
		} else if module := modules[0]; module.Name != "app" || module.IDPrefix != "APP" || module.Dir != "packages/app" || len(module.DependsOn) != 0 || !reflect.DeepEqual(module.Zones.Interface, []string{"api.go"}) || !reflect.DeepEqual(module.Zones.Evals, []string{"evals/**"}) || !reflect.DeepEqual(module.Zones.Disposable, []string{"internal/**"}) {
			t.Errorf("APP-050 generated app graph = %+v", module)
		}
	}
	if status, stdout, stderr := runAppBinary(t, bin, root, "check"); status != 0 || stderr != "" {
		t.Errorf("APP-050 generated Go graph check status=%d stdout=%q stderr=%q", status, stdout, stderr)
	}
}

func Test_APP_043_APP_044_APP_047_APP_051_NoneInit(t *testing.T) {
	bin := buildAppBinary(t)
	root := t.TempDir()
	status, stdout, stderr := runAppBinary(t, bin, root, "init", "--recipe", "none", "--agent", "claude", "--audit-agent", "codex", "--no-app")
	if status != 0 || stderr != "" {
		t.Errorf("none init status=%d stderr=%q", status, stderr)
	}
	for _, word := range []string{"eval", "boundaries", "setup"} {
		if !strings.Contains(strings.ToLower(stdout), word) {
			t.Errorf("APP-047 instructions omit %q: %q", word, stdout)
		}
	}
	config := readInitConfig(t, root)
	if len(config.Hooks) != 0 || len(config.ModuleDefaults.Protected)+len(config.ModuleDefaults.Interface)+len(config.ModuleDefaults.Evals)+len(config.ModuleDefaults.AppendOnly)+len(config.ModuleDefaults.Preserve)+len(config.ModuleDefaults.Disposable) != 0 {
		t.Errorf("APP-047 none has language defaults or hooks: %+v", config)
	}
	if _, err := os.Stat(filepath.Join(root, ".fireweed/templates")); !os.IsNotExist(err) {
		t.Errorf("APP-047 none created templates: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, ".fireweed/canaries")); !os.IsNotExist(err) {
		t.Errorf("APP-047 none created canaries: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "packages/app")); !os.IsNotExist(err) {
		t.Errorf("APP-051 --no-app created module: %v", err)
	}
	if got := initCommand(config.Agent.Command); !reflect.DeepEqual(got, []string{"claude", "-p", "{prompt}", "--permission-mode", "acceptEdits"}) {
		t.Errorf("APP-053 claude command = %q", got)
	}
	if got := initCommand(config.Audit.Command); !config.Audit.Enabled || !reflect.DeepEqual(got, []string{"codex", "exec", "{prompt}"}) {
		t.Errorf("APP-054 audit preset = %+v", config.Audit)
	}
	if system := readInitFile(t, root, "SYSTEM.md"); strings.TrimSpace(system) == "" {
		t.Error("APP-043 missing editable SYSTEM.md scaffold")
	}
	if agents := readInitFile(t, root, "AGENTS.md"); !strings.Contains(agents, "playbook author-change") || !strings.Contains(agents, "context") {
		t.Errorf("APP-044 missing managed guidance: %q", agents)
	}
}

func Test_APP_050_APP_052_NoneDefaultApp(t *testing.T) {
	bin := buildAppBinary(t)
	root := t.TempDir()
	status, _, stderr := runAppBinary(t, bin, root, "init", "--recipe", "none", "--agent", "codex")
	if status != 0 || stderr != "" {
		t.Errorf("none init status=%d stderr=%q", status, stderr)
	}
	spec := readInitFile(t, root, "packages/app/SPEC.md")
	for _, part := range []string{"name: app", "id_prefix: APP", "depends_on: []", "## Specs", "## Decisions"} {
		if !strings.Contains(spec, part) {
			t.Errorf("APP-050/052 minimal SPEC lacks %q: %q", part, spec)
		}
	}
	if strings.Contains(spec, "### ") {
		t.Errorf("APP-052 minimal SPEC is not empty: %q", spec)
	}
}

func Test_APP_042_APP_048_APP_049_APP_050_APP_053_APP_054_APP_060_LocalInit(t *testing.T) {
	bin := buildAppBinary(t)
	root := t.TempDir()
	recipeDir := filepath.Join(root, "recipe")
	instructions := "# Local setup\nConfigure the eval, boundaries and setup hooks.\n"
	appPut(t, recipeDir, "recipe.yaml", "module_defaults:\n  interface: [api.txt]\n  evals: [evals/**]\n  disposable: [internal/**]\n  preserve: [data/**]\n  append_only: [log/**]\nhooks:\n  eval: 'touch HOOK_RAN'\n  boundaries: 'touch BOUNDARIES_RAN'\n  setup: 'touch SETUP_RAN'\nprotected: [.fireweed/**, SYSTEM.md]\nworker_writable: [local.lock]\n")
	appPut(t, recipeDir, "recipe.md", instructions)
	appPut(t, recipeDir, "files/assets/banner.txt", "sentinel payload\n")
	appPut(t, recipeDir, "files/.fireweed/templates/module/SPEC.md", "---\nname: {{name}}\nid_prefix: {{prefix}}\ndepends_on: []\n---\n# {{name}}\n\n## Specs\n\n## Decisions\n")
	appPut(t, recipeDir, "files/.fireweed/templates/module/api.txt.tmpl", "package={{package}} name={{name}} prefix={{prefix}}\n")
	agent := "touch AGENT_RAN && printf '%s' {prompt}"
	auditor := "touch AUDITOR_RAN && printf '%s' {prompt}"
	status, stdout, stderr := runAppBinary(t, bin, root, "init", "--recipe", "./recipe", "--agent", agent, "--audit-agent", auditor)
	assertInitOutput(t, status, stdout, stderr, instructions)
	config := readInitConfig(t, root)
	if !reflect.DeepEqual(config.ModuleDefaults.Interface, []string{"api.txt"}) || !reflect.DeepEqual(config.ModuleDefaults.Evals, []string{"evals/**"}) || !reflect.DeepEqual(config.ModuleDefaults.Disposable, []string{"internal/**"}) || !reflect.DeepEqual(config.ModuleDefaults.Preserve, []string{"data/**"}) || !reflect.DeepEqual(config.ModuleDefaults.AppendOnly, []string{"log/**"}) || config.Hooks["eval"] != "touch HOOK_RAN" || config.Hooks["boundaries"] != "touch BOUNDARIES_RAN" || config.Hooks["setup"] != "touch SETUP_RAN" || !reflect.DeepEqual(config.Protected, []string{".fireweed/**", "SYSTEM.md"}) || !reflect.DeepEqual(config.WorkerWritable, []string{"local.lock"}) {
		t.Errorf("APP-042/048 local metadata differs: %+v", config)
	}
	if got := readInitFile(t, root, "assets/banner.txt"); got != "sentinel payload\n" {
		t.Errorf("APP-049 local payload = %q", got)
	}
	if got := readInitFile(t, root, ".fireweed/recipes/local.md"); got != instructions {
		t.Errorf("APP-049 copied instructions = %q", got)
	}
	if got := readInitFile(t, root, "packages/app/api.txt"); got != "package=app name=app prefix=APP\n" {
		t.Errorf("APP-050 rendered template = %q", got)
	}
	if spec := readInitFile(t, root, "packages/app/SPEC.md"); !strings.Contains(spec, "name: app") || !strings.Contains(spec, "id_prefix: APP") {
		t.Errorf("APP-050 rendered SPEC = %q", spec)
	}
	if config.Agent.Command != agent {
		t.Errorf("APP-053 custom command = %#v", config.Agent.Command)
	}
	if !config.Audit.Enabled || config.Audit.Command != auditor {
		t.Errorf("APP-054 custom audit command = %+v", config.Audit)
	}
	for _, marker := range []string{"HOOK_RAN", "BOUNDARIES_RAN", "SETUP_RAN", "AGENT_RAN", "AUDITOR_RAN"} {
		if _, err := os.Stat(filepath.Join(root, marker)); !os.IsNotExist(err) {
			t.Errorf("APP-060 init executed %s: %v", marker, err)
		}
	}
}

func Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name    string
		args    []string
		code    string
		choices bool
	}{
		{"missing-recipe", []string{"--agent", "codex"}, "U01", true},
		{"missing-agent", []string{"--recipe", "go"}, "U01", true},
		{"missing-value", []string{"--recipe", "go", "--agent"}, "U01", true},
		{"invalid-agent", []string{"--recipe", "go", "--agent", "fake-agent"}, "U01", false},
		{"malformed-option", []string{"--recipe", "go", "--agent", "codex", "--bad"}, "U01", false},
		{"extra-position", []string{"--recipe", "go", "--agent", "codex", "extra"}, "U01", false},
		{"unknown-recipe", []string{"--recipe", "ruby", "--agent", "codex"}, "U02", false},
		{"malformed-document", []string{"--recipe", "./recipe", "--agent", "codex"}, "C03", false},
		{"missing-document", []string{"--recipe", "./recipe", "--agent", "codex"}, "C03", false},
		{"invalid-metadata-type", []string{"--recipe", "./recipe", "--agent", "codex"}, "C12", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			switch tc.name {
			case "malformed-document":
				appPut(t, root, "recipe/recipe.yaml", "module_defaults: [broken\n")
				appPut(t, root, "recipe/recipe.md", "instructions\n")
			case "missing-document":
				appPut(t, root, "recipe/recipe.yaml", "module_defaults: {}\nhooks: {}\nprotected: []\nworker_writable: []\n")
			case "invalid-metadata-type":
				appPut(t, root, "recipe/recipe.yaml", "module_defaults: [wrong-type]\nhooks: {}\nprotected: []\nworker_writable: []\n")
				appPut(t, root, "recipe/recipe.md", "instructions\n")
			}
			args := append([]string{"init"}, tc.args...)
			status, stdout, stderr := runAppBinary(t, bin, root, args...)
			if status != 2 || !strings.HasPrefix(strings.TrimSpace(stdout+stderr), tc.code+":") || !corrective(stdout+stderr) {
				t.Errorf("status=%d output=%q; want corrective %s", status, stdout+stderr, tc.code)
			}
			if tc.choices {
				for _, choice := range []string{"go", "none", "codex", "claude"} {
					if !strings.Contains(stdout+stderr, choice) {
						t.Errorf("APP-055 missing-input diagnostic lacks %q: %q", choice, stdout+stderr)
					}
				}
			}
			assertNoInitConfig(t, root)
			for _, name := range []string{"SYSTEM.md", "AGENTS.md", ".gitignore", ".fireweed", "packages/app"} {
				if _, err := os.Stat(filepath.Join(root, name)); !os.IsNotExist(err) {
					t.Errorf("rejected init created %s: %v", name, err)
				}
			}
			jsonArgs := append(args, "--json")
			jsonStatus, jsonOut, jsonErr := runAppBinary(t, bin, root, jsonArgs...)
			if jsonStatus != 2 || jsonErr != "" {
				t.Errorf("APP-061 JSON status=%d stderr=%q", jsonStatus, jsonErr)
			}
			assertFailedJSON(t, jsonOut, tc.code)
		})
	}
}

func Test_APP_057_APP_058_InitConflicts(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct{ name, path, body string }{
		{"config", "fireweed.yaml", "existing config\n"},
		{"recipe", ".fireweed/hooks.go", "existing hook\n"},
		{"module", "packages/app/SPEC.md", "existing spec\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			appPut(t, root, tc.path, tc.body)
			appPut(t, root, "SYSTEM.md", "keep system\n")
			status, stdout, stderr := runAppBinary(t, bin, root, "init", "--recipe", "go", "--agent", "codex")
			if status != 2 || !strings.HasPrefix(strings.TrimSpace(stdout+stderr), "P01:") {
				t.Errorf("status=%d output=%q; want P01", status, stdout+stderr)
			}
			if got := readInitFile(t, root, tc.path); got != tc.body {
				t.Errorf("conflict file changed: %q", got)
			}
			if got := readInitFile(t, root, "SYSTEM.md"); got != "keep system\n" {
				t.Errorf("SYSTEM.md changed: %q", got)
			}
			if tc.name != "config" {
				assertNoInitConfig(t, root)
			}
		})
	}
}

func Test_APP_061_InitJSONSuccess(t *testing.T) {
	bin := buildAppBinary(t)
	root := t.TempDir()
	status, stdout, stderr := runAppBinary(t, bin, root, "init", "--recipe", "go", "--agent", "codex", "--json")
	if status != 0 || stderr != "" {
		t.Errorf("JSON init status=%d stderr=%q", status, stderr)
	}
	envelope := decodeEnvelope(t, stdout)
	var instructions string
	if err := json.Unmarshal(envelope["data"], &instructions); err != nil {
		t.Errorf("data is not instruction string: %v", err)
	}
	recipe, err := assets.BuiltinRecipe("go")
	if err != nil {
		t.Fatal(err)
	}
	if instructions != recipe.Instructions {
		t.Errorf("JSON instructions differ from recipe: %q", instructions)
	}
	var diagnostics []json.RawMessage
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
		t.Errorf("diagnostics=%s; want []", envelope["diagnostics"])
	}
}
