package workercontext

import (
	"errors"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func Build(snapshot project.Snapshot, slug string) (model.WorkerContext, error) {
	return model.WorkerContext{}, errors.New("worker context is not implemented")
}
