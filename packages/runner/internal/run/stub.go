package run

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner/internal/next"
	"github.com/lurkshark/fireweed/packages/runner/model"
)

func Execute(ctx context.Context, request model.RunRequest) (model.RunResult, error) {
	var result model.RunResult
	budget := &next.UsageBudget{}
	if request.For > 0 {
		stop := &next.TimeStopError{Limit: request.For}
		deadline := time.Now().Add(request.For)
		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadlineCause(ctx, deadline, stop)
		defer cancel()
		budget.SetTimeLimit(deadline, stop)
	}
	nextRequest := model.NextRequest{
		Root:          request.Root,
		Bin:           request.Bin,
		AllowUnmerged: request.AllowUnmerged,
		Timeout:       request.Timeout,
	}
	var blockedErr error
	for {
		change, err := next.RunWithBudget(ctx, nextRequest, budget)
		if err != nil {
			var blocked *next.BlockedError
			if errors.As(err, &blocked) {
				blockedErr = errors.Join(blockedErr, err)
				continue
			}
			if blockedErr != nil && next.IsUsageStop(err) {
				return result, &next.CombinedUsageFailure{Blocked: blockedErr, Stop: err}
			}
			if blockedErr != nil && IsTimeStop(err) {
				return result, &combinedTimeFailure{blocked: blockedErr, stop: err}
			}
			return result, errors.Join(blockedErr, err)
		}
		if change.Empty {
			queue, err := project.Queue(request.Root)
			if err != nil {
				return result, errors.Join(blockedErr, fmt.Errorf("read change queue: %w", err))
			}
			result.Empty = len(queue) == 0
			return result, blockedErr
		}
		result.Applied = append(result.Applied, change)
		if request.Max > 0 && len(result.Applied) == request.Max {
			queue, err := project.Queue(request.Root)
			if err != nil {
				return result, errors.Join(blockedErr, fmt.Errorf("read change queue: %w", err))
			}
			result.Empty = len(queue) == 0
			result.Limited = !result.Empty
			return result, blockedErr
		}
	}
}
