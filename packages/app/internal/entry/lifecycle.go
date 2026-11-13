package entry

import (
	"fmt"
	"slices"
	"strings"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func (c command) workflowCommand(root string) response {
	selector := c.arg(0)
	if c.name == "status" {
		return status(root)
	}
	if c.name == "pr-body" {
		report, err := workflow.PRBody(root, selector)
		if err != nil {
			if workflow.IsEndpointConfigError(err) {
				return configFailure(err)
			}
			return operationFailure(err, "Use a range whose endpoints resolve to local commits.")
		}
		result := verbatim(report.Body)
		result.diagnostics = report.Diagnostics
		return result
	}
	if _, err := project.Load(root); err != nil && cause(err, "") == "C01" {
		return configFailure(err)
	}
	action := "Use fireweed list and fireweed context to check the change and correct its state."
	switch c.name {
	case "start":
		var state workflow.State
		var err error
		if c.attempt > 0 {
			state, err = workflow.StartFromAttempt(root, selector, c.attempt)
		} else {
			state, err = workflow.Start(root, selector)
		}
		var result response
		if err == nil {
			var body string
			body, err = workflow.Context(root, selector, false)
			result = verbatim(body)
		}
		if err != nil {
			result = configuredOperationFailure(err, action)
		}
		result.diagnostics = append(result.diagnostics, state.Diagnostics...)
		return result
	case "context":
		body, err := workflow.Context(root, selector, c.has("--full"))
		if err != nil {
			return configuredOperationFailure(err, action)
		}
		return verbatim(body)
	case "verify", "done":
		request := workflow.Request{Root: root, Change: selector, Bin: invokingBin(), Quick: c.has("--quick")}
		check := workflow.Verify
		if c.name == "done" {
			check = workflow.Done
		}
		report, err := check(background(), request)
		diagnostics := slices.Concat(report.Diagnostics, report.State.Diagnostics)
		if err != nil {
			result := configuredOperationFailure(err, "Correct the reported issue and run fireweed verify again.")
			// The workflow can return the same cause both as a diagnostic and
			// as the command error. Retain the report without repeating it.
			for _, d := range diagnostics {
				if !d.Warning && d.Code == result.diagnostics[0].Code && strings.Contains(err.Error(), d.Message) {
					result.diagnostics = nil
					break
				}
			}
			result.diagnostics = append(result.diagnostics, diagnostics...)
			return result
		}
		result := output(fmt.Sprintf("%s passed (%s).\n", c.name, report.State.Phase), report)
		result.diagnostics = diagnostics
		return result
	case "release":
		if err := workflow.Release(root, selector); err != nil {
			return configuredOperationFailure(err, "Use fireweed release from the worktree that owns the claim and check the change selector.")
		}
		return message("Released " + selector + ".")
	case "commits", "abandon":
		var err error
		if c.name == "commits" {
			var commits []string
			commits, err = workflow.Commits(root, selector)
			if err == nil {
				var plain strings.Builder
				for _, commit := range commits {
					fmt.Fprintln(&plain, commit)
				}
				return output(plain.String(), array(commits))
			}
		} else {
			var commit string
			commit, err = workflow.Abandon(root, selector)
			if err == nil {
				return message(fmt.Sprintf("Abandoned %s in commit %s.", selector, commit))
			}
		}
		if cause(err, "") == "U02" {
			return failed(2, err, "U02", "Use fireweed list to check the change selector.")
		}
		return configuredOperationFailure(err, action)
	case "gc":
		if _, err := project.Load(root); err != nil {
			return configFailure(err)
		}
		result, err := workflow.GC(background(), workflow.GCRequest{Root: root})
		if err != nil {
			return configuredOperationFailure(err, "Correct the reported cleanup problem and run fireweed gc again.")
		}
		slices.Sort(result.Pruned)
		var plain strings.Builder
		if len(result.Pruned) == 0 {
			fmt.Fprintln(&plain, "No refs pruned.")
		}
		for _, ref := range result.Pruned {
			fmt.Fprintln(&plain, ref)
		}
		return output(plain.String(), struct {
			Pruned []string `json:"pruned"`
		}{array(result.Pruned)})
	}
	panic("unexpected workflow command: " + c.name)
}

type changeSummary struct {
	Change string `json:"change"`
	Run    string `json:"run"`
	Phase  string `json:"phase"`
}

func summarize(state workflow.State) changeSummary {
	return changeSummary{state.Change.Slug, state.Run, state.Phase}
}

func (c command) runCommand(root string) response {
	if c.name == "drill" {
		return drill(root, c.arg(0))
	}
	bin, allow := invokingBin(), c.has("--allow-unmerged")
	action := "Correct the reported problem and use fireweed " + c.name + " to resume."
	if c.name == "next" {
		step, err := runner.Next(background(), runner.NextRequest{Root: root, Bin: bin, AllowUnmerged: allow})
		var reason string
		var plain string
		status := 0
		switch {
		case runner.IsUsageStop(err):
			status, reason = 3, err.Error()
			plain = fmt.Sprintf("Stopped on %s: %s. Use fireweed next to resume.\n", step.State.Change.Slug, reason)
		case err != nil:
			result := configuredOperationFailure(err, action)
			result.diagnostics = append(result.diagnostics, step.State.Diagnostics...)
			return result
		case step.Empty:
			plain = "Queue is empty.\n"
		default:
			plain = fmt.Sprintf("Applied %s.\n", step.State.Change.Slug)
		}
		data := struct {
			changeSummary
			Empty  bool   `json:"empty"`
			Reason string `json:"reason,omitempty"`
		}{summarize(step.State), step.Empty, reason}
		return response{status: status, plain: plain, data: data, diagnostics: step.State.Diagnostics}
	}

	run, err := runner.Run(background(), runner.RunRequest{Root: root, Bin: bin, AllowUnmerged: allow, Max: c.max, For: c.duration})
	status, reason := 0, ""
	switch {
	case runner.IsUsageStop(err) || runner.IsTimeStop(err):
		status, reason = 3, err.Error()
	case err != nil:
		return configuredOperationFailure(err, action)
	}
	applied := make([]changeSummary, 0, len(run.Applied))
	var diagnostics []project.Diagnostic
	var plain strings.Builder
	fmt.Fprintf(&plain, "Applied %d changes.\n", len(run.Applied))
	for _, step := range run.Applied {
		summary := summarize(step.State)
		applied = append(applied, summary)
		diagnostics = append(diagnostics, step.State.Diagnostics...)
		fmt.Fprintln(&plain, summary.Change)
	}
	switch {
	case status == 3:
		fmt.Fprintf(&plain, "Stopped: %s. Use fireweed run to resume.\n", reason)
	case run.Empty:
		fmt.Fprintln(&plain, "Queue is empty.")
	case run.Limited:
		fmt.Fprintln(&plain, "Reached the maximum. Changes remain queued.")
	}
	data := struct {
		Applied []changeSummary `json:"applied"`
		Empty   bool            `json:"empty"`
		Limited bool            `json:"limited"`
		Reason  string          `json:"reason,omitempty"`
	}{applied, run.Empty, run.Limited, reason}
	return response{status: status, plain: plain.String(), data: data, diagnostics: diagnostics}
}

func drill(root, name string) response {
	snapshot, err := loadGraph(root)
	if err != nil {
		return configFailure(err)
	}
	if _, found := moduleByName(snapshot, name); !found {
		return failure(2, "U02", "", fmt.Sprintf("Unknown module %q. Use fireweed modules to list modules.", name))
	}
	result, err := runner.Drill(background(), runner.DrillRequest{Root: root, Module: name, Bin: invokingBin()})
	data := struct {
		Module string `json:"module"`
		Result string `json:"result"`
		Report string `json:"report"`
		Ref    string `json:"ref"`
		Base   string `json:"base_commit"`
		Reason string `json:"reason,omitempty"`
	}{result.Module, result.Result, result.ReportPath, result.Ref, result.Base, ""}
	var artifacts string
	if result.ReportPath != "" {
		artifacts += " Report: " + result.ReportPath + "."
	}
	if result.Ref != "" {
		artifacts += " Rebuilt ref: " + result.Ref + "."
	}
	if runner.IsUsageStop(err) && result.Result == "stopped" {
		data.Reason = err.Error()
		return response{status: 3, data: data, plain: fmt.Sprintf("Drill %s stopped: %s.%s\n", name, data.Reason, artifacts)}
	}
	if err != nil {
		return configuredOperationFailure(err, "Correct the reported issue before retrying."+artifacts)
	}
	return output(fmt.Sprintf("Drill %s: %s.%s\n", name, result.Result, artifacts), data)
}
