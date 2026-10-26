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

	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

// Reuse the calibrated observer transport and pre-import Go-entry trap, with
// an additional read-only policy. No product arguments or private code appear
// in the policy. Native Git protocol locality is checked separately.
func heldNativeGuard(t *testing.T) string {
	t.Helper()
	text := startFaultGuard
	replace := func(old, new string) {
		t.Helper()
		if strings.Count(text, old) != 1 {
			t.Fatal("native observer source anchor changed", old)
		}
		text = strings.Replace(text, old, new, 1)
	}
	replace("policy->mode>8", "policy->mode>9")
	replace("static void mutation(const char*p){char path[4096];p=fullpath(AT_FDCWD,p,path);inspection(p);", "static void mutation(const char*p){char path[4096];p=fullpath(AT_FDCWD,p,path);inspection(p);if(armed==9)violation();")
	replace("char path[4096];p=fullpath(AT_FDCWD,p,path);inspection(p);\n if(flags", `char path[4096];p=fullpath(AT_FDCWD,p,path);inspection(p);
 if(armed==9){
  char canonical[4096];if(p&&realpath(p,canonical))p=canonical;
  // os/exec opens the null device when streams are intentionally discarded.
  if(p&&!strcmp(p,"/dev/null")&&!(flags&(O_CREAT|O_TRUNC|O_APPEND)))return 0;
  if(flags&(O_WRONLY|O_RDWR|O_CREAT|O_TRUNC|O_APPEND))violation();
  struct stat info;int directory=p&&stat(p,&info)==0&&S_ISDIR(info.st_mode);
  if(!directory&&within(p,root)){const char*q=p+strlen(root);if(strcmp(q,"/.git")&&strncmp(q,"/.git/",6))violation();}
 }
 if(flags`)
	replace("static ssize_t fw_write(int d,const void*b,size_t n){", "static ssize_t fw_write(int d,const void*b,size_t n){if(armed==9&&d>2){struct stat s;if(fstat(d,&s)==0&&(S_ISREG(s.st_mode)||S_ISDIR(s.st_mode)))violation();}")
	// Native denial details are emitted through the underlying system call so
	// an observer failure identifies its cause without reentering stdio guards.
	text = strings.ReplaceAll(text, "violation()", "held_denied(__LINE__)")
	text = strings.Replace(text, "static int armed=", "#include <sys/syscall.h>\nstatic void held_denied(int);\nstatic int armed=", 1)
	return text + `
static void held_denied(int line){char text[160];int n=snprintf(text,sizeof(text),"held guard denial line=%d pid=%d owner=%d supervisor=%d mode=%d\n",line,getpid(),owner,supervisor,armed);syscall(SYS_write,2,text,n);violation();}
static int held_fchmod(int fd,mode_t m){if(armed==9)violation();return fchmod(fd,m);}
static int held_ftruncate(int fd,off_t n){struct stat s;if(armed==9&&fstat(fd,&s)==0&&(S_ISREG(s.st_mode)||S_ISDIR(s.st_mode)))violation();return ftruncate(fd,n);}
static int held_symlink(const char*a,const char*b){if(armed==9)violation();return symlink(a,b);}
static int held_link(const char*a,const char*b){if(armed==9)violation();return link(a,b);}
// The contract preserves the original caller's cwd/environment. A child may
// establish its own execution context and remove observer transport metadata.
static int held_chdir(const char*p){if(armed==9&&getpid()==owner)violation();return chdir(p);}
static int held_fchdir(int fd){if(armed==9&&getpid()==owner)violation();return fchdir(fd);}
static int held_setenv(const char*k,const char*v,int o){if(armed==9&&getpid()==owner)violation();return setenv(k,v,o);}
static int held_unsetenv(const char*k){if(armed==9&&getpid()==owner)violation();return unsetenv(k);}
I(held_fchmod,fchmod) I(held_ftruncate,ftruncate) I(held_symlink,symlink) I(held_link,link)
I(held_chdir,chdir) I(held_fchdir,fchdir) I(held_setenv,setenv) I(held_unsetenv,unsetenv)
`
}

type heldNativeInput struct {
	Root, Start, Private, Current, Address, Mode string
}
type heldNativeOutput struct {
	Snapshot    p.Snapshot
	Claim       w.Claim
	ClaimObject string
	Present     bool
	Error       string
	StateSame   bool
}

func Test_FLOW048_FLOW049_FLOW051HeldNativeObservation(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("macOS native interaction observer; portable source/state cases run separately")
	}
	build := stateCanonicalDirectory(t, t.TempDir())
	worktreePut(t, build, "guard.c", heldNativeGuard(t))
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	lib := filepath.Join(build, "libheldguard.dylib")
	command := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", lib, filepath.Join(build, "guard.c"), "-Wl,-install_name,"+lib)
	if out, err := command.CombinedOutput(); err != nil {
		t.Fatalf("native held observer compile: %v %s", err, out)
	}
	program := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -lheldguard
void fw_start_arm(int,const char*,const char*,const char*,const char*,const char*);
int fw_start_cleanup(int);
*/
import "C"
import("encoding/json";"fmt";"net";"os";"os/exec";"reflect";"strconv";"syscall";p "github.com/lurkshark/fireweed/packages/project";w "github.com/lurkshark/fireweed/packages/workflow")
type input struct{Root,Start,Private,Current,Address,Mode string}
type output struct{Snapshot p.Snapshot;Claim w.Claim;ClaimObject string;Present bool;Error string;StateSame bool}
func main(){
 if len(os.Args)==3&&os.Args[1]=="--cleanup-held-policy"{pid,e:=strconv.Atoi(os.Args[2]);if e!=nil{os.Exit(80)};os.Exit(int(C.fw_start_cleanup(C.int(pid))))}
 if len(os.Args)!=2{os.Exit(81)};raw,e:=os.ReadFile(os.Args[1]);var in input;if e!=nil||json.Unmarshal(raw,&in)!=nil{os.Exit(82)}
 cwd,e:=os.Getwd();if e!=nil{os.Exit(83)};env:=os.Environ();var result output
 var file *os.File
 if in.Mode=="fd-write"||in.Mode=="fd-truncate"||in.Mode=="fd-chmod"{file,e=os.OpenFile(in.Private,os.O_RDWR,0);if e!=nil{os.Exit(84)};defer file.Close()}
 C.fw_start_arm(9,C.CString(in.Root),C.CString(in.Current),nil,C.CString(in.Private),C.CString(in.Address))
 switch in.Mode {
 case "current-read":_,_=os.ReadFile(in.Current)
 case "private-read":_,_=os.ReadFile(in.Private)
 case "write":_=os.WriteFile(in.Private,[]byte("forbidden"),0600)
 case "create-readonly":f,_:=os.OpenFile(in.Private+"-created",os.O_RDONLY|os.O_CREATE,0600);if f!=nil{f.Close()}
 case "fd-write":_,_=file.Write([]byte("forbidden"))
 case "fd-truncate":_=file.Truncate(0)
 case "fd-chmod":_=file.Chmod(0777)
 case "rename":_=os.Rename(in.Private,in.Private+"-renamed")
 case "unlink":_=os.Remove(in.Private)
 case "mkdir":_=os.Mkdir(in.Private+"-directory",0700)
 case "symlink":_=os.Symlink(in.Private,in.Private+"-link")
 case "link":_=os.Link(in.Private,in.Private+"-link")
 case "chdir":_=os.Chdir(in.Root)
 case "setenv":_=os.Setenv("HELD_GUARD_FORBIDDEN","yes")
 case "stdin":var b [1]byte;_,_=os.Stdin.Read(b[:])
 case "stdout":fmt.Print("forbidden prompt")
 case "stderr":fmt.Fprint(os.Stderr,"forbidden prompt")
 case "network":c,_:=net.Dial("tcp",in.Address);if c!=nil{c.Close()}
 case "socket":fd,_:=syscall.Socket(syscall.AF_INET,syscall.SOCK_STREAM,0);if fd>=0{syscall.Close(fd)}
 case "hook":_=exec.Command("context-hook-trap").Run()
 case "agent":_=exec.Command("context-agent-trap").Run()
 case "bootstrap":_=exec.Command(os.Getenv("FIREWEED_BIN")).Run()
 case "go-self":self,_:=os.Executable();_=exec.Command(self,"--unexpected-go-entry").Run()
 case "native-only":self,_:=os.Executable();e=exec.Command(self,"__fw_start_native_control__").Run();if e!=nil{result.Error=e.Error()}
 case "approval-wire":result.Snapshot,result.Claim,e=w.LoadApproval(in.Start,"current");result.Present=e==nil;if e!=nil{result.Error=e.Error()}
 case "object-only-wire":result.Claim.Object="held-object-only-wire-control";result.Error="restore source for wire control"
 default:result.Snapshot,result.Claim,result.Present,e=w.LoadHeldApproval(in.Start);if e!=nil{result.Error=e.Error()}
 }
 C.fw_start_arm(0,nil,nil,nil,nil,nil);after,err:=os.Getwd();result.StateSame=err==nil&&cwd==after&&reflect.DeepEqual(env,os.Environ());result.ClaimObject=result.Claim.Object;if json.NewEncoder(os.Stdout).Encode(result)!=nil{os.Exit(85)}
}
`, build)
	worktreePut(t, build, "probe.go", program)
	source, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	worktreePut(t, build, "go.mod", fmt.Sprintf("module heldsourceprobe\n\ngo 1.23\nrequire github.com/lurkshark/fireweed v0.0.0\nreplace github.com/lurkshark/fireweed => %q\n", source))
	binary := filepath.Join(build, "probe")
	command = exec.CommandContext(ctx, "go", "build", "-mod=mod", "-o", binary, filepath.Join(build, "probe.go"))
	command.Dir = build
	command.Env = append(os.Environ(), "CGO_ENABLED=1", "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := command.CombinedOutput(); err != nil {
		t.Fatalf("held public probe compile: %v %s", err, out)
	}
	traps := filepath.Join(build, "traps")
	if err := os.Mkdir(traps, 0755); err != nil {
		t.Fatal(err)
	}
	bytes, err := os.ReadFile(binary)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"context-hook-trap", "context-agent-trap", "fireweed-bootstrap"} {
		if err := os.WriteFile(filepath.Join(traps, name), bytes, 0755); err != nil {
			t.Fatal(err)
		}
	}
	environment := []string{}
	for _, item := range os.Environ() {
		if !strings.HasPrefix(item, "__FIREWEED_TEST_START_POLICY_V1=") {
			environment = append(environment, item)
		}
	}
	environment = append(environment, "PATH="+traps+string(os.PathListSeparator)+os.Getenv("PATH"), "FIREWEED_BIN="+filepath.Join(traps, "fireweed-bootstrap"))
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	run := func(t *testing.T, in heldNativeInput) (heldNativeOutput, int) {
		t.Helper()
		path := filepath.Join(t.TempDir(), "input.json")
		encoded, err := json.Marshal(in)
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, encoded, 0600); err != nil {
			t.Fatal(err)
		}
		probeContext, probeCancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer probeCancel()
		child := exec.CommandContext(probeContext, binary, path)
		child.Dir = in.Root
		child.Env = append([]string(nil), environment...)
		out, runErr := child.CombinedOutput()
		if child.Process != nil {
			cleanupContext, cleanupCancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cleanupCancel()
			cleanup := exec.CommandContext(cleanupContext, binary, "--cleanup-held-policy", fmt.Sprint(child.Process.Pid))
			cleanup.Env = append([]string(nil), environment...)
			if detail, err := cleanup.CombinedOutput(); err != nil {
				t.Fatalf("held policy cleanup: %v %s", err, detail)
			}
		}
		if probeContext.Err() != nil {
			t.Fatal("held native probe timed out", probeContext.Err())
		}
		if runErr != nil {
			if exit, ok := runErr.(*exec.ExitError); ok {
				return heldNativeOutput{Error: string(out)}, exit.ExitCode()
			}
			t.Fatal(runErr)
		}
		var got heldNativeOutput
		if err := json.Unmarshal(out, &got); err != nil {
			t.Fatalf("held probe wire: %v %q", err, out)
		}
		got.Claim.Object = got.ClaimObject
		return got, 0
	}
	input := func(t *testing.T, f approvalFixtureData) heldNativeInput {
		t.Helper()
		return heldNativeInput{Root: f.Root, Start: filepath.Join(f.Root, "units/a/e"), Private: filepath.Join(f.Root, "units/a/impl/private.txt"), Current: filepath.Join(f.Root, "fireweed.yaml"), Address: listener.Addr().String()}
	}
	for _, mode := range []string{"approval-wire", "object-only-wire"} {
		t.Run("control/"+mode, func(t *testing.T) {
			f := heldSourceFixture(t, "impl", "sha1", false)
			in := input(t, f)
			in.Mode = mode
			got, code := run(t, in)
			if code != 0 {
				t.Fatal("native claim wire control", mode, code, got.Error)
			}
			if mode == "approval-wire" {
				if got.Error != "" || !got.Present || got.Claim != f.Claim || !reflect.DeepEqual(got.Snapshot, f.Want) {
					t.Fatal("native nonzero exact-token wire control", got.Error, got.Claim)
				}
			} else if got.Claim != (w.Claim{Object: "held-object-only-wire-control"}) || got.ClaimObject == "" {
				t.Fatal("native wire concealed an Object-only partial claim", got)
			}
		})
	}
	for _, mode := range []string{"native-only", "current-read", "private-read", "write", "create-readonly", "fd-write", "fd-truncate", "fd-chmod", "rename", "unlink", "mkdir", "symlink", "link", "chdir", "setenv", "stdin", "stdout", "stderr", "network", "socket", "hook", "agent", "bootstrap", "go-self"} {
		t.Run("control/"+mode, func(t *testing.T) {
			f := heldSourceFixture(t, "impl", "sha1", false)
			in := input(t, f)
			in.Mode = mode
			got, code := run(t, in)
			if mode == "native-only" {
				if code != 0 || got.Error != "" || !got.StateSame {
					t.Fatal("native-only control", code, got)
				}
			} else if code != 42 {
				t.Fatal("ignored-error denial observer not live", mode, code, got)
			}
		})
	}
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			for _, mode := range []string{"held", "none", "foreign", "malformed"} {
				t.Run(format+"/"+fmt.Sprint(linked)+"/"+mode, func(t *testing.T) {
					f := heldSourceFixture(t, "impl", format, true)
					if linked {
						f.Root = heldPeer(t, f.Root)
						f.Claim.Worktree = f.Root
						f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
						f.Want.Root = f.Root
					}
					switch mode {
					case "none":
						claimGit(t, f.Root, "update-ref", "-d", "refs/fireweed/claims/current", f.Claim.Object)
					case "foreign":
						foreign := f.Claim
						foreign.Worktree = heldPeer(t, f.Root)
						claimRaw(t, f.Root, "current", claimPayload(t, foreign))
					case "malformed":
						claimRaw(t, f.Root, "current", "{not a claim")
					}
					heldPoisonCurrent(t, f.Root)
					got, code := run(t, input(t, f))
					if code != 0 || !got.StateSame {
						t.Fatal("API performed forbidden native interaction or changed parent state", code, got)
					}
					switch mode {
					case "held":
						if got.Error != "" || !got.Present || got.Claim != f.Claim || !reflect.DeepEqual(got.Snapshot, f.Want) {
							t.Error("approved native source mismatch", got.Error, got.Present, got.Claim)
						}
					case "none", "foreign":
						if got.Error != "" || got.Present || got.Claim != (w.Claim{}) || !reflect.DeepEqual(got.Snapshot, p.Snapshot{}) {
							t.Error("native absence shape", got.Error, got.Present)
						}
					case "malformed":
						if got.Error == "" || !policyCorrective(got.Error) || got.Present || got.Claim != (w.Claim{}) || !reflect.DeepEqual(got.Snapshot, p.Snapshot{}) {
							t.Error("native rejection shape", got.Error, got.Present)
						}
					}
				})
			}
		}
	}
}
