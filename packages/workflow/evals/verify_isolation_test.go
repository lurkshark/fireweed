package evals

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func Test_FLOW041VerifyInvalidSelectorBeforeNativeAccess(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("macOS post-startup native no-access observer; portable selection tests run elsewhere")
	}
	dir := stateCanonicalDirectory(t, t.TempDir())
	guard := `#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>
#include <fcntl.h>
#include <stdarg.h>
#include <stdint.h>
#include <dirent.h>
static int armed=0;
void fw_verify_arm(int n){armed=n;}
static void denied(void){if(armed)_exit(42);}
static int fw_open(const char*p,int f,...){denied();mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return open(p,f,m);}
static int fw_openat(int d,const char*p,int f,...){denied();mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return openat(d,p,f,m);}
static ssize_t fw_read(int d,void*b,size_t n){denied();return read(d,b,n);}
static ssize_t fw_write(int d,const void*b,size_t n){denied();return write(d,b,n);}
static int fw_stat(const char*p,struct stat*b){denied();return stat(p,b);}
static int fw_lstat(const char*p,struct stat*b){denied();return lstat(p,b);}
static int fw_access(const char*p,int m){denied();return access(p,m);}
static int fw_fstat(int d,struct stat*b){denied();return fstat(d,b);}
static ssize_t fw_readlink(const char*p,char*b,size_t n){denied();return readlink(p,b,n);}
static char*fw_getcwd(char*b,size_t n){denied();return getcwd(b,n);}
static DIR*fw_opendir(const char*p){denied();return opendir(p);}
#define I(r,o) __attribute__((used)) static struct{const void*replacement;const void*original;} fw_##o##_link __attribute__((section("__DATA,__interpose")))={(const void*)(uintptr_t)&r,(const void*)(uintptr_t)&o};
I(fw_open,open) I(fw_openat,openat) I(fw_read,read) I(fw_write,write) I(fw_stat,stat) I(fw_lstat,lstat) I(fw_access,access) I(fw_fstat,fstat) I(fw_readlink,readlink) I(fw_getcwd,getcwd) I(fw_opendir,opendir)
`
	worktreePut(t, dir, "guard.c", guard)
	lib := filepath.Join(dir, "libverifyguard.dylib")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	cmd := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", lib, filepath.Join(dir, "guard.c"), "-Wl,-install_name,"+lib)
	if out, e := cmd.CombinedOutput(); e != nil {
		t.Fatalf("native observer build %v %s", e, out)
	}
	source := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -lverifyguard
void fw_verify_arm(int n);
*/
import "C"
import("context";"fmt";"os";"os/exec";"reflect";"strings";w "github.com/lurkshark/fireweed/packages/workflow")
func main(){mode,root,slug:=os.Args[1],os.Args[2],os.Args[3];C.fw_verify_arm(1);switch mode{case "metadata-control":os.Stat(root);os.Exit(70);case "read-control":os.ReadFile(root);os.Exit(71);case "git-control":exec.Command("git","-C",root,"rev-parse","--show-toplevel").Run();os.Exit(72)};v,e:=w.Verify(context.Background(),root,slug);C.fw_verify_arm(0);if !reflect.DeepEqual(v,w.Verification{})||e==nil||!strings.Contains(e.Error(),"C05"){fmt.Fprintf(os.Stderr,"invalid selector accepted or wrong error: %%+v %%v",v,e);os.Exit(10)}}
`, dir)
	worktreePut(t, dir, "probe.go", source)
	binary := filepath.Join(dir, "probe")
	root, e := filepath.Abs("../../..")
	if e != nil {
		t.Fatal(e)
	}
	cmd = exec.CommandContext(ctx, "go", "build", "-o", binary, filepath.Join(dir, "probe.go"))
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "CGO_ENABLED=1", "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, e := cmd.CombinedOutput(); e != nil {
		t.Fatalf("public no-access probe build %v %s", e, out)
	}
	fixture := claimRepo(t, "sha1")
	sentinel := filepath.Join(fixture, "retained")
	for _, mode := range []string{"metadata-control", "read-control", "git-control"} {
		target := fixture
		if mode == "read-control" {
			target = sentinel
		}
		cmd = exec.CommandContext(ctx, binary, mode, target, "../bad")
		cmd.Dir = fixture
		out, e := cmd.CombinedOutput()
		exit, ok := e.(*exec.ExitError)
		if !ok || exit.ExitCode() != 42 || len(out) != 0 {
			t.Fatalf("ignored-error native control %s not live: %v %s", mode, e, out)
		}
	}
	for _, slug := range []string{"UPPER", "../bad", "a/b", "-a", "a b", "a\n"} {
		t.Run(slug, func(t *testing.T) {
			cmd := exec.CommandContext(ctx, binary, "verify", fixture, slug)
			cmd.Dir = fixture
			if out, e := cmd.CombinedOutput(); e != nil {
				t.Errorf("invalid selector performed native access or rejected incorrectly: %v %s", e, out)
			}
		})
	}
}
