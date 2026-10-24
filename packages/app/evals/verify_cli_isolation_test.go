package evals

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	w "github.com/lurkshark/fireweed/packages/workflow"
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
		{"literal-json", []string{"--json", "verify", "--", "--json"}, 1, "C05"},
		{"bad-arity", []string{"--json", "verify", "a", "b"}, 2, "U01"},
		{"bad-flag", []string{"--json", "verify", "--unknown"}, 2, "U01"},
		{"valued-boolean", []string{"--json", "verify", "--json=true"}, 2, "U01"},
		{"show-flag", []string{"--json", "verify", "--projected"}, 2, "U01"},
		{"coverage-flag", []string{"--json", "verify", "--with-change", "other"}, 2, "U01"},
		{"help", []string{"--json", "verify", "--help"}, 0, ""},
		{"short-help", []string{"--json", "verify", "-h"}, 0, ""},
		{"help-excess-long", []string{"--json", "verify", "current", "extra", "--help"}, 0, ""},
		{"help-excess-short", []string{"--json", "verify", "current", "extra", "-h"}, 0, ""},
		{"bad-with-help", []string{"--json", "verify", "--unknown", "--help"}, 2, "U01"},
		{"valued-with-help", []string{"--json", "verify", "--help", "--json=true"}, 2, "U01"},
		{"wrong-command-flag-with-help", []string{"--json", "verify", "--projected", "--help"}, 2, "U01"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			mode := 3
			if tc.name == "invalid-slug" || tc.name == "exact-whitespace" || tc.name == "literal-json" {
				mode = 2
			}
			got := run(t, record{where, filepath.Join(where, "private"), "call", tc.args, mode}, 0)
			if _, e := os.Stat(marker); !os.IsNotExist(e) {
				t.Fatal("help/usage/invalid selector accessed Git repository", e)
			}
			command := "verify"
			if tc.code == 0 {
				command = "help"
			}
			x := decode(t, got, command)
			if tc.code == 0 {
				var help string
				if e := json.Unmarshal(x.Data, &help); e != nil || !strings.Contains(help, "verify [<change>]") {
					t.Error("native help content", e, got)
				}
			}
			if got.code != tc.code {
				t.Error("before-access result", got)
			}
			if tc.diagnostic != "" && (len(x.Diagnostics) != 1 || x.Diagnostics[0].Code != tc.diagnostic || string(x.Data) != "null") {
				t.Error("before-access diagnostic", got)
			}
		})
		t.Run(tc.name+"-plain", func(t *testing.T) {
			args := []string{}
			literal := false
			for _, arg := range tc.args {
				if arg == "--" {
					literal = true
				}
				if arg != "--json" || literal {
					args = append(args, arg)
				}
			}
			mode := 3
			if tc.name == "invalid-slug" || tc.name == "exact-whitespace" || tc.name == "literal-json" {
				mode = 2
			}
			got := run(t, record{where, filepath.Join(where, "private"), "call", args, mode}, 0)
			if _, e := os.Stat(marker); !os.IsNotExist(e) {
				t.Fatal("plain help/usage/selector accessed Git repository", e)
			}
			if got.code != tc.code {
				t.Error("plain before-access result", got)
			}
			if tc.code == 0 {
				if got.err != "" || !strings.Contains(got.out, "verify [<change>]") {
					t.Error("plain before-access help streams", got)
				}
			} else {
				if got.out != "" || !strings.Contains(got.err, tc.diagnostic) {
					t.Error("plain before-access diagnostic", got)
				}
				if tc.diagnostic == "C05" {
					_, e := w.Verify(context.Background(), where, args[len(args)-1])
					if e == nil {
						t.Fatal("public invalid selector premise")
					}
					cliVerifyFatalPlain(t, got, "C05", e.Error(), tc.code)
				}
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
 args:=[]string{"verify","current"};if os.Args[3]=="json"{args=append(args,"--json")};var out,errOut bytes.Buffer;code:=app.Run(args,&out,&errOut);rest,e:=io.ReadAll(os.Stdin);if e!=nil{os.Exit(81)};after,e:=os.Getwd();if e!=nil||after!=cwd||!reflect.DeepEqual(env,os.Environ()){os.Exit(82)}
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
	run := func(t *testing.T, f cliVerifyCase, mode string, formats ...string) ([]byte, string) {
		t.Helper()
		ctx, stop := context.WithTimeout(context.Background(), 5*time.Second)
		defer stop()
		format := "json"
		if len(formats) > 0 {
			format = formats[0]
		}
		cmd := exec.CommandContext(ctx, binary, f.Where, mode, format)
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

	for _, mode := range []string{"success", "structured-rejection", "coverage-failure", "fatal-source", "fatal-discovery"} {
		for _, format := range []string{"json", "plain"} {
			name := mode
			if format == "plain" {
				name += "-plain"
			}
			t.Run(name, func(t *testing.T) {
				f := cliVerifyFixture(t, "impl", "sha1", "all", false, false)
				switch mode {
				case "structured-rejection":
					projectPut(t, f.Root, "units/a/evals/case.txt", "forbidden")
				case "coverage-failure":
					cliVerifyReport(t, f, "alpha", true)
				case "fatal-source":
					if e := w.ReleaseClaim(f.Root, "current", f.Claim.Object); e != nil {
						t.Fatal(e)
					}
				}
				var reset func()
				if mode == "fatal-discovery" {
					reset = cliVerifyDiscoveryFault(t, f)
				}
				want, publicErr := w.Verify(context.Background(), f.Where, "current")
				reports := cliVerifyReports(t, f)
				cliVerifyClear(t, f)
				if reset != nil {
					if publicErr == nil || !strings.HasPrefix(publicErr.Error(), "C01:") {
						t.Fatal("stdio C01 public premise", want, publicErr)
					}
					reset()
				}
				if (mode == "fatal-source" || mode == "fatal-discovery") != (publicErr != nil) {
					t.Fatal("stdio public fatal premise", mode, publicErr)
				}
				if mode == "coverage-failure" && (len(want.Coverage) == 0 || cliVerifyExit(want) != 1) {
					t.Fatal("stdio coverage failure premise", want)
				}
				raw, stderr := run(t, f, "call", format)
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
				actual := cliVerifyReports(t, f)
				if publicErr != nil {
					code := strings.SplitN(publicErr.Error(), ":", 2)[0]
					exit := 1
					if code == "C01" || code == "D05" {
						exit = 2
					}
					if format == "plain" {
						cliVerifyFatalPlain(t, r, code, publicErr.Error(), exit)
					} else {
						x := decode(t, r, "verify")
						var envelope map[string]any
						if e := json.Unmarshal([]byte(r.out), &envelope); e != nil {
							t.Fatal(e)
						}
						ds, ok := envelope["diagnostics"].([]any)
						if !ok || len(ds) != 1 {
							t.Fatal("stdio fatal diagnostic count", r)
						}
						diag := ds[0].(map[string]any)
						if r.code != exit || string(x.Data) != "null" || diag["code"] != code || diag["severity"] != "error" || diag["message"] != publicErr.Error() {
							t.Error("stdio fatal classification", r, publicErr)
						}
						if _, ok := diag["path"]; ok {
							t.Error("stdio fatal invented path", diag)
						}
					}
					cliVerifyObserve(t, f, []string{}, binary)
				} else {
					if r.code != cliVerifyExit(want) {
						t.Error("stdio exit", r)
					}
					if format == "json" {
						x := decode(t, r, "verify")
						if !reflect.DeepEqual(cliVerifyNormalize(decodeData(t, x.Data), actual), cliVerifyWire(t, want, reports)) {
							t.Error("stdio call changed evidence", r)
						}
						var env map[string]any
						json.Unmarshal([]byte(r.out), &env)
						if !reflect.DeepEqual(cliVerifyNormalize(env["diagnostics"], actual), cliVerifyNormalize(canonical(t, want.Diagnostics), reports)) {
							t.Error("stdio diagnostics", r)
						}
					} else {
						var data any
						if e := json.Unmarshal([]byte(r.out), &data); e != nil {
							t.Fatal("stdio plain JSON", e, r)
						}
						if !strings.Contains(r.out, "\n  ") || !reflect.DeepEqual(cliVerifyNormalize(data, actual), cliVerifyWire(t, want, reports)) {
							t.Error("stdio plain evidence", r)
						}
						cliVerifyPlainStreams(t, r.err, want, actual, reports)
					}
					cliVerifyObserveVerification(t, f, want, binary)
				}
			})
		}
	}
}

// Failed output writers run in a bounded child with live caller streams.
func cliVerifyWriterProbe(t *testing.T) (string, func(*testing.T, cliVerifyCase, string) result) {
	t.Helper()
	temp := t.TempDir()
	source := `package main
import("bytes";"encoding/json";"fmt";"io";"os";"reflect";app "github.com/lurkshark/fireweed/packages/app")
type failedWriter struct{}
func(failedWriter)Write(p []byte)(int,error){return 0,fmt.Errorf("controlled output failure")}
func main(){if os.Chdir(os.Args[1])!=nil{os.Exit(80)};mode:=os.Args[2];cwd,_:=os.Getwd();env:=os.Environ();var out,errOut bytes.Buffer;var code int
switch mode{case "fatal-writer":code=app.Run([]string{"verify","missing"},&out,failedWriter{});case "usage-writer":code=app.Run([]string{"verify","a","b"},&out,failedWriter{});case "diagnostic-writer":code=app.Run([]string{"verify","current"},&out,failedWriter{});default:code=app.Run([]string{"verify","current","--json"},failedWriter{},&errOut)}
rest,e:=io.ReadAll(os.Stdin);if e!=nil{os.Exit(81)};after,e:=os.Getwd();if e!=nil||after!=cwd||!reflect.DeepEqual(env,os.Environ()){os.Exit(82)}
json.NewEncoder(os.Stdout).Encode(struct{Code int;Out,Err,Input string}{code,out.String(),errOut.String(),string(rest)})}`
	projectPut(t, temp, "writer.go", source)
	binary := filepath.Join(temp, "writer")
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}
	root, e := filepath.Abs("../../..")
	if e != nil {
		t.Fatal(e)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "build", "-o", binary, filepath.Join(temp, "writer.go"))
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, e := cmd.CombinedOutput(); e != nil {
		t.Fatal("writer probe build", e, string(out))
	}
	return binary, func(t *testing.T, f cliVerifyCase, mode string) result {
		t.Helper()
		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancel()
		cmd := exec.CommandContext(ctx, binary, f.Where, mode)
		input := "writer failure stdin sentinel\n"
		cmd.Stdin = strings.NewReader(input)
		var out, errOut bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &errOut
		if e := cmd.Run(); e != nil {
			t.Fatal("writer probe process", e, out.String(), errOut.String())
		}
		if errOut.Len() != 0 {
			t.Error("writer failure emitted raw stderr prompt", errOut.String())
		}
		var record struct {
			Code            int
			Out, Err, Input string
		}
		decoder := json.NewDecoder(&out)
		if e := decoder.Decode(&record); e != nil {
			t.Fatal("writer failure raw stdout prompt", e, out.String())
		}
		var extra any
		if e := decoder.Decode(&extra); e != io.EOF {
			t.Error("writer failure added stdout", e)
		}
		if record.Input != input {
			t.Error("writer failure consumed stdin", record.Input)
		}
		return result{record.Code, record.Out, record.Err}
	}
}
