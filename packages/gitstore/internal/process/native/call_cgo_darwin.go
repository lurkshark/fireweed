//go:build cgo

package native

/*
int fireweed_spawn_anchor(const char*,const char*,const char*,int,int,int,int,int,int*);
*/
import "C"
import "unsafe"

func Spawn(path, cwd, nonce *byte, fds [5]int, pid *int32) int {
	return int(C.fireweed_spawn_anchor((*C.char)(unsafe.Pointer(path)), (*C.char)(unsafe.Pointer(cwd)), (*C.char)(unsafe.Pointer(nonce)), C.int(fds[0]), C.int(fds[1]), C.int(fds[2]), C.int(fds[3]), C.int(fds[4]), (*C.int)(unsafe.Pointer(pid))))
}
