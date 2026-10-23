package evals

import (
	"context"
	"encoding/json"
	"fmt"
	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

const eligibilityNativeGuard = `#include <sys/types.h>
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
void fw_eligible_arm(int mode,const char*p){armed=mode;forbidden=p;}
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

func Test_FLOW024_FLOW026_FLOW027EligibleNativeNoIO(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("macOS post-startup native guard; portable pure eligibility assertions run separately")
	}
	root, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	worktreePut(t, root, "guard.c", eligibilityNativeGuard)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	lib := filepath.Join(root, "libeligibleguard.dylib")
	cmd := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", lib, filepath.Join(root, "guard.c"), "-Wl,-install_name,"+lib)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("native guard compile %v %s", err, out)
	}
	program := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -leligibleguard
void fw_eligible_arm(int m,const char*p);
*/
import "C"
import("os";"os/exec";"net";"time";"fmt";"encoding/json";"reflect";"strings";p "github.com/lurkshark/fireweed/packages/project";c "github.com/lurkshark/fireweed/packages/contracts";w "github.com/lurkshark/fireweed/packages/workflow")
type record struct{Snapshot p.Snapshot;Slug,Private,Address string;Want,Error bool}
func main(){
 raw,err:=os.ReadFile(os.Args[2]);if err!=nil{os.Exit(80)};var in record;if json.Unmarshal(raw,&in)!=nil{os.Exit(81)}
 mode:=os.Args[1];before,_:=json.Marshal(in.Snapshot);cwd,err:=os.Getwd();if err!=nil{os.Exit(82)};env:=os.Environ()
 C.fw_eligible_arm(2,nil)
 switch mode{
 case "read-control":os.ReadFile(in.Private)
 case "missing-control":os.ReadFile(in.Private+"/missing")
 case "stat-control":os.Stat(in.Private)
 case "cwd-control":os.Getwd()
 case "directory-control":os.ReadDir(in.Private)
 case "stdin-control":os.Stdin.Read(make([]byte,1))
 case "prompt-control":fmt.Fprint(os.Stdout,"prompt")
 case "write-control":os.WriteFile(in.Private,[]byte("bad"),0644)
 case "network-control":net.DialTimeout("tcp",in.Address,time.Second)
 case "git-control":exec.Command("git","status").Run()
 case "hook-control":exec.Command("hook-trap").Run()
 case "source-control":
  s,err:=p.Project(in.Snapshot,in.Slug);if err!=nil{os.Exit(83)};var ch c.Change;for _,v:=range in.Snapshot.Changes{if v.Slug==in.Slug{ch=v}}
  found:=false;for _,m:=range s.Modules{if m.Name==ch.Module{m,err=p.ApplyChange(m,ch);if err!=nil||len(m.Specs)!=2{os.Exit(84)};found=true}}
  if !found{os.Exit(85)}
  C.fw_eligible_arm(0,nil);return
 default:
  got,err:=w.Eligible(in.Snapshot,in.Slug)
  if got!=in.Want||(err!=nil)!=in.Error{C.fw_eligible_arm(0,nil);fmt.Fprintf(os.Stderr,"got %%v,%%v want %%v error=%%v",got,err,in.Want,in.Error);os.Exit(10)}
  if err!=nil&&!strings.HasPrefix(err.Error(),"C05:"){os.Exit(11)}
  after,_:=json.Marshal(in.Snapshot);if string(before)!=string(after)||!reflect.DeepEqual(env,os.Environ()){os.Exit(12)}
  C.fw_eligible_arm(0,nil);afterCwd,err:=os.Getwd();if err!=nil||cwd!=afterCwd{os.Exit(13)};return
 }
 C.fw_eligible_arm(0,nil);os.Exit(70)
}`, root)
	worktreePut(t, root, "probe.go", program)
	source, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	binary := filepath.Join(root, "probe")
	cmd = exec.CommandContext(ctx, "go", "build", "-o", binary, filepath.Join(root, "probe.go"))
	cmd.Dir = source
	cmd.Env = append(os.Environ(), "CGO_ENABLED=1", "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("native public eligibility probe %v %s", err, out)
	}
	private := filepath.Join(root, "sentinel")
	worktreePut(t, root, "sentinel", "Untouched private bytes.")
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	type record struct {
		Snapshot               p.Snapshot
		Slug, Private, Address string
		Want, Error            bool
	}
	s := eligibleSource(t)
	s.Changes = []c.Change{eligibleChange(t, 1, "candidate", "delta", "set: {depends_on: []}\n", "## Added\n\n### new:retained [review]\nThe target MUST retain this local handle.\n")}
	in := record{Snapshot: s, Slug: "candidate", Private: private, Address: listener.Addr().String(), Want: true}
	run := func(t *testing.T, mode string, in record, want int) {
		t.Helper()
		raw, err := json.Marshal(in)
		if err != nil {
			t.Fatal(err)
		}
		worktreePut(t, root, "record.json", string(raw))
		cmd := exec.CommandContext(ctx, binary, mode, filepath.Join(root, "record.json"))
		cmd.Dir = root
		out, err := cmd.CombinedOutput()
		code := 0
		if err != nil {
			exit, ok := err.(*exec.ExitError)
			if !ok {
				t.Fatal(err)
			}
			code = exit.ExitCode()
		}
		if code != want {
			t.Fatalf("native %s exit%d want%d: %s", mode, code, want, out)
		}
	}
	for _, mode := range []string{"read-control", "missing-control", "stat-control", "cwd-control", "directory-control", "stdin-control", "prompt-control", "write-control", "network-control", "git-control", "hook-control"} {
		t.Run(mode, func(t *testing.T) { run(t, mode, in, 42) })
	}
	t.Run("source-control", func(t *testing.T) { run(t, "source-control", in, 0) })
	t.Run("eligible", func(t *testing.T) { run(t, "subject", in, 0) })
	t.Run("blocked", func(t *testing.T) {
		x := in
		x.Snapshot = eligibleCopy(t, in.Snapshot)
		x.Snapshot.Changes[0].Blocked = "waiting"
		x.Want = false
		run(t, "subject", x, 0)
	})
	t.Run("invalid-selector", func(t *testing.T) { x := in; x.Slug = "../bad"; x.Want, x.Error = false, true; run(t, "subject", x, 0) })
	t.Run("missing-selector", func(t *testing.T) { x := in; x.Slug = "absent"; x.Want, x.Error = false, true; run(t, "subject", x, 0) })
	t.Run("invalid-graph", func(t *testing.T) {
		x := in
		x.Snapshot = eligibleCopy(t, in.Snapshot)
		x.Snapshot.Modules[0].DependsOn = []string{"absent"}
		x.Want, x.Error = false, true
		run(t, "subject", x, 0)
	})
	if raw, err := os.ReadFile(private); err != nil || string(raw) != "Untouched private bytes." {
		t.Fatal("private sentinel changed", err)
	}
}
