package next

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner/internal/invoke"
	"github.com/lurkshark/fireweed/packages/runner/internal/workspace"
	"github.com/lurkshark/fireweed/packages/runner/model"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func audit(ctx context.Context, state workflow.State, request model.AgentRequest, feedback string) (agent model.AgentResult, invoked bool, invocationErr, stepErr error) {
	root, err := os.MkdirTemp("", "fireweed-audit-*")
	if err != nil {
		return agent, false, nil, err
	}
	defer func() {
		if err := os.RemoveAll(root); err != nil {
			stepErr = errors.Join(stepErr, fmt.Errorf("remove audit workspace: %w", err))
		}
	}()

	var target project.Module
	for _, module := range state.Snapshot.Modules {
		if module.Name == state.Change.Module {
			target = module
			break
		}
	}
	if target.Name == "" {
		return agent, false, nil, fmt.Errorf("audit module %q is missing", state.Change.Module)
	}
	source := state.Snapshot.Root
	changePath := state.Change.Path
	change, err := workspace.Read(source, changePath)
	if err != nil || !change.Mode.IsRegular() {
		return agent, false, nil, fmt.Errorf("read regular audit change %s: %w", changePath, errors.Join(err, fs.ErrInvalid))
	}
	spec, err := contracts.RenderDocument(target.Document)
	if err != nil {
		return agent, false, nil, err
	}
	book, err := auditPlaybook(source)
	if err != nil {
		return agent, false, nil, err
	}

	specPath := path.Join(target.Dir, "SPEC.md")
	inputs := map[string]workspace.File{changePath: change, specPath: {Data: spec, Mode: 0644}}
	var prompt strings.Builder
	fmt.Fprintf(&prompt, "# Change (%s)\n%s\n\n# Projected spec (%s)\n%s\n\n# Playbook\n%s\n\n# Evals\n",
		changePath, change.Data, specPath, spec, book)
	err = filepath.WalkDir(workspace.Path(source, target.Dir), func(full string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(source, full)
		if err != nil {
			return err
		}
		name := filepath.ToSlash(rel)
		if entry.IsDir() {
			if name != target.Dir && (path.Base(name) == ".git" || workspace.Owner(state.Snapshot, name+"/").Name != target.Name) {
				return filepath.SkipDir
			}
			return nil
		}
		if workspace.Zone(state.Snapshot, target, name) != "evals" {
			return nil
		}
		file, err := workspace.Read(source, name)
		if err != nil {
			return err
		}
		inputs[name] = file
		fmt.Fprintf(&prompt, "\n## %s\n%s\n", name, file.Data)
		return nil
	})
	if err != nil {
		return agent, false, nil, err
	}
	if err := workspace.Init(ctx, root); err != nil {
		return agent, false, nil, err
	}
	for name, file := range inputs {
		if err := file.Put(root, name); err != nil {
			return agent, false, nil, err
		}
	}
	if err := workspace.Seed(ctx, root); err != nil {
		return agent, false, nil, err
	}

	request.Snapshot.Root = root
	request.Prompt = prompt.String() + feedback
	agent, invocationErr = invoke.Agent(ctx, request)
	stepErr = importAudit(root, source, changePath, inputs)
	return agent, true, invocationErr, stepErr
}

func auditPlaybook(root string) (string, error) {
	data, err := os.ReadFile(workspace.Path(root, ".fireweed/playbooks/audit.md"))
	if os.IsNotExist(err) {
		return assets.Playbook("audit")
	}
	return string(data), err
}

func importAudit(root, heldRoot, changePath string, inputs map[string]workspace.File) error {
	reject := func(err error) error {
		return contracts.CodedError{Code: "P09", Subject: changePath, Cause: fmt.Errorf("invalid audit workspace submission: %w", err)}
	}
	var candidate string
	seen := 0
	err := workspace.Walk(root, func(name string, entry fs.DirEntry) error {
		if entry.IsDir() {
			return nil
		}
		original, ok := inputs[name]
		if !ok {
			return fmt.Errorf("added file %s", name)
		}
		file, err := workspace.Read(root, name)
		if err != nil {
			return err
		}
		seen++
		if name == changePath && file.Mode == original.Mode {
			candidate = file.Data
			return nil
		}
		if original != file {
			return fmt.Errorf("modified input %s", name)
		}
		return nil
	})
	if err != nil {
		return reject(err)
	}
	if seen != len(inputs) {
		return reject(fmt.Errorf("deleted input files"))
	}
	original := inputs[changePath]
	if candidate == original.Data {
		return nil
	}
	approved, err := contracts.ParseDocument(original.Data)
	if err != nil {
		return reject(err)
	}
	submitted, err := contracts.ParseDocument(candidate)
	if err != nil {
		return reject(err)
	}
	frontmatter := original.Data[:len(original.Data)-len(approved.Body)]
	if frontmatter != candidate[:len(candidate)-len(submitted.Body)] {
		return reject(fmt.Errorf("change frontmatter was modified"))
	}
	body, err := workflow.ApplyAuditEdit(approved.Body, submitted.Body)
	if err != nil {
		return reject(err)
	}
	current, err := workspace.Read(heldRoot, changePath)
	if err != nil {
		return err
	}
	if current != original {
		return reject(fmt.Errorf("held change was modified during audit"))
	}
	return workspace.Write(heldRoot, changePath, []byte(frontmatter+body), original.Mode.Perm())
}
