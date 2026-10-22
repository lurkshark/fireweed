package policy

import (
	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

func Check(snapshot project.Snapshot, module contracts.Module, change contracts.Change, edits []gitstore.Edit) []checks.Diagnostic {
	return nil
}
