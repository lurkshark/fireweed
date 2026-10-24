package evals

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	k "github.com/lurkshark/fireweed/packages/checks"
	c "github.com/lurkshark/fireweed/packages/contracts"
	g "github.com/lurkshark/fireweed/packages/gitstore"
	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
)

// The external tool is a deterministic project hook, not an agent or controller.
const verifyToolSource = `package main
import("encoding/json";"fmt";"os";"path/filepath";"strconv";"os/exec";"strings";"time")
func main(){if marker:=os.Getenv("FW_VERIFY_FORBIDDEN_MARKER");marker!=""&&strings.TrimSuffix(filepath.Base(os.Args[0]),".exe")!="hook"{os.WriteFile(marker,[]byte("forbidden command"),0600);os.Exit(95)};if len(os.Args)!=5{os.Exit(95)};kind,module,report,plan:=os.Args[1],os.Args[2],os.Args[3],os.Args[4];root,_:=os.Getwd();rec,_:=json.Marshal(map[string]string{"kind":kind,"module":module,"root":root,"report":report,"envmodule":os.Getenv("FIREWEED_MODULE"),"bin":os.Getenv("FIREWEED_BIN")});f,e:=os.OpenFile(filepath.Join(plan,"calls"),os.O_CREATE|os.O_APPEND|os.O_WRONLY,0600);if e!=nil{os.Exit(96)};f.Write(append(rec,'\n'));f.Close();key:=module;if kind=="boundary"{key="boundary"};if delay,e:=os.ReadFile(filepath.Join(plan,key+".hold"));e==nil{fmt.Println("hook-ready:"+module);d,_:=time.ParseDuration(strings.TrimSpace(string(delay)));os.WriteFile(filepath.Join(plan,key+".ready"),[]byte("ready"),0600);time.Sleep(d);os.WriteFile(filepath.Join(plan,key+".late"),[]byte("late"),0600)};if data,e:=os.ReadFile(filepath.Join(plan,key+".mutation"));e==nil{var m map[string]string;if json.Unmarshal(data,&m)!=nil{os.Exit(97)};for name,text:=range m{os.MkdirAll(filepath.Dir(name),0755);if os.WriteFile(name,[]byte(text),0644)!=nil{os.Exit(98)}}};if data,e:=os.ReadFile(filepath.Join(plan,key+".gitargs"));e==nil{var args []string;if json.Unmarshal(data,&args)!=nil{os.Exit(94)};cmd:=exec.Command("git",args...);if out,e:=cmd.CombinedOutput();e!=nil{fmt.Print(string(out));os.Exit(93)}};if kind=="eval"{data,e:=os.ReadFile(filepath.Join(plan,module+".xml"));if e==nil{if os.WriteFile(report,data,0600)!=nil{os.Exit(99)}};fmt.Println("eval-output:"+module)}else{if data,e:=os.ReadFile(filepath.Join(plan,"boundary.output"));e==nil{fmt.Print(string(data))}};code:=0;if data,e:=os.ReadFile(filepath.Join(plan,key+".exit"));e==nil{code,_=strconv.Atoi(strings.TrimSpace(string(data)))};os.Exit(code)}
`

func verifyTool(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	src := filepath.Join(dir, "hook.go")
	if e := os.WriteFile(src, []byte(verifyToolSource), 0600); e != nil {
		t.Fatal(e)
	}
	bin := filepath.Join(dir, "hook")
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "build", "-o", bin, src)
	cmd.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, e := cmd.CombinedOutput(); e != nil {
		t.Fatalf("project hook build %v %s", e, out)
	}
	return bin
}
func verifyCommand(tool, kind, plan string) string {
	return claimQuote(tool) + " " + kind + " {module} {report} " + claimQuote(plan)
}

type verifyFixtureData struct {
	Root, Path, Plan, Raw string
	Claim                 w.Claim
	Snapshot              p.Snapshot
	Module                c.Module
	Change                c.Change
}

func verifyFixture(t *testing.T, tool, phase, mode, format string, linked bool) verifyFixtureData {
	t.Helper()
	root := claimRepo(t, format)
	plan := stateCanonicalDirectory(t, t.TempDir())
	eval := verifyCommand(tool, "eval", plan)
	boundary := verifyCommand(tool, "boundary", plan)
	cfg := fmt.Sprintf("version: 1\nfireweed: '>=0.4, <0.5'\nbase: main\nmodules: [units/*]\nmodule_defaults:\n  evals: [e/**]\n  interface: [api/**]\n  disposable: [impl/**]\nhooks:\n  eval: %s\n  boundaries: %s\nprotected: [private/**]\nworker_writable: [shared/**]\nagent: {command: [forbidden-verify-agent]}\ntimeout: 5s\n", strconv.Quote(eval), strconv.Quote(boundary))
	if mode != "" {
		cfg += "verify_dependents: " + mode + "\n"
	}
	worktreePut(t, root, "fireweed.yaml", cfg)
	worktreePut(t, root, ".gitignore", "ignored*\n")
	worktreePut(t, root, "SYSTEM.md", "Fixture authority.\n")
	for _, m := range []struct{ name, prefix, dep string }{{"alpha", "ALP", ""}, {"beta", "BET", ""}, {"gamma", "GAM", "alpha"}, {"delta", "DEL", "beta, gamma"}, {"extra", "EXT", ""}} {
		meta := ""
		if m.dep != "" {
			meta = "depends_on: [" + m.dep + "]\n"
		}
		body := "---\nname: " + m.name + "\nid_prefix: " + m.prefix + "\n" + meta + "---\n\n## Specs\n\n### " + m.prefix + "-001\nThe retained behavior MUST remain.\n"
		if m.name == "alpha" {
			body += "\n### ALP-002\nThe removed behavior MUST disappear.\n\n### ALP-003 [review]\nThe reviewed behavior MUST remain visible.\n\n### ALP-004 [quarantine: fixture]\nThe quarantined behavior MUST remain visible.\n\n### ALP-006\nThe fixed behavior MUST remain.\n"
		}
		worktreePut(t, root, "units/"+m.name+"/SPEC.md", body)
		worktreePut(t, root, "units/"+m.name+"/impl/old", "retained implementation\n")
		worktreePut(t, root, "units/"+m.name+"/e/old", "retained eval\n")
		worktreePut(t, root, "units/"+m.name+"/api/old", "retained interface\n")
	}
	before := "---\nmodule: beta\nset:\n  depends_on: [alpha]\n---\n\n## Motivation\n\nPreceding approved intent.\n\n## Added\n\n### BET-002\nThe preceding behavior MUST be included.\n"
	later := "---\nmodule: beta\nset:\n  depends_on: []\n---\n\n## Motivation\n\nLater effects remain excluded.\n"
	worktreePut(t, root, "changes/0001-before.md", before)
	worktreePut(t, root, "changes/0003-later.md", later)
	path := "changes/0002-current.md"
	raw := "---\nmodule: alpha\nset:\n  evals: [new-e/**]\n---\n\n## Motivation\n\nVerify approved phases.\n\n## Interface\n\nDeclare retained contract changes.\n\n## Added\n\n### new:feature\nThe new behavior MUST be checked.\n"
	worktreePut(t, root, path, raw)
	startCommit(t, root, "Authored verification source")
	if linked {
		root = claimPeer(t, root)
	}
	s, e := p.Load(root)
	if e != nil {
		t.Fatal(e)
	}
	ch, _, e := p.Allocate(s, "current")
	if e != nil {
		t.Fatal(e)
	}
	ch.Document.Metadata["run"] = "verify-run"
	ch.Document.Metadata["phase"] = phase
	raw, e = c.RenderMarkdown(ch.Document)
	if e != nil {
		t.Fatal(e)
	}
	ch, e = c.ParseChange(raw, path, "ALP")
	if e != nil {
		t.Fatal(e)
	}
	// Allocation reserves ALP-006 before assigning the next ID.
	if len(ch.Added) != 1 {
		t.Fatal("fixture allocation", ch)
	}
	worktreePut(t, root, path, raw)
	cl, e := w.CreateClaim(root, "current", "verify-run", "HEAD", raw)
	if e != nil {
		t.Fatal(e)
	}
	snapshot, owned, e := w.LoadApproval(root, "current")
	if e != nil || owned != cl {
		t.Fatal("approval fixture", e)
	}
	projected, e := p.Project(snapshot, "current")
	if e != nil {
		t.Fatal(e)
	}
	var target c.Module
	for i, m := range projected.Modules {
		if m.Name == ch.Module {
			target, e = p.ApplyChange(m, ch)
			if e != nil {
				t.Fatal(e)
			}
			projected.Modules[i] = target
		}
	}
	for _, d := range p.Check(projected) {
		if d.Severity == "error" {
			t.Fatalf("inclusive fixture invalid %+v", d)
		}
	}
	guard := w.CheckWorktree(projected, target, path, cl)
	if len(guard.Diagnostics) != 0 {
		t.Fatal("valid guard premise", guard)
	}
	f := verifyFixtureData{root, path, plan, raw, cl, snapshot, target, ch}
	verifyDefaultReports(t, f, phase == "evals")
	worktreePut(t, root, "ignored-cache", "retained ignored data")
	return f
}
func verifyDefaultReports(t *testing.T, f verifyFixtureData, fail bool) {
	t.Helper()
	for _, m := range f.Snapshot.Modules {
		ids := []string{m.IDPrefix + "-001"}
		failed := map[string]bool{}
		if m.Name == "alpha" {
			ids = []string{"ALP-001", "ALP-002", "ALP-006", f.Change.Added[0].ID}
			if fail {
				failed[f.Change.Added[0].ID] = true
			}
		}
		if m.Name == "beta" {
			ids = append(ids, "BET-002")
		}
		verifyReport(t, f, m.Name, ids, failed)
	}
}
func verifyReport(t *testing.T, f verifyFixtureData, module string, ids []string, failed map[string]bool) {
	t.Helper()
	var b strings.Builder
	b.WriteString("<testsuite>")
	n := 0
	for _, id := range ids {
		fmt.Fprintf(&b, "<testcase name=\"case_%s\">", id)
		if failed[id] {
			b.WriteString("<failure message=\"required behavior absent\"/>")
			n++
		}
		b.WriteString("</testcase>")
	}
	b.WriteString("</testsuite>")
	worktreePut(t, f.Plan, module+".xml", b.String())
	code := "0"
	if n > 0 {
		code = "1"
	}
	worktreePut(t, f.Plan, module+".exit", code)
}

type verifyCall struct{ Kind, Module, Root, Report, Envmodule, Bin string }

func verifyCalls(t *testing.T, plan string) []verifyCall {
	t.Helper()
	raw, e := os.ReadFile(filepath.Join(plan, "calls"))
	if os.IsNotExist(e) {
		return []verifyCall{}
	}
	if e != nil {
		t.Fatal(e)
	}
	out := []verifyCall{}
	for _, line := range strings.Split(strings.TrimSpace(string(raw)), "\n") {
		var x verifyCall
		if json.Unmarshal([]byte(line), &x) != nil {
			t.Fatal("bad hook log", line)
		}
		out = append(out, x)
	}
	return out
}
func verifyClear(t *testing.T, f verifyFixtureData) {
	t.Helper()
	if e := os.Remove(filepath.Join(f.Plan, "calls")); e != nil && !os.IsNotExist(e) {
		t.Fatal(e)
	}
}
func verifyShape(t *testing.T, f verifyFixtureData, v w.Verification) {
	t.Helper()
	if v.Change != "current" || v.Phase != f.Change.Phase || v.Diagnostics == nil || v.Coverage == nil || v.Edits == nil || v.Findings == nil {
		t.Errorf("incomplete verification %+v", v)
	}
}
func verifyNoErrors(t *testing.T, ds []k.Diagnostic) {
	t.Helper()
	for _, d := range ds {
		if d.Severity == "error" {
			t.Errorf("unexpected verification error %+v", d)
		}
	}
}
func verifyCode(t *testing.T, v w.Verification, code, path string) {
	t.Helper()
	for _, d := range v.Diagnostics {
		if d.Code == code && (path == "" || d.Path == path) && d.Severity == "error" && policyCorrective(d.Message) {
			return
		}
	}
	t.Errorf("missing useful %s at %s in %+v", code, path, v.Diagnostics)
}
func verifyInvoke(t *testing.T, f verifyFixtureData, slug string) w.Verification {
	t.Helper()
	v, e := w.Verify(context.Background(), f.Root, slug)
	if e != nil {
		t.Errorf("valid selected source returned fatal error: %v", e)
	}
	verifyShape(t, f, v)
	return v
}
func verifyExpectedCoverage(t *testing.T, f verifyFixtureData, modules []string, v w.Verification) {
	t.Helper()
	if len(v.Coverage) != len(modules) {
		t.Errorf("coverage modules %+v want %v", v.Coverage, modules)
		return
	}
	seen := map[string]bool{}
	for i, cov := range v.Coverage {
		if cov.Module != modules[i] || seen[cov.Module] {
			t.Errorf("coverage order/duplication got %s want %s", cov.Module, modules[i])
		}
		seen[cov.Module] = true
	}
}
func verifyReseed(t *testing.T, f *verifyFixtureData, raw string) {
	t.Helper()
	cl, e := w.UpdateClaim(f.Root, f.Claim, f.Claim.Base, raw)
	if e != nil {
		t.Fatal(e)
	}
	f.Claim = cl
	f.Raw = raw
	worktreePut(t, f.Root, f.Path, raw)
	f.Change, e = c.ParseChange(raw, f.Path, "ALP")
	if e != nil {
		t.Fatal(e)
	}
	f.Snapshot, _, e = w.LoadApproval(f.Root, "current")
	if e != nil {
		t.Fatal(e)
	}
	pre, e := p.Project(f.Snapshot, "current")
	if e != nil {
		t.Fatal(e)
	}
	for _, m := range pre.Modules {
		if m.Name == "alpha" {
			f.Module, e = p.ApplyChange(m, f.Change)
			if e != nil {
				t.Fatal(e)
			}
		}
	}
}

func Test_FLOW041VerifySelection(t *testing.T) {
	tool := verifyTool(t)
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/linked_%v", format, linked), func(t *testing.T) {
				f := verifyFixture(t, tool, "audit", "", format, linked)
				worktreePut(t, f.Root, f.Path, f.Raw+"\n## Audit\n\nNo findings. Checked "+f.Change.Added[0].ID+".\n")
				claimRaw(t, f.Root, "unrelated-malformed", "invalid JSON")
				v, e := w.Verify(context.Background(), filepath.Join(f.Root, "units/alpha/e"), "current")
				if e != nil {
					t.Fatalf("explicit independent selection %v", e)
				}
				verifyShape(t, f, v)
				verifyNoErrors(t, v.Diagnostics)
				if len(verifyCalls(t, f.Plan)) != 0 {
					t.Error("audit executed hooks")
				}
				if e := w.ReleaseClaim(f.Root, "unrelated-malformed", claimGit(t, f.Root, "rev-parse", "refs/fireweed/claims/unrelated-malformed")); e != nil {
					t.Fatal(e)
				}
				v = verifyInvoke(t, f, "")
				verifyNoErrors(t, v.Diagnostics)
			})
		}
	}
	for _, slug := range []string{"UPPER", "a/b", "-a", "a b", "a\n"} {
		v, e := w.Verify(context.Background(), "/nonexistent/verification-repo", slug)
		if !reflect.DeepEqual(v, w.Verification{}) || e == nil || !strings.Contains(e.Error(), "C05") || !policyCorrective(e.Error()) {
			t.Errorf("invalid selector %q %+v %v", slug, v, e)
		}
	}
}
func Test_FLOW041VerifyClaimRejections(t *testing.T) {
	tool := verifyTool(t)
	for _, mode := range []string{"absent", "missing-selected", "malformed", "foreign", "multiple-owned", "invalid-state", "missing-base"} {
		t.Run(mode, func(t *testing.T) {
			f := verifyFixture(t, tool, "impl", "none", "sha1", false)
			slug := "current"
			switch mode {
			case "absent":
				w.ReleaseClaim(f.Root, "current", f.Claim.Object)
			case "missing-selected":
				slug = "absent"
			case "malformed":
				claimRaw(t, f.Root, "current", "not JSON")
			case "foreign":
				x := f.Claim
				x.Worktree = claimPeer(t, f.Root)
				claimRaw(t, f.Root, "current", claimPayload(t, x))
			case "multiple-owned":
				claimSeed(t, f.Root, "other", f.Raw)
				slug = ""
			case "invalid-state":
				x := f.Claim
				x.State = strings.Replace(x.State, "phase: impl", "phase: applied", 1)
				claimRaw(t, f.Root, "current", claimPayload(t, x))
			case "missing-base":
				x := f.Claim
				x.Base = strings.Repeat("a", len(x.Base))
				claimRaw(t, f.Root, "current", claimPayload(t, x))
			}
			files := startTree(t, f.Root)
			meta := claimTree(t, claimCommon(t, f.Root), func(string) bool { return false })
			v, e := w.Verify(context.Background(), f.Root, slug)
			if !reflect.DeepEqual(v, w.Verification{}) || e == nil || !strings.Contains(e.Error(), "P02") || !policyCorrective(e.Error()) {
				t.Errorf("claim rejection %+v %v", v, e)
			}
			if len(verifyCalls(t, f.Plan)) != 0 || !reflect.DeepEqual(files, startTree(t, f.Root)) || !reflect.DeepEqual(meta, claimTree(t, claimCommon(t, f.Root), func(string) bool { return false })) {
				t.Error("rejected claim executed hooks or changed files/metadata")
			}
		})
	}
}

func Test_FLOW042VerifyGuardsBeforeHooks(t *testing.T) {
	tool := verifyTool(t)
	for _, mode := range []string{"protected", "foreign-module", "module-spec", "other-queue", "frozen-eval", "change-metadata", "missing-change", "unsafe-change", "blocked", "moved-head", "invalid-graph", "projected-eval", "projected-disposable"} {
		t.Run(mode, func(t *testing.T) {
			f := verifyFixture(t, tool, "impl", "all", "sha1", false)
			code, path := "P2-01", f.Path
			switch mode {
			case "protected":
				path = "private/forbidden"
				worktreePut(t, f.Root, path, "forbidden")
			case "module-spec":
				path = "units/alpha/SPEC.md"
				worktreePut(t, f.Root, path, "invalid worker source")
			case "other-queue":
				path = "changes/0001-before.md"
				worktreePut(t, f.Root, path, "invalid worker source")
			case "foreign-module":
				path = "units/beta/impl/old"
				worktreePut(t, f.Root, path, "forbidden")
			case "frozen-eval":
				path = "units/alpha/new-e/frozen"
				worktreePut(t, f.Root, path, "forbidden")
			case "projected-eval":
				path = "units/alpha/new-e/new"
				worktreePut(t, f.Root, path, "frozen despite stale eval zone")
			case "projected-disposable":
				path = "units/alpha/e/old"
				worktreePut(t, f.Root, path, "now ordinary owned file")
				code = ""
			case "change-metadata":
				worktreePut(t, f.Root, f.Path, strings.Replace(f.Raw, "phase: impl", "phase: evals", 1))
			case "missing-change":
				if e := os.Remove(filepath.Join(f.Root, f.Path)); e != nil {
					t.Fatal(e)
				}
			case "unsafe-change":
				outside := filepath.Join(t.TempDir(), "private")
				if e := os.WriteFile(outside, []byte(f.Raw), 0600); e != nil {
					t.Fatal(e)
				}
				if e := os.Remove(filepath.Join(f.Root, f.Path)); e != nil {
					t.Fatal(e)
				}
				if e := os.Symlink(outside, filepath.Join(f.Root, f.Path)); e != nil {
					t.Fatal(e)
				}
			case "blocked":
				code = "P05"
				worktreePut(t, f.Root, f.Path, strings.Replace(f.Raw, "module: alpha\n", "module: alpha\nblocked: awaiting an answer\n", 1))
			case "moved-head":
				code = "P02"
				claimGit(t, f.Root, "-c", "user.name=Verify Fixture", "-c", "user.email=verify@example.invalid", "commit", "--allow-empty", "-qm", "worker commit")
			case "invalid-graph":
				code = "C05"
				doc := f.Change.Document
				set := doc.Metadata["set"].(map[string]any)
				set["depends_on"] = []any{"missing"}
				raw, err := c.RenderMarkdown(doc)
				if err != nil {
					t.Fatal(err)
				}
				verifyReseed(t, &f, raw)
			}
			before := startTree(t, f.Root)
			meta := claimTree(t, claimCommon(t, f.Root), func(string) bool { return false })
			v := verifyInvoke(t, f, "current")
			if mode != "invalid-graph" {
				verifyGuardEvidence(t, f, v)
			}
			if code != "" {
				if mode == "invalid-graph" {
					found := false
					for _, d := range v.Diagnostics {
						if (d.Code == "C05" || d.Code == "C06") && d.Severity == "error" && policyCorrective(d.Message) && strings.Contains(d.Message, "missing") {
							found = true
						}
					}
					if !found {
						t.Errorf("missing useful effective graph rejection: %+v", v.Diagnostics)
					}
				} else {
					verifyCode(t, v, code, path)
				}
				if len(v.Coverage) != 0 || len(verifyCalls(t, f.Plan)) != 0 {
					t.Error("failed guards executed verification hooks")
				}
			} else {
				verifyNoErrors(t, v.Diagnostics)
				verifyExpectedCoverage(t, f, []string{"alpha", "beta", "delta", "gamma"}, v)
			}
			if !reflect.DeepEqual(before, startTree(t, f.Root)) || !reflect.DeepEqual(meta, claimTree(t, claimCommon(t, f.Root), func(string) bool { return false })) {
				t.Error("guard changed original files/metadata")
			}
		})
	}
}

func verifyManualCoverage(t *testing.T, f verifyFixtureData, ids, already []string) k.Coverage {
	t.Helper()
	s, e := p.Project(f.Snapshot, "current")
	if e != nil {
		t.Fatal(e)
	}
	for i, m := range s.Modules {
		if m.Name == "alpha" {
			s.Modules[i] = f.Module
		}
	}
	v := k.Evaluate(context.Background(), s, f.Module, ids, already)
	v = verifyNormalizeCoverage(t, f, v)
	verifyClear(t, f)
	return v
}
func Test_FLOW043VerifyEvalsEvidence(t *testing.T) {
	tool := verifyTool(t)
	for _, mode := range []string{"failing-new", "passing-new", "missing-id", "unknown-id", "skipped-new", "missing-report", "malformed-report", "ordinary-nonzero", "already-satisfied", "worker-hook-edit"} {
		t.Run(mode, func(t *testing.T) {
			f := verifyFixture(t, tool, "evals", "all", "sha1", false)
			id := f.Change.Added[0].ID
			ids := []string{"ALP-001", "ALP-002", "ALP-006", id}
			already := []string{}
			switch mode {
			case "passing-new":
				verifyReport(t, f, "alpha", ids, nil)
			case "missing-id":
				verifyReport(t, f, "alpha", ids[:3], nil)
			case "unknown-id":
				verifyReport(t, f, "alpha", append(ids, "ALP-999"), map[string]bool{id: true})
			case "skipped-new":
				worktreePut(t, f.Plan, "alpha.xml", strings.Replace(mustReadVerify(t, filepath.Join(f.Plan, "alpha.xml")), "<failure message=\"required behavior absent\"/>", "<skipped/>", 1))
				worktreePut(t, f.Plan, "alpha.exit", "0")
			case "missing-report":
				os.Remove(filepath.Join(f.Plan, "alpha.xml"))
			case "malformed-report":
				worktreePut(t, f.Plan, "alpha.xml", "<testsuite>")
			case "ordinary-nonzero":
				worktreePut(t, f.Plan, "alpha.exit", "7")
			case "already-satisfied":
				verifyReseed(t, &f, strings.Replace(f.Raw, "module: alpha\n", "module: alpha\nalready_satisfied: ["+id+"]\n", 1))
				verifyReport(t, f, "alpha", ids, nil)
				already = []string{id}
			case "worker-hook-edit":
				worktreePut(t, f.Root, "fireweed.yaml", "version: 1\nmodules: [outside/*]\nhooks: {eval: forbidden-verify-agent}\n")
			}
			if mode == "worker-hook-edit" {
				v := verifyInvoke(t, f, "current")
				verifyCode(t, v, "P2-01", "fireweed.yaml")
				if len(verifyCalls(t, f.Plan)) != 0 {
					t.Error("worker configuration executed")
				}
				return
			}
			want := verifyManualCoverage(t, f, []string{id}, already)
			v := verifyInvoke(t, f, "current")
			verifyExpectedCoverage(t, f, []string{"alpha"}, v)
			if len(v.Coverage) == 1 && !reflect.DeepEqual(verifyNormalizeCoverage(t, f, v.Coverage[0]), want) {
				t.Errorf("lost exact Evaluate evidence\n got %+v\nwant %+v", v.Coverage[0], want)
			}
			if !reflect.DeepEqual(verifyNormalizeDiagnostics(t, f, v.Diagnostics), want.Diagnostics) {
				t.Errorf("coverage diagnostics differ: got %+v want %+v", v.Diagnostics, want.Diagnostics)
			}
			calls := verifyCalls(t, f.Plan)
			if len(calls) != 1 || calls[0].Kind != "eval" || calls[0].Module != "alpha" {
				t.Errorf("evals invoked extra/wrong checks %+v", calls)
			}
			if mode == "failing-new" || mode == "already-satisfied" || mode == "ordinary-nonzero" {
				verifyNoErrors(t, v.Diagnostics)
			} else if len(v.Diagnostics) == 0 {
				t.Error("invalid fail-first evidence accepted")
			}
		})
	}
}
func mustReadVerify(t *testing.T, path string) string {
	t.Helper()
	b, e := os.ReadFile(path)
	if e != nil {
		t.Fatal(e)
	}
	return string(b)
}

func Test_FLOW044VerifyAuditFindings(t *testing.T) {
	tool := verifyTool(t)
	for _, mode := range []string{"valid", "untestable", "unknown", "inexact", "empty-clause", "review", "quarantine", "malformed", "absent", "empty", "missing-id", "prefix-id", "passing", "case-passing"} {
		t.Run(mode, func(t *testing.T) {
			f := verifyFixture(t, tool, "audit", "all", "sha1", false)
			id := f.Change.Added[0].ID
			line := "- " + id + ": \"MUST be checked\" — missing edge assertion"
			audit := line
			wantFinding := true
			switch mode {
			case "untestable":
				audit = "- [untestable] " + id + ": \"MUST be checked\" — impossible claim"
			case "unknown":
				audit = strings.Replace(line, id, "ALP-999", 1)
				wantFinding = false
			case "inexact":
				audit = strings.Replace(line, "MUST be checked", "must be checked", 1)
				wantFinding = false
			case "empty-clause":
				audit = strings.Replace(line, "MUST be checked", "", 1)
				wantFinding = false
			case "review":
				audit = "- ALP-003: \"MUST remain visible\" — no assertion"
				wantFinding = false
			case "quarantine":
				audit = "- ALP-004: \"MUST remain visible\" — no assertion"
				wantFinding = false
			case "malformed":
				audit = id + ": \"MUST be checked\" — missing list marker"
				wantFinding = false
			case "absent", "empty":
				audit = ""
				wantFinding = false
			case "missing-id":
				audit = "No findings. Checked ALP-001."
				wantFinding = false
			case "prefix-id":
				audit = "No findings. Checked " + id + "0."
				wantFinding = false
			case "passing":
				audit = "No findings. Checked case_" + id + "_edge."
				wantFinding = false
			case "case-passing":
				audit = "nO fInDiNgS. Checked " + id + "."
				wantFinding = false
			}
			text := f.Raw
			if mode != "absent" {
				text += "\n## Audit\n\n" + audit + "\n"
			}
			worktreePut(t, f.Root, f.Path, text)
			before := startTree(t, f.Root)
			meta := claimTree(t, claimCommon(t, f.Root), func(string) bool { return false })
			v := verifyInvoke(t, f, "current")
			if len(v.Coverage) != 0 || len(verifyCalls(t, f.Plan)) != 0 {
				t.Error("audit executed hooks")
			}
			if wantFinding {
				want := []w.Finding{{ID: id, Clause: "MUST be checked", Text: audit, Untestable: mode == "untestable"}}
				if !reflect.DeepEqual(v.Findings, want) {
					t.Errorf("findings %+v want %+v", v.Findings, want)
				}
				verifyNoErrors(t, v.Diagnostics)
			} else if mode == "passing" || mode == "case-passing" {
				if len(v.Findings) != 0 {
					t.Error("spurious finding")
				}
				verifyNoErrors(t, v.Diagnostics)
			} else {
				if len(v.Findings) != 0 {
					t.Error("retained invalid finding")
				}
				verifyCode(t, v, "P1-03", f.Path)
			}
			if !reflect.DeepEqual(before, startTree(t, f.Root)) || !reflect.DeepEqual(meta, claimTree(t, claimCommon(t, f.Root), func(string) bool { return false })) {
				t.Error("audit advanced/blocked/wrote questions or mutated source")
			}
		})
	}
}

func Test_FLOW045VerifyDependentComposition(t *testing.T) {
	tool := verifyTool(t)
	for _, mode := range []string{"", "all", "direct", "none"} {
		for _, failing := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/failing_%v", mode, failing), func(t *testing.T) {
				f := verifyFixture(t, tool, "impl", mode, "sha1", false)
				if failing {
					verifyReport(t, f, "alpha", []string{"ALP-001", "ALP-002", "ALP-006", f.Change.Added[0].ID}, map[string]bool{"ALP-001": true})
					verifyReport(t, f, "beta", []string{"BET-001", "BET-002"}, map[string]bool{"BET-002": true})
					worktreePut(t, f.Plan, "boundary.output", "alpha: illegal dependency\n")
				}
				mods := []string{"alpha", "beta", "delta", "gamma"}
				if mode == "direct" {
					mods = []string{"alpha", "beta", "gamma"}
				} else if mode == "none" {
					mods = []string{"alpha"}
				}
				expected := verifyModuleCoverage(t, f, mods)
				v := verifyInvoke(t, f, "current")
				verifyExpectedCoverage(t, f, mods, v)
				for i, cov := range v.Coverage {
					if i < len(expected) && !reflect.DeepEqual(verifyNormalizeCoverage(t, f, cov), expected[i]) {
						t.Errorf("dependent evidence lost for %s: got %+v want %+v", cov.Module, cov, expected[i])
					}
				}
				verifyDiagnosticOrder(t, v.Diagnostics)
				calls := verifyCalls(t, f.Plan)
				want := []string{}
				for _, m := range mods {
					want = append(want, "eval/"+m)
				}
				want = append(want, "boundary/")
				actual := []string{}
				for _, call := range calls {
					actual = append(actual, call.Kind+"/"+call.Module)
					if call.Root != f.Root || call.Envmodule != call.Module {
						t.Errorf("hook identity %+v", call)
					}
				}
				if !reflect.DeepEqual(actual, want) {
					t.Errorf("hook composition got %v want %v", actual, want)
				}
				if failing {
					verifyCode(t, v, "C14", "units/alpha")
					verifyCode(t, v, "C15", "units/alpha")
					if mode != "none" {
						verifyCode(t, v, "C14", "units/beta")
					}
				} else {
					verifyNoErrors(t, v.Diagnostics)
				}
			})
		}
	}
}

func Test_FLOW046VerifyHookEditsAndOwnership(t *testing.T) {
	tool := verifyTool(t)
	for _, mode := range []string{"protected", "frozen-eval", "state", "head", "claim"} {
		t.Run(mode, func(t *testing.T) {
			f := verifyFixture(t, tool, "impl", "none", "sha1", false)
			path, code := "private/from-hook", "P2-01"
			mutations := map[string]string{}
			switch mode {
			case "protected":
				mutations[path] = "hook must remain observable"
			case "frozen-eval":
				path = "units/alpha/new-e/forbidden"
				mutations[path] = "forbidden frozen bytes"
			case "state":
				path = f.Path
				mutations[path] = strings.Replace(f.Raw, "Verify approved phases.", "Hook changed motivation.", 1)
			case "head":
				code = "P02"
				path = f.Path
				gitArgs := []string{"update-ref", "HEAD", claimGit(t, f.Root, "rev-parse", "HEAD~1")}
				raw, _ := json.Marshal(gitArgs)
				worktreePut(t, f.Plan, "alpha.gitargs", string(raw))
			case "claim":
				code = "P02"
				path = f.Path
				gitArgs := []string{"update-ref", "-d", "refs/fireweed/claims/current"}
				raw, _ := json.Marshal(gitArgs)
				worktreePut(t, f.Plan, "alpha.gitargs", string(raw))
			}
			if len(mutations) > 0 {
				raw, _ := json.Marshal(mutations)
				worktreePut(t, f.Plan, "alpha.mutation", string(raw))
			}
			v := verifyInvoke(t, f, "current")
			verifyGuardEvidence(t, f, v)
			verifyExpectedCoverage(t, f, []string{"alpha"}, v)
			verifyCode(t, v, code, path)
			if len(v.Coverage) == 1 && len(v.Coverage[0].Tests) == 0 {
				t.Error("hook mutation hid otherwise available coverage")
			}
			for name, want := range mutations {
				if got := mustReadVerify(t, filepath.Join(f.Root, name)); got != want {
					t.Error("verification repaired/erased hook changes")
				}
				if name != f.Path {
					found := false
					for _, ed := range v.Edits {
						found = found || ed.Path == name
					}
					if !found {
						t.Error("result Edits omit hook mutation", name)
					}
				}
			}
			if mode == "claim" {
				if _, e := w.LoadClaim(f.Root, "current"); e == nil {
					t.Error("verification repaired released claim")
				}
			}
			if mode == "head" && claimBase(t, f.Root) == f.Claim.Base {
				t.Error("verification reset hook-moved HEAD")
			}
		})
	}
}

func Test_FLOW046VerifyCancellation(t *testing.T) {
	tool := verifyTool(t)
	t.Run("precancelled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		v, e := w.Verify(ctx, "/nonexistent/verify-repository", "")
		if !reflect.DeepEqual(v, w.Verification{}) || !errors.Is(e, context.Canceled) {
			t.Errorf("cancel before source %+v %v", v, e)
		}
	})
	for _, deadline := range []bool{false, true} {
		t.Run(fmt.Sprint(deadline), func(t *testing.T) {
			f := verifyFixture(t, tool, "impl", "all", "sha1", false)
			worktreePut(t, f.Plan, "alpha.hold", "6s")
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			if deadline {
				var timed context.CancelFunc
				ctx, timed = context.WithTimeout(ctx, 3*time.Second)
				defer timed()
			}
			type answer struct {
				v w.Verification
				e error
			}
			ch := make(chan answer, 1)
			go func() { v, e := w.Verify(ctx, f.Root, "current"); ch <- answer{v, e} }()
			until := time.Now().Add(5 * time.Second)
			for {
				if _, e := os.Stat(filepath.Join(f.Plan, "alpha.ready")); e == nil {
					break
				}
				select {
				case got := <-ch:
					t.Fatalf("verification ended before active hook control: %+v %v", got.v, got.e)
				default:
				}
				if time.Now().After(until) {
					t.Fatal("hook readiness not observed")
				}
				time.Sleep(10 * time.Millisecond)
			}
			if !deadline {
				cancel()
			}
			select {
			case got := <-ch:
				want := context.Canceled
				if deadline {
					want = context.DeadlineExceeded
				}
				if !errors.Is(got.e, want) {
					t.Errorf("lost cancellation error: %v", got.e)
				}
				verifyShape(t, f, got.v)
				if len(got.v.Coverage) != 1 {
					t.Errorf("cancelled own evidence missing %+v", got.v.Coverage)
				}
			case <-time.After(5 * time.Second):
				t.Fatal("cancelled verification did not terminate")
			}
			calls := verifyCalls(t, f.Plan)
			if len(calls) != 1 || calls[0].Module != "alpha" {
				t.Errorf("later checks executed after cancellation %+v", calls)
			}
			time.Sleep(6100 * time.Millisecond)
			if _, e := os.Stat(filepath.Join(f.Plan, "alpha.late")); !os.IsNotExist(e) {
				t.Error("cancelled hook survived")
			}
		})
	}
}

func Test_FLOW047VerifyLocalReadOnlyAndResultOwnership(t *testing.T) {
	tool := verifyTool(t)
	for _, phase := range []string{"evals", "audit", "impl"} {
		t.Run(phase, func(t *testing.T) {
			f := verifyFixture(t, tool, phase, "all", "sha1", false)
			if phase == "audit" {
				worktreePut(t, f.Root, f.Path, f.Raw+"\n## Audit\n\n- "+f.Change.Added[0].ID+": \"MUST be checked\" — retained finding\n")
			}
			if phase == "impl" {
				verifyReport(t, f, "alpha", []string{"ALP-001", "ALP-002", "ALP-006", f.Change.Added[0].ID}, map[string]bool{"ALP-001": true})
			}
			temp := stateCanonicalDirectory(t, t.TempDir())
			t.Setenv("TMPDIR", temp)
			t.Setenv("TMP", temp)
			t.Setenv("TEMP", temp)
			claimGit(t, f.Root, "--version")
			files := startTree(t, f.Root)
			meta := claimTree(t, claimCommon(t, f.Root), func(string) bool { return false })
			scratch := claimTree(t, temp, func(string) bool { return false })
			cwd, _ := os.Getwd()
			env := os.Environ()
			first := verifyInvoke(t, f, "current")
			second := verifyInvoke(t, f, "current")
			saved := policyJSON(t, second)
			if len(first.Edits) > 0 {
				first.Edits[0].Path = "mutated"
			}
			if len(first.Diagnostics) > 0 {
				first.Diagnostics[0].Message = "mutated"
			}
			if len(first.Findings) > 0 {
				first.Findings[0].Clause = "mutated"
			}
			if len(first.Coverage) > 0 {
				first.Coverage[0].Module = "mutated"
				if len(first.Coverage[0].Diagnostics) > 0 {
					first.Coverage[0].Diagnostics[0].Message = "mutated nested coverage diagnostic"
				}
				if len(first.Coverage[0].Tests) > 0 {
					first.Coverage[0].Tests[0].Name = "mutated"
					if len(first.Coverage[0].Tests[0].IDs) > 0 {
						first.Coverage[0].Tests[0].IDs[0] = "mutated"
					}
					if len(first.Coverage[0].Tests[0].Failures) > 0 {
						first.Coverage[0].Tests[0].Failures[0].Text = "mutated"
					}
				}
				for id := range first.Coverage[0].BySpec {
					if len(first.Coverage[0].BySpec[id]) > 0 {
						first.Coverage[0].BySpec[id][0] = -1
					}
					first.Coverage[0].BySpec[id] = []int{-2}
					break
				}
			}
			first.Edits = append(first.Edits, g.Edit{Path: "mutated"})
			first.Diagnostics = append(first.Diagnostics, k.Diagnostic{Message: "mutated"})
			first.Findings = append(first.Findings, w.Finding{Text: "mutated"})
			if saved != policyJSON(t, second) {
				t.Error("results alias")
			}
			after, _ := os.Getwd()
			if after != cwd || !reflect.DeepEqual(env, os.Environ()) || !reflect.DeepEqual(files, startTree(t, f.Root)) || !reflect.DeepEqual(meta, claimTree(t, claimCommon(t, f.Root), func(string) bool { return false })) || !reflect.DeepEqual(scratch, claimTree(t, temp, func(string) bool { return false })) {
				t.Error("verification changed parent state/source/Git metadata or leaked owned temporary storage")
			}
			for _, call := range verifyCalls(t, f.Plan) {
				if call.Kind == "eval" {
					if call.Report == "" {
						t.Error("no report identity")
					}
					if _, e := os.Stat(filepath.Dir(call.Report)); !os.IsNotExist(e) {
						t.Error("owned report directory leaked")
					}
				}
			}
			b, e := json.Marshal(second)
			if e != nil {
				t.Fatal(e)
			}
			var wire map[string]json.RawMessage
			if json.Unmarshal(b, &wire) != nil || len(wire) < 5 || wire["change"] == nil || wire["phase"] == nil || string(wire["diagnostics"]) == "null" || string(wire["coverage"]) == "null" || string(wire["edits"]) == "null" {
				t.Error("verification wire shape", string(b))
			}
			sorted := append([]k.Diagnostic{}, second.Diagnostics...)
			sort.Slice(sorted, func(i, j int) bool {
				a, b := sorted[i], sorted[j]
				if a.Path != b.Path {
					return a.Path < b.Path
				}
				if a.Code != b.Code {
					return a.Code < b.Code
				}
				if a.Severity != b.Severity {
					return a.Severity < b.Severity
				}
				return a.Message < b.Message
			})
			if !reflect.DeepEqual(sorted, second.Diagnostics) {
				t.Error("diagnostic order")
			}
			for i := 1; i < len(sorted); i++ {
				if sorted[i] == sorted[i-1] {
					t.Error("duplicate diagnostics")
				}
			}
		})
	}
}

func Test_FLOW041_FLOW042_FLOW043_FLOW044_FLOW045_FLOW046_FLOW047VerifyFixtureControls(t *testing.T) {
	tool := verifyTool(t)
	for _, phase := range []string{"evals", "impl", "audit"} {
		f := verifyFixture(t, tool, phase, "all", "sha1", false)
		s, e := p.Project(f.Snapshot, "current")
		if e != nil {
			t.Fatal(e)
		}
		mods := map[string]c.Module{}
		for i, m := range s.Modules {
			if m.Name == "alpha" {
				s.Modules[i] = f.Module
			}
			mods[s.Modules[i].Name] = s.Modules[i]
		}
		if !reflect.DeepEqual(mods["beta"].DependsOn, []string{"alpha"}) || len(mods["beta"].Specs) != 2 || len(mods["alpha"].Specs) != 6 || !reflect.DeepEqual(mods["alpha"].Evals, []string{"new-e/**"}) {
			t.Fatal("preceding/current/later fixture does not distinguish once-only inclusive projection", mods)
		}
		guard := w.CheckWorktree(s, f.Module, f.Path, f.Claim)
		if len(guard.Diagnostics) != 0 {
			t.Fatal(guard)
		}
		if phase != "audit" {
			ids := []string{}
			if phase == "evals" {
				ids = []string{f.Change.Added[0].ID}
			}
			cov := k.Evaluate(context.Background(), s, f.Module, ids, nil)
			if cov.HookExit != map[bool]int{true: 1, false: 0}[phase == "evals"] || len(cov.Tests) != 4 || len(cov.Diagnostics) != 0 {
				t.Fatal("deterministic local hook/fail-first control", cov)
			}
			for _, d := range k.Boundaries(context.Background(), s) {
				if d.Severity == "error" {
					t.Fatal("boundary fixture control", d)
				}
			}
			calls := verifyCalls(t, f.Plan)
			if len(calls) != 2 || calls[0].Module != "alpha" || calls[1].Kind != "boundary" {
				t.Fatal("direct adapter fixture liveness", calls)
			}
		}
	}
}
func verifyForbiddenCommands(t *testing.T, tool string) string {
	t.Helper()
	dir := t.TempDir()
	marker := filepath.Join(dir, "marker")
	raw, e := os.ReadFile(tool)
	if e != nil {
		t.Fatal(e)
	}
	for _, name := range []string{"forbidden-verify-agent", "codex", "claude", "fireweed", "fireweed-bootstrap", "curl", "wget", "ssh", "nc"} {
		if runtime.GOOS == "windows" {
			name += ".exe"
		}
		if e := os.WriteFile(filepath.Join(dir, name), raw, 0755); e != nil {
			t.Fatal(e)
		}
	}
	t.Setenv("FW_VERIFY_FORBIDDEN_MARKER", marker)
	controller := "fireweed-bootstrap"
	if runtime.GOOS == "windows" {
		controller += ".exe"
	}
	t.Setenv("FIREWEED_BIN", filepath.Join(dir, controller))
	t.Setenv("PATH", dir+string(os.PathListSeparator)+os.Getenv("PATH"))
	if e := exec.Command("forbidden-verify-agent").Run(); e == nil {
		t.Fatal("forbidden-command control succeeded")
	}
	if _, e := os.Stat(marker); e != nil {
		t.Fatal("forbidden-command marker not live", e)
	}
	if e := os.Remove(marker); e != nil {
		t.Fatal(e)
	}
	return marker
}
func Test_FLOW047VerifyNoHiddenCommands(t *testing.T) {
	tool := verifyTool(t)
	for _, mode := range []string{"success", "blocked", "zone-error", "missing-claim"} {
		t.Run(mode, func(t *testing.T) {
			f := verifyFixture(t, tool, "impl", "none", "sha1", false)
			switch mode {
			case "blocked":
				worktreePut(t, f.Root, f.Path, strings.Replace(f.Raw, "module: alpha\n", "module: alpha\nblocked: pending author\n", 1))
			case "zone-error":
				worktreePut(t, f.Root, "private/forbidden", "worker mutation")
			case "missing-claim":
				if e := w.ReleaseClaim(f.Root, "current", f.Claim.Object); e != nil {
					t.Fatal(e)
				}
			}
			marker := verifyForbiddenCommands(t, tool)
			trace := filepath.Join(t.TempDir(), "trace")
			t.Setenv("GIT_TRACE2_EVENT", trace)
			t.Setenv("GIT_TRACE2_ENV_VARS", "GIT_DIR")
			verifyGitRecorder(t, trace)
			owned := stateCanonicalDirectory(t, t.TempDir())
			t.Setenv("TMPDIR", owned)
			t.Setenv("TMP", owned)
			t.Setenv("TEMP", owned)
			v, e := w.Verify(context.Background(), f.Root, "current")
			if mode == "missing-claim" {
				if !reflect.DeepEqual(v, w.Verification{}) || e == nil {
					t.Error("missing claim unexpectedly accepted")
				}
			} else {
				if e != nil {
					t.Error(e)
				}
				verifyShape(t, f, v)
				if mode == "success" {
					verifyNoErrors(t, v.Diagnostics)
				} else if len(v.Diagnostics) == 0 {
					t.Error("rejected case passed")
				}
			}
			if _, e := os.Stat(marker); !os.IsNotExist(e) {
				t.Error("verification invoked hidden agent/controller/network command", e)
			}
			verifyNoForbiddenGit(t, trace, f.Root, owned)
			if _, e := os.Stat(trace); e == nil && claimLocalFetch(t, trace) {
				t.Error("verification attempted fetch")
			}
		})
	}
}

func verifyAllEditKinds(t *testing.T, f *verifyFixtureData) {
	t.Helper()
	doc := f.Change.Document
	doc.Metadata["fixes"] = []any{"ALP-006"}
	bases := map[string]any{}
	for _, m := range f.Snapshot.Modules {
		if m.Name == "alpha" {
			for _, sp := range m.Specs {
				if sp.ID == "ALP-001" || sp.ID == "ALP-002" {
					bases[sp.ID] = c.SpecHash(sp)
				}
			}
		}
	}
	if len(bases) != 2 {
		t.Fatal("modified/removed baseline premise")
	}
	doc.Metadata["bases"] = bases
	doc.Sections = append(doc.Sections, c.Section{Name: "Modified", Body: "### ALP-001\nThe target MUST use the revised behavior."}, c.Section{Name: "Removed", Body: "### ALP-002"})
	raw, e := c.RenderMarkdown(doc)
	if e != nil {
		t.Fatal(e)
	}
	verifyReseed(t, f, raw)
	if len(f.Change.Modified) != 1 || len(f.Change.Removed) != 1 || len(f.Change.Fixes) != 1 || len(f.Module.Specs) != 5 {
		t.Fatal("edit-kind inclusive fixture", f.Change, f.Module.Specs)
	}
}
func Test_FLOW043VerifyEveryChangedIDAndRemovedCoverage(t *testing.T) {
	tool := verifyTool(t)
	for _, mode := range []string{"all-fail", "added-pass", "modified-pass", "fix-pass", "removed-reference"} {
		t.Run(mode, func(t *testing.T) {
			f := verifyFixture(t, tool, "evals", "all", "sha1", false)
			verifyAllEditKinds(t, &f)
			id := f.Change.Added[0].ID
			ids := []string{"ALP-001", "ALP-006", id}
			failed := map[string]bool{"ALP-001": true, "ALP-006": true, id: true}
			missing := ""
			switch mode {
			case "added-pass":
				missing = id
			case "modified-pass":
				missing = "ALP-001"
			case "fix-pass":
				missing = "ALP-006"
			case "removed-reference":
				ids = append(ids, "ALP-002")
			}
			if missing != "" {
				delete(failed, missing)
			}
			verifyReport(t, f, "alpha", ids, failed)
			want := verifyManualCoverage(t, f, []string{id, "ALP-001", "ALP-006"}, nil)
			v := verifyInvoke(t, f, "current")
			verifyExpectedCoverage(t, f, []string{"alpha"}, v)
			if len(v.Coverage) == 1 && !reflect.DeepEqual(verifyNormalizeCoverage(t, f, v.Coverage[0]), want) {
				t.Errorf("modified/fix/removed evidence differs %+v vs %+v", v.Coverage[0], want)
			}
			if !reflect.DeepEqual(verifyNormalizeDiagnostics(t, f, v.Diagnostics), want.Diagnostics) {
				t.Errorf("coverage diagnostics differ: got %+v want %+v", v.Diagnostics, want.Diagnostics)
			}
			if mode == "all-fail" {
				verifyNoErrors(t, v.Diagnostics)
			} else {
				if missing != "" {
					verifyCode(t, v, "P1-02", "units/alpha")
				} else {
					verifyCode(t, v, "C12", "units/alpha")
				}
			}
		})
	}
}
func Test_FLOW044VerifyAuditChangedTargetsAndOrder(t *testing.T) {
	tool := verifyTool(t)
	for _, mode := range []string{"passing", "missing-modified", "missing-fix", "findings-order", "old-body"} {
		t.Run(mode, func(t *testing.T) {
			f := verifyFixture(t, tool, "audit", "all", "sha1", false)
			verifyAllEditKinds(t, &f)
			id := f.Change.Added[0].ID
			audit := "No findings. Checked " + id + ", ALP-001, ALP-006."
			want := []w.Finding{}
			switch mode {
			case "missing-modified":
				audit = "No findings. Checked " + id + ", ALP-006."
			case "missing-fix":
				audit = "No findings. Checked " + id + ", ALP-001."
			case "old-body":
				audit = "- ALP-001: \"MUST remain\" — stale body quote"
			case "findings-order":
				lines := []string{"- ALP-006: \"MUST remain\" — fix coverage gap", "- " + id + ": \"MUST be checked\" — added coverage gap", "- ALP-001: \"MUST use the revised behavior\" — modified coverage gap"}
				audit = strings.Join(lines, "\n")
				want = []w.Finding{{ID: "ALP-006", Clause: "MUST remain", Text: lines[0]}, {ID: id, Clause: "MUST be checked", Text: lines[1]}, {ID: "ALP-001", Clause: "MUST use the revised behavior", Text: lines[2]}}
			}
			worktreePut(t, f.Root, f.Path, f.Raw+"\n## Audit\n\n"+audit+"\n")
			v := verifyInvoke(t, f, "current")
			if !reflect.DeepEqual(v.Findings, want) {
				t.Errorf("ordered current-spec findings %+v want %+v", v.Findings, want)
			}
			if mode == "passing" || mode == "findings-order" {
				verifyNoErrors(t, v.Diagnostics)
			} else {
				verifyCode(t, v, "P1-03", f.Path)
			}
			if len(verifyCalls(t, f.Plan)) != 0 {
				t.Error("audit launched hooks")
			}
		})
	}
}

// Normalize only the exact unique report path recorded by this fixture hook.
// Report contents, statuses, IDs, ordering and other paths remain significant.
func verifyNormalizeDiagnostics(t *testing.T, f verifyFixtureData, ds []k.Diagnostic) []k.Diagnostic {
	t.Helper()
	out := append([]k.Diagnostic{}, ds...)
	for _, call := range verifyCalls(t, f.Plan) {
		if call.Report != "" {
			for i := range out {
				out[i].Message = strings.ReplaceAll(out[i].Message, call.Report, "<owned-report>")
			}
		}
	}
	return out
}
func verifyNormalizeCoverage(t *testing.T, f verifyFixtureData, cov k.Coverage) k.Coverage {
	t.Helper()
	cov.Diagnostics = verifyNormalizeDiagnostics(t, f, cov.Diagnostics)
	for _, call := range verifyCalls(t, f.Plan) {
		if call.Report != "" {
			cov.Output = strings.ReplaceAll(cov.Output, call.Report, "<owned-report>")
		}
	}
	return cov
}
