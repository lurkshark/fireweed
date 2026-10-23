package contentapproval

import (
	"fmt"
	"github.com/lurkshark/fireweed/packages/contracts"
)

func Matches(module contracts.Module, baseText, currentText, path string) (bool, error) {
	return false, fmt.Errorf("content approval not implemented")
}
