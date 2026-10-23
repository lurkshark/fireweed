package evals_test

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

// Shared only by fixture interceptors. Words in values and object-type arguments
// are not operations. Query forms of branch/worktree remain available.
const projectGitOperationPolicy = `
func fwGitCommand(args []string)(string,[]string){
 for i:=0;i<len(args);i++{a:=args[i];switch a{case "-c","-C","--git-dir","--work-tree","--namespace","--exec-path","--config-env","--super-prefix":i++;continue};if strings.HasPrefix(a,"-"){continue};return a,args[i+1:]};return "",nil
}
func fwGitForbidden(args []string)bool{
 command,tail:=fwGitCommand(args)
 switch command{
 case "fetch","fetch-pack","push","clone","ls-remote","commit","checkout","switch","reset","restore","add","update-ref":return true
 case "hash-object":for i:=0;i<len(tail);i++{a:=tail[i];if a=="--"{break};if a=="-t"||a=="--path"{i++;continue};if strings.HasPrefix(a,"-")&&!strings.HasPrefix(a,"--")&&strings.Contains(a[1:],"w"){return true}};return false
 case "worktree":return len(tail)==0||tail[0]!="list"
 case "branch":
  query:=false;operand:=false
  for i:=0;i<len(tail);i++{a:=tail[i]
   if a=="--"{if i+1<len(tail){operand=true};break}
   switch a{
   case "--list","--show-current","--all","--remotes":query=true;continue
   case "--verbose","--quiet","--color","--no-color","--column","--no-column","--ignore-case","--omit-empty":continue
   case "--sort","--format":if i+1<len(tail){i++};continue
   case "--contains","--no-contains","--merged","--no-merged","--points-at":query=true;if i+1<len(tail)&&!strings.HasPrefix(tail[i+1],"-"){i++};continue
   }
   if strings.HasPrefix(a,"--sort=")||strings.HasPrefix(a,"--format="){continue}
   if strings.HasPrefix(a,"--contains=")||strings.HasPrefix(a,"--no-contains=")||strings.HasPrefix(a,"--merged=")||strings.HasPrefix(a,"--no-merged=")||strings.HasPrefix(a,"--points-at="){query=true;continue}
   if strings.HasPrefix(a,"--color=")||strings.HasPrefix(a,"--column="){continue}
   if a=="--delete"||a=="--force"||a=="--move"||a=="--copy"||a=="--edit-description"||a=="--unset-upstream"||a=="--set-upstream-to"||a=="--set-upstream"||a=="--track"||strings.HasPrefix(a,"--set-upstream-to=")||strings.HasPrefix(a,"--track="){return true}
   if strings.HasPrefix(a,"--"){continue}
   if strings.HasPrefix(a,"-"){if strings.ContainsAny(a[1:],"dDmMcCfFtu"){return true};if strings.ContainsAny(a,"ar"){query=true};continue}
   operand=true
  }
  // No operand with only formatting options is an ordinary branch listing.
  return operand&&!query
 };return false
}
`

var projectGitTrapMu sync.Mutex
var projectGitTrapBytes []byte

func installProjectGitTrap(t *testing.T, dir, real, marker string) {
	t.Helper()
	projectGitTrapMu.Lock()
	defer projectGitTrapMu.Unlock()
	if projectGitTrapBytes == nil {
		build := t.TempDir()
		source := filepath.Join(build, "git.go")
		binary := filepath.Join(build, "git-trap")
		program := `package main
import("encoding/json";"os";"os/exec";"path/filepath";"strings")
` + projectGitOperationPolicy + `
func main(){var in struct{Real,Marker string};self,e:=os.Executable();if e!=nil{os.Exit(79)};raw,e:=os.ReadFile(filepath.Join(filepath.Dir(self),"git-trap.json"));if e!=nil||json.Unmarshal(raw,&in)!=nil{os.Exit(80)};if fwGitForbidden(os.Args[1:]){os.WriteFile(in.Marker,[]byte("attempted forbidden Git operation"),0600);os.Exit(97)};child:=exec.Command(in.Real,os.Args[1:]...);child.Stdin=os.Stdin;child.Stdout=os.Stdout;child.Stderr=os.Stderr;if e:=child.Run();e!=nil{if x,ok:=e.(*exec.ExitError);ok{os.Exit(x.ExitCode())};os.Exit(81)}}
`
		if err := os.WriteFile(source, []byte(program), 0600); err != nil {
			t.Fatal(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
		defer cancel()
		cmd := exec.CommandContext(ctx, "go", "build", "-o", binary, source)
		cmd.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("operation trap build: %v %s", err, out)
		}
		var err error
		projectGitTrapBytes, err = os.ReadFile(binary)
		if err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(dir, "git"), projectGitTrapBytes, 0755); err != nil {
		t.Fatal(err)
	}
	raw, err := json.Marshal(struct{ Real, Marker string }{real, marker})
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "git-trap.json"), raw, 0600); err != nil {
		t.Fatal(err)
	}
}
