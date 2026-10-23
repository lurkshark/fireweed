package evals

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

type approvalGitPolicy struct {
	Real, Marker, Base, Selected, Root string
	RejectAll                          bool
	Content, Trees, Paths              map[string]bool
}

var approvalGitProxyMu sync.Mutex
var approvalGitProxyBytes []byte

func installApprovalGitProxy(t *testing.T, directory string, policy approvalGitPolicy) {
	t.Helper()
	approvalGitProxyMu.Lock()
	defer approvalGitProxyMu.Unlock()
	if approvalGitProxyBytes == nil {
		temp := t.TempDir()
		source := filepath.Join(temp, "proxy.go")
		binary := filepath.Join(temp, "proxy")
		program := `package main
import("bufio";"encoding/json";"io";"os";"os/exec";"path/filepath";"strings")
var p struct{Real,Marker,Base,Selected,Root string;Content,Trees,Paths map[string]bool;RejectAll bool}
func bad(){os.WriteFile(p.Marker,[]byte(strings.Join(os.Args[1:],"\n")),0600);os.Exit(93)}
func selected(q string)string{if p.Content[q]||p.Trees[q]{return q};ref,_,path:=strings.Cut(q,":");root:=ref;for _,suffix:=range []string{"^{object}","^{commit}","^{tree}","^{}","^0"}{root=strings.TrimSuffix(root,suffix)};if root!=p.Base&&!p.Trees[root]&&!(root!=ref&&p.Content[root]){return ""};if !path&&root==ref{return ""};out,e:=exec.Command(p.Real,"rev-parse","--verify",q).Output();if e!=nil{return ""};id:=strings.TrimSpace(string(out));if p.Content[id]||p.Trees[id]{return id};return ""}
func claim(q string)bool{return q==p.Selected||q==p.Selected+"^{object}"||q==p.Selected+"^{}"}
func body(q string)bool{return claim(q)||selected(q)!=""}
func tree(q string)bool{id:=selected(q);return q==p.Base||p.Trees[id]}
func main(){self,e:=os.Executable();if e!=nil{os.Exit(80)};raw,e:=os.ReadFile(filepath.Join(filepath.Dir(self),"approval-git-policy.json"));if e!=nil||json.Unmarshal(raw,&p)!=nil{os.Exit(81)}
 if p.RejectAll{bad()};args:=os.Args[1:];op:="";var tail []string
 for i:=0;i<len(args);i++{a:=args[i];if a=="-c"||a=="--git-dir"||a=="--work-tree"||a=="--namespace"||a=="--config-env"{i++;continue};if a=="-C"{if i+1>=len(args){bad()};i++;full,e:=filepath.Abs(args[i]);if e!=nil{bad()};full,e=filepath.EvalSymlinks(full);if e!=nil||(full!=p.Root&&!strings.HasPrefix(full,p.Root+string(os.PathSeparator))){bad()};check:=exec.Command(p.Real,"-C",full,"rev-parse","--show-toplevel");out,e:=check.Output();if e!=nil||strings.TrimSpace(string(out))!=p.Root{bad()};continue};if strings.HasPrefix(a,"-"){continue};op=a;tail=args[i+1:];break}
 for _,a:=range args{
  if strings.HasPrefix(a,"refs/fireweed/claims")&&a!=p.Selected&&a!=p.Selected+"^{object}"&&a!=p.Selected+"^{}"{bad()}
  if a=="--all"||a=="--glob"||strings.HasPrefix(a,"--glob=")||strings.HasPrefix(a,"--branches")||strings.HasPrefix(a,"--tags")||strings.HasPrefix(a,"--remotes")||(strings.HasPrefix(a,"--exclude")&&a!="--exclude-standard"){bad()}
 }
 mode:="";all:=false;delimiter:=byte('\n')
 switch op{
 case "rev-parse","check-ref-format","ls-files":
 case "symbolic-ref":n:=0;for _,a:=range tail{if a=="--delete"||a=="-d"{bad()};if !strings.HasPrefix(a,"-"){n++}};if n!=1{bad()}
 case "show-ref":verify:=false;for _,a:=range tail{if a=="--verify"{verify=true}};if !verify||len(tail)==0{bad()};last:=tail[len(tail)-1];if !strings.HasPrefix(last,"refs/")&&last!="HEAD"{bad()}
 case "rev-list":
  noWalk,selected:=false,false
  for i:=0;i<len(tail);i++{a:=tail[i];switch a{case "--no-walk","--no-walk=sorted","--no-walk=unsorted":noWalk=true;continue;case "--missing=print","--missing=allow-any":continue;case "--max-count":i++;continue};if a==p.Base||a==p.Base+"^{commit}"||a==p.Base+"^0"{selected=true;continue};if a=="--end-of-options"||strings.HasPrefix(a,"--max-count="){continue};bad()};if !noWalk||!selected{bad()}
 case "ls-tree":
  selected:=false;for _,a:=range tail{if strings.HasPrefix(a,"-"){continue};if tree(a){selected=true;break};bad()};if !selected{bad()}
 case "show":
  if len(tail)!=1||(!strings.Contains(tail[0],":")&&!claim(tail[0]))||!body(tail[0]){bad()}
 case "cat-file":
  for _,a:=range tail{if a=="--batch-all-objects"{all=true};if a=="--batch"||strings.HasPrefix(a,"--batch="){mode="body"};if a=="--batch-check"||strings.HasPrefix(a,"--batch-check="){mode="info"};if strings.HasPrefix(a,"--batch-command"){mode="command"};if a=="-z"||a=="-Z"{delimiter=0}}
  if all&&mode!="info"{bad()}
  if mode==""&&(len(tail)==0||!body(tail[len(tail)-1])){bad()}
 default:bad()
 }
 child:=exec.Command(p.Real,args...);child.Stdout=os.Stdout;child.Stderr=os.Stderr
 if mode=="body"||mode=="command"||(mode=="info"&&!all){
  pipe,e:=child.StdinPipe();if e!=nil{os.Exit(82)};if child.Start()!=nil{os.Exit(83)};reader:=bufio.NewReader(os.Stdin)
  for{record,e:=reader.ReadString(delimiter);if len(record)>0{q:=strings.TrimSuffix(record,string(delimiter));check:=true;if mode=="command"{if q=="flush"{check=false}else if strings.HasPrefix(q,"info "){q=strings.TrimPrefix(q,"info ")}else{q=strings.TrimPrefix(q,"contents ")}};if check&&!body(q){child.Process.Kill();pipe.Close();child.Wait();bad()};if _,err:=io.WriteString(pipe,record);err!=nil{break}};if e!=nil{break}};pipe.Close()
 }else{child.Stdin=os.Stdin;if child.Start()!=nil{os.Exit(84)}}
 if e:=child.Wait();e!=nil{if x,ok:=e.(*exec.ExitError);ok{os.Exit(x.ExitCode())};os.Exit(85)}
}
`
		if err := os.WriteFile(source, []byte(program), 0600); err != nil {
			t.Fatal(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
		defer cancel()
		cmd := exec.CommandContext(ctx, "go", "build", "-o", binary, source)
		cmd.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("context Git observer build %v %s", err, out)
		}
		var err error
		approvalGitProxyBytes, err = os.ReadFile(binary)
		if err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(directory, "git"), approvalGitProxyBytes, 0755); err != nil {
		t.Fatal(err)
	}
	raw, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(directory, "approval-git-policy.json"), raw, 0600); err != nil {
		t.Fatal(err)
	}
}
