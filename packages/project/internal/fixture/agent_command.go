package fixture

import "errors"

func ResolveAgentCommand(config map[string]any, role, alias string) (any, bool, error) {
	return nil, false, errors.New("agent command resolution is not implemented")
}
