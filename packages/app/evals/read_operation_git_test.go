package evals

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"
)

var readOperationGitMu sync.Mutex
var readOperationGitBytes []byte

// This observer classifies the operation, never words used as data. The separate
// context observer enforces which immutable source contents may be opened.
func installReadOperationGitProxy(t *testing.T, directory, marker string) {
	t.Helper()
	readOperationGitMu.Lock()
	defer readOperationGitMu.Unlock()
	real, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	binary := filepath.Join(directory, "git")
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}
	if readOperationGitBytes == nil {
		source := filepath.Join(directory, "observer.go")
		program := `package main
import("encoding/json";"os";"os/exec";"path/filepath";"strings")
func main(){self,e:=os.Executable();if e!=nil{os.Exit(80)};raw,e:=os.ReadFile(filepath.Join(filepath.Dir(self),"read-git.json"));var p struct{Real,Marker string};if e!=nil||json.Unmarshal(raw,&p)!=nil{os.Exit(81)}
args:=os.Args[1:];op:="";var tail []string
for i:=0;i<len(args);i++{a:=args[i];switch a{case "-c","-C","--git-dir","--work-tree","--namespace","--config-env","--exec-path","--super-prefix":i++;continue};if strings.HasPrefix(a,"-"){continue};op=a;tail=args[i+1:];break}
deny:=func(){os.WriteFile(p.Marker,[]byte(strings.Join(args,"\n")),0600);os.Exit(97)}
switch op{
case "","version","rev-parse","check-ref-format","ls-files","status","diff","diff-files","diff-index","diff-tree","ls-tree","cat-file","show","show-ref","for-each-ref","rev-list":
case "symbolic-ref":n:=0;for _,a:=range tail{if a=="--delete"||a=="-d"{deny()};if !strings.HasPrefix(a,"-"){n++}};if n!=1{deny()}
case "config":read:=false;for i:=0;i<len(tail);i++{a:=tail[i];switch a{
case "--file","-f","--blob","--default","--type","-t":i++;if i>=len(tail){deny()}
case "--get","--get-all","--get-regexp","--get-urlmatch","--list","-l":read=true
case "--local","--global","--system","--worktree","--includes","--no-includes","--null","-z","--name-only","--show-origin","--show-scope","--bool","--int","--bool-or-int","--path","--expiry-date","--fixed-value":
default:if strings.HasPrefix(a,"--file=")||strings.HasPrefix(a,"--blob=")||strings.HasPrefix(a,"--default=")||strings.HasPrefix(a,"--type="){continue};if strings.HasPrefix(a,"-"){deny()}
}};if !read{deny()}
case "remote","submodule":sub:="";for _,a:=range tail{if strings.HasPrefix(a,"-"){continue};sub=a;break};if op=="remote"&&sub!=""&&sub!="get-url"||op=="submodule"&&sub!=""&&sub!="status"{deny()}
default:deny()
}
c:=exec.Command(p.Real,args...);c.Stdin=os.Stdin;c.Stdout=os.Stdout;c.Stderr=os.Stderr;if e:=c.Run();e!=nil{if x,ok:=e.(*exec.ExitError);ok{os.Exit(x.ExitCode())};os.Exit(82)}}`
		if err = os.WriteFile(source, []byte(program), 0600); err != nil {
			t.Fatal(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		c := exec.CommandContext(ctx, "go", "build", "-o", binary, source)
		c.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
		if out, e := c.CombinedOutput(); e != nil {
			t.Fatalf("observer build %v %s", e, out)
		}
		readOperationGitBytes, err = os.ReadFile(binary)
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = os.WriteFile(binary, readOperationGitBytes, 0755); err != nil {
		t.Fatal(err)
	}
	raw, _ := json.Marshal(struct{ Real, Marker string }{real, marker})
	if err = os.WriteFile(filepath.Join(directory, "read-git.json"), raw, 0600); err != nil {
		t.Fatal(err)
	}
	// Every allowed control must succeed through real Git and return identical
	// bytes through the observer. An ignored Git error is never a passing control.
	control := t.TempDir()
	realRun := func(args ...string) []byte {
		c := exec.Command(real, args...)
		c.Dir = control
		out, e := c.CombinedOutput()
		if e != nil {
			t.Fatalf("real Git premise %v: %v %s", args, e, out)
		}
		return out
	}
	realRun("init", "-q", "--object-format=sha1", "commit")
	projectPut(t, control, "commit/nested/retained", "selected immutable bytes\n")
	realRun("-C", "commit", "add", ".")
	realRun("-C", "commit", "remote", "add", "update", "/local/only")
	// -C takes the exact command-looking operand "commit" relative to control.
	realRun("-C", "commit", "-c", "user.name=Eval", "-c", "user.email=eval@local.invalid", "commit", "--allow-empty", "-qm", "fixture")
	base := strings.TrimSpace(string(realRun("-C", "commit", "rev-parse", "HEAD")))
	tree := strings.TrimSpace(string(realRun("-C", "commit", "rev-parse", "HEAD^{tree}")))
	subtree := strings.TrimSpace(string(realRun("-C", "commit", "rev-parse", "HEAD:nested")))
	allowed := [][]string{
		{"-C", "commit", "cat-file", "commit", base},
		{"-C", "commit", "-c", "example.value=fetch", "cat-file", "commit", base},
		{"-C", "commit", "--namespace", "push", "rev-parse", "--verify", base},
		{"-C", "commit", "cat-file", "--batch-all-objects", "--batch-check=%(objectname)"},
		{"-C", "commit", "rev-list", "--no-walk", "--missing=print", "--end-of-options", base},
		{"-C", "commit", "ls-tree", tree},
		{"-C", "commit", "ls-tree", subtree},
		{"-C", "commit", "remote", "get-url", "update"},
		{"-C", "commit", "config", "--get", "remote.update.url"},
		{"-C", "commit", "config", "--file", ".git/config", "--get", "remote.update.url"},
	}
	for _, args := range allowed {
		want := realRun(args...)
		c := exec.Command(binary, args...)
		c.Dir = control
		got, e := c.CombinedOutput()
		if e != nil || string(got) != string(want) {
			t.Fatalf("allowed Git %v: %v %s/%s", args, e, got, want)
		}
		if _, e = os.Stat(marker); !os.IsNotExist(e) {
			t.Fatal("allowed read flagged", args, e)
		}
	}
	for _, op := range []string{"fetch", "push", "pull", "clone", "ls-remote", "commit", "branch", "checkout", "switch", "reset", "restore", "add", "worktree", "update-ref", "hash-object", "stash", "clean", "rm", "mv", "notes", "bisect", "tag", "remote", "submodule"} {
		args := []string{"-C", "commit", "-c", "example.value=cat-file", op}
		if op == "remote" || op == "submodule" {
			args = append(args, "update")
		}
		c := exec.Command(binary, args...)
		c.Dir = control
		out, e := c.CombinedOutput()
		x, ok := e.(*exec.ExitError)
		if !ok || x.ExitCode() != 97 || len(out) != 0 {
			t.Fatalf("denied Git %v: %v %s", args, e, out)
		}
		if _, e = os.Stat(marker); e != nil {
			t.Fatal("denied control marker absent", e)
		}
		if e = os.Remove(marker); e != nil {
			t.Fatal(e)
		}
	}
	for _, tail := range [][]string{{"remote", "add", "control", "/local/only"}, {"remote", "rm", "update"}, {"notes", "add", "-m", "control"}, {"bisect", "reset"}, {"submodule", "add", "/local/only", "control"}, {"sparse-checkout", "add", "control"}, {"submodule", "foreach", "git", "commit"}, {"config", "example.key", "commit"}, {"symbolic-ref", "HEAD", "refs/heads/control"}, {"config", "--file", "--get", "example.key", "commit"}, {"config", "-f", "--get", "example.key", "commit"}} {
		args := append([]string{"-C", "commit"}, tail...)
		c := exec.Command(binary, args...)
		c.Dir = control
		out, e := c.CombinedOutput()
		x, ok := e.(*exec.ExitError)
		if !ok || x.ExitCode() != 97 || len(out) != 0 {
			t.Fatalf("nested mutation %v: %v %s", args, e, out)
		}
		if _, e = os.Stat(marker); e != nil {
			t.Fatal("nested mutation marker absent", e)
		}
		if e = os.Remove(marker); e != nil {
			t.Fatal(e)
		}
	}
}

func Test_APP021ReadOperationObserverControls(t *testing.T) {
	directory := t.TempDir()
	installReadOperationGitProxy(t, directory, filepath.Join(directory, "attempted"))
}
