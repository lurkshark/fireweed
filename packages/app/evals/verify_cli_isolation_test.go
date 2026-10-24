package evals

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

func Test_APP027_APP028VerifyBeforeAccess(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("macOS post-startup native I/O control; portable executable cases run separately")
	}
	temp, e := filepath.EvalSymlinks(t.TempDir())
	if e != nil {
		t.Fatal(e)
	}
	// Mode 3 permits harmless cwd/metadata but still forbids content, streams and network.
	// Invalid nonempty selectors retain the stronger all-I/O mode 2 contract.
	guard := strings.Replace(cliContextNativeGuard, "content_at(int d,const char*p){if(armed==2)", "content_at(int d,const char*p){if(armed==2||armed==3)", 1)
	metadataGuard := `static void verify_metadata(const char*p){if(armed!=3||!p)return;const char*b=strrchr(p,'/');b=b?b+1:p;if(!strcmp(b,".git")||!strcmp(b,"fireweed.yaml")||!strcmp(b,"SPEC.md")||!strcmp(b,"SYSTEM.md")||!strcmp(b,"AGENTS.md")||!strcmp(b,".fireweed")||!strncmp(p,".git/",5)||!strncmp(p,".fireweed/",10)||strstr(p,"/.git/")||strstr(p,"/.fireweed/"))_exit(42);}
`
	guard = strings.Replace(guard, "static int fw_stat(", metadataGuard+"static int fw_stat(", 1)
	guard = strings.ReplaceAll(guard, "if(armed==2)_exit(42);return stat(p,b);", "if(armed==2)_exit(42);verify_metadata(p);return stat(p,b);")
	guard = strings.ReplaceAll(guard, "if(armed==2)_exit(42);return lstat(p,b);", "if(armed==2)_exit(42);verify_metadata(p);return lstat(p,b);")
	guard = strings.ReplaceAll(guard, "if(armed==2)_exit(42);return readlink(p,b,n);", "if(armed==2)_exit(42);verify_metadata(p);return readlink(p,b,n);")
	projectPut(t, temp, "guard.c", guard)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	lib := filepath.Join(temp, "libcliverify.dylib")
	cmd := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", lib, filepath.Join(temp, "guard.c"), "-Wl,-install_name,"+lib)
	if out, e := cmd.CombinedOutput(); e != nil {
		t.Fatalf("native Verify guard build %v %s", e, out)
	}
	program := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -lcliverify
#include <stdlib.h>
void fw_cli_context_arm(int,const char*,const char*,const char*);
*/
import "C"
import("bytes";"encoding/json";"os";"fmt";"net";"reflect";"time";"os/exec";app "github.com/lurkshark/fireweed/packages/app")
type record struct{Root,Private,Mode string;Args []string;Guard int}
func main(){raw,e:=os.ReadFile(os.Args[1]);var in record;if e!=nil||json.Unmarshal(raw,&in)!=nil||os.Chdir(in.Root)!=nil{os.Exit(80)};cwd,_:=os.Getwd();env:=os.Environ()
 C.fw_cli_context_arm(C.int(in.Guard),nil,nil,nil)
 switch in.Mode{case "git-control":C.fw_cli_context_arm(0,nil,nil,nil);exec.Command("git","rev-parse","HEAD").Run();os.Exit(65);case "read-control":os.ReadFile(in.Private);os.Exit(60);case "cwd-control":os.Getwd();os.Exit(61);case "cwd-allowed":if _,e:=os.Getwd();e!=nil{os.Exit(83)};if _,e:=os.Stat(".");e!=nil{os.Exit(84)};case "repo-stat-control":os.Stat(".git");os.Exit(66);case "repo-lstat-control":os.Lstat("fireweed.yaml");os.Exit(67);case "repo-link-control":os.Readlink(".git");os.Exit(68);case "stdin-control":os.Stdin.Read(make([]byte,1));os.Exit(62);case "prompt-control":fmt.Print("prompt");os.Exit(63);case "network-control":net.DialTimeout("tcp","127.0.0.1:1",time.Second);os.Exit(64)}
 var out,errOut bytes.Buffer;code:=app.Run(in.Args,&out,&errOut)
 C.fw_cli_context_arm(0,nil,nil,nil);after,e:=os.Getwd();if e!=nil||after!=cwd||!reflect.DeepEqual(env,os.Environ()){os.Exit(81)}
 json.NewEncoder(os.Stdout).Encode(struct{Code int;Out,Err string}{code,out.String(),errOut.String()})
}`, temp)
	projectPut(t, temp, "probe.go", program)
	binary := filepath.Join(temp, "probe")
	root, e := filepath.Abs("../../..")
	if e != nil {
		t.Fatal(e)
	}
	cmd = exec.CommandContext(ctx, "go", "build", "-o", binary, filepath.Join(temp, "probe.go"))
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "CGO_ENABLED=1", "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, e := cmd.CombinedOutput(); e != nil {
		t.Fatalf("native Verify probe build %v %s", e, out)
	}
	where, e := filepath.EvalSymlinks(t.TempDir())
	if e != nil {
		t.Fatal(e)
	}
	projectPut(t, where, "private", "retained private input")
	type record struct {
		Root, Private, Mode string
		Args                []string
		Guard               int
	}
	run := func(t *testing.T, in record, want int) result {
		t.Helper()
		raw, e := json.Marshal(in)
		if e != nil {
			t.Fatal(e)
		}
		projectPut(t, temp, "record.json", string(raw))
		cmd := exec.CommandContext(ctx, binary, filepath.Join(temp, "record.json"))
		cmd.Stdin = strings.NewReader("native stdin sentinel")
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
			t.Fatalf("native %s exit%d want%d: %s", in.Mode, code, want, raw)
		}
		if want != 0 {
			return result{code, string(raw), ""}
		}
		var got struct {
			Code     int
			Out, Err string
		}
		if e = json.Unmarshal(raw, &got); e != nil {
			t.Fatal(e, string(raw))
		}
		return result{got.Code, got.Out, got.Err}
	}
	marker := filepath.Join(temp, "repository-access")
	proxy := filepath.Join(temp, "bin")
	real, e := exec.LookPath("git")
	if e != nil {
		t.Fatal(e)
	}
	projectPut(t, proxy, "git", "#!/bin/sh\ncase \"$*\" in --version|version) exec "+coverageQuote(real)+" \"$@\";; esac\nprintf observed > "+coverageQuote(marker)+"\nexit 94\n")
	if e := os.Chmod(filepath.Join(proxy, "git"), 0755); e != nil {
		t.Fatal(e)
	}
	t.Setenv("PATH", proxy+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Run("repository-observer-control", func(t *testing.T) {
		run(t, record{where, filepath.Join(where, "private"), "git-control", nil, 3}, 65)
		if _, e := os.Stat(marker); e != nil {
			t.Fatal("repository access observer not live", e)
		}
		if e := os.Remove(marker); e != nil {
			t.Fatal(e)
		}
	})
	for _, mode := range []string{"read-control", "cwd-control", "stdin-control", "prompt-control", "network-control"} {
		t.Run(mode, func(t *testing.T) { run(t, record{where, filepath.Join(where, "private"), mode, nil, 2}, 42) })
	}
	for _, mode := range []string{"read-control", "stdin-control", "prompt-control", "network-control", "repo-stat-control", "repo-lstat-control", "repo-link-control"} {
		t.Run("content-mode/"+mode, func(t *testing.T) { run(t, record{where, filepath.Join(where, "private"), mode, nil, 3}, 42) })
	}
	t.Run("cwd-allowed", func(t *testing.T) {
		run(t, record{where, filepath.Join(where, "private"), "cwd-allowed", []string{"version"}, 3}, 0)
	})
	for _, tc := range []struct {
		name       string
		args       []string
		code       int
		diagnostic string
	}{
		{"invalid-slug", []string{"--json", "verify", "../unsafe"}, 1, "C05"},
		{"exact-whitespace", []string{"verify", " leading", "--json"}, 1, "C05"},
		{"bad-arity", []string{"--json", "verify", "a", "b"}, 2, "U01"},
		{"bad-flag", []string{"--json", "verify", "--unknown"}, 2, "U01"},
		{"valued-boolean", []string{"--json", "verify", "--json=true"}, 2, "U01"},
		{"show-flag", []string{"--json", "verify", "--projected"}, 2, "U01"},
		{"coverage-flag", []string{"--json", "verify", "--with-change", "other"}, 2, "U01"},
		{"help", []string{"--json", "verify", "--help"}, 0, ""},
		{"short-help", []string{"--json", "verify", "-h"}, 0, ""},
		{"bad-with-help", []string{"--json", "verify", "--unknown", "--help"}, 2, "U01"},
		{"valued-with-help", []string{"--json", "verify", "--help", "--json=true"}, 2, "U01"},
		{"wrong-command-flag-with-help", []string{"--json", "verify", "--projected", "--help"}, 2, "U01"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			mode := 3
			if tc.name == "invalid-slug" || tc.name == "exact-whitespace" {
				mode = 2
			}
			got := run(t, record{where, filepath.Join(where, "private"), "call", tc.args, mode}, 0)
			if _, e := os.Stat(marker); !os.IsNotExist(e) {
				t.Fatal("help/usage/invalid selector accessed Git repository", e)
			}
			command := "verify"
			if tc.name == "help" || tc.name == "short-help" {
				command = "help"
			}
			x := decode(t, got, command)
			if got.code != tc.code {
				t.Error("before-access result", got)
			}
			if tc.diagnostic != "" && (len(x.Diagnostics) != 1 || x.Diagnostics[0].Code != tc.diagnostic || string(x.Data) != "null") {
				t.Error("before-access diagnostic", got)
			}
		})
	}
}

func Test_APP028VerifyNoninteractiveStdio(t *testing.T) {
	temp := t.TempDir()
	source := `package main
import("bytes";"encoding/json";"fmt";"io";"os";"reflect";app "github.com/lurkshark/fireweed/packages/app")
func main(){if os.Chdir(os.Args[1])!=nil{os.Exit(80)};mode:=os.Args[2];cwd,_:=os.Getwd();env:=os.Environ()
 switch mode{case "stdin-control":os.Stdin.Read(make([]byte,1));case "stdout-control":fmt.Print("prompt");case "stderr-control":fmt.Fprint(os.Stderr,"prompt")}
 var out,errOut bytes.Buffer;code:=app.Run([]string{"verify","current","--json"},&out,&errOut);rest,e:=io.ReadAll(os.Stdin);if e!=nil{os.Exit(81)};after,e:=os.Getwd();if e!=nil||after!=cwd||!reflect.DeepEqual(env,os.Environ()){os.Exit(82)}
 json.NewEncoder(os.Stdout).Encode(struct{Code int;Out,Err,Input string}{code,out.String(),errOut.String(),string(rest)})
}`
	projectPut(t, temp, "probe.go", source)
	binary := filepath.Join(temp, "probe")
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}
	root, e := filepath.Abs("../../..")
	if e != nil {
		t.Fatal(e)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "build", "-o", binary, filepath.Join(temp, "probe.go"))
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, e := cmd.CombinedOutput(); e != nil {
		t.Fatalf("stdio probe build %v %s", e, out)
	}
	input := "verify stdin sentinel\n"
	type message struct {
		Code            int
		Out, Err, Input string
	}
	run := func(t *testing.T, f cliVerifyCase, mode string) ([]byte, string) {
		t.Helper()
		ctx, stop := context.WithTimeout(context.Background(), 5*time.Second)
		defer stop()
		cmd := exec.CommandContext(ctx, binary, f.Where, mode)
		cmd.Stdin = strings.NewReader(input)
		var out, errOut strings.Builder
		cmd.Stdout = &out
		cmd.Stderr = &errOut
		if e := cmd.Run(); e != nil {
			t.Fatal(e, out.String(), errOut.String())
		}
		return []byte(out.String()), errOut.String()
	}
	f := cliVerifyFixture(t, "impl", "sha1", "all", false, false)
	for _, mode := range []string{"stdin-control", "stdout-control", "stderr-control"} {
		t.Run(mode, func(t *testing.T) {
			raw, stderr := run(t, f, mode)
			if mode == "stdout-control" {
				if !strings.HasPrefix(string(raw), "prompt") {
					t.Fatal("ignored stdout control not observed")
				}
				return
			}
			if mode == "stderr-control" {
				if !strings.Contains(stderr, "prompt") {
					t.Fatal("ignored stderr control not observed")
				}
				return
			}
			var got message
			if e := json.Unmarshal(raw, &got); e != nil {
				t.Fatal(e, string(raw))
			}
			if got.Input != input[1:] {
				t.Fatal("ignored stdin consumption control not observed", got.Input)
			}
		})
	}
	for _, mode := range []string{"success", "structured-rejection"} {
		t.Run(mode, func(t *testing.T) {
			f := cliVerifyFixture(t, "impl", "sha1", "all", false, false)
			if mode == "structured-rejection" {
				projectPut(t, f.Root, "units/a/evals/case.txt", "forbidden")
			}
			want, reports := cliVerifyExpected(t, f, "current")
			raw, stderr := run(t, f, "call")
			if stderr != "" {
				t.Fatal("independent prompt/output", stderr)
			}
			var got message
			decoder := json.NewDecoder(bytes.NewReader(raw))
			if e := decoder.Decode(&got); e != nil {
				t.Fatal("extra stdout/prompt", e, string(raw))
			}
			var extra any
			if e := decoder.Decode(&extra); e != io.EOF {
				t.Fatal("extra stdout value", e)
			}
			if got.Input != input {
				t.Error("Verify consumed caller stdin", got.Input)
			}
			r := result{got.Code, got.Out, got.Err}
			x := decode(t, r, "verify")
			if r.code != cliVerifyExit(want) || !reflect.DeepEqual(cliVerifyNormalize(decodeData(t, x.Data), cliVerifyReports(t, f)), cliVerifyWire(t, want, reports)) {
				t.Error("stdio call changed evidence", r)
			}
		})
	}
}
