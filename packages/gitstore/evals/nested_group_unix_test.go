//go:build linux || darwin

package evals

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func nestedSeparateGroup(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

func nestedSelfGroup() int { return syscall.Getpgrp() }

func nestedPS(pid int, field string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "ps", "-ww", "-p", strconv.Itoa(pid), "-o", field+"=")
	raw, err := cmd.Output()
	if err != nil {
		var exit *exec.ExitError
		if errors.As(err, &exit) && exit.ExitCode() == 1 && len(raw) == 0 {
			return "", nil
		}
		return "", err
	}
	return strings.TrimSpace(string(raw)), nil
}

func nestedAlive(pid int) (bool, error) {
	state, err := nestedPS(pid, "stat")
	return state != "" && !strings.HasPrefix(state, "Z"), err
}

func nestedKillOwned(pid int, root string) error {
	alive, err := nestedAlive(pid)
	if err != nil || !alive {
		return err
	}
	command, err := nestedPS(pid, "command")
	if err != nil {
		return err
	}
	if command == "" {
		return nil
	}
	if !strings.Contains(command, root) || !strings.Contains(command, nestedHelperName) {
		return fmt.Errorf("refusing to kill PID %d without helper/root identity", pid)
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	err = process.Kill()
	if errors.Is(err, os.ErrProcessDone) || errors.Is(err, syscall.ESRCH) {
		return nil
	}
	return err
}

func nestedPin(pid int, root string) (*nestedOwned, error) {
	return &nestedOwned{
		alive: func() (bool, error) { return nestedAlive(pid) },
		kill:  func() error { return nestedKillOwned(pid, root) },
		close: func() {},
	}, nil
}
