package evals_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func Test_PRJ024_PRJ025_PRJ026_PRJ027BaselineNativeIsolation(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("native ignored-read/mutation/network interception runs on macOS")
	}
	temp := baselineCanonical(t, t.TempDir())
	sourceRoot, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	sourceRoot = baselineCanonical(t, sourceRoot)
	code := `#include <sys/types.h>
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
static int armed=0;static const char*root=0;static const char*git=0;static const char*extra=0;
void fw_baseline_arm(int m,const char*r,const char*g,const char*b){armed=m;root=r;git=g;extra=b;}
static int prefix(const char*p,const char*q,size_t n){return p&&q&&n&&strncmp(p,q,n)==0&&(p[n]==0||p[n]=='/');}
static void content(const char*p){if(!armed||!p)return;if(root&&prefix(p,root,strlen(root))&&!(git&&prefix(p,git,strlen(git))))_exit(42);const char*c=extra;while(c&&*c){const char*e=strchr(c,'\n');size_t n=e?(size_t)(e-c):strlen(c);if(prefix(p,c,n))_exit(42);if(!e)break;c=e+1;}}
static void mutation(void){if(armed)_exit(42);}
static int fw_open(const char*p,int f,...){content(p);if(f&(O_WRONLY|O_RDWR|O_CREAT|O_TRUNC|O_APPEND))mutation();mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return open(p,f,m);}
static int fw_openat(int d,const char*p,int f,...){content(p);if(f&(O_WRONLY|O_RDWR|O_CREAT|O_TRUNC|O_APPEND))mutation();mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return openat(d,p,f,m);}
static ssize_t fw_read(int d,void*b,size_t n){if(armed&&d==0)_exit(42);return read(d,b,n);}
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
#define I(r,o) __attribute__((used)) static struct{const void*replacement;const void*original;} bi_##o __attribute__((section("__DATA,__interpose")))={(const void*)(uintptr_t)&r,(const void*)(uintptr_t)&o};
I(fw_open,open) I(fw_openat,openat) I(fw_read,read) I(fw_write,write) I(fw_socket,socket) I(fw_connect,connect) I(fw_mkdir,mkdir) I(fw_unlink,unlink) I(fw_rmdir,rmdir) I(fw_rename,rename) I(fw_symlink,symlink) I(fw_chmod,chmod) I(fw_opendir,opendir)
`
	if err = os.WriteFile(filepath.Join(temp, "guard.c"), []byte(code), 0644); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Minute)
	defer cancel()
	lib := filepath.Join(temp, "libbaselineguard.dylib")
	cmd := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", lib, filepath.Join(temp, "guard.c"), "-Wl,-install_name,"+lib)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("native baseline guard build %v %s", err, out)
	}
	program := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -lbaselineguard
void fw_baseline_arm(int m,const char*r,const char*g,const char*b);
*/
import "C"
import("encoding/json";"os";"net";"time";"reflect";"strings";"fmt";p "github.com/lurkshark/fireweed/packages/project")
type record struct{Root,Git,Revision,Blocked string;Expected p.Snapshot;Error bool}
func main(){var in record;raw,err:=os.ReadFile(os.Args[2]);if err!=nil{os.Exit(80)};if json.Unmarshal(raw,&in)!=nil{os.Exit(81)}
 if os.Args[1]=="fixture-control"{if in.Expected.Root!=in.Root||in.Expected.Modules==nil||in.Expected.Diagnostics==nil{os.Exit(82)};return}
 root:=C.CString(in.Root);git:=C.CString(in.Git);blocked:=C.CString(in.Blocked);C.fw_baseline_arm(1,root,git,blocked)
 switch os.Args[1]{case "read-control":os.ReadFile(os.Args[3]);os.Exit(70);case "network-control":net.DialTimeout("tcp",os.Args[3],time.Second);os.Exit(71);case "stdin-control":b:=make([]byte,1);os.Stdin.Read(b);os.Exit(72);case "prompt-control":fmt.Print("prompt");os.Exit(73);case "write-control":os.WriteFile(os.Args[3],[]byte("transient"),0600);os.Exit(74)}
 got,err:=p.LoadAt(in.Root,in.Revision)
 C.fw_baseline_arm(0,nil,nil,nil)
 if in.Error{if err==nil||!strings.Contains(err.Error(),"C01")||!reflect.DeepEqual(got,p.Snapshot{}){fmt.Fprintf(os.Stderr,"wrong fatal result %%v %%+v",err,got);os.Exit(10)};return}
 if err!=nil{fmt.Fprint(os.Stderr,err);os.Exit(11)}
 if got.Modules==nil||got.Changes==nil||got.Archive==nil||got.Diagnostics==nil{os.Exit(12)}
 for i:=range got.Diagnostics{if got.Diagnostics[i].Message==""{os.Exit(13)};got.Diagnostics[i].Message=""};for i:=range in.Expected.Diagnostics{in.Expected.Diagnostics[i].Message=""}
 if !reflect.DeepEqual(got,in.Expected){fmt.Fprint(os.Stderr,"selected metadata differs");os.Exit(14)}
}`, temp)
	source := filepath.Join(temp, "probe.go")
	if err = os.WriteFile(source, []byte(program), 0644); err != nil {
		t.Fatal(err)
	}
	binary := filepath.Join(temp, "probe")
	cmd = exec.CommandContext(ctx, "go", "build", "-o", binary, source)
	cmd.Dir = sourceRoot
	cmd.Env = append(os.Environ(), "CGO_ENABLED=1", "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("public baseline probe build %v %s", err, out)
	}
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	for _, kind := range []string{"normal", "dirty", "deleted", "outside-current", "config-link", "invalid-config", "diagnosed", "unknown-revision", "escape", "missing-outside-escape"} {
		t.Run(kind, func(t *testing.T) {
			f := baselineFixture(t, "sha1")
			privateRoot := baselineCanonical(t, t.TempDir())
			private := filepath.Join(privateRoot, "private")
			put(t, privateRoot, "private", "outside sentinel")
			blocked := []string{private, sourceRoot}
			extra := []string{}
			revision := f.Commit
			fatal := false
			switch kind {
			case "dirty":
				put(t, f.Root, "fireweed.yaml", "current malformed: [")
				put(t, f.Root, "units/a/SPEC.md", "current invalid module")
			case "deleted":
				for _, name := range []string{"fireweed.yaml", "units/a/SPEC.md", "changes/0002-alpha.md"} {
					if err = os.Remove(filepath.Join(f.Root, name)); err != nil {
						t.Fatal(err)
					}
				}
			case "outside-current":
				for _, name := range []string{"fireweed.yaml", "units/a/SPEC.md", "changes/0002-alpha.md"} {
					full := filepath.Join(f.Root, name)
					if err = os.Remove(full); err != nil {
						t.Fatal(err)
					}
					if err = os.Symlink(private, full); err != nil {
						t.Fatal(err)
					}
				}
			case "config-link":
				put(t, f.Root, "metadata/config.yaml", f.Files["fireweed.yaml"])
				if err = os.Remove(filepath.Join(f.Root, "fireweed.yaml")); err != nil {
					t.Fatal(err)
				}
				if err = os.Symlink("metadata/config.yaml", filepath.Join(f.Root, "fireweed.yaml")); err != nil {
					t.Fatal(err)
				}
				baselineFresh(t, &f)
				revision = f.Commit
				extra = append(extra, "metadata/config.yaml")
			case "invalid-config":
				put(t, f.Root, "fireweed.yaml", "selected malformed: [")
				f.Commit = baselineCommit(t, f.Root)
				revision = f.Commit
				put(t, f.Root, "fireweed.yaml", f.Files["fireweed.yaml"])
				fatal = true
			case "diagnosed":
				put(t, f.Root, "units/bad/SPEC.md", module("invalid name", "BAD", ""))
				baselineFresh(t, &f)
				revision = f.Commit
			case "unknown-revision":
				revision = "missing"
				fatal = true
			case "escape", "missing-outside-escape":
				target := private
				if kind == "missing-outside-escape" {
					target = filepath.Join(privateRoot, "missing", "configuration.yaml")
					blocked = append(blocked, target)
				}
				full := filepath.Join(f.Root, "fireweed.yaml")
				if err = os.Remove(full); err != nil {
					t.Fatal(err)
				}
				if err = os.Symlink(target, full); err != nil {
					t.Fatal(err)
				}
				f.Commit = baselineCommit(t, f.Root)
				revision = f.Commit
				if err = os.Remove(full); err != nil {
					t.Fatal(err)
				}
				put(t, f.Root, "fireweed.yaml", f.Files["fireweed.yaml"])
				fatal = true
			}
			common := baselineGit(t, f.Root, "rev-parse", "--git-common-dir")
			if !filepath.IsAbs(common) {
				common = filepath.Join(f.Root, common)
			}
			common = baselineCanonical(t, common)
			privateObject := baselineGit(t, f.Root, "rev-parse", f.Commit+":private/secret.txt")
			loose := filepath.Join(common, "objects", privateObject[:2], privateObject[2:])
			if _, err = os.Stat(loose); err != nil {
				t.Fatal("loose object fixture", err)
			}
			blocked = append(blocked, loose)
			policyFixture := f
			if fatal && kind != "escape" && kind != "missing-outside-escape" && kind != "unknown-revision" {
				policyFixture.Expected.Modules = nil
				policyFixture.Expected.Changes = nil
				policyFixture.Expected.Archive = nil
			}
			bodyMarker, _ := baselineGuard(t, policyFixture, "", extra...)
			commandMarker := playbookTraps(t)
			values := map[string]any{"Root": f.Root, "Git": filepath.Join(f.Root, ".git"), "Revision": revision, "Blocked": strings.Join(blocked, "\n"), "Expected": f.Expected, "Error": fatal}
			raw, err := json.Marshal(values)
			if err != nil {
				t.Fatal(err)
			}
			input := filepath.Join(temp, "record.json")
			if err = os.WriteFile(input, raw, 0644); err != nil {
				t.Fatal(err)
			}
			run := func(mode, target string) ([]byte, error) {
				cmd := exec.CommandContext(ctx, binary, mode, input, target)
				cmd.Dir = f.Root
				return cmd.CombinedOutput()
			}
			if out, err := run("fixture-control", ""); err != nil {
				t.Fatalf("serialized fixture control %v %s", err, out)
			}
			for _, target := range []string{filepath.Join(f.Root, "fireweed.yaml"), filepath.Join(f.Root, "units/a/SPEC.md"), filepath.Join(f.Root, "changes/0002-alpha.md"), private, loose, filepath.Join(sourceRoot, "packages/project/api.go")} {
				out, err := run("read-control", target)
				exit, ok := err.(*exec.ExitError)
				if !ok || exit.ExitCode() != 42 || len(out) != 0 {
					t.Fatalf("ignored-error read control %q: %v %s", target, err, out)
				}
			}
			if kind == "missing-outside-escape" {
				target := filepath.Join(privateRoot, "missing", "configuration.yaml")
				out, err := run("read-control", target)
				exit, ok := err.(*exec.ExitError)
				if !ok || exit.ExitCode() != 42 || len(out) != 0 {
					t.Fatalf("canonical ignored-ENOENT control %v %s", err, out)
				}
			}
			for _, mode := range []string{"network-control", "stdin-control", "prompt-control", "write-control"} {
				target := filepath.Join(privateRoot, "transient")
				if mode == "network-control" {
					target = listener.Addr().String()
				}
				out, err := run(mode, target)
				exit, ok := err.(*exec.ExitError)
				if !ok || exit.ExitCode() != 42 || len(out) != 0 {
					t.Fatalf("I/O control %s: %v %s", mode, err, out)
				}
			}
			if out, err := run("load", ""); err != nil || len(out) != 0 {
				t.Fatalf("baseline read attempted source/private I/O or failed contract: %v %s", err, out)
			}
			baselineNoMarker(t, bodyMarker)
			baselineNoMarker(t, commandMarker)
			if b, err := os.ReadFile(private); err != nil || string(b) != "outside sentinel" {
				t.Fatal("outside data changed")
			}
		})
	}
}
