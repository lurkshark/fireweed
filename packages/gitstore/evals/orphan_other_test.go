//go:build !linux && !darwin && !windows

package evals

import "os/exec"

func orphanModes() []string                                         { return nil }
func orphanSeparate(cmd *exec.Cmd, mode string)                     {}
func orphanSession() (int, error)                                   { return 0, nil }
func orphanIndependent(parent, leaf orphanReady, mode string) error { return nil }
func orphanReparented(pid, parent int) error                        { return nil }
func orphanPin(pid int, root, label string) (*nestedOwned, error)   { return nestedPin(pid, root) }
