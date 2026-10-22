package process

import (
	"bytes"
	"context"
)

func Capture(ctx context.Context, dir string, argv, overrides []string) (string, string, int, error) {
	var stdout, stderr bytes.Buffer
	code, err := execute(ctx, dir, argv, overrides, &stdout, &stderr)
	return stdout.String(), stderr.String(), code, err
}
