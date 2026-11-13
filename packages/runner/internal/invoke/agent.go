package invoke

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/runner/internal/policy"
	"github.com/lurkshark/fireweed/packages/runner/model"
)

func NewRun() (string, error) {
	var nonce [12]byte
	if _, err := rand.Read(nonce[:]); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s-%x", time.Now().UTC().Format("20060102T150405"), nonce), nil
}

func Agent(ctx context.Context, request model.AgentRequest) (result model.AgentResult, err error) {
	result.LogPath, err = logPath(request)
	if err != nil {
		return result, err
	}
	defer func() {
		if logErr := appendLog(result); logErr != nil {
			err = errors.Join(err, fmt.Errorf("write agent log: %w", logErr))
		}
	}()

	config := request.Snapshot.Config
	command := Command(config, request.Role)
	wrap, err := Strings(Mapping(config["agent"])["wrap"])
	if err != nil {
		return result, fmt.Errorf("agent.wrap: %w", err)
	}
	timeout, err := policy.Timeout(config, request.Timeout)
	if err != nil {
		return result, err
	}
	templates, _ := Strings(command)
	if shell, ok := command.(string); ok {
		templates = []string{shell}
	}
	uses := func(token string) bool {
		for _, group := range [][]string{templates, wrap} {
			for _, template := range group {
				if strings.Contains(template, token) {
					return true
				}
			}
		}
		return false
	}

	var common, promptFile string
	if uses("{git_common_dir}") {
		repo, err := gitstore.Inspect(request.Snapshot.Root)
		if err != nil {
			return result, err
		}
		common = repo.CommonDir
	}
	if uses("{prompt_file}") {
		file, err := os.CreateTemp("", "fireweed-prompt-*")
		if err != nil {
			return result, err
		}
		defer func() {
			if removeErr := os.Remove(file.Name()); removeErr != nil && !os.IsNotExist(removeErr) {
				err = errors.Join(err, removeErr)
			}
		}()
		_, writeErr := file.WriteString(request.Prompt)
		if err := errors.Join(writeErr, file.Close()); err != nil {
			return result, err
		}
		promptFile, err = filepath.Abs(file.Name())
		if err != nil {
			return result, err
		}
	}
	tokens := []string{"{prompt}", request.Prompt, "{prompt_file}", promptFile, "{git_common_dir}", common}
	args, err := Args(command, tokens...)
	if err != nil {
		return result, fmt.Errorf("%s command: %w", request.Role, err)
	}
	patterns, err := Strings(config["limit_patterns"])
	if err != nil {
		return result, fmt.Errorf("limit_patterns: %w", err)
	}

	process, err := gitstore.Run(ctx, gitstore.Request{
		Args: append(expand(wrap, strings.NewReplacer(tokens...)), args...),
		Dir:  request.Snapshot.Root, Timeout: timeout,
		Env: map[string]string{"FIREWEED_BIN": request.Bin},
	})
	result.Stdout, result.Stderr, result.ExitCode = process.Stdout, process.Stderr, process.ExitCode
	if result.ExitCode != 0 {
		result.UsageLimit = matchesLimit(result, patterns)
	}
	if err != nil {
		return result, fmt.Errorf("%s agent %s: %w", request.Role, Name(config, request.Role), err)
	}
	return result, nil
}

func matchesLimit(result model.AgentResult, patterns []string) bool {
	var lines []string
	for _, stream := range []string{result.Stdout, result.Stderr} {
		if stream = strings.TrimRight(stream, "\r\n"); stream != "" {
			lines = append(lines, strings.Split(stream, "\n")...)
		}
	}
	tail := strings.ToLower(strings.Join(lines[max(0, len(lines)-20):], "\n"))
	for _, pattern := range patterns {
		if strings.Contains(tail, strings.ToLower(pattern)) {
			return true
		}
	}
	return false
}

func logPath(request model.AgentRequest) (string, error) {
	root, run, phase, attempt := request.LogRoot, request.Run, request.Phase, request.Attempt
	if root == "" {
		root = request.Snapshot.Root
	}
	if run == "" {
		var err error
		run, err = NewRun()
		if err != nil {
			return "", err
		}
	}
	if phase == "" {
		phase = request.Role
	}
	if attempt <= 0 {
		attempt = 1
	}
	for _, part := range []string{run, phase} {
		if filepath.Base(part) != part || part == ".." || part == "." {
			return "", fmt.Errorf("invalid agent log run or phase %q", part)
		}
	}
	return filepath.Abs(filepath.Join(root, ".fireweed", "runs", run, fmt.Sprintf("%s-%d.log", phase, attempt)))
}

func appendLog(result model.AgentResult) error {
	if err := os.MkdirAll(filepath.Dir(result.LogPath), 0755); err != nil {
		return err
	}
	file, err := os.OpenFile(result.LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	_, err = file.WriteString(fmt.Sprintf("stdout %d bytes\n%s\nstderr %d bytes\n%s\n", len(result.Stdout), result.Stdout, len(result.Stderr), result.Stderr))
	return errors.Join(err, file.Close())
}
