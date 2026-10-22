package evals

import (
	"context"
	"fmt"
	a "github.com/lurkshark/fireweed/packages/assets"
	g "github.com/lurkshark/fireweed/packages/gitstore"
	p "github.com/lurkshark/fireweed/packages/project"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func Test_APP002_APP003_APP017_APP018NativePlaybookReadSafety(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("post-startup native CLI I/O interception executes on macOS")
	}
	root := boundPlaybookRoot(t)
	outside, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	projectPut(t, outside, "private", "PRIVATE-UNCHANGED")
	private := filepath.Join(outside, "private")
	projectPut(t, root, "units/private/SPEC.md", "Unrelated retained module bytes.")
	projectPut(t, root, "changes/0001-private.md", "Unrelated retained change bytes.")
	projectPut(t, root, ".fireweed/playbooks/audit.md", "Unselected retained instructions.")
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
static void pathcheck(const char*p){if(armed==2||(armed&&forbidden&&p&&(strcmp(p,forbidden)==0||strstr(p,"units/private/SPEC.md")||strstr(p,"changes/0001-private.md")||strstr(p,"playbooks/audit.md")||((armed==3||armed==5)&&strstr(p,"playbooks/evals.md"))||((armed==4||armed==5)&&strstr(p,"fireweed.yaml")))))_exit(42);}
static void iocheck(void){if(armed==2)_exit(42);}
static void netcheck(void){if(armed)_exit(42);}
static void mutatecheck(void){if(armed)_exit(42);}
static void opencheck(int f){if(armed&&(f&(O_WRONLY|O_RDWR|O_CREAT|O_TRUNC|O_APPEND)))_exit(42);}
static int fw_open(const char*p,int f,...){pathcheck(p);opencheck(f);mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return open(p,f,m);}
static int fw_openat(int d,const char*p,int f,...){pathcheck(p);opencheck(f);mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return openat(d,p,f,m);}
static ssize_t fw_read(int d,void*b,size_t n){if(armed&&d==STDIN_FILENO)_exit(42);iocheck();return read(d,b,n);}
static ssize_t fw_write(int d,const void*b,size_t n){mutatecheck();return write(d,b,n);}
static int fw_stat(const char*p,struct stat*b){iocheck();return stat(p,b);}
static int fw_lstat(const char*p,struct stat*b){iocheck();return lstat(p,b);}
static int fw_socket(int d,int t,int p){netcheck();return socket(d,t,p);}
static int fw_connect(int s,const struct sockaddr*a,socklen_t n){netcheck();return connect(s,a,n);}
static ssize_t fw_readlink(const char*p,char*b,size_t n){iocheck();return readlink(p,b,n);}
static int fw_access(const char*p,int m){iocheck();return access(p,m);}
static int fw_fstat(int d,struct stat*b){iocheck();return fstat(d,b);}
static int fw_getattrlist(const char*p,struct attrlist*a,void*b,size_t n,unsigned long o){iocheck();return getattrlist(p,a,b,n,o);}
static int fw_mkdir(const char*p,mode_t m){mutatecheck();return mkdir(p,m);}
static int fw_unlink(const char*p){mutatecheck();return unlink(p);}
static int fw_rmdir(const char*p){mutatecheck();return rmdir(p);}
static int fw_rename(const char*p,const char*q){mutatecheck();return rename(p,q);}
static int fw_symlink(const char*p,const char*q){mutatecheck();return symlink(p,q);}
static int fw_chmod(const char*p,mode_t m){mutatecheck();return chmod(p,m);}
static char*fw_getcwd(char*b,size_t n){iocheck();return getcwd(b,n);}
static DIR*fw_opendir(const char*p){iocheck();return opendir(p);}
static struct dirent*fw_readdir(DIR*d){iocheck();return readdir(d);}
#define I(r,o) __attribute__((used)) static struct{const void*replacement;const void*original;} fw_interpose_##o __attribute__((section("__DATA,__interpose")))={(const void*)(uintptr_t)&r,(const void*)(uintptr_t)&o};
I(fw_open,open) I(fw_openat,openat) I(fw_read,read) I(fw_write,write) I(fw_stat,stat) I(fw_lstat,lstat) I(fw_socket,socket) I(fw_connect,connect)
I(fw_readlink,readlink) I(fw_access,access) I(fw_fstat,fstat) I(fw_getattrlist,getattrlist)
I(fw_mkdir,mkdir) I(fw_unlink,unlink) I(fw_rmdir,rmdir) I(fw_rename,rename) I(fw_symlink,symlink) I(fw_chmod,chmod) I(fw_getcwd,getcwd) I(fw_opendir,opendir) I(fw_readdir,readdir)
`
	projectPut(t, temp, "guard.c", guard)
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
import("os";"net";"time";"bytes";"encoding/json";"io";"crypto/sha256";"fmt";"path/filepath";app "github.com/lurkshark/fireweed/packages/app";a "github.com/lurkshark/fireweed/packages/assets")
func main(){
 mode,root,private,address:=os.Args[1],os.Args[2],os.Args[3],os.Args[4];asset,err:=a.Playbook("evals");if err!=nil{os.Exit(80)}
 C.fw_guard_arm(1,C.CString(private))
 switch mode{
 case "private-control":os.ReadFile(private);os.Exit(70)
 case "module-control":os.ReadFile(filepath.Join(root,"units/private/SPEC.md"));os.Exit(71)
 case "queue-control":os.ReadFile(filepath.Join(root,"changes/0001-private.md"));os.Exit(72)
 case "unselected-control":os.ReadFile(filepath.Join(root,".fireweed/playbooks/audit.md"));os.Exit(73)
 case "stdin-control":os.Stdin.Read(make([]byte,1));os.Exit(74)
 case "prompt-control":fmt.Fprint(os.Stdout,"prompt");os.Exit(75)
 case "write-control":os.WriteFile(filepath.Join(root,"transient"),[]byte("bad"),0644);os.Remove(filepath.Join(root,"transient"));os.Exit(76)
 case "network-control":net.DialTimeout("tcp",address,time.Second);os.Exit(77)
 case "selected-control":C.fw_guard_arm(3,C.CString(private));os.ReadFile(filepath.Join(root,".fireweed/playbooks/evals.md"));os.Exit(78)
 case "config-control":C.fw_guard_arm(4,C.CString(private));os.ReadFile(filepath.Join(root,"fireweed.yaml"));os.Exit(79)
 case "standalone-config-control":C.fw_guard_arm(5,C.CString(private));os.ReadFile(filepath.Join(root,"fireweed.yaml"));os.Exit(82)
 case "standalone-selected-control":C.fw_guard_arm(5,C.CString(private));os.ReadFile(filepath.Join(root,".fireweed/playbooks/evals.md"));os.Exit(83)
 }
 args:=[]string{"playbook","evals","--json"};command:="playbook";wantCode:=0;errorCode,errorPath:="",""
 if mode=="unknown"{C.fw_guard_arm(2,C.CString(private));args=[]string{"playbook","../impl","--json"};wantCode=2;errorCode="U01"}
 switch mode {
 case "missing-operand":C.fw_guard_arm(2,C.CString(private));args=[]string{"playbook","--json"};wantCode=2;errorCode="U01"
 case "excess-operands":C.fw_guard_arm(2,C.CString(private));args=[]string{"playbook","evals","extra","--json"};wantCode=2;errorCode="U01"
 case "unknown-flag":C.fw_guard_arm(2,C.CString(private));args=[]string{"playbook","evals","--unknown","--json"};wantCode=2;errorCode="U01"
 case "valued-flag":C.fw_guard_arm(2,C.CString(private));args=[]string{"playbook","evals","--json=true","--json"};wantCode=2;errorCode="U01"
 }

 if mode=="help"{C.fw_guard_arm(2,C.CString(private));args=[]string{"playbook","--help","--json"};command="help"}
 if mode=="override_escape"{C.fw_guard_arm(3,C.CString(private));wantCode=2;errorCode="D05";errorPath=".fireweed/playbooks/evals.md"}
 if mode=="config_escape"{C.fw_guard_arm(4,C.CString(private));wantCode=2;errorCode="C01";errorPath="fireweed.yaml"}
 if mode=="config_error"{wantCode=2;errorCode="C01";errorPath="fireweed.yaml"}
 if mode=="no_config"||mode=="outside_repo"{C.fw_guard_arm(5,C.CString(private))}
 var out,errOut bytes.Buffer;code:=app.Run(args,&out,&errOut);if code!=wantCode||errOut.Len()!=0{os.Exit(10)}
 var result struct{Schema int `+"`json:\"schema_version\"`"+`;Command string;Data json.RawMessage;Diagnostics []struct{Code,Severity,Message,Path string}}
 decoder:=json.NewDecoder(&out);if err:=decoder.Decode(&result);err!=nil||result.Schema!=1||result.Command!=command||result.Diagnostics==nil{os.Exit(11)};var extra any;if decoder.Decode(&extra)!=io.EOF{os.Exit(12)}
 if errorCode!=""{if string(result.Data)!="null"||len(result.Diagnostics)!=1||result.Diagnostics[0].Code!=errorCode||result.Diagnostics[0].Severity!="error"||result.Diagnostics[0].Path!=errorPath{os.Exit(13)}}else if len(result.Diagnostics)!=0{os.Exit(14)}else if command=="playbook"{
  text:=asset.Text;overridden:=false;if mode=="override"{text="---\nbased_on: '"+asset.Hash+"'\n---\nGuarded CLI instructions.\n";overridden=true};if mode=="stale"{text="Guarded stale CLI instructions.\n";overridden=true}
  var data map[string]any;if err:=json.Unmarshal(result.Data,&data);err!=nil||len(data)!=5||data["name"]!="evals"||data["text"]!=text||data["hash"]!=fmt.Sprintf("%%x",sha256.Sum256([]byte(text)))||data["embedded_hash"]!=asset.Hash||data["overridden"]!=overridden{os.Exit(15)}
 }
 C.fw_guard_arm(0,nil)
}`, temp)
	projectPut(t, temp, "probe.go", program)
	sourceRoot, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	binary := filepath.Join(temp, "probe")
	cmd = exec.CommandContext(ctx, "go", "build", "-o", binary, filepath.Join(temp, "probe.go"))
	cmd.Dir = sourceRoot
	cmd.Env = append(os.Environ(), "CGO_ENABLED=1", "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build public CLI probe %v %s", err, out)
	}
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	runAt := func(mode, where, target string) ([]byte, error) {
		cmd := exec.CommandContext(ctx, binary, mode, where, target, listener.Addr().String())
		cmd.Dir = where
		return cmd.CombinedOutput()
	}
	run := func(mode, target string) ([]byte, error) { return runAt(mode, root, target) }
	control := func(t *testing.T, mode, target string) {
		t.Helper()
		out, err := run(mode, target)
		exit, ok := err.(*exec.ExitError)
		if !ok || exit.ExitCode() != 42 {
			t.Fatalf("native ignored-error control %s not intercepted: %v %s", mode, err, out)
		}
	}
	check := func(t *testing.T, mode, target string) {
		t.Helper()
		if out, err := run(mode, target); err != nil || len(out) != 0 {
			t.Fatalf("native CLI safety or output failure %s: %v %s", mode, err, out)
		}
	}
	for _, mode := range []string{"private-control", "module-control", "queue-control", "unselected-control", "stdin-control", "prompt-control", "write-control", "network-control"} {
		t.Run(mode, func(t *testing.T) { control(t, mode, private) })
	}
	for _, mode := range []string{"embedded", "unknown", "help", "missing-operand", "excess-operands", "unknown-flag", "valued-flag"} {
		t.Run(mode, func(t *testing.T) { check(t, mode, private) })
	}
	asset, _ := a.Playbook("evals")
	selected := filepath.Join(root, ".fireweed/playbooks/evals.md")
	for _, mode := range []string{"override", "stale"} {
		text := "Guarded stale CLI instructions.\n"
		if mode == "override" {
			text = "---\nbased_on: '" + asset.Hash + "'\n---\nGuarded CLI instructions.\n"
		}
		projectPut(t, root, ".fireweed/playbooks/evals.md", text)
		fixture, err := p.ResolvePlaybook(root, "evals")
		if err != nil || fixture.Text != text || fixture.Stale != (mode == "stale") {
			t.Fatalf("native current/stale fixture %+v %v", fixture, err)
		}
		t.Run(mode, func(t *testing.T) { check(t, mode, private) })
	}
	if err := os.Remove(selected); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(private, selected); err != nil {
		t.Fatal(err)
	}
	t.Run("selected-control", func(t *testing.T) { control(t, "selected-control", private) })
	t.Run("override_escape", func(t *testing.T) { check(t, "override_escape", private) })
	if err := os.RemoveAll(filepath.Dir(selected)); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, filepath.Dir(selected)); err != nil {
		t.Fatal(err)
	}
	missing := filepath.Join(outside, "evals.md")
	if _, err := os.Stat(missing); !os.IsNotExist(err) {
		t.Fatal("missing endpoint premise")
	}
	t.Run("missing-direct-control", func(t *testing.T) { control(t, "private-control", missing) })
	t.Run("missing-alias-control", func(t *testing.T) { control(t, "selected-control", missing) })
	t.Run("missing-escape", func(t *testing.T) { check(t, "override_escape", missing) })
	fullConfig := filepath.Join(root, "fireweed.yaml")
	if err := os.Remove(fullConfig); err != nil {
		t.Fatal(err)
	}
	t.Run("no_config", func(t *testing.T) { check(t, "no_config", missing) })
	projectPut(t, root, "fireweed.yaml", "invalid config [")
	t.Run("config_error", func(t *testing.T) { check(t, "config_error", private) })
	if err := os.Remove(fullConfig); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(private, fullConfig); err != nil {
		t.Fatal(err)
	}
	t.Run("config-control", func(t *testing.T) { control(t, "config-control", private) })
	t.Run("config_escape", func(t *testing.T) { check(t, "config_escape", private) })
	standalone, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := g.Open(standalone); err == nil {
		t.Fatal("outside-repository native premise")
	}
	projectPut(t, standalone, ".fireweed/playbooks/.keep", "")
	if err := os.Symlink(private, filepath.Join(standalone, "fireweed.yaml")); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(private, filepath.Join(standalone, ".fireweed/playbooks/evals.md")); err != nil {
		t.Fatal(err)
	}
	for _, mode := range []string{"standalone-config-control", "standalone-selected-control"} {
		t.Run(mode, func(t *testing.T) {
			out, err := runAt(mode, standalone, private)
			exit, ok := err.(*exec.ExitError)
			if !ok || exit.ExitCode() != 42 {
				t.Fatalf("standalone content control %s: %v %s", mode, err, out)
			}
		})
	}
	t.Run("outside_repo", func(t *testing.T) {
		if out, err := runAt("outside_repo", standalone, private); err != nil || len(out) != 0 {
			t.Fatalf("outside project read contents or failed embedded fallback: %v %s", err, out)
		}
	})
	if raw, err := os.ReadFile(private); err != nil || string(raw) != "PRIVATE-UNCHANGED" {
		t.Fatal("outside private content changed")
	}
}
