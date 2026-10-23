package evals

import (
	"context"
	"encoding/json"
	"fmt"
	w "github.com/lurkshark/fireweed/packages/workflow"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

// Reuse the independently controlled canonical-path policy of our fresh public
// source evals. This embedded test-only interposer has no product dependency.
const cliContextNativeGuard = `#include <sys/types.h>
#include <sys/stat.h>
#include <sys/socket.h>
#include <sys/attr.h>
#include <fcntl.h>
#include <unistd.h>
#include <stdarg.h>
#include <stdint.h>
#include <string.h>
#include <dirent.h>
#include <stdio.h>
#include <stdlib.h>
#include <limits.h>
static int armed=0;static const char*root=0;static const char*git=0;static const char*extra=0;
void fw_cli_context_arm(int m,const char*r,const char*g,const char*b){armed=m;root=r;git=g;extra=b;}
static int prefix(const char*p,const char*q,size_t n){return p&&q&&n&&strncmp(p,q,n)==0&&(p[n]==0||p[n]=='/');}
static void path_for(int d,const char*p,char*out){
 char abs[PATH_MAX],base[PATH_MAX],probe[PATH_MAX],tail[PATH_MAX]="",resolved[PATH_MAX];
 if(p[0]=='/'){if(snprintf(abs,sizeof(abs),"%s",p)>=(int)sizeof(abs))_exit(42);}
 else {if(d==AT_FDCWD){if(!getcwd(base,sizeof(base)))_exit(42);}else{if(fcntl(d,F_GETPATH,base)<0)_exit(42);}
  if(snprintf(abs,sizeof(abs),"%s/%s",base,p)>=(int)sizeof(abs))_exit(42);}
 if(!realpath(abs,resolved)){
  snprintf(probe,sizeof(probe),"%s",abs);
  while(!realpath(probe,resolved)){
   char*slash=strrchr(probe,'/');if(!slash)_exit(42);
   char next[PATH_MAX];if(snprintf(next,sizeof(next),"/%s%s",slash+1,tail)>=(int)sizeof(next))_exit(42);
   snprintf(tail,sizeof(tail),"%s",next);if(slash==probe)probe[1]=0;else *slash=0;
  }
  if(snprintf(abs,sizeof(abs),"%s%s",resolved,tail)>=(int)sizeof(abs))_exit(42);
 }else{snprintf(abs,sizeof(abs),"%s",resolved);}
 // Normalize missing suffix components as well as existing aliases.
 char*save=0;size_t n=0;out[n++]='/';out[n]=0;
 for(char*v=strtok_r(abs,"/",&save);v;v=strtok_r(0,"/",&save)){
  if(!strcmp(v,"."))continue;
  if(!strcmp(v,"..")){while(n>1&&out[n-1]!='/')n--;if(n>1)n--;out[n]=0;continue;}
  if(n>1)out[n++]='/';size_t len=strlen(v);if(n+len>=PATH_MAX)_exit(42);memcpy(out+n,v,len+1);n+=len;
 }
}
static void content_at(int d,const char*p){if(armed==2)_exit(42);if(!armed||!p)return;char path[PATH_MAX];path_for(d,p,path);p=path;const char*c=extra;while(c&&*c){const char*e=strchr(c,'\n');size_t n=e?(size_t)(e-c):strlen(c);if(prefix(p,c,n))_exit(42);if(!e)break;c=e+1;}}
static void content(const char*p){content_at(AT_FDCWD,p);}
int fw_cli_context_prepare(const char*p){char dir[PATH_MAX];snprintf(dir,sizeof(dir),"%s",p);char*slash=strrchr(dir,'/');if(!slash)return -1;*slash=0;return open(dir,O_RDONLY|O_DIRECTORY);}

static void mutation(void){if(armed)_exit(42);}
static int fw_open(const char*p,int f,...){content(p);if(f&(O_WRONLY|O_RDWR|O_CREAT|O_TRUNC|O_APPEND))mutation();mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return open(p,f,m);}
static int fw_openat(int d,const char*p,int f,...){content_at(d,p);if(f&(O_WRONLY|O_RDWR|O_CREAT|O_TRUNC|O_APPEND))mutation();mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return openat(d,p,f,m);}
static ssize_t fw_read(int d,void*b,size_t n){if(armed==2||(armed&&d==0))_exit(42);return read(d,b,n);}
static ssize_t fw_write(int d,const void*b,size_t n){mutation();return write(d,b,n);}
static int fw_socket(int d,int t,int p){mutation();return socket(d,t,p);}
static int fw_connect(int s,const struct sockaddr*a,socklen_t n){mutation();return connect(s,a,n);}
static int fw_mkdir(const char*p,mode_t m){mutation();return mkdir(p,m);}
static int fw_unlink(const char*p){mutation();return unlink(p);}
static int fw_rmdir(const char*p){mutation();return rmdir(p);}
static int fw_rename(const char*p,const char*q){mutation();return rename(p,q);}
static int fw_symlink(const char*p,const char*q){mutation();return symlink(p,q);}
static int fw_chmod(const char*p,mode_t m){mutation();return chmod(p,m);}
static DIR*fw_opendir(const char*p){content(p);return opendir(p);}
#define I(r,o) __attribute__((used)) static struct{const void*replacement;const void*original;} cx_##o __attribute__((section("__DATA,__interpose")))={(const void*)(uintptr_t)&r,(const void*)(uintptr_t)&o};
I(fw_open,open) I(fw_openat,openat) I(fw_read,read) I(fw_write,write) I(fw_socket,socket) I(fw_connect,connect) I(fw_mkdir,mkdir) I(fw_unlink,unlink) I(fw_rmdir,rmdir) I(fw_rename,rename) I(fw_symlink,symlink) I(fw_chmod,chmod) I(fw_opendir,opendir)
static int fw_stat(const char*p,struct stat*b){if(armed==2)_exit(42);return stat(p,b);}
static int fw_lstat(const char*p,struct stat*b){if(armed==2)_exit(42);return lstat(p,b);}
static char*fw_getcwd(char*b,size_t n){if(armed==2)_exit(42);return getcwd(b,n);}
static ssize_t fw_readlink(const char*p,char*b,size_t n){if(armed==2)_exit(42);return readlink(p,b,n);}
I(fw_stat,stat) I(fw_lstat,lstat) I(fw_getcwd,getcwd) I(fw_readlink,readlink)

`

func Test_APP019_APP020_APP021ContextNativeIsolation(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("macOS native post-startup CLI I/O guard; portable composition executes separately")
	}
	temp, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	projectPut(t, temp, "guard.c", cliContextNativeGuard)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	lib := filepath.Join(temp, "libclicontext.dylib")
	cmd := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", lib, filepath.Join(temp, "guard.c"), "-Wl,-install_name,"+lib)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("native guard build %v %s", err, out)
	}
	program := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -lclicontext
#include <fcntl.h>
#include <stdlib.h>
#include <unistd.h>
void fw_cli_context_arm(int m,const char*r,const char*g,const char*b);
int fw_cli_context_prepare(const char*p);
int cli_context_openat(int d,const char*p){return openat(d,p,O_RDONLY);}
*/
import "C"
import("bytes";"encoding/json";"os";"fmt";"io";"strings";"path/filepath";"reflect";"net";"time";app "github.com/lurkshark/fireweed/packages/app";w "github.com/lurkshark/fireweed/packages/workflow")
type record struct {Root,Blocked,Error,Address string;Args []string;Code int;Want w.WorkerContext}
func same(a,b any)bool{x,e:=json.Marshal(a);y,f:=json.Marshal(b);return e==nil&&f==nil&&bytes.Equal(x,y)}
func main(){
 raw,err:=os.ReadFile(os.Args[2]);if err!=nil{os.Exit(80)};var in record;if json.Unmarshal(raw,&in)!=nil{os.Exit(81)}
 if os.Chdir(in.Root)!=nil{os.Exit(82)};mode:=os.Args[1]
 private:=strings.Split(in.Blocked,"\n")[0];fd:=C.fw_cli_context_prepare(C.CString(private));if fd<0{os.Exit(83)}
 beforeCwd,err:=os.Getwd();if err!=nil{os.Exit(86)};beforeEnv:=os.Environ()
 arm:=1;if mode=="no-io"{arm=2}
 C.fw_cli_context_arm(C.int(arm),nil,nil,C.CString(in.Blocked))
 switch mode{
 case "absolute-control":os.ReadFile(private);os.Exit(60)
 case "relative-control":r,_:=filepath.Rel(in.Root,private);os.ReadFile("./"+r);os.Exit(61)
 case "fd-control":C.cli_context_openat(fd,C.CString(filepath.Base(private)));os.Exit(62)
 case "missing-control":os.ReadFile(filepath.Join(private,"missing","deeper"));os.Exit(63)
 case "allowed-control":if _,e:=os.ReadFile(filepath.Join(in.Root,"units/a/public/new.txt"));e!=nil{os.Exit(64)};C.fw_cli_context_arm(0,nil,nil,nil);return
 case "stdin-control":os.Stdin.Read(make([]byte,1));os.Exit(65)
 case "prompt-control":fmt.Fprint(os.Stdout,"prompt");os.Exit(66)
 case "mutation-control":os.WriteFile(filepath.Join(in.Root,"transient"),[]byte("bad"),0644);os.Remove(filepath.Join(in.Root,"transient"));os.Exit(67)
 case "network-control":net.DialTimeout("tcp",in.Address,time.Second);os.Exit(68)
 case "public-control":
  s,_,e:=w.LoadApproval(in.Root,"current");if e!=nil{C.fw_cli_context_arm(0,nil,nil,nil);fmt.Fprintln(os.Stderr,e);os.Exit(84)};got,e:=w.BuildContext(s,"current");if e!=nil||!same(got,in.Want){os.Exit(85)};C.fw_cli_context_arm(0,nil,nil,nil);return
 }
 var out,errOut bytes.Buffer;code:=app.Run(in.Args,&out,&errOut)
 C.fw_cli_context_arm(0,nil,nil,nil)
 afterCwd,err:=os.Getwd();if err!=nil||beforeCwd!=afterCwd||!reflect.DeepEqual(beforeEnv,os.Environ()){os.Exit(17)}
 if code!=in.Code||errOut.Len()!=0{fmt.Fprintf(os.Stderr,"code=%%d want=%%d stdout=%%s stderr=%%s",code,in.Code,out.String(),errOut.String());os.Exit(10)}
 var result struct{Schema int `+"`json:\"schema_version\"`"+`;Command string;Data json.RawMessage;Diagnostics []struct{Code,Severity,Message string}}
 d:=json.NewDecoder(&out);if d.Decode(&result)!=nil||result.Schema!=1||result.Diagnostics==nil{os.Exit(11)};var more any;if d.Decode(&more)!=io.EOF{os.Exit(12)}
 if len(in.Args)>1&&in.Args[1]=="--help"{if result.Command!="help"||len(result.Diagnostics)!=0{os.Exit(13)};return}
 if result.Command!="context"{os.Exit(14)}
 if in.Error!=""{if string(result.Data)!="null"||len(result.Diagnostics)!=1||result.Diagnostics[0].Code!=in.Error||result.Diagnostics[0].Severity!="error"||result.Diagnostics[0].Message==""{os.Exit(15)};return}
 var got w.WorkerContext;if json.Unmarshal(result.Data,&got)!=nil||len(result.Diagnostics)!=0||!same(got,in.Want){os.Exit(16)}
}`, temp)
	projectPut(t, temp, "probe.go", program)
	source, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	binary := filepath.Join(temp, "probe")
	cmd = exec.CommandContext(ctx, "go", "build", "-o", binary, filepath.Join(temp, "probe.go"))
	cmd.Dir = source
	cmd.Env = append(os.Environ(), "CGO_ENABLED=1", "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("native public CLI probe build %v %s", err, out)
	}
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	type record struct {
		Root, Blocked, Error, Address string
		Args                          []string
		Code                          int
		Want                          w.WorkerContext
	}
	proxyDir, proxyMarker := "", ""
	run := func(t *testing.T, mode string, in record, exit int) {
		t.Helper()
		path := filepath.Join(temp, "record.json")
		raw, err := json.Marshal(in)
		if err != nil {
			t.Fatal(err)
		}
		projectPut(t, temp, "record.json", string(raw))
		c := exec.CommandContext(ctx, binary, mode, path)
		c.Dir = in.Root
		if proxyDir != "" {
			c.Env = append(os.Environ(), "PATH="+proxyDir+string(os.PathListSeparator)+os.Getenv("PATH"))
		}
		c.Stdin = strings.NewReader("")
		out, err := c.CombinedOutput()
		actual := 0
		if err != nil {
			e, ok := err.(*exec.ExitError)
			if !ok {
				t.Fatal(err)
			}
			actual = e.ExitCode()
		}
		if actual != exit {
			t.Errorf("guard %s exit%d want%d: %s", mode, actual, exit, out)
		}
		if proxyMarker != "" {
			if _, err := os.Stat(proxyMarker); !os.IsNotExist(err) {
				raw, _ := os.ReadFile(proxyMarker)
				t.Fatalf("forbidden Git read/command in %s: %v %s", mode, err, raw)
			}
		}
	}
	root := cliContextFixture(t, "sha1")
	private := filepath.Join(root, "private/secret.txt")
	blocked := []string{private, filepath.Join(root, "units/a/internal/private.txt"), filepath.Join(root, "units/g/api.go"), filepath.Join(root, "units/a/api.go")}
	in := record{Root: root, Blocked: strings.Join(blocked, "\n"), Address: listener.Addr().String(), Args: []string{"context", "current", "--json"}, Want: cliContextExpected(t, root)}
	proxyDir, proxyMarker, loose := cliContextGitProxy(t, root)
	in.Blocked += "\n" + loose
	t.Run("loose-object-control", func(t *testing.T) {
		x := in
		x.Blocked = loose + "\n" + in.Blocked
		run(t, "absolute-control", x, 42)
		run(t, "fd-control", x, 42)
	})
	for _, mode := range []string{"absolute-control", "relative-control", "fd-control", "missing-control", "stdin-control", "prompt-control", "mutation-control", "network-control"} {
		t.Run(mode, func(t *testing.T) { run(t, mode, in, 42) })
	}
	t.Run("allowed-control", func(t *testing.T) { run(t, "allowed-control", in, 0) })
	t.Run("unclaimed-public-control", func(t *testing.T) { run(t, "public-control", in, 0) })
	t.Run("unclaimed", func(t *testing.T) { run(t, "subject", in, 0) })
	for _, test := range []struct {
		name    string
		args    []string
		code    int
		problem string
	}{{"usage", []string{"context", "--json"}, 2, "U01"}, {"excess", []string{"context", "current", "extra", "--json"}, 2, "U01"}, {"flag", []string{"context", "current", "--bad", "--json"}, 2, "U01"}, {"valued", []string{"context", "current", "--json=true", "--json"}, 2, "U01"}, {"projected", []string{"context", "current", "--projected", "--json"}, 2, "U01"}, {"history", []string{"context", "current", "--history", "--json"}, 2, "U01"}, {"coverage", []string{"context", "current", "--with-change=other", "--json"}, 2, "U01"}, {"help", []string{"context", "--help", "--json"}, 0, ""}, {"selector", []string{"context", "../bad", "--json"}, 1, "C05"}} {
		t.Run(test.name, func(t *testing.T) {
			x := in
			x.Args, x.Code, x.Error = test.args, test.code, test.problem
			run(t, "no-io", x, 0)
		})
	}
	cliContextClaim(t, root, "impl", false)
	proxyDir, proxyMarker, loose = cliContextGitProxy(t, root)
	blocked = append(blocked, loose)
	blocked = append(blocked, filepath.Join(root, "fireweed.yaml"), filepath.Join(root, "units/b/SPEC.md"), filepath.Join(root, "changes/0001-before.md"), filepath.Join(root, "changes/0003-later.md"), filepath.Join(root, "changes/archive"))
	in.Blocked = strings.Join(blocked, "\n")
	in.Want = cliContextExpected(t, root)
	t.Run("claimed-public-control", func(t *testing.T) { run(t, "public-control", in, 0) })
	t.Run("claimed", func(t *testing.T) { run(t, "subject", in, 0) })
	// Every newly denied source class is independently proven before the subject.
	for i, path := range blocked[4:] {
		t.Run(fmt.Sprintf("source-control-%d", i), func(t *testing.T) { x := in; x.Blocked = path + "\n" + in.Blocked; run(t, "absolute-control", x, 42) })
	}
	outside, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	projectPut(t, outside, "secret", "Outside read trap")
	selected := filepath.Join(root, "units/a/public/new.txt")
	if err := os.Remove(selected); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(filepath.Join(outside, "secret"), selected); err != nil {
		t.Fatal(err)
	}
	in.Blocked += "\n" + outside
	in.Want = w.WorkerContext{}
	in.Code, in.Error = 2, "C01"
	t.Run("outside-control", func(t *testing.T) {
		x := in
		x.Blocked = filepath.Join(outside, "secret") + "\n" + in.Blocked
		run(t, "absolute-control", x, 42)
	})
	t.Run("escape", func(t *testing.T) { run(t, "subject", in, 0) })
	if err := os.Remove(selected); err != nil {
		t.Fatal(err)
	}
	projectPut(t, root, "units/a/public/new.txt", "Restored selected feedback")
	if err := os.MkdirAll(filepath.Join(root, ".fireweed/playbooks/impl.md"), 0755); err != nil {
		t.Fatal(err)
	}
	in.Code, in.Error = 2, "D05"
	t.Run("override-error-process-state", func(t *testing.T) { run(t, "subject", in, 0) })
	projectPut(t, root, "corrupt-selected-claim", "invalid claim")
	oid := strings.TrimSpace(gitFixtureCommand(t, root, "hash-object", "-w", "corrupt-selected-claim"))
	gitFixtureCommand(t, root, "update-ref", "refs/fireweed/claims/current", oid)
	proxyDir, proxyMarker, _ = cliContextGitProxy(t, root)
	in.Code, in.Error = 1, "P02"
	t.Run("claim-error-process-state", func(t *testing.T) { run(t, "subject", in, 0) })
}
