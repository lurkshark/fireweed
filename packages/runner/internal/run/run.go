package run

import (
	"context"
	"errors"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner/internal/next"
	"github.com/lurkshark/fireweed/packages/runner/internal/policy"
	"github.com/lurkshark/fireweed/packages/runner/model"
)

func IsTimeStop(err error) bool { return policy.IsTimeStop(err) }

func Execute(ctx context.Context, request model.RunRequest) (model.RunResult, error) {
	var result model.RunResult
	var budget policy.Budget
	ctx, cancel := budget.Bound(ctx, request.For)
	defer cancel()

	var blocked error
	for request.Max <= 0 || len(result.Applied) < request.Max {
		step, err := next.WithBudget(ctx, model.NextRequest{
			Root: request.Root, Bin: request.Bin, AllowUnmerged: request.AllowUnmerged, Timeout: request.Timeout,
		}, &budget)
		if err != nil {
			var saved next.Blocked
			if errors.As(err, &saved) {
				blocked = policy.WithSecondary(blocked, err)
				continue
			}
			if blocked != nil {
				return result, policy.WithSecondary(blocked, err)
			}
			return result, err
		}
		if step.Empty {
			break
		}
		result.Applied = append(result.Applied, step)
	}
	if ctx.Err() != nil {
		if blocked != nil {
			return result, policy.WithSecondary(blocked, context.Cause(ctx))
		}
		return result, context.Cause(ctx)
	}
	if blocked != nil {
		return result, blocked
	}
	queue, err := project.Queue(request.Root)
	if err != nil {
		return result, err
	}
	result.Empty = len(queue) == 0
	result.Limited = !result.Empty && request.Max > 0 && len(result.Applied) == request.Max
	return result, nil
}
