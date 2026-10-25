//go:build windows

package evals

import (
	"os/exec"
	"syscall"
)

func orphanModes() []string { return []string{"group"} }
func orphanSeparate(cmd *exec.Cmd, mode string) {
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP}
}
func orphanSession() (int, error)                                   { return 0, nil }
func orphanIndependent(parent, leaf orphanReady, mode string) error { return nil }
func orphanReparented(pid, parent int) error                        { return nil } // Windows keeps the inherited parent ID; the pinned parent handle proves termination.
func orphanPin(pid int, root, label string) (*nestedOwned, error)   { return nestedPin(pid, root) }
