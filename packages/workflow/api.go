// Package workflow provides checked change lifecycle operations.
package workflow

import (
	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/internal/policy"
	"github.com/lurkshark/fireweed/packages/workflow/internal/state"
)

func CheckState(module contracts.Module, approvedText, currentText, path string) []checks.Diagnostic {
	return state.Check(module, approvedText, currentText, path)
}

func CheckEdits(snapshot project.Snapshot, approvedModule contracts.Module, approvedChange contracts.Change, edits []gitstore.Edit) []checks.Diagnostic {
	return policy.Check(snapshot, approvedModule, approvedChange, edits)
}
