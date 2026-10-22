package evals_test

import (
	"context"
	"fmt"
	a "github.com/lurkshark/fireweed/packages/assets"
	p "github.com/lurkshark/fireweed/packages/project"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func Test_PRJ001RunningVersion(t *testing.T) {
	version := a.Info().Version
	root := fixture(t)
	put(t, root, "fireweed.yaml", strings.Replace(config, ">=0.4, <0.5", "="+version, 1))
	if _, err := p.Load(root); err != nil {
		t.Fatalf("running version %s: %v", version, err)
	}
	put(t, root, "fireweed.yaml", strings.Replace(config, ">=0.4, <0.5", ">"+version, 1))
	if _, err := p.Load(root); err == nil || !strings.Contains(err.Error(), "C01") {
		t.Fatalf("incompatible running version %s: %v", version, err)
	}
}
func shellQuote(s string) string                                  { return "'" + strings.ReplaceAll(s, "'", "'\"'\"'") + "'" }
func Test_PRJ001_PRJ002_PRJ004NoExternalIO(t *testing.T)          { runProjectIsolation(t, false) }
func Test_PRJ007_PRJ008_PRJ009_PRJ010ProjectionNoIO(t *testing.T) { runProjectIsolation(t, true) }
func runProjectIsolation(t *testing.T, projection bool) {
	if runtime.GOOS != "darwin" {
		t.Skip("Darwin syscall interposition probe; portable fixtures exercise content elsewhere")
	}
	sourceRoot, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	temp := t.TempDir()
	fixtureRoot := fixture(t)
	outside := t.TempDir()
	forbidden := filepath.Join(outside, "forbidden-read")
	put(t, outside, "forbidden-read", module("secret", "SEC", ""))
	forbidden, err = filepath.EvalSymlinks(forbidden)
	if err != nil {
		t.Fatal(err)
	}
	marker := filepath.Join(temp, "EXECUTED")
	put(t, fixtureRoot, "fireweed.yaml", config+"  setup: touch "+shellQuote(marker)+"\nagent:\n  command: [agent-trap]\naudit:\n  command: [agent-trap]\n")
	put(t, fixtureRoot, "units/a/SPEC.md", module("alpha", "ALP", ""))
	put(t, fixtureRoot, "units/escape/.keep", "")
	put(t, fixtureRoot, "changes/.keep", "")
	put(t, fixtureRoot, "changes/archive/.keep", "")
	if err = os.Symlink(forbidden, filepath.Join(fixtureRoot, "units/escape/SPEC.md")); err != nil {
		t.Fatal(err)
	}
	if err = os.Symlink(forbidden, filepath.Join(fixtureRoot, "changes/0001-escape.md")); err != nil {
		t.Fatal(err)
	}
	if err = os.Symlink(forbidden, filepath.Join(fixtureRoot, "changes/archive/0001-escape.md")); err != nil {
		t.Fatal(err)
	}
	put(t, fixtureRoot, ".gitignore", "ignored\n")
	put(t, fixtureRoot, "ignored", "KEEP-IGNORED-BYTES")
	git(t, fixtureRoot, "add", ".")
	git(t, fixtureRoot, "commit", "-m", "fixture")
	git(t, fixtureRoot, "remote", "add", "origin", "ssh://example.invalid/never-contact")
	before := git(t, fixtureRoot, "status", "--porcelain=v1", "--untracked-files=all") + git(t, fixtureRoot, "show-ref")
	trapDir := filepath.Join(temp, "traps")
	realGit, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	put(t, temp, "traps/git", "#!/bin/sh\ncase \" $* \" in *' fetch '*|*' push '*|*' ls-remote '*|*' clone '*) echo git-network >> "+shellQuote(marker)+";exit 99;;esac\nexec "+shellQuote(realGit)+" \"$@\"\n")
	for _, name := range []string{"agent-trap", "curl", "wget", "nc", "ssh", "codex", "fireweed", "fireweed-bootstrap"} {
		put(t, temp, "traps/"+name, "#!/bin/sh\necho "+name+" >> "+shellQuote(marker)+"\nexit 99\n")
	}
	entries, _ := os.ReadDir(trapDir)
	for _, e := range entries {
		if err = os.Chmod(filepath.Join(trapDir, e.Name()), 0755); err != nil {
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
	put(t, temp, "guard.c", guard)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	lib := filepath.Join(temp, "libfwguard.dylib")
	cmd := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", lib, filepath.Join(temp, "guard.c"), "-Wl,-install_name,"+lib)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("compile guard: %v %s", err, out)
	}
	program := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -lfwguard
void fw_guard_arm(int mode,const char*p);
*/
import "C"
import("os";"net";"time";p "github.com/lurkshark/fireweed/packages/project";c "github.com/lurkshark/fireweed/packages/contracts")
func main(){
 root,target:=os.Args[2],os.Args[3]
 C.fw_guard_arm(1,C.CString(target))
 switch os.Args[1]{case "read-control":os.ReadFile(target);os.Exit(70);case "network-control":net.DialTimeout("tcp",root,time.Second);os.Exit(71);case "check-control", "readlink-control", "stat-control", "write-control", "dir-control", "mkdir-control", "unlink-control":
 C.fw_guard_arm(2,C.CString(target))
 switch os.Args[1]{case "check-control":os.ReadFile(root);case "readlink-control":os.Readlink(target);case "stat-control":os.Stat(target);case "write-control":os.WriteFile(root,[]byte("bad"),0644);case "dir-control":os.ReadDir(root);case "mkdir-control":os.Mkdir(root,0755);case "unlink-control":os.Remove(root)};os.Exit(72)}
 s,err:=p.Load(root);if err!=nil{os.Exit(10)}
 if len(s.Modules)!=1||len(s.Changes)!=0||len(s.Archive)!=0{os.Exit(11)}
 found:=map[string]bool{};for _,d:=range s.Diagnostics{if d.Severity=="error"&&((d.Code=="C03"&&d.Path=="units/escape/SPEC.md")||(d.Code=="C05"&&(d.Path=="changes/0001-escape.md"||d.Path=="changes/archive/0001-escape.md"))){found[d.Path]=true}};if len(found)!=3{os.Exit(12)}
 if os.Args[1]=="check"{
 C.fw_guard_arm(2,C.CString(target));ds:=p.Check(s);if len(ds)!=3{os.Exit(13)}
 if len(os.Args)>4&&os.Args[4]=="projection"{
 ch:=c.Change{Module:s.Modules[0].Name,Slug:"native-projection",Number:1,Path:"changes/0001-native-projection.md",Set:map[string]any{"preserve":[]any{"native/**"}},Added:[]c.Spec{{ID:"new:native",Body:"It MUST remain pure."}}}
 applied,err:=p.ApplyChange(s.Modules[0],ch);if err!=nil||len(applied.Specs)!=2||applied.Preserve[0]!="native/**"{os.Exit(14)}
 s.Changes=[]c.Change{ch};projected,err:=p.Project(s,"");if err!=nil||len(projected.Modules[0].Specs)!=2{os.Exit(15)}
 hashes,err:=p.BaseHashes(s,ch.Slug);if err!=nil||hashes==nil||len(hashes)!=0{os.Exit(16)}
 ch.Module="missing";if _,err=p.ApplyChange(s.Modules[0],ch);err==nil{os.Exit(17)}
 if _,err=p.Project(s,"missing");err==nil{os.Exit(18)};if _,err=p.BaseHashes(s,"missing");err==nil{os.Exit(19)}
 }
 C.fw_guard_arm(0,nil)
 }
}`, temp)
	put(t, temp, "probe.go", program)
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
	for _, args := range [][]string{{"read-control", fixtureRoot, forbidden}, {"read-control", fixtureRoot, filepath.Join(fixtureRoot, "units/escape/SPEC.md")}, {"read-control", fixtureRoot, filepath.Join(fixtureRoot, "changes/0001-escape.md")}, {"network-control", listener.Addr().String(), forbidden}, {"read-control", fixtureRoot, filepath.Join(fixtureRoot, "changes/archive/0001-escape.md")},
		{"check-control", filepath.Join(fixtureRoot, "fireweed.yaml"), forbidden},
		{"readlink-control", fixtureRoot, filepath.Join(fixtureRoot, "changes/0001-escape.md")},
		{"stat-control", fixtureRoot, forbidden},
		{"write-control", filepath.Join(temp, "mutation-control"), forbidden},
		{"dir-control", fixtureRoot, forbidden},
		{"mkdir-control", filepath.Join(temp, "mutation-directory"), forbidden},
		{"unlink-control", forbidden, forbidden}} {
		cmd = exec.CommandContext(ctx, binary, args...)
		out, err := cmd.CombinedOutput()
		exit, ok := err.(*exec.ExitError)
		if !ok || exit.ExitCode() != 42 {
			t.Fatalf("guard positive control %v: %v %s", args, err, out)
		}
	}
	for _, mode := range []string{"load", "check"} {
		args := []string{mode, fixtureRoot, forbidden}
		if projection {
			args = append(args, "projection")
		}
		cmd = exec.CommandContext(ctx, binary, args...)
		cmd.Dir = fixtureRoot
		cmd.Env = append(os.Environ(), "PATH="+trapDir+string(os.PathListSeparator)+os.Getenv("PATH"))
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("%s performed prohibited I/O or failed: %v %s", mode, err, out)
		}
		if _, err = os.Stat(marker); !os.IsNotExist(err) {
			t.Fatalf("%s executed hook/agent/network command", mode)
		}
		if _, err = os.Stat(filepath.Join(fixtureRoot, "MUST-NOT-RUN")); !os.IsNotExist(err) {
			t.Fatalf("%s executed hook", mode)
		}
		content, err := os.ReadFile(filepath.Join(fixtureRoot, "ignored"))
		if err != nil || string(content) != "KEEP-IGNORED-BYTES" {
			t.Fatalf("%s mutated ignored content", mode)
		}
		after := git(t, fixtureRoot, "status", "--porcelain=v1", "--untracked-files=all") + git(t, fixtureRoot, "show-ref")
		if before != after {
			t.Fatalf("%s mutated repository", mode)
		}
	}
}
