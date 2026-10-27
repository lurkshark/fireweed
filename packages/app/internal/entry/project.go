package entry

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/lurkshark/fireweed/packages/project"
)

type moduleResult struct {
	Name      string     `json:"name"`
	IDPrefix  string     `json:"id_prefix"`
	Dir       string     `json:"dir"`
	DependsOn []string   `json:"depends_on"`
	Zones     zoneResult `json:"zones"`
}

type zoneResult struct {
	Protected  []string `json:"protected"`
	Evals      []string `json:"evals"`
	Interface  []string `json:"interface"`
	AppendOnly []string `json:"append_only"`
	Preserve   []string `json:"preserve"`
	Disposable []string `json:"disposable"`
}

func loadValidatedProject() (project.Snapshot, error) {
	root, err := findProjectRoot()
	if err != nil {
		return project.Snapshot{}, err
	}
	snapshot, err := project.Load(root)
	if err != nil {
		return project.Snapshot{}, err
	}
	if err := project.ValidateGraph(snapshot.Modules); err != nil {
		return project.Snapshot{}, err
	}
	return snapshot, nil
}

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("find current directory: %w", err)
	}
	for {
		_, err := os.Stat(filepath.Join(dir, "fireweed.yaml"))
		if err == nil {
			return dir, nil
		}
		if !os.IsNotExist(err) {
			return "", fmt.Errorf("inspect fireweed.yaml in %s: %w", dir, err)
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("fireweed.yaml not found in the current directory or its ancestors")
		}
		dir = parent
	}
}

func moduleResults(modules []project.Module) []moduleResult {
	result := make([]moduleResult, 0, len(modules))
	for _, module := range modules {
		result = append(result, moduleResult{
			Name: module.Name, IDPrefix: module.IDPrefix, Dir: module.Dir,
			DependsOn: nonNil(module.DependsOn),
			Zones: zoneResult{
				Protected: nonNil(module.Zones.Protected), Evals: nonNil(module.Zones.Evals),
				Interface: nonNil(module.Zones.Interface), AppendOnly: nonNil(module.Zones.AppendOnly),
				Preserve: nonNil(module.Zones.Preserve), Disposable: nonNil(module.Zones.Disposable),
			},
		})
	}
	return result
}

func nonNil(items []string) []string {
	if items == nil {
		return []string{}
	}
	return items
}

func printModules(out io.Writer, modules []moduleResult) {
	for _, module := range modules {
		fmt.Fprintf(out, "%s (%s) %s\n", module.Name, module.IDPrefix, module.Dir)
		fmt.Fprintf(out, "  depends_on: %v\n", module.DependsOn)
		fmt.Fprintf(out, "  protected: %v\n  evals: %v\n  interface: %v\n", module.Zones.Protected, module.Zones.Evals, module.Zones.Interface)
		fmt.Fprintf(out, "  append_only: %v\n  preserve: %v\n  disposable: %v\n", module.Zones.AppendOnly, module.Zones.Preserve, module.Zones.Disposable)
	}
}
