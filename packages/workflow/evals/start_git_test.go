package evals

import (
	"bytes"
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

	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

type startProxyConfig struct{ Real, Mode, Marker, WinnerObject, Change, WinnerText, BaseRef, BaseOld, BaseNew, Blob, BaseText string }

func startProxyBuild(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	source := `package main
import("os";"os/exec";"context";"time";"encoding/json";"bytes";"io";"strings";"path/filepath";c "github.com/lurkshark/fireweed/packages/contracts")
type config struct{Real,Mode,Marker,WinnerObject,Change,WinnerText,BaseRef,BaseOld,BaseNew,Blob,BaseText string}
func main(){
 raw,e:=os.ReadFile(os.Getenv("START_PROXY_CONFIG"));if e!=nil{os.Exit(80)};var in config;if json.Unmarshal(raw,&in)!=nil{os.Exit(81)}
 if strings.HasPrefix(filepath.Base(os.Args[0]),"git-remote-"){os.WriteFile(in.Marker+"-transport",[]byte("attempted lazy transport\n"),0600);os.Exit(93)}
 if in.Mode=="hang"&&os.Getenv("START_PROXY_ARM_HANG")=="1"{os.WriteFile(in.Marker,[]byte("fired\n"),0600);pid,_:=json.Marshal(os.Getpid());os.WriteFile(in.Marker+"-pid",pid,0600);time.Sleep(5*time.Minute);os.Exit(94)}
 input,e:=io.ReadAll(os.Stdin);if e!=nil{os.Exit(82)};args:=os.Args[1:]
 run:=func(a []string,b []byte)([]byte,[]byte,int){ctx,cancel:=context.WithTimeout(context.Background(),8*time.Second);defer cancel();cmd:=exec.CommandContext(ctx,in.Real,a...);cmd.Stdin=bytes.NewReader(b);cmd.Env=os.Environ();var out,err bytes.Buffer;cmd.Stdout=&out;cmd.Stderr=&err;e:=cmd.Run();code:=0;if e!=nil{if x,ok:=e.(*exec.ExitError);ok{code=x.ExitCode()}else{code=83}};return out.Bytes(),err.Bytes(),code}
 seen:=func()bool{_,e:=os.Stat(in.Marker);return e==nil}
 mark:=func(){f,e:=os.OpenFile(in.Marker,os.O_WRONLY|os.O_CREATE|os.O_EXCL,0600);if e!=nil{os.Exit(84)};f.WriteString("fired\n");f.Close()}
 contains:=func(s string)bool{for _,a:=range args{if a==s{return true}};return false}
 forbidden:=false;for _,op:=range []string{"commit","commit-tree","add","reset","restore","checkout","switch","branch","worktree","fetch","push","clean","gc","update-index"}{if contains(op){forbidden=true}}
 if forbidden{os.WriteFile(in.Marker+"-forbidden",[]byte(strings.Join(args,"\n")),0600);os.Stderr.WriteString("injected forbidden Git operation\n");os.Exit(91)}
 if in.Mode=="warn"{os.Stderr.WriteString("warning: independent successful Git stderr must stay outside stdout protocol\n")}
 if in.Mode=="precondition"&&contains("update-ref"){os.WriteFile(in.Marker+"-precondition",[]byte(strings.Join(args,"\n")),0600);os.Stderr.WriteString("injected prohibited precondition claim mutation\n");os.Exit(95)}
 out,detail,code:=run(args,input)
 if code==0&&in.Mode=="move-base"&&!seen()&&(in.Blob!=""&&bytes.Contains(out,[]byte(in.Blob))||in.BaseText!=""&&bytes.Contains(out,[]byte(in.BaseText))){
  _,d,status:=run([]string{"update-ref",in.BaseRef,in.BaseNew},nil);if status!=0{os.Stderr.Write(d);os.Exit(85)};mark()
 }
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
	for _, mode := range []string{"replace-reservation", "replace-state", "replace-change", "lost-response"} {
		t.Run(mode, func(t *testing.T) {
			f := startFixture(t, "sha1", "spec", "")
			winner, raw := startWinner(t, f)
			cfg := startProxyConfig{Mode: mode, WinnerObject: winner}
			if mode == "replace-change" {
				cfg.WinnerObject = ""
				cfg.Change = filepath.Join(f.Root, f.Path)
				cfg.WinnerText = raw
			}
			cfg = startProxyInstall(t, binary, cfg)
			var err error
			_, err = claimCommand(t, f.Root, "", "update-ref", "refs/fireweed/claims/current", winner, strings.Repeat("0", 40))
			if (err != nil) != (mode == "lost-response") {
				t.Error("proxy response control", mode, err)
			}
			startProxyFired(t, cfg)
			switch mode {
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
			checkStorage := startStorage(t, f.Root, "current")
			defer checkStorage()
			tree := startTree(t, f.Root)
			git := startGitSnapshot(t, f.Root)
			got, err := w.Start(f.Root, "current")
			startProxyFired(t, cfg)
			if err == nil || !policyCorrective(err.Error()) || !reflect.DeepEqual(got, w.StartResult{}) {
				t.Errorf("ownership/failure did not reject with zero result: %+v %v", got, err)
			}
			if mode != "lost-response" && (err == nil || !startRecoveryDetail(err.Error())) {
				t.Error("lost ownership lacks recovery details", err)
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
			if mode == "move-base" {
				f, cfg = startPinnedFixture(t)
				old = cfg.BaseOld
			}
			if mode == "replacement-base" {
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
			if err := os.Chmod(filepath.Join(f.Root, f.Path), f.Mode); err != nil {
				t.Fatal(err)
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

func startRecoveryDetail(detail string) bool {
	text := strings.ToLower(detail)
	for _, term := range []string{"recover", "restore", "ownership", "preserv", "replacement", "replaced", "concurrent"} {
		if strings.Contains(text, term) {
			return true
		}
	}
	return false
}

func startReferenceHook(t *testing.T, root string) string {
	t.Helper()
	marker := filepath.Join(t.TempDir(), "reference-transaction-called")
	hook := filepath.Join(claimCommon(t, root), "hooks", "reference-transaction")
	worktreePut(t, filepath.Dir(hook), filepath.Base(hook), "#!/bin/sh\nprintf '%s\\n' \"$1\" >> "+claimQuote(marker)+"\ncat >/dev/null\n")
	if err := os.Chmod(hook, 0755); err != nil {
		t.Fatal(err)
	}
	claimGit(t, root, "update-ref", "refs/start-hook-control", "HEAD")
	raw, err := os.ReadFile(marker)
	if err != nil || !strings.Contains(string(raw), "prepared") || !strings.Contains(string(raw), "committed") {
		t.Fatal("reference-transaction hook not live", string(raw), err)
	}
	claimGit(t, root, "update-ref", "-d", "refs/start-hook-control")
	if err := os.Remove(marker); err != nil {
		t.Fatal(err)
	}
	disabled := t.TempDir()
	claimGit(t, root, "-c", "core.hooksPath="+disabled, "update-ref", "refs/start-hook-control", "HEAD")
	claimGit(t, root, "-c", "core.hooksPath="+disabled, "update-ref", "-d", "refs/start-hook-control")
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Fatal("reference hook suppressed control failed", err)
	}
	return marker
}
func Test_FLOW040StartReferenceTransactionHookControl(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("POSIX executable hook fixture; portable ref protocol controls run separately")
	}
	f := startFixture(t, "sha1", "spec", "")
	startReferenceHook(t, f.Root)
}
func Test_FLOW037_FLOW038_FLOW040StartSuppressesReferenceTransactionHook(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("POSIX executable hook fixture; portable ref protocol controls run separately")
	}
	for _, failed := range []bool{false, true} {
		t.Run(map[bool]string{false: "success", true: "rollback"}[failed], func(t *testing.T) {
			f := startFixture(t, "sha1", "spec", "")
			if failed {
				startBrokenBook(t, f, "evals")
			}
			marker := startReferenceHook(t, f.Root)
			if failed {
				startFailure(t, f.Root, "current", "")
			} else {
				startSuccess(t, f, f.Root, "evals")
			}
			if raw, err := os.ReadFile(marker); err == nil {
				t.Error("Start invoked proven reference-transaction hook", string(raw))
			} else if !os.IsNotExist(err) {
				t.Error(err)
			}
		})
	}
}

// The Git proxy deliberately outlives the outer ceiling. Only the subject can
// return an ordinary error before that ceiling; the test kills any leaked proxy.
func Test_FLOW040StartBoundsStalledGit(t *testing.T) {
	binary := startProxyBuild(t)
	build := t.TempDir()
	probe := filepath.Join(build, "start-probe")
	if runtime.GOOS == "windows" {
		probe += ".exe"
	}
	worktreePut(t, build, "main.go", `package main
import("encoding/json";"os";w "github.com/lurkshark/fireweed/packages/workflow")
func main(){result,err:=w.Start(os.Args[1],"current");detail:="";if err!=nil{detail=err.Error()};json.NewEncoder(os.Stdout).Encode(struct{Result w.StartResult;Error string}{result,detail})}`)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "build", "-o", probe, filepath.Join(build, "main.go"))
	cmd.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatal("public stalled-Git probe build", err, string(out))
	}
	for _, control := range []bool{true, false} {
		t.Run(map[bool]string{true: "live-control", false: "subject"}[control], func(t *testing.T) {
			f := startFixture(t, "sha1", "spec", "")
			cfg := startProxyInstall(t, binary, startProxyConfig{Mode: "hang"})
			checkStorage := startStorage(t, f.Root, "current")
			defer checkStorage()
			// Snapshot helpers use real Git: the stall is armed only in the child env.
			before := startTree(t, f.Root)
			git := startGitSnapshot(t, f.Root)
			ceiling := 90 * time.Second
			if control {
				ceiling = 350 * time.Millisecond
			}
			ctx, cancel := context.WithTimeout(context.Background(), ceiling)
			defer cancel()
			subject := exec.CommandContext(ctx, probe, f.Root)
			if control {
				subject = exec.CommandContext(ctx, binary, "--version")
			}
			subject.Dir = f.Root
			subject.Env = append(os.Environ(), "START_PROXY_ARM_HANG=1")
			defer func() {
				raw, e := os.ReadFile(cfg.Marker + "-pid")
				if e == nil {
					var pid int
					if json.Unmarshal(raw, &pid) == nil {
						process, e := os.FindProcess(pid)
						if e == nil {
							_ = process.Kill()
						}
					}
				}
			}()
			out, err := subject.CombinedOutput()
			startProxyFired(t, cfg)
			if control {
				if ctx.Err() != context.DeadlineExceeded || err == nil {
					t.Error("proxy did not actually stall until external cancellation", err)
				}
				return
			}
			if ctx.Err() != nil || err != nil {
				t.Fatalf("Start failed to return before generous external completion ceiling: %v %s", err, out)
			}
			var got struct {
				Result w.StartResult
				Error  string
			}
			if err := json.Unmarshal(out, &got); err != nil {
				t.Fatal("stalled subject JSON", err, string(out))
			}
			if !reflect.DeepEqual(got.Result, w.StartResult{}) || got.Error == "" || !policyCorrective(got.Error) {
				t.Error("stalled Git did not produce useful zero-result error", got)
			}
			detail := strings.ToLower(got.Error)
			if !strings.Contains(detail, "deadline") && !strings.Contains(detail, "timeout") && !strings.Contains(detail, "timed out") && !strings.Contains(detail, "cancel") && !strings.Contains(detail, "killed") && !strings.Contains(detail, "signal") {
				t.Error("missing stalled-command failure detail", got.Error)
			}
			if !reflect.DeepEqual(before, startTree(t, f.Root)) || !reflect.DeepEqual(git, startGitSnapshot(t, f.Root)) {
				t.Error("stalled subprocess changed repository state")
			}
		})
	}
}

func startPinnedFixture(t *testing.T) (startFixtureData, startProxyConfig) {
	t.Helper()
	f := startFixture(t, "sha1", "spec", "")
	baseText := strings.ReplaceAll(f.Raw, "\n", "\r\n")
	worktreePut(t, f.Root, f.Path, baseText)
	startCommit(t, f.Root, "Approved CRLF source")
	old := claimBase(t, f.Root)
	claimGit(t, f.Root, "update-ref", "refs/heads/approval-base", old)
	worktreePut(t, f.Root, f.Path, f.Raw)
	startCommit(t, f.Root, "Equivalent current LF source")
	current := claimBase(t, f.Root)
	blob := claimGit(t, f.Root, "rev-parse", old+":"+f.Path)
	if blob == claimGit(t, f.Root, "rev-parse", current+":"+f.Path) {
		t.Fatal("base and current selected blobs are not unique")
	}
	same, err := w.ApprovalMatches(c.Module{Name: "alpha", IDPrefix: "ALP"}, baseText, f.Raw, f.Path)
	if err != nil || !same {
		t.Fatal("LF/CRLF author approval premise", same, err)
	}
	worktreePut(t, f.Root, f.Path, strings.Replace(f.Raw, "Start the exact authored change.", "Unapproved replacement base content.", 1))
	startCommit(t, f.Root, "Unapproved alternate base")
	alternate := claimBase(t, f.Root)
	claimGit(t, f.Root, "reset", "--hard", current)
	if err := os.Chmod(filepath.Join(f.Root, f.Path), f.Mode); err != nil {
		t.Fatal(err)
	}
	return f, startProxyConfig{Mode: "move-base", BaseRef: "refs/heads/approval-base", BaseOld: old, BaseNew: alternate, Blob: blob, BaseText: baseText}
}
func startProxyOutput(t *testing.T, root, input string, args ...string) []byte {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = root
	cmd.Stdin = strings.NewReader(input)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		t.Fatal("observer Git control", err, stderr.String())
	}
	return out
}
func Test_FLOW036_FLOW040StartPinnedBaseObserverControls(t *testing.T) {
	binary := startProxyBuild(t)
	for _, shape := range []string{"show", "blob-oid", "ls-tree", "cat-file-blob", "cat-file-batch"} {
		t.Run(shape, func(t *testing.T) {
			f, cfg := startPinnedFixture(t)
			args := []string{"show", cfg.BaseOld + ":" + f.Path}
			input := ""
			switch shape {
			case "blob-oid":
				args = []string{"rev-parse", cfg.BaseOld + ":" + f.Path}
			case "ls-tree":
				args = []string{"ls-tree", cfg.BaseOld, "--", f.Path}
			case "cat-file-blob":
				args = []string{"cat-file", "blob", cfg.Blob}
			case "cat-file-batch":
				args = []string{"cat-file", "--batch"}
				input = cfg.Blob + "\n"
			}
			expected := startProxyOutput(t, f.Root, input, args...)
			cfg = startProxyInstall(t, binary, cfg)
			if _, err := p.Load(f.Root); err != nil {
				t.Fatal("current project control", err)
			}
			if _, err := p.LoadAt(f.Root, "HEAD"); err != nil {
				t.Fatal("current committed project control", err)
			}
			if _, err := os.Stat(cfg.Marker); !os.IsNotExist(err) {
				t.Fatal("current-project load triggered base event", err)
			}
			if tip := claimGit(t, f.Root, "rev-parse", cfg.BaseRef); tip != cfg.BaseOld {
				t.Fatal("base moved before base-source observation", tip)
			}
			got := startProxyOutput(t, f.Root, input, args...)
			if !bytes.Equal(got, expected) {
				t.Error("observer changed successful stdout", string(got), string(expected))
			}
			startProxyFired(t, cfg)
			if tip := claimGit(t, f.Root, "rev-parse", cfg.BaseRef); tip != cfg.BaseNew {
				t.Error("event did not move base", tip)
			}
			if shape == "blob-oid" {
				content := startProxyOutput(t, f.Root, "", "cat-file", "blob", strings.TrimSpace(string(got)))
				if string(content) != cfg.BaseText {
					t.Error("resolved pinned blob changed")
				}
			}
		})
	}
}
func Test_FLOW032_FLOW035StartPreconditionProxyControl(t *testing.T) {
	binary := startProxyBuild(t)
	f := startFixture(t, "sha1", "spec", "")
	cfg := startProxyInstall(t, binary, startProxyConfig{Mode: "precondition"})
	out, err := claimCommand(t, f.Root, "", "update-ref", "refs/fireweed/claims/current", "HEAD")
	if err == nil || !strings.Contains(out, "injected prohibited precondition claim mutation") {
		t.Error("precondition Git mutation guard not live", err)
	}
	if _, err := os.Stat(cfg.Marker + "-precondition"); err != nil {
		t.Error("precondition mutation control unobserved", err)
	}
}
func startPreconditionProxy(t *testing.T, binary string, f startFixtureData) func() {
	t.Helper()
	cfg := startProxyInstall(t, binary, startProxyConfig{Mode: "precondition"})
	return func() {
		t.Helper()
		if raw, err := os.ReadFile(cfg.Marker + "-precondition"); err == nil {
			t.Error("claim mutation attempted before rejecting precondition", string(raw))
		} else if !os.IsNotExist(err) {
			t.Error(err)
		}
	}
}

func Test_FLOW036_FLOW040StartLocalObjectInventoryControl(t *testing.T) {
	binary := startProxyBuild(t)
	f, cfg := startMissingPromisorFixture(t, binary)
	cmd := exec.Command(cfg.Real, "cat-file", "--batch-all-objects", "--batch-check=%(objectname)")
	cmd.Dir = f.Root
	cmd.Env = append(os.Environ(), "GIT_NO_LAZY_FETCH=1", "GIT_NO_REPLACE_OBJECTS=1")
	out, e := cmd.CombinedOutput()
	if e != nil || len(out) == 0 {
		t.Fatal("local object inventory is unavailable", e, string(out))
	}
	if b, e := os.ReadFile(cfg.Marker + "-transport"); e == nil {
		t.Fatal("local object inventory attempted transport", string(b))
	} else if !os.IsNotExist(e) {
		t.Fatal(e)
	}
}
