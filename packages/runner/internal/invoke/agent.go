package invoke

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/runner/model"
)

const defaultTimeout = 90 * time.Minute

func Agent(ctx context.Context, request model.AgentRequest) (result model.AgentResult, retErr error) {
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
	promptFile := ""
	if usesPromptFile(command) || usesPromptFile(agent["wrap"]) {
		promptFile, err = writePromptFile(request.Prompt)
		if err != nil {
			return result, err
		}
		defer func() {
			if err := os.Remove(promptFile); err != nil {
				retErr = errors.Join(retErr, fmt.Errorf("remove agent prompt file: %w", err))
			}
		}()
	}
	argv, err := commandArgs(command, request.Prompt, promptFile, repo.CommonDir)
	if err != nil {
		return result, err
	}
	wrap, err := argvValue(agent["wrap"], request.Prompt, promptFile, repo.CommonDir)
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
	path, logErr := logPath(request)
	if logErr == nil {
		result.LogPath = path
		logErr = appendLog(path, process.Stdout, process.Stderr)
	}
	if logErr != nil {
		logErr = &LogError{Cause: logErr}
	}
	if process.ExitCode != 0 {
		patterns, err := stringValues(request.Snapshot.Config["limit_patterns"])
		if err != nil {
			return result, errors.Join(runErr, logErr, fmt.Errorf("limit_patterns: %w", err))
		}
		result.UsageLimit = matchesLimit(process.Stdout, process.Stderr, patterns)
	}
	return result, errors.Join(runErr, logErr)
}

// LogError distinguishes a persistence failure from an agent's usage response.
type LogError struct{ Cause error }

func (e *LogError) Error() string { return fmt.Sprintf("persist agent log: %v", e.Cause) }
func (e *LogError) Unwrap() error { return e.Cause }

func IsLogError(err error) bool {
	var logError *LogError
	return errors.As(err, &logError)
}

// NewRunID creates a log identity for one runner operation.
func NewRunID() (string, error) {
	var random [16]byte
	if _, err := rand.Read(random[:]); err != nil {
		return "", fmt.Errorf("create agent log run: %w", err)
	}
	return hex.EncodeToString(random[:]), nil
}

func logPath(request model.AgentRequest) (string, error) {
	root := request.LogRoot
	if root == "" {
		root = request.Snapshot.Root
	}
	root, err := filepath.Abs(root)
	if err != nil {
		return "", fmt.Errorf("agent log root: %w", err)
	}
	run := request.Run
	if run == "" {
		run, err = NewRunID()
		if err != nil {
			return "", err
		}
	}
	phase := request.Phase
	if phase == "" {
		phase = request.Role
	}
	if !pathComponent(run) || !pathComponent(phase) {
		return "", fmt.Errorf("invalid agent log run or phase")
	}
	attempt := request.Attempt
	if attempt <= 0 {
		attempt = 1
	}
	return filepath.Join(root, ".fireweed", "runs", run, fmt.Sprintf("%s-%d.log", phase, attempt)), nil
}

func pathComponent(value string) bool {
	return value != "" && value != "." && value != ".." && !strings.ContainsAny(value, "/\\")
}

func appendLog(path, stdout, stderr string) (err error) {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("create agent log directory: %w", err)
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("open agent log: %w", err)
	}
	defer func() { err = errors.Join(err, file.Close()) }()
	// Lengths make each stream unambiguous even when it contains no final newline.
	if _, err = fmt.Fprintf(file, "stdout %d bytes\n", len(stdout)); err != nil {
		return fmt.Errorf("write agent log: %w", err)
	}
	if _, err = file.WriteString(stdout); err != nil {
		return fmt.Errorf("write agent log: %w", err)
	}
	if _, err = fmt.Fprintf(file, "\nstderr %d bytes\n", len(stderr)); err != nil {
		return fmt.Errorf("write agent log: %w", err)
	}
	if _, err = file.WriteString(stderr); err != nil {
		return fmt.Errorf("write agent log: %w", err)
	}
	if _, err = file.WriteString("\n"); err != nil {
		return fmt.Errorf("write agent log: %w", err)
	}
	return nil
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

func commandArgs(value any, prompt, promptFile, commonDir string) ([]string, error) {
	if shell, ok := value.(string); ok {
		return []string{"/bin/sh", "-c", expand(shell, shellQuote(prompt), shellQuote(promptFile), shellQuote(commonDir))}, nil
	}
	return argvValue(value, prompt, promptFile, commonDir)
}

func argvValue(value any, prompt, promptFile, commonDir string) ([]string, error) {
	parts, err := stringValues(value)
	if err != nil {
		return nil, err
	}
	args := make([]string, len(parts))
	for i, part := range parts {
		args[i] = expand(part, prompt, promptFile, commonDir)
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

func expand(template, prompt, promptFile, commonDir string) string {
	return strings.NewReplacer("{prompt}", prompt, "{prompt_file}", promptFile, "{git_common_dir}", commonDir).Replace(template)
}

func usesPromptFile(value any) bool {
	switch typed := value.(type) {
	case string:
		return strings.Contains(typed, "{prompt_file}")
	case []string:
		for _, part := range typed {
			if strings.Contains(part, "{prompt_file}") {
				return true
			}
		}
	case []any:
		for _, part := range typed {
			if text, ok := part.(string); ok && strings.Contains(text, "{prompt_file}") {
				return true
			}
		}
	}
	return false
}

func writePromptFile(prompt string) (string, error) {
	file, err := os.CreateTemp("", "fireweed-prompt-")
	if err != nil {
		return "", fmt.Errorf("create agent prompt file: %w", err)
	}
	keep := false
	defer func() {
		if !keep {
			_ = os.Remove(file.Name())
		}
	}()
	path, err := filepath.Abs(file.Name())
	if err == nil {
		_, err = file.WriteString(prompt)
	}
	err = errors.Join(err, file.Close())
	if err != nil {
		return "", fmt.Errorf("write agent prompt file: %w", err)
	}
	keep = true
	return path, nil
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
