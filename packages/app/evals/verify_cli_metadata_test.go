package evals

import (
	"context"
	"encoding/json"
	"fmt"
	w "github.com/lurkshark/fireweed/packages/workflow"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

// This mode permits legitimate owned metadata/index/report writes. It denies
// current project metadata contents in the caller, and keeps native stdio and
// network controls armed. Git children may perform the approved public checks.
func Test_APP025_APP028VerifyCurrentMetadataReads(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("macOS selective post-startup content observer; portable executable cases remain")
	}
	temp, e := filepath.EvalSymlinks(t.TempDir())
	if e != nil {
		t.Fatal(e)
	}
	guard := strings.Replace(cliContextNativeGuard, "static void mutation(void){if(armed)_exit(42);}", "static void mutation(void){if(armed&&armed!=4)_exit(42);}", 1)
	guard = strings.Replace(guard, "static ssize_t fw_write(int d,const void*b,size_t n){mutation();", "static ssize_t fw_write(int d,const void*b,size_t n){if(armed==4&&(d==1||d==2))_exit(42);mutation();", 1)
	guard = strings.Replace(guard, "static int fw_socket(int d,int t,int p){mutation();", "static int fw_socket(int d,int t,int p){if(armed==4)_exit(42);mutation();", 1)
	guard = strings.Replace(guard, "static int fw_connect(int s,const struct sockaddr*a,socklen_t n){mutation();", "static int fw_connect(int s,const struct sockaddr*a,socklen_t n){if(armed==4)_exit(42);mutation();", 1)
	projectPut(t, temp, "guard.c", guard)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	lib := filepath.Join(temp, "libmetadata.dylib")
	cmd := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", lib, filepath.Join(temp, "guard.c"), "-Wl,-install_name,"+lib)
	if out, e := cmd.CombinedOutput(); e != nil {
		t.Fatalf("selective metadata guard: %v %s", e, out)
	}
	source := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -lmetadata
#include <stdlib.h>
void fw_cli_context_arm(int,const char*,const char*,const char*);
*/
import "C"
import("bytes";"context";"encoding/json";"fmt";"net";"os";"reflect";"strings";"time";"unsafe";app "github.com/lurkshark/fireweed/packages/app";p "github.com/lurkshark/fireweed/packages/project";w "github.com/lurkshark/fireweed/packages/workflow")
type input struct{Where,Mode string;Metadata []string}
type output struct{Code int;Out,Err,PublicError string;Public w.Verification}
func main(){raw,e:=os.ReadFile(os.Args[1]);var in input;if e!=nil||json.Unmarshal(raw,&in)!=nil||os.Chdir(in.Where)!=nil{os.Exit(80)};cwd,_:=os.Getwd();env:=os.Environ();blocked:=C.CString(strings.Join(in.Metadata,"\n"));defer C.free(unsafe.Pointer(blocked));C.fw_cli_context_arm(4,nil,nil,blocked)
 switch in.Mode{case "config-control":os.ReadFile(in.Metadata[0]);os.Exit(61);case "spec-control":os.ReadFile(in.Metadata[1]);os.Exit(62);case "load-control":p.Load(".");os.Exit(63);case "stdin-control":os.Stdin.Read(make([]byte,1));os.Exit(64);case "prompt-control":fmt.Print("prompt");os.Exit(65);case "network-control":net.DialTimeout("tcp","127.0.0.1:1",time.Second);os.Exit(66)}
 var result output
 if in.Mode=="public"{result.Public,e=w.Verify(context.Background(),".","current");if e!=nil{result.PublicError=e.Error()}}else{args:=[]string{"verify","current"};if in.Mode=="json"{args=append(args,"--json")};var out,errOut bytes.Buffer;result.Code=app.Run(args,&out,&errOut);result.Out=out.String();result.Err=errOut.String()}
 C.fw_cli_context_arm(0,nil,nil,nil);after,e:=os.Getwd();if e!=nil||after!=cwd||!reflect.DeepEqual(env,os.Environ()){os.Exit(81)};json.NewEncoder(os.Stdout).Encode(result)
}`, temp)
	projectPut(t, temp, "probe.go", source)
	probe := filepath.Join(temp, "probe")
	root, e := filepath.Abs("../../..")
	if e != nil {
		t.Fatal(e)
	}
	cmd = exec.CommandContext(ctx, "go", "build", "-o", probe, filepath.Join(temp, "probe.go"))
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "CGO_ENABLED=1", "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, e := cmd.CombinedOutput(); e != nil {
		t.Fatalf("selective metadata probe: %v %s", e, out)
	}
	type input struct {
		Where, Mode string
		Metadata    []string
	}
	type output struct {
		Code                  int
		Out, Err, PublicError string
		Public                w.Verification
	}
	run := func(t *testing.T, f cliVerifyCase, mode string, want int) output {
		t.Helper()
		in := input{f.Where, mode, []string{filepath.Join(f.Root, "fireweed.yaml"), filepath.Join(f.Root, "units/a/SPEC.md"), filepath.Join(f.Root, "units/b/SPEC.md"), filepath.Join(f.Root, "units/gamma/SPEC.md"), filepath.Join(f.Root, "units/delta/SPEC.md")}}
		raw, e := json.Marshal(in)
		if e != nil {
			t.Fatal(e)
		}
		projectPut(t, temp, "input.json", string(raw))
		call, stop := context.WithTimeout(context.Background(), 15*time.Second)
		defer stop()
		cmd := exec.CommandContext(call, probe, filepath.Join(temp, "input.json"))
		cmd.Stdin = strings.NewReader("metadata stdin sentinel")
		raw, e = cmd.CombinedOutput()
		code := 0
		if e != nil {
			x, ok := e.(*exec.ExitError)
			if !ok {
				t.Fatal(e)
			}
			code = x.ExitCode()
		}
		if code != want {
			t.Fatalf("selective metadata %s exit%d want%d: %s", mode, code, want, raw)
		}
		var out output
		if want == 0 {
			if e := json.Unmarshal(raw, &out); e != nil {
				t.Fatal("raw prompt/extra output", e, string(raw))
			}
		}
		return out
	}
	control := cliVerifyFixture(t, "impl", "sha1", "all", false, false)
	for _, mode := range []string{"config-control", "spec-control", "load-control", "stdin-control", "prompt-control", "network-control"} {
		t.Run(mode, func(t *testing.T) { run(t, control, mode, 42) })
	}
	for _, scenario := range []string{"evals", "audit", "impl", "current-config", "forbidden-zone"} {
		t.Run(scenario, func(t *testing.T) {
			phase := scenario
			if phase == "current-config" || phase == "forbidden-zone" {
				phase = "impl"
			}
			f := cliVerifyFixture(t, phase, "sha1", "all", false, false)
			if scenario == "current-config" {
				projectPut(t, f.Root, "fireweed.yaml", "invalid: [current metadata")
			}
			if scenario == "forbidden-zone" {
				projectPut(t, f.Root, "units/a/evals/case.txt", "forbidden current edit")
			}
			want, reports := cliVerifyExpected(t, f, "current")
			names := []string{}
			if len(want.Coverage) > 0 {
				names = cliVerifyHookNames(phase, "all")
			}
			// Public Verify must succeed under this observer before CLI evidence counts.
			t.Run("public", func(t *testing.T) {
				cliVerifyReadonly(t, f, func() {
					out := run(t, f, "public", 0)
					actual := cliVerifyReports(t, f)
					if out.PublicError != "" || !reflect.DeepEqual(cliVerifyWire(t, out.Public, actual), cliVerifyWire(t, want, reports)) {
						t.Fatal("legitimate public observer calibration", out.PublicError, out.Public, want)
					}
					cliVerifyObserve(t, f, names, probe)
				})
			})
			cliVerifyClear(t, f)
			// Rejected pre-hook cases have no coverage; do not invent expected hooks.
			for _, mode := range []string{"json", "plain"} {
				t.Run(mode, func(t *testing.T) {
					cliVerifyClear(t, f)
					cliVerifyReadonly(t, f, func() {
						out := run(t, f, mode, 0)
						got := result{out.Code, out.Out, out.Err}
						actual := cliVerifyReports(t, f)
						if mode == "json" {
							x := decode(t, got, "verify")
							var envelope map[string]any
							if e := json.Unmarshal([]byte(got.out), &envelope); e != nil {
								t.Fatal(e)
							}
							if !reflect.DeepEqual(cliVerifyNormalize(envelope["diagnostics"], actual), cliVerifyNormalize(canonical(t, want.Diagnostics), reports)) {
								t.Error("selective-read envelope diagnostics", got)
							}
							if got.code != cliVerifyExit(want) || !reflect.DeepEqual(cliVerifyNormalize(decodeData(t, x.Data), actual), cliVerifyWire(t, want, reports)) {
								t.Error("selective-read CLI JSON", got)
							}
						} else {
							if len(want.Diagnostics) == 0 && got.err != "" {
								t.Error("selective-read plain extra stderr", got)
							}
							if got.code != cliVerifyExit(want) || !reflect.DeepEqual(cliVerifyNormalize(decodeData(t, json.RawMessage(got.out)), actual), cliVerifyWire(t, want, reports)) {
								t.Error("selective-read CLI plain", got)
							}
							for _, d := range want.Diagnostics {
								if !strings.Contains(got.err, d.Code) || d.Path != "" && !strings.Contains(got.err, d.Path) {
									t.Error("selective-read plain diagnostic", got, d)
								}
							}
						}
						cliVerifyObserve(t, f, names, probe)
					})
				})
			}
		})
	}
}
