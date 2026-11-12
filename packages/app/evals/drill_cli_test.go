package evals_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// The scripts are outside the repository. Only the public CLI starts a drill;
// the hooks and agent record the executable and produce deterministic work.
type drillCLISetup struct{ root, nested, observations, base, bin string }

func drillCLIProject(t *testing.T, bin string) drillCLISetup {
	t.Helper()
	outer := appProjectFixture(t)
	outer, err := filepath.EvalSymlinks(outer)
	if err != nil {
		t.Fatal(err)
	}
	root := filepath.Join(outer, "inner-project")
	observations, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	appPut(t, observations, "setup.sh", `#!/bin/sh
printf '%s\n%s\n' "$PWD" "$FIREWEED_BIN" > "$1/setup"
if [ -f "$1/setup-fail" ]; then echo FIXTURE_SETUP_FAILURE >&2; exit 17; fi
`)
	appPut(t, observations, "agent.sh", `#!/bin/sh
printf '%s\n%s\n' "$PWD" "$FIREWEED_BIN" > "$1/agent"
case "$(cat "$1/mode" 2>/dev/null)" in
  usage) echo 'RATE LIMIT' >&2; exit 9;;
  usage-after-failure) if [ -f "$1/attempted" ]; then echo 'RATE LIMIT' >&2; exit 9; fi; : > "$1/attempted"; value=wrong;;
  timeout) sleep 3; exit 0;;
  own-fail) value=wrong;;
  *) value=new;;
esac
mkdir -p packages/alpha/internal
printf '%s\n' "$value" > packages/alpha/internal/value.txt
`)
	appPut(t, observations, "eval.sh", `#!/bin/sh
printf '%s|%s\n' "$1" "$FIREWEED_BIN" >> "$3/evals"
case "$1" in alpha) id=ALP-001;; beta) id=BET-001;; *) exit 2;; esac
failure=''
if [ "$1" = alpha ] && [ "$(cat packages/alpha/internal/value.txt 2>/dev/null)" != new ]; then failure='own fixture failure'; fi
if [ "$1" = beta ] && [ -f "$3/dependent-fail" ]; then failure='dependent fixture failure'; fi
if [ -z "$failure" ]; then
  printf '<testsuite><testcase name="%s_case"/></testsuite>' "$id" > "$2"
else
  printf '<testsuite><testcase name="%s_case"><failure message="%s"/></testcase></testsuite>' "$id" "$failure" > "$2"
  exit 1
fi
`)
	appPut(t, root, ".gitignore", ".fireweed/drills/\n")
	appPut(t, root, "SYSTEM.md", "# Fixture\n")
	appPut(t, root, "fireweed.yaml", fmt.Sprintf(`fireweed: ">=0.1, <0.2"
version: 1
base: main
modules: ["packages/*"]
module_defaults:
  interface: [api.go]
  evals: [evals/**]
  disposable: [internal/**]
hooks:
  setup: [sh, %q, %q]
  eval: [sh, %q, "{module}", "{report}", %q]
agent:
  command: [bash, %q, %q, "{prompt}"]
attempts: 1
timeout: 2s
limit_patterns: [RATE LIMIT]
limit_wait: 5ms
limit_budget: 0s
`, filepath.Join(observations, "setup.sh"), observations, filepath.Join(observations, "eval.sh"), observations, filepath.Join(observations, "agent.sh"), observations))
	for _, m := range []struct{ name, prefix, deps string }{{"alpha", "ALP", "[]"}, {"beta", "BET", "[alpha]"}} {
		appPut(t, root, "packages/"+m.name+"/SPEC.md", fmt.Sprintf("---\nname: %s\nid_prefix: %s\ndepends_on: %s\n---\n# %s\n\n## Specs\n\n### %s-001\n%s MUST preserve a value.\n", m.name, m.prefix, m.deps, m.name, m.prefix, m.name))
		appPut(t, root, "packages/"+m.name+"/api.go", "package "+m.name+"\n")
		appPut(t, root, "packages/"+m.name+"/evals/fixture_test.go", "package evals\n")
	}
	appPut(t, root, "packages/alpha/internal/value.txt", "old\n")
	cliGit(t, root, "init", "-q", "-b", "main")
	cliGit(t, root, "config", "user.name", "Drill CLI Fixture")
	cliGit(t, root, "config", "user.email", "drill-cli@example.invalid")
	cliCommit(t, root, "drill CLI fixture")
	return drillCLISetup{root: root, nested: appNestedDir(t, root), observations: observations, base: cliGit(t, root, "rev-parse", "HEAD"), bin: bin}
}

func drillCLIObserve(t *testing.T, f drillCLISetup, name string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(f.observations, name))
	if os.IsNotExist(err) {
		return ""
	}
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func drillCLINoInvocation(t *testing.T, f drillCLISetup) {
	t.Helper()
	for _, name := range []string{"setup", "agent", "evals"} {
		if got := drillCLIObserve(t, f, name); got != "" {
			t.Errorf("rejected drill invoked %s: %q", name, got)
		}
	}
}

func drillCLICleaned(t *testing.T, f drillCLISetup) {
	t.Helper()
	for _, name := range []string{"setup", "agent"} {
		first := strings.SplitN(drillCLIObserve(t, f, name), "\n", 2)[0]
		if first != "" {
			if _, err := os.Stat(first); !os.IsNotExist(err) {
				t.Errorf("%s scratch directory %s remains: %v", name, first, err)
			}
		}
	}
	if worktrees := cliGit(t, f.root, "worktree", "list", "--porcelain"); strings.Count(worktrees, "worktree ") != 1 {
		t.Errorf("scratch worktree remains: %q", worktrees)
	}
}

func drillCLIJSONData(t *testing.T, output string, stopped bool) map[string]string {
	t.Helper()
	env := decodeEnvelope(t, output)
	var diagnostics []json.RawMessage
	if err := json.Unmarshal(env["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
		t.Errorf("drill diagnostics=%s; want []: %v", env["diagnostics"], err)
	}
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(env["data"], &raw); err != nil || raw == nil {
		t.Fatalf("drill data=%s; want object: %v", env["data"], err)
	}
	fields := []string{"module", "result", "report", "ref", "base_commit"}
	if stopped {
		fields = append(fields, "reason")
	}
	data := map[string]string{}
	for _, field := range fields {
		var value string
		if err := json.Unmarshal(raw[field], &value); err != nil {
			t.Errorf("drill %s=%s; want string: %v", field, raw[field], err)
		}
		data[field] = value
	}
	return data
}

func Test_APP_136_APP_137_DrillPlainNearestProjectAndInvokingBin(t *testing.T) {
	bin := buildAppBinary(t)
	f := drillCLIProject(t, bin)
	status, out, stderr := runNextCLI(t, bin, f.nested, "drill", "alpha")
	if status != 0 || stderr != "" {
		t.Fatalf("plain drill: status=%d stdout=%q stderr=%q", status, out, stderr)
	}
	report := ".fireweed/drills/alpha.md"
	ref := strings.TrimSpace(cliGit(t, f.root, "for-each-ref", "--format=%(refname)", "refs/fireweed/drills/alpha"))
	for _, marker := range []string{"alpha", "pass", report, ref} {
		if marker == "" || !strings.Contains(out, marker) {
			t.Errorf("plain drill output %q lacks %q", out, marker)
		}
	}
	if got := drillCLIObserve(t, f, "setup"); !strings.Contains(got, bin) {
		t.Errorf("setup FIREWEED_BIN=%q; want %s", got, bin)
	}
	if got := drillCLIObserve(t, f, "agent"); !strings.Contains(got, bin) {
		t.Errorf("agent FIREWEED_BIN=%q; want %s", got, bin)
	}
	if got := drillCLIObserve(t, f, "evals"); !strings.Contains(got, "alpha|"+bin) || !strings.Contains(got, "beta|"+bin) {
		t.Errorf("own/dependent evals and Bin=%q", got)
	}
	if body := string(mustReadNextCLI(t, filepath.Join(f.root, report))); !strings.Contains(body, "base_commit: "+f.base) || !strings.Contains(body, "result: pass") {
		t.Errorf("native report=%q", body)
	}
	if got := cliGit(t, f.root, "show", ref+":packages/alpha/internal/value.txt"); got != "new" {
		t.Errorf("retained rebuilt code=%q", got)
	}
	if got := cliGit(t, f.root, "rev-parse", "HEAD"); got != f.base {
		t.Errorf("drill moved invoking HEAD: %s -> %s", f.base, got)
	}
	drillCLICleaned(t, f)
}

func Test_APP_138_DrillJSONSuccess(t *testing.T) {
	bin := buildAppBinary(t)
	f := drillCLIProject(t, bin)
	status, out, stderr := runNextCLI(t, bin, f.nested, "drill", "alpha", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("JSON drill: status=%d stdout=%q stderr=%q", status, out, stderr)
	}
	data := drillCLIJSONData(t, out, false)
	if data["module"] != "alpha" || data["result"] != "pass" || data["report"] != ".fireweed/drills/alpha.md" || data["base_commit"] != f.base || !strings.HasPrefix(data["ref"], "refs/fireweed/drills/alpha/") {
		t.Errorf("JSON drill result=%v", data)
	}
	if body := string(mustReadNextCLI(t, filepath.Join(f.root, data["report"]))); !strings.Contains(body, "result: pass") {
		t.Errorf("JSON report path does not identify native report: %q", body)
	}
	if got := cliGit(t, f.root, "show", data["ref"]+":packages/alpha/internal/value.txt"); got != "new" {
		t.Errorf("JSON ref does not identify rebuilt code: %q", got)
	}
	drillCLICleaned(t, f)
}

func Test_APP_139_DrillArgumentAndModuleRejectionsDoNotInvokeHooks(t *testing.T) {
	bin := buildAppBinary(t)
	f := drillCLIProject(t, bin)
	for _, tc := range []struct {
		args        []string
		code, named string
	}{
		{[]string{"drill"}, "U01", "drill"},
		{[]string{"drill", "alpha", "extra", "--json"}, "U01", "drill"},
		{[]string{"drill", "missing", "--json"}, "U02", "missing"},
	} {
		status, out, stderr := runNextCLI(t, bin, f.nested, tc.args...)
		jsonMode := strings.Contains(strings.Join(tc.args, " "), "--json")
		drillCLIFailure(t, status, 2, out, stderr, jsonMode, tc.code, tc.named)
		drillCLINoInvocation(t, f)
	}
}

func drillCLIFailure(t *testing.T, status, want int, out, stderr string, jsonMode bool, code, named string) {
	t.Helper()
	if status != want {
		t.Errorf("drill exit=%d; want %d; stdout=%q stderr=%q", status, want, out, stderr)
	}
	message := stderr
	if jsonMode {
		if stderr != "" {
			t.Errorf("JSON stderr=%q; want empty", stderr)
		}
		message = assertFailedJSON(t, out, code)
	} else if out != "" || !strings.HasPrefix(strings.TrimSpace(stderr), code+":") {
		t.Errorf("plain failure stdout=%q stderr=%q; want empty stdout and %s", out, stderr, code)
	}
	if !corrective(message) || named != "" && !strings.Contains(message, named) {
		t.Errorf("diagnostic %q lacks correction or %q", message, named)
	}
}

func Test_APP_140_DrillProjectAndConfigurationFailures(t *testing.T) {
	bin := buildAppBinary(t)
	missing := t.TempDir()
	status, out, stderr := runNextCLI(t, bin, missing, "drill", "alpha")
	drillCLIFailure(t, status, 2, out, stderr, false, "C01", "fireweed.yaml")
	f := drillCLIProject(t, bin)
	appPut(t, f.root, "fireweed.yaml", "modules: [broken\n")
	status, out, stderr = runNextCLI(t, bin, f.nested, "drill", "alpha", "--json")
	drillCLIFailure(t, status, 2, out, stderr, true, "C01", "fireweed.yaml")
	drillCLINoInvocation(t, f)
	graph := drillCLIProject(t, bin)
	appPut(t, graph.root, "packages/beta/SPEC.md", strings.Replace(string(mustReadNextCLI(t, filepath.Join(graph.root, "packages/beta/SPEC.md"))), "depends_on: [alpha]", "depends_on: [missing]", 1))
	status, out, stderr = runNextCLI(t, bin, graph.nested, "drill", "alpha", "--json")
	drillCLIFailure(t, status, 2, out, stderr, true, "C02", "missing")
	drillCLINoInvocation(t, graph)
}

func Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name, mode, marker string
		jsonMode           bool
	}{
		{"own-plain", "own-fail", "own fixture failure", false},
		{"dependent-json", "dependent-fail", "dependent fixture failure", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := drillCLIProject(t, bin)
			if tc.mode == "dependent-fail" {
				appPut(t, f.observations, "dependent-fail", "yes\n")
			} else {
				appPut(t, f.observations, "mode", tc.mode+"\n")
			}
			args := []string{"drill", "alpha"}
			if tc.jsonMode {
				args = append(args, "--json")
			}
			status, out, stderr := runNextCLI(t, bin, f.nested, args...)
			drillCLIFailure(t, status, 1, out, stderr, tc.jsonMode, "E01", "")
			ref := cliGit(t, f.root, "for-each-ref", "--format=%(refname)", "refs/fireweed/drills/alpha")
			if ref == "" {
				t.Fatal("failed drill did not retain rebuilt ref")
			}
			for _, artifact := range []string{".fireweed/drills/alpha.md", ref} {
				if !strings.Contains(out+stderr, artifact) {
					t.Errorf("failure diagnostic omits artifact %s: %q %q", artifact, out, stderr)
				}
			}
			body := string(mustReadNextCLI(t, filepath.Join(f.root, ".fireweed/drills/alpha.md")))
			if !strings.Contains(body, "result: fail") || !strings.Contains(body, tc.marker) {
				t.Errorf("failed native report=%q", body)
			}
			if got := cliGit(t, f.root, "show", ref+":packages/alpha/internal/value.txt"); got != map[string]string{"own-fail": "wrong", "dependent-fail": "new"}[tc.mode] {
				t.Errorf("retained ref value=%q", got)
			}
			drillCLICleaned(t, f)
		})
	}
}

func Test_APP_141_DrillSetupFailureHasNoArtifacts(t *testing.T) {
	bin := buildAppBinary(t)
	f := drillCLIProject(t, bin)
	appPut(t, f.observations, "setup-fail", "yes\n")
	status, out, stderr := runNextCLI(t, bin, f.nested, "drill", "alpha", "--json")
	drillCLIFailure(t, status, 1, out, stderr, true, "P01", "FIXTURE_SETUP_FAILURE")
	for _, artifact := range []string{".fireweed/drills/alpha.md", "refs/fireweed/drills/alpha/"} {
		if strings.Contains(out, artifact) {
			t.Errorf("setup failure invented artifact %s: %q", artifact, out)
		}
	}
	if _, err := os.Stat(filepath.Join(f.root, ".fireweed/drills/alpha.md")); !os.IsNotExist(err) {
		t.Errorf("setup failure wrote report: %v", err)
	}
	if refs := cliGit(t, f.root, "for-each-ref", "--format=%(refname)", "refs/fireweed/drills/alpha"); refs != "" {
		t.Errorf("setup failure retained ref: %q", refs)
	}
	drillCLICleaned(t, f)
}

func Test_APP_142_DrillCleanUsageStopPlainAndJSON(t *testing.T) {
	bin := buildAppBinary(t)
	for _, jsonMode := range []bool{false, true} {
		f := drillCLIProject(t, bin)
		if jsonMode {
			config := string(mustReadNextCLI(t, filepath.Join(f.root, "fireweed.yaml")))
			appPut(t, f.root, "fireweed.yaml", strings.Replace(config, "attempts: 1", "attempts: 2", 1))
			cliCommit(t, f.root, "allow verification then usage stop")
			f.base = cliGit(t, f.root, "rev-parse", "HEAD")
			appPut(t, f.observations, "mode", "usage-after-failure\n")
		} else {
			appPut(t, f.observations, "mode", "usage\n")
		}
		args := []string{"drill", "alpha"}
		if jsonMode {
			args = append(args, "--json")
		}
		status, out, stderr := runNextCLI(t, bin, f.nested, args...)
		if status != 3 || stderr != "" {
			t.Errorf("clean stop exit=%d stdout=%q stderr=%q", status, out, stderr)
		}
		if jsonMode {
			data := drillCLIJSONData(t, out, true)
			if data["module"] != "alpha" || data["result"] != "stopped" || data["report"] != ".fireweed/drills/alpha.md" || !strings.HasPrefix(data["ref"], "refs/fireweed/drills/alpha/") || data["base_commit"] != f.base || !strings.Contains(data["reason"], "budget") {
				t.Errorf("clean stop data=%v", data)
			}
			if ref := cliGit(t, f.root, "for-each-ref", "--format=%(refname)", "refs/fireweed/drills/alpha"); ref != data["ref"] {
				t.Errorf("stop ref=%q; want JSON ref %q", ref, data["ref"])
			}
		} else if !strings.Contains(out, "alpha") || !strings.Contains(strings.ToLower(out), "stopped") || !strings.Contains(out, "budget") || !strings.Contains(out, ".fireweed/drills/alpha.md") {
			t.Errorf("plain clean stop omits module, reason, or report: %q", out)
		}
		if body := string(mustReadNextCLI(t, filepath.Join(f.root, ".fireweed/drills/alpha.md"))); !strings.Contains(body, "result: stopped") {
			t.Errorf("stopped report=%q", body)
		}
		if !jsonMode {
			if ref := cliGit(t, f.root, "for-each-ref", "--format=%(refname)", "refs/fireweed/drills/alpha"); ref != "" {
				t.Errorf("usage stop retained ref: %q", ref)
			}
		}
		drillCLICleaned(t, f)
	}
}

func Test_APP_136_DrillUsesConfiguredTimeout(t *testing.T) {
	bin := buildAppBinary(t)
	f := drillCLIProject(t, bin)
	config := string(mustReadNextCLI(t, filepath.Join(f.root, "fireweed.yaml")))
	appPut(t, f.root, "fireweed.yaml", strings.Replace(config, "timeout: 2s", "timeout: 200ms", 1))
	cliCommit(t, f.root, "short configured drill timeout")
	appPut(t, f.observations, "mode", "timeout\n")
	start := time.Now()
	status, out, stderr := runNextCLI(t, bin, f.nested, "drill", "alpha", "--json")
	if status != 1 || stderr != "" || time.Since(start) > 4*time.Second {
		t.Errorf("configured timeout exit=%d after %s: stdout=%q stderr=%q", status, time.Since(start), out, stderr)
	}
	if message := assertFailedJSON(t, out, "P01"); !corrective(message) {
		t.Errorf("timeout diagnostic lacks correction: %q", message)
	}
	drillCLICleaned(t, f)
}
