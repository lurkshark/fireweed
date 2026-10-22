package boundaries

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks/internal/hooks"
	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

const baselinePath = ".fireweed/boundaries-baseline"

func diagnostic(code, path, message string) model.Diagnostic {
	return model.Diagnostic{Code: code, Severity: "error", Path: path, Message: message}
}
func completed(ds []model.Diagnostic) []model.Diagnostic {
	result := make([]model.Diagnostic, 0, len(ds))
	seen := map[model.Diagnostic]bool{}
	for _, d := range ds {
		if !seen[d] {
			result = append(result, d)
			seen[d] = true
		}
	}
	sort.Slice(result, func(i, j int) bool {
		a, b := result[i], result[j]
		if a.Path != b.Path {
			return a.Path < b.Path
		}
		if a.Code != b.Code {
			return a.Code < b.Code
		}
		if a.Severity != b.Severity {
			return a.Severity < b.Severity
		}
		return a.Message < b.Message
	})
	return result
}
func failed(exit int, output string, err error) model.Diagnostic {
	message := fmt.Sprintf("Boundary hook exited %d. Captured output: %s", exit, output)
	if err != nil {
		message += fmt.Sprintf(" (%v)", err)
	}
	return diagnostic("C15", "", message+". Repair the checker and correct module dependency/interface violations before trusting boundary evidence.")
}
func optedOut(m contracts.Module) bool { return strings.TrimSpace(m.BoundariesOptOut) != "" }

func Check(ctx context.Context, snapshot project.Snapshot) []model.Diagnostic {
	if ctx == nil {
		return completed([]model.Diagnostic{failed(2, "", errors.New("provide a non-nil boundary context"))})
	}
	if err := ctx.Err(); err != nil {
		return completed([]model.Diagnostic{failed(-1, "", err)})
	}
	ds := []model.Diagnostic{}
	commands := map[string]bool{}
	modules := map[string]contracts.Module{}
	for _, module := range snapshot.Modules {
		modules[module.Name] = module
		if optedOut(module) {
			continue
		}
		command := module.Hooks.Boundaries
		if strings.TrimSpace(command) == "" {
			if len(module.Disposable) > 0 {
				ds = append(ds, diagnostic("C07", module.Dir, "Configure a boundaries hook for this disposable module or declare a reasoned boundaries_opt_out."))
			}
		} else {
			commands[command] = true
		}
	}
	if len(ds) > 0 || len(commands) == 0 {
		return completed(ds)
	}
	repo, err := gitstore.Open(snapshot.Root)
	if err != nil {
		return completed([]model.Diagnostic{failed(2, "", err)})
	}
	path, err := gitstore.SafePath(repo.Root, baselinePath)
	if err != nil {
		return completed([]model.Diagnostic{diagnostic("C15", baselinePath, "Cannot safely resolve boundary baseline: "+err.Error()+". Repair the baseline path inside the repository.")})
	}
	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil && !os.IsNotExist(err) {
		return completed([]model.Diagnostic{diagnostic("C15", baselinePath, "Cannot read boundary baseline: "+err.Error()+". Repair the baseline as a readable repository file.")})
	}
	baseline := map[string]bool{}
	for _, line := range strings.Split(string(data), "\n") {
		if line = strings.TrimSpace(line); line != "" {
			baseline[line] = true
		}
	}
	ordered := make([]string, 0, len(commands))
	for command := range commands {
		ordered = append(ordered, command)
	}
	sort.Strings(ordered)
	for _, command := range ordered {
		result := hooks.Run(ctx, snapshot, nil, command, "")
		var exitError *exec.ExitError
		ordinaryFailure := result.Err != nil && errors.As(result.Err, &exitError) && result.Exit >= 0
		if result.Exit < 0 || result.Err != nil && !ordinaryFailure {
			ds = append(ds, failed(result.Exit, result.Output, result.Err))
			continue
		}
		attributed, unattributed := 0, false
		for _, raw := range strings.Split(result.Output, "\n") {
			line := strings.TrimSpace(raw)
			if line == "" {
				continue
			}
			name, _, colon := strings.Cut(line, ":")
			if colon {
				name = strings.TrimSpace(name)
			} else {
				name = strings.Fields(line)[0]
			}
			module, known := modules[name]
			if !known {
				unattributed = true
				continue
			}
			attributed++
			if optedOut(module) || baseline[line] {
				continue
			}
			ds = append(ds, diagnostic("C15", module.Dir, "Boundary violation: "+line+". Correct imports to respect depends_on and keep evals within the public interface."))
		}
		if result.Exit != 0 && (attributed == 0 || unattributed) {
			ds = append(ds, failed(result.Exit, result.Output, result.Err))
		}
	}
	return completed(ds)
}
