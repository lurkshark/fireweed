package lifecycle

import (
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

// dependsOn follows the dependency graph in the candidate's projected snapshot.
func dependsOn(graph map[string][]string, module, dependency string) bool {
	seen := map[string]bool{}
	var visit func(string) bool
	visit = func(name string) bool {
		if seen[name] {
			return false
		}
		seen[name] = true
		for _, next := range graph[name] {
			if next == dependency || visit(next) {
				return true
			}
		}
		return false
	}
	return visit(module)
}

func Eligible(root string) ([]project.Change, error) {
	queue, err := project.Queue(root)
	if err != nil {
		return nil, err
	}
	eligible := make([]project.Change, 0, len(queue))
	var blocked []string
	for _, candidate := range queue {
		if str(candidate.Document.Frontmatter["blocked"]) != "" {
			blocked = append(blocked, candidate.Module)
			continue
		}
		projected, err := project.Project(root, candidate.Slug)
		if err != nil {
			return nil, err
		}
		graph := make(map[string][]string, len(projected.Modules))
		for _, module := range projected.Modules {
			graph[module.Name] = module.DependsOn
		}
		section, _ := contracts.Section(candidate.Document.Body, "Interface")
		interfaceChanged := strings.TrimSpace(section) != ""
		waiting := false
		for _, earlier := range blocked {
			if candidate.Module == earlier || dependsOn(graph, candidate.Module, earlier) ||
				interfaceChanged && dependsOn(graph, earlier, candidate.Module) {
				waiting = true
				break
			}
		}
		if !waiting {
			eligible = append(eligible, candidate)
		}
	}
	return eligible, nil
}
