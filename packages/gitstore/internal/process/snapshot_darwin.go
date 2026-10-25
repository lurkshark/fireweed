package process

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
	"unsafe"
)

// Darwin's proc_bsdinfo ABI is shared by arm64 and amd64. Creation time
// distinguishes the observed process from a later reuse of its PID.
type bsdProcessInfo struct {
	fields    [12]uint32
	names     [48]byte
	tail      [6]uint32
	sec, usec uint64
}

type lifetimeInfo struct {
	uuid                   [16]byte
	unique, parent         uint64
	version, parentVersion int32
	reserved               [2]uint64
}

func nativeProcess(pid int) (processRef, error) {
	var combined struct {
		bsd      bsdProcessInfo
		lifetime lifetimeInfo
	}
	// Include zombies: started() runs before our Wait, so even an already
	// exited child still supplies its exact lifetime identity safely.
	n, _, errno := syscall.Syscall6(syscall.SYS_PROC_INFO, 2, uintptr(pid), 18, 1, uintptr(unsafe.Pointer(&combined)), unsafe.Sizeof(combined))
	runtime.KeepAlive(&combined)
	info := combined.bsd
	runtime.KeepAlive(&info)
	if errno != 0 {
		return processRef{}, errno
	}
	if n == 0 {
		return processRef{}, os.ErrProcessDone
	}
	if n != unsafe.Sizeof(combined) || int(info.fields[3]) != pid {
		return processRef{}, fmt.Errorf("incomplete process metadata for %d", pid)
	}
	return processRef{pid: pid, parent: int(info.fields[4]), birth: [2]uint64{info.sec, info.usec}, unique: combined.lifetime.unique, parentUnique: combined.lifetime.parent, stopped: info.fields[1] == 4, dead: info.fields[1] == 5}, nil
}

func nativeChildren(pid int) ([]int, error) {
	return nativePIDs(6, pid)
}

func nativePIDs(kind, pid int) ([]int, error) {
	// PROC_INFO_CALL_LISTPIDS / PROC_PPID_ONLY: no external command or
	// repository input is consulted while inspecting the stopped parent.
	for size := 32; size <= 1<<20; size *= 2 {
		pids := make([]int32, size)
		n, _, errno := syscall.Syscall6(syscall.SYS_PROC_INFO, 1, uintptr(kind), uintptr(pid), 0, uintptr(unsafe.Pointer(&pids[0])), uintptr(len(pids)*4))
		runtime.KeepAlive(pids)
		if errno != 0 {
			return nil, errno
		}
		if n%4 != 0 || n > uintptr(len(pids)*4) {
			return nil, fmt.Errorf("invalid child process inventory for %d", pid)
		}
		if n == uintptr(len(pids)*4) {
			continue
		}
		result := make([]int, 0, int(n)/4)
		for _, child := range pids[:int(n)/4] {
			if child > 0 {
				result = append(result, int(child))
			}
		}
		return result, nil
	}
	return nil, fmt.Errorf("child process inventory exceeded capacity for %d", pid)
}

func nativeOwned(ancestors []processRef) ([]processRef, error) {
	identities := map[uint64]bool{}
	for _, ref := range ancestors {
		if ref.unique != 0 {
			identities[ref.unique] = true
		}
	}
	if len(identities) == 0 {
		return nil, nil
	}
	pids, err := nativePIDs(1, 0)
	if err != nil {
		return nil, err
	}
	refs := make([]processRef, 0, len(pids))
	for _, pid := range pids {
		ref, err := nativeProcess(pid)
		if err == nil {
			refs = append(refs, ref)
		}
	}
	result, seen := []processRef{}, map[uint64]bool{}
	for {
		added := false
		for _, ref := range refs {
			if ref.unique == 0 || seen[ref.unique] || !identities[ref.parentUnique] {
				continue
			}
			seen[ref.unique], identities[ref.unique], added = true, true, true
			result = append(result, ref)
		}
		if !added {
			return result, nil
		}
	}
}
