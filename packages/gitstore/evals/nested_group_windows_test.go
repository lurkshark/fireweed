//go:build windows

package evals

import (
	"os/exec"
	"syscall"
)

func nestedSeparateGroup(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP}
}

func nestedSelfGroup() int { return 0 }

func nestedPin(pid int, root string) (*nestedOwned, error) {
	// Pin while every helper is live and awaiting the caller's arm file.
	// Keep this handle for later observation/cleanup instead of reopening a PID.
	handle, err := syscall.OpenProcess(syscall.PROCESS_QUERY_INFORMATION|syscall.PROCESS_TERMINATE, false, uint32(pid))
	if err != nil {
		return nil, err
	}
	alive := func() (bool, error) {
		var code uint32
		if err := syscall.GetExitCodeProcess(handle, &code); err != nil {
			return false, err
		}
		return code == 259, nil // STILL_ACTIVE
	}
	return &nestedOwned{
		alive: alive,
		kill: func() error {
			err := syscall.TerminateProcess(handle, 101)
			if err != nil {
				if active, check := alive(); check == nil && !active {
					return nil
				}
			}
			return err
		},
		close: func() { syscall.CloseHandle(handle) },
	}, nil
}
