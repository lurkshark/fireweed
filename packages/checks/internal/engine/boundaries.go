package engine

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/project"
)

func boundaryReason(module project.Module) string {
	reason, _ := module.Document.Frontmatter["boundaries_opt_out"].(string)
	return strings.TrimSpace(reason)
}

func Boundaries(ctx context.Context, request model.BoundaryRequest) ([]project.Diagnostic, error) {
	snapshot := request.Snapshot
	baseline := map[string]bool{}
	data, err := os.ReadFile(filepath.Join(snapshot.Root, ".fireweed", "boundaries-baseline"))
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("read boundaries baseline: %w", err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			baseline[line] = true
		}
	}
	var diagnostics []project.Diagnostic
	var violations int
	global := mapValue(snapshot.Config, "hooks")["boundaries"]
	configured := global != nil
	for _, module := range snapshot.Modules {
		moduleHook := mapValue(module.Document.Frontmatter, "hooks")["boundaries"]
		if moduleHook != nil {
			configured = true
		}
		if global == nil && moduleHook == nil && len(module.Zones.Disposable) > 0 && boundaryReason(module) == "" {
			violations++
			diagnostics = append(diagnostics, project.Diagnostic{Code: "B05", Message: module.Name + ": missing boundaries hook"})
		}
		if reason := boundaryReason(module); reason != "" {
			diagnostics = append(diagnostics, project.Diagnostic{Code: "B05", Message: module.Name + ": boundaries opt-out: " + reason, Warning: true})
		}
	}
	if !configured {
		if violations == 0 {
			diagnostics = append(diagnostics, project.Diagnostic{Code: "B05", Message: "no boundary enforcement required; boundaries hook absent", Warning: true})
			return diagnostics, nil
		}
		return diagnostics, fmt.Errorf("missing required boundaries hook")
	}
	commands := []struct {
		command any
		module  project.Module
	}{}
	if global != nil {
		commands = append(commands, struct {
			command any
			module  project.Module
		}{command: global})
	}
	for _, module := range snapshot.Modules {
		if command := mapValue(module.Document.Frontmatter, "hooks")["boundaries"]; command != nil {
			commands = append(commands, struct {
				command any
				module  project.Module
			}{command: command, module: module})
		}
	}
	for _, hook := range commands {
		values := map[string]string{"module": hook.module.Name, "module_dir": hook.module.Dir}
		result, runErr := runHook(ctx, snapshot.Root, hook.command, values, request.Bin, request.Timeout)
		if runErr != nil {
			return diagnostics, fmt.Errorf("boundary hook: %w", runErr)
		}
		for _, output := range []string{result.Stdout, result.Stderr} {
			for _, line := range strings.Split(output, "\n") {
				line = strings.TrimSuffix(line, "\r")
				if line == "" || baseline[line] {
					continue
				}
				owner := ""
				for _, module := range snapshot.Modules {
					if strings.HasPrefix(line, module.Name+":") {
						owner = module.Name
						if boundaryReason(module) != "" {
							owner = "suppressed"
						}
						break
					}
				}
				if owner == "suppressed" {
					continue
				}
				if owner != "" {
					violations++
					diagnostics = append(diagnostics, project.Diagnostic{Code: "B05", Message: line})
				}
			}
		}
	}
	if violations > 0 {
		return diagnostics, fmt.Errorf("%d boundary violation(s)", violations)
	}
	return diagnostics, nil
}
