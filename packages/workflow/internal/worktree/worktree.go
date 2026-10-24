package worktree

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/internal/claims"
	"github.com/lurkshark/fireweed/packages/workflow/internal/policy"
	"github.com/lurkshark/fireweed/packages/workflow/internal/state"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func Check(snapshot project.Snapshot, module contracts.Module, path string, claim model.Claim) model.WorktreeCheck {
	result := model.WorktreeCheck{Edits: make([]gitstore.Edit, 0), Diagnostics: make([]checks.Diagnostic, 0)}
	add := func(code, at, message string) {
		result.Diagnostics = append(result.Diagnostics, checks.Diagnostic{Code: code, Severity: "error", Path: at, Message: message})
	}
	precondition := func(reason string) model.WorktreeCheck {
		add("P02", path, reason+". Restore the approved local claim and HEAD, or release and start the change again after checking ownership.")
		return result
	}
	repo, err := gitstore.Open(snapshot.Root)
	if err != nil {
		return precondition("Cannot discover the checkout")
	}
	stored, err := claims.Load(repo.Root, claim.Change)
	if err != nil || stored != claim || claim.Worktree != repo.Root {
		return precondition("The supplied claim is missing, stale, altered or owned by another checkout")
	}
	// Parse the retained state, never worker metadata, to establish identity and phase.
	approved, err := contracts.ParseChange(claim.State, path, module.IDPrefix)
	if err != nil || !strings.HasPrefix(path, "changes/") || strings.Count(path, "/") != 1 {
		return precondition("The approved state must name a valid queued change directly under changes/")
	}
	if approved.Slug != claim.Change || approved.Run != claim.Run || approved.Module != module.Name || len(state.Check(module, claim.State, claim.State, path)) != 0 {
		return precondition("The approved change identity, run or phase does not match its claim and selected module")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	head := gitstore.Capture(ctx, repo.Root, []string{"git", "rev-parse", "--verify", "--end-of-options", "HEAD"}, []string{"GIT_TERMINAL_PROMPT=0", "GIT_OPTIONAL_LOCKS=0", "GIT_NO_LAZY_FETCH=1", "GIT_NO_REPLACE_OBJECTS=1"})
	if head.Err != nil || head.Exit != 0 || strings.TrimSpace(head.Stdout) != claim.Base {
		return precondition("HEAD no longer equals the approved step baseline")
	}
	// Current text is useful even if collecting edits subsequently fails. Containment
	// is checked before opening it, including when the open would otherwise fail.
	current, err := gitstore.SafePath(repo.Root, path)
	if err != nil {
		add("P2-01", path, "Restore a contained current change path before comparing worker edits.")
	} else if info, err := os.Stat(current); err != nil || !info.Mode().IsRegular() {
		add("P2-01", path, "Restore a readable regular current change file before comparing worker edits.")
	} else if text, err := os.ReadFile(current); err != nil {
		add("P2-01", path, "Restore a readable current change file before comparing worker edits.")
	} else {
		result.Diagnostics = append(result.Diagnostics, state.Check(module, claim.State, string(text), path)...)
	}
	edits, err := gitstore.Edits(repo.Root, claim.Base)
	if err != nil {
		add("P02", path, "Cannot collect local edits from the approved baseline. Restore readable local Git state and retry verification.")
	} else {
		result.Edits = append(result.Edits, edits...)
		contain := func(endpoint string) {
			if _, err := gitstore.SafePath(repo.Root, endpoint); err != nil {
				add("P2-01", endpoint, fmt.Sprintf("Restore the contained edit endpoint %q; worker paths must remain inside the checkout.", endpoint))
			}
		}
		for _, edit := range edits {
			contain(edit.Path)
			if edit.Status == "R" || edit.Status == "C" {
				contain(edit.OldPath)
			}
		}
		snapshot.Root = repo.Root
		result.Diagnostics = append(result.Diagnostics, policy.Check(snapshot, module, approved, edits)...)
	}
	sort.Slice(result.Diagnostics, func(i, j int) bool {
		a, b := result.Diagnostics[i], result.Diagnostics[j]
		if a.Path != b.Path {
			return a.Path < b.Path
		}
		if a.Code != b.Code {
			return a.Code < b.Code
		}
		if a.Severity != b.Severity {
			return a.Severity < b.Severity
		}
		return a.Message < b.Message
	})
	unique := result.Diagnostics[:0]
	for _, diagnostic := range result.Diagnostics {
		if len(unique) == 0 || unique[len(unique)-1] != diagnostic {
			unique = append(unique, diagnostic)
		}
	}
	result.Diagnostics = unique
	return result
}
