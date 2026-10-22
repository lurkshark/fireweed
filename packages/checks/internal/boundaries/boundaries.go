package boundaries

import (
	"context"
	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/project"
)

func Check(ctx context.Context, snapshot project.Snapshot) []model.Diagnostic { return nil }
