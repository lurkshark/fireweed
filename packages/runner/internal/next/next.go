package next

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner/internal/invoke"
	"github.com/lurkshark/fireweed/packages/runner/internal/policy"
	"github.com/lurkshark/fireweed/packages/runner/internal/workspace"
	"github.com/lurkshark/fireweed/packages/runner/model"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func IsUsageStop(err error) bool { return policy.IsUsageStop(err) }

type approvalRequired struct{ path string }

func (e approvalRequired) Error() string {
	return fmt.Sprintf("%s requires approval on the local or shared base", e.path)
}

func IsApprovalRequired(err error) bool {
	var approval approvalRequired
	return errors.As(err, &approval)
}

type Blocked struct{ Cause error }

func (e Blocked) Error() string { return e.Cause.Error() }
func (e Blocked) Unwrap() error { return e.Cause }

func Run(ctx context.Context, request model.NextRequest) (model.NextResult, error) {
	return WithBudget(ctx, request, &policy.Budget{})
}

func WithBudget(ctx context.Context, request model.NextRequest, budget *policy.Budget) (model.NextResult, error) {
	var result model.NextResult
	state, usage, empty, err := acquire(ctx, request)
	result.State, result.Empty = state, empty
	if err != nil || empty {
		return result, err
	}
	config := state.Snapshot.Config
	maxAttempts := policy.Attempts(config)
	evalsAttempts := 0
	var warnings []project.Diagnostic
	for _, diagnostic := range state.Diagnostics {
		if diagnostic.Warning {
			warnings = append(warnings, diagnostic)
		}
	}
	for result.State.Phase != "applied" {
		if ctx.Err() != nil {
			return result, context.Cause(ctx)
		}
		state := result.State
		role := "worker"
		attempt := 1
		switch state.Phase {
		case "audit":
			role = "audit"
		case "evals":
			if evalsAttempts == maxAttempts {
				cause := fmt.Errorf("evals attempt allowance exhausted after %d attempts", maxAttempts)
				ref, err := blockStep(&result, warnings, cause.Error())
				if err != nil {
					return result, policy.WithSecondary(cause, err)
				}
				return result, Blocked{Cause: fmt.Errorf("%w (saved attempt %s)", cause, ref)}
			}
			attempt += evalsAttempts
		case "impl":
		default:
			return result, contracts.CodedError{Code: "P01", Subject: state.Change.Path, Cause: fmt.Errorf("unsupported phase %s", state.Phase)}
		}
		var prompt, feedback string
		if role == "worker" {
			var err error
			prompt, err = workflow.Context(state.Snapshot.Root, state.Change.Path, true)
			if err != nil {
				return result, err
			}
		}
		for {
			if ctx.Err() != nil {
				return result, context.Cause(ctx)
			}
			agentRequest := model.AgentRequest{
				Snapshot: state.Snapshot, Role: role, Bin: request.Bin, Timeout: request.Timeout,
				LogRoot: state.Snapshot.Root, Run: state.Run, Phase: state.Phase, Attempt: attempt,
			}
			var agent model.AgentResult
			var invocationErr, stepErr error
			var invoked bool
			if role == "audit" {
				agent, invoked, invocationErr, stepErr = audit(ctx, state, agentRequest, feedback)
			} else {
				agentRequest.Prompt = prompt + feedback
				agent, invocationErr = invoke.Agent(ctx, agentRequest)
				invoked = true
				var recovered bool
				recovered, stepErr = recoverHead(state)
				if recovered {
					warning := project.Diagnostic{Code: "P05", Warning: true, Path: state.Change.Path,
						Message: fmt.Sprintf("soft reset %s worker commits to the phase base and retained the work", state.Phase)}
					warnings = append(warnings, warning)
					result.State.Diagnostics = append(result.State.Diagnostics, warning)
					result.State.Head = state.PhaseBase
				}
			}
			if invoked {
				result.Steps = append(result.Steps, model.StepResult{Phase: state.Phase, Role: role, Agent: agent})
			}
			if stepErr != nil {
				return result, policy.WithSecondary(stepErr, invocationErr)
			}
			if ctx.Err() != nil {
				return result, context.Cause(ctx)
			}
			if invocationErr != nil {
				if !agent.UsageLimit || errors.Is(invocationErr, context.DeadlineExceeded) || errors.Is(invocationErr, context.Canceled) {
					return result, fmt.Errorf("%s: %w", state.Phase, invocationErr)
				}
				if err := budget.Wait(ctx, usage); err != nil {
					return result, err
				}
				continue
			}

			timeout, err := policy.Timeout(config, request.Timeout)
			if err != nil {
				return result, err
			}
			done, err := workflow.Done(ctx, workflow.Request{
				Root: state.Snapshot.Root, Change: state.Change.Path, Bin: request.Bin,
				Agent: invoke.Name(config, role), Timeout: timeout,
			})
			if done.State.Run != "" {
				result.State = done.State
				result.State.Diagnostics = append(result.State.Diagnostics, warnings...)
				result.State.Diagnostics = append(result.State.Diagnostics, done.Diagnostics...)
			}
			if err == nil {
				if state.Phase == "evals" {
					evalsAttempts = attempt
				}
				break
			}
			if ctx.Err() != nil {
				return result, context.Cause(ctx)
			}
			if !workflow.IsVerificationFailure(err) {
				if code, _, _ := contracts.ErrorCode(err); code == "P06" && result.State.ClaimOID == "" {
					return result, Blocked{Cause: err}
				}
				return result, err
			}
			if attempt == maxAttempts {
				reason := fmt.Sprintf("%s verification failed after %d attempts: %v", state.Phase, attempt, err)
				ref, blockErr := blockStep(&result, warnings, reason)
				if blockErr != nil {
					return result, policy.WithSecondary(err, blockErr)
				}
				return result, Blocked{Cause: fmt.Errorf("%s verification exhausted (saved attempt %s): %w", state.Phase, ref, err)}
			}
			feedback = verificationFeedback(done, err)
			attempt++
		}
	}
	return result, nil
}

func blockStep(result *model.NextResult, warnings []project.Diagnostic, reason string) (string, error) {
	state := result.State
	block, err := workflow.Block(workflow.BlockRequest{
		Root: state.Snapshot.Root, Change: state.Change.Path, Reason: reason,
	})
	if err != nil {
		return "", err
	}
	result.State = block.State
	result.State.Diagnostics = append(result.State.Diagnostics, warnings...)
	return block.AttemptRef, nil
}

func acquire(ctx context.Context, request model.NextRequest) (workflow.State, policy.Usage, bool, error) {
	var state workflow.State
	var usage policy.Usage
	if ctx.Err() != nil {
		return state, usage, false, context.Cause(ctx)
	}
	approved, held, err := workflow.Approved(request.Root)
	if err != nil {
		return state, usage, false, err
	}
	if held {
		state, err = workflow.Current(approved.Snapshot.Root, approved.Path)
		if err == nil {
			usage, err = policy.UsagePolicy(state.Snapshot.Config)
		}
		return state, usage, false, err
	}

	eligible, err := workflow.Eligible(request.Root)
	if err != nil || len(eligible) == 0 {
		return state, usage, err == nil, err
	}
	change := eligible[0]
	snapshot, err := project.Load(request.Root)
	if err != nil {
		return state, usage, false, err
	}
	usage, err = policy.UsagePolicy(snapshot.Config)
	if err != nil {
		return state, usage, false, err
	}
	if !request.AllowUnmerged {
		if err := requireApproval(ctx, snapshot, change.Path); err != nil {
			return state, usage, false, err
		}
	}
	if ctx.Err() != nil {
		return state, usage, false, context.Cause(ctx)
	}
	state, err = workflow.Start(snapshot.Root, change.Path)
	return state, usage, false, err
}

func requireApproval(ctx context.Context, snapshot project.Snapshot, change string) error {
	remote, _ := snapshot.Config["remote"].(string)
	base, _ := snapshot.Config["base"].(string)
	if remote == "" {
		remote = "origin"
	}
	shared, err := gitstore.BaseRef(snapshot.Root, remote, base)
	if err != nil {
		return err
	}
	refs := []string{shared}
	if branch, remoteRef := strings.CutPrefix(shared, "refs/remotes/"+remote+"/"); remoteRef {
		refs = append(refs, "refs/heads/"+branch)
	}
	for _, ref := range refs {
		if _, err := workspace.Git(ctx, snapshot.Root, "cat-file", "-e", ref+":"+change); err == nil {
			return nil
		}
	}
	if ctx.Err() != nil {
		return context.Cause(ctx)
	}
	return contracts.CodedError{Code: "P01", Subject: change, Cause: approvalRequired{change}}
}

func recoverHead(state workflow.State) (bool, error) {
	ctx := context.Background()
	root := state.Snapshot.Root
	head, err := workspace.Git(ctx, root, "rev-parse", "--verify", "HEAD")
	if err != nil || strings.TrimSpace(head) == state.PhaseBase {
		return false, err
	}
	// Recover commits made during this invocation. A pre-existing HEAD
	// mismatch must go through workflow's normal stale-base checks.
	if state.Head != state.PhaseBase {
		return false, nil
	}
	_, err = workspace.Git(ctx, root, "reset", "--soft", state.PhaseBase)
	return err == nil, err
}
