package process

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type treeControl interface {
	started(context.Context, *exec.Cmd) error
	stop(*exec.Cmd) error
	close()
}

func Run(ctx context.Context, dir string, argv, overrides []string) (string, int, error) {
	var output bytes.Buffer
	code, err := execute(ctx, dir, argv, overrides, &output, &output)
	return output.String(), code, err
}

// execute owns the subprocess lifecycle; each caller supplies its own output sinks.
func execute(ctx context.Context, dir string, argv, overrides []string, stdout, stderr io.Writer) (int, error) {
	if ctx == nil {
		return 2, fmt.Errorf("process requires a non-nil context")
	}
	if err := ctx.Err(); err != nil {
		return -1, err
	}
	if len(argv) == 0 || argv[0] == "" {
		return 2, fmt.Errorf("provide a nonempty executable argument vector")
	}
	for _, entry := range overrides {
		if strings.IndexByte(entry, '=') <= 0 || strings.ContainsRune(entry, '\x00') {
			return 2, fmt.Errorf("environment override must be a nonempty KEY=value without NUL bytes: %q", entry)
		}
	}
	args := append([]string(nil), argv...)
	command := exec.CommandContext(ctx, args[0], args[1:]...)
	command.Dir = dir
	command.Env = append(os.Environ(), overrides...)
	command.Stdout = stdout
	command.Stderr = stderr
	command.WaitDelay = time.Second
	control, err := newTree(command)
	if err != nil {
		return 2, fmt.Errorf("prepare process tree: %w", err)
	}
	defer control.close()
	if native, ok := control.(interface {
		run(context.Context, *exec.Cmd) (int, error)
	}); ok {
		return native.run(ctx, command)
	}
	// Cancellation cannot race assignment/resumption of a suspended Windows child.
	var startup sync.Mutex
	stop := func() error { startup.Lock(); defer startup.Unlock(); return control.stop(command) }
	command.Cancel = stop
	startup.Lock()
	startErr := command.Start()
	var setupErr error
	if startErr == nil {
		setupErr = control.started(ctx, command)
		if setupErr != nil {
			_ = control.stop(command)
			_ = command.Process.Kill()
		}
	}
	startup.Unlock()
	if startErr != nil {
		if err := ctx.Err(); err != nil {
			return -1, err
		}
		return 2, fmt.Errorf("start executable %q: %w", args[0], startErr)
	}
	// Cmd's context watcher can finish when its direct child exits even while
	// inherited pipes are still open. Ownership lasts until this whole call ends.
	finished, watcherDone := make(chan struct{}), make(chan struct{})
	go func() {
		defer close(watcherDone)
		select {
		case <-ctx.Done():
			_ = stop()
		case <-finished:
		}
	}()
	waitErr := command.Wait()
	close(finished)
	<-watcherDone
	if err := ctx.Err(); err != nil {
		_ = stop()
		return -1, err
	}
	if setupErr != nil {
		return 2, fmt.Errorf("establish process tree: %w", setupErr)
	}
	if result, ok := control.(interface{ outcome(error) (int, error) }); ok {
		return result.outcome(waitErr)
	}
	if waitErr == nil {
		return 0, nil
	}
	var exited *exec.ExitError
	if errors.As(waitErr, &exited) {
		return exited.ExitCode(), waitErr
	}
	return 2, fmt.Errorf("wait for executable %q: %w", args[0], waitErr)
}
