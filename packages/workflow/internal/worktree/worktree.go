package worktree

import (
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func Check(snapshot project.Snapshot, module contracts.Module, path string, claim model.Claim) model.WorktreeCheck {
	return model.WorktreeCheck{}
}
