package next

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner/internal/invoke"
	"github.com/lurkshark/fireweed/packages/runner/model"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func Run(ctx context.Context, request model.NextRequest) (model.NextResult, error) {
	return RunWithBudget(ctx, request, &UsageBudget{})
}

// RunWithBudget shares completed usage waits across changes in one Run call.
func RunWithBudget(ctx context.Context, request model.NextRequest, budget *UsageBudget) (model.NextResult, error) {
	var result model.NextResult
	if stop := budget.TimeStop(ctx); stop != nil {
		return result, stop
	}
	state, currentErr := workflow.Current(request.Root, "")
	if currentErr != nil && !strings.Contains(currentErr.Error(), "P02: worktree holds no change") {
		return result, fmt.Errorf("read held change: %w", currentErr)
	}
	if currentErr != nil {
		queue, err := workflow.Eligible(request.Root)
		if err != nil {
			return result, fmt.Errorf("read eligible changes: %w", err)
		}
		if len(queue) == 0 {
			result.Empty = true
			return result, nil
		}
		selected := &queue[0]
		snapshot, err := project.Load(request.Root)
		if err != nil {
			return result, fmt.Errorf("load usage policy: %w", err)
		}
		if _, err := readUsagePolicy(snapshot.Config); err != nil {
			return result, err
		}
		if !request.AllowUnmerged {
			if err := requireApproval(request.Root, selected.Path); err != nil {
				return result, err
			}
		}
		state, err = workflow.Start(request.Root, selected.Path)
		result.State = state
		if err != nil {
			return result, fmt.Errorf("start %s: %w", selected.Path, err)
		}
	}
	result.State = state
	if stop := budget.TimeStop(ctx); stop != nil {
		return result, stop
	}
	policy, err := readUsagePolicy(state.Snapshot.Config)
	if err != nil {
		return result, err
	}
	var recoveryWarnings []project.Diagnostic
	for state.Phase != "applied" {
		if stop := budget.TimeStop(ctx); stop != nil {
			return result, stop
		}
		phase := state.Phase
		role := "worker"
		var prompt string
		switch phase {
		case "evals", "impl":
			prompt, err = workflow.Context(state.Snapshot.Root, state.Change.Path, true)
		case "audit":
			role = "audit"
			prompt, err = auditPrompt(state)
		default:
			return result, fmt.Errorf("unsupported next phase %q", phase)
		}
		if err != nil {
			return result, fmt.Errorf("prepare %s context: %w", phase, err)
		}
		step := workflow.Request{
			Root: state.Snapshot.Root, Change: state.Change.Path, Bin: request.Bin,
			Agent: configuredCommand(state.Snapshot.Config, role), Timeout: request.Timeout,
		}
		attempts := configuredAttempts(state.Snapshot.Config)
		feedback := ""
		for attempt := 0; attempt < attempts; {
			if stop := budget.TimeStop(ctx); stop != nil {
				return result, stop
			}
			preInvoke, inspectErr := gitstore.Inspect(state.Snapshot.Root)
			if inspectErr != nil {
				return result, fmt.Errorf("inspect held phase HEAD: %w", inspectErr)
			}
			if preInvoke.Head != state.PhaseBase {
				return result, fmt.Errorf("P05: held %s phase HEAD moved from %s to %s before agent invocation", phase, state.PhaseBase, preInvoke.Head)
			}
			agent, invokeErr := invoke.Agent(ctx, model.AgentRequest{
				Snapshot: state.Snapshot, Role: role, Prompt: prompt + feedback,
				Bin: request.Bin, Timeout: request.Timeout,
			})
			result.Steps = append(result.Steps, model.StepResult{Phase: phase, Role: role, Agent: agent})
			recovered, recoveryErr := recoverAgentCommits(state.Snapshot.Root, state.PhaseBase)
			if recoveryErr != nil {
				return result, errors.Join(fmt.Errorf("recover %s agent commits: %w", phase, recoveryErr), invokeErr)
			}
			if recovered {
				recoveryWarnings = append(recoveryWarnings, project.Diagnostic{
					Code: "P05", Warning: true,
					Message: fmt.Sprintf("%s agent created commits; soft-reset HEAD to the held phase base and retained index and worktree work", phase),
				})
				result.State = withRecoveryWarnings(state, recoveryWarnings)
			}
			if stop := budget.TimeStop(ctx); stop != nil {
				return result, stop
			}
			if agent.UsageLimit && ctx.Err() == nil && !errors.Is(invokeErr, context.Canceled) && !errors.Is(invokeErr, context.DeadlineExceeded) {
				if err := budget.Wait(ctx, policy); err != nil {
					return result, err
				}
				continue
			}
			if invokeErr != nil {
				return result, fmt.Errorf("invoke %s agent: %w", phase, invokeErr)
			}
			if agent.ExitCode != 0 {
				return result, fmt.Errorf("%s agent exited with code %d", phase, agent.ExitCode)
			}
			attempt++
			verification, verifyErr := workflow.Verify(ctx, step)
			if stop := budget.TimeStop(ctx); stop != nil {
				return result, stop
			}
			if verifyErr == nil {
				break
			}
			if attempt == attempts {
				blocked, blockErr := workflow.Block(workflow.BlockRequest{
					Root: state.Snapshot.Root, Change: state.Change.Path,
					Reason: phase + " verification attempts exhausted",
				})
				if blockErr != nil {
					return result, fmt.Errorf("verify %s: %w; block exhausted step: %w", phase, verifyErr, blockErr)
				}
				result.State = withRecoveryWarnings(blocked.State, recoveryWarnings)
				return result, &BlockedError{Phase: phase, AttemptRef: blocked.AttemptRef, Cause: verifyErr}
			}
			feedback = verificationFeedback(phase, verification, verifyErr)
		}
		if stop := budget.TimeStop(ctx); stop != nil {
			return result, stop
		}
		done, err := workflow.Done(ctx, step)
		if err != nil {
			if done.State.Phase != "" {
				result.State = withRecoveryWarnings(done.State, recoveryWarnings)
			}
			if stop := budget.TimeStop(ctx); stop != nil {
				return result, stop
			}
			return result, fmt.Errorf("complete %s: %w", phase, err)
		}
		state = done.State
		result.State = withRecoveryWarnings(state, recoveryWarnings)
	}
	return result, nil
}

func recoverAgentCommits(root, phaseBase string) (bool, error) {
	repo, err := gitstore.Inspect(root)
	if err != nil {
		return false, err
	}
	if repo.Head == phaseBase {
		return false, nil
	}
	reset, err := gitstore.Run(context.Background(), gitstore.Request{
		Args:    []string{"git", "-C", root, "-c", "core.hooksPath=/dev/null", "reset", "--soft", phaseBase},
		Timeout: 30 * time.Second,
		Env:     map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0"},
	})
	if err != nil {
		return false, fmt.Errorf("git reset --soft: %w: %s", err, strings.TrimSpace(reset.Stderr))
	}
	return true, nil
}

func withRecoveryWarnings(state workflow.State, warnings []project.Diagnostic) workflow.State {
	state.Diagnostics = append(state.Diagnostics, warnings...)
	return state
}

// BlockedError identifies the verification failure returned after Block succeeds.
type BlockedError struct {
	Phase      string
	AttemptRef string
	Cause      error
}

func (e *BlockedError) Error() string {
	return fmt.Sprintf("verify %s after exhausted attempts; saved work at %s: %v", e.Phase, e.AttemptRef, e.Cause)
}

func (e *BlockedError) Unwrap() error { return e.Cause }

func configuredAttempts(config map[string]any) int {
	switch attempts := config["attempts"].(type) {
	case int:
		if attempts > 0 {
			return attempts
		}
	case float64:
		if attempts > 0 && attempts <= float64(int(^uint(0)>>1)) && attempts == float64(int(attempts)) {
			return int(attempts)
		}
	}
	return 2
}

func verificationFeedback(phase string, verification workflow.Result, verifyErr error) string {
	var feedback strings.Builder
	fmt.Fprintf(&feedback, "\n\n# Previous %s verification failed\n%s\n", phase, verifyErr)
	for _, diagnostic := range verification.Diagnostics {
		if diagnostic.Code != "" {
			fmt.Fprintf(&feedback, "%s: %s\n", diagnostic.Code, diagnostic.Message)
		} else {
			fmt.Fprintln(&feedback, diagnostic.Message)
		}
	}
	names := make([]string, 0, len(verification.Reports))
	for name := range verification.Reports {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		report := verification.Reports[name]
		fmt.Fprintf(&feedback, "\n# %s verification output\n", name)
		feedback.WriteString(report.Stdout)
		if !strings.HasSuffix(report.Stdout, "\n") {
			feedback.WriteByte('\n')
		}
		feedback.WriteString(report.Stderr)
		if !strings.HasSuffix(report.Stderr, "\n") {
			feedback.WriteByte('\n')
		}
	}
	return feedback.String()
}

func requireApproval(root, path string) error {
	snapshot, err := project.Load(root)
	if err != nil {
		return fmt.Errorf("load approval configuration: %w", err)
	}
	remote, _ := snapshot.Config["remote"].(string)
	if remote == "" {
		remote = "origin"
	}
	base, _ := snapshot.Config["base"].(string)
	ref, err := gitstore.BaseRef(root, remote, base)
	if err != nil {
		return fmt.Errorf("resolve approval base: %w", err)
	}
	refs := []string{ref}
	if strings.HasPrefix(ref, "refs/remotes/"+remote+"/") {
		branch := strings.TrimPrefix(ref, "refs/remotes/"+remote+"/")
		refs = append(refs, "refs/heads/"+branch)
	}
	for _, candidate := range refs {
		if _, err := gitstore.ReadAt(root, candidate, path); err == nil {
			return nil
		}
	}
	return &approvalRequiredError{path: path}
}

func configuredCommand(config map[string]any, role string) string {
	agent, _ := config["agent"].(map[string]any)
	command := agent["command"]
	if role == "audit" {
		audit, _ := config["audit"].(map[string]any)
		if !unsetCommand(audit["command"]) {
			command = audit["command"]
		}
	}
	switch value := command.(type) {
	case string:
		return value
	case []string:
		return strings.Join(value, " ")
	case []any:
		parts := make([]string, len(value))
		for i, part := range value {
			parts[i] = fmt.Sprint(part)
		}
		return strings.Join(parts, " ")
	default:
		return fmt.Sprint(command)
	}
}

func unsetCommand(command any) bool {
	switch value := command.(type) {
	case nil:
		return true
	case string:
		return value == ""
	case []string:
		return len(value) == 0
	case []any:
		return len(value) == 0
	default:
		return false
	}
}

func auditPrompt(state workflow.State) (string, error) {
	change, err := os.ReadFile(filepath.Join(state.Snapshot.Root, filepath.FromSlash(state.Change.Path)))
	if err != nil {
		return "", fmt.Errorf("read held change: %w", err)
	}
	var module *project.Module
	for i := range state.Snapshot.Modules {
		if state.Snapshot.Modules[i].Name == state.Change.Module {
			module = &state.Snapshot.Modules[i]
			break
		}
	}
	if module == nil {
		return "", fmt.Errorf("held module %q is absent", state.Change.Module)
	}
	spec, err := contracts.RenderDocument(module.Document)
	if err != nil {
		return "", fmt.Errorf("render projected spec: %w", err)
	}
	playbook, err := assets.Playbook("audit")
	if err != nil {
		return "", fmt.Errorf("load audit playbook: %w", err)
	}
	var prompt strings.Builder
	fmt.Fprintf(&prompt, "# Audit evals\n\nSubmit findings in %s.\n\n# Playbook\n%s\n\n# Change: %s\n%s\n\n# Projected spec: %s/SPEC.md\n%s\n", state.Change.Path, playbook, state.Change.Path, change, module.Dir, spec)

	root := filepath.Join(state.Snapshot.Root, filepath.FromSlash(module.Dir))
	var evals []string
	err = filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		for _, pattern := range module.Zones.Evals {
			if contracts.MatchGlob(pattern, rel) {
				evals = append(evals, path)
				break
			}
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("list audit evals: %w", err)
	}
	sort.Strings(evals)
	for _, path := range evals {
		contents, err := os.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("read audit eval %s: %w", path, err)
		}
		rel, err := filepath.Rel(state.Snapshot.Root, path)
		if err != nil {
			return "", err
		}
		fmt.Fprintf(&prompt, "\n# Eval: %s\n%s\n", filepath.ToSlash(rel), contents)
	}
	return prompt.String(), nil
}
