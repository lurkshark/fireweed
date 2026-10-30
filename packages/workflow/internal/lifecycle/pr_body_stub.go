package lifecycle

import (
	"errors"

	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func PRBody(root, revisionRange string) (model.PRBodyReport, error) {
	return model.PRBodyReport{}, errors.New("PRBody is not implemented")
}
