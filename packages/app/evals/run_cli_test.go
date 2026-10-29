package evals_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/checks"
)

type runCLISetup struct{ root, logs string }

var runCLIModules = []struct{ name, prefix string }{{"alpha", "ALP"}, {"beta", "BET"}}

// The agents only edit their assigned eval, implementation, or audit document.
// The CLI under test drives every lifecycle transition.
func newRunCLISetup(t *testing.T, bin string, queued int, unmerged bool) runCLISetup {
	t.Helper()
	root, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	logs, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	worker := filepath.Join(logs, "worker.sh")
	auditor := filepath.Join(logs, "auditor.sh")
	appPut(t, logs, "worker.sh", `#!/bin/sh
held=''
for file in changes/*.md; do
  [ -f "$file" ] || continue
  case "$(sed -n 's/^phase: //p' "$file" | head -1)" in evals|impl) held="$file"; break;; esac
done
[ -n "$held" ] || exit 30
phase=$(sed -n 's/^phase: //p' "$held" | head -1)
module=$(sed -n 's/^module: //p' "$held" | head -1)
case "$module" in alpha) prefix=ALP;; beta) prefix=BET;; *) exit 31;; esac
printf '%s|%s|%s|%s\n' "$module" "$phase" "$PWD" "$FIREWEED_BIN" >> "$1/invocations"
case "$phase" in
  evals) cat > "packages/$module/evals/added_test.go" <<GO
package evals
import ("testing"; "example.invalid/run-cli/packages/$module")
func Test_${prefix}_002_Added(t *testing.T) { if $module.Value() != "new" { t.Fatal("expected new value") } }
GO
    ;;
  impl) cat > "packages/$module/internal/value.go" <<'GO'
package internal
func Value() string { return "new" }
GO
    ;;
  *) exit 32;;
esac
`)
	appPut(t, logs, "auditor.sh", `#!/bin/sh
held=''
for file in changes/*.md; do
  [ -f "$file" ] || continue
  if [ "$(sed -n 's/^phase: //p' "$file" | head -1)" = audit ]; then held="$file"; break; fi
done
[ -n "$held" ] || exit 33
module=$(sed -n 's/^module: //p' "$held" | head -1)
case "$module" in alpha) prefix=ALP;; beta) prefix=BET;; *) exit 34;; esac
printf '%s|audit|%s|%s\n' "$module" "$PWD" "$FIREWEED_BIN" >> "$1/invocations"
printf '\n## Audit\nNo findings\nChecked IDs: %s-002\n' "$prefix" >> "$held"
`)
	appPut(t, root, ".gitignore", ".fireweed/runs/\n.fireweed/go-test.log\n")
	appPut(t, root, "go.mod", "module example.invalid/run-cli\n\ngo 1.22\n")
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
  eval: [sh, .fireweed/eval.sh, "{module}", "{report}"]
  boundaries: [sh, .fireweed/boundaries.sh]
agent:
  command: [sh, %q, %q, "{prompt}"]
audit:
  command: [sh, %q, %q, "{prompt}"]
`, worker, logs, auditor, logs))
	appPut(t, root, ".fireweed/eval.sh", `#!/bin/sh
case "$1" in alpha) prefix=ALP;; beta) prefix=BET;; *) exit 35;; esac
expected=$(cat .fireweed/expected-bin)
test "$FIREWEED_BIN" -ef "$expected" || exit 36
go test "./packages/$1/evals" > .fireweed/go-test.log 2>&1
code=$?
if [ "$code" -eq 0 ]; then result=''; else result='<failure message="Go eval failed"/>'; fi
if [ -f .fireweed/omit-added ]; then
  printf '<testsuite><testcase name="%s-001" classname="fixture"/></testsuite>' "$prefix" > "$2"
  exit 0
else
  printf '<testsuite><testcase name="%s-001" classname="fixture"/><testcase name="%s-002" classname="fixture">%s</testcase></testsuite>' "$prefix" "$prefix" "$result" > "$2"
fi
exit "$code"
`)
	appPut(t, root, ".fireweed/expected-bin", bin+"\n")
	appPut(t, root, ".fireweed/boundaries.sh", "#!/bin/sh\nexit 0\n")
	for _, m := range runCLIModules {
		appPut(t, root, "packages/"+m.name+"/SPEC.md", fmt.Sprintf("---\nname: %s\nid_prefix: %s\ndepends_on: []\n---\n# %s\n\n## Specs\n\n### %s-001\n%s MUST return a value.\n\n## Decisions\n", m.name, m.prefix, m.name, m.prefix, m.name))
		appPut(t, root, "packages/"+m.name+"/api.go", fmt.Sprintf("package %s\nimport \"example.invalid/run-cli/packages/%s/internal\"\nfunc Value() string { return internal.Value() }\n", m.name, m.name))
		appPut(t, root, "packages/"+m.name+"/internal/value.go", "package internal\nfunc Value() string { return \"old\" }\n")
		appPut(t, root, "packages/"+m.name+"/evals/existing_test.go", fmt.Sprintf("package evals\nimport (\"testing\"; \"example.invalid/run-cli/packages/%s\")\nfunc Test_%s_001_Value(t *testing.T) { if %s.Value() == \"\" { t.Fatal(\"empty\") } }\n", m.name, m.prefix, m.name))
	}
	cliGit(t, root, "init", "-q")
	cliGit(t, root, "checkout", "-q", "-b", "main")
	cliGit(t, root, "config", "user.name", "Fixture")
	cliGit(t, root, "config", "user.email", "fixture@example.invalid")
	if unmerged {
		cliCommit(t, root, "base")
		cliGit(t, root, "checkout", "-q", "-b", "feature")
	}
	for i, m := range runCLIModules[:queued] {
		appPut(t, root, fmt.Sprintf("changes/%04d-%s.md", i+1, m.name), fmt.Sprintf("---\nmodule: %s\nbases: {}\n---\n\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### new:value\n%s MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n", m.name, m.name))
	}
	cliCommit(t, root, "fixture")
	return runCLISetup{root, logs}
}

func runCLICalls(t *testing.T, f runCLISetup) []string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(f.logs, "invocations"))
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		t.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func runCLIData(t *testing.T, stdout string) (map[string]json.RawMessage, []map[string]json.RawMessage) {
	t.Helper()
	env := decodeEnvelope(t, stdout)
	var data map[string]json.RawMessage
	if err := json.Unmarshal(env["data"], &data); err != nil || data == nil {
		t.Fatalf("data=%s: %v", env["data"], err)
	}
	var applied []map[string]json.RawMessage
	if err := json.Unmarshal(data["applied"], &applied); err != nil || applied == nil {
		t.Fatalf("applied=%s: %v", data["applied"], err)
	}
	return data, applied
}

func runCLIFlag(t *testing.T, data map[string]json.RawMessage, key string, want bool) {
	t.Helper()
	var got bool
	if err := json.Unmarshal(data[key], &got); err != nil || got != want {
		t.Errorf("%s=%s; want %v: %v", key, data[key], want, err)
	}
}

func runCLISummaries(t *testing.T, applied []map[string]json.RawMessage, names ...string) {
	t.Helper()
	if len(applied) != len(names) {
		t.Fatalf("applied=%v; want %v", applied, names)
	}
	for i, name := range names {
		var change, run, phase string
		for field, target := range map[string]*string{"change": &change, "run": &run, "phase": &phase} {
			if err := json.Unmarshal(applied[i][field], target); err != nil {
				t.Errorf("applied[%d].%s=%s: %v", i, field, applied[i][field], err)
			}
		}
		if !strings.Contains(change, name) || run == "" || phase != "applied" {
			t.Errorf("applied[%d] = change %q run %q phase %q", i, change, run, phase)
		}
	}
}

func runCLISuccess(t *testing.T, status int, out, errOut string, jsonMode bool) {
	t.Helper()
	if status != 0 || (jsonMode && errOut != "") {
		t.Fatalf("run exit=%d stdout=%q stderr=%q", status, out, errOut)
	}
}

func runCLIPlain(t *testing.T, out string, count int, name, state string) {
	t.Helper()
	lower := strings.ToLower(out)
	if name != "" && !strings.Contains(lower, name) {
		t.Errorf("plain run lacks %s: %q", name, out)
	}
	counted := regexp.MustCompile(fmt.Sprintf(`(?i)(?:applied|count)\D*%d\b|\b%d\s+changes?\b`, count, count))
	if !counted.MatchString(out) {
		t.Errorf("plain run lacks applied count %d: %q", count, out)
	}
	if state == "empty" && !strings.Contains(lower, "empty") {
		t.Errorf("plain run lacks empty queue: %q", out)
	}
	if state == "limited" && !strings.Contains(lower, "limit") && !strings.Contains(lower, "remain") && !strings.Contains(lower, "queued") {
		t.Errorf("plain run lacks queued work: %q", out)
	}
}

func runCLIReport(t *testing.T, f runCLISetup, name, prefix, phase, want string) {
	t.Helper()
	paths, err := filepath.Glob(filepath.Join(f.root, ".fireweed/runs/*/", phase, name+".json"))
	if err != nil || len(paths) != 1 {
		t.Fatalf("%s %s reports=%v: %v", name, phase, paths, err)
	}
	var report checks.Report
	if err := json.Unmarshal(mustReadNextCLI(t, paths[0]), &report); err != nil {
		t.Fatal(err)
	}
	for _, c := range report.Cases {
		if c.Result == want {
			for _, id := range c.IDs {
				if id == prefix+"-002" {
					return
				}
			}
		}
	}
	t.Errorf("%s %s report lacks %s-002 %s: %#v", name, phase, prefix, want, report)
}

func runCLIAssertCalls(t *testing.T, f runCLISetup, bin string, names ...string) {
	t.Helper()
	calls := runCLICalls(t, f)
	if len(calls) != len(names)*3 {
		t.Fatalf("agent calls=%v; want %d", calls, len(names)*3)
	}
	for i, name := range names {
		for j, phase := range []string{"evals", "audit", "impl"} {
			want := fmt.Sprintf("%s|%s|%s|%s", name, phase, f.root, bin)
			if calls[i*3+j] != want {
				t.Errorf("call %d=%q; want %q", i*3+j, calls[i*3+j], want)
			}
		}
	}
}

func Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation(t *testing.T) {
	bin := buildAppBinary(t)
	f := newRunCLISetup(t, bin, 2, false)
	base := cliGit(t, f.root, "rev-parse", "HEAD")
	status, out, errOut := runNextCLI(t, bin, appNestedDir(t, f.root), "run", "--max", "1")
	runCLISuccess(t, status, out, errOut, false)
	runCLIPlain(t, out, 1, "alpha", "limited")
	if _, err := os.Stat(filepath.Join(f.root, "changes/archive/0001-alpha.md")); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(f.root, "changes/0002-beta.md")); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(f.root, "changes/archive/0002-beta.md")); !os.IsNotExist(err) {
		t.Errorf("beta applied before continuation: %v", err)
	}
	if got := cliGit(t, f.root, "rev-list", "--count", base+"..HEAD"); got != "3" {
		t.Errorf("first phase commits=%s", got)
	}
	runCLIAssertCalls(t, f, bin, "alpha")
	firstArchive := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/archive/0001-alpha.md")))
	for _, marker := range []string{"phase: applied", "### ALP-002", "No findings; checked ALP-002"} {
		if !strings.Contains(firstArchive, marker) {
			t.Errorf("alpha archive lacks %q", marker)
		}
	}
	runCLIReport(t, f, "alpha", "ALP", "evals", "fail")
	runCLIReport(t, f, "alpha", "ALP", "impl", "pass")
	status, out, errOut = runNextCLI(t, bin, appNestedDir(t, f.root), "run", "--max", "2", "--json")
	runCLISuccess(t, status, out, errOut, true)
	data, applied := runCLIData(t, out)
	runCLISummaries(t, applied, "beta")
	runCLIFlag(t, data, "empty", true)
	runCLIFlag(t, data, "limited", false)
	if got := cliGit(t, f.root, "rev-list", "--count", base+"..HEAD"); got != "6" {
		t.Errorf("total phase commits=%s", got)
	}
	runCLIAssertCalls(t, f, bin, "alpha", "beta")
	if got := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/archive/0001-alpha.md"))); got != firstArchive {
		t.Error("alpha reapplied on continuation")
	}
	if _, err := os.Stat(filepath.Join(f.root, "changes/0002-beta.md")); !os.IsNotExist(err) {
		t.Errorf("beta remains queued: %v", err)
	}
	if archive := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/archive/0002-beta.md"))); !strings.Contains(archive, "phase: applied") {
		t.Errorf("beta archive=%q", archive)
	}
	runCLIReport(t, f, "beta", "BET", "evals", "fail")
	runCLIReport(t, f, "beta", "BET", "impl", "pass")
	if claims := cliGit(t, f.root, "for-each-ref", "refs/fireweed/claims"); claims != "" {
		t.Errorf("retained claim: %s", claims)
	}
	prior := cliGit(t, f.root, "rev-parse", "HEAD")
	status, out, errOut = runNextCLI(t, bin, f.root, "run")
	runCLISuccess(t, status, out, errOut, false)
	runCLIPlain(t, out, 0, "", "empty")
	if head := cliGit(t, f.root, "rev-parse", "HEAD"); head != prior {
		t.Error("empty run committed")
	}
	runCLIAssertCalls(t, f, bin, "alpha", "beta")
}

func Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON(t *testing.T) {
	bin := buildAppBinary(t)
	f := newRunCLISetup(t, bin, 2, false)
	status, out, errOut := runNextCLI(t, bin, f.root, "run", "--json")
	runCLISuccess(t, status, out, errOut, true)
	data, applied := runCLIData(t, out)
	runCLISummaries(t, applied, "alpha", "beta")
	runCLIFlag(t, data, "empty", true)
	runCLIFlag(t, data, "limited", false)
	var diagnostics []json.RawMessage
	if env := decodeEnvelope(t, out); json.Unmarshal(env["diagnostics"], &diagnostics) != nil || diagnostics == nil || len(diagnostics) != 0 {
		t.Errorf("success diagnostics=%v", diagnostics)
	}
	runCLIAssertCalls(t, f, bin, "alpha", "beta")
	status, out, errOut = runNextCLI(t, bin, f.root, "run", "--json")
	runCLISuccess(t, status, out, errOut, true)
	data, applied = runCLIData(t, out)
	if len(applied) != 0 {
		t.Errorf("empty applied=%v", applied)
	}
	runCLIFlag(t, data, "empty", true)
	runCLIFlag(t, data, "limited", false)
}

func Test_APP_083_RunPlainMultiple(t *testing.T) {
	bin := buildAppBinary(t)
	f := newRunCLISetup(t, bin, 2, false)
	status, out, errOut := runNextCLI(t, bin, f.root, "run")
	runCLISuccess(t, status, out, errOut, false)
	runCLIPlain(t, out, 2, "alpha", "empty")
	if !strings.Contains(strings.ToLower(out), "beta") {
		t.Errorf("plain run lacks beta: %q", out)
	}
	runCLIAssertCalls(t, f, bin, "alpha", "beta")
}

func Test_APP_086_RunJSONLimited(t *testing.T) {
	bin := buildAppBinary(t)
	f := newRunCLISetup(t, bin, 2, false)
	status, out, errOut := runNextCLI(t, bin, f.root, "run", "--max", "1", "--json")
	runCLISuccess(t, status, out, errOut, true)
	data, applied := runCLIData(t, out)
	runCLISummaries(t, applied, "alpha")
	runCLIFlag(t, data, "empty", false)
	runCLIFlag(t, data, "limited", true)
	runCLIAssertCalls(t, f, bin, "alpha")
}

func Test_APP_082_RunAllowUnmerged(t *testing.T) {
	bin := buildAppBinary(t)
	f := newRunCLISetup(t, bin, 1, true)
	status, out, errOut := runNextCLI(t, bin, f.root, "run")
	nextCLIDiagnostic(t, status, 1, out, errOut, false, "P01")
	if len(runCLICalls(t, f)) != 0 {
		t.Error("default run invoked agents without approval")
	}
	status, out, errOut = runNextCLI(t, bin, f.root, "run", "--allow-unmerged")
	runCLISuccess(t, status, out, errOut, false)
	runCLIPlain(t, out, 1, "alpha", "empty")
	runCLIAssertCalls(t, f, bin, "alpha")
}

func Test_APP_081_APP_084_RunUsageAndConfiguration(t *testing.T) {
	bin := buildAppBinary(t)
	f := newRunCLISetup(t, bin, 0, false)
	for _, args := range [][]string{
		{"run", "--max"}, {"run", "--max", "oops"}, {"run", "--max", "0"},
		{"run", "--max", "-1"}, {"run", "extra"}, {"run", "--bogus"},
		{"modules", "--max", "1"}, {"modules", "--allow-unmerged"},
	} {
		status, out, errOut := runNextCLI(t, bin, f.root, args...)
		nextCLIDiagnostic(t, status, 2, out, errOut, false, "U01")
	}
	appPut(t, f.root, "fireweed.yaml", "modules: [broken\n")
	status, out, errOut := runNextCLI(t, bin, f.root, "run")
	nextCLIDiagnostic(t, status, 2, out, errOut, false, "C01")
}

func Test_APP_084_APP_085_APP_086_RunWrappedFailure(t *testing.T) {
	bin := buildAppBinary(t)
	for _, jsonMode := range []bool{false, true} {
		f := newRunCLISetup(t, bin, 1, false)
		appPut(t, f.root, ".fireweed/omit-added", "yes\n")
		cliCommit(t, f.root, "omit projected coverage")
		args := []string{"run"}
		if jsonMode {
			args = append(args, "--json")
		}
		status, out, errOut := runNextCLI(t, bin, f.root, args...)
		nextCLIDiagnostic(t, status, 1, out, errOut, jsonMode, "C10")
		if !strings.Contains(out+errOut, "ALP-002") {
			t.Errorf("wrapped cause lost ID: stdout=%q stderr=%q", out, errOut)
		}
	}
	f := newRunCLISetup(t, bin, 0, false)
	status, out, errOut := runNextCLI(t, bin, f.root, "run", "--max", "bad", "--json")
	nextCLIDiagnostic(t, status, 2, out, errOut, true, "U01")
}
