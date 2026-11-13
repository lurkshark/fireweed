package lifecycle

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func Verify(ctx context.Context, request model.Request) (model.Result, error) {
	state, held, err := loadHeld(request.Root, request.Change)
	result := model.Result{State: state}
	if err != nil {
		return result, err
	}
	return verify(ctx, request, state, held)
}

func supported(state model.State) error {
	switch state.Phase {
	case "evals", "audit", "impl":
		return nil
	default:
		return failure("P01", state.Change.Path, "unsupported phase %s", state.Phase)
	}
}

func verify(ctx context.Context, request model.Request, state model.State, held claim) (model.Result, error) {
	result := model.Result{State: state}
	if err := supported(state); err != nil {
		return result, err
	}
	if err := verifyApproval(state, held, false); err != nil {
		return result, err
	}
	for _, d := range project.Check(state.Snapshot.Root) {
		if d.Code == "C08" && !d.Warning {
			return result, staticFailure([]project.Diagnostic{d})
		}
	}
	if _, err := checkZones(&result); err != nil {
		return result, err
	}
	if state.Phase == "audit" {
		_, err := auditSubmission(state, held)
		return result, err
	}
	return verifyEvals(ctx, request, result)
}

func checkZones(result *model.Result) ([]gitstore.Change, error) {
	state := result.State
	changes, err := gitstore.Changes(state.Snapshot.Root, state.PhaseBase)
	if err != nil {
		return nil, err
	}
	changes = slices.DeleteFunc(changes, func(c gitstore.Change) bool {
		return isRunArtifact(c.Path) && (c.OldPath == "" || isRunArtifact(c.OldPath))
	})
	zone := checks.ZoneRequest{
		Snapshot: state.Snapshot, Module: state.Change.Module, Phase: state.Phase,
		ChangePath: state.Change.Path, Base: state.PhaseBase, Changes: changes,
		InterfaceDeclared: strings.TrimSpace(section(state.Change.Document, "Interface")) != "",
	}
	if evalFix(state.Change) {
		zone.Phase = "audit"
		if state.Phase == "evals" {
			module := moduleNamed(state.Snapshot, state.Change.Module)
			zone.Changes = nil
			for _, change := range changes {
				if fileZone(state.Snapshot, module, change.Path) == "evals" &&
					(change.OldPath == "" || fileZone(state.Snapshot, module, change.OldPath) == "evals") {
					continue
				}
				zone.Changes = append(zone.Changes, change)
			}
		}
	}
	diagnostics := checks.Zones(zone)
	result.Diagnostics = append(result.Diagnostics, diagnostics...)
	return changes, staticFailure(diagnostics)
}

func evalFix(change project.Change) bool {
	return len(stringsOf(change.Document.Frontmatter["eval_fix"])) > 0
}

func changedIDs(change project.Change, specs []contracts.Item) []string {
	var ids []string
	for _, item := range slices.Concat(change.Added, change.Modified) {
		ids = append(ids, item.ID)
	}
	for _, field := range []string{"fixes", "eval_fix"} {
		ids = append(ids, stringsOf(change.Document.Frontmatter[field])...)
	}
	var enforced []string
	for _, id := range ids {
		for _, spec := range specs {
			tag := strings.ToLower(spec.Tag)
			if spec.ID == id && !strings.HasPrefix(tag, "review") && !strings.HasPrefix(tag, "quarantine") &&
				!slices.Contains(enforced, id) {
				enforced = append(enforced, id)
			}
		}
	}
	return enforced
}

func verifyEvals(ctx context.Context, request model.Request, result model.Result) (model.Result, error) {
	state := result.State
	snapshot := state.Snapshot
	selected := []project.Module{moduleNamed(snapshot, state.Change.Module)}
	fullImpl := state.Phase == "impl" && !request.Quick
	mode := str(snapshot.Config["verify_dependents"])
	if fullImpl && mode != "none" {
		for _, module := range snapshot.Modules {
			if dependsOn(snapshot, module.Name, state.Change.Module, mode == "all") {
				selected = append(selected, module)
			}
		}
	}

	result.Reports = make(map[string]checks.Report)
	var verdict, executionErr error
	for i, module := range selected {
		report, err := checks.Evaluate(ctx, checks.HookRequest{
			Snapshot: snapshot, Module: module.Name, Bin: request.Bin, Timeout: request.Timeout,
		})
		if err != nil {
			executionErr = err
			break
		}
		result.Reports[module.Name] = report
		cost, err := measure(snapshot, module, report)
		if err != nil {
			executionErr = err
			break
		}
		result.Costs = append(result.Costs, cost)
		if err := retainReport(state, module.Name, report); err != nil {
			executionErr = err
			break
		}

		specs, err := items(module.Document, "Specs")
		if err != nil {
			executionErr = err
			break
		}
		if i == 0 {
			err = checks.RequireCoverage(report, specs)
		}
		if err == nil {
			err = requirePhaseResult(state, report, specs, i == 0)
		}
		if verdict == nil {
			verdict = err
		}
	}

	if len(result.Costs) > 0 {
		budgets := budgetWarnings(snapshot, result.Costs)
		result.Diagnostics = append(result.Diagnostics, budgets...)
		if err := retainCosts(state, result.Costs, budgets); err != nil && executionErr == nil {
			executionErr = err
		}
	}
	if executionErr != nil {
		return result, executionErr
	}
	if fullImpl {
		diagnostics, err := checks.Boundaries(ctx, checks.BoundaryRequest{
			Snapshot: snapshot, Bin: request.Bin, Timeout: request.Timeout,
		})
		result.Diagnostics = append(result.Diagnostics, diagnostics...)
		if verdict == nil {
			verdict = boundaryFailure(diagnostics, err)
		}
	}
	return result, verdict
}

func requirePhaseResult(state model.State, report checks.Report, specs []contracts.Item, target bool) error {
	if state.Phase == "evals" && !evalFix(state.Change) {
		return checks.RequireFailFirst(report, changedIDs(state.Change, specs),
			stringsOf(state.Change.Document.Frontmatter["already_satisfied"]))
	}
	if target && state.Phase == "evals" {
		for _, id := range stringsOf(state.Change.Document.Frontmatter["eval_fix"]) {
			if !slices.ContainsFunc(report.Cases, func(c checks.Case) bool { return slices.Contains(c.IDs, id) }) {
				return failure("C10", id, "no coverage for %s", id)
			}
		}
		for _, c := range report.Cases {
			if c.Result == "fail" || c.Result == "error" {
				return failure("E01", c.Name, "failed case %q for %s", c.Name, strings.Join(c.IDs, ", "))
			}
		}
		return nil
	}
	if !target {
		specs = slices.DeleteFunc(specs, func(item contracts.Item) bool {
			return strings.HasPrefix(strings.ToLower(item.Tag), "review")
		})
	}
	return checks.RequirePass(report, specs)
}

func boundaryFailure(diagnostics []project.Diagnostic, err error) error {
	if err == nil {
		return nil
	}
	code, _, assigned := contracts.ErrorCode(err)
	if assigned && code != "X01" {
		return err
	}
	for _, wanted := range []string{"X02", "X01"} {
		for _, d := range diagnostics {
			if d.Code != wanted || d.Warning {
				continue
			}
			if wanted == "X01" {
				return coded(d.Code, d.Path, fmt.Errorf("%s: %w", d.Message, err))
			}
			return coded(d.Code, d.Path, err)
		}
	}
	return err
}

type verificationFailure struct{ cause error }

func (e verificationFailure) Error() string { return e.cause.Error() }
func (e verificationFailure) Unwrap() error { return e.cause }

func IsVerificationFailure(err error) bool {
	var marked verificationFailure
	return errors.As(err, &marked)
}
