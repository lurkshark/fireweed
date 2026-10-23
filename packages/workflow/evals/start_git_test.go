package evals

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	w "github.com/lurkshark/fireweed/packages/workflow"
)

type startProxyConfig struct{ Real, Mode, Marker, WinnerObject, Change, WinnerText, BaseRef, BaseOld, BaseNew, Blob string }

func startProxyBuild(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	source := `package main
import("os";"os/exec";"context";"time";"encoding/json";"bytes";"io";"strings";"path/filepath";c "github.com/lurkshark/fireweed/packages/contracts")
type config struct{Real,Mode,Marker,WinnerObject,Change,WinnerText,BaseRef,BaseOld,BaseNew,Blob string}
func main(){
 raw,e:=os.ReadFile(os.Getenv("START_PROXY_CONFIG"));if e!=nil{os.Exit(80)};var in config;if json.Unmarshal(raw,&in)!=nil{os.Exit(81)}
 if strings.HasPrefix(filepath.Base(os.Args[0]),"git-remote-"){os.WriteFile(in.Marker+"-transport",[]byte("attempted lazy transport\n"),0600);os.Exit(93)}
 input,e:=io.ReadAll(os.Stdin);if e!=nil{os.Exit(82)};args:=os.Args[1:]
 run:=func(a []string,b []byte)([]byte,[]byte,int){ctx,cancel:=context.WithTimeout(context.Background(),8*time.Second);defer cancel();cmd:=exec.CommandContext(ctx,in.Real,a...);cmd.Stdin=bytes.NewReader(b);cmd.Env=os.Environ();var out,err bytes.Buffer;cmd.Stdout=&out;cmd.Stderr=&err;e:=cmd.Run();code:=0;if e!=nil{if x,ok:=e.(*exec.ExitError);ok{code=x.ExitCode()}else{code=83}};return out.Bytes(),err.Bytes(),code}
 seen:=func()bool{_,e:=os.Stat(in.Marker);return e==nil}
 mark:=func(){f,e:=os.OpenFile(in.Marker,os.O_WRONLY|os.O_CREATE|os.O_EXCL,0600);if e!=nil{os.Exit(84)};f.WriteString("fired\n");f.Close()}
 contains:=func(s string)bool{for _,a:=range args{if a==s{return true}};return false}
 forbidden:=false;for _,op:=range []string{"commit","commit-tree","add","reset","restore","checkout","switch","branch","worktree","fetch","push","clean","gc","update-index"}{if contains(op){forbidden=true}}
 if forbidden{os.WriteFile(in.Marker+"-forbidden",[]byte(strings.Join(args,"\n")),0600);os.Stderr.WriteString("injected forbidden Git operation\n");os.Exit(91)}
 if in.Mode=="move-base"&&!seen()&&(contains("show")||contains("cat-file")&&!contains("-e")&&!contains("-t")&&!contains("--batch-check")||contains("ls-tree")){
  query:=strings.Join(args," ")+" "+string(input)
  if strings.Contains(query,in.BaseOld)||strings.Contains(query,in.BaseRef)||in.Blob!=""&&strings.Contains(query,in.Blob){
   _,detail,code:=run([]string{"update-ref",in.BaseRef,in.BaseNew},nil);if code!=0{os.Stderr.Write(detail);os.Exit(85)};mark()
  }
 }
 if in.Mode=="warn"{os.Stderr.WriteString("warning: independent successful Git stderr must stay outside stdout protocol\n")}
 out,detail,code:=run(args,input)
 if code==0&&!seen()&&(in.Mode=="replace-reservation"||in.Mode=="replace-state"||in.Mode=="replace-change"||in.Mode=="lost-response"){
  ref:="refs/fireweed/claims/current";oid,_,status:=run([]string{"rev-parse","--verify","--end-of-options",ref},nil)
  if status==0{
   fire:=true
   if in.Mode=="replace-state"||in.Mode=="replace-change"{
    payload,_,status:=run([]string{"cat-file","blob",strings.TrimSpace(string(oid))},nil);var cl struct{State string};json.Unmarshal(payload,&cl);ch,e:=c.ParseChange(cl.State,"changes/0001-current.md","ALP");fire=status==0&&e==nil&&ch.Run!=""&&ch.Phase!=""
   }
   if fire{
    if in.Change!=""{if e:=os.WriteFile(in.Change,[]byte(in.WinnerText),0600);e!=nil{os.Exit(86)}}
    if in.WinnerObject!=""{_,d,s:=run([]string{"update-ref",ref,in.WinnerObject},nil);if s!=0{os.Stderr.Write(d);os.Exit(87)}}
    mark();if in.Mode=="lost-response"{detail=append(detail,[]byte("injected reservation response lost after successful atomic creation\n")...);code=92}
   }
  }
 }
 os.Stdout.Write(out);os.Stderr.Write(detail);os.Exit(code)
}`
	worktreePut(t, root, "proxy.go", source)
	name := "git"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	binary := filepath.Join(root, name)
	workspace, err := filepath.Abs("../../..")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "build", "-o", binary, filepath.Join(root, "proxy.go"))
	cmd.Dir = workspace
	cmd.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Git proxy compile %v %s", err, out)
	}
	return binary
}
func startProxyInstall(t *testing.T, binary string, cfg startProxyConfig) startProxyConfig {
	t.Helper()
	real, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	cfg.Real = real
	cfg.Marker = filepath.Join(t.TempDir(), "event")
	path := filepath.Join(t.TempDir(), "config.json")
	b, err := json.Marshal(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, b, 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("START_PROXY_CONFIG", path)
	t.Setenv("PATH", filepath.Dir(binary)+string(os.PathListSeparator)+os.Getenv("PATH"))
	return cfg
}
func startProxyFired(t *testing.T, cfg startProxyConfig) {
	t.Helper()
	b, err := os.ReadFile(cfg.Marker)
	if err != nil || string(b) != "fired\n" {
		t.Error("required externally controlled event not reached", err)
	}
	if b, err := os.ReadFile(cfg.Marker + "-forbidden"); err == nil {
		t.Error("unexpected Git mutation or network attempt", string(b))
	}
}
func startWinner(t *testing.T, f startFixtureData) (string, string) {
	t.Helper()
	raw := strings.Replace(f.Raw, "module: alpha\n", "module: alpha\nrun: external-winner\nphase: impl\n", 1)
	raw = strings.ReplaceAll(raw, "new:feature", "ALP-051")
	raw = strings.Replace(raw, "Start the exact authored change.", "Externally retained winning content.", 1)
	cl := w.Claim{Worktree: f.Root, Change: "current", Run: "external-winner", Base: claimBase(t, f.Root), State: raw}
	b, err := json.Marshal(cl)
	if err != nil {
		t.Fatal(err)
	}
	oid, err := claimCommand(t, f.Root, string(b), "hash-object", "-w", "--stdin")
	if err != nil {
		t.Fatal(err)
	}
	return oid, raw
}
func Test_FLOW038_FLOW040StartGitProxyControls(t *testing.T) {
	binary := startProxyBuild(t)
	for _, mode := range []string{"replace-reservation", "replace-state", "replace-change", "lost-response", "move-base"} {
		t.Run(mode, func(t *testing.T) {
			f := startFixture(t, "sha1", "spec", "")
			winner, raw := startWinner(t, f)
			base := claimBase(t, f.Root)
			cfg := startProxyConfig{Mode: mode, WinnerObject: winner}
			if mode == "replace-change" {
				cfg.WinnerObject = ""
				cfg.Change = filepath.Join(f.Root, f.Path)
				cfg.WinnerText = raw
			}
			if mode == "move-base" {
				cfg.BaseOld = base
				cfg.BaseNew = claimGit(t, f.Root, "rev-parse", "HEAD^")
				cfg.BaseRef = "refs/heads/approval-base"
			}
			cfg = startProxyInstall(t, binary, cfg)
			var err error
			if mode == "move-base" {
				_, err = claimCommand(t, f.Root, "", "show", base+":"+f.Path)
			} else {
				_, err = claimCommand(t, f.Root, "", "update-ref", "refs/fireweed/claims/current", winner, strings.Repeat("0", 40))
			}
			if (err != nil) != (mode == "lost-response") {
				t.Error("proxy response control", mode, err)
			}
			startProxyFired(t, cfg)
			switch mode {
			case "move-base":
				if tip := claimGit(t, f.Root, "rev-parse", cfg.BaseRef); tip != cfg.BaseNew {
					t.Error("base event did not move ref", tip)
				}
			case "replace-change":
				b, e := os.ReadFile(cfg.Change)
				if e != nil || string(b) != raw {
					t.Error("change event did not replace exact bytes", e)
				}
			default:
				if oid := claimGit(t, f.Root, "rev-parse", "refs/fireweed/claims/current"); oid != winner {
					t.Error("winner/lost-response ref missing", oid)
				}
			}
		})
	}
	t.Run("stderr-protocol", func(t *testing.T) {
		f := startFixture(t, "sha1", "spec", "")
		cfg := startProxyInstall(t, binary, startProxyConfig{Mode: "warn"})
		out, err := claimCommand(t, f.Root, "", "rev-parse", "HEAD")
		if err != nil || !strings.Contains(out, "independent successful Git stderr") {
			t.Error("warn control not live", out, err)
		}
		if _, err := os.Stat(cfg.Marker + "-forbidden"); !os.IsNotExist(err) {
			t.Error("unexpected proxy operation")
		}
	})
}
func Test_FLOW038StartConcurrentReplacementAndLostResponse(t *testing.T) {
	binary := startProxyBuild(t)
	for _, mode := range []string{"replace-reservation", "replace-state", "replace-change", "lost-response"} {
		t.Run(mode, func(t *testing.T) {
			f := startFixture(t, "sha1", "spec", "")
			winner, raw := startWinner(t, f)
			cfg := startProxyConfig{Mode: mode, WinnerObject: winner}
			if mode == "replace-state" || mode == "replace-change" {
				cfg.Change = filepath.Join(f.Root, f.Path)
				cfg.WinnerText = raw
			}
			if mode == "replace-change" || mode == "lost-response" {
				cfg.WinnerObject = ""
			}
			cfg = startProxyInstall(t, binary, cfg)
			tree := startTree(t, f.Root)
			git := startGitSnapshot(t, f.Root)
			got, err := w.Start(f.Root, "current")
			startProxyFired(t, cfg)
			if err == nil || !policyCorrective(err.Error()) || !reflect.DeepEqual(got, w.StartResult{}) {
				t.Errorf("ownership/failure did not reject with zero result: %+v %v", got, err)
			}
			if mode == "lost-response" && (err == nil || !strings.Contains(err.Error(), "injected reservation response lost")) {
				t.Error("lost original failed-command details", err)
			}
			if cfg.Change != "" {
				node := tree[f.Path]
				node.Bytes = raw
				tree[f.Path] = node
			}
			if !reflect.DeepEqual(tree, startTree(t, f.Root)) {
				t.Error("rollback overwrote external or unrelated file state")
			}
			after := startGitSnapshot(t, f.Root)
			if cfg.WinnerObject != "" {
				git.Refs += "\n" + winner + " refs/fireweed/claims/current"
				lines := strings.Split(git.Refs, "\n")
				sortStartRefLines(lines)
				git.Refs = strings.Join(lines, "\n")
			}
			if !reflect.DeepEqual(git, after) {
				t.Error("rollback deleted replacement ref or left owned claim/index/HEAD changes")
			}
		})
	}
}
func sortStartRefLines(lines []string) {
	// Git show-ref orders by ref name, rather than by the object column.
	for i := 1; i < len(lines); i++ {
		for j := i; j > 0; j-- {
			a := strings.SplitN(lines[j-1], " ", 2)
			b := strings.SplitN(lines[j], " ", 2)
			if len(a) != 2 || len(b) != 2 || a[1] <= b[1] {
				break
			}
			lines[j-1], lines[j] = lines[j], lines[j-1]
		}
	}
}
func Test_FLOW036_FLOW040StartPinnedBaseAndWarnedGit(t *testing.T) {
	binary := startProxyBuild(t)
	for _, mode := range []string{"warn", "move-base", "replacement-base"} {
		t.Run(mode, func(t *testing.T) {
			f := startFixture(t, "sha1", "spec", "")
			old := claimBase(t, f.Root)
			cfg := startProxyConfig{Mode: mode}
			if mode == "move-base" || mode == "replacement-base" {
				worktreePut(t, f.Root, f.Path, strings.Replace(f.Raw, "Start the exact authored change.", "Unapproved replacement base content.", 1))
				startCommit(t, f.Root, "Temporary alternate base content")
				alternate := claimBase(t, f.Root)
				claimGit(t, f.Root, "reset", "--hard", old)
				cfg.BaseOld = old
				cfg.BaseNew = alternate
				cfg.BaseRef = "refs/heads/approval-base"
				cfg.Blob = claimGit(t, f.Root, "rev-parse", old+":"+f.Path)
				if mode == "replacement-base" {
					replacement := claimGit(t, f.Root, "-c", "user.name=Start Fixture", "-c", "user.email=start@example.invalid", "commit-tree", claimGit(t, f.Root, "rev-parse", alternate+"^{tree}"), "-p", old+"^", "-m", "Replacement view with an acyclic parent")
					claimGit(t, f.Root, "-c", "user.name=Start Fixture", "-c", "user.email=start@example.invalid", "commit", "--allow-empty", "-m", "Current HEAD with unchanged authored tree")
					claimGit(t, f.Root, "replace", old, replacement)
					if text := claimGit(t, f.Root, "show", old+":"+f.Path); !strings.Contains(text, "Unapproved replacement base content.") {
						t.Fatal("real replacement view control")
					}
					if status := claimGit(t, f.Root, "status", "--porcelain"); status != "" {
						t.Fatal("replacement fixture not clean", status)
					}
				}
			}
			cfg = startProxyInstall(t, binary, cfg)
			before := startGitSnapshot(t, f.Root)
			got, err := w.Start(f.Root, "current")
			if err != nil {
				t.Fatalf("valid warned/pinned start: %v", err)
			}
			if mode == "move-base" {
				startProxyFired(t, cfg)
				before.Refs = strings.Replace(before.Refs, old+" "+cfg.BaseRef, cfg.BaseNew+" "+cfg.BaseRef, 1)
			}
			startAssertSuccess(t, f, "evals", before, got)
			for _, d := range got.Diagnostics {
				if d.Code == "P04" {
					t.Error("re-read moved base or mixed successful stderr into approval", d)
				}
			}
			if b, e := os.ReadFile(cfg.Marker + "-forbidden"); e == nil {
				t.Error("forbidden subprocess operation", string(b))
			}
		})
	}
}
func Test_FLOW040StartNativeGitHookControl(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("POSIX executable hook fixture; portable Git protocol controls run separately")
	}
	f := startFixture(t, "sha1", "spec", "")
	marker := contextNativeHook(t, f.Root)
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Error("native hook suppression control did not remain clean")
	}
}
func Test_FLOW037_FLOW040StartSuppressesNativeGitHooks(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("POSIX executable hook fixture; portable start state assertions run separately")
	}
	f := startFixture(t, "sha1", "spec", "")
	marker := contextNativeHook(t, f.Root)
	startSuccess(t, f, f.Root, "evals")
	if b, err := os.ReadFile(marker); err == nil {
		t.Error("start invoked proven fsmonitor hook", string(b))
	} else if !os.IsNotExist(err) {
		t.Error(err)
	}
}
func startMissingPromisorFixture(t *testing.T, binary string) (startFixtureData, startProxyConfig) {
	t.Helper()
	f := startFixture(t, "sha1", "spec", "")
	base := claimBase(t, f.Root)
	blob := claimGit(t, f.Root, "rev-parse", base+":"+f.Path)
	f.Raw += "\nRuntime feedback differs from the unavailable base blob.\n"
	worktreePut(t, f.Root, f.Path, f.Raw)
	startCommit(t, f.Root, "Current source uses a distinct available blob")
	object := filepath.Join(claimCommon(t, f.Root), "objects", blob[:2], blob[2:])
	payload, err := os.ReadFile(object)
	if err != nil {
		t.Fatal("required loose base blob premise", err)
	}
	if err := os.Remove(object); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.MkdirAll(filepath.Dir(object), 0755); os.WriteFile(object, payload, 0444) })
	for key, value := range map[string]string{"extensions.partialClone": "origin", "remote.origin.promisor": "true", "remote.origin.partialclonefilter": "blob:none", "remote.origin.url": "startfetch::controlled", "protocol.startfetch.allow": "always"} {
		claimGit(t, f.Root, "config", key, value)
	}
	data, err := os.ReadFile(binary)
	if err != nil {
		t.Fatal(err)
	}
	name := "git-remote-startfetch"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	if err := os.WriteFile(filepath.Join(filepath.Dir(binary), name), data, 0755); err != nil {
		t.Fatal(err)
	}
	cfg := startProxyInstall(t, binary, startProxyConfig{Mode: "delegate"})
	t.Setenv("GIT_NO_LAZY_FETCH", "0")
	_, err = claimCommand(t, f.Root, "", "cat-file", "blob", blob)
	if err == nil {
		t.Fatal("missing blob unexpectedly readable")
	}
	marker := cfg.Marker + "-transport"
	if b, e := os.ReadFile(marker); e != nil || string(b) != "attempted lazy transport\n" {
		t.Fatal("real promisor transport positive control failed", e, string(b))
	}
	if err := os.Remove(marker); err != nil {
		t.Fatal(err)
	}
	return f, cfg
}
func Test_FLOW036_FLOW040StartPromisorTransportControl(t *testing.T) {
	binary := startProxyBuild(t)
	f, cfg := startMissingPromisorFixture(t, binary)
	cmd := exec.Command(cfg.Real, "-c", "core.fsmonitor=false", "status", "--porcelain")
	cmd.Dir = f.Root
	out, err := cmd.CombinedOutput()
	if err != nil || len(out) != 0 {
		t.Fatal("current committed source must remain readable/clean", err, string(out))
	}
	if _, err := os.Stat(cfg.Marker + "-transport"); !os.IsNotExist(err) {
		t.Error("current clean check attempted unavailable unrelated historical blob")
	}
}
func Test_FLOW036_FLOW040StartDoesNotLazyFetchBase(t *testing.T) {
	binary := startProxyBuild(t)
	f, cfg := startMissingPromisorFixture(t, binary)
	got := startSuccess(t, f, f.Root, "evals")
	unapproved := false
	for _, d := range got.Diagnostics {
		if d.Code == "P04" && d.Severity == "warning" && (strings.Contains(strings.ToLower(d.Message), "not approved") || strings.Contains(strings.ToLower(d.Message), "unapproved")) {
			unapproved = true
		}
	}
	if !unapproved {
		t.Error("unavailable base content implied approval", got.Diagnostics)
	}
	if b, err := os.ReadFile(cfg.Marker + "-transport"); err == nil {
		t.Error("start triggered proven promisor lazy fetch", string(b))
	} else if !os.IsNotExist(err) {
		t.Error(err)
	}
}
