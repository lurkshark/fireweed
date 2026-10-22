package evals_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	p "github.com/lurkshark/fireweed/packages/project"
)

type baselinePolicy struct {
	Real, Marker     string
	Source, RootTree string
	NoTree           bool
	Allowed          map[string]bool
	Paths            map[string]bool
	Fail             string
	FailBodies       map[string]bool
}

func baselineGuard(t *testing.T, f baselineFixtureData, fail string, extra ...string) (string, string) {
	t.Helper()
	real, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	temp := baselineCanonical(t, t.TempDir())
	marker := filepath.Join(temp, "forbidden")
	policy := baselinePolicy{Real: real, Marker: marker, Allowed: map[string]bool{f.Commit: true}, Paths: map[string]bool{}, Fail: fail, FailBodies: map[string]bool{}}
	policy.Source = f.Commit
	policy.NoTree = fail == "before-tree"
	if policy.NoTree {
		policy.Fail = ""
	}
	failedPath := ""
	if strings.HasPrefix(fail, "body:") {
		failedPath = strings.TrimPrefix(fail, "body:")
		policy.Fail = ""
		policy.FailBodies[failedPath] = true
		policy.FailBodies[f.Commit+":"+failedPath] = true
	}
	paths := []string{"fireweed.yaml"}
	for _, m := range f.Expected.Modules {
		paths = append(paths, filepath.ToSlash(filepath.Join(m.Dir, "SPEC.md")))
	}
	for _, list := range [][]string{extra} {
		paths = append(paths, list...)
	}
	for _, ch := range f.Expected.Changes {
		paths = append(paths, ch.Path)
	}
	for _, ch := range f.Expected.Archive {
		paths = append(paths, ch.Path)
	}
	for _, path := range paths {
		policy.Paths[path] = true
	}
	raw := baselineGit(t, f.Root, "ls-tree", "-r", "-t", "-z", "--full-tree", f.Commit)
	for _, record := range strings.Split(raw, "\x00") {
		if record == "" {
			continue
		}
		pieces := strings.SplitN(record, "\t", 2)
		if len(pieces) != 2 {
			t.Fatal("tree fixture record")
		}
		header := strings.Fields(pieces[0])
		if len(header) != 3 {
			t.Fatal("tree fixture header")
		}
		if pieces[1] == failedPath {
			policy.FailBodies[header[2]] = true
		}
		if header[1] == "tree" || policy.Paths[pieces[1]] {
			policy.Allowed[header[2]] = true
		}
	}
	policy.RootTree = baselineGit(t, f.Root, "rev-parse", f.Commit+"^{tree}")
	policy.Allowed[policy.RootTree] = true
	input := filepath.Join(temp, "policy.json")
	bytes, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	if err = os.WriteFile(input, bytes, 0644); err != nil {
		t.Fatal(err)
	}
	program := `package main
import("bufio";"encoding/json";"io";"os";"os/exec";"strings")
type policy struct{Real,Marker,Source,RootTree string;NoTree bool;Allowed,Paths,FailBodies map[string]bool;Fail string}
var p policy
func bad(reason string){os.WriteFile(p.Marker,[]byte(reason),0600);os.Exit(97)}
func allowed(query string)bool{
 q:=strings.TrimSuffix(strings.TrimSuffix(query,"\n"),"\x00");q=strings.TrimSpace(q)
 if p.FailBodies[q]{os.Exit(73)};if idx:=strings.Index(q,":");idx>=0&&p.FailBodies[q[idx+1:]]{os.Exit(73)};fields:=strings.Fields(q);if len(fields)>0&&p.FailBodies[fields[0]]{os.Exit(73)}
 if p.Allowed[q]{return true};if q==p.Source+"^{tree}"||q==p.Source+"^{commit}"{return true}
 if idx:=strings.Index(q,":");idx>=0&&p.Paths[q[idx+1:]]&&(q[:idx]==p.Source||q[:idx]==p.RootTree){return true}
 first:=strings.Fields(q);if len(first)>0&&p.Allowed[strings.TrimSuffix(first[0],"^{blob}")]{return true};return false
}
func main(){raw,e:=os.ReadFile(os.Getenv("FW_BASELINE_GIT_POLICY"));if e!=nil{os.Exit(80)};if json.Unmarshal(raw,&p)!=nil{os.Exit(81)};args:=os.Args[1:];command:="";at:=0
 for i:=0;i<len(args);i++{a:=args[i];if a=="-c"||a=="-C"||a=="--git-dir"||a=="--work-tree"{i++;continue};if strings.HasPrefix(a,"-"){continue};command=a;at=i;break}
 tail:=args[at+1:];batch:=false;batchCommand:=false;metadataOnly:=false
 if p.NoTree&&command=="ls-tree"{bad("tree read before valid explicit revision: "+command)}
 switch command{case "fetch","push","clone","ls-remote","commit","branch","checkout","switch","reset","restore","add","worktree","update-ref","hash-object","archive","grep","diff","status":bad("forbidden command "+command)}
 if p.Fail==command{os.Exit(73)}
 if command=="cat-file"{for _,a:=range tail{if a=="-t"||a=="-s"||a=="-e"||strings.HasPrefix(a,"--batch-check"){metadataOnly=true};if strings.HasPrefix(a,"--batch-command"){batch=true;batchCommand=true};if a=="--batch"||strings.HasPrefix(a,"--batch="){batch=true}};if p.NoTree&&!metadataOnly&&!batch{bad("body read before valid explicit revision")};if !metadataOnly&&!batch{queries:=[]string{};for _,a:=range tail{if a=="blob"||a=="tree"||a=="commit"||strings.HasPrefix(a,"-"){continue};queries=append(queries,a)};for _,q:=range queries{if !allowed(q){bad("unrelated body "+q)}}}}
 if command=="show"||command=="log"{noPatch:=false;query:=false;for _,a:=range tail{if a=="--no-patch"||a=="-s"{noPatch=true};if strings.Contains(a,":")&&allowed(a){query=true}};if p.NoTree&&!noPatch{bad("historical content before valid explicit revision")};if !noPatch&&!query{bad("unbounded historical content "+command)}}
 child:=exec.Command(p.Real,args...);child.Stdout=os.Stdout;child.Stderr=os.Stderr;child.Env=os.Environ()
 if batch&&!metadataOnly{pipe,e:=child.StdinPipe();if e!=nil{os.Exit(82)};if child.Start()!=nil{os.Exit(83)};reader:=bufio.NewReader(os.Stdin);delimiter:=byte('\n');for _,a:=range tail{if a=="-z"||a=="-Z"{delimiter=0}}
  for{record,e:=reader.ReadString(delimiter);if len(record)>0{q:=strings.TrimSuffix(strings.TrimSuffix(record,"\n"),"\x00");check:=true;if batchCommand{if strings.HasPrefix(q,"info ")||q=="flush"{check=false}else{q=strings.TrimPrefix(q,"contents ")}};if check&&p.NoTree{child.Process.Kill();pipe.Close();child.Wait();bad("batch body before valid explicit revision")};if check&&!allowed(q){child.Process.Kill();pipe.Close();child.Wait();bad("unrelated batch body "+q)};if _,err:=io.WriteString(pipe,record);err!=nil{break}};if e!=nil{break}};pipe.Close()
 }else{child.Stdin=os.Stdin;if child.Start()!=nil{os.Exit(84)}}
 if err:=child.Wait();err!=nil{if e,ok:=err.(*exec.ExitError);ok{os.Exit(e.ExitCode())};os.Exit(85)}
}`
	source := filepath.Join(temp, "guard.go")
	if err = os.WriteFile(source, []byte(program), 0644); err != nil {
		t.Fatal(err)
	}
	binary := filepath.Join(temp, "git")
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "build", "-o", binary, source)
	cmd.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Git body guard build %v %s", err, out)
	}
	t.Setenv("FW_BASELINE_GIT_POLICY", input)
	t.Setenv("PATH", temp+string(os.PathListSeparator)+os.Getenv("PATH"))
	return marker, real
}
func baselineNoMarker(t *testing.T, marker string) {
	t.Helper()
	if raw, err := os.ReadFile(marker); !os.IsNotExist(err) {
		t.Fatalf("unexpected forbidden operation: %v %s", err, raw)
	}
}
func Test_PRJ027HistoricalBodyGuardControls(t *testing.T) {
	f := baselineFixture(t, "sha1")
	marker, _ := baselineGuard(t, f, "")
	allowed := baselineGit(t, f.Root, "rev-parse", f.Commit+":fireweed.yaml")
	denied := baselineGit(t, f.Root, "rev-parse", f.Commit+":private/secret.txt")
	if allowed == denied {
		t.Fatal("distinct content control premise")
	}
	for _, query := range []string{denied, f.Commit + ":private/secret.txt"} {
		if _, err := baselineInput(t, f.Root, "", "cat-file", "blob", query); err == nil {
			t.Fatal("unrelated direct body guard control passed")
		}
		if _, err := os.Stat(marker); err != nil {
			t.Fatal("missing ignored-error body marker")
		}
		if err := os.Remove(marker); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := baselineInput(t, f.Root, denied+"\n", "cat-file", "--batch"); err == nil {
		t.Fatal("batch body guard control passed")
	}
	if _, err := os.Stat(marker); err != nil {
		t.Fatal("missing batch marker")
	}
	_ = os.Remove(marker)
	raw, err := baselineInput(t, f.Root, "", "cat-file", "blob", allowed)
	if err != nil || raw != f.Files["fireweed.yaml"] {
		t.Fatalf("allowed object control %v %q", err, raw)
	}
	raw, err = baselineInput(t, f.Root, allowed+"\n", "cat-file", "--batch")
	if err != nil || !strings.Contains(raw, f.Files["fireweed.yaml"]) {
		t.Fatalf("allowed batch control %v", err)
	}
	baselineNoMarker(t, marker)
}
func Test_PRJ025_PRJ027OnlySelectedHistoricalBodies(t *testing.T) {
	for _, mode := range []string{"normal", "diagnosed", "config-error", "contained-link", "old-revision", "hidden-tree", "gitlink-module"} {
		t.Run(mode, func(t *testing.T) {
			f := baselineFixture(t, "sha1")
			extra := []string{}
			switch mode {
			case "diagnosed":
				put(t, f.Root, "outside/SPEC.md", "private out-of-glob source must not be read")
				baselineFresh(t, &f)
			case "config-error":
				put(t, f.Root, "fireweed.yaml", "invalid configuration: [")
				f.Commit = baselineCommit(t, f.Root)
			case "contained-link":
				put(t, f.Root, "metadata/config.yaml", f.Files["fireweed.yaml"])
				if err := os.Remove(filepath.Join(f.Root, "fireweed.yaml")); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink("metadata/config.yaml", filepath.Join(f.Root, "fireweed.yaml")); err != nil {
					t.Fatal(err)
				}
				baselineFresh(t, &f)
				extra = append(extra, "metadata/config.yaml")
			case "hidden-tree":
				baselineHiddenTree(t, &f, "root-dotgit")
			case "gitlink-module":
				baselineHiddenTree(t, &f, "gitlink-module")
			case "old-revision":
				put(t, f.Root, "fireweed.yaml", "invalid newer configuration: [")
				baselineCommit(t, f.Root)
			}
			policyFixture := f
			if mode == "config-error" {
				policyFixture.Expected.Modules = nil
				policyFixture.Expected.Changes = nil
				policyFixture.Expected.Archive = nil
			}
			marker, _ := baselineGuard(t, policyFixture, "", extra...)
			if mode == "config-error" {
				baselineError(t, f.Root, f.Commit, "fireweed.yaml")
			} else {
				baselineAssert(t, f, f.Root, f.Commit)
			}
			baselineNoMarker(t, marker)
		})
	}
}
func Test_PRJ024EnumerationAndObjectReadErrors(t *testing.T) {
	for _, command := range []string{"rev-parse", "ls-tree", "cat-file"} {
		t.Run(command, func(t *testing.T) {
			f := baselineFixture(t, "sha1")
			baselineGuard(t, f, command)
			baselineError(t, f.Root, f.Commit, "")
		})
	}
}
func Test_PRJ024ResolvedCommitSurvivesMovingRef(t *testing.T) {
	if os.PathSeparator == '\\' {
		t.Skip("POSIX local ref movement fixture")
	}
	f := baselineFixture(t, "sha1")
	baselineGit(t, f.Root, "branch", "moving", f.Commit)
	put(t, f.Root, "fireweed.yaml", "invalid later version: [")
	newer := baselineCommit(t, f.Root)
	real, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	temp := t.TempDir()
	marker := filepath.Join(temp, "moved")
	script := "#!/bin/sh\nresolve=no\ntree=no\nfor a in \"$@\"; do case \"$a\" in moving|moving^\\{commit\\}|refs/heads/moving|refs/heads/moving^\\{commit\\}) resolve=yes;;ls-tree) tree=yes;;esac;done\nif [ ! -e " + shellQuote(marker) + " ] && [ \"$resolve\" = yes ]; then\n" + shellQuote(real) + " \"$@\" > " + shellQuote(filepath.Join(temp, "resolved")) + "\ncode=$?\n" + shellQuote(real) + " update-ref refs/heads/moving " + shellQuote(newer) + "\nprintf moved > " + shellQuote(marker) + "\ncat " + shellQuote(filepath.Join(temp, "resolved")) + "\nexit $code\nfi\nif [ ! -e " + shellQuote(marker) + " ] && [ \"$tree\" = yes ]; then\n" + shellQuote(real) + " update-ref refs/heads/moving " + shellQuote(newer) + "\nprintf moved > " + shellQuote(marker) + "\nfi\nexec " + shellQuote(real) + " \"$@\"\n"
	put(t, temp, "git", script)
	if err = os.Chmod(filepath.Join(temp, "git"), 0755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", temp+string(os.PathListSeparator)+os.Getenv("PATH"))
	control := baselineGit(t, f.Root, "rev-parse", "--verify", "moving^{commit}")
	if control != f.Commit {
		t.Fatalf("pinning control did not return the old commit: %s", control)
	}
	if _, err = os.Stat(marker); err != nil {
		t.Fatal("ref movement control did not activate")
	}
	check := exec.Command(real, "rev-parse", "moving")
	check.Dir = f.Root
	if out, err := check.Output(); err != nil || strings.TrimSpace(string(out)) != newer {
		t.Fatalf("moved ref control %v %s", err, out)
	}
	reset := exec.Command(real, "update-ref", "refs/heads/moving", f.Commit)
	reset.Dir = f.Root
	if out, err := reset.CombinedOutput(); err != nil {
		t.Fatalf("reset fixture ref %v %s", err, out)
	}
	if err = os.Remove(marker); err != nil {
		t.Fatal(err)
	}
	baselineAssert(t, f, f.Root, "moving")
	if _, err = os.Stat(marker); err != nil {
		t.Fatal("ref-movement injection did not execute")
	}
	if actual := baselineGit(t, f.Root, "rev-parse", "moving"); actual != newer {
		t.Fatal("moving ref premise")
	}
	// A later call resolves the newly moved name afresh; no cache by mutable name.
	baselineError(t, f.Root, "moving", "fireweed.yaml")
}
func Test_PRJ024ReplacementRefsCannotSubstituteCommit(t *testing.T) {
	f := baselineFixture(t, "sha1")
	put(t, f.Root, "fireweed.yaml", "invalid replacement source: [")
	newer := baselineCommit(t, f.Root)
	baselineGit(t, f.Root, "replace", f.Commit, newer)
	raw, err := baselineInput(t, f.Root, "", "show", f.Commit+":fireweed.yaml")
	if err != nil || raw != "invalid replacement source: [" {
		t.Fatalf("active replacement fixture %v %q", err, raw)
	}
	baselineAssert(t, f, f.Root, f.Commit)
}
func baselineNativeHook(t *testing.T, root string) string {
	t.Helper()
	common := baselineGit(t, root, "rev-parse", "--git-common-dir")
	if !filepath.IsAbs(common) {
		common = filepath.Join(root, common)
	}
	hook := filepath.Join(common, "hooks", "baseline-fsmonitor")
	marker := filepath.Join(t.TempDir(), "hook")
	put(t, filepath.Dir(hook), filepath.Base(hook), "#!/bin/sh\nprintf called >> "+shellQuote(marker)+"\nprintf 'baseline-token\\000/\\000'\n")
	if err := os.Chmod(hook, 0755); err != nil {
		t.Fatal(err)
	}
	baselineGit(t, root, "config", "core.fsmonitor", shellQuote(hook))
	baselineGit(t, root, "config", "core.fsmonitorHookVersion", "2")
	baselineGit(t, root, "status", "--porcelain=v1")
	if raw, err := os.ReadFile(marker); err != nil || len(raw) == 0 {
		t.Fatalf("native hook control %v %q", err, raw)
	}
	if err := os.Remove(marker); err != nil {
		t.Fatal(err)
	}
	baselineGit(t, root, "-c", "core.fsmonitor=false", "status", "--porcelain=v1")
	baselineNoMarker(t, marker)
	return marker
}
func Test_PRJ027NativeHookAndStorageControls(t *testing.T) {
	if os.PathSeparator == '\\' {
		t.Skip("POSIX hook fixture")
	}
	baselineNativeHook(t, baselineFixture(t, "sha1").Root)
}
func Test_PRJ027CompleteReadOnlyState(t *testing.T) {
	if os.PathSeparator == '\\' {
		t.Skip("POSIX command/hook traps")
	}
	for _, linked := range []bool{false, true} {
		for _, mode := range []string{"success", "dirty", "invalid-revision", "config-error", "diagnosed"} {
			t.Run(fmt.Sprint(linked)+"/"+mode, func(t *testing.T) {
				f := baselineFixture(t, "sha1")
				if mode == "config-error" {
					put(t, f.Root, "fireweed.yaml", "invalid selected config: [")
					f.Commit = baselineCommit(t, f.Root)
				}
				if mode == "diagnosed" {
					put(t, f.Root, "units/bad/SPEC.md", module("invalid name", "BAD", ""))
					baselineFresh(t, &f)
				}
				if linked {
					peer := filepath.Join(t.TempDir(), "linked ' peer")
					baselineGit(t, f.Root, "worktree", "add", "-b", "peer", peer)
					f.Root = baselineCanonical(t, peer)
					f.Expected.Root = f.Root
				}
				if mode == "dirty" {
					put(t, f.Root, "fireweed.yaml", "invalid current config: [")
					put(t, f.Root, "untracked", "retain untracked")
					put(t, f.Root, "ignored", "retain ignored")
				}
				baselineGit(t, f.Root, "update-ref", "refs/fireweed/claims/seeded", f.Commit)
				baselineGit(t, f.Root, "update-ref", "refs/fireweed/drills/retained", f.Commit)
				hook := baselineNativeHook(t, f.Root)
				marker := playbookTraps(t)
				trapDir := filepath.Dir(marker)
				for _, name := range []string{"baseline-hook-trap", "baseline-agent-trap"} {
					put(t, trapDir, name, "#!/bin/sh\nprintf attempted >> "+shellQuote(marker)+"\nexit 97\n")
					if err := os.Chmod(filepath.Join(trapDir, name), 0755); err != nil {
						t.Fatal(err)
					}
				}
				temporary := baselineCanonical(t, t.TempDir())
				t.Setenv("TMPDIR", temporary)
				t.Setenv("TMP", temporary)
				t.Setenv("TEMP", temporary)
				common := baselineGit(t, f.Root, "rev-parse", "--git-common-dir")
				if !filepath.IsAbs(common) {
					common = filepath.Join(f.Root, common)
				}
				common = baselineCanonical(t, common)
				files := playbookTree(t, f.Root)
				metadata := playbookTree(t, common)
				temps := playbookTree(t, temporary)
				env := os.Environ()
				cwd, err := os.Getwd()
				if err != nil {
					t.Fatal(err)
				}
				if mode == "invalid-revision" {
					baselineError(t, f.Root, "missing-ref", "")
				} else if mode == "config-error" {
					baselineError(t, f.Root, f.Commit, "fireweed.yaml")
				} else {
					got, err := baselineCall(t, f.Root, f.Commit)
					if err != nil {
						t.Fatal(err)
					}
					if mode == "diagnosed" {
						baselineCategories(t, got, f.Expected)
					} else if !reflect.DeepEqual(got, f.Expected) {
						t.Fatal("baseline data changed")
					}
				}
				if !reflect.DeepEqual(files, playbookTree(t, f.Root)) || !reflect.DeepEqual(metadata, playbookTree(t, common)) || !reflect.DeepEqual(temps, playbookTree(t, temporary)) {
					t.Fatal("snapshot read mutated files/common Git/temp")
				}
				after, err := os.Getwd()
				if err != nil || after != cwd || !reflect.DeepEqual(env, os.Environ()) {
					t.Fatal("snapshot changed parent state")
				}
				baselineNoMarker(t, marker)
				baselineNoMarker(t, hook)
			})
		}
	}
}

func Test_PRJ025_PRJ027SelectedBodyFailureControl(t *testing.T) {
	f := baselineFixture(t, "sha1")
	marker, _ := baselineGuard(t, f, "body:units/a/SPEC.md")
	id := baselineGit(t, f.Root, "rev-parse", f.Commit+":units/a/SPEC.md")
	if _, err := baselineInput(t, f.Root, "", "cat-file", "blob", id); err == nil {
		t.Fatal("selected object failure did not activate")
	}
	baselineNoMarker(t, marker)
	raw, err := baselineInput(t, f.Root, "", "cat-file", "blob", f.Commit+":fireweed.yaml")
	if err != nil || raw != f.Files["fireweed.yaml"] {
		t.Fatal("selected object failure affected config control")
	}
}
func Test_PRJ025FileScopedObjectFailuresRemainDiagnosed(t *testing.T) {
	for _, target := range []string{"units/a/SPEC.md", "changes/0002-alpha.md", "changes/archive/0009-old.md"} {
		t.Run(target, func(t *testing.T) {
			f := baselineFixture(t, "sha1")
			marker, _ := baselineGuard(t, f, "body:"+target)
			want := f.Expected
			want.Diagnostics = nil
			switch target {
			case "units/a/SPEC.md":
				want.Modules = want.Modules[1:]
				want.Changes = want.Changes[:2]
				want.Diagnostics = []p.Diagnostic{{Code: "C03", Severity: "error", Path: target}, {Code: "C05", Severity: "error", Path: "changes/0000010-last.md"}}
			case "changes/0002-alpha.md":
				want.Changes = want.Changes[1:]
				want.Diagnostics = []p.Diagnostic{{Code: "C05", Severity: "error", Path: target}}
			default:
				want.Archive = want.Archive[:0]
				want.Diagnostics = []p.Diagnostic{{Code: "C05", Severity: "error", Path: target}}
			}
			got, err := baselineCall(t, f.Root, f.Commit)
			if err != nil {
				t.Fatalf("file-scoped object failure became fatal: %v", err)
			}
			baselineCategories(t, got, want)
			baselineNoMarker(t, marker)
		})
	}
}

func Test_PRJ024_PRJ025GitWarningsAreNotMetadata(t *testing.T) {
	if os.PathSeparator == '\\' {
		t.Skip("POSIX stderr warning fixture")
	}
	f := baselineFixture(t, "sha1")
	oid := baselineGit(t, f.Root, "rev-parse", f.Commit+":fireweed.yaml")
	real, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	temp := t.TempDir()
	put(t, temp, "git", "#!/bin/sh\nprintf 'local Git warning\\n' >&2\nexec "+shellQuote(real)+" \"$@\"\n")
	if err := os.Chmod(filepath.Join(temp, "git"), 0755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", temp+string(os.PathListSeparator)+os.Getenv("PATH"))
	// Separate direct process streams prove the warning is real and stdout exact.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", "cat-file", "blob", oid)
	cmd.Dir = f.Root
	var warning strings.Builder
	cmd.Stderr = &warning
	raw, err := cmd.Output()
	if err != nil || string(raw) != f.Files["fireweed.yaml"] || warning.String() != "local Git warning\n" {
		t.Fatalf("warning control %v stdout %q stderr %q", err, raw, warning.String())
	}
	baselineAssert(t, f, f.Root, f.Commit)
}

func Test_PRJ024_PRJ027PinnedBodyWhitelistControls(t *testing.T) {
	f := baselineFixture(t, "sha1")
	selected := f.Commit
	put(t, f.Root, "newer-only", "different commit with identical selected config")
	newer := baselineCommit(t, f.Root)
	marker, _ := baselineGuard(t, f, "")
	for _, q := range []string{"HEAD:fireweed.yaml", newer + ":fireweed.yaml", "refs/heads/main:fireweed.yaml"} {
		for _, batch := range []bool{false, true} {
			var err error
			if batch {
				_, err = baselineInput(t, f.Root, q+"\n", "cat-file", "--batch")
			} else {
				_, err = baselineInput(t, f.Root, "", "cat-file", "blob", q)
			}
			if err == nil {
				t.Fatalf("unpinned query control passed %q batch %v", q, batch)
			}
			if _, err := os.Stat(marker); err != nil {
				t.Fatal("missing unselected revision marker", err)
			}
			if err := os.Remove(marker); err != nil {
				t.Fatal(err)
			}
		}
	}
	raw, err := baselineInput(t, f.Root, "", "cat-file", "blob", selected+":fireweed.yaml")
	if err != nil || raw != f.Files["fireweed.yaml"] {
		t.Fatalf("selected revision control %v %q", err, raw)
	}
	baselineNoMarker(t, marker)
	// Identical config bytes make ignored reads from a newer ref distinguishable only by the guard.
	baselineAssert(t, f, f.Root, selected)
	baselineNoMarker(t, marker)
}
func Test_PRJ024_PRJ027RejectInvalidRevisionBeforeTree(t *testing.T) {
	f := baselineFixture(t, "sha1")
	blob := baselineGit(t, f.Root, "rev-parse", f.Commit+":fireweed.yaml")
	tree := baselineGit(t, f.Root, "rev-parse", f.Commit+"^{tree}")
	marker, _ := baselineGuard(t, f, "before-tree")
	for _, args := range [][]string{{"ls-tree", f.Commit}, {"cat-file", "blob", blob}} {
		if _, err := baselineInput(t, f.Root, "", args...); err == nil {
			t.Fatal("pre-resolution tree control passed")
		}
		if _, err := os.Stat(marker); err != nil {
			t.Fatal("missing premature read marker", err)
		}
		if err := os.Remove(marker); err != nil {
			t.Fatal(err)
		}
	}
	for _, revision := range []string{"", " ", "missing-ref", "--all", blob, tree} {
		t.Run(revision, func(t *testing.T) { baselineError(t, f.Root, revision, ""); baselineNoMarker(t, marker) })
	}
}
