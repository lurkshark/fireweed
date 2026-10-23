package entry

import (
	"io"
	"strings"

	"github.com/lurkshark/fireweed/packages/workflow"
)

func workerContext(out, errOut io.Writer, option options) int {
	slug := option.operands[0]
	source, _, err := workflow.LoadApproval(".", slug)
	if err == nil {
		var result workflow.WorkerContext
		result, err = workflow.BuildContext(source, slug)
		if err == nil {
			return emit(out, errOut, option, result, nil, 0)
		}
	}
	// The public source and context contracts carry their stable code in the
	// error prefix. Keep the complete detail; paths are not parsed from prose.
	code, _, _ := strings.Cut(err.Error(), ":")
	exit := 2
	switch code {
	case "C05", "P02":
		exit = 1
	case "C01", "D05":
	default:
		code = "C01"
	}
	return emit(out, errOut, option, nil, []diagnostic{{Code: code, Severity: "error", Message: err.Error()}}, exit)
}
