package run

import (
	"context"
	"errors"
	"fmt"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner/internal/next"
	"github.com/lurkshark/fireweed/packages/runner/model"
)

func Execute(ctx context.Context, request model.RunRequest) (model.RunResult, error) {
	var result model.RunResult
	nextRequest := model.NextRequest{
		Root:          request.Root,
		Bin:           request.Bin,
		AllowUnmerged: request.AllowUnmerged,
		Timeout:       request.Timeout,
	}
	var blockedErr error
	for {
		change, err := next.Run(ctx, nextRequest)
		if err != nil {
			var blocked *next.BlockedError
			if errors.As(err, &blocked) {
				blockedErr = errors.Join(blockedErr, err)
				continue
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
