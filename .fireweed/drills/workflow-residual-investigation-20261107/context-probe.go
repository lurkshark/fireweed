package main
import("encoding/json";"fmt";"os";"os/exec";"path/filepath";"strings";"github.com/lurkshark/fireweed/packages/assets";"github.com/lurkshark/fireweed/packages/workflow")
func main(){root,e:=os.MkdirTemp("","fireweed-recorded-context-probe-");must(e);defer os.RemoveAll(root);root,e=filepath.EvalSymlinks(root);must(e)
files:=map[string]string{
"SYSTEM.md":"# Fixture system\n",
"fireweed.yaml":"fireweed: '>=0.1, <0.2'\nversion: 1\nbase: main\nmodules: [packages/*]\nmodule_defaults:\n  interface: [api.go]\n  evals: [evals/**]\n  disposable: [internal/**]\nhooks:\n  eval: [sh, -c, 'exit 1']\n  boundaries: [sh, -c, 'exit 0']\n",
"packages/alpha/SPEC.md":"---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n",
"packages/alpha/api.go":"package alpha\n",
"packages/alpha/internal/ready":"recorded phase implementation\n",
"changes/0001-probe.md":"---\nmodule: alpha\nbases: {}\nregenerate: true\nphase: impl\nrun: fixture-recorded\n---\n\n## Motivation\nResume the recorded implementation.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"}
for n,s:=range files{p:=filepath.Join(root,n);must(os.MkdirAll(filepath.Dir(p),0755));must(os.WriteFile(p,[]byte(s),0644))};git:=func(args ...string){c:=exec.Command("git",args...);c.Dir=root;c.Env=append(os.Environ(),"GIT_CONFIG_NOSYSTEM=1","GIT_AUTHOR_NAME=Fixture","GIT_AUTHOR_EMAIL=fixture@localhost","GIT_COMMITTER_NAME=Fixture","GIT_COMMITTER_EMAIL=fixture@localhost");b,e:=c.CombinedOutput();if e!=nil{panic(fmt.Sprintf("git %v: %v: %s",args,e,b))}};git("init","-q");git("checkout","-q","-b","main");git("add","-A");git("commit","-q","-m","recorded regenerate fixture")
s,e:=workflow.Start(root,"probe");must(e);_,fileErr:=os.Stat(filepath.Join(root,"packages/alpha/internal/ready"));context,e:=workflow.Context(root,"probe",false);must(e);regen,e:=assets.Playbook("regenerate");must(e);impl,e:=assets.Playbook("impl");must(e);must(workflow.Release(root,"probe"));must(json.NewEncoder(os.Stdout).Encode(map[string]any{"phase":s.Phase,"run":s.Run,"recorded_file_preserved":fileErr==nil,"context_contains_regenerate":strings.Contains(context,regen),"context_contains_ordinary_impl":strings.Contains(context,impl),"fixture_claim_released":true,"native_drills":0}))}
func must(e error){if e!=nil{panic(e)}}
