package local

import (
	"context"
	"errors"

	"github.com/lurkshark/fireweed/packages/gitstore/model"
)

var errUnimplemented = errors.New("gitstore implementation pending")

func Inspect(string) (model.Repository, error)        { return model.Repository{}, errUnimplemented }
func BaseRef(string, string, string) (string, error)  { return "", errUnimplemented }
func Changes(string, string) ([]model.Change, error)  { return nil, errUnimplemented }
func ReadAt(string, string, string) ([]byte, error)   { return nil, errUnimplemented }
func Commit(string, string, []string) (string, error) { return "", errUnimplemented }
func WriteBlob(string, []byte) (string, error)        { return "", errUnimplemented }
func ReadBlob(string, string) ([]byte, error)         { return nil, errUnimplemented }
func Ref(string, string) (string, bool, error)        { return "", false, errUnimplemented }
func CASRef(string, string, string, string) error     { return errUnimplemented }
func Run(context.Context, model.Request) (model.Result, error) {
	return model.Result{}, errUnimplemented
}
