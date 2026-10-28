package invoke

import (
	"context"
	"errors"

	"github.com/lurkshark/fireweed/packages/runner/model"
)

func Agent(context.Context, model.AgentRequest) (model.AgentResult, error) {
	return model.AgentResult{}, errors.New("agent invocation is not implemented")
}
