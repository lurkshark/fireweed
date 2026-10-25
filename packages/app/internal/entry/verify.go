package entry

import (
	"context"
	"io"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/workflow"
)

type verificationEdit struct {
	Status  string `json:"status"`
	Path    string `json:"path"`
	OldPath string `json:"old_path,omitempty"`
}

type verificationData struct {
	Change      string              `json:"change"`
	Phase       string              `json:"phase"`
	Diagnostics []checks.Diagnostic `json:"diagnostics"`
	Coverage    []checks.Coverage   `json:"coverage"`
	Edits       []verificationEdit  `json:"edits"`
	Findings    []workflow.Finding  `json:"findings,omitempty"`
}

func verify(out, errOut io.Writer, option options) int {
	slug := ""
	if len(option.operands) == 1 {
		slug = option.operands[0]
	}
	result, err := workflow.Verify(context.Background(), ".", slug)
	if err != nil {
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
	data := verificationData{
		Change: result.Change, Phase: result.Phase,
		Diagnostics: result.Diagnostics, Coverage: result.Coverage,
		Edits: make([]verificationEdit, 0, len(result.Edits)), Findings: result.Findings,
	}
	for _, edit := range result.Edits {
		data.Edits = append(data.Edits, verificationEdit{edit.Status, edit.Path, edit.OldPath})
	}
	diagnostics := make([]diagnostic, 0, len(result.Diagnostics))
	exit := 0
	for _, d := range result.Diagnostics {
		diagnostics = append(diagnostics, diagnostic{Code: d.Code, Severity: d.Severity, Message: d.Message, Path: d.Path})
		if d.Severity == "error" {
			exit = 1
		}
	}
	return emit(out, errOut, option, data, diagnostics, exit)
}
