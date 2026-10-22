package allocation

import (
	"fmt"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/inspection"
)

func Allocate(s inspection.Data, slug string) (contracts.Change, map[string]string, error) {
	return contracts.Change{}, nil, fmt.Errorf("allocation not implemented")
}
