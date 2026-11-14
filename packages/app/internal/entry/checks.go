package entry

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func (c command) checkCommand(root string) response {
	// Static checks and doctor report their own module and queue errors.
	snapshot, err := project.Load(root)
	if err != nil && (cause(err, "") == "C01" || c.name == "check-commits") {
		return configFailure(err)
	}
	switch c.name {
	case "check":
		if c.has("--fix") {
			if err := stampQueue(root); err != nil {
				return configuredOperationFailure(err, "Correct the reported issue and use fireweed check --fix again.")
			}
		}
		return checked(verbatim("Check passed.\n"), project.Check(root), "Correct the reported issue and run fireweed check again.")
	case "check-commits":
		return history(root, c.arg(0))
	case "doctor":
		return doctor(root, c.has("--no-agent"))
	case "ci":
		return ci(root, c.arg(0), c.has("--no-doctor"), snapshot)
	}
	panic("unexpected check command: " + c.name)
}

func stampQueue(root string) error {
	queue, err := project.Queue(root)
	if err != nil {
		return err
	}
	for _, change := range queue {
		if err := project.StampMissing(root, change.Path); err != nil {
			return err
		}
	}
	return nil
}

type squashRow struct {
	Commit string `json:"commit"`
	Change string `json:"change"`
}

func history(root, revisionRange string) response {
	report, err := checks.History(root, revisionRange)
	if err != nil {
		return operationFailure(err, "Use a valid local commit range and correct any malformed project files at its endpoints.")
	}
	var plain strings.Builder
	fmt.Fprintf(&plain, "Checked %d commits.\n", len(report.Commits))
	rows := make([]squashRow, 0, len(report.Squashed))
	for _, entry := range report.Squashed {
		rows = append(rows, squashRow(entry))
		fmt.Fprintf(&plain, "%s %s\n", entry.Commit, entry.Change)
	}
	return checked(output(plain.String(), struct {
		Commits  []string    `json:"commits"`
		Squashed []squashRow `json:"squashed"`
	}{array(report.Commits), rows}), report.Diagnostics, "Correct the commit to satisfy the reported history rule.")
}

type doctorRow struct {
	Name   string `json:"name"`
	Module string `json:"module"`
	Passed bool   `json:"passed"`
	Detail string `json:"detail"`
}

func doctor(root string, noAgent bool) response {
	report := checks.Doctor(background(), checks.DoctorRequest{Root: root, Bin: invokingBin(), NoAgent: noAgent})
	rows := make([]doctorRow, 0, len(report.Checks))
	var plain strings.Builder
	for _, check := range report.Checks {
		row := doctorRow(check)
		name, scope, scoped := strings.Cut(check.Name, " ")
		switch name {
		case "baseline", "undeclared", "internal_eval", "failing_eval", "missing_id", "recovery":
			if scoped {
				row.Name, row.Detail = name, scope+": "+check.Detail
			}
		}
		rows = append(rows, row)
		outcome := "FAIL"
		if check.Passed {
			outcome = "PASS"
		}
		fmt.Fprintf(&plain, "%s %s %s %s\n", outcome, check.Name, check.Module, check.Detail)
	}
	return checked(output(plain.String(), struct {
		Checks []doctorRow `json:"checks"`
	}{rows}), report.Diagnostics, "")
}

func ci(root, revisionRange string, noDoctor bool, snapshot project.Snapshot) response {
	const action = "Correct the reported issue and run fireweed ci again."
	diagnostics := project.Check(root)
	finish := func() response {
		confirmation := response{}
		if !hasErrors(diagnostics) {
			confirmation = message("CI passed.")
		}
		return checked(confirmation, diagnostics, action)
	}
	if hasErrors(diagnostics) {
		return finish()
	}
	add := func(err error) {
		_, subject, _ := contracts.ErrorCode(err)
		diagnostics = append(diagnostics, project.Diagnostic{Code: cause(err, "P01"), Path: subject, Message: err.Error()})
	}

	report, err := checks.History(root, revisionRange)
	if err != nil {
		add(err)
		return finish()
	}
	diagnostics = append(diagnostics, report.Diagnostics...)
	if hasErrors(diagnostics) {
		return finish()
	}

	bin := invokingBin()
	timeout, err := hookTimeout(snapshot)
	if err != nil {
		return configFailure(err)
	}
	for _, module := range snapshot.Modules {
		items, err := specItems(module)
		if err == nil {
			_, err = evaluate(snapshot, module.Name, bin, timeout, items)
		}
		if err != nil {
			add(err)
		}
	}
	boundaries, err := checks.Boundaries(background(), checks.BoundaryRequest{Snapshot: snapshot, Bin: bin, Timeout: timeout})
	diagnostics = append(diagnostics, boundaries...)
	if err != nil && !hasErrors(boundaries) {
		add(err)
	}
	if !noDoctor {
		report := checks.Doctor(background(), checks.DoctorRequest{Root: root, Bin: bin, NoAgent: true})
		for _, d := range report.Diagnostics {
			found := false
			for _, existing := range diagnostics {
				if existing == d {
					found = true
					break
				}
			}
			if !found {
				diagnostics = append(diagnostics, d)
			}
		}
	}
	return finish()
}

func hookTimeout(snapshot project.Snapshot) (time.Duration, error) {
	if raw, exists := snapshot.Config["timeout"]; exists {
		value, _ := raw.(string)
		duration, err := time.ParseDuration(value)
		if err != nil || duration <= 0 {
			return 0, coded("C01", "fireweed.yaml", fmt.Errorf("timeout must be a positive Go duration, got %v", raw))
		}
		return duration, nil
	}
	return 90 * time.Minute, nil
}

func evaluate(snapshot project.Snapshot, module, bin string, timeout time.Duration, items []contracts.Item) (checks.Report, error) {
	report, err := checks.Evaluate(background(), checks.HookRequest{Snapshot: snapshot, Module: module, Bin: bin, Timeout: timeout})
	if err == nil {
		err = checks.RequireCoverage(report, items)
	}
	if err == nil {
		err = checks.RequirePass(report, items)
	}
	return report, err
}

func (c command) coverage(root string, approved *workflow.ApprovedSnapshot) response {
	name, selector := c.arg(0), c.options["--with-change"]
	snapshot, items, err := coverageInput(root, name, selector, approved)
	if err != nil {
		if cause(err, "") == "C01" {
			return configFailure(err)
		}
		return failed(2, err, "C01", "Correct the project inputs or use fireweed modules and fireweed list to check the selection.")
	}
	timeout, err := hookTimeout(snapshot)
	if err != nil {
		return configFailure(err)
	}
	report, err := evaluate(snapshot, name, invokingBin(), timeout, items)
	if err != nil {
		return operationFailure(err, "Correct the eval hook or testcases and run fireweed coverage again.")
	}
	var plain strings.Builder
	for _, item := range items {
		outcome := "uncovered"
		for _, testcase := range report.Cases {
			for _, id := range testcase.IDs {
				if id == item.ID {
					switch {
					case testcase.Result == "fail":
						outcome = "fail"
					case testcase.Result == "pass" && outcome != "fail":
						outcome = "pass"
					case testcase.Result == "skip" && outcome == "uncovered":
						outcome = "skip"
					}
				}
			}
		}
		fmt.Fprintf(&plain, "%s: %s\n", item.ID, outcome)
	}
	return output(plain.String(), report)
}

func coverageInput(root, name, selector string, approved *workflow.ApprovedSnapshot) (project.Snapshot, []contracts.Item, error) {
	unknown := func() error {
		return coded("U02", name, fmt.Errorf("Unknown module %q", name))
	}
	var snapshot project.Snapshot
	var selected project.Module
	var found bool
	if approved == nil {
		var err error
		if selector != "" {
			snapshot, err = project.Project(root, selector)
		} else {
			snapshot, err = loadGraph(root)
		}
		if err != nil {
			return snapshot, nil, err
		}
		selected, found = moduleByName(snapshot, name)
		if !found {
			return snapshot, nil, unknown()
		}
	} else {
		snapshot = approved.Snapshot
		if err := validateGraph(snapshot.Modules); err != nil {
			return snapshot, nil, err
		}
		selected, found = moduleByName(snapshot, name)
		if !found {
			return snapshot, nil, unknown()
		}
		switch selector {
		case "":
			file := path.Join(selected.Dir, "SPEC.md")
			body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(file)))
			if err != nil {
				return snapshot, nil, coded("C03", file, fmt.Errorf("%s: %w", file, err))
			}
			doc, err := contracts.ParseDocument(string(body))
			if err != nil {
				return snapshot, nil, coded(cause(err, "C03"), file, fmt.Errorf("%s: %w", file, err))
			}
			selected.Document.Body = doc.Body
		case approved.Slug, path.Base(approved.Path), approved.Path:
			// The approved document supplies both the spec text and hook metadata.
		default:
			live, err := project.Project(root, selector)
			if err != nil {
				return snapshot, nil, err
			}
			module, found := moduleByName(live, name)
			if !found {
				return snapshot, nil, unknown()
			}
			selected.Document.Body = module.Document.Body
		}
		// Preserve the held graph and hook configuration and replace only the
		// selected module's spec text for JUnit mapping.
		snapshot.Modules = append([]project.Module(nil), snapshot.Modules...)
		for i := range snapshot.Modules {
			if snapshot.Modules[i].Name == name {
				snapshot.Modules[i] = selected
				break
			}
		}
	}
	items, err := specItems(selected)
	return snapshot, items, err
}
