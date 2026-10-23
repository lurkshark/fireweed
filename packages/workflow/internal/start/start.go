package start

import (
	"fmt"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func Start(start, slug string) (model.StartResult, error) {
	return model.StartResult{}, fmt.Errorf("manual start is not implemented")
}
