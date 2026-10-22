package process

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type treeControl interface {
	started(*exec.Cmd) error
	stop(*exec.Cmd) error
	close()
}

func Run(ctx context.Context, dir string, argv, overrides []string) (string, int, error) {
	if ctx == nil {
		return "", 2, fmt.Errorf("process requires a non-nil context")
	}
	if err := ctx.Err(); err != nil {
		return "", -1, err
	}
	if len(argv) == 0 || argv[0] == "" {
		return "", 2, fmt.Errorf("provide a nonempty executable argument vector")
	}
	for _, entry := range overrides {
		if strings.IndexByte(entry, '=') <= 0 || strings.ContainsRune(entry, '\x00') {
			return "", 2, fmt.Errorf("environment override must be a nonempty KEY=value without NUL bytes: %q", entry)
		}
	}
	args := append([]string(nil), argv...)
	command := exec.CommandContext(ctx, args[0], args[1:]...)
	command.Dir = dir
	command.Env = append(os.Environ(), overrides...)
	var output bytes.Buffer
	command.Stdout = &output
	command.Stderr = &output
	command.WaitDelay = time.Second
	control, err := newTree(command)
	if err != nil {
		return "", 2, fmt.Errorf("prepare process tree: %w", err)
	}
	defer control.close()
	// Cancellation cannot race assignment/resumption of a suspended Windows child.
	var startup sync.Mutex
	command.Cancel = func() error { startup.Lock(); defer startup.Unlock(); return control.stop(command) }
	startup.Lock()
	startErr := command.Start()
	var setupErr error
	if startErr == nil {
		setupErr = control.started(command)
		if setupErr != nil {
			_ = control.stop(command)
			_ = command.Process.Kill()
		}
	}
	startup.Unlock()
	if startErr != nil {
		if err := ctx.Err(); err != nil {
			return output.String(), -1, err
		}
		return output.String(), 2, fmt.Errorf("start executable %q: %w", args[0], startErr)
	}
	waitErr := command.Wait()
	if err := ctx.Err(); err != nil {
		return output.String(), -1, err
	}
	if setupErr != nil {
		return output.String(), 2, fmt.Errorf("establish process tree: %w", setupErr)
	}
	if waitErr == nil {
		return output.String(), 0, nil
	}
	var exited *exec.ExitError
	if errors.As(waitErr, &exited) {
		return output.String(), exited.ExitCode(), waitErr
	}
	return output.String(), 2, fmt.Errorf("wait for executable %q: %w", args[0], waitErr)
}
