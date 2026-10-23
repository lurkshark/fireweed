// Package workflow provides checked change lifecycle operations.
package workflow

import (
	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/internal/approval"
	"github.com/lurkshark/fireweed/packages/workflow/internal/claims"
	"github.com/lurkshark/fireweed/packages/workflow/internal/contentapproval"
	workercontext "github.com/lurkshark/fireweed/packages/workflow/internal/context"
	"github.com/lurkshark/fireweed/packages/workflow/internal/eligibility"
	"github.com/lurkshark/fireweed/packages/workflow/internal/policy"
	manualstart "github.com/lurkshark/fireweed/packages/workflow/internal/start"
	"github.com/lurkshark/fireweed/packages/workflow/internal/state"
	"github.com/lurkshark/fireweed/packages/workflow/internal/worktree"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func CheckState(module contracts.Module, approvedText, currentText, path string) []checks.Diagnostic {
	return state.Check(module, approvedText, currentText, path)
}

func CheckEdits(snapshot project.Snapshot, approvedModule contracts.Module, approvedChange contracts.Change, edits []gitstore.Edit) []checks.Diagnostic {
	return policy.Check(snapshot, approvedModule, approvedChange, edits)
}

type Claim = model.Claim

func CreateClaim(start, slug, run, base, state string) (Claim, error) {
	return claims.Create(start, slug, run, base, state)
}
func LoadClaim(start, slug string) (Claim, error) { return claims.Load(start, slug) }
func Claims(start string) ([]Claim, error)        { return claims.List(start) }
func UpdateClaim(start string, previous Claim, base, state string) (Claim, error) {
	return claims.Update(start, previous, base, state)
}
func ReleaseClaim(start, slug, expectedObject string) error {
	return claims.Release(start, slug, expectedObject)
}

type WorktreeCheck = model.WorktreeCheck

func CheckWorktree(snapshot project.Snapshot, approvedModule contracts.Module, changePath string, claim Claim) WorktreeCheck {
	return worktree.Check(snapshot, approvedModule, changePath, claim)
}

type FileContent = model.FileContent
type WorkerContext = model.WorkerContext

func BuildContext(snapshot project.Snapshot, slug string) (WorkerContext, error) {
	return workercontext.Build(snapshot, slug)
}

func LoadApproval(start, slug string) (project.Snapshot, Claim, error) {
	return approval.Load(start, slug)
}

func Eligible(snapshot project.Snapshot, slug string) (bool, error) {
	return eligibility.Check(snapshot, slug)
}

func ApprovalMatches(module contracts.Module, baseText, currentText, path string) (bool, error) {
	return contentapproval.Matches(module, baseText, currentText, path)
}

type StartResult = model.StartResult

func Start(start, slug string) (StartResult, error) { return manualstart.Start(start, slug) }
