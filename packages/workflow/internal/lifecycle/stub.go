package lifecycle

import (
	"context"
	"errors"

	"github.com/lurkshark/fireweed/packages/workflow/model"
)

var errUnimplemented = errors.New("workflow lifecycle is not implemented")

func Start(string, string) (model.State, error)   { return model.State{}, errUnimplemented }
func Current(string, string) (model.State, error) { return model.State{}, errUnimplemented }
func Verify(context.Context, model.Request) (model.Result, error) {
	return model.Result{}, errUnimplemented
}
func Done(context.Context, model.Request) (model.Result, error) {
	return model.Result{}, errUnimplemented
}
func Context(string, string, bool) (string, error) { return "", errUnimplemented }
