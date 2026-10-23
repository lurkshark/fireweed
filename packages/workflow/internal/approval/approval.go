package approval

import (
	"fmt"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func Load(start, slug string) (project.Snapshot, model.Claim, error) {
	return project.Snapshot{}, model.Claim{}, fmt.Errorf("approved snapshot restoration is not implemented")
}
