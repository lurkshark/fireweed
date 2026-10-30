package lifecycle

import (
	"errors"

	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func Status(root string) (model.StatusReport, error) {
	return model.StatusReport{}, errors.New("Status not implemented")
}
