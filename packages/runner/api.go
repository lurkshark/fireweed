package runner

import (
	"context"

	"github.com/lurkshark/fireweed/packages/runner/internal/invoke"
	"github.com/lurkshark/fireweed/packages/runner/internal/next"
	"github.com/lurkshark/fireweed/packages/runner/model"
)

type AgentRequest = model.AgentRequest
type AgentResult = model.AgentResult
type NextRequest = model.NextRequest
type NextResult = model.NextResult
type StepResult = model.StepResult

func InvokeAgent(ctx context.Context, request AgentRequest) (AgentResult, error) {
	return invoke.Agent(ctx, request)
}

func Next(ctx context.Context, request NextRequest) (NextResult, error) {
	return next.Run(ctx, request)
}
