package evals

import (
	"context"
	"encoding/json"
	"fmt"
	checks "github.com/lurkshark/fireweed/packages/checks"
	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

type cliVerifyCase struct {
	cliStartCase
	Plan, Raw string
	Claim     w.Claim
	Outside   []string
}
type cliVerifyCall struct{ Kind, Report, Module, Dir, Bin, Root string }

func cliVerifyFixture(t *testing.T, phase, format, policy string, linked, warning bool) cliVerifyCase {
	t.Helper()
	f := cliStartFixture(t, format, "spec", phase, linked, true, false)
	plan := filepath.Join(t.TempDir(), "verify ' plan")
	if e := os.Mkdir(plan, 0755); e != nil {
		t.Fatal(e)
	}
	plan, e := filepath.EvalSymlinks(plan)
	if e != nil {
		t.Fatal(e)
	}
	eval := "sh hooks/verify.sh {module} {report} " + coverageQuote(plan)
	boundary := "sh hooks/verify.sh boundary {report} " + coverageQuote(plan)
	a, _ := json.Marshal(eval)
	b, _ := json.Marshal(boundary)
	cfg := "version: 1\nfireweed: '>=0.4, <0.5'\nbase: main\nmodules: [units/*]\nmodule_defaults:\n  interface: [api.go]\n  evals: [evals/**]\n  disposable: [internal/**]\nhooks:\n  eval: " + string(a) + "\n  boundaries: " + string(b) + "\n  setup: touch MUST-NOT-RUN\nprotected: [private/**]\nagent: {command: [agent-trap]}\ntimeout: 2s\nverify_dependents: " + policy + "\n"
	projectPut(t, f.Root, "fireweed.yaml", cfg)
	projectModule(t, f.Root, "units/gamma", "gamma", "GAM", "depends_on: [alpha]\n")
	projectModule(t, f.Root, "units/delta", "delta", "DEL", "depends_on: [gamma]\n")
	delta, e := os.ReadFile(filepath.Join(f.Root, "units/delta/SPEC.md"))
	if e != nil {
		t.Fatal(e)
	}
	projectPut(t, f.Root, "units/delta/SPEC.md", strings.Replace(string(delta), "DEL-001 [review]", "DEL-001", 1))
	raw, e := os.ReadFile(filepath.Join(f.Root, "units/gamma/SPEC.md"))
	if e != nil {
		t.Fatal(e)
	}
	projectPut(t, f.Root, "units/gamma/SPEC.md", strings.Replace(string(raw), "GAM-001 [review]", "GAM-001", 1))
	if warning {
		projectModule(t, f.Root, "units/b", "beta", "BET", "boundaries_opt_out: reviewed CLI fixture\n")
	}
	script := `#!/bin/sh
set -eu
role=$1;report=$2;plan=$3
printf '%s\t%s\t%s\t%s\t%s\t%s\n' "$role" "$report" "$FIREWEED_MODULE" "$FIREWEED_MODULE_DIR" "$FIREWEED_BIN" "$(pwd -P)" >> "$plan/calls"
if [ "$role" = boundary ]; then
 if [ -f "$plan/boundary.output" ];then cat "$plan/boundary.output";fi
 if [ -f "$plan/boundary.fail" ];then exit 1;fi
 exit 0
fi
mode=$(cat "$plan/$role.mode")
case "$mode" in
 missing) ;;
 malformed) printf '<testsuite>' > "$report" ;;
 *) cat "$plan/$role.xml" > "$report" ;;
esac
printf '%s-out\n' "$role"
printf '%s-err\n' "$role" >&2
if [ -f "$plan/$role.mutate" ];then cp "$plan/$role.mutate" private/from-hook;fi
if [ -f "$plan/$role.head" ];then git -c core.hooksPath=/dev/null commit --allow-empty -qm 'configured verification hook';fi
if [ -f "$plan/$role.release" ];then git update-ref -d refs/fireweed/claims/current;fi
if [ -f "$plan/$role.replace" ];then git update-ref refs/fireweed/claims/current "$(cat "$plan/$role.replace")";fi
if [ "$mode" = timeout ];then sleep 2;printf late > "$plan/LATE";fi
if [ "$mode" = fail ];then exit 1;fi
exit 0
`
	projectPut(t, f.Root, "hooks/verify.sh", script)
	gitFixtureCommand(t, f.Root, "add", ".")
	gitFixtureCommand(t, f.Root, "commit", "-qm", "Approved CLI verification inputs")
	f.Head = strings.TrimSpace(gitFixtureCommand(t, f.Root, "rev-parse", "HEAD"))
	raw, e = os.ReadFile(filepath.Join(f.Root, f.Path))
	if e != nil {
		t.Fatal(e)
	}
	claim, e := w.CreateClaim(f.Root, "current", "retained-cli-run", "HEAD", string(raw))
	if e != nil {
		t.Fatal(e)
	}
	out := cliVerifyCase{cliStartCase: f, Plan: plan, Raw: string(raw), Claim: claim}
	cliVerifyReport(t, out, "alpha", phase == "evals")
	cliVerifyReport(t, out, "gamma", false)
	cliVerifyReport(t, out, "delta", false)
	cliVerifyReport(t, out, "beta", false)
	if phase == "audit" {
		projectPut(t, f.Root, f.Path, string(raw)+"\n## Audit\n\nNo findings. Checked ALP-002.\n")
	}
	v, e := w.Verify(context.Background(), f.Where, "current")
	if e != nil || cliVerifyExit(v) != 0 {
		t.Fatalf("invalid public verification fixture: %v %+v", e, v)
	}
	cliVerifyClear(t, out)
	return out
}
func cliVerifyReport(t *testing.T, f cliVerifyCase, module string, failed bool) {
	t.Helper()
	ids := []string{"ALP001", "ALP002"}
	if module == "delta" {
		ids = []string{"DEL001"}
	}
	if module == "gamma" {
		ids = []string{"GAM001"}
	}
	if module == "beta" {
		ids = []string{"BET001"}
	}
	var s strings.Builder
	s.WriteString("<testsuite>")
	for i, id := range ids {
		s.WriteString("<testcase name=\"" + id + "\" classname=\"cli.verify\">")
		if failed && i == len(ids)-1 {
			s.WriteString("<failure message=\"required behavior absent\">retained failure text</failure>")
		}
		s.WriteString("</testcase>")
	}
	s.WriteString("</testsuite>")
	projectPut(t, f.Plan, module+".xml", s.String())
	mode := "pass"
	if failed {
		mode = "fail"
	}
	projectPut(t, f.Plan, module+".mode", mode)
}
func cliVerifyCalls(t *testing.T, f cliVerifyCase) []cliVerifyCall {
	t.Helper()
	data, e := os.ReadFile(filepath.Join(f.Plan, "calls"))
	if os.IsNotExist(e) {
		return []cliVerifyCall{}
	}
	if e != nil {
		t.Fatal(e)
	}
	out := []cliVerifyCall{}
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		p := strings.Split(line, "\t")
		if len(p) != 6 {
			t.Fatalf("bad observer record %q", line)
		}
		out = append(out, cliVerifyCall{p[0], p[1], p[2], p[3], p[4], p[5]})
	}
	return out
}
func cliVerifyClear(t *testing.T, f cliVerifyCase) {
	t.Helper()
	if e := os.Remove(filepath.Join(f.Plan, "calls")); e != nil && !os.IsNotExist(e) {
		t.Fatal(e)
	}
}
func cliVerifyExit(v w.Verification) int {
	for _, d := range v.Diagnostics {
		if d.Severity == "error" {
			return 1
		}
	}
	return 0
}
func cliVerifyReports(t *testing.T, f cliVerifyCase) []string {
	t.Helper()
	out := []string{}
	for _, call := range cliVerifyCalls(t, f) {
		if call.Report != "" {
			out = append(out, call.Report)
		}
	}
	return out
}
func cliVerifyNormalize(value any, reports []string) any {
	switch v := value.(type) {
	case string:
		for _, p := range reports {
			v = strings.ReplaceAll(v, p, "<owned-report>")
		}
		return v
	case []any:
		out := make([]any, len(v))
		for i, x := range v {
			out[i] = cliVerifyNormalize(x, reports)
		}
		return out
	case map[string]any:
		out := map[string]any{}
		for k, x := range v {
			out[k] = cliVerifyNormalize(x, reports)
		}
		return out
	default:
		return value
	}
}
func cliVerifyWire(t *testing.T, v w.Verification, reports []string) any {
	t.Helper()
	data := canonical(t, v).(map[string]any)
	edits := []any{}
	for _, e := range v.Edits {
		item := map[string]any{"status": e.Status, "path": e.Path}
		if e.OldPath != "" {
			item["old_path"] = e.OldPath
		}
		edits = append(edits, item)
	}
	data["edits"] = edits
	return cliVerifyNormalize(data, reports)
}
func cliVerifyPlainLeaks(stderr string, want w.Verification) bool {
	for _, cov := range want.Coverage {
		for _, suffix := range []string{"-out", "-err"} {
			if strings.Contains(stderr, cov.Module+suffix) {
				return true
			}
		}
	}
	return false
}
func cliVerifyPlainStreams(t *testing.T, stderr string, want w.Verification, actualReports, reports []string) {
	t.Helper()
	for _, cov := range want.Coverage {
		for _, suffix := range []string{"-out", "-err"} {
			if !strings.Contains(cov.Output, cov.Module+suffix) {
				t.Fatal("public hook output marker premise absent", cov)
			}
		}
	}
	if cliVerifyPlainLeaks(stderr, want) {
		t.Error("plain stderr leaked captured hook stdout/stderr", stderr)
	}
	if len(want.Diagnostics) == 0 && stderr != "" {
		t.Error("plain verification added stderr", stderr)
	}
	if e := cliVerifyPlainDiagnosticText(stderr, want.Diagnostics, actualReports, reports); e != nil {
		t.Error(e)
	}
}
func Test_APP026VerifyPlainStreams(t *testing.T) {
	t.Run("control", func(t *testing.T) {
		f := cliVerifyFixture(t, "impl", "sha1", "all", false, false)
		want, _ := cliVerifyExpected(t, f, "current")
		if len(want.Coverage) != 3 || cliVerifyPlainLeaks("diagnostics only\n", want) {
			t.Fatal("plain stream control premise", want)
		}
		cliVerifyPlainStreams(t, "", want, nil, nil)
		for _, cov := range want.Coverage {
			for _, suffix := range []string{"-out", "-err"} {
				if !cliVerifyPlainLeaks("diagnostics only\n"+cov.Module+suffix+"\n", want) {
					t.Fatal("plain stream observer missed individual hook stream", cov.Module, suffix)
				}
			}
			if !cliVerifyPlainLeaks("diagnostics only\n"+cov.Output, want) {
				t.Fatal("plain stream observer missed copied Coverage.Output", cov)
			}
		}
		cliVerifyReport(t, f, "alpha", true)
		failed, reports := cliVerifyExpected(t, f, "current")
		if len(failed.Coverage) == 0 || len(failed.Diagnostics) == 0 {
			t.Fatal("plain diagnostic control premise", failed)
		}
		text := ""
		for _, d := range failed.Diagnostics {
			text += fmt.Sprintf("%s %s (%s): %s\n", d.Severity, d.Code, d.Path, d.Message)
		}
		if e := cliVerifyPlainDiagnosticText(text, failed.Diagnostics, reports, reports); e != nil {
			t.Fatal("public plain diagnostic calibration", e)
		}
		for _, message := range []string{"C05: invalid exact selector", "P02: held ownership missing", "C01: controlled discovery failure"} {
			code := strings.SplitN(message, ":", 2)[0]
			ds := []checks.Diagnostic{{Code: code, Severity: "error", Message: message}}
			for _, plain := range []string{"error " + message + "\n", "error " + code + ": " + message + "\n"} {
				if e := cliVerifyPlainDiagnosticText(plain, ds, nil, nil); e != nil {
					t.Fatal("fatal plain presentation calibration", e)
				}
				if cliVerifyPlainDiagnosticText(plain+plain, ds, nil, nil) == nil {
					t.Fatal("duplicated fatal diagnostic escaped observer")
				}
			}
		}
		for _, extra := range []string{"prompt> ", text} {
			if cliVerifyPlainDiagnosticText(text+extra, failed.Diagnostics, reports, reports) == nil {
				t.Fatal("plain diagnostic observer missed appended text", extra)
			}
		}

	})
}
func cliVerifyExpected(t *testing.T, f cliVerifyCase, slug string) (w.Verification, []string) {
	t.Helper()
	v, e := w.Verify(context.Background(), f.Where, slug)
	if e != nil {
		t.Fatalf("expected selected public data: %v", e)
	}
	reports := cliVerifyReports(t, f)
	cliVerifyClear(t, f)
	return v, reports
}
func cliVerifyObserve(t *testing.T, f cliVerifyCase, expected []string, binary string) {
	t.Helper()
	calls := cliVerifyCalls(t, f)
	names := []string{}
	for _, call := range calls {
		names = append(names, call.Kind)
		if call.Bin != binary || call.Root != f.Root {
			t.Errorf("wrong subject hook identity %+v want %s", call, binary)
		}
		if call.Kind == "boundary" {
			if call.Module != "" || call.Dir != "" || call.Report != "" {
				t.Error("boundary received selected-module identity", call)
			}
		} else {
			dir := "units/" + call.Kind
			if call.Kind == "alpha" {
				dir = "units/a"
			}
			if call.Kind == "beta" {
				dir = "units/b"
			}
			if call.Module != call.Kind || call.Dir != dir || !filepath.IsAbs(call.Report) {
				t.Error("eval identity", call)
			}
			if _, e := os.Stat(filepath.Dir(call.Report)); !os.IsNotExist(e) {
				t.Error("owned report storage survived", call.Report, e)
			}
		}
	}
	if !reflect.DeepEqual(names, expected) {
		t.Errorf("hook order/count %+v want %+v", names, expected)
	}
}
func cliVerifyObserveVerification(t *testing.T, f cliVerifyCase, want w.Verification, binary string) {
	t.Helper()
	names := []string{}
	for _, cov := range want.Coverage {
		names = append(names, cov.Module)
	}
	if want.Phase == "impl" && len(want.Coverage) > 0 {
		names = append(names, "boundary")
	}
	cliVerifyObserve(t, f, names, binary)
}
func cliVerifyJSON(t *testing.T, f cliVerifyCase, args []string, want w.Verification, reports []string) result {
	t.Helper()
	got := invoke(t, f.Where, args...)
	x := decode(t, got, "verify")
	actualReports := cliVerifyReports(t, f)
	cliVerifyObserveVerification(t, f, want, subject)
	if got.code != cliVerifyExit(want) || !reflect.DeepEqual(cliVerifyNormalize(decodeData(t, x.Data), actualReports), cliVerifyWire(t, want, reports)) {
		t.Errorf("verification data/exit %+v\nwant %+v", got, cliVerifyWire(t, want, reports))
	}
	var envelope map[string]any
	if e := json.Unmarshal([]byte(got.out), &envelope); e != nil {
		t.Fatal(e)
	}
	if !reflect.DeepEqual(cliVerifyNormalize(envelope["diagnostics"], actualReports), cliVerifyNormalize(canonical(t, want.Diagnostics), reports)) {
		t.Error("envelope dropped/reordered diagnostics", envelope["diagnostics"], want.Diagnostics)
	}
	return got
}
func cliVerifyHookNames(phase, policy string) []string {
	if phase == "audit" {
		return []string{}
	}
	if phase == "evals" {
		return []string{"alpha"}
	}
	if policy == "none" {
		return []string{"alpha", "boundary"}
	}
	if policy == "direct" {
		return []string{"alpha", "gamma", "boundary"}
	}
	return []string{"alpha", "delta", "gamma", "boundary"}
}
func cliVerifyOutside(t *testing.T, f cliVerifyCase) map[string]map[string]string {
	t.Helper()
	out := map[string]map[string]string{}
	for _, path := range f.Outside {
		out[path] = queueFiles(t, filepath.Dir(path))
	}
	return out
}

func cliVerifyReadonly(t *testing.T, f cliVerifyCase, run func()) {
	t.Helper()
	scratch, e := filepath.EvalSymlinks(t.TempDir())
	if e != nil {
		t.Fatal(e)
	}
	t.Setenv("TMPDIR", scratch)
	t.Setenv("TMP", scratch)
	t.Setenv("TEMP", scratch)
	proxy := cliStartLocalGit(t, filepath.Join(scratch, "forbidden-git"))
	t.Setenv("PATH", proxy+string(os.PathListSeparator)+os.Getenv("PATH"))
	traps := t.TempDir()
	marker := filepath.Join(scratch, "forbidden-command")
	if runtime.GOOS != "windows" {
		for _, name := range []string{"agent-trap", "codex", "claude", "fireweed", "fireweed-bootstrap", "curl", "wget", "ssh"} {
			p := filepath.Join(traps, name)
			if e := os.WriteFile(p, []byte("#!/bin/sh\nprintf denied > "+coverageQuote(marker)+"\nexit 94\n"), 0755); e != nil {
				t.Fatal(e)
			}
			cmd := exec.Command(p)
			if cmd.Run() == nil {
				t.Fatal("command trap control did not fail")
			}
			if _, e := os.Stat(marker); e != nil {
				t.Fatal("command trap was not observed")
			}
			os.Remove(marker)
		}
		t.Setenv("PATH", traps+string(os.PathListSeparator)+os.Getenv("PATH"))
	}
	t.Setenv("FIREWEED_BIN", filepath.Join(traps, "fireweed-bootstrap"))
	gitFixtureCommand(t, f.Root, "--version")
	common := strings.TrimSpace(gitFixtureCommand(t, f.Root, "rev-parse", "--git-common-dir"))
	if !filepath.IsAbs(common) {
		common = filepath.Join(f.Root, common)
	}
	trace := filepath.Join(t.TempDir(), "actual-git-trace")
	t.Setenv("GIT_TRACE2_EVENT", trace)
	t.Setenv("GIT_TRACE2_ENV_VARS", "GIT_DIR,GIT_INDEX_FILE")
	cliVerifyGitRecorder(t, trace)
	files, meta, temp := queueFiles(t, f.Root), queueFiles(t, common), queueFiles(t, scratch)
	outside := cliVerifyOutside(t, f)
	cwd, e := os.Getwd()
	if e != nil {
		t.Fatal(e)
	}
	env := os.Environ()
	run()
	cliVerifyNoForbiddenGit(t, trace, f.Root, scratch)
	after, e := os.Getwd()
	if e != nil || cwd != after || !reflect.DeepEqual(env, os.Environ()) || !reflect.DeepEqual(files, queueFiles(t, f.Root)) || !reflect.DeepEqual(meta, queueFiles(t, common)) || !reflect.DeepEqual(temp, queueFiles(t, scratch)) || !reflect.DeepEqual(outside, cliVerifyOutside(t, f)) {
		t.Error("verify changed original files/ignored/index/HEAD/refs/common Git/temp or parent state")
	}
}

func Test_APP025_APP026_APP028VerifyExecutablePhases(t *testing.T) {
	for _, phase := range []string{"evals", "audit", "impl"} {
		for _, format := range []string{"sha1", "sha256"} {
			for _, linked := range []bool{false, true} {
				t.Run(fmt.Sprintf("%s/%s/linked-%v", phase, format, linked), func(t *testing.T) {
					f := cliVerifyFixture(t, phase, format, "all", linked, false)
					slug := "current"
					args := []string{"verify", slug, "--json"}
					if linked {
						slug = ""
						args = []string{"--json", "verify"}
					}
					want, reports := cliVerifyExpected(t, f, slug)
					cliVerifyReadonly(t, f, func() {
						cliVerifyJSON(t, f, args, want, reports)
						cliVerifyObserve(t, f, cliVerifyHookNames(phase, "all"), subject)
					})
					cliVerifyClear(t, f)
					plainArgs := []string{"verify"}
					if slug != "" {
						plainArgs = append(plainArgs, slug)
					}
					var got result
					cliVerifyReadonly(t, f, func() { got = invoke(t, f.Where, plainArgs...) })
					actualReports := cliVerifyReports(t, f)
					if got.code != 0 || got.err != "" || !reflect.DeepEqual(cliVerifyNormalize(decodeData(t, json.RawMessage(got.out)), actualReports), cliVerifyWire(t, want, reports)) {
						t.Error("plain complete evidence", got)
					}
					coveragePretty(t, got.out)
					cliVerifyPlainStreams(t, got.err, want, actualReports, reports)
					cliVerifyObserve(t, f, cliVerifyHookNames(phase, "all"), subject)
				})
			}
		}
	}
}
func Test_APP025_APP026VerifyEvidenceAndPolicies(t *testing.T) {
	for _, mode := range []string{"passing-failfirst-rejected", "already-satisfied", "own-failure", "dependent-failure", "transitive-failure", "direct", "none", "warning", "missing-report", "malformed-report", "rich-report", "audit-multiple", "audit-finding", "audit-untestable", "audit-incomplete", "rename", "worker-config", "worker-valid-config", "worker-module-hooks", "blocked", "forbidden-eval", "unsafe-current", "malformed-unrelated"} {
		t.Run(mode, func(t *testing.T) {
			phase := "impl"
			if mode == "passing-failfirst-rejected" || mode == "already-satisfied" {
				phase = "evals"
			}
			if strings.HasPrefix(mode, "audit-") {
				phase = "audit"
			}
			policy := "all"
			if mode == "direct" || mode == "none" {
				policy = mode
			}
			f := cliVerifyFixture(t, phase, "sha1", policy, false, mode == "warning")
			switch mode {
			case "passing-failfirst-rejected":
				cliVerifyReport(t, f, "alpha", false)
			case "already-satisfied":
				changed := strings.Replace(f.Raw, "module: alpha\n", "module: alpha\nalready_satisfied: [ALP-002]\n", 1)
				projectPut(t, f.Root, f.Path, changed)
				approved, err := w.UpdateClaim(f.Root, f.Claim, f.Head, changed)
				if err != nil {
					t.Fatal(err)
				}
				f.Claim, f.Raw = approved, changed
				cliVerifyReport(t, f, "alpha", false)
			case "transitive-failure":
				cliVerifyReport(t, f, "delta", true)
			case "own-failure":
				cliVerifyReport(t, f, "alpha", true)
			case "dependent-failure":
				cliVerifyReport(t, f, "gamma", true)
			case "missing-report":
				projectPut(t, f.Plan, "alpha.mode", "missing")
			case "rich-report":
				projectPut(t, f.Plan, "alpha.xml", `<testsuite><testcase name="ALP001 skipped" classname="rich.first"><skipped message="retained skip reason">retained skip text</skipped></testcase><testcase name="ALP002 first" classname="rich.second"><error message="retained error reason">retained error text</error><failure message="retained failure reason">retained failure text</failure></testcase><testcase name="ALP002 second" classname="rich.third"><failure message="second failure">second text</failure></testcase></testsuite>`)
				projectPut(t, f.Plan, "alpha.mode", "fail")
			case "audit-multiple":
				projectPut(t, f.Root, f.Path, f.Raw+"\n## Audit\n\n- ALP-002: \"MUST retain\" — first assertion missing\n- ALP-002: \"new behavior\" — second assertion missing\n")
			case "malformed-report":
				projectPut(t, f.Plan, "alpha.mode", "malformed")
			case "audit-finding", "audit-untestable":
				marker := ""
				if mode == "audit-untestable" {
					marker = "[untestable] "
				}
				projectPut(t, f.Root, f.Path, f.Raw+"\n## Audit\n\n- "+marker+"ALP-002: \"MUST retain\" — missing edge assertion\n")
			case "audit-incomplete":
				projectPut(t, f.Root, f.Path, f.Raw+"\n## Audit\n\nNo findings. Checked ALP-001.\n")
			case "rename":
				gitFixtureCommand(t, f.Root, "mv", "units/a/internal/private.txt", "units/a/internal/renamed.txt")
			case "worker-config":
				projectPut(t, f.Root, "fireweed.yaml", "invalid: [current worker metadata")
			case "worker-valid-config":
				raw, e := os.ReadFile(filepath.Join(f.Root, "fireweed.yaml"))
				if e != nil {
					t.Fatal(e)
				}
				lines := strings.Split(string(raw), "\n")
				for i, line := range lines {
					if strings.HasPrefix(line, "  eval:") {
						lines[i] = "  eval: 'agent-trap current-eval'"
					}
					if strings.HasPrefix(line, "  boundaries:") {
						lines[i] = "  boundaries: 'agent-trap current-boundaries'"
					}
					if strings.HasPrefix(line, "verify_dependents:") {
						lines[i] = "verify_dependents: none"
					}
				}
				projectPut(t, f.Root, "fireweed.yaml", strings.Join(lines, "\n"))
				current, e := p.Load(f.Where)
				if e != nil || current.Config.Hooks.Eval != "agent-trap current-eval" || current.Config.VerifyDependents != "none" {
					t.Fatal("valid conflicting config premise", e, current.Config)
				}
			case "worker-module-hooks":
				raw, e := os.ReadFile(filepath.Join(f.Root, "units/a/SPEC.md"))
				if e != nil {
					t.Fatal(e)
				}
				projectPut(t, f.Root, "units/a/SPEC.md", strings.Replace(string(raw), "name: alpha\n", "name: alpha\nhooks: {eval: 'agent-trap current-module-eval', boundaries: 'agent-trap current-module-boundaries'}\n", 1))
				current, e := p.Load(f.Where)
				if e != nil {
					t.Fatal("valid module override premise", e)
				}
				found := false
				for _, m := range current.Modules {
					if m.Name == "alpha" {
						found = m.Hooks.Eval == "agent-trap current-module-eval" && m.Hooks.Boundaries == "agent-trap current-module-boundaries"
					}
				}
				if !found {
					t.Fatal("valid conflicting module hook premise", current.Modules)
				}
			case "blocked":
				projectPut(t, f.Root, f.Path, strings.Replace(f.Raw, "module: alpha\n", "module: alpha\nblocked: resolve the question\n", 1))
			case "forbidden-eval":
				projectPut(t, f.Root, "units/a/evals/case.txt", "forbidden current mutation")
			case "unsafe-current":
				outside := filepath.Join(t.TempDir(), "outside")
				projectPut(t, filepath.Dir(outside), "outside", f.Raw)
				if e := os.Remove(filepath.Join(f.Root, f.Path)); e != nil {
					t.Fatal(e)
				}
				f.Outside = []string{outside}
				before := cliVerifyOutside(t, f)
				projectPut(t, filepath.Dir(outside), "outside", "observer mutation")
				if reflect.DeepEqual(before, cliVerifyOutside(t, f)) {
					t.Fatal("outside content observer not live")
				}
				projectPut(t, filepath.Dir(outside), "outside", f.Raw)
				if e := os.Symlink(outside, filepath.Join(f.Root, f.Path)); e != nil {
					t.Fatal(e)
				}
			case "malformed-unrelated":
				projectPut(t, f.Plan, "bad-claim", "not JSON")
				oid := strings.TrimSpace(gitFixtureCommand(t, f.Root, "hash-object", "-w", filepath.Join(f.Plan, "bad-claim")))
				gitFixtureCommand(t, f.Root, "update-ref", "refs/fireweed/claims/unrelated", oid)
			}
			outsideBeforeOracle := cliVerifyOutside(t, f)
			want, reports := cliVerifyExpected(t, f, "current")
			if !reflect.DeepEqual(outsideBeforeOracle, cliVerifyOutside(t, f)) {
				t.Fatal("public rejected call changed outside content")
			}
			if mode == "worker-valid-config" || mode == "worker-module-hooks" {
				if cliVerifyExit(want) != 1 || len(want.Coverage) != 0 || len(cliVerifyCalls(t, f)) != 0 {
					t.Fatal("approved current-metadata guard premise", want)
				}
			}
			if mode == "rich-report" {
				cov := want.Coverage[0]
				if len(cov.Tests) != 3 || len(cov.Tests[1].Errors) != 1 || cov.Tests[0].Skipped == nil || len(cov.BySpec["ALP-002"]) != 2 || cov.Tests[1].Errors[0].Text != "retained error text" || cov.Tests[0].Skipped.Text != "retained skip text" {
					t.Fatal("rich public evidence control", cov)
				}
			}
			if mode == "audit-multiple" && (len(want.Findings) != 2 || want.Findings[0].Clause != "MUST retain" || want.Findings[1].Clause != "new behavior") {
				t.Fatal("ordered finding control", want.Findings)
			}
			if mode == "rename" {
				found := false
				for _, e := range want.Edits {
					found = found || e.OldPath != ""
				}
				if !found {
					t.Fatal("native rename premise absent", want.Edits)
				}
			}
			if mode == "audit-finding" || mode == "audit-untestable" {
				if len(want.Findings) != 1 || cliVerifyExit(want) != 0 {
					t.Fatal("finding premise", want)
				}
			}
			cliVerifyReadonly(t, f, func() { cliVerifyJSON(t, f, []string{"verify", "--json", "current"}, want, reports) })
			cliVerifyClear(t, f)
			outsideBeforePlain := cliVerifyOutside(t, f)
			var got result
			cliVerifyReadonly(t, f, func() { got = invoke(t, f.Where, "verify", "current") })
			cliVerifyObserveVerification(t, f, want, subject)
			if !reflect.DeepEqual(outsideBeforePlain, cliVerifyOutside(t, f)) {
				t.Error("plain rejected call changed outside content")
			}
			actualReports := cliVerifyReports(t, f)
			if got.code != cliVerifyExit(want) || !reflect.DeepEqual(cliVerifyNormalize(decodeData(t, json.RawMessage(got.out)), actualReports), cliVerifyWire(t, want, reports)) {
				t.Error("plain structured failure evidence", got)
			}
			coveragePretty(t, got.out)
			cliVerifyPlainStreams(t, got.err, want, actualReports, reports)
		})
	}
}
func Test_APP026_APP027VerifyFatalAndArguments(t *testing.T) {
	outside := t.TempDir()
	for _, args := range [][]string{{"verify", "current", "extra"}, {"verify", "--unknown"}, {"verify", "--json=true"}, {"verify", "--help=true"}, {"verify", "--projected"}, {"verify", "--history"}, {"verify", "--with-change", "current"}, {"verify", "--", "--json", "extra"}, {"verify", "--unknown", "--help"}, {"verify", "--help", "--json=true"}, {"verify", "--help=true", "--help"}, {"verify", "--projected", "--help"}, {"verify", "--help", "--with-change", "other"}} {
		assertUsage(t, cliVerifyOutsideInvoke(t, outside, args...))
		got := cliVerifyOutsideInvoke(t, outside, append([]string{"--json"}, args...)...)
		x := decode(t, got, "verify")
		if got.code != 2 || string(x.Data) != "null" || len(x.Diagnostics) != 1 || x.Diagnostics[0].Code != "U01" || !strings.Contains(x.Diagnostics[0].Message, "help") {
			t.Error("JSON usage", got)
		}
	}
	for _, args := range [][]string{{"verify", "--help"}, {"verify", "-h"}, {"--json", "verify", "--help"}, {"verify", "current", "extra", "--help"}, {"verify", "current", "extra", "-h"}, {"--json", "verify", "current", "extra", "--help"}, {"--json", "verify", "current", "extra", "-h"}} {
		got := cliVerifyOutsideInvoke(t, outside, args...)
		if got.code != 0 {
			t.Error("verify help", got)
		}
		if args[0] != "--json" && !strings.Contains(got.out, "verify [<change>]") {
			t.Error("plain help content", got)
		}
		if args[0] == "--json" {
			successJSON(t, got, "help")
		}
	}
	for _, mode := range []string{"invalid-slug", "missing-claim", "implicit-unowned", "malformed-claim", "foreign-claim"} {
		t.Run(mode, func(t *testing.T) {
			f := cliVerifyFixture(t, "audit", "sha1", "all", false, false)
			slug := "current"
			args := []string{"verify", slug, "--json"}
			switch mode {
			case "invalid-slug":
				slug = "../unsafe"
				args = []string{"verify", slug, "--json"}
			case "missing-claim", "implicit-unowned":
				if e := w.ReleaseClaim(f.Root, "current", f.Claim.Object); e != nil {
					t.Fatal(e)
				}
				if mode == "implicit-unowned" {
					slug = ""
					args = []string{"--json", "verify"}
				}
			case "malformed-claim":
				projectPut(t, f.Plan, "bad", "invalid claim")
				oid := strings.TrimSpace(gitFixtureCommand(t, f.Root, "hash-object", "-w", filepath.Join(f.Plan, "bad")))
				gitFixtureCommand(t, f.Root, "update-ref", "refs/fireweed/claims/current", oid)
			case "foreign-claim":
				claim := f.Claim
				other, e := w.LoadClaim(f.Root, "unrelated")
				if e != nil {
					t.Fatal(e)
				}
				claim.Worktree = other.Worktree
				raw, _ := json.Marshal(claim)
				projectPut(t, f.Plan, "foreign", string(raw))
				oid := strings.TrimSpace(gitFixtureCommand(t, f.Root, "hash-object", "-w", filepath.Join(f.Plan, "foreign")))
				gitFixtureCommand(t, f.Root, "update-ref", "refs/fireweed/claims/current", oid)
			}
			_, err := w.Verify(context.Background(), f.Where, slug)
			if err == nil {
				t.Fatal("public fatal premise absent")
			}
			code := strings.SplitN(err.Error(), ":", 2)[0]
			cliVerifyReadonly(t, f, func() {
				got := invoke(t, f.Where, args...)
				x := decode(t, got, "verify")
				var raw map[string]any
				json.Unmarshal([]byte(got.out), &raw)
				ds := raw["diagnostics"].([]any)
				if len(ds) != 1 {
					t.Fatalf("expected one fatal diagnostic: %+v", got)
				}
				diag := ds[0].(map[string]any)
				if got.code != 1 || string(x.Data) != "null" || len(ds) != 1 || diag["code"] != code || diag["message"] != err.Error() || diag["severity"] != "error" {
					t.Error("fatal source changed", got, err)
				}
				if _, exists := diag["path"]; exists {
					t.Error("fatal invented path")
				}
			})
			plainArgs := []string{"verify"}
			if slug != "" {
				plainArgs = append(plainArgs, slug)
			}
			var got result
			cliVerifyReadonly(t, f, func() { got = invoke(t, f.Where, plainArgs...) })
			cliVerifyFatalPlain(t, got, code, err.Error(), 1)
		})
	}
	got := cliVerifyOutsideInvoke(t, outside, "verify", "--", "--json")
	_, err := w.Verify(context.Background(), outside, "--json")
	if err == nil {
		t.Fatal("literal selector public error absent")
	}
	cliVerifyFatalPlain(t, got, "C05", err.Error(), 1)
}
func Test_APP026_APP028VerifyWritersAndHookChanges(t *testing.T) {
	probeBinary, probe := cliVerifyWriterProbe(t)
	for _, mode := range []string{"data-writer", "diagnostic-writer", "fatal-writer", "usage-writer", "hook-file", "hook-head"} {
		t.Run(mode, func(t *testing.T) {
			f := cliVerifyFixture(t, "impl", "sha1", "all", false, false)
			if mode == "hook-file" {
				projectPut(t, f.Plan, "alpha.mutate", "retained configured hook bytes")
			}
			if mode == "hook-head" {
				projectPut(t, f.Plan, "alpha.head", "yes")
			}
			if mode == "diagnostic-writer" {
				projectPut(t, f.Root, "units/a/evals/case.txt", "forbidden")
			}

			watch := cliVerifyWatchEffects(t, f)
			public := func() (w.Verification, error) {
				if mode == "usage-writer" {
					return w.Verification{}, nil
				}
				slug := "current"
				if mode == "fatal-writer" {
					slug = "missing"
				}
				return w.Verify(context.Background(), f.Where, slug)
			}
			restore := func() {
				if mode == "hook-file" {
					if e := os.Remove(filepath.Join(f.Root, "private/from-hook")); e != nil {
						t.Fatal(e)
					}
				}
				if mode == "hook-head" {
					gitFixtureCommand(t, f.Root, "update-ref", "HEAD", f.Head)
				}
			}
			beforePublic := watch.capture(t)
			want, publicErr := public()
			if (mode == "fatal-writer") != (publicErr != nil) {
				t.Fatal("writer public fatal premise", mode, publicErr)
			}
			reports := cliVerifyReports(t, f)
			afterPublic := watch.capture(t)
			cliVerifyClear(t, f)
			restore()
			if !t.Run("public-control", func(t *testing.T) {
				before := watch.capture(t)
				again, e := public()
				if (e != nil) != (publicErr != nil) {
					t.Fatal("writer repeat fatal premise", e, publicErr)
				}
				if e != nil && e.Error() != publicErr.Error() {
					t.Fatal("writer repeat fatal message", e, publicErr)
				}
				actual := cliVerifyReports(t, f)
				if !reflect.DeepEqual(cliVerifyWire(t, again, actual), cliVerifyWire(t, want, reports)) {
					t.Fatal("writer repeat public evidence", again, want)
				}
				bin, e := os.Executable()
				if e != nil {
					t.Fatal(e)
				}
				cliVerifyObserveVerification(t, f, want, bin)
				cliVerifyAssertEffects(t, cliVerifyReplayEffects(t, beforePublic, afterPublic, before), watch.capture(t))
			}) {
				return
			}
			cliVerifyClear(t, f)
			restore()
			beforeCLI := watch.capture(t)
			got := probe(t, f, mode)
			if got.code != 2 {
				t.Error("output failure exit", got)
			}
			cliVerifyObserveVerification(t, f, want, probeBinary)
			cliVerifyAssertEffects(t, cliVerifyReplayEffects(t, beforePublic, afterPublic, beforeCLI), watch.capture(t))
			switch mode {
			case "fatal-writer", "usage-writer":
				if got.out != "" {
					t.Error("failed fatal/usage writer stdout", got)
				}
			case "diagnostic-writer":
				if got.out == "" {
					break
				}
				if !strings.Contains(got.out, "\n  ") {
					t.Error("failed diagnostic writer output lost pretty formatting", got)
				}
				var data any
				if e := json.Unmarshal([]byte(got.out), &data); e != nil {
					t.Error("failed diagnostic writer lost data", e, got)
				} else if !reflect.DeepEqual(cliVerifyNormalize(data, cliVerifyReports(t, f)), cliVerifyWire(t, want, reports)) {
					t.Error("failed diagnostic writer changed data", got)
				}
			default:
				if got.err != "" {
					t.Error("failed JSON data writer added stderr", got)
				}
			}

			claim, e := w.LoadClaim(f.Root, "current")
			if e != nil || claim != f.Claim {
				t.Error("output failure lost held state", claim, e)
			}
			if mode == "data-writer" || mode == "hook-file" || mode == "hook-head" {
				if len(cliVerifyCalls(t, f)) != 4 {
					t.Error("writer failure skipped public verification")
				}
			}
			if mode == "hook-file" {
				raw, e := os.ReadFile(filepath.Join(f.Root, "private/from-hook"))
				if e != nil || string(raw) != "retained configured hook bytes" {
					t.Error("output failure erased hook mutation", e)
				}
			}
			head := strings.TrimSpace(gitFixtureCommand(t, f.Root, "rev-parse", "HEAD"))
			if (mode == "hook-head") != (head != f.Head) {
				t.Error("writer failure changed/repaired HEAD", mode, head, f.Head)
			}
			text, e := os.ReadFile(filepath.Join(f.Root, f.Path))
			if e != nil || string(text) != f.Raw {
				t.Error("writer failure advanced or discarded change")
			}
		})
	}
}
func Test_APP025_APP026_APP027_APP028VerifyPublicControls(t *testing.T) {
	f := cliVerifyFixture(t, "evals", "sha1", "all", false, false)
	v, reports := cliVerifyExpected(t, f, "current")
	if cliVerifyExit(v) != 0 || len(v.Coverage) != 1 || len(v.Coverage[0].Tests) != 2 || len(v.Coverage[0].Tests[1].Failures) != 1 {
		t.Fatal("real failed-before evidence missing", v)
	}
	wire := cliVerifyWire(t, v, reports).(map[string]any)
	if wire["change"] != "current" || wire["phase"] != "evals" || wire["edits"] == nil {
		t.Fatal("wire oracle control", wire)
	}
	if _, exists := wire["findings"]; exists {
		t.Fatal("empty findings must omit")
	}
	// Test the exact native rename evidence and prescribed edit wire independently.
	gitFixtureCommand(t, f.Root, "mv", "units/a/internal/private.txt", "units/a/internal/new.txt")
	raw := strings.Replace(f.Raw, "phase: evals", "phase: impl", 1)
	projectPut(t, f.Root, f.Path, raw)
	claim, e := w.UpdateClaim(f.Root, f.Claim, f.Claim.Base, raw)
	if e != nil {
		t.Fatal(e)
	}
	f.Claim = claim
	v, e = w.Verify(context.Background(), f.Where, "current")
	if e != nil {
		t.Fatal(e)
	}
	wire = cliVerifyWire(t, v, cliVerifyReports(t, f)).(map[string]any)
	found := false
	for _, item := range wire["edits"].([]any) {
		m := item.(map[string]any)
		if m["old_path"] == "units/a/internal/private.txt" {
			found = true
			if m["status"] != "R" || m["path"] != "units/a/internal/new.txt" {
				t.Error("native rename values lost", m)
			}
		}
		if _, exists := m["OldPath"]; exists {
			t.Error("Go field spelling leaked")
		}
	}
	if !found {
		t.Fatal("native old_path control not observed", wire)
	}
}

func Test_APP025_APP026_APP028VerifyRetainsHookMutations(t *testing.T) {
	for _, kind := range []string{"file", "head", "release", "replace"} {
		for _, jsonMode := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/json-%v", kind, jsonMode), func(t *testing.T) {
				f := cliVerifyFixture(t, "impl", "sha1", "all", false, false)
				watch := cliVerifyWatchEffects(t, f)
				replacement := ""
				switch kind {
				case "file":
					projectPut(t, f.Plan, "alpha.mutate", "retained configured hook edit")
				case "head":
					projectPut(t, f.Plan, "alpha.head", "yes")
				case "release":
					projectPut(t, f.Plan, "alpha.release", "yes")
				case "replace":
					projectPut(t, f.Plan, "replacement", "replacement hook-owned opaque payload")
					replacement = strings.TrimSpace(gitFixtureCommand(t, f.Root, "hash-object", "-w", filepath.Join(f.Plan, "replacement")))
					projectPut(t, f.Plan, "alpha.replace", replacement)
				}
				beforePublic := watch.capture(t)
				want, err := w.Verify(context.Background(), f.Where, "current")
				if err != nil {
					t.Fatal("post-hook public source", err)
				}
				reports := cliVerifyReports(t, f)
				afterPublic := watch.capture(t)
				cliVerifyClear(t, f)
				if cliVerifyExit(want) != 1 || len(want.Diagnostics) == 0 || len(want.Coverage) != 3 {
					t.Fatal("post-hook public guard premise", want)
				}
				restore := func() {
					if kind == "file" {
						if e := os.Remove(filepath.Join(f.Root, "private/from-hook")); e != nil {
							t.Fatal(e)
						}
					}
					if kind == "head" {
						gitFixtureCommand(t, f.Root, "update-ref", "HEAD", f.Head)
					}
					if kind == "release" || kind == "replace" {
						gitFixtureCommand(t, f.Root, "update-ref", "refs/fireweed/claims/current", f.Claim.Object)
					}
				}
				if kind == "file" {
					found := false
					for _, edit := range want.Edits {
						found = found || edit.Path == "private/from-hook"
					}
					if !found {
						t.Fatal("post-hook edit missing from public evidence", want.Edits)
					}
				}
				if kind == "head" {
					if strings.TrimSpace(gitFixtureCommand(t, f.Root, "rev-parse", "HEAD")) == f.Head {
						t.Fatal("native HEAD mutation absent")
					}
				}
				restore()
				if !t.Run("public-control", func(t *testing.T) {
					before := watch.capture(t)
					again, e := w.Verify(context.Background(), f.Where, "current")
					if e != nil {
						t.Fatal("repeated public mutation calibration", e)
					}
					actualReports := cliVerifyReports(t, f)
					if !reflect.DeepEqual(cliVerifyWire(t, again, actualReports), cliVerifyWire(t, want, reports)) {
						t.Fatal("public mutation evidence not repeatable", again, want)
					}
					bin, e := os.Executable()
					if e != nil {
						t.Fatal(e)
					}
					cliVerifyObserveVerification(t, f, want, bin)
					cliVerifyAssertEffects(t, cliVerifyReplayEffects(t, beforePublic, afterPublic, before), watch.capture(t))
					stable := watch.capture(t)
					marker := "ignored/extra-effect-control"
					if _, e := os.Lstat(filepath.Join(f.Root, marker)); !os.IsNotExist(e) {
						t.Fatal("extra-effect control path already exists", e)
					}
					projectPut(t, f.Root, marker, "additional adapter effect")
					if strings.TrimSpace(gitFixtureCommand(t, f.Root, "check-ignore", "--", marker)) != marker || reflect.DeepEqual(stable, watch.capture(t)) {
						t.Fatal("effect observer missed an additional ignored file")
					}
					if e := os.Remove(filepath.Join(f.Root, marker)); e != nil {
						t.Fatal(e)
					}
					cliVerifyAssertEffects(t, stable, watch.capture(t))
				}) {
					return
				}
				cliVerifyClear(t, f)
				restore()
				beforeCLI := watch.capture(t)
				if jsonMode {
					cliVerifyJSON(t, f, []string{"verify", "current", "--json"}, want, reports)
				} else {
					got := invoke(t, f.Where, "verify", "current")
					cliVerifyObserveVerification(t, f, want, subject)
					actualReports := cliVerifyReports(t, f)
					if got.code != 1 || !reflect.DeepEqual(cliVerifyNormalize(decodeData(t, json.RawMessage(got.out)), actualReports), cliVerifyWire(t, want, reports)) {
						t.Error("plain post-hook complete evidence", got)
					}
					cliVerifyPlainStreams(t, got.err, want, actualReports, reports)
				}
				cliVerifyAssertEffects(t, cliVerifyReplayEffects(t, beforePublic, afterPublic, beforeCLI), watch.capture(t))
				switch kind {
				case "file":
					raw, e := os.ReadFile(filepath.Join(f.Root, "private/from-hook"))
					if e != nil || string(raw) != "retained configured hook edit" {
						t.Error("CLI erased configured file edit", e)
					}
				case "head":
					if strings.TrimSpace(gitFixtureCommand(t, f.Root, "rev-parse", "HEAD")) == f.Head {
						t.Error("CLI repaired configured HEAD edit")
					}
				case "release":
					cmd := exec.Command("git", "show-ref", "--verify", "refs/fireweed/claims/current")
					cmd.Dir = f.Root
					if cmd.Run() == nil {
						t.Error("CLI recreated released ownership")
					}
				case "replace":
					if got := strings.TrimSpace(gitFixtureCommand(t, f.Root, "rev-parse", "refs/fireweed/claims/current")); got != replacement {
						t.Error("CLI erased replacement ownership", got, replacement)
					}
				}
			})
		}
	}
}

func Test_APP025_APP027VerifyExplicitEmptySelector(t *testing.T) {
	f := cliVerifyFixture(t, "impl", "sha1", "none", false, false)
	want, reports := cliVerifyExpected(t, f, "")
	cliVerifyReadonly(t, f, func() { cliVerifyJSON(t, f, []string{"verify", "", "--json"}, want, reports) })
}
