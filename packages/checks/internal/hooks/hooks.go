package hooks

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/checks/internal/evidence"
	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

func invalid(message string) gitstore.CommandResult {
	return gitstore.CommandResult{Exit: 2, Err: errors.New(message)}
}
func quote(value string) string { return "'" + strings.ReplaceAll(value, "'", `'"'"'`) + "'" }
func Run(ctx context.Context, snapshot project.Snapshot, module *contracts.Module, command, report string) gitstore.CommandResult {
	if ctx == nil {
		return invalid("provide a non-nil hook context")
	}
	if err := ctx.Err(); err != nil {
		return gitstore.CommandResult{Exit: -1, Err: err}
	}
	if strings.TrimSpace(command) == "" {
		return invalid("configure a nonempty hook command")
	}
	timeout, err := time.ParseDuration(snapshot.Config.Timeout)
	if err != nil || timeout <= 0 {
		return invalid("configure a valid positive hook timeout")
	}
	repo, err := gitstore.Open(snapshot.Root)
	if err != nil {
		return invalid("discover hook repository: " + err.Error())
	}
	binary, err := os.Executable()
	if err != nil {
		return invalid("resolve current hook executable: " + err.Error())
	}
	name, dir := "", ""
	if module != nil {
		name, dir = module.Name, module.Dir
	}
	// NewReplacer scans the original command once; inserted values are never rescanned.
	expand := strings.NewReplacer("{module}", quote(name), "{module_dir}", quote(dir), "{report}", quote(report), "{git_common_dir}", quote(repo.CommonDir))
	bounded, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	env := []string{"FIREWEED_MODULE=" + name, "FIREWEED_MODULE_DIR=" + dir, "FIREWEED_REPORT=" + report, "FIREWEED_BIN=" + binary}
	return gitstore.Run(bounded, repo.Root, []string{"sh", "-c", expand.Replace(command)}, env)
}
func emptyCoverage(module contracts.Module) model.Coverage {
	return model.Coverage{Module: module.Name, Tests: []model.TestCase{}, BySpec: map[string][]int{}, Diagnostics: []model.Diagnostic{}}
}
func problem(cov model.Coverage, module contracts.Module, code, message string) model.Coverage {
	cov.Diagnostics = append(cov.Diagnostics, model.Diagnostic{Code: code, Severity: "error", Path: module.Dir, Message: message})
	return cov
}
func Evaluate(ctx context.Context, snapshot project.Snapshot, module contracts.Module, failFirst, already []string) model.Coverage {
	cov := emptyCoverage(module)
	if strings.TrimSpace(module.Hooks.Eval) == "" {
		return problem(cov, module, "C07", "Configure the module's eval hook before running coverage.")
	}
	directory, err := os.MkdirTemp("", "fireweed-eval-")
	if err != nil {
		return problem(cov, module, "C11", "Cannot create an owned report directory: "+err.Error()+". Repair the temporary directory configuration.")
	}
	defer os.RemoveAll(directory)
	directory, err = filepath.Abs(directory)
	if err != nil {
		return problem(cov, module, "C11", "Cannot resolve the report directory: "+err.Error()+". Repair the temporary directory configuration.")
	}
	result := Run(ctx, snapshot, &module, module.Hooks.Eval, filepath.Join(directory, "junit.xml"))
	cov.Output, cov.HookExit = result.Output, result.Exit
	var exited *exec.ExitError
	ordinaryFailure := result.Err != nil && errors.As(result.Err, &exited) && result.Exit >= 0
	if result.Exit < 0 || result.Err != nil && !ordinaryFailure {
		return problem(cov, module, "C11", fmt.Sprintf("Cannot complete eval hook: %v. Repair hook configuration or cancellation before trusting its report.", result.Err))
	}
	report, err := os.ReadFile(filepath.Join(directory, "junit.xml"))
	if err != nil {
		return problem(cov, module, "C11", "Cannot read the fresh JUnit report: "+err.Error()+". Configure the eval hook to write the supplied report path.")
	}
	tests, err := evidence.Parse(report)
	if err != nil {
		return problem(cov, module, "C11", "Invalid JUnit report: "+err.Error()+". Repair the reporter to produce one complete XML document.")
	}
	prefixes := make([]string, 0, len(snapshot.Modules))
	for _, m := range snapshot.Modules {
		prefixes = append(prefixes, m.IDPrefix)
	}
	cov = evidence.Analyze(module, prefixes, tests, result.Exit, failFirst, already)
	cov.Output, cov.HookExit = result.Output, result.Exit
	return cov
}
