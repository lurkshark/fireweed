//go:build linux || darwin

package process

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

type processRef struct {
	pid, parent   int
	birth         [2]uint64
	stopped, dead bool
}

func sameProcess(a, b processRef) bool { return a.pid == b.pid && a.birth == b.birth }
func processGone(err error) bool {
	return errors.Is(err, os.ErrNotExist) || errors.Is(err, syscall.ESRCH) || errors.Is(err, os.ErrProcessDone)
}

type groupControl struct {
	root  processRef
	known bool
}

func newTree(command *exec.Cmd) (treeControl, error) {
	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	return &groupControl{}, nil
}
func (control *groupControl) started(command *exec.Cmd) error {
	// A short-lived child may already have exited. Optional metadata must
	// not replace an otherwise ordinary successful launch with an error.
	if ref, err := nativeProcess(command.Process.Pid); err == nil {
		control.root, control.known = ref, true
	}
	return nil
}
func (*groupControl) close() {}

func (control *groupControl) stop(command *exec.Cmd) error {
	if command.Process == nil {
		return os.ErrProcessDone
	}
	root, err := nativeProcess(command.Process.Pid)
	if err == nil && control.known && !sameProcess(root, control.root) {
		return os.ErrProcessDone
	}
	var held []processRef
	seen := map[int]bool{}
	var discover func(processRef, bool) error
	discover = func(ref processRef, initial bool) error {
		if seen[ref.pid] || ref.dead {
			return nil
		}
		seen[ref.pid] = true
		current, err := nativeProcess(ref.pid)
		if processGone(err) || err == nil && !sameProcess(current, ref) {
			return nil
		}
		if err != nil {
			return err
		}
		if initial {
			err = command.Process.Signal(syscall.SIGSTOP)
		} else {
			err = syscall.Kill(ref.pid, syscall.SIGSTOP)
		}
		if processGone(err) {
			return nil
		}
		if err != nil {
			return err
		}
		// Freeze a parent before enumerating children: their PIDs cannot be
		// reaped/reused by that parent while crossing independent groups.
		held = append(held, ref)
		waitUntil := time.Now().Add(500 * time.Millisecond)
		var stoppedErr error
		for {
			current, err = nativeProcess(ref.pid)
			if processGone(err) || err == nil && !sameProcess(current, ref) {
				return nil
			}
			if err != nil {
				return err
			}
			if current.stopped || current.dead {
				break
			}
			if time.Now().After(waitUntil) {
				stoppedErr = fmt.Errorf("process %d did not acknowledge stop", ref.pid)
				break
			}
			time.Sleep(time.Millisecond)
		}
		children, err := nativeChildren(ref.pid)
		if err != nil && !processGone(err) {
			return errors.Join(stoppedErr, err)
		}
		for _, pid := range children {
			child, err := nativeProcess(pid)
			if processGone(err) {
				continue
			}
			if err != nil {
				stoppedErr = errors.Join(stoppedErr, err)
				continue
			}
			if child.parent == ref.pid {
				stoppedErr = errors.Join(stoppedErr, discover(child, false))
			}
		}
		return stoppedErr
	}
	var stopErr error
	if err == nil {
		stopErr = discover(root, true)
	} else if !processGone(err) {
		stopErr = err
	}
	// Kill descendants before stopped parents, rechecking creation identity.
	// Cleanup continues even if a stop or metadata operation failed.
	for i := len(held) - 1; i >= 0; i-- {
		ref := held[i]
		if ref.pid == command.Process.Pid {
			continue
		}
		current, err := nativeProcess(ref.pid)
		if processGone(err) || err == nil && !sameProcess(current, ref) {
			continue
		}
		if err == nil {
			err = syscall.Kill(ref.pid, syscall.SIGKILL)
		}
		if err != nil && !processGone(err) {
			stopErr = errors.Join(stopErr, err)
		}
	}
	// The original group also covers inherited pipes and same-group children
	// whose direct parent already exited. Kill the root separately in case
	// the executable changed its own group after launch.
	groupErr := syscall.Kill(-command.Process.Pid, syscall.SIGKILL)
	rootErr := command.Process.Kill()
	if groupErr != nil && !processGone(groupErr) {
		stopErr = errors.Join(stopErr, groupErr)
	}
	if rootErr != nil && !processGone(rootErr) {
		stopErr = errors.Join(stopErr, rootErr)
	}
	if stopErr == nil && processGone(groupErr) && processGone(rootErr) {
		return os.ErrProcessDone
	}
	return stopErr
}
