package next

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner/internal/invoke"
	"github.com/lurkshark/fireweed/packages/runner/model"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func Run(ctx context.Context, request model.NextRequest) (model.NextResult, error) {
	var result model.NextResult
	queue, err := project.Queue(request.Root)
	if err != nil {
		return result, fmt.Errorf("read change queue: %w", err)
	}
	if len(queue) == 0 {
		result.Empty = true
		return result, nil
	}
	var selected *project.Change
	for i := range queue {
		blocked, _ := queue[i].Document.Frontmatter["blocked"].(string)
		if strings.TrimSpace(blocked) == "" {
			selected = &queue[i]
			break
		}
	}
	if selected == nil {
		return result, fmt.Errorf("all queued changes are blocked")
	}
	if !request.AllowUnmerged {
		if err := requireApproval(request.Root, selected.Path); err != nil {
			return result, err
		}
	}

	state, err := workflow.Start(request.Root, selected.Path)
	result.State = state
	if err != nil {
		return result, fmt.Errorf("start %s: %w", selected.Path, err)
	}
	for state.Phase != "applied" {
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
		for attempt := 0; attempt < attempts; attempt++ {
			agent, invokeErr := invoke.Agent(ctx, model.AgentRequest{
				Snapshot: state.Snapshot, Role: role, Prompt: prompt + feedback,
				Bin: request.Bin, Timeout: request.Timeout,
			})
			result.Steps = append(result.Steps, model.StepResult{Phase: phase, Role: role, Agent: agent})
			if invokeErr != nil {
				return result, fmt.Errorf("invoke %s agent: %w", phase, invokeErr)
			}
			if agent.ExitCode != 0 {
				return result, fmt.Errorf("%s agent exited with code %d", phase, agent.ExitCode)
			}
			verification, verifyErr := workflow.Verify(ctx, step)
			if verifyErr == nil {
				break
			}
			if attempt+1 == attempts {
				return result, fmt.Errorf("verify %s: %w", phase, verifyErr)
			}
			feedback = verificationFeedback(phase, verification, verifyErr)
		}
		done, err := workflow.Done(ctx, step)
		if err != nil {
			if done.State.Phase != "" {
				result.State = done.State
			}
			return result, fmt.Errorf("complete %s: %w", phase, err)
		}
		state = done.State
		result.State = state
	}
	return result, nil
}

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
	return fmt.Errorf("change %s requires approval on the local or shared base", path)
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
