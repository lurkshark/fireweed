package recipe

import (
	"embed"
	"fmt"

	"github.com/lurkshark/fireweed/packages/assets/model"
	"github.com/lurkshark/fireweed/packages/contracts"
)

//go:embed files/*
var bundle embed.FS

func BuiltinRecipe(name string) (model.Recipe, error) {
	if name != "go" && name != "typescript-pnpm" && name != "python-uv" {
		return model.Recipe{}, fmt.Errorf("unknown recipe %q", name)
	}
	if name != "go" {
		return additionalRecipe(name)
	}
	paths := map[string]string{
		".fireweed/recipes/go.md":                "files/go.md",
		".fireweed/hooks.go":                     "files/hooks.go",
		".fireweed/templates/module/SPEC.md":     "files/SPEC.md",
		".fireweed/templates/module/api.go.tmpl": "files/api.go.tmpl",
		".fireweed/canaries/canaries.yaml":       "files/canaries.yaml",
	}
	files := make(map[string]string, len(paths))
	for output, resource := range paths {
		data, err := bundle.ReadFile(resource)
		if err != nil {
			return model.Recipe{}, fmt.Errorf("read Go recipe %s: %w", resource, err)
		}
		files[output] = string(data)
	}
	return model.Recipe{
		Name:         "go",
		Instructions: files[".fireweed/recipes/go.md"],
		ModuleDefaults: contracts.Zones{
			Interface:  []string{"api.go"},
			Evals:      []string{"evals/**"},
			Disposable: []string{"internal/**"},
		},
		Hooks: map[string]any{
			"eval":       "go run .fireweed/hooks.go eval {module_dir} {report}",
			"boundaries": "go run .fireweed/hooks.go boundaries",
			"setup":      "go mod download",
		},
		Protected:      []string{".fireweed/**", "fireweed.yaml", "AGENTS.md", "SYSTEM.md"},
		WorkerWritable: []string{"go.sum"},
		Files:          files,
	}, nil
}

func additionalRecipe(name string) (model.Recipe, error) {
	paths := map[string]string{
		".fireweed/recipes/" + name + ".md": "files/" + name + ".md",
		".fireweed/templates/module/SPEC.md": "files/" + name + "-SPEC.md",
		".fireweed/canaries/canaries.yaml": "files/" + name + "-canaries.yaml",
	}
	var zones contracts.Zones
	var hooks map[string]any
	var writable []string
	if name == "typescript-pnpm" {
		paths[".fireweed/hooks.mjs"] = "files/typescript-pnpm-hooks.mjs"
		paths[".fireweed/templates/module/package.json.tmpl"] = "files/typescript-pnpm-package.json.tmpl"
		paths[".fireweed/templates/module/src/index.ts.tmpl"] = "files/typescript-pnpm-index.ts.tmpl"
		zones = contracts.Zones{Interface: []string{"src/index.ts"}, Evals: []string{"evals/**"}, Disposable: []string{"src/internal/**"}}
		hooks = map[string]any{"eval": "pnpm exec node .fireweed/hooks.mjs eval {module_dir} {report}", "boundaries": "pnpm exec node .fireweed/hooks.mjs boundaries", "setup": "pnpm install --offline"}
		writable = []string{"pnpm-lock.yaml"}
	} else {
		paths[".fireweed/hooks.py"] = "files/python-uv-hooks.py"
		paths[".fireweed/templates/module/pyproject.toml.tmpl"] = "files/python-uv-pyproject.toml.tmpl"
		paths[".fireweed/templates/module/__init__.py.tmpl"] = "files/python-uv-init.py.tmpl"
		zones = contracts.Zones{Interface: []string{"__init__.py"}, Evals: []string{"evals/**"}, Disposable: []string{"internal/**"}}
		hooks = map[string]any{"eval": "uv run --no-project python .fireweed/hooks.py eval {module_dir} {report}", "boundaries": "uv run --no-project python .fireweed/hooks.py boundaries", "setup": "uv sync --offline"}
		writable = []string{"uv.lock"}
	}
	files := make(map[string]string, len(paths))
	for output, resource := range paths {
		data, err := bundle.ReadFile(resource)
		if err != nil {
			return model.Recipe{}, fmt.Errorf("read %s recipe %s: %w", name, resource, err)
		}
		files[output] = string(data)
	}
	return model.Recipe{Name: name, Instructions: files[".fireweed/recipes/"+name+".md"], ModuleDefaults: zones, Hooks: hooks,
		Protected: []string{".fireweed/**", "fireweed.yaml", "AGENTS.md", "SYSTEM.md"}, WorkerWritable: writable, Files: files}, nil
}
