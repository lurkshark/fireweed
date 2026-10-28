package entry

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/project"
)

type doctorCheckResult struct {
	Name   string `json:"name"`
	Module string `json:"module"`
	Passed bool   `json:"passed"`
	Detail string `json:"detail"`
}

func runDoctor(noAgent, jsonMode bool, out, errOut io.Writer) int {
	root, err := findProjectRoot()
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, "C01", fmt.Sprintf("%v. Provide a valid fireweed.yaml and retry doctor.", err))
	}
	// Configuration loading is a usage failure. Other static failures belong to
	// Doctor's diagnostic report and retain their assigned causes.
	if _, err := project.Load(root); err != nil && projectFailureCode(err, "C01") == "C01" {
		return cliFailure(2, jsonMode, out, errOut, "C01", fmt.Sprintf("%v. Correct fireweed.yaml and retry doctor.", err))
	}
	bin, err := os.Executable()
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, "P01", fmt.Sprintf("Find invoking executable: %v. Run the installed fireweed executable.", err))
	}
	report := checks.Doctor(context.Background(), checks.DoctorRequest{Root: root, Bin: bin, NoAgent: noAgent})
	checksOut := make([]doctorCheckResult, 0, len(report.Checks))
	for _, check := range report.Checks {
		name := strings.Fields(check.Name)
		if len(name) == 0 {
			name = []string{check.Name}
		}
		checksOut = append(checksOut, doctorCheckResult{Name: name[0], Module: check.Module, Passed: check.Passed, Detail: check.Name + ": " + check.Detail})
	}
	diagnostics := reported(report.Diagnostics)
	status := 0
	for _, d := range diagnostics {
		if !d.Warning {
			status = 1
		}
	}
	if jsonMode {
		var data any
		if status == 0 {
			data = struct {
				Checks []doctorCheckResult `json:"checks"`
			}{Checks: checksOut}
		}
		_ = json.NewEncoder(out).Encode(envelope{Schema: 1, Data: data, Diagnostics: diagnostics})
		return status
	}
	for _, check := range checksOut {
		outcome := "fail"
		if check.Passed {
			outcome = "pass"
		}
		name := check.Name
		if check.Module != "" {
			name += " (" + check.Module + ")"
		}
		fmt.Fprintf(out, "%s: %s", name, outcome)
		if check.Detail != "" {
			fmt.Fprintf(out, ": %s", check.Detail)
		}
		fmt.Fprintln(out)
	}
	for _, d := range diagnostics {
		kind := "error"
		if d.Warning {
			kind = "warning"
		}
		path := ""
		if d.Path != "" {
			path = " (" + d.Path + ")"
		}
		fmt.Fprintf(errOut, "%s %s%s: %s\n", d.Code, kind, path, d.Message)
	}
	return status
}

func runAgentsMD(jsonMode bool, out, errOut io.Writer) int {
	root, err := findProjectRoot()
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, "C01", fmt.Sprintf("%v. Provide a valid fireweed.yaml and retry agents-md --write.", err))
	}
	if _, err := project.Load(root); err != nil {
		code := projectFailureCode(err, "C01")
		return cliFailure(2, jsonMode, out, errOut, code, fmt.Sprintf("%v. Correct the project configuration and retry agents-md --write.", err))
	}
	data, err := updateManagedFile(root)
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, "P01", fmt.Sprintf("Cannot update AGENTS.md: %v. Use paired managed-block markers and a readable file, then retry.", err))
	}
	if err := os.WriteFile(filepath.Join(root, "AGENTS.md"), data, 0644); err != nil {
		return cliFailure(2, jsonMode, out, errOut, "P01", fmt.Sprintf("Cannot write AGENTS.md: %v. Use a writable file and retry.", err))
	}
	confirmation := "Updated AGENTS.md"
	return cliResponse(0, jsonMode, out, errOut, confirmation, strings.TrimSpace(confirmation)+"\n", nil)
}
