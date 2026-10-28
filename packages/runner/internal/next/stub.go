package next

import (
	"context"
	"errors"

	"github.com/lurkshark/fireweed/packages/runner/model"
)

func Run(context.Context, model.NextRequest) (model.NextResult, error) {
	return model.NextResult{}, errors.New("next is not implemented")
}
