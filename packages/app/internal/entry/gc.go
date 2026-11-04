package entry

import (
	"context"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/lurkshark/fireweed/packages/workflow"
)

type gcData struct {
	Pruned []string `json:"pruned"`
}

func runGC(args []string, jsonMode bool, out, errOut io.Writer) int {
	if len(args) != 0 {
		return cliFailure(2, jsonMode, out, errOut, "U01", "gc takes no arguments. Use gc [--json].")
	}
	snapshot, err := loadValidatedProject()
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, projectFailureCode(err, "C01"), fmt.Sprintf("%v. Correct fireweed.yaml and module declarations, then retry gc.", err))
	}
	result, err := workflow.GC(context.Background(), workflow.GCRequest{Root: snapshot.Root})
	if err != nil {
		diagnostics := lifecycleError(err, nil, "Correct the reported project configuration or Git state, then retry gc.")
		status := 1
		if diagnostics[0].Code == "C01" {
			status = 2
		}
		return cliResponse(status, jsonMode, out, errOut, nil, "", diagnostics)
	}
	pruned := nonNil(result.Pruned)
	sort.Strings(pruned)
	plain := "No refs pruned.\n"
	if len(pruned) != 0 {
		plain = strings.Join(pruned, "\n") + "\n"
	}
	return cliResponse(0, jsonMode, out, errOut, gcData{Pruned: pruned}, plain, nil)
}
