package process

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"
	"unsafe"

	"github.com/lurkshark/fireweed/packages/gitstore/internal/process/native"
)

type darwinControl struct{}

func newTree(command *exec.Cmd) (treeControl, error) {
	if command.Err != nil {
		return newGroupTree(command)
	}
	return &darwinControl{}, nil
}
func (*darwinControl) started(context.Context, *exec.Cmd) error { return nil }
func (*darwinControl) stop(*exec.Cmd) error                     { return os.ErrProcessDone }
func (*darwinControl) close()                                   {}

// run keeps a distinct native responsibility anchor alive through root exit and
// inherited stream drain. Its entry precedes Go initializers/main. Request data
// crosses exec through anonymous POSIX shared memory, never a temporary file.
func (*darwinControl) run(ctx context.Context, command *exec.Cmd) (int, error) {
	env := command.Environ()
	values := append([]string{command.Path}, command.Args...)
	values = append(values, env...)
	size := 64
	for _, value := range values {
		if strings.ContainsRune(value, 0) {
			return 2, errors.New("process argument or environment contains NUL")
		}
		size += len(value) + 1
	}
	if size > 16<<20 {
		return 2, errors.New("process request exceeds native argument capacity")
	}
	executable, err := os.Executable()
	if err != nil {
		return 2, fmt.Errorf("locate native process anchor: %w", err)
	}
	path, err := syscall.BytePtrFromString(executable)
	if err != nil {
		return 2, err
	}
	cwd, err := syscall.BytePtrFromString(command.Dir)
	if err != nil {
		return 2, err
	}
	var random [24]byte
	if _, err = rand.Read(random[:]); err != nil {
		return 2, err
	}
	nonce, _ := syscall.BytePtrFromString(hex.EncodeToString(random[:]))
	name, _ := syscall.BytePtrFromString("/fw" + hex.EncodeToString(random[:12]))
	fd, _, errno := syscall.Syscall(syscall.SYS_SHM_OPEN, uintptr(unsafe.Pointer(name)), syscall.O_CREAT|syscall.O_EXCL|syscall.O_RDWR, 0600)
	runtime.KeepAlive(name)
	if errno != 0 {
		return 2, fmt.Errorf("allocate anonymous process request: %w", errno)
	}
	request := os.NewFile(fd, "fireweed-process-request")
	defer request.Close()
	syscall.CloseOnExec(int(fd))
	_, _, errno = syscall.Syscall(syscall.SYS_SHM_UNLINK, uintptr(unsafe.Pointer(name)), 0, 0)
	runtime.KeepAlive(name)
	if errno != 0 {
		return 2, fmt.Errorf("unlink anonymous process request: %w", errno)
	}
	if err = syscall.Ftruncate(int(fd), int64(size)); err != nil {
		return 2, err
	}
	mapped, err := syscall.Mmap(int(fd), 0, size, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		return 2, err
	}
	defer syscall.Munmap(mapped)
	put := func(offset int, value uint64) { binary.LittleEndian.PutUint64(mapped[offset:offset+8], value) }
	put(0, 0x667770726f633031)
	put(8, uint64(size))
	put(16, uint64(len(command.Args)))
	put(24, uint64(len(env)))
	if command.Stdout == command.Stderr {
		put(32, 1)
	}
	copy(mapped[40:64], random[:])
	offset := 64
	for _, value := range values {
		offset += copy(mapped[offset:], value)
		mapped[offset] = 0
		offset++
	}
	var files []*os.File
	defer func() {
		for _, file := range files {
			_ = file.Close()
		}
	}()
	pipe := func() (*os.File, *os.File, error) {
		r, w, e := os.Pipe()
		if e == nil {
			files = append(files, r, w)
		}
		return r, w, e
	}
	cancelR, cancelW, err := pipe()
	if err != nil {
		return 2, err
	}
	replyR, replyW, err := pipe()
	if err != nil {
		return 2, err
	}
	outR, outW, err := pipe()
	if err != nil {
		return 2, err
	}
	errR, errW := outR, outW
	if command.Stdout != command.Stderr {
		errR, errW, err = pipe()
		if err != nil {
			return 2, err
		}
	}
	if err = ctx.Err(); err != nil {
		return -1, err
	}
	var pid int32
	launch := native.Spawn(path, cwd, nonce, [5]int{int(fd), int(cancelR.Fd()), int(replyW.Fd()), int(outW.Fd()), int(errW.Fd())}, &pid)
	runtime.KeepAlive(path)
	runtime.KeepAlive(cwd)
	runtime.KeepAlive(nonce)
	if launch != 0 {
		if err = ctx.Err(); err != nil {
			return -1, err
		}
		return 2, fmt.Errorf("start executable %q: %w", command.Args[0], syscall.Errno(launch))
	}
	process, err := os.FindProcess(int(pid))
	if err != nil {
		_ = syscall.Kill(int(pid), syscall.SIGKILL)
		return 2, err
	}
	command.Process = process
	fallback := groupControl{}
	_ = fallback.started(ctx, command)
	_ = cancelR.Close()
	_ = replyW.Close()
	_ = outW.Close()
	if errW != outW {
		_ = errW.Close()
	}
	_ = request.Close()
	var cancelOnce sync.Once
	cancel := func() { cancelOnce.Do(func() { _ = cancelW.Close() }) }
	streams := make(chan error, 2)
	copyStream := func(reader *os.File, writer io.Writer) {
		_, e := io.Copy(writer, reader)
		if e != nil {
			cancel()
		}
		streams <- e
	}
	go copyStream(outR, command.Stdout)
	streamCount := 1
	if errR != outR {
		streamCount++
		go copyStream(errR, command.Stderr)
	}
	type outcome struct {
		state *os.ProcessState
		err   error
	}
	finished := make(chan outcome, 1)
	go func() { state, e := process.Wait(); finished <- outcome{state, e} }()
	var result outcome
	select {
	case result = <-finished:
	case <-ctx.Done():
		cancel()
		select {
		case result = <-finished:
		case <-time.After(5 * time.Second):
			_ = fallback.stop(command)
			result = <-finished
		}
	}
	var streamErr error
	for i := 0; i < streamCount; i++ {
		streamErr = errors.Join(streamErr, <-streams)
	}
	raw, replyErr := io.ReadAll(io.LimitReader(replyR, 121))
	if err = ctx.Err(); err != nil {
		return -1, err
	}
	if result.err != nil {
		return 2, fmt.Errorf("wait for native process anchor: %w", result.err)
	}
	if result.state == nil {
		return 2, errors.New("native process anchor omitted its process state")
	}
	if streamErr != nil {
		return 2, fmt.Errorf("collect process streams: %w", streamErr)
	}
	if replyErr != nil {
		return 2, fmt.Errorf("read native process outcome: %w", replyErr)
	}
	if len(raw) != 40 && len(raw) != 80 {
		return 2, errors.New("native process anchor omitted its outcome")
	}
	ready := false
	for len(raw) > 0 {
		if binary.LittleEndian.Uint64(raw[:8]) != 0x66777265706c3031 {
			return 2, errors.New("invalid native process outcome")
		}
		stage := binary.LittleEndian.Uint64(raw[8:16])
		code := int(int64(binary.LittleEndian.Uint64(raw[16:24])))
		problem := binary.LittleEndian.Uint64(raw[24:32])
		raw = raw[40:]
		if stage == 1 && !ready && len(raw) > 0 {
			ready = true
			continue
		}
		if stage != 2 || len(raw) != 0 {
			return 2, errors.New("invalid native process outcome sequence")
		}
		if problem != 0 {
			return code, fmt.Errorf("execute %q: %w", command.Args[0], syscall.Errno(problem))
		}
		if !ready {
			return 2, errors.New("native process anchor omitted startup readiness")
		}
		if result.state.ExitCode() != code {
			return 2, fmt.Errorf("native process outcome disagrees with its exit state: %v", result.state)
		}
		if code != 0 {
			return code, &exec.ExitError{ProcessState: result.state}
		}
		return 0, nil
	}
	return 2, errors.New("native process anchor omitted its final outcome")
}
