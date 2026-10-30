package evals_test

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func prBodyAppRun(t *testing.T, dir string, args ...string) (int, string, string) {
	t.Helper()
	prior, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(prior); err != nil {
			t.Fatal(err)
		}
	}()
	return run(args...)
}

func prBodyAppFixture(t *testing.T) (root, nested, revisionRange, endpoint string, report workflow.PRBodyReport) {
	t.Helper()
	root = appProjectFixture(t)
	appPut(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001 [review]\nAlpha MUST retain the selected behavior.\n")
	cliGit(t, root, "init", "-q", "-b", "main")
	cliCommit(t, root, "PR body base")
	base := cliGit(t, root, "rev-parse", "HEAD")
	statusArchive(t, root, "0001-review", "alpha", "2026-10-30T12:00:00Z")
	appPut(t, root, "packages/alpha/internal/opaque.dat", "\x00\x01\x02\x03")
	appPut(t, root, "packages/alpha/internal/ordinary.txt", "ordinary addition\n")
	snapshot := workflow.CostSnapshot{
		Run: "pr-body-fixture", Phase: "impl", RecordedAt: time.Date(2026, 10, 30, 12, 0, 0, 0, time.UTC),
		Modules: []workflow.ModuleCost{{Module: "alpha", Seconds: 2.5, CasesPerSpec: map[string]int{"ALP-001": 4}, EvalLines: 15, DisposableLines: 7}},
		Diagnostics: []project.Diagnostic{
			{Code: "B01", Message: "SAVED B01 time", Warning: true},
			{Code: "B02", Message: "ALP-001: SAVED B02 cases", Warning: true},
			{Code: "B03", Message: "alpha: SAVED B03 ratio", Warning: true},
		},
	}
	data, err := json.Marshal(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	appPut(t, root, ".fireweed/runs/pr-body-fixture/impl/costs.json", string(data))
	cliCommit(t, root, "selected PR body endpoint")
	endpoint = cliGit(t, root, "rev-parse", "HEAD")
	revisionRange = base + ".." + endpoint
	appPut(t, root, "fireweed.yaml", strings.Replace(appFixtureConfig, `modules: ["packages/*"]`, "modules: [] # later HEAD", 1))
	cliCommit(t, root, "later HEAD config")
	appPut(t, root, "fireweed.yaml", "[invalid live worktree config\n")
	nested = appNestedDir(t, root)
	report, err = workflow.PRBody(root, revisionRange)
	if err != nil {
		t.Fatalf("public PRBody fixture: %v", err)
	}
	if !strings.Contains(report.Body, "0001-review") || !strings.Contains(report.Body, "opaque.dat") || !strings.Contains(report.Body, "SAVED B01") || strings.Contains(report.Body, "later HEAD") {
		t.Fatalf("public PRBody fixture omitted selected history or used later metadata: %s", report.Body)
	}
	for _, code := range []string{"B01", "B02", "B03", "B04"} {
		found := false
		for _, d := range report.Diagnostics {
			if d.Code == code && d.Warning {
				found = true
				if code == "B04" && d.Path != "packages/alpha/internal/opaque.dat" {
					t.Errorf("B04 path=%q; want opaque artifact path", d.Path)
				}
			}
		}
		if !found {
			t.Fatalf("public PRBody fixture lacks %s warning: %#v", code, report.Diagnostics)
		}
	}
	return
}

func prBodyAppSuccess(t *testing.T, dir, revisionRange string, report workflow.PRBodyReport) {
	t.Helper()
	plainStatus, plain, plainErr := prBodyAppRun(t, dir, "pr-body", revisionRange)
	if plainStatus != 0 {
		t.Errorf("plain status=%d stdout=%q stderr=%q; want 0", plainStatus, plain, plainErr)
	}
	if !strings.Contains(plain, report.Body) {
		t.Errorf("plain body differs from public PRBody: got %q want %q", plain, report.Body)
	}
	for _, diagnostic := range report.Diagnostics {
		if !strings.Contains(plain+plainErr, diagnostic.Code) || !strings.Contains(plain+plainErr, diagnostic.Message) {
			t.Errorf("plain output omits diagnostic %#v: stdout=%s stderr=%s", diagnostic, plain, plainErr)
		}
	}
	status, output, stderr := prBodyAppRun(t, dir, "pr-body", revisionRange, "--json")
	if status != 0 || stderr != "" {
		t.Errorf("JSON status=%d stderr=%q; want 0 and empty", status, stderr)
	}
	envelope := decodeEnvelope(t, output)
	var body string
	if err := json.Unmarshal(envelope["data"], &body); err != nil || body != report.Body {
		t.Errorf("JSON body=%q, error=%v; want public PRBody Body %q", body, err, report.Body)
	}
	var diagnostics []project.Diagnostic
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil {
		t.Fatalf("JSON diagnostics=%s, error=%v; want array", envelope["diagnostics"], err)
	}
	if len(diagnostics) != len(report.Diagnostics) {
		t.Errorf("JSON diagnostics=%#v; want %#v", diagnostics, report.Diagnostics)
	}
	for _, want := range report.Diagnostics {
		found := false
		for _, got := range diagnostics {
			found = found || got == want
		}
		if !found {
			t.Errorf("JSON diagnostics omit public PRBody diagnostic %#v: %#v", want, diagnostics)
		}
	}
}

func Test_APP_127_APP_128_APP_129_PRBodyExplicitHistoricalEndpoint(t *testing.T) {
	_, nested, revisionRange, _, report := prBodyAppFixture(t)
	prBodyAppSuccess(t, nested, revisionRange, report)
}

func Test_APP_128_APP_129_PRBodyValidEmptyRange(t *testing.T) {
	root, nested, _, endpoint, _ := prBodyAppFixture(t)
	empty := endpoint + ".." + endpoint
	report, err := workflow.PRBody(root, empty)
	if err != nil {
		t.Fatalf("public PRBody empty range: %v", err)
	}
	if strings.TrimSpace(report.Body) == "" || !(strings.Contains(strings.ToLower(report.Body), "none") || strings.Contains(strings.ToLower(report.Body), "no change")) {
		t.Fatalf("public PRBody empty range lacks no-selection text: %q", report.Body)
	}
	prBodyAppSuccess(t, nested, empty, report)
}

func Test_APP_130_APP_131_PRBodyFailures(t *testing.T) {
	root, nested, revisionRange, _, _ := prBodyAppFixture(t)
	invalidRoot := t.TempDir()
	badConfig := appProjectFixture(t)
	cliGit(t, badConfig, "init", "-q", "-b", "main")
	cliCommit(t, badConfig, "valid base")
	badBase := cliGit(t, badConfig, "rev-parse", "HEAD")
	appPut(t, badConfig, "fireweed.yaml", "[invalid endpoint config\n")
	cliCommit(t, badConfig, "bad endpoint config")
	badEndpoint := cliGit(t, badConfig, "rev-parse", "HEAD")
	badRange := badBase + ".." + badEndpoint
	if _, err := workflow.PRBody(badConfig, badRange); err == nil || !strings.Contains(err.Error(), "fireweed.yaml") {
		t.Fatalf("public PRBody bad endpoint configuration error=%v; want configuration cause", err)
	}
	unavailableRange := strings.SplitN(revisionRange, "..", 2)[0] + "..not-a-commit"
	if _, err := workflow.PRBody(root, unavailableRange); err == nil || !strings.Contains(err.Error(), "P01") {
		t.Fatalf("public PRBody unavailable endpoint error=%v; want P01", err)
	}
	for _, tc := range []struct {
		name, dir, code, cause string
		args                   []string
		status                 int
	}{
		{"missing-range", nested, "U01", "pr-body", []string{"pr-body"}, 2},
		{"extra-range", nested, "U01", "pr-body", []string{"pr-body", revisionRange, "HEAD"}, 2},
		{"invalid-option", nested, "U01", "--bogus", []string{"pr-body", revisionRange, "--bogus"}, 2},
		{"absent-project", invalidRoot, "C01", "", []string{"pr-body", revisionRange}, 2},
		{"bad-endpoint-config", badConfig, "C01", "", []string{"pr-body", badRange}, 2},
		{"unavailable-range", nested, "P01", "not-a-commit", []string{"pr-body", unavailableRange}, 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, jsonMode := range []bool{false, true} {
				args := append([]string(nil), tc.args...)
				if jsonMode {
					args = append(args, "--json")
				}
				status, stdout, stderr := prBodyAppRun(t, tc.dir, args...)
				message := commandFailure(t, status, tc.status, stdout, stderr, jsonMode)
				if jsonMode {
					message = assertFailedJSON(t, stdout, tc.code)
				} else if !strings.HasPrefix(strings.TrimSpace(stdout+stderr), tc.code+":") {
					t.Errorf("plain diagnostic=%q; want %s", stdout+stderr, tc.code)
				}
				if !corrective(message) || (tc.cause != "" && !strings.Contains(message, tc.cause)) {
					t.Errorf("diagnostic %q lacks cause %q or corrective action", message, tc.cause)
				}
				if jsonMode && stderr != "" {
					t.Errorf("JSON stderr=%q; want empty", stderr)
				}
			}
		})
	}
}
