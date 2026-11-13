package invoke

import (
	"fmt"
	"path/filepath"
	"strings"
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

func Command(config map[string]any, role string) any {
	if role == "audit" {
		if command := Mapping(config["audit"])["command"]; command != nil {
			return command
		}
	}
	return Mapping(config["agent"])["command"]
}

func Name(config map[string]any, role string) string {
	command := Command(config, role)
	var args []string
	if shell, ok := command.(string); ok {
		args = strings.Fields(shell)
	} else {
		args, _ = Strings(command)
	}
	if len(args) > 0 {
		return filepath.Base(args[0])
	}
	return "unknown"
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
