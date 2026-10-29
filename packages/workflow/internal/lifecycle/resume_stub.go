package lifecycle

import (
	"errors"

	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func StartFromAttempt(root, change string, attempt int) (model.State, error) {
	return model.State{}, errors.New("start from attempt is not implemented")
}
