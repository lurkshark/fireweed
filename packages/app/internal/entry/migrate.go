package entry

import (
	"fmt"
	"io"

	"github.com/lurkshark/fireweed/packages/project"
)

const migrateConfirmation = "Project format 1 is already current.\n"

func runMigrate(args []string, jsonMode bool, out, errOut io.Writer) int {
	if len(args) != 0 {
		return cliFailure(2, jsonMode, out, errOut, "U01", "migrate takes no arguments. Use migrate [--json].")
	}
	root, err := findProjectRoot()
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, "C01", fmt.Sprintf("%v. Use migrate from a project containing fireweed.yaml.", err))
	}
	snapshot, err := project.Load(root)
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, errorCode(err, "C01"), fmt.Sprintf("%v. Correct fireweed.yaml and module declarations, then retry migrate.", err))
	}
	if version, ok := snapshot.Config["version"].(int); !ok || version != 1 {
		return cliFailure(2, jsonMode, out, errOut, "C01", fmt.Sprintf("Unsupported project format version %v in fireweed.yaml. Use a project with explicit version: 1, then retry migrate.", snapshot.Config["version"]))
	}
	diagnostics := reported(project.Check(snapshot.Root))
	status := 0
	for i := range diagnostics {
		if !diagnostics[i].Warning {
			status = 1
			diagnostics[i].Message += ". Correct the reported project document, then retry migrate."
		}
	}
	if status != 0 {
		return cliResponse(status, jsonMode, out, errOut, nil, "", diagnostics)
	}
	return cliResponse(0, jsonMode, out, errOut, migrateConfirmation, migrateConfirmation, diagnostics)
}
