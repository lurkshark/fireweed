package entry

import (
	"fmt"
	"io"

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
		status, fallback := 1, "P01"
		action := "Use a resolvable commit range, then retry pr-body."
		if workflow.IsEndpointConfigError(err) {
			status, fallback = 2, "C01"
			action = "Correct fireweed.yaml and module declarations at the selected endpoint, then retry pr-body."
		}
		diagnostics = appendErrorDiagnostic(diagnostics, err, fallback, action)
		return cliResponse(status, jsonMode, out, errOut, nil, "", diagnostics)
	}
	return cliResponse(0, jsonMode, out, errOut, report.Body, report.Body, diagnostics)
}
