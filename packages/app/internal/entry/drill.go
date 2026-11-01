package entry

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/lurkshark/fireweed/packages/runner"
)

type drillData struct {
	Module     string `json:"module"`
	Result     string `json:"result"`
	Report     string `json:"report"`
	Ref        string `json:"ref"`
	BaseCommit string `json:"base_commit"`
	Reason     string `json:"reason,omitempty"`
}

func runDrill(args []string, jsonMode bool, out, errOut io.Writer) int {
	if len(args) != 1 {
		return cliFailure(2, jsonMode, out, errOut, "U01", "drill requires exactly one module. Use drill <module> [--json].")
	}
	snapshot, err := loadValidatedProject()
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, projectFailureCode(err, "C01"), fmt.Sprintf("%v. Provide a valid fireweed.yaml and module declarations, then retry drill.", err))
	}
	found := false
	for _, module := range snapshot.Modules {
		if module.Name == args[0] {
			found = true
			break
		}
	}
	if !found {
		return cliFailure(2, jsonMode, out, errOut, "U02", fmt.Sprintf("Unknown module %q. Use modules to choose a valid module, then retry drill.", args[0]))
	}
	bin, err := os.Executable()
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, "P01", fmt.Sprintf("Find invoking executable: %v. Run the installed fireweed executable.", err))
	}
	result, err := runner.Drill(context.Background(), runner.DrillRequest{
		Root: snapshot.Root, Module: args[0], Bin: bin,
	})
	data := drillData{Module: result.Module, Result: result.Result, Report: result.ReportPath, Ref: result.Ref, BaseCommit: result.Base}
	if err != nil && runner.IsUsageStop(err) && result.Result == "stopped" {
		data.Reason = err.Error()
		plain := fmt.Sprintf("Drill %s stopped: %s", data.Module, data.Reason)
		if data.Report != "" {
			plain += ". Report: " + data.Report
		}
		if data.Ref != "" {
			plain += ". Ref: " + data.Ref
		}
		return cliResponse(3, jsonMode, out, errOut, data, plain+"\n", nil)
	}
	if err != nil {
		code := "P01"
		for _, match := range wrappedCause.FindAllStringSubmatch(err.Error(), -1) {
			code = match[1]
		}
		var artifacts []string
		if data.Report != "" {
			if _, statErr := os.Stat(filepath.Join(snapshot.Root, filepath.FromSlash(data.Report))); statErr == nil {
				artifacts = append(artifacts, "report "+data.Report)
			}
		}
		if data.Ref != "" {
			artifacts = append(artifacts, "rebuilt ref "+data.Ref)
		}
		message := err.Error()
		if len(artifacts) > 0 {
			message += ". Retained " + strings.Join(artifacts, " and ")
		}
		message += ". Use the reported cause to correct the drill inputs or checks, then retry drill."
		return cliFailure(1, jsonMode, out, errOut, code, message)
	}
	plain := fmt.Sprintf("Drill %s: %s\nReport: %s\nRef: %s\n", data.Module, data.Result, data.Report, data.Ref)
	return cliResponse(0, jsonMode, out, errOut, data, plain, nil)
}
