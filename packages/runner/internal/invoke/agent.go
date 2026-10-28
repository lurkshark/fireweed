package invoke

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/runner/model"
)

const defaultTimeout = 90 * time.Minute

func Agent(ctx context.Context, request model.AgentRequest) (model.AgentResult, error) {
	var result model.AgentResult
	agent := section(request.Snapshot.Config, "agent")
	command := agent["command"]
	switch request.Role {
	case "worker":
	case "audit":
		if auditCommand := section(request.Snapshot.Config, "audit")["command"]; !unset(auditCommand) {
			command = auditCommand
		}
	default:
		return result, fmt.Errorf("invalid agent role %q", request.Role)
	}
	if unset(command) {
		return result, fmt.Errorf("agent command is not configured")
	}

	repo, err := gitstore.Inspect(request.Snapshot.Root)
	if err != nil {
		return result, fmt.Errorf("inspect agent repository: %w", err)
	}
	argv, err := commandArgs(command, request.Prompt, repo.CommonDir)
	if err != nil {
		return result, err
	}
	wrap, err := argvValue(agent["wrap"], request.Prompt, repo.CommonDir)
	if err != nil {
		return result, fmt.Errorf("agent wrap: %w", err)
	}
	argv = append(wrap, argv...)
	if len(argv) == 0 || argv[0] == "" {
		return result, fmt.Errorf("agent command is empty")
	}

	timeout := request.Timeout
	if timeout <= 0 {
		timeout, err = configTimeout(request.Snapshot.Config["timeout"])
		if err != nil {
			return result, err
		}
	}
	process, runErr := gitstore.Run(ctx, gitstore.Request{
		Args: argv, Dir: request.Snapshot.Root,
		Env: map[string]string{"FIREWEED_BIN": request.Bin}, Timeout: timeout,
	})
	result = model.AgentResult{Stdout: process.Stdout, Stderr: process.Stderr, ExitCode: process.ExitCode}
	if process.ExitCode != 0 {
		patterns, err := stringValues(request.Snapshot.Config["limit_patterns"])
		if err != nil {
			return result, fmt.Errorf("limit_patterns: %w", err)
		}
		result.UsageLimit = matchesLimit(process.Stdout, process.Stderr, patterns)
	}
	return result, runErr
}

func section(config map[string]any, key string) map[string]any {
	value, _ := config[key].(map[string]any)
	return value
}

func unset(value any) bool {
	if value == nil {
		return true
	}
	switch typed := value.(type) {
	case string:
		return typed == ""
	case []any:
		return len(typed) == 0
	case []string:
		return len(typed) == 0
	}
	return false
}

func commandArgs(value any, prompt, commonDir string) ([]string, error) {
	if shell, ok := value.(string); ok {
		return []string{"/bin/sh", "-c", expand(shell, shellQuote(prompt), shellQuote(commonDir))}, nil
	}
	return argvValue(value, prompt, commonDir)
}

func argvValue(value any, prompt, commonDir string) ([]string, error) {
	parts, err := stringValues(value)
	if err != nil {
		return nil, err
	}
	args := make([]string, len(parts))
	for i, part := range parts {
		args[i] = expand(part, prompt, commonDir)
	}
	return args, nil
}

func stringValues(value any) ([]string, error) {
	switch typed := value.(type) {
	case nil:
		return nil, nil
	case []string:
		return typed, nil
	case []any:
		values := make([]string, len(typed))
		for i, item := range typed {
			text, ok := item.(string)
			if !ok {
				return nil, fmt.Errorf("entry %d is not a string", i)
			}
			values[i] = text
		}
		return values, nil
	default:
		return nil, fmt.Errorf("expected an argv or string list, got %T", value)
	}
}

func expand(template, prompt, commonDir string) string {
	return strings.NewReplacer("{prompt}", prompt, "{git_common_dir}", commonDir).Replace(template)
}

func shellQuote(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "'\\''") + "'"
}

func configTimeout(value any) (time.Duration, error) {
	if value == nil || value == "" {
		return defaultTimeout, nil
	}
	text, ok := value.(string)
	if !ok {
		return 0, fmt.Errorf("timeout must be a duration string, got %T", value)
	}
	duration, err := time.ParseDuration(text)
	if err != nil {
		return 0, fmt.Errorf("invalid timeout %q: %w", text, err)
	}
	if duration <= 0 {
		return defaultTimeout, nil
	}
	return duration, nil
}

func matchesLimit(stdout, stderr string, patterns []string) bool {
	stdout = strings.TrimRight(stdout, "\r\n")
	stderr = strings.TrimRight(stderr, "\r\n")
	combined := stdout
	if stderr != "" {
		if combined != "" {
			combined += "\n"
		}
		combined += stderr
	}
	if combined == "" {
		return false
	}
	lines := strings.Split(combined, "\n")
	if len(lines) > 20 {
		lines = lines[len(lines)-20:]
	}
	for _, line := range lines {
		line = strings.ToLower(line)
		for _, pattern := range patterns {
			if pattern != "" && strings.Contains(line, strings.ToLower(pattern)) {
				return true
			}
		}
	}
	return false
}
