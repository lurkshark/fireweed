package entry

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/lurkshark/fireweed/packages/runner"
)

type runSummary struct {
	Change string `json:"change"`
	Run    string `json:"run"`
	Phase  string `json:"phase"`
}

type runData struct {
	Applied []runSummary `json:"applied"`
	Empty   bool         `json:"empty"`
	Limited bool         `json:"limited"`
}

func runQueue(max int, allowUnmerged, jsonMode bool, out, errOut io.Writer) int {
	snapshot, err := loadValidatedProject()
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, projectFailureCode(err, "C01"), fmt.Sprintf("%v. Provide a valid fireweed.yaml and module declarations, then retry run.", err))
	}
	bin, err := os.Executable()
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, "P01", fmt.Sprintf("Find invoking executable: %v. Run the installed fireweed executable.", err))
	}
	result, err := runner.Run(context.Background(), runner.RunRequest{
		Root: snapshot.Root, Bin: bin, Max: max, AllowUnmerged: allowUnmerged,
	})
	diagnostics := make([]diagnostic, 0)
	for _, applied := range result.Applied {
		diagnostics = append(diagnostics, reported(applied.State.Diagnostics)...)
	}
	if err != nil {
		code := "P01"
		for _, match := range wrappedCause.FindAllStringSubmatch(err.Error(), -1) {
			code = match[1]
		}
		action := "Use the reported cause to correct the change or project checks, then retry run."
		if strings.Contains(err.Error(), "requires approval") {
			action = "Approve the change on the configured base or use --allow-unmerged for authorized local work, then retry run."
		}
		message := fmt.Sprintf("%v. %s", err, action)
		diagnostics = append([]diagnostic{{Code: code, Message: message}}, diagnostics...)
		return cliResponse(1, jsonMode, out, errOut, nil, "", diagnostics)
	}
	data := runData{Applied: make([]runSummary, 0, len(result.Applied)), Empty: result.Empty, Limited: result.Limited}
	var plain strings.Builder
	for _, applied := range result.Applied {
		state := applied.State
		data.Applied = append(data.Applied, runSummary{Change: state.Change.Path, Run: state.Run, Phase: state.Phase})
		fmt.Fprintf(&plain, "Applied %s\n", state.Change.Path)
	}
	fmt.Fprintf(&plain, "Applied %d changes.\n", len(data.Applied))
	if data.Limited {
		plain.WriteString("Max limit reached; changes remain queued.\n")
	} else if data.Empty {
		plain.WriteString("Queue empty.\n")
	}
	return cliResponse(0, jsonMode, out, errOut, data, plain.String(), diagnostics)
}
