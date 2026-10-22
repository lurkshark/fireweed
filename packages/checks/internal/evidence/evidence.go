package evidence

import (
	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
)

func Parse(report []byte) ([]model.TestCase, error) { return nil, nil }
func Analyze(module contracts.Module, prefixes []string, tests []model.TestCase, hookExit int, failFirst, already []string) model.Coverage {
	return model.Coverage{}
}
