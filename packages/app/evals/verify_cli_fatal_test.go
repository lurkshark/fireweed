package evals

import (
	"context"
	"encoding/json"
	w "github.com/lurkshark/fireweed/packages/workflow"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

// This controlled local Git failure occurs after successful initial discovery,
// at the held-approval loader's public discovery boundary. The public Verify
// result independently establishes the C01 premise before the CLI is tested.
func cliVerifyDiscoveryFault(t *testing.T, f cliVerifyCase) func() {
	t.Helper()
	real, e := exec.LookPath("git")
	if e != nil {
		t.Fatal(e)
	}
	dir := t.TempDir()
	counter := filepath.Join(f.Plan, "discovery-count")
	cfg, _ := json.Marshal(struct{ Real, Counter string }{real, counter})
	projectPut(t, dir, "fault.json", string(cfg))
	source := `package main
import("encoding/json";"fmt";"os";"os/exec";"path/filepath"
 "strconv";"strings")
func main(){self,e:=os.Executable();if e!=nil{os.Exit(80)};raw,e:=os.ReadFile(filepath.Join(filepath.Dir(self),"fault.json"));var p struct{Real,Counter string};if e!=nil||json.Unmarshal(raw,&p)!=nil{os.Exit(81)}
 for _,a:=range os.Args[1:]{if a=="--show-toplevel"{raw,_:=os.ReadFile(p.Counter);n,_:=strconv.Atoi(strings.TrimSpace(string(raw)));n++;if os.WriteFile(p.Counter,[]byte(strconv.Itoa(n)),0600)!=nil{os.Exit(82)};if n==2{fmt.Fprintln(os.Stderr,"controlled local approval discovery failure");os.Exit(93)};break}}
 cmd:=exec.Command(p.Real,os.Args[1:]...);cmd.Stdin=os.Stdin;cmd.Stdout=os.Stdout;cmd.Stderr=os.Stderr;if e:=cmd.Run();e!=nil{if x,ok:=e.(*exec.ExitError);ok{os.Exit(x.ExitCode())};os.Exit(83)}}`
	projectPut(t, dir, "fault.go", source)
	bin := filepath.Join(dir, "git")
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "build", "-o", bin, filepath.Join(dir, "fault.go"))
	cmd.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, e := cmd.CombinedOutput(); e != nil {
		t.Fatalf("local discovery control build: %v %s", e, out)
	}
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
	reset := func() { t.Helper(); projectPut(t, f.Plan, "discovery-count", "0") }
	reset()
	return reset
}

func Test_APP026_APP028VerifyFatalDiscovery(t *testing.T) {
	f := cliVerifyFixture(t, "impl", "sha1", "all", false, false)
	reset := cliVerifyDiscoveryFault(t, f)
	var err error
	if !t.Run("public-control", func(t *testing.T) {
		v, e := w.Verify(context.Background(), f.Where, "current")
		err = e
		if err == nil || !strings.HasPrefix(err.Error(), "C01:") || !strings.Contains(err.Error(), "controlled local approval discovery failure") || v.Change != "" || len(cliVerifyCalls(t, f)) != 0 {
			t.Fatal("public C01 discovery control", v, err)
		}
	}) {
		return
	}
	reset()
	cliVerifyReadonly(t, f, func() {
		raw, e := os.ReadFile(filepath.Join(f.Plan, "discovery-count"))
		if e != nil || string(raw) != "0" {
			t.Fatal("discovery fixture was consumed by harness setup", e, string(raw))
		}
		got := invoke(t, f.Where, "verify", "current", "--json")
		x := decode(t, got, "verify")
		var envelope map[string]any
		if e := json.Unmarshal([]byte(got.out), &envelope); e != nil {
			t.Fatal(e)
		}
		ds, ok := envelope["diagnostics"].([]any)
		if !ok || len(ds) != 1 {
			t.Fatal("fatal discovery diagnostic count", got)
		}
		diag := ds[0].(map[string]any)
		if got.code != 2 || string(x.Data) != "null" || diag["code"] != "C01" || diag["severity"] != "error" || diag["message"] != err.Error() {
			t.Error("fatal discovery JSON classification", got, err)
		}
		if _, ok := diag["path"]; ok {
			t.Error("fatal discovery invented path", diag)
		}
	})
	reset()
	var got result
	cliVerifyReadonly(t, f, func() { got = invoke(t, f.Where, "verify", "current") })
	cliVerifyFatalPlain(t, got, "C01", err.Error(), 2)
}
