package evals_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/workflow"
)

// The fixture agent records its call before this branch. A usage exit leaves
// real, uncommitted eval work; its next invocation follows the ordinary script.
func nextCLIUsageAgent(t *testing.T, logs string) {
	t.Helper()
	path := filepath.Join(logs, "worker.sh")
	script := string(mustReadNextCLI(t, path))
	anchor := "case \"$phase\" in\n"
	addition := `if [ "$phase" = evals ] && [ ! -f "$1/usage-once" ]; then
  : > "$1/usage-once"
  printf 'package evals\n// USAGE_PARTIAL_WORK\n' > packages/alpha/evals/partial_test.go
  printf 'RATE LIMIT' >&2
  exit 9
fi
`
	if !strings.Contains(script, anchor) {
		t.Fatal("next fixture agent insertion point missing")
	}
	appPut(t, logs, "worker.sh", strings.Replace(script, anchor, addition+anchor, 1))
}

func runCLIUsageAgent(t *testing.T, f runCLISetup) {
	t.Helper()
	path := filepath.Join(f.logs, "worker.sh")
	script := string(mustReadNextCLI(t, path))
	anchor := "case \"$phase\" in\n"
	addition := `if [ "$phase" = evals ] && [ ! -f "$1/usage-$module" ]; then
  : > "$1/usage-$module"
  printf 'package evals\n// USAGE_PARTIAL_WORK\n' > "packages/$module/evals/partial_test.go"
  printf 'RATE LIMIT' >&2
  exit 9
fi
`
	if !strings.Contains(script, anchor) {
		t.Fatal("run fixture agent insertion point missing")
	}
	appPut(t, f.logs, "worker.sh", strings.Replace(script, anchor, addition+anchor, 1))
}

func cliUsagePolicy(t *testing.T, root, budget string) {
	t.Helper()
	config := string(mustReadNextCLI(t, filepath.Join(root, "fireweed.yaml")))
	appPut(t, root, "fireweed.yaml", config+"limit_patterns: [RATE LIMIT]\nlimit_wait: 5ms\nlimit_budget: "+budget+"\nattempts: 1\n")
	cliCommit(t, root, "fixture usage policy")
}

func cliStopEnvelope(t *testing.T, output, wait, budget string) (map[string]json.RawMessage, string) {
	t.Helper()
	env := decodeEnvelope(t, output)
	var data map[string]json.RawMessage
	if err := json.Unmarshal(env["data"], &data); err != nil || data == nil {
		t.Errorf("clean stop data=%s; want object: %v", env["data"], err)
		return nil, ""
	}
	var reason string
	if err := json.Unmarshal(data["reason"], &reason); err != nil {
		t.Errorf("clean stop reason=%s: %v", data["reason"], err)
	}
	cliStopReason(t, reason, wait, budget)
	var diagnostics []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Warning bool   `json:"warning"`
	}
	if err := json.Unmarshal(env["diagnostics"], &diagnostics); err != nil || diagnostics == nil {
		t.Errorf("clean stop diagnostics=%s: %v", env["diagnostics"], err)
	}
	for _, d := range diagnostics {
		if !d.Warning || d.Code == "" || d.Message == "" {
			t.Errorf("clean stop has error/invalid diagnostic: %+v", d)
		}
	}
	return data, reason
}

func cliStopReason(t *testing.T, reason, wait, budget string) {
	t.Helper()
	lower := strings.ToLower(reason)
	for _, marker := range []string{"completed", "next", "budget", wait, budget} {
		if !strings.Contains(lower, strings.ToLower(marker)) {
			t.Errorf("stop reason %q lacks %q", reason, marker)
		}
	}
}

func cliHeldState(t *testing.T, root, slug, prefix, head string) workflow.State {
	t.Helper()
	state, err := workflow.Current(root, slug)
	if err != nil {
		t.Fatalf("held %s state: %v", slug, err)
	}
	if state.Phase != "evals" || state.Run == "" || state.ClaimOID == "" || state.Head != head || len(state.Change.Added) != 1 || state.Change.Added[0].ID != prefix+"-002" {
		t.Errorf("held %s identity: phase=%q run=%q claim=%q head=%q added=%v", slug, state.Phase, state.Run, state.ClaimOID, state.Head, state.Change.Added)
	}
	if got := cliGit(t, root, "rev-parse", "HEAD"); got != head {
		t.Errorf("clean stop committed: HEAD=%s, want %s", got, head)
	}
	if refs := cliGit(t, root, "for-each-ref", "--format=%(refname)", "refs/fireweed/claims"); refs != "refs/fireweed/claims/"+slug {
		t.Errorf("held claim refs=%q", refs)
	}
	change := string(mustReadNextCLI(t, filepath.Join(root, fmt.Sprintf("changes/%04d-%s.md", map[string]int{"next": 1, "beta": 2}[slug], slug))))
	if strings.Contains(change, "blocked:") || !strings.Contains(change, "phase: evals") || !strings.Contains(change, "run: "+state.Run) {
		t.Errorf("held change lost clean state: %q", change)
	}
	return state
}

func Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume(t *testing.T) {
	bin := buildAppBinary(t)
	root, logs := nextCLIProject(t, bin, true, true)
	cliUsagePolicy(t, root, "0s")
	nextCLIUsageAgent(t, logs)
	head := cliGit(t, root, "rev-parse", "HEAD")
	status, out, stderr := runNextCLI(t, bin, appNestedDir(t, root), "next", "--json")
	if status != 3 || stderr != "" {
		t.Errorf("clean next stop exit=%d stderr=%q stdout=%q", status, stderr, out)
	}
	data, _ := cliStopEnvelope(t, out, "5ms", "0s")
	if data != nil {
		var got struct {
			Change string `json:"change"`
			Run    string `json:"run"`
			Phase  string `json:"phase"`
			Empty  bool   `json:"empty"`
		}
		for _, key := range []string{"change", "run", "phase", "empty"} {
			if _, ok := data[key]; !ok {
				t.Errorf("stop data lacks %s: %v", key, data)
			}
		}
		if err := json.Unmarshal(decodeEnvelope(t, out)["data"], &got); err != nil || !strings.Contains(got.Change, "next") || got.Run == "" || got.Phase != "evals" || got.Empty {
			t.Errorf("stop next data=%+v: %v", got, err)
		}
	}
	held := cliHeldState(t, root, "next", "ALP", head)
	partial := string(mustReadNextCLI(t, filepath.Join(root, "packages/alpha/evals/partial_test.go")))
	if !strings.Contains(partial, "USAGE_PARTIAL_WORK") {
		t.Error("usage work missing")
	}
	if calls := strings.Split(strings.TrimSpace(string(mustReadNextCLI(t, filepath.Join(logs, "invocations")))), "\n"); len(calls) != 1 || calls[0] != "evals|"+root+"|"+bin {
		t.Errorf("stop calls=%v", calls)
	}
	status, out, stderr = runNextCLI(t, bin, root, "next", "--json")
	if status != 0 || stderr != "" {
		t.Errorf("owned resume exit=%d stderr=%q stdout=%q", status, stderr, out)
	}
	env := decodeEnvelope(t, out)
	var resumed struct {
		Change, Run, Phase string
		Empty              bool
	}
	if err := json.Unmarshal(env["data"], &resumed); err != nil || resumed.Run != held.Run || resumed.Phase != "applied" || !strings.Contains(resumed.Change, "next") || resumed.Empty {
		t.Errorf("resumed next data=%+v: %v", resumed, err)
	}
	archive := string(mustReadNextCLI(t, filepath.Join(root, "changes/archive/0001-next.md")))
	for _, marker := range []string{"phase: applied", "run: " + held.Run, "### ALP-002", "No findings"} {
		if !strings.Contains(archive, marker) {
			t.Errorf("archive lacks %q", marker)
		}
	}
	if got := string(mustReadNextCLI(t, filepath.Join(root, "packages/alpha/evals/partial_test.go"))); got != partial {
		t.Error("owned resume changed partial work")
	}
	if got := cliGit(t, root, "rev-list", "--count", head+"..HEAD"); got != "3" {
		t.Errorf("resume commits=%s", got)
	}
	if refs := cliGit(t, root, "for-each-ref", "refs/fireweed/claims"); refs != "" {
		t.Errorf("resume retained claim: %s", refs)
	}
	matches, err := filepath.Glob(filepath.Join(root, ".fireweed/runs/*/evals/alpha.json"))
	if err != nil || len(matches) != 1 {
		t.Errorf("eval reports=%v: %v", matches, err)
	} else if body := string(mustReadNextCLI(t, matches[0])); !strings.Contains(body, "ALP-002") || !strings.Contains(body, "fail") {
		t.Errorf("eval report=%q", body)
	}
	matches, err = filepath.Glob(filepath.Join(root, ".fireweed/runs/*/impl/alpha.json"))
	if err != nil || len(matches) != 1 {
		t.Errorf("impl reports=%v: %v", matches, err)
	} else if body := string(mustReadNextCLI(t, matches[0])); !strings.Contains(body, "ALP-002") || !strings.Contains(body, "pass") {
		t.Errorf("impl report=%q", body)
	}
	if calls := strings.Split(strings.TrimSpace(string(mustReadNextCLI(t, filepath.Join(logs, "invocations")))), "\n"); len(calls) != 4 {
		t.Errorf("owned resume calls=%v, want 4", calls)
	} else {
		for i, phase := range []string{"evals", "evals", "audit", "impl"} {
			if want := phase + "|" + root + "|" + bin; calls[i] != want {
				t.Errorf("owned resume call %d=%q, want %q", i, calls[i], want)
			}
		}
	}
}

func Test_APP_076_APP_077_NextPlainUsageStop(t *testing.T) {
	bin := buildAppBinary(t)
	root, logs := nextCLIProject(t, bin, true, true)
	cliUsagePolicy(t, root, "0s")
	nextCLIUsageAgent(t, logs)
	head := cliGit(t, root, "rev-parse", "HEAD")
	status, out, stderr := runNextCLI(t, bin, root, "next")
	if status != 3 || stderr != "" {
		t.Errorf("plain next stop exit=%d stdout=%q stderr=%q", status, out, stderr)
	}
	if !strings.Contains(strings.ToLower(out), "next") || !regexp.MustCompile(`(?i)resum|retry|run next`).MatchString(out) {
		t.Errorf("plain stop lacks change/resume guidance: %q", out)
	}
	cliStopReason(t, out, "5ms", "0s")
	cliHeldState(t, root, "next", "ALP", head)
}

func Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume(t *testing.T) {
	bin := buildAppBinary(t)
	f := newRunCLISetup(t, bin, 2, false)
	cliUsagePolicy(t, f.root, "5ms")
	runCLIUsageAgent(t, f)
	base := cliGit(t, f.root, "rev-parse", "HEAD")
	status, out, stderr := runNextCLI(t, bin, f.root, "run", "--json")
	if status != 3 || stderr != "" {
		t.Errorf("run stop exit=%d stdout=%q stderr=%q", status, out, stderr)
	}
	data, _ := cliStopEnvelope(t, out, "5ms", "5ms")
	if data != nil {
		for _, key := range []string{"applied", "empty", "limited"} {
			if _, ok := data[key]; !ok {
				t.Errorf("run stop data lacks %s: %v", key, data)
			}
		}
		var applied []map[string]json.RawMessage
		if err := json.Unmarshal(data["applied"], &applied); err != nil || len(applied) != 1 {
			t.Errorf("stop applied=%s: %v", data["applied"], err)
		} else {
			runCLISummaries(t, applied, "alpha")
		}
		runCLIFlag(t, data, "empty", false)
		runCLIFlag(t, data, "limited", false)
	}
	alphaArchive := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/archive/0001-alpha.md")))
	if !strings.Contains(alphaArchive, "phase: applied") || !strings.Contains(alphaArchive, "### ALP-002") {
		t.Errorf("alpha archive=%q", alphaArchive)
	}
	runCLIReport(t, f, "alpha", "ALP", "evals", "fail")
	runCLIReport(t, f, "alpha", "ALP", "impl", "pass")
	head := cliGit(t, f.root, "rev-parse", "HEAD")
	if count := cliGit(t, f.root, "rev-list", "--count", base+"..HEAD"); count != "3" {
		t.Errorf("stop phase commits=%s", count)
	}
	held := cliHeldState(t, f.root, "beta", "BET", head)
	partial := string(mustReadNextCLI(t, filepath.Join(f.root, "packages/beta/evals/partial_test.go")))
	if !strings.Contains(partial, "USAGE_PARTIAL_WORK") {
		t.Error("beta partial work missing")
	}
	if calls := runCLICalls(t, f); len(calls) != 5 {
		t.Errorf("shared budget calls=%v, want 5", calls)
	} else {
		for i, call := range []string{"alpha|evals", "alpha|evals", "alpha|audit", "alpha|impl", "beta|evals"} {
			if want := call + "|" + f.root + "|" + bin; calls[i] != want {
				t.Errorf("shared budget call %d=%q, want %q", i, calls[i], want)
			}
		}
	}
	status, out, stderr = runNextCLI(t, bin, f.root, "next", "--json")
	if status != 0 || stderr != "" {
		t.Errorf("beta resume exit=%d stdout=%q stderr=%q", status, out, stderr)
	}
	env := decodeEnvelope(t, out)
	var resumed struct {
		Change, Run, Phase string
		Empty              bool
	}
	if err := json.Unmarshal(env["data"], &resumed); err != nil || !strings.Contains(resumed.Change, "beta") || resumed.Run != held.Run || resumed.Phase != "applied" || resumed.Empty {
		t.Errorf("beta resume=%+v: %v", resumed, err)
	}
	if archive := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/archive/0002-beta.md"))); !strings.Contains(archive, "### BET-002") || !strings.Contains(archive, "run: "+held.Run) {
		t.Errorf("beta archive=%q", archive)
	}
	if got := string(mustReadNextCLI(t, filepath.Join(f.root, "packages/beta/evals/partial_test.go"))); got != partial {
		t.Error("beta work changed on resume")
	}
	if got := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/archive/0001-alpha.md"))); got != alphaArchive {
		t.Error("alpha reapplied")
	}
	runCLIReport(t, f, "beta", "BET", "evals", "fail")
	runCLIReport(t, f, "beta", "BET", "impl", "pass")
	if calls := runCLICalls(t, f); len(calls) != 8 {
		t.Errorf("resume calls=%v, want 8", calls)
	}
	if refs := cliGit(t, f.root, "for-each-ref", "refs/fireweed/claims"); refs != "" {
		t.Errorf("resume retained claim: %s", refs)
	}
}

func Test_APP_083_APP_084_RunPlainSharedUsageStop(t *testing.T) {
	bin := buildAppBinary(t)
	f := newRunCLISetup(t, bin, 2, false)
	cliUsagePolicy(t, f.root, "5ms")
	runCLIUsageAgent(t, f)
	status, out, stderr := runNextCLI(t, bin, f.root, "run")
	if status != 3 || stderr != "" {
		t.Errorf("plain run stop exit=%d stdout=%q stderr=%q", status, out, stderr)
	}
	runCLIPlain(t, out, 1, "alpha", "")
	if !regexp.MustCompile(`(?i)resum|retry|run next|run run`).MatchString(out) {
		t.Errorf("plain run lacks resume guidance: %q", out)
	}
	cliStopReason(t, out, "5ms", "5ms")
	if _, err := os.Stat(filepath.Join(f.root, "changes/archive/0001-alpha.md")); err != nil {
		t.Errorf("alpha not applied: %v", err)
	}
	cliHeldState(t, f.root, "beta", "BET", cliGit(t, f.root, "rev-parse", "HEAD"))
	if calls := runCLICalls(t, f); len(calls) != 5 {
		t.Errorf("plain run calls=%v, want 5", calls)
	}
}

func Test_APP_084_APP_086_EarlierBlockPrecedesUsageStop(t *testing.T) {
	bin := buildAppBinary(t)
	f := newRunCLISetup(t, bin, 2, false)
	cliUsagePolicy(t, f.root, "0s")
	appPut(t, f.root, ".fireweed/omit-added", "yes\n")
	cliCommit(t, f.root, "omit projected coverage")
	path := filepath.Join(f.logs, "worker.sh")
	script := string(mustReadNextCLI(t, path))
	anchor := "case \"$phase\" in\n"
	usage := `if [ "$module" = alpha ] && [ "$phase" = evals ]; then
  exit 0
fi
if [ "$module" = beta ] && [ "$phase" = evals ]; then
  printf 'package evals\n// USAGE_PARTIAL_WORK\n' > packages/beta/evals/partial_test.go
  printf 'RATE LIMIT'
  printf 'RATE LIMIT' >&2
  exit 9
fi
`
	if !strings.Contains(script, anchor) {
		t.Fatal("run fixture agent insertion point missing")
	}
	selection := "  [ -f \"$file\" ] || continue\n"
	if !strings.Contains(script, selection) {
		t.Fatal("run fixture selection point missing")
	}
	script = strings.Replace(script, selection, selection+"  [ -z \"$(sed -n 's/^blocked: //p' \"$file\" | head -1)\" ] || continue\n", 1)
	appPut(t, f.logs, "worker.sh", strings.Replace(script, anchor, usage+anchor, 1))

	status, out, stderr := runNextCLI(t, bin, f.root, "run", "--json")
	nextCLIDiagnostic(t, status, 1, out, stderr, true, "C10")
	env := decodeEnvelope(t, out)
	var diagnostics []struct {
		Message string `json:"message"`
	}
	if err := json.Unmarshal(env["diagnostics"], &diagnostics); err != nil || len(diagnostics) == 0 {
		t.Fatalf("combined diagnostics=%s: %v", env["diagnostics"], err)
	}
	message := diagnostics[0].Message
	for _, marker := range []string{"ALP-002", "refs/fireweed/attempts/alpha/1", "5ms", "0s", "budget"} {
		if !strings.Contains(message, marker) {
			t.Errorf("combined failure lacks %q: %q", marker, message)
		}
	}
	alpha := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/0001-alpha.md")))
	if !strings.Contains(alpha, "blocked:") || !strings.Contains(alpha, "phase: evals") || !strings.Contains(alpha, "### ALP-002") {
		t.Errorf("alpha block state=%q", alpha)
	}
	ref := "refs/fireweed/attempts/alpha/1"
	if saved := cliGit(t, f.root, "for-each-ref", "--format=%(objectname)", ref); saved == "" {
		t.Errorf("missing saved attempt %s", ref)
	} else if kind := cliGit(t, f.root, "cat-file", "-t", saved); kind != "commit" {
		t.Errorf("saved attempt kind=%q", kind)
	}
	cliHeldState(t, f.root, "beta", "BET", cliGit(t, f.root, "rev-parse", "HEAD"))
	partial := string(mustReadNextCLI(t, filepath.Join(f.root, "packages/beta/evals/partial_test.go")))
	if !strings.Contains(partial, "USAGE_PARTIAL_WORK") {
		t.Error("beta partial work missing")
	}
	if calls := runCLICalls(t, f); len(calls) != 2 {
		t.Errorf("combined failure calls=%v, want two", calls)
	} else {
		for i, call := range []string{"alpha|evals", "beta|evals"} {
			if want := call + "|" + f.root + "|" + bin; calls[i] != want {
				t.Errorf("combined failure call %d=%q, want %q", i, calls[i], want)
			}
		}
	}
}

func Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation(t *testing.T) {
	bin := buildAppBinary(t)
	for _, command := range []string{"next", "run"} {
		t.Run(command, func(t *testing.T) {
			f := newRunCLISetup(t, bin, 1, false)
			config := string(mustReadNextCLI(t, filepath.Join(f.root, "fireweed.yaml")))
			appPut(t, f.root, "fireweed.yaml", config+"limit_wait: nonsense\n")
			cliCommit(t, f.root, "invalid usage policy")
			head := cliGit(t, f.root, "rev-parse", "HEAD")
			queued := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/0001-alpha.md")))
			status, out, stderr := runNextCLI(t, bin, f.root, command, "--json")
			nextCLIDiagnostic(t, status, 2, out, stderr, true, "C01")
			if got := cliGit(t, f.root, "rev-parse", "HEAD"); got != head {
				t.Errorf("invalid policy committed: %s", got)
			}
			if got := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/0001-alpha.md"))); got != queued {
				t.Error("invalid policy changed queued document")
			}
			if refs := cliGit(t, f.root, "for-each-ref", "refs/fireweed/claims"); refs != "" {
				t.Errorf("invalid policy claimed work: %s", refs)
			}
			if calls := runCLICalls(t, f); len(calls) != 0 {
				t.Errorf("invalid policy invoked agents: %v", calls)
			}
		})
	}
}
