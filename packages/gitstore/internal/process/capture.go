package process

import (
	"context"
	"errors"
)

func Capture(ctx context.Context, dir string, argv, overrides []string) (string, string, int, error) {
	return "", "", 2, errors.New("separate process capture is not implemented")
}
