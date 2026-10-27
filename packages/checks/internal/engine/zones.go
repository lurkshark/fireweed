package engine

import (
	"fmt"
	"path"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

func patterns(value any) []string {
	switch list := value.(type) {
	case []string:
		return list
	case []any:
		result := make([]string, 0, len(list))
		for _, item := range list {
			if s, ok := item.(string); ok {
				result = append(result, s)
			}
		}
		return result
	default:
		return nil
	}
}

func matches(patterns []string, file string) bool {
	for _, pattern := range patterns {
		if contracts.MatchGlob(pattern, file) {
			return true
		}
	}
	return false
}

func phaseBaseHas(request model.ZoneRequest, file string) bool {
	if request.Base == "" {
		return false
	}
	_, err := gitstore.ReadAt(request.Snapshot.Root, request.Base, file)
	return err == nil
}

func zonePath(request model.ZoneRequest, file string) error {
	file = path.Clean(strings.TrimPrefix(file, "./"))
	if file == "." || file == ".." || strings.HasPrefix(file, "../") || path.IsAbs(file) {
		return fmt.Errorf("invalid path %q", file)
	}
	protected := []string{".git", ".git/**", ".fireweed", ".fireweed/**", "SYSTEM.md", "fireweed.yaml"}
	protected = append(protected, patterns(request.Snapshot.Config["protected"])...)
	if matches(protected, file) || path.Base(file) == "AGENTS.md" || path.Base(file) == "CLAUDE.md" {
		return fmt.Errorf("protected path %s", file)
	}
	if file == request.ChangePath {
		return nil
	}
	if strings.HasPrefix(file, "changes/") {
		return fmt.Errorf("foreign change %s", file)
	}
	if request.Phase == "audit" {
		return fmt.Errorf("audit may edit only current change: %s", file)
	}
	var owner *project.Module
	for i := range request.Snapshot.Modules {
		module := &request.Snapshot.Modules[i]
		if file == module.Dir || strings.HasPrefix(file, strings.TrimSuffix(module.Dir, "/")+"/") {
			if owner == nil || len(module.Dir) > len(owner.Dir) {
				owner = module
			}
		}
	}
	if owner == nil {
		if matches(patterns(request.Snapshot.Config["worker_writable"]), file) && (request.Phase == "evals" || request.Phase == "impl") {
			return nil
		}
		return fmt.Errorf("outside module: %s", file)
	}
	if owner.Name != request.Module {
		return fmt.Errorf("foreign module: %s", file)
	}
	rel := strings.TrimPrefix(file, strings.TrimSuffix(owner.Dir, "/")+"/")
	zone := contracts.ResolveZone(rel, owner.Zones)
	switch zone {
	case "protected", "spec":
		return fmt.Errorf("protected module path %s", file)
	case "evals":
		if request.Phase == "evals" {
			return nil
		}
	case "interface":
		if request.Phase == "evals" && request.InterfaceDeclared {
			return nil
		}
	case "disposable":
		if request.Phase == "impl" || request.Phase == "evals" && !phaseBaseHas(request, file) {
			return nil
		}
	case "append_only":
		if request.Phase == "impl" && !phaseBaseHas(request, file) {
			return nil
		}
	case "preserve", "other":
		if request.Phase == "impl" || zone == "other" && request.Phase == "evals" {
			return nil
		}
	}
	return fmt.Errorf("%s phase cannot edit %s (%s)", request.Phase, file, zone)
}

func Zones(request model.ZoneRequest) []project.Diagnostic {
	var diagnostics []project.Diagnostic
	for _, change := range request.Changes {
		if change.OldPath != "" {
			for _, module := range request.Snapshot.Modules {
				prefix := strings.TrimSuffix(module.Dir, "/") + "/"
				if strings.HasPrefix(change.Path, prefix) && matches(module.Zones.AppendOnly, strings.TrimPrefix(change.Path, prefix)) {
					diagnostics = append(diagnostics, project.Diagnostic{Code: "Z01", Path: change.Path, Message: "append-only path cannot be a rename destination"})
				}
			}
		}
		for _, file := range []string{change.OldPath, change.Path} {
			if file == "" {
				continue
			}
			if err := zonePath(request, file); err != nil {
				diagnostics = append(diagnostics, project.Diagnostic{Code: "Z01", Path: file, Message: err.Error()})
			}
		}
	}
	return diagnostics
}
