package runner

import (
	"context"

	"github.com/lurkshark/fireweed/packages/runner/internal/invoke"
	"github.com/lurkshark/fireweed/packages/runner/model"
)

type AgentRequest = model.AgentRequest
type AgentResult = model.AgentResult

func InvokeAgent(ctx context.Context, request AgentRequest) (AgentResult, error) {
	return invoke.Agent(ctx, request)
}
