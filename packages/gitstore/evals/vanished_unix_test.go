//go:build linux || darwin

package evals

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func vanishedPinAncestor(pid int, root, label string) (*nestedOwned, error) {
	alive := func() (bool, error) {
		live, err := nestedAlive(pid)
		if err != nil || !live {
			return live, err
		}
		command, err := nestedPS(pid, "command")
		if err != nil {
			return false, err
		}
		return strings.Contains(command, vanishedHelperName) && strings.Contains(command, "-- ancestor "+root+" "+label+" "), nil
	}
	live, err := alive()
	if err != nil || !live {
		return nil, fmt.Errorf("pin exact ancestor %d/%s: %t/%v", pid, label, live, err)
	}
	return &nestedOwned{alive: alive, kill: func() error {
		live, err := alive()
		if err != nil || !live {
			return err
		}
		err = syscall.Kill(pid, syscall.SIGKILL)
		if err == syscall.ESRCH {
			return nil
		}
		return err
	}, close: func() {}}, nil
}

// Dead/zombie is insufficient here: a retained zombie still exposes its parent
// identity. Prove this exact PID has actually been reaped before cancellation.
func vanishedReaped(pid int, owned *nestedOwned) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	output, err := exec.CommandContext(ctx, "ps", "-p", strconv.Itoa(pid), "-o", "pid=").Output()
	var exited *exec.ExitError
	if strings.TrimSpace(string(output)) == "" && (err == nil || errors.As(err, &exited) && exited.ExitCode() == 1) {
		return nil
	}
	return fmt.Errorf("ancestor %d still has native process metadata: %q/%v", pid, output, err)
}
