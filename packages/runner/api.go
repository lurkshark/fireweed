package runner

import (
	"context"

	"github.com/lurkshark/fireweed/packages/runner/internal/drill"
	"github.com/lurkshark/fireweed/packages/runner/internal/invoke"
	"github.com/lurkshark/fireweed/packages/runner/internal/next"
	"github.com/lurkshark/fireweed/packages/runner/internal/run"
	"github.com/lurkshark/fireweed/packages/runner/model"
)

type AgentRequest = model.AgentRequest
type AgentResult = model.AgentResult
type NextRequest = model.NextRequest
type NextResult = model.NextResult
type StepResult = model.StepResult
type RunRequest = model.RunRequest
type RunResult = model.RunResult
type DrillRequest = model.DrillRequest
type DrillResult = model.DrillResult

func InvokeAgent(ctx context.Context, request AgentRequest) (AgentResult, error) {
	return invoke.Agent(ctx, request)
}

func Next(ctx context.Context, request NextRequest) (NextResult, error) {
	return next.Run(ctx, request)
}

func Run(ctx context.Context, request RunRequest) (RunResult, error) {
	return run.Execute(ctx, request)
}

func Drill(ctx context.Context, request DrillRequest) (DrillResult, error) {
	return drill.Run(ctx, request)
}

func IsUsageStop(err error) bool {
	return next.IsUsageStop(err)
}

func IsApprovalRequired(err error) bool {
	return next.IsApprovalRequired(err)
}

func IsTimeStop(err error) bool {
	return run.IsTimeStop(err)
}
