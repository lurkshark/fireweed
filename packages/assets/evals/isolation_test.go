package evals

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

// This probe uses only the public API. A test-only interposer is armed after
// the OS/Go loader starts; any later file/network attempt exits the process,
// even if a caller would otherwise ignore the error. Product builds remain Go-only.
func Test_AST001_AST002SandboxIndependence(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("macOS native/interposition probe; portable content checks run on every platform")
	}
	root, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	root, err = filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatal(err)
	}
	temp := t.TempDir()
	source := filepath.Join(temp, "probe.go")
	binary := filepath.Join(temp, "probe")
	library := filepath.Join(temp, "libfwguard.dylib")
	csource := filepath.Join(temp, "guard.c")
	guard := `#include <sys/types.h>
#include <sys/socket.h>
#include <sys/stat.h>
#include <unistd.h>
#include <fcntl.h>
#include <stdarg.h>
#include <stdint.h>
#include <stdatomic.h>
static _Atomic int armed=0;
void fw_guard_arm(void){atomic_store(&armed,1);}
static void check(void){if(atomic_load(&armed))_exit(42);}
static int fw_open(const char*p,int f,...){check();mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return open(p,f,m);}
static int fw_openat(int d,const char*p,int f,...){check();mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return openat(d,p,f,m);}
static ssize_t fw_read(int d,void*b,size_t n){check();return read(d,b,n);}
static ssize_t fw_pread(int d,void*b,size_t n,off_t o){check();return pread(d,b,n,o);}
static int fw_socket(int d,int t,int p){check();return socket(d,t,p);}
static int fw_connect(int s,const struct sockaddr*a,socklen_t n){check();return connect(s,a,n);}
static ssize_t fw_recv(int s,void*b,size_t n,int f){check();return recv(s,b,n,f);}
static ssize_t fw_recvfrom(int s,void*b,size_t n,int f,struct sockaddr*a,socklen_t*l){check();return recvfrom(s,b,n,f,a,l);}
static ssize_t fw_recvmsg(int s,struct msghdr*m,int f){check();return recvmsg(s,m,f);}
static int fw_stat(const char*p,struct stat*b){check();return stat(p,b);}
static int fw_lstat(const char*p,struct stat*b){check();return lstat(p,b);}
#define I(r,o) __attribute__((used)) static struct{const void*replacement;const void*original;} fw_interpose_##o __attribute__((section("__DATA,__interpose")))={(const void*)(uintptr_t)&r,(const void*)(uintptr_t)&o};
I(fw_open,open) I(fw_openat,openat) I(fw_read,read) I(fw_pread,pread)
I(fw_socket,socket) I(fw_connect,connect) I(fw_recv,recv) I(fw_recvfrom,recvfrom) I(fw_recvmsg,recvmsg)
I(fw_stat,stat) I(fw_lstat,lstat)
`
	if err := os.WriteFile(csource, []byte(guard), 0644); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	compile := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", library, csource, "-Wl,-install_name,"+library)
	if output, err := compile.CombinedOutput(); err != nil {
		t.Fatalf("compile test-only I/O interposer: %v %s", err, output)
	}
	program := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -lfwguard
void fw_guard_arm(void);
*/
import "C"
import("fmt";"net";"os";"time"; a "github.com/lurkshark/fireweed/packages/assets")
func main(){
 if len(os.Args)>1&&os.Args[1]=="--guard"{
  if _,err:=os.ReadFile(os.Args[2]);err==nil{os.Exit(3)}
  if _,err:=os.ReadFile(os.Args[4]);err==nil{os.Exit(9)}
  if conn,err:=net.DialTimeout("tcp",os.Args[3],time.Second);err==nil{conn.Close();os.Exit(4)}
  fmt.Println("DENIAL_GUARD_OK");return
 }
 C.fw_guard_arm()
 if len(os.Args)>1&&os.Args[1]=="--file-attempt"{os.ReadFile(os.Args[2]);os.Exit(10)}
 if len(os.Args)>1&&os.Args[1]=="--network-attempt"{net.DialTimeout("tcp",os.Args[2],time.Second);os.Exit(11)}
 info:=a.Info();if info.Version!="0.4.0"||info.CurrentFormat!=1||info.PreviousFormat!=0{os.Exit(5)}
 for _,name:=range a.Playbooks(){text,err:=a.Playbook(name);if err!=nil||text.Text==""||text.Hash==""{os.Exit(6)}}
 if len(a.Playbooks())!=5{os.Exit(7)}
 for _,code:=range a.Codes(){if text,err:=a.Explain(code);err!=nil||text==""{os.Exit(8)}}
 for _,name:=range []string{"","missing","../author-change","AUDIT"}{if _,err:=a.Playbook(name);err==nil{os.Exit(12)}}
 for _,code:=range []string{"","unknown","C00","H05","c01"}{if _,err:=a.Explain(code);err==nil{os.Exit(13)}}
 fmt.Println("EMBEDDED_ASSETS_OK")
}`, temp)
	if err := os.WriteFile(source, []byte(program), 0644); err != nil {
		t.Fatal(err)
	}
	build := exec.CommandContext(ctx, "go", "build", "-o", binary, source)
	build.Dir = root
	build.Env = append(os.Environ(), "CGO_ENABLED=1", "GOPROXY=off", "GOTOOLCHAIN=local")
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build public-API probe: %v %s", err, output)
	}
	binary, err = filepath.EvalSymlinks(binary)
	if err != nil {
		t.Fatal(err)
	}
	sentinel := filepath.Join(temp, "forbidden-source-data")
	if err := os.WriteFile(sentinel, []byte("private data"), 0644); err != nil {
		t.Fatal(err)
	}
	sentinel, err = filepath.EvalSymlinks(sentinel)
	if err != nil {
		t.Fatal(err)
	}
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	for _, attempt := range [][]string{{"--file-attempt", sentinel}, {"--file-attempt", filepath.Join(root, "go.mod")}, {"--network-attempt", listener.Addr().String()}} {
		output, err := exec.CommandContext(ctx, binary, attempt...).CombinedOutput()
		exit, ok := err.(*exec.ExitError)
		if !ok || exit.ExitCode() != 42 {
			t.Fatalf("I/O attempt guard ineffective %v: %v %s", attempt, err, output)
		}
	}
	profile := "(version 1)(allow default)(deny file-read-data (subpath " + strconv.Quote(filepath.Dir(root)) + ") (literal " + strconv.Quote(sentinel) + "))(deny network*)"
	if err := exec.CommandContext(ctx, binary, "--guard", sentinel, listener.Addr().String(), filepath.Join(root, "go.mod")).Run(); err == nil {
		t.Fatal("negative native control allowed file access")
	}
	guarded := exec.CommandContext(ctx, "/usr/bin/sandbox-exec", "-p", profile, binary, "--guard", sentinel, listener.Addr().String(), filepath.Join(root, "go.mod"))
	guarded.Dir = temp
	if output, err := guarded.CombinedOutput(); err != nil || !strings.Contains(string(output), "DENIAL_GUARD_OK") {
		t.Fatalf("native guard unavailable: %v %s", err, output)
	}
	command := exec.CommandContext(ctx, "/usr/bin/sandbox-exec", "-p", profile, binary)
	command.Dir = temp
	if output, err := command.CombinedOutput(); err != nil || !strings.Contains(string(output), "EMBEDDED_ASSETS_OK") {
		t.Fatalf("assets depend on runtime I/O: %v %s", err, output)
	}
}
