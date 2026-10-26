package native

import (
	"encoding/binary"
	"encoding/hex"
	"os"
	"syscall"
)

// Fail closed if a linker/toolchain omits the early entry. Authenticate an
// anonymous request before recognizing the private invocation, so ordinary
// callers can still pass any argument vector to their own executable.
func init() {
	if len(os.Args) != 3 || os.Args[1] != "__fireweed_process_anchor_v1__" || len(os.Args[2]) != 48 {
		return
	}
	var info syscall.Stat_t
	if syscall.Fstat(3, &info) != nil || info.Mode&syscall.S_IFMT != 0 || info.Size < 64 || info.Size > 16<<20 {
		return
	}
	header, err := syscall.Mmap(3, 0, 64, syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		return
	}
	matches := binary.LittleEndian.Uint64(header[:8]) == 0x667770726f633031 && hex.EncodeToString(header[40:64]) == os.Args[2]
	_ = syscall.Munmap(header)
	if matches {
		os.Exit(125)
	}
}
