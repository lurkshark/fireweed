// Protected, project-owned Go hook. No third-party dependencies.
package main
import (
 "bytes"
 "encoding/json"
 "encoding/xml"
 "fmt"
 "go/parser"
 "go/token"
 "io/fs"
 "os"
 "os/exec"
 "path/filepath"
 "strconv"
 "strings"
)
type Case struct { Name string `xml:"name,attr"`; Class string `xml:"classname,attr"`; Failure *string `xml:"failure,omitempty"`; Skipped *string `xml:"skipped,omitempty"` }
type Suite struct { XMLName xml.Name `xml:"testsuite"`; Cases []Case `xml:"testcase"` }
func main() {
 if len(os.Args)<2 { os.Exit(2) }
 if os.Args[1]=="eval" {
  cmd:=exec.Command("go","test","-json","./"+os.Args[2]+"/..."); var output bytes.Buffer; cmd.Stdout=&output; cmd.Stderr=os.Stderr; err:=cmd.Run()
  dec:=json.NewDecoder(&output); suite:=Suite{}; indexes:=map[string]int{}; logs:=map[string]string{}
  for { var ev struct { Action,Package,Test,Output string }; if dec.Decode(&ev)!=nil { break }; if ev.Test=="" { if ev.Output!="" { fmt.Print(ev.Output) }; continue }; key:=ev.Package+"/"+ev.Test
   if ev.Action=="run" { indexes[key]=len(suite.Cases);suite.Cases=append(suite.Cases,Case{Name:ev.Test,Class:ev.Package}) }
   logs[key]+=ev.Output
   if ev.Action=="fail" { if i,ok:=indexes[key];ok { text:=logs[key];suite.Cases[i].Failure=&text } }
   if ev.Action=="skip" { if i,ok:=indexes[key];ok { text:=logs[key];suite.Cases[i].Skipped=&text } }
  }; b,e:=xml.Marshal(suite);if e!=nil {panic(e)};if e=os.WriteFile(os.Args[3],b,0644);e!=nil{panic(e)};if err!=nil {os.Exit(1)};return
 }
 bin:=os.Getenv("FIREWEED_BIN");if bin==""{bin="fireweed"};cmd:=exec.Command(bin,"modules","--json");data,err:=cmd.Output();if err!=nil{panic(err)}
 var envelope struct { Data []struct { Name,Dir string; DependsOn []string `json:"depends_on"` } };if err=json.Unmarshal(data,&envelope);err!=nil{panic(err)}
 mod,err:=os.ReadFile("go.mod");if err!=nil{panic(err)};modulePath:="";for _,line:=range strings.Split(string(mod),"\n"){f:=strings.Fields(line);if len(f)==2&&f[0]=="module"{modulePath=strings.Trim(f[1],"\"")}}
 violations:=0
 for _,m:=range envelope.Data { err=filepath.WalkDir(m.Dir,func(file string,d fs.DirEntry,err error)error{
  if err!=nil{return err};if d.IsDir()||!strings.HasSuffix(file,".go"){return nil};tree,err:=parser.ParseFile(token.NewFileSet(),file,nil,parser.ImportsOnly);if err!=nil{return err}
  for _,imp:=range tree.Imports { imported,_:=strconv.Unquote(imp.Path.Value);for _,target:=range envelope.Data { base:=modulePath+"/"+target.Dir;if imported!=base&&!strings.HasPrefix(imported,base+"/"){continue};allowed:=m.Name==target.Name;for _,dep:=range m.DependsOn{if dep==target.Name{allowed=true}}
   if !allowed || (strings.Contains(filepath.ToSlash(file),"/evals/")&&imported!=base) || (m.Name!=target.Name&&imported!=base) {fmt.Printf("%s: %s imports %s outside its declared interface/dependencies\n",m.Name,file,imported);violations++}
  }};return nil });if err!=nil{fmt.Fprintln(os.Stderr,err);os.Exit(2)} }
 if violations>0{os.Exit(1)}
}
