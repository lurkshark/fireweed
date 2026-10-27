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
	if name != "go" {
		return model.Recipe{}, fmt.Errorf("unknown recipe %q", name)
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
