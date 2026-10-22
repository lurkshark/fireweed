package evals

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func Test_FLOW012_FLOW013_FLOW015EscapingReadGuard(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("native post-startup read/network interception runs on macOS")
	}
	f := worktreeFixture(t, "impl", "sha1", "escape")
	outside := t.TempDir()
	private := filepath.Join(outside, "private")
	if err := os.WriteFile(private, []byte("PRIVATE-UNTOUCHED"), 0644); err != nil {
		t.Fatal(err)
	}
	private = stateCanonicalDirectory(t, private)
	full := filepath.Join(f.Root, f.Path)
	if err := os.Remove(full); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(private, full); err != nil {
		t.Fatal(err)
	}
	temp := t.TempDir()
	record := map[string]any{"Snapshot": f.Snapshot, "Module": f.Module, "Claim": f.Claim, "Path": f.Path, "Object": f.Claim.Object}
	raw, err := json.Marshal(record)
	if err != nil {
		t.Fatal(err)
	}
	input := filepath.Join(temp, "input.json")
	if err = os.WriteFile(input, raw, 0644); err != nil {
		t.Fatal(err)
	}
	sourceRoot, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	_, marker := worktreeTraps(t, false)
	guard := `#include <sys/types.h>
#include <sys/socket.h>
#include <sys/stat.h>
#include <unistd.h>
#include <fcntl.h>
#include <stdarg.h>
#include <stdint.h>
#include <string.h>
#include <dirent.h>
#include <stdio.h>
#include <sys/attr.h>
static int armed=0;static const char*forbidden=0;
void fw_guard_arm(int mode,const char*p){armed=mode;forbidden=p;}
static void pathcheck(const char*p){if(armed==2||(armed&&forbidden&&p&&(strcmp(p,forbidden)==0||strstr(p,"units/escape/SPEC.md")||strstr(p,"changes/0001-escape.md")||strstr(p,"changes/archive/0001-escape.md"))))_exit(42);}
static void iocheck(void){if(armed==2)_exit(42);}
static void netcheck(void){if(armed)_exit(42);}
static int fw_open(const char*p,int f,...){pathcheck(p);mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return open(p,f,m);}
static int fw_openat(int d,const char*p,int f,...){pathcheck(p);mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return openat(d,p,f,m);}
static ssize_t fw_read(int d,void*b,size_t n){iocheck();return read(d,b,n);}
static ssize_t fw_write(int d,const void*b,size_t n){iocheck();return write(d,b,n);}
static int fw_stat(const char*p,struct stat*b){iocheck();return stat(p,b);}
static int fw_lstat(const char*p,struct stat*b){iocheck();return lstat(p,b);}
static int fw_socket(int d,int t,int p){netcheck();return socket(d,t,p);}
static int fw_connect(int s,const struct sockaddr*a,socklen_t n){netcheck();return connect(s,a,n);}
static ssize_t fw_readlink(const char*p,char*b,size_t n){iocheck();return readlink(p,b,n);}
static int fw_access(const char*p,int m){iocheck();return access(p,m);}
static int fw_fstat(int d,struct stat*b){iocheck();return fstat(d,b);}
static int fw_getattrlist(const char*p,struct attrlist*a,void*b,size_t n,unsigned long o){iocheck();return getattrlist(p,a,b,n,o);}
static int fw_mkdir(const char*p,mode_t m){iocheck();return mkdir(p,m);}
static int fw_unlink(const char*p){iocheck();return unlink(p);}
static int fw_rmdir(const char*p){iocheck();return rmdir(p);}
static int fw_rename(const char*p,const char*q){iocheck();return rename(p,q);}
static int fw_symlink(const char*p,const char*q){iocheck();return symlink(p,q);}
static int fw_chmod(const char*p,mode_t m){iocheck();return chmod(p,m);}
static char*fw_getcwd(char*b,size_t n){iocheck();return getcwd(b,n);}
static DIR*fw_opendir(const char*p){iocheck();return opendir(p);}
static struct dirent*fw_readdir(DIR*d){iocheck();return readdir(d);}
#define I(r,o) __attribute__((used)) static struct{const void*replacement;const void*original;} fw_interpose_##o __attribute__((section("__DATA,__interpose")))={(const void*)(uintptr_t)&r,(const void*)(uintptr_t)&o};
I(fw_open,open) I(fw_openat,openat) I(fw_read,read) I(fw_write,write) I(fw_stat,stat) I(fw_lstat,lstat) I(fw_socket,socket) I(fw_connect,connect)
I(fw_readlink,readlink) I(fw_access,access) I(fw_fstat,fstat) I(fw_getattrlist,getattrlist)
I(fw_mkdir,mkdir) I(fw_unlink,unlink) I(fw_rmdir,rmdir) I(fw_rename,rename) I(fw_symlink,symlink) I(fw_chmod,chmod) I(fw_getcwd,getcwd) I(fw_opendir,opendir) I(fw_readdir,readdir)
`
	if err := os.WriteFile(filepath.Join(temp, "guard.c"), []byte(guard), 0644); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	lib := filepath.Join(temp, "libfwguard.dylib")
	cmd := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", lib, filepath.Join(temp, "guard.c"), "-Wl,-install_name,"+lib)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("compile read guard: %v %s", err, out)
	}
	program := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -lfwguard
void fw_guard_arm(int mode,const char*p);
*/
import "C"
import("encoding/json";"os";"path/filepath";"net";"time";"strings";p "github.com/lurkshark/fireweed/packages/project";c "github.com/lurkshark/fireweed/packages/contracts";w "github.com/lurkshark/fireweed/packages/workflow")
func main(){
 var in struct{Snapshot p.Snapshot;Module c.Module;Claim w.Claim;Path,Object string}
 raw,err:=os.ReadFile(os.Args[2]);if err!=nil{os.Exit(80)};if err=json.Unmarshal(raw,&in);err!=nil{os.Exit(81)};in.Claim.Object=in.Object
 loaded,err:=w.LoadClaim(in.Snapshot.Root,in.Claim.Change);if err!=nil||loaded!=in.Claim{os.Exit(82)}
 if os.Args[1]=="fixture-control"{return}
 if os.Args[1]=="bad-claim"{in.Claim.Object=strings.Repeat("f",len(in.Object));if in.Claim.Object==loaded.Object{os.Exit(83)}}
 C.fw_guard_arm(1,C.CString(os.Args[3]))
 switch os.Args[1]{
 case "read-private-control":os.ReadFile(os.Args[3]);os.Exit(70)
 case "read-link-control":os.ReadFile(filepath.Join(in.Snapshot.Root,in.Path));os.Exit(71)
 case "network-control":net.DialTimeout("tcp",os.Args[4],time.Second);os.Exit(72)
 default:
  result:=w.CheckWorktree(in.Snapshot,in.Module,in.Path,in.Claim)
  if result.Edits==nil||result.Diagnostics==nil||len(result.Diagnostics)==0{os.Exit(10)}
  code:="P2-01";if os.Args[1]=="bad-claim"{code="P02";if len(result.Edits)!=0||len(result.Diagnostics)!=1{os.Exit(12)}}
  for _,d:=range result.Diagnostics{if d.Code!=code||d.Path!=in.Path||d.Severity!="error"||d.Message==""{os.Exit(11)}}
 }
 C.fw_guard_arm(0,nil)
}`, temp)
	probeSource := filepath.Join(temp, "probe.go")
	if err := os.WriteFile(probeSource, []byte(program), 0644); err != nil {
		t.Fatal(err)
	}
	binary := filepath.Join(temp, "probe")
	cmd = exec.CommandContext(ctx, "go", "build", "-o", binary, probeSource)
	cmd.Dir = sourceRoot
	cmd.Env = append(os.Environ(), "CGO_ENABLED=1", "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build public worktree probe: %v %s", err, out)
	}
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	run := func(mode string) ([]byte, error) {
		cmd := exec.CommandContext(ctx, binary, mode, input, private, listener.Addr().String())
		cmd.Dir = f.Root
		return cmd.CombinedOutput()
	}
	t.Run("fixture-control", func(t *testing.T) {
		if out, err := run("fixture-control"); err != nil {
			t.Fatalf("serialized approved-state fixture failed: %v %s", err, out)
		}
	})
	for _, mode := range []string{"read-private-control", "read-link-control", "network-control"} {
		t.Run(mode, func(t *testing.T) {
			out, err := run(mode)
			exit, ok := err.(*exec.ExitError)
			if !ok || exit.ExitCode() != 42 {
				t.Fatalf("native guard did not detect ignored-error attempt: %v %s", err, out)
			}
		})
	}
	t.Run("safe-rejection", func(t *testing.T) {
		if out, err := run("check"); err != nil {
			t.Fatalf("worktree inspected private contents or failed to reject escape: %v %s", err, out)
		}
	})
	t.Run("bad-claim-short-circuits-current-read", func(t *testing.T) {
		if out, err := run("bad-claim"); err != nil {
			t.Fatalf("rejected claim read current/private bytes or failed short-circuit contract: %v %s", err, out)
		}
	})
	if raw, err := os.ReadFile(private); err != nil || string(raw) != "PRIVATE-UNTOUCHED" {
		t.Fatal("changed private outside file")
	}
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Fatal("worktree probe executed a forbidden hook/agent/Git/network command")
	}
}
