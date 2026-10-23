package evals

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
	"testing"
	"time"
)

var startGitMu sync.Mutex
var startGitBytes []byte

// Observe actual network Git verbs, independently of command result or config
// and path argument text. All other operations are delegated without change.
func cliStartLocalGit(t *testing.T, marker string) string {
	t.Helper()
	startGitMu.Lock()
	defer startGitMu.Unlock()
	real, e := exec.LookPath("git")
	if e != nil {
		t.Fatal(e)
	}
	dir := t.TempDir()
	binary := filepath.Join(dir, "git")
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}
	if startGitBytes == nil {
		program := `package main
import("encoding/json";"os";"os/exec";"path/filepath";"strings")
func main(){self,e:=os.Executable();if e!=nil{os.Exit(80)};raw,e:=os.ReadFile(filepath.Join(filepath.Dir(self),"start-git.json"));var p struct{Real,Marker string};if e!=nil||json.Unmarshal(raw,&p)!=nil{os.Exit(81)}
 deny:=func(){os.WriteFile(p.Marker,[]byte(strings.Join(os.Args[1:],"\n")),0600);os.Exit(93)}
 args:=os.Args[1:];op:="";var rest []string
 for i:=0;i<len(args);i++{a:=args[i];if a=="-c"||a=="-C"||a=="--git-dir"||a=="--work-tree"||a=="--namespace"||a=="--config-env"{i++;continue};if strings.HasPrefix(a,"-"){continue};op=a;rest=args[i+1:];break}
 switch op{case "fetch","push","pull","clone","ls-remote":deny();case "remote","submodule":sub:="";for _,a:=range rest{if strings.HasPrefix(a,"-"){continue};sub=a;break};if op=="remote"&&sub=="update"||op=="submodule"&&(sub=="update"||sub=="add"){deny()}}
 cmd:=exec.Command(p.Real,args...);cmd.Stdin=os.Stdin;cmd.Stdout=os.Stdout;cmd.Stderr=os.Stderr;if e:=cmd.Run();e!=nil{if x,ok:=e.(*exec.ExitError);ok{os.Exit(x.ExitCode())};os.Exit(82)}
}`
		source := filepath.Join(dir, "proxy.go")
		if e = os.WriteFile(source, []byte(program), 0600); e != nil {
			t.Fatal(e)
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		cmd := exec.CommandContext(ctx, "go", "build", "-o", binary, source)
		cmd.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
		if out, e := cmd.CombinedOutput(); e != nil {
			t.Fatalf("local Git observer build %v %s", e, out)
		}
		startGitBytes, e = os.ReadFile(binary)
		if e != nil {
			t.Fatal(e)
		}
	}
	if e = os.WriteFile(binary, startGitBytes, 0755); e != nil {
		t.Fatal(e)
	}
	raw, _ := json.Marshal(struct{ Real, Marker string }{real, marker})
	if e = os.WriteFile(filepath.Join(dir, "start-git.json"), raw, 0600); e != nil {
		t.Fatal(e)
	}
	for _, args := range [][]string{{"fetch", "--dry-run", "invalid"}, {"-c", "example.value=harmless", "push", "invalid"}, {"ls-remote", "invalid"}, {"remote", "update"}} {
		cmd := exec.Command(binary, args...)
		out, e := cmd.CombinedOutput()
		x, ok := e.(*exec.ExitError)
		if !ok || x.ExitCode() != 93 || len(out) != 0 {
			t.Fatalf("Git denied control %v: %v %s", args, e, out)
		}
		if _, e = os.Stat(marker); e != nil {
			t.Fatal("Git trap not observable", e)
		}
		if e = os.Remove(marker); e != nil {
			t.Fatal(e)
		}
	}
	// Exercise harmless subcommands with command-looking remote/path names.
	control := filepath.Join(dir, "controls")
	if e = os.Mkdir(control, 0755); e != nil {
		t.Fatal(e)
	}
	local := func(args ...string) {
		cmd := exec.Command(real, args...)
		cmd.Dir = control
		if out, e := cmd.CombinedOutput(); e != nil {
			t.Fatalf("Git calibration %v %v %s", args, e, out)
		}
	}
	local("init", "-q", "--object-format=sha1")
	local("remote", "add", "update", "/local/only")
	projectPut(t, control, ".gitmodules", "[submodule \"update\"]\n path = update\n url = /local/only\n")
	local("add", ".gitmodules")
	local("update-index", "--add", "--cacheinfo", "160000,1111111111111111111111111111111111111111,update")
	// A word resembling a command inside a config value is not an operation.
	for _, args := range [][]string{{"--version"}, {"-c", "example.value=fetch", "--version"}, {"remote", "get-url", "update"}, {"submodule", "status", "update"}} {
		realCmd := exec.Command(real, args...)
		realCmd.Dir = control
		want, e := realCmd.CombinedOutput()
		if e != nil {
			t.Fatalf("real allowed Git premise %v %v %s", args, e, want)
		}
		proxyCmd := exec.Command(binary, args...)
		proxyCmd.Dir = control
		got, e := proxyCmd.CombinedOutput()
		if e != nil || string(got) != string(want) {
			t.Fatalf("Git allowed control %v: %v %s/%s", args, e, got, want)
		}
		if _, e = os.Stat(marker); !os.IsNotExist(e) {
			t.Fatal("allowed Git flagged", e)
		}
	}
	return dir
}
