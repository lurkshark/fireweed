package verification

import (
	"context"
	"fmt"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func Verify(ctx context.Context, start, slug string) (model.Verification, error) {
	return model.Verification{}, fmt.Errorf("checked verification is not implemented")
}
