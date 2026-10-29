package run

import (
	"context"
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
	for {
		change, err := next.Run(ctx, nextRequest)
		if err != nil {
			return result, err
		}
		if change.Empty {
			result.Empty = true
			return result, nil
		}
		result.Applied = append(result.Applied, change)
		if request.Max > 0 && len(result.Applied) == request.Max {
			queue, err := project.Queue(request.Root)
			if err != nil {
				return result, fmt.Errorf("read change queue: %w", err)
			}
			result.Empty = len(queue) == 0
			result.Limited = !result.Empty
			return result, nil
		}
	}
}
