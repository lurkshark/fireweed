package process

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/rand"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"syscall"
	"time"
	"unsafe"

	"github.com/lurkshark/fireweed/packages/gitstore/internal/supervision"
)

type reaperControl struct {
	request                   supervision.Request
	files                     []*os.File
	requestW, replyR, cancelW *os.File
	ready, settled            chan struct{}
	endOnce, cancelOnce       sync.Once
	mu                        sync.Mutex
	reply                     supervision.Reply
	problem                   error
	fallback                  groupControl
}

func newTree(command *exec.Cmd) (treeControl, error) {
	// Preserve exec's original lookup errors rather than launching a helper.
	if command.Err != nil {
		return newGroupTree(command)
	}
	image, err := reaperImage()
	if err != nil {
		return nil, err
	}
	c := &reaperControl{files: []*os.File{image}, ready: make(chan struct{}), settled: make(chan struct{})}
	fail := func(err error) (treeControl, error) { c.close(); return nil, err }
	pipe := func() (*os.File, *os.File, error) {
		r, w, e := os.Pipe()
		if e == nil {
			c.files = append(c.files, r, w)
		}
		return r, w, e
	}
	requestR, requestW, err := pipe()
	if err != nil {
		return fail(err)
	}
	replyR, replyW, err := pipe()
	if err != nil {
		return fail(err)
	}
	cancelR, cancelW, err := pipe()
	if err != nil {
		return fail(err)
	}
	var random [24]byte
	if _, err := rand.Read(random[:]); err != nil {
		return fail(err)
	}
	nonce := hex.EncodeToString(random[:])
	c.request = supervision.Request{Version: supervision.Version, Nonce: nonce, Executable: command.Path, Args: command.Args, Env: command.Env, Combined: command.Stdout == command.Stderr}
	c.requestW, c.replyR, c.cancelW = requestW, replyR, cancelW
	command.Path = "/proc/self/fd/3"
	command.Args = []string{"fireweed-process-reaper", nonce}
	// This standalone helper must not inherit loader/runtime diagnostics or
	// arbitrary initializers from the executable which called Run/Capture.
	command.Env = []string{}
	command.ExtraFiles = []*os.File{image, requestR, replyW, cancelR}
	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	return c, nil
}

func (c *reaperControl) started(ctx context.Context, command *exec.Cmd) error {
	_ = c.fallback.started(ctx, command)
	for _, f := range command.ExtraFiles {
		_ = f.Close()
	}
	go func() {
		decoder := gob.NewDecoder(c.replyR)
		var problem error
		for {
			var reply supervision.Reply
			if err := decoder.Decode(&reply); err != nil {
				problem = fmt.Errorf("read process supervisor: %w", err)
				break
			}
			if reply.Version != supervision.Version || reply.Nonce != c.request.Nonce || reply.Ready == reply.Final {
				problem = errors.New("invalid process supervisor response")
				break
			}
			if reply.Ready {
				c.endOnce.Do(func() { close(c.ready) })
				continue
			}
			c.mu.Lock()
			c.reply = reply
			c.mu.Unlock()
			break
		}
		c.mu.Lock()
		c.problem = errors.Join(c.problem, problem)
		c.mu.Unlock()
		close(c.settled)
	}()
	written := make(chan error, 1)
	go func() { err := gob.NewEncoder(c.requestW).Encode(c.request); _ = c.requestW.Close(); written <- err }()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-c.settled:
		select {
		case <-c.ready:
		default:
			c.mu.Lock()
			defer c.mu.Unlock()
			if c.problem != nil {
				return c.problem
			}
			if c.reply.LaunchError {
				return c.replyError()
			}
			return errors.New("process supervisor ended before readiness")
		}
	case <-c.ready:
	}
	select {
	case err := <-written:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (c *reaperControl) stop(command *exec.Cmd) error {
	if command.Process == nil {
		return os.ErrProcessDone
	}
	c.cancelOnce.Do(func() { _ = c.cancelW.Close() })
	select {
	case <-c.settled:
		c.mu.Lock()
		final := c.reply.Final
		c.mu.Unlock()
		if final {
			return nil
		}
	case <-time.After(5 * time.Second):
	}
	// A failed helper must not strand descendants. Freeze its still-owned
	// tree before terminating the helper itself, retaining normal PID checks.
	return c.fallback.stop(command)
}

func (c *reaperControl) replyError() error {
	if c.reply.Error == "" {
		return nil
	}
	if c.reply.Errno != 0 {
		return fmt.Errorf("%s: %w", c.reply.Error, syscall.Errno(c.reply.Errno))
	}
	return errors.New(c.reply.Error)
}
func (c *reaperControl) outcome(waitErr error) (int, error) {
	<-c.settled
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.problem != nil {
		return 2, c.problem
	}
	if !c.reply.Final {
		return 2, errors.New("process supervisor did not report an outcome")
	}
	err := c.replyError()
	if c.reply.Exit != 0 && err == nil {
		err = fmt.Errorf("executable exited with status %d", c.reply.Exit)
	}
	if c.reply.Exit == 0 && waitErr != nil {
		return 2, waitErr
	}
	return c.reply.Exit, err
}
func (c *reaperControl) close() {
	for _, f := range c.files {
		_ = f.Close()
	}
}

func reaperImage() (*os.File, error) {
	reader, err := gzip.NewReader(bytes.NewReader(compressedReaper))
	if err != nil {
		return nil, fmt.Errorf("read embedded process supervisor: %w", err)
	}
	defer reader.Close()
	// memfd avoids executable files in the caller's checkout or temporary
	// directories, and works with noexec mounts. Only FD3 crosses exec.
	numbers := map[string]uintptr{"386": 356, "amd64": 319, "arm": 385, "arm64": 279, "loong64": 279, "mips": 4354, "mipsle": 4354, "mips64": 5314, "mips64le": 5314, "ppc64": 360, "ppc64le": 360, "riscv64": 279, "s390x": 350}
	number, ok := numbers[runtime.GOARCH]
	if !ok {
		return nil, errors.New("unsupported Linux process supervisor architecture")
	}
	name, _ := syscall.BytePtrFromString("fireweed-process-reaper")
	fd, _, errno := syscall.Syscall(number, uintptr(unsafe.Pointer(name)), 1, 0)
	runtime.KeepAlive(name)
	if errno != 0 {
		return nil, fmt.Errorf("allocate process supervisor: %w", errno)
	}
	file := os.NewFile(fd, "fireweed-process-reaper")
	if _, err := io.Copy(file, reader); err != nil {
		file.Close()
		return nil, err
	}
	if _, err := file.Seek(0, 0); err != nil {
		file.Close()
		return nil, err
	}
	return file, nil
}
