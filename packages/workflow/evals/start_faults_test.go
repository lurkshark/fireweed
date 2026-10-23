package evals

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	a "github.com/lurkshark/fireweed/packages/assets"
	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

const startFaultGuard = `#include <sys/types.h>
#include <sys/socket.h>
#include <sys/stat.h>
#include <spawn.h>
#include <unistd.h>
#include <fcntl.h>
#include <stdarg.h>
#include <stdint.h>
#include <string.h>
#include <stdio.h>
#include <errno.h>
#include <stdlib.h>
#include <signal.h>
static int armed=0,fired=0,deletions=0;
static pid_t owner=0;
static void stopped(int signal_number){_exit(42);}
static void violation(void){if(owner&&getpid()!=owner)kill(owner,SIGUSR2);_exit(42);}
static const char *root=0,*change=0,*context_file=0,*private_file=0;
static char failed_delete[4096],recreated[4096];
void fw_start_arm(int m,const char*r,const char*c,const char*x,const char*p){armed=m;root=r;change=c;context_file=x;private_file=p;if(m){owner=getpid();signal(SIGUSR2,stopped);}}
int fw_start_fired(void){return fired;}
const char*fw_start_recreated(void){return recreated;}
static int exact(const char*a,const char*b){return a&&b&&strcmp(a,b)==0;}
static int within(const char*p,const char*d){return p&&d&&strncmp(p,d,strlen(d))==0&&(p[strlen(d)]==0||p[strlen(d)]=='/');}
static const char*fullpath(int fd,const char*p,char*b){if(!p||p[0]=='/'||fd==AT_FDCWD)return p;char d[4096];if(fcntl(fd,F_GETPATH,d)==0){snprintf(b,4096,"%s/%s",d,p);return b;}return p;}
static void mutation(const char*p){if(armed&&within(p,private_file))_exit(42);}
static int opening(const char*p,int flags){
 if(!armed)return 0;
 if(flags&(O_WRONLY|O_RDWR)){mutation(p);if(armed==2&&exact(p,change)){fired=1;errno=EACCES;return -1;}}
 if(armed==5&&exact(p,context_file)){fired=1;errno=EACCES;return -1;}
 return 0;
}
static int fw_open(const char*p,int f,...){if(opening(p,f)<0)return -1;mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return open(p,f,m);}
static int fw_openat(int d,const char*p,int f,...){char b[4096];const char*q=fullpath(d,p,b);if(opening(q,f)<0)return -1;mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return openat(d,p,f,m);}
static ssize_t fw_read(int d,void*b,size_t n){if(armed&&d==0)_exit(42);return read(d,b,n);}
static ssize_t fw_write(int d,const void*b,size_t n){if(armed&&(d==1||d==2))_exit(42);return write(d,b,n);}
static int fw_socket(int d,int t,int p){if(armed)_exit(42);return socket(d,t,p);}
static int fw_connect(int s,const struct sockaddr*a,socklen_t n){if(armed)_exit(42);return connect(s,a,n);}
static void execution(const char*p){if(armed){const char*b=strrchr(p,'/');b=b?b+1:p;if(strcmp(b,"git")!=0)violation();}}
static int fw_execve(const char*p,char*const a[],char*const e[]){execution(p);return execve(p,a,e);}
static int fw_posix_spawn(pid_t*p,const char*s,const posix_spawn_file_actions_t*f,const posix_spawnattr_t*t,char*const a[],char*const e[]){execution(s);return posix_spawn(p,s,f,t,a,e);}
static int fw_posix_spawnp(pid_t*p,const char*s,const posix_spawn_file_actions_t*f,const posix_spawnattr_t*t,char*const a[],char*const e[]){execution(s);return posix_spawnp(p,s,f,t,a,e);}
static int fw_rename(const char*p,const char*q){mutation(p);mutation(q);if(armed==2&&exact(q,change)){fired=1;errno=EACCES;return -1;}return rename(p,q);}
static int disposable(const char*p){return root&&p&&within(p,root)&&strstr(p,"/units/a/impl/")&&(strstr(p,"/code.txt")||strstr(p,"/second.txt"));}
static int deletion(const char*p){
 mutation(p);
 if((armed==3||armed==4)&&disposable(p)){
  if(failed_delete[0]&&exact(p,failed_delete)){errno=EACCES;return -1;}
  if(deletions){snprintf(failed_delete,sizeof(failed_delete),"%s",p);fired=1;errno=EACCES;return -1;}
 }
 return 0;
}
static void deleted(const char*p,int result){
 if(result==0&&(armed==3||armed==4)&&disposable(p)){
  deletions++;
  if(armed==4){int f=open(p,O_WRONLY|O_CREAT|O_EXCL,0640);if(f<0)_exit(86);const char*v="Concurrent winner bytes.\n";if(write(f,v,strlen(v))!=(ssize_t)strlen(v))_exit(87);fchmod(f,0640);close(f);snprintf(recreated,sizeof(recreated),"%s",p);}
 }
}
static int fw_unlink(const char*p){if(deletion(p)<0)return -1;int r=unlink(p);deleted(p,r);return r;}
static int fw_unlinkat(int d,const char*p,int f){char b[4096];const char*q=fullpath(d,p,b);if(deletion(q)<0)return -1;int r=unlinkat(d,p,f);deleted(q,r);return r;}
static int fw_rmdir(const char*p){mutation(p);return rmdir(p);}
static int fw_mkdir(const char*p,mode_t m){mutation(p);return mkdir(p,m);}
static int fw_chmod(const char*p,mode_t m){mutation(p);return chmod(p,m);}
#define I(r,o) __attribute__((used)) static struct{const void*replacement;const void*original;} fw_interpose_##o __attribute__((section("__DATA,__interpose")))={(const void*)(uintptr_t)&r,(const void*)(uintptr_t)&o};
I(fw_open,open) I(fw_openat,openat) I(fw_read,read) I(fw_write,write) I(fw_socket,socket) I(fw_connect,connect)
I(fw_execve,execve) I(fw_posix_spawn,posix_spawn) I(fw_posix_spawnp,posix_spawnp) I(fw_rename,rename)
I(fw_unlink,unlink) I(fw_unlinkat,unlinkat) I(fw_rmdir,rmdir) I(fw_mkdir,mkdir) I(fw_chmod,chmod)
`

type startFaultInput struct {
	Root, Change, Context, Private, Address, Mode string
	Fault                                         int
}
type startFaultOutput struct {
	Result    w.StartResult
	Error     string
	Fired     bool
	Recreated string
	Control   bool
	StateSame bool
}

func Test_FLOW032_FLOW037_FLOW038_FLOW039_FLOW040StartNativeFailures(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("macOS armed native fault/interaction controls; portable start assertions run separately")
	}
	build, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	worktreePut(t, build, "guard.c", startFaultGuard)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	lib := filepath.Join(build, "libstartfault.dylib")
	cmd := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", lib, filepath.Join(build, "guard.c"), "-Wl,-install_name,"+lib)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("native fault guard compile: %v %s", err, out)
	}
	program := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -lstartfault
void fw_start_arm(int m,const char*r,const char*c,const char*x,const char*p);
int fw_start_fired(void);
const char*fw_start_recreated(void);
*/
import "C"
import("encoding/json";"os";"os/exec";"net";"time";"fmt";"reflect";w "github.com/lurkshark/fireweed/packages/workflow")
type input struct{Root,Change,Context,Private,Address,Mode string;Fault int}
type output struct{Result w.StartResult;Error string;Fired bool;Recreated string;Control,StateSame bool}
func main(){
 raw,e:=os.ReadFile(os.Args[1]);if e!=nil{os.Exit(80)};var in input;if json.Unmarshal(raw,&in)!=nil{os.Exit(81)}
 cwd,e:=os.Getwd();if e!=nil{os.Exit(82)};environment:=os.Environ()
 C.fw_start_arm(C.int(in.Fault),C.CString(in.Root),C.CString(in.Change),C.CString(in.Context),C.CString(in.Private))
 out:=output{}
 switch in.Mode{
 case "stdin-control":os.Stdin.Read(make([]byte,1))
 case "prompt-control":fmt.Fprint(os.Stdout,"prompt")
 case "stderr-control":fmt.Fprint(os.Stderr,"prompt")
 case "network-control":net.DialTimeout("tcp",in.Address,time.Second)
 case "process-control":exec.Command("/usr/bin/true").Run()
 case "hook-control":exec.Command("start-hook-trap").Run()
 case "agent-control":exec.Command("start-agent-trap").Run()
 case "private-write-control":os.WriteFile(in.Private,[]byte("bad"),0644)
 case "private-chmod-control":os.Chmod(in.Private,0777)
 case "private-unlink-control":os.Remove(in.Private)
 case "git-control":_,e=exec.Command("git","--version").CombinedOutput();out.Control=e==nil
 case "write-fault-control":e=os.WriteFile(in.Change,[]byte("bad"),0600);out.Control=e!=nil
 case "context-fault-control":_,e=os.ReadFile(in.Context);out.Control=e!=nil
 case "delete-fault-control":
  a:=os.Remove(in.Root+"/units/a/impl/code.txt");b:=os.Remove(in.Root+"/units/a/impl/second.txt");out.Control=a==nil&&b!=nil
 default:
  out.Result,e=w.Start(in.Root,"current");if e!=nil{out.Error=e.Error()}
 }
 out.Fired=C.fw_start_fired()!=0;out.Recreated=C.GoString(C.fw_start_recreated());C.fw_start_arm(0,nil,nil,nil,nil)
 afterCwd,e:=os.Getwd();out.StateSame=e==nil&&cwd==afterCwd&&reflect.DeepEqual(environment,os.Environ())
 switch in.Mode{case "stdin-control","prompt-control","stderr-control","network-control","process-control","hook-control","agent-control","private-write-control","private-chmod-control","private-unlink-control":os.Exit(70)}
 json.NewEncoder(os.Stdout).Encode(out)
}`, build)
	worktreePut(t, build, "probe.go", program)
	source, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	binary := filepath.Join(build, "probe")
	cmd = exec.CommandContext(ctx, "go", "build", "-o", binary, filepath.Join(build, "probe.go"))
	cmd.Dir = source
	cmd.Env = append(os.Environ(), "CGO_ENABLED=1", "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("public start native probe compile: %v %s", err, out)
	}
	trapDir := filepath.Join(build, "traps")
	if err := os.Mkdir(trapDir, 0755); err != nil {
		t.Fatal(err)
	}
	executableBytes, err := os.ReadFile(binary)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"start-hook-trap", "start-agent-trap", "fireweed-bootstrap"} {
		if err := os.WriteFile(filepath.Join(trapDir, name), executableBytes, 0755); err != nil {
			t.Fatal(err)
		}
	}
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	run := func(t *testing.T, in startFaultInput) (startFaultOutput, int) {
		t.Helper()
		path := filepath.Join(t.TempDir(), "input.json")
		b, err := json.Marshal(in)
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, b, 0600); err != nil {
			t.Fatal(err)
		}
		child := exec.CommandContext(ctx, binary, path)
		child.Dir = in.Root
		child.Env = append(os.Environ(), "PATH="+trapDir+string(os.PathListSeparator)+os.Getenv("PATH"))
		out, err := child.CombinedOutput()
		code := 0
		if err != nil {
			var ok bool
			var exit *exec.ExitError
			if exit, ok = err.(*exec.ExitError); ok {
				code = exit.ExitCode()
			} else {
				t.Fatal("native probe execution", err)
			}
		}
		if code != 0 {
			return startFaultOutput{}, code
		}
		var got startFaultOutput
		if err := json.Unmarshal(out, &got); err != nil {
			t.Fatalf("native probe JSON: %v %q", err, out)
		}
		return got, code
	}
	fixture := func(t *testing.T, kind string) (startFixtureData, startFaultInput) {
		t.Helper()
		f := startFixture(t, "sha1", kind, "")
		worktreePut(t, f.Root, "units/a/impl/second.txt", "Second original disposable bytes.\n")
		phase := "evals"
		if kind == "regenerate" {
			phase = "impl"
		}
		book, err := a.Playbook(phase)
		if err != nil {
			t.Fatal(err)
		}
		target := ".fireweed/playbooks/" + phase + ".md"
		worktreePut(t, f.Root, target, "---\nbased_on: '"+book.Hash+"'\n---\n\nExact selected phase instructions.\n")
		startCommit(t, f.Root, "Fault fixture retained sources")
		selected, err := p.ResolvePlaybook(f.Root, phase)
		if err != nil || !selected.Overridden {
			t.Fatal("public override fixture premise", selected, err)
		}
		return f, startFaultInput{Root: f.Root, Change: filepath.Join(f.Root, f.Path), Context: filepath.Join(f.Root, target), Private: filepath.Join(f.Root, "private/secret.txt"), Address: listener.Addr().String(), Fault: 1}
	}
	for _, mode := range []string{"stdin-control", "prompt-control", "stderr-control", "network-control", "process-control", "hook-control", "agent-control", "private-write-control", "private-chmod-control", "private-unlink-control"} {
		t.Run(mode, func(t *testing.T) {
			_, in := fixture(t, "spec")
			in.Mode = mode
			if _, code := run(t, in); code != 42 {
				t.Errorf("live ignored-error control exit %d want42", code)
			}
		})
	}
	for _, control := range []struct {
		Mode  string
		Fault int
	}{{"git-control", 1}, {"write-fault-control", 2}, {"context-fault-control", 5}, {"delete-fault-control", 3}, {"delete-fault-control", 4}} {
		t.Run(fmt.Sprintf("%s-%d", control.Mode, control.Fault), func(t *testing.T) {
			_, in := fixture(t, "regenerate")
			in.Mode = control.Mode
			in.Fault = control.Fault
			got, code := run(t, in)
			if code != 0 || !got.Control || control.Fault != 1 && !got.Fired {
				t.Errorf("native fault positive control exit%d %+v", code, got)
			}
		})
	}
	for _, test := range []struct {
		Name, Kind string
		Fault      int
	}{{"success", "spec", 1}, {"source-read", "spec", 5}, {"interface-read", "spec", 5}, {"drill-read", "regenerate", 5}, {"selected-write", "spec", 2}, {"context-read", "spec", 5}, {"regenerate-context", "regenerate", 5}, {"second-delete", "regenerate", 3}, {"recreated-delete", "regenerate", 4}} {
		t.Run(test.Name, func(t *testing.T) {
			f, in := fixture(t, test.Kind)
			in.Mode = "subject"
			if test.Name == "context-read" {
				if err := os.Chmod(filepath.Join(f.Root, f.Path), 0751); err != nil {
					t.Fatal(err)
				}
				f.Mode = 0751
				startCommit(t, f.Root, "Executable original change mode for rollback")
			}
			switch test.Name {
			case "source-read":
				in.Context = in.Change
			case "interface-read":
				in.Context = filepath.Join(f.Root, "units/a/api.txt")
			case "drill-read":
				startReport(t, f, "pass", claimBase(t, f.Root))
				in.Context = filepath.Join(f.Root, ".fireweed/drills/alpha.md")
			}
			in.Fault = test.Fault
			before := startTree(t, f.Root)
			git := startGitSnapshot(t, f.Root)
			got, code := run(t, in)
			if code != 0 {
				t.Fatalf("subject attempted forbidden interaction or crashed: exit%d", code)
			}
			if !got.StateSame {
				t.Error("fault/success changed the subject process cwd/environment")
			}
			if test.Fault == 1 || test.Name == "drill-read" {
				if got.Error != "" || got.Result.Claim.Object == "" {
					t.Fatalf("valid guarded Start: %+v", got)
				}
				loaded, err := w.LoadClaim(f.Root, "current")
				if err != nil || loaded != got.Result.Claim {
					t.Error("guarded claim mismatch", loaded, err)
				}
				delete(before, f.Path)
				if test.Name == "drill-read" {
					if !got.Fired {
						t.Error("report read fault not reached")
					}
					warning := false
					for _, d := range got.Result.Diagnostics {
						if d.Code == "D01" && d.Severity == "warning" {
							warning = true
						}
					}
					if !warning {
						t.Error("unreadable passing report did not warn", got.Result.Diagnostics)
					}
					delete(before, "units/a/impl/code.txt")
					delete(before, "units/a/impl/second.txt")
				}
				after := startTree(t, f.Root)
				delete(after, f.Path)
				if !reflect.DeepEqual(before, after) {
					t.Error("guarded success changed unrelated files")
				}
				startUnrelated(t, f.Root, git, "current")
				return
			}
			if !got.Fired {
				t.Error("start did not reach its required public write/deletion/context-read boundary")
			}
			if got.Error == "" || !policyCorrective(got.Error) || !reflect.DeepEqual(got.Result, w.StartResult{}) {
				t.Errorf("fault did not reject atomically: %+v", got)
			}
			after := startTree(t, f.Root)
			if test.Fault == 4 {
				if got.Recreated == "" {
					t.Fatal("no externally recreated deletion premise")
				}
				rel, err := filepath.Rel(f.Root, got.Recreated)
				if err != nil {
					t.Fatal(err)
				}
				rel = filepath.ToSlash(rel)
				node, ok := after[rel]
				if !ok || node.Bytes != "Concurrent winner bytes.\n" || node.Mode.Perm() != 0640 {
					t.Error("overwrote concurrent recreated path", node)
				}
				before[rel] = node
				text := strings.ToLower(got.Error)
				if !strings.Contains(text, "recover") && !strings.Contains(text, "restore") && !strings.Contains(text, "ownership") && !strings.Contains(text, "preserv") {
					t.Error("lost incomplete-recovery detail", got.Error)
				}
			}
			if !reflect.DeepEqual(before, after) {
				t.Error("fault rollback changed exact original files/type/modes")
			}
			if !reflect.DeepEqual(git, startGitSnapshot(t, f.Root)) {
				t.Error("fault rollback changed claim/refs/index/HEAD/parent state")
			}
		})
	}
}
