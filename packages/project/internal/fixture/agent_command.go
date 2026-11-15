package fixture

import (
	"fmt"
	"strings"

	"github.com/lurkshark/fireweed/packages/project/internal/source"
)

func ResolveAgentCommand(config map[string]any, role, alias string) (any, bool, error) {
	agent, _ := config["agent"].(map[string]any)
	if role != "audit" && alias != "" {
		command, err := namedAgentCommand(agent, alias, "")
		return command, false, err
	}

	setting, reference := "agent.default", agent["default"]
	if role == "audit" {
		if raw, present := config["audit"]; present {
			audit, ok := raw.(map[string]any)
			if !ok {
				return nil, false, source.Failure("C01", "audit", fmt.Errorf("audit must be a mapping to resolve audit.agent"))
			}
			if value, configured := audit["agent"]; configured {
				setting, reference = "audit.agent", value
			}
		}
	}
	alias, ok := reference.(string)
	if !ok || alias == "" {
		return nil, false, source.Failure("C01", setting, fmt.Errorf("%s must name an agent command with a nonempty string", setting))
	}
	command, err := namedAgentCommand(agent, alias, setting)
	return command, err == nil && role == "audit" && setting == "agent.default", err
}

func namedAgentCommand(agent map[string]any, alias, setting string) (any, error) {
	commands, ok := agent["commands"].(map[string]any)
	if !ok {
		return nil, source.Failure("C01", "agent.commands", fmt.Errorf("agent.commands must be a mapping"))
	}
	command, exists := commands[alias]
	if !exists {
		if setting == "" {
			return nil, source.Failure("U02", alias, fmt.Errorf("unknown agent command alias %q", alias))
		}
		return nil, source.Failure("C01", setting, fmt.Errorf("%s names unlisted agent command %q", setting, alias))
	}
	if !validAgentCommand(command) {
		commandSetting := fmt.Sprintf("agent.commands[%q]", alias)
		if setting != "" {
			commandSetting = fmt.Sprintf("%s references %s, which", setting, commandSetting)
		} else {
			setting = commandSetting
		}
		return nil, source.Failure("C01", setting, fmt.Errorf("%s must be a nonempty shell template or a string list with a nonempty executable", commandSetting))
	}
	return command, nil
}

func validAgentCommand(command any) bool {
	switch command := command.(type) {
	case string:
		return strings.TrimSpace(command) != ""
	case []string:
		return len(command) > 0 && command[0] != ""
	case []any:
		argv, err := stringList("agent command", command)
		return err == nil && len(argv) > 0 && argv[0] != ""
	default:
		return false
	}
}
