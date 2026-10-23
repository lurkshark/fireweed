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
)

func Test_APP022_APP023_APP024StartNativeIsolation(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("macOS post-startup native I/O observer; portable executable cases run separately")
	}
	temp, e := filepath.EvalSymlinks(t.TempDir())
	if e != nil {
		t.Fatal(e)
	}
	guard := strings.Replace(cliContextNativeGuard, "static void mutation(void){if(armed)_exit(42);}", "static void mutation(void){if(armed==2)_exit(42);}", 1)
	guard = strings.Replace(guard, "mutation();return write(d,b,n);", "if(armed&&(d==1||d==2))_exit(42);mutation();return write(d,b,n);", 1)
	guard = strings.Replace(guard, "mutation();return socket(d,t,p);", "if(armed)_exit(42);return socket(d,t,p);", 1)
	guard = strings.Replace(guard, "mutation();return connect(s,a,n);", "if(armed)_exit(42);return connect(s,a,n);", 1)
	projectPut(t, temp, "guard.c", guard)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	lib := filepath.Join(temp, "libclistart.dylib")
	cmd := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", lib, filepath.Join(temp, "guard.c"), "-Wl,-install_name,"+lib)
	if out, e := cmd.CombinedOutput(); e != nil {
		t.Fatalf("native Start guard build %v %s", e, out)
	}
	program := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -lclistart
#include <stdlib.h>
void fw_cli_context_arm(int,const char*,const char*,const char*);
*/
import "C"
import("bytes";"encoding/json";"fmt";"net";"os";"reflect";"time";app "github.com/lurkshark/fireweed/packages/app";w "github.com/lurkshark/fireweed/packages/workflow")
type record struct{Root,Blocked,Address,Mode,Expected string;Args []string;NoIO bool;Writer string}
type failed struct{};func(failed)Write([]byte)(int,error){return 0,fmt.Errorf("controlled output failure")}
func main(){raw,e:=os.ReadFile(os.Args[1]);if e!=nil{os.Exit(80)};var in record;if json.Unmarshal(raw,&in)!=nil||os.Chdir(in.Root)!=nil{os.Exit(81)}
 cwd,_:=os.Getwd();env:=os.Environ();arm:=1;if in.NoIO{arm=2};C.fw_cli_context_arm(C.int(arm),nil,nil,C.CString(in.Blocked))
 switch in.Mode {case "read":os.ReadFile(in.Blocked);os.Exit(60);case "cwd":os.Getwd();os.Exit(61);case "stdin":os.Stdin.Read(make([]byte,1));os.Exit(62);case "prompt":fmt.Print("prompt");os.Exit(63);case "network":net.DialTimeout("tcp",in.Address,time.Second);os.Exit(64);case "public":_,e:=w.Start(".","current");C.fw_cli_context_arm(0,nil,nil,nil);if (e==nil&&in.Expected!="")||(e!=nil&&e.Error()!=in.Expected){fmt.Fprintln(os.Stderr,e);os.Exit(82)};return}
 var out,errOut bytes.Buffer;var code int
 if in.Writer=="out"{code=app.Run(in.Args,failed{},&errOut)}else if in.Writer=="err"{code=app.Run(in.Args,&out,failed{})}else{code=app.Run(in.Args,&out,&errOut)}
 C.fw_cli_context_arm(0,nil,nil,nil);after,e:=os.Getwd();if e!=nil||after!=cwd||!reflect.DeepEqual(env,os.Environ()){os.Exit(83)}
 json.NewEncoder(os.Stdout).Encode(struct{Code int;Out,Err string}{code,out.String(),errOut.String()})
}`, temp)
	projectPut(t, temp, "probe.go", program)
	source, e := filepath.Abs("../../..")
	if e != nil {
		t.Fatal(e)
	}
	binary := filepath.Join(temp, "probe")
	cmd = exec.CommandContext(ctx, "go", "build", "-o", binary, filepath.Join(temp, "probe.go"))
	cmd.Dir = source
	cmd.Env = append(os.Environ(), "CGO_ENABLED=1", "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, e := cmd.CombinedOutput(); e != nil {
		t.Fatalf("public Start CLI probe build %v %s", e, out)
	}
	listener, e := net.Listen("tcp", "127.0.0.1:0")
	if e != nil {
		t.Fatal(e)
	}
	defer listener.Close()
	type record struct {
		Root, Blocked, Address, Mode, Expected string
		Args                                   []string
		NoIO                                   bool
		Writer                                 string
	}
	run := func(t *testing.T, in record, exit int) result {
		t.Helper()
		raw, e := json.Marshal(in)
		if e != nil {
			t.Fatal(e)
		}
		projectPut(t, temp, "record.json", string(raw))
		cmd := exec.CommandContext(ctx, binary, filepath.Join(temp, "record.json"))
		cmd.Stdin = strings.NewReader("")
		raw, e = cmd.CombinedOutput()
		actual := 0
		if e != nil {
			x, ok := e.(*exec.ExitError)
			if !ok {
				t.Fatal(e)
			}
			actual = x.ExitCode()
		}
		if actual != exit {
			t.Fatalf("native %s exit%d want%d: %s", in.Mode, actual, exit, raw)
		}
		if exit != 0 || in.Mode == "public" {
			return result{}
		}
		var got struct {
			Code     int
			Out, Err string
		}
		if e = json.Unmarshal(raw, &got); e != nil {
			t.Fatalf("probe record %v %s", e, raw)
		}
		return result{got.Code, got.Out, got.Err}
	}
	root, e := filepath.EvalSymlinks(t.TempDir())
	if e != nil {
		t.Fatal(e)
	}
	projectPut(t, root, "secret", "private")
	in := record{Root: root, Blocked: filepath.Join(root, "secret"), Address: listener.Addr().String(), NoIO: true}
	for _, mode := range []string{"read", "cwd", "stdin", "prompt", "network"} {
		t.Run("no-io-control/"+mode, func(t *testing.T) { x := in; x.Mode = mode; run(t, x, 42) })
	}
	for _, args := range [][]string{{"start"}, {"start", "a", "b"}, {"start", "a", "--bad"}, {"start", "a", "--json=true"}, {"start", "a", "--help=true"}, {"start", "a", "--with-change", "other"}, {"start", "a", "--projected"}, {"start", "a", "--history"}, {"start", "a", "--with-change=other"}, {"start", "--", "--json"}, {"start", "../invalid"}} {
		t.Run("no-io/"+strings.Join(args, "_"), func(t *testing.T) {
			x := in
			x.Args = append(append([]string{}, args...), "--json")
			if len(args) > 1 && args[1] == "--" {
				x.Args = append([]string{"--json"}, args...)
			}
			got := run(t, x, 0)
			v := decode(t, got, "start")
			want := "U01"
			code := 2
			if args[len(args)-1] == "../invalid" || (len(args) > 1 && args[1] == "--") {
				want = "C05"
				code = 1
			}
			if got.code != code || string(v.Data) != "null" || len(v.Diagnostics) != 1 || v.Diagnostics[0].Code != want {
				t.Errorf("no-I/O rejection %+v", got)
			}
		})
	}
	for _, args := range [][]string{{"start", "--help", "--json"}, {"--json", "start", "-h"}} {
		t.Run("no-io/help/"+strings.Join(args, "_"), func(t *testing.T) { x := in; x.Args = args; successJSON(t, run(t, x, 0), "help") })
	}
	for _, mode := range []string{"read", "stdin", "prompt", "network"} {
		t.Run("success-control/"+mode, func(t *testing.T) { x := in; x.NoIO = false; x.Mode = mode; run(t, x, 42) })
	}
	for _, mode := range []string{"public", "subject", "writer"} {
		t.Run("start/"+mode, func(t *testing.T) {
			f := cliStartFixture(t, "sha1", "regenerate", "", true, true, false)
			x := record{Root: f.Where, Blocked: strings.Join([]string{filepath.Join(f.Root, "private"), filepath.Join(source, "packages"), filepath.Join(source, ".git")}, "\n"), Address: in.Address, Mode: mode, Args: []string{"start", "current", "--json"}}
			if mode == "writer" {
				x.Writer = "out"
			}
			before := cliStartSnapshot(t, f)
			got := run(t, x, 0)
			cliStartPreserved(t, f, before, mode == "public" || got.code == 0 || mode == "writer" && got.code == 2)
			if mode == "public" {
				cliStartOwned(t, f)
				return
			}
			if mode == "writer" {
				if got.code != 2 {
					t.Errorf("writer exit%d", got.code)
				}
				cliStartOwned(t, f)
				return
			}
			if got.code != 0 {
				t.Errorf("native subject failed %+v", got)
				return
			}
			v := decode(t, got, "start")
			if !reflectEqualData(t, v.Data, cliStartOwned(t, f)) {
				t.Error("native full context mismatch")
			}
		})
	}
	for _, mode := range []string{"public", "subject"} {
		t.Run("dirty-config/"+mode, func(t *testing.T) {
			f, slug, expected := cliStartError(t, "dirty-config")
			x := record{Root: f.Where, Blocked: filepath.Join(f.Root, "fireweed.yaml"), Address: in.Address, Mode: mode, Expected: expected.Error(), Args: []string{"start", slug, "--json"}}
			before := cliStartSnapshot(t, f)
			got := run(t, x, 0)
			cliStartPreserved(t, f, before, false)
			if mode == "public" {
				return
			}
			diag, code := cliStartErrorCode(expected)
			v := decode(t, got, "start")
			ds := detailedDiagnostics(t, got)
			if got.code != code || string(v.Data) != "null" || len(ds) != 1 || ds[0].Code != diag || ds[0].Message != expected.Error() {
				t.Errorf("dirty current config was independently loaded %+v", got)
			}
		})
	}

}
func reflectEqualData(t *testing.T, raw json.RawMessage, v any) bool {
	return reflect.DeepEqual(decodeData(t, raw), canonical(t, v))
}
