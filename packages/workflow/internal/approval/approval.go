package approval

import (
	"fmt"
	"path"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/internal/claims"
	"github.com/lurkshark/fireweed/packages/workflow/internal/state"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func problem(code, format string, args ...any) error {
	return fmt.Errorf("%s: %s. Check the exact queued slug and restore the approved source, or use fireweed start or release after inspecting its owner.", code, fmt.Sprintf(format, args...))
}

func Load(start, slug string) (project.Snapshot, model.Claim, error) {
	return load(start, slug, false)
}

// LoadHeld restores only committed claim authority, without an unclaimed fallback.
func LoadHeld(start, slug string) (project.Snapshot, model.Claim, error) {
	return load(start, slug, true)
}

func load(start, slug string, required bool) (project.Snapshot, model.Claim, error) {
	zero, noClaim := project.Snapshot{}, model.Claim{}
	if !claims.ValidSlug(slug) {
		return zero, noClaim, problem("C05", "invalid exact change selector %q", slug)
	}
	repo, err := gitstore.Open(start)
	if err != nil {
		return zero, noClaim, problem("C01", "cannot discover the local approval repository: %v", err)
	}
	claim, present, err := claims.Lookup(repo.Root, slug)
	if err != nil {
		return zero, noClaim, problem("P02", "cannot inspect the selected approval: %v", err)
	}
	if required && !present {
		return zero, noClaim, problem("P02", "change %s has no held claim", slug)
	}
	var source project.Snapshot
	if present {
		if claim.Worktree != repo.Root {
			return zero, noClaim, problem("P02", "claim %s belongs to worktree %q rather than %q", slug, claim.Worktree, repo.Root)
		}
		source, err = project.LoadAt(repo.Root, claim.Base)
		if err != nil {
			return zero, noClaim, problem("P02", "cannot restore claim %s baseline %s: %v", slug, claim.Base, err)
		}
	} else {
		source, err = project.Load(repo.Root)
		if err != nil {
			return zero, noClaim, problem("C01", "cannot discover current approval metadata: %v", err)
		}
	}
	index, module, err := selectChange(source, slug)
	if err != nil {
		return zero, noClaim, err
	}
	if !present {
		return source, noClaim, nil
	}
	selected := source.Changes[index]
	if ds := state.Check(module, claim.State, claim.State, selected.Path); len(ds) != 0 {
		messages := make([]string, len(ds))
		for i, d := range ds {
			messages[i] = d.Message
		}
		return zero, noClaim, problem("P02", "invalid approved state for %s: %s", slug, strings.Join(messages, "; "))
	}
	approved, err := contracts.ParseChange(claim.State, selected.Path, module.IDPrefix)
	if err != nil {
		return zero, noClaim, problem("P02", "cannot parse the approved state: %v", err)
	}
	if approved.Module != selected.Module || approved.Run != claim.Run {
		return zero, noClaim, problem("P02", "approved module/run %q/%q do not match baseline target and claim %q/%q", approved.Module, approved.Run, selected.Module, claim.Run)
	}
	// Loading is intentionally not projection or live-worktree verification.
	// The caller retains the complete committed source and applies queue effects
	// once through Project/ApplyChange, even when current feedback is malformed.
	source.Changes[index] = approved
	return source, claim, nil
}

func selectChange(source project.Snapshot, slug string) (int, contracts.Module, error) {
	index := -1
	for i, change := range source.Changes {
		if change.Slug == slug {
			if index >= 0 {
				return 0, contracts.Module{}, problem("C05", "queued selector %q is ambiguous", slug)
			}
			index = i
		}
	}
	if index < 0 {
		return 0, contracts.Module{}, problem("C05", "no queued change has exact slug %q", slug)
	}
	selected := source.Changes[index]
	var module contracts.Module
	count := 0
	for _, candidate := range source.Modules {
		if candidate.Name == selected.Module {
			module = candidate
			count++
		}
	}
	if count != 1 {
		return 0, contracts.Module{}, problem("C05", "change %s requires exactly one target module %q; found %d", slug, selected.Module, count)
	}
	if path.Dir(selected.Path) != "changes" || path.Clean(selected.Path) != selected.Path || strings.ContainsAny(selected.Path, "\\\x00") {
		return 0, contracts.Module{}, problem("C05", "change %s has invalid immediate queued path %q", slug, selected.Path)
	}
	// Reuse the public filename grammar rather than maintaining a second naming
	// rule here. Rendering supplied parsed data does not open current contents.
	raw, err := contracts.RenderMarkdown(selected.Document)
	if err != nil {
		return 0, contracts.Module{}, problem("C05", "invalid source change %s: %v", slug, err)
	}
	parsed, err := contracts.ParseChange(raw, selected.Path, module.IDPrefix)
	if err != nil || parsed.Slug != slug {
		return 0, contracts.Module{}, problem("C05", "source path %q does not identify exact queued slug %q: %v", selected.Path, slug, err)
	}
	return index, module, nil
}
