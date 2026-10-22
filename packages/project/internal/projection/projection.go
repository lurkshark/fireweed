package projection

import (
	"fmt"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/inspection"
)

func Apply(m contracts.Module, ch contracts.Change) (contracts.Module, error) {
	return contracts.Module{}, fmt.Errorf("application not implemented")
}
func Project(s inspection.Data, before string) (inspection.Data, error) {
	return inspection.Data{}, fmt.Errorf("projection not implemented")
}
func Bases(s inspection.Data, slug string) (map[string]string, error) {
	return nil, fmt.Errorf("bases not implemented")
}
