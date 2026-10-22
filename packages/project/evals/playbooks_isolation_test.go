package evals_test

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func Test_PRJ021_PRJ022_PRJ023NativeReadAndNetworkGuards(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("post-startup native read/network interception executes on macOS")
	}
	root := playbookRoot(t)
	outside := t.TempDir()
	put(t, outside, "private", "PRIVATE-UNCHANGED")
	private, err := filepath.EvalSymlinks(filepath.Join(outside, "private"))
	if err != nil {
		t.Fatal(err)
	}
	put(t, root, ".fireweed/playbooks/.keep", "")
	if err := os.Symlink(private, filepath.Join(root, ".fireweed/playbooks/audit.md")); err != nil {
		t.Fatal(err)
	}
	temp := t.TempDir()
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
static void pathcheck(const char*p){if(armed==2||(armed&&forbidden&&p&&(strcmp(p,forbidden)==0||strstr(p,"fireweed.yaml")||strstr(p,"playbooks/audit.md")||(armed==3&&strstr(p,"playbooks/evals.md")))))_exit(42);}
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
	put(t, temp, "guard.c", guard)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	lib := filepath.Join(temp, "libfwguard.dylib")
	cmd := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", lib, filepath.Join(temp, "guard.c"), "-Wl,-install_name,"+lib)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("compile guard %v %s", err, out)
	}
	program := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -lfwguard
void fw_guard_arm(int mode,const char*p);
*/
import "C"
import("os";"net";"time";"strings";"path/filepath";a "github.com/lurkshark/fireweed/packages/assets";p "github.com/lurkshark/fireweed/packages/project")
func main(){
 mode,root,private,address:=os.Args[1],os.Args[2],os.Args[3],os.Args[4]
 asset,err:=a.Playbook("evals");if err!=nil{os.Exit(80)}
 C.fw_guard_arm(1,C.CString(private))
 switch mode {
 case "private-control":os.ReadFile(private);os.Exit(70)
 case "config-control":os.ReadFile(filepath.Join(root,"fireweed.yaml"));os.Exit(71)
 case "unselected-control":os.ReadFile(filepath.Join(root,".fireweed/playbooks/audit.md"));os.Exit(72)
 case "network-control":net.DialTimeout("tcp",address,time.Second);os.Exit(73)
 case "selected-control":C.fw_guard_arm(3,C.CString(private));os.ReadFile(filepath.Join(root,".fireweed/playbooks/evals.md"));os.Exit(74)
 case "unknown":C.fw_guard_arm(2,C.CString(private));got,err:=p.ResolvePlaybook(root,"../impl");if err==nil||got!=(p.Playbook{}){os.Exit(10)}
 case "escape":C.fw_guard_arm(3,C.CString(private));got,err:=p.ResolvePlaybook(root,"evals");if err==nil||got!=(p.Playbook{})||!strings.Contains(err.Error(),"D05"){os.Exit(11)}
 default:got,err:=p.ResolvePlaybook(root,"evals");if err!=nil||got.Name!="evals"||got.Text!=asset.Text||got.Hash!=asset.Hash||got.EmbeddedHash!=asset.Hash||got.Overridden||got.Stale||got.BasedOn!=""{os.Exit(12)}
 }
 C.fw_guard_arm(0,nil)
}`, temp)
	put(t, temp, "probe.go", program)
	sourceRoot, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	binary := filepath.Join(temp, "probe")
	cmd = exec.CommandContext(ctx, "go", "build", "-o", binary, filepath.Join(temp, "probe.go"))
	cmd.Dir = sourceRoot
	cmd.Env = append(os.Environ(), "CGO_ENABLED=1", "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build public probe %v %s", err, out)
	}
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	run := func(mode string) ([]byte, error) {
		cmd := exec.CommandContext(ctx, binary, mode, root, private, listener.Addr().String())
		return cmd.CombinedOutput()
	}
	for _, mode := range []string{"private-control", "config-control", "unselected-control", "network-control"} {
		t.Run(mode, func(t *testing.T) {
			out, err := run(mode)
			exit, ok := err.(*exec.ExitError)
			if !ok || exit.ExitCode() != 42 {
				t.Fatalf("native ignored-error control %s not intercepted: %v %s", mode, err, out)
			}
		})
	}
	for _, mode := range []string{"embedded", "unknown"} {
		t.Run(mode, func(t *testing.T) {
			if out, err := run(mode); err != nil {
				t.Fatalf("unselected/config read or resolution failure: %v %s", err, out)
			}
		})
	}
	if err := os.Symlink(private, filepath.Join(root, ".fireweed/playbooks/evals.md")); err != nil {
		t.Fatal(err)
	}
	t.Run("selected-control", func(t *testing.T) {
		out, err := run("selected-control")
		exit, ok := err.(*exec.ExitError)
		if !ok || exit.ExitCode() != 42 {
			t.Fatalf("selected-link read control not intercepted: %v %s", err, out)
		}
	})
	t.Run("escape", func(t *testing.T) {
		if out, err := run("escape"); err != nil {
			t.Fatalf("unsafe selected read or rejection failure: %v %s", err, out)
		}
	})
	if raw, err := os.ReadFile(private); err != nil || string(raw) != "PRIVATE-UNCHANGED" {
		t.Fatal("changed outside private contents")
	}
}
