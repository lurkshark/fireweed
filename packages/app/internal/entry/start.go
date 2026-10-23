package entry

import (
	"io"
	"strings"

	"github.com/lurkshark/fireweed/packages/workflow"
)

func start(out, errOut io.Writer, option options) int {
	result, err := workflow.Start(".", option.operands[0])
	if err != nil {
		code, _, _ := strings.Cut(err.Error(), ":")
		exit := 2
		switch code {
		case "P01", "P02", "P03", "C05":
			exit = 1
		case "C01", "D05":
		default:
			code = "C01"
		}
		return emit(out, errOut, option, nil, []diagnostic{{Code: code, Severity: "error", Message: err.Error()}}, exit)
	}
	diagnostics := make([]diagnostic, 0, len(result.Diagnostics))
	for _, d := range result.Diagnostics {
		diagnostics = append(diagnostics, diagnostic{Code: d.Code, Severity: d.Severity, Message: d.Message, Path: d.Path})
	}
	return emit(out, errOut, option, result.Context, diagnostics, 0)
}
