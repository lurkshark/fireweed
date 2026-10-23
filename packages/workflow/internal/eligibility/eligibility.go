package eligibility

import (
	"fmt"

	"github.com/lurkshark/fireweed/packages/project"
)

func Check(snapshot project.Snapshot, slug string) (bool, error) {
	return false, fmt.Errorf("queue eligibility not implemented")
}
