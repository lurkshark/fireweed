//go:build !linux && !darwin && !windows

package process

import (
	"os"
	"os/exec"
)

type directControl struct{}

func newTree(command *exec.Cmd) (treeControl, error) { return directControl{}, nil }
func (directControl) started(*exec.Cmd) error        { return nil }
func (directControl) close()                         {}
func (directControl) stop(command *exec.Cmd) error {
	if command.Process == nil {
		return os.ErrProcessDone
	}
	return command.Process.Kill()
}
