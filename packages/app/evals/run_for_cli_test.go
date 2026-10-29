package evals_test

import (
	"encoding/json"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/workflow"
)

// The worker records Beta's public claim and document after writing work, then
// waits longer than run's bound. Its next invocation takes the ordinary path.
func runForBetaWorker(t *testing.T, f runCLISetup, blockedAlpha bool) {
	t.Helper()
	path := filepath.Join(f.logs, "worker.sh")
	script := string(mustReadNextCLI(t, path))
	if blockedAlpha {
		selection := "  [ -f \"$file\" ] || continue\n"
		if !strings.Contains(script, selection) {
			t.Fatal("run fixture selection point missing")
		}
		script = strings.Replace(script, selection, selection+"  [ -z \"$(sed -n 's/^blocked: //p' \"$file\" | head -1)\" ] || continue\n", 1)
	}
	anchor := "case \"$phase\" in\n"
	if !strings.Contains(script, anchor) {
		t.Fatal("run fixture agent insertion point missing")
	}
	addition := `if [ "$module" = alpha ] && [ "$phase" = evals ] && [ -f .fireweed/omit-added ]; then
  exit 0
fi
if [ "$module" = beta ] && [ "$phase" = evals ] && [ ! -f "$1/time-beta-once" ]; then
  : > "$1/time-beta-once"
  printf 'package evals\n// TIME_PARTIAL_WORK\n' > packages/beta/evals/partial_test.go
  cp "$held" "$1/original-beta.md"
  git rev-parse refs/fireweed/claims/beta > "$1/original-beta-claim"
  sleep 18
elif [ "$module" = beta ] && [ "$phase" = evals ]; then
  cp "$held" "$1/resumed-beta.md"
  git rev-parse refs/fireweed/claims/beta > "$1/resumed-beta-claim"
fi
`
	appPut(t, f.logs, "worker.sh", strings.Replace(script, anchor, addition+anchor, 1))
}

func runForReason(t *testing.T, reason, limit string) {
	t.Helper()
	if strings.TrimSpace(reason) == "" || !strings.Contains(reason, limit) || !regexp.MustCompile(`(?i)time|duration|deadline`).MatchString(reason) {
		t.Errorf("time-stop reason=%q; want %s and time limit", reason, limit)
	}
}

func runForCalls(t *testing.T, f runCLISetup, bin string, want ...string) {
	t.Helper()
	calls := runCLICalls(t, f)
	if len(calls) != len(want) {
		t.Fatalf("calls=%v; want %v", calls, want)
	}
	for i, phase := range want {
		if expected := phase + "|" + f.root + "|" + bin; calls[i] != expected {
			t.Errorf("call %d=%q; want %q", i, calls[i], expected)
		}
	}
}

func runForBetaHeld(t *testing.T, f runCLISetup, head string) (workflow.State, string) {
	t.Helper()
	held := cliHeldState(t, f.root, "beta", "BET", head)
	original := string(mustReadNextCLI(t, filepath.Join(f.logs, "original-beta.md")))
	queued := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/0002-beta.md")))
	if queued != original || !strings.Contains(original, "run: "+held.Run) || !strings.Contains(original, "phase: evals") || !strings.Contains(original, "### BET-002") {
		t.Errorf("Beta identity/document changed from agent capture: original=%q held=%q", original, queued)
	}
	claim := strings.TrimSpace(string(mustReadNextCLI(t, filepath.Join(f.logs, "original-beta-claim"))))
	if claim == "" || claim != held.ClaimOID {
		t.Errorf("Beta claim changed: original=%q held=%q", claim, held.ClaimOID)
	}
	partial := string(mustReadNextCLI(t, filepath.Join(f.root, "packages/beta/evals/partial_test.go")))
	if !strings.Contains(partial, "TIME_PARTIAL_WORK") {
		t.Errorf("Beta agent work missing: %q", partial)
	}
	return held, partial
}

func Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes(t *testing.T) {
	bin := buildAppBinary(t)
	f := newRunCLISetup(t, bin, 2, false)
	runForBetaWorker(t, f, false)
	base := cliGit(t, f.root, "rev-parse", "HEAD")
	status, out, stderr := runNextCLI(t, bin, f.root, "run", "--for", "8s", "--json")
	if status != 3 || stderr != "" {
		t.Fatalf("run stop exit=%d stdout=%q stderr=%q", status, out, stderr)
	}
	data, applied := runCLIData(t, out)
	runCLISummaries(t, applied, "alpha")
	runCLIFlag(t, data, "empty", false)
	runCLIFlag(t, data, "limited", false)
	var reason string
	if err := json.Unmarshal(data["reason"], &reason); err != nil {
		t.Errorf("reason=%s: %v", data["reason"], err)
	}
	runForReason(t, reason, "8s")
	var diagnostics []struct {
		Warning bool `json:"warning"`
	}
	if err := json.Unmarshal(decodeEnvelope(t, out)["diagnostics"], &diagnostics); err != nil || diagnostics == nil {
		t.Fatalf("diagnostics malformed: %v", err)
	}
	for _, d := range diagnostics {
		if !d.Warning {
			t.Errorf("clean stop has error diagnostic: %v", d)
		}
	}
	alphaArchive := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/archive/0001-alpha.md")))
	runCLIReport(t, f, "alpha", "ALP", "evals", "fail")
	runCLIReport(t, f, "alpha", "ALP", "impl", "pass")
	head := cliGit(t, f.root, "rev-parse", "HEAD")
	if count := cliGit(t, f.root, "rev-list", "--count", base+"..HEAD"); count != "3" {
		t.Errorf("stop phase commits=%s; want Alpha's three", count)
	}
	held, partial := runForBetaHeld(t, f, head)
	runForCalls(t, f, bin, "alpha|evals", "alpha|audit", "alpha|impl", "beta|evals")

	status, out, stderr = runNextCLI(t, bin, f.root, "next", "--json")
	if status != 0 || stderr != "" {
		t.Errorf("Beta resume exit=%d stdout=%q stderr=%q", status, out, stderr)
	}
	var resumed struct {
		Change, Run, Phase string
		Empty              bool
	}
	if err := json.Unmarshal(decodeEnvelope(t, out)["data"], &resumed); err != nil || !strings.Contains(resumed.Change, "beta") || resumed.Run != held.Run || resumed.Phase != "applied" || resumed.Empty {
		t.Errorf("Beta resume=%+v: %v", resumed, err)
	}
	if got := string(mustReadNextCLI(t, filepath.Join(f.logs, "resumed-beta.md"))); !strings.Contains(got, "run: "+held.Run) || !strings.Contains(got, "phase: evals") || !strings.Contains(got, "### BET-002") {
		t.Errorf("Beta resume agent saw changed run, phase or ID: %q", got)
	}
	if got := strings.TrimSpace(string(mustReadNextCLI(t, filepath.Join(f.logs, "resumed-beta-claim")))); got != held.ClaimOID {
		t.Errorf("Beta resume claim=%q; want original %q", got, held.ClaimOID)
	}
	archive := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/archive/0002-beta.md")))
	for _, marker := range []string{"phase: applied", "run: " + held.Run, "### BET-002"} {
		if !strings.Contains(archive, marker) {
			t.Errorf("Beta archive lacks %q", marker)
		}
	}
	if got := string(mustReadNextCLI(t, filepath.Join(f.root, "packages/beta/evals/partial_test.go"))); got != partial {
		t.Error("Beta original work changed on resume")
	}
	if got := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/archive/0001-alpha.md"))); got != alphaArchive {
		t.Error("Alpha archive changed on Beta resume")
	}
	runCLIReport(t, f, "beta", "BET", "evals", "fail")
	runCLIReport(t, f, "beta", "BET", "impl", "pass")
	if count := cliGit(t, f.root, "rev-list", "--count", base+"..HEAD"); count != "6" {
		t.Errorf("total phase commits=%s; want six", count)
	}
	if refs := cliGit(t, f.root, "for-each-ref", "refs/fireweed/claims"); refs != "" {
		t.Errorf("claim remains after resume: %q", refs)
	}
	runForCalls(t, f, bin, "alpha|evals", "alpha|audit", "alpha|impl", "beta|evals", "beta|evals", "beta|audit", "beta|impl")
}

func Test_APP_090_APP_084_RunForPlainStop(t *testing.T) {
	bin := buildAppBinary(t)
	f := newRunCLISetup(t, bin, 1, false)
	path := filepath.Join(f.logs, "worker.sh")
	script := string(mustReadNextCLI(t, path))
	anchor := "case \"$phase\" in\n"
	if !strings.Contains(script, anchor) {
		t.Fatal("run fixture agent insertion point missing")
	}
	addition := `if [ "$module" = alpha ] && [ "$phase" = evals ]; then
  printf 'package evals\n// TIME_PARTIAL_WORK\n' > packages/alpha/evals/partial_test.go
  sleep 9
fi
`
	appPut(t, f.logs, "worker.sh", strings.Replace(script, anchor, addition+anchor, 1))
	head := cliGit(t, f.root, "rev-parse", "HEAD")
	status, out, stderr := runNextCLI(t, bin, f.root, "run", "--for", "3s")
	if status != 3 || stderr != "" {
		t.Fatalf("plain stop exit=%d stdout=%q stderr=%q", status, out, stderr)
	}
	runCLIPlain(t, out, 0, "", "")
	runForReason(t, out, "3s")
	if !regexp.MustCompile(`(?i)resum|retry|run next|run again`).MatchString(out) {
		t.Errorf("plain stop lacks resume guidance: %q", out)
	}
	state, err := workflow.Current(f.root, "alpha")
	if err != nil || state.Phase != "evals" || state.Run == "" || state.ClaimOID == "" || len(state.Change.Added) != 1 || state.Change.Added[0].ID != "ALP-002" {
		t.Errorf("Alpha held state=%+v: %v", state, err)
	}
	if got := cliGit(t, f.root, "rev-parse", "HEAD"); got != head {
		t.Errorf("interrupted Alpha committed: %s", got)
	}
	if work := string(mustReadNextCLI(t, filepath.Join(f.root, "packages/alpha/evals/partial_test.go"))); !strings.Contains(work, "TIME_PARTIAL_WORK") {
		t.Errorf("Alpha work missing: %q", work)
	}
	runForCalls(t, f, bin, "alpha|evals")
}

func Test_APP_084_RunForPriorBlockStillFails(t *testing.T) {
	bin := buildAppBinary(t)
	f := newRunCLISetup(t, bin, 2, false)
	config := string(mustReadNextCLI(t, filepath.Join(f.root, "fireweed.yaml")))
	appPut(t, f.root, "fireweed.yaml", config+"attempts: 1\n")
	appPut(t, f.root, ".fireweed/omit-added", "yes\n")
	cliCommit(t, f.root, "omit projected coverage")
	runForBetaWorker(t, f, true)
	base := cliGit(t, f.root, "rev-parse", "HEAD")
	status, out, stderr := runNextCLI(t, bin, f.root, "run", "--for", "8s", "--json")
	if status != 1 || stderr != "" {
		t.Fatalf("mixed run exit=%d stdout=%q stderr=%q", status, out, stderr)
	}
	nextCLIDiagnostic(t, status, 1, out, stderr, true, "C10")
	env := decodeEnvelope(t, out)
	if string(env["data"]) != "null" {
		t.Errorf("mixed failure data=%s; want null", env["data"])
	}
	var diagnostics []struct {
		Message string `json:"message"`
	}
	if err := json.Unmarshal(env["diagnostics"], &diagnostics); err != nil || len(diagnostics) == 0 {
		t.Fatalf("mixed diagnostics=%s: %v", env["diagnostics"], err)
	}
	for _, marker := range []string{"ALP-002", "refs/fireweed/attempts/alpha/1", "8s"} {
		if !strings.Contains(diagnostics[0].Message, marker) {
			t.Errorf("mixed failure lacks %q: %q", marker, diagnostics[0].Message)
		}
	}
	runForReason(t, diagnostics[0].Message, "8s")
	alpha := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/0001-alpha.md")))
	if !strings.Contains(alpha, "blocked:") || !strings.Contains(alpha, "phase: evals") || !strings.Contains(alpha, "### ALP-002") {
		t.Errorf("Alpha block state=%q", alpha)
	}
	ref := "refs/fireweed/attempts/alpha/1"
	if saved := cliGit(t, f.root, "for-each-ref", "--format=%(objectname)", ref); saved == "" || cliGit(t, f.root, "cat-file", "-t", saved) != "commit" {
		t.Errorf("missing Alpha saved attempt %s", ref)
	}
	head := cliGit(t, f.root, "rev-parse", "HEAD")
	if count := cliGit(t, f.root, "rev-list", "--count", base+"..HEAD"); count != "1" {
		t.Errorf("stop phase commits=%s; want Alpha's block only", count)
	}
	runForBetaHeld(t, f, head)
	runForCalls(t, f, bin, "alpha|evals", "beta|evals")
}

func Test_APP_089_RunForInvalidOptionsBeforeMutation(t *testing.T) {
	bin := buildAppBinary(t)
	f := newRunCLISetup(t, bin, 1, false)
	head := cliGit(t, f.root, "rev-parse", "HEAD")
	queued := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/0001-alpha.md")))
	statusBefore := cliGit(t, f.root, "status", "--porcelain", "--untracked-files=all")
	for _, tc := range []struct {
		name string
		args []string
	}{
		{"missing", []string{"run", "--for"}},
		{"malformed", []string{"run", "--for", "nonsense"}},
		{"zero", []string{"run", "--for", "0s"}},
		{"negative", []string{"run", "--for", "-1s"}},
		{"repeated", []string{"run", "--for", "1s", "--for", "2s"}},
		{"next", []string{"next", "--for", "1s"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			args := append(append([]string{}, tc.args...), "--json")
			status, out, stderr := runNextCLI(t, bin, f.root, args...)
			if status != 2 || stderr != "" {
				t.Errorf("invalid For exit=%d stdout=%q stderr=%q", status, out, stderr)
			}
			message := assertFailedJSON(t, out, "U01")
			if !strings.Contains(message, "run") || !strings.Contains(message, "--for") {
				t.Errorf("diagnostic lacks valid run --for usage: %q", message)
			}
			if got := cliGit(t, f.root, "rev-parse", "HEAD"); got != head {
				t.Errorf("invalid For committed: %s", got)
			}
			if got := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/0001-alpha.md"))); got != queued {
				t.Error("invalid For changed queued file")
			}
			if got := cliGit(t, f.root, "status", "--porcelain", "--untracked-files=all"); got != statusBefore {
				t.Errorf("invalid For changed files: %q", got)
			}
			if claims := cliGit(t, f.root, "for-each-ref", "refs/fireweed/claims"); claims != "" {
				t.Errorf("invalid For claimed work: %q", claims)
			}
			if calls := runCLICalls(t, f); len(calls) != 0 {
				t.Errorf("invalid For invoked agents: %v", calls)
			}
		})
	}
}
