//go:build !cgo

package native

import (
	"runtime"
	"syscall"
	"unsafe"
)

//go:linkname call9 syscall.syscall9
func call9(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (uintptr, uintptr, syscall.Errno)
func spawnAddress() uintptr

func Spawn(path, cwd, nonce *byte, fds [5]int, pid *int32) int {
	result, _, _ := call9(spawnAddress(), uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(cwd)), uintptr(unsafe.Pointer(nonce)), uintptr(fds[0]), uintptr(fds[1]), uintptr(fds[2]), uintptr(fds[3]), uintptr(fds[4]), uintptr(unsafe.Pointer(pid)))
	runtime.KeepAlive(path)
	runtime.KeepAlive(cwd)
	runtime.KeepAlive(nonce)
	runtime.KeepAlive(pid)
	return int(result)
}
