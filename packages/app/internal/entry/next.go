package entry

import (
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/lurkshark/fireweed/packages/runner"
)

type nextData struct {
	Change string `json:"change"`
	Run    string `json:"run"`
	Phase  string `json:"phase"`
	Empty  bool   `json:"empty"`
	Reason string `json:"reason,omitempty"`
}

var wrappedCause = regexp.MustCompile(`(?:^|: )([A-Z][A-Z0-9-]+):\s*`)

func runNext(allowUnmerged, jsonMode bool, out, errOut io.Writer) int {
	snapshot, err := loadValidatedProject()
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, projectFailureCode(err, "C01"), fmt.Sprintf("%v. Provide a valid fireweed.yaml and module declarations, then retry next.", err))
	}
	bin, err := os.Executable()
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, "P01", fmt.Sprintf("Find invoking executable: %v. Run the installed fireweed executable.", err))
	}
	result, err := runner.Next(context.Background(), runner.NextRequest{
		Root: snapshot.Root, Bin: bin, AllowUnmerged: allowUnmerged,
	})
	if runner.IsUsageStop(err) {
		data := nextData{Change: result.State.Change.Path, Run: result.State.Run, Phase: result.State.Phase, Empty: result.Empty, Reason: err.Error()}
		plain := fmt.Sprintf("Held %s. %s. Run next to resume this change.\n", data.Change, data.Reason)
		return cliResponse(3, jsonMode, out, errOut, data, plain, reported(result.State.Diagnostics))
	}
	if err != nil {
		code := "P01"
		for _, match := range wrappedCause.FindAllStringSubmatch(err.Error(), -1) {
			code = match[1]
		}
		action := "Use the reported cause to correct the change or project checks, then retry next."
		if strings.Contains(err.Error(), "requires approval") {
			action = "Approve the change on the configured base or use --allow-unmerged for authorized local work, then retry next."
		}
		status := 1
		if code == "C01" {
			status = 2
			action = "Use a valid usage policy in fireweed.yaml, then retry next."
		}
		message := fmt.Sprintf("%v. %s", err, action)
		diagnostics := append([]diagnostic{{Code: code, Message: message}}, reported(result.State.Diagnostics)...)
		return cliResponse(status, jsonMode, out, errOut, nil, "", diagnostics)
	}
	data := nextData{Change: result.State.Change.Path, Run: result.State.Run, Phase: result.State.Phase, Empty: result.Empty}
	plain := fmt.Sprintf("Applied %s\n", data.Change)
	if data.Empty {
		plain = "Queue empty\n"
	}
	return cliResponse(0, jsonMode, out, errOut, data, plain, reported(result.State.Diagnostics))
}
