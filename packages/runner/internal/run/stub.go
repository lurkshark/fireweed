package run

import (
	"context"

	"github.com/lurkshark/fireweed/packages/runner/model"
)

// Execute is a disposable compiling stub. Queue behavior is implemented in impl.
func Execute(context.Context, model.RunRequest) (model.RunResult, error) {
	return model.RunResult{}, nil
}
