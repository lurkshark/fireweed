package evals_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/app"
	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/project"
)

// Scratch doctor hooks invoke FIREWEED_BIN modules --json. In this file the
// invoking executable is the Go test binary, dispatched through public app.Run.
func TestMain(m *testing.M) {
	if os.Getenv("FIREWEED_APP_EVAL_CLI") == "1" {
		os.Exit(app.Run(os.Args[1:], os.Stdout, os.Stderr))
	}
	os.Exit(m.Run())
}

func ciCLIFixture(t *testing.T) (root, nested, base string) {
	t.Helper()
	outer := t.TempDir()
	appPut(t, outer, "fireweed.yaml", appFixtureConfig)
	appPut(t, outer, "packages/ancestor/SPEC.md", "---\nname: ancestor\nid_prefix: ANC\ndepends_on: []\n---\n# Ancestor\n\n## Specs\n\n### ANC-001\nThe ancestor MUST pass.\n")
	appPut(t, outer, "packages/ancestor/api.go", "package ancestor\n")
	root = filepath.Join(outer, "inner")
	log := filepath.Join(root, ".fireweed", "hooks.log")
	config := strings.Replace(appFixtureConfig, "base: main", "base: review-base", 1) + `verify_dependents: none
protected: [locked/**]
hooks:
  eval: [sh, .fireweed/eval.sh, "{module}", "{report}"]
  boundaries: [sh, .fireweed/boundaries.sh]
`
	appPut(t, root, "fireweed.yaml", config)
	appPut(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST pass.\n")
	appPut(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n\n### BET-001\nBeta MUST pass.\n")
	appPut(t, root, "packages/alpha/api.go", "package alpha\n")
	appPut(t, root, "packages/beta/api.go", "package beta\n")
	appPut(t, root, "locked/keep.txt", "initial\n")
	appPut(t, root, ".fireweed/eval.sh", fmt.Sprintf(`#!/bin/sh
printf 'eval:%%s:%%s\n' "$1" "$FIREWEED_BIN" >> %q
case "$1" in alpha) id=ALP-001;; beta) id=BET-001;; *) exit 2;; esac
mode=$(cat ".fireweed/$1.mode" 2>/dev/null)
case "$mode" in
  missing) printf '<testsuite/>' > "$2";;
  failing) printf '<testsuite><testcase name="%%s"><failure message="fixture failed"/></testcase></testsuite>' "$id" > "$2"; exit 1;;
  *) printf '<testsuite><testcase name="%%s"/></testsuite>' "$id" > "$2";;
esac
`, log))
	appPut(t, root, ".fireweed/boundaries.sh", fmt.Sprintf("#!/bin/sh\nprintf 'boundaries:%%s\\n' \"$FIREWEED_BIN\" >> %q\nif test -f .fireweed/boundary-fail; then printf 'beta: configured boundary violation\\n'; fi\n", log))
	cliGit(t, root, "init", "-q", "-b", "main")
	cliGit(t, root, "config", "user.name", "CI Fixture")
	cliGit(t, root, "config", "user.email", "ci-fixture@example.invalid")
	cliCommit(t, root, "initial project")
	cliGit(t, root, "checkout", "-qb", "review-base")
	appPut(t, root, "notes/base.txt", "base\n")
	cliCommit(t, root, "review base")
	base = cliGit(t, root, "rev-parse", "HEAD")
	cliGit(t, root, "checkout", "-qb", "feature")
	appPut(t, root, "notes/feature.txt", "feature\n")
	cliCommit(t, root, "ordinary feature note")
	for _, d := range project.Check(root) {
		if !d.Warning {
			t.Fatalf("invalid CI fixture static check: %+v", d)
		}
	}
	if report, err := checks.History(root, base+"..HEAD"); err != nil || len(report.Diagnostics) != 0 || len(report.Commits) != 1 {
		t.Fatalf("invalid CI fixture history: %+v, %v", report, err)
	}
	nested = appNestedDir(t, root)
	return root, nested, base
}

func ciCLISuccess(t *testing.T, output string) string {
	t.Helper()
	envelope := decodeEnvelope(t, output)
	var confirmation string
	if err := json.Unmarshal(envelope["data"], &confirmation); err != nil || strings.TrimSpace(confirmation) == "" {
		t.Fatalf("CI confirmation=%s: %v", envelope["data"], err)
	}
	return confirmation
}

func ciCLIDiagnostics(t *testing.T, output string) []struct {
	Code, Message string
	Warning       bool
} {
	t.Helper()
	envelope := decodeEnvelope(t, output)
	var diagnostics []struct {
		Code, Message string
		Warning       bool
	}
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil {
		t.Fatalf("CI diagnostics=%s: %v", envelope["diagnostics"], err)
	}
	return diagnostics
}

func ciCLIFailure(t *testing.T, nested, code string, status int, args ...string) string {
	t.Helper()
	got, output, stderr := runSavedAttemptApp(t, nested, args...)
	if got != status || stderr != "" {
		t.Errorf("ci %v: status=%d stdout=%q stderr=%q; want %d and empty stderr", args, got, output, stderr, status)
	}
	assertFailedJSON(t, output)
	found := false
	for _, d := range ciCLIDiagnostics(t, output) {
		if d.Code == code {
			found = true
			if !corrective(d.Message) {
				t.Errorf("CI %s diagnostic lacks corrective action: %q", code, d.Message)
			}
		}
	}
	if !found {
		t.Errorf("CI failure lacks %s diagnostic: %s", code, output)
	}
	return output
}

func Test_APP_109_APP_110_APP_112_APP_113_CIHistoryAndAllModules(t *testing.T) {
	root, nested, base := ciCLIFixture(t)
	status, plain, stderr := runSavedAttemptApp(t, nested, "ci", "--no-doctor")
	if status != 0 || stderr != "" || strings.TrimSpace(plain) == "" {
		t.Errorf("configured-base CI: status=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	log := doctorCLIFile(t, filepath.Join(root, ".fireweed", "hooks.log"))
	for _, marker := range []string{"eval:alpha:", "eval:beta:", "boundaries:"} {
		if !strings.Contains(log, marker) {
			t.Errorf("CI omitted %s with verify_dependents none: %q", marker, log)
		}
	}
	status, output, stderr := runSavedAttemptApp(t, nested, "ci", base+"..HEAD", "--no-doctor", "--json")
	if status != 0 || stderr != "" || ciCLISuccess(t, output) != strings.TrimSpace(plain) {
		t.Errorf("explicit-range JSON CI: status=%d stdout=%q stderr=%q", status, output, stderr)
	}
	if diagnostics := ciCLIDiagnostics(t, output); len(diagnostics) != 0 {
		t.Errorf("clean CI diagnostics=%+v", diagnostics)
	}

	pending := "changes/0001-work.md"
	appPut(t, root, pending, strings.Replace(checkCommitsChange, "%s", "phase: impl\n", 1))
	cliCommit(t, root, "ready for application")
	if err := os.MkdirAll(filepath.Join(root, "changes", "archive"), 0755); err != nil {
		t.Fatal(err)
	}
	cliGit(t, root, "mv", pending, "changes/archive/0001-work.md")
	appPut(t, root, "changes/archive/0001-work.md", strings.Replace(checkCommitsChange, "%s", "phase: applied\n", 1))
	appPut(t, root, "locked/keep.txt", "changed in application\n")
	cliCommit(t, root, "application with protected edit")
	if report, err := checks.History(root, base+"..HEAD"); err != nil || len(report.Diagnostics) == 0 || report.Diagnostics[0].Code != "H01" {
		t.Fatalf("history failure fixture: %+v, %v", report, err)
	}
	ciCLIFailure(t, nested, "H01", 1, "ci", "--no-doctor", "--json")
	status, output, stderr = runSavedAttemptApp(t, nested, "ci", "HEAD..HEAD", "--no-doctor", "--json")
	if status != 0 || stderr != "" {
		t.Errorf("explicit empty range should exclude earlier violation: status=%d stdout=%q stderr=%q", status, output, stderr)
	} else {
		ciCLISuccess(t, output)
	}
}

func Test_APP_110_APP_112_APP_113_CIEvalAndBoundaryFailures(t *testing.T) {
	root, nested, _ := ciCLIFixture(t)
	appPut(t, root, ".fireweed/beta.mode", "missing\n")
	ciCLIFailure(t, nested, "C10", 1, "ci", "--no-doctor", "--json")
	appPut(t, root, ".fireweed/beta.mode", "failing\n")
	ciCLIFailure(t, nested, "E01", 1, "ci", "--no-doctor", "--json")
	status, plain, stderr := runSavedAttemptApp(t, nested, "ci", "--no-doctor")
	if status != 1 || !strings.Contains(plain+stderr, "E01") || !corrective(plain+stderr) {
		t.Errorf("plain failing eval: status=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	appPut(t, root, ".fireweed/beta.mode", "pass\n")
	appPut(t, root, ".fireweed/boundary-fail", "yes\n")
	ciCLIFailure(t, nested, "X01", 1, "ci", "--no-doctor", "--json")
}

func Test_APP_109_APP_112_APP_113_CIStaticUsageAndConfiguration(t *testing.T) {
	root, nested, _ := ciCLIFixture(t)
	for _, args := range [][]string{{"ci", "one", "two", "--no-doctor", "--json"}, {"ci", "--unexpected", "--json"}} {
		ciCLIFailure(t, nested, "U01", 2, args...)
	}
	appPut(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n\n### BET-001\nBeta MUST pass.\n\n### BET-001\nDuplicate MUST fail.\n")
	ciCLIFailure(t, nested, "C05", 1, "ci", "--no-doctor", "--json")
	appPut(t, root, "fireweed.yaml", "modules: [broken\n")
	ciCLIFailure(t, nested, "C01", 2, "ci", "--no-doctor", "--json")
}

func Test_APP_111_APP_112_APP_113_CIDoctorCanariesNoAgentAndSkip(t *testing.T) {
	t.Setenv("FIREWEED_APP_EVAL_CLI", "1")
	bin, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	root, nested, log := doctorCLIFixture(t, bin)
	appPut(t, root, "packages/alpha/evals/alpha_test.go", "package alpha\nimport \"testing\"\nfunc Test_ALP_001_Fixture(t *testing.T) {}\n")
	appPut(t, root, "AGENTS.md", "<!-- fireweed:start -->\nstale\n<!-- fireweed:end -->\n")
	status, output, stderr := runSavedAttemptApp(t, nested, "ci", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("default doctor CI: status=%d stdout=%q stderr=%q", status, output, stderr)
	}
	ciCLISuccess(t, output)
	foundWarning := false
	for _, d := range ciCLIDiagnostics(t, output) {
		if d.Code == "C16" && d.Warning && strings.Contains(d.Message, "AGENTS.md") {
			foundWarning = true
		}
	}
	if !foundWarning {
		t.Errorf("CI lost doctor warning: %s", output)
	}
	if got := doctorCLIFile(t, log); !strings.Contains(got, "setup:") || !strings.Contains(got, bin) || strings.Contains(got, "worker:") || strings.Contains(got, "audit:") {
		t.Errorf("CI doctor did not run canaries with NoAgent: %q", got)
	}
	doctorCLIConfig(t, root, func(config map[string]any) {
		config["hooks"].(map[string]any)["boundaries"] = "sh -c 'exit 0'"
	})
	canaryOutput := ciCLIFailure(t, nested, "C14", 1, "ci", "--json")
	foundMissedCanary := false
	for _, d := range ciCLIDiagnostics(t, canaryOutput) {
		if d.Code == "C14" && (strings.Contains(d.Message, "undeclared") || strings.Contains(d.Message, "internal_eval")) {
			foundMissedCanary = true
		}
	}
	if !foundMissedCanary {
		t.Errorf("default CI did not report a missed boundary canary: %s", canaryOutput)
	}
	appPut(t, root, ".fireweed/setup.sh", "#!/bin/sh\necho named-ci-canary-failure >&2\nexit 17\n")
	before := doctorCLIFile(t, log)
	status, output, stderr = runSavedAttemptApp(t, nested, "ci", "--no-doctor", "--json")
	if status != 0 || stderr != "" {
		t.Errorf("--no-doctor CI: status=%d stdout=%q stderr=%q", status, output, stderr)
	} else {
		ciCLISuccess(t, output)
		for _, d := range ciCLIDiagnostics(t, output) {
			if d.Code == "C16" {
				t.Errorf("--no-doctor retained doctor warning: %+v", d)
			}
		}
	}
	if after := doctorCLIFile(t, log); strings.Count(after, "setup:") != strings.Count(before, "setup:") || strings.Contains(after, "worker:") || strings.Contains(after, "audit:") {
		t.Errorf("--no-doctor ran doctor or agents: before=%q after=%q", before, after)
	}
}
