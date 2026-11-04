package lifecycle

import (
	"context"
	"encoding/json"
	"os"
	"strings"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

// GC has a deliberately narrow temporary implementation for evals-phase
// compilation. The implementation step replaces it after fail-first evidence.
func GC(ctx context.Context, request model.GCRequest) (model.GCResult, error) {
	repo, err := gitstore.Inspect(request.Root)
	if err != nil {
		return model.GCResult{}, err
	}
	result, err := gitstore.Run(ctx, gitstore.Request{Args: []string{"git", "-C", repo.Root, "for-each-ref", "--format=%(refname) %(objectname)", "refs/fireweed/claims"}})
	if err != nil {
		return model.GCResult{}, err
	}
	for _, line := range strings.Split(strings.TrimSpace(result.Stdout), "\n") {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}
		data, err := gitstore.ReadBlob(repo.Root, fields[1])
		if err != nil {
			continue
		}
		var recorded struct {
			Worktree string `json:"worktree"`
		}
		if json.Unmarshal(data, &recorded) != nil || recorded.Worktree == "" {
			continue
		}
		if _, err := os.Stat(recorded.Worktree); os.IsNotExist(err) {
			if err := gitstore.CASRef(repo.Root, fields[0], fields[1], ""); err != nil {
				return model.GCResult{}, err
			}
		}
	}
	return model.GCResult{}, nil
}
