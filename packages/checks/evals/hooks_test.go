package evals_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

const passXML = `<testsuite><testcase name="CHK-003" classname="hook" time="0.25"/></testsuite>`
const failXML = `<testsuite><testcase name="CHK-003" classname="hook"><failure message="expected"/></testcase></testsuite>`

func hookSnapshot(t *testing.T, global, override any) project.Snapshot {
	t.Helper()
	root := t.TempDir()
	module := project.Module{Name: "alpha", Dir: "packages/alpha", Document: contracts.Document{
		Frontmatter: map[string]any{},
		Body:        "# alpha\n\n## Specs\n\n### CHK-003\nEvaluate MUST run the module eval override or global eval hook, accepting shell strings or argv with one-pass module/module_dir/report placeholders (shell values quoted) and FIREWEED_BIN set to request Bin. It MUST parse the written JUnit and retain process output/exit. Launch/timeout, missing report and exit/report disagreement MUST fail; a nonzero exited hook with mapped failures is valid report evidence.\n",
	}}
	if override != nil {
		module.Document.Frontmatter["hooks"] = map[string]any{"eval": override}
	}
	config := map[string]any{}
	if global != nil {
		config["hooks"] = map[string]any{"eval": global}
	}
	return project.Snapshot{Root: root, Config: config, Modules: []project.Module{module}}
}

func Test_CHK_003_CHK_032_CHK_033_CHK_034_CHK_035_GlobalShellHookQuotesAndOnePassExpansion(t *testing.T) {
	cmd := `printf '%s|%s|%s' {module} {module_dir} "$FIREWEED_BIN"; printf '%s' '` + passXML + `' > {report}`
	snap := hookSnapshot(t, cmd, nil)
	snap.Modules[0].Name = "alpha {report}'s"
	snap.Modules[0].Dir = "module {module}'s dir"
	report, err := checks.Evaluate(context.Background(), checks.HookRequest{Snapshot: snap, Module: snap.Modules[0].Name, Bin: "/tmp/new-fireweed", Timeout: 2 * time.Second})
	if err != nil {
		t.Fatal(err)
	}
	want := snap.Modules[0].Name + "|" + snap.Modules[0].Dir + "|/tmp/new-fireweed"
	if report.Stdout != want {
		t.Errorf("stdout = %q; want %q", report.Stdout, want)
	}
	if report.HookExit != 0 || len(report.Cases) != 1 || report.Cases[0].Result != "pass" || report.Cases[0].Seconds != 0.25 {
		t.Errorf("parsed report = %#v", report)
	}
}

func Test_CHK_003_CHK_031_CHK_032_CHK_033_CHK_034_CHK_035_ModuleOverrideArgvAndOutput(t *testing.T) {
	snap := hookSnapshot(t, `echo global-was-run; exit 4`, nil)
	script := filepath.Join(snap.Root, "fixture.sh")
	code := `#!/bin/sh
printf '%s|%s|%s' "$1" "$2" "$FIREWEED_BIN"
printf '%s' 'stderr-marker' >&2
printf '%s' '` + passXML + `' > "$3"
`
	if err := os.WriteFile(script, []byte(code), 0700); err != nil {
		t.Fatal(err)
	}
	snap.Modules[0].Document.Frontmatter["hooks"] = map[string]any{"eval": []any{script, "{module}", "{module_dir}", "{report}"}}
	report, err := checks.Evaluate(context.Background(), checks.HookRequest{Snapshot: snap, Module: "alpha", Bin: "/tmp/subject", Timeout: 2 * time.Second})
	if err != nil {
		t.Fatal(err)
	}
	want := "alpha|" + snap.Modules[0].Dir + "|/tmp/subject"
	if report.Stdout != want || report.Stderr != "stderr-marker" || report.HookExit != 0 || len(report.Cases) != 1 {
		t.Errorf("override report = %#v; want stdout %q", report, want)
	}
}

func Test_CHK_038_ExitAndReportAgreement(t *testing.T) {
	for _, tc := range []struct {
		name, xml, exit string
		valid           bool
		wantExit        int
	}{
		{"zero-with-failure", failXML, "0", false, 0},
		{"nonzero-with-pass", passXML, "7", false, 7},
		{"nonzero-with-failure", failXML, "7", true, 7},
	} {
		snap := hookSnapshot(t, `printf '%s' '`+tc.xml+`' > {report}; exit `+tc.exit, nil)
		report, err := checks.Evaluate(context.Background(), checks.HookRequest{Snapshot: snap, Module: "alpha", Timeout: 2 * time.Second})
		if (err == nil) != tc.valid {
			t.Errorf("%s: report %#v, err %v; valid want %v", tc.name, report, err, tc.valid)
		}
		if !tc.valid {
			requireDiagnosticError(t, err, "C11", "eval hook exit "+tc.exit+" disagrees with JUnit failures")
		}
		if tc.valid && (report.HookExit != tc.wantExit || len(report.Cases) != 1 || report.Cases[0].Result != "fail") {
			t.Errorf("%s: nonzero mapped failure lost: %#v", tc.name, report)
		}
	}
}

func Test_CHK_036_CHK_037_MissingReportAndLaunchFailure(t *testing.T) {
	for _, tc := range []struct {
		command any
		cause   string
	}{
		{`exit 0`, "read JUnit report:"},
		{[]any{filepath.Join(t.TempDir(), "missing-command"), "{report}"}, "run hook:"},
	} {
		snap := hookSnapshot(t, tc.command, nil)
		_, err := checks.Evaluate(context.Background(), checks.HookRequest{Snapshot: snap, Module: "alpha", Timeout: time.Second})
		requireDiagnosticError(t, err, "C11", tc.cause)
	}
}

func Test_CHK_036_TimeoutFails(t *testing.T) {
	snap := hookSnapshot(t, `sleep 2; printf '%s' '`+passXML+`' > {report}`, nil)
	start := time.Now()
	_, err := checks.Evaluate(context.Background(), checks.HookRequest{Snapshot: snap, Module: "alpha", Timeout: 80 * time.Millisecond})
	requireDiagnosticError(t, err, "C11", "run hook:")
	if time.Since(start) > 3*time.Second {
		t.Error("hook timeout did not bound execution")
	}
}

func evaluateSetupCause(t *testing.T, command any, module, body, code, cause string) {
	t.Helper()
	snap := hookSnapshot(t, command, nil)
	if body != "" {
		snap.Modules[0].Document.Body = body
	}
	_, err := checks.Evaluate(context.Background(), checks.HookRequest{Snapshot: snap, Module: module, Timeout: time.Second})
	requireDiagnosticError(t, err, code, cause)
}

func Test_CHK_046_EvaluateHookSetupCauses(t *testing.T) {
	for _, tc := range []struct {
		name, cause string
		command     any
	}{
		{name: "missing hook", cause: "missing eval hook for alpha"},
		{name: "invalid hook command", command: []any{42}, cause: "hook argv contains non-string"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			evaluateSetupCause(t, tc.command, "alpha", "", "C11", tc.cause)
		})
	}
}

func Test_CHK_047_EvaluateUnknownModuleCause(t *testing.T) {
	evaluateSetupCause(t, nil, "missing", "", "U02", `unknown module "missing"`)
}

func Test_CHK_048_EvaluateModuleItemCauses(t *testing.T) {
	for _, tc := range []struct {
		name, body, code, cause string
	}{
		{name: "malformed item", body: "# alpha\n\n## Specs\n\n### malformed\nBad.\n", code: "C03", cause: `malformed item heading: "### malformed"`},
		{name: "duplicate item", body: "# alpha\n\n## Specs\n\n### CHK-003\nOne.\n\n### CHK-003\nTwo.\n", code: "C05", cause: "duplicate item ID: CHK-003"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			evaluateSetupCause(t, `exit 0`, "alpha", tc.body, tc.code, tc.cause)
		})
	}
}

func Test_CHK_028_EvaluatePreservesJUnitDiagnostic(t *testing.T) {
	snap := hookSnapshot(t, `printf '%s' '<testsuite><testcase name="ordinary"/></testsuite>' > {report}`, nil)
	_, err := checks.Evaluate(context.Background(), checks.HookRequest{Snapshot: snap, Module: "alpha", Timeout: time.Second})
	requireDiagnosticError(t, err, "C11", `testcase "ordinary" has no known spec ID`)
}

func Test_CHK_003_CHK_032_ArgvPlaceholdersAreOnePass(t *testing.T) {
	snap := hookSnapshot(t, nil, nil)
	snap.Modules[0].Name = "alpha-{report}"
	snap.Modules[0].Dir = "dir-{module}"
	script := filepath.Join(snap.Root, "fixture.sh")
	code := `#!/bin/sh
printf '%s|%s' "$1" "$2"
printf '%s' '` + passXML + `' > "$3"
`
	if err := os.WriteFile(script, []byte(code), 0700); err != nil {
		t.Fatal(err)
	}
	snap.Modules[0].Document.Frontmatter["hooks"] = map[string]any{"eval": []any{script, "{module}", "{module_dir}", "{report}"}}
	report, err := checks.Evaluate(context.Background(), checks.HookRequest{Snapshot: snap, Module: snap.Modules[0].Name, Timeout: time.Second})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(report.Stdout, "alpha-{report}|"+snap.Modules[0].Dir) {
		t.Errorf("placeholders expanded recursively: %q", report.Stdout)
	}
}
