package invoke

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/lurkshark/fireweed/packages/project"
)

func Mapping(value any) map[string]any {
	mapping, _ := value.(map[string]any)
	return mapping
}

func Strings(value any) ([]string, error) {
	switch value := value.(type) {
	case nil:
		return nil, nil
	case []string:
		return value, nil
	case []any:
		args := make([]string, 0, len(value))
		for _, item := range value {
			arg, ok := item.(string)
			if !ok {
				return nil, fmt.Errorf("expected a string in argv or string list")
			}
			args = append(args, arg)
		}
		return args, nil
	default:
		return nil, fmt.Errorf("expected a string list, got %T", value)
	}
}

func Name(config map[string]any, role string) string {
	command, _, _ := project.ResolveAgentCommand(config, role, "")
	return CommandName(command)
}

func CommandName(command any) string {
	if args := commandWords(command); len(args) > 0 {
		return filepath.Base(args[0])
	}
	return "unknown"
}

func Model(command any) string {
	args := commandWords(command)
	for i, arg := range args {
		switch {
		case (arg == "-m" || arg == "--model") && i+1 < len(args):
			return args[i+1]
		case strings.HasPrefix(arg, "--model="):
			return strings.TrimPrefix(arg, "--model=")
		}
	}
	return "unknown"
}

func commandWords(command any) []string {
	if shell, ok := command.(string); ok {
		return strings.Fields(shell)
	}
	args, _ := Strings(command)
	return args
}

func Quote(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "'\"'\"'") + "'"
}

func Args(command any, tokens ...string) ([]string, error) {
	if shell, ok := command.(string); ok {
		if strings.TrimSpace(shell) == "" {
			return nil, fmt.Errorf("empty shell command")
		}
		quoted := append([]string(nil), tokens...)
		for i := 1; i < len(quoted); i += 2 {
			quoted[i] = Quote(quoted[i])
		}
		return []string{"sh", "-c", strings.NewReplacer(quoted...).Replace(shell)}, nil
	}
	args, err := Strings(command)
	if err != nil {
		return nil, err
	}
	if len(args) == 0 || args[0] == "" {
		return nil, fmt.Errorf("empty command argv")
	}
	return expand(args, strings.NewReplacer(tokens...)), nil
}

func expand(args []string, replacer *strings.Replacer) []string {
	result := make([]string, len(args))
	for i, arg := range args {
		result[i] = replacer.Replace(arg)
	}
	return result
}
