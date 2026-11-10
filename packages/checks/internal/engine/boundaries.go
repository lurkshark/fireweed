package engine

import (
	"context"
	"errors"
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
			diagnostics = append(diagnostics, project.Diagnostic{Code: "X02", Message: module.Name + ": missing boundaries hook"})
		}
		if reason := boundaryReason(module); reason != "" {
			diagnostics = append(diagnostics, project.Diagnostic{Code: "X03", Message: module.Name + ": boundaries opt-out: " + reason, Warning: true})
		}
	}
	if !configured {
		if violations == 0 {
			diagnostics = append(diagnostics, project.Diagnostic{Code: "X04", Message: "no boundary enforcement required; boundaries hook absent", Warning: true})
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
	var hookErrors []error
	for _, hook := range commands {
		values := map[string]string{"module": hook.module.Name, "module_dir": hook.module.Dir}
		result, runErr := runHook(ctx, snapshot.Root, hook.command, values, request.Bin, request.Timeout)
		if runErr != nil {
			diagnostics = append(diagnostics, project.Diagnostic{Code: "X02", Message: "boundary hook: " + runErr.Error()})
			hookErrors = append(hookErrors, fmt.Errorf("boundary hook: %w", runErr))
			break
		}
		attributed := false
		for _, output := range []string{result.Stdout, result.Stderr} {
			for _, line := range strings.Split(output, "\n") {
				line = strings.TrimSuffix(line, "\r")
				for _, module := range snapshot.Modules {
					if strings.HasPrefix(line, module.Name+":") {
						attributed = true
						if !baseline[line] && boundaryReason(module) == "" {
							violations++
							diagnostics = append(diagnostics, project.Diagnostic{Code: "X01", Message: line})
						}
						break
					}
				}
			}
		}
		if result.ExitCode != 0 && !attributed {
			name := "global boundary hook"
			if hook.module.Name != "" {
				name = hook.module.Name + " boundary hook"
			}
			err := fmt.Errorf("%s: exit status %d without a module-attributed violation", name, result.ExitCode)
			diagnostics = append(diagnostics, project.Diagnostic{Code: "X02", Message: err.Error()})
			hookErrors = append(hookErrors, err)
		}
	}
	if len(hookErrors) == 1 {
		return diagnostics, hookErrors[0]
	}
	if len(hookErrors) > 1 {
		return diagnostics, errors.Join(hookErrors...)
	}
	if violations > 0 {
		return diagnostics, fmt.Errorf("%d boundary violation(s)", violations)
	}
	return diagnostics, nil
}
