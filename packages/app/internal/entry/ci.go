package entry

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func runCI(args []string, noDoctor, jsonMode bool, out, errOut io.Writer) int {
	if len(args) > 1 {
		return cliFailure(2, jsonMode, out, errOut, "U01", "Use ci [range] [--no-doctor] with at most one revision range.")
	}
	snapshot, err := loadValidatedProject()
	if err != nil {
		code := projectFailureCode(err, "C01")
		status := 1
		if code == "C01" || code == "C02" || code == "C12" {
			status = 2
		}
		return cliFailure(status, jsonMode, out, errOut, code, fmt.Sprintf("%v. Correct fireweed.yaml and module declarations, then retry ci.", err))
	}
	bin, err := os.Executable()
	if err != nil {
		return cliFailure(2, jsonMode, out, errOut, "P01", fmt.Sprintf("Find invoking executable: %v. Run the installed fireweed executable.", err))
	}
	ctx := context.Background()
	var diagnostics []diagnostic
	add := func(items []project.Diagnostic) {
		for _, item := range reported(items) {
			if !item.Warning {
				item.Message += ". Use explain " + item.Code + " for corrective guidance."
			}
			diagnostics = append(diagnostics, item)
		}
	}
	addError := func(err error, action string) {
		diagnostics = lifecycleError(err, diagnostics, action)
	}
	add(project.Check(snapshot.Root))

	revisionRange := ""
	if len(args) == 1 {
		revisionRange = args[0]
	}
	history, err := checks.History(snapshot.Root, revisionRange)
	if err != nil {
		code := projectFailureCode(err, "P01")
		if strings.HasPrefix(err.Error(), "parse ") && strings.Contains(err.Error(), "/SPEC.md:") {
			code = "C03"
		}
		addError(fmt.Errorf("%s: %w", code, err), "Use a valid Git revision range and repair recorded documents, then retry ci.")
	} else {
		add(history.Diagnostics)
	}

	for _, module := range snapshot.Modules {
		report, evalErr := checks.Evaluate(ctx, checks.HookRequest{
			Snapshot: snapshot, Module: module.Name, Bin: bin, Timeout: hookTimeout(snapshot),
		})
		if evalErr != nil {
			addError(evalErr, "Use valid Specs and an eval hook for "+module.Name+", then retry ci.")
			continue
		}
		section, _ := contracts.Section(module.Document.Body, "Specs")
		specs, parseErr := contracts.ParseItems(section)
		if parseErr != nil {
			code := "C03"
			if strings.HasPrefix(parseErr.Error(), "duplicate item ID:") {
				code = "C05"
			}
			addError(fmt.Errorf("%s: %w", code, parseErr), "Use valid "+module.Name+" Specs, then retry ci.")
			continue
		}
		if coverageErr := checks.RequireCoverage(report, specs); coverageErr != nil {
			addError(coverageErr, "Use a passing case for each current "+module.Name+" spec, then retry ci.")
		}
		if passErr := checks.RequirePass(report, specs); passErr != nil {
			addError(passErr, "Use corrected "+module.Name+" cases, then retry ci.")
		}
	}
	boundaryDiagnostics, boundaryErr := checks.Boundaries(ctx, checks.BoundaryRequest{
		Snapshot: snapshot, Bin: bin, Timeout: hookTimeout(snapshot),
	})
	add(boundaryDiagnostics)
	if boundaryErr != nil && !hasError(reported(boundaryDiagnostics)) {
		addError(boundaryErr, "Use a valid boundaries hook, then retry ci.")
	}
	if !noDoctor {
		doctor := checks.Doctor(ctx, checks.DoctorRequest{
			Root: snapshot.Root, Bin: bin, NoAgent: true, Timeout: hookTimeout(snapshot),
		})
		add(doctor.Diagnostics)
	}
	if hasError(diagnostics) {
		return cliResponse(1, jsonMode, out, errOut, nil, "", diagnostics)
	}
	return cliResponse(0, jsonMode, out, errOut, "CI passed", "CI passed\n", diagnostics)
}

func hasError(diagnostics []diagnostic) bool {
	for _, item := range diagnostics {
		if !item.Warning {
			return true
		}
	}
	return false
}
