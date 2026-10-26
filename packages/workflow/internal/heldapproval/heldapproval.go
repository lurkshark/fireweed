package heldapproval

import (
	"errors"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func Load(start string) (project.Snapshot, model.Claim, bool, error) {
	return project.Snapshot{}, model.Claim{}, false, errors.New("P02: held source is not implemented; use the approved selection contract and restore the required implementation")
}
