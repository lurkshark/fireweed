package entry

import (
	"fmt"
	"io"
	"strings"

	"github.com/lurkshark/fireweed/packages/workflow"
)

func runPRBody(args []string, jsonMode bool, out, errOut io.Writer) int {
	if len(args) != 1 {
		return cliFailure(2, jsonMode, out, errOut, "U01", "pr-body requires exactly one range. Use pr-body <range> [--json].")
	}
	root, err := findProjectRoot()
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, "C01", fmt.Sprintf("%v. Use pr-body <range> from a project containing fireweed.yaml.", err))
	}
	report, err := workflow.PRBody(root, args[0])
	diagnostics := reported(report.Diagnostics)
	if err != nil {
		endpointConfig := strings.HasPrefix(err.Error(), "load endpoint project at ")
		action := "Use a resolvable commit range, then retry pr-body."
		if endpointConfig {
			action = "Correct fireweed.yaml and module declarations at the selected endpoint, then retry pr-body."
		}
		diagnostics = lifecycleError(err, diagnostics, action)
		status := 1
		if endpointConfig {
			status = 2
			last := len(diagnostics) - 1
			if diagnostics[last].Code == "P01" {
				diagnostics[last].Code = "C01"
				for _, part := range strings.SplitAfter(err.Error(), ": ") {
					if assigned := diagnosticPrefix.FindStringSubmatch(part); assigned != nil {
						diagnostics[last].Code = assigned[1]
						break
					}
				}
			}
		}
		return cliResponse(status, jsonMode, out, errOut, nil, "", diagnostics)
	}
	return cliResponse(0, jsonMode, out, errOut, report.Body, report.Body, diagnostics)
}
