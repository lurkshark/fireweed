package workercontext

import (
	"context"
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/internal/claims"
	"github.com/lurkshark/fireweed/packages/workflow/internal/state"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

// Build consumes a captured approved snapshot. Source text is feedback, never a
// replacement for its permission metadata. Failed assembly never exposes a
// partially populated prompt.
func Build(snapshot project.Snapshot, slug string) (out model.WorkerContext, err error) {
	defer func() {
		if err != nil {
			out = model.WorkerContext{}
		} else {
			out = owned(out)
		}
	}()
	var selected contracts.Change
	matches := 0
	for _, change := range snapshot.Changes {
		if change.Slug == slug {
			selected = change
			matches++
		}
	}
	if matches != 1 {
		return out, invalid("select exactly one queued change by its complete slug")
	}
	if selected.Phase != "" && selected.Phase != "evals" && selected.Phase != "audit" && selected.Phase != "impl" {
		return out, invalid("restore an empty, evals, audit or impl phase")
	}
	var baseModule contracts.Module
	found := false
	for _, module := range snapshot.Modules {
		if module.Name == selected.Module {
			baseModule = module
			found = true
			break
		}
	}
	if !found {
		return out, invalid("restore the selected change's target module")
	}
	if path.Dir(selected.Path) != "changes" || strings.ContainsAny(selected.Path, "\\\x00") {
		return out, invalid("restore a direct changes/<number>-<slug>.md queued path")
	}
	parsedPath, parseErr := contracts.ParseChange("---\nmodule: context\n---\n\n## Motivation\n\nValidate the queued filename.\n", selected.Path, baseModule.IDPrefix)
	if parseErr != nil || parsedPath.Slug != selected.Slug {
		return out, invalid("restore the queued filename to match its selected slug")
	}
	before, projectErr := project.Project(snapshot, slug)
	if projectErr != nil {
		return out, invalid("review preceding queue projection: %v", projectErr)
	}
	for _, module := range before.Modules {
		if module.Name == selected.Module {
			baseModule = module
			break
		}
	}
	target, projectErr := project.ApplyChange(baseModule, selected)
	if projectErr != nil {
		return out, invalid("review the selected projected contract: %v", projectErr)
	}
	out.Change = selected
	out.Module = target
	out.Dependencies = make([]contracts.Module, 0, len(target.DependsOn))
	dependencyNames := map[string]bool{}
	for _, name := range target.DependsOn {
		present := false
		for _, module := range before.Modules {
			if module.Name == name {
				out.Dependencies = append(out.Dependencies, module)
				present = true
				break
			}
		}
		if !present {
			return out, invalid("restore the declared direct dependency %q", name)
		}
		dependencyNames[name] = true
	}
	repo, openErr := gitstore.Open(snapshot.Root)
	if openErr != nil {
		return out, readError(snapshot.Root, openErr)
	}
	out.GitCommonDir = repo.CommonDir
	text, missing, readErr := readSelected(repo.Root, selected.Path)
	if readErr != nil {
		if !missing {
			return out, readError(selected.Path, readErr)
		}
		text, readErr = recoverChange(repo.Root, target, selected)
		if readErr != nil {
			return out, readError(selected.Path, readErr)
		}
	}
	out.ChangeText = text
	specPath := path.Join(target.Dir, "SPEC.md")
	out.SpecText, _, readErr = readSelected(repo.Root, specPath)
	if readErr != nil {
		return out, readError(specPath, readErr)
	}
	out.System, missing, readErr = readSelected(repo.Root, "SYSTEM.md")
	if readErr != nil && !missing {
		return out, readError("SYSTEM.md", readErr)
	}
	bookName := selected.Phase
	if bookName == "" {
		bookName = "author-change"
	}
	if selected.Regenerate {
		bookName = "regenerate"
	}
	book, bookErr := project.ResolvePlaybook(repo.Root, bookName)
	if bookErr != nil {
		if strings.HasPrefix(bookErr.Error(), "D05:") {
			return out, bookErr
		}
		return out, readError(".fireweed/playbooks/"+bookName+".md", bookErr)
	}
	out.Playbook, out.PlaybookHash = book.Text, book.Hash
	inventoryContext, cancelInventory := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelInventory()
	inventory := gitstore.Run(inventoryContext, repo.Root, []string{"git", "--no-optional-locks", "-c", "core.fsmonitor=false", "-c", "core.hooksPath=" + os.DevNull, "ls-files", "-z", "--cached", "--others", "--exclude-standard"}, []string{"GIT_OPTIONAL_LOCKS=0", "GIT_TERMINAL_PROMPT=0"})
	if inventory.Err != nil || inventory.Exit != 0 {
		return out, readError(repo.Root, fmt.Errorf("Git file enumeration failed (exit %d): %v %s", inventory.Exit, inventory.Err, inventory.Output))
	}
	if inventory.Output != "" && !strings.HasSuffix(inventory.Output, "\x00") {
		return out, readError(repo.Root, fmt.Errorf("Git file enumeration is not a complete NUL-delimited inventory"))
	}
	paths := map[string]bool{}
	for _, file := range strings.Split(inventory.Output, "\x00") {
		if file != "" {
			paths[file] = true
		}
	}
	names := make([]string, 0, len(paths))
	for file := range paths {
		names = append(names, file)
	}
	sort.Strings(names)
	out.Files = make([]model.FileContent, 0)
	for _, file := range names {
		include, selectionErr := selectFile(file, before.Modules, target, dependencyNames, snapshot.Config.Protected, selected.Phase == "audit")
		if selectionErr != nil {
			return out, readError(file, selectionErr)
		}
		if !include {
			continue
		}
		content, _, fileErr := readSelected(repo.Root, file)
		if fileErr != nil {
			return out, readError(file, fileErr)
		}
		out.Files = append(out.Files, model.FileContent{Path: file, Content: content})
	}
	return out, nil
}

func invalid(message string, args ...any) error {
	return fmt.Errorf("C05: check the approved queue context: %s", fmt.Sprintf(message, args...))
}
func readError(file string, err error) error {
	return fmt.Errorf("C01: restore readable contained context input %s and review local Git configuration: %w", file, err)
}

// missing means a safe ordinary path was absent at stat time. A containment
// failure (including a dangling link) is never eligible for fallback.
func readSelected(root, relative string) (text string, missing bool, err error) {
	full, err := gitstore.SafePath(root, relative)
	if err != nil {
		return "", false, err
	}
	info, err := os.Stat(full)
	if err != nil {
		return "", os.IsNotExist(err), err
	}
	if !info.Mode().IsRegular() {
		return "", false, fmt.Errorf("restore a regular file at %s", relative)
	}
	raw, err := os.ReadFile(full)
	if err != nil {
		return "", false, err
	}
	return string(raw), false, nil
}

func recoverChange(root string, module contracts.Module, selected contracts.Change) (string, error) {
	if strings.TrimSpace(selected.Run) == "" || selected.Phase == "" {
		return "", fmt.Errorf("restore the missing queued change; only a started change can recover from its owned claim")
	}
	claim, err := claims.Load(root, selected.Slug)
	if err != nil {
		return "", fmt.Errorf("restore the missing current change or recover its matching owned claim: %w", err)
	}
	if claim.Worktree != root || claim.Change != selected.Slug || claim.Run != selected.Run {
		return "", fmt.Errorf("restore the missing current change or recover its matching owned worktree/run claim")
	}
	change, err := contracts.ParseChange(claim.State, selected.Path, module.IDPrefix)
	if err != nil || change.Module != selected.Module || change.Run != selected.Run || change.Phase != selected.Phase || change.Slug != selected.Slug || len(state.Check(module, claim.State, claim.State, selected.Path)) != 0 {
		return "", fmt.Errorf("restore the missing current change; the stored claim must contain its valid matching approved target, run and phase")
	}
	return claim.State, nil
}

func selectFile(file string, modules []contracts.Module, target contracts.Module, deps map[string]bool, protected []string, audit bool) (bool, error) {
	for _, module := range modules {
		if module.Name == target.Name {
			module = target
		}
		zone, err := contracts.ZoneFor(file, module.Dir, module.Zones, protected)
		if err != nil {
			return false, err
		}
		if zone == contracts.Protected || zone == contracts.SpecZone {
			return false, nil
		}
		if zone == contracts.OutsideZone {
			continue
		}
		if module.Name == target.Name {
			return zone == contracts.EvalZone || zone == contracts.InterfaceZone || (!audit && (zone == contracts.PreserveZone || zone == contracts.AppendOnlyZone)), nil
		}
		return !audit && deps[module.Name] && zone == contracts.InterfaceZone, nil
	}
	return false, nil
}
