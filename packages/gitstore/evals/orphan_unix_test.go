//go:build linux || darwin

package evals

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func orphanModes() []string { return []string{"group", "session"} }
func orphanSeparate(cmd *exec.Cmd, mode string) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: mode == "group", Setsid: mode == "session"}
}
func orphanSession() (int, error) {
	value, _, err := syscall.Syscall(syscall.SYS_GETSID, 0, 0, 0)
	if err != 0 {
		return 0, err
	}
	return int(value), nil
}
func orphanIndependent(parent, leaf orphanReady, mode string) error {
	if leaf.Group != leaf.PID || leaf.Group == parent.Group || leaf.Session <= 0 || parent.Session <= 0 {
		return fmt.Errorf("independent group absent: parent=%+v leaf=%+v", parent, leaf)
	}
	if mode == "session" && (leaf.Session != leaf.PID || leaf.Session == parent.Session) {
		return fmt.Errorf("native new session absent: parent=%+v leaf=%+v", parent, leaf)
	}
	if mode == "group" && leaf.Session != parent.Session {
		return fmt.Errorf("group control changed session: parent=%+v leaf=%+v", parent, leaf)
	}
	return nil
}
func orphanReparented(pid, parent int) error {
	value, err := nestedPS(pid, "ppid")
	if err != nil {
		return err
	}
	ppid, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	if ppid <= 0 || ppid == parent {
		return fmt.Errorf("PID%d not reparented from%d: parent%d", pid, parent, ppid)
	}
	return nil
}
func orphanPin(pid int, root, label string) (*nestedOwned, error) {
	stage, argument := "leaf", label
	if strings.HasSuffix(label, "-parent") {
		stage, argument = "bridge", strings.TrimSuffix(label, "-parent")
	}
	alive := func() (bool, error) {
		active, err := nestedAlive(pid)
		if err != nil || !active {
			return active, err
		}
		command, err := nestedPS(pid, "command")
		if err != nil {
			return false, err
		}
		return strings.Contains(command, orphanHelperName) && strings.Contains(command, "-- "+stage+" "+root+" "+argument+" "), nil
	}
	active, err := alive()
	if err != nil || !active {
		return nil, fmt.Errorf("live exact helper%d/%s: %t/%v", pid, label, active, err)
	}
	return &nestedOwned{alive: alive, kill: func() error {
		active, err := alive()
		if err != nil || !active {
			return err
		}
		err = syscall.Kill(pid, syscall.SIGKILL)
		if err == syscall.ESRCH {
			return nil
		}
		return err
	}, close: func() {}}, nil
}
