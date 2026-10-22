package hooks

import (
	"context"
	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

func Run(ctx context.Context, snapshot project.Snapshot, module *contracts.Module, command, report string) gitstore.CommandResult {
	return gitstore.CommandResult{}
}
func Evaluate(ctx context.Context, snapshot project.Snapshot, module contracts.Module, failFirst, already []string) model.Coverage {
	return model.Coverage{}
}
