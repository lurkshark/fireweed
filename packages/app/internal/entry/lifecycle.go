package entry

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func contains(values []string, value string) bool {
	for _, item := range values {
		if item == value {
			return true
		}
	}
	return false
}

func cliFailure(status int, jsonMode bool, out, errOut io.Writer, code, message string) int {
	return cliResponse(status, jsonMode, out, errOut, nil, "", []diagnostic{{Code: code, Message: message}})
}

func cliResponse(status int, jsonMode bool, out, errOut io.Writer, data any, plain string, diagnostics []diagnostic) int {
	if diagnostics == nil {
		diagnostics = []diagnostic{}
	}
	if jsonMode {
		_ = json.NewEncoder(out).Encode(envelope{Schema: 1, Data: data, Diagnostics: diagnostics})
	} else {
		if plain != "" {
			_, _ = io.WriteString(out, plain)
		}
		for _, d := range diagnostics {
			fmt.Fprintf(errOut, "%s: %s\n", d.Code, d.Message)
		}
	}
	return status
}

func reported(ds []project.Diagnostic) []diagnostic {
	result := make([]diagnostic, 0, len(ds))
	for _, d := range ds {
		result = append(result, diagnostic{Code: d.Code, Message: d.Message, Path: d.Path, Warning: d.Warning})
	}
	return result
}

var diagnosticPrefix = regexp.MustCompile(`^([A-Z][A-Z0-9-]+):\s*`)

func lifecycleError(err error, ds []diagnostic, action string) []diagnostic {
	message := err.Error()
	code := "P01"
	if parts := diagnosticPrefix.FindStringSubmatch(message); parts != nil {
		code = parts[1]
		message = strings.TrimPrefix(message, parts[0])
	}
	// Keep the source diagnostics, including warnings, and add an actionable error.
	return append(ds, diagnostic{Code: code, Message: message + ". " + action})
}

func loadModuleProject() (project.Snapshot, error) {
	snapshot, err := loadValidatedProject()
	if err != nil {
		return project.Snapshot{}, err
	}
	// A plain project fixture is valid without Git. Only inspect claims in a
	// repository; Current supplies the approved graph saved with the claim.
	if _, err := os.Stat(filepath.Join(snapshot.Root, ".git")); err == nil {
		if state, err := workflow.Current(snapshot.Root, ""); err == nil {
			return state.Snapshot, nil
		}
	}
	return snapshot, nil
}

func hookTimeout(snapshot project.Snapshot) time.Duration {
	if value, ok := snapshot.Config["timeout"].(string); ok {
		if duration, err := time.ParseDuration(value); err == nil && duration > 0 {
			return duration
		}
	}
	return 2 * time.Minute
}

func runLifecycle(command string, args []string, jsonMode, quick, full bool, withChange string, fromAttempt int, out, errOut io.Writer) int {
	fail := func(status int, code, message string) int {
		return cliFailure(status, jsonMode, out, errOut, code, message)
	}
	if command == "start" || command == "context" || command == "release" || command == "commits" || command == "abandon" {
		if len(args) != 1 {
			return fail(2, "U01", "Use "+command+" <change> with exactly one change.")
		}
	} else if command == "coverage" {
		if len(args) != 1 {
			return fail(2, "U01", "Use coverage <module> [--with-change <change>].")
		}
	} else if len(args) > 1 {
		return fail(2, "U01", "Use "+command+" [change] with at most one change.")
	}
	snapshot, err := loadValidatedProject()
	if err != nil {
		return fail(2, projectFailureCode(err, "C01"), fmt.Sprintf("%v. Provide a valid fireweed.yaml and module declarations, then retry.", err))
	}
	root := snapshot.Root
	if command == "start" {
		var state workflow.State
		if fromAttempt > 0 {
			state, err = workflow.StartFromAttempt(root, args[0], fromAttempt)
		} else {
			state, err = workflow.Start(root, args[0])
		}
		ds := reported(state.Diagnostics)
		if err != nil {
			return cliResponse(1, jsonMode, out, errOut, nil, "", lifecycleError(err, ds, "Use an eligible queued change and clear conflicting claims, then retry start."))
		}
		body, err := workflow.Context(root, args[0], false)
		if err != nil {
			return cliResponse(1, jsonMode, out, errOut, nil, "", lifecycleError(err, ds, "Use context <change> after checking the held claim."))
		}
		return cliResponse(0, jsonMode, out, errOut, body, body, ds)
	}
	if command == "context" {
		body, err := workflow.Context(root, args[0], full)
		if err != nil {
			return cliResponse(1, jsonMode, out, errOut, nil, "", lifecycleError(err, nil, "Use a change held by this worktree and retry context."))
		}
		return cliResponse(0, jsonMode, out, errOut, body, body, nil)
	}
	if command == "release" {
		if err := workflow.Release(root, args[0]); err != nil {
			return cliResponse(1, jsonMode, out, errOut, nil, "", lifecycleError(err, nil, "Use a change claimed by this worktree and retry release."))
		}
		body := fmt.Sprintf("Released %s\n", args[0])
		return cliResponse(0, jsonMode, out, errOut, body, body, nil)
	}
	if command == "commits" {
		ids, err := workflow.Commits(root, args[0])
		if err != nil {
			ds := lifecycleError(err, nil, "Use a valid pending or archived change selector in a Git project, then retry commits.")
			status := 1
			if ds[0].Code == "U02" {
				status = 2
			}
			return cliResponse(status, jsonMode, out, errOut, nil, "", ds)
		}
		if ids == nil {
			ids = []string{}
		}
		plain := ""
		if len(ids) > 0 {
			plain = strings.Join(ids, "\n") + "\n"
		}
		return cliResponse(0, jsonMode, out, errOut, ids, plain, nil)
	}
	if command == "abandon" {
		oid, err := workflow.Abandon(root, args[0])
		if err != nil {
			ds := lifecycleError(err, nil, "Use an eligible pending change in a clean Git project, resolve any conflicting claim, then retry abandon.")
			status := 1
			if ds[0].Code == "U02" {
				status = 2
			}
			return cliResponse(status, jsonMode, out, errOut, nil, "", ds)
		}
		body := fmt.Sprintf("Abandoned %s in %s\n", args[0], oid)
		return cliResponse(0, jsonMode, out, errOut, body, body, nil)
	}
	bin, err := os.Executable()
	if err != nil {
		return fail(2, "P01", fmt.Sprintf("Find invoking executable: %v. Run the installed fireweed executable.", err))
	}
	if command == "coverage" {
		if withChange != "" {
			if _, statErr := os.Stat(filepath.Join(root, ".git")); statErr == nil {
				if state, currentErr := workflow.Current(root, withChange); currentErr == nil {
					snapshot = state.Snapshot
				} else {
					snapshot, err = project.Project(root, withChange)
				}
			} else {
				snapshot, err = project.Project(root, withChange)
			}
			if err != nil {
				return fail(2, projectFailureCode(err, "P01"), fmt.Sprintf("%v. Use a valid queued change alias.", err))
			}
		}
		var module *project.Module
		for i := range snapshot.Modules {
			if snapshot.Modules[i].Name == args[0] {
				module = &snapshot.Modules[i]
				break
			}
		}
		if module == nil {
			return fail(2, "U02", fmt.Sprintf("Unknown module %q. Use modules to choose a valid module.", args[0]))
		}
		if withChange == "" {
			if _, statErr := os.Stat(filepath.Join(root, ".git")); statErr == nil {
				if state, currentErr := workflow.Current(root, ""); currentErr == nil {
					snapshot.Config = state.Snapshot.Config
					for _, approved := range state.Snapshot.Modules {
						if approved.Name == module.Name {
							currentBody := module.Document.Body
							*module = approved
							module.Document.Body = currentBody
							break
						}
					}
				}
			}
		}
		section, _ := contracts.Section(module.Document.Body, "Specs")
		specs, err := contracts.ParseItems(section)
		if err != nil {
			code := "C03"
			if assigned, _, ok := contracts.ErrorCode(err); ok {
				code = assigned
			}
			return fail(2, code, fmt.Sprintf("%v. Correct the module Specs and retry coverage.", err))
		}
		report, err := checks.Evaluate(context.Background(), checks.HookRequest{Snapshot: snapshot, Module: module.Name, Bin: bin, Timeout: hookTimeout(snapshot)})
		if err == nil {
			err = checks.RequireCoverage(report, specs)
		}
		if err == nil {
			err = checks.RequirePass(report, specs)
		}
		if err != nil {
			return cliResponse(1, jsonMode, out, errOut, nil, "", lifecycleError(err, nil, "Use a valid eval hook with covered, passing spec cases, then retry coverage."))
		}
		var plain strings.Builder
		for _, c := range report.Cases {
			fmt.Fprintf(&plain, "%s: %s\n", strings.Join(c.IDs, ", "), c.Result)
		}
		return cliResponse(0, jsonMode, out, errOut, report, plain.String(), nil)
	}
	change := ""
	if len(args) == 1 {
		change = args[0]
	}
	request := workflow.Request{Root: root, Change: change, Bin: bin, Quick: quick, Timeout: hookTimeout(snapshot)}
	var result workflow.Result
	if command == "done" {
		result, err = workflow.Done(context.Background(), request)
	} else {
		result, err = workflow.Verify(context.Background(), request)
	}
	ds := reported(result.Diagnostics)
	if err != nil {
		return cliResponse(1, jsonMode, out, errOut, nil, "", lifecycleError(err, ds, "Use the held change and correct its checks, then retry "+command+"."))
	}
	return cliResponse(0, jsonMode, out, errOut, result, fmt.Sprintf("%s: %s\n", command, result.State.Phase), ds)
}
