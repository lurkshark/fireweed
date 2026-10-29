package workflow

import (
	"context"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/internal/lifecycle"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

type State = model.State
type Request = model.Request
type Result = model.Result
type BlockRequest = model.BlockRequest
type BlockResult = model.BlockResult

func Start(root, change string) (State, error)       { return lifecycle.Start(root, change) }
func Eligible(root string) ([]project.Change, error) { return lifecycle.Eligible(root) }
func Current(root, change string) (State, error)     { return lifecycle.Current(root, change) }
func Verify(ctx context.Context, request Request) (Result, error) {
	return lifecycle.Verify(ctx, request)
}
func Done(ctx context.Context, request Request) (Result, error) {
	return lifecycle.Done(ctx, request)
}
func Block(request BlockRequest) (BlockResult, error) {
	return lifecycle.Block(request)
}
func Context(root, change string, full bool) (string, error) {
	return lifecycle.Context(root, change, full)
}
