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

	c "github.com/lurkshark/fireweed/packages/contracts"
)

func Test_FLOW028_FLOW029_FLOW030_FLOW031ApprovalNativeNoIO(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("macOS post-startup native guard; portable content assertions run separately")
	}
	root, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	worktreePut(t, root, "guard.c", strings.ReplaceAll(eligibilityNativeGuard, "fw_eligible_arm", "fw_approval_arm"))
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	lib := filepath.Join(root, "libapprovalguard.dylib")
	cmd := exec.CommandContext(ctx, "cc", "-dynamiclib", "-o", lib, filepath.Join(root, "guard.c"), "-Wl,-install_name,"+lib)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("native guard compile: %v %s", err, out)
	}
	program := fmt.Sprintf(`package main
/*
#cgo LDFLAGS: -L%s -lapprovalguard
void fw_approval_arm(int m,const char*p);
*/
import "C"
import("os";"os/exec";"net";"time";"fmt";"encoding/json";"reflect";"strings";c "github.com/lurkshark/fireweed/packages/contracts";w "github.com/lurkshark/fireweed/packages/workflow")
type record struct{Module c.Module;Base,Current,Path,Private,Address string;Want,Error bool}
func main(){
 raw,err:=os.ReadFile(os.Args[2]);if err!=nil{os.Exit(80)};var in record;if json.Unmarshal(raw,&in)!=nil{os.Exit(81)}
 mode:=os.Args[1];before,_:=json.Marshal(in.Module);cwd,err:=os.Getwd();if err!=nil{os.Exit(82)};env:=os.Environ()
 C.fw_approval_arm(2,nil)
 switch mode{
 case "read-control":os.ReadFile(in.Private)
 case "missing-control":os.ReadFile(in.Private+"/missing")
 case "stat-control":os.Stat(in.Private)
 case "cwd-control":os.Getwd()
 case "directory-control":os.ReadDir(in.Private)
 case "stdin-control":os.Stdin.Read(make([]byte,1))
 case "prompt-control":fmt.Fprint(os.Stdout,"prompt")
 case "write-control":os.WriteFile(in.Private,[]byte("bad"),0644)
 case "network-control":net.DialTimeout("tcp",in.Address,time.Second)
 case "git-control":exec.Command("git","status").Run()
 case "hook-control":exec.Command("hook-trap").Run()
 case "absolute-process-control":exec.Command("/usr/bin/true").Run()
 case "source-control":
  for _,text:=range []string{in.Base,in.Current}{ch,err:=c.ParseChange(text,in.Path,in.Module.IDPrefix);if err!=nil||len(ch.Added)!=2{os.Exit(83)};if _,err=c.RenderMarkdown(ch.Document);err!=nil{os.Exit(84)}}
  C.fw_approval_arm(0,nil);return
 default:
  got,err:=w.ApprovalMatches(in.Module,in.Base,in.Current,in.Path)
  if got!=in.Want||(err!=nil)!=in.Error{C.fw_approval_arm(0,nil);fmt.Fprintf(os.Stderr,"got %%v,%%v want %%v error=%%v",got,err,in.Want,in.Error);os.Exit(10)}
  if err!=nil&&!strings.HasPrefix(err.Error(),"C05:"){os.Exit(11)}
  after,_:=json.Marshal(in.Module);if string(before)!=string(after)||!reflect.DeepEqual(env,os.Environ()){os.Exit(12)}
  C.fw_approval_arm(0,nil);afterCwd,err:=os.Getwd();if err!=nil||cwd!=afterCwd{os.Exit(13)};return
 }
 C.fw_approval_arm(0,nil);os.Exit(70)
}`, root)
	worktreePut(t, root, "probe.go", program)
	source, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	binary := filepath.Join(root, "probe")
	cmd = exec.CommandContext(ctx, "go", "build", "-o", binary, filepath.Join(root, "probe.go"))
	cmd.Dir = source
	cmd.Env = append(os.Environ(), "CGO_ENABLED=1", "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("public native approval probe: %v %s", err, out)
	}
	private := filepath.Join(root, "private")
	worktreePut(t, root, "private", "Untouched private bytes.")
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()
	type record struct {
		Module                                c.Module
		Base, Current, Path, Private, Address string
		Want, Error                           bool
	}
	m, base, current := contentApprovalPair(t)
	in := record{Module: m, Base: base, Current: current, Path: contentApprovalPath, Private: private, Address: listener.Addr().String(), Want: true}
	run := func(t *testing.T, mode string, in record, want int) {
		t.Helper()
		raw, err := json.Marshal(in)
		if err != nil {
			t.Fatal(err)
		}
		worktreePut(t, root, "record.json", string(raw))
		cmd := exec.CommandContext(ctx, binary, mode, filepath.Join(root, "record.json"))
		cmd.Dir = root
		out, err := cmd.CombinedOutput()
		code := 0
		if err != nil {
			exit, ok := err.(*exec.ExitError)
			if !ok {
				t.Fatal(err)
			}
			code = exit.ExitCode()
		}
		if code != want {
			t.Fatalf("native %s exit%d want%d: %s", mode, code, want, out)
		}
	}
	for _, mode := range []string{"read-control", "missing-control", "stat-control", "cwd-control", "directory-control", "stdin-control", "prompt-control", "write-control", "network-control", "git-control", "hook-control", "absolute-process-control"} {
		t.Run(mode, func(t *testing.T) { run(t, mode, in, 42) })
	}
	t.Run("source-control", func(t *testing.T) { run(t, "source-control", in, 0) })
	t.Run("matching", func(t *testing.T) { run(t, "subject", in, 0) })
	t.Run("mismatch", func(t *testing.T) {
		x := in
		x.Current = strings.Replace(current, "Keep approval tied", "Changed approval tied", 1)
		contentApprovalParse(t, x.Current)
		x.Want = false
		run(t, "subject", x, 0)
	})
	t.Run("invalid-base", func(t *testing.T) { x := in; x.Base = "malformed"; x.Want = false; run(t, "subject", x, 0) })
	t.Run("invalid-current", func(t *testing.T) {
		x := in
		x.Current = "malformed"
		x.Want, x.Error = false, true
		run(t, "subject", x, 0)
	})
	t.Run("wrong-target", func(t *testing.T) {
		x := in
		x.Current = strings.Replace(current, "module: alpha", "module: bravo", 1)
		contentApprovalParse(t, x.Current)
		x.Want, x.Error = false, true
		run(t, "subject", x, 0)
	})
	if raw, err := os.ReadFile(private); err != nil || string(raw) != "Untouched private bytes." {
		t.Fatal("private bytes changed", err)
	}
}
