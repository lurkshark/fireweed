// Package checks evaluates recorded test evidence and repository enforcement.
package checks

import (
	"context"
	"github.com/lurkshark/fireweed/packages/checks/internal/boundaries"
	"github.com/lurkshark/fireweed/packages/checks/internal/evidence"
	"github.com/lurkshark/fireweed/packages/checks/internal/hooks"
	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

type Message = model.Message
type TestCase = model.TestCase
type Diagnostic = model.Diagnostic
type Coverage = model.Coverage

func ParseJUnit(report []byte) ([]TestCase, error) { return evidence.Parse(report) }
func Analyze(module contracts.Module, prefixes []string, tests []TestCase, hookExit int, failFirst, already []string) Coverage {
	return evidence.Analyze(module, prefixes, tests, hookExit, failFirst, already)
}

func RunHook(ctx context.Context, snapshot project.Snapshot, module *contracts.Module, command, report string) gitstore.CommandResult {
	return hooks.Run(ctx, snapshot, module, command, report)
}
func Evaluate(ctx context.Context, snapshot project.Snapshot, module contracts.Module, failFirst, already []string) Coverage {
	return hooks.Evaluate(ctx, snapshot, module, failFirst, already)
}

func Boundaries(ctx context.Context, snapshot project.Snapshot) []Diagnostic {
	return boundaries.Check(ctx, snapshot)
}
