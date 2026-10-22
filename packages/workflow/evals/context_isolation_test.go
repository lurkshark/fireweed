package evals

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

	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

func Test_FLOW016_FLOW017_FLOW018_FLOW019ContextNativeIsolation(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("macOS post-startup native I/O interception; portable context evals run elsewhere")
	}
	temp := stateCanonicalDirectory(t, t.TempDir())
	sourceRoot, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
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
#include <sys/attr.h>
#include <stdio.h>
static int armed=0;static const char*blocked=0;
void fw_context_arm(int m,const char*p){armed=m;blocked=p;}
static void anyio(void){if(armed==2)_exit(42);}
static void mutate(void){if(armed)_exit(42);}
static void content(const char*p){anyio();if(!armed||!blocked||!p)return;const char*c=blocked;while(*c){const char*e=strchr(c,'\n');size_t n=e?(size_t)(e-c):strlen(c);if(n&&strncmp(p,c,n)==0&&(p[n]==0||p[n]=='/'))_exit(42);if(!e)break;c=e+1;}}
static int fw_open(const char*p,int f,...){content(p);if(f&(O_WRONLY|O_RDWR|O_CREAT|O_TRUNC|O_APPEND))mutate();mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return open(p,f,m);}
static int fw_openat(int d,const char*p,int f,...){content(p);if(f&(O_WRONLY|O_RDWR|O_CREAT|O_TRUNC|O_APPEND))mutate();mode_t m=0;if(f&O_CREAT){va_list a;va_start(a,f);m=va_arg(a,int);va_end(a);}return openat(d,p,f,m);}
static ssize_t fw_read(int d,void*b,size_t n){anyio();if(armed&&d==0)_exit(42);return read(d,b,n);}
static ssize_t fw_write(int d,const void*b,size_t n){mutate();return write(d,b,n);}
static int fw_stat(const char*p,struct stat*b){anyio();return stat(p,b);}
static int fw_lstat(const char*p,struct stat*b){anyio();return lstat(p,b);}
static int fw_socket(int d,int t,int p){mutate();return socket(d,t,p);}
static int fw_connect(int s,const struct sockaddr*a,socklen_t n){mutate();return connect(s,a,n);}
static ssize_t fw_readlink(const char*p,char*b,size_t n){anyio();return readlink(p,b,n);}
static int fw_access(const char*p,int m){anyio();return access(p,m);}
static int fw_fstat(int d,struct stat*b){anyio();return fstat(d,b);}
static int fw_getattrlist(const char*p,struct attrlist*a,void*b,size_t n,unsigned long o){anyio();return getattrlist(p,a,b,n,o);}
static int fw_mkdir(const char*p,mode_t m){mutate();return mkdir(p,m);}
static int fw_unlink(const char*p){mutate();return unlink(p);}
static int fw_rmdir(const char*p){mutate();return rmdir(p);}
static int fw_rename(const char*p,const char*q){mutate();return rename(p,q);}
static int fw_symlink(const char*p,const char*q){mutate();return symlink(p,q);}
static int fw_chmod(const char*p,mode_t m){mutate();return chmod(p,m);}
static char*fw_getcwd(char*b,size_t n){anyio();return getcwd(b,n);}
static DIR*fw_opendir(const char*p){anyio();return opendir(p);}
static struct dirent*fw_readdir(DIR*d){anyio();return readdir(d);}
#define I(r,o) __attribute__((used)) static struct{const void*replacement;const void*original;} ci_##o __attribute__((section("__DATA,__interpose")))={(const void*)(uintptr_t)&r,(const void*)(uintptr_t)&o};
I(fw_open,open) I(fw_openat,openat) I(fw_read,read) I(fw_write,write) I(fw_stat,stat) I(fw_lstat,lstat) I(fw_socket,socket) I(fw_connect,connect) I(fw_readlink,readlink) I(fw_access,access) I(fw_fstat,fstat) I(fw_getattrlist,getattrlist) I(fw_mkdir,mkdir) I(fw_unlink,unlink) I(fw_rmdir,rmdir) I(fw_rename,rename) I(fw_symlink,symlink) I(fw_chmod,chmod) I(fw_getcwd,getcwd) I(fw_opendir,opendir) I(fw_readdir,readdir)
`
	if err = os.WriteFile(filepath.Join(temp, "guard.c"), []byte(guard), 0644); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	lib := filepath.Join(temp, "libcontextguard.dylib")
	cmd := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", lib, filepath.Join(temp, "guard.c"), "-Wl,-install_name,"+lib)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("native control build %v %s", err, out)
	}
	program := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -lcontextguard
void fw_context_arm(int m,const char*p);
*/
import "C"
import("encoding/json";"os";"fmt";"net";"time";"reflect";"strings";p "github.com/lurkshark/fireweed/packages/project"
 w "github.com/lurkshark/fireweed/packages/workflow")
type record struct{Snapshot p.Snapshot;Slug,Error,Raw,Spec,System,Playbook,Hash,Common string;Files []w.FileContent}
func main(){var in record;raw,err:=os.ReadFile(os.Args[2]);if err!=nil{os.Exit(80)};if err=json.Unmarshal(raw,&in);err!=nil{os.Exit(81)}
 if os.Args[1]=="fixture-control"{if _,err=p.Project(in.Snapshot,"current");err!=nil{os.Exit(82)};return}
 arm:=1;if os.Args[1]=="no-io"{arm=2};blocked:=C.CString(os.Args[3]);C.fw_context_arm(C.int(arm),blocked)
 switch os.Args[1]{case "metadata-control":C.fw_context_arm(2,blocked);os.Stat(os.Args[2]);os.Exit(75);case "read-control":os.ReadFile(os.Args[4]);os.Exit(70);case "network-control":net.DialTimeout("tcp",os.Args[4],time.Second);os.Exit(71);case "stdin-control":b:=make([]byte,1);os.Stdin.Read(b);os.Exit(72);case "prompt-control":fmt.Print("prompt");os.Exit(73);case "write-control":os.WriteFile(os.Args[4],[]byte("forbidden"),0600);os.Exit(74)}
 got,err:=w.BuildContext(in.Snapshot,in.Slug)
 C.fw_context_arm(0,nil)
 if in.Error!=""{if err==nil||!strings.Contains(err.Error(),in.Error)||!reflect.DeepEqual(got,w.WorkerContext{}){fmt.Fprintf(os.Stderr,"wrong rejection %%v %%+v",err,got);os.Exit(10)};return}
 if err!=nil{fmt.Fprint(os.Stderr,err);os.Exit(11)}
 if got.Change.Slug!="current"||got.ChangeText!=in.Raw||got.SpecText!=in.Spec||got.System!=in.System||got.Playbook!=in.Playbook||got.PlaybookHash!=in.Hash||got.GitCommonDir!=in.Common||got.Files==nil||got.Dependencies==nil||!reflect.DeepEqual(got.Files,in.Files)||len(got.Module.Specs)!=2||got.Module.Specs[1].ID!="ALP-002"{fmt.Fprint(os.Stderr,"incomplete or leaked context");os.Exit(12)}
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
		t.Fatalf("public context probe build %v %s", err, out)
	}
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	record := func(f contextFixtureData, phase, slug, code string) string {
		t.Helper()
		name := phase
		if name == "" {
			name = "author-change"
		}
		if f.Snapshot.Changes[1].Regenerate {
			name = "regenerate"
		}
		book, err := p.ResolvePlaybook(f.Root, name)
		if err != nil && code == "" {
			t.Fatal(err)
		}
		values := map[string]any{"Snapshot": f.Snapshot, "Slug": slug, "Error": code, "Raw": f.Raw, "Spec": f.Spec, "System": f.System, "Files": contextExpectedFiles(f, phase), "Playbook": book.Text, "Hash": book.Hash, "Common": claimCommon(t, f.Root)}
		raw, err := json.Marshal(values)
		if err != nil {
			t.Fatal(err)
		}
		input := filepath.Join(temp, "record.json")
		if err = os.WriteFile(input, raw, 0644); err != nil {
			t.Fatal(err)
		}
		return input
	}
	blocked := func(f contextFixtureData, phase string, extras ...string) string {
		names := []string{"fireweed.yaml", "private/secret.txt", "units/a/impl/private.txt", "units/a/manifest.txt", "units/a/api/old.txt", "units/a/keep/old.txt", "units/a/later/future.txt", "units/a/e/protected.txt", "units/b/api/old.txt", "units/b/e/dependency.txt", "units/b/keep/private.txt", "units/b/impl/private.txt", "units/d/api/transitive.txt", "units/x/api/unrelated.txt", "shared/contract.txt", "units/a/e/ignored-new.txt"}
		if phase == "audit" {
			names = append(names, "units/a/saved/keep.txt", "units/a/append/001.txt", "units/c/api/gamma.txt", "units/b/future/beta.txt")
		}
		selected := phase
		if selected == "" {
			selected = "author-change"
		}
		if f.Snapshot.Changes[1].Regenerate {
			selected = "regenerate"
		}
		for _, name := range []string{"author-change", "audit", "evals", "impl", "regenerate"} {
			if name != selected {
				names = append(names, ".fireweed/playbooks/"+name+".md")
			}
		}
		out := append([]string{}, extras...)
		for _, name := range names {
			out = append(out, name, filepath.Join(f.Root, name))
		}
		return strings.Join(out, "\n")
	}
	run := func(f contextFixtureData, mode, input, deny, target string) ([]byte, error) {
		cmd := exec.CommandContext(ctx, binary, mode, input, deny, target)
		cmd.Dir = f.Root
		return cmd.CombinedOutput()
	}
	control := contextFixture(t, "audit", "sha1")
	input := record(control, "audit", "current", "")
	deny := blocked(control, "audit")
	t.Run("serialized-fixture-control", func(t *testing.T) {
		if out, err := run(control, "fixture-control", input, deny, ""); err != nil {
			t.Fatalf("fixture control %v %s", err, out)
		}
	})
	for _, path := range []string{"fireweed.yaml", "private/secret.txt", "units/a/impl/private.txt", "units/a/e/protected.txt", "units/a/manifest.txt", "units/a/e/ignored-new.txt", "units/a/saved/keep.txt", "units/a/append/001.txt", "units/b/future/beta.txt", "units/b/e/dependency.txt", "units/d/api/transitive.txt", "units/x/api/unrelated.txt", ".fireweed/playbooks/impl.md"} {
		t.Run("ignored-read-control/"+path, func(t *testing.T) {
			out, err := run(control, "read-control", input, deny, filepath.Join(control.Root, path))
			exit, ok := err.(*exec.ExitError)
			if !ok || exit.ExitCode() != 42 || len(out) != 0 {
				t.Fatalf("native ignored-read positive control %v %s", err, out)
			}
		})
	}
	for _, mode := range []string{"network-control", "stdin-control", "prompt-control", "write-control", "metadata-control"} {
		t.Run(mode, func(t *testing.T) {
			target := filepath.Join(temp, "transient")
			if mode == "network-control" {
				target = listener.Addr().String()
			}
			out, err := run(control, mode, input, deny, target)
			exit, ok := err.(*exec.ExitError)
			if !ok || exit.ExitCode() != 42 || len(out) != 0 {
				t.Fatalf("native I/O positive control %v %s", err, out)
			}
		})
	}
	for _, kind := range []string{"queued", "evals", "impl", "audit", "stale", "empty", "regenerate", "claim-recovery", "selected-missing", "worker-metadata", "nested-spec", "claim-existing-empty", "selection", "phase", "projection", "ambiguous", "missing-module", "archive-path", "nested-path", "mismatched-path", "dependency", "system-escape", "spec-escape", "change-escape", "own-escape", "dependency-escape", "playbook-escape", "resolved-missing-escape", "directory-missing-escape"} {
		t.Run(kind, func(t *testing.T) {
			phase := "impl"
			if kind == "queued" {
				phase = ""
			}
			if kind == "evals" || kind == "audit" {
				phase = kind
			}
			f := contextFixture(t, phase, "sha1")
			mode := "check"
			slug := "current"
			code := ""
			extras := []string{}
			switch kind {
			case "ambiguous":
				f.Snapshot.Changes = append(f.Snapshot.Changes, f.Snapshot.Changes[1])
				code = "C05"
				mode = "no-io"
			case "missing-module":
				f.Snapshot.Changes[1].Module = "absent"
				code = "C05"
				mode = "no-io"
			case "archive-path":
				f.Snapshot.Changes[1].Path = "changes/archive/0002-current.md"
				code = "C05"
				mode = "no-io"
			case "nested-path":
				f.Snapshot.Changes[1].Path = "changes/nested/0002-current.md"
				code = "C05"
				mode = "no-io"
			case "mismatched-path":
				f.Snapshot.Changes[1].Path = "changes/0002-other.md"
				code = "C05"
				mode = "no-io"
			case "dependency":
				f.Snapshot.Changes[1].Set["depends_on"] = []string{"absent"}
				code = "C05"
				mode = "no-io"
			case "worker-metadata":
				f.Spec = strings.Replace(f.Spec, "id_prefix: ALP\n", "id_prefix: ALP\nevals: [leaked/**]\n", 1)
				worktreePut(t, f.Root, "units/a/SPEC.md", f.Spec)
				worktreePut(t, f.Root, "units/a/leaked/private.txt", "forbidden by approved zone map")
				extras = append(extras, "units/a/leaked/private.txt", filepath.Join(f.Root, "units/a/leaked/private.txt"))
			case "nested-spec":
				worktreePut(t, f.Root, "units/a/e/SPEC.md", "excluded nested spec bytes")
				extras = append(extras, "units/a/e/SPEC.md", filepath.Join(f.Root, "units/a/e/SPEC.md"))
			case "claim-existing-empty":
				claim := w.Claim{Worktree: f.Root, Change: "current", Run: "context-run", Base: claimBase(t, f.Root), State: f.Raw}
				claimRaw(t, f.Root, "current", claimPayload(t, claim))
				worktreePut(t, f.Root, f.Path, "")
				f.Raw = ""
			case "claim-recovery":
				claim := w.Claim{Worktree: f.Root, Change: "current", Run: "context-run", Base: claimBase(t, f.Root), State: f.Raw}
				claimRaw(t, f.Root, "current", claimPayload(t, claim))
				if err = os.Remove(filepath.Join(f.Root, f.Path)); err != nil {
					t.Fatal(err)
				}
			case "selected-missing":
				if err = os.Remove(filepath.Join(f.Root, "units/a/e/current.txt")); err != nil {
					t.Fatal(err)
				}
				code = "C01"
			case "selection":
				slug = "absent"
				code = "C05"
				mode = "no-io"
			case "phase":
				f.Snapshot.Changes[1].Phase = "applied"
				code = "C05"
				mode = "no-io"
			case "projection":
				f.Snapshot.Changes[0].Added = append(f.Snapshot.Changes[0].Added, f.Snapshot.Modules[1].Specs[0])
				code = "C05"
				mode = "no-io"
			case "stale":
				worktreePut(t, f.Root, ".fireweed/playbooks/impl.md", "---\nbased_on: stale\n---\nexact override")
			case "empty":
				worktreePut(t, f.Root, ".fireweed/playbooks/impl.md", "")
			case "regenerate":
				f.Snapshot.Changes[1].Regenerate = true
				worktreePut(t, f.Root, ".fireweed/playbooks/regenerate.md", "regeneration instructions")
			}

			if kind == "directory-missing-escape" {
				target := "units/a/pub/contract ' λ.txt"
				full := filepath.Join(f.Root, target)
				parent := filepath.Dir(full)
				if err = os.Remove(full); err != nil {
					t.Fatal(err)
				}
				if err = os.Remove(parent); err != nil {
					t.Fatal(err)
				}
				outside := stateCanonicalDirectory(t, t.TempDir())
				if err = os.Symlink(outside, parent); err != nil {
					t.Fatal(err)
				}
				private := filepath.Join(outside, filepath.Base(full))
				extras = append(extras, private, full, target)
				code = "C01"
				input := record(f, phase, slug, code)
				deny := blocked(f, phase, extras...)
				for _, attempt := range []string{private, full} {
					out, err := run(f, "read-control", input, deny, attempt)
					exit, ok := err.(*exec.ExitError)
					if !ok || exit.ExitCode() != 42 || len(out) != 0 {
						t.Fatalf("directory escape canonical missing endpoint control %q %v %s", attempt, err, out)
					}
				}
			}
			targets := map[string]string{"system-escape": "SYSTEM.md", "spec-escape": "units/a/SPEC.md", "change-escape": f.Path, "own-escape": "units/a/pub/contract ' λ.txt", "dependency-escape": "units/c/api/gamma.txt", "playbook-escape": ".fireweed/playbooks/impl.md", "resolved-missing-escape": "units/a/pub/contract ' λ.txt"}
			if target, ok := targets[kind]; ok {
				if kind == "change-escape" {
					claim := w.Claim{Worktree: f.Root, Change: "current", Run: "context-run", Base: claimBase(t, f.Root), State: f.Raw}
					claimRaw(t, f.Root, "current", claimPayload(t, claim))
				}

				outside := stateCanonicalDirectory(t, t.TempDir())
				private := filepath.Join(outside, "private")
				if err = os.WriteFile(private, []byte("private"), 0600); err != nil {
					t.Fatal(err)
				}
				if kind == "resolved-missing-escape" {
					private = filepath.Join(outside, "missing", "file")
				}
				full := filepath.Join(f.Root, target)
				if kind == "playbook-escape" {
					worktreePut(t, f.Root, target, "selected")
				}
				if err = os.Remove(full); err != nil {
					t.Fatal(err)
				}
				if err = os.Symlink(private, full); err != nil {
					t.Fatal(err)
				}
				extras = append(extras, private, full, target)
				code = "C01"
				if kind == "playbook-escape" {
					code = "D05"
				}
				input := record(f, phase, slug, code)
				deny := blocked(f, phase, extras...)
				for _, attempt := range []string{private, full} {
					out, err := run(f, "read-control", input, deny, attempt)
					exit, ok := err.(*exec.ExitError)
					if !ok || exit.ExitCode() != 42 || len(out) != 0 {
						t.Fatalf("canonical/alias ignored-read control %q %v %s", attempt, err, out)
					}
				}
			}
			input := record(f, phase, slug, code)
			deny := blocked(f, phase, extras...)
			if out, err := run(f, mode, input, deny, ""); err != nil || len(out) != 0 {
				t.Fatalf("context attempted forbidden I/O or failed result contract: %v %s", err, out)
			}
		})
	}
}
