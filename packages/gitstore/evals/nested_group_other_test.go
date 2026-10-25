//go:build !linux && !darwin && !windows

package evals

import "os/exec"

func nestedSeparateGroup(cmd *exec.Cmd) {}
func nestedSelfGroup() int              { return 0 }
func nestedPin(pid int, root string) (*nestedOwned, error) {
	return &nestedOwned{
		alive: func() (bool, error) { return false, nil },
		kill:  func() error { return nil },
		close: func() {},
	}, nil
}
