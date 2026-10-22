package evals

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

// This uses the established test-only native interception fixture. The compiled
// subject imports public workflow contracts; the guard is never product code.
func Test_FLOW007PolicyNoIO(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("native I/O interception runs on macOS; portable policy cases run elsewhere")
	}
	root := t.TempDir()
	sourceRoot, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	write := func(name, text string) {
		t.Helper()
		if err := os.WriteFile(filepath.Join(root, name), []byte(text), 0644); err != nil {
			t.Fatal(err)
		}
	}
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
	write("guard.c", guard)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	lib := filepath.Join(root, "libfwguard.dylib")
	cmd := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", lib, filepath.Join(root, "guard.c"), "-Wl,-install_name,"+lib)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("compile guard: %v %s", err, out)
	}
	program := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -lfwguard
void fw_guard_arm(int mode,const char*p);
*/
import "C"
import("os";"net";"time";"reflect";"encoding/json";c "github.com/lurkshark/fireweed/packages/contracts";g "github.com/lurkshark/fireweed/packages/gitstore";p "github.com/lurkshark/fireweed/packages/project";w "github.com/lurkshark/fireweed/packages/workflow")
func main(){
 mode,target,address:=os.Args[1],os.Args[2],os.Args[3]
 m:=c.Module{Name:"alpha",IDPrefix:"ALP",Dir:"units/a",Zones:c.Zones{Evals:[]string{"evals/**"},Disposable:[]string{"impl/**"},AppendOnly:[]string{"migrations/**"}}}
 s:=p.Snapshot{Root:target,Config:c.Config{WorkerWritable:[]string{"shared/**"},Hooks:c.Hooks{Eval:"hook-trap",Boundaries:"hook-trap"},Agent:c.Agent{Command:[]string{"agent-trap"}}},Modules:[]c.Module{m}}
 ch:=c.Change{Module:"alpha",Run:"run",Phase:"impl",Path:"changes/0001-native.md",Set:map[string]any{"interface":[]any{"api.go"}},Document:c.Document{Sections:[]c.Section{{Name:"Interface",Body:"Retained contract."}}}}
 edits:=[]g.Edit{{Status:"M",Path:"units/a/impl/code.go"}}
 if mode=="denied"{edits=[]g.Edit{{Status:"R",OldPath:"units/a/migrations/old.sql",Path:"units/a/evals/new.go"}}}
 if mode=="invalid"{ch.Run=""};if mode=="malformed"{edits[0].Path="../escape"}
 before,_:=json.Marshal([]any{s,m,ch,edits});env:=os.Environ()
 C.fw_guard_arm(2,nil)
 switch mode {
 case "read-control":os.ReadFile(target)
 case "stat-control":os.Stat(target)
 case "readlink-control":os.Readlink(target)
 case "write-control":os.WriteFile(target,[]byte("changed"),0644)
 case "directory-control":os.ReadDir(target)
 case "mkdir-control":os.Mkdir(target+"-new",0755)
 case "unlink-control":os.Remove(target)
 case "network-control":net.DialTimeout("tcp",address,time.Second)
 default:
  ds:=w.CheckEdits(s,m,ch,edits);if ds==nil{os.Exit(10)}
  counts:=map[string]int{};for _,d:=range ds{if d.Severity!="error"||d.Message==""{os.Exit(11)};counts[d.Code]++}
  switch mode{case "allowed":if len(ds)!=0{os.Exit(12)};case "denied":if len(ds)!=2||counts["H03"]!=1||counts["P2-01"]!=1{os.Exit(13)};case "invalid":if len(ds)!=1||counts["P02"]!=1{os.Exit(14)};case "malformed":if len(ds)!=1||counts["P2-01"]!=1{os.Exit(15)}}
  after,_:=json.Marshal([]any{s,m,ch,edits});if string(before)!=string(after)||!reflect.DeepEqual(env,os.Environ()){os.Exit(16)}
  C.fw_guard_arm(0,nil);return
 }
 C.fw_guard_arm(0,nil);os.Exit(70)
}`, root)
	write("probe.go", program)
	binary := filepath.Join(root, "probe")
	cmd = exec.CommandContext(ctx, "go", "build", "-o", binary, filepath.Join(root, "probe.go"))
	cmd.Dir = sourceRoot
	cmd.Env = append(os.Environ(), "CGO_ENABLED=1", "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build public subject: %v %s", err, out)
	}
	target := filepath.Join(root, "sentinel")
	if err := os.WriteFile(target, []byte("retain"), 0644); err != nil {
		t.Fatal(err)
	}
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	for _, mode := range []string{"read-control", "stat-control", "readlink-control", "write-control", "directory-control", "mkdir-control", "unlink-control", "network-control"} {
		t.Run(mode, func(t *testing.T) {
			cmd := exec.CommandContext(ctx, binary, mode, target, listener.Addr().String())
			out, err := cmd.CombinedOutput()
			exit, ok := err.(*exec.ExitError)
			if !ok || exit.ExitCode() != 42 {
				t.Fatalf("native control failed to intercept attempt despite ignored error: %v %s", err, out)
			}
		})
	}
	for _, mode := range []string{"allowed", "denied", "invalid", "malformed"} {
		t.Run(mode, func(t *testing.T) {
			cmd := exec.CommandContext(ctx, binary, mode, target, listener.Addr().String())
			cmd.Dir = root
			if out, err := cmd.CombinedOutput(); err != nil {
				t.Fatalf("public pure policy failed or attempted I/O: %v %s", err, out)
			}
		})
	}
	if b, err := os.ReadFile(target); err != nil || string(b) != "retain" {
		t.Fatalf("changed sentinel: %q %v", b, err)
	}
}
