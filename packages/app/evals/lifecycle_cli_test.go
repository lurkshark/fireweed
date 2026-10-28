package evals_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func runLifecycleBinary(t *testing.T, bin, dir string, args ...string) (int, string, string) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Dir = dir
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err := cmd.Run()
	if ctx.Err() != nil {
		t.Fatalf("NEW CLI %v exceeded 8 seconds: %v", args, ctx.Err())
	}
	if err == nil {
		return 0, stdout.String(), stderr.String()
	}
	var exit *exec.ExitError
	if !errors.As(err, &exit) {
		t.Fatalf("run NEW CLI %v: %v", args, err)
	}
	return exit.ExitCode(), stdout.String(), stderr.String()
}

// lifecycleCLIProject runs the installed NEW executable against a disposable Git
// project. Hook reports name the exact current and projected spec IDs.
func lifecycleCLIProject(t *testing.T, bin string, audit bool) string {
	t.Helper()
	root := t.TempDir()
	appPut(t, root, ".gitignore", ".fireweed/runs/\n")
	appPut(t, root, "SYSTEM.md", "# Fixture\nLifecycle CLI fixture.\n")
	config := appFixtureConfig + `hooks:
  eval: [sh, .fireweed/eval.sh, "{module}", "{report}"]
  boundaries: [sh, .fireweed/boundaries.sh]
`
	if !audit {
		config += "audit:\n  enabled: false\n"
	}
	appPut(t, root, "fireweed.yaml", config)
	appPut(t, root, ".fireweed/expected-bin", bin+"\n")
	appPut(t, root, ".fireweed/eval.sh", `#!/bin/sh
module="$1"
report="$2"
expected=$(cat .fireweed/expected-bin)
if [ ! "$FIREWEED_BIN" -ef "$expected" ]; then
  printf 'FIREWEED_BIN mismatch: %s expected %s\n' "$FIREWEED_BIN" "$expected" >&2
  exit 3
fi
if [ "$module" = alpha ]; then
  result=''
  exitcode=0
  if [ ! -f packages/alpha/internal/ready ] && [ ! -f .fireweed/runs/pass-new ] && [ -f .fireweed/emit-new ]; then
    result='<failure message="expected before implementation"/>'
    exitcode=1
  fi
  extra=''
  if [ -f .fireweed/emit-new ]; then
    extra="<testcase name=\"ALP-002\" classname=\"fixture\">$result</testcase>"
  fi
  printf '<testsuite><testcase name="ALP-001" classname="fixture"/>%s</testsuite>' "$extra" > "$report"
  exit "$exitcode"
fi
if [ "$module" = beta ]; then
  result=''; exitcode=0
  if [ -f .fireweed/runs/beta-fail ]; then result='<failure message="dependent failed"/>'; exitcode=1; fi
  printf '<testsuite><testcase name="BET-001" classname="fixture">%s</testcase></testsuite>' "$result" > "$report"
  exit "$exitcode"
fi
exit 2
`)
	appPut(t, root, ".fireweed/boundaries.sh", "#!/bin/sh\nexit 0\n")
	appPut(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST retain the existing value.\n")
	appPut(t, root, "packages/alpha/api.go", "package alpha\n")
	appPut(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n")
	appPut(t, root, "changes/0001-manual.md", "---\nmodule: alpha\nset:\n  interface: [api.go, extra.go]\n---\n\n## Motivation\nUsers need an additional value.\n\n## Interface\nAlpha exposes the additional value.\n\n## Added\n### new:extra\nAlpha MUST return the additional value.\n\n## Modified\n\n## Removed\n\n## Decisions\n")
	cliGit(t, root, "init")
	cliGit(t, root, "checkout", "-b", "main")
	cliCommit(t, root, "approved fixture")
	for _, diagnostic := range project.Check(root) {
		if !diagnostic.Warning {
			t.Fatalf("invalid lifecycle fixture: %#v", diagnostic)
		}
	}
	return root
}

func cliHold(t *testing.T, root string) {
	t.Helper()
	state, err := workflow.Start(root, "manual")
	if err != nil || state.Phase != "evals" {
		t.Fatalf("fixture claim: %#v, %v", state, err)
	}
}

func cliGit(t *testing.T, root string, args ...string) string {
	t.Helper()
	cmd := exec.Command("git", args...)
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_AUTHOR_NAME=Fixture", "GIT_AUTHOR_EMAIL=fixture@example.invalid", "GIT_COMMITTER_NAME=Fixture", "GIT_COMMITTER_EMAIL=fixture@example.invalid")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v: %s", args, err, out)
	}
	return strings.TrimSpace(string(out))
}

func cliCommit(t *testing.T, root, message string) {
	t.Helper()
	cliGit(t, root, "add", "-A")
	cliGit(t, root, "commit", "-m", message)
}

func cliPhase(t *testing.T, root, want string) {
	t.Helper()
	state, err := workflow.Current(root, "manual")
	if err != nil || state.Phase != want {
		t.Fatalf("held phase = %q, %v; want %q", state.Phase, err, want)
	}
}

func cliString(t *testing.T, output string) string {
	t.Helper()
	envelope := decodeEnvelope(t, output)
	var value string
	if err := json.Unmarshal(envelope["data"], &value); err != nil {
		t.Fatalf("data=%s; want string: %v", envelope["data"], err)
	}
	return value
}

func cliJSONResult(t *testing.T, output string) map[string]json.RawMessage {
	t.Helper()
	envelope := decodeEnvelope(t, output)
	if string(envelope["data"]) == "null" {
		t.Fatalf("result data is null: %s", output)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(envelope["data"], &result); err != nil || result == nil {
		t.Fatalf("result data=%s: %v", envelope["data"], err)
	}
	return result
}

func cliResultPhase(t *testing.T, output, phase string) {
	t.Helper()
	envelope := decodeEnvelope(t, output)
	var result struct {
		State struct {
			Phase string `json:"Phase"`
		} `json:"State"`
	}
	if err := json.Unmarshal(envelope["data"], &result); err != nil || result.State.Phase != phase {
		t.Errorf("result phase=%s; want %s: %v", envelope["data"], phase, err)
	}
}

func lifecycleCommandFailure(t *testing.T, status, want int, stdout, stderr string, jsonMode bool) {
	t.Helper()
	message := commandFailure(t, status, want, stdout, stderr, jsonMode)
	if !corrective(message) {
		t.Errorf("diagnostic %q must suggest a correction", message)
	}
}

func Test_APP_009_APP_027_APP_028_StartContextAssignmentAndRejection(t *testing.T) {
	bin := buildAppBinary(t)
	root := lifecycleCLIProject(t, bin, true)
	before := cliGit(t, root, "rev-parse", "HEAD")
	status, plain, stderr := runLifecycleBinary(t, bin, appNestedDir(t, root), "start", "manual")
	if status != 0 || stderr != "" || !strings.Contains(plain, "# Write evals") || !strings.Contains(plain, "ALP-002") {
		t.Fatalf("start: exit=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	if after := cliGit(t, root, "rev-parse", "HEAD"); after != before {
		t.Errorf("start committed: %s -> %s", before, after)
	}
	state, err := workflow.Current(root, "manual")
	if err != nil || state.Phase != "evals" || state.Run == "" {
		t.Fatalf("start state=%#v, %v", state, err)
	}
	change, err := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if err != nil || !strings.Contains(string(change), "run:") || !strings.Contains(string(change), "phase: evals") || !strings.Contains(string(change), "### ALP-002") || strings.Contains(string(change), "new:extra") {
		t.Errorf("assigned change=%q, %v", change, err)
	}
	expected, err := workflow.Context(root, "manual", false)
	if err != nil || plain != expected {
		t.Errorf("first-step context differs from workflow.Context: %v", err)
	}
	for _, tc := range []struct {
		args []string
		want int
	}{
		{[]string{"start"}, 2},
		{[]string{"start", "manual", "extra"}, 2},
		{[]string{"start", "missing"}, 1},
		{[]string{"start", "manual"}, 1},
	} {
		args := append(append([]string{}, tc.args...), "--json")
		status, out, errOut := runLifecycleBinary(t, bin, root, args...)
		lifecycleCommandFailure(t, status, tc.want, out, errOut, true)
	}
	bad := lifecycleCLIProject(t, bin, false)
	appPut(t, bad, "fireweed.yaml", "modules: [broken\n")
	status, out, errOut := runLifecycleBinary(t, bin, bad, "start", "manual", "--json")
	lifecycleCommandFailure(t, status, 2, out, errOut, true)
}

func Test_APP_009_APP_027_APP_028_APP_040_StartPreservesApprovalWarning(t *testing.T) {
	bin := buildAppBinary(t)
	for _, jsonMode := range []bool{false, true} {
		root := lifecycleCLIProject(t, bin, false)
		cliGit(t, root, "checkout", "-b", "feature")
		cliGit(t, root, "checkout", "main")
		cliGit(t, root, "rm", "changes/0001-manual.md")
		cliCommit(t, root, "base without queued change")
		cliGit(t, root, "update-ref", "refs/remotes/origin/main", cliGit(t, root, "rev-parse", "HEAD"))
		cliGit(t, root, "checkout", "feature")
		args := []string{"start", "manual"}
		if jsonMode {
			args = append(args, "--json")
		}
		status, output, stderr := runLifecycleBinary(t, bin, root, args...)
		if status != 0 {
			t.Fatalf("start with approval warning (JSON=%v): %d %q %q", jsonMode, status, output, stderr)
		}
		if jsonMode {
			if stderr != "" || !strings.Contains(cliString(t, output), "# Write evals") {
				t.Errorf("JSON start omitted context or wrote stderr: %q %q", output, stderr)
			}
			var diagnostics []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
				Warning bool   `json:"warning"`
			}
			if err := json.Unmarshal(decodeEnvelope(t, output)["diagnostics"], &diagnostics); err != nil || len(diagnostics) == 0 || diagnostics[0].Code != "P04" || diagnostics[0].Message == "" || !diagnostics[0].Warning {
				t.Errorf("approval warning missing: %s, %v", output, err)
			}
		} else if !strings.Contains(output, "# Write evals") || !strings.Contains(output+stderr, "P04") {
			t.Errorf("plain start omitted context or P04 warning: %q %q", output, stderr)
		}
	}
}

func Test_APP_004_HeldModulesUseApprovedProjection(t *testing.T) {
	bin := buildAppBinary(t)
	root := lifecycleCLIProject(t, bin, false)
	appPut(t, root, ".fireweed/emit-new", "yes\n")
	cliCommit(t, root, "projected report")
	cliHold(t, root)
	changePath := filepath.Join(root, "changes/0001-manual.md")
	change, err := os.ReadFile(changePath)
	if err != nil {
		t.Fatal(err)
	}
	appPut(t, root, "changes/0001-manual.md", strings.Replace(string(change), "interface: [api.go, extra.go]", "interface: [api.go, unapproved.go]", 1))
	var status int
	var stderr string
	status, plain, stderr := runLifecycleBinary(t, bin, appNestedDir(t, root), "modules")
	if status != 0 || stderr != "" || !strings.Contains(plain, "extra.go") || strings.Contains(plain, "unapproved.go") {
		t.Errorf("held plain graph: exit=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	status, output, stderr := runLifecycleBinary(t, bin, root, "modules", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("held JSON graph: exit=%d stdout=%q stderr=%q", status, output, stderr)
	}
	envelope := decodeEnvelope(t, output)
	var modules []struct {
		Name  string `json:"name"`
		Zones struct {
			Interface []string `json:"interface"`
		} `json:"zones"`
	}
	if err := json.Unmarshal(envelope["data"], &modules); err != nil || len(modules) != 2 || modules[0].Name != "alpha" || !strings.Contains(strings.Join(modules[0].Zones.Interface, ","), "extra.go") || strings.Contains(strings.Join(modules[0].Zones.Interface, ","), "unapproved.go") {
		t.Errorf("held JSON graph=%s, %v", envelope["data"], err)
	}
	var diagnostics []json.RawMessage
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || len(diagnostics) != 0 {
		t.Errorf("held diagnostics=%s, %v", envelope["diagnostics"], err)
	}
}

func Test_APP_011_APP_032_APP_033_ContextCompactFullAndErrors(t *testing.T) {
	bin := buildAppBinary(t)
	root := lifecycleCLIProject(t, bin, false)
	appPut(t, root, "packages/alpha/internal/context-note.txt", "Full context marker.\n")
	cliCommit(t, root, "context content")
	cliHold(t, root)
	for _, full := range []bool{false, true} {
		args := []string{"context", "manual"}
		if full {
			args = append(args, "--full")
		}
		status, plain, stderr := runLifecycleBinary(t, bin, appNestedDir(t, root), args...)
		if status != 0 || stderr != "" {
			t.Fatalf("context %v: %d %q", full, status, stderr)
		}
		expected, err := workflow.Context(root, "manual", full)
		if err != nil || plain != expected {
			t.Errorf("context %v differs from public workflow: %v", full, err)
		}
		status, output, stderr := runLifecycleBinary(t, bin, root, append(args, "--json")...)
		if status != 0 || stderr != "" || cliString(t, output) != plain {
			t.Errorf("context JSON %v: %d %q %q", full, status, output, stderr)
		}
		var diagnostics []json.RawMessage
		diagnosticJSON := decodeEnvelope(t, output)["diagnostics"]
		if err := json.Unmarshal(diagnosticJSON, &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
			t.Errorf("context diagnostics=%s; want empty array: %v", diagnosticJSON, err)
		}
		if full && !strings.Contains(plain, "Full context marker.") || !full && strings.Contains(plain, "Full context marker.") {
			t.Errorf("context full selection: %v", full)
		}
	}
	for _, tc := range []struct {
		args []string
		want int
	}{
		{[]string{"context"}, 2}, {[]string{"context", "manual", "extra"}, 2}, {[]string{"context", "missing"}, 1},
	} {
		status, out, errOut := runLifecycleBinary(t, bin, root, append(tc.args, "--json")...)
		lifecycleCommandFailure(t, status, tc.want, out, errOut, true)
	}
	bad := lifecycleCLIProject(t, bin, false)
	appPut(t, bad, "fireweed.yaml", "modules: [broken\n")
	status, out, errOut := runLifecycleBinary(t, bin, bad, "context", "manual", "--json")
	lifecycleCommandFailure(t, status, 2, out, errOut, true)
}

func Test_APP_012_APP_034_APP_035_APP_036_CoverageCurrentProjectionAndErrors(t *testing.T) {
	bin := buildAppBinary(t)
	root := lifecycleCLIProject(t, bin, false)
	status, plain, stderr := runLifecycleBinary(t, bin, root, "coverage", "alpha")
	if status != 0 || stderr != "" || !strings.Contains(plain, "ALP-001") {
		t.Errorf("plain current coverage: %d %q %q", status, plain, stderr)
	}
	status, output, stderr := runLifecycleBinary(t, bin, root, "coverage", "alpha", "--json")
	if status != 0 || stderr != "" || !strings.Contains(output, "ALP-001") {
		t.Errorf("current coverage: %d %q %q", status, output, stderr)
	}
	cliJSONResult(t, output)
	status, out, errOut := runLifecycleBinary(t, bin, root, "coverage", "alpha", "--with-change", "manual", "--json")
	lifecycleCommandFailure(t, status, 1, out, errOut, true) // projected ALP-002 is uncovered
	appPut(t, root, ".fireweed/emit-new", "yes\n")
	cliCommit(t, root, "projected JUnit case")
	cliHold(t, root)
	status, out, errOut = runLifecycleBinary(t, bin, root, "coverage", "alpha", "--with-change", "manual", "--json")
	lifecycleCommandFailure(t, status, 1, out, errOut, true) // mapped ALP-002 fails before implementation
	appPut(t, root, "packages/alpha/internal/ready", "ready\n")
	status, output, stderr = runLifecycleBinary(t, bin, root, "coverage", "alpha", "--with-change", "manual", "--json")
	if status != 0 || stderr != "" || !strings.Contains(output, "ALP-002") {
		t.Errorf("projected coverage: %d %q %q", status, output, stderr)
	}
	cliJSONResult(t, output)
	held := lifecycleCLIProject(t, bin, false)
	appPut(t, held, ".fireweed/emit-new", "yes\n")
	appPut(t, held, ".fireweed/no-report.sh", "#!/bin/sh\nexit 7\n")
	cliCommit(t, held, "projected JUnit case")
	cliHold(t, held)
	appPut(t, held, "packages/alpha/internal/ready", "ready\n")
	changePath := filepath.Join(held, "changes/0001-manual.md")
	change, err := os.ReadFile(changePath)
	if err != nil {
		t.Fatal(err)
	}
	appPut(t, held, "changes/0001-manual.md", strings.Replace(string(change), "  interface: [api.go, extra.go]\n", "  interface: [api.go, extra.go]\n  hooks:\n    eval: [sh, .fireweed/no-report.sh]\n", 1))
	if _, err := project.Project(held, "manual"); err != nil {
		t.Fatalf("live hook override must be a valid projection: %v", err)
	}
	status, output, stderr = runLifecycleBinary(t, bin, held, "coverage", "alpha", "--with-change", "manual", "--json")
	if status != 0 || stderr != "" || !strings.Contains(output, "ALP-002") {
		t.Errorf("held approved coverage: %d %q %q", status, output, stderr)
	}
	cliJSONResult(t, output)
	for _, args := range [][]string{{"coverage"}, {"coverage", "missing"}, {"coverage", "alpha", "--with-change", "missing"}, {"coverage", "alpha", "extra"}} {
		status, out, errOut := runLifecycleBinary(t, bin, root, append(args, "--json")...)
		lifecycleCommandFailure(t, status, 2, out, errOut, true)
	}
	bad := lifecycleCLIProject(t, bin, false)
	appPut(t, bad, "fireweed.yaml", "modules: [broken\n")
	status, out, errOut = runLifecycleBinary(t, bin, bad, "coverage", "alpha", "--json")
	lifecycleCommandFailure(t, status, 2, out, errOut, true)
}

func Test_APP_010_APP_029_APP_030_APP_031_VerifyDoneAuditAndFullChecks(t *testing.T) {
	bin := buildAppBinary(t)
	root := lifecycleCLIProject(t, bin, true)
	appPut(t, root, ".fireweed/emit-new", "yes\n")
	cliCommit(t, root, "failing projected case")
	cliHold(t, root)
	appPut(t, root, "packages/alpha/extra.go", "package alpha\n")
	var status int
	var stderr string
	before := cliGit(t, root, "rev-parse", "HEAD")
	status, output, stderr := runLifecycleBinary(t, bin, root, "verify", "manual", "--quick", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("evals quick verify: %d %q %q", status, output, stderr)
	}
	cliResultPhase(t, output, "evals")
	status, plain, stderr := runLifecycleBinary(t, bin, root, "verify", "manual", "--quick")
	if status != 0 || stderr != "" || !strings.Contains(strings.ToLower(plain), "evals") {
		t.Errorf("plain evals verify: %d %q %q", status, plain, stderr)
	}
	if cliGit(t, root, "rev-parse", "HEAD") != before {
		t.Error("verify committed")
	}
	status, plain, stderr = runLifecycleBinary(t, bin, root, "done", "manual")
	if status != 0 || stderr != "" || !strings.Contains(strings.ToLower(plain), "audit") {
		t.Fatalf("evals done: %d %q %q", status, plain, stderr)
	}
	cliPhase(t, root, "audit")
	changePath := filepath.Join(root, "changes/0001-manual.md")
	change, err := os.ReadFile(changePath)
	if err != nil {
		t.Fatal(err)
	}
	appPut(t, root, "changes/0001-manual.md", strings.Replace(string(change), "## Decisions", "## Audit\nChecked IDs: ALP-002\nNo findings\n\n## Decisions", 1))
	status, output, stderr = runLifecycleBinary(t, bin, root, "done", "manual", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("audit done: %d %q %q", status, output, stderr)
	}
	cliResultPhase(t, output, "impl")
	cliPhase(t, root, "impl")
	appPut(t, root, "packages/alpha/internal/ready", "ready\n")
	appPut(t, root, ".fireweed/runs/beta-fail", "yes\n")
	before = cliGit(t, root, "rev-parse", "HEAD")
	status, output, stderr = runLifecycleBinary(t, bin, root, "verify", "--quick", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("impl quick verify: %d %q %q", status, output, stderr)
	}
	cliResultPhase(t, output, "impl")
	status, plain, stderr = runLifecycleBinary(t, bin, root, "verify", "--quick")
	if status != 0 || stderr != "" || !strings.Contains(strings.ToLower(plain), "impl") {
		t.Errorf("plain impl verify: %d %q %q", status, plain, stderr)
	}
	status, out, errOut := runLifecycleBinary(t, bin, root, "done", "manual", "--json")
	commandFailure(t, status, 1, out, errOut, true) // done must include dependent beta
	if cliGit(t, root, "rev-parse", "HEAD") != before {
		t.Error("failed done committed")
	}
	cliPhase(t, root, "impl")
	if err := os.Remove(filepath.Join(root, ".fireweed/runs/beta-fail")); err != nil {
		t.Fatal(err)
	}
	// Omitted change must resolve the held change and complete its full checks.
	status, plain, stderr = runLifecycleBinary(t, bin, root, "done")
	if status != 0 || stderr != "" || !strings.Contains(strings.ToLower(plain), "applied") {
		t.Fatalf("impl done: %d %q %q", status, plain, stderr)
	}
	if _, err := os.Stat(filepath.Join(root, "changes/archive/0001-manual.md")); err != nil {
		t.Errorf("missing archive: %v", err)
	}
}

func Test_APP_010_APP_029_APP_030_APP_031_VerifyRejectsFalseFailFirstAndFrozenEdits(t *testing.T) {
	bin := buildAppBinary(t)
	root := lifecycleCLIProject(t, bin, false)
	appPut(t, root, ".fireweed/emit-new", "yes\n")
	cliCommit(t, root, "projected case")
	cliHold(t, root)
	appPut(t, root, "packages/alpha/extra.go", "package alpha\n")
	var status int
	appPut(t, root, ".fireweed/runs/pass-new", "yes\n")
	status, out, errOut := runLifecycleBinary(t, bin, root, "verify", "manual", "--quick", "--json")
	commandFailure(t, status, 1, out, errOut, true)
	if err := os.Remove(filepath.Join(root, ".fireweed/runs/pass-new")); err != nil {
		t.Fatal(err)
	}
	status, _, _ = runLifecycleBinary(t, bin, root, "done", "manual")
	if status != 0 {
		t.Fatal("evals done failed")
	}
	cliPhase(t, root, "impl")
	appPut(t, root, "packages/alpha/internal/ready", "ready\n")
	appPut(t, root, "packages/alpha/evals/new_test.go", "package evals\n")
	status, out, errOut = runLifecycleBinary(t, bin, root, "verify", "manual", "--quick", "--json")
	commandFailure(t, status, 1, out, errOut, true)
	if err := os.Remove(filepath.Join(root, "packages/alpha/evals/new_test.go")); err != nil {
		t.Fatal(err)
	}
	appPut(t, root, "packages/alpha/api.go", "package alpha // forbidden edit\n")
	status, out, errOut = runLifecycleBinary(t, bin, root, "verify", "manual", "--quick", "--json")
	commandFailure(t, status, 1, out, errOut, true)
	for _, args := range [][]string{{"verify", "missing"}, {"done", "missing"}} {
		status, out, errOut = runLifecycleBinary(t, bin, root, append(args, "--json")...)
		commandFailure(t, status, 1, out, errOut, true)
	}
	for _, args := range [][]string{{"verify", "manual", "extra"}, {"done", "manual", "extra"}} {
		status, out, errOut = runLifecycleBinary(t, bin, root, append(args, "--json")...)
		commandFailure(t, status, 2, out, errOut, true)
	}
	bad := lifecycleCLIProject(t, bin, false)
	appPut(t, bad, "fireweed.yaml", "modules: [broken\n")
	for _, command := range []string{"verify", "done"} {
		status, out, errOut = runLifecycleBinary(t, bin, bad, command, "manual", "--json")
		commandFailure(t, status, 2, out, errOut, true)
	}
}

func Test_APP_010_APP_029_APP_031_DoneReportsBlockedAudit(t *testing.T) {
	bin := buildAppBinary(t)
	root := lifecycleCLIProject(t, bin, true)
	appPut(t, root, ".fireweed/emit-new", "yes\n")
	cliCommit(t, root, "failing projected case")
	cliHold(t, root)
	appPut(t, root, "packages/alpha/extra.go", "package alpha\n")
	result, err := workflow.Done(context.Background(), workflow.Request{Root: root, Change: "manual", Bin: bin, Timeout: 2 * time.Second})
	if err != nil || result.State.Phase != "audit" {
		t.Fatalf("fixture audit phase: %#v, %v", result.State, err)
	}
	changePath := filepath.Join(root, "changes/0001-manual.md")
	change, err := os.ReadFile(changePath)
	if err != nil {
		t.Fatal(err)
	}
	appPut(t, root, "changes/0001-manual.md", strings.Replace(string(change), "## Decisions", "## Audit\n- [untestable] ALP-002: \"Alpha MUST return the additional value.\" — cannot observe the value\n\n## Decisions", 1))
	status, output, stderr := runLifecycleBinary(t, bin, root, "done", "manual", "--json")
	commandFailure(t, status, 1, output, stderr, true)
	blocked, err := os.ReadFile(changePath)
	if err != nil || !strings.Contains(string(blocked), "blocked:") {
		t.Errorf("blocked audit state missing: %q, %v", blocked, err)
	}
	status, output, stderr = runLifecycleBinary(t, bin, root, "verify", "--json")
	commandFailure(t, status, 1, output, stderr, true)
}
