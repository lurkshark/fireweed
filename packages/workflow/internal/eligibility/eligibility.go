package eligibility

import (
	"fmt"
	"sort"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/internal/claims"
)

func problem(format string, args ...any) error {
	return fmt.Errorf("C05: %s. Fix the queued selection or projected dependency graph and retry.", fmt.Sprintf(format, args...))
}

func Check(snapshot project.Snapshot, slug string) (bool, error) {
	if !claims.ValidSlug(slug) {
		return false, problem("invalid exact change selector %q", slug)
	}
	matches := 0
	for _, change := range snapshot.Changes {
		if change.Slug == slug {
			matches++
		}
	}
	if matches != 1 {
		return false, problem("exact queued selector %q matches %d changes", slug, matches)
	}

	// Project also orders intent, but this copy defines which candidate is next
	// without changing the caller's queue or its spare backing storage.
	snapshot.Changes = append([]contracts.Change(nil), snapshot.Changes...)
	sort.SliceStable(snapshot.Changes, func(i, j int) bool {
		a, b := snapshot.Changes[i], snapshot.Changes[j]
		if a.Number != b.Number {
			return a.Number < b.Number
		}
		if a.Slug != b.Slug {
			return a.Slug < b.Slug
		}
		return a.Path < b.Path
	})
	selected := 0
	for i, change := range snapshot.Changes {
		if change.Slug == slug {
			selected = i
			break
		}
	}
	// Validate selected intent first, even if its ordinary ordering result would
	// already be false. Invalid intent must never masquerade as a held change.
	selectedGraph, err := inclusive(snapshot, snapshot.Changes[selected])
	if err != nil {
		return false, err
	}
	if snapshot.Changes[selected].Blocked != "" {
		return false, nil
	}

	var blockers []contracts.Change
	for i, candidate := range snapshot.Changes[:selected+1] {
		if candidate.Blocked != "" {
			blockers = append(blockers, candidate)
			continue
		}
		graph := selectedGraph
		if i != selected {
			graph, err = inclusive(snapshot, candidate)
			if err != nil {
				return false, err
			}
		}
		held := false
		for _, blocker := range blockers {
			if graph.holds(blocker, candidate) {
				held = true
				break
			}
		}
		if !held {
			return i == selected, nil
		}
	}
	return false, nil
}

type dependencyGraph map[string][]string

func inclusive(snapshot project.Snapshot, candidate contracts.Change) (dependencyGraph, error) {
	before, err := project.Project(snapshot, candidate.Slug)
	if err != nil {
		return nil, problem("cannot project prefix for %s: %v", candidate.Slug, err)
	}
	index, count := 0, 0
	for i, module := range before.Modules {
		if module.Name == candidate.Module {
			index, count = i, count+1
		}
	}
	if count != 1 {
		return nil, problem("change %s requires exactly one target module %q; found %d", candidate.Slug, candidate.Module, count)
	}
	module, err := project.ApplyChange(before.Modules[index], candidate)
	if err != nil {
		return nil, problem("cannot apply change %s: %v", candidate.Slug, err)
	}
	before.Modules[index] = module

	graph := make(dependencyGraph, len(before.Modules))
	for _, module := range before.Modules {
		if _, exists := graph[module.Name]; exists {
			return nil, problem("projected module %q is ambiguous", module.Name)
		}
		graph[module.Name] = module.DependsOn
	}
	// Validate the entire graph, including components unrelated to the selected
	// target. Iterate supplied order for deterministic errors, never map order.
	for _, module := range before.Modules {
		for _, dependency := range module.DependsOn {
			if _, exists := graph[dependency]; !exists {
				return nil, problem("module %s depends on missing module %q", module.Name, dependency)
			}
		}
	}
	state := make(map[string]uint8, len(graph))
	var visit func(string) error
	visit = func(name string) error {
		if state[name] == 1 {
			return problem("projected dependency cycle includes module %q", name)
		}
		if state[name] == 2 {
			return nil
		}
		state[name] = 1
		for _, dependency := range graph[name] {
			if err := visit(dependency); err != nil {
				return err
			}
		}
		state[name] = 2
		return nil
	}
	for _, module := range before.Modules {
		if err := visit(module.Name); err != nil {
			return nil, err
		}
	}
	return graph, nil
}

func (graph dependencyGraph) holds(blocker, candidate contracts.Change) bool {
	if blocker.Module == candidate.Module || graph.depends(candidate.Module, blocker.Module) {
		return true
	}
	for _, section := range candidate.Document.Sections {
		if section.Name == "Interface" && strings.TrimSpace(section.Body) != "" {
			return graph.depends(blocker.Module, candidate.Module)
		}
	}
	return false
}

func (graph dependencyGraph) depends(from, target string) bool {
	seen := make(map[string]bool)
	pending := append([]string(nil), graph[from]...)
	for len(pending) != 0 {
		name := pending[len(pending)-1]
		pending = pending[:len(pending)-1]
		if name == target {
			return true
		}
		if !seen[name] {
			seen[name] = true
			pending = append(pending, graph[name]...)
		}
	}
	return false
}
