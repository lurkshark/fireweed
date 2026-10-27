package engine

import (
	"context"

	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func ParseJUnit([]byte, []contracts.Item) (model.Report, error)         { return model.Report{}, nil }
func RequireCoverage(model.Report, []contracts.Item) error              { return nil }
func RequireFailFirst(model.Report, []string, []string) error           { return nil }
func RequirePass(model.Report, []contracts.Item) error                  { return nil }
func Evaluate(context.Context, model.HookRequest) (model.Report, error) { return model.Report{}, nil }
func Boundaries(context.Context, model.BoundaryRequest) ([]project.Diagnostic, error) {
	return nil, nil
}
func Zones(model.ZoneRequest) []project.Diagnostic { return nil }
