package engine

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

func mapValue(parent map[string]any, key string) map[string]any {
	value, _ := parent[key].(map[string]any)
	return value
}

var placeholder = regexp.MustCompile(`\{(?:module_dir|module|report)\}`)

func shellQuote(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "'\"'\"'") + "'"
}

func hookArgs(command any, values map[string]string) ([]string, error) {
	replace := func(s string, quote bool) string {
		return placeholder.ReplaceAllStringFunc(s, func(token string) string {
			value := values[strings.Trim(token, "{}")]
			if quote {
				return shellQuote(value)
			}
			return value
		})
	}
	switch v := command.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			return nil, fmt.Errorf("empty hook command")
		}
		return []string{"sh", "-c", replace(v, true)}, nil
	case []any:
		args := make([]string, 0, len(v))
		for _, arg := range v {
			s, ok := arg.(string)
			if !ok {
				return nil, fmt.Errorf("hook argv contains non-string")
			}
			args = append(args, replace(s, false))
		}
		if len(args) == 0 || args[0] == "" {
			return nil, fmt.Errorf("empty hook argv")
		}
		return args, nil
	case []string:
		args := make([]string, len(v))
		for i, s := range v {
			args[i] = replace(s, false)
		}
		if len(args) == 0 || args[0] == "" {
			return nil, fmt.Errorf("empty hook argv")
		}
		return args, nil
	default:
		return nil, fmt.Errorf("invalid hook command %T", command)
	}
}

func runHook(ctx context.Context, root string, command any, values map[string]string, bin string, timeout time.Duration) (gitstore.Result, error) {
	args, err := hookArgs(command, values)
	if err != nil {
		return gitstore.Result{}, err
	}
	result, err := gitstore.Run(ctx, gitstore.Request{Args: args, Dir: root, Env: map[string]string{"FIREWEED_BIN": bin}, Timeout: timeout})
	if err != nil {
		var exit *exec.ExitError
		if !errors.As(err, &exit) {
			return result, fmt.Errorf("run hook: %w", err)
		}
	}
	return result, nil
}

func moduleSpecs(module project.Module) ([]contracts.Item, error) {
	section, ok := contracts.Section(module.Document.Body, "Specs")
	if !ok {
		return nil, nil
	}
	items, err := contracts.ParseItems(section)
	if err != nil {
		code, _, ok := contracts.ErrorCode(err)
		if !ok {
			return nil, fmt.Errorf("parse module specs: %w", err)
		}
		return nil, fmt.Errorf("%s: %w", code, err)
	}
	return items, nil
}

func Evaluate(ctx context.Context, request model.HookRequest) (model.Report, error) {
	var module *project.Module
	for i := range request.Snapshot.Modules {
		if request.Snapshot.Modules[i].Name == request.Module {
			module = &request.Snapshot.Modules[i]
			break
		}
	}
	if module == nil {
		return model.Report{}, fmt.Errorf("U02: unknown module %q", request.Module)
	}
	command := mapValue(request.Snapshot.Config, "hooks")["eval"]
	if override, ok := mapValue(module.Document.Frontmatter, "hooks")["eval"]; ok {
		command = override
	}
	if command == nil {
		return model.Report{}, fmt.Errorf("C11: missing eval hook for %s", module.Name)
	}
	specs, err := moduleSpecs(*module)
	if err != nil {
		return model.Report{}, err
	}
	temp, err := os.MkdirTemp("", "fireweed-junit-")
	if err != nil {
		return model.Report{}, fmt.Errorf("C11: %w", err)
	}
	defer os.RemoveAll(temp)
	reportPath := filepath.Join(temp, "report.xml")
	values := map[string]string{"module": module.Name, "module_dir": module.Dir, "report": reportPath}
	process, err := runHook(ctx, request.Snapshot.Root, command, values, request.Bin, request.Timeout)
	if err != nil {
		return model.Report{}, fmt.Errorf("C11: %w", err)
	}
	data, err := os.ReadFile(reportPath)
	if err != nil {
		return model.Report{}, fmt.Errorf("C11: read JUnit report: %w", err)
	}
	report, err := ParseJUnit(data, specs)
	if err != nil {
		return model.Report{}, err
	}
	report.HookExit, report.Stdout, report.Stderr = process.ExitCode, process.Stdout, process.Stderr
	failed := false
	for _, c := range report.Cases {
		if c.Result == "fail" {
			failed = true
			break
		}
	}
	if (process.ExitCode != 0) != failed {
		return report, fmt.Errorf("C11: eval hook exit %d disagrees with JUnit failures", process.ExitCode)
	}
	return report, nil
}
