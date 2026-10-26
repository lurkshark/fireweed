package evals

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

type heldGitPolicy struct {
	Real, Trace, Violation, Stderr string
	Fail                           bool
}

// The proxy passes streams through without buffering stdin or prescribing a
// show/cat-file/body strategy. It observes Git operations, not implementation
// functions. A native implementation using no Git process remains permitted.
const heldGitObserverProgram = `package main
import("encoding/json";"fmt";"os";"os/exec";"path/filepath";"strings")
type policy struct{Real,Trace,Violation,Stderr string;Fail bool}
func main(){
 self,e:=os.Executable();if e!=nil{os.Exit(80)};raw,e:=os.ReadFile(filepath.Join(filepath.Dir(self),"held-git-policy.json"));var p policy;if e!=nil||json.Unmarshal(raw,&p)!=nil{os.Exit(81)}
 args:=os.Args[1:];log,e:=os.OpenFile(p.Trace,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0600);if e!=nil{os.Exit(82)};json.NewEncoder(log).Encode(args);log.Close()
 op:="";tail:=[]string{}
 for i:=0;i<len(args);i++{a:=args[i];if a=="-c"||a=="-C"||a=="--git-dir"||a=="--work-tree"||a=="--namespace"||a=="--config-env"{i++;continue};if strings.HasPrefix(a,"-"){continue};op=a;tail=args[i+1:];break}
 allowed:=false
 switch op{
 case "rev-parse","show-ref","for-each-ref","ls-tree","cat-file","show","rev-list","ls-files","diff","status","check-ref-format","merge-base","archive","version":allowed=true
 case "symbolic-ref":n:=0;allowed=true;for _,a:=range tail{if a=="-d"||a=="--delete"{allowed=false};if !strings.HasPrefix(a,"-"){n++}};allowed=allowed&&n==1
 case "config":for _,a:=range tail{if a=="--get"||a=="--get-all"||a=="--get-regexp"||a=="--get-urlmatch"||a=="--list"||a=="-l"{allowed=true}};for _,a:=range tail{if a=="--add"||a=="--replace-all"||a=="--unset"||a=="--unset-all"||a=="--remove-section"||a=="--rename-section"||a=="--edit"{allowed=false}}
 case "hash-object":allowed=true;for _,a:=range tail{if a=="-w"||a=="--literally"{allowed=false}}
 }
 if !allowed{os.WriteFile(p.Violation,[]byte(strings.Join(args,"\n")),0600);fmt.Fprint(os.Stderr,"held observer denied mutating/network Git operation");os.Exit(93)}
 if p.Fail{fmt.Fprint(os.Stderr,p.Stderr);os.Exit(92)}
 command:=exec.Command(p.Real,args...);command.Stdin=os.Stdin;command.Stdout=os.Stdout;command.Stderr=os.Stderr;err:=command.Run();if p.Stderr!=""{fmt.Fprint(os.Stderr,p.Stderr)}
 if err!=nil{if x,ok:=err.(*exec.ExitError);ok{code:=x.ExitCode();if code<0{code=255};os.Exit(code)};fmt.Fprint(os.Stderr,err);os.Exit(83)}
}
`

func heldBuildGitObserver(t *testing.T) (string, string) {
	t.Helper()
	directory := stateCanonicalDirectory(t, t.TempDir())
	worktreePut(t, directory, "proxy.go", heldGitObserverProgram)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	binary := filepath.Join(directory, "git")
	if os.PathSeparator == '\\' {
		binary += ".exe"
	}
	command := exec.CommandContext(ctx, "go", "build", "-o", binary, filepath.Join(directory, "proxy.go"))
	command.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := command.CombinedOutput(); err != nil {
		t.Fatalf("held Git observer build: %v %s", err, out)
	}
	return directory, binary
}

func heldWriteGitPolicy(t *testing.T, directory string, policy heldGitPolicy) {
	t.Helper()
	raw, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(directory, "held-git-policy.json"), raw, 0600); err != nil {
		t.Fatal(err)
	}
}

func heldRunGitControl(t *testing.T, binary, root, input string, args ...string) (string, string, int) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	command := exec.CommandContext(ctx, binary, args...)
	command.Dir = root
	command.Stdin = strings.NewReader(input)
	var out, errOut strings.Builder
	command.Stdout, command.Stderr = &out, &errOut
	err := command.Run()
	if ctx.Err() != nil {
		t.Fatal(ctx.Err())
	}
	code := 0
	if err != nil {
		if exit, ok := err.(*exec.ExitError); ok {
			code = exit.ExitCode()
		} else {
			t.Fatal(err)
		}
	}
	return out.String(), errOut.String(), code
}

func Test_FLOW048_FLOW049_FLOW051HeldGitStreamsAndOperations(t *testing.T) {
	directory, binary := heldBuildGitObserver(t)
	real := protocolRealGit(t)
	for _, format := range []string{"sha1", "sha256"} {
		for _, mode := range []string{"held", "none", "foreign", "malformed"} {
			t.Run(format+"/"+mode, func(t *testing.T) {
				f := heldSourceFixture(t, "impl", format, true)
				policyDir := t.TempDir()
				policy := heldGitPolicy{Real: real, Trace: filepath.Join(policyDir, "trace.jsonl"), Violation: filepath.Join(policyDir, "violation"), Stderr: "valid-looking stderr: refs/fireweed/claims/current\n"}
				heldWriteGitPolicy(t, directory, policy)
				for _, args := range [][]string{{"rev-parse", "--show-toplevel"}, {"cat-file", "blob", f.Claim.Object}, {"for-each-ref", "--format=%(refname) %(objectname)", "refs/fireweed/claims/"}} {
					want, _, code := heldRunGitControl(t, real, f.Root, "", args...)
					got, diagnostics, actual := heldRunGitControl(t, binary, f.Root, "", args...)
					if code != actual || got != want || diagnostics != policy.Stderr {
						t.Fatal("native stream-preserving observer control", args, code, actual, got, want, diagnostics)
					}
				}
				for _, args := range [][]string{{"fetch", "--all"}, {"push"}, {"checkout", "-b", "forbidden"}, {"update-ref", "refs/fireweed/claims/forbidden", f.Claim.Object}, {"worktree", "add", "forbidden"}} {
					_, _, code := heldRunGitControl(t, binary, f.Root, "", args...)
					if code != 93 {
						t.Fatal("forbidden Git operation control", args, code)
					}
					if _, err := os.Stat(policy.Violation); err != nil {
						t.Fatal(err)
					}
					if err := os.Remove(policy.Violation); err != nil {
						t.Fatal(err)
					}
				}
				switch mode {
				case "none":
					claimGit(t, f.Root, "update-ref", "-d", "refs/fireweed/claims/current", f.Claim.Object)
				case "foreign":
					other := f.Claim
					other.Worktree = heldPeer(t, f.Root)
					claimRaw(t, f.Root, "current", claimPayload(t, other))
				case "malformed":
					claimRaw(t, f.Root, "current", "{broken unclassified source")
				}
				if err := os.Remove(policy.Trace); err != nil {
					t.Fatal(err)
				}
				heldPoisonCurrent(t, f.Root)
				t.Setenv("PATH", directory+string(os.PathListSeparator)+os.Getenv("PATH"))
				switch mode {
				case "held":
					heldSourceResult(t, f.Root, f)
				case "none", "foreign":
					heldAbsentResult(t, f.Root)
				case "malformed":
					heldRejectedResult(t, f.Root)
				}
				if _, err := os.Stat(policy.Violation); !os.IsNotExist(err) {
					t.Fatal("held source attempted forbidden Git operation", err)
				}
			})
		}
	}
}

func Test_FLOW051HeldRequiredGitFailureDetails(t *testing.T) {
	directory, binary := heldBuildGitObserver(t)
	real := protocolRealGit(t)
	f := heldSourceFixture(t, "impl", "sha1", true)
	artifacts := t.TempDir()
	policy := heldGitPolicy{Real: real, Trace: filepath.Join(artifacts, "trace.jsonl"), Violation: filepath.Join(artifacts, "violation"), Fail: true, Stderr: "held-selected-local-git-failure-9284"}
	heldWriteGitPolicy(t, directory, policy)
	defer func() {
		if _, err := os.Stat(policy.Violation); !os.IsNotExist(err) {
			t.Error("held source fault recovery attempted forbidden Git operation", err)
		}
	}()
	for _, args := range [][]string{{"rev-parse", "--show-toplevel"}, {"cat-file", "blob", f.Claim.Object}} {
		if out, diagnostic, code := heldRunGitControl(t, binary, f.Root, "", args...); out != "" || diagnostic != policy.Stderr || code != 92 {
			t.Fatal("live native failure control", args, out, diagnostic, code)
		}
	}
	controlTrace, err := os.ReadFile(policy.Trace)
	if err != nil {
		t.Fatal(err)
	}
	operations := heldGitTraceOperations(t, controlTrace)
	if !heldGitCommandDetail("git cat-file failed; restore objects", operations) || !heldGitCommandDetail("git rev-parse failed; use a valid checkout", operations) || heldGitCommandDetail("git imaginary-command failed; restore objects", operations) {
		t.Fatal("later required-command diagnostic observer control", operations)
	}
	if _, _, code := heldRunGitControl(t, binary, f.Root, "", "fetch", "--all"); code != 93 {
		t.Fatal("fault-policy forbidden operation observer control", code)
	}
	if marker, err := os.ReadFile(policy.Violation); err != nil || !strings.Contains(string(marker), "fetch") {
		t.Fatal("fault-policy forbidden operation marker control", string(marker), err)
	}
	if err := os.Remove(policy.Violation); err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(policy.Trace); err != nil {
		t.Fatal(err)
	}
	heldPoisonCurrent(t, f.Root)
	t.Setenv("PATH", directory+string(os.PathListSeparator)+os.Getenv("PATH"))
	source, claim, present, err := w.LoadHeldApproval(f.Root)
	trace, traceErr := os.ReadFile(policy.Trace)
	if os.IsNotExist(traceErr) {
		// A conforming native reader need not launch Git. Its ordinary valid
		// source result is still required; an inert stub cannot pass here.
		if err != nil || !present || claim != f.Claim || !reflect.DeepEqual(source, f.Want) {
			t.Error("no Git failure occurred, but available authority was rejected", err, present)
		}
		return
	}
	if traceErr != nil || len(trace) == 0 {
		t.Fatal("Git invocation trace", traceErr, string(trace))
	}
	if err == nil && present && claim == f.Claim && reflect.DeepEqual(source, f.Want) {
		// A failed optional Git probe may legitimately fall back to a complete
		// native reader. Observing a command does not make it required.
		return
	}
	operations = heldGitTraceOperations(t, trace)
	if err == nil || present || claim != (w.Claim{}) || !reflect.DeepEqual(source, p.Snapshot{}) || !strings.Contains(err.Error(), policy.Stderr) || !policyCorrective(err.Error()) {
		t.Error("required Git failure lost useful command/stderr or leaked authority", err, present, claim)
	}
	if err != nil && !heldGitCommandDetail(err.Error(), operations) {
		t.Error("required Git failure omitted actual observed command details", operations, err)
	}

}

func heldGitTraceOperations(t *testing.T, trace []byte) []string {
	t.Helper()
	operations := []string{}
	for _, line := range strings.Split(strings.TrimSpace(string(trace)), "\n") {
		var args []string
		if err := json.Unmarshal([]byte(line), &args); err != nil || len(args) == 0 {
			t.Fatal("native command trace shape", string(trace), err)
		}
		operation := ""
		for index := 0; index < len(args); index++ {
			argument := args[index]
			if argument == "-c" || argument == "-C" || argument == "--git-dir" || argument == "--work-tree" || argument == "--namespace" || argument == "--config-env" {
				index++
				continue
			}
			if strings.HasPrefix(argument, "-") {
				continue
			}
			operation = argument
			break
		}
		if operation == "" {
			t.Fatal("native command trace has no operation", args)
		}
		operations = append(operations, operation)
	}
	return operations
}

func heldGitCommandDetail(message string, operations []string) bool {
	for _, operation := range operations {
		if strings.Contains(message, operation) {
			return true
		}
	}
	return false
}
