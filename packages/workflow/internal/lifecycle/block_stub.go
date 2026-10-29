package lifecycle

import (
	"errors"

	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func Block(request model.BlockRequest) (model.BlockResult, error) {
	return model.BlockResult{}, errors.New("Block is not implemented")
}
