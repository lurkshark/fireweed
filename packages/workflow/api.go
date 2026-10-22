// Package workflow provides checked change lifecycle operations.
package workflow

import (
	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/internal/claims"
	"github.com/lurkshark/fireweed/packages/workflow/internal/policy"
	"github.com/lurkshark/fireweed/packages/workflow/internal/state"
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
