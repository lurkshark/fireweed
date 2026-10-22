package finalization

import (
	"fmt"
	"github.com/lurkshark/fireweed/packages/contracts"
)

func Finalize(module contracts.Module, change contracts.Change) (contracts.Module, error) {
	return contracts.Module{}, fmt.Errorf("finalization not implemented")
}
