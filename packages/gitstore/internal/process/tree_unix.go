//go:build linux || darwin

package process

import (
	"errors"
	"os"
	"os/exec"
	"syscall"
)

type groupControl struct{}

func newTree(command *exec.Cmd) (treeControl, error) {
	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	return groupControl{}, nil
}
func (groupControl) started(*exec.Cmd) error { return nil }
func (groupControl) close()                  {}
func (groupControl) stop(command *exec.Cmd) error {
	if command.Process == nil {
		return os.ErrProcessDone
	}
	err := syscall.Kill(-command.Process.Pid, syscall.SIGKILL)
	if errors.Is(err, syscall.ESRCH) {
		return os.ErrProcessDone
	}
	return err
}
