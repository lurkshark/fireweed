package engine

import (
	"context"

	"github.com/lurkshark/fireweed/packages/checks/model"
)

func Doctor(context.Context, model.DoctorRequest) model.DoctorReport { return model.DoctorReport{} }
