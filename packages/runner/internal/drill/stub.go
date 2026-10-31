package drill

import (
	"context"
	"fmt"

	"github.com/lurkshark/fireweed/packages/runner/model"
)

func Run(_ context.Context, _ model.DrillRequest) (model.DrillResult, error) {
	return model.DrillResult{}, fmt.Errorf("drill is not implemented")
}
