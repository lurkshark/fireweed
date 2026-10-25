package supervision

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type rootResult struct {
	state *os.ProcessState
	err   error
}

// Run is used by a separate, statically linked helper, not by the calling
// program's main or initializers. Only this isolated process becomes a reaper.
func Run() int {
	if len(os.Args) != 2 {
		return 2
	}
	nonce := os.Args[1]
	requestFile := os.NewFile(RequestFD, "request")
	replyFile := os.NewFile(ReplyFD, "reply")
	cancelFile := os.NewFile(CancelFD, "cancel")
	if requestFile == nil || replyFile == nil || cancelFile == nil {
		return 2
	}
	defer requestFile.Close()
	defer replyFile.Close()
	defer cancelFile.Close()
	// These descriptors belong only to the supervisor protocol. Never let
	// the requested executable or its descendants inherit them.
	for _, fd := range []int{RequestFD, ReplyFD, CancelFD} {
		syscall.CloseOnExec(fd)
	}
	_ = os.NewFile(ImageFD, "image").Close()
	encoder := gob.NewEncoder(replyFile)
	finish := func(reply Reply) int {
		reply.Version, reply.Nonce, reply.Final = Version, nonce, true
		_ = encoder.Encode(reply)
		if reply.Exit > 0 && reply.Exit < 256 {
			return reply.Exit
		}
		if reply.Error != "" {
			return 1
		}
		return 0
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() { var one [1]byte; _, _ = cancelFile.Read(one[:]); cancel() }()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, os.Interrupt)
	defer signal.Stop(signals)
	go func() {
		select {
		case <-signals:
			cancel()
		case <-ctx.Done():
		}
	}()
	var request Request
	if err := gob.NewDecoder(io.LimitReader(requestFile, 16<<20)).Decode(&request); err != nil {
		return finish(Reply{Exit: 2, LaunchError: true, Error: fmt.Sprintf("read process request: %v", err)})
	}
	requestFile.Close()
	if request.Version != Version || request.Nonce != nonce || nonce == "" || request.Executable == "" || len(request.Args) == 0 {
		return finish(Reply{Exit: 2, LaunchError: true, Error: "invalid process supervisor request"})
	}
	// PR_SET_CHILD_SUBREAPER is local to this helper. Reparented descendants
	// remain its children even if they create new groups or sessions.
	if _, _, errno := syscall.RawSyscall6(syscall.SYS_PRCTL, 36, 1, 0, 0, 0, 0); errno != 0 {
		return finish(Reply{Exit: 2, LaunchError: true, Errno: int(errno), Error: "establish isolated child reaper: " + errno.Error()})
	}
	if err := encoder.Encode(Reply{Version: Version, Nonce: nonce, Ready: true}); err != nil {
		return 2
	}
	if ctx.Err() != nil {
		return finish(Reply{Exit: -1, Canceled: true, Error: "process canceled before startup"})
	}
	outR, outW, err := os.Pipe()
	if err != nil {
		return finish(Reply{Exit: 2, LaunchError: true, Error: err.Error()})
	}
	defer outR.Close()
	defer outW.Close()
	errR, errW := outR, outW
	if !request.Combined {
		errR, errW, err = os.Pipe()
		if err != nil {
			return finish(Reply{Exit: 2, LaunchError: true, Error: err.Error()})
		}
		defer errR.Close()
		defer errW.Close()
	}
	command := &exec.Cmd{Path: request.Executable, Args: request.Args, Env: request.Env, Stdout: outW, Stderr: errW}
	if err := command.Start(); err != nil {
		var errno syscall.Errno
		errors.As(err, &errno)
		return finish(Reply{Exit: 2, LaunchError: true, Errno: int(errno), Error: err.Error()})
	}
	rootPID := command.Process.Pid
	outW.Close()
	if errW != outW {
		errW.Close()
	}
	copies := 1
	if !request.Combined {
		copies = 2
	}
	copyDone := make(chan error, copies)
	go func() { _, err := io.Copy(os.Stdout, outR); outR.Close(); copyDone <- err }()
	if !request.Combined {
		go func() { _, err := io.Copy(os.Stderr, errR); errR.Close(); copyDone <- err }()
	}
	rootDone := make(chan rootResult, 1)
	go func() {
		state, err := command.Process.Wait()
		command.ProcessState = state
		rootDone <- rootResult{state, err}
	}()
	var result rootResult
	rootExited, forwarded := false, 0
	var problem error
	ticker := time.NewTicker(2 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case result = <-rootDone:
			rootExited = true
			problem = errors.Join(problem, result.err)
		default:
		}
		for forwarded < copies {
			select {
			case err := <-copyDone:
				forwarded++
				problem = errors.Join(problem, err)
			default:
				goto copiesObserved
			}
		}
	copiesObserved:
		stopping := ctx.Err() != nil || problem != nil
		children, err := directChildren()
		problem = errors.Join(problem, err)
		if stopping {
			for _, pid := range children {
				// The root has a concurrent Go waiter: signal its lifetime handle,
				// not a raw PID which that waiter might already have released.
				if pid == rootPID && !rootExited {
					err = command.Process.Kill()
				} else {
					err = syscall.Kill(pid, syscall.SIGKILL)
				}
				if err != nil && !errors.Is(err, os.ErrProcessDone) && !errors.Is(err, syscall.ESRCH) {
					problem = errors.Join(problem, err)
				}
			}
		}
		for _, pid := range children {
			if pid == rootPID && !rootExited {
				continue
			}
			var status syscall.WaitStatus
			_, err := syscall.Wait4(pid, &status, syscall.WNOHANG, nil)
			if err != nil && err != syscall.ECHILD && err != syscall.EINTR {
				problem = errors.Join(problem, err)
			}
		}
		if rootExited && forwarded == copies {
			// Preserve ordinary exec semantics: detached children with closed
			// streams do not extend a completed, uncanceled call.
			if !stopping {
				return finish(rootReply(result, false, problem))
			}
			remaining, err := directChildren()
			problem = errors.Join(problem, err)
			if err == nil && len(remaining) == 0 {
				return finish(rootReply(result, ctx.Err() != nil, problem))
			}
		}
		select {
		case <-ticker.C:
		case <-ctx.Done():
			time.Sleep(time.Millisecond)
		}
	}
}

func rootReply(result rootResult, canceled bool, problem error) Reply {
	reply := Reply{HasRoot: result.state != nil, Canceled: canceled, Exit: 2}
	if result.state != nil {
		reply.Exit = result.state.ExitCode()
		if !result.state.Success() {
			reply.Error = result.state.String()
		}
	}
	if problem != nil {
		if reply.Error == "" {
			reply.Error = problem.Error()
		} else {
			reply.Error = errors.Join(errors.New(reply.Error), problem).Error()
		}
		if reply.Exit == 0 {
			reply.Exit = 2
		}
	}
	return reply
}

func directChildren() ([]int, error) {
	tasks, err := os.ReadDir("/proc/self/task")
	if err != nil {
		return nil, err
	}
	result, seen := []int{}, map[int]bool{}
	for _, task := range tasks {
		if _, err := strconv.Atoi(task.Name()); err != nil {
			continue
		}
		raw, err := os.ReadFile(filepath.Join("/proc/self/task", task.Name(), "children"))
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return nil, err
		}
		for _, field := range strings.Fields(string(raw)) {
			pid, err := strconv.Atoi(field)
			if err != nil || pid <= 0 {
				return nil, fmt.Errorf("invalid owned child inventory")
			}
			if !seen[pid] {
				seen[pid] = true
				result = append(result, pid)
			}
		}
	}
	return result, nil
}
