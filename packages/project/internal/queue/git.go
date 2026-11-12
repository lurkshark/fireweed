package queue

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
)

func readGit(root string, args ...string) ([]byte, error) {
	argv := append([]string{"git", "-C", root, "-c", "core.hooksPath=/dev/null", "-c", "core.fsmonitor=false"}, args...)
	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Args: argv, Timeout: 30 * time.Second,
		Env: map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0"},
	})
	if err != nil {
		return nil, fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, result.Stderr)
	}
	return []byte(result.Stdout), nil
}
