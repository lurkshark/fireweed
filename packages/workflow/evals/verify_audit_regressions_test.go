package evals

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	k "github.com/lurkshark/fireweed/packages/checks"
	c "github.com/lurkshark/fireweed/packages/contracts"
	g "github.com/lurkshark/fireweed/packages/gitstore"
	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

func Test_FLOW041VerifyImplicitSelectionFiltersForeignClaims(t *testing.T) {
	tool := verifyTool(t)
	f := verifyFixture(t, tool, "audit", "none", "sha1", false)
	peer := claimPeer(t, f.Root)
	foreign, e := w.CreateClaim(peer, "foreign", "foreign-run", "HEAD", "Unrelated foreign state")
	if e != nil {
		t.Fatal(e)
	}
	inventory, e := w.Claims(f.Root)
	if e != nil || len(inventory) != 2 || foreign.Worktree == f.Root {
		t.Fatal("foreign ownership premise", inventory, e)
	}
	worktreePut(t, f.Root, f.Path, f.Raw+"\n## Audit\n\nNo findings. Checked "+f.Change.Added[0].ID+".\n")
	verifyReadOnly(t, f, func() {
		v := verifyInvoke(t, f, "")
		verifyNoErrors(t, v.Diagnostics)
		if len(v.Coverage) != 0 || len(v.Findings) != 0 || len(verifyCalls(t, f.Plan)) != 0 {
			t.Error("implicit selection used a foreign claim or ran audit hooks", v)
		}
	})
}

func Test_FLOW042VerifyStaticProjectionValidationBeforeHooks(t *testing.T) {
	tool := verifyTool(t)
	f := verifyFixture(t, tool, "impl", "all", "sha1", false)
	doc := f.Change.Document
	set := doc.Metadata["set"].(map[string]any)
	set["evals"] = []any{"overlap/**"}
	set["disposable"] = []any{"overlap/**"}
	raw, e := c.RenderMarkdown(doc)
	if e != nil {
		t.Fatal(e)
	}
	verifyReseed(t, &f, raw)
	expected := p.Check(verifyInclusive(t, f))
	found := false
	for _, d := range expected {
		if d.Code == "C06" && d.Severity == "error" && strings.Contains(d.Message, "intersect") {
			found = true
		}
	}
	if !found {
		t.Fatal("static effective-project rejection premise", expected)
	}
	verifyReadOnly(t, f, func() {
		v := verifyInvoke(t, f, "current")
		verifyGuardEvidence(t, f, v)
		for _, d := range expected {
			want := k.Diagnostic{Code: d.Code, Severity: d.Severity, Message: d.Message, Path: d.Path}
			have := false
			for _, got := range v.Diagnostics {
				have = have || got == want
			}
			if !have {
				t.Error("lost effective-project diagnostic", want, v.Diagnostics)
			}
		}
		if len(v.Coverage) != 0 || len(verifyCalls(t, f.Plan)) != 0 {
			t.Error("invalid static projection executed hooks")
		}
	})
}

func Test_FLOW046VerifyPreHookWarningsSurviveOtherEvidence(t *testing.T) {
	tool := verifyTool(t)
	f := verifyFixture(t, tool, "impl", "none", "sha1", false)
	verifyPutModuleMetadata(t, f, "extra", "boundaries_opt_out", "Reviewed fixture warning must remain visible")
	verifyRefreshFixture(t, &f)
	expected := p.Check(verifyInclusive(t, f))
	warnings := []k.Diagnostic{}
	for _, d := range expected {
		if d.Code == "D04" && d.Severity == "warning" {
			warnings = append(warnings, k.Diagnostic{Code: d.Code, Severity: d.Severity, Message: d.Message, Path: d.Path})
		}
	}
	if len(warnings) != 1 {
		t.Fatal("independent approved pre-hook warning premise", expected)
	}
	guard := w.CheckWorktree(verifyInclusive(t, f), f.Module, f.Path, f.Claim)
	verifyNoErrors(t, guard.Diagnostics)
	verifyReport(t, f, "alpha", []string{"ALP-001", "ALP-002", "ALP-006", f.Change.Added[0].ID}, map[string]bool{"ALP-001": true})
	mutations, _ := json.Marshal(map[string]string{"private/from-warning-hook": "retained mutation"})
	worktreePut(t, f.Plan, "alpha.mutation", string(mutations))
	v := verifyInvoke(t, f, "current")
	verifyExpectedCoverage(t, f, []string{"alpha"}, v)
	verifyGuardEvidence(t, f, v)
	for _, want := range warnings {
		have := false
		for _, got := range v.Diagnostics {
			have = have || got == want
		}
		if !have {
			t.Error("lost independent pre-hook warning", want, v.Diagnostics)
		}
	}
	verifyCode(t, v, "C14", "units/alpha")
	verifyCode(t, v, "P2-01", "private/from-warning-hook")
	if mustReadVerify(t, filepath.Join(f.Root, "private/from-warning-hook")) != "retained mutation" {
		t.Error("hook mutation repaired")
	}
	verifyDiagnosticOrder(t, v.Diagnostics)
}

func Test_FLOW047VerifyErrorEvidenceHasIndependentStorage(t *testing.T) {
	tool := verifyTool(t)
	f := verifyFixture(t, tool, "impl", "none", "sha1", false)
	verifyReport(t, f, "alpha", []string{"ALP-001", "ALP-002", "ALP-006", f.Change.Added[0].ID}, map[string]bool{"ALP-001": true})
	raw := mustReadVerify(t, filepath.Join(f.Plan, "alpha.xml"))
	raw = strings.Replace(raw, "<failure message=\"required behavior absent\"/>", "<error message=\"error evidence\">retained error text</error>", 1)
	worktreePut(t, f.Plan, "alpha.xml", raw)
	control := verifyManualCoverage(t, f, nil, nil)
	index := -1
	for i, tc := range control.Tests {
		if len(tc.Errors) > 0 {
			index = i
			break
		}
	}
	if index < 0 {
		t.Fatal("nonempty public Errors premise", control)
	}
	a := verifyInvoke(t, f, "current")
	b := verifyInvoke(t, f, "current")
	verifyExpectedCoverage(t, f, []string{"alpha"}, a)
	verifyExpectedCoverage(t, f, []string{"alpha"}, b)
	if len(a.Coverage) != 1 || index >= len(a.Coverage[0].Tests) || len(a.Coverage[0].Tests[index].Errors) == 0 {
		t.Fatal("lost actual Errors evidence", a)
	}
	saved := policyJSON(t, b)
	a.Coverage[0].Tests[index].Errors[0].Text = "mutated Errors text"
	a.Coverage[0].Tests[index].Errors[0].Message = "mutated Errors message"
	if saved != policyJSON(t, b) {
		t.Error("Errors backing storage shared between results")
	}
	again := verifyInvoke(t, f, "current")
	verifyExpectedCoverage(t, f, []string{"alpha"}, again)
	if len(again.Coverage) != 1 {
		t.Fatal("subsequent Errors coverage missing")
	}
	if !reflect.DeepEqual(verifyNormalizeCoverage(t, f, again.Coverage[0]), control) {
		t.Error("Errors mutation leaked to subsequent result", again.Coverage[0], control)
	}
}

func Test_FLOW047VerifyEmptyCollectionsHaveIndependentStorage(t *testing.T) {
	tool := verifyTool(t)
	f := verifyFixture(t, tool, "audit", "none", "sha1", false)
	worktreePut(t, f.Root, f.Path, f.Raw+"\n## Audit\n\nNo findings. Checked "+f.Change.Added[0].ID+".\n")
	verifyRefreshFixture(t, &f)
	guard := w.CheckWorktree(verifyInclusive(t, f), f.Module, f.Path, f.Claim)
	if len(guard.Edits) != 0 || len(guard.Diagnostics) != 0 {
		t.Fatal("empty guard premise", guard)
	}
	a := verifyInvoke(t, f, "current")
	b := verifyInvoke(t, f, "current")
	verifyNoErrors(t, a.Diagnostics)
	verifyNoErrors(t, b.Diagnostics)
	if len(a.Edits)+len(a.Diagnostics)+len(a.Coverage)+len(a.Findings)+len(b.Edits)+len(b.Diagnostics)+len(b.Coverage)+len(b.Findings) != 0 {
		t.Fatal("empty collection premise", a, b)
	}
	a.Diagnostics = append(a.Diagnostics, k.Diagnostic{Message: "first"})
	b.Diagnostics = append(b.Diagnostics, k.Diagnostic{Message: "second"})
	a.Coverage = append(a.Coverage, k.Coverage{Module: "first"})
	b.Coverage = append(b.Coverage, k.Coverage{Module: "second"})
	a.Edits = append(a.Edits, g.Edit{Path: "first"})
	b.Edits = append(b.Edits, g.Edit{Path: "second"})
	a.Findings = append(a.Findings, w.Finding{Text: "first"})
	b.Findings = append(b.Findings, w.Finding{Text: "second"})
	if a.Diagnostics[0].Message != "first" || b.Diagnostics[0].Message != "second" || a.Coverage[0].Module != "first" || b.Coverage[0].Module != "second" || a.Edits[0].Path != "first" || b.Edits[0].Path != "second" || a.Findings[0].Text != "first" || b.Findings[0].Text != "second" {
		t.Error("two-sided append exposes shared empty backing storage", a, b)
	}
	a.Diagnostics[0].Message = "changed"
	a.Coverage[0].Module = "changed"
	a.Edits[0].Path = "changed"
	a.Findings[0].Text = "changed"
	if b.Diagnostics[0].Message != "second" || b.Coverage[0].Module != "second" || b.Edits[0].Path != "second" || b.Findings[0].Text != "second" {
		t.Error("two-sided mutation changed second result")
	}
}

// The parent supplies stdin only after the public probe has completed startup.
// Both streams are captured, and unread pipe bytes remain observable after exit.
const verifyStdioSource = `package main
import("context";"encoding/json";"fmt";"os";"path/filepath";"time";w "github.com/lurkshark/fireweed/packages/workflow")
func main(){mode,root,dir:=os.Args[1],os.Args[2],os.Args[3];os.WriteFile(filepath.Join(dir,"ready"),[]byte("ready"),0600);for{if _,e:=os.Stat(filepath.Join(dir,"gate"));e==nil{break};time.Sleep(time.Millisecond)};switch mode{case "stdin-control":b:=make([]byte,1);os.Stdin.Read(b);return;case "stdout-control":fmt.Print("live prompt");return;case "stderr-control":fmt.Fprint(os.Stderr,"live prompt");return};v,e:=w.Verify(context.Background(),root,"current");message:="";if e!=nil{message=e.Error()};record:=struct{Verification w.Verification;Findings []w.Finding;Error string}{v,v.Findings,message};raw,_:=json.Marshal(record);if os.WriteFile(filepath.Join(dir,"result"),raw,0600)!=nil{os.Exit(92)}}
`

type verifyStdioRecord struct {
	Verification w.Verification
	Findings     []w.Finding
	Error        string
}

func verifyStdioProbe(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	src := filepath.Join(dir, "probe.go")
	worktreePut(t, dir, "probe.go", verifyStdioSource)
	binary := filepath.Join(dir, "probe")
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	cmd := exec.CommandContext(ctx, "go", "build", "-o", binary, src)
	cmd.Env = append(os.Environ(), "GOPROXY=off", "GOTOOLCHAIN=local")
	if out, e := cmd.CombinedOutput(); e != nil {
		t.Fatalf("public stdio probe build: %v %s", e, out)
	}
	return binary
}
func verifyObserveStdio(t *testing.T, binary, root, mode string) (string, string, verifyStdioRecord) {
	t.Helper()
	dir := t.TempDir()
	reader, writer, e := os.Pipe()
	if e != nil {
		t.Fatal(e)
	}
	defer reader.Close()
	defer writer.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, binary, mode, root, dir)
	cmd.Dir = root
	cmd.Stdin = reader
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	if e = cmd.Start(); e != nil {
		t.Fatal(e)
	}
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	until := time.Now().Add(5 * time.Second)
	for {
		if _, e := os.Stat(filepath.Join(dir, "ready")); e == nil {
			break
		}
		select {
		case e := <-done:
			t.Fatalf("stdio observer ended before startup: %v %s", e, output.String())
		default:
		}
		if time.Now().After(until) {
			t.Fatal("stdio startup marker absent")
		}
		time.Sleep(time.Millisecond)
	}
	payload := "unchanged stdin bytes\n"
	if _, e = io.WriteString(writer, payload); e != nil {
		t.Fatal(e)
	}
	if e = writer.Close(); e != nil {
		t.Fatal(e)
	}
	worktreePut(t, dir, "gate", "run")
	if e = <-done; e != nil {
		t.Fatalf("public stdio probe failed: %v %s", e, output.String())
	}
	remaining, e := io.ReadAll(reader)
	if e != nil {
		t.Fatal(e)
	}
	result := verifyStdioRecord{}
	if mode == "verify" {
		if e = json.Unmarshal([]byte(mustReadVerify(t, filepath.Join(dir, "result"))), &result); e != nil {
			t.Fatal(e)
		}
		result.Verification.Findings = result.Findings
	}
	return output.String(), string(remaining), result
}
func Test_FLOW047VerifyNoninteractiveStdio(t *testing.T) {
	binary := verifyStdioProbe(t)
	controlRoot := claimRepo(t, "sha1")
	for _, mode := range []string{"stdin-control", "stdout-control", "stderr-control"} {
		t.Run(mode, func(t *testing.T) {
			out, remaining, _ := verifyObserveStdio(t, binary, controlRoot, mode)
			if mode == "stdin-control" {
				if remaining != "nchanged stdin bytes\n" || out != "" {
					t.Fatal("ignored stdin read observer not live", fmt.Sprintf("%q %q", remaining, out))
				}
			} else if out != "live prompt" || remaining != "unchanged stdin bytes\n" {
				t.Fatal("prompt stream observer not live", fmt.Sprintf("%q %q", remaining, out))
			}
		})
	}
	tool := verifyTool(t)
	for _, rejected := range []bool{false, true} {
		t.Run(fmt.Sprint(rejected), func(t *testing.T) {
			f := verifyFixture(t, tool, "impl", "none", "sha1", false)
			if rejected {
				worktreePut(t, f.Root, "private/forbidden", "structured rejection")
			}
			verifyReadOnly(t, f, func() {
				out, remaining, result := verifyObserveStdio(t, binary, f.Root, "verify")
				if out != "" || remaining != "unchanged stdin bytes\n" {
					t.Error("Verify read stdin or emitted a prompt", fmt.Sprintf("%q %q", remaining, out))
				}
				if result.Error != "" {
					t.Error("unexpected Verify error", result.Error)
				}
				v := result.Verification
				verifyShape(t, f, v)
				if rejected {
					verifyCode(t, v, "P2-01", "private/forbidden")
					if len(v.Coverage) != 0 || len(verifyCalls(t, f.Plan)) != 0 {
						t.Error("structured rejection executed hooks")
					}
				} else {
					verifyNoErrors(t, v.Diagnostics)
					verifyExpectedCoverage(t, f, []string{"alpha"}, v)
				}
			})
		})
	}
}

// A corrective action is semantic guidance, not a prescribed verb vocabulary.
// Keep the older policy fixtures unchanged while calibrating Verify's oracle.
func verifyCorrective(message string) bool {
	if policyCorrective(message) {
		return true
	}
	lower := strings.ToLower(message)
	for _, action := range []string{"add ", "configure ", "declare ", "implement ", "repair "} {
		if strings.Contains(lower, action) {
			return true
		}
	}
	return false
}

func Test_FLOW043VerifyCorrectiveGuidanceControls(t *testing.T) {
	for _, message := range []string{
		"ALP-001 has failing evals; implement its required behavior without weakening the tests.",
		"Cannot complete eval hook: context deadline exceeded. Repair hook configuration or cancellation before trusting its report.",
		"Cannot read the fresh JUnit report. Configure the eval hook to write the supplied report path.",
		"Dependency missing is missing; declare an existing module.",
		"ALP-007 must have an executed failing eval before implementation; add a regression that fails against existing behavior.",
		"Restore the approved state.", "Use the queued slug.", "Remove the unsupported field.", "Resolve Questions before retrying.",
	} {
		t.Run(message, func(t *testing.T) {
			if !verifyCorrective(message) {
				t.Error("valid corrective guidance rejected")
			}
		})
	}
	for _, message := range []string{"", " ", "error", "C11 failed", "ALP-001 missing evidence"} {
		t.Run("reject-"+message, func(t *testing.T) {
			if verifyCorrective(message) {
				t.Error("missing corrective guidance accepted")
			}
		})
	}
}

func Test_FLOW043VerifyLaunchReportObserverControl(t *testing.T) {
	tool := verifyTool(t)
	f := verifyFixture(t, tool, "evals", "all", "sha1", false)
	doc := f.Change.Document
	doc.Metadata["set"].(map[string]any)["hooks"] = map[string]any{"eval": verifyCommand(tool, "observe", f.Plan) + " && " + claimQuote(filepath.Join(f.Plan, "missing-hook-executable"))}
	raw, err := c.RenderMarkdown(doc)
	if err != nil {
		t.Fatal(err)
	}
	verifyReseed(t, &f, raw)
	snapshot := verifyInclusive(t, f)
	a := k.Evaluate(context.Background(), snapshot, f.Module, []string{f.Change.Added[0].ID}, nil)
	b := k.Evaluate(context.Background(), snapshot, f.Module, []string{f.Change.Added[0].ID}, nil)
	calls := verifyCalls(t, f.Plan)
	if len(calls) != 2 {
		t.Fatalf("missing exact report observations: %+v", calls)
	}
	if calls[0].Report == calls[1].Report {
		t.Fatal("public fresh report identities were reused")
	}
	for _, call := range calls {
		if call.Kind != "observe" || !filepath.IsAbs(call.Report) {
			t.Fatal("invalid observation", call)
		}
		if _, err := os.Stat(filepath.Dir(call.Report)); !os.IsNotExist(err) {
			t.Fatal("owned report storage survived", call.Report, err)
		}
	}
	for _, cov := range []k.Coverage{a, b} {
		if cov.HookExit != 127 || len(cov.Tests) != 0 || len(cov.BySpec) != 0 || !strings.Contains(cov.Output, "missing-hook-executable") || len(cov.Diagnostics) != 1 || cov.Diagnostics[0].Code != "C11" {
			t.Fatal("missing-executable public failure control", cov)
		}
	}
	normalizedA, normalizedB := verifyNormalizeCoverage(t, f, a), verifyNormalizeCoverage(t, f, b)
	if !reflect.DeepEqual(normalizedA, normalizedB) {
		t.Fatal("exact observed-report normalization changed semantic evidence", normalizedA, normalizedB)
	}
	// Evidence outside the observed report identity must remain significant.
	changed := b
	changed.Diagnostics = append([]k.Diagnostic{}, b.Diagnostics...)
	changed.Diagnostics[0].Message += " distinct diagnostic"
	if reflect.DeepEqual(normalizedA, verifyNormalizeCoverage(t, f, changed)) {
		t.Fatal("normalization erased meaningful diagnostic text")
	}
}
