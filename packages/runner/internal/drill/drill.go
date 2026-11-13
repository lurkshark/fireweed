package drill

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"syscall"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner/internal/invoke"
	"github.com/lurkshark/fireweed/packages/runner/internal/policy"
	"github.com/lurkshark/fireweed/packages/runner/internal/workspace"
	"github.com/lurkshark/fireweed/packages/runner/model"
)

func Run(ctx context.Context, request model.DrillRequest) (result model.DrillResult, err error) {
	result = model.DrillResult{Module: request.Module, Result: "fail"}
	repo, snapshot, target, err := preflight(ctx, request)
	if err != nil {
		return result, err
	}
	usage, err := policy.UsagePolicy(snapshot.Config)
	if err != nil {
		return result, err
	}
	timeout, err := policy.Timeout(snapshot.Config, request.Timeout)
	if err != nil {
		return result, err
	}
	run, err := invoke.NewRun()
	if err != nil {
		return result, err
	}
	now := time.Now().UTC()
	result.Base, result.Date = repo.Head, now.Format("20060102T150405.000000000")
	result.Agent = invoke.Name(snapshot.Config, "worker")
	result.Model = configuredModel(snapshot.Config)
	result.Dependents = make(map[string]checks.Report)
	dependents := dependentModules(snapshot, target.Name)

	var scratch, input, notes string
	var observations []string
	defer func() {
		var cleanupErr error
		if input != "" {
			if removeErr := os.RemoveAll(input); removeErr != nil {
				cleanupErr = fmt.Errorf("remove reconstruction workspace: %w", removeErr)
			}
		}
		if scratch != "" {
			cleanupErr = policy.WithSecondary(cleanupErr, removeWorktree(repo.Root, scratch))
		}
		if cleanupErr != nil {
			result.Result = "fail"
			err = policy.WithSecondary(cleanupErr, err)
		}
		if len(result.Steps) > 0 {
			if reportErr := writeReport(repo.Root, now, &result, dependents, notes, observations); reportErr != nil {
				result.Result = "fail"
				err = policy.WithSecondary(reportErr, err)
			}
		}
	}()

	scratch, err = addWorktree(ctx, repo, run)
	if err != nil {
		return result, err
	}
	files, err := trackedFiles(ctx, scratch, repo.Head)
	if err != nil {
		return result, err
	}
	input, err = os.MkdirTemp("", "fireweed-drill-input-*")
	if err != nil {
		return result, err
	}
	disposable, err := prepareInput(ctx, snapshot, target, dependents, files, scratch, input)
	if err != nil {
		return result, err
	}
	if err := setup(ctx, snapshot.Config, scratch, request.Bin, timeout); err != nil {
		return result, err
	}
	if err := deleteDisposable(snapshot, target, scratch); err != nil {
		return result, err
	}
	if err := setup(ctx, snapshot.Config, input, request.Bin, timeout); err != nil {
		return result, err
	}
	initial, err := inputFiles(input)
	if err != nil {
		return result, err
	}
	prompt, err := reconstructionPrompt(snapshot, target, initial, request.Bin)
	if err != nil {
		return result, err
	}

	inputSnapshot, scratchSnapshot := snapshot, snapshot
	inputSnapshot.Root, scratchSnapshot.Root = input, scratch
	agentRequest := model.AgentRequest{
		Snapshot: inputSnapshot, Role: "worker", Bin: request.Bin, Timeout: timeout,
		LogRoot: repo.Root, Run: run, Phase: "drill",
	}
	var budget policy.Budget
	var feedback, retained string
	attempts := policy.Attempts(snapshot.Config)
	for attempt := 1; attempt <= attempts; {
		if ctx.Err() != nil {
			return result, context.Cause(ctx)
		}
		agentRequest.Attempt, agentRequest.Prompt = attempt, prompt+feedback
		agent, agentErr := invoke.Agent(ctx, agentRequest)
		result.Steps = append(result.Steps, agent)
		if file, readErr := workspace.Read(input, "DRILL_NOTES.md"); readErr == nil && file.Mode.IsRegular() {
			notes = file.Data
		}
		if ctx.Err() != nil {
			return result, context.Cause(ctx)
		}
		if errors.Is(agentErr, context.DeadlineExceeded) || errors.Is(agentErr, context.Canceled) {
			observations = append(observations, agentErr.Error())
			return result, agentErr
		}
		submission, err := inputFiles(input)
		if err == nil {
			err = unchangedInputs(initial, submission)
		}
		if err != nil {
			observations = append(observations, err.Error())
			return result, err
		}
		if agentErr != nil {
			failure := fmt.Errorf("reconstruction failed: %w\n%s\n%s", agentErr, agent.Stdout, agent.Stderr)
			observations = append(observations, failure.Error())
			if agent.UsageLimit {
				if err := budget.Wait(ctx, usage); err != nil {
					if policy.IsUsageStop(err) {
						result.Result = "stopped"
					}
					observations = append(observations, err.Error())
					return result, err
				}
				continue
			}
			if attempt == attempts {
				return result, failure
			}
			feedback = "\n\n## Last reconstruction failure\n" + failure.Error()
			attempt++
			continue
		}

		oid, err := accept(ctx, snapshot, target, scratch, repo.Head, disposable, initial, submission)
		if err != nil {
			observations = append(observations, err.Error())
			return result, err
		}
		ref := "refs/fireweed/drills/" + target.Name + "/" + result.Date
		if err := gitstore.CASRef(repo.Root, ref, retained, oid); err != nil {
			observations = append(observations, err.Error())
			return result, err
		}
		retained, result.Ref = oid, ref

		hook := checks.HookRequest{Snapshot: scratchSnapshot, Bin: request.Bin, Timeout: timeout}
		var ownErr error
		result.Own, ownErr = evaluate(ctx, hook, target)
		var evalErr error
		if ownErr != nil {
			evalErr = fmt.Errorf("%s: %w", target.Name, ownErr)
		}
		for _, module := range dependents {
			if ctx.Err() != nil {
				return result, context.Cause(ctx)
			}
			report, err := evaluate(ctx, hook, module)
			result.Dependents[module.Name] = report
			if err != nil {
				evalErr = policy.WithSecondary(evalErr, fmt.Errorf("%s: %w", module.Name, err))
			}
		}
		if ctx.Err() != nil {
			return result, context.Cause(ctx)
		}
		if evalErr == nil {
			result.Result = "pass"
			return result, nil
		}
		observations = append(observations, evalErr.Error())
		if ownErr == nil || attempt == attempts {
			return result, evalErr
		}
		feedback = "\n\n## Last own-eval failure\n" + ownErr.Error()
		attempt++
	}
	panic("unreachable reconstruction attempt count")
}

func preflight(ctx context.Context, request model.DrillRequest) (gitstore.Repository, project.Snapshot, project.Module, error) {
	var target project.Module
	repo, err := gitstore.Inspect(request.Root)
	if err != nil {
		return repo, project.Snapshot{}, target, err
	}
	snapshot, err := project.Load(repo.Root)
	if err != nil {
		return repo, snapshot, target, err
	}
	if err := project.ValidateGraph(snapshot.Modules); err != nil {
		return repo, snapshot, target, err
	}
	for _, module := range snapshot.Modules {
		if module.Name == request.Module {
			target = module
			break
		}
	}
	if target.Name == "" {
		return repo, snapshot, target, contracts.CodedError{Code: "U02", Subject: request.Module, Cause: fmt.Errorf("unknown module %q", request.Module)}
	}
	status, err := workspace.Git(ctx, repo.Root, "--no-optional-locks", "status", "--porcelain", "--untracked-files=all")
	if err != nil {
		return repo, snapshot, target, err
	}
	if status != "" {
		return repo, snapshot, target, contracts.CodedError{Code: "P01", Subject: repo.Root, Cause: fmt.Errorf("drill requires a clean checkout; %s is dirty", repo.Root)}
	}
	claims, err := workspace.Git(ctx, repo.Root, "for-each-ref", "--format=%(refname)", "refs/fireweed/claims/")
	if err == nil && claims != "" {
		err = contracts.CodedError{Code: "P02", Subject: repo.Root, Cause: fmt.Errorf("drill requires no live shared claims: %s", strings.TrimSpace(claims))}
	}
	return repo, snapshot, target, err
}

func addWorktree(ctx context.Context, repo gitstore.Repository, run string) (string, error) {
	parent := filepath.Join(repo.CommonDir, "fireweed", "worktrees")
	if err := os.MkdirAll(parent, 0755); err != nil {
		return "", err
	}
	root := filepath.Join(parent, "drill-"+run)
	if _, err := workspace.Git(ctx, repo.Root, "worktree", "add", "--quiet", "--detach", root, repo.Head); err != nil {
		if _, statErr := os.Stat(root); statErr == nil {
			_, cleanup := workspace.Git(context.Background(), repo.Root, "worktree", "remove", "--force", root)
			err = errors.Join(err, cleanup)
		}
		return "", err
	}
	return root, nil
}

func removeWorktree(root, scratch string) error {
	// Git may drop the registration even when it cannot remove the directory.
	// Check access first so the caller can still retry cleanup after a failure.
	const writeAndSearch = 3
	if err := syscall.Access(filepath.Dir(scratch), writeAndSearch); err != nil {
		return fmt.Errorf("remove verification worktree %s: %w", scratch, err)
	}
	_, err := workspace.Git(context.Background(), root, "worktree", "remove", "--force", scratch)
	return err
}

func setup(ctx context.Context, config map[string]any, root, bin string, timeout time.Duration) error {
	command := invoke.Mapping(config["hooks"])["setup"]
	if command == nil {
		return nil
	}
	args, err := invoke.Args(command)
	if err != nil {
		return fmt.Errorf("setup hook: %w", err)
	}
	process, err := gitstore.Run(ctx, gitstore.Request{
		Args: args, Dir: root, Timeout: timeout, Env: map[string]string{"FIREWEED_BIN": bin},
	})
	if err != nil {
		return fmt.Errorf("setup hook: %w\n%s%s", err, process.Stdout, process.Stderr)
	}
	return nil
}

func dependentModules(snapshot project.Snapshot, target string) []project.Module {
	names := map[string]bool{target: true}
	var result []project.Module
	for {
		count := len(result)
		for _, module := range snapshot.Modules {
			if names[module.Name] || !slices.ContainsFunc(module.DependsOn, func(name string) bool { return names[name] }) {
				continue
			}
			names[module.Name] = true
			result = append(result, module)
		}
		if len(result) == count {
			break
		}
	}
	slices.SortFunc(result, func(a, b project.Module) int { return strings.Compare(a.Name, b.Name) })
	return result
}

func evaluate(ctx context.Context, request checks.HookRequest, module project.Module) (checks.Report, error) {
	request.Module = module.Name
	report, err := checks.Evaluate(ctx, request)
	if err == nil {
		body, _ := contracts.Section(module.Document.Body, "Specs")
		var specs []contracts.Item
		specs, err = contracts.ParseItems(body)
		if err == nil {
			err = policy.WithSecondary(checks.RequirePass(report, specs), checks.RequireCoverage(report, specs))
		}
	}
	if err != nil {
		var detail strings.Builder
		for _, c := range report.Cases {
			if c.Result == "fail" || c.Result == "error" {
				fmt.Fprintf(&detail, "\n%s: %s", c.Name, c.Failure)
			}
		}
		if report.Stdout != "" || report.Stderr != "" {
			fmt.Fprintf(&detail, "\n%s\n%s", report.Stdout, report.Stderr)
		}
		err = fmt.Errorf("%w%s", err, detail.String())
	}
	return report, err
}
