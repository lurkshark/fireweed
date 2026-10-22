// Package checks evaluates recorded test evidence and repository enforcement.
package checks

import (
	"github.com/lurkshark/fireweed/packages/checks/internal/evidence"
	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
)

type Message = model.Message
type TestCase = model.TestCase
type Diagnostic = model.Diagnostic
type Coverage = model.Coverage

func ParseJUnit(report []byte) ([]TestCase, error) { return evidence.Parse(report) }
func Analyze(module contracts.Module, prefixes []string, tests []TestCase, hookExit int, failFirst, already []string) Coverage {
	return evidence.Analyze(module, prefixes, tests, hookExit, failFirst, already)
}
