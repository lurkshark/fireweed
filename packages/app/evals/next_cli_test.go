package evals_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// These fixtures exercise the installed NEW executable and ordinary Git, hooks,
// and separate bounded agents. The scripts cannot drive workflow transitions.
func nextCLIProject(t *testing.T, bin string, queued, approved bool) (string, string) {
	t.Helper()
	root, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	logs := t.TempDir()
	worker := filepath.Join(logs, "worker.sh")
	auditor := filepath.Join(logs, "auditor.sh")
	appPut(t, logs, "worker.sh", `#!/bin/sh
phase=$(sed -n 's/^phase: //p' changes/0001-next.md | head -1)
test "$FIREWEED_BIN" -ef "$3" || exit 12
printf '%s|%s|%s\n' "$phase" "$PWD" "$FIREWEED_BIN" >> "$1/invocations"
case "$phase" in
  evals) cat > packages/alpha/evals/next_test.go <<'GO'
package evals
import ("testing"; "example.invalid/next-cli/packages/alpha")
func Test_ALP_002_NewValue(t *testing.T) { if alpha.Value() != "new" { t.Fatal("expected new value") } }
GO
    ;;
  impl) cat > packages/alpha/internal/value.go <<'GO'
package internal
func Value() string { return "new" }
GO
    ;;
  *) exit 13 ;;
esac
`)
	appPut(t, logs, "auditor.sh", `#!/bin/sh
test "$FIREWEED_BIN" -ef "$3" || exit 14
printf 'audit|%s|%s\n' "$PWD" "$FIREWEED_BIN" >> "$1/invocations"
cat >> changes/0001-next.md <<'AUDIT'

## Audit
No findings
Checked IDs: ALP-002
AUDIT
`)
	appPut(t, root, ".gitignore", ".fireweed/runs/\n.fireweed/go-test.log\n")
	appPut(t, root, "go.mod", "module example.invalid/next-cli\n\ngo 1.22\n")
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
  command: [sh, %q, %q, "{prompt}", %q]
audit:
  command: [sh, %q, %q, "{prompt}", %q]
`, worker, logs, bin, auditor, logs, bin))
	appPut(t, root, ".fireweed/expected-bin", bin+"\n")
	appPut(t, root, ".fireweed/eval.sh", `#!/bin/sh
expected=$(cat .fireweed/expected-bin)
test "$FIREWEED_BIN" -ef "$expected" || exit 15
go test ./packages/alpha/evals > .fireweed/go-test.log 2>&1
code=$?
if [ "$code" -eq 0 ]; then result=''; else result='<failure message="Go eval failed"/>'; fi
if [ -f .fireweed/omit-new ]; then
  printf '<testsuite><testcase name="ALP-001" classname="fixture"/></testsuite>' > "$2"
  exit 0
else
  printf '<testsuite><testcase name="ALP-001" classname="fixture"/><testcase name="ALP-002" classname="fixture">%s</testcase></testsuite>' "$result" > "$2"
fi
exit "$code"
`)
	appPut(t, root, ".fireweed/boundaries.sh", "#!/bin/sh\nexit 0\n")
	appPut(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return a value.\n\n## Decisions\n")
	appPut(t, root, "packages/alpha/api.go", "package alpha\nimport \"example.invalid/next-cli/packages/alpha/internal\"\nfunc Value() string { return internal.Value() }\n")
	appPut(t, root, "packages/alpha/internal/value.go", "package internal\nfunc Value() string { return \"old\" }\n")
	appPut(t, root, "packages/alpha/evals/existing_test.go", "package evals\nimport (\"testing\"; \"example.invalid/next-cli/packages/alpha\")\nfunc Test_ALP_001_Value(t *testing.T) { if alpha.Value() == \"\" { t.Fatal(\"empty\") } }\n")
	cliGit(t, root, "init", "-q")
	cliGit(t, root, "checkout", "-q", "-b", "main")
	cliGit(t, root, "config", "user.name", "Fixture")
	cliGit(t, root, "config", "user.email", "fixture@example.invalid")
	if queued && !approved {
		cliCommit(t, root, "base")
		cliGit(t, root, "checkout", "-q", "-b", "feature")
	}
	if queued {
		appPut(t, root, "changes/0001-next.md", "---\nmodule: alpha\nbases: {}\n---\n\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### new:value\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n")
	}
	cliCommit(t, root, "fixture")
	return root, logs
}

func runNextCLI(t *testing.T, bin, dir string, args ...string) (int, string, string) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 75*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Dir = dir
	var out, errOut bytes.Buffer
	cmd.Stdout, cmd.Stderr = &out, &errOut
	err := cmd.Run()
	if ctx.Err() != nil {
		t.Fatalf("NEW %v timed out: %v", args, ctx.Err())
	}
	if err == nil {
		return 0, out.String(), errOut.String()
	}
	var exit *exec.ExitError
	if !errors.As(err, &exit) {
		t.Fatalf("NEW %v: %v", args, err)
	}
	return exit.ExitCode(), out.String(), errOut.String()
}

func nextCLIDiagnostic(t *testing.T, status, want int, out, errOut string, jsonMode bool, code string) {
	t.Helper()
	if status != want {
		t.Errorf("exit=%d, want %d; stdout=%q stderr=%q", status, want, out, errOut)
	}
	if jsonMode {
		if errOut != "" {
			t.Errorf("JSON stderr=%q", errOut)
		}
		env := decodeEnvelope(t, out)
		if string(env["data"]) != "null" {
			t.Errorf("failure data=%s", env["data"])
		}
		var ds []struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}
		if err := json.Unmarshal(env["diagnostics"], &ds); err != nil || len(ds) == 0 {
			t.Fatalf("diagnostics=%s: %v", env["diagnostics"], err)
		}
		if ds[0].Code != code || !corrective(ds[0].Message) {
			t.Errorf("diagnostic=%+v; want %s and action", ds[0], code)
		}
	} else if out != "" || !strings.Contains(errOut, code+":") || !corrective(errOut) {
		t.Errorf("plain diagnostic stdout=%q stderr=%q; want %s and action", out, errOut, code)
	}
}

func Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange(t *testing.T) {
	bin := buildAppBinary(t)
	root, logs := nextCLIProject(t, bin, true, true)
	base := cliGit(t, root, "rev-parse", "HEAD")
	status, out, errOut := runNextCLI(t, bin, appNestedDir(t, root), "next")
	if status != 0 {
		t.Fatalf("next exit=%d stdout=%q stderr=%q", status, out, errOut)
	}
	if !strings.Contains(out, "next") {
		t.Errorf("plain success did not identify applied change: %q", out)
	}
	for _, phase := range []string{"evals", "audit", "impl"} {
		want := phase + "|" + root + "|" + bin
		if phase == "audit" {
			want = "audit|" + root + "|" + bin
		}
		if !strings.Contains(string(mustReadNextCLI(t, filepath.Join(logs, "invocations"))), want) {
			t.Errorf("missing invocation %q", want)
		}
	}
	if count := cliGit(t, root, "rev-list", "--count", base+"..HEAD"); count != "3" {
		t.Errorf("phase commits=%s", count)
	}
	for _, item := range []struct{ rev, path, phase string }{{"HEAD~2", "changes/0001-next.md", "audit"}, {"HEAD~1", "changes/0001-next.md", "impl"}, {"HEAD", "changes/archive/0001-next.md", "applied"}} {
		doc := cliGit(t, root, "show", item.rev+":"+item.path)
		if !strings.Contains(doc, "phase: "+item.phase) {
			t.Errorf("%s change phase=%q; want %s", item.rev, doc, item.phase)
		}
	}
	archived := string(mustReadNextCLI(t, filepath.Join(root, "changes/archive/0001-next.md")))
	for _, marker := range []string{"phase: applied", "### ALP-002", "No findings", "- evals ", "- audit ", "- impl "} {
		if !strings.Contains(archived, marker) {
			t.Errorf("archive missing %q", marker)
		}
	}
	if _, err := os.Stat(filepath.Join(root, "changes/0001-next.md")); !os.IsNotExist(err) {
		t.Errorf("queued change remains: %v", err)
	}
	if spec := string(mustReadNextCLI(t, filepath.Join(root, "packages/alpha/SPEC.md"))); !strings.Contains(spec, "Alpha MUST return the new value.") {
		t.Errorf("SPEC missing applied item: %q", spec)
	}
	if claims := cliGit(t, root, "for-each-ref", "refs/fireweed/claims"); claims != "" {
		t.Errorf("claim retained: %s", claims)
	}
	for _, item := range []struct{ phase, result string }{{"evals", "fail"}, {"impl", "pass"}} {
		matches, err := filepath.Glob(filepath.Join(root, ".fireweed/runs/*/", item.phase, "alpha.json"))
		if err != nil || len(matches) != 1 {
			t.Errorf("%s reports=%v: %v", item.phase, matches, err)
			continue
		}
		body := string(mustReadNextCLI(t, matches[0]))
		if !strings.Contains(body, "ALP-002") || !strings.Contains(body, item.result) {
			t.Errorf("%s report=%q", item.phase, body)
		}
	}
}

func mustReadNextCLI(t *testing.T, path string) []byte {
	t.Helper()
	body, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return body
}

func Test_APP_075_NextAllowUnmergedAndUsage(t *testing.T) {
	bin := buildAppBinary(t)
	root, logs := nextCLIProject(t, bin, true, false)
	status, out, errOut := runNextCLI(t, bin, root, "next")
	nextCLIDiagnostic(t, status, 1, out, errOut, false, "P01")
	if _, err := os.Stat(filepath.Join(logs, "invocations")); !os.IsNotExist(err) {
		t.Errorf("unapproved work invoked agent: %v", err)
	}
	status, out, errOut = runNextCLI(t, bin, root, "next", "--allow-unmerged")
	if status != 0 || !strings.Contains(out, "next") {
		t.Errorf("allowed next exit=%d stdout=%q stderr=%q", status, out, errOut)
	}
	for _, args := range [][]string{{"next", "extra"}, {"next", "--bogus"}, {"modules", "--allow-unmerged"}} {
		status, out, errOut = runNextCLI(t, bin, root, args...)
		nextCLIDiagnostic(t, status, 2, out, errOut, false, "U01")
	}
}

func Test_APP_076_NextPlainEmptyQueue(t *testing.T) {
	bin := buildAppBinary(t)
	root, _ := nextCLIProject(t, bin, false, true)
	status, out, errOut := runNextCLI(t, bin, root, "next")
	if status != 0 || !strings.Contains(strings.ToLower(out), "empty") || errOut != "" {
		t.Errorf("empty next exit=%d stdout=%q stderr=%q", status, out, errOut)
	}
}

func Test_APP_077_NextExitClasses(t *testing.T) {
	bin := buildAppBinary(t)
	root, _ := nextCLIProject(t, bin, false, true)
	status, out, errOut := runNextCLI(t, bin, root, "next")
	if status != 0 {
		t.Errorf("empty queue exit=%d stdout=%q stderr=%q", status, out, errOut)
	}
	appPut(t, root, "fireweed.yaml", "modules: [broken\n")
	status, out, errOut = runNextCLI(t, bin, root, "next")
	nextCLIDiagnostic(t, status, 2, out, errOut, false, "C01")
	status, out, errOut = runNextCLI(t, bin, root, "next", "extra")
	nextCLIDiagnostic(t, status, 2, out, errOut, false, "U01")
	failed, _ := nextCLIProject(t, bin, true, false)
	status, out, errOut = runNextCLI(t, bin, failed, "next")
	nextCLIDiagnostic(t, status, 1, out, errOut, false, "P01")
}

func Test_APP_078_NextWrappedCauseAndAction(t *testing.T) {
	bin := buildAppBinary(t)
	root, _ := nextCLIProject(t, bin, true, true)
	appPut(t, root, ".fireweed/omit-new", "yes\n")
	cliCommit(t, root, "omit projected coverage")
	status, out, errOut := runNextCLI(t, bin, root, "next", "--json")
	nextCLIDiagnostic(t, status, 1, out, errOut, true, "C10")
	if !strings.Contains(out, "ALP-002") {
		t.Errorf("wrapped cause lost ID: %q", out)
	}
}

func Test_APP_079_NextJSONSuccessAndFailure(t *testing.T) {
	bin := buildAppBinary(t)
	root, _ := nextCLIProject(t, bin, false, true)
	status, out, errOut := runNextCLI(t, bin, root, "next", "--json")
	if status != 0 || errOut != "" {
		t.Fatalf("JSON empty exit=%d stdout=%q stderr=%q", status, out, errOut)
	}
	env := decodeEnvelope(t, out)
	var data struct {
		Change string `json:"change"`
		Run    string `json:"run"`
		Phase  string `json:"phase"`
		Empty  bool   `json:"empty"`
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(env["data"], &fields); err != nil {
		t.Errorf("empty data is not an object: %s: %v", env["data"], err)
	}
	for _, key := range []string{"change", "run", "phase", "empty"} {
		if _, ok := fields[key]; !ok {
			t.Errorf("empty data missing %q: %s", key, env["data"])
		}
	}
	if err := json.Unmarshal(env["data"], &data); err != nil || !data.Empty {
		t.Errorf("empty data=%s: %v", env["data"], err)
	}
	var diagnostics []json.RawMessage
	if err := json.Unmarshal(env["diagnostics"], &diagnostics); err != nil || len(diagnostics) != 0 {
		t.Errorf("success diagnostics=%s: %v", env["diagnostics"], err)
	}
	status, out, errOut = runNextCLI(t, bin, root, "next", "extra", "--json")
	nextCLIDiagnostic(t, status, 2, out, errOut, true, "U01")
	applied, _ := nextCLIProject(t, bin, true, true)
	status, out, errOut = runNextCLI(t, bin, applied, "next", "--json")
	if status != 0 || errOut != "" {
		t.Fatalf("JSON applied exit=%d stdout=%q stderr=%q", status, out, errOut)
	}
	env = decodeEnvelope(t, out)
	fields = nil
	if err := json.Unmarshal(env["data"], &fields); err != nil {
		t.Errorf("applied data is not an object: %s: %v", env["data"], err)
	}
	for _, key := range []string{"change", "run", "phase", "empty"} {
		if _, ok := fields[key]; !ok {
			t.Errorf("applied data missing %q: %s", key, env["data"])
		}
	}
	if err := json.Unmarshal(env["data"], &data); err != nil || data.Change == "" || data.Run == "" || data.Phase != "applied" || data.Empty {
		t.Errorf("applied data=%s: %v", env["data"], err)
	}
}
