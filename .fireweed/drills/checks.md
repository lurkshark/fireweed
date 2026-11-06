---
date: 2026-11-06T15:54:22.780454Z
base_commit: 2a48ccfbba646da4cd32c4bbc52dfde2599f2f9b
agent: codex
model: gpt-6-sol
result: fail
---

## Own evals

## Dependent evals

## Contradicted Decisions
None identified.
## Gaps
- CHK-055 (checks): The supplied workspace has no `go.mod` or dependent internal implementations, so the retained Go evals cannot run here. The prior `PackageFailure`/coverage result cannot be reproduced or cleared in this input alone.
- no spec (app, runner, workflow): The prior verification also reported `PackageFailure` for these modules. Their implementations are outside the authorized checks rebuild paths.
## Agent observations

### Attempt 1 (exit 0)
Reconstruction held for outside-controller input/verifier inspection.
Outside-controller inspection released prescribed reconstruction worker.
Rebuilt checks under [packages/checks/internal/engine](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562/packages/checks/internal/engine) and wrote [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562/DRILL_NOTES.md). Retained inputs were untouched; nothing was committed.

The new Go files pass `gofmt`. I could not run the retained evals here: the supplied workspace has no Go manifest or dependency implementations, so `go test` stops before compilation.

Reading additional input from stdin...
OpenAI Codex v0.165.0
--------
workdir: /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
model: gpt-6-sol
provider: openai
approval: never
sandbox: workspace-write [workdir, /tmp, $TMPDIR]
reasoning effort: high
reasoning summaries: none
session id: 01a1b191-7d65-7df0-b0ca-eccb50c1486d
--------
user
# Rebuild checks

Target spec and Decisions:
---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
id_prefix: CHK
interface:
    - api.go
    - model/**
name: checks
---

# checks

## Specs

### CHK-001
ParseJUnit MUST accept testsuite or testsuites documents, including nested suites, returning testcase name, classname and time.

### CHK-002
RequireCoverage MUST cover each current spec except review or quarantine with a non-skipped referencing case. Missing coverage MUST return a C10-prefixed error.

### CHK-003
Evaluate MUST run the configured eval hook, accepting shell strings or argv.

### CHK-004
During audit, Zones MUST allow edits only to the current change file.

### CHK-005
Boundaries MUST run applicable global and per-module boundary hooks.

### CHK-006
Zones MUST reject edits to protected paths in every phase.

### CHK-007
Zones MUST reject edits to module SPEC.md files in every phase.

### CHK-008
Zones MUST reject edits within another module in every phase.

### CHK-009
Zones MUST reject edits to change files other than the current change in every phase.

### CHK-010
During evals and impl, Zones MUST allow edits outside modules only when the path matches worker_writable, except for the current change file.

### CHK-011
During evals, Zones MUST allow edits in the target module's evals zone.

### CHK-012
During evals, Zones MUST allow edits in the target module's interface zone only when InterfaceDeclared is true.

### CHK-013
During evals, Zones MUST allow edits in the target module's disposable zone only for files absent from the phase base.

### CHK-014
During evals, Zones MUST allow edits in the target module's other zone.

### CHK-015
During evals, Zones MUST reject edits in the target module's preserve zone.

### CHK-016
During impl, Zones MUST reject edits in the target module's evals zone.

### CHK-017
During impl, Zones MUST reject edits in the target module's interface zone.

### CHK-018
During impl, Zones MUST allow edits in the target module's disposable zone.

### CHK-019
During impl, Zones MUST allow edits in the target module's preserve zone.

### CHK-020
During impl, Zones MUST allow edits in the target module's other zone.

### CHK-021
During evals and impl, Zones MUST allow edits to the current change file.

### CHK-022
During impl, Zones MUST allow only newly added files in the target module's append_only zone.

### CHK-023
During evals, Zones MUST reject edits in the target module's append_only zone.

### CHK-024
Zones MUST check both the source and destination of a rename against phase permissions.

### CHK-025
RequireFailFirst MUST find a failed or errored referencing case for every requested enforced ID except already_satisfied; passing, skipped, empty or unrelated failures do not qualify. Missing genuine fail-first evidence MUST return a P2-01-prefixed error.

### CHK-026
RequirePass MUST reject failed or errored cases referencing any non-quarantined spec with an E01-prefixed error.

### CHK-027
ParseJUnit MUST reject malformed XML with a C11-prefixed error.

### CHK-028
ParseJUnit MUST reject testcases with missing or unknown spec IDs with a C11-prefixed error.

### CHK-029
ParseJUnit MUST map testcase name and classname to current spec IDs using contracts.MatchesID.

### CHK-030
ParseJUnit MUST classify failure and error elements as fail, and skipped elements as skip.

### CHK-031
Evaluate MUST prefer a module's eval hook override over the global eval hook.

### CHK-032
Evaluate MUST expand module, module_dir and report placeholders once in shell strings or argv, quoting substituted shell values.

### CHK-033
Evaluate MUST set the hook's FIREWEED_BIN environment variable to the request Bin.

### CHK-034
Evaluate MUST retain the hook process's stdout, stderr and exit code in its returned report.

### CHK-035
Evaluate MUST parse the hook's written JUnit report into the returned report's cases.

### CHK-036
Evaluate MUST fail with a C11-prefixed error when the eval hook cannot launch or exceeds its timeout.

### CHK-037
Evaluate MUST fail with a C11-prefixed error when the eval hook does not write its JUnit report.

### CHK-038
Evaluate MUST reject a zero exit with mapped failures or a nonzero exit without mapped failures with a C11-prefixed error, while accepting a nonzero exit with mapped failures as valid report evidence.

### CHK-039
Boundaries MUST return errors with diagnostic X01 for unsuppressed module-attributed violations.

### CHK-040
Boundaries MUST return a visible X03 warning containing the reason for a reasoned boundaries_opt_out.

### CHK-041
Boundaries MUST suppress only exact baseline lines from normally exited hooks.

### CHK-042
Boundaries MUST suppress only the opted-out module's violation lines through a reasoned opt-out after a hook exits normally.

### CHK-043
Boundaries MUST report hook launch, cancellation or timeout failures as errors with a diagnostic whose Code is X02. The returned error MUST preserve the hook failure cause without a leading X02 prefix.

### CHK-044
Boundaries MUST fail with diagnostic X02 for a missing hook required by a module with disposable globs and no reasoned opt-out.

### CHK-045
Boundaries MUST warn with diagnostic X04 without an error for a missing hook when no boundary enforcement is required.

### CHK-046
Evaluate MUST return C11-prefixed errors for existing eval-hook setup failures, including missing or invalid hook commands.

### CHK-047
Evaluate MUST return a U02-prefixed error for an unknown requested module.

### CHK-048
Evaluate MUST return C03-prefixed errors for malformed module item syntax and C05-prefixed errors for duplicate module item identities.

### CHK-049
Doctor MUST retain project.Check diagnostics and prevent hook or agent probes when the current project has a static error.

### CHK-050
Doctor MUST run probes using current project files in a detached scratch worktree below the common Git directory's fireweed/worktrees, leaving the caller's checkout unchanged.

### CHK-051
Doctor MUST remove its scratch worktree and Git registration after successful or failed probes.

### CHK-052
Doctor MUST run the configured setup hook in scratch before probes; a setup failure MUST produce C14 naming the setup hook and prevent later probes.

### CHK-053
Doctor MUST render two uniquely named throwaway modules from project-owned module templates and the version-1 canary baseline, substituting name, prefix, package, go_module, module_dir and other_dir.

### CHK-054
Doctor MUST report C14 naming missing or malformed module templates or required canary maps and skip the affected probes.

### CHK-055
Doctor MUST require passing baseline evals and coverage for the throwaway canary modules, and applicable boundaries, before planting canaries and after removing each planted canary. This baseline coverage requirement applies to the canary modules used to prove hooks, rather than every existing module. Recovery failures MUST produce C14 naming the failed hook and canary.

### CHK-056
Doctor MUST plant the supplied undeclared canary and require its boundary hook to report a module-attributed violation; a missed canary MUST produce C14 naming the hook and canary.

### CHK-057
Doctor MUST plant the supplied internal_eval canary and require its boundary hook to report a module-attributed interface violation; a missed canary MUST produce C14 naming the hook and canary.

### CHK-058
Doctor MUST require the supplied failing_eval canary to produce a failed referencing testcase and nonzero eval exit; compile errors, missing reports and unrelated failures MUST NOT count as detection.

### CHK-059
Doctor MUST require the supplied missing_id canary to be rejected for its unmapped testcase ID; unrelated hook or report failures MUST NOT count as detection.

### CHK-060
Doctor MUST test applicable global eval and boundary hooks and every per-module override through their corresponding canary probes.

### CHK-061
Doctor MUST set probe hooks' FIREWEED_BIN to the request Bin and make the scratch canary modules discoverable through that executable's modules output.

### CHK-062
Doctor MUST retain each real module's reasoned X03 boundaries-opt-out warning and applicable X04 warning when no boundary enforcement is required.

### CHK-063
Doctor MUST run the worker and audit commands once each in scratch through the configured agent.wrap with a fixed-reply prompt, requiring exit 0 and the requested reply for success.

### CHK-064
Doctor MUST use agent.command for the audit probe when audit.command is unset and emit a C15 warning identifying the shared-command blind spot.

### CHK-065
Doctor MUST bound each agent probe by the request Timeout or three minutes when unset, and report C15 failures with the command and its final 20 output lines.

### CHK-066
Doctor MUST classify a configured limit_patterns match in an agent probe's final 20 output lines as a C15 usage-limit diagnostic rather than an ordinary command failure.

### CHK-067
Doctor with NoAgent MUST skip worker and audit command execution while retaining hook canary and recovery checks.

### CHK-068
Doctor MUST emit a C16 warning when the AGENTS.md managed block is missing or differs from the current fireweed guidance, while accepting a current block regardless of surrounding text.

### CHK-069
Doctor MUST emit a C16 warning when a project playbook override's embedded_base is missing or differs from the SHA-256 of its current embedded playbook; a matching base MUST avoid that warning.

### CHK-070
History MUST list full commit IDs selected by the supplied Git revision range in oldest-first topological order.
For an empty range it MUST use the merge base of HEAD and the locally resolved configured base through HEAD.

### CHK-071
History MUST identify worker commits from new phase assignments in change-file contents across pending/archive moves, rather than commit subjects.
Deleting a change document without assigning a phase MUST not by itself classify the commit as a worker commit.

### CHK-072
History MUST report H01 when a worker commit changes a protected path or a project playbook override.

### CHK-073
History MUST report H02 when an impl-to-applied phase commit changes evals or interface zones.

### CHK-074
History MUST report H03 when any commit modifies, deletes or renames an existing append_only path, including merge-resolution changes identified by combined Git diffs.

### CHK-075
History MUST report H04 when a commit modifies an existing file under changes/archive.
Deleting an archived change file MUST remain allowed by H04.

### CHK-076
History MUST recognize each new applied phase whose parents lack an impl phase as squashed, listing its commit and change without a warning.
H01 and H02 MUST not apply to those squashed commits.
H03 and H04 MUST still apply.

### CHK-077
History MUST return each rule finding as a Diagnostic with its H01–H04 code, repository-relative Path and full commit ID in Message, using the recorded before/after zone maps and the worker change's projected set values.

### CHK-078
History MUST return an error when a revision range or required recorded project/change document cannot be read or parsed.

### CHK-079
ParseJUnit MUST retain the message attributes and text bodies of JUnit failure and error elements in each affected Case.Failure, including message-only and body-only diagnostics.
A testcase without such text MUST have an empty Failure.

### CHK-080
Evaluate MUST retain parsed Case.Failure strings in its returned Report.Cases, including when failing testcases make the eval hook exit nonzero.

### CHK-081
Doctor MUST persist complete stdout and stderr from every executed worker and audit probe, including failed, timed-out and cancelled probes, under the invoking project's .fireweed/runs/<run>/<phase>-1.log. Each Doctor invocation MUST use a fresh shared run identity, with phases doctor-worker and doctor-audit. Each executed probe's Detail MUST identify its log file by absolute path.

### CHK-082
Doctor MUST expand {prompt}, {prompt_file} and {git_common_dir} once in selected worker/audit command templates and wrapper argv, quoting shell substitutions and preserving argv values literally.

### CHK-083
Doctor MUST create one temporary file containing the exact probe prompt only when a selected command or wrapper template uses {prompt_file}, and expand that placeholder to its absolute path. The file MUST remain available through execution and be removed afterward, including after launch failure, nonzero exit, timeout or cancellation.

## Decisions

[0071-doctor-probes] Doctor requests the literal reply FIREWEED_DOCTOR_OK; limit_patterns match
case-insensitive literal substrings. Project playbook overrides use frontmatter
embedded_base containing the hexadecimal SHA-256 of the embedded playbook text.

[0114-history-validation] A merge merely retaining a phase value already held by a parent is not a new
phase assignment; combined merge diffs still enforce append-only files.
Squashed entries identify the change's repository-relative after-path; ordinary
non-merge phase transitions compare the sole parent.


Retained interfaces and files:
packages/checks/SPEC.md
packages/checks/api.go
packages/checks/model/model.go

Retained evals:
packages/checks/evals/.gitkeep
packages/checks/evals/boundaries_test.go
packages/checks/evals/doctor_agent_logs_test.go
packages/checks/evals/doctor_prompt_file_test.go
packages/checks/evals/doctor_test.go
packages/checks/evals/history_test.go
packages/checks/evals/hooks_test.go
packages/checks/evals/junit_test.go
packages/checks/evals/zones_test.go

Rebuild only target disposable paths. Do not edit retained inputs. Do not read outside the supplied input workspace, external trees, old code, source history, binaries, saved attempts, or dependent internal implementation. Do not repair dependent internals. Write DRILL_NOTES.md with separate ## Contradicted Decisions and ## Gaps sections. Every gap must name the nearest spec ID or "no spec" and a target module. Do not commit.

codex
I’ll inspect the checks interface, retained evals, and allowed project context, then rebuild the disposable checks files and run the relevant tests. I’ll keep the retained inputs untouched.
exec
/bin/zsh -lc 'find packages/checks -maxdepth 3 -type f -print' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc "pwd && rg --files -g 'AGENTS.md' -g 'packages/checks/**' -g 'go.mod' -g 'DRILL_NOTES.md'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc 'git status --short' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 620ms:
packages/checks/evals/history_test.go
packages/checks/evals/doctor_agent_logs_test.go
packages/checks/evals/.gitkeep
packages/checks/evals/doctor_prompt_file_test.go
packages/checks/evals/junit_test.go
packages/checks/evals/hooks_test.go
packages/checks/evals/zones_test.go
packages/checks/evals/doctor_test.go
packages/checks/evals/boundaries_test.go
packages/checks/model/model.go
packages/checks/api.go
packages/checks/SPEC.md

 succeeded in 638ms:
 succeeded in 663ms:
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
packages/checks/SPEC.md
packages/checks/api.go
packages/checks/model/model.go
packages/checks/evals/boundaries_test.go
packages/checks/evals/doctor_test.go
packages/checks/evals/zones_test.go
packages/checks/evals/hooks_test.go
packages/checks/evals/junit_test.go
packages/checks/evals/doctor_prompt_file_test.go
packages/checks/evals/.gitkeep
packages/checks/evals/doctor_agent_logs_test.go
packages/checks/evals/history_test.go

exec
/bin/zsh -lc 'cat packages/checks/api.go packages/checks/model/model.go packages/checks/SPEC.md' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc 'ls -la && find . -maxdepth 2 -type f -print' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc 'cat packages/checks/evals/junit_test.go packages/checks/evals/hooks_test.go packages/checks/evals/zones_test.go packages/checks/evals/boundaries_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc 'cat packages/checks/evals/doctor_test.go packages/checks/evals/doctor_agent_logs_test.go packages/checks/evals/doctor_prompt_file_test.go packages/checks/evals/history_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 663ms:
package evals_test

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
)

func requireDiagnosticError(t *testing.T, err error, code, cause string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected %s error containing %q", code, cause)
	}
	if !strings.HasPrefix(err.Error(), code+": ") || !strings.Contains(err.Error(), cause) || strings.HasPrefix(strings.TrimPrefix(err.Error(), code+": "), code+":") {
		t.Errorf("error = %q; want %s prefix and cause %q", err, code, cause)
	}
}

func checkSpecs() []contracts.Item {
	return []contracts.Item{{ID: "CHK-001"}, {ID: "CHK-002"}, {ID: "CHK-003"}}
}

func Test_CHK_001_CHK_029_CHK_030_NestedJUnitIDsResultsAndTime(t *testing.T) {
	data := []byte(`<testsuites><testsuite name="outer"><testsuite name="inner">
<testcase classname="suite.ChK_001" name="passes" time="1.25"/>
<testcase classname="suite" name="test_chk002_failed" time="0.5"><failure message="bad"/></testcase>
<testcase classname="suite" name="test_chk_003_error" time="0.75"><error message="broken"/></testcase>
<testcase classname="CHK-002" name="CHK-001_both" time="0.125"/>
</testsuite></testsuite></testsuites>`)
	report, err := checks.ParseJUnit(data, checkSpecs())
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]checks.Case{
		"passes":             {Name: "passes", Classname: "suite.ChK_001", IDs: []string{"CHK-001"}, Result: "pass", Seconds: 1.25},
		"test_chk002_failed": {Name: "test_chk002_failed", Classname: "suite", IDs: []string{"CHK-002"}, Result: "fail", Seconds: 0.5},
		"test_chk_003_error": {Name: "test_chk_003_error", Classname: "suite", IDs: []string{"CHK-003"}, Result: "fail", Seconds: 0.75},
		"CHK-001_both":       {Name: "CHK-001_both", Classname: "CHK-002", IDs: []string{"CHK-001", "CHK-002"}, Result: "pass", Seconds: 0.125},
	}
	if len(report.Cases) != len(want) {
		t.Fatalf("cases = %#v; want %d", report.Cases, len(want))
	}
	for _, got := range report.Cases {
		sort.Strings(got.IDs)
		got.Failure = ""
		if !reflect.DeepEqual(got, want[got.Name]) {
			t.Errorf("case = %#v; want %#v", got, want[got.Name])
		}
	}
}

func Test_CHK_001_CHK_029_CHK_030_SingleSuiteSkipped(t *testing.T) {
	report, err := checks.ParseJUnit([]byte(`<testsuite><testcase name="CHK-001" classname="unit" time="0.125"><skipped/></testcase></testsuite>`), checkSpecs())
	if err != nil {
		t.Fatal(err)
	}
	if len(report.Cases) != 1 || report.Cases[0].Result != "skip" || report.Cases[0].Seconds != 0.125 || !reflect.DeepEqual(report.Cases[0].IDs, []string{"CHK-001"}) {
		t.Fatalf("skipped case = %#v", report.Cases)
	}
}

func Test_CHK_079_ParseJUnitRetainsFailureAndErrorDetails(t *testing.T) {
	data := []byte(`<testsuite>
<testcase name="CHK-001_failure_combined"><failure message="want 2 &amp; got 3">assertion at example_test.go:17</failure></testcase>
<testcase name="CHK-001_failure_message"><failure message="expected true"/></testcase>
<testcase name="CHK-002_failure_body"><failure>actual value was false</failure></testcase>
<testcase name="CHK-002_error_combined"><error message="panic: nil pointer">stack at example_test.go:29</error></testcase>
<testcase name="CHK-003_error_message"><error message="setup failed"/></testcase>
<testcase name="CHK-003_error_body"><error>missing fixture</error></testcase>
<testcase name="CHK-001_empty_failure"><failure/></testcase>
<testcase name="CHK-002_pass"/>
<testcase name="CHK-003_skip"><skipped/></testcase>
</testsuite>`)
	report, err := checks.ParseJUnit(data, checkSpecs())
	if err != nil {
		t.Fatal(err)
	}
	want := map[string][]string{
		"CHK-001_failure_combined": {"want 2 & got 3", "assertion at example_test.go:17"},
		"CHK-001_failure_message":  {"expected true"},
		"CHK-002_failure_body":     {"actual value was false"},
		"CHK-002_error_combined":   {"panic: nil pointer", "stack at example_test.go:29"},
		"CHK-003_error_message":    {"setup failed"},
		"CHK-003_error_body":       {"missing fixture"},
		"CHK-001_empty_failure":    nil,
		"CHK-002_pass":             nil,
		"CHK-003_skip":             nil,
	}
	if len(report.Cases) != len(want) {
		t.Fatalf("cases = %#v; want %d", report.Cases, len(want))
	}
	for _, got := range report.Cases {
		parts, ok := want[got.Name]
		if !ok {
			t.Errorf("unexpected case %#v", got)
			continue
		}
		if len(parts) == 0 && got.Failure != "" {
			t.Errorf("%s: failure = %q; want empty", got.Name, got.Failure)
		}
		for _, part := range parts {
			if !strings.Contains(got.Failure, part) {
				t.Errorf("%s: failure = %q; missing %q", got.Name, got.Failure, part)
			}
		}
	}
}

func Test_CHK_002_SkippedJUnitCaseDoesNotCover(t *testing.T) {
	report, err := checks.ParseJUnit([]byte(`<testsuite><testcase name="CHK-001" classname="unit" time="0.125"><skipped/></testcase></testsuite>`), checkSpecs())
	if err != nil {
		t.Fatal(err)
	}
	requireDiagnosticError(t, checks.RequireCoverage(report, []contracts.Item{{ID: "CHK-001"}}), "C10", "no coverage for CHK-001")
}

func Test_CHK_027_RejectMalformedXML(t *testing.T) {
	for _, tc := range []struct{ xml, cause string }{
		{`<testsuite><testcase name="CHK-001"></testsuite>`, "parse JUnit:"},
		{`<document/>`, `invalid JUnit root "document"`},
	} {
		_, err := checks.ParseJUnit([]byte(tc.xml), checkSpecs())
		requireDiagnosticError(t, err, "C11", tc.cause)
	}
}

func Test_CHK_028_RejectMissingAndUnknownIDs(t *testing.T) {
	for _, xml := range []string{
		`<testsuite><testcase name="ordinary" classname="unit"/></testsuite>`,
		`<testsuite><testcase name="ordinary" classname="unit"><skipped/></testcase></testsuite>`,
		`<testsuite><testcase name="CHK-999" classname="unit"/></testsuite>`,
		`<testsuite><testcase name="OCHK001" classname="unit"/></testsuite>`,
		`<testsuite><testcase name="CHK0010" classname="unit"/></testsuite>`,
	} {
		_, err := checks.ParseJUnit([]byte(xml), checkSpecs())
		cause := "has no known spec ID"
		if strings.Contains(xml, "CHK-999") {
			cause = `references unknown spec ID "CHK-999"`
		} else if strings.Contains(xml, "CHK0010") {
			cause = `references unknown spec ID "CHK0010"`
		}
		requireDiagnosticError(t, err, "C11", cause)
	}
}

func Test_CHK_002_CoverageRequiresCurrentNonSkippedCases(t *testing.T) {
	specs := []contracts.Item{{ID: "CHK-001"}, {ID: "CHK-002"}}
	requireDiagnosticError(t, checks.RequireCoverage(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001"}, Result: "pass"}}}, specs), "C10", "no coverage for CHK-002")
	if err := checks.RequireCoverage(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001", "CHK-002"}, Result: "fail"}}}, specs); err != nil {
		t.Errorf("failed referencing case should cover both IDs: %v", err)
	}
}

func Test_CHK_002_CoverageExemptsReviewAndQuarantine(t *testing.T) {
	specs := []contracts.Item{{ID: "CHK-001"}, {ID: "CHK-002", Tag: "review"}, {ID: "CHK-003", Tag: "quarantine: flaky"}}
	if err := checks.RequireCoverage(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001"}, Result: "pass"}}}, specs); err != nil {
		t.Errorf("review/quarantine required coverage: %v", err)
	}
}

func Test_CHK_025_FailFirstRequiresEachRequestedFailure(t *testing.T) {
	ids := []string{"CHK-001", "CHK-002"}
	bad := []checks.Case{
		{IDs: []string{"CHK-001"}, Result: "pass"},
		{IDs: []string{"CHK-002"}, Result: "skip"},
		{IDs: []string{"CHK-003"}, Result: "fail"},
	}
	requireDiagnosticError(t, checks.RequireFailFirst(checks.Report{Cases: bad}, ids, nil), "P2-01", "no fail-first case for CHK-001")
	requireDiagnosticError(t, checks.RequireFailFirst(checks.Report{}, ids, nil), "P2-01", "no fail-first case for CHK-001")
	requireDiagnosticError(t, checks.RequireFailFirst(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001"}, Result: "fail"}, {IDs: []string{"CHK-002"}, Result: "skip"}}}, ids, nil), "P2-01", "no fail-first case for CHK-002")
	good := checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001"}, Result: "fail"}, {IDs: []string{"CHK-002"}, Result: "fail"}}}
	if err := checks.RequireFailFirst(good, ids, nil); err != nil {
		t.Errorf("both requested IDs failed: %v", err)
	}
}

func Test_CHK_025_FailFirstAlreadySatisfiedExemption(t *testing.T) {
	report := checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-002"}, Result: "fail"}}}
	if err := checks.RequireFailFirst(report, []string{"CHK-001", "CHK-002"}, []string{"CHK-001"}); err != nil {
		t.Errorf("already_satisfied ID incorrectly required: %v", err)
	}
}

func Test_CHK_026_PassRejectsNonQuarantinedFailures(t *testing.T) {
	specs := []contracts.Item{{ID: "CHK-001"}, {ID: "CHK-002", Tag: "review"}, {ID: "CHK-003", Tag: "quarantine: flaky"}}
	for _, id := range []string{"CHK-001", "CHK-002"} {
		requireDiagnosticError(t, checks.RequirePass(checks.Report{Cases: []checks.Case{{IDs: []string{id}, Result: "fail"}}}, specs), "E01", `failed case "" for `+id)
	}
	if err := checks.RequirePass(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-003"}, Result: "fail"}}}, specs); err != nil {
		t.Errorf("quarantined failure rejected: %v", err)
	}
}
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

func Test_CHK_080_EvaluateRetainsParsedFailureDetails(t *testing.T) {
	for _, tc := range []struct {
		name, xml, exit, wantFailure string
		wantExit                     int
	}{
		{"nonzero failure", `<testsuite><testcase name="CHK-003_assertion"><failure message="want 2">got 3 at example_test.go:17</failure></testcase></testsuite>`, "7", "got 3 at example_test.go:17", 7},
		{"zero passing", `<testsuite><testcase name="CHK-003_pass"/></testsuite>`, "0", "", 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			snap := hookSnapshot(t, `printf '%s' '`+tc.xml+`' > {report}; exit `+tc.exit, nil)
			report, err := checks.Evaluate(context.Background(), checks.HookRequest{Snapshot: snap, Module: "alpha", Timeout: 2 * time.Second})
			if err != nil {
				t.Fatal(err)
			}
			if len(report.Cases) != 1 {
				t.Fatalf("cases = %#v; want one", report.Cases)
			}
			if report.HookExit != tc.wantExit {
				t.Errorf("hook exit = %d; want %s", report.HookExit, tc.exit)
			}
			if !strings.Contains(report.Cases[0].Failure, tc.wantFailure) {
				t.Errorf("failure = %q; missing %q", report.Cases[0].Failure, tc.wantFailure)
			}
			if tc.wantFailure == "" && report.Cases[0].Failure != "" {
				t.Errorf("passing case failure = %q; want empty", report.Cases[0].Failure)
			}
			if tc.wantFailure != "" && !strings.Contains(report.Cases[0].Failure, "want 2") {
				t.Errorf("failure = %q; missing message", report.Cases[0].Failure)
			}
		})
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
package evals_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

func zoneFixture(t *testing.T) (project.Snapshot, string) {
	t.Helper()
	root := t.TempDir()
	for path := range map[string]string{
		"packages/alpha/SPEC.md": "alpha", "packages/beta/SPEC.md": "beta",
		"packages/alpha/evals/old_test.go": "old", "packages/alpha/api.go": "api",
		"packages/alpha/internal/existing.go": "old", "packages/alpha/gen/old.sql": "old",
		"packages/alpha/keep.txt": "old", "changes/0011-current.md": "current",
	} {
		full := filepath.Join(root, path)
		if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte("fixture"), 0644); err != nil {
			t.Fatal(err)
		}
	}
	run := func(args ...string) string {
		cmd := exec.Command("git", args...)
		cmd.Dir = root
		cmd.Env = append(os.Environ(), "GIT_AUTHOR_NAME=Fixture", "GIT_AUTHOR_EMAIL=fixture@example.test", "GIT_COMMITTER_NAME=Fixture", "GIT_COMMITTER_EMAIL=fixture@example.test")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("git %v: %v: %s", args, err, out)
		}
		return string(out)
	}
	run("init", "-q")
	run("add", ".")
	run("commit", "-qm", "base")
	base := run("rev-parse", "HEAD")
	zones := contracts.Zones{Evals: []string{"evals/**"}, Interface: []string{"api.go"}, Disposable: []string{"internal/**"}, Preserve: []string{"keep.txt"}, AppendOnly: []string{"gen/**"}}
	snap := project.Snapshot{Root: root, Config: map[string]any{"protected": []any{"locked/**"}, "worker_writable": []any{"go.sum"}}, Modules: []project.Module{
		{Name: "alpha", Dir: "packages/alpha", Zones: zones},
		{Name: "beta", Dir: "packages/beta", Zones: zones},
	}}
	return snap, base[:len(base)-1]
}

func zoneRequest(snap project.Snapshot, base, phase string, changes ...gitstore.Change) checks.ZoneRequest {
	return checks.ZoneRequest{Snapshot: snap, Module: "alpha", Phase: phase, ChangePath: "changes/0011-current.md", Base: base, InterfaceDeclared: true, Changes: changes}
}

func zoneAllowed(t *testing.T, request checks.ZoneRequest) {
	t.Helper()
	if diags := checks.Zones(request); len(diags) != 0 {
		t.Errorf("wanted allowed change, got %#v", diags)
	}
}

func zoneRejected(t *testing.T, request checks.ZoneRequest) {
	t.Helper()
	if diags := checks.Zones(request); len(diags) == 0 {
		t.Errorf("forbidden change accepted: %#v", request.Changes)
	}
}

func Test_CHK_004_AuditOnlyCurrentChange(t *testing.T) {
	snap, base := zoneFixture(t)
	zoneAllowed(t, zoneRequest(snap, base, "audit", gitstore.Change{Path: "changes/0011-current.md", Status: "M"}))
	for _, path := range []string{"changes/0012-other.md", "packages/alpha/evals/old_test.go", "go.sum"} {
		zoneRejected(t, zoneRequest(snap, base, "audit", gitstore.Change{Path: path, Status: "M"}))
	}
}

func Test_CHK_006_CHK_007_CHK_008_CHK_009_RejectProtectedSpecForeignModuleAndChangeEveryPhase(t *testing.T) {
	snap, base := zoneFixture(t)
	for _, phase := range []string{"evals", "impl", "audit"} {
		for _, path := range []string{"locked/config", ".fireweed/hooks.go", "packages/alpha/SPEC.md", "packages/beta/internal/code.go", "changes/0012-other.md"} {
			zoneRejected(t, zoneRequest(snap, base, phase, gitstore.Change{Path: path, Status: "M"}))
		}
	}
}

func Test_CHK_011_CHK_012_CHK_013_CHK_014_CHK_015_CHK_021_EvalsAllowedZonesAndNewStubs(t *testing.T) {
	snap, base := zoneFixture(t)
	for _, change := range []gitstore.Change{
		{Path: "packages/alpha/evals/new_test.go", Status: "A"},
		{Path: "packages/alpha/api.go", Status: "M"},
		{Path: "packages/alpha/internal/new_stub.go", Status: "A"},
		{Path: "packages/alpha/manifest.json", Status: "M"},
		{Path: "changes/0011-current.md", Status: "M"},
	} {
		zoneAllowed(t, zoneRequest(snap, base, "evals", change))
	}
	zoneRejected(t, zoneRequest(snap, base, "evals", gitstore.Change{Path: "packages/alpha/internal/existing.go", Status: "M"}))
	zoneRejected(t, zoneRequest(snap, base, "evals", gitstore.Change{Path: "packages/alpha/keep.txt", Status: "M"}))
	request := zoneRequest(snap, base, "evals", gitstore.Change{Path: "packages/alpha/api.go", Status: "M"})
	request.InterfaceDeclared = false
	zoneRejected(t, request)
}

func Test_CHK_016_CHK_017_CHK_018_CHK_019_CHK_020_CHK_021_ImplFreezesEvalsAndInterfaces(t *testing.T) {
	snap, base := zoneFixture(t)
	for _, path := range []string{"packages/alpha/evals/old_test.go", "packages/alpha/api.go"} {
		zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: path, Status: "M"}))
	}
	for _, path := range []string{"packages/alpha/internal/existing.go", "packages/alpha/keep.txt", "packages/alpha/manifest.json", "changes/0011-current.md"} {
		zoneAllowed(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: path, Status: "M"}))
	}
}

func Test_CHK_022_CHK_023_ImplAppendOnlyNewVersusExisting(t *testing.T) {
	snap, base := zoneFixture(t)
	zoneAllowed(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: "packages/alpha/gen/new.sql", Status: "A"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: "packages/alpha/gen/old.sql", Status: "M"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: "packages/alpha/gen/old.sql", Status: "D"}))
	zoneRejected(t, zoneRequest(snap, base, "evals", gitstore.Change{Path: "packages/alpha/gen/new.sql", Status: "A"}))
}

func Test_CHK_024_RenamesCheckBothPaths(t *testing.T) {
	snap, base := zoneFixture(t)
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{OldPath: "packages/beta/internal/old.go", Path: "packages/alpha/internal/new.go", Status: "R100"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{OldPath: "packages/alpha/internal/existing.go", Path: "packages/beta/internal/new.go", Status: "R100"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{OldPath: "packages/alpha/gen/old.sql", Path: "packages/alpha/internal/new.go", Status: "R100"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{OldPath: "packages/alpha/internal/existing.go", Path: "packages/alpha/gen/new.sql", Status: "R100"}))
}

func Test_CHK_010_WorkerWritableOnlyEvalsAndImpl(t *testing.T) {
	snap, base := zoneFixture(t)
	for _, phase := range []string{"evals", "impl"} {
		zoneAllowed(t, zoneRequest(snap, base, phase, gitstore.Change{Path: "go.sum", Status: "M"}))
		zoneRejected(t, zoneRequest(snap, base, phase, gitstore.Change{Path: "outside.txt", Status: "M"}))
	}
	zoneRejected(t, zoneRequest(snap, base, "audit", gitstore.Change{Path: "go.sum", Status: "M"}))
}
package evals_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func boundarySnapshot(t *testing.T, hook any) project.Snapshot {
	t.Helper()
	root := t.TempDir()
	snap := project.Snapshot{Root: root, Config: map[string]any{}, Modules: []project.Module{{
		Name: "alpha", Dir: "packages/alpha",
		Zones:    contracts.Zones{Disposable: []string{"internal/**"}},
		Document: contracts.Document{Frontmatter: map[string]any{}},
	}}}
	if hook != nil {
		snap.Config["hooks"] = map[string]any{"boundaries": hook}
	}
	return snap
}

func boundaryScript(t *testing.T, root, body string) []any {
	t.Helper()
	path := filepath.Join(root, "boundary-hook.sh")
	if err := os.WriteFile(path, []byte("#!/bin/sh\n"+body+"\n"), 0700); err != nil {
		t.Fatal(err)
	}
	return []any{path}
}

func boundaryHasError(diags []project.Diagnostic, err error) bool {
	if err != nil {
		return true
	}
	for _, d := range diags {
		if !d.Warning {
			return true
		}
	}
	return false
}

func boundaryHasWarning(diags []project.Diagnostic) bool {
	for _, d := range diags {
		if d.Warning {
			return true
		}
	}
	return false
}

func requireBoundaryDiagnostic(t *testing.T, diags []project.Diagnostic, code string, warning bool, message string) {
	t.Helper()
	for _, d := range diags {
		if message == "" {
			if d.Code == code && d.Warning == warning {
				return
			}
			continue
		}
		if strings.Contains(d.Message, message) {
			if d.Code != code || d.Warning != warning {
				t.Errorf("boundary diagnostic for %q = %#v; want code %s, warning %t", message, d, code, warning)
			}
			return
		}
	}
	t.Errorf("boundary diagnostic for %q absent: %#v", message, diags)
}

func Test_CHK_005_GlobalAndModuleBoundaryHooks(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	globalMarker := filepath.Join(snap.Root, "global-ran")
	moduleMarker := filepath.Join(snap.Root, "module-ran")
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, fmt.Sprintf("printf run > %q; printf 'alpha: global violation\\n'", globalMarker))}
	moduleHook := filepath.Join(snap.Root, "module-hook.sh")
	if err := os.WriteFile(moduleHook, []byte(fmt.Sprintf("#!/bin/sh\nprintf run > %q; printf 'beta: module violation\\n'\n", moduleMarker)), 0700); err != nil {
		t.Fatal(err)
	}
	snap.Modules = append(snap.Modules, project.Module{Name: "beta", Dir: "packages/beta", Zones: contracts.Zones{Disposable: []string{"internal/**"}}, Document: contracts.Document{Frontmatter: map[string]any{"hooks": map[string]any{"boundaries": []any{moduleHook}}}}})
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: 2 * time.Second})
	if !boundaryHasError(diags, err) {
		t.Errorf("unsuppressed violations accepted: %#v, %v", diags, err)
	}
	for _, module := range []string{"alpha", "beta"} {
		if !strings.Contains(fmt.Sprint(diags, err), module) {
			t.Errorf("violation for %s not attributed: %#v, %v", module, diags, err)
		}
	}
	for _, path := range []string{globalMarker, moduleMarker} {
		if _, statErr := os.Stat(path); statErr != nil {
			t.Errorf("applicable hook did not run: %s: %v", path, statErr)
		}
	}
}

func Test_CHK_041_ExactBaselineOnly(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	if err := os.MkdirAll(filepath.Join(snap.Root, ".fireweed"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(snap.Root, ".fireweed", "boundaries-baseline"), []byte("alpha: old violation\n"), 0644); err != nil {
		t.Fatal(err)
	}
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: old violation\n'`)}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) {
		t.Errorf("exact baseline line rejected: %#v, %v", diags, err)
	}
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: old violation changed\n'`)}
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("non-exact baseline line suppressed")
	}
}

func Test_CHK_039_CHK_040_OptOutSuppressesOnlyItsModuleAndWarns(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	snap.Modules[0].Document.Frontmatter["boundaries_opt_out"] = "legacy dependency"
	snap.Modules = append(snap.Modules, project.Module{Name: "beta", Dir: "packages/beta", Zones: contracts.Zones{Disposable: []string{"internal/**"}}, Document: contracts.Document{Frontmatter: map[string]any{}}})
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: ignored\n'`)}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) || !boundaryHasWarning(diags) {
		t.Errorf("reasoned opt-out = %#v, %v", diags, err)
	}
	if !strings.Contains(fmt.Sprint(diags), "legacy dependency") {
		t.Errorf("opt-out reason absent from warning: %#v", diags)
	}
	requireBoundaryDiagnostic(t, diags, "X03", true, "legacy dependency")
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: ignored\nbeta: enforced\n'`)}
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("opt-out hid another module's violation")
	}
	if err == nil {
		t.Error("unsuppressed beta violation returned no error")
	}
	requireBoundaryDiagnostic(t, diags, "X01", false, "beta: enforced")
}

func Test_CHK_044_CHK_045_MissingHookRequiredVersusOptional(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("missing hook for disposable module accepted")
	}
	if err == nil {
		t.Error("missing required hook returned no error")
	}
	requireBoundaryDiagnostic(t, diags, "X02", false, "alpha: missing boundaries hook")
	snap.Modules[0].Zones.Disposable = nil
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) || !boundaryHasWarning(diags) {
		t.Errorf("missing optional hook = %#v, %v", diags, err)
	}
	requireBoundaryDiagnostic(t, diags, "X04", true, "no boundary enforcement required")
	snap.Modules[0].Zones.Disposable = []string{"internal/**"}
	snap.Modules[0].Document.Frontmatter["boundaries_opt_out"] = "legacy"
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) || !boundaryHasWarning(diags) {
		t.Errorf("missing hook with reasoned opt-out = %#v, %v", diags, err)
	}
	requireBoundaryDiagnostic(t, diags, "X03", true, "legacy")
	requireBoundaryDiagnostic(t, diags, "X04", true, "no boundary enforcement required")
}

func Test_CHK_041_CHK_042_NormalNonzeroExitCanUseBaselineAndOptOut(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	if err := os.MkdirAll(filepath.Join(snap.Root, ".fireweed"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(snap.Root, ".fireweed", "boundaries-baseline"), []byte("alpha: known\n"), 0644); err != nil {
		t.Fatal(err)
	}
	snap.Modules[0].Document.Frontmatter["boundaries_opt_out"] = "legacy"
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: known\nalpha: opted out\n'; exit 3`)}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) {
		t.Errorf("normally exited hook with fully suppressed lines rejected: %#v, %v", diags, err)
	}
}

func Test_CHK_043_LaunchCancellationAndTimeoutStayErrors(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	snap.Modules[0].Document.Frontmatter["boundaries_opt_out"] = "legacy"
	snap.Config["hooks"] = map[string]any{"boundaries": []any{filepath.Join(snap.Root, "missing-executable")}}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("launch failure hidden by opt-out")
	}
	if err == nil {
		t.Error("launch failure returned no error")
	} else if !strings.Contains(err.Error(), "missing-executable") || strings.HasPrefix(err.Error(), "X02:") {
		t.Errorf("launch error = %q; want missing executable cause without X02 prefix", err)
	}
	requireBoundaryDiagnostic(t, diags, "X02", false, "")
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `sleep 2`)}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	diags, err = checks.Boundaries(ctx, checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("cancelled hook hidden by opt-out")
	}
	if err == nil {
		t.Error("cancelled hook returned no error")
	} else if !errors.Is(err, context.Canceled) || strings.HasPrefix(err.Error(), "X02:") {
		t.Errorf("cancellation error = %q; want canceled cause without X02 prefix", err)
	}
	requireBoundaryDiagnostic(t, diags, "X02", false, "")
	start := time.Now()
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: 80 * time.Millisecond})
	if !boundaryHasError(diags, err) {
		t.Error("timed-out hook hidden by opt-out")
	}
	if err == nil {
		t.Error("timed-out hook returned no error")
	} else if !errors.Is(err, context.DeadlineExceeded) || strings.HasPrefix(err.Error(), "X02:") {
		t.Errorf("timeout error = %q; want deadline cause without X02 prefix", err)
	}
	requireBoundaryDiagnostic(t, diags, "X02", false, "")
	if time.Since(start) > 3*time.Second {
		t.Error("boundary timeout did not bound execution")
	}
}

 succeeded in 663ms:
package checks

import (
	"context"

	"github.com/lurkshark/fireweed/packages/checks/internal/engine"
	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

type Case = model.Case
type Report = model.Report
type HookRequest = model.HookRequest
type BoundaryRequest = model.BoundaryRequest
type ZoneRequest = model.ZoneRequest
type DoctorRequest = model.DoctorRequest
type DoctorCheck = model.DoctorCheck
type DoctorReport = model.DoctorReport
type HistoryReport = model.HistoryReport
type SquashedChange = model.SquashedChange

func ParseJUnit(data []byte, specs []contracts.Item) (Report, error) {
	return engine.ParseJUnit(data, specs)
}
func RequireCoverage(report Report, specs []contracts.Item) error {
	return engine.RequireCoverage(report, specs)
}
func RequireFailFirst(report Report, ids, already []string) error {
	return engine.RequireFailFirst(report, ids, already)
}
func RequirePass(report Report, specs []contracts.Item) error {
	return engine.RequirePass(report, specs)
}
func Evaluate(ctx context.Context, request HookRequest) (Report, error) {
	return engine.Evaluate(ctx, request)
}
func Boundaries(ctx context.Context, request BoundaryRequest) ([]project.Diagnostic, error) {
	return engine.Boundaries(ctx, request)
}
func Zones(request ZoneRequest) []project.Diagnostic { return engine.Zones(request) }
func Doctor(ctx context.Context, request DoctorRequest) DoctorReport {
	return engine.Doctor(ctx, request)
}
func History(root, revisionRange string) (HistoryReport, error) {
	return engine.History(root, revisionRange)
}
package model

import (
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

type Case struct {
	Name      string
	Classname string
	IDs       []string
	Result    string
	Failure   string
	Seconds   float64
}

type Report struct {
	Cases    []Case
	Seconds  float64
	HookExit int
	Stdout   string
	Stderr   string
}

type HookRequest struct {
	Snapshot project.Snapshot
	Module   string
	Bin      string
	Timeout  time.Duration
}

type BoundaryRequest struct {
	Snapshot project.Snapshot
	Bin      string
	Timeout  time.Duration
}

type ZoneRequest struct {
	Snapshot          project.Snapshot
	Module            string
	Phase             string
	ChangePath        string
	Base              string
	InterfaceDeclared bool
	Changes           []gitstore.Change
}

type DoctorRequest struct {
	Root    string
	Bin     string
	NoAgent bool
	Timeout time.Duration
}

type DoctorCheck struct {
	Name   string
	Module string
	Passed bool
	Detail string
}

type DoctorReport struct {
	Checks      []DoctorCheck
	Diagnostics []project.Diagnostic
}

type HistoryReport struct {
	Commits     []string
	Squashed    []SquashedChange
	Diagnostics []project.Diagnostic
}

type SquashedChange struct {
	Commit string
	Change string
}
---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
id_prefix: CHK
interface:
    - api.go
    - model/**
name: checks
---

# checks

## Specs

### CHK-001
ParseJUnit MUST accept testsuite or testsuites documents, including nested suites, returning testcase name, classname and time.

### CHK-002
RequireCoverage MUST cover each current spec except review or quarantine with a non-skipped referencing case. Missing coverage MUST return a C10-prefixed error.

### CHK-003
Evaluate MUST run the configured eval hook, accepting shell strings or argv.

### CHK-004
During audit, Zones MUST allow edits only to the current change file.

### CHK-005
Boundaries MUST run applicable global and per-module boundary hooks.

### CHK-006
Zones MUST reject edits to protected paths in every phase.

### CHK-007
Zones MUST reject edits to module SPEC.md files in every phase.

### CHK-008
Zones MUST reject edits within another module in every phase.

### CHK-009
Zones MUST reject edits to change files other than the current change in every phase.

### CHK-010
During evals and impl, Zones MUST allow edits outside modules only when the path matches worker_writable, except for the current change file.

### CHK-011
During evals, Zones MUST allow edits in the target module's evals zone.

### CHK-012
During evals, Zones MUST allow edits in the target module's interface zone only when InterfaceDeclared is true.

### CHK-013
During evals, Zones MUST allow edits in the target module's disposable zone only for files absent from the phase base.

### CHK-014
During evals, Zones MUST allow edits in the target module's other zone.

### CHK-015
During evals, Zones MUST reject edits in the target module's preserve zone.

### CHK-016
During impl, Zones MUST reject edits in the target module's evals zone.

### CHK-017
During impl, Zones MUST reject edits in the target module's interface zone.

### CHK-018
During impl, Zones MUST allow edits in the target module's disposable zone.

### CHK-019
During impl, Zones MUST allow edits in the target module's preserve zone.

### CHK-020
During impl, Zones MUST allow edits in the target module's other zone.

### CHK-021
During evals and impl, Zones MUST allow edits to the current change file.

### CHK-022
During impl, Zones MUST allow only newly added files in the target module's append_only zone.

### CHK-023
During evals, Zones MUST reject edits in the target module's append_only zone.

### CHK-024
Zones MUST check both the source and destination of a rename against phase permissions.

### CHK-025
RequireFailFirst MUST find a failed or errored referencing case for every requested enforced ID except already_satisfied; passing, skipped, empty or unrelated failures do not qualify. Missing genuine fail-first evidence MUST return a P2-01-prefixed error.

### CHK-026
RequirePass MUST reject failed or errored cases referencing any non-quarantined spec with an E01-prefixed error.

### CHK-027
ParseJUnit MUST reject malformed XML with a C11-prefixed error.

### CHK-028
ParseJUnit MUST reject testcases with missing or unknown spec IDs with a C11-prefixed error.

### CHK-029
ParseJUnit MUST map testcase name and classname to current spec IDs using contracts.MatchesID.

### CHK-030
ParseJUnit MUST classify failure and error elements as fail, and skipped elements as skip.

### CHK-031
Evaluate MUST prefer a module's eval hook override over the global eval hook.

### CHK-032
Evaluate MUST expand module, module_dir and report placeholders once in shell strings or argv, quoting substituted shell values.

### CHK-033
Evaluate MUST set the hook's FIREWEED_BIN environment variable to the request Bin.

### CHK-034
Evaluate MUST retain the hook process's stdout, stderr and exit code in its returned report.

### CHK-035
Evaluate MUST parse the hook's written JUnit report into the returned report's cases.

### CHK-036
Evaluate MUST fail with a C11-prefixed error when the eval hook cannot launch or exceeds its timeout.

### CHK-037
Evaluate MUST fail with a C11-prefixed error when the eval hook does not write its JUnit report.

### CHK-038
Evaluate MUST reject a zero exit with mapped failures or a nonzero exit without mapped failures with a C11-prefixed error, while accepting a nonzero exit with mapped failures as valid report evidence.

### CHK-039
Boundaries MUST return errors with diagnostic X01 for unsuppressed module-attributed violations.

### CHK-040
Boundaries MUST return a visible X03 warning containing the reason for a reasoned boundaries_opt_out.

### CHK-041
Boundaries MUST suppress only exact baseline lines from normally exited hooks.

### CHK-042
Boundaries MUST suppress only the opted-out module's violation lines through a reasoned opt-out after a hook exits normally.

### CHK-043
Boundaries MUST report hook launch, cancellation or timeout failures as errors with a diagnostic whose Code is X02. The returned error MUST preserve the hook failure cause without a leading X02 prefix.

### CHK-044
Boundaries MUST fail with diagnostic X02 for a missing hook required by a module with disposable globs and no reasoned opt-out.

### CHK-045
Boundaries MUST warn with diagnostic X04 without an error for a missing hook when no boundary enforcement is required.

### CHK-046
Evaluate MUST return C11-prefixed errors for existing eval-hook setup failures, including missing or invalid hook commands.

### CHK-047
Evaluate MUST return a U02-prefixed error for an unknown requested module.

### CHK-048
Evaluate MUST return C03-prefixed errors for malformed module item syntax and C05-prefixed errors for duplicate module item identities.

### CHK-049
Doctor MUST retain project.Check diagnostics and prevent hook or agent probes when the current project has a static error.

### CHK-050
Doctor MUST run probes using current project files in a detached scratch worktree below the common Git directory's fireweed/worktrees, leaving the caller's checkout unchanged.

### CHK-051
Doctor MUST remove its scratch worktree and Git registration after successful or failed probes.

### CHK-052
Doctor MUST run the configured setup hook in scratch before probes; a setup failure MUST produce C14 naming the setup hook and prevent later probes.

### CHK-053
Doctor MUST render two uniquely named throwaway modules from project-owned module templates and the version-1 canary baseline, substituting name, prefix, package, go_module, module_dir and other_dir.

### CHK-054
Doctor MUST report C14 naming missing or malformed module templates or required canary maps and skip the affected probes.

### CHK-055
Doctor MUST require passing baseline evals and coverage for the throwaway canary modules, and applicable boundaries, before planting canaries and after removing each planted canary. This baseline coverage requirement applies to the canary modules used to prove hooks, rather than every existing module. Recovery failures MUST produce C14 naming the failed hook and canary.

### CHK-056
Doctor MUST plant the supplied undeclared canary and require its boundary hook to report a module-attributed violation; a missed canary MUST produce C14 naming the hook and canary.

### CHK-057
Doctor MUST plant the supplied internal_eval canary and require its boundary hook to report a module-attributed interface violation; a missed canary MUST produce C14 naming the hook and canary.

### CHK-058
Doctor MUST require the supplied failing_eval canary to produce a failed referencing testcase and nonzero eval exit; compile errors, missing reports and unrelated failures MUST NOT count as detection.

### CHK-059
Doctor MUST require the supplied missing_id canary to be rejected for its unmapped testcase ID; unrelated hook or report failures MUST NOT count as detection.

### CHK-060
Doctor MUST test applicable global eval and boundary hooks and every per-module override through their corresponding canary probes.

### CHK-061
Doctor MUST set probe hooks' FIREWEED_BIN to the request Bin and make the scratch canary modules discoverable through that executable's modules output.

### CHK-062
Doctor MUST retain each real module's reasoned X03 boundaries-opt-out warning and applicable X04 warning when no boundary enforcement is required.

### CHK-063
Doctor MUST run the worker and audit commands once each in scratch through the configured agent.wrap with a fixed-reply prompt, requiring exit 0 and the requested reply for success.

### CHK-064
Doctor MUST use agent.command for the audit probe when audit.command is unset and emit a C15 warning identifying the shared-command blind spot.

### CHK-065
Doctor MUST bound each agent probe by the request Timeout or three minutes when unset, and report C15 failures with the command and its final 20 output lines.

### CHK-066
Doctor MUST classify a configured limit_patterns match in an agent probe's final 20 output lines as a C15 usage-limit diagnostic rather than an ordinary command failure.

### CHK-067
Doctor with NoAgent MUST skip worker and audit command execution while retaining hook canary and recovery checks.

### CHK-068
Doctor MUST emit a C16 warning when the AGENTS.md managed block is missing or differs from the current fireweed guidance, while accepting a current block regardless of surrounding text.

### CHK-069
Doctor MUST emit a C16 warning when a project playbook override's embedded_base is missing or differs from the SHA-256 of its current embedded playbook; a matching base MUST avoid that warning.

### CHK-070
History MUST list full commit IDs selected by the supplied Git revision range in oldest-first topological order.
For an empty range it MUST use the merge base of HEAD and the locally resolved configured base through HEAD.

### CHK-071
History MUST identify worker commits from new phase assignments in change-file contents across pending/archive moves, rather than commit subjects.
Deleting a change document without assigning a phase MUST not by itself classify the commit as a worker commit.

### CHK-072
History MUST report H01 when a worker commit changes a protected path or a project playbook override.

### CHK-073
History MUST report H02 when an impl-to-applied phase commit changes evals or interface zones.

### CHK-074
History MUST report H03 when any commit modifies, deletes or renames an existing append_only path, including merge-resolution changes identified by combined Git diffs.

### CHK-075
History MUST report H04 when a commit modifies an existing file under changes/archive.
Deleting an archived change file MUST remain allowed by H04.

### CHK-076
History MUST recognize each new applied phase whose parents lack an impl phase as squashed, listing its commit and change without a warning.
H01 and H02 MUST not apply to those squashed commits.
H03 and H04 MUST still apply.

### CHK-077
History MUST return each rule finding as a Diagnostic with its H01–H04 code, repository-relative Path and full commit ID in Message, using the recorded before/after zone maps and the worker change's projected set values.

### CHK-078
History MUST return an error when a revision range or required recorded project/change document cannot be read or parsed.

### CHK-079
ParseJUnit MUST retain the message attributes and text bodies of JUnit failure and error elements in each affected Case.Failure, including message-only and body-only diagnostics.
A testcase without such text MUST have an empty Failure.

### CHK-080
Evaluate MUST retain parsed Case.Failure strings in its returned Report.Cases, including when failing testcases make the eval hook exit nonzero.

### CHK-081
Doctor MUST persist complete stdout and stderr from every executed worker and audit probe, including failed, timed-out and cancelled probes, under the invoking project's .fireweed/runs/<run>/<phase>-1.log. Each Doctor invocation MUST use a fresh shared run identity, with phases doctor-worker and doctor-audit. Each executed probe's Detail MUST identify its log file by absolute path.

### CHK-082
Doctor MUST expand {prompt}, {prompt_file} and {git_common_dir} once in selected worker/audit command templates and wrapper argv, quoting shell substitutions and preserving argv values literally.

### CHK-083
Doctor MUST create one temporary file containing the exact probe prompt only when a selected command or wrapper template uses {prompt_file}, and expand that placeholder to its absolute path. The file MUST remain available through execution and be removed afterward, including after launch failure, nonzero exit, timeout or cancellation.

## Decisions

[0071-doctor-probes] Doctor requests the literal reply FIREWEED_DOCTOR_OK; limit_patterns match
case-insensitive literal substrings. Project playbook overrides use frontmatter
embedded_base containing the hexadecimal SHA-256 of the embedded playbook text.

[0114-history-validation] A merge merely retaining a phase value already held by a parent is not a new
phase assignment; combined merge diffs still enforce append-only files.
Squashed entries identify the change's repository-relative after-path; ordinary
non-merge phase transitions compare the sole parent.

 succeeded in 662ms:
package evals_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"gopkg.in/yaml.v3"
)

const doctorGuidance = "<!-- fireweed:start -->\nThis project uses fireweed. Specs and change files are the source of truth.\nRun `fireweed playbook author-change` before authoring changes and\n`fireweed context <change>` before each step.\n<!-- fireweed:end -->\n"

type doctorFixture struct {
	t    *testing.T
	root string
	log  string
	bin  string
}

func (f doctorFixture) put(name, body string) {
	f.t.Helper()
	path := filepath.Join(f.root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		f.t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		f.t.Fatal(err)
	}
}

func (f doctorFixture) git(args ...string) string {
	f.t.Helper()
	cmd := exec.Command("git", args...)
	cmd.Dir = f.root
	cmd.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_TERMINAL_PROMPT=0")
	out, err := cmd.CombinedOutput()
	if err != nil {
		f.t.Fatalf("git %v: %v: %s", args, err, out)
	}
	return string(out)
}

func (f doctorFixture) commit() {
	f.t.Helper()
	f.git("init", "-b", "main")
	f.git("config", "user.name", "Doctor Fixture")
	f.git("config", "user.email", "doctor@example.test")
	f.git("add", ".")
	f.git("commit", "-m", "fixture")
}

func (f doctorFixture) request(noAgent bool) checks.DoctorRequest {
	return checks.DoctorRequest{Root: f.root, Bin: f.bin, NoAgent: noAgent, Timeout: 15 * time.Second}
}

func doctorConfig(t *testing.T, recipe assets.Recipe) string {
	t.Helper()
	config := map[string]any{
		"fireweed": ">=0.1, <0.2", "version": 1, "base": "main", "modules": []string{"packages/*"},
		"limit_patterns":  []string{"rate limit"},
		"module_defaults": map[string]any{"interface": recipe.ModuleDefaults.Interface, "evals": recipe.ModuleDefaults.Evals, "disposable": recipe.ModuleDefaults.Disposable},
		"hooks":           recipe.Hooks,
		"agent":           map[string]any{"command": []string{"bash", ".fireweed/agent.sh", "{prompt}", "{git_common_dir}"}, "wrap": []string{"bash", ".fireweed/wrap.sh"}},
		"audit":           map[string]any{"enabled": true, "command": []string{"bash", ".fireweed/audit.sh", "{prompt}"}},
	}
	data, err := yaml.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func (f doctorFixture) editConfig(change func(map[string]any)) {
	f.t.Helper()
	var config map[string]any
	if err := yaml.Unmarshal(mustRead(f.t, filepath.Join(f.root, "fireweed.yaml")), &config); err != nil {
		f.t.Fatal(err)
	}
	change(config)
	data, err := yaml.Marshal(config)
	if err != nil {
		f.t.Fatal(err)
	}
	f.put("fireweed.yaml", string(data))
}

func newDoctorFixture(t *testing.T) doctorFixture {
	t.Helper()
	r, err := assets.BuiltinRecipe("go")
	if err != nil {
		t.Fatal(err)
	}
	f := doctorFixture{t: t, root: t.TempDir(), log: filepath.Join(t.TempDir(), "probes.log")}
	f.put("go.mod", "module example.test/doctorfixture\n\ngo 1.23\n")
	f.put("fireweed.yaml", doctorConfig(t, r))
	f.put("AGENTS.md", "Before.\n"+doctorGuidance+"After.\n")
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nThe fixture MUST pass.\n")
	f.put("packages/alpha/api.go", "package alpha\n")
	for path, body := range r.Files {
		f.put(path, body)
	}
	f.put(".fireweed/wrap.sh", fmt.Sprintf("#!/bin/bash\nprintf 'wrap:%%s\n' \"$PWD\" >> %q\nexec \"$@\"\n", f.log))
	f.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/bash\nprintf 'worker:%%s:%%s:%%s\n' \"$PWD\" \"$1\" \"$2\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	f.put(".fireweed/audit.sh", fmt.Sprintf("#!/bin/bash\nprintf 'audit:%%s:%%s\n' \"$PWD\" \"$1\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	return f
}

func localDoctorFixture(t *testing.T) doctorFixture {
	t.Helper()
	f := doctorFixture{t: t, root: t.TempDir(), log: filepath.Join(t.TempDir(), "probes.log")}
	f.bin = filepath.Join(t.TempDir(), "fake-fireweed")
	if err := os.WriteFile(f.bin, []byte(`#!/bin/sh
if test "$1" != modules || test "$2" != --json; then exit 2; fi
printf '{"data":['
first=1
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test "$first" = 0; then printf ','; fi; first=0
  printf '{"name":"%s","dir":"%s","depends_on":[],"zones":{"interface":["api.txt"],"evals":["evals/**"],"disposable":["internal/**"]}}' "$name" "$dir"
done
printf ']}'
`), 0700); err != nil {
		t.Fatal(err)
	}
	recipe := assets.Recipe{ModuleDefaults: structZones(), Hooks: map[string]any{
		"eval":       "sh .fireweed/eval.sh {module} {module_dir} {report}",
		"boundaries": "sh .fireweed/boundary.sh",
		"setup":      "sh .fireweed/setup.sh",
	}}
	f.put("fireweed.yaml", doctorConfig(t, recipe))
	f.put(".gitignore", ".fireweed/runs/\n")
	f.put("AGENTS.md", doctorGuidance)
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nThe fixture MUST pass.\n")
	f.put("packages/alpha/api.txt", "public\n")
	f.put(".fireweed/templates/module/SPEC.md", "---\nname: {{name}}\nid_prefix: {{prefix}}\ndepends_on: []\n---\n# {{name}}\n\n## Specs\n\n### {{prefix}}-001\nThe canary MUST pass.\n")
	f.put(".fireweed/templates/module/api.txt.tmpl", "{{package}} {{go_module}} {{module_dir}} {{other_dir}}\n")
	f.put(".fireweed/canaries/canaries.yaml", `version: 1
baseline:
  SPEC.md: "---\nname: {{name}}\nid_prefix: {{prefix}}\ndepends_on: []\n---\n# {{name}}\n\n## Specs\n\n### {{prefix}}-001\nThe canary MUST pass.\n"
  api.txt: "{{package}} {{go_module}} {{module_dir}} {{other_dir}}"
undeclared:
  undeclared.txt: "{{other_dir}}"
internal_eval:
  internal/private.txt: "private"
  evals/internal.txt: "{{module_dir}}"
failing_eval:
  evals/failing.txt: "{{prefix}}-001"
missing_id:
  evals/missing.txt: "UNMAPPED-999"
`)
	f.put(".fireweed/setup.sh", fmt.Sprintf("#!/bin/sh\nprintf 'setup:%%s\n' \"$PWD\" >> %q\n", f.log))
	f.put(".fireweed/eval.sh", fmt.Sprintf(`#!/bin/sh
printf 'eval:%%s:%%s:%%s\n' "$PWD" "$1" "$FIREWEED_BIN" >> %q
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
if test -f "$2/evals/failing.txt"; then printf '<testsuite><testcase name="%%s-001"><failure/></testcase></testsuite>' "$prefix" > "$3"; exit 1; fi
if test -f "$2/evals/missing.txt"; then printf '<testsuite><testcase name="UNMAPPED-999"/></testsuite>' > "$3"; exit 0; fi
printf '<testsuite><testcase name="%%s-001"/></testsuite>' "$prefix" > "$3"
`, f.log))
	f.put(".fireweed/boundary.sh", fmt.Sprintf(`#!/bin/sh
printf 'boundaries:%%s:%%s\n' "$PWD" "$FIREWEED_BIN" >> %q
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test -f "$dir/undeclared.txt"; then printf '%%s: undeclared import\n' "$name"; fi
  if test -f "$dir/evals/internal.txt"; then printf '%%s: interface violation\n' "$name"; fi
done
`, f.log))
	f.put(".fireweed/wrap.sh", fmt.Sprintf("#!/bin/bash\nprintf 'wrap:%%s\n' \"$PWD\" >> %q\nexec \"$@\"\n", f.log))
	f.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/bash\nprintf 'worker:%%s:%%s:%%s\n' \"$PWD\" \"$1\" \"$2\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	f.put(".fireweed/audit.sh", fmt.Sprintf("#!/bin/bash\nprintf 'audit:%%s:%%s\n' \"$PWD\" \"$1\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	return f
}

func structZones() (z contracts.Zones) {
	z.Interface = []string{"api.txt"}
	z.Evals = []string{"evals/**"}
	z.Disposable = []string{"internal/**"}
	return
}

func doctorText(report checks.DoctorReport) string {
	return fmt.Sprintf("checks=%+v diagnostics=%+v", report.Checks, report.Diagnostics)
}
func hasDoctorDiagnostic(report checks.DoctorReport, code string, warning bool, terms ...string) bool {
	for _, d := range report.Diagnostics {
		if d.Code != code || d.Warning != warning {
			continue
		}
		ok := true
		for _, s := range terms {
			if !strings.Contains(strings.ToLower(d.Message), strings.ToLower(s)) {
				ok = false
			}
		}
		if ok {
			return true
		}
	}
	return false
}
func hasDoctorCheck(report checks.DoctorReport, name string, passed bool) bool {
	for _, c := range report.Checks {
		if strings.Contains(c.Name, name) && c.Passed == passed {
			return true
		}
	}
	return false
}
func doctorErrors(report checks.DoctorReport) bool {
	for _, d := range report.Diagnostics {
		if !d.Warning {
			return true
		}
	}
	return false
}
func doctorLog(f doctorFixture) string { b, _ := os.ReadFile(f.log); return string(b) }
func worktrees(f doctorFixture) string { return f.git("worktree", "list", "--porcelain") }

func logDoctorCanaryProbes(t *testing.T, f doctorFixture) {
	t.Helper()
	logProbe := fmt.Sprintf(`printf 'probe:%%s\n' "$PWD" >> %q
for marker in undeclared.txt internal.txt failing.txt missing.txt; do
  if find packages -name "$marker" | grep -q .; then printf 'planted:%%s\n' "$marker" >> %q; fi
done
`, f.log, f.log)
	for _, path := range []string{".fireweed/eval.sh", ".fireweed/boundary.sh"} {
		body := string(mustRead(t, filepath.Join(f.root, path)))
		if !strings.HasPrefix(body, "#!/bin/sh\n") {
			t.Fatalf("fixture hook %s has no shell header", path)
		}
		f.put(path, strings.Replace(body, "#!/bin/sh\n", "#!/bin/sh\n"+logProbe, 1))
	}
}

func assertScratchRemoved(t *testing.T, f doctorFixture) {
	t.Helper()
	common := strings.TrimSpace(f.git("rev-parse", "--git-common-dir"))
	if !filepath.IsAbs(common) {
		common = filepath.Join(f.root, common)
	}
	common, err := filepath.EvalSymlinks(common)
	if err != nil {
		t.Fatal(err)
	}
	base := filepath.Join(common, "fireweed", "worktrees")
	for _, line := range strings.Split(doctorLog(f), "\n") {
		if !strings.HasPrefix(line, "setup:") {
			continue
		}
		path := strings.TrimPrefix(line, "setup:")
		for parent := path; ; parent = filepath.Dir(parent) {
			resolved, err := filepath.EvalSymlinks(parent)
			if err == nil {
				suffix, err := filepath.Rel(parent, path)
				if err != nil {
					t.Fatal(err)
				}
				path = filepath.Join(resolved, suffix)
				break
			}
			if parent == filepath.Dir(parent) {
				t.Fatal(err)
			}
		}
		rel, err := filepath.Rel(base, path)
		if err != nil || rel == "." || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
			t.Errorf("scratch %q is not below common Git directory %q", path, base)
		}
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Errorf("physical scratch directory %q remains or cannot be checked: %v", path, err)
		}
		return
	}
	t.Error("setup did not record a scratch path")
}

func Test_CHK_049_StaticDiagnosticsStopProbes(t *testing.T) {
	f := localDoctorFixture(t)
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n## Specs\n### ALP-001\nA.\n### ALP-001\nB.\n")
	f.commit()
	want := project.Check(f.root)
	if len(want) == 0 || !doctorErrors(checks.DoctorReport{Diagnostics: want}) {
		t.Fatal("fixture has no static diagnostic")
	}
	report := checks.Doctor(context.Background(), f.request(false))
	remaining := map[project.Diagnostic]int{}
	for _, d := range report.Diagnostics {
		remaining[d]++
	}
	for _, d := range want {
		if remaining[d] == 0 {
			t.Errorf("complete project diagnostic %+v lost: %s", d, doctorText(report))
		}
		remaining[d]--
	}
	if strings.Contains(doctorLog(f), "setup:") || strings.Contains(doctorLog(f), "eval:") || strings.Contains(doctorLog(f), "boundaries:") || strings.Contains(doctorLog(f), "worker:") || strings.Contains(doctorLog(f), "audit:") {
		t.Errorf("static error ran probes: %s", doctorLog(f))
	}
}

func Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_CHK_060_CHK_061_CHK_067_RealGoRecipe(t *testing.T) {
	f := newDoctorFixture(t)
	f.commit()
	f.bin = doctorBinary(t)
	f.editConfig(func(config map[string]any) {
		config["modules"] = []string{"packages/alpha"}
		config["hooks"].(map[string]any)["setup"] = fmt.Sprintf("printf 'setup:%%s\\n' \"$PWD\" >> %q; test -z \"$(git symbolic-ref -q HEAD)\" && test -f .fireweed/doctor-untracked && go mod download", f.log)
	})
	f.put(".fireweed/doctor-untracked", "ready\n")
	for _, d := range project.Check(f.root) {
		if !d.Warning {
			t.Fatalf("invalid Go doctor fixture: %+v", d)
		}
	}
	beforeStatus, beforeWorktrees := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) {
		t.Errorf("real Go recipe doctor failed: %s", doctorText(report))
	}
	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id", "recovery"} {
		if !hasDoctorCheck(report, name, true) {
			t.Errorf("missing passing %s outcome: %s", name, doctorText(report))
		}
	}
	recoveries := 0
	for _, check := range report.Checks {
		if strings.Contains(strings.ToLower(check.Name), "recovery") && check.Passed {
			recoveries++
		}
	}
	if recoveries < 4 {
		t.Errorf("clean recovery not reported after each canary: %s", doctorText(report))
	}
	modules := map[string]bool{}
	for _, check := range report.Checks {
		if check.Module != "" && check.Module != "alpha" {
			modules[check.Module] = true
		}
	}
	if len(modules) < 2 {
		t.Errorf("two distinct throwaway modules absent: %s", doctorText(report))
	}
	if !strings.Contains(doctorLog(f), "setup:") || !strings.Contains(doctorLog(f), "/fireweed/worktrees/") {
		t.Errorf("setup did not run in scratch: %s", doctorLog(f))
	}
	if strings.Contains(doctorLog(f), "worker:") || strings.Contains(doctorLog(f), "audit:") {
		t.Errorf("NoAgent ran agents: %s", doctorLog(f))
	}
	if got := f.git("status", "--porcelain", "--untracked-files=all"); got != beforeStatus {
		t.Errorf("caller changed: before %q after %q", beforeStatus, got)
	}
	if got := worktrees(f); got != beforeWorktrees {
		t.Errorf("scratch registration leaked: before %q after %q", beforeWorktrees, got)
	}
	assertScratchRemoved(t, f)
}

func Test_CHK_053_NestedModuleTemplatesRenderForBothCanaries(t *testing.T) {
	f := localDoctorFixture(t)
	f.put("go.mod", "module example.test/doctorfixture\n\ngo 1.23\n")
	fields := "{{name}}|{{prefix}}|{{package}}|{{go_module}}|{{module_dir}}|{{other_dir}}\n"
	f.put(".fireweed/templates/module/evals/nested.txt.tmpl", "eval|"+fields)
	f.put(".fireweed/templates/module/internal/nested.txt", "internal|"+fields)
	canaries := string(mustRead(t, filepath.Join(f.root, ".fireweed/canaries/canaries.yaml")))
	f.put(".fireweed/canaries/canaries.yaml", strings.Replace(canaries, "undeclared:\n", "  evals/baseline.txt: \"baseline|"+strings.TrimSpace(fields)+"\"\nundeclared:\n", 1))
	eval := string(mustRead(t, filepath.Join(f.root, ".fireweed/eval.sh")))
	check := fmt.Sprintf(`case "$1" in
  doctor*)
    prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
    other=
    for spec in packages/*/SPEC.md; do
      dir=${spec%%/SPEC.md}
      case "$dir" in packages/alpha|"$2") ;; *) other=$dir ;; esac
    done
    expected="$1|$prefix|$1|example.test/doctorfixture|$2|$other"
    test -n "$other" &&
      test "$(cat "$2/evals/nested.txt")" = "eval|$expected" &&
      test "$(cat "$2/internal/nested.txt")" = "internal|$expected" &&
      test "$(cat "$2/evals/baseline.txt")" = "baseline|$expected" &&
      test ! -e "$2/evals/nested.txt.tmpl" || exit 17
    printf 'rendered:%%s\n' "$1" >> %q
    ;;
esac
`, f.log)
	f.put(".fireweed/eval.sh", strings.Replace(eval, "#!/bin/sh\n", "#!/bin/sh\n"+check, 1))
	f.commit()
	before := worktrees(f)
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) {
		t.Errorf("nested template probes failed: %s", doctorText(report))
	}
	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id", "recovery"} {
		if !hasDoctorCheck(report, name, true) {
			t.Errorf("missing passing %s probe: %s", name, doctorText(report))
		}
	}
	modules := map[string]bool{}
	for _, line := range strings.Split(doctorLog(f), "\n") {
		if strings.HasPrefix(line, "rendered:") {
			modules[strings.TrimPrefix(line, "rendered:")] = true
		}
	}
	if len(modules) != 2 {
		t.Errorf("expected both throwaway modules to pass nested template checks, got %v: %s", modules, doctorText(report))
	}
	if got := worktrees(f); got != before {
		t.Errorf("scratch registration leaked: before %q after %q", before, got)
	}
	if strings.Contains(doctorLog(f), "setup:") {
		assertScratchRemoved(t, f)
	}
}

func Test_CHK_051_CHK_052_SetupFailureStopsAndCleans(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/setup.sh", fmt.Sprintf("#!/bin/sh\nprintf 'setup:%%s\\n' \"$PWD\" >> %q\necho failed setup >&2\nexit 17\n", f.log))
	f.commit()
	before := worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if !hasDoctorDiagnostic(report, "C14", false, "setup") {
		t.Errorf("setup failure absent: %s", doctorText(report))
	}
	if strings.Contains(doctorLog(f), "eval:") || strings.Contains(doctorLog(f), "boundaries:") || strings.Contains(doctorLog(f), "worker:") {
		t.Errorf("setup failure ran later probes: %s", doctorLog(f))
	}
	if got := worktrees(f); got != before {
		t.Errorf("scratch worktree leaked: %q", got)
	}
	assertScratchRemoved(t, f)
}

func Test_CHK_051_CancellationCleansScratchAndStopsSetupGroup(t *testing.T) {
	f := localDoctorFixture(t)
	ready := filepath.Join(t.TempDir(), "ready")
	childFinished := filepath.Join(t.TempDir(), "child-finished")
	f.put(".fireweed/setup.sh", fmt.Sprintf(`#!/bin/sh
printf 'setup:%%s\n' "$PWD" >> %q
sh -c 'sleep 2; printf finished > "$1"' sh %q &
printf ready > %q
wait
`, f.log, childFinished, ready))
	f.commit()
	before := worktrees(f)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reports := make(chan checks.DoctorReport, 1)
	go func() { reports <- checks.Doctor(ctx, f.request(true)) }()
	deadline := time.After(8 * time.Second)
	for {
		if _, err := os.Stat(ready); err == nil {
			break
		} else if !os.IsNotExist(err) {
			t.Fatal(err)
		}
		select {
		case report := <-reports:
			t.Fatalf("setup exited before cancellation: %s", doctorText(report))
		case <-deadline:
			t.Fatal("setup never signalled readiness")
		case <-time.After(10 * time.Millisecond):
		}
	}
	cancel()
	select {
	case report := <-reports:
		if !hasDoctorDiagnostic(report, "C14", false, "setup") {
			t.Errorf("cancelled setup lacked C14: %s", doctorText(report))
		}
	case <-time.After(5 * time.Second):
		t.Fatal("cancelled doctor did not return promptly")
	}
	if got := worktrees(f); got != before {
		t.Errorf("cancelled doctor retained Git registration: before %q after %q", before, got)
	}
	assertScratchRemoved(t, f)
	time.Sleep(2200 * time.Millisecond)
	if _, err := os.Stat(childFinished); !os.IsNotExist(err) {
		t.Errorf("cancelled setup descendant continued: %v", err)
	}
}

func Test_CHK_054_MissingTemplatesAndCanaryMaps(t *testing.T) {
	for _, tc := range []struct{ name, path, body, term, input, affected string }{
		{"missing template", ".fireweed/templates/module/SPEC.md", "", "template", "SPEC.md", "undeclared"},
		{"malformed template", ".fireweed/templates/module/SPEC.md", "---\nname: [broken\nid_prefix: {{prefix}}\n---\n", "template", "SPEC.md", "undeclared"},
		{"missing baseline", ".fireweed/canaries/canaries.yaml", "version: 1\nbaseline: {}\n", "canary", "baseline", "undeclared"},
		{"missing internal_eval", ".fireweed/canaries/canaries.yaml", "missing_id:\n  evals/missing.txt: missing\n", "internal_eval", "internal_eval", "internal_eval"},
		{"wrong version", ".fireweed/canaries/canaries.yaml", "version: 2\nbaseline: {}\n", "canary", "canaries.yaml", "undeclared"},
		{"malformed canary", ".fireweed/canaries/canaries.yaml", "version: [bad\n", "canary", "canaries.yaml", "undeclared"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			if tc.name == "missing baseline" {
				body := string(mustRead(t, filepath.Join(f.root, tc.path)))
				start, end := strings.Index(body, "baseline:\n"), strings.Index(body, "undeclared:\n")
				if start < 0 || end <= start {
					t.Fatal("fixture cannot isolate baseline map")
				}
				tc.body = body[:start] + "baseline: {}\n" + body[end:]
			}
			if tc.name == "missing internal_eval" {
				body := string(mustRead(t, filepath.Join(f.root, tc.path)))
				tc.body = strings.Replace(body, "internal_eval:\n  internal/private.txt: \"private\"\n  evals/internal.txt: \"{{module_dir}}\"\n", "", 1)
				if tc.body == body {
					t.Fatal("fixture did not remove internal_eval map")
				}
			}
			if tc.body == "" {
				if err := os.Remove(filepath.Join(f.root, tc.path)); err != nil {
					t.Fatal(err)
				}
			} else {
				f.put(tc.path, tc.body)
			}
			f.commit()
			report := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(report, "C14", false, tc.term) {
				t.Errorf("missing C14: %s", doctorText(report))
			}
			identified := false
			for _, d := range report.Diagnostics {
				if d.Code == "C14" && !d.Warning && strings.Contains(strings.ToLower(d.Message+" "+d.Path), strings.ToLower(tc.input)) {
					identified = true
				}
			}
			if !identified {
				t.Errorf("C14 did not identify %s: %s", tc.input, doctorText(report))
			}
			if hasDoctorCheck(report, tc.affected, true) {
				t.Errorf("invalid recipe passed canary: %s", doctorText(report))
			}
		})
	}
}

func Test_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_ZeroHookAndRecovery(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/boundary.sh", "#!/bin/sh\nexit 0\n")
	f.put(".fireweed/eval.sh", `#!/bin/sh
printf '<testsuite><testcase name="ALP-001"/></testsuite>' > "$3"
`)
	f.commit()
	before := worktrees(f)
	report := checks.Doctor(context.Background(), f.request(true))
	for _, tc := range []struct{ name, hook string }{{"undeclared", "boundar"}, {"internal_eval", "boundar"}, {"failing_eval", "eval"}, {"missing_id", "eval"}} {
		if !hasDoctorDiagnostic(report, "C14", false, tc.name, tc.hook) {
			t.Errorf("missed %s canary not diagnosed with hook: %s", tc.name, doctorText(report))
		}
	}
	if !doctorErrors(report) {
		t.Errorf("zero hooks accepted: %s", doctorText(report))
	}
	if got := worktrees(f); got != before {
		t.Errorf("failed canary leaked worktree: %s", got)
	}

	g := localDoctorFixture(t)
	g.put(".fireweed/boundary.sh", fmt.Sprintf(`#!/bin/sh
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test -f "$dir/undeclared.txt"; then printf '%%s' "$name" > .fireweed/recovery-seen; printf '%%s: undeclared import\n' "$name"; fi
done
if test -f .fireweed/recovery-seen && ! find packages -name undeclared.txt | grep -q .; then printf '%%s: lingering violation\n' "$(cat .fireweed/recovery-seen)"; fi
`))
	g.commit()
	r := checks.Doctor(context.Background(), g.request(true))
	if !hasDoctorDiagnostic(r, "C14", false, "recovery", "undeclared", "boundar") {
		t.Errorf("missing failed recovery diagnostic: %s", doctorText(r))
	}
}

func Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific(t *testing.T) {
	for _, tc := range []struct{ name, script, canary string }{
		{"missing baseline coverage", `#!/bin/sh
printf '<testsuite/>' > "$3"
`, "baseline"},
		{"failing eval compile error without report", `#!/bin/sh
if test -f "$2/evals/failing.txt"; then echo 'compile error' >&2; exit 2; fi
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "failing_eval"},
		{"failing eval with unrelated failure", `#!/bin/sh
if test -f "$2/evals/failing.txt"; then printf '<testsuite><testcase name="ALP-001"><failure/></testcase></testsuite>' > "$3"; exit 1; fi
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "failing_eval"},
		{"missing id with unrelated failure", `#!/bin/sh
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
if test -f "$2/evals/missing.txt"; then printf '<testsuite><testcase name="%s-001"><failure/></testcase></testsuite>' "$prefix" > "$3"; exit 1; fi
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "missing_id"},
		{"missing id without report", `#!/bin/sh
if test -f "$2/evals/missing.txt"; then echo 'hook failed' >&2; exit 2; fi
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "missing_id"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			f.put(".fireweed/eval.sh", tc.script)
			if tc.name == "missing baseline coverage" {
				logDoctorCanaryProbes(t, f)
			}
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(r, "C14", false, tc.canary, "eval") {
				t.Errorf("unrelated eval outcome accepted: %s", doctorText(r))
			}
			if tc.name == "missing baseline coverage" && !strings.Contains(doctorLog(f), "probe:") {
				t.Errorf("failed baseline coverage hook did not run: %s", doctorLog(f))
			}
			if tc.name == "missing baseline coverage" && strings.Contains(doctorLog(f), "planted:") {
				t.Errorf("failed baseline coverage continued into canary probes: %s", doctorLog(f))
			}
		})
	}
}

func Test_CHK_055_BaselineAndRecoveryFailures(t *testing.T) {
	for _, tc := range []struct{ name, path, script, phase, hook string }{
		{"baseline eval failure", ".fireweed/eval.sh", "#!/bin/sh\necho baseline-eval-failed >&2\nexit 9\n", "baseline", "eval"},
		{"baseline boundary failure", ".fireweed/boundary.sh", "#!/bin/sh\nprintf 'alpha: baseline violation\\n'\n", "baseline", "boundar"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			f.put(tc.path, tc.script)
			logDoctorCanaryProbes(t, f)
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(r, "C14", false, tc.phase, tc.hook) {
				t.Errorf("failed baseline accepted: %s", doctorText(r))
			}
			if !strings.Contains(doctorLog(f), "probe:") {
				t.Errorf("failed baseline hook did not run: %s", doctorLog(f))
			}
			if strings.Contains(doctorLog(f), "planted:") {
				t.Errorf("failed baseline continued into canary probes: %s", doctorLog(f))
			}
		})
	}
	for _, tc := range []struct{ name, failure string }{
		{"recovery eval failure", "echo recovery-eval-failed >&2; exit 9"},
		{"recovery coverage failure", "printf '<testsuite/>' > \"$3\"; exit 0"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			original := string(mustRead(t, filepath.Join(f.root, ".fireweed/eval.sh")))
			injection := "if test -f \"$2/evals/failing.txt\"; then printf '%s' \"$2\" > .fireweed/eval-recovery-module; fi\n" +
				"if test -f .fireweed/eval-recovery-module && test \"$(cat .fireweed/eval-recovery-module)\" = \"$2\" && ! test -f \"$2/evals/failing.txt\"; then " + tc.failure + "; fi\n"
			f.put(".fireweed/eval.sh", strings.Replace(original, "#!/bin/sh\n", "#!/bin/sh\n"+injection, 1))
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(r, "C14", false, "recovery", "failing_eval", "eval") {
				t.Errorf("failed %s accepted: %s", tc.name, doctorText(r))
			}
		})
	}
}

func Test_CHK_056_CHK_057_BoundaryViolationMustNameCanaryModule(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/boundary.sh", `#!/bin/sh
if find packages \( -name undeclared.txt -o -name internal.txt \) | grep -q .; then printf 'alpha: unrelated violation\n'; fi
`)
	f.commit()
	r := checks.Doctor(context.Background(), f.request(true))
	for _, name := range []string{"undeclared", "internal_eval"} {
		if !hasDoctorDiagnostic(r, "C14", false, name, "boundar") {
			t.Errorf("unrelated module violation accepted for %s: %s", name, doctorText(r))
		}
	}
}

func Test_CHK_057_InternalEvalRequiresInterfaceViolation(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/boundary.sh", `#!/bin/sh
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test -f "$dir/undeclared.txt"; then printf '%s: undeclared import\n' "$name"; fi
  if test -f "$dir/evals/internal.txt"; then printf '%s: unrelated boundary violation\n' "$name"; fi
done
`)
	f.commit()
	r := checks.Doctor(context.Background(), f.request(true))
	passedProbe := false
	for _, check := range r.Checks {
		if check.Passed && (check.Name == "internal_eval" || strings.HasPrefix(check.Name, "internal_eval ")) {
			passedProbe = true
		}
	}
	if !hasDoctorDiagnostic(r, "C14", false, "internal_eval", "boundar") || passedProbe {
		t.Errorf("same-module non-interface violation accepted: %s", doctorText(r))
	}
	if !hasDoctorCheck(r, "undeclared", true) {
		t.Errorf("fixture did not detect the independent undeclared canary: %s", doctorText(r))
	}
}

func Test_CHK_055_CHK_056_CHK_057_MissingRequiredBoundaryHook(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	f.put("packages/alpha/internal/live.txt", "disposable implementation\n")
	logDoctorCanaryProbes(t, f)
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(report, "X02", false, "alpha", "boundar") {
		t.Errorf("required boundary hook was not diagnosed: %s", doctorText(report))
	}
	if !hasDoctorDiagnostic(report, "C14", false, "baseline", "boundar") {
		t.Errorf("missing boundary hook was accepted as a clean baseline: %s", doctorText(report))
	}
	if strings.Contains(doctorLog(f), "planted:") {
		t.Errorf("missing baseline boundary hook continued into canary probes: %s", doctorLog(f))
	}
	for _, canary := range []string{"undeclared", "internal_eval"} {
		if hasDoctorCheck(report, canary, true) {
			t.Errorf("missing boundary hook counted as %s detection: %s", canary, doctorText(report))
		}
		if !hasDoctorDiagnostic(report, "C14", false, canary, "boundar") {
			t.Errorf("missing boundary hook lacked named %s canary failure: %s", canary, doctorText(report))
		}
	}
}

func Test_CHK_055_MissingRequiredHookBesideOverride(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	f.put(".fireweed/override-boundary.sh", fmt.Sprintf("#!/bin/sh\nprintf 'override-boundaries:%%s\\n' \"$PWD\" >> %q\nexec sh .fireweed/boundary.sh\n", f.log))
	f.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(f.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nhooks:\n  boundaries: sh .fireweed/override-boundary.sh", 1))
	f.put("packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: []\n---\n# Beta\n\n## Specs\n\n### BET-001\nThe fixture MUST pass.\n")
	f.put("packages/beta/internal/live.txt", "disposable implementation\n")
	logDoctorCanaryProbes(t, f)
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(report, "X02", false, "beta", "boundar") {
		t.Errorf("module without its own hook escaped required boundary enforcement: %s", doctorText(report))
	}
	if strings.Contains(doctorLog(f), "planted:") {
		t.Errorf("failed baseline beside override continued into canary probes: %s", doctorLog(f))
	}
	if !strings.Contains(doctorLog(f), "override-boundaries:") {
		t.Errorf("applicable override was not run: %s", doctorLog(f))
	}
}

func Test_CHK_055_ApplicableOverrideNeedsNoGlobalHook(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	f.put(".fireweed/override-boundary.sh", "#!/bin/sh\nexec sh .fireweed/boundary.sh\n")
	f.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(f.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nhooks:\n  boundaries: sh .fireweed/override-boundary.sh", 1))
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) || !hasDoctorCheck(report, "undeclared", true) || !hasDoctorCheck(report, "internal_eval", true) {
		t.Errorf("applicable override was rejected without a global hook: %s", doctorText(report))
	}
}

func Test_CHK_060_CHK_062_ModuleOverridesAndBoundaryWarnings(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/override-eval.sh", fmt.Sprintf("#!/bin/sh\nprintf 'override-eval:%%s\\n' \"$PWD\" >> %q\nexec sh .fireweed/eval.sh \"$@\"\n", f.log))
	f.put(".fireweed/override-boundary.sh", fmt.Sprintf("#!/bin/sh\nprintf 'override-boundaries:%%s:%%s\\n' \"$PWD\" \"$1\" >> %q\nexec sh .fireweed/boundary.sh\n", f.log))
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\nhooks:\n  eval: sh .fireweed/override-eval.sh {module} {module_dir} {report}\n  boundaries: sh .fireweed/override-boundary.sh {module}\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nThe fixture MUST pass.\n")
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) || strings.Count(doctorLog(f), "override-eval:") < 2 || strings.Count(doctorLog(f), "override-boundaries:") < 2 {
		t.Errorf("override probes absent: %s", doctorLog(f))
	}
	g := localDoctorFixture(t)
	g.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	g.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(g.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nboundaries_opt_out: legacy", 1))
	g.commit()
	r := checks.Doctor(context.Background(), g.request(true))
	if doctorErrors(r) || !hasDoctorDiagnostic(r, "X03", true, "legacy") || !hasDoctorDiagnostic(r, "X04", true, "no boundary enforcement required") {
		t.Errorf("optional boundaries warning lost: %s", doctorText(r))
	}
}

func Test_CHK_060_OverridesMustDetectTheirCanaries(t *testing.T) {
	for _, tc := range []struct {
		name, hook, command, script, canary string
	}{
		{"eval", "eval", "sh .fireweed/override-eval.sh {module} {module_dir} {report}", `#!/bin/sh
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "failing_eval"},
		{"boundary", "boundaries", "sh .fireweed/override-boundary.sh {module}", "#!/bin/sh\nexit 0\n", "internal_eval"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			path := ".fireweed/override-" + tc.name + ".sh"
			f.put(path, tc.script)
			f.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(f.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nhooks:\n  "+tc.hook+": "+tc.command, 1))
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			diagnosed := false
			for _, d := range r.Diagnostics {
				message := strings.ToLower(d.Message)
				boundaryHook := tc.name == "boundary" && !strings.Contains(message, "global") &&
					(strings.Contains(message, "alpha boundary hook") ||
						strings.Contains(message, "alpha boundaries hook") ||
						strings.Contains(message, "alpha boundaries override boundary hook") ||
						strings.Contains(message, "alpha boundaries override boundaries hook"))
				if d.Code == "C14" && !d.Warning && strings.Contains(message, tc.canary) &&
					(tc.name == "eval" && strings.Contains(message, "alpha "+tc.hook+" override") || boundaryHook) {
					diagnosed = true
				}
			}
			if !diagnosed {
				t.Errorf("broken %s override missed its %s canary without C14: %s", tc.name, tc.canary, doctorText(r))
			}
		})
	}
}

func Test_CHK_063_CHK_064_CHK_065_CHK_066_CHK_067_AgentProbes(t *testing.T) {
	f := localDoctorFixture(t)
	f.commit()
	for _, d := range project.Check(f.root) {
		if !d.Warning {
			t.Fatalf("invalid agent fixture: %+v", d)
		}
	}
	r := checks.Doctor(context.Background(), f.request(false))
	log := doctorLog(f)
	if doctorErrors(r) || strings.Count(log, "worker:") != 1 || strings.Count(log, "audit:") != 1 || strings.Count(log, "wrap:") != 2 || !strings.Contains(log, "FIREWEED_DOCTOR_OK") {
		t.Errorf("worker/audit/wrap probes: %s; log=%s", doctorText(r), log)
	}
	workerPrompt, auditPrompt := false, false
	for _, line := range strings.Split(log, "\n") {
		if fields := strings.SplitN(line, ":", 3); len(fields) == 3 && strings.Contains(fields[2], "FIREWEED_DOCTOR_OK") {
			switch fields[0] {
			case "worker":
				workerPrompt = true
			case "audit":
				auditPrompt = true
			}
		}
	}
	if !workerPrompt {
		t.Errorf("worker did not receive its fixed-reply prompt: %s", log)
	}
	if !auditPrompt {
		t.Errorf("audit did not receive its fixed-reply prompt: %s", log)
	}
	for _, part := range []string{"/fireweed/worktrees/", "worker:", "audit:"} {
		if !strings.Contains(log, part) {
			t.Errorf("scratch/role marker %s absent: %s", part, log)
		}
	}

	g := localDoctorFixture(t)
	g.editConfig(func(config map[string]any) {
		delete(config["audit"].(map[string]any), "command")
		config["agent"].(map[string]any)["command"] = "bash .fireweed/agent.sh {prompt} {git_common_dir}"
	})
	g.commit()
	rr := checks.Doctor(context.Background(), g.request(false))
	if !hasDoctorDiagnostic(rr, "C15", true, "shared") || strings.Count(doctorLog(g), "worker:") != 2 {
		t.Errorf("audit fallback: %s; log=%s", doctorText(rr), doctorLog(g))
	}

	h := localDoctorFixture(t)
	h.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/bash\nfor i in $(seq 1 25); do printf 'line-%%02d\\n' \"$i\"; done\necho RATE LIMIT reached\nexit 9\n"))
	h.commit()
	bad := checks.Doctor(context.Background(), h.request(false))
	if !hasDoctorDiagnostic(bad, "C15", false, "usage limit", "rate limit", "line-25") {
		t.Errorf("usage-limit tail missing: %s", doctorText(bad))
	}
	if strings.Contains(doctorText(bad), "line-01") {
		t.Errorf("agent output was not trimmed to final 20 lines: %s", doctorText(bad))
	}
	outOfTail := localDoctorFixture(t)
	outOfTail.put(".fireweed/agent.sh", "#!/bin/bash\necho RATE LIMIT early\nfor i in $(seq 1 25); do printf 'later-%02d\\n' \"$i\"; done\nexit 9\n")
	outOfTail.commit()
	ordinary := checks.Doctor(context.Background(), outOfTail.request(false))
	if !hasDoctorDiagnostic(ordinary, "C15", false, "agent.sh", "later-25") {
		t.Errorf("ordinary out-of-tail failure absent: %s", doctorText(ordinary))
	}
	for _, d := range ordinary.Diagnostics {
		if d.Code == "C15" && !d.Warning && strings.Contains(strings.ToLower(d.Message), "usage limit") {
			t.Errorf("out-of-tail pattern classified as usage limit: %s", doctorText(ordinary))
		}
	}
	for _, tc := range []struct{ name, script string }{{"wrong reply", "#!/bin/bash\necho WRONG_REPLY\n"}, {"nonzero with reply", "#!/bin/bash\necho FIREWEED_DOCTOR_OK\nexit 7\n"}, {"ordinary failure", "#!/bin/bash\nfor i in $(seq 1 25); do printf 'ordinary-%02d\\n' \"$i\"; done\nexit 8\n"}} {
		t.Run(tc.name, func(t *testing.T) {
			q := localDoctorFixture(t)
			q.put(".fireweed/agent.sh", tc.script)
			q.commit()
			r := checks.Doctor(context.Background(), q.request(false))
			if !hasDoctorDiagnostic(r, "C15", false, "agent.sh") {
				t.Errorf("agent failure accepted: %s", doctorText(r))
			}
			if tc.name == "ordinary failure" && (!hasDoctorDiagnostic(r, "C15", false, "ordinary-25") || strings.Contains(doctorText(r), "ordinary-01")) {
				t.Errorf("ordinary failure tail incorrect: %s", doctorText(r))
			}
		})
	}
	wrongAudit := localDoctorFixture(t)
	wrongAudit.put(".fireweed/audit.sh", "#!/bin/bash\necho WRONG_REPLY\n")
	wrongAudit.commit()
	auditFailure := checks.Doctor(context.Background(), wrongAudit.request(false))
	if !hasDoctorDiagnostic(auditFailure, "C15", false, "audit.sh") || hasDoctorCheck(auditFailure, "audit", true) {
		t.Errorf("wrong audit reply accepted: %s", doctorText(auditFailure))
	}

	i := localDoctorFixture(t)
	i.put(".fireweed/audit.sh", "#!/bin/bash\nsleep 2\n")
	i.commit()
	req := i.request(false)
	req.Timeout = 100 * time.Millisecond
	start := time.Now()
	timed := checks.Doctor(context.Background(), req)
	if !hasDoctorDiagnostic(timed, "C15", false, "audit") || time.Since(start) > 4*time.Second {
		t.Errorf("audit timeout not bounded: %s; elapsed=%s", doctorText(timed), time.Since(start))
	}

	defaulted := localDoctorFixture(t)
	defaulted.put(".fireweed/agent.sh", "#!/bin/bash\nsleep 0.2\nprintf 'FIREWEED_DOCTOR_OK\\n'\n")
	defaulted.put(".fireweed/audit.sh", "#!/bin/bash\nsleep 0.2\nprintf 'FIREWEED_DOCTOR_OK\\n'\n")
	defaulted.commit()
	defaultRequest := defaulted.request(false)
	defaultRequest.Timeout = 0
	defaultReport := checks.Doctor(context.Background(), defaultRequest)
	if doctorErrors(defaultReport) || !hasDoctorCheck(defaultReport, "worker", true) || !hasDoctorCheck(defaultReport, "audit", true) {
		t.Errorf("unset timeout did not give both agent probes their default allowance: %s", doctorText(defaultReport))
	}

	j := localDoctorFixture(t)
	j.commit()
	skip := checks.Doctor(context.Background(), j.request(true))
	if strings.Contains(doctorLog(j), "worker:") || strings.Contains(doctorLog(j), "audit:") || !hasDoctorCheck(skip, "undeclared", true) || !strings.Contains(doctorLog(j), j.bin) {
		t.Errorf("NoAgent skipped hooks or ran agent: %s; log=%s", doctorText(skip), doctorLog(j))
	}
}

func Test_CHK_065_UnsetTimeoutBoundsAuditProbe(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/audit.sh", fmt.Sprintf("#!/bin/bash\nprintf 'audit:%%s:%%s\n' \"$PWD\" \"$1\" >> %q\nprintf 'audit-timeout-started\n'\nsleep 210\nprintf 'FIREWEED_DOCTOR_OK\n'\n", f.log))
	f.commit()
	req := f.request(false)
	req.Timeout = 0

	start := time.Now()
	report := checks.Doctor(context.Background(), req)
	elapsed := time.Since(start)
	log := doctorLog(f)
	if !hasDoctorCheck(report, "worker", true) || strings.Count(log, "worker:") != 1 || strings.Count(log, "audit:") != 1 {
		t.Errorf("agent probes did not run once with immediate worker reply: %s; log=%s", doctorText(report), log)
	}
	if !hasDoctorCheck(report, "audit", false) || !hasDoctorDiagnostic(report, "C15", false, "audit.sh", "audit-timeout-started") {
		t.Errorf("long audit probe lacked a C15 timeout with its command and output: %s", doctorText(report))
	}
	if elapsed < 165*time.Second || elapsed > 200*time.Second {
		t.Errorf("unset audit timeout took %s, want a bound near three minutes", elapsed)
	}
}

func Test_CHK_068_CHK_069_ManagedGuidanceAndPlaybookBase(t *testing.T) {
	f := localDoctorFixture(t)
	text, err := assets.Playbook("evals")
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256([]byte(text))
	base := hex.EncodeToString(sum[:])
	f.put(".fireweed/playbooks/evals.md", "---\nembedded_base: "+base+"\n---\n# Local override\n")
	f.commit()
	r := checks.Doctor(context.Background(), f.request(true))
	if hasDoctorDiagnostic(r, "C16", true) {
		t.Errorf("current guidance/base warned: %s", doctorText(r))
	}
	f.put("AGENTS.md", "Before.\n"+doctorGuidance+"After.\n")
	surrounded := checks.Doctor(context.Background(), f.request(true))
	if hasDoctorDiagnostic(surrounded, "C16", true) {
		t.Errorf("current managed guidance with surrounding text warned: %s", doctorText(surrounded))
	}
	f.put("AGENTS.md", "Before.\n<!-- fireweed:start -->\nold guidance\n<!-- fireweed:end -->\nAfter.\n")
	f.put(".fireweed/playbooks/evals.md", "---\nembedded_base: stale\n---\n# Local override\n")
	stale := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(stale, "C16", true, "AGENTS.md") || !hasDoctorDiagnostic(stale, "C16", true, "evals") {
		t.Errorf("stale guidance/base not warned: %s", doctorText(stale))
	}
	if err := os.Remove(filepath.Join(f.root, "AGENTS.md")); err != nil {
		t.Fatal(err)
	}
	f.put(".fireweed/playbooks/evals.md", "# No embedded base\n")
	missing := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(missing, "C16", true, "AGENTS.md") || !hasDoctorDiagnostic(missing, "C16", true, "evals") {
		t.Errorf("missing guidance/base not warned: %s", doctorText(missing))
	}
}

func mustRead(t *testing.T, path string) []byte {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return b
}
func doctorBinary(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("caller")
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(file), "..", "..", ".."))
	bin := filepath.Join(t.TempDir(), "fireweed-new")
	cmd := exec.Command("go", "build", "-o", bin, ".")
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GOCACHE=/tmp/fireweed-rebuild-cache")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build product: %v: %s", err, out)
	}
	return bin
}
package evals_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
)

func doctorProbeLog(t *testing.T, f doctorFixture, report checks.DoctorReport, role string, markers ...string) string {
	t.Helper()
	var detail string
	for _, check := range report.Checks {
		if check.Name == role {
			detail = check.Detail
			break
		}
	}
	if detail == "" {
		t.Errorf("executed %s probe has no check: %s", role, doctorText(report))
		return ""
	}
	at := strings.Index(detail, "log: ")
	if at < 0 {
		t.Errorf("%s detail lacks log path: %q", role, detail)
		return ""
	}
	fields := strings.Fields(detail[at+len("log: "):])
	if len(fields) == 0 {
		t.Errorf("%s detail has empty log path: %q", role, detail)
		return ""
	}
	path := fields[0]
	if !filepath.IsAbs(path) {
		t.Errorf("%s log path is not absolute: %q", role, path)
	}
	rel, err := filepath.Rel(filepath.Join(f.root, ".fireweed", "runs"), path)
	if err != nil {
		t.Errorf("%s log path %q: %v", role, path, err)
		return ""
	}
	parts := strings.Split(filepath.ToSlash(rel), "/")
	if len(parts) != 2 || parts[0] == "" || parts[0] == "." || parts[0] == ".." || parts[1] != "doctor-"+role+"-1.log" {
		t.Errorf("%s log path = %q, want .fireweed/runs/<run>/doctor-%s-1.log", role, path, role)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("read %s log %q: %v", role, path, err)
		return path
	}
	for _, marker := range markers {
		if !strings.Contains(string(data), marker) {
			t.Errorf("%s log lacks %q: %q", role, marker, data)
		}
	}
	return path
}

func doctorLogRun(t *testing.T, f doctorFixture, worker, audit string, report checks.DoctorReport) string {
	t.Helper()
	base := filepath.Join(f.root, ".fireweed", "runs")
	w, ew := filepath.Rel(base, worker)
	a, ea := filepath.Rel(base, audit)
	if ew != nil || ea != nil {
		t.Errorf("log paths outside run root: %q %q: %v %v", worker, audit, ew, ea)
		return ""
	}
	wRun, aRun := filepath.Dir(w), filepath.Dir(a)
	if wRun != aRun || wRun == "." || wRun == ".." {
		t.Errorf("worker and audit have different run identities: %q %q: %s", worker, audit, doctorText(report))
	}
	return wRun
}

func doctorCallerClean(t *testing.T, f doctorFixture, status, registrations string) {
	t.Helper()
	if got := f.git("status", "--porcelain", "--untracked-files=all"); got != status {
		t.Errorf("caller status changed: before %q after %q", status, got)
	}
	if got := worktrees(f); got != registrations {
		t.Errorf("scratch registration changed: before %q after %q", registrations, got)
	}
	assertScratchRemoved(t, f)
}

func Test_CHK_081_DoctorProbeLogsShareFreshRunAndPreserveStreams(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/agent.sh", "#!/bin/bash\nprintf 'worker stdout Ω without newline'; printf 'worker stderr β without newline' >&2; printf 'FIREWEED_DOCTOR_OK\\n'\n")
	f.put(".fireweed/audit.sh", "#!/bin/bash\nprintf 'audit stdout Ω without newline'; printf 'audit stderr β without newline' >&2; printf 'FIREWEED_DOCTOR_OK\\n'\n")
	f.commit()
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	var firstRun string
	for i := 0; i < 2; i++ {
		report := checks.Doctor(context.Background(), f.request(false))
		if doctorErrors(report) || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
			t.Errorf("successful probes changed outcome: %s", doctorText(report))
		}
		worker := doctorProbeLog(t, f, report, "worker", "worker stdout Ω without newlineFIREWEED_DOCTOR_OK\n", "worker stderr β without newline")
		audit := doctorProbeLog(t, f, report, "audit", "audit stdout Ω without newlineFIREWEED_DOCTOR_OK\n", "audit stderr β without newline")
		if worker != "" && audit != "" {
			run := doctorLogRun(t, f, worker, audit, report)
			if i == 1 && run == firstRun {
				t.Errorf("consecutive Doctor invocations reused run identity %q", run)
			}
			firstRun = run
		}
		doctorCallerClean(t, f, status, registrations)
	}
}

func Test_CHK_081_DoctorPersistsFailedAndTimedOutProbeOutput(t *testing.T) {
	failureStdout := "early worker stdout Ω\n"
	for i := 1; i <= 25; i++ {
		failureStdout += fmt.Sprintf("later-%02d\n", i)
	}
	for _, tc := range []struct {
		name, role, script, stdout, stderr string
		timeout                            time.Duration
	}{
		{"failure", "worker", "#!/bin/bash\nprintf 'early worker stdout Ω\\n'; for i in $(seq 1 25); do printf 'later-%02d\\n' \"$i\"; done; printf 'failed worker stderr β' >&2; exit 7\n", failureStdout, "failed worker stderr β", 2 * time.Second},
		{"timeout", "audit", "#!/bin/bash\nprintf 'partial audit stdout Ω'; printf 'partial audit stderr β' >&2; sleep 5\n", "partial audit stdout Ω", "partial audit stderr β", 150 * time.Millisecond},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			if tc.role == "worker" {
				f.put(".fireweed/agent.sh", tc.script)
			} else {
				f.put(".fireweed/audit.sh", tc.script)
			}
			f.commit()
			status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
			req := f.request(false)
			req.Timeout = tc.timeout
			start := time.Now()
			report := checks.Doctor(context.Background(), req)
			if time.Since(start) > 4*time.Second {
				t.Errorf("%s probe did not honor bounded request timeout", tc.name)
			}
			if !hasDoctorCheck(report, tc.role, false) || !hasDoctorDiagnostic(report, "C15", false, tc.role, tc.stderr) {
				t.Errorf("%s probe outcome/diagnostic lost: %s", tc.name, doctorText(report))
			}
			path := doctorProbeLog(t, f, report, tc.role, tc.stdout, tc.stderr)
			if tc.name == "failure" && path != "" {
				if !strings.Contains(string(mustRead(t, path)), "later-25") || strings.Contains(doctorText(report), "early worker stdout Ω") {
					t.Errorf("failed probe log or final-20-line diagnostic changed: %s", doctorText(report))
				}
			}
			doctorCallerClean(t, f, status, registrations)
		})
	}
}

func Test_CHK_081_DoctorPersistsCancelledProbeOutput(t *testing.T) {
	f := localDoctorFixture(t)
	ready := filepath.Join(t.TempDir(), "worker-ready")
	f.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/bash\nprintf 'cancel worker stdout Ω'; printf 'cancel worker stderr β' >&2; printf ready > %q; sleep 5\n", ready))
	f.commit()
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reports := make(chan checks.DoctorReport, 1)
	go func() { reports <- checks.Doctor(ctx, f.request(false)) }()
	deadline := time.After(8 * time.Second)
	for {
		if _, err := os.Stat(ready); err == nil {
			break
		} else if !os.IsNotExist(err) {
			t.Fatal(err)
		}
		select {
		case report := <-reports:
			t.Fatalf("worker exited before cancellation: %s", doctorText(report))
		case <-deadline:
			t.Fatal("worker never signalled readiness")
		case <-time.After(10 * time.Millisecond):
		}
	}
	cancel()
	select {
	case report := <-reports:
		if !hasDoctorCheck(report, "worker", false) || !hasDoctorDiagnostic(report, "C15", false, "worker") {
			t.Errorf("cancelled probe outcome lost: %s", doctorText(report))
		}
		doctorProbeLog(t, f, report, "worker", "cancel worker stdout Ω", "cancel worker stderr β")
	case <-time.After(4 * time.Second):
		t.Fatal("cancelled Doctor did not return promptly")
	}
	doctorCallerClean(t, f, status, registrations)
}
package evals_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
)

const doctorProbePrompt = "Reply with exactly FIREWEED_DOCTOR_OK"

func doctorRecordedPath(t *testing.T, log, label string) string {
	t.Helper()
	return doctorRecordedPaths(t, log, label, 1)[0]
}

func doctorRecordedPaths(t *testing.T, log, label string, count int) []string {
	t.Helper()
	var paths []string
	for _, line := range strings.Split(log, "\n") {
		if strings.HasPrefix(line, label+":") {
			paths = append(paths, strings.TrimPrefix(line, label+":"))
		}
	}
	if len(paths) != count {
		t.Fatalf("%s paths = %q; want %d absolute paths", label, paths, count)
	}
	for _, path := range paths {
		if !filepath.IsAbs(path) {
			t.Fatalf("%s path = %q; want an absolute path", label, path)
		}
	}
	return paths
}

func doctorPromptFileRemoved(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Errorf("prompt file %q remains after probe: %v", path, err)
	}
}

func Test_CHK_082_CHK_083_DoctorExpandsSelectedCommandsAndWrapperOnce(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) {
		config["agent"].(map[string]any)["command"] = []any{"bash", ".fireweed/agent.sh", "pre{prompt}post", "file={prompt_file}", "common={git_common_dir}", "literal {unknown}"}
		config["audit"].(map[string]any)["command"] = "bash .fireweed/audit.sh {prompt} {prompt_file} {git_common_dir}"
		config["agent"].(map[string]any)["wrap"] = []any{"bash", ".fireweed/prompt-wrap.sh", "{prompt_file}", "{prompt}", "{git_common_dir}"}
	})
	f.put(".fireweed/prompt-wrap.sh", fmt.Sprintf(`#!/bin/bash
printf 'wrap-file:%%s\nwrap-prompt:%%s\nwrap-common:%%s\n' "$1" "$2" "$3" >> %q
test -f "$1" && printf '%%s' "$2" | cmp -s - "$1" || exit 11
shift 3
exec "$@"
`, f.log))
	f.put(".fireweed/agent.sh", fmt.Sprintf(`#!/bin/bash
printf 'worker-prompt:%%s\nworker-file:%%s\nworker-common:%%s\nworker-literal:%%s\n' "$1" "${2#file=}" "${3#common=}" "$4" >> %q
test "$1" = "pre%spost" || exit 12
test "$2" != "file=" || exit 13
test -f "${2#file=}" && printf '%%s' %q | cmp -s - "${2#file=}" || exit 14
printf 'FIREWEED_DOCTOR_OK\n'
`, f.log, doctorProbePrompt, doctorProbePrompt))
	f.put(".fireweed/audit.sh", fmt.Sprintf(`#!/bin/bash
printf 'audit-prompt:%%s\naudit-file:%%s\naudit-common:%%s\n' "$1" "$2" "$3" >> %q
test -f "$2" && printf '%%s' "$1" | cmp -s - "$2" || exit 15
printf 'FIREWEED_DOCTOR_OK\n'
`, f.log))
	f.commit()
	// Relocate the Git directory so its absolute path contains a quote, a space,
	// and another placeholder. The expanded value must remain literal.
	commonDir := filepath.Join(t.TempDir(), "space 'quote' {prompt}", "repo.git")
	if err := os.MkdirAll(filepath.Dir(commonDir), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Rename(filepath.Join(f.root, ".git"), commonDir); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(f.root, ".git"), []byte("gitdir: "+commonDir+"\n"), 0644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("TMPDIR", filepath.Dir(commonDir))
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if doctorErrors(report) || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
		t.Errorf("prompt-file probes failed: %s", doctorText(report))
	}
	log := doctorLog(f)
	common, err := filepath.EvalSymlinks(commonDir)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"worker-prompt:pre" + doctorProbePrompt + "post", "audit-prompt:" + doctorProbePrompt,
		"worker-common:" + common, "audit-common:" + common,
		"worker-literal:literal {unknown}",
		"wrap-prompt:" + doctorProbePrompt, "wrap-common:" + common,
	} {
		if !strings.Contains(log, want+"\n") {
			t.Errorf("missing exact argument %q", want)
		}
	}
	workerFile := doctorRecordedPath(t, log, "worker-file")
	auditFile := doctorRecordedPath(t, log, "audit-file")
	if workerFile == auditFile {
		t.Errorf("worker and audit reused prompt file %q", workerFile)
	}
	for _, path := range []string{workerFile, auditFile} {
		if !strings.Contains(log, "wrap-file:"+path+"\n") {
			t.Errorf("wrapper did not receive selected probe file %q", path)
		}
		doctorPromptFileRemoved(t, path)
	}
	doctorProbeLog(t, f, report, "worker", "FIREWEED_DOCTOR_OK")
	doctorProbeLog(t, f, report, "audit", "FIREWEED_DOCTOR_OK")
	doctorCallerClean(t, f, status, registrations)
}

func Test_CHK_083_DoctorRemovesPromptFileOnProbeFailures(t *testing.T) {
	for _, mode := range []string{"launch", "nonzero", "timeout", "cancel"} {
		t.Run(mode, func(t *testing.T) {
			f := localDoctorFixture(t)
			ready := filepath.Join(t.TempDir(), "ready")
			f.editConfig(func(config map[string]any) {
				config["agent"].(map[string]any)["command"] = []any{"bash", ".fireweed/agent.sh", "{prompt_file}"}
				if mode == "launch" {
					config["agent"].(map[string]any)["command"] = []any{filepath.Join(f.root, "missing-agent"), "{prompt_file}"}
					config["agent"].(map[string]any)["wrap"] = []any{"bash", ".fireweed/prompt-wrap.sh", "{prompt_file}"}
				}
			})
			f.put(".fireweed/prompt-wrap.sh", fmt.Sprintf("#!/bin/bash\nprintf 'worker-file:%%s\\n' \"$1\" >> %q\ntest -f \"$1\" || exit 11\nshift\nexec \"$@\"\n", f.log))
			script := fmt.Sprintf("#!/bin/bash\nprintf 'worker-file:%%s\\nworker-content:%%s\\n' \"$1\" \"$(cat \"$1\")\" >> %q\ntest -f \"$1\" && printf '%%s' %q | cmp -s - \"$1\" || exit 18\nprintf ready > %q\n", f.log, doctorProbePrompt, ready)
			switch mode {
			case "nonzero":
				script += "printf 'failed worker stderr\\n' >&2\nexit 7\n"
			case "timeout", "cancel":
				script += "printf 'waiting worker stdout\\n'\nsleep 5\n"
			}
			f.put(".fireweed/agent.sh", script)
			f.commit()
			status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
			req := f.request(false)
			if mode == "timeout" {
				req.Timeout = 2 * time.Second
			}
			ctx := context.Background()
			var cancel context.CancelFunc
			if mode == "cancel" {
				ctx, cancel = context.WithCancel(ctx)
				defer cancel()
			}
			var report checks.DoctorReport
			if mode == "cancel" {
				reports := make(chan checks.DoctorReport, 1)
				go func() { reports <- checks.Doctor(ctx, req) }()
				deadline := time.After(8 * time.Second)
				for {
					if _, err := os.Stat(ready); err == nil {
						break
					} else if !os.IsNotExist(err) {
						t.Fatal(err)
					}
					select {
					case report = <-reports:
						t.Fatalf("worker exited before cancellation: %s", doctorText(report))
					case <-deadline:
						t.Fatal("worker never signalled readiness")
					case <-time.After(10 * time.Millisecond):
					}
				}
				livePath := doctorRecordedPath(t, doctorLog(f), "worker-file")
				if data, err := os.ReadFile(livePath); err != nil || string(data) != doctorProbePrompt {
					t.Errorf("prompt file unavailable before cancellation: %q, %v", data, err)
				}
				cancel()
				select {
				case report = <-reports:
				case <-time.After(4 * time.Second):
					t.Fatal("cancelled Doctor did not return")
				}
			} else {
				report = checks.Doctor(ctx, req)
			}
			if !hasDoctorCheck(report, "worker", false) || !hasDoctorDiagnostic(report, "C15", false, "worker") {
				t.Errorf("%s probe outcome changed: %s", mode, doctorText(report))
			}
			log := doctorLog(f)
			count := 1
			if mode == "launch" {
				// The shared wrapper records both the failed worker launch and
				// the subsequent audit probe; both files must be cleaned up.
				count = 2
			}
			paths := doctorRecordedPaths(t, log, "worker-file", count)
			if mode != "launch" && !strings.Contains(log, "worker-content:"+doctorProbePrompt+"\n") {
				t.Errorf("%s worker could not read exact prompt: %q", mode, log)
			}
			for _, path := range paths {
				doctorPromptFileRemoved(t, path)
			}
			doctorProbeLog(t, f, report, "worker")
			doctorCallerClean(t, f, status, registrations)
		})
	}
}

func Test_CHK_083_DoctorDoesNotCreatePromptFileWithoutSelectedPlaceholder(t *testing.T) {
	f := localDoctorFixture(t)
	expected := filepath.Join(t.TempDir(), "expected-prompt")
	if err := os.WriteFile(expected, []byte(doctorProbePrompt), 0644); err != nil {
		t.Fatal(err)
	}
	script := fmt.Sprintf(`#!/bin/bash
while IFS= read -r candidate; do
  if test -f "$candidate" && cmp -s "$candidate" %q; then
    printf 'unexpected prompt file: %%s\n' "$candidate" >&2
    exit 19
  fi
done < <(find "$TMPDIR" "$PWD" -type f -print)
printf 'FIREWEED_DOCTOR_OK\n'
`, expected)
	f.put(".fireweed/agent.sh", script)
	f.put(".fireweed/audit.sh", script)
	f.commit()
	temp := t.TempDir()
	t.Setenv("TMPDIR", temp)
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if doctorErrors(report) || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
		t.Errorf("unneeded prompt file appeared: %s", doctorText(report))
	}
	doctorProbeLog(t, f, report, "worker", "FIREWEED_DOCTOR_OK")
	doctorProbeLog(t, f, report, "audit", "FIREWEED_DOCTOR_OK")
	doctorCallerClean(t, f, status, registrations)
}

func Test_CHK_082_CHK_083_DoctorWrapperOnlyFileWithAuditFallback(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) {
		delete(config["audit"].(map[string]any), "command")
		config["agent"].(map[string]any)["command"] = []any{"bash", ".fireweed/agent.sh", "{prompt}"}
		config["agent"].(map[string]any)["wrap"] = []any{"bash", ".fireweed/prompt-wrap.sh", "{prompt_file}"}
	})
	f.put(".fireweed/prompt-wrap.sh", fmt.Sprintf(`#!/bin/bash
printf 'wrap-file:%%s\n' "$1" >> %q
test -f "$1" && printf '%%s' %q | cmp -s - "$1" || exit 11
shift
exec "$@"
`, f.log, doctorProbePrompt))
	f.put(".fireweed/agent.sh", fmt.Sprintf(`#!/bin/bash
printf 'fallback-prompt:%%s\n' "$1" >> %q
test "$1" = %q || exit 12
printf 'FIREWEED_DOCTOR_OK\n'
`, f.log, doctorProbePrompt))
	f.commit()
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if doctorErrors(report) || !hasDoctorDiagnostic(report, "C15", true, "shared") || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
		t.Errorf("wrapper-only fallback changed: %s", doctorText(report))
	}
	log := doctorLog(f)
	if strings.Count(log, "fallback-prompt:"+doctorProbePrompt+"\n") != 2 {
		t.Errorf("fallback did not receive fixed prompt twice")
	}
	var paths []string
	for _, line := range strings.Split(log, "\n") {
		if strings.HasPrefix(line, "wrap-file:") {
			paths = append(paths, strings.TrimPrefix(line, "wrap-file:"))
		}
	}
	if len(paths) != 2 || paths[0] == paths[1] {
		t.Errorf("wrapper prompt files = %q; want distinct worker/audit paths", paths)
	}
	for _, path := range paths {
		if !filepath.IsAbs(path) {
			t.Errorf("wrapper prompt path is relative: %q", path)
		}
		doctorPromptFileRemoved(t, path)
	}
	doctorProbeLog(t, f, report, "worker", "FIREWEED_DOCTOR_OK")
	doctorProbeLog(t, f, report, "audit", "FIREWEED_DOCTOR_OK")
	doctorCallerClean(t, f, status, registrations)
}
package evals_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

const historyConfig = `version: 1
base: main
modules: ["packages/*"]
protected: [locked/**]
module_defaults:
  evals: [evals/**]
  interface: [api.go]
  append_only: [ledger/**]
`

type historyFixture struct {
	t    *testing.T
	root string
	base string
}

func (f historyFixture) git(args ...string) string {
	f.t.Helper()
	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Args: append([]string{"git", "-C", f.root}, args...),
		Dir:  f.root,
		Env: map[string]string{
			"GIT_CONFIG_GLOBAL": "/dev/null", "GIT_CONFIG_NOSYSTEM": "1",
			"GIT_AUTHOR_NAME": "History Fixture", "GIT_AUTHOR_EMAIL": "fixture@example.invalid",
			"GIT_COMMITTER_NAME": "History Fixture", "GIT_COMMITTER_EMAIL": "fixture@example.invalid",
		},
	})
	if err != nil {
		f.t.Fatalf("git %v: %v: %s", args, err, result.Stderr)
	}
	return strings.TrimSpace(result.Stdout)
}

func (f historyFixture) put(path, body string) {
	f.t.Helper()
	full := filepath.Join(f.root, filepath.FromSlash(path))
	if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
		f.t.Fatal(err)
	}
	if err := os.WriteFile(full, []byte(body), 0644); err != nil {
		f.t.Fatal(err)
	}
}

func (f historyFixture) remove(path string) {
	f.t.Helper()
	if err := os.Remove(filepath.Join(f.root, filepath.FromSlash(path))); err != nil {
		f.t.Fatal(err)
	}
}

func (f historyFixture) move(oldPath, newPath string) {
	f.t.Helper()
	if err := os.MkdirAll(filepath.Dir(filepath.Join(f.root, filepath.FromSlash(newPath))), 0755); err != nil {
		f.t.Fatal(err)
	}
	f.git("mv", oldPath, newPath)
}

func (f historyFixture) commit(subject string) string {
	f.t.Helper()
	f.git("add", "-A")
	f.git("commit", "-qm", subject)
	return f.git("rev-parse", "HEAD")
}

func (f historyFixture) change(path, phase, set string) {
	f.t.Helper()
	body := fmt.Sprintf("---\nmodule: alpha\nphase: %s\n%s---\n\n## Motivation\nHistory fixture.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n\n## Notes\n", phase, set)
	if _, err := contracts.ParseDocument(body); err != nil {
		f.t.Fatalf("fixture change: %v", err)
	}
	f.put(path, body)
}

func newHistoryFixture(t *testing.T) historyFixture {
	t.Helper()
	f := historyFixture{t: t, root: t.TempDir()}
	f.git("init", "-q", "-b", "main")
	f.put("fireweed.yaml", historyConfig)
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha behavior.\n")
	f.put("packages/alpha/api.go", "package alpha\n")
	f.put("packages/alpha/evals/existing_test.go", "package evals\n")
	f.put("packages/alpha/ledger/entry.txt", "initial\n")
	f.put("packages/alpha/newledger/entry.txt", "initial\n")
	f.put("locked/config.txt", "initial\n")
	f.base = f.commit("base")
	if _, err := project.LoadAt(f.root, f.base); err != nil {
		t.Fatalf("recorded base project: %v", err)
	}
	return f
}

func (f historyFixture) report(revisionRange string) checks.HistoryReport {
	f.t.Helper()
	report, err := checks.History(f.root, revisionRange)
	if err != nil {
		f.t.Fatalf("History(%q): %v", revisionRange, err)
	}
	return report
}

type historyFinding struct{ code, path, commit string }

func expectHistoryFindings(t *testing.T, report checks.HistoryReport, want ...historyFinding) {
	t.Helper()
	if len(report.Diagnostics) != len(want) {
		t.Fatalf("diagnostics = %#v; want %d findings", report.Diagnostics, len(want))
	}
	used := make([]bool, len(report.Diagnostics))
	for _, expected := range want {
		found := false
		for i, d := range report.Diagnostics {
			if !used[i] && d.Code == expected.code && d.Path == expected.path && strings.Contains(d.Message, expected.commit) {
				used[i], found = true, true
				break
			}
		}
		if !found {
			t.Errorf("missing %s on %s at %s in %#v", expected.code, expected.path, expected.commit, report.Diagnostics)
		}
	}
}

func Test_CHK_070_HistoryExplicitRangeOrderAndFullOIDs(t *testing.T) {
	f := newHistoryFixture(t)
	f.put("notes/one.txt", "one")
	one := f.commit("first")
	f.put("notes/two.txt", "two")
	two := f.commit("second")
	got := f.report(f.base + ".." + two)
	if !reflect.DeepEqual(got.Commits, []string{one, two}) {
		t.Errorf("commits = %#v; want [%s %s]", got.Commits, one, two)
	}
	expectHistoryFindings(t, got)
}

func Test_CHK_070_HistoryBranchingRangeTopologicalOrder(t *testing.T) {
	f := newHistoryFixture(t)
	f.git("branch", "side")
	f.put("notes/main-one.txt", "main one\n")
	mainOne := f.commit("main first")
	f.put("notes/main-two.txt", "main two\n")
	mainTwo := f.commit("main second")
	f.git("checkout", "-q", "side")
	f.put("notes/side-one.txt", "side one\n")
	sideOne := f.commit("side first")
	f.put("notes/side-two.txt", "side two\n")
	sideTwo := f.commit("side second")
	f.git("checkout", "-q", "main")
	f.git("merge", "--no-ff", "-m", "merge side", "side")
	merge := f.git("rev-parse", "HEAD")

	got := f.report(f.base + ".." + merge)
	want := map[string]bool{mainOne: true, mainTwo: true, sideOne: true, sideTwo: true, merge: true}
	if len(got.Commits) != len(want) {
		t.Errorf("branching commits = %#v; want exactly five commits", got.Commits)
	}
	positions := make(map[string]int)
	for i, commit := range got.Commits {
		if !want[commit] {
			t.Errorf("unexpected commit %s in %#v", commit, got.Commits)
		}
		if _, duplicate := positions[commit]; duplicate {
			t.Errorf("duplicate commit %s in %#v", commit, got.Commits)
		}
		positions[commit] = i
	}
	for commit := range want {
		if _, present := positions[commit]; !present {
			t.Errorf("missing commit %s in %#v", commit, got.Commits)
		}
	}
	for _, edge := range [][2]string{{mainOne, mainTwo}, {sideOne, sideTwo}, {mainTwo, merge}, {sideTwo, merge}} {
		parent, hasParent := positions[edge[0]]
		child, hasChild := positions[edge[1]]
		if hasParent && hasChild && parent >= child {
			t.Errorf("parent %s must precede child %s in %#v", edge[0], edge[1], got.Commits)
		}
	}
	expectHistoryFindings(t, got)
}

func Test_CHK_070_HistoryDefaultRangeUsesLocalConfiguredBase(t *testing.T) {
	f := newHistoryFixture(t)
	f.put("fireweed.yaml", strings.Replace(historyConfig, "base: main", "base: review-base", 1))
	f.commit("configure review base")
	f.git("branch", "review-base")
	f.put("notes/common.txt", "common\n")
	common := f.commit("common after review base")
	f.git("branch", "feature")
	f.git("checkout", "-q", "feature")
	f.put("notes/one.txt", "one")
	one := f.commit("first")
	f.put("notes/two.txt", "two")
	two := f.commit("second")
	f.git("checkout", "-q", "review-base")
	f.put("notes/review-only.txt", "review only\n")
	reviewOnly := f.commit("review base advances independently")
	f.git("checkout", "-q", "feature")
	got := f.report("")
	if !reflect.DeepEqual(got.Commits, []string{common, one, two}) {
		t.Errorf("default commits = %#v; want [%s %s %s]", got.Commits, common, one, two)
	}
	for _, commit := range got.Commits {
		if commit == reviewOnly {
			t.Errorf("default commits include configured-base-only commit %s", reviewOnly)
		}
	}
}

func Test_CHK_071_CHK_072_HistoryPhaseContentsAndDeletionControl(t *testing.T) {
	f := newHistoryFixture(t)
	path := "changes/0001-work.md"
	f.change(path, "evals", "")
	f.commit("add queued change")
	f.change(path, "audit", "")
	f.put("locked/config.txt", "worker edit\n")
	worker := f.commit("ordinary subject")
	f.remove(path)
	f.put("locked/config.txt", "author edit\n")
	f.commit("fireweed impl 0001-work")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got, historyFinding{"H01", "locked/config.txt", worker})
}

func Test_CHK_072_CHK_073_CHK_077_HistoryPlaybookOverrideAndProjectedInterface(t *testing.T) {
	f := newHistoryFixture(t)
	pending := "changes/0001-work.md"
	archive := "changes/archive/0001-work.md"
	set := "set:\n  interface: [newguard/**]\n"
	f.change(pending, "impl", set)
	f.commit("ready")
	projected, err := project.Project(f.root, "")
	if err != nil {
		t.Fatalf("Project with interface set: %v", err)
	}
	if len(projected.Modules) != 1 || projected.Modules[0].Name != "alpha" || !reflect.DeepEqual(projected.Modules[0].Zones.Interface, []string{"newguard/**"}) {
		t.Fatalf("projected alpha interface = %#v; want [newguard/**]", projected.Modules)
	}
	f.move(pending, archive)
	f.change(archive, "applied", set)
	f.put("packages/alpha/newguard/api.go", "package newguard\n")
	f.put(".fireweed/playbooks/evals.md", "local override\n")
	applied := f.commit("phase assigned")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got,
		historyFinding{"H02", "packages/alpha/newguard/api.go", applied},
		historyFinding{"H01", ".fireweed/playbooks/evals.md", applied})
}

func Test_CHK_071_CHK_073_HistoryAppliedImplTouchesEvalsAndInterface(t *testing.T) {
	f := newHistoryFixture(t)
	pending := "changes/0001-work.md"
	archive := "changes/archive/0001-work.md"
	f.change(pending, "impl", "")
	f.commit("ready for implementation")
	f.move(pending, archive)
	f.change(archive, "applied", "")
	f.put("packages/alpha/evals/existing_test.go", "changed test\n")
	f.put("packages/alpha/api.go", "changed api\n")
	applied := f.commit("unrelated subject")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got,
		historyFinding{"H02", "packages/alpha/evals/existing_test.go", applied},
		historyFinding{"H02", "packages/alpha/api.go", applied})
}

func Test_CHK_074_HistoryAuthorAppendOnlyModifyDeleteRename(t *testing.T) {
	for _, action := range []string{"modify", "delete", "rename"} {
		t.Run(action, func(t *testing.T) {
			f := newHistoryFixture(t)
			path := "packages/alpha/ledger/entry.txt"
			switch action {
			case "modify":
				f.put(path, "modified\n")
			case "delete":
				f.remove(path)
			case "rename":
				f.move(path, "packages/alpha/ledger/renamed.txt")
			}
			commit := f.commit("human maintenance")
			got := f.report(f.base + "..HEAD")
			if len(got.Diagnostics) == 0 {
				t.Errorf("%s diagnostics = %#v; want H03", action, got.Diagnostics)
			}
			for _, d := range got.Diagnostics {
				if d.Code != "H03" || !strings.Contains(d.Message, commit) || (d.Path != path && d.Path != "packages/alpha/ledger/renamed.txt") {
					t.Errorf("%s unexpected diagnostic %#v", action, d)
				}
			}
		})
	}
}

func Test_CHK_074_HistoryMergeResolutionInMergeOnlyRange(t *testing.T) {
	f := newHistoryFixture(t)
	f.git("checkout", "-qb", "side")
	f.put("packages/alpha/ledger/entry.txt", "side\n")
	f.commit("side edits ledger")
	f.git("checkout", "-q", "main")
	f.put("packages/alpha/ledger/entry.txt", "main\n")
	f.commit("main edits ledger")
	result, err := gitstore.Run(context.Background(), gitstore.Request{Args: []string{"git", "-C", f.root, "merge", "--no-ff", "--no-commit", "side"}, Dir: f.root})
	if err == nil || !strings.Contains(result.Stdout+result.Stderr, "CONFLICT") {
		t.Fatalf("expected ordinary merge conflict, got %v: %s %s", err, result.Stdout, result.Stderr)
	}
	f.put("packages/alpha/ledger/entry.txt", "resolved uniquely\n")
	merge := f.commit("resolve conflict")
	got := f.report(merge + "^!")
	if !reflect.DeepEqual(got.Commits, []string{merge}) {
		t.Errorf("merge-only selection = %#v; want [%s]", got.Commits, merge)
	}
	expectHistoryFindings(t, got, historyFinding{"H03", "packages/alpha/ledger/entry.txt", merge})
}

func Test_CHK_075_HistoryArchiveModifyAndDeletionControl(t *testing.T) {
	f := newHistoryFixture(t)
	path := "changes/archive/0001-old.md"
	f.change(path, "applied", "")
	f.commit("archived baseline")
	f.put(path, "---\nmodule: alpha\nphase: applied\n---\n\nUpdated archive.\n")
	modified := f.commit("edit archive")
	f.remove(path)
	f.commit("delete archive")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got, historyFinding{"H04", path, modified})
}

func Test_CHK_076_HistoryMultipleSquashedChangesKeepH03H04(t *testing.T) {
	f := newHistoryFixture(t)
	archiveOld := "changes/archive/0000-old.md"
	f.change(archiveOld, "applied", "")
	archiveBaseline := f.commit("existing archive")
	for _, path := range []string{"changes/archive/0001-one.md", "changes/archive/0002-two.md"} {
		f.change(path, "applied", "")
	}
	f.put("locked/config.txt", "squash protection\n")
	f.put("packages/alpha/evals/existing_test.go", "squash eval\n")
	f.put("packages/alpha/ledger/entry.txt", "squash ledger\n")
	f.put(archiveOld, "---\nmodule: alpha\nphase: applied\n---\n\nChanged old archive.\n")
	squash := f.commit("squashed changes")
	got := f.report(archiveBaseline + "..HEAD")
	wantSquashed := map[checks.SquashedChange]bool{
		{Commit: squash, Change: "changes/archive/0001-one.md"}: true,
		{Commit: squash, Change: "changes/archive/0002-two.md"}: true,
	}
	if len(got.Squashed) != len(wantSquashed) {
		t.Errorf("squashed = %#v; want %#v", got.Squashed, wantSquashed)
	} else {
		for _, entry := range got.Squashed {
			if !wantSquashed[entry] {
				t.Errorf("unexpected squashed entry %#v; want %#v", entry, wantSquashed)
			}
			delete(wantSquashed, entry)
		}
		if len(wantSquashed) != 0 {
			t.Errorf("missing squashed entries %#v", wantSquashed)
		}
	}
	expectHistoryFindings(t, got,
		historyFinding{"H03", "packages/alpha/ledger/entry.txt", squash},
		historyFinding{"H04", archiveOld, squash})
}

func Test_CHK_077_HistoryUsesRecordedBeforeAfterZones(t *testing.T) {
	f := newHistoryFixture(t)
	pending := "changes/0001-work.md"
	f.change(pending, "impl", "")
	f.commit("ready")
	f.move(pending, "changes/archive/0001-work.md")
	f.change("changes/archive/0001-work.md", "applied", "")
	f.put("packages/alpha/evals/new_test.go", "new eval\n")
	applied := f.commit("apply")
	f.put("fireweed.yaml", strings.Replace(historyConfig, "evals: [evals/**]", "evals: [later/**]", 1))
	f.commit("later config")
	got := f.report(f.base + ".." + applied)
	expectHistoryFindings(t, got, historyFinding{"H02", "packages/alpha/evals/new_test.go", applied})
}

func Test_CHK_077_HistoryChecksBothRecordedZoneMaps(t *testing.T) {
	f := newHistoryFixture(t)
	f.put("fireweed.yaml", strings.Replace(historyConfig, "append_only: [ledger/**]", "append_only: [newledger/**]", 1))
	f.put("packages/alpha/ledger/entry.txt", "changed under before zone\n")
	f.put("packages/alpha/newledger/entry.txt", "changed under after zone\n")
	commit := f.commit("author changes zone and both existing files")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got,
		historyFinding{"H03", "packages/alpha/ledger/entry.txt", commit},
		historyFinding{"H03", "packages/alpha/newledger/entry.txt", commit})
}

func Test_CHK_078_HistoryRejectsUnreadableRangeAndRecordedDocuments(t *testing.T) {
	for _, tc := range []struct {
		name       string
		edit       func(historyFixture)
		rangeValue string
	}{
		{"invalid_range", nil, "missing-ref..HEAD"},
		{"missing_config", func(f historyFixture) { f.remove("fireweed.yaml") }, ""},
		{"malformed_config", func(f historyFixture) { f.put("fireweed.yaml", "modules: [broken\n") }, ""},
		{"missing_SPEC", func(f historyFixture) { f.remove("packages/alpha/SPEC.md") }, ""},
		{"malformed_change", func(f historyFixture) { f.put("changes/0001-work.md", "---\nmodule: [broken\n---\n") }, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := newHistoryFixture(t)
			if tc.edit != nil {
				tc.edit(f)
				f.commit("record invalid input")
			}
			revisionRange := tc.rangeValue
			if revisionRange == "" {
				revisionRange = f.base + "..HEAD"
			}
			if _, err := checks.History(f.root, revisionRange); err == nil {
				t.Errorf("History(%q) accepted %s", revisionRange, tc.name)
			}
		})
	}
}

 succeeded in 667ms:
total 8
drwx------@     5 joel  staff     160 Nov  6 07:54 .
drwx------@ 10623 joel  staff  339936 Nov  6 07:54 ..
drwxr-xr-x@     9 joel  staff     288 Nov  6 07:54 .git
-rw-r--r--@     1 joel  staff    3605 Nov  6 07:54 SYSTEM.md
drwxr-xr-x@     7 joel  staff     224 Nov  6 07:54 packages
./SYSTEM.md
./.git/config
./.git/HEAD
./.git/index
./.git/COMMIT_EDITMSG

exec
/bin/zsh -lc "sed -n '1,260p' packages/checks/evals/hooks_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc 'cat SYSTEM.md && rg --files packages | sort' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc "sed -n '1,300p' packages/checks/evals/zones_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc "sed -n '1,280p' packages/checks/evals/boundaries_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 689ms:
package evals_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func boundarySnapshot(t *testing.T, hook any) project.Snapshot {
	t.Helper()
	root := t.TempDir()
	snap := project.Snapshot{Root: root, Config: map[string]any{}, Modules: []project.Module{{
		Name: "alpha", Dir: "packages/alpha",
		Zones:    contracts.Zones{Disposable: []string{"internal/**"}},
		Document: contracts.Document{Frontmatter: map[string]any{}},
	}}}
	if hook != nil {
		snap.Config["hooks"] = map[string]any{"boundaries": hook}
	}
	return snap
}

func boundaryScript(t *testing.T, root, body string) []any {
	t.Helper()
	path := filepath.Join(root, "boundary-hook.sh")
	if err := os.WriteFile(path, []byte("#!/bin/sh\n"+body+"\n"), 0700); err != nil {
		t.Fatal(err)
	}
	return []any{path}
}

func boundaryHasError(diags []project.Diagnostic, err error) bool {
	if err != nil {
		return true
	}
	for _, d := range diags {
		if !d.Warning {
			return true
		}
	}
	return false
}

func boundaryHasWarning(diags []project.Diagnostic) bool {
	for _, d := range diags {
		if d.Warning {
			return true
		}
	}
	return false
}

func requireBoundaryDiagnostic(t *testing.T, diags []project.Diagnostic, code string, warning bool, message string) {
	t.Helper()
	for _, d := range diags {
		if message == "" {
			if d.Code == code && d.Warning == warning {
				return
			}
			continue
		}
		if strings.Contains(d.Message, message) {
			if d.Code != code || d.Warning != warning {
				t.Errorf("boundary diagnostic for %q = %#v; want code %s, warning %t", message, d, code, warning)
			}
			return
		}
	}
	t.Errorf("boundary diagnostic for %q absent: %#v", message, diags)
}

func Test_CHK_005_GlobalAndModuleBoundaryHooks(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	globalMarker := filepath.Join(snap.Root, "global-ran")
	moduleMarker := filepath.Join(snap.Root, "module-ran")
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, fmt.Sprintf("printf run > %q; printf 'alpha: global violation\\n'", globalMarker))}
	moduleHook := filepath.Join(snap.Root, "module-hook.sh")
	if err := os.WriteFile(moduleHook, []byte(fmt.Sprintf("#!/bin/sh\nprintf run > %q; printf 'beta: module violation\\n'\n", moduleMarker)), 0700); err != nil {
		t.Fatal(err)
	}
	snap.Modules = append(snap.Modules, project.Module{Name: "beta", Dir: "packages/beta", Zones: contracts.Zones{Disposable: []string{"internal/**"}}, Document: contracts.Document{Frontmatter: map[string]any{"hooks": map[string]any{"boundaries": []any{moduleHook}}}}})
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: 2 * time.Second})
	if !boundaryHasError(diags, err) {
		t.Errorf("unsuppressed violations accepted: %#v, %v", diags, err)
	}
	for _, module := range []string{"alpha", "beta"} {
		if !strings.Contains(fmt.Sprint(diags, err), module) {
			t.Errorf("violation for %s not attributed: %#v, %v", module, diags, err)
		}
	}
	for _, path := range []string{globalMarker, moduleMarker} {
		if _, statErr := os.Stat(path); statErr != nil {
			t.Errorf("applicable hook did not run: %s: %v", path, statErr)
		}
	}
}

func Test_CHK_041_ExactBaselineOnly(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	if err := os.MkdirAll(filepath.Join(snap.Root, ".fireweed"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(snap.Root, ".fireweed", "boundaries-baseline"), []byte("alpha: old violation\n"), 0644); err != nil {
		t.Fatal(err)
	}
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: old violation\n'`)}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) {
		t.Errorf("exact baseline line rejected: %#v, %v", diags, err)
	}
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: old violation changed\n'`)}
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("non-exact baseline line suppressed")
	}
}

func Test_CHK_039_CHK_040_OptOutSuppressesOnlyItsModuleAndWarns(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	snap.Modules[0].Document.Frontmatter["boundaries_opt_out"] = "legacy dependency"
	snap.Modules = append(snap.Modules, project.Module{Name: "beta", Dir: "packages/beta", Zones: contracts.Zones{Disposable: []string{"internal/**"}}, Document: contracts.Document{Frontmatter: map[string]any{}}})
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: ignored\n'`)}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) || !boundaryHasWarning(diags) {
		t.Errorf("reasoned opt-out = %#v, %v", diags, err)
	}
	if !strings.Contains(fmt.Sprint(diags), "legacy dependency") {
		t.Errorf("opt-out reason absent from warning: %#v", diags)
	}
	requireBoundaryDiagnostic(t, diags, "X03", true, "legacy dependency")
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: ignored\nbeta: enforced\n'`)}
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("opt-out hid another module's violation")
	}
	if err == nil {
		t.Error("unsuppressed beta violation returned no error")
	}
	requireBoundaryDiagnostic(t, diags, "X01", false, "beta: enforced")
}

func Test_CHK_044_CHK_045_MissingHookRequiredVersusOptional(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("missing hook for disposable module accepted")
	}
	if err == nil {
		t.Error("missing required hook returned no error")
	}
	requireBoundaryDiagnostic(t, diags, "X02", false, "alpha: missing boundaries hook")
	snap.Modules[0].Zones.Disposable = nil
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) || !boundaryHasWarning(diags) {
		t.Errorf("missing optional hook = %#v, %v", diags, err)
	}
	requireBoundaryDiagnostic(t, diags, "X04", true, "no boundary enforcement required")
	snap.Modules[0].Zones.Disposable = []string{"internal/**"}
	snap.Modules[0].Document.Frontmatter["boundaries_opt_out"] = "legacy"
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) || !boundaryHasWarning(diags) {
		t.Errorf("missing hook with reasoned opt-out = %#v, %v", diags, err)
	}
	requireBoundaryDiagnostic(t, diags, "X03", true, "legacy")
	requireBoundaryDiagnostic(t, diags, "X04", true, "no boundary enforcement required")
}

func Test_CHK_041_CHK_042_NormalNonzeroExitCanUseBaselineAndOptOut(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	if err := os.MkdirAll(filepath.Join(snap.Root, ".fireweed"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(snap.Root, ".fireweed", "boundaries-baseline"), []byte("alpha: known\n"), 0644); err != nil {
		t.Fatal(err)
	}
	snap.Modules[0].Document.Frontmatter["boundaries_opt_out"] = "legacy"
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: known\nalpha: opted out\n'; exit 3`)}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) {
		t.Errorf("normally exited hook with fully suppressed lines rejected: %#v, %v", diags, err)
	}
}

func Test_CHK_043_LaunchCancellationAndTimeoutStayErrors(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	snap.Modules[0].Document.Frontmatter["boundaries_opt_out"] = "legacy"
	snap.Config["hooks"] = map[string]any{"boundaries": []any{filepath.Join(snap.Root, "missing-executable")}}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("launch failure hidden by opt-out")
	}
	if err == nil {
		t.Error("launch failure returned no error")
	} else if !strings.Contains(err.Error(), "missing-executable") || strings.HasPrefix(err.Error(), "X02:") {
		t.Errorf("launch error = %q; want missing executable cause without X02 prefix", err)
	}
	requireBoundaryDiagnostic(t, diags, "X02", false, "")
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `sleep 2`)}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	diags, err = checks.Boundaries(ctx, checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("cancelled hook hidden by opt-out")
	}
	if err == nil {
		t.Error("cancelled hook returned no error")
	} else if !errors.Is(err, context.Canceled) || strings.HasPrefix(err.Error(), "X02:") {
		t.Errorf("cancellation error = %q; want canceled cause without X02 prefix", err)
	}
	requireBoundaryDiagnostic(t, diags, "X02", false, "")
	start := time.Now()
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: 80 * time.Millisecond})
	if !boundaryHasError(diags, err) {
		t.Error("timed-out hook hidden by opt-out")
	}
	if err == nil {
		t.Error("timed-out hook returned no error")
	} else if !errors.Is(err, context.DeadlineExceeded) || strings.HasPrefix(err.Error(), "X02:") {
		t.Errorf("timeout error = %q; want deadline cause without X02 prefix", err)
	}
	requireBoundaryDiagnostic(t, diags, "X02", false, "")
	if time.Since(start) > 3*time.Second {
		t.Error("boundary timeout did not bound execution")
	}
}

 succeeded in 693ms:
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

func Test_CHK_080_EvaluateRetainsParsedFailureDetails(t *testing.T) {
	for _, tc := range []struct {
		name, xml, exit, wantFailure string
		wantExit                     int
	}{
		{"nonzero failure", `<testsuite><testcase name="CHK-003_assertion"><failure message="want 2">got 3 at example_test.go:17</failure></testcase></testsuite>`, "7", "got 3 at example_test.go:17", 7},
		{"zero passing", `<testsuite><testcase name="CHK-003_pass"/></testsuite>`, "0", "", 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			snap := hookSnapshot(t, `printf '%s' '`+tc.xml+`' > {report}; exit `+tc.exit, nil)
			report, err := checks.Evaluate(context.Background(), checks.HookRequest{Snapshot: snap, Module: "alpha", Timeout: 2 * time.Second})
			if err != nil {
				t.Fatal(err)
			}
			if len(report.Cases) != 1 {
				t.Fatalf("cases = %#v; want one", report.Cases)
			}
			if report.HookExit != tc.wantExit {
				t.Errorf("hook exit = %d; want %s", report.HookExit, tc.exit)
			}
			if !strings.Contains(report.Cases[0].Failure, tc.wantFailure) {
				t.Errorf("failure = %q; missing %q", report.Cases[0].Failure, tc.wantFailure)
			}
			if tc.wantFailure == "" && report.Cases[0].Failure != "" {
				t.Errorf("passing case failure = %q; want empty", report.Cases[0].Failure)
			}
			if tc.wantFailure != "" && !strings.Contains(report.Cases[0].Failure, "want 2") {
				t.Errorf("failure = %q; missing message", report.Cases[0].Failure)
			}
		})
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

 succeeded in 693ms:
package evals_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

func zoneFixture(t *testing.T) (project.Snapshot, string) {
	t.Helper()
	root := t.TempDir()
	for path := range map[string]string{
		"packages/alpha/SPEC.md": "alpha", "packages/beta/SPEC.md": "beta",
		"packages/alpha/evals/old_test.go": "old", "packages/alpha/api.go": "api",
		"packages/alpha/internal/existing.go": "old", "packages/alpha/gen/old.sql": "old",
		"packages/alpha/keep.txt": "old", "changes/0011-current.md": "current",
	} {
		full := filepath.Join(root, path)
		if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte("fixture"), 0644); err != nil {
			t.Fatal(err)
		}
	}
	run := func(args ...string) string {
		cmd := exec.Command("git", args...)
		cmd.Dir = root
		cmd.Env = append(os.Environ(), "GIT_AUTHOR_NAME=Fixture", "GIT_AUTHOR_EMAIL=fixture@example.test", "GIT_COMMITTER_NAME=Fixture", "GIT_COMMITTER_EMAIL=fixture@example.test")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("git %v: %v: %s", args, err, out)
		}
		return string(out)
	}
	run("init", "-q")
	run("add", ".")
	run("commit", "-qm", "base")
	base := run("rev-parse", "HEAD")
	zones := contracts.Zones{Evals: []string{"evals/**"}, Interface: []string{"api.go"}, Disposable: []string{"internal/**"}, Preserve: []string{"keep.txt"}, AppendOnly: []string{"gen/**"}}
	snap := project.Snapshot{Root: root, Config: map[string]any{"protected": []any{"locked/**"}, "worker_writable": []any{"go.sum"}}, Modules: []project.Module{
		{Name: "alpha", Dir: "packages/alpha", Zones: zones},
		{Name: "beta", Dir: "packages/beta", Zones: zones},
	}}
	return snap, base[:len(base)-1]
}

func zoneRequest(snap project.Snapshot, base, phase string, changes ...gitstore.Change) checks.ZoneRequest {
	return checks.ZoneRequest{Snapshot: snap, Module: "alpha", Phase: phase, ChangePath: "changes/0011-current.md", Base: base, InterfaceDeclared: true, Changes: changes}
}

func zoneAllowed(t *testing.T, request checks.ZoneRequest) {
	t.Helper()
	if diags := checks.Zones(request); len(diags) != 0 {
		t.Errorf("wanted allowed change, got %#v", diags)
	}
}

func zoneRejected(t *testing.T, request checks.ZoneRequest) {
	t.Helper()
	if diags := checks.Zones(request); len(diags) == 0 {
		t.Errorf("forbidden change accepted: %#v", request.Changes)
	}
}

func Test_CHK_004_AuditOnlyCurrentChange(t *testing.T) {
	snap, base := zoneFixture(t)
	zoneAllowed(t, zoneRequest(snap, base, "audit", gitstore.Change{Path: "changes/0011-current.md", Status: "M"}))
	for _, path := range []string{"changes/0012-other.md", "packages/alpha/evals/old_test.go", "go.sum"} {
		zoneRejected(t, zoneRequest(snap, base, "audit", gitstore.Change{Path: path, Status: "M"}))
	}
}

func Test_CHK_006_CHK_007_CHK_008_CHK_009_RejectProtectedSpecForeignModuleAndChangeEveryPhase(t *testing.T) {
	snap, base := zoneFixture(t)
	for _, phase := range []string{"evals", "impl", "audit"} {
		for _, path := range []string{"locked/config", ".fireweed/hooks.go", "packages/alpha/SPEC.md", "packages/beta/internal/code.go", "changes/0012-other.md"} {
			zoneRejected(t, zoneRequest(snap, base, phase, gitstore.Change{Path: path, Status: "M"}))
		}
	}
}

func Test_CHK_011_CHK_012_CHK_013_CHK_014_CHK_015_CHK_021_EvalsAllowedZonesAndNewStubs(t *testing.T) {
	snap, base := zoneFixture(t)
	for _, change := range []gitstore.Change{
		{Path: "packages/alpha/evals/new_test.go", Status: "A"},
		{Path: "packages/alpha/api.go", Status: "M"},
		{Path: "packages/alpha/internal/new_stub.go", Status: "A"},
		{Path: "packages/alpha/manifest.json", Status: "M"},
		{Path: "changes/0011-current.md", Status: "M"},
	} {
		zoneAllowed(t, zoneRequest(snap, base, "evals", change))
	}
	zoneRejected(t, zoneRequest(snap, base, "evals", gitstore.Change{Path: "packages/alpha/internal/existing.go", Status: "M"}))
	zoneRejected(t, zoneRequest(snap, base, "evals", gitstore.Change{Path: "packages/alpha/keep.txt", Status: "M"}))
	request := zoneRequest(snap, base, "evals", gitstore.Change{Path: "packages/alpha/api.go", Status: "M"})
	request.InterfaceDeclared = false
	zoneRejected(t, request)
}

func Test_CHK_016_CHK_017_CHK_018_CHK_019_CHK_020_CHK_021_ImplFreezesEvalsAndInterfaces(t *testing.T) {
	snap, base := zoneFixture(t)
	for _, path := range []string{"packages/alpha/evals/old_test.go", "packages/alpha/api.go"} {
		zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: path, Status: "M"}))
	}
	for _, path := range []string{"packages/alpha/internal/existing.go", "packages/alpha/keep.txt", "packages/alpha/manifest.json", "changes/0011-current.md"} {
		zoneAllowed(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: path, Status: "M"}))
	}
}

func Test_CHK_022_CHK_023_ImplAppendOnlyNewVersusExisting(t *testing.T) {
	snap, base := zoneFixture(t)
	zoneAllowed(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: "packages/alpha/gen/new.sql", Status: "A"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: "packages/alpha/gen/old.sql", Status: "M"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: "packages/alpha/gen/old.sql", Status: "D"}))
	zoneRejected(t, zoneRequest(snap, base, "evals", gitstore.Change{Path: "packages/alpha/gen/new.sql", Status: "A"}))
}

func Test_CHK_024_RenamesCheckBothPaths(t *testing.T) {
	snap, base := zoneFixture(t)
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{OldPath: "packages/beta/internal/old.go", Path: "packages/alpha/internal/new.go", Status: "R100"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{OldPath: "packages/alpha/internal/existing.go", Path: "packages/beta/internal/new.go", Status: "R100"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{OldPath: "packages/alpha/gen/old.sql", Path: "packages/alpha/internal/new.go", Status: "R100"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{OldPath: "packages/alpha/internal/existing.go", Path: "packages/alpha/gen/new.sql", Status: "R100"}))
}

func Test_CHK_010_WorkerWritableOnlyEvalsAndImpl(t *testing.T) {
	snap, base := zoneFixture(t)
	for _, phase := range []string{"evals", "impl"} {
		zoneAllowed(t, zoneRequest(snap, base, phase, gitstore.Change{Path: "go.sum", Status: "M"}))
		zoneRejected(t, zoneRequest(snap, base, phase, gitstore.Change{Path: "outside.txt", Status: "M"}))
	}
	zoneRejected(t, zoneRequest(snap, base, "audit", gitstore.Change{Path: "go.sum", Status: "M"}))
}

 succeeded in 709ms:
# Checks reconstruction workspace

Reconstruct checks from its supplied SPEC.md, Decisions, retained interfaces and
evals. The module path is github.com/lurkshark/fireweed. Retain the thin api.go
facade and model types; put all reconstructed behavior under
packages/checks/internal/**. Never import the public checks parent from its own
internal implementation. Keep dependencies minimal; the existing YAML parser is
permitted. The product performs no runtime networking, fetch, push or telemetry.
Use ordinary Go source without cgo, native objects, assembly or embedded binaries.

Project conventions remain in force: read commands run no hooks or agents,
change no files, index or refs, and produce deterministic output. SPEC.md
frontmatter is the sole module graph. Evals use public interfaces, local fixtures
and bounded fake agents, without runtime network calls or agent authentication.

## Input isolation

Read only this input workspace and files you create here. Do not read external
trees, old implementation, old history, saved attempts, other-module internals,
prior reports, product binaries or bootstrap. Do not query the invoking controller
as an implementation oracle, including its help, recipes, playbooks or behavior.
Do not invoke other agents. Do not commit or run lifecycle/authoring mutators.

Retained files must remain byte-unchanged. Write only target disposable files
under packages/checks/internal/** and DRILL_NOTES.md. Do not edit specs, Decisions,
evals, interfaces, or dependency files. Do not create root manifests or supply
replacement dependency implementations. Remove temporary probe files before
returning. Preserve every existing assertion and behavior.

The only supplied dependency material is the declared contracts, gitstore,
assets and project public interfaces, including their model interfaces. Facades
may import deliberately absent internal packages. Do not search for or replace
those internals. Full manifests, dependency and dependent implementations, hooks
and the invoking executable are outside reconstruction inputs in the separate
native verifier. No old target code, source history, remotes or alternates are
available in input Git.

## Verification and notes

The outside owning controller runs the native verifier against the reconstructed
checks plus real dependencies, and enforces checks and all transitive dependent
evals and coverage (workflow, runner and app). Retained inputs must stay unchanged.
Do not claim those full checks passed from this input-only workspace. A local
build limitation caused by the deliberate absence of full manifests/dependency
implementations is a verification limitation, not permission to obtain them.
Use gofmt and bounded ordinary local Go/Git probes when useful; do not inspect
tool caches as source or change enforcement to make probes pass.

Write DRILL_NOTES.md with separate Contradicted Decisions and Gaps sections.
Every behavioral gap names the nearest spec ID (or no spec) and target module.
State accurately when there are no contradicted Decisions or behavioral gaps.
Do not author corrections from gap suggestions or repair dependent code.

Keep this input tree's .fireweed/bin first on worker PATH and use
GOCACHE=/tmp/fireweed-rebuild-cache. No binary is supplied as an input. The outside
controller selects exactly codex exec --ephemeral --ignore-user-config -m
gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt}; do not
substitute the model or effort. Return work to that controller; it owns native
bounded retries, verification, report/ref retention and scratch cleanup.
packages/assets/api.go
packages/assets/model/model.go
packages/checks/SPEC.md
packages/checks/api.go
packages/checks/evals/boundaries_test.go
packages/checks/evals/doctor_agent_logs_test.go
packages/checks/evals/doctor_prompt_file_test.go
packages/checks/evals/doctor_test.go
packages/checks/evals/history_test.go
packages/checks/evals/hooks_test.go
packages/checks/evals/junit_test.go
packages/checks/evals/zones_test.go
packages/checks/model/model.go
packages/contracts/api.go
packages/contracts/model/model.go
packages/gitstore/api.go
packages/gitstore/model/model.go
packages/project/api.go
packages/project/model/model.go

exec
/bin/zsh -lc 'cat packages/project/api.go packages/project/model/model.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc 'cat packages/assets/api.go packages/assets/model/model.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc 'cat packages/gitstore/api.go packages/gitstore/model/model.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc "sed -n '1,280p' packages/checks/evals/doctor_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc 'cat packages/contracts/api.go packages/contracts/model/model.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 784ms:
package project

import (
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/internal/history"
	"github.com/lurkshark/fireweed/packages/project/internal/queue"
	"github.com/lurkshark/fireweed/packages/project/model"
)

type Snapshot = model.Snapshot
type Module = model.Module
type Change = model.Change
type SpecHistoryEntry = model.SpecHistoryEntry
type StaleSpec = model.StaleSpec
type Diagnostic = model.Diagnostic

func Load(root string) (Snapshot, error) { return fixture.Load(root) }

func LoadAt(root, revision string) (Snapshot, error) { return fixture.LoadAt(root, revision) }

func ValidateGraph(modules []Module) error { return fixture.ValidateGraph(modules) }

func NewChange(root, module, slug string) (string, error) {
	return fixture.NewChange(root, module, slug)
}

func Queue(root string) ([]Change, error) { return queue.Queue(root) }

func Project(root, through string) (Snapshot, error) { return queue.Project(root, through) }

func Stamp(root, change string) error { return queue.Stamp(root, change) }

func StampMissing(root, change string) error { return queue.StampMissing(root, change) }

func StaleSpecs(root, change string) ([]StaleSpec, error) { return queue.StaleSpecs(root, change) }

func Check(root string) []Diagnostic { return queue.Check(root) }

func SpecHistory(root, module, id string) ([]SpecHistoryEntry, error) {
	return history.SpecHistory(root, module, id)
}
package model

import (
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
)

// Snapshot is the project data read from a repository root.
type Snapshot struct {
	Root    string
	Config  map[string]any
	Modules []Module
}

// Module is a discovered module and its declared document and zones.
type Module struct {
	Name      string
	IDPrefix  string
	Dir       string
	DependsOn []string
	Zones     contracts.Zones
	Document  contracts.Document
}

// Change is one pending change document and its spec edits.
type Change struct {
	Number   int
	Slug     string
	Path     string
	Module   string
	Document contracts.Document
	Added    []contracts.Item
	Modified []contracts.Item
	Removed  []contracts.Item
}

// SpecHistoryEntry describes one archived change affecting selected specs.
type SpecHistoryEntry struct {
	Path               string
	Module             string
	IDs                []string
	Motivation         string
	Decisions          string
	RemovedDecisions   string
	Provenance         string
	ImplementationDate time.Time
}

// StaleSpec describes a recorded base that differs from the preceding projection.
type StaleSpec struct {
	ID           string
	RecordedHash string
	Before       string
	Current      string
	Recovered    bool
}

// Diagnostic reports a static authoring check.
type Diagnostic struct {
	Code    string
	Message string
	Path    string
	Warning bool
}

 succeeded in 782ms:
package contracts

import (
	"github.com/lurkshark/fireweed/packages/contracts/internal/primitives"
	"github.com/lurkshark/fireweed/packages/contracts/model"
)

type Document = model.Document
type Item = model.Item
type Zones = model.Zones
type CodedError = model.CodedError

// ErrorCode returns the first coded diagnostic in an error chain.
func ErrorCode(err error) (code, subject string, ok bool) { return primitives.ErrorCode(err) }

func ParseDocument(source string) (Document, error) { return primitives.ParseDocument(source) }
func RenderDocument(doc Document) (string, error)   { return primitives.RenderDocument(doc) }

// Section returns the section body without its level-two heading.
func Section(markdown, name string) (string, bool) { return primitives.Section(markdown, name) }

// ReplaceSection replaces the body of a level-two section, or appends it.
func ReplaceSection(markdown, name, body string) string {
	return primitives.ReplaceSection(markdown, name, body)
}

func ParseItems(section string) ([]Item, error) { return primitives.ParseItems(section) }
func RenderItem(item Item) string               { return primitives.RenderItem(item) }
func SpecHash(item Item) string                 { return primitives.SpecHash(item) }

func MatchesID(id, testcaseName, classname string) bool {
	return primitives.MatchesID(id, testcaseName, classname)
}

func MatchGlob(pattern, path string) bool { return primitives.MatchGlob(pattern, path) }

// ResolveZone classifies a module-relative path.
func ResolveZone(path string, zones Zones) string { return primitives.ResolveZone(path, zones) }
package model

// Document is a Markdown document with YAML frontmatter.
type Document struct {
	Frontmatter map[string]any
	Body        string
}

// Item is one level-three spec item. Body excludes the heading line.
type Item struct {
	ID   string
	Tag  string
	Body string
}

// CodedError carries a stable diagnostic code while preserving the cause's text.
type CodedError struct {
	Code    string
	Subject string
	Cause   error
}

func (e CodedError) Error() string {
	if e.Cause != nil {
		return e.Cause.Error()
	}
	return e.Code
}

func (e CodedError) Unwrap() error { return e.Cause }

// Zones holds module-relative path patterns for each editable zone.
type Zones struct {
	Protected  []string
	Evals      []string
	Interface  []string
	AppendOnly []string
	Preserve   []string
	Disposable []string
}

 succeeded in 783ms:
package evals_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"gopkg.in/yaml.v3"
)

const doctorGuidance = "<!-- fireweed:start -->\nThis project uses fireweed. Specs and change files are the source of truth.\nRun `fireweed playbook author-change` before authoring changes and\n`fireweed context <change>` before each step.\n<!-- fireweed:end -->\n"

type doctorFixture struct {
	t    *testing.T
	root string
	log  string
	bin  string
}

func (f doctorFixture) put(name, body string) {
	f.t.Helper()
	path := filepath.Join(f.root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		f.t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		f.t.Fatal(err)
	}
}

func (f doctorFixture) git(args ...string) string {
	f.t.Helper()
	cmd := exec.Command("git", args...)
	cmd.Dir = f.root
	cmd.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_TERMINAL_PROMPT=0")
	out, err := cmd.CombinedOutput()
	if err != nil {
		f.t.Fatalf("git %v: %v: %s", args, err, out)
	}
	return string(out)
}

func (f doctorFixture) commit() {
	f.t.Helper()
	f.git("init", "-b", "main")
	f.git("config", "user.name", "Doctor Fixture")
	f.git("config", "user.email", "doctor@example.test")
	f.git("add", ".")
	f.git("commit", "-m", "fixture")
}

func (f doctorFixture) request(noAgent bool) checks.DoctorRequest {
	return checks.DoctorRequest{Root: f.root, Bin: f.bin, NoAgent: noAgent, Timeout: 15 * time.Second}
}

func doctorConfig(t *testing.T, recipe assets.Recipe) string {
	t.Helper()
	config := map[string]any{
		"fireweed": ">=0.1, <0.2", "version": 1, "base": "main", "modules": []string{"packages/*"},
		"limit_patterns":  []string{"rate limit"},
		"module_defaults": map[string]any{"interface": recipe.ModuleDefaults.Interface, "evals": recipe.ModuleDefaults.Evals, "disposable": recipe.ModuleDefaults.Disposable},
		"hooks":           recipe.Hooks,
		"agent":           map[string]any{"command": []string{"bash", ".fireweed/agent.sh", "{prompt}", "{git_common_dir}"}, "wrap": []string{"bash", ".fireweed/wrap.sh"}},
		"audit":           map[string]any{"enabled": true, "command": []string{"bash", ".fireweed/audit.sh", "{prompt}"}},
	}
	data, err := yaml.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func (f doctorFixture) editConfig(change func(map[string]any)) {
	f.t.Helper()
	var config map[string]any
	if err := yaml.Unmarshal(mustRead(f.t, filepath.Join(f.root, "fireweed.yaml")), &config); err != nil {
		f.t.Fatal(err)
	}
	change(config)
	data, err := yaml.Marshal(config)
	if err != nil {
		f.t.Fatal(err)
	}
	f.put("fireweed.yaml", string(data))
}

func newDoctorFixture(t *testing.T) doctorFixture {
	t.Helper()
	r, err := assets.BuiltinRecipe("go")
	if err != nil {
		t.Fatal(err)
	}
	f := doctorFixture{t: t, root: t.TempDir(), log: filepath.Join(t.TempDir(), "probes.log")}
	f.put("go.mod", "module example.test/doctorfixture\n\ngo 1.23\n")
	f.put("fireweed.yaml", doctorConfig(t, r))
	f.put("AGENTS.md", "Before.\n"+doctorGuidance+"After.\n")
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nThe fixture MUST pass.\n")
	f.put("packages/alpha/api.go", "package alpha\n")
	for path, body := range r.Files {
		f.put(path, body)
	}
	f.put(".fireweed/wrap.sh", fmt.Sprintf("#!/bin/bash\nprintf 'wrap:%%s\n' \"$PWD\" >> %q\nexec \"$@\"\n", f.log))
	f.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/bash\nprintf 'worker:%%s:%%s:%%s\n' \"$PWD\" \"$1\" \"$2\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	f.put(".fireweed/audit.sh", fmt.Sprintf("#!/bin/bash\nprintf 'audit:%%s:%%s\n' \"$PWD\" \"$1\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	return f
}

func localDoctorFixture(t *testing.T) doctorFixture {
	t.Helper()
	f := doctorFixture{t: t, root: t.TempDir(), log: filepath.Join(t.TempDir(), "probes.log")}
	f.bin = filepath.Join(t.TempDir(), "fake-fireweed")
	if err := os.WriteFile(f.bin, []byte(`#!/bin/sh
if test "$1" != modules || test "$2" != --json; then exit 2; fi
printf '{"data":['
first=1
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test "$first" = 0; then printf ','; fi; first=0
  printf '{"name":"%s","dir":"%s","depends_on":[],"zones":{"interface":["api.txt"],"evals":["evals/**"],"disposable":["internal/**"]}}' "$name" "$dir"
done
printf ']}'
`), 0700); err != nil {
		t.Fatal(err)
	}
	recipe := assets.Recipe{ModuleDefaults: structZones(), Hooks: map[string]any{
		"eval":       "sh .fireweed/eval.sh {module} {module_dir} {report}",
		"boundaries": "sh .fireweed/boundary.sh",
		"setup":      "sh .fireweed/setup.sh",
	}}
	f.put("fireweed.yaml", doctorConfig(t, recipe))
	f.put(".gitignore", ".fireweed/runs/\n")
	f.put("AGENTS.md", doctorGuidance)
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nThe fixture MUST pass.\n")
	f.put("packages/alpha/api.txt", "public\n")
	f.put(".fireweed/templates/module/SPEC.md", "---\nname: {{name}}\nid_prefix: {{prefix}}\ndepends_on: []\n---\n# {{name}}\n\n## Specs\n\n### {{prefix}}-001\nThe canary MUST pass.\n")
	f.put(".fireweed/templates/module/api.txt.tmpl", "{{package}} {{go_module}} {{module_dir}} {{other_dir}}\n")
	f.put(".fireweed/canaries/canaries.yaml", `version: 1
baseline:
  SPEC.md: "---\nname: {{name}}\nid_prefix: {{prefix}}\ndepends_on: []\n---\n# {{name}}\n\n## Specs\n\n### {{prefix}}-001\nThe canary MUST pass.\n"
  api.txt: "{{package}} {{go_module}} {{module_dir}} {{other_dir}}"
undeclared:
  undeclared.txt: "{{other_dir}}"
internal_eval:
  internal/private.txt: "private"
  evals/internal.txt: "{{module_dir}}"
failing_eval:
  evals/failing.txt: "{{prefix}}-001"
missing_id:
  evals/missing.txt: "UNMAPPED-999"
`)
	f.put(".fireweed/setup.sh", fmt.Sprintf("#!/bin/sh\nprintf 'setup:%%s\n' \"$PWD\" >> %q\n", f.log))
	f.put(".fireweed/eval.sh", fmt.Sprintf(`#!/bin/sh
printf 'eval:%%s:%%s:%%s\n' "$PWD" "$1" "$FIREWEED_BIN" >> %q
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
if test -f "$2/evals/failing.txt"; then printf '<testsuite><testcase name="%%s-001"><failure/></testcase></testsuite>' "$prefix" > "$3"; exit 1; fi
if test -f "$2/evals/missing.txt"; then printf '<testsuite><testcase name="UNMAPPED-999"/></testsuite>' > "$3"; exit 0; fi
printf '<testsuite><testcase name="%%s-001"/></testsuite>' "$prefix" > "$3"
`, f.log))
	f.put(".fireweed/boundary.sh", fmt.Sprintf(`#!/bin/sh
printf 'boundaries:%%s:%%s\n' "$PWD" "$FIREWEED_BIN" >> %q
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test -f "$dir/undeclared.txt"; then printf '%%s: undeclared import\n' "$name"; fi
  if test -f "$dir/evals/internal.txt"; then printf '%%s: interface violation\n' "$name"; fi
done
`, f.log))
	f.put(".fireweed/wrap.sh", fmt.Sprintf("#!/bin/bash\nprintf 'wrap:%%s\n' \"$PWD\" >> %q\nexec \"$@\"\n", f.log))
	f.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/bash\nprintf 'worker:%%s:%%s:%%s\n' \"$PWD\" \"$1\" \"$2\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	f.put(".fireweed/audit.sh", fmt.Sprintf("#!/bin/bash\nprintf 'audit:%%s:%%s\n' \"$PWD\" \"$1\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	return f
}

func structZones() (z contracts.Zones) {
	z.Interface = []string{"api.txt"}
	z.Evals = []string{"evals/**"}
	z.Disposable = []string{"internal/**"}
	return
}

func doctorText(report checks.DoctorReport) string {
	return fmt.Sprintf("checks=%+v diagnostics=%+v", report.Checks, report.Diagnostics)
}
func hasDoctorDiagnostic(report checks.DoctorReport, code string, warning bool, terms ...string) bool {
	for _, d := range report.Diagnostics {
		if d.Code != code || d.Warning != warning {
			continue
		}
		ok := true
		for _, s := range terms {
			if !strings.Contains(strings.ToLower(d.Message), strings.ToLower(s)) {
				ok = false
			}
		}
		if ok {
			return true
		}
	}
	return false
}
func hasDoctorCheck(report checks.DoctorReport, name string, passed bool) bool {
	for _, c := range report.Checks {
		if strings.Contains(c.Name, name) && c.Passed == passed {
			return true
		}
	}
	return false
}
func doctorErrors(report checks.DoctorReport) bool {
	for _, d := range report.Diagnostics {
		if !d.Warning {
			return true
		}
	}
	return false
}
func doctorLog(f doctorFixture) string { b, _ := os.ReadFile(f.log); return string(b) }
func worktrees(f doctorFixture) string { return f.git("worktree", "list", "--porcelain") }

func logDoctorCanaryProbes(t *testing.T, f doctorFixture) {
	t.Helper()
	logProbe := fmt.Sprintf(`printf 'probe:%%s\n' "$PWD" >> %q
for marker in undeclared.txt internal.txt failing.txt missing.txt; do
  if find packages -name "$marker" | grep -q .; then printf 'planted:%%s\n' "$marker" >> %q; fi
done
`, f.log, f.log)
	for _, path := range []string{".fireweed/eval.sh", ".fireweed/boundary.sh"} {
		body := string(mustRead(t, filepath.Join(f.root, path)))
		if !strings.HasPrefix(body, "#!/bin/sh\n") {
			t.Fatalf("fixture hook %s has no shell header", path)
		}
		f.put(path, strings.Replace(body, "#!/bin/sh\n", "#!/bin/sh\n"+logProbe, 1))
	}
}

func assertScratchRemoved(t *testing.T, f doctorFixture) {
	t.Helper()
	common := strings.TrimSpace(f.git("rev-parse", "--git-common-dir"))
	if !filepath.IsAbs(common) {
		common = filepath.Join(f.root, common)
	}
	common, err := filepath.EvalSymlinks(common)
	if err != nil {
		t.Fatal(err)
	}
	base := filepath.Join(common, "fireweed", "worktrees")
	for _, line := range strings.Split(doctorLog(f), "\n") {
		if !strings.HasPrefix(line, "setup:") {
			continue
		}
		path := strings.TrimPrefix(line, "setup:")
		for parent := path; ; parent = filepath.Dir(parent) {
			resolved, err := filepath.EvalSymlinks(parent)
			if err == nil {
				suffix, err := filepath.Rel(parent, path)
				if err != nil {
					t.Fatal(err)
				}
				path = filepath.Join(resolved, suffix)
				break
			}
			if parent == filepath.Dir(parent) {
				t.Fatal(err)
			}
		}
		rel, err := filepath.Rel(base, path)
		if err != nil || rel == "." || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
			t.Errorf("scratch %q is not below common Git directory %q", path, base)
		}

 succeeded in 799ms:
package assets

import "github.com/lurkshark/fireweed/packages/assets/internal/content"
import "github.com/lurkshark/fireweed/packages/assets/internal/recipe"
import "github.com/lurkshark/fireweed/packages/assets/model"

type Recipe = model.Recipe

func Version() string { return content.Version() }

func Help() string { return content.Help() }

func Playbook(name string) (string, error) { return content.Playbook(name) }

func Explain(code string) (string, error) { return content.Explain(code) }

func BuiltinRecipe(name string) (Recipe, error) { return recipe.BuiltinRecipe(name) }
package model

import "github.com/lurkshark/fireweed/packages/contracts"

// Recipe is the data copied into a project during setup.
type Recipe struct {
	Name           string
	Instructions   string
	ModuleDefaults contracts.Zones
	Hooks          map[string]any
	Protected      []string
	WorkerWritable []string
	Files          map[string]string
}

 succeeded in 809ms:
package gitstore

import (
	"context"

	"github.com/lurkshark/fireweed/packages/gitstore/internal/local"
	"github.com/lurkshark/fireweed/packages/gitstore/model"
)

type Repository = model.Repository
type Change = model.Change
type Request = model.Request
type Result = model.Result

func Inspect(dir string) (Repository, error)             { return local.Inspect(dir) }
func BaseRef(root, remote, base string) (string, error)  { return local.BaseRef(root, remote, base) }
func Changes(root, base string) ([]Change, error)        { return local.Changes(root, base) }
func ReadAt(root, revision, path string) ([]byte, error) { return local.ReadAt(root, revision, path) }
func Commit(root, subject string, paths []string) (string, error) {
	return local.Commit(root, subject, paths)
}
func WriteBlob(root string, data []byte) (string, error)       { return local.WriteBlob(root, data) }
func ReadBlob(root, oid string) ([]byte, error)                { return local.ReadBlob(root, oid) }
func Ref(root, name string) (string, bool, error)              { return local.Ref(root, name) }
func CASRef(root, name, old, new string) error                 { return local.CASRef(root, name, old, new) }
func Run(ctx context.Context, request Request) (Result, error) { return local.Run(ctx, request) }
package model

import "time"

type Repository struct {
	Root      string
	CommonDir string
	Head      string
	Branch    string
}

type Change struct {
	Path    string
	OldPath string
	Status  string
}

type Request struct {
	Args    []string
	Dir     string
	Env     map[string]string
	Input   []byte
	Timeout time.Duration
}

type Result struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

codex
The workspace contains the retained facades and evals, but no dependency implementations or Go manifest, as the supplied instructions describe. I’m building against the public interfaces and using the eval fixtures to guide behavior.
exec
/bin/zsh -lc "sed -n '280,700p' packages/checks/evals/doctor_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc 'cat packages/checks/evals/doctor_agent_logs_test.go packages/checks/evals/doctor_prompt_file_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc "sed -n '1,290p' packages/checks/evals/junit_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 703ms:
package evals_test

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
)

func requireDiagnosticError(t *testing.T, err error, code, cause string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected %s error containing %q", code, cause)
	}
	if !strings.HasPrefix(err.Error(), code+": ") || !strings.Contains(err.Error(), cause) || strings.HasPrefix(strings.TrimPrefix(err.Error(), code+": "), code+":") {
		t.Errorf("error = %q; want %s prefix and cause %q", err, code, cause)
	}
}

func checkSpecs() []contracts.Item {
	return []contracts.Item{{ID: "CHK-001"}, {ID: "CHK-002"}, {ID: "CHK-003"}}
}

func Test_CHK_001_CHK_029_CHK_030_NestedJUnitIDsResultsAndTime(t *testing.T) {
	data := []byte(`<testsuites><testsuite name="outer"><testsuite name="inner">
<testcase classname="suite.ChK_001" name="passes" time="1.25"/>
<testcase classname="suite" name="test_chk002_failed" time="0.5"><failure message="bad"/></testcase>
<testcase classname="suite" name="test_chk_003_error" time="0.75"><error message="broken"/></testcase>
<testcase classname="CHK-002" name="CHK-001_both" time="0.125"/>
</testsuite></testsuite></testsuites>`)
	report, err := checks.ParseJUnit(data, checkSpecs())
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]checks.Case{
		"passes":             {Name: "passes", Classname: "suite.ChK_001", IDs: []string{"CHK-001"}, Result: "pass", Seconds: 1.25},
		"test_chk002_failed": {Name: "test_chk002_failed", Classname: "suite", IDs: []string{"CHK-002"}, Result: "fail", Seconds: 0.5},
		"test_chk_003_error": {Name: "test_chk_003_error", Classname: "suite", IDs: []string{"CHK-003"}, Result: "fail", Seconds: 0.75},
		"CHK-001_both":       {Name: "CHK-001_both", Classname: "CHK-002", IDs: []string{"CHK-001", "CHK-002"}, Result: "pass", Seconds: 0.125},
	}
	if len(report.Cases) != len(want) {
		t.Fatalf("cases = %#v; want %d", report.Cases, len(want))
	}
	for _, got := range report.Cases {
		sort.Strings(got.IDs)
		got.Failure = ""
		if !reflect.DeepEqual(got, want[got.Name]) {
			t.Errorf("case = %#v; want %#v", got, want[got.Name])
		}
	}
}

func Test_CHK_001_CHK_029_CHK_030_SingleSuiteSkipped(t *testing.T) {
	report, err := checks.ParseJUnit([]byte(`<testsuite><testcase name="CHK-001" classname="unit" time="0.125"><skipped/></testcase></testsuite>`), checkSpecs())
	if err != nil {
		t.Fatal(err)
	}
	if len(report.Cases) != 1 || report.Cases[0].Result != "skip" || report.Cases[0].Seconds != 0.125 || !reflect.DeepEqual(report.Cases[0].IDs, []string{"CHK-001"}) {
		t.Fatalf("skipped case = %#v", report.Cases)
	}
}

func Test_CHK_079_ParseJUnitRetainsFailureAndErrorDetails(t *testing.T) {
	data := []byte(`<testsuite>
<testcase name="CHK-001_failure_combined"><failure message="want 2 &amp; got 3">assertion at example_test.go:17</failure></testcase>
<testcase name="CHK-001_failure_message"><failure message="expected true"/></testcase>
<testcase name="CHK-002_failure_body"><failure>actual value was false</failure></testcase>
<testcase name="CHK-002_error_combined"><error message="panic: nil pointer">stack at example_test.go:29</error></testcase>
<testcase name="CHK-003_error_message"><error message="setup failed"/></testcase>
<testcase name="CHK-003_error_body"><error>missing fixture</error></testcase>
<testcase name="CHK-001_empty_failure"><failure/></testcase>
<testcase name="CHK-002_pass"/>
<testcase name="CHK-003_skip"><skipped/></testcase>
</testsuite>`)
	report, err := checks.ParseJUnit(data, checkSpecs())
	if err != nil {
		t.Fatal(err)
	}
	want := map[string][]string{
		"CHK-001_failure_combined": {"want 2 & got 3", "assertion at example_test.go:17"},
		"CHK-001_failure_message":  {"expected true"},
		"CHK-002_failure_body":     {"actual value was false"},
		"CHK-002_error_combined":   {"panic: nil pointer", "stack at example_test.go:29"},
		"CHK-003_error_message":    {"setup failed"},
		"CHK-003_error_body":       {"missing fixture"},
		"CHK-001_empty_failure":    nil,
		"CHK-002_pass":             nil,
		"CHK-003_skip":             nil,
	}
	if len(report.Cases) != len(want) {
		t.Fatalf("cases = %#v; want %d", report.Cases, len(want))
	}
	for _, got := range report.Cases {
		parts, ok := want[got.Name]
		if !ok {
			t.Errorf("unexpected case %#v", got)
			continue
		}
		if len(parts) == 0 && got.Failure != "" {
			t.Errorf("%s: failure = %q; want empty", got.Name, got.Failure)
		}
		for _, part := range parts {
			if !strings.Contains(got.Failure, part) {
				t.Errorf("%s: failure = %q; missing %q", got.Name, got.Failure, part)
			}
		}
	}
}

func Test_CHK_002_SkippedJUnitCaseDoesNotCover(t *testing.T) {
	report, err := checks.ParseJUnit([]byte(`<testsuite><testcase name="CHK-001" classname="unit" time="0.125"><skipped/></testcase></testsuite>`), checkSpecs())
	if err != nil {
		t.Fatal(err)
	}
	requireDiagnosticError(t, checks.RequireCoverage(report, []contracts.Item{{ID: "CHK-001"}}), "C10", "no coverage for CHK-001")
}

func Test_CHK_027_RejectMalformedXML(t *testing.T) {
	for _, tc := range []struct{ xml, cause string }{
		{`<testsuite><testcase name="CHK-001"></testsuite>`, "parse JUnit:"},
		{`<document/>`, `invalid JUnit root "document"`},
	} {
		_, err := checks.ParseJUnit([]byte(tc.xml), checkSpecs())
		requireDiagnosticError(t, err, "C11", tc.cause)
	}
}

func Test_CHK_028_RejectMissingAndUnknownIDs(t *testing.T) {
	for _, xml := range []string{
		`<testsuite><testcase name="ordinary" classname="unit"/></testsuite>`,
		`<testsuite><testcase name="ordinary" classname="unit"><skipped/></testcase></testsuite>`,
		`<testsuite><testcase name="CHK-999" classname="unit"/></testsuite>`,
		`<testsuite><testcase name="OCHK001" classname="unit"/></testsuite>`,
		`<testsuite><testcase name="CHK0010" classname="unit"/></testsuite>`,
	} {
		_, err := checks.ParseJUnit([]byte(xml), checkSpecs())
		cause := "has no known spec ID"
		if strings.Contains(xml, "CHK-999") {
			cause = `references unknown spec ID "CHK-999"`
		} else if strings.Contains(xml, "CHK0010") {
			cause = `references unknown spec ID "CHK0010"`
		}
		requireDiagnosticError(t, err, "C11", cause)
	}
}

func Test_CHK_002_CoverageRequiresCurrentNonSkippedCases(t *testing.T) {
	specs := []contracts.Item{{ID: "CHK-001"}, {ID: "CHK-002"}}
	requireDiagnosticError(t, checks.RequireCoverage(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001"}, Result: "pass"}}}, specs), "C10", "no coverage for CHK-002")
	if err := checks.RequireCoverage(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001", "CHK-002"}, Result: "fail"}}}, specs); err != nil {
		t.Errorf("failed referencing case should cover both IDs: %v", err)
	}
}

func Test_CHK_002_CoverageExemptsReviewAndQuarantine(t *testing.T) {
	specs := []contracts.Item{{ID: "CHK-001"}, {ID: "CHK-002", Tag: "review"}, {ID: "CHK-003", Tag: "quarantine: flaky"}}
	if err := checks.RequireCoverage(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001"}, Result: "pass"}}}, specs); err != nil {
		t.Errorf("review/quarantine required coverage: %v", err)
	}
}

func Test_CHK_025_FailFirstRequiresEachRequestedFailure(t *testing.T) {
	ids := []string{"CHK-001", "CHK-002"}
	bad := []checks.Case{
		{IDs: []string{"CHK-001"}, Result: "pass"},
		{IDs: []string{"CHK-002"}, Result: "skip"},
		{IDs: []string{"CHK-003"}, Result: "fail"},
	}
	requireDiagnosticError(t, checks.RequireFailFirst(checks.Report{Cases: bad}, ids, nil), "P2-01", "no fail-first case for CHK-001")
	requireDiagnosticError(t, checks.RequireFailFirst(checks.Report{}, ids, nil), "P2-01", "no fail-first case for CHK-001")
	requireDiagnosticError(t, checks.RequireFailFirst(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001"}, Result: "fail"}, {IDs: []string{"CHK-002"}, Result: "skip"}}}, ids, nil), "P2-01", "no fail-first case for CHK-002")
	good := checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001"}, Result: "fail"}, {IDs: []string{"CHK-002"}, Result: "fail"}}}
	if err := checks.RequireFailFirst(good, ids, nil); err != nil {
		t.Errorf("both requested IDs failed: %v", err)
	}
}

func Test_CHK_025_FailFirstAlreadySatisfiedExemption(t *testing.T) {
	report := checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-002"}, Result: "fail"}}}
	if err := checks.RequireFailFirst(report, []string{"CHK-001", "CHK-002"}, []string{"CHK-001"}); err != nil {
		t.Errorf("already_satisfied ID incorrectly required: %v", err)
	}
}

func Test_CHK_026_PassRejectsNonQuarantinedFailures(t *testing.T) {
	specs := []contracts.Item{{ID: "CHK-001"}, {ID: "CHK-002", Tag: "review"}, {ID: "CHK-003", Tag: "quarantine: flaky"}}
	for _, id := range []string{"CHK-001", "CHK-002"} {
		requireDiagnosticError(t, checks.RequirePass(checks.Report{Cases: []checks.Case{{IDs: []string{id}, Result: "fail"}}}, specs), "E01", `failed case "" for `+id)
	}
	if err := checks.RequirePass(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-003"}, Result: "fail"}}}, specs); err != nil {
		t.Errorf("quarantined failure rejected: %v", err)
	}
}

 succeeded in 703ms:
package evals_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
)

func doctorProbeLog(t *testing.T, f doctorFixture, report checks.DoctorReport, role string, markers ...string) string {
	t.Helper()
	var detail string
	for _, check := range report.Checks {
		if check.Name == role {
			detail = check.Detail
			break
		}
	}
	if detail == "" {
		t.Errorf("executed %s probe has no check: %s", role, doctorText(report))
		return ""
	}
	at := strings.Index(detail, "log: ")
	if at < 0 {
		t.Errorf("%s detail lacks log path: %q", role, detail)
		return ""
	}
	fields := strings.Fields(detail[at+len("log: "):])
	if len(fields) == 0 {
		t.Errorf("%s detail has empty log path: %q", role, detail)
		return ""
	}
	path := fields[0]
	if !filepath.IsAbs(path) {
		t.Errorf("%s log path is not absolute: %q", role, path)
	}
	rel, err := filepath.Rel(filepath.Join(f.root, ".fireweed", "runs"), path)
	if err != nil {
		t.Errorf("%s log path %q: %v", role, path, err)
		return ""
	}
	parts := strings.Split(filepath.ToSlash(rel), "/")
	if len(parts) != 2 || parts[0] == "" || parts[0] == "." || parts[0] == ".." || parts[1] != "doctor-"+role+"-1.log" {
		t.Errorf("%s log path = %q, want .fireweed/runs/<run>/doctor-%s-1.log", role, path, role)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("read %s log %q: %v", role, path, err)
		return path
	}
	for _, marker := range markers {
		if !strings.Contains(string(data), marker) {
			t.Errorf("%s log lacks %q: %q", role, marker, data)
		}
	}
	return path
}

func doctorLogRun(t *testing.T, f doctorFixture, worker, audit string, report checks.DoctorReport) string {
	t.Helper()
	base := filepath.Join(f.root, ".fireweed", "runs")
	w, ew := filepath.Rel(base, worker)
	a, ea := filepath.Rel(base, audit)
	if ew != nil || ea != nil {
		t.Errorf("log paths outside run root: %q %q: %v %v", worker, audit, ew, ea)
		return ""
	}
	wRun, aRun := filepath.Dir(w), filepath.Dir(a)
	if wRun != aRun || wRun == "." || wRun == ".." {
		t.Errorf("worker and audit have different run identities: %q %q: %s", worker, audit, doctorText(report))
	}
	return wRun
}

func doctorCallerClean(t *testing.T, f doctorFixture, status, registrations string) {
	t.Helper()
	if got := f.git("status", "--porcelain", "--untracked-files=all"); got != status {
		t.Errorf("caller status changed: before %q after %q", status, got)
	}
	if got := worktrees(f); got != registrations {
		t.Errorf("scratch registration changed: before %q after %q", registrations, got)
	}
	assertScratchRemoved(t, f)
}

func Test_CHK_081_DoctorProbeLogsShareFreshRunAndPreserveStreams(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/agent.sh", "#!/bin/bash\nprintf 'worker stdout Ω without newline'; printf 'worker stderr β without newline' >&2; printf 'FIREWEED_DOCTOR_OK\\n'\n")
	f.put(".fireweed/audit.sh", "#!/bin/bash\nprintf 'audit stdout Ω without newline'; printf 'audit stderr β without newline' >&2; printf 'FIREWEED_DOCTOR_OK\\n'\n")
	f.commit()
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	var firstRun string
	for i := 0; i < 2; i++ {
		report := checks.Doctor(context.Background(), f.request(false))
		if doctorErrors(report) || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
			t.Errorf("successful probes changed outcome: %s", doctorText(report))
		}
		worker := doctorProbeLog(t, f, report, "worker", "worker stdout Ω without newlineFIREWEED_DOCTOR_OK\n", "worker stderr β without newline")
		audit := doctorProbeLog(t, f, report, "audit", "audit stdout Ω without newlineFIREWEED_DOCTOR_OK\n", "audit stderr β without newline")
		if worker != "" && audit != "" {
			run := doctorLogRun(t, f, worker, audit, report)
			if i == 1 && run == firstRun {
				t.Errorf("consecutive Doctor invocations reused run identity %q", run)
			}
			firstRun = run
		}
		doctorCallerClean(t, f, status, registrations)
	}
}

func Test_CHK_081_DoctorPersistsFailedAndTimedOutProbeOutput(t *testing.T) {
	failureStdout := "early worker stdout Ω\n"
	for i := 1; i <= 25; i++ {
		failureStdout += fmt.Sprintf("later-%02d\n", i)
	}
	for _, tc := range []struct {
		name, role, script, stdout, stderr string
		timeout                            time.Duration
	}{
		{"failure", "worker", "#!/bin/bash\nprintf 'early worker stdout Ω\\n'; for i in $(seq 1 25); do printf 'later-%02d\\n' \"$i\"; done; printf 'failed worker stderr β' >&2; exit 7\n", failureStdout, "failed worker stderr β", 2 * time.Second},
		{"timeout", "audit", "#!/bin/bash\nprintf 'partial audit stdout Ω'; printf 'partial audit stderr β' >&2; sleep 5\n", "partial audit stdout Ω", "partial audit stderr β", 150 * time.Millisecond},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			if tc.role == "worker" {
				f.put(".fireweed/agent.sh", tc.script)
			} else {
				f.put(".fireweed/audit.sh", tc.script)
			}
			f.commit()
			status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
			req := f.request(false)
			req.Timeout = tc.timeout
			start := time.Now()
			report := checks.Doctor(context.Background(), req)
			if time.Since(start) > 4*time.Second {
				t.Errorf("%s probe did not honor bounded request timeout", tc.name)
			}
			if !hasDoctorCheck(report, tc.role, false) || !hasDoctorDiagnostic(report, "C15", false, tc.role, tc.stderr) {
				t.Errorf("%s probe outcome/diagnostic lost: %s", tc.name, doctorText(report))
			}
			path := doctorProbeLog(t, f, report, tc.role, tc.stdout, tc.stderr)
			if tc.name == "failure" && path != "" {
				if !strings.Contains(string(mustRead(t, path)), "later-25") || strings.Contains(doctorText(report), "early worker stdout Ω") {
					t.Errorf("failed probe log or final-20-line diagnostic changed: %s", doctorText(report))
				}
			}
			doctorCallerClean(t, f, status, registrations)
		})
	}
}

func Test_CHK_081_DoctorPersistsCancelledProbeOutput(t *testing.T) {
	f := localDoctorFixture(t)
	ready := filepath.Join(t.TempDir(), "worker-ready")
	f.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/bash\nprintf 'cancel worker stdout Ω'; printf 'cancel worker stderr β' >&2; printf ready > %q; sleep 5\n", ready))
	f.commit()
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reports := make(chan checks.DoctorReport, 1)
	go func() { reports <- checks.Doctor(ctx, f.request(false)) }()
	deadline := time.After(8 * time.Second)
	for {
		if _, err := os.Stat(ready); err == nil {
			break
		} else if !os.IsNotExist(err) {
			t.Fatal(err)
		}
		select {
		case report := <-reports:
			t.Fatalf("worker exited before cancellation: %s", doctorText(report))
		case <-deadline:
			t.Fatal("worker never signalled readiness")
		case <-time.After(10 * time.Millisecond):
		}
	}
	cancel()
	select {
	case report := <-reports:
		if !hasDoctorCheck(report, "worker", false) || !hasDoctorDiagnostic(report, "C15", false, "worker") {
			t.Errorf("cancelled probe outcome lost: %s", doctorText(report))
		}
		doctorProbeLog(t, f, report, "worker", "cancel worker stdout Ω", "cancel worker stderr β")
	case <-time.After(4 * time.Second):
		t.Fatal("cancelled Doctor did not return promptly")
	}
	doctorCallerClean(t, f, status, registrations)
}
package evals_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
)

const doctorProbePrompt = "Reply with exactly FIREWEED_DOCTOR_OK"

func doctorRecordedPath(t *testing.T, log, label string) string {
	t.Helper()
	return doctorRecordedPaths(t, log, label, 1)[0]
}

func doctorRecordedPaths(t *testing.T, log, label string, count int) []string {
	t.Helper()
	var paths []string
	for _, line := range strings.Split(log, "\n") {
		if strings.HasPrefix(line, label+":") {
			paths = append(paths, strings.TrimPrefix(line, label+":"))
		}
	}
	if len(paths) != count {
		t.Fatalf("%s paths = %q; want %d absolute paths", label, paths, count)
	}
	for _, path := range paths {
		if !filepath.IsAbs(path) {
			t.Fatalf("%s path = %q; want an absolute path", label, path)
		}
	}
	return paths
}

func doctorPromptFileRemoved(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Errorf("prompt file %q remains after probe: %v", path, err)
	}
}

func Test_CHK_082_CHK_083_DoctorExpandsSelectedCommandsAndWrapperOnce(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) {
		config["agent"].(map[string]any)["command"] = []any{"bash", ".fireweed/agent.sh", "pre{prompt}post", "file={prompt_file}", "common={git_common_dir}", "literal {unknown}"}
		config["audit"].(map[string]any)["command"] = "bash .fireweed/audit.sh {prompt} {prompt_file} {git_common_dir}"
		config["agent"].(map[string]any)["wrap"] = []any{"bash", ".fireweed/prompt-wrap.sh", "{prompt_file}", "{prompt}", "{git_common_dir}"}
	})
	f.put(".fireweed/prompt-wrap.sh", fmt.Sprintf(`#!/bin/bash
printf 'wrap-file:%%s\nwrap-prompt:%%s\nwrap-common:%%s\n' "$1" "$2" "$3" >> %q
test -f "$1" && printf '%%s' "$2" | cmp -s - "$1" || exit 11
shift 3
exec "$@"
`, f.log))
	f.put(".fireweed/agent.sh", fmt.Sprintf(`#!/bin/bash
printf 'worker-prompt:%%s\nworker-file:%%s\nworker-common:%%s\nworker-literal:%%s\n' "$1" "${2#file=}" "${3#common=}" "$4" >> %q
test "$1" = "pre%spost" || exit 12
test "$2" != "file=" || exit 13
test -f "${2#file=}" && printf '%%s' %q | cmp -s - "${2#file=}" || exit 14
printf 'FIREWEED_DOCTOR_OK\n'
`, f.log, doctorProbePrompt, doctorProbePrompt))
	f.put(".fireweed/audit.sh", fmt.Sprintf(`#!/bin/bash
printf 'audit-prompt:%%s\naudit-file:%%s\naudit-common:%%s\n' "$1" "$2" "$3" >> %q
test -f "$2" && printf '%%s' "$1" | cmp -s - "$2" || exit 15
printf 'FIREWEED_DOCTOR_OK\n'
`, f.log))
	f.commit()
	// Relocate the Git directory so its absolute path contains a quote, a space,
	// and another placeholder. The expanded value must remain literal.
	commonDir := filepath.Join(t.TempDir(), "space 'quote' {prompt}", "repo.git")
	if err := os.MkdirAll(filepath.Dir(commonDir), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Rename(filepath.Join(f.root, ".git"), commonDir); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(f.root, ".git"), []byte("gitdir: "+commonDir+"\n"), 0644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("TMPDIR", filepath.Dir(commonDir))
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if doctorErrors(report) || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
		t.Errorf("prompt-file probes failed: %s", doctorText(report))
	}
	log := doctorLog(f)
	common, err := filepath.EvalSymlinks(commonDir)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"worker-prompt:pre" + doctorProbePrompt + "post", "audit-prompt:" + doctorProbePrompt,
		"worker-common:" + common, "audit-common:" + common,
		"worker-literal:literal {unknown}",
		"wrap-prompt:" + doctorProbePrompt, "wrap-common:" + common,
	} {
		if !strings.Contains(log, want+"\n") {
			t.Errorf("missing exact argument %q", want)
		}
	}
	workerFile := doctorRecordedPath(t, log, "worker-file")
	auditFile := doctorRecordedPath(t, log, "audit-file")
	if workerFile == auditFile {
		t.Errorf("worker and audit reused prompt file %q", workerFile)
	}
	for _, path := range []string{workerFile, auditFile} {
		if !strings.Contains(log, "wrap-file:"+path+"\n") {
			t.Errorf("wrapper did not receive selected probe file %q", path)
		}
		doctorPromptFileRemoved(t, path)
	}
	doctorProbeLog(t, f, report, "worker", "FIREWEED_DOCTOR_OK")
	doctorProbeLog(t, f, report, "audit", "FIREWEED_DOCTOR_OK")
	doctorCallerClean(t, f, status, registrations)
}

func Test_CHK_083_DoctorRemovesPromptFileOnProbeFailures(t *testing.T) {
	for _, mode := range []string{"launch", "nonzero", "timeout", "cancel"} {
		t.Run(mode, func(t *testing.T) {
			f := localDoctorFixture(t)
			ready := filepath.Join(t.TempDir(), "ready")
			f.editConfig(func(config map[string]any) {
				config["agent"].(map[string]any)["command"] = []any{"bash", ".fireweed/agent.sh", "{prompt_file}"}
				if mode == "launch" {
					config["agent"].(map[string]any)["command"] = []any{filepath.Join(f.root, "missing-agent"), "{prompt_file}"}
					config["agent"].(map[string]any)["wrap"] = []any{"bash", ".fireweed/prompt-wrap.sh", "{prompt_file}"}
				}
			})
			f.put(".fireweed/prompt-wrap.sh", fmt.Sprintf("#!/bin/bash\nprintf 'worker-file:%%s\\n' \"$1\" >> %q\ntest -f \"$1\" || exit 11\nshift\nexec \"$@\"\n", f.log))
			script := fmt.Sprintf("#!/bin/bash\nprintf 'worker-file:%%s\\nworker-content:%%s\\n' \"$1\" \"$(cat \"$1\")\" >> %q\ntest -f \"$1\" && printf '%%s' %q | cmp -s - \"$1\" || exit 18\nprintf ready > %q\n", f.log, doctorProbePrompt, ready)
			switch mode {
			case "nonzero":
				script += "printf 'failed worker stderr\\n' >&2\nexit 7\n"
			case "timeout", "cancel":
				script += "printf 'waiting worker stdout\\n'\nsleep 5\n"
			}
			f.put(".fireweed/agent.sh", script)
			f.commit()
			status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
			req := f.request(false)
			if mode == "timeout" {
				req.Timeout = 2 * time.Second
			}
			ctx := context.Background()
			var cancel context.CancelFunc
			if mode == "cancel" {
				ctx, cancel = context.WithCancel(ctx)
				defer cancel()
			}
			var report checks.DoctorReport
			if mode == "cancel" {
				reports := make(chan checks.DoctorReport, 1)
				go func() { reports <- checks.Doctor(ctx, req) }()
				deadline := time.After(8 * time.Second)
				for {
					if _, err := os.Stat(ready); err == nil {
						break
					} else if !os.IsNotExist(err) {
						t.Fatal(err)
					}
					select {
					case report = <-reports:
						t.Fatalf("worker exited before cancellation: %s", doctorText(report))
					case <-deadline:
						t.Fatal("worker never signalled readiness")
					case <-time.After(10 * time.Millisecond):
					}
				}
				livePath := doctorRecordedPath(t, doctorLog(f), "worker-file")
				if data, err := os.ReadFile(livePath); err != nil || string(data) != doctorProbePrompt {
					t.Errorf("prompt file unavailable before cancellation: %q, %v", data, err)
				}
				cancel()
				select {
				case report = <-reports:
				case <-time.After(4 * time.Second):
					t.Fatal("cancelled Doctor did not return")
				}
			} else {
				report = checks.Doctor(ctx, req)
			}
			if !hasDoctorCheck(report, "worker", false) || !hasDoctorDiagnostic(report, "C15", false, "worker") {
				t.Errorf("%s probe outcome changed: %s", mode, doctorText(report))
			}
			log := doctorLog(f)
			count := 1
			if mode == "launch" {
				// The shared wrapper records both the failed worker launch and
				// the subsequent audit probe; both files must be cleaned up.
				count = 2
			}
			paths := doctorRecordedPaths(t, log, "worker-file", count)
			if mode != "launch" && !strings.Contains(log, "worker-content:"+doctorProbePrompt+"\n") {
				t.Errorf("%s worker could not read exact prompt: %q", mode, log)
			}
			for _, path := range paths {
				doctorPromptFileRemoved(t, path)
			}
			doctorProbeLog(t, f, report, "worker")
			doctorCallerClean(t, f, status, registrations)
		})
	}
}

func Test_CHK_083_DoctorDoesNotCreatePromptFileWithoutSelectedPlaceholder(t *testing.T) {
	f := localDoctorFixture(t)
	expected := filepath.Join(t.TempDir(), "expected-prompt")
	if err := os.WriteFile(expected, []byte(doctorProbePrompt), 0644); err != nil {
		t.Fatal(err)
	}
	script := fmt.Sprintf(`#!/bin/bash
while IFS= read -r candidate; do
  if test -f "$candidate" && cmp -s "$candidate" %q; then
    printf 'unexpected prompt file: %%s\n' "$candidate" >&2
    exit 19
  fi
done < <(find "$TMPDIR" "$PWD" -type f -print)
printf 'FIREWEED_DOCTOR_OK\n'
`, expected)
	f.put(".fireweed/agent.sh", script)
	f.put(".fireweed/audit.sh", script)
	f.commit()
	temp := t.TempDir()
	t.Setenv("TMPDIR", temp)
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if doctorErrors(report) || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
		t.Errorf("unneeded prompt file appeared: %s", doctorText(report))
	}
	doctorProbeLog(t, f, report, "worker", "FIREWEED_DOCTOR_OK")
	doctorProbeLog(t, f, report, "audit", "FIREWEED_DOCTOR_OK")
	doctorCallerClean(t, f, status, registrations)
}

func Test_CHK_082_CHK_083_DoctorWrapperOnlyFileWithAuditFallback(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) {
		delete(config["audit"].(map[string]any), "command")
		config["agent"].(map[string]any)["command"] = []any{"bash", ".fireweed/agent.sh", "{prompt}"}
		config["agent"].(map[string]any)["wrap"] = []any{"bash", ".fireweed/prompt-wrap.sh", "{prompt_file}"}
	})
	f.put(".fireweed/prompt-wrap.sh", fmt.Sprintf(`#!/bin/bash
printf 'wrap-file:%%s\n' "$1" >> %q
test -f "$1" && printf '%%s' %q | cmp -s - "$1" || exit 11
shift
exec "$@"
`, f.log, doctorProbePrompt))
	f.put(".fireweed/agent.sh", fmt.Sprintf(`#!/bin/bash
printf 'fallback-prompt:%%s\n' "$1" >> %q
test "$1" = %q || exit 12
printf 'FIREWEED_DOCTOR_OK\n'
`, f.log, doctorProbePrompt))
	f.commit()
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if doctorErrors(report) || !hasDoctorDiagnostic(report, "C15", true, "shared") || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
		t.Errorf("wrapper-only fallback changed: %s", doctorText(report))
	}
	log := doctorLog(f)
	if strings.Count(log, "fallback-prompt:"+doctorProbePrompt+"\n") != 2 {
		t.Errorf("fallback did not receive fixed prompt twice")
	}
	var paths []string
	for _, line := range strings.Split(log, "\n") {
		if strings.HasPrefix(line, "wrap-file:") {
			paths = append(paths, strings.TrimPrefix(line, "wrap-file:"))
		}
	}
	if len(paths) != 2 || paths[0] == paths[1] {
		t.Errorf("wrapper prompt files = %q; want distinct worker/audit paths", paths)
	}
	for _, path := range paths {
		if !filepath.IsAbs(path) {
			t.Errorf("wrapper prompt path is relative: %q", path)
		}
		doctorPromptFileRemoved(t, path)
	}
	doctorProbeLog(t, f, report, "worker", "FIREWEED_DOCTOR_OK")
	doctorProbeLog(t, f, report, "audit", "FIREWEED_DOCTOR_OK")
	doctorCallerClean(t, f, status, registrations)
}

 succeeded in 703ms:
		}
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Errorf("physical scratch directory %q remains or cannot be checked: %v", path, err)
		}
		return
	}
	t.Error("setup did not record a scratch path")
}

func Test_CHK_049_StaticDiagnosticsStopProbes(t *testing.T) {
	f := localDoctorFixture(t)
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n## Specs\n### ALP-001\nA.\n### ALP-001\nB.\n")
	f.commit()
	want := project.Check(f.root)
	if len(want) == 0 || !doctorErrors(checks.DoctorReport{Diagnostics: want}) {
		t.Fatal("fixture has no static diagnostic")
	}
	report := checks.Doctor(context.Background(), f.request(false))
	remaining := map[project.Diagnostic]int{}
	for _, d := range report.Diagnostics {
		remaining[d]++
	}
	for _, d := range want {
		if remaining[d] == 0 {
			t.Errorf("complete project diagnostic %+v lost: %s", d, doctorText(report))
		}
		remaining[d]--
	}
	if strings.Contains(doctorLog(f), "setup:") || strings.Contains(doctorLog(f), "eval:") || strings.Contains(doctorLog(f), "boundaries:") || strings.Contains(doctorLog(f), "worker:") || strings.Contains(doctorLog(f), "audit:") {
		t.Errorf("static error ran probes: %s", doctorLog(f))
	}
}

func Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_CHK_060_CHK_061_CHK_067_RealGoRecipe(t *testing.T) {
	f := newDoctorFixture(t)
	f.commit()
	f.bin = doctorBinary(t)
	f.editConfig(func(config map[string]any) {
		config["modules"] = []string{"packages/alpha"}
		config["hooks"].(map[string]any)["setup"] = fmt.Sprintf("printf 'setup:%%s\\n' \"$PWD\" >> %q; test -z \"$(git symbolic-ref -q HEAD)\" && test -f .fireweed/doctor-untracked && go mod download", f.log)
	})
	f.put(".fireweed/doctor-untracked", "ready\n")
	for _, d := range project.Check(f.root) {
		if !d.Warning {
			t.Fatalf("invalid Go doctor fixture: %+v", d)
		}
	}
	beforeStatus, beforeWorktrees := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) {
		t.Errorf("real Go recipe doctor failed: %s", doctorText(report))
	}
	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id", "recovery"} {
		if !hasDoctorCheck(report, name, true) {
			t.Errorf("missing passing %s outcome: %s", name, doctorText(report))
		}
	}
	recoveries := 0
	for _, check := range report.Checks {
		if strings.Contains(strings.ToLower(check.Name), "recovery") && check.Passed {
			recoveries++
		}
	}
	if recoveries < 4 {
		t.Errorf("clean recovery not reported after each canary: %s", doctorText(report))
	}
	modules := map[string]bool{}
	for _, check := range report.Checks {
		if check.Module != "" && check.Module != "alpha" {
			modules[check.Module] = true
		}
	}
	if len(modules) < 2 {
		t.Errorf("two distinct throwaway modules absent: %s", doctorText(report))
	}
	if !strings.Contains(doctorLog(f), "setup:") || !strings.Contains(doctorLog(f), "/fireweed/worktrees/") {
		t.Errorf("setup did not run in scratch: %s", doctorLog(f))
	}
	if strings.Contains(doctorLog(f), "worker:") || strings.Contains(doctorLog(f), "audit:") {
		t.Errorf("NoAgent ran agents: %s", doctorLog(f))
	}
	if got := f.git("status", "--porcelain", "--untracked-files=all"); got != beforeStatus {
		t.Errorf("caller changed: before %q after %q", beforeStatus, got)
	}
	if got := worktrees(f); got != beforeWorktrees {
		t.Errorf("scratch registration leaked: before %q after %q", beforeWorktrees, got)
	}
	assertScratchRemoved(t, f)
}

func Test_CHK_053_NestedModuleTemplatesRenderForBothCanaries(t *testing.T) {
	f := localDoctorFixture(t)
	f.put("go.mod", "module example.test/doctorfixture\n\ngo 1.23\n")
	fields := "{{name}}|{{prefix}}|{{package}}|{{go_module}}|{{module_dir}}|{{other_dir}}\n"
	f.put(".fireweed/templates/module/evals/nested.txt.tmpl", "eval|"+fields)
	f.put(".fireweed/templates/module/internal/nested.txt", "internal|"+fields)
	canaries := string(mustRead(t, filepath.Join(f.root, ".fireweed/canaries/canaries.yaml")))
	f.put(".fireweed/canaries/canaries.yaml", strings.Replace(canaries, "undeclared:\n", "  evals/baseline.txt: \"baseline|"+strings.TrimSpace(fields)+"\"\nundeclared:\n", 1))
	eval := string(mustRead(t, filepath.Join(f.root, ".fireweed/eval.sh")))
	check := fmt.Sprintf(`case "$1" in
  doctor*)
    prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
    other=
    for spec in packages/*/SPEC.md; do
      dir=${spec%%/SPEC.md}
      case "$dir" in packages/alpha|"$2") ;; *) other=$dir ;; esac
    done
    expected="$1|$prefix|$1|example.test/doctorfixture|$2|$other"
    test -n "$other" &&
      test "$(cat "$2/evals/nested.txt")" = "eval|$expected" &&
      test "$(cat "$2/internal/nested.txt")" = "internal|$expected" &&
      test "$(cat "$2/evals/baseline.txt")" = "baseline|$expected" &&
      test ! -e "$2/evals/nested.txt.tmpl" || exit 17
    printf 'rendered:%%s\n' "$1" >> %q
    ;;
esac
`, f.log)
	f.put(".fireweed/eval.sh", strings.Replace(eval, "#!/bin/sh\n", "#!/bin/sh\n"+check, 1))
	f.commit()
	before := worktrees(f)
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) {
		t.Errorf("nested template probes failed: %s", doctorText(report))
	}
	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id", "recovery"} {
		if !hasDoctorCheck(report, name, true) {
			t.Errorf("missing passing %s probe: %s", name, doctorText(report))
		}
	}
	modules := map[string]bool{}
	for _, line := range strings.Split(doctorLog(f), "\n") {
		if strings.HasPrefix(line, "rendered:") {
			modules[strings.TrimPrefix(line, "rendered:")] = true
		}
	}
	if len(modules) != 2 {
		t.Errorf("expected both throwaway modules to pass nested template checks, got %v: %s", modules, doctorText(report))
	}
	if got := worktrees(f); got != before {
		t.Errorf("scratch registration leaked: before %q after %q", before, got)
	}
	if strings.Contains(doctorLog(f), "setup:") {
		assertScratchRemoved(t, f)
	}
}

func Test_CHK_051_CHK_052_SetupFailureStopsAndCleans(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/setup.sh", fmt.Sprintf("#!/bin/sh\nprintf 'setup:%%s\\n' \"$PWD\" >> %q\necho failed setup >&2\nexit 17\n", f.log))
	f.commit()
	before := worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if !hasDoctorDiagnostic(report, "C14", false, "setup") {
		t.Errorf("setup failure absent: %s", doctorText(report))
	}
	if strings.Contains(doctorLog(f), "eval:") || strings.Contains(doctorLog(f), "boundaries:") || strings.Contains(doctorLog(f), "worker:") {
		t.Errorf("setup failure ran later probes: %s", doctorLog(f))
	}
	if got := worktrees(f); got != before {
		t.Errorf("scratch worktree leaked: %q", got)
	}
	assertScratchRemoved(t, f)
}

func Test_CHK_051_CancellationCleansScratchAndStopsSetupGroup(t *testing.T) {
	f := localDoctorFixture(t)
	ready := filepath.Join(t.TempDir(), "ready")
	childFinished := filepath.Join(t.TempDir(), "child-finished")
	f.put(".fireweed/setup.sh", fmt.Sprintf(`#!/bin/sh
printf 'setup:%%s\n' "$PWD" >> %q
sh -c 'sleep 2; printf finished > "$1"' sh %q &
printf ready > %q
wait
`, f.log, childFinished, ready))
	f.commit()
	before := worktrees(f)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reports := make(chan checks.DoctorReport, 1)
	go func() { reports <- checks.Doctor(ctx, f.request(true)) }()
	deadline := time.After(8 * time.Second)
	for {
		if _, err := os.Stat(ready); err == nil {
			break
		} else if !os.IsNotExist(err) {
			t.Fatal(err)
		}
		select {
		case report := <-reports:
			t.Fatalf("setup exited before cancellation: %s", doctorText(report))
		case <-deadline:
			t.Fatal("setup never signalled readiness")
		case <-time.After(10 * time.Millisecond):
		}
	}
	cancel()
	select {
	case report := <-reports:
		if !hasDoctorDiagnostic(report, "C14", false, "setup") {
			t.Errorf("cancelled setup lacked C14: %s", doctorText(report))
		}
	case <-time.After(5 * time.Second):
		t.Fatal("cancelled doctor did not return promptly")
	}
	if got := worktrees(f); got != before {
		t.Errorf("cancelled doctor retained Git registration: before %q after %q", before, got)
	}
	assertScratchRemoved(t, f)
	time.Sleep(2200 * time.Millisecond)
	if _, err := os.Stat(childFinished); !os.IsNotExist(err) {
		t.Errorf("cancelled setup descendant continued: %v", err)
	}
}

func Test_CHK_054_MissingTemplatesAndCanaryMaps(t *testing.T) {
	for _, tc := range []struct{ name, path, body, term, input, affected string }{
		{"missing template", ".fireweed/templates/module/SPEC.md", "", "template", "SPEC.md", "undeclared"},
		{"malformed template", ".fireweed/templates/module/SPEC.md", "---\nname: [broken\nid_prefix: {{prefix}}\n---\n", "template", "SPEC.md", "undeclared"},
		{"missing baseline", ".fireweed/canaries/canaries.yaml", "version: 1\nbaseline: {}\n", "canary", "baseline", "undeclared"},
		{"missing internal_eval", ".fireweed/canaries/canaries.yaml", "missing_id:\n  evals/missing.txt: missing\n", "internal_eval", "internal_eval", "internal_eval"},
		{"wrong version", ".fireweed/canaries/canaries.yaml", "version: 2\nbaseline: {}\n", "canary", "canaries.yaml", "undeclared"},
		{"malformed canary", ".fireweed/canaries/canaries.yaml", "version: [bad\n", "canary", "canaries.yaml", "undeclared"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			if tc.name == "missing baseline" {
				body := string(mustRead(t, filepath.Join(f.root, tc.path)))
				start, end := strings.Index(body, "baseline:\n"), strings.Index(body, "undeclared:\n")
				if start < 0 || end <= start {
					t.Fatal("fixture cannot isolate baseline map")
				}
				tc.body = body[:start] + "baseline: {}\n" + body[end:]
			}
			if tc.name == "missing internal_eval" {
				body := string(mustRead(t, filepath.Join(f.root, tc.path)))
				tc.body = strings.Replace(body, "internal_eval:\n  internal/private.txt: \"private\"\n  evals/internal.txt: \"{{module_dir}}\"\n", "", 1)
				if tc.body == body {
					t.Fatal("fixture did not remove internal_eval map")
				}
			}
			if tc.body == "" {
				if err := os.Remove(filepath.Join(f.root, tc.path)); err != nil {
					t.Fatal(err)
				}
			} else {
				f.put(tc.path, tc.body)
			}
			f.commit()
			report := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(report, "C14", false, tc.term) {
				t.Errorf("missing C14: %s", doctorText(report))
			}
			identified := false
			for _, d := range report.Diagnostics {
				if d.Code == "C14" && !d.Warning && strings.Contains(strings.ToLower(d.Message+" "+d.Path), strings.ToLower(tc.input)) {
					identified = true
				}
			}
			if !identified {
				t.Errorf("C14 did not identify %s: %s", tc.input, doctorText(report))
			}
			if hasDoctorCheck(report, tc.affected, true) {
				t.Errorf("invalid recipe passed canary: %s", doctorText(report))
			}
		})
	}
}

func Test_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_ZeroHookAndRecovery(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/boundary.sh", "#!/bin/sh\nexit 0\n")
	f.put(".fireweed/eval.sh", `#!/bin/sh
printf '<testsuite><testcase name="ALP-001"/></testsuite>' > "$3"
`)
	f.commit()
	before := worktrees(f)
	report := checks.Doctor(context.Background(), f.request(true))
	for _, tc := range []struct{ name, hook string }{{"undeclared", "boundar"}, {"internal_eval", "boundar"}, {"failing_eval", "eval"}, {"missing_id", "eval"}} {
		if !hasDoctorDiagnostic(report, "C14", false, tc.name, tc.hook) {
			t.Errorf("missed %s canary not diagnosed with hook: %s", tc.name, doctorText(report))
		}
	}
	if !doctorErrors(report) {
		t.Errorf("zero hooks accepted: %s", doctorText(report))
	}
	if got := worktrees(f); got != before {
		t.Errorf("failed canary leaked worktree: %s", got)
	}

	g := localDoctorFixture(t)
	g.put(".fireweed/boundary.sh", fmt.Sprintf(`#!/bin/sh
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test -f "$dir/undeclared.txt"; then printf '%%s' "$name" > .fireweed/recovery-seen; printf '%%s: undeclared import\n' "$name"; fi
done
if test -f .fireweed/recovery-seen && ! find packages -name undeclared.txt | grep -q .; then printf '%%s: lingering violation\n' "$(cat .fireweed/recovery-seen)"; fi
`))
	g.commit()
	r := checks.Doctor(context.Background(), g.request(true))
	if !hasDoctorDiagnostic(r, "C14", false, "recovery", "undeclared", "boundar") {
		t.Errorf("missing failed recovery diagnostic: %s", doctorText(r))
	}
}

func Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific(t *testing.T) {
	for _, tc := range []struct{ name, script, canary string }{
		{"missing baseline coverage", `#!/bin/sh
printf '<testsuite/>' > "$3"
`, "baseline"},
		{"failing eval compile error without report", `#!/bin/sh
if test -f "$2/evals/failing.txt"; then echo 'compile error' >&2; exit 2; fi
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "failing_eval"},
		{"failing eval with unrelated failure", `#!/bin/sh
if test -f "$2/evals/failing.txt"; then printf '<testsuite><testcase name="ALP-001"><failure/></testcase></testsuite>' > "$3"; exit 1; fi
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "failing_eval"},
		{"missing id with unrelated failure", `#!/bin/sh
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
if test -f "$2/evals/missing.txt"; then printf '<testsuite><testcase name="%s-001"><failure/></testcase></testsuite>' "$prefix" > "$3"; exit 1; fi
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "missing_id"},
		{"missing id without report", `#!/bin/sh
if test -f "$2/evals/missing.txt"; then echo 'hook failed' >&2; exit 2; fi
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "missing_id"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			f.put(".fireweed/eval.sh", tc.script)
			if tc.name == "missing baseline coverage" {
				logDoctorCanaryProbes(t, f)
			}
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(r, "C14", false, tc.canary, "eval") {
				t.Errorf("unrelated eval outcome accepted: %s", doctorText(r))
			}
			if tc.name == "missing baseline coverage" && !strings.Contains(doctorLog(f), "probe:") {
				t.Errorf("failed baseline coverage hook did not run: %s", doctorLog(f))
			}
			if tc.name == "missing baseline coverage" && strings.Contains(doctorLog(f), "planted:") {
				t.Errorf("failed baseline coverage continued into canary probes: %s", doctorLog(f))
			}
		})
	}
}

func Test_CHK_055_BaselineAndRecoveryFailures(t *testing.T) {
	for _, tc := range []struct{ name, path, script, phase, hook string }{
		{"baseline eval failure", ".fireweed/eval.sh", "#!/bin/sh\necho baseline-eval-failed >&2\nexit 9\n", "baseline", "eval"},
		{"baseline boundary failure", ".fireweed/boundary.sh", "#!/bin/sh\nprintf 'alpha: baseline violation\\n'\n", "baseline", "boundar"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			f.put(tc.path, tc.script)
			logDoctorCanaryProbes(t, f)
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(r, "C14", false, tc.phase, tc.hook) {
				t.Errorf("failed baseline accepted: %s", doctorText(r))
			}
			if !strings.Contains(doctorLog(f), "probe:") {
				t.Errorf("failed baseline hook did not run: %s", doctorLog(f))
			}
			if strings.Contains(doctorLog(f), "planted:") {
				t.Errorf("failed baseline continued into canary probes: %s", doctorLog(f))
			}
		})
	}
	for _, tc := range []struct{ name, failure string }{
		{"recovery eval failure", "echo recovery-eval-failed >&2; exit 9"},
		{"recovery coverage failure", "printf '<testsuite/>' > \"$3\"; exit 0"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			original := string(mustRead(t, filepath.Join(f.root, ".fireweed/eval.sh")))
			injection := "if test -f \"$2/evals/failing.txt\"; then printf '%s' \"$2\" > .fireweed/eval-recovery-module; fi\n" +
				"if test -f .fireweed/eval-recovery-module && test \"$(cat .fireweed/eval-recovery-module)\" = \"$2\" && ! test -f \"$2/evals/failing.txt\"; then " + tc.failure + "; fi\n"
			f.put(".fireweed/eval.sh", strings.Replace(original, "#!/bin/sh\n", "#!/bin/sh\n"+injection, 1))
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(r, "C14", false, "recovery", "failing_eval", "eval") {
				t.Errorf("failed %s accepted: %s", tc.name, doctorText(r))
			}
		})
	}
}

func Test_CHK_056_CHK_057_BoundaryViolationMustNameCanaryModule(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/boundary.sh", `#!/bin/sh
if find packages \( -name undeclared.txt -o -name internal.txt \) | grep -q .; then printf 'alpha: unrelated violation\n'; fi
`)
	f.commit()
	r := checks.Doctor(context.Background(), f.request(true))
	for _, name := range []string{"undeclared", "internal_eval"} {
		if !hasDoctorDiagnostic(r, "C14", false, name, "boundar") {
			t.Errorf("unrelated module violation accepted for %s: %s", name, doctorText(r))
		}
	}
}

func Test_CHK_057_InternalEvalRequiresInterfaceViolation(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/boundary.sh", `#!/bin/sh
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test -f "$dir/undeclared.txt"; then printf '%s: undeclared import\n' "$name"; fi
  if test -f "$dir/evals/internal.txt"; then printf '%s: unrelated boundary violation\n' "$name"; fi
done
`)
	f.commit()
	r := checks.Doctor(context.Background(), f.request(true))
	passedProbe := false
	for _, check := range r.Checks {
		if check.Passed && (check.Name == "internal_eval" || strings.HasPrefix(check.Name, "internal_eval ")) {
			passedProbe = true

exec
/bin/zsh -lc "mkdir -p packages/checks/internal/engine && cat > packages/checks/internal/engine/core.go <<'EOF'
package engine

import (
 \"context\"
 \"encoding/xml\"
 \"errors\"
 \"fmt\"
 \"io\"
 \"os\"
 \"os/exec\"
 \"path/filepath\"
 \"regexp\"
 \"strconv\"
 \"strings\"
 \"time\"

 \"github.com/lurkshark/fireweed/packages/checks/model\"
 \"github.com/lurkshark/fireweed/packages/contracts\"
 \"github.com/lurkshark/fireweed/packages/project\"
)

func pref(code string, err error) error { return fmt.Errorf(\"%s: %w\", code, err) }
func stringMap(v any) map[string]any { if m,ok:=v.(map[string]any);ok{return m}; return nil }
func str(v any) string { s,_:=v.(string); return s }
func strSlice(v any) []string { switch x:=v.(type) { case []string:return x; case []any: var a []string;for _,v:=range x {if s,ok:=v.(string);ok {a=append(a,s)}};return a };return nil }
func configured(m map[string]any, key string) any { return stringMap(m[\"hooks\"])[key] }
func shellQuote(s string) string {return \"'\"+strings.ReplaceAll(s,\"'\",\"'\\\\''\")+\"'\"}
func expand(s string, vals map[string]string, shell bool) string { re:=regexp.MustCompile("'`'"\\{[a-z_]+\\}"'`);return re.ReplaceAllStringFunc(s,func(tok string)string{v,ok:=vals[tok[1:len(tok)-1]];if !ok{return tok};if shell{return shellQuote(v)};return v}) }
func command(value any, vals map[string]string) ([]string,error) {switch x:=value.(type){case string:if strings.TrimSpace(x)=="" {return nil,errors.New("empty hook command")};return []string{"sh","-c",expand(x,vals,true)},nil;case []string:if len(x)==0{return nil,errors.New("empty hook argv")};a:=make([]string,len(x));for i,v:=range x{a[i]=expand(v,vals,false)};return a,nil;case []any:if len(x)==0{return nil,errors.New("empty hook argv")};a:=make([]string,len(x));for i,v:=range x{s,ok:=v.(string);if !ok{return nil,errors.New("hook argv contains non-string")};a[i]=expand(s,vals,false)};return a,nil;default:return nil,errors.New("invalid hook command")}}
func run(ctx context.Context, dir string, argv []string, env map[string]string, timeout time.Duration) (string,string,int,error) { if timeout>0{var cancel context.CancelFunc;ctx,cancel=context.WithTimeout(ctx,timeout);defer cancel()};c:=exec.CommandContext(ctx,argv[0],argv[1:]...);c.Dir=dir;c.Env=os.Environ();for k,v:=range env{c.Env=append(c.Env,k+"="+v)};var out,errout strings.Builder;c.Stdout=&out;c.Stderr=&errout;err:=c.Run();exit:=0;if c.ProcessState!=nil{exit=c.ProcessState.ExitCode()};if ctx.Err()!=nil{return out.String(),errout.String(),exit,ctx.Err()};if err!=nil {if _,ok:=err.(*exec.ExitError);!ok{return out.String(),errout.String(),exit,err}};return out.String(),errout.String(),exit,nil }

type junitNode struct {XMLName xml.Name;Name string `xml:"name,attr"`;Classname string `xml:"classname,attr"`;Time string `xml:"time,attr"`;Message string `xml:"message,attr"`;Body string `xml:",chardata"`;Children []junitNode `xml:",any"`}
var unknownID=regexp.MustCompile(`(?i)[A-Z][A-Z0-9]*[-_]?'"\\d{3,}"'`)
func ParseJUnit(data []byte,specs []contracts.Item)(model.Report,error){var root junitNode;dec:=xml.NewDecoder(strings.NewReader(string(data)));if err:=dec.Decode(&root);err!=nil{return model.Report{},pref("C11",fmt.Errorf("parse JUnit: %w",err))};var extra any;if err:=dec.Decode(&extra);err!=io.EOF{return model.Report{},pref("C11",fmt.Errorf("parse JUnit: trailing XML"))};if root.XMLName.Local!="testsuite"&&root.XMLName.Local!="testsuites"{return model.Report{},pref("C11",fmt.Errorf("invalid JUnit root %q",root.XMLName.Local))};known:=map[string]bool{};for _,s:=range specs{known[strings.ToUpper(s.ID)]=true};var report model.Report;var walk func(junitNode)error;walk=func(n junitNode)error{if n.XMLName.Local=="testcase" {c:=model.Case{Name:n.Name,Classname:n.Classname,Result:"pass"};if n.Time!=""{v,e:=strconv.ParseFloat(n.Time,64);if e!=nil{return fmt.Errorf("testcase %q invalid time: %w",n.Name,e)};c.Seconds=v};for _,s:=range specs{if contracts.MatchesID(s.ID,n.Name,n.Classname){c.IDs=append(c.IDs,s.ID)}};if len(c.IDs)==0 {for _,id:=range unknownID.FindAllString(n.Name+" "+n.Classname,-1){normalized:=strings.ToUpper(strings.ReplaceAll(id,"_","-"));if !known[normalized] {return fmt.Errorf("testcase %q references unknown spec ID %q",n.Name,id)}};return fmt.Errorf("testcase %q has no known spec ID",n.Name)};var failures []string;for _,ch:=range n.Children{switch ch.XMLName.Local{case "failure","error":c.Result="fail";if strings.TrimSpace(ch.Message)!=""{failures=append(failures,strings.TrimSpace(ch.Message))};if strings.TrimSpace(ch.Body)!=""{failures=append(failures,strings.TrimSpace(ch.Body))};case "skipped":if c.Result!="fail"{c.Result="skip"}}};c.Failure=strings.Join(failures,"'"\\n\");report.Cases=append(report.Cases,c);report.Seconds+=c.Seconds;return nil};for _,ch:=range n.Children{if err:=walk(ch);err"'!=nil{return err}};return nil};if err:=walk(root);err!=nil{return model.Report{},pref("C11",err)};return report,nil}
func hasID(ids []string,id string)bool{for _,v:=range ids{if v==id{return true}};return false}
func RequireCoverage(r model.Report,specs []contracts.Item)error{for _,s:=range specs{tag:=strings.ToLower(s.Tag);if strings.HasPrefix(tag,"review")||strings.HasPrefix(tag,"quarantine"){continue};found:=false;for _,c:=range r.Cases{if c.Result!="skip"&&hasID(c.IDs,s.ID){found=true;break}};if !found{return pref("C10",fmt.Errorf("no coverage for %s",s.ID))}};return nil}
func RequireFailFirst(r model.Report,ids,already []string)error{for _,id:=range ids{if hasID(already,id){continue};found:=false;for _,c:=range r.Cases{if (c.Result=="fail"||c.Result=="error")&&hasID(c.IDs,id){found=true;break}};if !found{return pref("P2-01",fmt.Errorf("no fail-first case for %s",id))}};return nil}
func RequirePass(r model.Report,specs []contracts.Item)error{tags:=map[string]string{};for _,s:=range specs{tags[s.ID]=s.Tag};for _,c:=range r.Cases{if c.Result!="fail"&&c.Result!="error"{continue};for _,id:=range c.IDs{tag,ok:=tags[id];if ok&&!strings.HasPrefix(strings.ToLower(tag),"quarantine"){return pref("E01",fmt.Errorf("failed case %q for %s",c.Name,id))}}};return nil}
func moduleSpecs(m project.Module)([]contracts.Item,error){section,ok:=contracts.Section(m.Document.Body,"Specs");if !ok{return nil,nil};return contracts.ParseItems(section)}
func Evaluate(ctx context.Context,req model.HookRequest)(model.Report,error){var mod *project.Module;for i:=range req.Snapshot.Modules{if req.Snapshot.Modules[i].Name==req.Module{mod=&req.Snapshot.Modules[i];break}};if mod==nil{return model.Report{},pref("U02",fmt.Errorf("unknown module %q",req.Module))};specs,err:=moduleSpecs(*mod);if err!=nil{return model.Report{},err};hook:=configured(req.Snapshot.Config,"eval");if h:=configured(mod.Document.Frontmatter,"eval");h!=nil{hook=h};if hook==nil{return model.Report{},pref("C11",fmt.Errorf("missing eval hook for %s",req.Module))};f,err:=os.CreateTemp("","fireweed-junit-*.xml");if err!=nil{return model.Report{},pref("C11",err)};reportPath:=f.Name();f.Close();os.Remove(reportPath);defer os.Remove(reportPath);args,err:=command(hook,map[string]string{"module":mod.Name,"module_dir":mod.Dir,"report":reportPath});if err!=nil{return model.Report{},pref("C11",err)};out,stderr,exit,err:=run(ctx,req.Snapshot.Root,args,map[string]string{"FIREWEED_BIN":req.Bin},req.Timeout);if err!=nil{return model.Report{},pref("C11",fmt.Errorf("run hook: %w",err))};data,err:=os.ReadFile(reportPath);if err!=nil{return model.Report{},pref("C11",fmt.Errorf("read JUnit report: %w",err))};report,err:=ParseJUnit(data,specs);if err!=nil{return model.Report{},err};report.Stdout=out;report.Stderr=stderr;report.HookExit=exit;failed:=false;for _,c:=range report.Cases{if c.Result=="fail"{failed=true}};if (exit==0)==failed{return report,pref("C11",fmt.Errorf("eval hook exit %d disagrees with JUnit failures",exit))};return report,nil}
EOF

gofmt -w packages/checks/internal/engine/core.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 584ms:
exec
/bin/zsh -lc "cat > packages/checks/internal/engine/zones_boundaries.go <<'EOF'
package engine

import (
 \"context\"
 \"fmt\"
 \"os\"
 \"path\"
 \"path/filepath\"
 \"strings\"
 \"time\"

 \"github.com/lurkshark/fireweed/packages/checks/model\"
 \"github.com/lurkshark/fireweed/packages/contracts\"
 \"github.com/lurkshark/fireweed/packages/project\"
)
func matchAny(globs []string,p string)bool{for _,g:=range globs{if contracts.MatchGlob(g,p){return true}};return false}
func baseExists(root,base,p string)bool{if base==\"\"{return false};_,err:=os.Stat(filepath.Join(root,p));_ = err; // recorded phase base is authoritative
 _,err2:=readAt(root,base,p);return err2==nil}
func zonePath(r model.ZoneRequest,p,status string) *project.Diagnostic {p=path.Clean(strings.TrimPrefix(p,\"./\"));bad:=func(reason string)*project.Diagnostic{return &project.Diagnostic{Code:\"Z01\",Path:p,Message:reason+\": \"+p}};if p==r.ChangePath {if r.Phase==\"audit\"||r.Phase==\"evals\"||r.Phase==\"impl\" {return nil}};if strings.HasPrefix(p,\"changes/\"){return bad(\"other change file\")};if p==\".fireweed\"||strings.HasPrefix(p,\".fireweed/\"){return bad(\"protected path\")};if matchAny(strSlice(r.Snapshot.Config[\"protected\"]),p){return bad(\"protected path\")};if r.Phase==\"audit\"{return bad(\"audit edit\")};var owner *project.Module;for i:=range r.Snapshot.Modules{m:=&r.Snapshot.Modules[i];if p==m.Dir||strings.HasPrefix(p,m.Dir+\"/\"){if owner==nil||len(m.Dir)>len(owner.Dir){owner=m}}};if owner==nil{if matchAny(strSlice(r.Snapshot.Config[\"worker_writable\"]),p){return nil};return bad(\"outside worker writable\")};if owner.Name"'!=r.Module{return bad("other module")};rel:=strings.TrimPrefix(p,owner.Dir+"/");if rel=="SPEC.md"||matchAny(owner.Zones.Protected,rel){return bad("protected module path")};zone:=contracts.ResolveZone(rel,owner.Zones);switch r.Phase {case "evals":switch zone{case "evals","other":return nil;case "interface":if r.InterfaceDeclared{return nil};case "disposable":if !baseExists(r.Snapshot.Root,r.Base,p){return nil}};case "impl":switch zone{case "other","disposable","preserve":return nil;case "append_only":if strings.HasPrefix(status,"A")&&!baseExists(r.Snapshot.Root,r.Base,p){return nil}}};return bad("phase zone")}
func Zones(r model.ZoneRequest)[]project.Diagnostic{var d []project.Diagnostic;for _,c:=range r.Changes{if c.OldPath!=""{if x:=zonePath(r,c.OldPath,"D");x!=nil{d=append(d,*x)}};if x:=zonePath(r,c.Path,c.Status);x!=nil{d=append(d,*x)}};return d}
func boundaryHook(s project.Snapshot,m *project.Module)any{if m!=nil{if v:=configured(m.Document.Frontmatter,"boundaries");v!=nil{return v}};return configured(s.Config,"boundaries")}
func boundaryLines(text string)[]string{var a []string;for _,l:=range strings.Split(strings.ReplaceAll(text,"'"\\r\\n\",\"\\n\"),\"\\n\"){if strings.TrimSpace(l)"'!=""{a=append(a,l)}};return a}
func boundaryOpt(m project.Module)string{return strings.TrimSpace(str(m.Document.Frontmatter["boundaries_opt_out"]))}
func Boundaries(ctx context.Context,r model.BoundaryRequest)([]project.Diagnostic,error){var diags []project.Diagnostic;var first error;global:=configured(r.Snapshot.Config,"boundaries");baseline:=map[string]bool{};if b,e:=os.ReadFile(filepath.Join(r.Snapshot.Root,".fireweed","boundaries-baseline"));e==nil{for _,l:=range boundaryLines(string(b)){baseline[l]=true}};for _,m:=range r.Snapshot.Modules{if reason:=boundaryOpt(m);reason!=""{diags=append(diags,project.Diagnostic{Code:"X03",Warning:true,Message:m.Name+": boundaries opt out: "+reason})}};type task struct{hook any;module *project.Module};var tasks []task;if global!=nil{tasks=append(tasks,task{hook:global})};for i:=range r.Snapshot.Modules{m:=&r.Snapshot.Modules[i];h:=configured(m.Document.Frontmatter,"boundaries");if h!=nil{tasks=append(tasks,task{hook:h,module:m});continue};if global==nil {if len(m.Zones.Disposable)>0&&boundaryOpt(*m)==""{d:=project.Diagnostic{Code:"X02",Message:m.Name+": missing boundaries hook"};diags=append(diags,d);if first==nil{first=fmt.Errorf(d.Message)}}else{diags=append(diags,project.Diagnostic{Code:"X04",Warning:true,Message:m.Name+": no boundary enforcement required"})}}};if len(r.Snapshot.Modules)==0&&global==nil{diags=append(diags,project.Diagnostic{Code:"X04",Warning:true,Message:"no boundary enforcement required"})};for _,t:=range tasks{vals:=map[string]string{};if t.module!=nil{vals["module"]=t.module.Name;vals["module_dir"]=t.module.Dir};args,e:=command(t.hook,vals);if e!=nil{diags=append(diags,project.Diagnostic{Code:"X02",Message:e.Error()});if first==nil{first=e};continue};out,stderr,_,e:=run(ctx,r.Snapshot.Root,args,map[string]string{"FIREWEED_BIN":r.Bin},r.Timeout);if e!=nil{diags=append(diags,project.Diagnostic{Code:"X02",Message:fmt.Sprintf("run boundaries hook: %v",e)});if first==nil{first=e};continue};for _,line:=range boundaryLines(out+"'"\\n\"+stderr){if baseline[line]{continue};owner:=\"\";for _,m:=range r.Snapshot.Modules{if strings.HasPrefix(line,m.Name+\":\"){owner=m.Name;if boundaryOpt(m)"'!=""{owner=""};break}};if owner==""{continue};diags=append(diags,project.Diagnostic{Code:"X01",Message:line});if first==nil{first=fmt.Errorf("boundary violation: %s",line)}}};return diags,first}
var _ time.Duration
EOF
python3 - <<'"'PY'
p='packages/checks/internal/engine/zones_boundaries.go'
s=open(p).read().replace('\"time\"','\"github.com/lurkshark/fireweed/packages/gitstore\"').replace('_,err:=os.Stat(filepath.Join(root,p));_ = err; // recorded phase base is authoritative\\n _,err2:=readAt(root,base,p);return err2==nil','_,err:=gitstore.ReadAt(root,base,p);return err==nil').replace('var _ time.Duration','')
open(p,'w').write(s)
PY
gofmt -w packages/checks/internal/engine/zones_boundaries.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 629ms:
exec
/bin/zsh -lc "rg -n '"'^func Test_|''^const doctorProbe|''^func doctorProbe'"' packages/checks/evals/doctor*" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 530ms:
packages/checks/evals/doctor_agent_logs_test.go:15:func doctorProbeLog(t *testing.T, f doctorFixture, report checks.DoctorReport, role string, markers ...string) string {
packages/checks/evals/doctor_agent_logs_test.go:91:func Test_CHK_081_DoctorProbeLogsShareFreshRunAndPreserveStreams(t *testing.T) {
packages/checks/evals/doctor_agent_logs_test.go:116:func Test_CHK_081_DoctorPersistsFailedAndTimedOutProbeOutput(t *testing.T) {
packages/checks/evals/doctor_agent_logs_test.go:158:func Test_CHK_081_DoctorPersistsCancelledProbeOutput(t *testing.T) {
packages/checks/evals/doctor_prompt_file_test.go:15:const doctorProbePrompt = "Reply with exactly FIREWEED_DOCTOR_OK"
packages/checks/evals/doctor_prompt_file_test.go:48:func Test_CHK_082_CHK_083_DoctorExpandsSelectedCommandsAndWrapperOnce(t *testing.T) {
packages/checks/evals/doctor_prompt_file_test.go:123:func Test_CHK_083_DoctorRemovesPromptFileOnProbeFailures(t *testing.T) {
packages/checks/evals/doctor_prompt_file_test.go:211:func Test_CHK_083_DoctorDoesNotCreatePromptFileWithoutSelectedPlaceholder(t *testing.T) {
packages/checks/evals/doctor_prompt_file_test.go:241:func Test_CHK_082_CHK_083_DoctorWrapperOnlyFileWithAuditFallback(t *testing.T) {
packages/checks/evals/doctor_test.go:289:func Test_CHK_049_StaticDiagnosticsStopProbes(t *testing.T) {
packages/checks/evals/doctor_test.go:313:func Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_CHK_060_CHK_061_CHK_067_RealGoRecipe(t *testing.T) {
packages/checks/evals/doctor_test.go:370:func Test_CHK_053_NestedModuleTemplatesRenderForBothCanaries(t *testing.T) {
packages/checks/evals/doctor_test.go:426:func Test_CHK_051_CHK_052_SetupFailureStopsAndCleans(t *testing.T) {
packages/checks/evals/doctor_test.go:444:func Test_CHK_051_CancellationCleansScratchAndStopsSetupGroup(t *testing.T) {
packages/checks/evals/doctor_test.go:494:func Test_CHK_054_MissingTemplatesAndCanaryMaps(t *testing.T) {
packages/checks/evals/doctor_test.go:548:func Test_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_ZeroHookAndRecovery(t *testing.T) {
packages/checks/evals/doctor_test.go:584:func Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific(t *testing.T) {
packages/checks/evals/doctor_test.go:631:func Test_CHK_055_BaselineAndRecoveryFailures(t *testing.T) {
packages/checks/evals/doctor_test.go:672:func Test_CHK_056_CHK_057_BoundaryViolationMustNameCanaryModule(t *testing.T) {
packages/checks/evals/doctor_test.go:686:func Test_CHK_057_InternalEvalRequiresInterfaceViolation(t *testing.T) {
packages/checks/evals/doctor_test.go:711:func Test_CHK_055_CHK_056_CHK_057_MissingRequiredBoundaryHook(t *testing.T) {
packages/checks/evals/doctor_test.go:737:func Test_CHK_055_MissingRequiredHookBesideOverride(t *testing.T) {
packages/checks/evals/doctor_test.go:758:func Test_CHK_055_ApplicableOverrideNeedsNoGlobalHook(t *testing.T) {
packages/checks/evals/doctor_test.go:770:func Test_CHK_060_CHK_062_ModuleOverridesAndBoundaryWarnings(t *testing.T) {
packages/checks/evals/doctor_test.go:790:func Test_CHK_060_OverridesMustDetectTheirCanaries(t *testing.T) {
packages/checks/evals/doctor_test.go:827:func Test_CHK_063_CHK_064_CHK_065_CHK_066_CHK_067_AgentProbes(t *testing.T) {
packages/checks/evals/doctor_test.go:948:func Test_CHK_065_UnsetTimeoutBoundsAuditProbe(t *testing.T) {
packages/checks/evals/doctor_test.go:970:func Test_CHK_068_CHK_069_ManagedGuidanceAndPlaybookBase(t *testing.T) {

codex
The JUnit, evaluation, zone, and boundary implementations are in place. I’m moving on to recorded Git history and Doctor’s scratch probes, where phase transitions and cleanup need the most care.
exec
/bin/zsh -lc "cat > packages/checks/internal/engine/history.go <<'EOF'
package engine

import (
 \"context\"
 \"fmt\"
 \"os/exec\"
 \"path\"
 \"strings\"

 \"github.com/lurkshark/fireweed/packages/checks/model\"
 \"github.com/lurkshark/fireweed/packages/contracts\"
 \"github.com/lurkshark/fireweed/packages/project\"
 \"github.com/lurkshark/fireweed/packages/gitstore\"
)
func git(root string,args ...string)(string,error){c:=exec.CommandContext(context.Background(),\"git\",append([]string{\"-C\",root},args...)...);b,e:=c.CombinedOutput();if e"'!=nil{return "",fmt.Errorf("git %v: %w: %s",args,e,b)};return strings.TrimSpace(string(b)),nil}
func historyChanges(root,rev string)(map[string]string,error){out:=map[string]string{};names,e:=git(root,"ls-tree","-r","--name-only",rev,"--","changes");if e!=nil{return nil,e};for _,p:=range strings.Split(names,"'"\\n\"){if p==\"\"||"'!strings.HasSuffix(p,".md"){continue};b,e:=gitstore.ReadAt(root,rev,p);if e!=nil{return nil,e};d,e:=contracts.ParseDocument(string(b));if e!=nil{return nil,fmt.Errorf("%s at %s: %w",p,rev,e)};out[p]=str(d.Frontmatter["phase"])};return out,nil}
type hchange struct{status,old,new string}
func historyDiff(root,from,to string)([]hchange,error){args:=[]string{"diff","--name-status","-M","--no-ext-diff"};if from==""{args=append(args,emptyTree)}else{args=append(args,from)};args=append(args,to);s,e:=git(root,args...);if e!=nil{return nil,e};var out []hchange;for _,line:=range strings.Split(s,"'"\\n\"){if line==\"\"{continue};f:=strings.Split(line,\"\\t\");if len(f)<2{continue};c:=hchange{status:f[0],new:f[1]};if strings.HasPrefix(f[0],\"R\")||strings.HasPrefix(f[0],\"C\"){if len(f)<3{continue};c.old=f[1];c.new=f[2]};out=append(out,c)};return out,nil}
const emptyTree=\"4b825dc642cb6eb9a060e54bf8d69288fbee4904\"
func recordedZone(p string, snaps ...project.Snapshot)string{for _,s:=range snaps{for _,m:=range s.Modules{if strings.HasPrefix(p,m.Dir+\"/\"){z:=contracts.ResolveZone(strings.TrimPrefix(p,m.Dir+\"/\"),m.Zones);if z"'!="other"{return z}}}};return "other"}
func historyProtected(p string,snaps ...project.Snapshot)bool{if strings.HasPrefix(p,".fireweed/playbooks/"){return true};for _,s:=range snaps{if matchAny(strSlice(s.Config["protected"]),p){return true};for _,m:=range s.Modules{if strings.HasPrefix(p,m.Dir+"/"){rel:=strings.TrimPrefix(p,m.Dir+"/");if rel=="SPEC.md"||matchAny(m.Zones.Protected,rel){return true}}}};return false}
func History(root,revisionRange string)(model.HistoryReport,error){var rep model.HistoryReport;if revisionRange==""{s,e:=project.Load(root);if e!=nil{return rep,e};base:=str(s.Config["base"]);remote:=str(s.Config["remote"]);ref,e:=gitstore.BaseRef(root,remote,base);if e!=nil{return rep,e};merge,e:=git(root,"merge-base","HEAD",ref);if e!=nil{return rep,e};revisionRange=merge+"..HEAD"};list,e:=git(root,"rev-list","--topo-order","--reverse",revisionRange);if e!=nil{return rep,e};if list==""{return rep,nil};rep.Commits=strings.Split(list,"'"\\n\");for _,oid:=range rep.Commits{after,e:=project.LoadAt(root,oid);if e"'!=nil{return rep,e};afterChanges,e:=historyChanges(root,oid);if e!=nil{return rep,e};parentText,e:=git(root,"rev-list","--parents","-n","1",oid);if e!=nil{return rep,e};parents:=strings.Fields(parentText);if len(parents)>0{parents=parents[1:]};var before []project.Snapshot;var beforeChanges []map[string]string;for _,parent:=range parents{s,e:=project.LoadAt(root,parent);if e!=nil{return rep,e};before=append(before,s);m,e:=historyChanges(root,parent);if e!=nil{return rep,e};beforeChanges=append(beforeChanges,m)};var diffs []hchange;if len(parents)==0{diffs,e=historyDiff(root,"",oid);if e!=nil{return rep,e}}else{seen:=map[string]bool{};for _,parent:=range parents{ds,e:=historyDiff(root,parent,oid);if e!=nil{return rep,e};for _,d:=range ds{k:=d.status+"|"+d.old+"|"+d.new;if !seen[k]{diffs=append(diffs,d);seen[k]=true}}}};worker:=false;implApplied:=false;squashed:=false;for p,phase:=range afterChanges{if phase==""{continue};oldPhase:="";held:=false;for _,m:=range beforeChanges{if x,ok:=m[p];ok{oldPhase=x;held=held||x==phase};for old,x:=range m{if path.Base(old)==path.Base(p){if oldPhase==""{oldPhase=x};held=held||x==phase}}};if held{continue};worker=true;if phase=="applied"{if oldPhase=="impl"&&len(parents)==1{implApplied=true}else{rep.Squashed=append(rep.Squashed,model.SquashedChange{Commit:oid,Change:p});squashed=true}}};snaps:=append([]project.Snapshot{after},before...);add:=func(code,p string){rep.Diagnostics=append(rep.Diagnostics,project.Diagnostic{Code:code,Path:p,Message:fmt.Sprintf("%s at %s",code,oid)})};seenFind:=map[string]bool{};for _,d:=range diffs{paths:=[]string{d.new};if d.old!=""{paths=append(paths,d.old)};for _,p:=range paths{if p==""{continue};if worker&&!squashed&&historyProtected(p,snaps...){k:="H01|"+p;if !seenFind[k]{add("H01",p);seenFind[k]=true}};if implApplied {z:=recordedZone(p,snaps...);if z=="evals"||z=="interface"{k:="H02|"+p;if !seenFind[k]{add("H02",p);seenFind[k]=true}}};if recordedZone(p,snaps...)=="append_only" {existing:=false;for _,parent:=range parents{if _,e:=gitstore.ReadAt(root,parent,p);e==nil{existing=true;break}};if existing {k:="H03|"+p;if !seenFind[k]{add("H03",p);seenFind[k]=true}}};if strings.HasPrefix(p,"changes/archive/")&&d.status!="D"{existing:=false;for _,parent:=range parents{if _,e:=gitstore.ReadAt(root,parent,p);e==nil{existing=true;break}};if existing{k:="H04|"+p;if !seenFind[k]{add("H04",p);seenFind[k]=true}}}}}};return rep,nil}
EOF
gofmt -w packages/checks/internal/engine/history.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 531ms:
exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/checks/internal/engine/core.go'
s=open(p).read().replace('\"time\"','\"time\"\\n \"syscall\"')
a=s.index('func run(ctx context.Context')
b=s.index('\\ntype junitNode',a)
s=s[:a]+'''func run(ctx context.Context, dir string, argv []string, env map[string]string, timeout time.Duration) (string, string, int, error) {
 if timeout>0 {var cancel context.CancelFunc;ctx,cancel=context.WithTimeout(ctx,timeout);defer cancel()}
 c:=exec.Command(argv[0],argv[1:]...);c.Dir=dir;c.Env=os.Environ();for k,v:=range env{c.Env=append(c.Env,k+\"=\"+v)}
 c.SysProcAttr=&syscall.SysProcAttr{Setpgid:true}
 var out,errout strings.Builder;c.Stdout=&out;c.Stderr=&errout
 if e:=c.Start();e"'!=nil{return "","",-1,e}
 done:=make(chan error,1);go func(){done<-c.Wait()}()
 var e error;select{case e=<-done:case <-ctx.Done():_ = syscall.Kill(-c.Process.Pid,syscall.SIGKILL);e=<-done;_ = e;return out.String(),errout.String(),-1,ctx.Err()}
 exit:=0;if c.ProcessState!=nil{exit=c.ProcessState.ExitCode()};if e!=nil {if _,ok:=e.(*exec.ExitError);!ok{return out.String(),errout.String(),exit,e}}
 return out.String(),errout.String(),exit,nil
}
'"''' +s[b:]
open(p,'w').write(s)
PY
gofmt -w packages/checks/internal/engine/core.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 601ms:
exec
/bin/zsh -lc "sed -n '770,1050p' packages/checks/evals/doctor_test.go && sed -n '1,110p' packages/checks/evals/doctor_agent_logs_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 493ms:
func Test_CHK_060_CHK_062_ModuleOverridesAndBoundaryWarnings(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/override-eval.sh", fmt.Sprintf("#!/bin/sh\nprintf 'override-eval:%%s\\n' \"$PWD\" >> %q\nexec sh .fireweed/eval.sh \"$@\"\n", f.log))
	f.put(".fireweed/override-boundary.sh", fmt.Sprintf("#!/bin/sh\nprintf 'override-boundaries:%%s:%%s\\n' \"$PWD\" \"$1\" >> %q\nexec sh .fireweed/boundary.sh\n", f.log))
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\nhooks:\n  eval: sh .fireweed/override-eval.sh {module} {module_dir} {report}\n  boundaries: sh .fireweed/override-boundary.sh {module}\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nThe fixture MUST pass.\n")
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) || strings.Count(doctorLog(f), "override-eval:") < 2 || strings.Count(doctorLog(f), "override-boundaries:") < 2 {
		t.Errorf("override probes absent: %s", doctorLog(f))
	}
	g := localDoctorFixture(t)
	g.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	g.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(g.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nboundaries_opt_out: legacy", 1))
	g.commit()
	r := checks.Doctor(context.Background(), g.request(true))
	if doctorErrors(r) || !hasDoctorDiagnostic(r, "X03", true, "legacy") || !hasDoctorDiagnostic(r, "X04", true, "no boundary enforcement required") {
		t.Errorf("optional boundaries warning lost: %s", doctorText(r))
	}
}

func Test_CHK_060_OverridesMustDetectTheirCanaries(t *testing.T) {
	for _, tc := range []struct {
		name, hook, command, script, canary string
	}{
		{"eval", "eval", "sh .fireweed/override-eval.sh {module} {module_dir} {report}", `#!/bin/sh
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "failing_eval"},
		{"boundary", "boundaries", "sh .fireweed/override-boundary.sh {module}", "#!/bin/sh\nexit 0\n", "internal_eval"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			path := ".fireweed/override-" + tc.name + ".sh"
			f.put(path, tc.script)
			f.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(f.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nhooks:\n  "+tc.hook+": "+tc.command, 1))
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			diagnosed := false
			for _, d := range r.Diagnostics {
				message := strings.ToLower(d.Message)
				boundaryHook := tc.name == "boundary" && !strings.Contains(message, "global") &&
					(strings.Contains(message, "alpha boundary hook") ||
						strings.Contains(message, "alpha boundaries hook") ||
						strings.Contains(message, "alpha boundaries override boundary hook") ||
						strings.Contains(message, "alpha boundaries override boundaries hook"))
				if d.Code == "C14" && !d.Warning && strings.Contains(message, tc.canary) &&
					(tc.name == "eval" && strings.Contains(message, "alpha "+tc.hook+" override") || boundaryHook) {
					diagnosed = true
				}
			}
			if !diagnosed {
				t.Errorf("broken %s override missed its %s canary without C14: %s", tc.name, tc.canary, doctorText(r))
			}
		})
	}
}

func Test_CHK_063_CHK_064_CHK_065_CHK_066_CHK_067_AgentProbes(t *testing.T) {
	f := localDoctorFixture(t)
	f.commit()
	for _, d := range project.Check(f.root) {
		if !d.Warning {
			t.Fatalf("invalid agent fixture: %+v", d)
		}
	}
	r := checks.Doctor(context.Background(), f.request(false))
	log := doctorLog(f)
	if doctorErrors(r) || strings.Count(log, "worker:") != 1 || strings.Count(log, "audit:") != 1 || strings.Count(log, "wrap:") != 2 || !strings.Contains(log, "FIREWEED_DOCTOR_OK") {
		t.Errorf("worker/audit/wrap probes: %s; log=%s", doctorText(r), log)
	}
	workerPrompt, auditPrompt := false, false
	for _, line := range strings.Split(log, "\n") {
		if fields := strings.SplitN(line, ":", 3); len(fields) == 3 && strings.Contains(fields[2], "FIREWEED_DOCTOR_OK") {
			switch fields[0] {
			case "worker":
				workerPrompt = true
			case "audit":
				auditPrompt = true
			}
		}
	}
	if !workerPrompt {
		t.Errorf("worker did not receive its fixed-reply prompt: %s", log)
	}
	if !auditPrompt {
		t.Errorf("audit did not receive its fixed-reply prompt: %s", log)
	}
	for _, part := range []string{"/fireweed/worktrees/", "worker:", "audit:"} {
		if !strings.Contains(log, part) {
			t.Errorf("scratch/role marker %s absent: %s", part, log)
		}
	}

	g := localDoctorFixture(t)
	g.editConfig(func(config map[string]any) {
		delete(config["audit"].(map[string]any), "command")
		config["agent"].(map[string]any)["command"] = "bash .fireweed/agent.sh {prompt} {git_common_dir}"
	})
	g.commit()
	rr := checks.Doctor(context.Background(), g.request(false))
	if !hasDoctorDiagnostic(rr, "C15", true, "shared") || strings.Count(doctorLog(g), "worker:") != 2 {
		t.Errorf("audit fallback: %s; log=%s", doctorText(rr), doctorLog(g))
	}

	h := localDoctorFixture(t)
	h.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/bash\nfor i in $(seq 1 25); do printf 'line-%%02d\\n' \"$i\"; done\necho RATE LIMIT reached\nexit 9\n"))
	h.commit()
	bad := checks.Doctor(context.Background(), h.request(false))
	if !hasDoctorDiagnostic(bad, "C15", false, "usage limit", "rate limit", "line-25") {
		t.Errorf("usage-limit tail missing: %s", doctorText(bad))
	}
	if strings.Contains(doctorText(bad), "line-01") {
		t.Errorf("agent output was not trimmed to final 20 lines: %s", doctorText(bad))
	}
	outOfTail := localDoctorFixture(t)
	outOfTail.put(".fireweed/agent.sh", "#!/bin/bash\necho RATE LIMIT early\nfor i in $(seq 1 25); do printf 'later-%02d\\n' \"$i\"; done\nexit 9\n")
	outOfTail.commit()
	ordinary := checks.Doctor(context.Background(), outOfTail.request(false))
	if !hasDoctorDiagnostic(ordinary, "C15", false, "agent.sh", "later-25") {
		t.Errorf("ordinary out-of-tail failure absent: %s", doctorText(ordinary))
	}
	for _, d := range ordinary.Diagnostics {
		if d.Code == "C15" && !d.Warning && strings.Contains(strings.ToLower(d.Message), "usage limit") {
			t.Errorf("out-of-tail pattern classified as usage limit: %s", doctorText(ordinary))
		}
	}
	for _, tc := range []struct{ name, script string }{{"wrong reply", "#!/bin/bash\necho WRONG_REPLY\n"}, {"nonzero with reply", "#!/bin/bash\necho FIREWEED_DOCTOR_OK\nexit 7\n"}, {"ordinary failure", "#!/bin/bash\nfor i in $(seq 1 25); do printf 'ordinary-%02d\\n' \"$i\"; done\nexit 8\n"}} {
		t.Run(tc.name, func(t *testing.T) {
			q := localDoctorFixture(t)
			q.put(".fireweed/agent.sh", tc.script)
			q.commit()
			r := checks.Doctor(context.Background(), q.request(false))
			if !hasDoctorDiagnostic(r, "C15", false, "agent.sh") {
				t.Errorf("agent failure accepted: %s", doctorText(r))
			}
			if tc.name == "ordinary failure" && (!hasDoctorDiagnostic(r, "C15", false, "ordinary-25") || strings.Contains(doctorText(r), "ordinary-01")) {
				t.Errorf("ordinary failure tail incorrect: %s", doctorText(r))
			}
		})
	}
	wrongAudit := localDoctorFixture(t)
	wrongAudit.put(".fireweed/audit.sh", "#!/bin/bash\necho WRONG_REPLY\n")
	wrongAudit.commit()
	auditFailure := checks.Doctor(context.Background(), wrongAudit.request(false))
	if !hasDoctorDiagnostic(auditFailure, "C15", false, "audit.sh") || hasDoctorCheck(auditFailure, "audit", true) {
		t.Errorf("wrong audit reply accepted: %s", doctorText(auditFailure))
	}

	i := localDoctorFixture(t)
	i.put(".fireweed/audit.sh", "#!/bin/bash\nsleep 2\n")
	i.commit()
	req := i.request(false)
	req.Timeout = 100 * time.Millisecond
	start := time.Now()
	timed := checks.Doctor(context.Background(), req)
	if !hasDoctorDiagnostic(timed, "C15", false, "audit") || time.Since(start) > 4*time.Second {
		t.Errorf("audit timeout not bounded: %s; elapsed=%s", doctorText(timed), time.Since(start))
	}

	defaulted := localDoctorFixture(t)
	defaulted.put(".fireweed/agent.sh", "#!/bin/bash\nsleep 0.2\nprintf 'FIREWEED_DOCTOR_OK\\n'\n")
	defaulted.put(".fireweed/audit.sh", "#!/bin/bash\nsleep 0.2\nprintf 'FIREWEED_DOCTOR_OK\\n'\n")
	defaulted.commit()
	defaultRequest := defaulted.request(false)
	defaultRequest.Timeout = 0
	defaultReport := checks.Doctor(context.Background(), defaultRequest)
	if doctorErrors(defaultReport) || !hasDoctorCheck(defaultReport, "worker", true) || !hasDoctorCheck(defaultReport, "audit", true) {
		t.Errorf("unset timeout did not give both agent probes their default allowance: %s", doctorText(defaultReport))
	}

	j := localDoctorFixture(t)
	j.commit()
	skip := checks.Doctor(context.Background(), j.request(true))
	if strings.Contains(doctorLog(j), "worker:") || strings.Contains(doctorLog(j), "audit:") || !hasDoctorCheck(skip, "undeclared", true) || !strings.Contains(doctorLog(j), j.bin) {
		t.Errorf("NoAgent skipped hooks or ran agent: %s; log=%s", doctorText(skip), doctorLog(j))
	}
}

func Test_CHK_065_UnsetTimeoutBoundsAuditProbe(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/audit.sh", fmt.Sprintf("#!/bin/bash\nprintf 'audit:%%s:%%s\n' \"$PWD\" \"$1\" >> %q\nprintf 'audit-timeout-started\n'\nsleep 210\nprintf 'FIREWEED_DOCTOR_OK\n'\n", f.log))
	f.commit()
	req := f.request(false)
	req.Timeout = 0

	start := time.Now()
	report := checks.Doctor(context.Background(), req)
	elapsed := time.Since(start)
	log := doctorLog(f)
	if !hasDoctorCheck(report, "worker", true) || strings.Count(log, "worker:") != 1 || strings.Count(log, "audit:") != 1 {
		t.Errorf("agent probes did not run once with immediate worker reply: %s; log=%s", doctorText(report), log)
	}
	if !hasDoctorCheck(report, "audit", false) || !hasDoctorDiagnostic(report, "C15", false, "audit.sh", "audit-timeout-started") {
		t.Errorf("long audit probe lacked a C15 timeout with its command and output: %s", doctorText(report))
	}
	if elapsed < 165*time.Second || elapsed > 200*time.Second {
		t.Errorf("unset audit timeout took %s, want a bound near three minutes", elapsed)
	}
}

func Test_CHK_068_CHK_069_ManagedGuidanceAndPlaybookBase(t *testing.T) {
	f := localDoctorFixture(t)
	text, err := assets.Playbook("evals")
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256([]byte(text))
	base := hex.EncodeToString(sum[:])
	f.put(".fireweed/playbooks/evals.md", "---\nembedded_base: "+base+"\n---\n# Local override\n")
	f.commit()
	r := checks.Doctor(context.Background(), f.request(true))
	if hasDoctorDiagnostic(r, "C16", true) {
		t.Errorf("current guidance/base warned: %s", doctorText(r))
	}
	f.put("AGENTS.md", "Before.\n"+doctorGuidance+"After.\n")
	surrounded := checks.Doctor(context.Background(), f.request(true))
	if hasDoctorDiagnostic(surrounded, "C16", true) {
		t.Errorf("current managed guidance with surrounding text warned: %s", doctorText(surrounded))
	}
	f.put("AGENTS.md", "Before.\n<!-- fireweed:start -->\nold guidance\n<!-- fireweed:end -->\nAfter.\n")
	f.put(".fireweed/playbooks/evals.md", "---\nembedded_base: stale\n---\n# Local override\n")
	stale := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(stale, "C16", true, "AGENTS.md") || !hasDoctorDiagnostic(stale, "C16", true, "evals") {
		t.Errorf("stale guidance/base not warned: %s", doctorText(stale))
	}
	if err := os.Remove(filepath.Join(f.root, "AGENTS.md")); err != nil {
		t.Fatal(err)
	}
	f.put(".fireweed/playbooks/evals.md", "# No embedded base\n")
	missing := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(missing, "C16", true, "AGENTS.md") || !hasDoctorDiagnostic(missing, "C16", true, "evals") {
		t.Errorf("missing guidance/base not warned: %s", doctorText(missing))
	}
}

func mustRead(t *testing.T, path string) []byte {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return b
}
func doctorBinary(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("caller")
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(file), "..", "..", ".."))
	bin := filepath.Join(t.TempDir(), "fireweed-new")
	cmd := exec.Command("go", "build", "-o", bin, ".")
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GOCACHE=/tmp/fireweed-rebuild-cache")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build product: %v: %s", err, out)
	}
	return bin
}
package evals_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
)

func doctorProbeLog(t *testing.T, f doctorFixture, report checks.DoctorReport, role string, markers ...string) string {
	t.Helper()
	var detail string
	for _, check := range report.Checks {
		if check.Name == role {
			detail = check.Detail
			break
		}
	}
	if detail == "" {
		t.Errorf("executed %s probe has no check: %s", role, doctorText(report))
		return ""
	}
	at := strings.Index(detail, "log: ")
	if at < 0 {
		t.Errorf("%s detail lacks log path: %q", role, detail)
		return ""
	}
	fields := strings.Fields(detail[at+len("log: "):])
	if len(fields) == 0 {
		t.Errorf("%s detail has empty log path: %q", role, detail)
		return ""
	}
	path := fields[0]
	if !filepath.IsAbs(path) {
		t.Errorf("%s log path is not absolute: %q", role, path)
	}
	rel, err := filepath.Rel(filepath.Join(f.root, ".fireweed", "runs"), path)
	if err != nil {
		t.Errorf("%s log path %q: %v", role, path, err)
		return ""
	}
	parts := strings.Split(filepath.ToSlash(rel), "/")
	if len(parts) != 2 || parts[0] == "" || parts[0] == "." || parts[0] == ".." || parts[1] != "doctor-"+role+"-1.log" {
		t.Errorf("%s log path = %q, want .fireweed/runs/<run>/doctor-%s-1.log", role, path, role)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("read %s log %q: %v", role, path, err)
		return path
	}
	for _, marker := range markers {
		if !strings.Contains(string(data), marker) {
			t.Errorf("%s log lacks %q: %q", role, marker, data)
		}
	}
	return path
}

func doctorLogRun(t *testing.T, f doctorFixture, worker, audit string, report checks.DoctorReport) string {
	t.Helper()
	base := filepath.Join(f.root, ".fireweed", "runs")
	w, ew := filepath.Rel(base, worker)
	a, ea := filepath.Rel(base, audit)
	if ew != nil || ea != nil {
		t.Errorf("log paths outside run root: %q %q: %v %v", worker, audit, ew, ea)
		return ""
	}
	wRun, aRun := filepath.Dir(w), filepath.Dir(a)
	if wRun != aRun || wRun == "." || wRun == ".." {
		t.Errorf("worker and audit have different run identities: %q %q: %s", worker, audit, doctorText(report))
	}
	return wRun
}

func doctorCallerClean(t *testing.T, f doctorFixture, status, registrations string) {
	t.Helper()
	if got := f.git("status", "--porcelain", "--untracked-files=all"); got != status {
		t.Errorf("caller status changed: before %q after %q", status, got)
	}
	if got := worktrees(f); got != registrations {
		t.Errorf("scratch registration changed: before %q after %q", registrations, got)
	}
	assertScratchRemoved(t, f)
}

func Test_CHK_081_DoctorProbeLogsShareFreshRunAndPreserveStreams(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/agent.sh", "#!/bin/bash\nprintf 'worker stdout Ω without newline'; printf 'worker stderr β without newline' >&2; printf 'FIREWEED_DOCTOR_OK\\n'\n")
	f.put(".fireweed/audit.sh", "#!/bin/bash\nprintf 'audit stdout Ω without newline'; printf 'audit stderr β without newline' >&2; printf 'FIREWEED_DOCTOR_OK\\n'\n")
	f.commit()
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	var firstRun string
	for i := 0; i < 2; i++ {
		report := checks.Doctor(context.Background(), f.request(false))
		if doctorErrors(report) || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
			t.Errorf("successful probes changed outcome: %s", doctorText(report))
		}
		worker := doctorProbeLog(t, f, report, "worker", "worker stdout Ω without newlineFIREWEED_DOCTOR_OK\n", "worker stderr β without newline")
		audit := doctorProbeLog(t, f, report, "audit", "audit stdout Ω without newlineFIREWEED_DOCTOR_OK\n", "audit stderr β without newline")
		if worker != "" && audit != "" {
			run := doctorLogRun(t, f, worker, audit, report)
			if i == 1 && run == firstRun {
				t.Errorf("consecutive Doctor invocations reused run identity %q", run)
			}
			firstRun = run

exec
/bin/zsh -lc "cat > packages/checks/internal/engine/doctor.go <<'EOF'
package engine

import (
 \"context\"
 \"crypto/sha256\"
 \"encoding/hex\"
 \"fmt\"
 \"io\"
 \"os\"
 \"path/filepath\"
 \"sort\"
 \"strings\"
 \"time\"

 \"github.com/lurkshark/fireweed/packages/assets\"
 \"github.com/lurkshark/fireweed/packages/checks/model\"
 \"github.com/lurkshark/fireweed/packages/contracts\"
 \"github.com/lurkshark/fireweed/packages/gitstore\"
 \"github.com/lurkshark/fireweed/packages/project\"
 \"gopkg.in/yaml.v3\"
)
const doctorPrompt=\"Reply with exactly FIREWEED_DOCTOR_OK\"
const doctorGuidance=\"<"'!-- fireweed:start -->'"\\nThis project uses fireweed. Specs and change files are the source of truth.\\nRun "'`fireweed playbook author-change` before authoring changes and'"\\n"'`fireweed context <change>` before each step.'"\\n<"'!-- fireweed:end -->'"\\n\"
func doctorDiag(r *model.DoctorReport,code,msg string,warning bool){r.Diagnostics=append(r.Diagnostics,project.Diagnostic{Code:code,Message:msg,Warning:warning})}
func doctorCheck(r *model.DoctorReport,name,module string,pass bool,detail string){r.Checks=append(r.Checks,model.DoctorCheck{Name:name,Module:module,Passed:pass,Detail:detail})}
func staticErr(d []project.Diagnostic)bool{for _,x:=range d{if "'!x.Warning{return true}};return false}
func copyTree(from,to string)error{return filepath.Walk(from,func(p string,info os.FileInfo,e error)error{if e!=nil{return e};rel,e:=filepath.Rel(from,p);if e!=nil{return e};if rel=="."{return nil};if rel==".git"||strings.HasPrefix(rel,".git"+string(os.PathSeparator))||rel==filepath.Join(".fireweed","runs")||strings.HasPrefix(rel,filepath.Join(".fireweed","runs")+string(os.PathSeparator)){if info.IsDir(){return filepath.SkipDir};return nil};dest:=filepath.Join(to,rel);if info.IsDir(){return os.MkdirAll(dest,info.Mode().Perm())};if !info.Mode().IsRegular(){return nil};if e:=os.MkdirAll(filepath.Dir(dest),0755);e!=nil{return e};src,e:=os.Open(p);if e!=nil{return e};defer src.Close();dst,e:=os.OpenFile(dest,os.O_CREATE|os.O_TRUNC|os.O_WRONLY,info.Mode().Perm());if e!=nil{return e};_,e=io.Copy(dst,src);closeErr:=dst.Close();if e!=nil{return e};return closeErr})}
func doctorWarnings(root string,r *model.DoctorReport){b,e:=os.ReadFile(filepath.Join(root,"AGENTS.md"));if e!=nil||!strings.Contains(string(b),doctorGuidance){doctorDiag(r,"C16","AGENTS.md managed guidance missing or stale",true)};files,_:=filepath.Glob(filepath.Join(root,".fireweed","playbooks","*.md"));for _,p:=range files{b,e:=os.ReadFile(p);if e!=nil{continue};d,e:=contracts.ParseDocument(string(b));name:=strings.TrimSuffix(filepath.Base(p),".md");text,e2:=assets.Playbook(name);if e!=nil||e2!=nil{doctorDiag(r,"C16",name+" playbook embedded_base missing or stale",true);continue};sum:=sha256.Sum256([]byte(text));if str(d.Frontmatter["embedded_base"])!=hex.EncodeToString(sum[:]){doctorDiag(r,"C16",name+" playbook embedded_base missing or stale",true)}}}
func renderCanary(text string,vars map[string]string)string{for _,key:=range []string{"name","prefix","package","go_module","module_dir","other_dir"}{text=strings.ReplaceAll(text,"{{"+key+"}}",vars[key])};return text}
func writeMap(root,dir string,m map[string]string,vars map[string]string)error{keys:=make([]string,0,len(m));for k:=range m{keys=append(keys,k)};sort.Strings(keys);for _,k:=range keys{p:=filepath.Join(root,dir,strings.TrimSuffix(k,".tmpl"));if e:=os.MkdirAll(filepath.Dir(p),0755);e!=nil{return e};if e:=os.WriteFile(p,[]byte(renderCanary(m[k],vars)),0644);e!=nil{return e}};return nil}
func templateFiles(root string)(map[string]string,error){base:=filepath.Join(root,".fireweed","templates","module");out:=map[string]string{};e:=filepath.Walk(base,func(p string,info os.FileInfo,e error)error{if e!=nil{return e};if info.IsDir(){return nil};rel,e:=filepath.Rel(base,p);if e!=nil{return e};b,e:=os.ReadFile(p);if e!=nil{return e};out[filepath.ToSlash(rel)]=string(b);return nil});if e!=nil{return nil,e};if _,ok:=out["SPEC.md"];!ok{return nil,fmt.Errorf("template SPEC.md missing")};return out,nil}
func canaryMaps(root string)(map[string]map[string]string,error){p:=filepath.Join(root,".fireweed","canaries","canaries.yaml");b,e:=os.ReadFile(p);if e!=nil{return nil,fmt.Errorf("canary %s: %w",p,e)};var raw struct{Version int `yaml:"version"`;Baseline map[string]string `yaml:"baseline"`;Undeclared map[string]string `yaml:"undeclared"`;InternalEval map[string]string `yaml:"internal_eval"`;FailingEval map[string]string `yaml:"failing_eval"`;MissingID map[string]string `yaml:"missing_id"`};if e:=yaml.Unmarshal(b,&raw);e!=nil{return nil,fmt.Errorf("canary canaries.yaml: %w",e)};if raw.Version!=1{return nil,fmt.Errorf("canary canaries.yaml version %d",raw.Version)};out:=map[string]map[string]string{"baseline":raw.Baseline,"undeclared":raw.Undeclared,"internal_eval":raw.InternalEval,"failing_eval":raw.FailingEval,"missing_id":raw.MissingID};for _,name:=range []string{"baseline","undeclared","internal_eval","failing_eval","missing_id"}{if len(out[name])==0{return nil,fmt.Errorf("canary %s missing",name)}};return out,nil}
func evalProbe(ctx context.Context,s project.Snapshot,m project.Module,hook any,bin string,timeout time.Duration)(model.Report,error){copy:=s;copy.Modules=[]project.Module{m};if hook!=nil{copy.Config=map[string]any{"hooks":map[string]any{"eval":hook}};m.Document.Frontmatter=map[string]any{};copy.Modules[0]=m};return Evaluate(ctx,model.HookRequest{Snapshot:copy,Module:m.Name,Bin:bin,Timeout:timeout})}
func boundaryProbe(ctx context.Context,s project.Snapshot,m project.Module,hook any,bin string,timeout time.Duration)([]project.Diagnostic,error){copy:=s;copy.Modules=[]project.Module{m};if hook!=nil{copy.Config=map[string]any{"hooks":map[string]any{"boundaries":hook}};m.Document.Frontmatter=map[string]any{};copy.Modules[0]=m};return Boundaries(ctx,model.BoundaryRequest{Snapshot:copy,Bin:bin,Timeout:timeout})}
func doctorBaseline(ctx context.Context,s project.Snapshot,m project.Module,evalHook,boundHook any,bin string,timeout time.Duration)error{rep,e:=evalProbe(ctx,s,m,evalHook,bin,timeout);if e!=nil{return fmt.Errorf("eval hook: %w",e)};specs,e:=moduleSpecs(m);if e!=nil{return fmt.Errorf("eval hook specs: %w",e)};if rep.HookExit!=0{return fmt.Errorf("eval hook exit %d",rep.HookExit)};if e=RequireCoverage(rep,specs);e!=nil{return fmt.Errorf("eval hook coverage: %w",e)};if e=RequirePass(rep,specs);e!=nil{return fmt.Errorf("eval hook pass: %w",e)};ds,e:=boundaryProbe(ctx,s,m,boundHook,bin,timeout);if e!=nil{return fmt.Errorf("boundaries hook: %w (%v)",e,ds)};return nil}
func doctorCanary(ctx context.Context,r *model.DoctorReport,s project.Snapshot,m project.Module,hookName string,evalHook,boundHook any,vars map[string]string,files map[string]string,bin string,timeout time.Duration,name string){if err:=writeMap(s.Root,m.Dir,files,vars);err!=nil{doctorDiag(r,"C14",name+" "+hookName+" canary write: "+err.Error(),false);return};pass:=false;detail:="";switch name{case "undeclared","internal_eval":ds,e:=boundaryProbe(ctx,s,m,boundHook,bin,timeout);if e==nil{for _,d:=range ds{if d.Code=="X01"&&strings.HasPrefix(d.Message,m.Name+":")&&(name!="internal_eval"||strings.Contains(strings.ToLower(d.Message),"interface")){pass=true}}};detail=fmt.Sprintf("%v %v",ds,e);case "failing_eval":rep,e:=evalProbe(ctx,s,m,evalHook,bin,timeout);if e==nil&&rep.HookExit!=0{for _,c:=range rep.Cases{if c.Result=="fail"&&hasID(c.IDs,m.IDPrefix+"-001"){pass=true}}};detail=fmt.Sprintf("%v %v",rep,e);case "missing_id":_,e:=evalProbe(ctx,s,m,evalHook,bin,timeout);if e!=nil&&strings.Contains(e.Error(),"unknown spec ID")&&strings.Contains(e.Error(),"UNMAPPED"){pass=true};detail=fmt.Sprint(e)};doctorCheck(r,name,m.Name,pass,detail);if !pass{doctorDiag(r,"C14",name+" missed by "+hookName+" for "+m.Name+": "+detail,false)};for p:=range files{_ = os.Remove(filepath.Join(s.Root,m.Dir,p))};if e:=doctorBaseline(ctx,s,m,evalHook,boundHook,bin,timeout);e!=nil{doctorCheck(r,"recovery "+name,m.Name,false,e.Error());doctorDiag(r,"C14","recovery after "+name+" failed "+hookName+": "+e.Error(),false)}else{doctorCheck(r,"recovery "+name,m.Name,true,"")}}
func tail20(s string)string{lines:=strings.Split(strings.TrimRight(s,"'"\\n\"),\"\\n\");if len(lines)>20{lines=lines[len(lines)-20:]};return strings.Join(lines,\"\\n\")}
func doctorAgent(ctx context.Context,r *model.DoctorReport,req model.DoctorRequest,scratch,common,runID,role string,cmdValue,wrapValue any){vars:=map[string]string{\"prompt\":doctorPrompt,\"git_common_dir\":common};needsFile:=strings.Contains(fmt.Sprint(cmdValue),\"{prompt_file}\")||strings.Contains(fmt.Sprint(wrapValue),\"{prompt_file}\");if needsFile{f,e:=os.CreateTemp(\"\",\"fireweed-doctor-prompt-*\");if e"'!=nil{doctorDiag(r,"C15",role+" prompt file: "+e.Error(),false);return};p:=f.Name();_,e=f.WriteString(doctorPrompt);f.Close();if e!=nil{os.Remove(p);doctorDiag(r,"C15",role+" prompt file: "+e.Error(),false);return};vars["prompt_file"]=p;defer os.Remove(p)};cmd,e:=command(cmdValue,vars);if e!=nil{doctorDiag(r,"C15",role+" command: "+e.Error(),false);return};wrap:=[]string{};if wrapValue!=nil{wrap,e=command(wrapValue,vars);if e!=nil{doctorDiag(r,"C15",role+" wrap: "+e.Error(),false);return}};argv:=append(wrap,cmd...);limit:=req.Timeout;if limit<=0{limit=3*time.Minute};out,stderr,exit,runErr:=run(ctx,scratch,argv,nil,limit);logPath:=filepath.Join(req.Root,".fireweed","runs",runID,"doctor-"+role+"-1.log");_ = os.MkdirAll(filepath.Dir(logPath),0755);_ = os.WriteFile(logPath,[]byte(out+stderr),0644);detail:="log: "+logPath;combined:=out+stderr;tail:=tail20(combined);passed:=runErr==nil&&exit==0&&strings.Contains(out,"FIREWEED_DOCTOR_OK");doctorCheck(r,role,"",passed,detail);if !passed{message:=fmt.Sprintf("%s probe command %v failed: %v exit %d'"\\n%s\",role,argv,runErr,exit,tail);for _,pat:=range strSlice(nil){_ = pat};_ = message;doctorDiag(r,\"C15\",message,false)}}
func Doctor(ctx context.Context,req model.DoctorRequest)model.DoctorReport{var r model.DoctorReport;r.Diagnostics=append(r.Diagnostics,project.Check(req.Root)...);doctorWarnings(req.Root,&r);if staticErr(r.Diagnostics){return r};repo,e:=gitstore.Inspect(req.Root);if e"'!=nil{doctorDiag(&r,"C14","scratch repository: "+e.Error(),false);return r};common:=repo.CommonDir;if !filepath.IsAbs(common){common=filepath.Join(req.Root,common)};common,_=filepath.EvalSymlinks(common);base:=filepath.Join(common,"fireweed","worktrees");_ = os.MkdirAll(base,0755);scratch,e:=os.MkdirTemp(base,"doctor-");if e!=nil{doctorDiag(&r,"C14","scratch worktree: "+e.Error(),false);return r};_ = os.Remove(scratch);if _,e=git(req.Root,"worktree","add","--detach",scratch,"HEAD");e!=nil{doctorDiag(&r,"C14","scratch worktree: "+e.Error(),false);return r};defer func(){_,_=git(req.Root,"worktree","remove","--force",scratch);_ = os.RemoveAll(scratch)}();if e=copyTree(req.Root,scratch);e!=nil{doctorDiag(&r,"C14","scratch copy: "+e.Error(),false);return r};snap,e:=project.Load(scratch);if e!=nil{doctorDiag(&r,"C14","scratch project: "+e.Error(),false);return r};setup:=configured(snap.Config,"setup");if setup!=nil{args,e:=command(setup,nil);if e==nil{_,_,exit,runErr:=run(ctx,scratch,args,map[string]string{"FIREWEED_BIN":req.Bin},req.Timeout);if runErr!=nil||exit!=0{e=fmt.Errorf("exit %d: %v",exit,runErr)}};if e!=nil{doctorDiag(&r,"C14","setup hook failed: "+e.Error(),false);return r}};for _,m:=range snap.Modules{if reason:=boundaryOpt(m);reason!=""{doctorDiag(&r,"X03",m.Name+": boundaries opt out: "+reason,true)};if boundaryHook(snap,&m)==nil&&(len(m.Zones.Disposable)==0||reason!=""){doctorDiag(&r,"X04",m.Name+": no boundary enforcement required",true)}};templates,e:=templateFiles(scratch);if e!=nil{doctorDiag(&r,"C14","module template SPEC.md: "+e.Error(),false);return r};maps,e:=canaryMaps(scratch);if e!=nil{doctorDiag(&r,"C14",e.Error(),false);return r};modText:="";if b,e:=os.ReadFile(filepath.Join(scratch,"go.mod"));e==nil{for _,l:=range strings.Split(string(b),"'"\\n\"){if strings.HasPrefix(l,\"module \"){modText=strings.TrimSpace(strings.TrimPrefix(l,\"module \"));break}}};names:=[]string{\"doctorprobeone\",\"doctorprobetwo\"};dirs:=[]string{\"packages/\"+names[0],\"packages/\"+names[1]};var canMods []project.Module;for i,name:=range names{vars:=map[string]string{\"name\":name,\"prefix\":fmt.Sprintf(\"DCT%d\",i+1),\"package\":name,\"go_module\":modText,\"module_dir\":dirs[i],\"other_dir\":dirs[1-i]};if e:=writeMap(scratch,dirs[i],templates,vars);e"'!=nil{doctorDiag(&r,"C14","module template: "+e.Error(),false);return r};if e:=writeMap(scratch,dirs[i],maps["baseline"],vars);e!=nil{doctorDiag(&r,"C14","canary baseline: "+e.Error(),false);return r};b,e:=os.ReadFile(filepath.Join(scratch,dirs[i],"SPEC.md"));if e!=nil{doctorDiag(&r,"C14","module template SPEC.md: "+e.Error(),false);return r};d,e:=contracts.ParseDocument(string(b));if e!=nil{doctorDiag(&r,"C14","module template SPEC.md: "+e.Error(),false);return r};m:=project.Module{Name:name,IDPrefix:vars["prefix"],Dir:dirs[i],Zones:contracts.Zones{Disposable:[]string{"internal/**"},Evals:[]string{"evals/**"},Interface:[]string{"api.txt"}},Document:d};canMods=append(canMods,m)};globalEval:=configured(snap.Config,"eval");globalBoundary:=configured(snap.Config,"boundaries");targets:=[]struct{m project.Module; eval,bound any;label string}{{canMods[0],globalEval,globalBoundary,"global"},{canMods[1],globalEval,globalBoundary,"global"}};for _,m:=range snap.Modules{eh:=configured(m.Document.Frontmatter,"eval");bh:=configured(m.Document.Frontmatter,"boundaries");if eh!=nil||bh!=nil{if eh==nil{eh=globalEval};if bh==nil{bh=globalBoundary};targets=append(targets,struct{m project.Module;eval,bound any;label string}{m,eh,bh,m.Name+" eval override"})}};for _,target:=range targets{m:=target.m;eh,bh:=target.eval,target.bound;if e:=doctorBaseline(ctx,snap,m,eh,bh,req.Bin,req.Timeout);e!=nil{doctorCheck(&r,"baseline",m.Name,false,e.Error());doctorDiag(&r,"C14","baseline "+target.label+" "+e.Error(),false);continue};doctorCheck(&r,"baseline",m.Name,true,"");vars:=map[string]string{"name":m.Name,"prefix":m.IDPrefix,"package":m.Name,"go_module":modText,"module_dir":m.Dir,"other_dir":dirs[0]};if m.Name==names[0]{vars["other_dir"]=dirs[1]};for _,name:=range []string{"undeclared","internal_eval","failing_eval","missing_id"}{hookName:=target.label+" boundary hook";if name=="failing_eval"||name=="missing_id"{hookName=target.label+" eval hook"};doctorCanary(ctx,&r,snap,m,hookName,eh,bh,vars,maps[name],req.Bin,req.Timeout,name)}};if !req.NoAgent{runID:=fmt.Sprintf("%d",time.Now().UnixNano());agent:=stringMap(snap.Config["agent"]);audit:=stringMap(snap.Config["audit"]);wrap:=agent["wrap"];doctorAgent(ctx,&r,req,scratch,common,runID,"worker",agent["command"],wrap);auditCmd:=audit["command"];if auditCmd==nil{auditCmd=agent["command"];doctorDiag(&r,"C15","audit shares worker command; audit command blind spot",true)};doctorAgent(ctx,&r,req,scratch,common,runID,"audit",auditCmd,wrap)};return r}
EOF
gofmt -w packages/checks/internal/engine/doctor.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 607ms:
exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/checks/internal/engine/doctor.go'
s=open(p).read()
s=s.replace('runID, role string, cmdValue, wrapValue any)', 'runID, role string, cmdValue, wrapValue any, patterns []string)')
s=s.replace('for _, pat := range strSlice(nil) {\\n\\t\\t\\t_ = pat\\n\\t\\t}\\n\\t\\t_ = message\\n\\t\\tdoctorDiag(r, \"C15\", message, false)', 'for _, pat := range patterns {\\n            if pat "'!= "" && strings.Contains(strings.ToLower(tail), strings.ToLower(pat)) { message = "usage limit: " + message; break }'"\\n        }\\n        doctorDiag(r, \"C15\", message, false)')
s=s.replace('agent[\"command\"], wrap)', 'agent[\"command\"], wrap, strSlice(snap.Config[\"limit_patterns\"]))')
s=s.replace('auditCmd, wrap)', 'auditCmd, wrap, strSlice(snap.Config[\"limit_patterns\"]))')
open(p,'w').write(s)
PY
rg -n 'func doctorAgent|for _, pat|doctorAgent\\(' packages/checks/internal/engine/doctor.go
gofmt -w packages/checks/internal/engine/doctor.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 647ms:
289:func doctorAgent(ctx context.Context, r *model.DoctorReport, req model.DoctorRequest, scratch, common, runID, role string, cmdValue, wrapValue any, patterns []string) {
338:		for _, pat := range patterns {
500:		doctorAgent(ctx, &r, req, scratch, common, runID, "worker", agent["command"], wrap, strSlice(snap.Config["limit_patterns"]))
506:		doctorAgent(ctx, &r, req, scratch, common, runID, "audit", auditCmd, wrap, strSlice(snap.Config["limit_patterns"]))

exec
/bin/zsh -lc "sed -n '320,520p' packages/checks/internal/engine/doctor.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 495ms:
		}
	}
	argv := append(wrap, cmd...)
	limit := req.Timeout
	if limit <= 0 {
		limit = 3 * time.Minute
	}
	out, stderr, exit, runErr := run(ctx, scratch, argv, nil, limit)
	logPath := filepath.Join(req.Root, ".fireweed", "runs", runID, "doctor-"+role+"-1.log")
	_ = os.MkdirAll(filepath.Dir(logPath), 0755)
	_ = os.WriteFile(logPath, []byte(out+stderr), 0644)
	detail := "log: " + logPath
	combined := out + stderr
	tail := tail20(combined)
	passed := runErr == nil && exit == 0 && strings.Contains(out, "FIREWEED_DOCTOR_OK")
	doctorCheck(r, role, "", passed, detail)
	if !passed {
		message := fmt.Sprintf("%s probe command %v failed: %v exit %d\n%s", role, argv, runErr, exit, tail)
		for _, pat := range patterns {
			if pat != "" && strings.Contains(strings.ToLower(tail), strings.ToLower(pat)) {
				message = "usage limit: " + message
				break
			}
		}
		doctorDiag(r, "C15", message, false)
	}
}
func Doctor(ctx context.Context, req model.DoctorRequest) model.DoctorReport {
	var r model.DoctorReport
	r.Diagnostics = append(r.Diagnostics, project.Check(req.Root)...)
	doctorWarnings(req.Root, &r)
	if staticErr(r.Diagnostics) {
		return r
	}
	repo, e := gitstore.Inspect(req.Root)
	if e != nil {
		doctorDiag(&r, "C14", "scratch repository: "+e.Error(), false)
		return r
	}
	common := repo.CommonDir
	if !filepath.IsAbs(common) {
		common = filepath.Join(req.Root, common)
	}
	common, _ = filepath.EvalSymlinks(common)
	base := filepath.Join(common, "fireweed", "worktrees")
	_ = os.MkdirAll(base, 0755)
	scratch, e := os.MkdirTemp(base, "doctor-")
	if e != nil {
		doctorDiag(&r, "C14", "scratch worktree: "+e.Error(), false)
		return r
	}
	_ = os.Remove(scratch)
	if _, e = git(req.Root, "worktree", "add", "--detach", scratch, "HEAD"); e != nil {
		doctorDiag(&r, "C14", "scratch worktree: "+e.Error(), false)
		return r
	}
	defer func() { _, _ = git(req.Root, "worktree", "remove", "--force", scratch); _ = os.RemoveAll(scratch) }()
	if e = copyTree(req.Root, scratch); e != nil {
		doctorDiag(&r, "C14", "scratch copy: "+e.Error(), false)
		return r
	}
	snap, e := project.Load(scratch)
	if e != nil {
		doctorDiag(&r, "C14", "scratch project: "+e.Error(), false)
		return r
	}
	setup := configured(snap.Config, "setup")
	if setup != nil {
		args, e := command(setup, nil)
		if e == nil {
			_, _, exit, runErr := run(ctx, scratch, args, map[string]string{"FIREWEED_BIN": req.Bin}, req.Timeout)
			if runErr != nil || exit != 0 {
				e = fmt.Errorf("exit %d: %v", exit, runErr)
			}
		}
		if e != nil {
			doctorDiag(&r, "C14", "setup hook failed: "+e.Error(), false)
			return r
		}
	}
	for _, m := range snap.Modules {
		if reason := boundaryOpt(m); reason != "" {
			doctorDiag(&r, "X03", m.Name+": boundaries opt out: "+reason, true)
		}
		if boundaryHook(snap, &m) == nil && (len(m.Zones.Disposable) == 0 || reason != "") {
			doctorDiag(&r, "X04", m.Name+": no boundary enforcement required", true)
		}
	}
	templates, e := templateFiles(scratch)
	if e != nil {
		doctorDiag(&r, "C14", "module template SPEC.md: "+e.Error(), false)
		return r
	}
	maps, e := canaryMaps(scratch)
	if e != nil {
		doctorDiag(&r, "C14", e.Error(), false)
		return r
	}
	modText := ""
	if b, e := os.ReadFile(filepath.Join(scratch, "go.mod")); e == nil {
		for _, l := range strings.Split(string(b), "\n") {
			if strings.HasPrefix(l, "module ") {
				modText = strings.TrimSpace(strings.TrimPrefix(l, "module "))
				break
			}
		}
	}
	names := []string{"doctorprobeone", "doctorprobetwo"}
	dirs := []string{"packages/" + names[0], "packages/" + names[1]}
	var canMods []project.Module
	for i, name := range names {
		vars := map[string]string{"name": name, "prefix": fmt.Sprintf("DCT%d", i+1), "package": name, "go_module": modText, "module_dir": dirs[i], "other_dir": dirs[1-i]}
		if e := writeMap(scratch, dirs[i], templates, vars); e != nil {
			doctorDiag(&r, "C14", "module template: "+e.Error(), false)
			return r
		}
		if e := writeMap(scratch, dirs[i], maps["baseline"], vars); e != nil {
			doctorDiag(&r, "C14", "canary baseline: "+e.Error(), false)
			return r
		}
		b, e := os.ReadFile(filepath.Join(scratch, dirs[i], "SPEC.md"))
		if e != nil {
			doctorDiag(&r, "C14", "module template SPEC.md: "+e.Error(), false)
			return r
		}
		d, e := contracts.ParseDocument(string(b))
		if e != nil {
			doctorDiag(&r, "C14", "module template SPEC.md: "+e.Error(), false)
			return r
		}
		m := project.Module{Name: name, IDPrefix: vars["prefix"], Dir: dirs[i], Zones: contracts.Zones{Disposable: []string{"internal/**"}, Evals: []string{"evals/**"}, Interface: []string{"api.txt"}}, Document: d}
		canMods = append(canMods, m)
	}
	globalEval := configured(snap.Config, "eval")
	globalBoundary := configured(snap.Config, "boundaries")
	targets := []struct {
		m           project.Module
		eval, bound any
		label       string
	}{{canMods[0], globalEval, globalBoundary, "global"}, {canMods[1], globalEval, globalBoundary, "global"}}
	for _, m := range snap.Modules {
		eh := configured(m.Document.Frontmatter, "eval")
		bh := configured(m.Document.Frontmatter, "boundaries")
		if eh != nil || bh != nil {
			if eh == nil {
				eh = globalEval
			}
			if bh == nil {
				bh = globalBoundary
			}
			targets = append(targets, struct {
				m           project.Module
				eval, bound any
				label       string
			}{m, eh, bh, m.Name + " eval override"})
		}
	}
	for _, target := range targets {
		m := target.m
		eh, bh := target.eval, target.bound
		if e := doctorBaseline(ctx, snap, m, eh, bh, req.Bin, req.Timeout); e != nil {
			doctorCheck(&r, "baseline", m.Name, false, e.Error())
			doctorDiag(&r, "C14", "baseline "+target.label+" "+e.Error(), false)
			continue
		}
		doctorCheck(&r, "baseline", m.Name, true, "")
		vars := map[string]string{"name": m.Name, "prefix": m.IDPrefix, "package": m.Name, "go_module": modText, "module_dir": m.Dir, "other_dir": dirs[0]}
		if m.Name == names[0] {
			vars["other_dir"] = dirs[1]
		}
		for _, name := range []string{"undeclared", "internal_eval", "failing_eval", "missing_id"} {
			hookName := target.label + " boundary hook"
			if name == "failing_eval" || name == "missing_id" {
				hookName = target.label + " eval hook"
			}
			doctorCanary(ctx, &r, snap, m, hookName, eh, bh, vars, maps[name], req.Bin, req.Timeout, name)
		}
	}
	if !req.NoAgent {
		runID := fmt.Sprintf("%d", time.Now().UnixNano())
		agent := stringMap(snap.Config["agent"])
		audit := stringMap(snap.Config["audit"])
		wrap := agent["wrap"]
		doctorAgent(ctx, &r, req, scratch, common, runID, "worker", agent["command"], wrap, strSlice(snap.Config["limit_patterns"]))
		auditCmd := audit["command"]
		if auditCmd == nil {
			auditCmd = agent["command"]
			doctorDiag(&r, "C15", "audit shares worker command; audit command blind spot", true)
		}
		doctorAgent(ctx, &r, req, scratch, common, runID, "audit", auditCmd, wrap, strSlice(snap.Config["limit_patterns"]))
	}
	return r
}

exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/checks/internal/engine/core.go'
s=open(p).read()
old='''if "'!known[normalized] {
					return fmt.Errorf("testcase %q references unknown spec ID %q", n.Name, id)
				}'"'''
new='''if "'!known[normalized] {
                    candidate := strings.ContainsAny(id, "-_")
                    if !candidate { for _, spec := range specs { prefix := strings.SplitN(spec.ID, "-", 2)[0]; if strings.HasPrefix(strings.ToUpper(id), strings.ToUpper(prefix)) { candidate = true; break } } }
                    if candidate { return fmt.Errorf("testcase %q references unknown spec ID %q", n.Name, id) }
                }'"'''
assert old in s
s=s.replace(old,new)
open(p,'w').write(s)
p='packages/checks/internal/engine/doctor.go';s=open(p).read();old='''if hook "'!= nil {
		copy.Config = map[string]any{"hooks": map[string]any{"boundaries": hook}}
		m.Document.Frontmatter = map[string]any{}
		copy.Modules[0] = m
	}'"''';new='''if hook "'!= nil {
        copy.Config = map[string]any{}
        m.Document.Frontmatter = map[string]any{"hooks":map[string]any{"boundaries":hook}}
        copy.Modules[0] = m
    }'"''';assert old in s;s=s.replace(old,new)
# Validate the source module template before baseline overlays it.
old='''if e := writeMap(scratch, dirs[i], maps[\"baseline\"], vars); e "'!= nil {'"'''
new='''templateSpec := renderCanary(templates[\"SPEC.md\"], vars)
        if _, e := contracts.ParseDocument(templateSpec); e "'!= nil { doctorDiag(&r,"C14","module template SPEC.md: "+e.Error(),false); return r }
        if e := writeMap(scratch, dirs[i], maps["baseline"], vars); e != nil {'"'''
assert old in s;s=s.replace(old,new)
open(p,'w').write(s)
PY
gofmt -w packages/checks/internal/engine/core.go packages/checks/internal/engine/doctor.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 605ms:
Traceback (most recent call last):
  File "<stdin>", line 11, in <module>
AssertionError

exec
/bin/zsh -lc "rg -n -C 5 'known\\[normalized\\]|copy.Config = map\\[string\\]any|writeMap\\(scratch, dirs\\[i\\], maps' packages/checks/internal/engine/core.go packages/checks/internal/engine/doctor.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 513ms:
packages/checks/internal/engine/core.go-180-				}
packages/checks/internal/engine/core.go-181-			}
packages/checks/internal/engine/core.go-182-			if len(c.IDs) == 0 {
packages/checks/internal/engine/core.go-183-				for _, id := range unknownID.FindAllString(n.Name+" "+n.Classname, -1) {
packages/checks/internal/engine/core.go-184-					normalized := strings.ToUpper(strings.ReplaceAll(id, "_", "-"))
packages/checks/internal/engine/core.go:185:					if !known[normalized] {
packages/checks/internal/engine/core.go-186-						return fmt.Errorf("testcase %q references unknown spec ID %q", n.Name, id)
packages/checks/internal/engine/core.go-187-					}
packages/checks/internal/engine/core.go-188-				}
packages/checks/internal/engine/core.go-189-				return fmt.Errorf("testcase %q has no known spec ID", n.Name)
packages/checks/internal/engine/core.go-190-			}
--
packages/checks/internal/engine/doctor.go-188-}
packages/checks/internal/engine/doctor.go-189-func evalProbe(ctx context.Context, s project.Snapshot, m project.Module, hook any, bin string, timeout time.Duration) (model.Report, error) {
packages/checks/internal/engine/doctor.go-190-	copy := s
packages/checks/internal/engine/doctor.go-191-	copy.Modules = []project.Module{m}
packages/checks/internal/engine/doctor.go-192-	if hook != nil {
packages/checks/internal/engine/doctor.go:193:		copy.Config = map[string]any{"hooks": map[string]any{"eval": hook}}
packages/checks/internal/engine/doctor.go-194-		m.Document.Frontmatter = map[string]any{}
packages/checks/internal/engine/doctor.go-195-		copy.Modules[0] = m
packages/checks/internal/engine/doctor.go-196-	}
packages/checks/internal/engine/doctor.go-197-	return Evaluate(ctx, model.HookRequest{Snapshot: copy, Module: m.Name, Bin: bin, Timeout: timeout})
packages/checks/internal/engine/doctor.go-198-}
packages/checks/internal/engine/doctor.go-199-func boundaryProbe(ctx context.Context, s project.Snapshot, m project.Module, hook any, bin string, timeout time.Duration) ([]project.Diagnostic, error) {
packages/checks/internal/engine/doctor.go-200-	copy := s
packages/checks/internal/engine/doctor.go-201-	copy.Modules = []project.Module{m}
packages/checks/internal/engine/doctor.go-202-	if hook != nil {
packages/checks/internal/engine/doctor.go:203:		copy.Config = map[string]any{"hooks": map[string]any{"boundaries": hook}}
packages/checks/internal/engine/doctor.go-204-		m.Document.Frontmatter = map[string]any{}
packages/checks/internal/engine/doctor.go-205-		copy.Modules[0] = m
packages/checks/internal/engine/doctor.go-206-	}
packages/checks/internal/engine/doctor.go-207-	return Boundaries(ctx, model.BoundaryRequest{Snapshot: copy, Bin: bin, Timeout: timeout})
packages/checks/internal/engine/doctor.go-208-}
--
packages/checks/internal/engine/doctor.go-431-		vars := map[string]string{"name": name, "prefix": fmt.Sprintf("DCT%d", i+1), "package": name, "go_module": modText, "module_dir": dirs[i], "other_dir": dirs[1-i]}
packages/checks/internal/engine/doctor.go-432-		if e := writeMap(scratch, dirs[i], templates, vars); e != nil {
packages/checks/internal/engine/doctor.go-433-			doctorDiag(&r, "C14", "module template: "+e.Error(), false)
packages/checks/internal/engine/doctor.go-434-			return r
packages/checks/internal/engine/doctor.go-435-		}
packages/checks/internal/engine/doctor.go:436:		if e := writeMap(scratch, dirs[i], maps["baseline"], vars); e != nil {
packages/checks/internal/engine/doctor.go-437-			doctorDiag(&r, "C14", "canary baseline: "+e.Error(), false)
packages/checks/internal/engine/doctor.go-438-			return r
packages/checks/internal/engine/doctor.go-439-		}
packages/checks/internal/engine/doctor.go-440-		b, e := os.ReadFile(filepath.Join(scratch, dirs[i], "SPEC.md"))
packages/checks/internal/engine/doctor.go-441-		if e != nil {

exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/checks/internal/engine/core.go';s=open(p).read();s=s.replace('''if "'!known[normalized] {
						return fmt.Errorf("testcase %q references unknown spec ID %q", n.Name, id)
					}'"''','''if "'!known[normalized] {
                        candidate := strings.ContainsAny(id, "-_")
                        if !candidate { for _, spec := range specs { prefix := strings.SplitN(spec.ID, "-", 2)[0]; if strings.HasPrefix(strings.ToUpper(id), strings.ToUpper(prefix)) { candidate = true; break } } }
                        if candidate { return fmt.Errorf("testcase %q references unknown spec ID %q", n.Name, id) }
                    }'"''');open(p,'w').write(s)
p='packages/checks/internal/engine/doctor.go';s=open(p).read();s=s.replace('''copy.Config = map[string]any{\"hooks\": map[string]any{\"boundaries\": hook}}
		m.Document.Frontmatter = map[string]any{}''','''copy.Config = map[string]any{}
        m.Document.Frontmatter = map[string]any{\"hooks\":map[string]any{\"boundaries\":hook}}''');s=s.replace('''		if e := writeMap(scratch, dirs[i], maps[\"baseline\"], vars); e "'!= nil {'"''','''        if _, e := contracts.ParseDocument(renderCanary(templates[\"SPEC.md\"],vars)); e "'!= nil { doctorDiag(&r,"C14","module template SPEC.md: "+e.Error(),false); return r }
		if e := writeMap(scratch, dirs[i], maps["baseline"], vars); e != nil {'"''');open(p,'w').write(s)
PY
gofmt -w packages/checks/internal/engine/core.go packages/checks/internal/engine/doctor.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 622ms:
exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/checks/internal/engine/doctor.go';s=open(p).read();needle='''	globalEval := configured(snap.Config, \"eval\")''';insert='''    // Include scratch canaries even when the project lists exact module paths.
    configPath := filepath.Join(scratch,\"fireweed.yaml\")
    if raw,err:=os.ReadFile(configPath);err==nil {
        var cfg map[string]any
        if yaml.Unmarshal(raw,&cfg)==nil {
            modules:=strSlice(cfg[\"modules\"])
            modules=append(modules,dirs...)
            cfg[\"modules\"]=modules
            if updated,err:=yaml.Marshal(cfg);err==nil { _=os.WriteFile(configPath,updated,0644) }
        }
    }
	globalEval := configured(snap.Config, \"eval\")''';assert needle in s;s=s.replace(needle,insert);open(p,'w').write(s)
PY
gofmt -w packages/checks/internal/engine/doctor.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 581ms:
exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/checks/internal/engine/core.go';s=open(p).read();old='''return contracts.ParseItems(section)''';new='''items, err := contracts.ParseItems(section)
    if err "'!= nil { if code,_,ok:=contracts.ErrorCode(err);ok { return nil,pref(code,err) }; return nil,err }
    return items,nil'"''';assert old in s;s=s.replace(old,new);open(p,'w').write(s)
PY
gofmt -w packages/checks/internal/engine/core.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 581ms:
exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/checks/internal/engine/doctor.go';s=open(p).read();old='''copy.Modules = []project.Module{m}
	if hook "'!= nil {
		copy.Config = map[string]any{}
		m.Document.Frontmatter = map[string]any{"hooks": map[string]any{"boundaries": hook}}
		copy.Modules[0] = m
	}'"''';new='''copy.Modules = []project.Module{m}
    for _, original := range s.Modules { if original.Name "'!= m.Name { copy.Modules=append(copy.Modules,project.Module{Name:original.Name,Dir:original.Dir,Document:contracts.Document{Frontmatter:map[string]any{}}}) } }
	if hook != nil {
		copy.Config = map[string]any{}
		m.Document.Frontmatter = map[string]any{"hooks": map[string]any{"boundaries": hook}}
		copy.Modules[0] = m
	}'"''';assert old in s;s=s.replace(old,new);open(p,'w').write(s)
PY
gofmt -w packages/checks/internal/engine/doctor.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 613ms:
exec
/bin/zsh -lc "sed -n '705,770p' packages/checks/evals/doctor_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 533ms:
	}
	if !hasDoctorCheck(r, "undeclared", true) {
		t.Errorf("fixture did not detect the independent undeclared canary: %s", doctorText(r))
	}
}

func Test_CHK_055_CHK_056_CHK_057_MissingRequiredBoundaryHook(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	f.put("packages/alpha/internal/live.txt", "disposable implementation\n")
	logDoctorCanaryProbes(t, f)
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(report, "X02", false, "alpha", "boundar") {
		t.Errorf("required boundary hook was not diagnosed: %s", doctorText(report))
	}
	if !hasDoctorDiagnostic(report, "C14", false, "baseline", "boundar") {
		t.Errorf("missing boundary hook was accepted as a clean baseline: %s", doctorText(report))
	}
	if strings.Contains(doctorLog(f), "planted:") {
		t.Errorf("missing baseline boundary hook continued into canary probes: %s", doctorLog(f))
	}
	for _, canary := range []string{"undeclared", "internal_eval"} {
		if hasDoctorCheck(report, canary, true) {
			t.Errorf("missing boundary hook counted as %s detection: %s", canary, doctorText(report))
		}
		if !hasDoctorDiagnostic(report, "C14", false, canary, "boundar") {
			t.Errorf("missing boundary hook lacked named %s canary failure: %s", canary, doctorText(report))
		}
	}
}

func Test_CHK_055_MissingRequiredHookBesideOverride(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	f.put(".fireweed/override-boundary.sh", fmt.Sprintf("#!/bin/sh\nprintf 'override-boundaries:%%s\\n' \"$PWD\" >> %q\nexec sh .fireweed/boundary.sh\n", f.log))
	f.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(f.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nhooks:\n  boundaries: sh .fireweed/override-boundary.sh", 1))
	f.put("packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: []\n---\n# Beta\n\n## Specs\n\n### BET-001\nThe fixture MUST pass.\n")
	f.put("packages/beta/internal/live.txt", "disposable implementation\n")
	logDoctorCanaryProbes(t, f)
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(report, "X02", false, "beta", "boundar") {
		t.Errorf("module without its own hook escaped required boundary enforcement: %s", doctorText(report))
	}
	if strings.Contains(doctorLog(f), "planted:") {
		t.Errorf("failed baseline beside override continued into canary probes: %s", doctorLog(f))
	}
	if !strings.Contains(doctorLog(f), "override-boundaries:") {
		t.Errorf("applicable override was not run: %s", doctorLog(f))
	}
}

func Test_CHK_055_ApplicableOverrideNeedsNoGlobalHook(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	f.put(".fireweed/override-boundary.sh", "#!/bin/sh\nexec sh .fireweed/boundary.sh\n")
	f.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(f.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nhooks:\n  boundaries: sh .fireweed/override-boundary.sh", 1))
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) || !hasDoctorCheck(report, "undeclared", true) || !hasDoctorCheck(report, "internal_eval", true) {
		t.Errorf("applicable override was rejected without a global hook: %s", doctorText(report))
	}
}

func Test_CHK_060_CHK_062_ModuleOverridesAndBoundaryWarnings(t *testing.T) {

exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/checks/internal/engine/doctor.go';s=open(p).read();old='''ds, e := boundaryProbe(ctx, s, m, boundHook, bin, timeout)
	if e "'!= nil {
		return fmt.Errorf("boundaries hook: %w (%v)", e, ds)
	}'"''';new='''if boundHook "'!= nil { ds, e := boundaryProbe(ctx, s, m, boundHook, bin, timeout)
        if e != nil { return fmt.Errorf("boundaries hook: %w (%v)", e, ds) } }'"''';assert old in s;s=s.replace(old,new)
needle='''	templates, e := templateFiles(scratch)''';insert='''    if ds, boundaryErr := Boundaries(ctx,model.BoundaryRequest{Snapshot:snap,Bin:req.Bin,Timeout:req.Timeout}); boundaryErr "'!= nil {
        r.Diagnostics=append(r.Diagnostics,ds...)
        doctorDiag(&r,"C14","baseline boundaries hook: "+boundaryErr.Error(),false)
        doctorDiag(&r,"C14","undeclared boundaries hook unavailable",false)
        doctorDiag(&r,"C14","internal_eval boundaries hook unavailable",false)
        return r
    }
	templates, e := templateFiles(scratch)'"''';assert needle in s;s=s.replace(needle,insert)
old='''for _, name := range []string{\"undeclared\", \"internal_eval\", \"failing_eval\", \"missing_id\"} {
			hookName := target.label + \" boundary hook\"''';new='''for _, name := range []string{\"undeclared\", \"internal_eval\", \"failing_eval\", \"missing_id\"} {
            if bh == nil && (name==\"undeclared\" || name==\"internal_eval\") { continue }
			hookName := target.label + \" boundary hook\"''';assert old in s;s=s.replace(old,new)
open(p,'w').write(s)
PY
gofmt -w packages/checks/internal/engine/doctor.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 692ms:
exec
/bin/zsh -lc 'GOCACHE=/tmp/fireweed-rebuild-cache go test ./packages/checks/...' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc 'git diff --stat && git status --short && gofmt -l packages/checks/internal/engine/*.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 exited 1 in 669ms:
FAIL	./packages/checks/... [setup failed]
# ./packages/checks/...
pattern ./packages/checks/...: directory prefix packages/checks does not contain main module or its selected dependencies
FAIL

 succeeded in 713ms:
?? packages/checks/internal/

exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/checks/internal/engine/history.go';s=open(p).read();needle='''func History(root, revisionRange string) (model.HistoryReport, error) {''';insert='''func projectedHistory(root,rev string,s project.Snapshot,changes map[string]string)(project.Snapshot,error){
    copy:=s;copy.Modules=append([]project.Module(nil),s.Modules...)
    for p:=range changes { b,e:=gitstore.ReadAt(root,rev,p);if e"'!=nil{return copy,e};d,e:=contracts.ParseDocument(string(b));if e!=nil{return copy,e};set:=stringMap(d.Frontmatter["set"]);if len(set)==0{continue};name:=str(d.Frontmatter["module"]);for i:=range copy.Modules{if copy.Modules[i].Name!=name{continue};z:=&copy.Modules[i].Zones
        if v,ok:=set["evals"];ok{z.Evals=strSlice(v)}
        if v,ok:=set["interface"];ok{z.Interface=strSlice(v)}
        if v,ok:=set["append_only"];ok{z.AppendOnly=strSlice(v)}
        if v,ok:=set["disposable"];ok{z.Disposable=strSlice(v)}
        if v,ok:=set["preserve"];ok{z.Preserve=strSlice(v)}
        if v,ok:=set["protected"];ok{z.Protected=strSlice(v)}
    }}
    return copy,nil
}
func History(root, revisionRange string) (model.HistoryReport, error) {'"''';assert needle in s;s=s.replace(needle,insert)
old='''snaps := append([]project.Snapshot{after}, before...)''';new='''projectedAfter,e:=projectedHistory(root,oid,after,afterChanges);if e"'!=nil{return rep,e}
        snaps := append([]project.Snapshot{projectedAfter}, before...)'"''';assert old in s;s=s.replace(old,new);open(p,'w').write(s)
PY
gofmt -w packages/checks/internal/engine/history.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 615ms:
exec
/bin/zsh -lc "sed -n '185,280p' packages/checks/internal/engine/history.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 505ms:
		return rep, nil
	}
	rep.Commits = strings.Split(list, "\n")
	for _, oid := range rep.Commits {
		after, e := project.LoadAt(root, oid)
		if e != nil {
			return rep, e
		}
		afterChanges, e := historyChanges(root, oid)
		if e != nil {
			return rep, e
		}
		parentText, e := git(root, "rev-list", "--parents", "-n", "1", oid)
		if e != nil {
			return rep, e
		}
		parents := strings.Fields(parentText)
		if len(parents) > 0 {
			parents = parents[1:]
		}
		var before []project.Snapshot
		var beforeChanges []map[string]string
		for _, parent := range parents {
			s, e := project.LoadAt(root, parent)
			if e != nil {
				return rep, e
			}
			before = append(before, s)
			m, e := historyChanges(root, parent)
			if e != nil {
				return rep, e
			}
			beforeChanges = append(beforeChanges, m)
		}
		var diffs []hchange
		if len(parents) == 0 {
			diffs, e = historyDiff(root, "", oid)
			if e != nil {
				return rep, e
			}
		} else {
			seen := map[string]bool{}
			for _, parent := range parents {
				ds, e := historyDiff(root, parent, oid)
				if e != nil {
					return rep, e
				}
				for _, d := range ds {
					k := d.status + "|" + d.old + "|" + d.new
					if !seen[k] {
						diffs = append(diffs, d)
						seen[k] = true
					}
				}
			}
		}
		worker := false
		implApplied := false
		squashed := false
		for p, phase := range afterChanges {
			if phase == "" {
				continue
			}
			oldPhase := ""
			held := false
			for _, m := range beforeChanges {
				if x, ok := m[p]; ok {
					oldPhase = x
					held = held || x == phase
				}
				for old, x := range m {
					if path.Base(old) == path.Base(p) {
						if oldPhase == "" {
							oldPhase = x
						}
						held = held || x == phase
					}
				}
			}
			if held {
				continue
			}
			worker = true
			if phase == "applied" {
				if oldPhase == "impl" && len(parents) == 1 {
					implApplied = true
				} else {
					rep.Squashed = append(rep.Squashed, model.SquashedChange{Commit: oid, Change: p})
					squashed = true
				}
			}
		}
		projectedAfter, e := projectedHistory(root, oid, after, afterChanges)
		if e != nil {
			return rep, e
		}

exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/checks/internal/engine/history.go';s=open(p).read();a=s.index('''		} else {
			seen := map[string]bool{}''');b=s.index('''
		worker := false''',a);new='''        } else {
            counts:=map[string]int{}
            choices:=map[string]hchange{}
            for _,parent:=range parents { ds,err:=historyDiff(root,parent,oid);if err"'!=nil{return rep,err};once:=map[string]bool{};for _,d:=range ds {key:=d.new;if d.old!=""{key=d.old+"|"+d.new};if !once[key]{counts[key]++;once[key]=true;choices[key]=d}} }
            for key,count:=range counts {if len(parents)==1 || count==len(parents){diffs=append(diffs,choices[key])}}
            sort.Slice(diffs,func(i,j int)bool{return diffs[i].new<diffs[j].new})
        }'"''';s=s[:a]+new+s[b:];s=s.replace('\"path\"','\"path\"\\n \"sort\"');open(p,'w').write(s)
PY
gofmt -w packages/checks/internal/engine/history.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 584ms:
exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/checks/internal/engine/history.go';s=open(p).read();needle='''		worker := false''';ins='''        for _,d:=range diffs {if d.status==\"D\" && path.Base(d.new)==\"SPEC.md\" { return rep,fmt.Errorf(\"required recorded module document missing: %s at %s\",d.new,oid) }}
		worker := false''';assert needle in s;s=s.replace(needle,ins);open(p,'w').write(s)
PY
gofmt -w packages/checks/internal/engine/history.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 586ms:
exec
/bin/zsh -lc 'head -35 packages/checks/internal/engine/zones_boundaries.go && head -33 packages/checks/internal/engine/history.go && head -35 packages/checks/internal/engine/core.go && git status --short' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 541ms:
package engine

import (
	"context"
	"fmt"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func matchAny(globs []string, p string) bool {
	for _, g := range globs {
		if contracts.MatchGlob(g, p) {
			return true
		}
	}
	return false
}
func baseExists(root, base, p string) bool {
	if base == "" {
		return false
	}
	_, err := gitstore.ReadAt(root, base, p)
	return err == nil
}
func zonePath(r model.ZoneRequest, p, status string) *project.Diagnostic {
	p = path.Clean(strings.TrimPrefix(p, "./"))
	bad := func(reason string) *project.Diagnostic {
		return &project.Diagnostic{Code: "Z01", Path: p, Message: reason + ": " + p}
package engine

import (
	"context"
	"fmt"
	"os/exec"
	"path"
	"sort"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

func git(root string, args ...string) (string, error) {
	c := exec.CommandContext(context.Background(), "git", append([]string{"-C", root}, args...)...)
	b, e := c.CombinedOutput()
	if e != nil {
		return "", fmt.Errorf("git %v: %w: %s", args, e, b)
	}
	return strings.TrimSpace(string(b)), nil
}
func historyChanges(root, rev string) (map[string]string, error) {
	out := map[string]string{}
	names, e := git(root, "ls-tree", "-r", "--name-only", rev, "--", "changes")
	if e != nil {
		return nil, e
	}
	for _, p := range strings.Split(names, "\n") {
		if p == "" || !strings.HasSuffix(p, ".md") {
			continue
package engine

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func pref(code string, err error) error { return fmt.Errorf("%s: %w", code, err) }
func stringMap(v any) map[string]any {
	if m, ok := v.(map[string]any); ok {
		return m
	}
	return nil
}
func str(v any) string { s, _ := v.(string); return s }
func strSlice(v any) []string {
	switch x := v.(type) {
	case []string:
		return x
	case []any:
?? packages/checks/internal/

exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/checks/internal/engine/zones_boundaries.go';s=open(p).read().replace('\\n\\t\"github.com/lurkshark/fireweed/packages/gitstore\"\\n\\t\"os\"','\\n\\t\"os\"').replace('\"github.com/lurkshark/fireweed/packages/contracts\"\\n','\"github.com/lurkshark/fireweed/packages/contracts\"\\n\\t\"github.com/lurkshark/fireweed/packages/gitstore\"\\n');open(p,'w').write(s)
PY
gofmt -w packages/checks/internal/engine/*.go
cat > DRILL_NOTES.md <<'EOF'
# Checks rebuild notes

## Contradicted Decisions

None identified.

## Gaps

No known behavioral gaps in the target checks module. The supplied workspace omits the Go manifest and dependency implementations, so the retained evals could not be run here. Native verification remains with the outside controller.
EOF
git status --short && gofmt -l packages/checks/internal/engine/*.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 614ms:
?? DRILL_NOTES.md
?? packages/checks/internal/

exec
/bin/zsh -lc "sed -n '30,90p' packages/checks/internal/engine/zones_boundaries.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 515ms:
	return err == nil
}
func zonePath(r model.ZoneRequest, p, status string) *project.Diagnostic {
	p = path.Clean(strings.TrimPrefix(p, "./"))
	bad := func(reason string) *project.Diagnostic {
		return &project.Diagnostic{Code: "Z01", Path: p, Message: reason + ": " + p}
	}
	if p == r.ChangePath {
		if r.Phase == "audit" || r.Phase == "evals" || r.Phase == "impl" {
			return nil
		}
	}
	if strings.HasPrefix(p, "changes/") {
		return bad("other change file")
	}
	if p == ".fireweed" || strings.HasPrefix(p, ".fireweed/") {
		return bad("protected path")
	}
	if matchAny(strSlice(r.Snapshot.Config["protected"]), p) {
		return bad("protected path")
	}
	if r.Phase == "audit" {
		return bad("audit edit")
	}
	var owner *project.Module
	for i := range r.Snapshot.Modules {
		m := &r.Snapshot.Modules[i]
		if p == m.Dir || strings.HasPrefix(p, m.Dir+"/") {
			if owner == nil || len(m.Dir) > len(owner.Dir) {
				owner = m
			}
		}
	}
	if owner == nil {
		if matchAny(strSlice(r.Snapshot.Config["worker_writable"]), p) {
			return nil
		}
		return bad("outside worker writable")
	}
	if owner.Name != r.Module {
		return bad("other module")
	}
	rel := strings.TrimPrefix(p, owner.Dir+"/")
	if rel == "SPEC.md" || matchAny(owner.Zones.Protected, rel) {
		return bad("protected module path")
	}
	zone := contracts.ResolveZone(rel, owner.Zones)
	switch r.Phase {
	case "evals":
		switch zone {
		case "evals", "other":
			return nil
		case "interface":
			if r.InterfaceDeclared {
				return nil
			}
		case "disposable":
			if !baseExists(r.Snapshot.Root, r.Base, p) {
				return nil
			}
		}

exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/checks/internal/engine/zones_boundaries.go';s=open(p).read();start=s.index('\\tif p == r.ChangePath {');end=s.index('\\tif r.Phase == \"audit\" {',start);s=s[:start]+'''	if p == \".fireweed\" || strings.HasPrefix(p, \".fireweed/\") || matchAny(strSlice(r.Snapshot.Config[\"protected\"]), p) { return bad(\"protected path\") }
    if p == r.ChangePath && (r.Phase==\"audit\"||r.Phase==\"evals\"||r.Phase==\"impl\") {return nil}
    if strings.HasPrefix(p,\"changes/\") {return bad(\"other change file\")}
'''+s[end:];open(p,'w').write(s)
PY
gofmt -w packages/checks/internal/engine/zones_boundaries.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 610ms:
exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/checks/internal/engine/doctor.go';s=open(p).read();old='''hookName := target.label + \" boundary hook\"
			if name == \"failing_eval\" || name == \"missing_id\" {
				hookName = target.label + \" eval hook\"
			}''';new='''hookName := target.label + \" boundary hook\"
            if configured(m.Document.Frontmatter,\"boundaries\")"'!=nil { hookName=m.Name+" boundaries override boundary hook" }
			if name == "failing_eval" || name == "missing_id" {
				hookName = target.label + " eval hook"
                if configured(m.Document.Frontmatter,"eval")!=nil {hookName=m.Name+" eval override eval hook"}
			}'"''';assert old in s;s=s.replace(old,new);open(p,'w').write(s)
PY
gofmt -w packages/checks/internal/engine/doctor.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 594ms:
exec
/bin/zsh -lc "rg -n 'fmt.Errorf\\(["'^"].*'"\\)' packages/checks/internal/engine/*.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 538ms:
packages/checks/internal/engine/zones_boundaries.go:167:					first = fmt.Errorf(d.Message)

exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/checks/internal/engine/zones_boundaries.go';s=open(p).read().replace('\"context\"','\"context\"\\n \"errors\"').replace('first = fmt.Errorf(d.Message)','first = errors.New(d.Message)');open(p,'w').write(s)
PY
gofmt -w packages/checks/internal/engine/zones_boundaries.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 1229ms:
codex
The rebuild files are formatted, and the Doctor probe paths now include scratch cleanup, canary recovery, and probe logs. I attempted the retained Go tests; this input workspace has no Go manifest or dependency implementations, so `go test` stops before compilation. I’m doing a final source and scope check.
exec
/bin/zsh -lc 'git status --short && find packages/checks/internal -type f -maxdepth 3 -print && wc -l packages/checks/internal/engine/*.go DRILL_NOTES.md' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc "gofmt -l packages/checks/internal/engine/*.go && rg -n 'TODO|FIXME|panic\\(' packages/checks/internal/engine DRILL_NOTES.md" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 exited 1 in 531ms:
 succeeded in 581ms:
?? DRILL_NOTES.md
?? packages/checks/internal/
packages/checks/internal/engine/zones_boundaries.go
packages/checks/internal/engine/core.go
packages/checks/internal/engine/history.go
packages/checks/internal/engine/doctor.go
     373 packages/checks/internal/engine/core.go
     552 packages/checks/internal/engine/doctor.go
     366 packages/checks/internal/engine/history.go
     224 packages/checks/internal/engine/zones_boundaries.go
       9 DRILL_NOTES.md
    1524 total

codex
Rebuilt checks under [packages/checks/internal/engine](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562/packages/checks/internal/engine) and wrote [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562/DRILL_NOTES.md). Retained inputs were untouched; nothing was committed.

The new Go files pass `gofmt`. I could not run the retained evals here: the supplied workspace has no Go manifest or dependency implementations, so `go test` stops before compilation.
tokens used
151,607


### Attempt 2 (exit 0)
Reconstruction held for outside-controller input/verifier inspection.
Outside-controller inspection released prescribed reconstruction worker.
The checks implementation is in the target disposable paths. I fixed Doctor’s boundary canary detection, made canary names unique, tightened JUnit unknown-ID handling, and corrected History’s phase and zone checks. I also updated [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562/DRILL_NOTES.md).

`gofmt` passed. The retained Go evals could not run because this supplied workspace has no `go.mod` or dependent internal implementations, so the prior `PackageFailure` and coverage failures remain unverified. No commit was made.

Reading additional input from stdin...
OpenAI Codex v0.165.0
--------
workdir: /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
model: gpt-6-sol
provider: openai
approval: never
sandbox: workspace-write [workdir, /tmp, $TMPDIR]
reasoning effort: high
reasoning summaries: none
session id: 01a1b19b-476a-7d30-8890-9be98a5042e2
--------
user
# Rebuild checks

Target spec and Decisions:
---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
id_prefix: CHK
interface:
    - api.go
    - model/**
name: checks
---

# checks

## Specs

### CHK-001
ParseJUnit MUST accept testsuite or testsuites documents, including nested suites, returning testcase name, classname and time.

### CHK-002
RequireCoverage MUST cover each current spec except review or quarantine with a non-skipped referencing case. Missing coverage MUST return a C10-prefixed error.

### CHK-003
Evaluate MUST run the configured eval hook, accepting shell strings or argv.

### CHK-004
During audit, Zones MUST allow edits only to the current change file.

### CHK-005
Boundaries MUST run applicable global and per-module boundary hooks.

### CHK-006
Zones MUST reject edits to protected paths in every phase.

### CHK-007
Zones MUST reject edits to module SPEC.md files in every phase.

### CHK-008
Zones MUST reject edits within another module in every phase.

### CHK-009
Zones MUST reject edits to change files other than the current change in every phase.

### CHK-010
During evals and impl, Zones MUST allow edits outside modules only when the path matches worker_writable, except for the current change file.

### CHK-011
During evals, Zones MUST allow edits in the target module's evals zone.

### CHK-012
During evals, Zones MUST allow edits in the target module's interface zone only when InterfaceDeclared is true.

### CHK-013
During evals, Zones MUST allow edits in the target module's disposable zone only for files absent from the phase base.

### CHK-014
During evals, Zones MUST allow edits in the target module's other zone.

### CHK-015
During evals, Zones MUST reject edits in the target module's preserve zone.

### CHK-016
During impl, Zones MUST reject edits in the target module's evals zone.

### CHK-017
During impl, Zones MUST reject edits in the target module's interface zone.

### CHK-018
During impl, Zones MUST allow edits in the target module's disposable zone.

### CHK-019
During impl, Zones MUST allow edits in the target module's preserve zone.

### CHK-020
During impl, Zones MUST allow edits in the target module's other zone.

### CHK-021
During evals and impl, Zones MUST allow edits to the current change file.

### CHK-022
During impl, Zones MUST allow only newly added files in the target module's append_only zone.

### CHK-023
During evals, Zones MUST reject edits in the target module's append_only zone.

### CHK-024
Zones MUST check both the source and destination of a rename against phase permissions.

### CHK-025
RequireFailFirst MUST find a failed or errored referencing case for every requested enforced ID except already_satisfied; passing, skipped, empty or unrelated failures do not qualify. Missing genuine fail-first evidence MUST return a P2-01-prefixed error.

### CHK-026
RequirePass MUST reject failed or errored cases referencing any non-quarantined spec with an E01-prefixed error.

### CHK-027
ParseJUnit MUST reject malformed XML with a C11-prefixed error.

### CHK-028
ParseJUnit MUST reject testcases with missing or unknown spec IDs with a C11-prefixed error.

### CHK-029
ParseJUnit MUST map testcase name and classname to current spec IDs using contracts.MatchesID.

### CHK-030
ParseJUnit MUST classify failure and error elements as fail, and skipped elements as skip.

### CHK-031
Evaluate MUST prefer a module's eval hook override over the global eval hook.

### CHK-032
Evaluate MUST expand module, module_dir and report placeholders once in shell strings or argv, quoting substituted shell values.

### CHK-033
Evaluate MUST set the hook's FIREWEED_BIN environment variable to the request Bin.

### CHK-034
Evaluate MUST retain the hook process's stdout, stderr and exit code in its returned report.

### CHK-035
Evaluate MUST parse the hook's written JUnit report into the returned report's cases.

### CHK-036
Evaluate MUST fail with a C11-prefixed error when the eval hook cannot launch or exceeds its timeout.

### CHK-037
Evaluate MUST fail with a C11-prefixed error when the eval hook does not write its JUnit report.

### CHK-038
Evaluate MUST reject a zero exit with mapped failures or a nonzero exit without mapped failures with a C11-prefixed error, while accepting a nonzero exit with mapped failures as valid report evidence.

### CHK-039
Boundaries MUST return errors with diagnostic X01 for unsuppressed module-attributed violations.

### CHK-040
Boundaries MUST return a visible X03 warning containing the reason for a reasoned boundaries_opt_out.

### CHK-041
Boundaries MUST suppress only exact baseline lines from normally exited hooks.

### CHK-042
Boundaries MUST suppress only the opted-out module's violation lines through a reasoned opt-out after a hook exits normally.

### CHK-043
Boundaries MUST report hook launch, cancellation or timeout failures as errors with a diagnostic whose Code is X02. The returned error MUST preserve the hook failure cause without a leading X02 prefix.

### CHK-044
Boundaries MUST fail with diagnostic X02 for a missing hook required by a module with disposable globs and no reasoned opt-out.

### CHK-045
Boundaries MUST warn with diagnostic X04 without an error for a missing hook when no boundary enforcement is required.

### CHK-046
Evaluate MUST return C11-prefixed errors for existing eval-hook setup failures, including missing or invalid hook commands.

### CHK-047
Evaluate MUST return a U02-prefixed error for an unknown requested module.

### CHK-048
Evaluate MUST return C03-prefixed errors for malformed module item syntax and C05-prefixed errors for duplicate module item identities.

### CHK-049
Doctor MUST retain project.Check diagnostics and prevent hook or agent probes when the current project has a static error.

### CHK-050
Doctor MUST run probes using current project files in a detached scratch worktree below the common Git directory's fireweed/worktrees, leaving the caller's checkout unchanged.

### CHK-051
Doctor MUST remove its scratch worktree and Git registration after successful or failed probes.

### CHK-052
Doctor MUST run the configured setup hook in scratch before probes; a setup failure MUST produce C14 naming the setup hook and prevent later probes.

### CHK-053
Doctor MUST render two uniquely named throwaway modules from project-owned module templates and the version-1 canary baseline, substituting name, prefix, package, go_module, module_dir and other_dir.

### CHK-054
Doctor MUST report C14 naming missing or malformed module templates or required canary maps and skip the affected probes.

### CHK-055
Doctor MUST require passing baseline evals and coverage for the throwaway canary modules, and applicable boundaries, before planting canaries and after removing each planted canary. This baseline coverage requirement applies to the canary modules used to prove hooks, rather than every existing module. Recovery failures MUST produce C14 naming the failed hook and canary.

### CHK-056
Doctor MUST plant the supplied undeclared canary and require its boundary hook to report a module-attributed violation; a missed canary MUST produce C14 naming the hook and canary.

### CHK-057
Doctor MUST plant the supplied internal_eval canary and require its boundary hook to report a module-attributed interface violation; a missed canary MUST produce C14 naming the hook and canary.

### CHK-058
Doctor MUST require the supplied failing_eval canary to produce a failed referencing testcase and nonzero eval exit; compile errors, missing reports and unrelated failures MUST NOT count as detection.

### CHK-059
Doctor MUST require the supplied missing_id canary to be rejected for its unmapped testcase ID; unrelated hook or report failures MUST NOT count as detection.

### CHK-060
Doctor MUST test applicable global eval and boundary hooks and every per-module override through their corresponding canary probes.

### CHK-061
Doctor MUST set probe hooks' FIREWEED_BIN to the request Bin and make the scratch canary modules discoverable through that executable's modules output.

### CHK-062
Doctor MUST retain each real module's reasoned X03 boundaries-opt-out warning and applicable X04 warning when no boundary enforcement is required.

### CHK-063
Doctor MUST run the worker and audit commands once each in scratch through the configured agent.wrap with a fixed-reply prompt, requiring exit 0 and the requested reply for success.

### CHK-064
Doctor MUST use agent.command for the audit probe when audit.command is unset and emit a C15 warning identifying the shared-command blind spot.

### CHK-065
Doctor MUST bound each agent probe by the request Timeout or three minutes when unset, and report C15 failures with the command and its final 20 output lines.

### CHK-066
Doctor MUST classify a configured limit_patterns match in an agent probe's final 20 output lines as a C15 usage-limit diagnostic rather than an ordinary command failure.

### CHK-067
Doctor with NoAgent MUST skip worker and audit command execution while retaining hook canary and recovery checks.

### CHK-068
Doctor MUST emit a C16 warning when the AGENTS.md managed block is missing or differs from the current fireweed guidance, while accepting a current block regardless of surrounding text.

### CHK-069
Doctor MUST emit a C16 warning when a project playbook override's embedded_base is missing or differs from the SHA-256 of its current embedded playbook; a matching base MUST avoid that warning.

### CHK-070
History MUST list full commit IDs selected by the supplied Git revision range in oldest-first topological order.
For an empty range it MUST use the merge base of HEAD and the locally resolved configured base through HEAD.

### CHK-071
History MUST identify worker commits from new phase assignments in change-file contents across pending/archive moves, rather than commit subjects.
Deleting a change document without assigning a phase MUST not by itself classify the commit as a worker commit.

### CHK-072
History MUST report H01 when a worker commit changes a protected path or a project playbook override.

### CHK-073
History MUST report H02 when an impl-to-applied phase commit changes evals or interface zones.

### CHK-074
History MUST report H03 when any commit modifies, deletes or renames an existing append_only path, including merge-resolution changes identified by combined Git diffs.

### CHK-075
History MUST report H04 when a commit modifies an existing file under changes/archive.
Deleting an archived change file MUST remain allowed by H04.

### CHK-076
History MUST recognize each new applied phase whose parents lack an impl phase as squashed, listing its commit and change without a warning.
H01 and H02 MUST not apply to those squashed commits.
H03 and H04 MUST still apply.

### CHK-077
History MUST return each rule finding as a Diagnostic with its H01–H04 code, repository-relative Path and full commit ID in Message, using the recorded before/after zone maps and the worker change's projected set values.

### CHK-078
History MUST return an error when a revision range or required recorded project/change document cannot be read or parsed.

### CHK-079
ParseJUnit MUST retain the message attributes and text bodies of JUnit failure and error elements in each affected Case.Failure, including message-only and body-only diagnostics.
A testcase without such text MUST have an empty Failure.

### CHK-080
Evaluate MUST retain parsed Case.Failure strings in its returned Report.Cases, including when failing testcases make the eval hook exit nonzero.

### CHK-081
Doctor MUST persist complete stdout and stderr from every executed worker and audit probe, including failed, timed-out and cancelled probes, under the invoking project's .fireweed/runs/<run>/<phase>-1.log. Each Doctor invocation MUST use a fresh shared run identity, with phases doctor-worker and doctor-audit. Each executed probe's Detail MUST identify its log file by absolute path.

### CHK-082
Doctor MUST expand {prompt}, {prompt_file} and {git_common_dir} once in selected worker/audit command templates and wrapper argv, quoting shell substitutions and preserving argv values literally.

### CHK-083
Doctor MUST create one temporary file containing the exact probe prompt only when a selected command or wrapper template uses {prompt_file}, and expand that placeholder to its absolute path. The file MUST remain available through execution and be removed afterward, including after launch failure, nonzero exit, timeout or cancellation.

## Decisions

[0071-doctor-probes] Doctor requests the literal reply FIREWEED_DOCTOR_OK; limit_patterns match
case-insensitive literal substrings. Project playbook overrides use frontmatter
embedded_base containing the hexadecimal SHA-256 of the embedded playbook text.

[0114-history-validation] A merge merely retaining a phase value already held by a parent is not a new
phase assignment; combined merge diffs still enforce append-only files.
Squashed entries identify the change's repository-relative after-path; ordinary
non-merge phase transitions compare the sole parent.


Retained interfaces and files:
packages/checks/SPEC.md
packages/checks/api.go
packages/checks/model/model.go

Retained evals:
packages/checks/evals/.gitkeep
packages/checks/evals/boundaries_test.go
packages/checks/evals/doctor_agent_logs_test.go
packages/checks/evals/doctor_prompt_file_test.go
packages/checks/evals/doctor_test.go
packages/checks/evals/history_test.go
packages/checks/evals/hooks_test.go
packages/checks/evals/junit_test.go
packages/checks/evals/zones_test.go

Rebuild only target disposable paths. Do not edit retained inputs. Do not read outside the supplied input workspace, external trees, old code, source history, binaries, saved attempts, or dependent internal implementation. Do not repair dependent internals. Write DRILL_NOTES.md with separate ## Contradicted Decisions and ## Gaps sections. Every gap must name the nearest spec ID or "no spec" and a target module. Do not commit.


# Previous verification failed
checks eval: C11: testcase "PackageFailure" has no known spec ID
C10: no coverage for CHK-001
app eval: C11: testcase "PackageFailure" has no known spec ID
C10: no coverage for APP-001
runner eval: C11: testcase "PackageFailure" has no known spec ID
C10: no coverage for RUN-001
workflow eval: C11: testcase "PackageFailure" has no known spec ID
C10: no coverage for FLOW-001
checks stdout: 
checks stderr: 

app stdout: 
app stderr: 

runner stdout: 
runner stderr: 

workflow stdout: 
workflow stderr: 

codex
I’ll inspect the checks interfaces, retained evals, and module layout, then rebuild the disposable implementation paths. I’ll run the relevant checks and record any remaining spec gaps in DRILL_NOTES.md.
exec
/bin/zsh -lc 'git status --short && find packages/checks -maxdepth 3 -type f | sort' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc "pwd && rg --files -g 'AGENTS.md' -g 'SPEC.md' -g 'api.go' -g 'model.go' -g 'go.mod' -g 'DRILL_NOTES.md' -g 'packages/checks/**' | sort" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 496ms:
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
DRILL_NOTES.md
packages/assets/api.go
packages/assets/model/model.go
packages/checks/SPEC.md
packages/checks/api.go
packages/checks/evals/.gitkeep
packages/checks/evals/boundaries_test.go
packages/checks/evals/doctor_agent_logs_test.go
packages/checks/evals/doctor_prompt_file_test.go
packages/checks/evals/doctor_test.go
packages/checks/evals/history_test.go
packages/checks/evals/hooks_test.go
packages/checks/evals/junit_test.go
packages/checks/evals/zones_test.go
packages/checks/internal/engine/core.go
packages/checks/internal/engine/doctor.go
packages/checks/internal/engine/history.go
packages/checks/internal/engine/zones_boundaries.go
packages/checks/model/model.go
packages/contracts/api.go
packages/contracts/model/model.go
packages/gitstore/api.go
packages/gitstore/model/model.go
packages/project/api.go
packages/project/model/model.go

 succeeded in 499ms:
?? DRILL_NOTES.md
?? packages/checks/internal/
packages/checks/SPEC.md
packages/checks/api.go
packages/checks/evals/.gitkeep
packages/checks/evals/boundaries_test.go
packages/checks/evals/doctor_agent_logs_test.go
packages/checks/evals/doctor_prompt_file_test.go
packages/checks/evals/doctor_test.go
packages/checks/evals/history_test.go
packages/checks/evals/hooks_test.go
packages/checks/evals/junit_test.go
packages/checks/evals/zones_test.go
packages/checks/internal/engine/core.go
packages/checks/internal/engine/doctor.go
packages/checks/internal/engine/history.go
packages/checks/internal/engine/zones_boundaries.go
packages/checks/model/model.go

exec
/bin/zsh -lc 'cat packages/checks/internal/engine/doctor.go packages/checks/internal/engine/history.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc 'cat packages/checks/internal/engine/core.go packages/checks/internal/engine/zones_boundaries.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc 'cat packages/checks/api.go packages/checks/model/model.go packages/checks/SPEC.md' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc "cat DRILL_NOTES.md; rg --files -g 'go.mod' -g 'go.work' -g 'Makefile' -g 'fireweed.yaml' -g 'fireweed.yml' -g 'AGENTS.md' ." in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 581ms:
package checks

import (
	"context"

	"github.com/lurkshark/fireweed/packages/checks/internal/engine"
	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

type Case = model.Case
type Report = model.Report
type HookRequest = model.HookRequest
type BoundaryRequest = model.BoundaryRequest
type ZoneRequest = model.ZoneRequest
type DoctorRequest = model.DoctorRequest
type DoctorCheck = model.DoctorCheck
type DoctorReport = model.DoctorReport
type HistoryReport = model.HistoryReport
type SquashedChange = model.SquashedChange

func ParseJUnit(data []byte, specs []contracts.Item) (Report, error) {
	return engine.ParseJUnit(data, specs)
}
func RequireCoverage(report Report, specs []contracts.Item) error {
	return engine.RequireCoverage(report, specs)
}
func RequireFailFirst(report Report, ids, already []string) error {
	return engine.RequireFailFirst(report, ids, already)
}
func RequirePass(report Report, specs []contracts.Item) error {
	return engine.RequirePass(report, specs)
}
func Evaluate(ctx context.Context, request HookRequest) (Report, error) {
	return engine.Evaluate(ctx, request)
}
func Boundaries(ctx context.Context, request BoundaryRequest) ([]project.Diagnostic, error) {
	return engine.Boundaries(ctx, request)
}
func Zones(request ZoneRequest) []project.Diagnostic { return engine.Zones(request) }
func Doctor(ctx context.Context, request DoctorRequest) DoctorReport {
	return engine.Doctor(ctx, request)
}
func History(root, revisionRange string) (HistoryReport, error) {
	return engine.History(root, revisionRange)
}
package model

import (
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

type Case struct {
	Name      string
	Classname string
	IDs       []string
	Result    string
	Failure   string
	Seconds   float64
}

type Report struct {
	Cases    []Case
	Seconds  float64
	HookExit int
	Stdout   string
	Stderr   string
}

type HookRequest struct {
	Snapshot project.Snapshot
	Module   string
	Bin      string
	Timeout  time.Duration
}

type BoundaryRequest struct {
	Snapshot project.Snapshot
	Bin      string
	Timeout  time.Duration
}

type ZoneRequest struct {
	Snapshot          project.Snapshot
	Module            string
	Phase             string
	ChangePath        string
	Base              string
	InterfaceDeclared bool
	Changes           []gitstore.Change
}

type DoctorRequest struct {
	Root    string
	Bin     string
	NoAgent bool
	Timeout time.Duration
}

type DoctorCheck struct {
	Name   string
	Module string
	Passed bool
	Detail string
}

type DoctorReport struct {
	Checks      []DoctorCheck
	Diagnostics []project.Diagnostic
}

type HistoryReport struct {
	Commits     []string
	Squashed    []SquashedChange
	Diagnostics []project.Diagnostic
}

type SquashedChange struct {
	Commit string
	Change string
}
---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
id_prefix: CHK
interface:
    - api.go
    - model/**
name: checks
---

# checks

## Specs

### CHK-001
ParseJUnit MUST accept testsuite or testsuites documents, including nested suites, returning testcase name, classname and time.

### CHK-002
RequireCoverage MUST cover each current spec except review or quarantine with a non-skipped referencing case. Missing coverage MUST return a C10-prefixed error.

### CHK-003
Evaluate MUST run the configured eval hook, accepting shell strings or argv.

### CHK-004
During audit, Zones MUST allow edits only to the current change file.

### CHK-005
Boundaries MUST run applicable global and per-module boundary hooks.

### CHK-006
Zones MUST reject edits to protected paths in every phase.

### CHK-007
Zones MUST reject edits to module SPEC.md files in every phase.

### CHK-008
Zones MUST reject edits within another module in every phase.

### CHK-009
Zones MUST reject edits to change files other than the current change in every phase.

### CHK-010
During evals and impl, Zones MUST allow edits outside modules only when the path matches worker_writable, except for the current change file.

### CHK-011
During evals, Zones MUST allow edits in the target module's evals zone.

### CHK-012
During evals, Zones MUST allow edits in the target module's interface zone only when InterfaceDeclared is true.

### CHK-013
During evals, Zones MUST allow edits in the target module's disposable zone only for files absent from the phase base.

### CHK-014
During evals, Zones MUST allow edits in the target module's other zone.

### CHK-015
During evals, Zones MUST reject edits in the target module's preserve zone.

### CHK-016
During impl, Zones MUST reject edits in the target module's evals zone.

### CHK-017
During impl, Zones MUST reject edits in the target module's interface zone.

### CHK-018
During impl, Zones MUST allow edits in the target module's disposable zone.

### CHK-019
During impl, Zones MUST allow edits in the target module's preserve zone.

### CHK-020
During impl, Zones MUST allow edits in the target module's other zone.

### CHK-021
During evals and impl, Zones MUST allow edits to the current change file.

### CHK-022
During impl, Zones MUST allow only newly added files in the target module's append_only zone.

### CHK-023
During evals, Zones MUST reject edits in the target module's append_only zone.

### CHK-024
Zones MUST check both the source and destination of a rename against phase permissions.

### CHK-025
RequireFailFirst MUST find a failed or errored referencing case for every requested enforced ID except already_satisfied; passing, skipped, empty or unrelated failures do not qualify. Missing genuine fail-first evidence MUST return a P2-01-prefixed error.

### CHK-026
RequirePass MUST reject failed or errored cases referencing any non-quarantined spec with an E01-prefixed error.

### CHK-027
ParseJUnit MUST reject malformed XML with a C11-prefixed error.

### CHK-028
ParseJUnit MUST reject testcases with missing or unknown spec IDs with a C11-prefixed error.

### CHK-029
ParseJUnit MUST map testcase name and classname to current spec IDs using contracts.MatchesID.

### CHK-030
ParseJUnit MUST classify failure and error elements as fail, and skipped elements as skip.

### CHK-031
Evaluate MUST prefer a module's eval hook override over the global eval hook.

### CHK-032
Evaluate MUST expand module, module_dir and report placeholders once in shell strings or argv, quoting substituted shell values.

### CHK-033
Evaluate MUST set the hook's FIREWEED_BIN environment variable to the request Bin.

### CHK-034
Evaluate MUST retain the hook process's stdout, stderr and exit code in its returned report.

### CHK-035
Evaluate MUST parse the hook's written JUnit report into the returned report's cases.

### CHK-036
Evaluate MUST fail with a C11-prefixed error when the eval hook cannot launch or exceeds its timeout.

### CHK-037
Evaluate MUST fail with a C11-prefixed error when the eval hook does not write its JUnit report.

### CHK-038
Evaluate MUST reject a zero exit with mapped failures or a nonzero exit without mapped failures with a C11-prefixed error, while accepting a nonzero exit with mapped failures as valid report evidence.

### CHK-039
Boundaries MUST return errors with diagnostic X01 for unsuppressed module-attributed violations.

### CHK-040
Boundaries MUST return a visible X03 warning containing the reason for a reasoned boundaries_opt_out.

### CHK-041
Boundaries MUST suppress only exact baseline lines from normally exited hooks.

### CHK-042
Boundaries MUST suppress only the opted-out module's violation lines through a reasoned opt-out after a hook exits normally.

### CHK-043
Boundaries MUST report hook launch, cancellation or timeout failures as errors with a diagnostic whose Code is X02. The returned error MUST preserve the hook failure cause without a leading X02 prefix.

### CHK-044
Boundaries MUST fail with diagnostic X02 for a missing hook required by a module with disposable globs and no reasoned opt-out.

### CHK-045
Boundaries MUST warn with diagnostic X04 without an error for a missing hook when no boundary enforcement is required.

### CHK-046
Evaluate MUST return C11-prefixed errors for existing eval-hook setup failures, including missing or invalid hook commands.

### CHK-047
Evaluate MUST return a U02-prefixed error for an unknown requested module.

### CHK-048
Evaluate MUST return C03-prefixed errors for malformed module item syntax and C05-prefixed errors for duplicate module item identities.

### CHK-049
Doctor MUST retain project.Check diagnostics and prevent hook or agent probes when the current project has a static error.

### CHK-050
Doctor MUST run probes using current project files in a detached scratch worktree below the common Git directory's fireweed/worktrees, leaving the caller's checkout unchanged.

### CHK-051
Doctor MUST remove its scratch worktree and Git registration after successful or failed probes.

### CHK-052
Doctor MUST run the configured setup hook in scratch before probes; a setup failure MUST produce C14 naming the setup hook and prevent later probes.

### CHK-053
Doctor MUST render two uniquely named throwaway modules from project-owned module templates and the version-1 canary baseline, substituting name, prefix, package, go_module, module_dir and other_dir.

### CHK-054
Doctor MUST report C14 naming missing or malformed module templates or required canary maps and skip the affected probes.

### CHK-055
Doctor MUST require passing baseline evals and coverage for the throwaway canary modules, and applicable boundaries, before planting canaries and after removing each planted canary. This baseline coverage requirement applies to the canary modules used to prove hooks, rather than every existing module. Recovery failures MUST produce C14 naming the failed hook and canary.

### CHK-056
Doctor MUST plant the supplied undeclared canary and require its boundary hook to report a module-attributed violation; a missed canary MUST produce C14 naming the hook and canary.

### CHK-057
Doctor MUST plant the supplied internal_eval canary and require its boundary hook to report a module-attributed interface violation; a missed canary MUST produce C14 naming the hook and canary.

### CHK-058
Doctor MUST require the supplied failing_eval canary to produce a failed referencing testcase and nonzero eval exit; compile errors, missing reports and unrelated failures MUST NOT count as detection.

### CHK-059
Doctor MUST require the supplied missing_id canary to be rejected for its unmapped testcase ID; unrelated hook or report failures MUST NOT count as detection.

### CHK-060
Doctor MUST test applicable global eval and boundary hooks and every per-module override through their corresponding canary probes.

### CHK-061
Doctor MUST set probe hooks' FIREWEED_BIN to the request Bin and make the scratch canary modules discoverable through that executable's modules output.

### CHK-062
Doctor MUST retain each real module's reasoned X03 boundaries-opt-out warning and applicable X04 warning when no boundary enforcement is required.

### CHK-063
Doctor MUST run the worker and audit commands once each in scratch through the configured agent.wrap with a fixed-reply prompt, requiring exit 0 and the requested reply for success.

### CHK-064
Doctor MUST use agent.command for the audit probe when audit.command is unset and emit a C15 warning identifying the shared-command blind spot.

### CHK-065
Doctor MUST bound each agent probe by the request Timeout or three minutes when unset, and report C15 failures with the command and its final 20 output lines.

### CHK-066
Doctor MUST classify a configured limit_patterns match in an agent probe's final 20 output lines as a C15 usage-limit diagnostic rather than an ordinary command failure.

### CHK-067
Doctor with NoAgent MUST skip worker and audit command execution while retaining hook canary and recovery checks.

### CHK-068
Doctor MUST emit a C16 warning when the AGENTS.md managed block is missing or differs from the current fireweed guidance, while accepting a current block regardless of surrounding text.

### CHK-069
Doctor MUST emit a C16 warning when a project playbook override's embedded_base is missing or differs from the SHA-256 of its current embedded playbook; a matching base MUST avoid that warning.

### CHK-070
History MUST list full commit IDs selected by the supplied Git revision range in oldest-first topological order.
For an empty range it MUST use the merge base of HEAD and the locally resolved configured base through HEAD.

### CHK-071
History MUST identify worker commits from new phase assignments in change-file contents across pending/archive moves, rather than commit subjects.
Deleting a change document without assigning a phase MUST not by itself classify the commit as a worker commit.

### CHK-072
History MUST report H01 when a worker commit changes a protected path or a project playbook override.

### CHK-073
History MUST report H02 when an impl-to-applied phase commit changes evals or interface zones.

### CHK-074
History MUST report H03 when any commit modifies, deletes or renames an existing append_only path, including merge-resolution changes identified by combined Git diffs.

### CHK-075
History MUST report H04 when a commit modifies an existing file under changes/archive.
Deleting an archived change file MUST remain allowed by H04.

### CHK-076
History MUST recognize each new applied phase whose parents lack an impl phase as squashed, listing its commit and change without a warning.
H01 and H02 MUST not apply to those squashed commits.
H03 and H04 MUST still apply.

### CHK-077
History MUST return each rule finding as a Diagnostic with its H01–H04 code, repository-relative Path and full commit ID in Message, using the recorded before/after zone maps and the worker change's projected set values.

### CHK-078
History MUST return an error when a revision range or required recorded project/change document cannot be read or parsed.

### CHK-079
ParseJUnit MUST retain the message attributes and text bodies of JUnit failure and error elements in each affected Case.Failure, including message-only and body-only diagnostics.
A testcase without such text MUST have an empty Failure.

### CHK-080
Evaluate MUST retain parsed Case.Failure strings in its returned Report.Cases, including when failing testcases make the eval hook exit nonzero.

### CHK-081
Doctor MUST persist complete stdout and stderr from every executed worker and audit probe, including failed, timed-out and cancelled probes, under the invoking project's .fireweed/runs/<run>/<phase>-1.log. Each Doctor invocation MUST use a fresh shared run identity, with phases doctor-worker and doctor-audit. Each executed probe's Detail MUST identify its log file by absolute path.

### CHK-082
Doctor MUST expand {prompt}, {prompt_file} and {git_common_dir} once in selected worker/audit command templates and wrapper argv, quoting shell substitutions and preserving argv values literally.

### CHK-083
Doctor MUST create one temporary file containing the exact probe prompt only when a selected command or wrapper template uses {prompt_file}, and expand that placeholder to its absolute path. The file MUST remain available through execution and be removed afterward, including after launch failure, nonzero exit, timeout or cancellation.

## Decisions

[0071-doctor-probes] Doctor requests the literal reply FIREWEED_DOCTOR_OK; limit_patterns match
case-insensitive literal substrings. Project playbook overrides use frontmatter
embedded_base containing the hexadecimal SHA-256 of the embedded playbook text.

[0114-history-validation] A merge merely retaining a phase value already held by a parent is not a new
phase assignment; combined merge diffs still enforce append-only files.
Squashed entries identify the change's repository-relative after-path; ordinary
non-merge phase transitions compare the sole parent.

 succeeded in 581ms:
package engine

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"gopkg.in/yaml.v3"
)

const doctorPrompt = "Reply with exactly FIREWEED_DOCTOR_OK"
const doctorGuidance = "<!-- fireweed:start -->\nThis project uses fireweed. Specs and change files are the source of truth.\nRun `fireweed playbook author-change` before authoring changes and\n`fireweed context <change>` before each step.\n<!-- fireweed:end -->\n"

func doctorDiag(r *model.DoctorReport, code, msg string, warning bool) {
	r.Diagnostics = append(r.Diagnostics, project.Diagnostic{Code: code, Message: msg, Warning: warning})
}
func doctorCheck(r *model.DoctorReport, name, module string, pass bool, detail string) {
	r.Checks = append(r.Checks, model.DoctorCheck{Name: name, Module: module, Passed: pass, Detail: detail})
}
func staticErr(d []project.Diagnostic) bool {
	for _, x := range d {
		if !x.Warning {
			return true
		}
	}
	return false
}
func copyTree(from, to string) error {
	return filepath.Walk(from, func(p string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}
		rel, e := filepath.Rel(from, p)
		if e != nil {
			return e
		}
		if rel == "." {
			return nil
		}
		if rel == ".git" || strings.HasPrefix(rel, ".git"+string(os.PathSeparator)) || rel == filepath.Join(".fireweed", "runs") || strings.HasPrefix(rel, filepath.Join(".fireweed", "runs")+string(os.PathSeparator)) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		dest := filepath.Join(to, rel)
		if info.IsDir() {
			return os.MkdirAll(dest, info.Mode().Perm())
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		if e := os.MkdirAll(filepath.Dir(dest), 0755); e != nil {
			return e
		}
		src, e := os.Open(p)
		if e != nil {
			return e
		}
		defer src.Close()
		dst, e := os.OpenFile(dest, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode().Perm())
		if e != nil {
			return e
		}
		_, e = io.Copy(dst, src)
		closeErr := dst.Close()
		if e != nil {
			return e
		}
		return closeErr
	})
}
func doctorWarnings(root string, r *model.DoctorReport) {
	b, e := os.ReadFile(filepath.Join(root, "AGENTS.md"))
	if e != nil || !strings.Contains(string(b), doctorGuidance) {
		doctorDiag(r, "C16", "AGENTS.md managed guidance missing or stale", true)
	}
	files, _ := filepath.Glob(filepath.Join(root, ".fireweed", "playbooks", "*.md"))
	for _, p := range files {
		b, e := os.ReadFile(p)
		if e != nil {
			continue
		}
		d, e := contracts.ParseDocument(string(b))
		name := strings.TrimSuffix(filepath.Base(p), ".md")
		text, e2 := assets.Playbook(name)
		if e != nil || e2 != nil {
			doctorDiag(r, "C16", name+" playbook embedded_base missing or stale", true)
			continue
		}
		sum := sha256.Sum256([]byte(text))
		if str(d.Frontmatter["embedded_base"]) != hex.EncodeToString(sum[:]) {
			doctorDiag(r, "C16", name+" playbook embedded_base missing or stale", true)
		}
	}
}
func renderCanary(text string, vars map[string]string) string {
	for _, key := range []string{"name", "prefix", "package", "go_module", "module_dir", "other_dir"} {
		text = strings.ReplaceAll(text, "{{"+key+"}}", vars[key])
	}
	return text
}
func writeMap(root, dir string, m map[string]string, vars map[string]string) error {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		p := filepath.Join(root, dir, strings.TrimSuffix(k, ".tmpl"))
		if e := os.MkdirAll(filepath.Dir(p), 0755); e != nil {
			return e
		}
		if e := os.WriteFile(p, []byte(renderCanary(m[k], vars)), 0644); e != nil {
			return e
		}
	}
	return nil
}
func templateFiles(root string) (map[string]string, error) {
	base := filepath.Join(root, ".fireweed", "templates", "module")
	out := map[string]string{}
	e := filepath.Walk(base, func(p string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}
		if info.IsDir() {
			return nil
		}
		rel, e := filepath.Rel(base, p)
		if e != nil {
			return e
		}
		b, e := os.ReadFile(p)
		if e != nil {
			return e
		}
		out[filepath.ToSlash(rel)] = string(b)
		return nil
	})
	if e != nil {
		return nil, e
	}
	if _, ok := out["SPEC.md"]; !ok {
		return nil, fmt.Errorf("template SPEC.md missing")
	}
	return out, nil
}
func canaryMaps(root string) (map[string]map[string]string, error) {
	p := filepath.Join(root, ".fireweed", "canaries", "canaries.yaml")
	b, e := os.ReadFile(p)
	if e != nil {
		return nil, fmt.Errorf("canary %s: %w", p, e)
	}
	var raw struct {
		Version      int               `yaml:"version"`
		Baseline     map[string]string `yaml:"baseline"`
		Undeclared   map[string]string `yaml:"undeclared"`
		InternalEval map[string]string `yaml:"internal_eval"`
		FailingEval  map[string]string `yaml:"failing_eval"`
		MissingID    map[string]string `yaml:"missing_id"`
	}
	if e := yaml.Unmarshal(b, &raw); e != nil {
		return nil, fmt.Errorf("canary canaries.yaml: %w", e)
	}
	if raw.Version != 1 {
		return nil, fmt.Errorf("canary canaries.yaml version %d", raw.Version)
	}
	out := map[string]map[string]string{"baseline": raw.Baseline, "undeclared": raw.Undeclared, "internal_eval": raw.InternalEval, "failing_eval": raw.FailingEval, "missing_id": raw.MissingID}
	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id"} {
		if len(out[name]) == 0 {
			return nil, fmt.Errorf("canary %s missing", name)
		}
	}
	return out, nil
}
func evalProbe(ctx context.Context, s project.Snapshot, m project.Module, hook any, bin string, timeout time.Duration) (model.Report, error) {
	copy := s
	copy.Modules = []project.Module{m}
	if hook != nil {
		copy.Config = map[string]any{"hooks": map[string]any{"eval": hook}}
		m.Document.Frontmatter = map[string]any{}
		copy.Modules[0] = m
	}
	return Evaluate(ctx, model.HookRequest{Snapshot: copy, Module: m.Name, Bin: bin, Timeout: timeout})
}
func boundaryProbe(ctx context.Context, s project.Snapshot, m project.Module, hook any, bin string, timeout time.Duration) ([]project.Diagnostic, error) {
	copy := s
	copy.Modules = []project.Module{m}
	for _, original := range s.Modules {
		if original.Name != m.Name {
			copy.Modules = append(copy.Modules, project.Module{Name: original.Name, Dir: original.Dir, Document: contracts.Document{Frontmatter: map[string]any{}}})
		}
	}
	if hook != nil {
		copy.Config = map[string]any{}
		m.Document.Frontmatter = map[string]any{"hooks": map[string]any{"boundaries": hook}}
		copy.Modules[0] = m
	}
	return Boundaries(ctx, model.BoundaryRequest{Snapshot: copy, Bin: bin, Timeout: timeout})
}
func doctorBaseline(ctx context.Context, s project.Snapshot, m project.Module, evalHook, boundHook any, bin string, timeout time.Duration) error {
	rep, e := evalProbe(ctx, s, m, evalHook, bin, timeout)
	if e != nil {
		return fmt.Errorf("eval hook: %w", e)
	}
	specs, e := moduleSpecs(m)
	if e != nil {
		return fmt.Errorf("eval hook specs: %w", e)
	}
	if rep.HookExit != 0 {
		return fmt.Errorf("eval hook exit %d", rep.HookExit)
	}
	if e = RequireCoverage(rep, specs); e != nil {
		return fmt.Errorf("eval hook coverage: %w", e)
	}
	if e = RequirePass(rep, specs); e != nil {
		return fmt.Errorf("eval hook pass: %w", e)
	}
	if boundHook != nil {
		ds, e := boundaryProbe(ctx, s, m, boundHook, bin, timeout)
		if e != nil {
			return fmt.Errorf("boundaries hook: %w (%v)", e, ds)
		}
	}
	return nil
}
func doctorCanary(ctx context.Context, r *model.DoctorReport, s project.Snapshot, m project.Module, hookName string, evalHook, boundHook any, vars map[string]string, files map[string]string, bin string, timeout time.Duration, name string) {
	if err := writeMap(s.Root, m.Dir, files, vars); err != nil {
		doctorDiag(r, "C14", name+" "+hookName+" canary write: "+err.Error(), false)
		return
	}
	pass := false
	detail := ""
	switch name {
	case "undeclared", "internal_eval":
		ds, e := boundaryProbe(ctx, s, m, boundHook, bin, timeout)
		if e == nil {
			for _, d := range ds {
				if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
					pass = true
				}
			}
		}
		detail = fmt.Sprintf("%v %v", ds, e)
	case "failing_eval":
		rep, e := evalProbe(ctx, s, m, evalHook, bin, timeout)
		if e == nil && rep.HookExit != 0 {
			for _, c := range rep.Cases {
				if c.Result == "fail" && hasID(c.IDs, m.IDPrefix+"-001") {
					pass = true
				}
			}
		}
		detail = fmt.Sprintf("%v %v", rep, e)
	case "missing_id":
		_, e := evalProbe(ctx, s, m, evalHook, bin, timeout)
		if e != nil && strings.Contains(e.Error(), "unknown spec ID") && strings.Contains(e.Error(), "UNMAPPED") {
			pass = true
		}
		detail = fmt.Sprint(e)
	}
	doctorCheck(r, name, m.Name, pass, detail)
	if !pass {
		doctorDiag(r, "C14", name+" missed by "+hookName+" for "+m.Name+": "+detail, false)
	}
	for p := range files {
		_ = os.Remove(filepath.Join(s.Root, m.Dir, p))
	}
	if e := doctorBaseline(ctx, s, m, evalHook, boundHook, bin, timeout); e != nil {
		doctorCheck(r, "recovery "+name, m.Name, false, e.Error())
		doctorDiag(r, "C14", "recovery after "+name+" failed "+hookName+": "+e.Error(), false)
	} else {
		doctorCheck(r, "recovery "+name, m.Name, true, "")
	}
}
func tail20(s string) string {
	lines := strings.Split(strings.TrimRight(s, "\n"), "\n")
	if len(lines) > 20 {
		lines = lines[len(lines)-20:]
	}
	return strings.Join(lines, "\n")
}
func doctorAgent(ctx context.Context, r *model.DoctorReport, req model.DoctorRequest, scratch, common, runID, role string, cmdValue, wrapValue any, patterns []string) {
	vars := map[string]string{"prompt": doctorPrompt, "git_common_dir": common}
	needsFile := strings.Contains(fmt.Sprint(cmdValue), "{prompt_file}") || strings.Contains(fmt.Sprint(wrapValue), "{prompt_file}")
	if needsFile {
		f, e := os.CreateTemp("", "fireweed-doctor-prompt-*")
		if e != nil {
			doctorDiag(r, "C15", role+" prompt file: "+e.Error(), false)
			return
		}
		p := f.Name()
		_, e = f.WriteString(doctorPrompt)
		f.Close()
		if e != nil {
			os.Remove(p)
			doctorDiag(r, "C15", role+" prompt file: "+e.Error(), false)
			return
		}
		vars["prompt_file"] = p
		defer os.Remove(p)
	}
	cmd, e := command(cmdValue, vars)
	if e != nil {
		doctorDiag(r, "C15", role+" command: "+e.Error(), false)
		return
	}
	wrap := []string{}
	if wrapValue != nil {
		wrap, e = command(wrapValue, vars)
		if e != nil {
			doctorDiag(r, "C15", role+" wrap: "+e.Error(), false)
			return
		}
	}
	argv := append(wrap, cmd...)
	limit := req.Timeout
	if limit <= 0 {
		limit = 3 * time.Minute
	}
	out, stderr, exit, runErr := run(ctx, scratch, argv, nil, limit)
	logPath := filepath.Join(req.Root, ".fireweed", "runs", runID, "doctor-"+role+"-1.log")
	_ = os.MkdirAll(filepath.Dir(logPath), 0755)
	_ = os.WriteFile(logPath, []byte(out+stderr), 0644)
	detail := "log: " + logPath
	combined := out + stderr
	tail := tail20(combined)
	passed := runErr == nil && exit == 0 && strings.Contains(out, "FIREWEED_DOCTOR_OK")
	doctorCheck(r, role, "", passed, detail)
	if !passed {
		message := fmt.Sprintf("%s probe command %v failed: %v exit %d\n%s", role, argv, runErr, exit, tail)
		for _, pat := range patterns {
			if pat != "" && strings.Contains(strings.ToLower(tail), strings.ToLower(pat)) {
				message = "usage limit: " + message
				break
			}
		}
		doctorDiag(r, "C15", message, false)
	}
}
func Doctor(ctx context.Context, req model.DoctorRequest) model.DoctorReport {
	var r model.DoctorReport
	r.Diagnostics = append(r.Diagnostics, project.Check(req.Root)...)
	doctorWarnings(req.Root, &r)
	if staticErr(r.Diagnostics) {
		return r
	}
	repo, e := gitstore.Inspect(req.Root)
	if e != nil {
		doctorDiag(&r, "C14", "scratch repository: "+e.Error(), false)
		return r
	}
	common := repo.CommonDir
	if !filepath.IsAbs(common) {
		common = filepath.Join(req.Root, common)
	}
	common, _ = filepath.EvalSymlinks(common)
	base := filepath.Join(common, "fireweed", "worktrees")
	_ = os.MkdirAll(base, 0755)
	scratch, e := os.MkdirTemp(base, "doctor-")
	if e != nil {
		doctorDiag(&r, "C14", "scratch worktree: "+e.Error(), false)
		return r
	}
	_ = os.Remove(scratch)
	if _, e = git(req.Root, "worktree", "add", "--detach", scratch, "HEAD"); e != nil {
		doctorDiag(&r, "C14", "scratch worktree: "+e.Error(), false)
		return r
	}
	defer func() { _, _ = git(req.Root, "worktree", "remove", "--force", scratch); _ = os.RemoveAll(scratch) }()
	if e = copyTree(req.Root, scratch); e != nil {
		doctorDiag(&r, "C14", "scratch copy: "+e.Error(), false)
		return r
	}
	snap, e := project.Load(scratch)
	if e != nil {
		doctorDiag(&r, "C14", "scratch project: "+e.Error(), false)
		return r
	}
	setup := configured(snap.Config, "setup")
	if setup != nil {
		args, e := command(setup, nil)
		if e == nil {
			_, _, exit, runErr := run(ctx, scratch, args, map[string]string{"FIREWEED_BIN": req.Bin}, req.Timeout)
			if runErr != nil || exit != 0 {
				e = fmt.Errorf("exit %d: %v", exit, runErr)
			}
		}
		if e != nil {
			doctorDiag(&r, "C14", "setup hook failed: "+e.Error(), false)
			return r
		}
	}
	for _, m := range snap.Modules {
		if reason := boundaryOpt(m); reason != "" {
			doctorDiag(&r, "X03", m.Name+": boundaries opt out: "+reason, true)
		}
		if boundaryHook(snap, &m) == nil && (len(m.Zones.Disposable) == 0 || reason != "") {
			doctorDiag(&r, "X04", m.Name+": no boundary enforcement required", true)
		}
	}
	if ds, boundaryErr := Boundaries(ctx, model.BoundaryRequest{Snapshot: snap, Bin: req.Bin, Timeout: req.Timeout}); boundaryErr != nil {
		r.Diagnostics = append(r.Diagnostics, ds...)
		doctorDiag(&r, "C14", "baseline boundaries hook: "+boundaryErr.Error(), false)
		doctorDiag(&r, "C14", "undeclared boundaries hook unavailable", false)
		doctorDiag(&r, "C14", "internal_eval boundaries hook unavailable", false)
		return r
	}
	templates, e := templateFiles(scratch)
	if e != nil {
		doctorDiag(&r, "C14", "module template SPEC.md: "+e.Error(), false)
		return r
	}
	maps, e := canaryMaps(scratch)
	if e != nil {
		doctorDiag(&r, "C14", e.Error(), false)
		return r
	}
	modText := ""
	if b, e := os.ReadFile(filepath.Join(scratch, "go.mod")); e == nil {
		for _, l := range strings.Split(string(b), "\n") {
			if strings.HasPrefix(l, "module ") {
				modText = strings.TrimSpace(strings.TrimPrefix(l, "module "))
				break
			}
		}
	}
	names := []string{"doctorprobeone", "doctorprobetwo"}
	dirs := []string{"packages/" + names[0], "packages/" + names[1]}
	var canMods []project.Module
	for i, name := range names {
		vars := map[string]string{"name": name, "prefix": fmt.Sprintf("DCT%d", i+1), "package": name, "go_module": modText, "module_dir": dirs[i], "other_dir": dirs[1-i]}
		if e := writeMap(scratch, dirs[i], templates, vars); e != nil {
			doctorDiag(&r, "C14", "module template: "+e.Error(), false)
			return r
		}
		if _, e := contracts.ParseDocument(renderCanary(templates["SPEC.md"], vars)); e != nil {
			doctorDiag(&r, "C14", "module template SPEC.md: "+e.Error(), false)
			return r
		}
		if e := writeMap(scratch, dirs[i], maps["baseline"], vars); e != nil {
			doctorDiag(&r, "C14", "canary baseline: "+e.Error(), false)
			return r
		}
		b, e := os.ReadFile(filepath.Join(scratch, dirs[i], "SPEC.md"))
		if e != nil {
			doctorDiag(&r, "C14", "module template SPEC.md: "+e.Error(), false)
			return r
		}
		d, e := contracts.ParseDocument(string(b))
		if e != nil {
			doctorDiag(&r, "C14", "module template SPEC.md: "+e.Error(), false)
			return r
		}
		m := project.Module{Name: name, IDPrefix: vars["prefix"], Dir: dirs[i], Zones: contracts.Zones{Disposable: []string{"internal/**"}, Evals: []string{"evals/**"}, Interface: []string{"api.txt"}}, Document: d}
		canMods = append(canMods, m)
	}
	// Include scratch canaries even when the project lists exact module paths.
	configPath := filepath.Join(scratch, "fireweed.yaml")
	if raw, err := os.ReadFile(configPath); err == nil {
		var cfg map[string]any
		if yaml.Unmarshal(raw, &cfg) == nil {
			modules := strSlice(cfg["modules"])
			modules = append(modules, dirs...)
			cfg["modules"] = modules
			if updated, err := yaml.Marshal(cfg); err == nil {
				_ = os.WriteFile(configPath, updated, 0644)
			}
		}
	}
	globalEval := configured(snap.Config, "eval")
	globalBoundary := configured(snap.Config, "boundaries")
	targets := []struct {
		m           project.Module
		eval, bound any
		label       string
	}{{canMods[0], globalEval, globalBoundary, "global"}, {canMods[1], globalEval, globalBoundary, "global"}}
	for _, m := range snap.Modules {
		eh := configured(m.Document.Frontmatter, "eval")
		bh := configured(m.Document.Frontmatter, "boundaries")
		if eh != nil || bh != nil {
			if eh == nil {
				eh = globalEval
			}
			if bh == nil {
				bh = globalBoundary
			}
			targets = append(targets, struct {
				m           project.Module
				eval, bound any
				label       string
			}{m, eh, bh, m.Name + " eval override"})
		}
	}
	for _, target := range targets {
		m := target.m
		eh, bh := target.eval, target.bound
		if e := doctorBaseline(ctx, snap, m, eh, bh, req.Bin, req.Timeout); e != nil {
			doctorCheck(&r, "baseline", m.Name, false, e.Error())
			doctorDiag(&r, "C14", "baseline "+target.label+" "+e.Error(), false)
			continue
		}
		doctorCheck(&r, "baseline", m.Name, true, "")
		vars := map[string]string{"name": m.Name, "prefix": m.IDPrefix, "package": m.Name, "go_module": modText, "module_dir": m.Dir, "other_dir": dirs[0]}
		if m.Name == names[0] {
			vars["other_dir"] = dirs[1]
		}
		for _, name := range []string{"undeclared", "internal_eval", "failing_eval", "missing_id"} {
			if bh == nil && (name == "undeclared" || name == "internal_eval") {
				continue
			}
			hookName := target.label + " boundary hook"
			if configured(m.Document.Frontmatter, "boundaries") != nil {
				hookName = m.Name + " boundaries override boundary hook"
			}
			if name == "failing_eval" || name == "missing_id" {
				hookName = target.label + " eval hook"
				if configured(m.Document.Frontmatter, "eval") != nil {
					hookName = m.Name + " eval override eval hook"
				}
			}
			doctorCanary(ctx, &r, snap, m, hookName, eh, bh, vars, maps[name], req.Bin, req.Timeout, name)
		}
	}
	if !req.NoAgent {
		runID := fmt.Sprintf("%d", time.Now().UnixNano())
		agent := stringMap(snap.Config["agent"])
		audit := stringMap(snap.Config["audit"])
		wrap := agent["wrap"]
		doctorAgent(ctx, &r, req, scratch, common, runID, "worker", agent["command"], wrap, strSlice(snap.Config["limit_patterns"]))
		auditCmd := audit["command"]
		if auditCmd == nil {
			auditCmd = agent["command"]
			doctorDiag(&r, "C15", "audit shares worker command; audit command blind spot", true)
		}
		doctorAgent(ctx, &r, req, scratch, common, runID, "audit", auditCmd, wrap, strSlice(snap.Config["limit_patterns"]))
	}
	return r
}
package engine

import (
	"context"
	"fmt"
	"os/exec"
	"path"
	"sort"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

func git(root string, args ...string) (string, error) {
	c := exec.CommandContext(context.Background(), "git", append([]string{"-C", root}, args...)...)
	b, e := c.CombinedOutput()
	if e != nil {
		return "", fmt.Errorf("git %v: %w: %s", args, e, b)
	}
	return strings.TrimSpace(string(b)), nil
}
func historyChanges(root, rev string) (map[string]string, error) {
	out := map[string]string{}
	names, e := git(root, "ls-tree", "-r", "--name-only", rev, "--", "changes")
	if e != nil {
		return nil, e
	}
	for _, p := range strings.Split(names, "\n") {
		if p == "" || !strings.HasSuffix(p, ".md") {
			continue
		}
		b, e := gitstore.ReadAt(root, rev, p)
		if e != nil {
			return nil, e
		}
		d, e := contracts.ParseDocument(string(b))
		if e != nil {
			return nil, fmt.Errorf("%s at %s: %w", p, rev, e)
		}
		out[p] = str(d.Frontmatter["phase"])
	}
	return out, nil
}

type hchange struct{ status, old, new string }

func historyDiff(root, from, to string) ([]hchange, error) {
	args := []string{"diff", "--name-status", "-M", "--no-ext-diff"}
	if from == "" {
		args = append(args, emptyTree)
	} else {
		args = append(args, from)
	}
	args = append(args, to)
	s, e := git(root, args...)
	if e != nil {
		return nil, e
	}
	var out []hchange
	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}
		f := strings.Split(line, "\t")
		if len(f) < 2 {
			continue
		}
		c := hchange{status: f[0], new: f[1]}
		if strings.HasPrefix(f[0], "R") || strings.HasPrefix(f[0], "C") {
			if len(f) < 3 {
				continue
			}
			c.old = f[1]
			c.new = f[2]
		}
		out = append(out, c)
	}
	return out, nil
}

const emptyTree = "4b825dc642cb6eb9a060e54bf8d69288fbee4904"

func recordedZone(p string, snaps ...project.Snapshot) string {
	for _, s := range snaps {
		for _, m := range s.Modules {
			if strings.HasPrefix(p, m.Dir+"/") {
				z := contracts.ResolveZone(strings.TrimPrefix(p, m.Dir+"/"), m.Zones)
				if z != "other" {
					return z
				}
			}
		}
	}
	return "other"
}
func historyProtected(p string, snaps ...project.Snapshot) bool {
	if strings.HasPrefix(p, ".fireweed/playbooks/") {
		return true
	}
	for _, s := range snaps {
		if matchAny(strSlice(s.Config["protected"]), p) {
			return true
		}
		for _, m := range s.Modules {
			if strings.HasPrefix(p, m.Dir+"/") {
				rel := strings.TrimPrefix(p, m.Dir+"/")
				if rel == "SPEC.md" || matchAny(m.Zones.Protected, rel) {
					return true
				}
			}
		}
	}
	return false
}
func projectedHistory(root, rev string, s project.Snapshot, changes map[string]string) (project.Snapshot, error) {
	copy := s
	copy.Modules = append([]project.Module(nil), s.Modules...)
	for p := range changes {
		b, e := gitstore.ReadAt(root, rev, p)
		if e != nil {
			return copy, e
		}
		d, e := contracts.ParseDocument(string(b))
		if e != nil {
			return copy, e
		}
		set := stringMap(d.Frontmatter["set"])
		if len(set) == 0 {
			continue
		}
		name := str(d.Frontmatter["module"])
		for i := range copy.Modules {
			if copy.Modules[i].Name != name {
				continue
			}
			z := &copy.Modules[i].Zones
			if v, ok := set["evals"]; ok {
				z.Evals = strSlice(v)
			}
			if v, ok := set["interface"]; ok {
				z.Interface = strSlice(v)
			}
			if v, ok := set["append_only"]; ok {
				z.AppendOnly = strSlice(v)
			}
			if v, ok := set["disposable"]; ok {
				z.Disposable = strSlice(v)
			}
			if v, ok := set["preserve"]; ok {
				z.Preserve = strSlice(v)
			}
			if v, ok := set["protected"]; ok {
				z.Protected = strSlice(v)
			}
		}
	}
	return copy, nil
}
func History(root, revisionRange string) (model.HistoryReport, error) {
	var rep model.HistoryReport
	if revisionRange == "" {
		s, e := project.Load(root)
		if e != nil {
			return rep, e
		}
		base := str(s.Config["base"])
		remote := str(s.Config["remote"])
		ref, e := gitstore.BaseRef(root, remote, base)
		if e != nil {
			return rep, e
		}
		merge, e := git(root, "merge-base", "HEAD", ref)
		if e != nil {
			return rep, e
		}
		revisionRange = merge + "..HEAD"
	}
	list, e := git(root, "rev-list", "--topo-order", "--reverse", revisionRange)
	if e != nil {
		return rep, e
	}
	if list == "" {
		return rep, nil
	}
	rep.Commits = strings.Split(list, "\n")
	for _, oid := range rep.Commits {
		after, e := project.LoadAt(root, oid)
		if e != nil {
			return rep, e
		}
		afterChanges, e := historyChanges(root, oid)
		if e != nil {
			return rep, e
		}
		parentText, e := git(root, "rev-list", "--parents", "-n", "1", oid)
		if e != nil {
			return rep, e
		}
		parents := strings.Fields(parentText)
		if len(parents) > 0 {
			parents = parents[1:]
		}
		var before []project.Snapshot
		var beforeChanges []map[string]string
		for _, parent := range parents {
			s, e := project.LoadAt(root, parent)
			if e != nil {
				return rep, e
			}
			before = append(before, s)
			m, e := historyChanges(root, parent)
			if e != nil {
				return rep, e
			}
			beforeChanges = append(beforeChanges, m)
		}
		var diffs []hchange
		if len(parents) == 0 {
			diffs, e = historyDiff(root, "", oid)
			if e != nil {
				return rep, e
			}
		} else {
			counts := map[string]int{}
			choices := map[string]hchange{}
			for _, parent := range parents {
				ds, err := historyDiff(root, parent, oid)
				if err != nil {
					return rep, err
				}
				once := map[string]bool{}
				for _, d := range ds {
					key := d.new
					if d.old != "" {
						key = d.old + "|" + d.new
					}
					if !once[key] {
						counts[key]++
						once[key] = true
						choices[key] = d
					}
				}
			}
			for key, count := range counts {
				if len(parents) == 1 || count == len(parents) {
					diffs = append(diffs, choices[key])
				}
			}
			sort.Slice(diffs, func(i, j int) bool { return diffs[i].new < diffs[j].new })
		}
		for _, d := range diffs {
			if d.status == "D" && path.Base(d.new) == "SPEC.md" {
				return rep, fmt.Errorf("required recorded module document missing: %s at %s", d.new, oid)
			}
		}
		worker := false
		implApplied := false
		squashed := false
		for p, phase := range afterChanges {
			if phase == "" {
				continue
			}
			oldPhase := ""
			held := false
			for _, m := range beforeChanges {
				if x, ok := m[p]; ok {
					oldPhase = x
					held = held || x == phase
				}
				for old, x := range m {
					if path.Base(old) == path.Base(p) {
						if oldPhase == "" {
							oldPhase = x
						}
						held = held || x == phase
					}
				}
			}
			if held {
				continue
			}
			worker = true
			if phase == "applied" {
				if oldPhase == "impl" && len(parents) == 1 {
					implApplied = true
				} else {
					rep.Squashed = append(rep.Squashed, model.SquashedChange{Commit: oid, Change: p})
					squashed = true
				}
			}
		}
		projectedAfter, e := projectedHistory(root, oid, after, afterChanges)
		if e != nil {
			return rep, e
		}
		snaps := append([]project.Snapshot{projectedAfter}, before...)
		add := func(code, p string) {
			rep.Diagnostics = append(rep.Diagnostics, project.Diagnostic{Code: code, Path: p, Message: fmt.Sprintf("%s at %s", code, oid)})
		}
		seenFind := map[string]bool{}
		for _, d := range diffs {
			paths := []string{d.new}
			if d.old != "" {
				paths = append(paths, d.old)
			}
			for _, p := range paths {
				if p == "" {
					continue
				}
				if worker && !squashed && historyProtected(p, snaps...) {
					k := "H01|" + p
					if !seenFind[k] {
						add("H01", p)
						seenFind[k] = true
					}
				}
				if implApplied {
					z := recordedZone(p, snaps...)
					if z == "evals" || z == "interface" {
						k := "H02|" + p
						if !seenFind[k] {
							add("H02", p)
							seenFind[k] = true
						}
					}
				}
				if recordedZone(p, snaps...) == "append_only" {
					existing := false
					for _, parent := range parents {
						if _, e := gitstore.ReadAt(root, parent, p); e == nil {
							existing = true
							break
						}
					}
					if existing {
						k := "H03|" + p
						if !seenFind[k] {
							add("H03", p)
							seenFind[k] = true
						}
					}
				}
				if strings.HasPrefix(p, "changes/archive/") && d.status != "D" {
					existing := false
					for _, parent := range parents {
						if _, e := gitstore.ReadAt(root, parent, p); e == nil {
							existing = true
							break
						}
					}
					if existing {
						k := "H04|" + p
						if !seenFind[k] {
							add("H04", p)
							seenFind[k] = true
						}
					}
				}
			}
		}
	}
	return rep, nil
}

 exited 1 in 583ms:
# Checks rebuild notes

## Contradicted Decisions

None identified.

## Gaps

No known behavioral gaps in the target checks module. The supplied workspace omits the Go manifest and dependency implementations, so the retained evals could not be run here. Native verification remains with the outside controller.

 succeeded in 587ms:
package engine

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func pref(code string, err error) error { return fmt.Errorf("%s: %w", code, err) }
func stringMap(v any) map[string]any {
	if m, ok := v.(map[string]any); ok {
		return m
	}
	return nil
}
func str(v any) string { s, _ := v.(string); return s }
func strSlice(v any) []string {
	switch x := v.(type) {
	case []string:
		return x
	case []any:
		var a []string
		for _, v := range x {
			if s, ok := v.(string); ok {
				a = append(a, s)
			}
		}
		return a
	}
	return nil
}
func configured(m map[string]any, key string) any { return stringMap(m["hooks"])[key] }
func shellQuote(s string) string                  { return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'" }
func expand(s string, vals map[string]string, shell bool) string {
	re := regexp.MustCompile(`\{[a-z_]+\}`)
	return re.ReplaceAllStringFunc(s, func(tok string) string {
		v, ok := vals[tok[1:len(tok)-1]]
		if !ok {
			return tok
		}
		if shell {
			return shellQuote(v)
		}
		return v
	})
}
func command(value any, vals map[string]string) ([]string, error) {
	switch x := value.(type) {
	case string:
		if strings.TrimSpace(x) == "" {
			return nil, errors.New("empty hook command")
		}
		return []string{"sh", "-c", expand(x, vals, true)}, nil
	case []string:
		if len(x) == 0 {
			return nil, errors.New("empty hook argv")
		}
		a := make([]string, len(x))
		for i, v := range x {
			a[i] = expand(v, vals, false)
		}
		return a, nil
	case []any:
		if len(x) == 0 {
			return nil, errors.New("empty hook argv")
		}
		a := make([]string, len(x))
		for i, v := range x {
			s, ok := v.(string)
			if !ok {
				return nil, errors.New("hook argv contains non-string")
			}
			a[i] = expand(s, vals, false)
		}
		return a, nil
	default:
		return nil, errors.New("invalid hook command")
	}
}
func run(ctx context.Context, dir string, argv []string, env map[string]string, timeout time.Duration) (string, string, int, error) {
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	c := exec.Command(argv[0], argv[1:]...)
	c.Dir = dir
	c.Env = os.Environ()
	for k, v := range env {
		c.Env = append(c.Env, k+"="+v)
	}
	c.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	var out, errout strings.Builder
	c.Stdout = &out
	c.Stderr = &errout
	if e := c.Start(); e != nil {
		return "", "", -1, e
	}
	done := make(chan error, 1)
	go func() { done <- c.Wait() }()
	var e error
	select {
	case e = <-done:
	case <-ctx.Done():
		_ = syscall.Kill(-c.Process.Pid, syscall.SIGKILL)
		e = <-done
		_ = e
		return out.String(), errout.String(), -1, ctx.Err()
	}
	exit := 0
	if c.ProcessState != nil {
		exit = c.ProcessState.ExitCode()
	}
	if e != nil {
		if _, ok := e.(*exec.ExitError); !ok {
			return out.String(), errout.String(), exit, e
		}
	}
	return out.String(), errout.String(), exit, nil
}

type junitNode struct {
	XMLName   xml.Name
	Name      string      `xml:"name,attr"`
	Classname string      `xml:"classname,attr"`
	Time      string      `xml:"time,attr"`
	Message   string      `xml:"message,attr"`
	Body      string      `xml:",chardata"`
	Children  []junitNode `xml:",any"`
}

var unknownID = regexp.MustCompile(`(?i)[A-Z][A-Z0-9]*[-_]?\d{3,}`)

func ParseJUnit(data []byte, specs []contracts.Item) (model.Report, error) {
	var root junitNode
	dec := xml.NewDecoder(strings.NewReader(string(data)))
	if err := dec.Decode(&root); err != nil {
		return model.Report{}, pref("C11", fmt.Errorf("parse JUnit: %w", err))
	}
	var extra any
	if err := dec.Decode(&extra); err != io.EOF {
		return model.Report{}, pref("C11", fmt.Errorf("parse JUnit: trailing XML"))
	}
	if root.XMLName.Local != "testsuite" && root.XMLName.Local != "testsuites" {
		return model.Report{}, pref("C11", fmt.Errorf("invalid JUnit root %q", root.XMLName.Local))
	}
	known := map[string]bool{}
	for _, s := range specs {
		known[strings.ToUpper(s.ID)] = true
	}
	var report model.Report
	var walk func(junitNode) error
	walk = func(n junitNode) error {
		if n.XMLName.Local == "testcase" {
			c := model.Case{Name: n.Name, Classname: n.Classname, Result: "pass"}
			if n.Time != "" {
				v, e := strconv.ParseFloat(n.Time, 64)
				if e != nil {
					return fmt.Errorf("testcase %q invalid time: %w", n.Name, e)
				}
				c.Seconds = v
			}
			for _, s := range specs {
				if contracts.MatchesID(s.ID, n.Name, n.Classname) {
					c.IDs = append(c.IDs, s.ID)
				}
			}
			if len(c.IDs) == 0 {
				for _, id := range unknownID.FindAllString(n.Name+" "+n.Classname, -1) {
					normalized := strings.ToUpper(strings.ReplaceAll(id, "_", "-"))
					if !known[normalized] {
						candidate := strings.ContainsAny(id, "-_")
						if !candidate {
							for _, spec := range specs {
								prefix := strings.SplitN(spec.ID, "-", 2)[0]
								if strings.HasPrefix(strings.ToUpper(id), strings.ToUpper(prefix)) {
									candidate = true
									break
								}
							}
						}
						if candidate {
							return fmt.Errorf("testcase %q references unknown spec ID %q", n.Name, id)
						}
					}
				}
				return fmt.Errorf("testcase %q has no known spec ID", n.Name)
			}
			var failures []string
			for _, ch := range n.Children {
				switch ch.XMLName.Local {
				case "failure", "error":
					c.Result = "fail"
					if strings.TrimSpace(ch.Message) != "" {
						failures = append(failures, strings.TrimSpace(ch.Message))
					}
					if strings.TrimSpace(ch.Body) != "" {
						failures = append(failures, strings.TrimSpace(ch.Body))
					}
				case "skipped":
					if c.Result != "fail" {
						c.Result = "skip"
					}
				}
			}
			c.Failure = strings.Join(failures, "\n")
			report.Cases = append(report.Cases, c)
			report.Seconds += c.Seconds
			return nil
		}
		for _, ch := range n.Children {
			if err := walk(ch); err != nil {
				return err
			}
		}
		return nil
	}
	if err := walk(root); err != nil {
		return model.Report{}, pref("C11", err)
	}
	return report, nil
}
func hasID(ids []string, id string) bool {
	for _, v := range ids {
		if v == id {
			return true
		}
	}
	return false
}
func RequireCoverage(r model.Report, specs []contracts.Item) error {
	for _, s := range specs {
		tag := strings.ToLower(s.Tag)
		if strings.HasPrefix(tag, "review") || strings.HasPrefix(tag, "quarantine") {
			continue
		}
		found := false
		for _, c := range r.Cases {
			if c.Result != "skip" && hasID(c.IDs, s.ID) {
				found = true
				break
			}
		}
		if !found {
			return pref("C10", fmt.Errorf("no coverage for %s", s.ID))
		}
	}
	return nil
}
func RequireFailFirst(r model.Report, ids, already []string) error {
	for _, id := range ids {
		if hasID(already, id) {
			continue
		}
		found := false
		for _, c := range r.Cases {
			if (c.Result == "fail" || c.Result == "error") && hasID(c.IDs, id) {
				found = true
				break
			}
		}
		if !found {
			return pref("P2-01", fmt.Errorf("no fail-first case for %s", id))
		}
	}
	return nil
}
func RequirePass(r model.Report, specs []contracts.Item) error {
	tags := map[string]string{}
	for _, s := range specs {
		tags[s.ID] = s.Tag
	}
	for _, c := range r.Cases {
		if c.Result != "fail" && c.Result != "error" {
			continue
		}
		for _, id := range c.IDs {
			tag, ok := tags[id]
			if ok && !strings.HasPrefix(strings.ToLower(tag), "quarantine") {
				return pref("E01", fmt.Errorf("failed case %q for %s", c.Name, id))
			}
		}
	}
	return nil
}
func moduleSpecs(m project.Module) ([]contracts.Item, error) {
	section, ok := contracts.Section(m.Document.Body, "Specs")
	if !ok {
		return nil, nil
	}
	items, err := contracts.ParseItems(section)
	if err != nil {
		if code, _, ok := contracts.ErrorCode(err); ok {
			return nil, pref(code, err)
		}
		return nil, err
	}
	return items, nil
}
func Evaluate(ctx context.Context, req model.HookRequest) (model.Report, error) {
	var mod *project.Module
	for i := range req.Snapshot.Modules {
		if req.Snapshot.Modules[i].Name == req.Module {
			mod = &req.Snapshot.Modules[i]
			break
		}
	}
	if mod == nil {
		return model.Report{}, pref("U02", fmt.Errorf("unknown module %q", req.Module))
	}
	specs, err := moduleSpecs(*mod)
	if err != nil {
		return model.Report{}, err
	}
	hook := configured(req.Snapshot.Config, "eval")
	if h := configured(mod.Document.Frontmatter, "eval"); h != nil {
		hook = h
	}
	if hook == nil {
		return model.Report{}, pref("C11", fmt.Errorf("missing eval hook for %s", req.Module))
	}
	f, err := os.CreateTemp("", "fireweed-junit-*.xml")
	if err != nil {
		return model.Report{}, pref("C11", err)
	}
	reportPath := f.Name()
	f.Close()
	os.Remove(reportPath)
	defer os.Remove(reportPath)
	args, err := command(hook, map[string]string{"module": mod.Name, "module_dir": mod.Dir, "report": reportPath})
	if err != nil {
		return model.Report{}, pref("C11", err)
	}
	out, stderr, exit, err := run(ctx, req.Snapshot.Root, args, map[string]string{"FIREWEED_BIN": req.Bin}, req.Timeout)
	if err != nil {
		return model.Report{}, pref("C11", fmt.Errorf("run hook: %w", err))
	}
	data, err := os.ReadFile(reportPath)
	if err != nil {
		return model.Report{}, pref("C11", fmt.Errorf("read JUnit report: %w", err))
	}
	report, err := ParseJUnit(data, specs)
	if err != nil {
		return model.Report{}, err
	}
	report.Stdout = out
	report.Stderr = stderr
	report.HookExit = exit
	failed := false
	for _, c := range report.Cases {
		if c.Result == "fail" {
			failed = true
		}
	}
	if (exit == 0) == failed {
		return report, pref("C11", fmt.Errorf("eval hook exit %d disagrees with JUnit failures", exit))
	}
	return report, nil
}
package engine

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

func matchAny(globs []string, p string) bool {
	for _, g := range globs {
		if contracts.MatchGlob(g, p) {
			return true
		}
	}
	return false
}
func baseExists(root, base, p string) bool {
	if base == "" {
		return false
	}
	_, err := gitstore.ReadAt(root, base, p)
	return err == nil
}
func zonePath(r model.ZoneRequest, p, status string) *project.Diagnostic {
	p = path.Clean(strings.TrimPrefix(p, "./"))
	bad := func(reason string) *project.Diagnostic {
		return &project.Diagnostic{Code: "Z01", Path: p, Message: reason + ": " + p}
	}
	if p == ".fireweed" || strings.HasPrefix(p, ".fireweed/") || matchAny(strSlice(r.Snapshot.Config["protected"]), p) {
		return bad("protected path")
	}
	if p == r.ChangePath && (r.Phase == "audit" || r.Phase == "evals" || r.Phase == "impl") {
		return nil
	}
	if strings.HasPrefix(p, "changes/") {
		return bad("other change file")
	}
	if r.Phase == "audit" {
		return bad("audit edit")
	}
	var owner *project.Module
	for i := range r.Snapshot.Modules {
		m := &r.Snapshot.Modules[i]
		if p == m.Dir || strings.HasPrefix(p, m.Dir+"/") {
			if owner == nil || len(m.Dir) > len(owner.Dir) {
				owner = m
			}
		}
	}
	if owner == nil {
		if matchAny(strSlice(r.Snapshot.Config["worker_writable"]), p) {
			return nil
		}
		return bad("outside worker writable")
	}
	if owner.Name != r.Module {
		return bad("other module")
	}
	rel := strings.TrimPrefix(p, owner.Dir+"/")
	if rel == "SPEC.md" || matchAny(owner.Zones.Protected, rel) {
		return bad("protected module path")
	}
	zone := contracts.ResolveZone(rel, owner.Zones)
	switch r.Phase {
	case "evals":
		switch zone {
		case "evals", "other":
			return nil
		case "interface":
			if r.InterfaceDeclared {
				return nil
			}
		case "disposable":
			if !baseExists(r.Snapshot.Root, r.Base, p) {
				return nil
			}
		}
	case "impl":
		switch zone {
		case "other", "disposable", "preserve":
			return nil
		case "append_only":
			if strings.HasPrefix(status, "A") && !baseExists(r.Snapshot.Root, r.Base, p) {
				return nil
			}
		}
	}
	return bad("phase zone")
}
func Zones(r model.ZoneRequest) []project.Diagnostic {
	var d []project.Diagnostic
	for _, c := range r.Changes {
		if c.OldPath != "" {
			if x := zonePath(r, c.OldPath, "D"); x != nil {
				d = append(d, *x)
			}
		}
		if x := zonePath(r, c.Path, c.Status); x != nil {
			d = append(d, *x)
		}
	}
	return d
}
func boundaryHook(s project.Snapshot, m *project.Module) any {
	if m != nil {
		if v := configured(m.Document.Frontmatter, "boundaries"); v != nil {
			return v
		}
	}
	return configured(s.Config, "boundaries")
}
func boundaryLines(text string) []string {
	var a []string
	for _, l := range strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n") {
		if strings.TrimSpace(l) != "" {
			a = append(a, l)
		}
	}
	return a
}
func boundaryOpt(m project.Module) string {
	return strings.TrimSpace(str(m.Document.Frontmatter["boundaries_opt_out"]))
}
func Boundaries(ctx context.Context, r model.BoundaryRequest) ([]project.Diagnostic, error) {
	var diags []project.Diagnostic
	var first error
	global := configured(r.Snapshot.Config, "boundaries")
	baseline := map[string]bool{}
	if b, e := os.ReadFile(filepath.Join(r.Snapshot.Root, ".fireweed", "boundaries-baseline")); e == nil {
		for _, l := range boundaryLines(string(b)) {
			baseline[l] = true
		}
	}
	for _, m := range r.Snapshot.Modules {
		if reason := boundaryOpt(m); reason != "" {
			diags = append(diags, project.Diagnostic{Code: "X03", Warning: true, Message: m.Name + ": boundaries opt out: " + reason})
		}
	}
	type task struct {
		hook   any
		module *project.Module
	}
	var tasks []task
	if global != nil {
		tasks = append(tasks, task{hook: global})
	}
	for i := range r.Snapshot.Modules {
		m := &r.Snapshot.Modules[i]
		h := configured(m.Document.Frontmatter, "boundaries")
		if h != nil {
			tasks = append(tasks, task{hook: h, module: m})
			continue
		}
		if global == nil {
			if len(m.Zones.Disposable) > 0 && boundaryOpt(*m) == "" {
				d := project.Diagnostic{Code: "X02", Message: m.Name + ": missing boundaries hook"}
				diags = append(diags, d)
				if first == nil {
					first = errors.New(d.Message)
				}
			} else {
				diags = append(diags, project.Diagnostic{Code: "X04", Warning: true, Message: m.Name + ": no boundary enforcement required"})
			}
		}
	}
	if len(r.Snapshot.Modules) == 0 && global == nil {
		diags = append(diags, project.Diagnostic{Code: "X04", Warning: true, Message: "no boundary enforcement required"})
	}
	for _, t := range tasks {
		vals := map[string]string{}
		if t.module != nil {
			vals["module"] = t.module.Name
			vals["module_dir"] = t.module.Dir
		}
		args, e := command(t.hook, vals)
		if e != nil {
			diags = append(diags, project.Diagnostic{Code: "X02", Message: e.Error()})
			if first == nil {
				first = e
			}
			continue
		}
		out, stderr, _, e := run(ctx, r.Snapshot.Root, args, map[string]string{"FIREWEED_BIN": r.Bin}, r.Timeout)
		if e != nil {
			diags = append(diags, project.Diagnostic{Code: "X02", Message: fmt.Sprintf("run boundaries hook: %v", e)})
			if first == nil {
				first = e
			}
			continue
		}
		for _, line := range boundaryLines(out + "\n" + stderr) {
			if baseline[line] {
				continue
			}
			owner := ""
			for _, m := range r.Snapshot.Modules {
				if strings.HasPrefix(line, m.Name+":") {
					owner = m.Name
					if boundaryOpt(m) != "" {
						owner = ""
					}
					break
				}
			}
			if owner == "" {
				continue
			}
			diags = append(diags, project.Diagnostic{Code: "X01", Message: line})
			if first == nil {
				first = fmt.Errorf("boundary violation: %s", line)
			}
		}
	}
	return diags, first
}

exec
/bin/zsh -lc "sed -n '1,380p' packages/checks/internal/engine/zones_boundaries.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc "sed -n '1,360p' packages/checks/internal/engine/doctor.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc 'rg --files packages | sort' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc "sed -n '1,360p' packages/checks/internal/engine/core.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 633ms:
package engine

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

func matchAny(globs []string, p string) bool {
	for _, g := range globs {
		if contracts.MatchGlob(g, p) {
			return true
		}
	}
	return false
}
func baseExists(root, base, p string) bool {
	if base == "" {
		return false
	}
	_, err := gitstore.ReadAt(root, base, p)
	return err == nil
}
func zonePath(r model.ZoneRequest, p, status string) *project.Diagnostic {
	p = path.Clean(strings.TrimPrefix(p, "./"))
	bad := func(reason string) *project.Diagnostic {
		return &project.Diagnostic{Code: "Z01", Path: p, Message: reason + ": " + p}
	}
	if p == ".fireweed" || strings.HasPrefix(p, ".fireweed/") || matchAny(strSlice(r.Snapshot.Config["protected"]), p) {
		return bad("protected path")
	}
	if p == r.ChangePath && (r.Phase == "audit" || r.Phase == "evals" || r.Phase == "impl") {
		return nil
	}
	if strings.HasPrefix(p, "changes/") {
		return bad("other change file")
	}
	if r.Phase == "audit" {
		return bad("audit edit")
	}
	var owner *project.Module
	for i := range r.Snapshot.Modules {
		m := &r.Snapshot.Modules[i]
		if p == m.Dir || strings.HasPrefix(p, m.Dir+"/") {
			if owner == nil || len(m.Dir) > len(owner.Dir) {
				owner = m
			}
		}
	}
	if owner == nil {
		if matchAny(strSlice(r.Snapshot.Config["worker_writable"]), p) {
			return nil
		}
		return bad("outside worker writable")
	}
	if owner.Name != r.Module {
		return bad("other module")
	}
	rel := strings.TrimPrefix(p, owner.Dir+"/")
	if rel == "SPEC.md" || matchAny(owner.Zones.Protected, rel) {
		return bad("protected module path")
	}
	zone := contracts.ResolveZone(rel, owner.Zones)
	switch r.Phase {
	case "evals":
		switch zone {
		case "evals", "other":
			return nil
		case "interface":
			if r.InterfaceDeclared {
				return nil
			}
		case "disposable":
			if !baseExists(r.Snapshot.Root, r.Base, p) {
				return nil
			}
		}
	case "impl":
		switch zone {
		case "other", "disposable", "preserve":
			return nil
		case "append_only":
			if strings.HasPrefix(status, "A") && !baseExists(r.Snapshot.Root, r.Base, p) {
				return nil
			}
		}
	}
	return bad("phase zone")
}
func Zones(r model.ZoneRequest) []project.Diagnostic {
	var d []project.Diagnostic
	for _, c := range r.Changes {
		if c.OldPath != "" {
			if x := zonePath(r, c.OldPath, "D"); x != nil {
				d = append(d, *x)
			}
		}
		if x := zonePath(r, c.Path, c.Status); x != nil {
			d = append(d, *x)
		}
	}
	return d
}
func boundaryHook(s project.Snapshot, m *project.Module) any {
	if m != nil {
		if v := configured(m.Document.Frontmatter, "boundaries"); v != nil {
			return v
		}
	}
	return configured(s.Config, "boundaries")
}
func boundaryLines(text string) []string {
	var a []string
	for _, l := range strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n") {
		if strings.TrimSpace(l) != "" {
			a = append(a, l)
		}
	}
	return a
}
func boundaryOpt(m project.Module) string {
	return strings.TrimSpace(str(m.Document.Frontmatter["boundaries_opt_out"]))
}
func Boundaries(ctx context.Context, r model.BoundaryRequest) ([]project.Diagnostic, error) {
	var diags []project.Diagnostic
	var first error
	global := configured(r.Snapshot.Config, "boundaries")
	baseline := map[string]bool{}
	if b, e := os.ReadFile(filepath.Join(r.Snapshot.Root, ".fireweed", "boundaries-baseline")); e == nil {
		for _, l := range boundaryLines(string(b)) {
			baseline[l] = true
		}
	}
	for _, m := range r.Snapshot.Modules {
		if reason := boundaryOpt(m); reason != "" {
			diags = append(diags, project.Diagnostic{Code: "X03", Warning: true, Message: m.Name + ": boundaries opt out: " + reason})
		}
	}
	type task struct {
		hook   any
		module *project.Module
	}
	var tasks []task
	if global != nil {
		tasks = append(tasks, task{hook: global})
	}
	for i := range r.Snapshot.Modules {
		m := &r.Snapshot.Modules[i]
		h := configured(m.Document.Frontmatter, "boundaries")
		if h != nil {
			tasks = append(tasks, task{hook: h, module: m})
			continue
		}
		if global == nil {
			if len(m.Zones.Disposable) > 0 && boundaryOpt(*m) == "" {
				d := project.Diagnostic{Code: "X02", Message: m.Name + ": missing boundaries hook"}
				diags = append(diags, d)
				if first == nil {
					first = errors.New(d.Message)
				}
			} else {
				diags = append(diags, project.Diagnostic{Code: "X04", Warning: true, Message: m.Name + ": no boundary enforcement required"})
			}
		}
	}
	if len(r.Snapshot.Modules) == 0 && global == nil {
		diags = append(diags, project.Diagnostic{Code: "X04", Warning: true, Message: "no boundary enforcement required"})
	}
	for _, t := range tasks {
		vals := map[string]string{}
		if t.module != nil {
			vals["module"] = t.module.Name
			vals["module_dir"] = t.module.Dir
		}
		args, e := command(t.hook, vals)
		if e != nil {
			diags = append(diags, project.Diagnostic{Code: "X02", Message: e.Error()})
			if first == nil {
				first = e
			}
			continue
		}
		out, stderr, _, e := run(ctx, r.Snapshot.Root, args, map[string]string{"FIREWEED_BIN": r.Bin}, r.Timeout)
		if e != nil {
			diags = append(diags, project.Diagnostic{Code: "X02", Message: fmt.Sprintf("run boundaries hook: %v", e)})
			if first == nil {
				first = e
			}
			continue
		}
		for _, line := range boundaryLines(out + "\n" + stderr) {
			if baseline[line] {
				continue
			}
			owner := ""
			for _, m := range r.Snapshot.Modules {
				if strings.HasPrefix(line, m.Name+":") {
					owner = m.Name
					if boundaryOpt(m) != "" {
						owner = ""
					}
					break
				}
			}
			if owner == "" {
				continue
			}
			diags = append(diags, project.Diagnostic{Code: "X01", Message: line})
			if first == nil {
				first = fmt.Errorf("boundary violation: %s", line)
			}
		}
	}
	return diags, first
}

 succeeded in 635ms:
package engine

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func pref(code string, err error) error { return fmt.Errorf("%s: %w", code, err) }
func stringMap(v any) map[string]any {
	if m, ok := v.(map[string]any); ok {
		return m
	}
	return nil
}
func str(v any) string { s, _ := v.(string); return s }
func strSlice(v any) []string {
	switch x := v.(type) {
	case []string:
		return x
	case []any:
		var a []string
		for _, v := range x {
			if s, ok := v.(string); ok {
				a = append(a, s)
			}
		}
		return a
	}
	return nil
}
func configured(m map[string]any, key string) any { return stringMap(m["hooks"])[key] }
func shellQuote(s string) string                  { return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'" }
func expand(s string, vals map[string]string, shell bool) string {
	re := regexp.MustCompile(`\{[a-z_]+\}`)
	return re.ReplaceAllStringFunc(s, func(tok string) string {
		v, ok := vals[tok[1:len(tok)-1]]
		if !ok {
			return tok
		}
		if shell {
			return shellQuote(v)
		}
		return v
	})
}
func command(value any, vals map[string]string) ([]string, error) {
	switch x := value.(type) {
	case string:
		if strings.TrimSpace(x) == "" {
			return nil, errors.New("empty hook command")
		}
		return []string{"sh", "-c", expand(x, vals, true)}, nil
	case []string:
		if len(x) == 0 {
			return nil, errors.New("empty hook argv")
		}
		a := make([]string, len(x))
		for i, v := range x {
			a[i] = expand(v, vals, false)
		}
		return a, nil
	case []any:
		if len(x) == 0 {
			return nil, errors.New("empty hook argv")
		}
		a := make([]string, len(x))
		for i, v := range x {
			s, ok := v.(string)
			if !ok {
				return nil, errors.New("hook argv contains non-string")
			}
			a[i] = expand(s, vals, false)
		}
		return a, nil
	default:
		return nil, errors.New("invalid hook command")
	}
}
func run(ctx context.Context, dir string, argv []string, env map[string]string, timeout time.Duration) (string, string, int, error) {
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	c := exec.Command(argv[0], argv[1:]...)
	c.Dir = dir
	c.Env = os.Environ()
	for k, v := range env {
		c.Env = append(c.Env, k+"="+v)
	}
	c.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	var out, errout strings.Builder
	c.Stdout = &out
	c.Stderr = &errout
	if e := c.Start(); e != nil {
		return "", "", -1, e
	}
	done := make(chan error, 1)
	go func() { done <- c.Wait() }()
	var e error
	select {
	case e = <-done:
	case <-ctx.Done():
		_ = syscall.Kill(-c.Process.Pid, syscall.SIGKILL)
		e = <-done
		_ = e
		return out.String(), errout.String(), -1, ctx.Err()
	}
	exit := 0
	if c.ProcessState != nil {
		exit = c.ProcessState.ExitCode()
	}
	if e != nil {
		if _, ok := e.(*exec.ExitError); !ok {
			return out.String(), errout.String(), exit, e
		}
	}
	return out.String(), errout.String(), exit, nil
}

type junitNode struct {
	XMLName   xml.Name
	Name      string      `xml:"name,attr"`
	Classname string      `xml:"classname,attr"`
	Time      string      `xml:"time,attr"`
	Message   string      `xml:"message,attr"`
	Body      string      `xml:",chardata"`
	Children  []junitNode `xml:",any"`
}

var unknownID = regexp.MustCompile(`(?i)[A-Z][A-Z0-9]*[-_]?\d{3,}`)

func ParseJUnit(data []byte, specs []contracts.Item) (model.Report, error) {
	var root junitNode
	dec := xml.NewDecoder(strings.NewReader(string(data)))
	if err := dec.Decode(&root); err != nil {
		return model.Report{}, pref("C11", fmt.Errorf("parse JUnit: %w", err))
	}
	var extra any
	if err := dec.Decode(&extra); err != io.EOF {
		return model.Report{}, pref("C11", fmt.Errorf("parse JUnit: trailing XML"))
	}
	if root.XMLName.Local != "testsuite" && root.XMLName.Local != "testsuites" {
		return model.Report{}, pref("C11", fmt.Errorf("invalid JUnit root %q", root.XMLName.Local))
	}
	known := map[string]bool{}
	for _, s := range specs {
		known[strings.ToUpper(s.ID)] = true
	}
	var report model.Report
	var walk func(junitNode) error
	walk = func(n junitNode) error {
		if n.XMLName.Local == "testcase" {
			c := model.Case{Name: n.Name, Classname: n.Classname, Result: "pass"}
			if n.Time != "" {
				v, e := strconv.ParseFloat(n.Time, 64)
				if e != nil {
					return fmt.Errorf("testcase %q invalid time: %w", n.Name, e)
				}
				c.Seconds = v
			}
			for _, s := range specs {
				if contracts.MatchesID(s.ID, n.Name, n.Classname) {
					c.IDs = append(c.IDs, s.ID)
				}
			}
			if len(c.IDs) == 0 {
				for _, id := range unknownID.FindAllString(n.Name+" "+n.Classname, -1) {
					normalized := strings.ToUpper(strings.ReplaceAll(id, "_", "-"))
					if !known[normalized] {
						candidate := strings.ContainsAny(id, "-_")
						if !candidate {
							for _, spec := range specs {
								prefix := strings.SplitN(spec.ID, "-", 2)[0]
								if strings.HasPrefix(strings.ToUpper(id), strings.ToUpper(prefix)) {
									candidate = true
									break
								}
							}
						}
						if candidate {
							return fmt.Errorf("testcase %q references unknown spec ID %q", n.Name, id)
						}
					}
				}
				return fmt.Errorf("testcase %q has no known spec ID", n.Name)
			}
			var failures []string
			for _, ch := range n.Children {
				switch ch.XMLName.Local {
				case "failure", "error":
					c.Result = "fail"
					if strings.TrimSpace(ch.Message) != "" {
						failures = append(failures, strings.TrimSpace(ch.Message))
					}
					if strings.TrimSpace(ch.Body) != "" {
						failures = append(failures, strings.TrimSpace(ch.Body))
					}
				case "skipped":
					if c.Result != "fail" {
						c.Result = "skip"
					}
				}
			}
			c.Failure = strings.Join(failures, "\n")
			report.Cases = append(report.Cases, c)
			report.Seconds += c.Seconds
			return nil
		}
		for _, ch := range n.Children {
			if err := walk(ch); err != nil {
				return err
			}
		}
		return nil
	}
	if err := walk(root); err != nil {
		return model.Report{}, pref("C11", err)
	}
	return report, nil
}
func hasID(ids []string, id string) bool {
	for _, v := range ids {
		if v == id {
			return true
		}
	}
	return false
}
func RequireCoverage(r model.Report, specs []contracts.Item) error {
	for _, s := range specs {
		tag := strings.ToLower(s.Tag)
		if strings.HasPrefix(tag, "review") || strings.HasPrefix(tag, "quarantine") {
			continue
		}
		found := false
		for _, c := range r.Cases {
			if c.Result != "skip" && hasID(c.IDs, s.ID) {
				found = true
				break
			}
		}
		if !found {
			return pref("C10", fmt.Errorf("no coverage for %s", s.ID))
		}
	}
	return nil
}
func RequireFailFirst(r model.Report, ids, already []string) error {
	for _, id := range ids {
		if hasID(already, id) {
			continue
		}
		found := false
		for _, c := range r.Cases {
			if (c.Result == "fail" || c.Result == "error") && hasID(c.IDs, id) {
				found = true
				break
			}
		}
		if !found {
			return pref("P2-01", fmt.Errorf("no fail-first case for %s", id))
		}
	}
	return nil
}
func RequirePass(r model.Report, specs []contracts.Item) error {
	tags := map[string]string{}
	for _, s := range specs {
		tags[s.ID] = s.Tag
	}
	for _, c := range r.Cases {
		if c.Result != "fail" && c.Result != "error" {
			continue
		}
		for _, id := range c.IDs {
			tag, ok := tags[id]
			if ok && !strings.HasPrefix(strings.ToLower(tag), "quarantine") {
				return pref("E01", fmt.Errorf("failed case %q for %s", c.Name, id))
			}
		}
	}
	return nil
}
func moduleSpecs(m project.Module) ([]contracts.Item, error) {
	section, ok := contracts.Section(m.Document.Body, "Specs")
	if !ok {
		return nil, nil
	}
	items, err := contracts.ParseItems(section)
	if err != nil {
		if code, _, ok := contracts.ErrorCode(err); ok {
			return nil, pref(code, err)
		}
		return nil, err
	}
	return items, nil
}
func Evaluate(ctx context.Context, req model.HookRequest) (model.Report, error) {
	var mod *project.Module
	for i := range req.Snapshot.Modules {
		if req.Snapshot.Modules[i].Name == req.Module {
			mod = &req.Snapshot.Modules[i]
			break
		}
	}
	if mod == nil {
		return model.Report{}, pref("U02", fmt.Errorf("unknown module %q", req.Module))
	}
	specs, err := moduleSpecs(*mod)
	if err != nil {
		return model.Report{}, err
	}
	hook := configured(req.Snapshot.Config, "eval")
	if h := configured(mod.Document.Frontmatter, "eval"); h != nil {
		hook = h
	}
	if hook == nil {
		return model.Report{}, pref("C11", fmt.Errorf("missing eval hook for %s", req.Module))
	}
	f, err := os.CreateTemp("", "fireweed-junit-*.xml")
	if err != nil {
		return model.Report{}, pref("C11", err)
	}
	reportPath := f.Name()
	f.Close()
	os.Remove(reportPath)
	defer os.Remove(reportPath)
	args, err := command(hook, map[string]string{"module": mod.Name, "module_dir": mod.Dir, "report": reportPath})
	if err != nil {
		return model.Report{}, pref("C11", err)
	}
	out, stderr, exit, err := run(ctx, req.Snapshot.Root, args, map[string]string{"FIREWEED_BIN": req.Bin}, req.Timeout)
	if err != nil {
		return model.Report{}, pref("C11", fmt.Errorf("run hook: %w", err))
	}
	data, err := os.ReadFile(reportPath)
	if err != nil {
		return model.Report{}, pref("C11", fmt.Errorf("read JUnit report: %w", err))
	}
	report, err := ParseJUnit(data, specs)
	if err != nil {
		return model.Report{}, err
	}
	report.Stdout = out

 succeeded in 635ms:
package engine

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"gopkg.in/yaml.v3"
)

const doctorPrompt = "Reply with exactly FIREWEED_DOCTOR_OK"
const doctorGuidance = "<!-- fireweed:start -->\nThis project uses fireweed. Specs and change files are the source of truth.\nRun `fireweed playbook author-change` before authoring changes and\n`fireweed context <change>` before each step.\n<!-- fireweed:end -->\n"

func doctorDiag(r *model.DoctorReport, code, msg string, warning bool) {
	r.Diagnostics = append(r.Diagnostics, project.Diagnostic{Code: code, Message: msg, Warning: warning})
}
func doctorCheck(r *model.DoctorReport, name, module string, pass bool, detail string) {
	r.Checks = append(r.Checks, model.DoctorCheck{Name: name, Module: module, Passed: pass, Detail: detail})
}
func staticErr(d []project.Diagnostic) bool {
	for _, x := range d {
		if !x.Warning {
			return true
		}
	}
	return false
}
func copyTree(from, to string) error {
	return filepath.Walk(from, func(p string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}
		rel, e := filepath.Rel(from, p)
		if e != nil {
			return e
		}
		if rel == "." {
			return nil
		}
		if rel == ".git" || strings.HasPrefix(rel, ".git"+string(os.PathSeparator)) || rel == filepath.Join(".fireweed", "runs") || strings.HasPrefix(rel, filepath.Join(".fireweed", "runs")+string(os.PathSeparator)) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		dest := filepath.Join(to, rel)
		if info.IsDir() {
			return os.MkdirAll(dest, info.Mode().Perm())
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		if e := os.MkdirAll(filepath.Dir(dest), 0755); e != nil {
			return e
		}
		src, e := os.Open(p)
		if e != nil {
			return e
		}
		defer src.Close()
		dst, e := os.OpenFile(dest, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode().Perm())
		if e != nil {
			return e
		}
		_, e = io.Copy(dst, src)
		closeErr := dst.Close()
		if e != nil {
			return e
		}
		return closeErr
	})
}
func doctorWarnings(root string, r *model.DoctorReport) {
	b, e := os.ReadFile(filepath.Join(root, "AGENTS.md"))
	if e != nil || !strings.Contains(string(b), doctorGuidance) {
		doctorDiag(r, "C16", "AGENTS.md managed guidance missing or stale", true)
	}
	files, _ := filepath.Glob(filepath.Join(root, ".fireweed", "playbooks", "*.md"))
	for _, p := range files {
		b, e := os.ReadFile(p)
		if e != nil {
			continue
		}
		d, e := contracts.ParseDocument(string(b))
		name := strings.TrimSuffix(filepath.Base(p), ".md")
		text, e2 := assets.Playbook(name)
		if e != nil || e2 != nil {
			doctorDiag(r, "C16", name+" playbook embedded_base missing or stale", true)
			continue
		}
		sum := sha256.Sum256([]byte(text))
		if str(d.Frontmatter["embedded_base"]) != hex.EncodeToString(sum[:]) {
			doctorDiag(r, "C16", name+" playbook embedded_base missing or stale", true)
		}
	}
}
func renderCanary(text string, vars map[string]string) string {
	for _, key := range []string{"name", "prefix", "package", "go_module", "module_dir", "other_dir"} {
		text = strings.ReplaceAll(text, "{{"+key+"}}", vars[key])
	}
	return text
}
func writeMap(root, dir string, m map[string]string, vars map[string]string) error {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		p := filepath.Join(root, dir, strings.TrimSuffix(k, ".tmpl"))
		if e := os.MkdirAll(filepath.Dir(p), 0755); e != nil {
			return e
		}
		if e := os.WriteFile(p, []byte(renderCanary(m[k], vars)), 0644); e != nil {
			return e
		}
	}
	return nil
}
func templateFiles(root string) (map[string]string, error) {
	base := filepath.Join(root, ".fireweed", "templates", "module")
	out := map[string]string{}
	e := filepath.Walk(base, func(p string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}
		if info.IsDir() {
			return nil
		}
		rel, e := filepath.Rel(base, p)
		if e != nil {
			return e
		}
		b, e := os.ReadFile(p)
		if e != nil {
			return e
		}
		out[filepath.ToSlash(rel)] = string(b)
		return nil
	})
	if e != nil {
		return nil, e
	}
	if _, ok := out["SPEC.md"]; !ok {
		return nil, fmt.Errorf("template SPEC.md missing")
	}
	return out, nil
}
func canaryMaps(root string) (map[string]map[string]string, error) {
	p := filepath.Join(root, ".fireweed", "canaries", "canaries.yaml")
	b, e := os.ReadFile(p)
	if e != nil {
		return nil, fmt.Errorf("canary %s: %w", p, e)
	}
	var raw struct {
		Version      int               `yaml:"version"`
		Baseline     map[string]string `yaml:"baseline"`
		Undeclared   map[string]string `yaml:"undeclared"`
		InternalEval map[string]string `yaml:"internal_eval"`
		FailingEval  map[string]string `yaml:"failing_eval"`
		MissingID    map[string]string `yaml:"missing_id"`
	}
	if e := yaml.Unmarshal(b, &raw); e != nil {
		return nil, fmt.Errorf("canary canaries.yaml: %w", e)
	}
	if raw.Version != 1 {
		return nil, fmt.Errorf("canary canaries.yaml version %d", raw.Version)
	}
	out := map[string]map[string]string{"baseline": raw.Baseline, "undeclared": raw.Undeclared, "internal_eval": raw.InternalEval, "failing_eval": raw.FailingEval, "missing_id": raw.MissingID}
	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id"} {
		if len(out[name]) == 0 {
			return nil, fmt.Errorf("canary %s missing", name)
		}
	}
	return out, nil
}
func evalProbe(ctx context.Context, s project.Snapshot, m project.Module, hook any, bin string, timeout time.Duration) (model.Report, error) {
	copy := s
	copy.Modules = []project.Module{m}
	if hook != nil {
		copy.Config = map[string]any{"hooks": map[string]any{"eval": hook}}
		m.Document.Frontmatter = map[string]any{}
		copy.Modules[0] = m
	}
	return Evaluate(ctx, model.HookRequest{Snapshot: copy, Module: m.Name, Bin: bin, Timeout: timeout})
}
func boundaryProbe(ctx context.Context, s project.Snapshot, m project.Module, hook any, bin string, timeout time.Duration) ([]project.Diagnostic, error) {
	copy := s
	copy.Modules = []project.Module{m}
	for _, original := range s.Modules {
		if original.Name != m.Name {
			copy.Modules = append(copy.Modules, project.Module{Name: original.Name, Dir: original.Dir, Document: contracts.Document{Frontmatter: map[string]any{}}})
		}
	}
	if hook != nil {
		copy.Config = map[string]any{}
		m.Document.Frontmatter = map[string]any{"hooks": map[string]any{"boundaries": hook}}
		copy.Modules[0] = m
	}
	return Boundaries(ctx, model.BoundaryRequest{Snapshot: copy, Bin: bin, Timeout: timeout})
}
func doctorBaseline(ctx context.Context, s project.Snapshot, m project.Module, evalHook, boundHook any, bin string, timeout time.Duration) error {
	rep, e := evalProbe(ctx, s, m, evalHook, bin, timeout)
	if e != nil {
		return fmt.Errorf("eval hook: %w", e)
	}
	specs, e := moduleSpecs(m)
	if e != nil {
		return fmt.Errorf("eval hook specs: %w", e)
	}
	if rep.HookExit != 0 {
		return fmt.Errorf("eval hook exit %d", rep.HookExit)
	}
	if e = RequireCoverage(rep, specs); e != nil {
		return fmt.Errorf("eval hook coverage: %w", e)
	}
	if e = RequirePass(rep, specs); e != nil {
		return fmt.Errorf("eval hook pass: %w", e)
	}
	if boundHook != nil {
		ds, e := boundaryProbe(ctx, s, m, boundHook, bin, timeout)
		if e != nil {
			return fmt.Errorf("boundaries hook: %w (%v)", e, ds)
		}
	}
	return nil
}
func doctorCanary(ctx context.Context, r *model.DoctorReport, s project.Snapshot, m project.Module, hookName string, evalHook, boundHook any, vars map[string]string, files map[string]string, bin string, timeout time.Duration, name string) {
	if err := writeMap(s.Root, m.Dir, files, vars); err != nil {
		doctorDiag(r, "C14", name+" "+hookName+" canary write: "+err.Error(), false)
		return
	}
	pass := false
	detail := ""
	switch name {
	case "undeclared", "internal_eval":
		ds, e := boundaryProbe(ctx, s, m, boundHook, bin, timeout)
		if e == nil {
			for _, d := range ds {
				if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
					pass = true
				}
			}
		}
		detail = fmt.Sprintf("%v %v", ds, e)
	case "failing_eval":
		rep, e := evalProbe(ctx, s, m, evalHook, bin, timeout)
		if e == nil && rep.HookExit != 0 {
			for _, c := range rep.Cases {
				if c.Result == "fail" && hasID(c.IDs, m.IDPrefix+"-001") {
					pass = true
				}
			}
		}
		detail = fmt.Sprintf("%v %v", rep, e)
	case "missing_id":
		_, e := evalProbe(ctx, s, m, evalHook, bin, timeout)
		if e != nil && strings.Contains(e.Error(), "unknown spec ID") && strings.Contains(e.Error(), "UNMAPPED") {
			pass = true
		}
		detail = fmt.Sprint(e)
	}
	doctorCheck(r, name, m.Name, pass, detail)
	if !pass {
		doctorDiag(r, "C14", name+" missed by "+hookName+" for "+m.Name+": "+detail, false)
	}
	for p := range files {
		_ = os.Remove(filepath.Join(s.Root, m.Dir, p))
	}
	if e := doctorBaseline(ctx, s, m, evalHook, boundHook, bin, timeout); e != nil {
		doctorCheck(r, "recovery "+name, m.Name, false, e.Error())
		doctorDiag(r, "C14", "recovery after "+name+" failed "+hookName+": "+e.Error(), false)
	} else {
		doctorCheck(r, "recovery "+name, m.Name, true, "")
	}
}
func tail20(s string) string {
	lines := strings.Split(strings.TrimRight(s, "\n"), "\n")
	if len(lines) > 20 {
		lines = lines[len(lines)-20:]
	}
	return strings.Join(lines, "\n")
}
func doctorAgent(ctx context.Context, r *model.DoctorReport, req model.DoctorRequest, scratch, common, runID, role string, cmdValue, wrapValue any, patterns []string) {
	vars := map[string]string{"prompt": doctorPrompt, "git_common_dir": common}
	needsFile := strings.Contains(fmt.Sprint(cmdValue), "{prompt_file}") || strings.Contains(fmt.Sprint(wrapValue), "{prompt_file}")
	if needsFile {
		f, e := os.CreateTemp("", "fireweed-doctor-prompt-*")
		if e != nil {
			doctorDiag(r, "C15", role+" prompt file: "+e.Error(), false)
			return
		}
		p := f.Name()
		_, e = f.WriteString(doctorPrompt)
		f.Close()
		if e != nil {
			os.Remove(p)
			doctorDiag(r, "C15", role+" prompt file: "+e.Error(), false)
			return
		}
		vars["prompt_file"] = p
		defer os.Remove(p)
	}
	cmd, e := command(cmdValue, vars)
	if e != nil {
		doctorDiag(r, "C15", role+" command: "+e.Error(), false)
		return
	}
	wrap := []string{}
	if wrapValue != nil {
		wrap, e = command(wrapValue, vars)
		if e != nil {
			doctorDiag(r, "C15", role+" wrap: "+e.Error(), false)
			return
		}
	}
	argv := append(wrap, cmd...)
	limit := req.Timeout
	if limit <= 0 {
		limit = 3 * time.Minute
	}
	out, stderr, exit, runErr := run(ctx, scratch, argv, nil, limit)
	logPath := filepath.Join(req.Root, ".fireweed", "runs", runID, "doctor-"+role+"-1.log")
	_ = os.MkdirAll(filepath.Dir(logPath), 0755)
	_ = os.WriteFile(logPath, []byte(out+stderr), 0644)
	detail := "log: " + logPath
	combined := out + stderr
	tail := tail20(combined)
	passed := runErr == nil && exit == 0 && strings.Contains(out, "FIREWEED_DOCTOR_OK")
	doctorCheck(r, role, "", passed, detail)
	if !passed {
		message := fmt.Sprintf("%s probe command %v failed: %v exit %d\n%s", role, argv, runErr, exit, tail)
		for _, pat := range patterns {
			if pat != "" && strings.Contains(strings.ToLower(tail), strings.ToLower(pat)) {
				message = "usage limit: " + message
				break
			}
		}
		doctorDiag(r, "C15", message, false)
	}
}
func Doctor(ctx context.Context, req model.DoctorRequest) model.DoctorReport {
	var r model.DoctorReport
	r.Diagnostics = append(r.Diagnostics, project.Check(req.Root)...)
	doctorWarnings(req.Root, &r)
	if staticErr(r.Diagnostics) {
		return r
	}

 succeeded in 639ms:
packages/assets/api.go
packages/assets/model/model.go
packages/checks/SPEC.md
packages/checks/api.go
packages/checks/evals/boundaries_test.go
packages/checks/evals/doctor_agent_logs_test.go
packages/checks/evals/doctor_prompt_file_test.go
packages/checks/evals/doctor_test.go
packages/checks/evals/history_test.go
packages/checks/evals/hooks_test.go
packages/checks/evals/junit_test.go
packages/checks/evals/zones_test.go
packages/checks/internal/engine/core.go
packages/checks/internal/engine/doctor.go
packages/checks/internal/engine/history.go
packages/checks/internal/engine/zones_boundaries.go
packages/checks/model/model.go
packages/contracts/api.go
packages/contracts/model/model.go
packages/gitstore/api.go
packages/gitstore/model/model.go
packages/project/api.go
packages/project/model/model.go

exec
/bin/zsh -lc 'cat packages/contracts/api.go packages/contracts/model/model.go packages/gitstore/api.go packages/gitstore/model/model.go packages/project/api.go packages/project/model/model.go packages/assets/api.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc 'cat packages/checks/evals/history_test.go packages/checks/evals/boundaries_test.go packages/checks/evals/zones_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc "sed -n '350,450p' packages/checks/internal/engine/core.go; cat packages/checks/evals/junit_test.go packages/checks/evals/hooks_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc 'cat packages/checks/evals/doctor_test.go packages/checks/evals/doctor_agent_logs_test.go packages/checks/evals/doctor_prompt_file_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 616ms:
		return model.Report{}, pref("C11", fmt.Errorf("run hook: %w", err))
	}
	data, err := os.ReadFile(reportPath)
	if err != nil {
		return model.Report{}, pref("C11", fmt.Errorf("read JUnit report: %w", err))
	}
	report, err := ParseJUnit(data, specs)
	if err != nil {
		return model.Report{}, err
	}
	report.Stdout = out
	report.Stderr = stderr
	report.HookExit = exit
	failed := false
	for _, c := range report.Cases {
		if c.Result == "fail" {
			failed = true
		}
	}
	if (exit == 0) == failed {
		return report, pref("C11", fmt.Errorf("eval hook exit %d disagrees with JUnit failures", exit))
	}
	return report, nil
}
package evals_test

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
)

func requireDiagnosticError(t *testing.T, err error, code, cause string) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected %s error containing %q", code, cause)
	}
	if !strings.HasPrefix(err.Error(), code+": ") || !strings.Contains(err.Error(), cause) || strings.HasPrefix(strings.TrimPrefix(err.Error(), code+": "), code+":") {
		t.Errorf("error = %q; want %s prefix and cause %q", err, code, cause)
	}
}

func checkSpecs() []contracts.Item {
	return []contracts.Item{{ID: "CHK-001"}, {ID: "CHK-002"}, {ID: "CHK-003"}}
}

func Test_CHK_001_CHK_029_CHK_030_NestedJUnitIDsResultsAndTime(t *testing.T) {
	data := []byte(`<testsuites><testsuite name="outer"><testsuite name="inner">
<testcase classname="suite.ChK_001" name="passes" time="1.25"/>
<testcase classname="suite" name="test_chk002_failed" time="0.5"><failure message="bad"/></testcase>
<testcase classname="suite" name="test_chk_003_error" time="0.75"><error message="broken"/></testcase>
<testcase classname="CHK-002" name="CHK-001_both" time="0.125"/>
</testsuite></testsuite></testsuites>`)
	report, err := checks.ParseJUnit(data, checkSpecs())
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]checks.Case{
		"passes":             {Name: "passes", Classname: "suite.ChK_001", IDs: []string{"CHK-001"}, Result: "pass", Seconds: 1.25},
		"test_chk002_failed": {Name: "test_chk002_failed", Classname: "suite", IDs: []string{"CHK-002"}, Result: "fail", Seconds: 0.5},
		"test_chk_003_error": {Name: "test_chk_003_error", Classname: "suite", IDs: []string{"CHK-003"}, Result: "fail", Seconds: 0.75},
		"CHK-001_both":       {Name: "CHK-001_both", Classname: "CHK-002", IDs: []string{"CHK-001", "CHK-002"}, Result: "pass", Seconds: 0.125},
	}
	if len(report.Cases) != len(want) {
		t.Fatalf("cases = %#v; want %d", report.Cases, len(want))
	}
	for _, got := range report.Cases {
		sort.Strings(got.IDs)
		got.Failure = ""
		if !reflect.DeepEqual(got, want[got.Name]) {
			t.Errorf("case = %#v; want %#v", got, want[got.Name])
		}
	}
}

func Test_CHK_001_CHK_029_CHK_030_SingleSuiteSkipped(t *testing.T) {
	report, err := checks.ParseJUnit([]byte(`<testsuite><testcase name="CHK-001" classname="unit" time="0.125"><skipped/></testcase></testsuite>`), checkSpecs())
	if err != nil {
		t.Fatal(err)
	}
	if len(report.Cases) != 1 || report.Cases[0].Result != "skip" || report.Cases[0].Seconds != 0.125 || !reflect.DeepEqual(report.Cases[0].IDs, []string{"CHK-001"}) {
		t.Fatalf("skipped case = %#v", report.Cases)
	}
}

func Test_CHK_079_ParseJUnitRetainsFailureAndErrorDetails(t *testing.T) {
	data := []byte(`<testsuite>
<testcase name="CHK-001_failure_combined"><failure message="want 2 &amp; got 3">assertion at example_test.go:17</failure></testcase>
<testcase name="CHK-001_failure_message"><failure message="expected true"/></testcase>
<testcase name="CHK-002_failure_body"><failure>actual value was false</failure></testcase>
<testcase name="CHK-002_error_combined"><error message="panic: nil pointer">stack at example_test.go:29</error></testcase>
<testcase name="CHK-003_error_message"><error message="setup failed"/></testcase>
<testcase name="CHK-003_error_body"><error>missing fixture</error></testcase>
<testcase name="CHK-001_empty_failure"><failure/></testcase>
<testcase name="CHK-002_pass"/>
<testcase name="CHK-003_skip"><skipped/></testcase>
</testsuite>`)
	report, err := checks.ParseJUnit(data, checkSpecs())
	if err != nil {
		t.Fatal(err)
	}
	want := map[string][]string{
		"CHK-001_failure_combined": {"want 2 & got 3", "assertion at example_test.go:17"},
		"CHK-001_failure_message":  {"expected true"},
		"CHK-002_failure_body":     {"actual value was false"},
		"CHK-002_error_combined":   {"panic: nil pointer", "stack at example_test.go:29"},
		"CHK-003_error_message":    {"setup failed"},
		"CHK-003_error_body":       {"missing fixture"},
		"CHK-001_empty_failure":    nil,
		"CHK-002_pass":             nil,
		"CHK-003_skip":             nil,
	}
	if len(report.Cases) != len(want) {
		t.Fatalf("cases = %#v; want %d", report.Cases, len(want))
	}
	for _, got := range report.Cases {
		parts, ok := want[got.Name]
		if !ok {
			t.Errorf("unexpected case %#v", got)
			continue
		}
		if len(parts) == 0 && got.Failure != "" {
			t.Errorf("%s: failure = %q; want empty", got.Name, got.Failure)
		}
		for _, part := range parts {
			if !strings.Contains(got.Failure, part) {
				t.Errorf("%s: failure = %q; missing %q", got.Name, got.Failure, part)
			}
		}
	}
}

func Test_CHK_002_SkippedJUnitCaseDoesNotCover(t *testing.T) {
	report, err := checks.ParseJUnit([]byte(`<testsuite><testcase name="CHK-001" classname="unit" time="0.125"><skipped/></testcase></testsuite>`), checkSpecs())
	if err != nil {
		t.Fatal(err)
	}
	requireDiagnosticError(t, checks.RequireCoverage(report, []contracts.Item{{ID: "CHK-001"}}), "C10", "no coverage for CHK-001")
}

func Test_CHK_027_RejectMalformedXML(t *testing.T) {
	for _, tc := range []struct{ xml, cause string }{
		{`<testsuite><testcase name="CHK-001"></testsuite>`, "parse JUnit:"},
		{`<document/>`, `invalid JUnit root "document"`},
	} {
		_, err := checks.ParseJUnit([]byte(tc.xml), checkSpecs())
		requireDiagnosticError(t, err, "C11", tc.cause)
	}
}

func Test_CHK_028_RejectMissingAndUnknownIDs(t *testing.T) {
	for _, xml := range []string{
		`<testsuite><testcase name="ordinary" classname="unit"/></testsuite>`,
		`<testsuite><testcase name="ordinary" classname="unit"><skipped/></testcase></testsuite>`,
		`<testsuite><testcase name="CHK-999" classname="unit"/></testsuite>`,
		`<testsuite><testcase name="OCHK001" classname="unit"/></testsuite>`,
		`<testsuite><testcase name="CHK0010" classname="unit"/></testsuite>`,
	} {
		_, err := checks.ParseJUnit([]byte(xml), checkSpecs())
		cause := "has no known spec ID"
		if strings.Contains(xml, "CHK-999") {
			cause = `references unknown spec ID "CHK-999"`
		} else if strings.Contains(xml, "CHK0010") {
			cause = `references unknown spec ID "CHK0010"`
		}
		requireDiagnosticError(t, err, "C11", cause)
	}
}

func Test_CHK_002_CoverageRequiresCurrentNonSkippedCases(t *testing.T) {
	specs := []contracts.Item{{ID: "CHK-001"}, {ID: "CHK-002"}}
	requireDiagnosticError(t, checks.RequireCoverage(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001"}, Result: "pass"}}}, specs), "C10", "no coverage for CHK-002")
	if err := checks.RequireCoverage(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001", "CHK-002"}, Result: "fail"}}}, specs); err != nil {
		t.Errorf("failed referencing case should cover both IDs: %v", err)
	}
}

func Test_CHK_002_CoverageExemptsReviewAndQuarantine(t *testing.T) {
	specs := []contracts.Item{{ID: "CHK-001"}, {ID: "CHK-002", Tag: "review"}, {ID: "CHK-003", Tag: "quarantine: flaky"}}
	if err := checks.RequireCoverage(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001"}, Result: "pass"}}}, specs); err != nil {
		t.Errorf("review/quarantine required coverage: %v", err)
	}
}

func Test_CHK_025_FailFirstRequiresEachRequestedFailure(t *testing.T) {
	ids := []string{"CHK-001", "CHK-002"}
	bad := []checks.Case{
		{IDs: []string{"CHK-001"}, Result: "pass"},
		{IDs: []string{"CHK-002"}, Result: "skip"},
		{IDs: []string{"CHK-003"}, Result: "fail"},
	}
	requireDiagnosticError(t, checks.RequireFailFirst(checks.Report{Cases: bad}, ids, nil), "P2-01", "no fail-first case for CHK-001")
	requireDiagnosticError(t, checks.RequireFailFirst(checks.Report{}, ids, nil), "P2-01", "no fail-first case for CHK-001")
	requireDiagnosticError(t, checks.RequireFailFirst(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001"}, Result: "fail"}, {IDs: []string{"CHK-002"}, Result: "skip"}}}, ids, nil), "P2-01", "no fail-first case for CHK-002")
	good := checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-001"}, Result: "fail"}, {IDs: []string{"CHK-002"}, Result: "fail"}}}
	if err := checks.RequireFailFirst(good, ids, nil); err != nil {
		t.Errorf("both requested IDs failed: %v", err)
	}
}

func Test_CHK_025_FailFirstAlreadySatisfiedExemption(t *testing.T) {
	report := checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-002"}, Result: "fail"}}}
	if err := checks.RequireFailFirst(report, []string{"CHK-001", "CHK-002"}, []string{"CHK-001"}); err != nil {
		t.Errorf("already_satisfied ID incorrectly required: %v", err)
	}
}

func Test_CHK_026_PassRejectsNonQuarantinedFailures(t *testing.T) {
	specs := []contracts.Item{{ID: "CHK-001"}, {ID: "CHK-002", Tag: "review"}, {ID: "CHK-003", Tag: "quarantine: flaky"}}
	for _, id := range []string{"CHK-001", "CHK-002"} {
		requireDiagnosticError(t, checks.RequirePass(checks.Report{Cases: []checks.Case{{IDs: []string{id}, Result: "fail"}}}, specs), "E01", `failed case "" for `+id)
	}
	if err := checks.RequirePass(checks.Report{Cases: []checks.Case{{IDs: []string{"CHK-003"}, Result: "fail"}}}, specs); err != nil {
		t.Errorf("quarantined failure rejected: %v", err)
	}
}
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

func Test_CHK_080_EvaluateRetainsParsedFailureDetails(t *testing.T) {
	for _, tc := range []struct {
		name, xml, exit, wantFailure string
		wantExit                     int
	}{
		{"nonzero failure", `<testsuite><testcase name="CHK-003_assertion"><failure message="want 2">got 3 at example_test.go:17</failure></testcase></testsuite>`, "7", "got 3 at example_test.go:17", 7},
		{"zero passing", `<testsuite><testcase name="CHK-003_pass"/></testsuite>`, "0", "", 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			snap := hookSnapshot(t, `printf '%s' '`+tc.xml+`' > {report}; exit `+tc.exit, nil)
			report, err := checks.Evaluate(context.Background(), checks.HookRequest{Snapshot: snap, Module: "alpha", Timeout: 2 * time.Second})
			if err != nil {
				t.Fatal(err)
			}
			if len(report.Cases) != 1 {
				t.Fatalf("cases = %#v; want one", report.Cases)
			}
			if report.HookExit != tc.wantExit {
				t.Errorf("hook exit = %d; want %s", report.HookExit, tc.exit)
			}
			if !strings.Contains(report.Cases[0].Failure, tc.wantFailure) {
				t.Errorf("failure = %q; missing %q", report.Cases[0].Failure, tc.wantFailure)
			}
			if tc.wantFailure == "" && report.Cases[0].Failure != "" {
				t.Errorf("passing case failure = %q; want empty", report.Cases[0].Failure)
			}
			if tc.wantFailure != "" && !strings.Contains(report.Cases[0].Failure, "want 2") {
				t.Errorf("failure = %q; missing message", report.Cases[0].Failure)
			}
		})
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

 succeeded in 614ms:
package evals_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"gopkg.in/yaml.v3"
)

const doctorGuidance = "<!-- fireweed:start -->\nThis project uses fireweed. Specs and change files are the source of truth.\nRun `fireweed playbook author-change` before authoring changes and\n`fireweed context <change>` before each step.\n<!-- fireweed:end -->\n"

type doctorFixture struct {
	t    *testing.T
	root string
	log  string
	bin  string
}

func (f doctorFixture) put(name, body string) {
	f.t.Helper()
	path := filepath.Join(f.root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		f.t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		f.t.Fatal(err)
	}
}

func (f doctorFixture) git(args ...string) string {
	f.t.Helper()
	cmd := exec.Command("git", args...)
	cmd.Dir = f.root
	cmd.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_TERMINAL_PROMPT=0")
	out, err := cmd.CombinedOutput()
	if err != nil {
		f.t.Fatalf("git %v: %v: %s", args, err, out)
	}
	return string(out)
}

func (f doctorFixture) commit() {
	f.t.Helper()
	f.git("init", "-b", "main")
	f.git("config", "user.name", "Doctor Fixture")
	f.git("config", "user.email", "doctor@example.test")
	f.git("add", ".")
	f.git("commit", "-m", "fixture")
}

func (f doctorFixture) request(noAgent bool) checks.DoctorRequest {
	return checks.DoctorRequest{Root: f.root, Bin: f.bin, NoAgent: noAgent, Timeout: 15 * time.Second}
}

func doctorConfig(t *testing.T, recipe assets.Recipe) string {
	t.Helper()
	config := map[string]any{
		"fireweed": ">=0.1, <0.2", "version": 1, "base": "main", "modules": []string{"packages/*"},
		"limit_patterns":  []string{"rate limit"},
		"module_defaults": map[string]any{"interface": recipe.ModuleDefaults.Interface, "evals": recipe.ModuleDefaults.Evals, "disposable": recipe.ModuleDefaults.Disposable},
		"hooks":           recipe.Hooks,
		"agent":           map[string]any{"command": []string{"bash", ".fireweed/agent.sh", "{prompt}", "{git_common_dir}"}, "wrap": []string{"bash", ".fireweed/wrap.sh"}},
		"audit":           map[string]any{"enabled": true, "command": []string{"bash", ".fireweed/audit.sh", "{prompt}"}},
	}
	data, err := yaml.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func (f doctorFixture) editConfig(change func(map[string]any)) {
	f.t.Helper()
	var config map[string]any
	if err := yaml.Unmarshal(mustRead(f.t, filepath.Join(f.root, "fireweed.yaml")), &config); err != nil {
		f.t.Fatal(err)
	}
	change(config)
	data, err := yaml.Marshal(config)
	if err != nil {
		f.t.Fatal(err)
	}
	f.put("fireweed.yaml", string(data))
}

func newDoctorFixture(t *testing.T) doctorFixture {
	t.Helper()
	r, err := assets.BuiltinRecipe("go")
	if err != nil {
		t.Fatal(err)
	}
	f := doctorFixture{t: t, root: t.TempDir(), log: filepath.Join(t.TempDir(), "probes.log")}
	f.put("go.mod", "module example.test/doctorfixture\n\ngo 1.23\n")
	f.put("fireweed.yaml", doctorConfig(t, r))
	f.put("AGENTS.md", "Before.\n"+doctorGuidance+"After.\n")
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nThe fixture MUST pass.\n")
	f.put("packages/alpha/api.go", "package alpha\n")
	for path, body := range r.Files {
		f.put(path, body)
	}
	f.put(".fireweed/wrap.sh", fmt.Sprintf("#!/bin/bash\nprintf 'wrap:%%s\n' \"$PWD\" >> %q\nexec \"$@\"\n", f.log))
	f.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/bash\nprintf 'worker:%%s:%%s:%%s\n' \"$PWD\" \"$1\" \"$2\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	f.put(".fireweed/audit.sh", fmt.Sprintf("#!/bin/bash\nprintf 'audit:%%s:%%s\n' \"$PWD\" \"$1\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	return f
}

func localDoctorFixture(t *testing.T) doctorFixture {
	t.Helper()
	f := doctorFixture{t: t, root: t.TempDir(), log: filepath.Join(t.TempDir(), "probes.log")}
	f.bin = filepath.Join(t.TempDir(), "fake-fireweed")
	if err := os.WriteFile(f.bin, []byte(`#!/bin/sh
if test "$1" != modules || test "$2" != --json; then exit 2; fi
printf '{"data":['
first=1
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test "$first" = 0; then printf ','; fi; first=0
  printf '{"name":"%s","dir":"%s","depends_on":[],"zones":{"interface":["api.txt"],"evals":["evals/**"],"disposable":["internal/**"]}}' "$name" "$dir"
done
printf ']}'
`), 0700); err != nil {
		t.Fatal(err)
	}
	recipe := assets.Recipe{ModuleDefaults: structZones(), Hooks: map[string]any{
		"eval":       "sh .fireweed/eval.sh {module} {module_dir} {report}",
		"boundaries": "sh .fireweed/boundary.sh",
		"setup":      "sh .fireweed/setup.sh",
	}}
	f.put("fireweed.yaml", doctorConfig(t, recipe))
	f.put(".gitignore", ".fireweed/runs/\n")
	f.put("AGENTS.md", doctorGuidance)
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nThe fixture MUST pass.\n")
	f.put("packages/alpha/api.txt", "public\n")
	f.put(".fireweed/templates/module/SPEC.md", "---\nname: {{name}}\nid_prefix: {{prefix}}\ndepends_on: []\n---\n# {{name}}\n\n## Specs\n\n### {{prefix}}-001\nThe canary MUST pass.\n")
	f.put(".fireweed/templates/module/api.txt.tmpl", "{{package}} {{go_module}} {{module_dir}} {{other_dir}}\n")
	f.put(".fireweed/canaries/canaries.yaml", `version: 1
baseline:
  SPEC.md: "---\nname: {{name}}\nid_prefix: {{prefix}}\ndepends_on: []\n---\n# {{name}}\n\n## Specs\n\n### {{prefix}}-001\nThe canary MUST pass.\n"
  api.txt: "{{package}} {{go_module}} {{module_dir}} {{other_dir}}"
undeclared:
  undeclared.txt: "{{other_dir}}"
internal_eval:
  internal/private.txt: "private"
  evals/internal.txt: "{{module_dir}}"
failing_eval:
  evals/failing.txt: "{{prefix}}-001"
missing_id:
  evals/missing.txt: "UNMAPPED-999"
`)
	f.put(".fireweed/setup.sh", fmt.Sprintf("#!/bin/sh\nprintf 'setup:%%s\n' \"$PWD\" >> %q\n", f.log))
	f.put(".fireweed/eval.sh", fmt.Sprintf(`#!/bin/sh
printf 'eval:%%s:%%s:%%s\n' "$PWD" "$1" "$FIREWEED_BIN" >> %q
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
if test -f "$2/evals/failing.txt"; then printf '<testsuite><testcase name="%%s-001"><failure/></testcase></testsuite>' "$prefix" > "$3"; exit 1; fi
if test -f "$2/evals/missing.txt"; then printf '<testsuite><testcase name="UNMAPPED-999"/></testsuite>' > "$3"; exit 0; fi
printf '<testsuite><testcase name="%%s-001"/></testsuite>' "$prefix" > "$3"
`, f.log))
	f.put(".fireweed/boundary.sh", fmt.Sprintf(`#!/bin/sh
printf 'boundaries:%%s:%%s\n' "$PWD" "$FIREWEED_BIN" >> %q
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test -f "$dir/undeclared.txt"; then printf '%%s: undeclared import\n' "$name"; fi
  if test -f "$dir/evals/internal.txt"; then printf '%%s: interface violation\n' "$name"; fi
done
`, f.log))
	f.put(".fireweed/wrap.sh", fmt.Sprintf("#!/bin/bash\nprintf 'wrap:%%s\n' \"$PWD\" >> %q\nexec \"$@\"\n", f.log))
	f.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/bash\nprintf 'worker:%%s:%%s:%%s\n' \"$PWD\" \"$1\" \"$2\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	f.put(".fireweed/audit.sh", fmt.Sprintf("#!/bin/bash\nprintf 'audit:%%s:%%s\n' \"$PWD\" \"$1\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	return f
}

func structZones() (z contracts.Zones) {
	z.Interface = []string{"api.txt"}
	z.Evals = []string{"evals/**"}
	z.Disposable = []string{"internal/**"}
	return
}

func doctorText(report checks.DoctorReport) string {
	return fmt.Sprintf("checks=%+v diagnostics=%+v", report.Checks, report.Diagnostics)
}
func hasDoctorDiagnostic(report checks.DoctorReport, code string, warning bool, terms ...string) bool {
	for _, d := range report.Diagnostics {
		if d.Code != code || d.Warning != warning {
			continue
		}
		ok := true
		for _, s := range terms {
			if !strings.Contains(strings.ToLower(d.Message), strings.ToLower(s)) {
				ok = false
			}
		}
		if ok {
			return true
		}
	}
	return false
}
func hasDoctorCheck(report checks.DoctorReport, name string, passed bool) bool {
	for _, c := range report.Checks {
		if strings.Contains(c.Name, name) && c.Passed == passed {
			return true
		}
	}
	return false
}
func doctorErrors(report checks.DoctorReport) bool {
	for _, d := range report.Diagnostics {
		if !d.Warning {
			return true
		}
	}
	return false
}
func doctorLog(f doctorFixture) string { b, _ := os.ReadFile(f.log); return string(b) }
func worktrees(f doctorFixture) string { return f.git("worktree", "list", "--porcelain") }

func logDoctorCanaryProbes(t *testing.T, f doctorFixture) {
	t.Helper()
	logProbe := fmt.Sprintf(`printf 'probe:%%s\n' "$PWD" >> %q
for marker in undeclared.txt internal.txt failing.txt missing.txt; do
  if find packages -name "$marker" | grep -q .; then printf 'planted:%%s\n' "$marker" >> %q; fi
done
`, f.log, f.log)
	for _, path := range []string{".fireweed/eval.sh", ".fireweed/boundary.sh"} {
		body := string(mustRead(t, filepath.Join(f.root, path)))
		if !strings.HasPrefix(body, "#!/bin/sh\n") {
			t.Fatalf("fixture hook %s has no shell header", path)
		}
		f.put(path, strings.Replace(body, "#!/bin/sh\n", "#!/bin/sh\n"+logProbe, 1))
	}
}

func assertScratchRemoved(t *testing.T, f doctorFixture) {
	t.Helper()
	common := strings.TrimSpace(f.git("rev-parse", "--git-common-dir"))
	if !filepath.IsAbs(common) {
		common = filepath.Join(f.root, common)
	}
	common, err := filepath.EvalSymlinks(common)
	if err != nil {
		t.Fatal(err)
	}
	base := filepath.Join(common, "fireweed", "worktrees")
	for _, line := range strings.Split(doctorLog(f), "\n") {
		if !strings.HasPrefix(line, "setup:") {
			continue
		}
		path := strings.TrimPrefix(line, "setup:")
		for parent := path; ; parent = filepath.Dir(parent) {
			resolved, err := filepath.EvalSymlinks(parent)
			if err == nil {
				suffix, err := filepath.Rel(parent, path)
				if err != nil {
					t.Fatal(err)
				}
				path = filepath.Join(resolved, suffix)
				break
			}
			if parent == filepath.Dir(parent) {
				t.Fatal(err)
			}
		}
		rel, err := filepath.Rel(base, path)
		if err != nil || rel == "." || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
			t.Errorf("scratch %q is not below common Git directory %q", path, base)
		}
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Errorf("physical scratch directory %q remains or cannot be checked: %v", path, err)
		}
		return
	}
	t.Error("setup did not record a scratch path")
}

func Test_CHK_049_StaticDiagnosticsStopProbes(t *testing.T) {
	f := localDoctorFixture(t)
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n## Specs\n### ALP-001\nA.\n### ALP-001\nB.\n")
	f.commit()
	want := project.Check(f.root)
	if len(want) == 0 || !doctorErrors(checks.DoctorReport{Diagnostics: want}) {
		t.Fatal("fixture has no static diagnostic")
	}
	report := checks.Doctor(context.Background(), f.request(false))
	remaining := map[project.Diagnostic]int{}
	for _, d := range report.Diagnostics {
		remaining[d]++
	}
	for _, d := range want {
		if remaining[d] == 0 {
			t.Errorf("complete project diagnostic %+v lost: %s", d, doctorText(report))
		}
		remaining[d]--
	}
	if strings.Contains(doctorLog(f), "setup:") || strings.Contains(doctorLog(f), "eval:") || strings.Contains(doctorLog(f), "boundaries:") || strings.Contains(doctorLog(f), "worker:") || strings.Contains(doctorLog(f), "audit:") {
		t.Errorf("static error ran probes: %s", doctorLog(f))
	}
}

func Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_CHK_060_CHK_061_CHK_067_RealGoRecipe(t *testing.T) {
	f := newDoctorFixture(t)
	f.commit()
	f.bin = doctorBinary(t)
	f.editConfig(func(config map[string]any) {
		config["modules"] = []string{"packages/alpha"}
		config["hooks"].(map[string]any)["setup"] = fmt.Sprintf("printf 'setup:%%s\\n' \"$PWD\" >> %q; test -z \"$(git symbolic-ref -q HEAD)\" && test -f .fireweed/doctor-untracked && go mod download", f.log)
	})
	f.put(".fireweed/doctor-untracked", "ready\n")
	for _, d := range project.Check(f.root) {
		if !d.Warning {
			t.Fatalf("invalid Go doctor fixture: %+v", d)
		}
	}
	beforeStatus, beforeWorktrees := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) {
		t.Errorf("real Go recipe doctor failed: %s", doctorText(report))
	}
	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id", "recovery"} {
		if !hasDoctorCheck(report, name, true) {
			t.Errorf("missing passing %s outcome: %s", name, doctorText(report))
		}
	}
	recoveries := 0
	for _, check := range report.Checks {
		if strings.Contains(strings.ToLower(check.Name), "recovery") && check.Passed {
			recoveries++
		}
	}
	if recoveries < 4 {
		t.Errorf("clean recovery not reported after each canary: %s", doctorText(report))
	}
	modules := map[string]bool{}
	for _, check := range report.Checks {
		if check.Module != "" && check.Module != "alpha" {
			modules[check.Module] = true
		}
	}
	if len(modules) < 2 {
		t.Errorf("two distinct throwaway modules absent: %s", doctorText(report))
	}
	if !strings.Contains(doctorLog(f), "setup:") || !strings.Contains(doctorLog(f), "/fireweed/worktrees/") {
		t.Errorf("setup did not run in scratch: %s", doctorLog(f))
	}
	if strings.Contains(doctorLog(f), "worker:") || strings.Contains(doctorLog(f), "audit:") {
		t.Errorf("NoAgent ran agents: %s", doctorLog(f))
	}
	if got := f.git("status", "--porcelain", "--untracked-files=all"); got != beforeStatus {
		t.Errorf("caller changed: before %q after %q", beforeStatus, got)
	}
	if got := worktrees(f); got != beforeWorktrees {
		t.Errorf("scratch registration leaked: before %q after %q", beforeWorktrees, got)
	}
	assertScratchRemoved(t, f)
}

func Test_CHK_053_NestedModuleTemplatesRenderForBothCanaries(t *testing.T) {
	f := localDoctorFixture(t)
	f.put("go.mod", "module example.test/doctorfixture\n\ngo 1.23\n")
	fields := "{{name}}|{{prefix}}|{{package}}|{{go_module}}|{{module_dir}}|{{other_dir}}\n"
	f.put(".fireweed/templates/module/evals/nested.txt.tmpl", "eval|"+fields)
	f.put(".fireweed/templates/module/internal/nested.txt", "internal|"+fields)
	canaries := string(mustRead(t, filepath.Join(f.root, ".fireweed/canaries/canaries.yaml")))
	f.put(".fireweed/canaries/canaries.yaml", strings.Replace(canaries, "undeclared:\n", "  evals/baseline.txt: \"baseline|"+strings.TrimSpace(fields)+"\"\nundeclared:\n", 1))
	eval := string(mustRead(t, filepath.Join(f.root, ".fireweed/eval.sh")))
	check := fmt.Sprintf(`case "$1" in
  doctor*)
    prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
    other=
    for spec in packages/*/SPEC.md; do
      dir=${spec%%/SPEC.md}
      case "$dir" in packages/alpha|"$2") ;; *) other=$dir ;; esac
    done
    expected="$1|$prefix|$1|example.test/doctorfixture|$2|$other"
    test -n "$other" &&
      test "$(cat "$2/evals/nested.txt")" = "eval|$expected" &&
      test "$(cat "$2/internal/nested.txt")" = "internal|$expected" &&
      test "$(cat "$2/evals/baseline.txt")" = "baseline|$expected" &&
      test ! -e "$2/evals/nested.txt.tmpl" || exit 17
    printf 'rendered:%%s\n' "$1" >> %q
    ;;
esac
`, f.log)
	f.put(".fireweed/eval.sh", strings.Replace(eval, "#!/bin/sh\n", "#!/bin/sh\n"+check, 1))
	f.commit()
	before := worktrees(f)
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) {
		t.Errorf("nested template probes failed: %s", doctorText(report))
	}
	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id", "recovery"} {
		if !hasDoctorCheck(report, name, true) {
			t.Errorf("missing passing %s probe: %s", name, doctorText(report))
		}
	}
	modules := map[string]bool{}
	for _, line := range strings.Split(doctorLog(f), "\n") {
		if strings.HasPrefix(line, "rendered:") {
			modules[strings.TrimPrefix(line, "rendered:")] = true
		}
	}
	if len(modules) != 2 {
		t.Errorf("expected both throwaway modules to pass nested template checks, got %v: %s", modules, doctorText(report))
	}
	if got := worktrees(f); got != before {
		t.Errorf("scratch registration leaked: before %q after %q", before, got)
	}
	if strings.Contains(doctorLog(f), "setup:") {
		assertScratchRemoved(t, f)
	}
}

func Test_CHK_051_CHK_052_SetupFailureStopsAndCleans(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/setup.sh", fmt.Sprintf("#!/bin/sh\nprintf 'setup:%%s\\n' \"$PWD\" >> %q\necho failed setup >&2\nexit 17\n", f.log))
	f.commit()
	before := worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if !hasDoctorDiagnostic(report, "C14", false, "setup") {
		t.Errorf("setup failure absent: %s", doctorText(report))
	}
	if strings.Contains(doctorLog(f), "eval:") || strings.Contains(doctorLog(f), "boundaries:") || strings.Contains(doctorLog(f), "worker:") {
		t.Errorf("setup failure ran later probes: %s", doctorLog(f))
	}
	if got := worktrees(f); got != before {
		t.Errorf("scratch worktree leaked: %q", got)
	}
	assertScratchRemoved(t, f)
}

func Test_CHK_051_CancellationCleansScratchAndStopsSetupGroup(t *testing.T) {
	f := localDoctorFixture(t)
	ready := filepath.Join(t.TempDir(), "ready")
	childFinished := filepath.Join(t.TempDir(), "child-finished")
	f.put(".fireweed/setup.sh", fmt.Sprintf(`#!/bin/sh
printf 'setup:%%s\n' "$PWD" >> %q
sh -c 'sleep 2; printf finished > "$1"' sh %q &
printf ready > %q
wait
`, f.log, childFinished, ready))
	f.commit()
	before := worktrees(f)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reports := make(chan checks.DoctorReport, 1)
	go func() { reports <- checks.Doctor(ctx, f.request(true)) }()
	deadline := time.After(8 * time.Second)
	for {
		if _, err := os.Stat(ready); err == nil {
			break
		} else if !os.IsNotExist(err) {
			t.Fatal(err)
		}
		select {
		case report := <-reports:
			t.Fatalf("setup exited before cancellation: %s", doctorText(report))
		case <-deadline:
			t.Fatal("setup never signalled readiness")
		case <-time.After(10 * time.Millisecond):
		}
	}
	cancel()
	select {
	case report := <-reports:
		if !hasDoctorDiagnostic(report, "C14", false, "setup") {
			t.Errorf("cancelled setup lacked C14: %s", doctorText(report))
		}
	case <-time.After(5 * time.Second):
		t.Fatal("cancelled doctor did not return promptly")
	}
	if got := worktrees(f); got != before {
		t.Errorf("cancelled doctor retained Git registration: before %q after %q", before, got)
	}
	assertScratchRemoved(t, f)
	time.Sleep(2200 * time.Millisecond)
	if _, err := os.Stat(childFinished); !os.IsNotExist(err) {
		t.Errorf("cancelled setup descendant continued: %v", err)
	}
}

func Test_CHK_054_MissingTemplatesAndCanaryMaps(t *testing.T) {
	for _, tc := range []struct{ name, path, body, term, input, affected string }{
		{"missing template", ".fireweed/templates/module/SPEC.md", "", "template", "SPEC.md", "undeclared"},
		{"malformed template", ".fireweed/templates/module/SPEC.md", "---\nname: [broken\nid_prefix: {{prefix}}\n---\n", "template", "SPEC.md", "undeclared"},
		{"missing baseline", ".fireweed/canaries/canaries.yaml", "version: 1\nbaseline: {}\n", "canary", "baseline", "undeclared"},
		{"missing internal_eval", ".fireweed/canaries/canaries.yaml", "missing_id:\n  evals/missing.txt: missing\n", "internal_eval", "internal_eval", "internal_eval"},
		{"wrong version", ".fireweed/canaries/canaries.yaml", "version: 2\nbaseline: {}\n", "canary", "canaries.yaml", "undeclared"},
		{"malformed canary", ".fireweed/canaries/canaries.yaml", "version: [bad\n", "canary", "canaries.yaml", "undeclared"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			if tc.name == "missing baseline" {
				body := string(mustRead(t, filepath.Join(f.root, tc.path)))
				start, end := strings.Index(body, "baseline:\n"), strings.Index(body, "undeclared:\n")
				if start < 0 || end <= start {
					t.Fatal("fixture cannot isolate baseline map")
				}
				tc.body = body[:start] + "baseline: {}\n" + body[end:]
			}
			if tc.name == "missing internal_eval" {
				body := string(mustRead(t, filepath.Join(f.root, tc.path)))
				tc.body = strings.Replace(body, "internal_eval:\n  internal/private.txt: \"private\"\n  evals/internal.txt: \"{{module_dir}}\"\n", "", 1)
				if tc.body == body {
					t.Fatal("fixture did not remove internal_eval map")
				}
			}
			if tc.body == "" {
				if err := os.Remove(filepath.Join(f.root, tc.path)); err != nil {
					t.Fatal(err)
				}
			} else {
				f.put(tc.path, tc.body)
			}
			f.commit()
			report := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(report, "C14", false, tc.term) {
				t.Errorf("missing C14: %s", doctorText(report))
			}
			identified := false
			for _, d := range report.Diagnostics {
				if d.Code == "C14" && !d.Warning && strings.Contains(strings.ToLower(d.Message+" "+d.Path), strings.ToLower(tc.input)) {
					identified = true
				}
			}
			if !identified {
				t.Errorf("C14 did not identify %s: %s", tc.input, doctorText(report))
			}
			if hasDoctorCheck(report, tc.affected, true) {
				t.Errorf("invalid recipe passed canary: %s", doctorText(report))
			}
		})
	}
}

func Test_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_ZeroHookAndRecovery(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/boundary.sh", "#!/bin/sh\nexit 0\n")
	f.put(".fireweed/eval.sh", `#!/bin/sh
printf '<testsuite><testcase name="ALP-001"/></testsuite>' > "$3"
`)
	f.commit()
	before := worktrees(f)
	report := checks.Doctor(context.Background(), f.request(true))
	for _, tc := range []struct{ name, hook string }{{"undeclared", "boundar"}, {"internal_eval", "boundar"}, {"failing_eval", "eval"}, {"missing_id", "eval"}} {
		if !hasDoctorDiagnostic(report, "C14", false, tc.name, tc.hook) {
			t.Errorf("missed %s canary not diagnosed with hook: %s", tc.name, doctorText(report))
		}
	}
	if !doctorErrors(report) {
		t.Errorf("zero hooks accepted: %s", doctorText(report))
	}
	if got := worktrees(f); got != before {
		t.Errorf("failed canary leaked worktree: %s", got)
	}

	g := localDoctorFixture(t)
	g.put(".fireweed/boundary.sh", fmt.Sprintf(`#!/bin/sh
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test -f "$dir/undeclared.txt"; then printf '%%s' "$name" > .fireweed/recovery-seen; printf '%%s: undeclared import\n' "$name"; fi
done
if test -f .fireweed/recovery-seen && ! find packages -name undeclared.txt | grep -q .; then printf '%%s: lingering violation\n' "$(cat .fireweed/recovery-seen)"; fi
`))
	g.commit()
	r := checks.Doctor(context.Background(), g.request(true))
	if !hasDoctorDiagnostic(r, "C14", false, "recovery", "undeclared", "boundar") {
		t.Errorf("missing failed recovery diagnostic: %s", doctorText(r))
	}
}

func Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific(t *testing.T) {
	for _, tc := range []struct{ name, script, canary string }{
		{"missing baseline coverage", `#!/bin/sh
printf '<testsuite/>' > "$3"
`, "baseline"},
		{"failing eval compile error without report", `#!/bin/sh
if test -f "$2/evals/failing.txt"; then echo 'compile error' >&2; exit 2; fi
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "failing_eval"},
		{"failing eval with unrelated failure", `#!/bin/sh
if test -f "$2/evals/failing.txt"; then printf '<testsuite><testcase name="ALP-001"><failure/></testcase></testsuite>' > "$3"; exit 1; fi
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "failing_eval"},
		{"missing id with unrelated failure", `#!/bin/sh
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
if test -f "$2/evals/missing.txt"; then printf '<testsuite><testcase name="%s-001"><failure/></testcase></testsuite>' "$prefix" > "$3"; exit 1; fi
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "missing_id"},
		{"missing id without report", `#!/bin/sh
if test -f "$2/evals/missing.txt"; then echo 'hook failed' >&2; exit 2; fi
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "missing_id"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			f.put(".fireweed/eval.sh", tc.script)
			if tc.name == "missing baseline coverage" {
				logDoctorCanaryProbes(t, f)
			}
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(r, "C14", false, tc.canary, "eval") {
				t.Errorf("unrelated eval outcome accepted: %s", doctorText(r))
			}
			if tc.name == "missing baseline coverage" && !strings.Contains(doctorLog(f), "probe:") {
				t.Errorf("failed baseline coverage hook did not run: %s", doctorLog(f))
			}
			if tc.name == "missing baseline coverage" && strings.Contains(doctorLog(f), "planted:") {
				t.Errorf("failed baseline coverage continued into canary probes: %s", doctorLog(f))
			}
		})
	}
}

func Test_CHK_055_BaselineAndRecoveryFailures(t *testing.T) {
	for _, tc := range []struct{ name, path, script, phase, hook string }{
		{"baseline eval failure", ".fireweed/eval.sh", "#!/bin/sh\necho baseline-eval-failed >&2\nexit 9\n", "baseline", "eval"},
		{"baseline boundary failure", ".fireweed/boundary.sh", "#!/bin/sh\nprintf 'alpha: baseline violation\\n'\n", "baseline", "boundar"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			f.put(tc.path, tc.script)
			logDoctorCanaryProbes(t, f)
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(r, "C14", false, tc.phase, tc.hook) {
				t.Errorf("failed baseline accepted: %s", doctorText(r))
			}
			if !strings.Contains(doctorLog(f), "probe:") {
				t.Errorf("failed baseline hook did not run: %s", doctorLog(f))
			}
			if strings.Contains(doctorLog(f), "planted:") {
				t.Errorf("failed baseline continued into canary probes: %s", doctorLog(f))
			}
		})
	}
	for _, tc := range []struct{ name, failure string }{
		{"recovery eval failure", "echo recovery-eval-failed >&2; exit 9"},
		{"recovery coverage failure", "printf '<testsuite/>' > \"$3\"; exit 0"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			original := string(mustRead(t, filepath.Join(f.root, ".fireweed/eval.sh")))
			injection := "if test -f \"$2/evals/failing.txt\"; then printf '%s' \"$2\" > .fireweed/eval-recovery-module; fi\n" +
				"if test -f .fireweed/eval-recovery-module && test \"$(cat .fireweed/eval-recovery-module)\" = \"$2\" && ! test -f \"$2/evals/failing.txt\"; then " + tc.failure + "; fi\n"
			f.put(".fireweed/eval.sh", strings.Replace(original, "#!/bin/sh\n", "#!/bin/sh\n"+injection, 1))
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(r, "C14", false, "recovery", "failing_eval", "eval") {
				t.Errorf("failed %s accepted: %s", tc.name, doctorText(r))
			}
		})
	}
}

func Test_CHK_056_CHK_057_BoundaryViolationMustNameCanaryModule(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/boundary.sh", `#!/bin/sh
if find packages \( -name undeclared.txt -o -name internal.txt \) | grep -q .; then printf 'alpha: unrelated violation\n'; fi
`)
	f.commit()
	r := checks.Doctor(context.Background(), f.request(true))
	for _, name := range []string{"undeclared", "internal_eval"} {
		if !hasDoctorDiagnostic(r, "C14", false, name, "boundar") {
			t.Errorf("unrelated module violation accepted for %s: %s", name, doctorText(r))
		}
	}
}

func Test_CHK_057_InternalEvalRequiresInterfaceViolation(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/boundary.sh", `#!/bin/sh
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test -f "$dir/undeclared.txt"; then printf '%s: undeclared import\n' "$name"; fi
  if test -f "$dir/evals/internal.txt"; then printf '%s: unrelated boundary violation\n' "$name"; fi
done
`)
	f.commit()
	r := checks.Doctor(context.Background(), f.request(true))
	passedProbe := false
	for _, check := range r.Checks {
		if check.Passed && (check.Name == "internal_eval" || strings.HasPrefix(check.Name, "internal_eval ")) {
			passedProbe = true
		}
	}
	if !hasDoctorDiagnostic(r, "C14", false, "internal_eval", "boundar") || passedProbe {
		t.Errorf("same-module non-interface violation accepted: %s", doctorText(r))
	}
	if !hasDoctorCheck(r, "undeclared", true) {
		t.Errorf("fixture did not detect the independent undeclared canary: %s", doctorText(r))
	}
}

func Test_CHK_055_CHK_056_CHK_057_MissingRequiredBoundaryHook(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	f.put("packages/alpha/internal/live.txt", "disposable implementation\n")
	logDoctorCanaryProbes(t, f)
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(report, "X02", false, "alpha", "boundar") {
		t.Errorf("required boundary hook was not diagnosed: %s", doctorText(report))
	}
	if !hasDoctorDiagnostic(report, "C14", false, "baseline", "boundar") {
		t.Errorf("missing boundary hook was accepted as a clean baseline: %s", doctorText(report))
	}
	if strings.Contains(doctorLog(f), "planted:") {
		t.Errorf("missing baseline boundary hook continued into canary probes: %s", doctorLog(f))
	}
	for _, canary := range []string{"undeclared", "internal_eval"} {
		if hasDoctorCheck(report, canary, true) {
			t.Errorf("missing boundary hook counted as %s detection: %s", canary, doctorText(report))
		}
		if !hasDoctorDiagnostic(report, "C14", false, canary, "boundar") {
			t.Errorf("missing boundary hook lacked named %s canary failure: %s", canary, doctorText(report))
		}
	}
}

func Test_CHK_055_MissingRequiredHookBesideOverride(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	f.put(".fireweed/override-boundary.sh", fmt.Sprintf("#!/bin/sh\nprintf 'override-boundaries:%%s\\n' \"$PWD\" >> %q\nexec sh .fireweed/boundary.sh\n", f.log))
	f.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(f.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nhooks:\n  boundaries: sh .fireweed/override-boundary.sh", 1))
	f.put("packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: []\n---\n# Beta\n\n## Specs\n\n### BET-001\nThe fixture MUST pass.\n")
	f.put("packages/beta/internal/live.txt", "disposable implementation\n")
	logDoctorCanaryProbes(t, f)
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(report, "X02", false, "beta", "boundar") {
		t.Errorf("module without its own hook escaped required boundary enforcement: %s", doctorText(report))
	}
	if strings.Contains(doctorLog(f), "planted:") {
		t.Errorf("failed baseline beside override continued into canary probes: %s", doctorLog(f))
	}
	if !strings.Contains(doctorLog(f), "override-boundaries:") {
		t.Errorf("applicable override was not run: %s", doctorLog(f))
	}
}

func Test_CHK_055_ApplicableOverrideNeedsNoGlobalHook(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	f.put(".fireweed/override-boundary.sh", "#!/bin/sh\nexec sh .fireweed/boundary.sh\n")
	f.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(f.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nhooks:\n  boundaries: sh .fireweed/override-boundary.sh", 1))
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) || !hasDoctorCheck(report, "undeclared", true) || !hasDoctorCheck(report, "internal_eval", true) {
		t.Errorf("applicable override was rejected without a global hook: %s", doctorText(report))
	}
}

func Test_CHK_060_CHK_062_ModuleOverridesAndBoundaryWarnings(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/override-eval.sh", fmt.Sprintf("#!/bin/sh\nprintf 'override-eval:%%s\\n' \"$PWD\" >> %q\nexec sh .fireweed/eval.sh \"$@\"\n", f.log))
	f.put(".fireweed/override-boundary.sh", fmt.Sprintf("#!/bin/sh\nprintf 'override-boundaries:%%s:%%s\\n' \"$PWD\" \"$1\" >> %q\nexec sh .fireweed/boundary.sh\n", f.log))
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\nhooks:\n  eval: sh .fireweed/override-eval.sh {module} {module_dir} {report}\n  boundaries: sh .fireweed/override-boundary.sh {module}\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nThe fixture MUST pass.\n")
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) || strings.Count(doctorLog(f), "override-eval:") < 2 || strings.Count(doctorLog(f), "override-boundaries:") < 2 {
		t.Errorf("override probes absent: %s", doctorLog(f))
	}
	g := localDoctorFixture(t)
	g.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	g.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(g.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nboundaries_opt_out: legacy", 1))
	g.commit()
	r := checks.Doctor(context.Background(), g.request(true))
	if doctorErrors(r) || !hasDoctorDiagnostic(r, "X03", true, "legacy") || !hasDoctorDiagnostic(r, "X04", true, "no boundary enforcement required") {
		t.Errorf("optional boundaries warning lost: %s", doctorText(r))
	}
}

func Test_CHK_060_OverridesMustDetectTheirCanaries(t *testing.T) {
	for _, tc := range []struct {
		name, hook, command, script, canary string
	}{
		{"eval", "eval", "sh .fireweed/override-eval.sh {module} {module_dir} {report}", `#!/bin/sh
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "failing_eval"},
		{"boundary", "boundaries", "sh .fireweed/override-boundary.sh {module}", "#!/bin/sh\nexit 0\n", "internal_eval"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			path := ".fireweed/override-" + tc.name + ".sh"
			f.put(path, tc.script)
			f.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(f.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nhooks:\n  "+tc.hook+": "+tc.command, 1))
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			diagnosed := false
			for _, d := range r.Diagnostics {
				message := strings.ToLower(d.Message)
				boundaryHook := tc.name == "boundary" && !strings.Contains(message, "global") &&
					(strings.Contains(message, "alpha boundary hook") ||
						strings.Contains(message, "alpha boundaries hook") ||
						strings.Contains(message, "alpha boundaries override boundary hook") ||
						strings.Contains(message, "alpha boundaries override boundaries hook"))
				if d.Code == "C14" && !d.Warning && strings.Contains(message, tc.canary) &&
					(tc.name == "eval" && strings.Contains(message, "alpha "+tc.hook+" override") || boundaryHook) {
					diagnosed = true
				}
			}
			if !diagnosed {
				t.Errorf("broken %s override missed its %s canary without C14: %s", tc.name, tc.canary, doctorText(r))
			}
		})
	}
}

func Test_CHK_063_CHK_064_CHK_065_CHK_066_CHK_067_AgentProbes(t *testing.T) {
	f := localDoctorFixture(t)
	f.commit()
	for _, d := range project.Check(f.root) {
		if !d.Warning {
			t.Fatalf("invalid agent fixture: %+v", d)
		}
	}
	r := checks.Doctor(context.Background(), f.request(false))
	log := doctorLog(f)
	if doctorErrors(r) || strings.Count(log, "worker:") != 1 || strings.Count(log, "audit:") != 1 || strings.Count(log, "wrap:") != 2 || !strings.Contains(log, "FIREWEED_DOCTOR_OK") {
		t.Errorf("worker/audit/wrap probes: %s; log=%s", doctorText(r), log)
	}
	workerPrompt, auditPrompt := false, false
	for _, line := range strings.Split(log, "\n") {
		if fields := strings.SplitN(line, ":", 3); len(fields) == 3 && strings.Contains(fields[2], "FIREWEED_DOCTOR_OK") {
			switch fields[0] {
			case "worker":
				workerPrompt = true
			case "audit":
				auditPrompt = true
			}
		}
	}
	if !workerPrompt {
		t.Errorf("worker did not receive its fixed-reply prompt: %s", log)
	}
	if !auditPrompt {
		t.Errorf("audit did not receive its fixed-reply prompt: %s", log)
	}
	for _, part := range []string{"/fireweed/worktrees/", "worker:", "audit:"} {
		if !strings.Contains(log, part) {
			t.Errorf("scratch/role marker %s absent: %s", part, log)
		}
	}

	g := localDoctorFixture(t)
	g.editConfig(func(config map[string]any) {
		delete(config["audit"].(map[string]any), "command")
		config["agent"].(map[string]any)["command"] = "bash .fireweed/agent.sh {prompt} {git_common_dir}"
	})
	g.commit()
	rr := checks.Doctor(context.Background(), g.request(false))
	if !hasDoctorDiagnostic(rr, "C15", true, "shared") || strings.Count(doctorLog(g), "worker:") != 2 {
		t.Errorf("audit fallback: %s; log=%s", doctorText(rr), doctorLog(g))
	}

	h := localDoctorFixture(t)
	h.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/bash\nfor i in $(seq 1 25); do printf 'line-%%02d\\n' \"$i\"; done\necho RATE LIMIT reached\nexit 9\n"))
	h.commit()
	bad := checks.Doctor(context.Background(), h.request(false))
	if !hasDoctorDiagnostic(bad, "C15", false, "usage limit", "rate limit", "line-25") {
		t.Errorf("usage-limit tail missing: %s", doctorText(bad))
	}
	if strings.Contains(doctorText(bad), "line-01") {
		t.Errorf("agent output was not trimmed to final 20 lines: %s", doctorText(bad))
	}
	outOfTail := localDoctorFixture(t)
	outOfTail.put(".fireweed/agent.sh", "#!/bin/bash\necho RATE LIMIT early\nfor i in $(seq 1 25); do printf 'later-%02d\\n' \"$i\"; done\nexit 9\n")
	outOfTail.commit()
	ordinary := checks.Doctor(context.Background(), outOfTail.request(false))
	if !hasDoctorDiagnostic(ordinary, "C15", false, "agent.sh", "later-25") {
		t.Errorf("ordinary out-of-tail failure absent: %s", doctorText(ordinary))
	}
	for _, d := range ordinary.Diagnostics {
		if d.Code == "C15" && !d.Warning && strings.Contains(strings.ToLower(d.Message), "usage limit") {
			t.Errorf("out-of-tail pattern classified as usage limit: %s", doctorText(ordinary))
		}
	}
	for _, tc := range []struct{ name, script string }{{"wrong reply", "#!/bin/bash\necho WRONG_REPLY\n"}, {"nonzero with reply", "#!/bin/bash\necho FIREWEED_DOCTOR_OK\nexit 7\n"}, {"ordinary failure", "#!/bin/bash\nfor i in $(seq 1 25); do printf 'ordinary-%02d\\n' \"$i\"; done\nexit 8\n"}} {
		t.Run(tc.name, func(t *testing.T) {
			q := localDoctorFixture(t)
			q.put(".fireweed/agent.sh", tc.script)
			q.commit()
			r := checks.Doctor(context.Background(), q.request(false))
			if !hasDoctorDiagnostic(r, "C15", false, "agent.sh") {
				t.Errorf("agent failure accepted: %s", doctorText(r))
			}
			if tc.name == "ordinary failure" && (!hasDoctorDiagnostic(r, "C15", false, "ordinary-25") || strings.Contains(doctorText(r), "ordinary-01")) {
				t.Errorf("ordinary failure tail incorrect: %s", doctorText(r))
			}
		})
	}
	wrongAudit := localDoctorFixture(t)
	wrongAudit.put(".fireweed/audit.sh", "#!/bin/bash\necho WRONG_REPLY\n")
	wrongAudit.commit()
	auditFailure := checks.Doctor(context.Background(), wrongAudit.request(false))
	if !hasDoctorDiagnostic(auditFailure, "C15", false, "audit.sh") || hasDoctorCheck(auditFailure, "audit", true) {
		t.Errorf("wrong audit reply accepted: %s", doctorText(auditFailure))
	}

	i := localDoctorFixture(t)
	i.put(".fireweed/audit.sh", "#!/bin/bash\nsleep 2\n")
	i.commit()
	req := i.request(false)
	req.Timeout = 100 * time.Millisecond
	start := time.Now()
	timed := checks.Doctor(context.Background(), req)
	if !hasDoctorDiagnostic(timed, "C15", false, "audit") || time.Since(start) > 4*time.Second {
		t.Errorf("audit timeout not bounded: %s; elapsed=%s", doctorText(timed), time.Since(start))
	}

	defaulted := localDoctorFixture(t)
	defaulted.put(".fireweed/agent.sh", "#!/bin/bash\nsleep 0.2\nprintf 'FIREWEED_DOCTOR_OK\\n'\n")
	defaulted.put(".fireweed/audit.sh", "#!/bin/bash\nsleep 0.2\nprintf 'FIREWEED_DOCTOR_OK\\n'\n")
	defaulted.commit()
	defaultRequest := defaulted.request(false)
	defaultRequest.Timeout = 0
	defaultReport := checks.Doctor(context.Background(), defaultRequest)
	if doctorErrors(defaultReport) || !hasDoctorCheck(defaultReport, "worker", true) || !hasDoctorCheck(defaultReport, "audit", true) {
		t.Errorf("unset timeout did not give both agent probes their default allowance: %s", doctorText(defaultReport))
	}

	j := localDoctorFixture(t)
	j.commit()
	skip := checks.Doctor(context.Background(), j.request(true))
	if strings.Contains(doctorLog(j), "worker:") || strings.Contains(doctorLog(j), "audit:") || !hasDoctorCheck(skip, "undeclared", true) || !strings.Contains(doctorLog(j), j.bin) {
		t.Errorf("NoAgent skipped hooks or ran agent: %s; log=%s", doctorText(skip), doctorLog(j))
	}
}

func Test_CHK_065_UnsetTimeoutBoundsAuditProbe(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/audit.sh", fmt.Sprintf("#!/bin/bash\nprintf 'audit:%%s:%%s\n' \"$PWD\" \"$1\" >> %q\nprintf 'audit-timeout-started\n'\nsleep 210\nprintf 'FIREWEED_DOCTOR_OK\n'\n", f.log))
	f.commit()
	req := f.request(false)
	req.Timeout = 0

	start := time.Now()
	report := checks.Doctor(context.Background(), req)
	elapsed := time.Since(start)
	log := doctorLog(f)
	if !hasDoctorCheck(report, "worker", true) || strings.Count(log, "worker:") != 1 || strings.Count(log, "audit:") != 1 {
		t.Errorf("agent probes did not run once with immediate worker reply: %s; log=%s", doctorText(report), log)
	}
	if !hasDoctorCheck(report, "audit", false) || !hasDoctorDiagnostic(report, "C15", false, "audit.sh", "audit-timeout-started") {
		t.Errorf("long audit probe lacked a C15 timeout with its command and output: %s", doctorText(report))
	}
	if elapsed < 165*time.Second || elapsed > 200*time.Second {
		t.Errorf("unset audit timeout took %s, want a bound near three minutes", elapsed)
	}
}

func Test_CHK_068_CHK_069_ManagedGuidanceAndPlaybookBase(t *testing.T) {
	f := localDoctorFixture(t)
	text, err := assets.Playbook("evals")
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256([]byte(text))
	base := hex.EncodeToString(sum[:])
	f.put(".fireweed/playbooks/evals.md", "---\nembedded_base: "+base+"\n---\n# Local override\n")
	f.commit()
	r := checks.Doctor(context.Background(), f.request(true))
	if hasDoctorDiagnostic(r, "C16", true) {
		t.Errorf("current guidance/base warned: %s", doctorText(r))
	}
	f.put("AGENTS.md", "Before.\n"+doctorGuidance+"After.\n")
	surrounded := checks.Doctor(context.Background(), f.request(true))
	if hasDoctorDiagnostic(surrounded, "C16", true) {
		t.Errorf("current managed guidance with surrounding text warned: %s", doctorText(surrounded))
	}
	f.put("AGENTS.md", "Before.\n<!-- fireweed:start -->\nold guidance\n<!-- fireweed:end -->\nAfter.\n")
	f.put(".fireweed/playbooks/evals.md", "---\nembedded_base: stale\n---\n# Local override\n")
	stale := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(stale, "C16", true, "AGENTS.md") || !hasDoctorDiagnostic(stale, "C16", true, "evals") {
		t.Errorf("stale guidance/base not warned: %s", doctorText(stale))
	}
	if err := os.Remove(filepath.Join(f.root, "AGENTS.md")); err != nil {
		t.Fatal(err)
	}
	f.put(".fireweed/playbooks/evals.md", "# No embedded base\n")
	missing := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(missing, "C16", true, "AGENTS.md") || !hasDoctorDiagnostic(missing, "C16", true, "evals") {
		t.Errorf("missing guidance/base not warned: %s", doctorText(missing))
	}
}

func mustRead(t *testing.T, path string) []byte {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return b
}
func doctorBinary(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("caller")
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(file), "..", "..", ".."))
	bin := filepath.Join(t.TempDir(), "fireweed-new")
	cmd := exec.Command("go", "build", "-o", bin, ".")
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GOCACHE=/tmp/fireweed-rebuild-cache")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build product: %v: %s", err, out)
	}
	return bin
}
package evals_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
)

func doctorProbeLog(t *testing.T, f doctorFixture, report checks.DoctorReport, role string, markers ...string) string {
	t.Helper()
	var detail string
	for _, check := range report.Checks {
		if check.Name == role {
			detail = check.Detail
			break
		}
	}
	if detail == "" {
		t.Errorf("executed %s probe has no check: %s", role, doctorText(report))
		return ""
	}
	at := strings.Index(detail, "log: ")
	if at < 0 {
		t.Errorf("%s detail lacks log path: %q", role, detail)
		return ""
	}
	fields := strings.Fields(detail[at+len("log: "):])
	if len(fields) == 0 {
		t.Errorf("%s detail has empty log path: %q", role, detail)
		return ""
	}
	path := fields[0]
	if !filepath.IsAbs(path) {
		t.Errorf("%s log path is not absolute: %q", role, path)
	}
	rel, err := filepath.Rel(filepath.Join(f.root, ".fireweed", "runs"), path)
	if err != nil {
		t.Errorf("%s log path %q: %v", role, path, err)
		return ""
	}
	parts := strings.Split(filepath.ToSlash(rel), "/")
	if len(parts) != 2 || parts[0] == "" || parts[0] == "." || parts[0] == ".." || parts[1] != "doctor-"+role+"-1.log" {
		t.Errorf("%s log path = %q, want .fireweed/runs/<run>/doctor-%s-1.log", role, path, role)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("read %s log %q: %v", role, path, err)
		return path
	}
	for _, marker := range markers {
		if !strings.Contains(string(data), marker) {
			t.Errorf("%s log lacks %q: %q", role, marker, data)
		}
	}
	return path
}

func doctorLogRun(t *testing.T, f doctorFixture, worker, audit string, report checks.DoctorReport) string {
	t.Helper()
	base := filepath.Join(f.root, ".fireweed", "runs")
	w, ew := filepath.Rel(base, worker)
	a, ea := filepath.Rel(base, audit)
	if ew != nil || ea != nil {
		t.Errorf("log paths outside run root: %q %q: %v %v", worker, audit, ew, ea)
		return ""
	}
	wRun, aRun := filepath.Dir(w), filepath.Dir(a)
	if wRun != aRun || wRun == "." || wRun == ".." {
		t.Errorf("worker and audit have different run identities: %q %q: %s", worker, audit, doctorText(report))
	}
	return wRun
}

func doctorCallerClean(t *testing.T, f doctorFixture, status, registrations string) {
	t.Helper()
	if got := f.git("status", "--porcelain", "--untracked-files=all"); got != status {
		t.Errorf("caller status changed: before %q after %q", status, got)
	}
	if got := worktrees(f); got != registrations {
		t.Errorf("scratch registration changed: before %q after %q", registrations, got)
	}
	assertScratchRemoved(t, f)
}

func Test_CHK_081_DoctorProbeLogsShareFreshRunAndPreserveStreams(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/agent.sh", "#!/bin/bash\nprintf 'worker stdout Ω without newline'; printf 'worker stderr β without newline' >&2; printf 'FIREWEED_DOCTOR_OK\\n'\n")
	f.put(".fireweed/audit.sh", "#!/bin/bash\nprintf 'audit stdout Ω without newline'; printf 'audit stderr β without newline' >&2; printf 'FIREWEED_DOCTOR_OK\\n'\n")
	f.commit()
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	var firstRun string
	for i := 0; i < 2; i++ {
		report := checks.Doctor(context.Background(), f.request(false))
		if doctorErrors(report) || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
			t.Errorf("successful probes changed outcome: %s", doctorText(report))
		}
		worker := doctorProbeLog(t, f, report, "worker", "worker stdout Ω without newlineFIREWEED_DOCTOR_OK\n", "worker stderr β without newline")
		audit := doctorProbeLog(t, f, report, "audit", "audit stdout Ω without newlineFIREWEED_DOCTOR_OK\n", "audit stderr β without newline")
		if worker != "" && audit != "" {
			run := doctorLogRun(t, f, worker, audit, report)
			if i == 1 && run == firstRun {
				t.Errorf("consecutive Doctor invocations reused run identity %q", run)
			}
			firstRun = run
		}
		doctorCallerClean(t, f, status, registrations)
	}
}

func Test_CHK_081_DoctorPersistsFailedAndTimedOutProbeOutput(t *testing.T) {
	failureStdout := "early worker stdout Ω\n"
	for i := 1; i <= 25; i++ {
		failureStdout += fmt.Sprintf("later-%02d\n", i)
	}
	for _, tc := range []struct {
		name, role, script, stdout, stderr string
		timeout                            time.Duration
	}{
		{"failure", "worker", "#!/bin/bash\nprintf 'early worker stdout Ω\\n'; for i in $(seq 1 25); do printf 'later-%02d\\n' \"$i\"; done; printf 'failed worker stderr β' >&2; exit 7\n", failureStdout, "failed worker stderr β", 2 * time.Second},
		{"timeout", "audit", "#!/bin/bash\nprintf 'partial audit stdout Ω'; printf 'partial audit stderr β' >&2; sleep 5\n", "partial audit stdout Ω", "partial audit stderr β", 150 * time.Millisecond},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			if tc.role == "worker" {
				f.put(".fireweed/agent.sh", tc.script)
			} else {
				f.put(".fireweed/audit.sh", tc.script)
			}
			f.commit()
			status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
			req := f.request(false)
			req.Timeout = tc.timeout
			start := time.Now()
			report := checks.Doctor(context.Background(), req)
			if time.Since(start) > 4*time.Second {
				t.Errorf("%s probe did not honor bounded request timeout", tc.name)
			}
			if !hasDoctorCheck(report, tc.role, false) || !hasDoctorDiagnostic(report, "C15", false, tc.role, tc.stderr) {
				t.Errorf("%s probe outcome/diagnostic lost: %s", tc.name, doctorText(report))
			}
			path := doctorProbeLog(t, f, report, tc.role, tc.stdout, tc.stderr)
			if tc.name == "failure" && path != "" {
				if !strings.Contains(string(mustRead(t, path)), "later-25") || strings.Contains(doctorText(report), "early worker stdout Ω") {
					t.Errorf("failed probe log or final-20-line diagnostic changed: %s", doctorText(report))
				}
			}
			doctorCallerClean(t, f, status, registrations)
		})
	}
}

func Test_CHK_081_DoctorPersistsCancelledProbeOutput(t *testing.T) {
	f := localDoctorFixture(t)
	ready := filepath.Join(t.TempDir(), "worker-ready")
	f.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/bash\nprintf 'cancel worker stdout Ω'; printf 'cancel worker stderr β' >&2; printf ready > %q; sleep 5\n", ready))
	f.commit()
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reports := make(chan checks.DoctorReport, 1)
	go func() { reports <- checks.Doctor(ctx, f.request(false)) }()
	deadline := time.After(8 * time.Second)
	for {
		if _, err := os.Stat(ready); err == nil {
			break
		} else if !os.IsNotExist(err) {
			t.Fatal(err)
		}
		select {
		case report := <-reports:
			t.Fatalf("worker exited before cancellation: %s", doctorText(report))
		case <-deadline:
			t.Fatal("worker never signalled readiness")
		case <-time.After(10 * time.Millisecond):
		}
	}
	cancel()
	select {
	case report := <-reports:
		if !hasDoctorCheck(report, "worker", false) || !hasDoctorDiagnostic(report, "C15", false, "worker") {
			t.Errorf("cancelled probe outcome lost: %s", doctorText(report))
		}
		doctorProbeLog(t, f, report, "worker", "cancel worker stdout Ω", "cancel worker stderr β")
	case <-time.After(4 * time.Second):
		t.Fatal("cancelled Doctor did not return promptly")
	}
	doctorCallerClean(t, f, status, registrations)
}
package evals_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
)

const doctorProbePrompt = "Reply with exactly FIREWEED_DOCTOR_OK"

func doctorRecordedPath(t *testing.T, log, label string) string {
	t.Helper()
	return doctorRecordedPaths(t, log, label, 1)[0]
}

func doctorRecordedPaths(t *testing.T, log, label string, count int) []string {
	t.Helper()
	var paths []string
	for _, line := range strings.Split(log, "\n") {
		if strings.HasPrefix(line, label+":") {
			paths = append(paths, strings.TrimPrefix(line, label+":"))
		}
	}
	if len(paths) != count {
		t.Fatalf("%s paths = %q; want %d absolute paths", label, paths, count)
	}
	for _, path := range paths {
		if !filepath.IsAbs(path) {
			t.Fatalf("%s path = %q; want an absolute path", label, path)
		}
	}
	return paths
}

func doctorPromptFileRemoved(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Errorf("prompt file %q remains after probe: %v", path, err)
	}
}

func Test_CHK_082_CHK_083_DoctorExpandsSelectedCommandsAndWrapperOnce(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) {
		config["agent"].(map[string]any)["command"] = []any{"bash", ".fireweed/agent.sh", "pre{prompt}post", "file={prompt_file}", "common={git_common_dir}", "literal {unknown}"}
		config["audit"].(map[string]any)["command"] = "bash .fireweed/audit.sh {prompt} {prompt_file} {git_common_dir}"
		config["agent"].(map[string]any)["wrap"] = []any{"bash", ".fireweed/prompt-wrap.sh", "{prompt_file}", "{prompt}", "{git_common_dir}"}
	})
	f.put(".fireweed/prompt-wrap.sh", fmt.Sprintf(`#!/bin/bash
printf 'wrap-file:%%s\nwrap-prompt:%%s\nwrap-common:%%s\n' "$1" "$2" "$3" >> %q
test -f "$1" && printf '%%s' "$2" | cmp -s - "$1" || exit 11
shift 3
exec "$@"
`, f.log))
	f.put(".fireweed/agent.sh", fmt.Sprintf(`#!/bin/bash
printf 'worker-prompt:%%s\nworker-file:%%s\nworker-common:%%s\nworker-literal:%%s\n' "$1" "${2#file=}" "${3#common=}" "$4" >> %q
test "$1" = "pre%spost" || exit 12
test "$2" != "file=" || exit 13
test -f "${2#file=}" && printf '%%s' %q | cmp -s - "${2#file=}" || exit 14
printf 'FIREWEED_DOCTOR_OK\n'
`, f.log, doctorProbePrompt, doctorProbePrompt))
	f.put(".fireweed/audit.sh", fmt.Sprintf(`#!/bin/bash
printf 'audit-prompt:%%s\naudit-file:%%s\naudit-common:%%s\n' "$1" "$2" "$3" >> %q
test -f "$2" && printf '%%s' "$1" | cmp -s - "$2" || exit 15
printf 'FIREWEED_DOCTOR_OK\n'
`, f.log))
	f.commit()
	// Relocate the Git directory so its absolute path contains a quote, a space,
	// and another placeholder. The expanded value must remain literal.
	commonDir := filepath.Join(t.TempDir(), "space 'quote' {prompt}", "repo.git")
	if err := os.MkdirAll(filepath.Dir(commonDir), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Rename(filepath.Join(f.root, ".git"), commonDir); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(f.root, ".git"), []byte("gitdir: "+commonDir+"\n"), 0644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("TMPDIR", filepath.Dir(commonDir))
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if doctorErrors(report) || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
		t.Errorf("prompt-file probes failed: %s", doctorText(report))
	}
	log := doctorLog(f)
	common, err := filepath.EvalSymlinks(commonDir)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"worker-prompt:pre" + doctorProbePrompt + "post", "audit-prompt:" + doctorProbePrompt,
		"worker-common:" + common, "audit-common:" + common,
		"worker-literal:literal {unknown}",
		"wrap-prompt:" + doctorProbePrompt, "wrap-common:" + common,
	} {
		if !strings.Contains(log, want+"\n") {
			t.Errorf("missing exact argument %q", want)
		}
	}
	workerFile := doctorRecordedPath(t, log, "worker-file")
	auditFile := doctorRecordedPath(t, log, "audit-file")
	if workerFile == auditFile {
		t.Errorf("worker and audit reused prompt file %q", workerFile)
	}
	for _, path := range []string{workerFile, auditFile} {
		if !strings.Contains(log, "wrap-file:"+path+"\n") {
			t.Errorf("wrapper did not receive selected probe file %q", path)
		}
		doctorPromptFileRemoved(t, path)
	}
	doctorProbeLog(t, f, report, "worker", "FIREWEED_DOCTOR_OK")
	doctorProbeLog(t, f, report, "audit", "FIREWEED_DOCTOR_OK")
	doctorCallerClean(t, f, status, registrations)
}

func Test_CHK_083_DoctorRemovesPromptFileOnProbeFailures(t *testing.T) {
	for _, mode := range []string{"launch", "nonzero", "timeout", "cancel"} {
		t.Run(mode, func(t *testing.T) {
			f := localDoctorFixture(t)
			ready := filepath.Join(t.TempDir(), "ready")
			f.editConfig(func(config map[string]any) {
				config["agent"].(map[string]any)["command"] = []any{"bash", ".fireweed/agent.sh", "{prompt_file}"}
				if mode == "launch" {
					config["agent"].(map[string]any)["command"] = []any{filepath.Join(f.root, "missing-agent"), "{prompt_file}"}
					config["agent"].(map[string]any)["wrap"] = []any{"bash", ".fireweed/prompt-wrap.sh", "{prompt_file}"}
				}
			})
			f.put(".fireweed/prompt-wrap.sh", fmt.Sprintf("#!/bin/bash\nprintf 'worker-file:%%s\\n' \"$1\" >> %q\ntest -f \"$1\" || exit 11\nshift\nexec \"$@\"\n", f.log))
			script := fmt.Sprintf("#!/bin/bash\nprintf 'worker-file:%%s\\nworker-content:%%s\\n' \"$1\" \"$(cat \"$1\")\" >> %q\ntest -f \"$1\" && printf '%%s' %q | cmp -s - \"$1\" || exit 18\nprintf ready > %q\n", f.log, doctorProbePrompt, ready)
			switch mode {
			case "nonzero":
				script += "printf 'failed worker stderr\\n' >&2\nexit 7\n"
			case "timeout", "cancel":
				script += "printf 'waiting worker stdout\\n'\nsleep 5\n"
			}
			f.put(".fireweed/agent.sh", script)
			f.commit()
			status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
			req := f.request(false)
			if mode == "timeout" {
				req.Timeout = 2 * time.Second
			}
			ctx := context.Background()
			var cancel context.CancelFunc
			if mode == "cancel" {
				ctx, cancel = context.WithCancel(ctx)
				defer cancel()
			}
			var report checks.DoctorReport
			if mode == "cancel" {
				reports := make(chan checks.DoctorReport, 1)
				go func() { reports <- checks.Doctor(ctx, req) }()
				deadline := time.After(8 * time.Second)
				for {
					if _, err := os.Stat(ready); err == nil {
						break
					} else if !os.IsNotExist(err) {
						t.Fatal(err)
					}
					select {
					case report = <-reports:
						t.Fatalf("worker exited before cancellation: %s", doctorText(report))
					case <-deadline:
						t.Fatal("worker never signalled readiness")
					case <-time.After(10 * time.Millisecond):
					}
				}
				livePath := doctorRecordedPath(t, doctorLog(f), "worker-file")
				if data, err := os.ReadFile(livePath); err != nil || string(data) != doctorProbePrompt {
					t.Errorf("prompt file unavailable before cancellation: %q, %v", data, err)
				}
				cancel()
				select {
				case report = <-reports:
				case <-time.After(4 * time.Second):
					t.Fatal("cancelled Doctor did not return")
				}
			} else {
				report = checks.Doctor(ctx, req)
			}
			if !hasDoctorCheck(report, "worker", false) || !hasDoctorDiagnostic(report, "C15", false, "worker") {
				t.Errorf("%s probe outcome changed: %s", mode, doctorText(report))
			}
			log := doctorLog(f)
			count := 1
			if mode == "launch" {
				// The shared wrapper records both the failed worker launch and
				// the subsequent audit probe; both files must be cleaned up.
				count = 2
			}
			paths := doctorRecordedPaths(t, log, "worker-file", count)
			if mode != "launch" && !strings.Contains(log, "worker-content:"+doctorProbePrompt+"\n") {
				t.Errorf("%s worker could not read exact prompt: %q", mode, log)
			}
			for _, path := range paths {
				doctorPromptFileRemoved(t, path)
			}
			doctorProbeLog(t, f, report, "worker")
			doctorCallerClean(t, f, status, registrations)
		})
	}
}

func Test_CHK_083_DoctorDoesNotCreatePromptFileWithoutSelectedPlaceholder(t *testing.T) {
	f := localDoctorFixture(t)
	expected := filepath.Join(t.TempDir(), "expected-prompt")
	if err := os.WriteFile(expected, []byte(doctorProbePrompt), 0644); err != nil {
		t.Fatal(err)
	}
	script := fmt.Sprintf(`#!/bin/bash
while IFS= read -r candidate; do
  if test -f "$candidate" && cmp -s "$candidate" %q; then
    printf 'unexpected prompt file: %%s\n' "$candidate" >&2
    exit 19
  fi
done < <(find "$TMPDIR" "$PWD" -type f -print)
printf 'FIREWEED_DOCTOR_OK\n'
`, expected)
	f.put(".fireweed/agent.sh", script)
	f.put(".fireweed/audit.sh", script)
	f.commit()
	temp := t.TempDir()
	t.Setenv("TMPDIR", temp)
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if doctorErrors(report) || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
		t.Errorf("unneeded prompt file appeared: %s", doctorText(report))
	}
	doctorProbeLog(t, f, report, "worker", "FIREWEED_DOCTOR_OK")
	doctorProbeLog(t, f, report, "audit", "FIREWEED_DOCTOR_OK")
	doctorCallerClean(t, f, status, registrations)
}

func Test_CHK_082_CHK_083_DoctorWrapperOnlyFileWithAuditFallback(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) {
		delete(config["audit"].(map[string]any), "command")
		config["agent"].(map[string]any)["command"] = []any{"bash", ".fireweed/agent.sh", "{prompt}"}
		config["agent"].(map[string]any)["wrap"] = []any{"bash", ".fireweed/prompt-wrap.sh", "{prompt_file}"}
	})
	f.put(".fireweed/prompt-wrap.sh", fmt.Sprintf(`#!/bin/bash
printf 'wrap-file:%%s\n' "$1" >> %q
test -f "$1" && printf '%%s' %q | cmp -s - "$1" || exit 11
shift
exec "$@"
`, f.log, doctorProbePrompt))
	f.put(".fireweed/agent.sh", fmt.Sprintf(`#!/bin/bash
printf 'fallback-prompt:%%s\n' "$1" >> %q
test "$1" = %q || exit 12
printf 'FIREWEED_DOCTOR_OK\n'
`, f.log, doctorProbePrompt))
	f.commit()
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if doctorErrors(report) || !hasDoctorDiagnostic(report, "C15", true, "shared") || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
		t.Errorf("wrapper-only fallback changed: %s", doctorText(report))
	}
	log := doctorLog(f)
	if strings.Count(log, "fallback-prompt:"+doctorProbePrompt+"\n") != 2 {
		t.Errorf("fallback did not receive fixed prompt twice")
	}
	var paths []string
	for _, line := range strings.Split(log, "\n") {
		if strings.HasPrefix(line, "wrap-file:") {
			paths = append(paths, strings.TrimPrefix(line, "wrap-file:"))
		}
	}
	if len(paths) != 2 || paths[0] == paths[1] {
		t.Errorf("wrapper prompt files = %q; want distinct worker/audit paths", paths)
	}
	for _, path := range paths {
		if !filepath.IsAbs(path) {
			t.Errorf("wrapper prompt path is relative: %q", path)
		}
		doctorPromptFileRemoved(t, path)
	}
	doctorProbeLog(t, f, report, "worker", "FIREWEED_DOCTOR_OK")
	doctorProbeLog(t, f, report, "audit", "FIREWEED_DOCTOR_OK")
	doctorCallerClean(t, f, status, registrations)
}

 succeeded in 619ms:
package contracts

import (
	"github.com/lurkshark/fireweed/packages/contracts/internal/primitives"
	"github.com/lurkshark/fireweed/packages/contracts/model"
)

type Document = model.Document
type Item = model.Item
type Zones = model.Zones
type CodedError = model.CodedError

// ErrorCode returns the first coded diagnostic in an error chain.
func ErrorCode(err error) (code, subject string, ok bool) { return primitives.ErrorCode(err) }

func ParseDocument(source string) (Document, error) { return primitives.ParseDocument(source) }
func RenderDocument(doc Document) (string, error)   { return primitives.RenderDocument(doc) }

// Section returns the section body without its level-two heading.
func Section(markdown, name string) (string, bool) { return primitives.Section(markdown, name) }

// ReplaceSection replaces the body of a level-two section, or appends it.
func ReplaceSection(markdown, name, body string) string {
	return primitives.ReplaceSection(markdown, name, body)
}

func ParseItems(section string) ([]Item, error) { return primitives.ParseItems(section) }
func RenderItem(item Item) string               { return primitives.RenderItem(item) }
func SpecHash(item Item) string                 { return primitives.SpecHash(item) }

func MatchesID(id, testcaseName, classname string) bool {
	return primitives.MatchesID(id, testcaseName, classname)
}

func MatchGlob(pattern, path string) bool { return primitives.MatchGlob(pattern, path) }

// ResolveZone classifies a module-relative path.
func ResolveZone(path string, zones Zones) string { return primitives.ResolveZone(path, zones) }
package model

// Document is a Markdown document with YAML frontmatter.
type Document struct {
	Frontmatter map[string]any
	Body        string
}

// Item is one level-three spec item. Body excludes the heading line.
type Item struct {
	ID   string
	Tag  string
	Body string
}

// CodedError carries a stable diagnostic code while preserving the cause's text.
type CodedError struct {
	Code    string
	Subject string
	Cause   error
}

func (e CodedError) Error() string {
	if e.Cause != nil {
		return e.Cause.Error()
	}
	return e.Code
}

func (e CodedError) Unwrap() error { return e.Cause }

// Zones holds module-relative path patterns for each editable zone.
type Zones struct {
	Protected  []string
	Evals      []string
	Interface  []string
	AppendOnly []string
	Preserve   []string
	Disposable []string
}
package gitstore

import (
	"context"

	"github.com/lurkshark/fireweed/packages/gitstore/internal/local"
	"github.com/lurkshark/fireweed/packages/gitstore/model"
)

type Repository = model.Repository
type Change = model.Change
type Request = model.Request
type Result = model.Result

func Inspect(dir string) (Repository, error)             { return local.Inspect(dir) }
func BaseRef(root, remote, base string) (string, error)  { return local.BaseRef(root, remote, base) }
func Changes(root, base string) ([]Change, error)        { return local.Changes(root, base) }
func ReadAt(root, revision, path string) ([]byte, error) { return local.ReadAt(root, revision, path) }
func Commit(root, subject string, paths []string) (string, error) {
	return local.Commit(root, subject, paths)
}
func WriteBlob(root string, data []byte) (string, error)       { return local.WriteBlob(root, data) }
func ReadBlob(root, oid string) ([]byte, error)                { return local.ReadBlob(root, oid) }
func Ref(root, name string) (string, bool, error)              { return local.Ref(root, name) }
func CASRef(root, name, old, new string) error                 { return local.CASRef(root, name, old, new) }
func Run(ctx context.Context, request Request) (Result, error) { return local.Run(ctx, request) }
package model

import "time"

type Repository struct {
	Root      string
	CommonDir string
	Head      string
	Branch    string
}

type Change struct {
	Path    string
	OldPath string
	Status  string
}

type Request struct {
	Args    []string
	Dir     string
	Env     map[string]string
	Input   []byte
	Timeout time.Duration
}

type Result struct {
	Stdout   string
	Stderr   string
	ExitCode int
}
package project

import (
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/internal/history"
	"github.com/lurkshark/fireweed/packages/project/internal/queue"
	"github.com/lurkshark/fireweed/packages/project/model"
)

type Snapshot = model.Snapshot
type Module = model.Module
type Change = model.Change
type SpecHistoryEntry = model.SpecHistoryEntry
type StaleSpec = model.StaleSpec
type Diagnostic = model.Diagnostic

func Load(root string) (Snapshot, error) { return fixture.Load(root) }

func LoadAt(root, revision string) (Snapshot, error) { return fixture.LoadAt(root, revision) }

func ValidateGraph(modules []Module) error { return fixture.ValidateGraph(modules) }

func NewChange(root, module, slug string) (string, error) {
	return fixture.NewChange(root, module, slug)
}

func Queue(root string) ([]Change, error) { return queue.Queue(root) }

func Project(root, through string) (Snapshot, error) { return queue.Project(root, through) }

func Stamp(root, change string) error { return queue.Stamp(root, change) }

func StampMissing(root, change string) error { return queue.StampMissing(root, change) }

func StaleSpecs(root, change string) ([]StaleSpec, error) { return queue.StaleSpecs(root, change) }

func Check(root string) []Diagnostic { return queue.Check(root) }

func SpecHistory(root, module, id string) ([]SpecHistoryEntry, error) {
	return history.SpecHistory(root, module, id)
}
package model

import (
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
)

// Snapshot is the project data read from a repository root.
type Snapshot struct {
	Root    string
	Config  map[string]any
	Modules []Module
}

// Module is a discovered module and its declared document and zones.
type Module struct {
	Name      string
	IDPrefix  string
	Dir       string
	DependsOn []string
	Zones     contracts.Zones
	Document  contracts.Document
}

// Change is one pending change document and its spec edits.
type Change struct {
	Number   int
	Slug     string
	Path     string
	Module   string
	Document contracts.Document
	Added    []contracts.Item
	Modified []contracts.Item
	Removed  []contracts.Item
}

// SpecHistoryEntry describes one archived change affecting selected specs.
type SpecHistoryEntry struct {
	Path               string
	Module             string
	IDs                []string
	Motivation         string
	Decisions          string
	RemovedDecisions   string
	Provenance         string
	ImplementationDate time.Time
}

// StaleSpec describes a recorded base that differs from the preceding projection.
type StaleSpec struct {
	ID           string
	RecordedHash string
	Before       string
	Current      string
	Recovered    bool
}

// Diagnostic reports a static authoring check.
type Diagnostic struct {
	Code    string
	Message string
	Path    string
	Warning bool
}
package assets

import "github.com/lurkshark/fireweed/packages/assets/internal/content"
import "github.com/lurkshark/fireweed/packages/assets/internal/recipe"
import "github.com/lurkshark/fireweed/packages/assets/model"

type Recipe = model.Recipe

func Version() string { return content.Version() }

func Help() string { return content.Help() }

func Playbook(name string) (string, error) { return content.Playbook(name) }

func Explain(code string) (string, error) { return content.Explain(code) }

func BuiltinRecipe(name string) (Recipe, error) { return recipe.BuiltinRecipe(name) }

 succeeded in 619ms:
package evals_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

const historyConfig = `version: 1
base: main
modules: ["packages/*"]
protected: [locked/**]
module_defaults:
  evals: [evals/**]
  interface: [api.go]
  append_only: [ledger/**]
`

type historyFixture struct {
	t    *testing.T
	root string
	base string
}

func (f historyFixture) git(args ...string) string {
	f.t.Helper()
	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Args: append([]string{"git", "-C", f.root}, args...),
		Dir:  f.root,
		Env: map[string]string{
			"GIT_CONFIG_GLOBAL": "/dev/null", "GIT_CONFIG_NOSYSTEM": "1",
			"GIT_AUTHOR_NAME": "History Fixture", "GIT_AUTHOR_EMAIL": "fixture@example.invalid",
			"GIT_COMMITTER_NAME": "History Fixture", "GIT_COMMITTER_EMAIL": "fixture@example.invalid",
		},
	})
	if err != nil {
		f.t.Fatalf("git %v: %v: %s", args, err, result.Stderr)
	}
	return strings.TrimSpace(result.Stdout)
}

func (f historyFixture) put(path, body string) {
	f.t.Helper()
	full := filepath.Join(f.root, filepath.FromSlash(path))
	if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
		f.t.Fatal(err)
	}
	if err := os.WriteFile(full, []byte(body), 0644); err != nil {
		f.t.Fatal(err)
	}
}

func (f historyFixture) remove(path string) {
	f.t.Helper()
	if err := os.Remove(filepath.Join(f.root, filepath.FromSlash(path))); err != nil {
		f.t.Fatal(err)
	}
}

func (f historyFixture) move(oldPath, newPath string) {
	f.t.Helper()
	if err := os.MkdirAll(filepath.Dir(filepath.Join(f.root, filepath.FromSlash(newPath))), 0755); err != nil {
		f.t.Fatal(err)
	}
	f.git("mv", oldPath, newPath)
}

func (f historyFixture) commit(subject string) string {
	f.t.Helper()
	f.git("add", "-A")
	f.git("commit", "-qm", subject)
	return f.git("rev-parse", "HEAD")
}

func (f historyFixture) change(path, phase, set string) {
	f.t.Helper()
	body := fmt.Sprintf("---\nmodule: alpha\nphase: %s\n%s---\n\n## Motivation\nHistory fixture.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n\n## Notes\n", phase, set)
	if _, err := contracts.ParseDocument(body); err != nil {
		f.t.Fatalf("fixture change: %v", err)
	}
	f.put(path, body)
}

func newHistoryFixture(t *testing.T) historyFixture {
	t.Helper()
	f := historyFixture{t: t, root: t.TempDir()}
	f.git("init", "-q", "-b", "main")
	f.put("fireweed.yaml", historyConfig)
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha behavior.\n")
	f.put("packages/alpha/api.go", "package alpha\n")
	f.put("packages/alpha/evals/existing_test.go", "package evals\n")
	f.put("packages/alpha/ledger/entry.txt", "initial\n")
	f.put("packages/alpha/newledger/entry.txt", "initial\n")
	f.put("locked/config.txt", "initial\n")
	f.base = f.commit("base")
	if _, err := project.LoadAt(f.root, f.base); err != nil {
		t.Fatalf("recorded base project: %v", err)
	}
	return f
}

func (f historyFixture) report(revisionRange string) checks.HistoryReport {
	f.t.Helper()
	report, err := checks.History(f.root, revisionRange)
	if err != nil {
		f.t.Fatalf("History(%q): %v", revisionRange, err)
	}
	return report
}

type historyFinding struct{ code, path, commit string }

func expectHistoryFindings(t *testing.T, report checks.HistoryReport, want ...historyFinding) {
	t.Helper()
	if len(report.Diagnostics) != len(want) {
		t.Fatalf("diagnostics = %#v; want %d findings", report.Diagnostics, len(want))
	}
	used := make([]bool, len(report.Diagnostics))
	for _, expected := range want {
		found := false
		for i, d := range report.Diagnostics {
			if !used[i] && d.Code == expected.code && d.Path == expected.path && strings.Contains(d.Message, expected.commit) {
				used[i], found = true, true
				break
			}
		}
		if !found {
			t.Errorf("missing %s on %s at %s in %#v", expected.code, expected.path, expected.commit, report.Diagnostics)
		}
	}
}

func Test_CHK_070_HistoryExplicitRangeOrderAndFullOIDs(t *testing.T) {
	f := newHistoryFixture(t)
	f.put("notes/one.txt", "one")
	one := f.commit("first")
	f.put("notes/two.txt", "two")
	two := f.commit("second")
	got := f.report(f.base + ".." + two)
	if !reflect.DeepEqual(got.Commits, []string{one, two}) {
		t.Errorf("commits = %#v; want [%s %s]", got.Commits, one, two)
	}
	expectHistoryFindings(t, got)
}

func Test_CHK_070_HistoryBranchingRangeTopologicalOrder(t *testing.T) {
	f := newHistoryFixture(t)
	f.git("branch", "side")
	f.put("notes/main-one.txt", "main one\n")
	mainOne := f.commit("main first")
	f.put("notes/main-two.txt", "main two\n")
	mainTwo := f.commit("main second")
	f.git("checkout", "-q", "side")
	f.put("notes/side-one.txt", "side one\n")
	sideOne := f.commit("side first")
	f.put("notes/side-two.txt", "side two\n")
	sideTwo := f.commit("side second")
	f.git("checkout", "-q", "main")
	f.git("merge", "--no-ff", "-m", "merge side", "side")
	merge := f.git("rev-parse", "HEAD")

	got := f.report(f.base + ".." + merge)
	want := map[string]bool{mainOne: true, mainTwo: true, sideOne: true, sideTwo: true, merge: true}
	if len(got.Commits) != len(want) {
		t.Errorf("branching commits = %#v; want exactly five commits", got.Commits)
	}
	positions := make(map[string]int)
	for i, commit := range got.Commits {
		if !want[commit] {
			t.Errorf("unexpected commit %s in %#v", commit, got.Commits)
		}
		if _, duplicate := positions[commit]; duplicate {
			t.Errorf("duplicate commit %s in %#v", commit, got.Commits)
		}
		positions[commit] = i
	}
	for commit := range want {
		if _, present := positions[commit]; !present {
			t.Errorf("missing commit %s in %#v", commit, got.Commits)
		}
	}
	for _, edge := range [][2]string{{mainOne, mainTwo}, {sideOne, sideTwo}, {mainTwo, merge}, {sideTwo, merge}} {
		parent, hasParent := positions[edge[0]]
		child, hasChild := positions[edge[1]]
		if hasParent && hasChild && parent >= child {
			t.Errorf("parent %s must precede child %s in %#v", edge[0], edge[1], got.Commits)
		}
	}
	expectHistoryFindings(t, got)
}

func Test_CHK_070_HistoryDefaultRangeUsesLocalConfiguredBase(t *testing.T) {
	f := newHistoryFixture(t)
	f.put("fireweed.yaml", strings.Replace(historyConfig, "base: main", "base: review-base", 1))
	f.commit("configure review base")
	f.git("branch", "review-base")
	f.put("notes/common.txt", "common\n")
	common := f.commit("common after review base")
	f.git("branch", "feature")
	f.git("checkout", "-q", "feature")
	f.put("notes/one.txt", "one")
	one := f.commit("first")
	f.put("notes/two.txt", "two")
	two := f.commit("second")
	f.git("checkout", "-q", "review-base")
	f.put("notes/review-only.txt", "review only\n")
	reviewOnly := f.commit("review base advances independently")
	f.git("checkout", "-q", "feature")
	got := f.report("")
	if !reflect.DeepEqual(got.Commits, []string{common, one, two}) {
		t.Errorf("default commits = %#v; want [%s %s %s]", got.Commits, common, one, two)
	}
	for _, commit := range got.Commits {
		if commit == reviewOnly {
			t.Errorf("default commits include configured-base-only commit %s", reviewOnly)
		}
	}
}

func Test_CHK_071_CHK_072_HistoryPhaseContentsAndDeletionControl(t *testing.T) {
	f := newHistoryFixture(t)
	path := "changes/0001-work.md"
	f.change(path, "evals", "")
	f.commit("add queued change")
	f.change(path, "audit", "")
	f.put("locked/config.txt", "worker edit\n")
	worker := f.commit("ordinary subject")
	f.remove(path)
	f.put("locked/config.txt", "author edit\n")
	f.commit("fireweed impl 0001-work")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got, historyFinding{"H01", "locked/config.txt", worker})
}

func Test_CHK_072_CHK_073_CHK_077_HistoryPlaybookOverrideAndProjectedInterface(t *testing.T) {
	f := newHistoryFixture(t)
	pending := "changes/0001-work.md"
	archive := "changes/archive/0001-work.md"
	set := "set:\n  interface: [newguard/**]\n"
	f.change(pending, "impl", set)
	f.commit("ready")
	projected, err := project.Project(f.root, "")
	if err != nil {
		t.Fatalf("Project with interface set: %v", err)
	}
	if len(projected.Modules) != 1 || projected.Modules[0].Name != "alpha" || !reflect.DeepEqual(projected.Modules[0].Zones.Interface, []string{"newguard/**"}) {
		t.Fatalf("projected alpha interface = %#v; want [newguard/**]", projected.Modules)
	}
	f.move(pending, archive)
	f.change(archive, "applied", set)
	f.put("packages/alpha/newguard/api.go", "package newguard\n")
	f.put(".fireweed/playbooks/evals.md", "local override\n")
	applied := f.commit("phase assigned")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got,
		historyFinding{"H02", "packages/alpha/newguard/api.go", applied},
		historyFinding{"H01", ".fireweed/playbooks/evals.md", applied})
}

func Test_CHK_071_CHK_073_HistoryAppliedImplTouchesEvalsAndInterface(t *testing.T) {
	f := newHistoryFixture(t)
	pending := "changes/0001-work.md"
	archive := "changes/archive/0001-work.md"
	f.change(pending, "impl", "")
	f.commit("ready for implementation")
	f.move(pending, archive)
	f.change(archive, "applied", "")
	f.put("packages/alpha/evals/existing_test.go", "changed test\n")
	f.put("packages/alpha/api.go", "changed api\n")
	applied := f.commit("unrelated subject")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got,
		historyFinding{"H02", "packages/alpha/evals/existing_test.go", applied},
		historyFinding{"H02", "packages/alpha/api.go", applied})
}

func Test_CHK_074_HistoryAuthorAppendOnlyModifyDeleteRename(t *testing.T) {
	for _, action := range []string{"modify", "delete", "rename"} {
		t.Run(action, func(t *testing.T) {
			f := newHistoryFixture(t)
			path := "packages/alpha/ledger/entry.txt"
			switch action {
			case "modify":
				f.put(path, "modified\n")
			case "delete":
				f.remove(path)
			case "rename":
				f.move(path, "packages/alpha/ledger/renamed.txt")
			}
			commit := f.commit("human maintenance")
			got := f.report(f.base + "..HEAD")
			if len(got.Diagnostics) == 0 {
				t.Errorf("%s diagnostics = %#v; want H03", action, got.Diagnostics)
			}
			for _, d := range got.Diagnostics {
				if d.Code != "H03" || !strings.Contains(d.Message, commit) || (d.Path != path && d.Path != "packages/alpha/ledger/renamed.txt") {
					t.Errorf("%s unexpected diagnostic %#v", action, d)
				}
			}
		})
	}
}

func Test_CHK_074_HistoryMergeResolutionInMergeOnlyRange(t *testing.T) {
	f := newHistoryFixture(t)
	f.git("checkout", "-qb", "side")
	f.put("packages/alpha/ledger/entry.txt", "side\n")
	f.commit("side edits ledger")
	f.git("checkout", "-q", "main")
	f.put("packages/alpha/ledger/entry.txt", "main\n")
	f.commit("main edits ledger")
	result, err := gitstore.Run(context.Background(), gitstore.Request{Args: []string{"git", "-C", f.root, "merge", "--no-ff", "--no-commit", "side"}, Dir: f.root})
	if err == nil || !strings.Contains(result.Stdout+result.Stderr, "CONFLICT") {
		t.Fatalf("expected ordinary merge conflict, got %v: %s %s", err, result.Stdout, result.Stderr)
	}
	f.put("packages/alpha/ledger/entry.txt", "resolved uniquely\n")
	merge := f.commit("resolve conflict")
	got := f.report(merge + "^!")
	if !reflect.DeepEqual(got.Commits, []string{merge}) {
		t.Errorf("merge-only selection = %#v; want [%s]", got.Commits, merge)
	}
	expectHistoryFindings(t, got, historyFinding{"H03", "packages/alpha/ledger/entry.txt", merge})
}

func Test_CHK_075_HistoryArchiveModifyAndDeletionControl(t *testing.T) {
	f := newHistoryFixture(t)
	path := "changes/archive/0001-old.md"
	f.change(path, "applied", "")
	f.commit("archived baseline")
	f.put(path, "---\nmodule: alpha\nphase: applied\n---\n\nUpdated archive.\n")
	modified := f.commit("edit archive")
	f.remove(path)
	f.commit("delete archive")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got, historyFinding{"H04", path, modified})
}

func Test_CHK_076_HistoryMultipleSquashedChangesKeepH03H04(t *testing.T) {
	f := newHistoryFixture(t)
	archiveOld := "changes/archive/0000-old.md"
	f.change(archiveOld, "applied", "")
	archiveBaseline := f.commit("existing archive")
	for _, path := range []string{"changes/archive/0001-one.md", "changes/archive/0002-two.md"} {
		f.change(path, "applied", "")
	}
	f.put("locked/config.txt", "squash protection\n")
	f.put("packages/alpha/evals/existing_test.go", "squash eval\n")
	f.put("packages/alpha/ledger/entry.txt", "squash ledger\n")
	f.put(archiveOld, "---\nmodule: alpha\nphase: applied\n---\n\nChanged old archive.\n")
	squash := f.commit("squashed changes")
	got := f.report(archiveBaseline + "..HEAD")
	wantSquashed := map[checks.SquashedChange]bool{
		{Commit: squash, Change: "changes/archive/0001-one.md"}: true,
		{Commit: squash, Change: "changes/archive/0002-two.md"}: true,
	}
	if len(got.Squashed) != len(wantSquashed) {
		t.Errorf("squashed = %#v; want %#v", got.Squashed, wantSquashed)
	} else {
		for _, entry := range got.Squashed {
			if !wantSquashed[entry] {
				t.Errorf("unexpected squashed entry %#v; want %#v", entry, wantSquashed)
			}
			delete(wantSquashed, entry)
		}
		if len(wantSquashed) != 0 {
			t.Errorf("missing squashed entries %#v", wantSquashed)
		}
	}
	expectHistoryFindings(t, got,
		historyFinding{"H03", "packages/alpha/ledger/entry.txt", squash},
		historyFinding{"H04", archiveOld, squash})
}

func Test_CHK_077_HistoryUsesRecordedBeforeAfterZones(t *testing.T) {
	f := newHistoryFixture(t)
	pending := "changes/0001-work.md"
	f.change(pending, "impl", "")
	f.commit("ready")
	f.move(pending, "changes/archive/0001-work.md")
	f.change("changes/archive/0001-work.md", "applied", "")
	f.put("packages/alpha/evals/new_test.go", "new eval\n")
	applied := f.commit("apply")
	f.put("fireweed.yaml", strings.Replace(historyConfig, "evals: [evals/**]", "evals: [later/**]", 1))
	f.commit("later config")
	got := f.report(f.base + ".." + applied)
	expectHistoryFindings(t, got, historyFinding{"H02", "packages/alpha/evals/new_test.go", applied})
}

func Test_CHK_077_HistoryChecksBothRecordedZoneMaps(t *testing.T) {
	f := newHistoryFixture(t)
	f.put("fireweed.yaml", strings.Replace(historyConfig, "append_only: [ledger/**]", "append_only: [newledger/**]", 1))
	f.put("packages/alpha/ledger/entry.txt", "changed under before zone\n")
	f.put("packages/alpha/newledger/entry.txt", "changed under after zone\n")
	commit := f.commit("author changes zone and both existing files")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got,
		historyFinding{"H03", "packages/alpha/ledger/entry.txt", commit},
		historyFinding{"H03", "packages/alpha/newledger/entry.txt", commit})
}

func Test_CHK_078_HistoryRejectsUnreadableRangeAndRecordedDocuments(t *testing.T) {
	for _, tc := range []struct {
		name       string
		edit       func(historyFixture)
		rangeValue string
	}{
		{"invalid_range", nil, "missing-ref..HEAD"},
		{"missing_config", func(f historyFixture) { f.remove("fireweed.yaml") }, ""},
		{"malformed_config", func(f historyFixture) { f.put("fireweed.yaml", "modules: [broken\n") }, ""},
		{"missing_SPEC", func(f historyFixture) { f.remove("packages/alpha/SPEC.md") }, ""},
		{"malformed_change", func(f historyFixture) { f.put("changes/0001-work.md", "---\nmodule: [broken\n---\n") }, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := newHistoryFixture(t)
			if tc.edit != nil {
				tc.edit(f)
				f.commit("record invalid input")
			}
			revisionRange := tc.rangeValue
			if revisionRange == "" {
				revisionRange = f.base + "..HEAD"
			}
			if _, err := checks.History(f.root, revisionRange); err == nil {
				t.Errorf("History(%q) accepted %s", revisionRange, tc.name)
			}
		})
	}
}
package evals_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func boundarySnapshot(t *testing.T, hook any) project.Snapshot {
	t.Helper()
	root := t.TempDir()
	snap := project.Snapshot{Root: root, Config: map[string]any{}, Modules: []project.Module{{
		Name: "alpha", Dir: "packages/alpha",
		Zones:    contracts.Zones{Disposable: []string{"internal/**"}},
		Document: contracts.Document{Frontmatter: map[string]any{}},
	}}}
	if hook != nil {
		snap.Config["hooks"] = map[string]any{"boundaries": hook}
	}
	return snap
}

func boundaryScript(t *testing.T, root, body string) []any {
	t.Helper()
	path := filepath.Join(root, "boundary-hook.sh")
	if err := os.WriteFile(path, []byte("#!/bin/sh\n"+body+"\n"), 0700); err != nil {
		t.Fatal(err)
	}
	return []any{path}
}

func boundaryHasError(diags []project.Diagnostic, err error) bool {
	if err != nil {
		return true
	}
	for _, d := range diags {
		if !d.Warning {
			return true
		}
	}
	return false
}

func boundaryHasWarning(diags []project.Diagnostic) bool {
	for _, d := range diags {
		if d.Warning {
			return true
		}
	}
	return false
}

func requireBoundaryDiagnostic(t *testing.T, diags []project.Diagnostic, code string, warning bool, message string) {
	t.Helper()
	for _, d := range diags {
		if message == "" {
			if d.Code == code && d.Warning == warning {
				return
			}
			continue
		}
		if strings.Contains(d.Message, message) {
			if d.Code != code || d.Warning != warning {
				t.Errorf("boundary diagnostic for %q = %#v; want code %s, warning %t", message, d, code, warning)
			}
			return
		}
	}
	t.Errorf("boundary diagnostic for %q absent: %#v", message, diags)
}

func Test_CHK_005_GlobalAndModuleBoundaryHooks(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	globalMarker := filepath.Join(snap.Root, "global-ran")
	moduleMarker := filepath.Join(snap.Root, "module-ran")
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, fmt.Sprintf("printf run > %q; printf 'alpha: global violation\\n'", globalMarker))}
	moduleHook := filepath.Join(snap.Root, "module-hook.sh")
	if err := os.WriteFile(moduleHook, []byte(fmt.Sprintf("#!/bin/sh\nprintf run > %q; printf 'beta: module violation\\n'\n", moduleMarker)), 0700); err != nil {
		t.Fatal(err)
	}
	snap.Modules = append(snap.Modules, project.Module{Name: "beta", Dir: "packages/beta", Zones: contracts.Zones{Disposable: []string{"internal/**"}}, Document: contracts.Document{Frontmatter: map[string]any{"hooks": map[string]any{"boundaries": []any{moduleHook}}}}})
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: 2 * time.Second})
	if !boundaryHasError(diags, err) {
		t.Errorf("unsuppressed violations accepted: %#v, %v", diags, err)
	}
	for _, module := range []string{"alpha", "beta"} {
		if !strings.Contains(fmt.Sprint(diags, err), module) {
			t.Errorf("violation for %s not attributed: %#v, %v", module, diags, err)
		}
	}
	for _, path := range []string{globalMarker, moduleMarker} {
		if _, statErr := os.Stat(path); statErr != nil {
			t.Errorf("applicable hook did not run: %s: %v", path, statErr)
		}
	}
}

func Test_CHK_041_ExactBaselineOnly(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	if err := os.MkdirAll(filepath.Join(snap.Root, ".fireweed"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(snap.Root, ".fireweed", "boundaries-baseline"), []byte("alpha: old violation\n"), 0644); err != nil {
		t.Fatal(err)
	}
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: old violation\n'`)}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) {
		t.Errorf("exact baseline line rejected: %#v, %v", diags, err)
	}
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: old violation changed\n'`)}
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("non-exact baseline line suppressed")
	}
}

func Test_CHK_039_CHK_040_OptOutSuppressesOnlyItsModuleAndWarns(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	snap.Modules[0].Document.Frontmatter["boundaries_opt_out"] = "legacy dependency"
	snap.Modules = append(snap.Modules, project.Module{Name: "beta", Dir: "packages/beta", Zones: contracts.Zones{Disposable: []string{"internal/**"}}, Document: contracts.Document{Frontmatter: map[string]any{}}})
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: ignored\n'`)}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) || !boundaryHasWarning(diags) {
		t.Errorf("reasoned opt-out = %#v, %v", diags, err)
	}
	if !strings.Contains(fmt.Sprint(diags), "legacy dependency") {
		t.Errorf("opt-out reason absent from warning: %#v", diags)
	}
	requireBoundaryDiagnostic(t, diags, "X03", true, "legacy dependency")
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: ignored\nbeta: enforced\n'`)}
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("opt-out hid another module's violation")
	}
	if err == nil {
		t.Error("unsuppressed beta violation returned no error")
	}
	requireBoundaryDiagnostic(t, diags, "X01", false, "beta: enforced")
}

func Test_CHK_044_CHK_045_MissingHookRequiredVersusOptional(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("missing hook for disposable module accepted")
	}
	if err == nil {
		t.Error("missing required hook returned no error")
	}
	requireBoundaryDiagnostic(t, diags, "X02", false, "alpha: missing boundaries hook")
	snap.Modules[0].Zones.Disposable = nil
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) || !boundaryHasWarning(diags) {
		t.Errorf("missing optional hook = %#v, %v", diags, err)
	}
	requireBoundaryDiagnostic(t, diags, "X04", true, "no boundary enforcement required")
	snap.Modules[0].Zones.Disposable = []string{"internal/**"}
	snap.Modules[0].Document.Frontmatter["boundaries_opt_out"] = "legacy"
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) || !boundaryHasWarning(diags) {
		t.Errorf("missing hook with reasoned opt-out = %#v, %v", diags, err)
	}
	requireBoundaryDiagnostic(t, diags, "X03", true, "legacy")
	requireBoundaryDiagnostic(t, diags, "X04", true, "no boundary enforcement required")
}

func Test_CHK_041_CHK_042_NormalNonzeroExitCanUseBaselineAndOptOut(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	if err := os.MkdirAll(filepath.Join(snap.Root, ".fireweed"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(snap.Root, ".fireweed", "boundaries-baseline"), []byte("alpha: known\n"), 0644); err != nil {
		t.Fatal(err)
	}
	snap.Modules[0].Document.Frontmatter["boundaries_opt_out"] = "legacy"
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `printf 'alpha: known\nalpha: opted out\n'; exit 3`)}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if boundaryHasError(diags, err) {
		t.Errorf("normally exited hook with fully suppressed lines rejected: %#v, %v", diags, err)
	}
}

func Test_CHK_043_LaunchCancellationAndTimeoutStayErrors(t *testing.T) {
	snap := boundarySnapshot(t, nil)
	snap.Modules[0].Document.Frontmatter["boundaries_opt_out"] = "legacy"
	snap.Config["hooks"] = map[string]any{"boundaries": []any{filepath.Join(snap.Root, "missing-executable")}}
	diags, err := checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("launch failure hidden by opt-out")
	}
	if err == nil {
		t.Error("launch failure returned no error")
	} else if !strings.Contains(err.Error(), "missing-executable") || strings.HasPrefix(err.Error(), "X02:") {
		t.Errorf("launch error = %q; want missing executable cause without X02 prefix", err)
	}
	requireBoundaryDiagnostic(t, diags, "X02", false, "")
	snap.Config["hooks"] = map[string]any{"boundaries": boundaryScript(t, snap.Root, `sleep 2`)}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	diags, err = checks.Boundaries(ctx, checks.BoundaryRequest{Snapshot: snap, Timeout: time.Second})
	if !boundaryHasError(diags, err) {
		t.Error("cancelled hook hidden by opt-out")
	}
	if err == nil {
		t.Error("cancelled hook returned no error")
	} else if !errors.Is(err, context.Canceled) || strings.HasPrefix(err.Error(), "X02:") {
		t.Errorf("cancellation error = %q; want canceled cause without X02 prefix", err)
	}
	requireBoundaryDiagnostic(t, diags, "X02", false, "")
	start := time.Now()
	diags, err = checks.Boundaries(context.Background(), checks.BoundaryRequest{Snapshot: snap, Timeout: 80 * time.Millisecond})
	if !boundaryHasError(diags, err) {
		t.Error("timed-out hook hidden by opt-out")
	}
	if err == nil {
		t.Error("timed-out hook returned no error")
	} else if !errors.Is(err, context.DeadlineExceeded) || strings.HasPrefix(err.Error(), "X02:") {
		t.Errorf("timeout error = %q; want deadline cause without X02 prefix", err)
	}
	requireBoundaryDiagnostic(t, diags, "X02", false, "")
	if time.Since(start) > 3*time.Second {
		t.Error("boundary timeout did not bound execution")
	}
}
package evals_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

func zoneFixture(t *testing.T) (project.Snapshot, string) {
	t.Helper()
	root := t.TempDir()
	for path := range map[string]string{
		"packages/alpha/SPEC.md": "alpha", "packages/beta/SPEC.md": "beta",
		"packages/alpha/evals/old_test.go": "old", "packages/alpha/api.go": "api",
		"packages/alpha/internal/existing.go": "old", "packages/alpha/gen/old.sql": "old",
		"packages/alpha/keep.txt": "old", "changes/0011-current.md": "current",
	} {
		full := filepath.Join(root, path)
		if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte("fixture"), 0644); err != nil {
			t.Fatal(err)
		}
	}
	run := func(args ...string) string {
		cmd := exec.Command("git", args...)
		cmd.Dir = root
		cmd.Env = append(os.Environ(), "GIT_AUTHOR_NAME=Fixture", "GIT_AUTHOR_EMAIL=fixture@example.test", "GIT_COMMITTER_NAME=Fixture", "GIT_COMMITTER_EMAIL=fixture@example.test")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("git %v: %v: %s", args, err, out)
		}
		return string(out)
	}
	run("init", "-q")
	run("add", ".")
	run("commit", "-qm", "base")
	base := run("rev-parse", "HEAD")
	zones := contracts.Zones{Evals: []string{"evals/**"}, Interface: []string{"api.go"}, Disposable: []string{"internal/**"}, Preserve: []string{"keep.txt"}, AppendOnly: []string{"gen/**"}}
	snap := project.Snapshot{Root: root, Config: map[string]any{"protected": []any{"locked/**"}, "worker_writable": []any{"go.sum"}}, Modules: []project.Module{
		{Name: "alpha", Dir: "packages/alpha", Zones: zones},
		{Name: "beta", Dir: "packages/beta", Zones: zones},
	}}
	return snap, base[:len(base)-1]
}

func zoneRequest(snap project.Snapshot, base, phase string, changes ...gitstore.Change) checks.ZoneRequest {
	return checks.ZoneRequest{Snapshot: snap, Module: "alpha", Phase: phase, ChangePath: "changes/0011-current.md", Base: base, InterfaceDeclared: true, Changes: changes}
}

func zoneAllowed(t *testing.T, request checks.ZoneRequest) {
	t.Helper()
	if diags := checks.Zones(request); len(diags) != 0 {
		t.Errorf("wanted allowed change, got %#v", diags)
	}
}

func zoneRejected(t *testing.T, request checks.ZoneRequest) {
	t.Helper()
	if diags := checks.Zones(request); len(diags) == 0 {
		t.Errorf("forbidden change accepted: %#v", request.Changes)
	}
}

func Test_CHK_004_AuditOnlyCurrentChange(t *testing.T) {
	snap, base := zoneFixture(t)
	zoneAllowed(t, zoneRequest(snap, base, "audit", gitstore.Change{Path: "changes/0011-current.md", Status: "M"}))
	for _, path := range []string{"changes/0012-other.md", "packages/alpha/evals/old_test.go", "go.sum"} {
		zoneRejected(t, zoneRequest(snap, base, "audit", gitstore.Change{Path: path, Status: "M"}))
	}
}

func Test_CHK_006_CHK_007_CHK_008_CHK_009_RejectProtectedSpecForeignModuleAndChangeEveryPhase(t *testing.T) {
	snap, base := zoneFixture(t)
	for _, phase := range []string{"evals", "impl", "audit"} {
		for _, path := range []string{"locked/config", ".fireweed/hooks.go", "packages/alpha/SPEC.md", "packages/beta/internal/code.go", "changes/0012-other.md"} {
			zoneRejected(t, zoneRequest(snap, base, phase, gitstore.Change{Path: path, Status: "M"}))
		}
	}
}

func Test_CHK_011_CHK_012_CHK_013_CHK_014_CHK_015_CHK_021_EvalsAllowedZonesAndNewStubs(t *testing.T) {
	snap, base := zoneFixture(t)
	for _, change := range []gitstore.Change{
		{Path: "packages/alpha/evals/new_test.go", Status: "A"},
		{Path: "packages/alpha/api.go", Status: "M"},
		{Path: "packages/alpha/internal/new_stub.go", Status: "A"},
		{Path: "packages/alpha/manifest.json", Status: "M"},
		{Path: "changes/0011-current.md", Status: "M"},
	} {
		zoneAllowed(t, zoneRequest(snap, base, "evals", change))
	}
	zoneRejected(t, zoneRequest(snap, base, "evals", gitstore.Change{Path: "packages/alpha/internal/existing.go", Status: "M"}))
	zoneRejected(t, zoneRequest(snap, base, "evals", gitstore.Change{Path: "packages/alpha/keep.txt", Status: "M"}))
	request := zoneRequest(snap, base, "evals", gitstore.Change{Path: "packages/alpha/api.go", Status: "M"})
	request.InterfaceDeclared = false
	zoneRejected(t, request)
}

func Test_CHK_016_CHK_017_CHK_018_CHK_019_CHK_020_CHK_021_ImplFreezesEvalsAndInterfaces(t *testing.T) {
	snap, base := zoneFixture(t)
	for _, path := range []string{"packages/alpha/evals/old_test.go", "packages/alpha/api.go"} {
		zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: path, Status: "M"}))
	}
	for _, path := range []string{"packages/alpha/internal/existing.go", "packages/alpha/keep.txt", "packages/alpha/manifest.json", "changes/0011-current.md"} {
		zoneAllowed(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: path, Status: "M"}))
	}
}

func Test_CHK_022_CHK_023_ImplAppendOnlyNewVersusExisting(t *testing.T) {
	snap, base := zoneFixture(t)
	zoneAllowed(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: "packages/alpha/gen/new.sql", Status: "A"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: "packages/alpha/gen/old.sql", Status: "M"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{Path: "packages/alpha/gen/old.sql", Status: "D"}))
	zoneRejected(t, zoneRequest(snap, base, "evals", gitstore.Change{Path: "packages/alpha/gen/new.sql", Status: "A"}))
}

func Test_CHK_024_RenamesCheckBothPaths(t *testing.T) {
	snap, base := zoneFixture(t)
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{OldPath: "packages/beta/internal/old.go", Path: "packages/alpha/internal/new.go", Status: "R100"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{OldPath: "packages/alpha/internal/existing.go", Path: "packages/beta/internal/new.go", Status: "R100"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{OldPath: "packages/alpha/gen/old.sql", Path: "packages/alpha/internal/new.go", Status: "R100"}))
	zoneRejected(t, zoneRequest(snap, base, "impl", gitstore.Change{OldPath: "packages/alpha/internal/existing.go", Path: "packages/alpha/gen/new.sql", Status: "R100"}))
}

func Test_CHK_010_WorkerWritableOnlyEvalsAndImpl(t *testing.T) {
	snap, base := zoneFixture(t)
	for _, phase := range []string{"evals", "impl"} {
		zoneAllowed(t, zoneRequest(snap, base, phase, gitstore.Change{Path: "go.sum", Status: "M"}))
		zoneRejected(t, zoneRequest(snap, base, phase, gitstore.Change{Path: "outside.txt", Status: "M"}))
	}
	zoneRejected(t, zoneRequest(snap, base, "audit", gitstore.Change{Path: "go.sum", Status: "M"}))
}

codex
The existing disposable implementation is present, but the workspace has no Go manifest or dependent internals, so the retained evals cannot run directly here. I’m checking the implementation against those evals and the spec, especially Doctor and History, where fixture details matter.
exec
/bin/zsh -lc "rg -n '"'^func Test|''^func new|''^func .*Fixture|''^func .*fixture|''^func .*probe|''^func .*want'"' packages/checks/evals/doctor*.go packages/checks/evals/history_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc "sed -n '300,620p' packages/checks/evals/doctor_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc "sed -n '1,300p' packages/checks/evals/doctor_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc 'cat packages/checks/evals/doctor_agent_logs_test.go packages/checks/evals/doctor_prompt_file_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 621ms:
package evals_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"gopkg.in/yaml.v3"
)

const doctorGuidance = "<!-- fireweed:start -->\nThis project uses fireweed. Specs and change files are the source of truth.\nRun `fireweed playbook author-change` before authoring changes and\n`fireweed context <change>` before each step.\n<!-- fireweed:end -->\n"

type doctorFixture struct {
	t    *testing.T
	root string
	log  string
	bin  string
}

func (f doctorFixture) put(name, body string) {
	f.t.Helper()
	path := filepath.Join(f.root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		f.t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		f.t.Fatal(err)
	}
}

func (f doctorFixture) git(args ...string) string {
	f.t.Helper()
	cmd := exec.Command("git", args...)
	cmd.Dir = f.root
	cmd.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_TERMINAL_PROMPT=0")
	out, err := cmd.CombinedOutput()
	if err != nil {
		f.t.Fatalf("git %v: %v: %s", args, err, out)
	}
	return string(out)
}

func (f doctorFixture) commit() {
	f.t.Helper()
	f.git("init", "-b", "main")
	f.git("config", "user.name", "Doctor Fixture")
	f.git("config", "user.email", "doctor@example.test")
	f.git("add", ".")
	f.git("commit", "-m", "fixture")
}

func (f doctorFixture) request(noAgent bool) checks.DoctorRequest {
	return checks.DoctorRequest{Root: f.root, Bin: f.bin, NoAgent: noAgent, Timeout: 15 * time.Second}
}

func doctorConfig(t *testing.T, recipe assets.Recipe) string {
	t.Helper()
	config := map[string]any{
		"fireweed": ">=0.1, <0.2", "version": 1, "base": "main", "modules": []string{"packages/*"},
		"limit_patterns":  []string{"rate limit"},
		"module_defaults": map[string]any{"interface": recipe.ModuleDefaults.Interface, "evals": recipe.ModuleDefaults.Evals, "disposable": recipe.ModuleDefaults.Disposable},
		"hooks":           recipe.Hooks,
		"agent":           map[string]any{"command": []string{"bash", ".fireweed/agent.sh", "{prompt}", "{git_common_dir}"}, "wrap": []string{"bash", ".fireweed/wrap.sh"}},
		"audit":           map[string]any{"enabled": true, "command": []string{"bash", ".fireweed/audit.sh", "{prompt}"}},
	}
	data, err := yaml.Marshal(config)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func (f doctorFixture) editConfig(change func(map[string]any)) {
	f.t.Helper()
	var config map[string]any
	if err := yaml.Unmarshal(mustRead(f.t, filepath.Join(f.root, "fireweed.yaml")), &config); err != nil {
		f.t.Fatal(err)
	}
	change(config)
	data, err := yaml.Marshal(config)
	if err != nil {
		f.t.Fatal(err)
	}
	f.put("fireweed.yaml", string(data))
}

func newDoctorFixture(t *testing.T) doctorFixture {
	t.Helper()
	r, err := assets.BuiltinRecipe("go")
	if err != nil {
		t.Fatal(err)
	}
	f := doctorFixture{t: t, root: t.TempDir(), log: filepath.Join(t.TempDir(), "probes.log")}
	f.put("go.mod", "module example.test/doctorfixture\n\ngo 1.23\n")
	f.put("fireweed.yaml", doctorConfig(t, r))
	f.put("AGENTS.md", "Before.\n"+doctorGuidance+"After.\n")
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nThe fixture MUST pass.\n")
	f.put("packages/alpha/api.go", "package alpha\n")
	for path, body := range r.Files {
		f.put(path, body)
	}
	f.put(".fireweed/wrap.sh", fmt.Sprintf("#!/bin/bash\nprintf 'wrap:%%s\n' \"$PWD\" >> %q\nexec \"$@\"\n", f.log))
	f.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/bash\nprintf 'worker:%%s:%%s:%%s\n' \"$PWD\" \"$1\" \"$2\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	f.put(".fireweed/audit.sh", fmt.Sprintf("#!/bin/bash\nprintf 'audit:%%s:%%s\n' \"$PWD\" \"$1\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	return f
}

func localDoctorFixture(t *testing.T) doctorFixture {
	t.Helper()
	f := doctorFixture{t: t, root: t.TempDir(), log: filepath.Join(t.TempDir(), "probes.log")}
	f.bin = filepath.Join(t.TempDir(), "fake-fireweed")
	if err := os.WriteFile(f.bin, []byte(`#!/bin/sh
if test "$1" != modules || test "$2" != --json; then exit 2; fi
printf '{"data":['
first=1
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test "$first" = 0; then printf ','; fi; first=0
  printf '{"name":"%s","dir":"%s","depends_on":[],"zones":{"interface":["api.txt"],"evals":["evals/**"],"disposable":["internal/**"]}}' "$name" "$dir"
done
printf ']}'
`), 0700); err != nil {
		t.Fatal(err)
	}
	recipe := assets.Recipe{ModuleDefaults: structZones(), Hooks: map[string]any{
		"eval":       "sh .fireweed/eval.sh {module} {module_dir} {report}",
		"boundaries": "sh .fireweed/boundary.sh",
		"setup":      "sh .fireweed/setup.sh",
	}}
	f.put("fireweed.yaml", doctorConfig(t, recipe))
	f.put(".gitignore", ".fireweed/runs/\n")
	f.put("AGENTS.md", doctorGuidance)
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nThe fixture MUST pass.\n")
	f.put("packages/alpha/api.txt", "public\n")
	f.put(".fireweed/templates/module/SPEC.md", "---\nname: {{name}}\nid_prefix: {{prefix}}\ndepends_on: []\n---\n# {{name}}\n\n## Specs\n\n### {{prefix}}-001\nThe canary MUST pass.\n")
	f.put(".fireweed/templates/module/api.txt.tmpl", "{{package}} {{go_module}} {{module_dir}} {{other_dir}}\n")
	f.put(".fireweed/canaries/canaries.yaml", `version: 1
baseline:
  SPEC.md: "---\nname: {{name}}\nid_prefix: {{prefix}}\ndepends_on: []\n---\n# {{name}}\n\n## Specs\n\n### {{prefix}}-001\nThe canary MUST pass.\n"
  api.txt: "{{package}} {{go_module}} {{module_dir}} {{other_dir}}"
undeclared:
  undeclared.txt: "{{other_dir}}"
internal_eval:
  internal/private.txt: "private"
  evals/internal.txt: "{{module_dir}}"
failing_eval:
  evals/failing.txt: "{{prefix}}-001"
missing_id:
  evals/missing.txt: "UNMAPPED-999"
`)
	f.put(".fireweed/setup.sh", fmt.Sprintf("#!/bin/sh\nprintf 'setup:%%s\n' \"$PWD\" >> %q\n", f.log))
	f.put(".fireweed/eval.sh", fmt.Sprintf(`#!/bin/sh
printf 'eval:%%s:%%s:%%s\n' "$PWD" "$1" "$FIREWEED_BIN" >> %q
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
if test -f "$2/evals/failing.txt"; then printf '<testsuite><testcase name="%%s-001"><failure/></testcase></testsuite>' "$prefix" > "$3"; exit 1; fi
if test -f "$2/evals/missing.txt"; then printf '<testsuite><testcase name="UNMAPPED-999"/></testsuite>' > "$3"; exit 0; fi
printf '<testsuite><testcase name="%%s-001"/></testsuite>' "$prefix" > "$3"
`, f.log))
	f.put(".fireweed/boundary.sh", fmt.Sprintf(`#!/bin/sh
printf 'boundaries:%%s:%%s\n' "$PWD" "$FIREWEED_BIN" >> %q
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test -f "$dir/undeclared.txt"; then printf '%%s: undeclared import\n' "$name"; fi
  if test -f "$dir/evals/internal.txt"; then printf '%%s: interface violation\n' "$name"; fi
done
`, f.log))
	f.put(".fireweed/wrap.sh", fmt.Sprintf("#!/bin/bash\nprintf 'wrap:%%s\n' \"$PWD\" >> %q\nexec \"$@\"\n", f.log))
	f.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/bash\nprintf 'worker:%%s:%%s:%%s\n' \"$PWD\" \"$1\" \"$2\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	f.put(".fireweed/audit.sh", fmt.Sprintf("#!/bin/bash\nprintf 'audit:%%s:%%s\n' \"$PWD\" \"$1\" >> %q\nprintf 'FIREWEED_DOCTOR_OK\\n'\n", f.log))
	return f
}

func structZones() (z contracts.Zones) {
	z.Interface = []string{"api.txt"}
	z.Evals = []string{"evals/**"}
	z.Disposable = []string{"internal/**"}
	return
}

func doctorText(report checks.DoctorReport) string {
	return fmt.Sprintf("checks=%+v diagnostics=%+v", report.Checks, report.Diagnostics)
}
func hasDoctorDiagnostic(report checks.DoctorReport, code string, warning bool, terms ...string) bool {
	for _, d := range report.Diagnostics {
		if d.Code != code || d.Warning != warning {
			continue
		}
		ok := true
		for _, s := range terms {
			if !strings.Contains(strings.ToLower(d.Message), strings.ToLower(s)) {
				ok = false
			}
		}
		if ok {
			return true
		}
	}
	return false
}
func hasDoctorCheck(report checks.DoctorReport, name string, passed bool) bool {
	for _, c := range report.Checks {
		if strings.Contains(c.Name, name) && c.Passed == passed {
			return true
		}
	}
	return false
}
func doctorErrors(report checks.DoctorReport) bool {
	for _, d := range report.Diagnostics {
		if !d.Warning {
			return true
		}
	}
	return false
}
func doctorLog(f doctorFixture) string { b, _ := os.ReadFile(f.log); return string(b) }
func worktrees(f doctorFixture) string { return f.git("worktree", "list", "--porcelain") }

func logDoctorCanaryProbes(t *testing.T, f doctorFixture) {
	t.Helper()
	logProbe := fmt.Sprintf(`printf 'probe:%%s\n' "$PWD" >> %q
for marker in undeclared.txt internal.txt failing.txt missing.txt; do
  if find packages -name "$marker" | grep -q .; then printf 'planted:%%s\n' "$marker" >> %q; fi
done
`, f.log, f.log)
	for _, path := range []string{".fireweed/eval.sh", ".fireweed/boundary.sh"} {
		body := string(mustRead(t, filepath.Join(f.root, path)))
		if !strings.HasPrefix(body, "#!/bin/sh\n") {
			t.Fatalf("fixture hook %s has no shell header", path)
		}
		f.put(path, strings.Replace(body, "#!/bin/sh\n", "#!/bin/sh\n"+logProbe, 1))
	}
}

func assertScratchRemoved(t *testing.T, f doctorFixture) {
	t.Helper()
	common := strings.TrimSpace(f.git("rev-parse", "--git-common-dir"))
	if !filepath.IsAbs(common) {
		common = filepath.Join(f.root, common)
	}
	common, err := filepath.EvalSymlinks(common)
	if err != nil {
		t.Fatal(err)
	}
	base := filepath.Join(common, "fireweed", "worktrees")
	for _, line := range strings.Split(doctorLog(f), "\n") {
		if !strings.HasPrefix(line, "setup:") {
			continue
		}
		path := strings.TrimPrefix(line, "setup:")
		for parent := path; ; parent = filepath.Dir(parent) {
			resolved, err := filepath.EvalSymlinks(parent)
			if err == nil {
				suffix, err := filepath.Rel(parent, path)
				if err != nil {
					t.Fatal(err)
				}
				path = filepath.Join(resolved, suffix)
				break
			}
			if parent == filepath.Dir(parent) {
				t.Fatal(err)
			}
		}
		rel, err := filepath.Rel(base, path)
		if err != nil || rel == "." || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
			t.Errorf("scratch %q is not below common Git directory %q", path, base)
		}
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Errorf("physical scratch directory %q remains or cannot be checked: %v", path, err)
		}
		return
	}
	t.Error("setup did not record a scratch path")
}

func Test_CHK_049_StaticDiagnosticsStopProbes(t *testing.T) {
	f := localDoctorFixture(t)
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n## Specs\n### ALP-001\nA.\n### ALP-001\nB.\n")
	f.commit()
	want := project.Check(f.root)
	if len(want) == 0 || !doctorErrors(checks.DoctorReport{Diagnostics: want}) {
		t.Fatal("fixture has no static diagnostic")
	}
	report := checks.Doctor(context.Background(), f.request(false))
	remaining := map[project.Diagnostic]int{}
	for _, d := range report.Diagnostics {
		remaining[d]++

 succeeded in 622ms:
		remaining[d]++
	}
	for _, d := range want {
		if remaining[d] == 0 {
			t.Errorf("complete project diagnostic %+v lost: %s", d, doctorText(report))
		}
		remaining[d]--
	}
	if strings.Contains(doctorLog(f), "setup:") || strings.Contains(doctorLog(f), "eval:") || strings.Contains(doctorLog(f), "boundaries:") || strings.Contains(doctorLog(f), "worker:") || strings.Contains(doctorLog(f), "audit:") {
		t.Errorf("static error ran probes: %s", doctorLog(f))
	}
}

func Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_CHK_060_CHK_061_CHK_067_RealGoRecipe(t *testing.T) {
	f := newDoctorFixture(t)
	f.commit()
	f.bin = doctorBinary(t)
	f.editConfig(func(config map[string]any) {
		config["modules"] = []string{"packages/alpha"}
		config["hooks"].(map[string]any)["setup"] = fmt.Sprintf("printf 'setup:%%s\\n' \"$PWD\" >> %q; test -z \"$(git symbolic-ref -q HEAD)\" && test -f .fireweed/doctor-untracked && go mod download", f.log)
	})
	f.put(".fireweed/doctor-untracked", "ready\n")
	for _, d := range project.Check(f.root) {
		if !d.Warning {
			t.Fatalf("invalid Go doctor fixture: %+v", d)
		}
	}
	beforeStatus, beforeWorktrees := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) {
		t.Errorf("real Go recipe doctor failed: %s", doctorText(report))
	}
	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id", "recovery"} {
		if !hasDoctorCheck(report, name, true) {
			t.Errorf("missing passing %s outcome: %s", name, doctorText(report))
		}
	}
	recoveries := 0
	for _, check := range report.Checks {
		if strings.Contains(strings.ToLower(check.Name), "recovery") && check.Passed {
			recoveries++
		}
	}
	if recoveries < 4 {
		t.Errorf("clean recovery not reported after each canary: %s", doctorText(report))
	}
	modules := map[string]bool{}
	for _, check := range report.Checks {
		if check.Module != "" && check.Module != "alpha" {
			modules[check.Module] = true
		}
	}
	if len(modules) < 2 {
		t.Errorf("two distinct throwaway modules absent: %s", doctorText(report))
	}
	if !strings.Contains(doctorLog(f), "setup:") || !strings.Contains(doctorLog(f), "/fireweed/worktrees/") {
		t.Errorf("setup did not run in scratch: %s", doctorLog(f))
	}
	if strings.Contains(doctorLog(f), "worker:") || strings.Contains(doctorLog(f), "audit:") {
		t.Errorf("NoAgent ran agents: %s", doctorLog(f))
	}
	if got := f.git("status", "--porcelain", "--untracked-files=all"); got != beforeStatus {
		t.Errorf("caller changed: before %q after %q", beforeStatus, got)
	}
	if got := worktrees(f); got != beforeWorktrees {
		t.Errorf("scratch registration leaked: before %q after %q", beforeWorktrees, got)
	}
	assertScratchRemoved(t, f)
}

func Test_CHK_053_NestedModuleTemplatesRenderForBothCanaries(t *testing.T) {
	f := localDoctorFixture(t)
	f.put("go.mod", "module example.test/doctorfixture\n\ngo 1.23\n")
	fields := "{{name}}|{{prefix}}|{{package}}|{{go_module}}|{{module_dir}}|{{other_dir}}\n"
	f.put(".fireweed/templates/module/evals/nested.txt.tmpl", "eval|"+fields)
	f.put(".fireweed/templates/module/internal/nested.txt", "internal|"+fields)
	canaries := string(mustRead(t, filepath.Join(f.root, ".fireweed/canaries/canaries.yaml")))
	f.put(".fireweed/canaries/canaries.yaml", strings.Replace(canaries, "undeclared:\n", "  evals/baseline.txt: \"baseline|"+strings.TrimSpace(fields)+"\"\nundeclared:\n", 1))
	eval := string(mustRead(t, filepath.Join(f.root, ".fireweed/eval.sh")))
	check := fmt.Sprintf(`case "$1" in
  doctor*)
    prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
    other=
    for spec in packages/*/SPEC.md; do
      dir=${spec%%/SPEC.md}
      case "$dir" in packages/alpha|"$2") ;; *) other=$dir ;; esac
    done
    expected="$1|$prefix|$1|example.test/doctorfixture|$2|$other"
    test -n "$other" &&
      test "$(cat "$2/evals/nested.txt")" = "eval|$expected" &&
      test "$(cat "$2/internal/nested.txt")" = "internal|$expected" &&
      test "$(cat "$2/evals/baseline.txt")" = "baseline|$expected" &&
      test ! -e "$2/evals/nested.txt.tmpl" || exit 17
    printf 'rendered:%%s\n' "$1" >> %q
    ;;
esac
`, f.log)
	f.put(".fireweed/eval.sh", strings.Replace(eval, "#!/bin/sh\n", "#!/bin/sh\n"+check, 1))
	f.commit()
	before := worktrees(f)
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) {
		t.Errorf("nested template probes failed: %s", doctorText(report))
	}
	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id", "recovery"} {
		if !hasDoctorCheck(report, name, true) {
			t.Errorf("missing passing %s probe: %s", name, doctorText(report))
		}
	}
	modules := map[string]bool{}
	for _, line := range strings.Split(doctorLog(f), "\n") {
		if strings.HasPrefix(line, "rendered:") {
			modules[strings.TrimPrefix(line, "rendered:")] = true
		}
	}
	if len(modules) != 2 {
		t.Errorf("expected both throwaway modules to pass nested template checks, got %v: %s", modules, doctorText(report))
	}
	if got := worktrees(f); got != before {
		t.Errorf("scratch registration leaked: before %q after %q", before, got)
	}
	if strings.Contains(doctorLog(f), "setup:") {
		assertScratchRemoved(t, f)
	}
}

func Test_CHK_051_CHK_052_SetupFailureStopsAndCleans(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/setup.sh", fmt.Sprintf("#!/bin/sh\nprintf 'setup:%%s\\n' \"$PWD\" >> %q\necho failed setup >&2\nexit 17\n", f.log))
	f.commit()
	before := worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if !hasDoctorDiagnostic(report, "C14", false, "setup") {
		t.Errorf("setup failure absent: %s", doctorText(report))
	}
	if strings.Contains(doctorLog(f), "eval:") || strings.Contains(doctorLog(f), "boundaries:") || strings.Contains(doctorLog(f), "worker:") {
		t.Errorf("setup failure ran later probes: %s", doctorLog(f))
	}
	if got := worktrees(f); got != before {
		t.Errorf("scratch worktree leaked: %q", got)
	}
	assertScratchRemoved(t, f)
}

func Test_CHK_051_CancellationCleansScratchAndStopsSetupGroup(t *testing.T) {
	f := localDoctorFixture(t)
	ready := filepath.Join(t.TempDir(), "ready")
	childFinished := filepath.Join(t.TempDir(), "child-finished")
	f.put(".fireweed/setup.sh", fmt.Sprintf(`#!/bin/sh
printf 'setup:%%s\n' "$PWD" >> %q
sh -c 'sleep 2; printf finished > "$1"' sh %q &
printf ready > %q
wait
`, f.log, childFinished, ready))
	f.commit()
	before := worktrees(f)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reports := make(chan checks.DoctorReport, 1)
	go func() { reports <- checks.Doctor(ctx, f.request(true)) }()
	deadline := time.After(8 * time.Second)
	for {
		if _, err := os.Stat(ready); err == nil {
			break
		} else if !os.IsNotExist(err) {
			t.Fatal(err)
		}
		select {
		case report := <-reports:
			t.Fatalf("setup exited before cancellation: %s", doctorText(report))
		case <-deadline:
			t.Fatal("setup never signalled readiness")
		case <-time.After(10 * time.Millisecond):
		}
	}
	cancel()
	select {
	case report := <-reports:
		if !hasDoctorDiagnostic(report, "C14", false, "setup") {
			t.Errorf("cancelled setup lacked C14: %s", doctorText(report))
		}
	case <-time.After(5 * time.Second):
		t.Fatal("cancelled doctor did not return promptly")
	}
	if got := worktrees(f); got != before {
		t.Errorf("cancelled doctor retained Git registration: before %q after %q", before, got)
	}
	assertScratchRemoved(t, f)
	time.Sleep(2200 * time.Millisecond)
	if _, err := os.Stat(childFinished); !os.IsNotExist(err) {
		t.Errorf("cancelled setup descendant continued: %v", err)
	}
}

func Test_CHK_054_MissingTemplatesAndCanaryMaps(t *testing.T) {
	for _, tc := range []struct{ name, path, body, term, input, affected string }{
		{"missing template", ".fireweed/templates/module/SPEC.md", "", "template", "SPEC.md", "undeclared"},
		{"malformed template", ".fireweed/templates/module/SPEC.md", "---\nname: [broken\nid_prefix: {{prefix}}\n---\n", "template", "SPEC.md", "undeclared"},
		{"missing baseline", ".fireweed/canaries/canaries.yaml", "version: 1\nbaseline: {}\n", "canary", "baseline", "undeclared"},
		{"missing internal_eval", ".fireweed/canaries/canaries.yaml", "missing_id:\n  evals/missing.txt: missing\n", "internal_eval", "internal_eval", "internal_eval"},
		{"wrong version", ".fireweed/canaries/canaries.yaml", "version: 2\nbaseline: {}\n", "canary", "canaries.yaml", "undeclared"},
		{"malformed canary", ".fireweed/canaries/canaries.yaml", "version: [bad\n", "canary", "canaries.yaml", "undeclared"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			if tc.name == "missing baseline" {
				body := string(mustRead(t, filepath.Join(f.root, tc.path)))
				start, end := strings.Index(body, "baseline:\n"), strings.Index(body, "undeclared:\n")
				if start < 0 || end <= start {
					t.Fatal("fixture cannot isolate baseline map")
				}
				tc.body = body[:start] + "baseline: {}\n" + body[end:]
			}
			if tc.name == "missing internal_eval" {
				body := string(mustRead(t, filepath.Join(f.root, tc.path)))
				tc.body = strings.Replace(body, "internal_eval:\n  internal/private.txt: \"private\"\n  evals/internal.txt: \"{{module_dir}}\"\n", "", 1)
				if tc.body == body {
					t.Fatal("fixture did not remove internal_eval map")
				}
			}
			if tc.body == "" {
				if err := os.Remove(filepath.Join(f.root, tc.path)); err != nil {
					t.Fatal(err)
				}
			} else {
				f.put(tc.path, tc.body)
			}
			f.commit()
			report := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(report, "C14", false, tc.term) {
				t.Errorf("missing C14: %s", doctorText(report))
			}
			identified := false
			for _, d := range report.Diagnostics {
				if d.Code == "C14" && !d.Warning && strings.Contains(strings.ToLower(d.Message+" "+d.Path), strings.ToLower(tc.input)) {
					identified = true
				}
			}
			if !identified {
				t.Errorf("C14 did not identify %s: %s", tc.input, doctorText(report))
			}
			if hasDoctorCheck(report, tc.affected, true) {
				t.Errorf("invalid recipe passed canary: %s", doctorText(report))
			}
		})
	}
}

func Test_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_ZeroHookAndRecovery(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/boundary.sh", "#!/bin/sh\nexit 0\n")
	f.put(".fireweed/eval.sh", `#!/bin/sh
printf '<testsuite><testcase name="ALP-001"/></testsuite>' > "$3"
`)
	f.commit()
	before := worktrees(f)
	report := checks.Doctor(context.Background(), f.request(true))
	for _, tc := range []struct{ name, hook string }{{"undeclared", "boundar"}, {"internal_eval", "boundar"}, {"failing_eval", "eval"}, {"missing_id", "eval"}} {
		if !hasDoctorDiagnostic(report, "C14", false, tc.name, tc.hook) {
			t.Errorf("missed %s canary not diagnosed with hook: %s", tc.name, doctorText(report))
		}
	}
	if !doctorErrors(report) {
		t.Errorf("zero hooks accepted: %s", doctorText(report))
	}
	if got := worktrees(f); got != before {
		t.Errorf("failed canary leaked worktree: %s", got)
	}

	g := localDoctorFixture(t)
	g.put(".fireweed/boundary.sh", fmt.Sprintf(`#!/bin/sh
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test -f "$dir/undeclared.txt"; then printf '%%s' "$name" > .fireweed/recovery-seen; printf '%%s: undeclared import\n' "$name"; fi
done
if test -f .fireweed/recovery-seen && ! find packages -name undeclared.txt | grep -q .; then printf '%%s: lingering violation\n' "$(cat .fireweed/recovery-seen)"; fi
`))
	g.commit()
	r := checks.Doctor(context.Background(), g.request(true))
	if !hasDoctorDiagnostic(r, "C14", false, "recovery", "undeclared", "boundar") {
		t.Errorf("missing failed recovery diagnostic: %s", doctorText(r))
	}
}

func Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific(t *testing.T) {
	for _, tc := range []struct{ name, script, canary string }{
		{"missing baseline coverage", `#!/bin/sh
printf '<testsuite/>' > "$3"
`, "baseline"},
		{"failing eval compile error without report", `#!/bin/sh
if test -f "$2/evals/failing.txt"; then echo 'compile error' >&2; exit 2; fi
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "failing_eval"},
		{"failing eval with unrelated failure", `#!/bin/sh
if test -f "$2/evals/failing.txt"; then printf '<testsuite><testcase name="ALP-001"><failure/></testcase></testsuite>' > "$3"; exit 1; fi
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "failing_eval"},
		{"missing id with unrelated failure", `#!/bin/sh
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
if test -f "$2/evals/missing.txt"; then printf '<testsuite><testcase name="%s-001"><failure/></testcase></testsuite>' "$prefix" > "$3"; exit 1; fi
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "missing_id"},
		{"missing id without report", `#!/bin/sh
if test -f "$2/evals/missing.txt"; then echo 'hook failed' >&2; exit 2; fi
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "missing_id"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			f.put(".fireweed/eval.sh", tc.script)
			if tc.name == "missing baseline coverage" {
				logDoctorCanaryProbes(t, f)
			}
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(r, "C14", false, tc.canary, "eval") {
				t.Errorf("unrelated eval outcome accepted: %s", doctorText(r))
			}

 succeeded in 622ms:
package evals_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
)

func doctorProbeLog(t *testing.T, f doctorFixture, report checks.DoctorReport, role string, markers ...string) string {
	t.Helper()
	var detail string
	for _, check := range report.Checks {
		if check.Name == role {
			detail = check.Detail
			break
		}
	}
	if detail == "" {
		t.Errorf("executed %s probe has no check: %s", role, doctorText(report))
		return ""
	}
	at := strings.Index(detail, "log: ")
	if at < 0 {
		t.Errorf("%s detail lacks log path: %q", role, detail)
		return ""
	}
	fields := strings.Fields(detail[at+len("log: "):])
	if len(fields) == 0 {
		t.Errorf("%s detail has empty log path: %q", role, detail)
		return ""
	}
	path := fields[0]
	if !filepath.IsAbs(path) {
		t.Errorf("%s log path is not absolute: %q", role, path)
	}
	rel, err := filepath.Rel(filepath.Join(f.root, ".fireweed", "runs"), path)
	if err != nil {
		t.Errorf("%s log path %q: %v", role, path, err)
		return ""
	}
	parts := strings.Split(filepath.ToSlash(rel), "/")
	if len(parts) != 2 || parts[0] == "" || parts[0] == "." || parts[0] == ".." || parts[1] != "doctor-"+role+"-1.log" {
		t.Errorf("%s log path = %q, want .fireweed/runs/<run>/doctor-%s-1.log", role, path, role)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("read %s log %q: %v", role, path, err)
		return path
	}
	for _, marker := range markers {
		if !strings.Contains(string(data), marker) {
			t.Errorf("%s log lacks %q: %q", role, marker, data)
		}
	}
	return path
}

func doctorLogRun(t *testing.T, f doctorFixture, worker, audit string, report checks.DoctorReport) string {
	t.Helper()
	base := filepath.Join(f.root, ".fireweed", "runs")
	w, ew := filepath.Rel(base, worker)
	a, ea := filepath.Rel(base, audit)
	if ew != nil || ea != nil {
		t.Errorf("log paths outside run root: %q %q: %v %v", worker, audit, ew, ea)
		return ""
	}
	wRun, aRun := filepath.Dir(w), filepath.Dir(a)
	if wRun != aRun || wRun == "." || wRun == ".." {
		t.Errorf("worker and audit have different run identities: %q %q: %s", worker, audit, doctorText(report))
	}
	return wRun
}

func doctorCallerClean(t *testing.T, f doctorFixture, status, registrations string) {
	t.Helper()
	if got := f.git("status", "--porcelain", "--untracked-files=all"); got != status {
		t.Errorf("caller status changed: before %q after %q", status, got)
	}
	if got := worktrees(f); got != registrations {
		t.Errorf("scratch registration changed: before %q after %q", registrations, got)
	}
	assertScratchRemoved(t, f)
}

func Test_CHK_081_DoctorProbeLogsShareFreshRunAndPreserveStreams(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/agent.sh", "#!/bin/bash\nprintf 'worker stdout Ω without newline'; printf 'worker stderr β without newline' >&2; printf 'FIREWEED_DOCTOR_OK\\n'\n")
	f.put(".fireweed/audit.sh", "#!/bin/bash\nprintf 'audit stdout Ω without newline'; printf 'audit stderr β without newline' >&2; printf 'FIREWEED_DOCTOR_OK\\n'\n")
	f.commit()
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	var firstRun string
	for i := 0; i < 2; i++ {
		report := checks.Doctor(context.Background(), f.request(false))
		if doctorErrors(report) || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
			t.Errorf("successful probes changed outcome: %s", doctorText(report))
		}
		worker := doctorProbeLog(t, f, report, "worker", "worker stdout Ω without newlineFIREWEED_DOCTOR_OK\n", "worker stderr β without newline")
		audit := doctorProbeLog(t, f, report, "audit", "audit stdout Ω without newlineFIREWEED_DOCTOR_OK\n", "audit stderr β without newline")
		if worker != "" && audit != "" {
			run := doctorLogRun(t, f, worker, audit, report)
			if i == 1 && run == firstRun {
				t.Errorf("consecutive Doctor invocations reused run identity %q", run)
			}
			firstRun = run
		}
		doctorCallerClean(t, f, status, registrations)
	}
}

func Test_CHK_081_DoctorPersistsFailedAndTimedOutProbeOutput(t *testing.T) {
	failureStdout := "early worker stdout Ω\n"
	for i := 1; i <= 25; i++ {
		failureStdout += fmt.Sprintf("later-%02d\n", i)
	}
	for _, tc := range []struct {
		name, role, script, stdout, stderr string
		timeout                            time.Duration
	}{
		{"failure", "worker", "#!/bin/bash\nprintf 'early worker stdout Ω\\n'; for i in $(seq 1 25); do printf 'later-%02d\\n' \"$i\"; done; printf 'failed worker stderr β' >&2; exit 7\n", failureStdout, "failed worker stderr β", 2 * time.Second},
		{"timeout", "audit", "#!/bin/bash\nprintf 'partial audit stdout Ω'; printf 'partial audit stderr β' >&2; sleep 5\n", "partial audit stdout Ω", "partial audit stderr β", 150 * time.Millisecond},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			if tc.role == "worker" {
				f.put(".fireweed/agent.sh", tc.script)
			} else {
				f.put(".fireweed/audit.sh", tc.script)
			}
			f.commit()
			status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
			req := f.request(false)
			req.Timeout = tc.timeout
			start := time.Now()
			report := checks.Doctor(context.Background(), req)
			if time.Since(start) > 4*time.Second {
				t.Errorf("%s probe did not honor bounded request timeout", tc.name)
			}
			if !hasDoctorCheck(report, tc.role, false) || !hasDoctorDiagnostic(report, "C15", false, tc.role, tc.stderr) {
				t.Errorf("%s probe outcome/diagnostic lost: %s", tc.name, doctorText(report))
			}
			path := doctorProbeLog(t, f, report, tc.role, tc.stdout, tc.stderr)
			if tc.name == "failure" && path != "" {
				if !strings.Contains(string(mustRead(t, path)), "later-25") || strings.Contains(doctorText(report), "early worker stdout Ω") {
					t.Errorf("failed probe log or final-20-line diagnostic changed: %s", doctorText(report))
				}
			}
			doctorCallerClean(t, f, status, registrations)
		})
	}
}

func Test_CHK_081_DoctorPersistsCancelledProbeOutput(t *testing.T) {
	f := localDoctorFixture(t)
	ready := filepath.Join(t.TempDir(), "worker-ready")
	f.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/bash\nprintf 'cancel worker stdout Ω'; printf 'cancel worker stderr β' >&2; printf ready > %q; sleep 5\n", ready))
	f.commit()
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reports := make(chan checks.DoctorReport, 1)
	go func() { reports <- checks.Doctor(ctx, f.request(false)) }()
	deadline := time.After(8 * time.Second)
	for {
		if _, err := os.Stat(ready); err == nil {
			break
		} else if !os.IsNotExist(err) {
			t.Fatal(err)
		}
		select {
		case report := <-reports:
			t.Fatalf("worker exited before cancellation: %s", doctorText(report))
		case <-deadline:
			t.Fatal("worker never signalled readiness")
		case <-time.After(10 * time.Millisecond):
		}
	}
	cancel()
	select {
	case report := <-reports:
		if !hasDoctorCheck(report, "worker", false) || !hasDoctorDiagnostic(report, "C15", false, "worker") {
			t.Errorf("cancelled probe outcome lost: %s", doctorText(report))
		}
		doctorProbeLog(t, f, report, "worker", "cancel worker stdout Ω", "cancel worker stderr β")
	case <-time.After(4 * time.Second):
		t.Fatal("cancelled Doctor did not return promptly")
	}
	doctorCallerClean(t, f, status, registrations)
}
package evals_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
)

const doctorProbePrompt = "Reply with exactly FIREWEED_DOCTOR_OK"

func doctorRecordedPath(t *testing.T, log, label string) string {
	t.Helper()
	return doctorRecordedPaths(t, log, label, 1)[0]
}

func doctorRecordedPaths(t *testing.T, log, label string, count int) []string {
	t.Helper()
	var paths []string
	for _, line := range strings.Split(log, "\n") {
		if strings.HasPrefix(line, label+":") {
			paths = append(paths, strings.TrimPrefix(line, label+":"))
		}
	}
	if len(paths) != count {
		t.Fatalf("%s paths = %q; want %d absolute paths", label, paths, count)
	}
	for _, path := range paths {
		if !filepath.IsAbs(path) {
			t.Fatalf("%s path = %q; want an absolute path", label, path)
		}
	}
	return paths
}

func doctorPromptFileRemoved(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Errorf("prompt file %q remains after probe: %v", path, err)
	}
}

func Test_CHK_082_CHK_083_DoctorExpandsSelectedCommandsAndWrapperOnce(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) {
		config["agent"].(map[string]any)["command"] = []any{"bash", ".fireweed/agent.sh", "pre{prompt}post", "file={prompt_file}", "common={git_common_dir}", "literal {unknown}"}
		config["audit"].(map[string]any)["command"] = "bash .fireweed/audit.sh {prompt} {prompt_file} {git_common_dir}"
		config["agent"].(map[string]any)["wrap"] = []any{"bash", ".fireweed/prompt-wrap.sh", "{prompt_file}", "{prompt}", "{git_common_dir}"}
	})
	f.put(".fireweed/prompt-wrap.sh", fmt.Sprintf(`#!/bin/bash
printf 'wrap-file:%%s\nwrap-prompt:%%s\nwrap-common:%%s\n' "$1" "$2" "$3" >> %q
test -f "$1" && printf '%%s' "$2" | cmp -s - "$1" || exit 11
shift 3
exec "$@"
`, f.log))
	f.put(".fireweed/agent.sh", fmt.Sprintf(`#!/bin/bash
printf 'worker-prompt:%%s\nworker-file:%%s\nworker-common:%%s\nworker-literal:%%s\n' "$1" "${2#file=}" "${3#common=}" "$4" >> %q
test "$1" = "pre%spost" || exit 12
test "$2" != "file=" || exit 13
test -f "${2#file=}" && printf '%%s' %q | cmp -s - "${2#file=}" || exit 14
printf 'FIREWEED_DOCTOR_OK\n'
`, f.log, doctorProbePrompt, doctorProbePrompt))
	f.put(".fireweed/audit.sh", fmt.Sprintf(`#!/bin/bash
printf 'audit-prompt:%%s\naudit-file:%%s\naudit-common:%%s\n' "$1" "$2" "$3" >> %q
test -f "$2" && printf '%%s' "$1" | cmp -s - "$2" || exit 15
printf 'FIREWEED_DOCTOR_OK\n'
`, f.log))
	f.commit()
	// Relocate the Git directory so its absolute path contains a quote, a space,
	// and another placeholder. The expanded value must remain literal.
	commonDir := filepath.Join(t.TempDir(), "space 'quote' {prompt}", "repo.git")
	if err := os.MkdirAll(filepath.Dir(commonDir), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Rename(filepath.Join(f.root, ".git"), commonDir); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(f.root, ".git"), []byte("gitdir: "+commonDir+"\n"), 0644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("TMPDIR", filepath.Dir(commonDir))
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if doctorErrors(report) || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
		t.Errorf("prompt-file probes failed: %s", doctorText(report))
	}
	log := doctorLog(f)
	common, err := filepath.EvalSymlinks(commonDir)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"worker-prompt:pre" + doctorProbePrompt + "post", "audit-prompt:" + doctorProbePrompt,
		"worker-common:" + common, "audit-common:" + common,
		"worker-literal:literal {unknown}",
		"wrap-prompt:" + doctorProbePrompt, "wrap-common:" + common,
	} {
		if !strings.Contains(log, want+"\n") {
			t.Errorf("missing exact argument %q", want)
		}
	}
	workerFile := doctorRecordedPath(t, log, "worker-file")
	auditFile := doctorRecordedPath(t, log, "audit-file")
	if workerFile == auditFile {
		t.Errorf("worker and audit reused prompt file %q", workerFile)
	}
	for _, path := range []string{workerFile, auditFile} {
		if !strings.Contains(log, "wrap-file:"+path+"\n") {
			t.Errorf("wrapper did not receive selected probe file %q", path)
		}
		doctorPromptFileRemoved(t, path)
	}
	doctorProbeLog(t, f, report, "worker", "FIREWEED_DOCTOR_OK")
	doctorProbeLog(t, f, report, "audit", "FIREWEED_DOCTOR_OK")
	doctorCallerClean(t, f, status, registrations)
}

func Test_CHK_083_DoctorRemovesPromptFileOnProbeFailures(t *testing.T) {
	for _, mode := range []string{"launch", "nonzero", "timeout", "cancel"} {
		t.Run(mode, func(t *testing.T) {
			f := localDoctorFixture(t)
			ready := filepath.Join(t.TempDir(), "ready")
			f.editConfig(func(config map[string]any) {
				config["agent"].(map[string]any)["command"] = []any{"bash", ".fireweed/agent.sh", "{prompt_file}"}
				if mode == "launch" {
					config["agent"].(map[string]any)["command"] = []any{filepath.Join(f.root, "missing-agent"), "{prompt_file}"}
					config["agent"].(map[string]any)["wrap"] = []any{"bash", ".fireweed/prompt-wrap.sh", "{prompt_file}"}
				}
			})
			f.put(".fireweed/prompt-wrap.sh", fmt.Sprintf("#!/bin/bash\nprintf 'worker-file:%%s\\n' \"$1\" >> %q\ntest -f \"$1\" || exit 11\nshift\nexec \"$@\"\n", f.log))
			script := fmt.Sprintf("#!/bin/bash\nprintf 'worker-file:%%s\\nworker-content:%%s\\n' \"$1\" \"$(cat \"$1\")\" >> %q\ntest -f \"$1\" && printf '%%s' %q | cmp -s - \"$1\" || exit 18\nprintf ready > %q\n", f.log, doctorProbePrompt, ready)
			switch mode {
			case "nonzero":
				script += "printf 'failed worker stderr\\n' >&2\nexit 7\n"
			case "timeout", "cancel":
				script += "printf 'waiting worker stdout\\n'\nsleep 5\n"
			}
			f.put(".fireweed/agent.sh", script)
			f.commit()
			status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
			req := f.request(false)
			if mode == "timeout" {
				req.Timeout = 2 * time.Second
			}
			ctx := context.Background()
			var cancel context.CancelFunc
			if mode == "cancel" {
				ctx, cancel = context.WithCancel(ctx)
				defer cancel()
			}
			var report checks.DoctorReport
			if mode == "cancel" {
				reports := make(chan checks.DoctorReport, 1)
				go func() { reports <- checks.Doctor(ctx, req) }()
				deadline := time.After(8 * time.Second)
				for {
					if _, err := os.Stat(ready); err == nil {
						break
					} else if !os.IsNotExist(err) {
						t.Fatal(err)
					}
					select {
					case report = <-reports:
						t.Fatalf("worker exited before cancellation: %s", doctorText(report))
					case <-deadline:
						t.Fatal("worker never signalled readiness")
					case <-time.After(10 * time.Millisecond):
					}
				}
				livePath := doctorRecordedPath(t, doctorLog(f), "worker-file")
				if data, err := os.ReadFile(livePath); err != nil || string(data) != doctorProbePrompt {
					t.Errorf("prompt file unavailable before cancellation: %q, %v", data, err)
				}
				cancel()
				select {
				case report = <-reports:
				case <-time.After(4 * time.Second):
					t.Fatal("cancelled Doctor did not return")
				}
			} else {
				report = checks.Doctor(ctx, req)
			}
			if !hasDoctorCheck(report, "worker", false) || !hasDoctorDiagnostic(report, "C15", false, "worker") {
				t.Errorf("%s probe outcome changed: %s", mode, doctorText(report))
			}
			log := doctorLog(f)
			count := 1
			if mode == "launch" {
				// The shared wrapper records both the failed worker launch and
				// the subsequent audit probe; both files must be cleaned up.
				count = 2
			}
			paths := doctorRecordedPaths(t, log, "worker-file", count)
			if mode != "launch" && !strings.Contains(log, "worker-content:"+doctorProbePrompt+"\n") {
				t.Errorf("%s worker could not read exact prompt: %q", mode, log)
			}
			for _, path := range paths {
				doctorPromptFileRemoved(t, path)
			}
			doctorProbeLog(t, f, report, "worker")
			doctorCallerClean(t, f, status, registrations)
		})
	}
}

func Test_CHK_083_DoctorDoesNotCreatePromptFileWithoutSelectedPlaceholder(t *testing.T) {
	f := localDoctorFixture(t)
	expected := filepath.Join(t.TempDir(), "expected-prompt")
	if err := os.WriteFile(expected, []byte(doctorProbePrompt), 0644); err != nil {
		t.Fatal(err)
	}
	script := fmt.Sprintf(`#!/bin/bash
while IFS= read -r candidate; do
  if test -f "$candidate" && cmp -s "$candidate" %q; then
    printf 'unexpected prompt file: %%s\n' "$candidate" >&2
    exit 19
  fi
done < <(find "$TMPDIR" "$PWD" -type f -print)
printf 'FIREWEED_DOCTOR_OK\n'
`, expected)
	f.put(".fireweed/agent.sh", script)
	f.put(".fireweed/audit.sh", script)
	f.commit()
	temp := t.TempDir()
	t.Setenv("TMPDIR", temp)
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if doctorErrors(report) || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
		t.Errorf("unneeded prompt file appeared: %s", doctorText(report))
	}
	doctorProbeLog(t, f, report, "worker", "FIREWEED_DOCTOR_OK")
	doctorProbeLog(t, f, report, "audit", "FIREWEED_DOCTOR_OK")
	doctorCallerClean(t, f, status, registrations)
}

func Test_CHK_082_CHK_083_DoctorWrapperOnlyFileWithAuditFallback(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) {
		delete(config["audit"].(map[string]any), "command")
		config["agent"].(map[string]any)["command"] = []any{"bash", ".fireweed/agent.sh", "{prompt}"}
		config["agent"].(map[string]any)["wrap"] = []any{"bash", ".fireweed/prompt-wrap.sh", "{prompt_file}"}
	})
	f.put(".fireweed/prompt-wrap.sh", fmt.Sprintf(`#!/bin/bash
printf 'wrap-file:%%s\n' "$1" >> %q
test -f "$1" && printf '%%s' %q | cmp -s - "$1" || exit 11
shift
exec "$@"
`, f.log, doctorProbePrompt))
	f.put(".fireweed/agent.sh", fmt.Sprintf(`#!/bin/bash
printf 'fallback-prompt:%%s\n' "$1" >> %q
test "$1" = %q || exit 12
printf 'FIREWEED_DOCTOR_OK\n'
`, f.log, doctorProbePrompt))
	f.commit()
	status, registrations := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if doctorErrors(report) || !hasDoctorDiagnostic(report, "C15", true, "shared") || !hasDoctorCheck(report, "worker", true) || !hasDoctorCheck(report, "audit", true) {
		t.Errorf("wrapper-only fallback changed: %s", doctorText(report))
	}
	log := doctorLog(f)
	if strings.Count(log, "fallback-prompt:"+doctorProbePrompt+"\n") != 2 {
		t.Errorf("fallback did not receive fixed prompt twice")
	}
	var paths []string
	for _, line := range strings.Split(log, "\n") {
		if strings.HasPrefix(line, "wrap-file:") {
			paths = append(paths, strings.TrimPrefix(line, "wrap-file:"))
		}
	}
	if len(paths) != 2 || paths[0] == paths[1] {
		t.Errorf("wrapper prompt files = %q; want distinct worker/audit paths", paths)
	}
	for _, path := range paths {
		if !filepath.IsAbs(path) {
			t.Errorf("wrapper prompt path is relative: %q", path)
		}
		doctorPromptFileRemoved(t, path)
	}
	doctorProbeLog(t, f, report, "worker", "FIREWEED_DOCTOR_OK")
	doctorProbeLog(t, f, report, "audit", "FIREWEED_DOCTOR_OK")
	doctorCallerClean(t, f, status, registrations)
}

 succeeded in 634ms:
packages/checks/evals/doctor_agent_logs_test.go:15:func doctorProbeLog(t *testing.T, f doctorFixture, report checks.DoctorReport, role string, markers ...string) string {
packages/checks/evals/doctor_agent_logs_test.go:64:func doctorLogRun(t *testing.T, f doctorFixture, worker, audit string, report checks.DoctorReport) string {
packages/checks/evals/doctor_agent_logs_test.go:80:func doctorCallerClean(t *testing.T, f doctorFixture, status, registrations string) {
packages/checks/evals/doctor_agent_logs_test.go:91:func Test_CHK_081_DoctorProbeLogsShareFreshRunAndPreserveStreams(t *testing.T) {
packages/checks/evals/doctor_agent_logs_test.go:116:func Test_CHK_081_DoctorPersistsFailedAndTimedOutProbeOutput(t *testing.T) {
packages/checks/evals/doctor_agent_logs_test.go:158:func Test_CHK_081_DoctorPersistsCancelledProbeOutput(t *testing.T) {
packages/checks/evals/doctor_prompt_file_test.go:48:func Test_CHK_082_CHK_083_DoctorExpandsSelectedCommandsAndWrapperOnce(t *testing.T) {
packages/checks/evals/doctor_prompt_file_test.go:123:func Test_CHK_083_DoctorRemovesPromptFileOnProbeFailures(t *testing.T) {
packages/checks/evals/doctor_prompt_file_test.go:211:func Test_CHK_083_DoctorDoesNotCreatePromptFileWithoutSelectedPlaceholder(t *testing.T) {
packages/checks/evals/doctor_prompt_file_test.go:241:func Test_CHK_082_CHK_083_DoctorWrapperOnlyFileWithAuditFallback(t *testing.T) {
packages/checks/evals/doctor_test.go:32:func (f doctorFixture) put(name, body string) {
packages/checks/evals/doctor_test.go:43:func (f doctorFixture) git(args ...string) string {
packages/checks/evals/doctor_test.go:55:func (f doctorFixture) commit() {
packages/checks/evals/doctor_test.go:64:func (f doctorFixture) request(noAgent bool) checks.DoctorRequest {
packages/checks/evals/doctor_test.go:85:func (f doctorFixture) editConfig(change func(map[string]any)) {
packages/checks/evals/doctor_test.go:99:func newDoctorFixture(t *testing.T) doctorFixture {
packages/checks/evals/doctor_test.go:120:func localDoctorFixture(t *testing.T) doctorFixture {
packages/checks/evals/doctor_test.go:228:func doctorLog(f doctorFixture) string { b, _ := os.ReadFile(f.log); return string(b) }
packages/checks/evals/doctor_test.go:229:func worktrees(f doctorFixture) string { return f.git("worktree", "list", "--porcelain") }
packages/checks/evals/doctor_test.go:231:func logDoctorCanaryProbes(t *testing.T, f doctorFixture) {
packages/checks/evals/doctor_test.go:247:func assertScratchRemoved(t *testing.T, f doctorFixture) {
packages/checks/evals/doctor_test.go:289:func Test_CHK_049_StaticDiagnosticsStopProbes(t *testing.T) {
packages/checks/evals/doctor_test.go:313:func Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_CHK_060_CHK_061_CHK_067_RealGoRecipe(t *testing.T) {
packages/checks/evals/doctor_test.go:370:func Test_CHK_053_NestedModuleTemplatesRenderForBothCanaries(t *testing.T) {
packages/checks/evals/doctor_test.go:426:func Test_CHK_051_CHK_052_SetupFailureStopsAndCleans(t *testing.T) {
packages/checks/evals/doctor_test.go:444:func Test_CHK_051_CancellationCleansScratchAndStopsSetupGroup(t *testing.T) {
packages/checks/evals/doctor_test.go:494:func Test_CHK_054_MissingTemplatesAndCanaryMaps(t *testing.T) {
packages/checks/evals/doctor_test.go:548:func Test_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_ZeroHookAndRecovery(t *testing.T) {
packages/checks/evals/doctor_test.go:584:func Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific(t *testing.T) {
packages/checks/evals/doctor_test.go:631:func Test_CHK_055_BaselineAndRecoveryFailures(t *testing.T) {
packages/checks/evals/doctor_test.go:672:func Test_CHK_056_CHK_057_BoundaryViolationMustNameCanaryModule(t *testing.T) {
packages/checks/evals/doctor_test.go:686:func Test_CHK_057_InternalEvalRequiresInterfaceViolation(t *testing.T) {
packages/checks/evals/doctor_test.go:711:func Test_CHK_055_CHK_056_CHK_057_MissingRequiredBoundaryHook(t *testing.T) {
packages/checks/evals/doctor_test.go:737:func Test_CHK_055_MissingRequiredHookBesideOverride(t *testing.T) {
packages/checks/evals/doctor_test.go:758:func Test_CHK_055_ApplicableOverrideNeedsNoGlobalHook(t *testing.T) {
packages/checks/evals/doctor_test.go:770:func Test_CHK_060_CHK_062_ModuleOverridesAndBoundaryWarnings(t *testing.T) {
packages/checks/evals/doctor_test.go:790:func Test_CHK_060_OverridesMustDetectTheirCanaries(t *testing.T) {
packages/checks/evals/doctor_test.go:827:func Test_CHK_063_CHK_064_CHK_065_CHK_066_CHK_067_AgentProbes(t *testing.T) {
packages/checks/evals/doctor_test.go:948:func Test_CHK_065_UnsetTimeoutBoundsAuditProbe(t *testing.T) {
packages/checks/evals/doctor_test.go:970:func Test_CHK_068_CHK_069_ManagedGuidanceAndPlaybookBase(t *testing.T) {
packages/checks/evals/history_test.go:34:func (f historyFixture) git(args ...string) string {
packages/checks/evals/history_test.go:51:func (f historyFixture) put(path, body string) {
packages/checks/evals/history_test.go:62:func (f historyFixture) remove(path string) {
packages/checks/evals/history_test.go:69:func (f historyFixture) move(oldPath, newPath string) {
packages/checks/evals/history_test.go:77:func (f historyFixture) commit(subject string) string {
packages/checks/evals/history_test.go:84:func (f historyFixture) change(path, phase, set string) {
packages/checks/evals/history_test.go:93:func newHistoryFixture(t *testing.T) historyFixture {
packages/checks/evals/history_test.go:111:func (f historyFixture) report(revisionRange string) checks.HistoryReport {
packages/checks/evals/history_test.go:122:func expectHistoryFindings(t *testing.T, report checks.HistoryReport, want ...historyFinding) {
packages/checks/evals/history_test.go:142:func Test_CHK_070_HistoryExplicitRangeOrderAndFullOIDs(t *testing.T) {
packages/checks/evals/history_test.go:155:func Test_CHK_070_HistoryBranchingRangeTopologicalOrder(t *testing.T) {
packages/checks/evals/history_test.go:201:func Test_CHK_070_HistoryDefaultRangeUsesLocalConfiguredBase(t *testing.T) {
packages/checks/evals/history_test.go:229:func Test_CHK_071_CHK_072_HistoryPhaseContentsAndDeletionControl(t *testing.T) {
packages/checks/evals/history_test.go:244:func Test_CHK_072_CHK_073_CHK_077_HistoryPlaybookOverrideAndProjectedInterface(t *testing.T) {
packages/checks/evals/history_test.go:269:func Test_CHK_071_CHK_073_HistoryAppliedImplTouchesEvalsAndInterface(t *testing.T) {
packages/checks/evals/history_test.go:286:func Test_CHK_074_HistoryAuthorAppendOnlyModifyDeleteRename(t *testing.T) {
packages/checks/evals/history_test.go:313:func Test_CHK_074_HistoryMergeResolutionInMergeOnlyRange(t *testing.T) {
packages/checks/evals/history_test.go:334:func Test_CHK_075_HistoryArchiveModifyAndDeletionControl(t *testing.T) {
packages/checks/evals/history_test.go:347:func Test_CHK_076_HistoryMultipleSquashedChangesKeepH03H04(t *testing.T) {
packages/checks/evals/history_test.go:383:func Test_CHK_077_HistoryUsesRecordedBeforeAfterZones(t *testing.T) {
packages/checks/evals/history_test.go:398:func Test_CHK_077_HistoryChecksBothRecordedZoneMaps(t *testing.T) {
packages/checks/evals/history_test.go:410:func Test_CHK_078_HistoryRejectsUnreadableRangeAndRecordedDocuments(t *testing.T) {

exec
/bin/zsh -lc "sed -n '584,830p' packages/checks/evals/doctor_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc "sed -n '827,1020p' packages/checks/evals/doctor_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc "sed -n '1,440p' packages/checks/evals/history_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc "sed -n '370,584p' packages/checks/evals/doctor_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 612ms:
func Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific(t *testing.T) {
	for _, tc := range []struct{ name, script, canary string }{
		{"missing baseline coverage", `#!/bin/sh
printf '<testsuite/>' > "$3"
`, "baseline"},
		{"failing eval compile error without report", `#!/bin/sh
if test -f "$2/evals/failing.txt"; then echo 'compile error' >&2; exit 2; fi
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "failing_eval"},
		{"failing eval with unrelated failure", `#!/bin/sh
if test -f "$2/evals/failing.txt"; then printf '<testsuite><testcase name="ALP-001"><failure/></testcase></testsuite>' > "$3"; exit 1; fi
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "failing_eval"},
		{"missing id with unrelated failure", `#!/bin/sh
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
if test -f "$2/evals/missing.txt"; then printf '<testsuite><testcase name="%s-001"><failure/></testcase></testsuite>' "$prefix" > "$3"; exit 1; fi
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "missing_id"},
		{"missing id without report", `#!/bin/sh
if test -f "$2/evals/missing.txt"; then echo 'hook failed' >&2; exit 2; fi
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "missing_id"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			f.put(".fireweed/eval.sh", tc.script)
			if tc.name == "missing baseline coverage" {
				logDoctorCanaryProbes(t, f)
			}
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(r, "C14", false, tc.canary, "eval") {
				t.Errorf("unrelated eval outcome accepted: %s", doctorText(r))
			}
			if tc.name == "missing baseline coverage" && !strings.Contains(doctorLog(f), "probe:") {
				t.Errorf("failed baseline coverage hook did not run: %s", doctorLog(f))
			}
			if tc.name == "missing baseline coverage" && strings.Contains(doctorLog(f), "planted:") {
				t.Errorf("failed baseline coverage continued into canary probes: %s", doctorLog(f))
			}
		})
	}
}

func Test_CHK_055_BaselineAndRecoveryFailures(t *testing.T) {
	for _, tc := range []struct{ name, path, script, phase, hook string }{
		{"baseline eval failure", ".fireweed/eval.sh", "#!/bin/sh\necho baseline-eval-failed >&2\nexit 9\n", "baseline", "eval"},
		{"baseline boundary failure", ".fireweed/boundary.sh", "#!/bin/sh\nprintf 'alpha: baseline violation\\n'\n", "baseline", "boundar"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			f.put(tc.path, tc.script)
			logDoctorCanaryProbes(t, f)
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(r, "C14", false, tc.phase, tc.hook) {
				t.Errorf("failed baseline accepted: %s", doctorText(r))
			}
			if !strings.Contains(doctorLog(f), "probe:") {
				t.Errorf("failed baseline hook did not run: %s", doctorLog(f))
			}
			if strings.Contains(doctorLog(f), "planted:") {
				t.Errorf("failed baseline continued into canary probes: %s", doctorLog(f))
			}
		})
	}
	for _, tc := range []struct{ name, failure string }{
		{"recovery eval failure", "echo recovery-eval-failed >&2; exit 9"},
		{"recovery coverage failure", "printf '<testsuite/>' > \"$3\"; exit 0"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			original := string(mustRead(t, filepath.Join(f.root, ".fireweed/eval.sh")))
			injection := "if test -f \"$2/evals/failing.txt\"; then printf '%s' \"$2\" > .fireweed/eval-recovery-module; fi\n" +
				"if test -f .fireweed/eval-recovery-module && test \"$(cat .fireweed/eval-recovery-module)\" = \"$2\" && ! test -f \"$2/evals/failing.txt\"; then " + tc.failure + "; fi\n"
			f.put(".fireweed/eval.sh", strings.Replace(original, "#!/bin/sh\n", "#!/bin/sh\n"+injection, 1))
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(r, "C14", false, "recovery", "failing_eval", "eval") {
				t.Errorf("failed %s accepted: %s", tc.name, doctorText(r))
			}
		})
	}
}

func Test_CHK_056_CHK_057_BoundaryViolationMustNameCanaryModule(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/boundary.sh", `#!/bin/sh
if find packages \( -name undeclared.txt -o -name internal.txt \) | grep -q .; then printf 'alpha: unrelated violation\n'; fi
`)
	f.commit()
	r := checks.Doctor(context.Background(), f.request(true))
	for _, name := range []string{"undeclared", "internal_eval"} {
		if !hasDoctorDiagnostic(r, "C14", false, name, "boundar") {
			t.Errorf("unrelated module violation accepted for %s: %s", name, doctorText(r))
		}
	}
}

func Test_CHK_057_InternalEvalRequiresInterfaceViolation(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/boundary.sh", `#!/bin/sh
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test -f "$dir/undeclared.txt"; then printf '%s: undeclared import\n' "$name"; fi
  if test -f "$dir/evals/internal.txt"; then printf '%s: unrelated boundary violation\n' "$name"; fi
done
`)
	f.commit()
	r := checks.Doctor(context.Background(), f.request(true))
	passedProbe := false
	for _, check := range r.Checks {
		if check.Passed && (check.Name == "internal_eval" || strings.HasPrefix(check.Name, "internal_eval ")) {
			passedProbe = true
		}
	}
	if !hasDoctorDiagnostic(r, "C14", false, "internal_eval", "boundar") || passedProbe {
		t.Errorf("same-module non-interface violation accepted: %s", doctorText(r))
	}
	if !hasDoctorCheck(r, "undeclared", true) {
		t.Errorf("fixture did not detect the independent undeclared canary: %s", doctorText(r))
	}
}

func Test_CHK_055_CHK_056_CHK_057_MissingRequiredBoundaryHook(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	f.put("packages/alpha/internal/live.txt", "disposable implementation\n")
	logDoctorCanaryProbes(t, f)
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(report, "X02", false, "alpha", "boundar") {
		t.Errorf("required boundary hook was not diagnosed: %s", doctorText(report))
	}
	if !hasDoctorDiagnostic(report, "C14", false, "baseline", "boundar") {
		t.Errorf("missing boundary hook was accepted as a clean baseline: %s", doctorText(report))
	}
	if strings.Contains(doctorLog(f), "planted:") {
		t.Errorf("missing baseline boundary hook continued into canary probes: %s", doctorLog(f))
	}
	for _, canary := range []string{"undeclared", "internal_eval"} {
		if hasDoctorCheck(report, canary, true) {
			t.Errorf("missing boundary hook counted as %s detection: %s", canary, doctorText(report))
		}
		if !hasDoctorDiagnostic(report, "C14", false, canary, "boundar") {
			t.Errorf("missing boundary hook lacked named %s canary failure: %s", canary, doctorText(report))
		}
	}
}

func Test_CHK_055_MissingRequiredHookBesideOverride(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	f.put(".fireweed/override-boundary.sh", fmt.Sprintf("#!/bin/sh\nprintf 'override-boundaries:%%s\\n' \"$PWD\" >> %q\nexec sh .fireweed/boundary.sh\n", f.log))
	f.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(f.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nhooks:\n  boundaries: sh .fireweed/override-boundary.sh", 1))
	f.put("packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: []\n---\n# Beta\n\n## Specs\n\n### BET-001\nThe fixture MUST pass.\n")
	f.put("packages/beta/internal/live.txt", "disposable implementation\n")
	logDoctorCanaryProbes(t, f)
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(report, "X02", false, "beta", "boundar") {
		t.Errorf("module without its own hook escaped required boundary enforcement: %s", doctorText(report))
	}
	if strings.Contains(doctorLog(f), "planted:") {
		t.Errorf("failed baseline beside override continued into canary probes: %s", doctorLog(f))
	}
	if !strings.Contains(doctorLog(f), "override-boundaries:") {
		t.Errorf("applicable override was not run: %s", doctorLog(f))
	}
}

func Test_CHK_055_ApplicableOverrideNeedsNoGlobalHook(t *testing.T) {
	f := localDoctorFixture(t)
	f.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	f.put(".fireweed/override-boundary.sh", "#!/bin/sh\nexec sh .fireweed/boundary.sh\n")
	f.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(f.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nhooks:\n  boundaries: sh .fireweed/override-boundary.sh", 1))
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) || !hasDoctorCheck(report, "undeclared", true) || !hasDoctorCheck(report, "internal_eval", true) {
		t.Errorf("applicable override was rejected without a global hook: %s", doctorText(report))
	}
}

func Test_CHK_060_CHK_062_ModuleOverridesAndBoundaryWarnings(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/override-eval.sh", fmt.Sprintf("#!/bin/sh\nprintf 'override-eval:%%s\\n' \"$PWD\" >> %q\nexec sh .fireweed/eval.sh \"$@\"\n", f.log))
	f.put(".fireweed/override-boundary.sh", fmt.Sprintf("#!/bin/sh\nprintf 'override-boundaries:%%s:%%s\\n' \"$PWD\" \"$1\" >> %q\nexec sh .fireweed/boundary.sh\n", f.log))
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\nhooks:\n  eval: sh .fireweed/override-eval.sh {module} {module_dir} {report}\n  boundaries: sh .fireweed/override-boundary.sh {module}\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nThe fixture MUST pass.\n")
	f.commit()
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) || strings.Count(doctorLog(f), "override-eval:") < 2 || strings.Count(doctorLog(f), "override-boundaries:") < 2 {
		t.Errorf("override probes absent: %s", doctorLog(f))
	}
	g := localDoctorFixture(t)
	g.editConfig(func(config map[string]any) { delete(config["hooks"].(map[string]any), "boundaries") })
	g.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(g.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nboundaries_opt_out: legacy", 1))
	g.commit()
	r := checks.Doctor(context.Background(), g.request(true))
	if doctorErrors(r) || !hasDoctorDiagnostic(r, "X03", true, "legacy") || !hasDoctorDiagnostic(r, "X04", true, "no boundary enforcement required") {
		t.Errorf("optional boundaries warning lost: %s", doctorText(r))
	}
}

func Test_CHK_060_OverridesMustDetectTheirCanaries(t *testing.T) {
	for _, tc := range []struct {
		name, hook, command, script, canary string
	}{
		{"eval", "eval", "sh .fireweed/override-eval.sh {module} {module_dir} {report}", `#!/bin/sh
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "failing_eval"},
		{"boundary", "boundaries", "sh .fireweed/override-boundary.sh {module}", "#!/bin/sh\nexit 0\n", "internal_eval"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			path := ".fireweed/override-" + tc.name + ".sh"
			f.put(path, tc.script)
			f.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(f.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nhooks:\n  "+tc.hook+": "+tc.command, 1))
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			diagnosed := false
			for _, d := range r.Diagnostics {
				message := strings.ToLower(d.Message)
				boundaryHook := tc.name == "boundary" && !strings.Contains(message, "global") &&
					(strings.Contains(message, "alpha boundary hook") ||
						strings.Contains(message, "alpha boundaries hook") ||
						strings.Contains(message, "alpha boundaries override boundary hook") ||
						strings.Contains(message, "alpha boundaries override boundaries hook"))
				if d.Code == "C14" && !d.Warning && strings.Contains(message, tc.canary) &&
					(tc.name == "eval" && strings.Contains(message, "alpha "+tc.hook+" override") || boundaryHook) {
					diagnosed = true
				}
			}
			if !diagnosed {
				t.Errorf("broken %s override missed its %s canary without C14: %s", tc.name, tc.canary, doctorText(r))
			}
		})
	}
}

func Test_CHK_063_CHK_064_CHK_065_CHK_066_CHK_067_AgentProbes(t *testing.T) {
	f := localDoctorFixture(t)
	f.commit()
	for _, d := range project.Check(f.root) {

 succeeded in 612ms:
func Test_CHK_053_NestedModuleTemplatesRenderForBothCanaries(t *testing.T) {
	f := localDoctorFixture(t)
	f.put("go.mod", "module example.test/doctorfixture\n\ngo 1.23\n")
	fields := "{{name}}|{{prefix}}|{{package}}|{{go_module}}|{{module_dir}}|{{other_dir}}\n"
	f.put(".fireweed/templates/module/evals/nested.txt.tmpl", "eval|"+fields)
	f.put(".fireweed/templates/module/internal/nested.txt", "internal|"+fields)
	canaries := string(mustRead(t, filepath.Join(f.root, ".fireweed/canaries/canaries.yaml")))
	f.put(".fireweed/canaries/canaries.yaml", strings.Replace(canaries, "undeclared:\n", "  evals/baseline.txt: \"baseline|"+strings.TrimSpace(fields)+"\"\nundeclared:\n", 1))
	eval := string(mustRead(t, filepath.Join(f.root, ".fireweed/eval.sh")))
	check := fmt.Sprintf(`case "$1" in
  doctor*)
    prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
    other=
    for spec in packages/*/SPEC.md; do
      dir=${spec%%/SPEC.md}
      case "$dir" in packages/alpha|"$2") ;; *) other=$dir ;; esac
    done
    expected="$1|$prefix|$1|example.test/doctorfixture|$2|$other"
    test -n "$other" &&
      test "$(cat "$2/evals/nested.txt")" = "eval|$expected" &&
      test "$(cat "$2/internal/nested.txt")" = "internal|$expected" &&
      test "$(cat "$2/evals/baseline.txt")" = "baseline|$expected" &&
      test ! -e "$2/evals/nested.txt.tmpl" || exit 17
    printf 'rendered:%%s\n' "$1" >> %q
    ;;
esac
`, f.log)
	f.put(".fireweed/eval.sh", strings.Replace(eval, "#!/bin/sh\n", "#!/bin/sh\n"+check, 1))
	f.commit()
	before := worktrees(f)
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) {
		t.Errorf("nested template probes failed: %s", doctorText(report))
	}
	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id", "recovery"} {
		if !hasDoctorCheck(report, name, true) {
			t.Errorf("missing passing %s probe: %s", name, doctorText(report))
		}
	}
	modules := map[string]bool{}
	for _, line := range strings.Split(doctorLog(f), "\n") {
		if strings.HasPrefix(line, "rendered:") {
			modules[strings.TrimPrefix(line, "rendered:")] = true
		}
	}
	if len(modules) != 2 {
		t.Errorf("expected both throwaway modules to pass nested template checks, got %v: %s", modules, doctorText(report))
	}
	if got := worktrees(f); got != before {
		t.Errorf("scratch registration leaked: before %q after %q", before, got)
	}
	if strings.Contains(doctorLog(f), "setup:") {
		assertScratchRemoved(t, f)
	}
}

func Test_CHK_051_CHK_052_SetupFailureStopsAndCleans(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/setup.sh", fmt.Sprintf("#!/bin/sh\nprintf 'setup:%%s\\n' \"$PWD\" >> %q\necho failed setup >&2\nexit 17\n", f.log))
	f.commit()
	before := worktrees(f)
	report := checks.Doctor(context.Background(), f.request(false))
	if !hasDoctorDiagnostic(report, "C14", false, "setup") {
		t.Errorf("setup failure absent: %s", doctorText(report))
	}
	if strings.Contains(doctorLog(f), "eval:") || strings.Contains(doctorLog(f), "boundaries:") || strings.Contains(doctorLog(f), "worker:") {
		t.Errorf("setup failure ran later probes: %s", doctorLog(f))
	}
	if got := worktrees(f); got != before {
		t.Errorf("scratch worktree leaked: %q", got)
	}
	assertScratchRemoved(t, f)
}

func Test_CHK_051_CancellationCleansScratchAndStopsSetupGroup(t *testing.T) {
	f := localDoctorFixture(t)
	ready := filepath.Join(t.TempDir(), "ready")
	childFinished := filepath.Join(t.TempDir(), "child-finished")
	f.put(".fireweed/setup.sh", fmt.Sprintf(`#!/bin/sh
printf 'setup:%%s\n' "$PWD" >> %q
sh -c 'sleep 2; printf finished > "$1"' sh %q &
printf ready > %q
wait
`, f.log, childFinished, ready))
	f.commit()
	before := worktrees(f)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reports := make(chan checks.DoctorReport, 1)
	go func() { reports <- checks.Doctor(ctx, f.request(true)) }()
	deadline := time.After(8 * time.Second)
	for {
		if _, err := os.Stat(ready); err == nil {
			break
		} else if !os.IsNotExist(err) {
			t.Fatal(err)
		}
		select {
		case report := <-reports:
			t.Fatalf("setup exited before cancellation: %s", doctorText(report))
		case <-deadline:
			t.Fatal("setup never signalled readiness")
		case <-time.After(10 * time.Millisecond):
		}
	}
	cancel()
	select {
	case report := <-reports:
		if !hasDoctorDiagnostic(report, "C14", false, "setup") {
			t.Errorf("cancelled setup lacked C14: %s", doctorText(report))
		}
	case <-time.After(5 * time.Second):
		t.Fatal("cancelled doctor did not return promptly")
	}
	if got := worktrees(f); got != before {
		t.Errorf("cancelled doctor retained Git registration: before %q after %q", before, got)
	}
	assertScratchRemoved(t, f)
	time.Sleep(2200 * time.Millisecond)
	if _, err := os.Stat(childFinished); !os.IsNotExist(err) {
		t.Errorf("cancelled setup descendant continued: %v", err)
	}
}

func Test_CHK_054_MissingTemplatesAndCanaryMaps(t *testing.T) {
	for _, tc := range []struct{ name, path, body, term, input, affected string }{
		{"missing template", ".fireweed/templates/module/SPEC.md", "", "template", "SPEC.md", "undeclared"},
		{"malformed template", ".fireweed/templates/module/SPEC.md", "---\nname: [broken\nid_prefix: {{prefix}}\n---\n", "template", "SPEC.md", "undeclared"},
		{"missing baseline", ".fireweed/canaries/canaries.yaml", "version: 1\nbaseline: {}\n", "canary", "baseline", "undeclared"},
		{"missing internal_eval", ".fireweed/canaries/canaries.yaml", "missing_id:\n  evals/missing.txt: missing\n", "internal_eval", "internal_eval", "internal_eval"},
		{"wrong version", ".fireweed/canaries/canaries.yaml", "version: 2\nbaseline: {}\n", "canary", "canaries.yaml", "undeclared"},
		{"malformed canary", ".fireweed/canaries/canaries.yaml", "version: [bad\n", "canary", "canaries.yaml", "undeclared"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			if tc.name == "missing baseline" {
				body := string(mustRead(t, filepath.Join(f.root, tc.path)))
				start, end := strings.Index(body, "baseline:\n"), strings.Index(body, "undeclared:\n")
				if start < 0 || end <= start {
					t.Fatal("fixture cannot isolate baseline map")
				}
				tc.body = body[:start] + "baseline: {}\n" + body[end:]
			}
			if tc.name == "missing internal_eval" {
				body := string(mustRead(t, filepath.Join(f.root, tc.path)))
				tc.body = strings.Replace(body, "internal_eval:\n  internal/private.txt: \"private\"\n  evals/internal.txt: \"{{module_dir}}\"\n", "", 1)
				if tc.body == body {
					t.Fatal("fixture did not remove internal_eval map")
				}
			}
			if tc.body == "" {
				if err := os.Remove(filepath.Join(f.root, tc.path)); err != nil {
					t.Fatal(err)
				}
			} else {
				f.put(tc.path, tc.body)
			}
			f.commit()
			report := checks.Doctor(context.Background(), f.request(true))
			if !hasDoctorDiagnostic(report, "C14", false, tc.term) {
				t.Errorf("missing C14: %s", doctorText(report))
			}
			identified := false
			for _, d := range report.Diagnostics {
				if d.Code == "C14" && !d.Warning && strings.Contains(strings.ToLower(d.Message+" "+d.Path), strings.ToLower(tc.input)) {
					identified = true
				}
			}
			if !identified {
				t.Errorf("C14 did not identify %s: %s", tc.input, doctorText(report))
			}
			if hasDoctorCheck(report, tc.affected, true) {
				t.Errorf("invalid recipe passed canary: %s", doctorText(report))
			}
		})
	}
}

func Test_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_ZeroHookAndRecovery(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/boundary.sh", "#!/bin/sh\nexit 0\n")
	f.put(".fireweed/eval.sh", `#!/bin/sh
printf '<testsuite><testcase name="ALP-001"/></testsuite>' > "$3"
`)
	f.commit()
	before := worktrees(f)
	report := checks.Doctor(context.Background(), f.request(true))
	for _, tc := range []struct{ name, hook string }{{"undeclared", "boundar"}, {"internal_eval", "boundar"}, {"failing_eval", "eval"}, {"missing_id", "eval"}} {
		if !hasDoctorDiagnostic(report, "C14", false, tc.name, tc.hook) {
			t.Errorf("missed %s canary not diagnosed with hook: %s", tc.name, doctorText(report))
		}
	}
	if !doctorErrors(report) {
		t.Errorf("zero hooks accepted: %s", doctorText(report))
	}
	if got := worktrees(f); got != before {
		t.Errorf("failed canary leaked worktree: %s", got)
	}

	g := localDoctorFixture(t)
	g.put(".fireweed/boundary.sh", fmt.Sprintf(`#!/bin/sh
find packages -name SPEC.md | while IFS= read -r spec; do
  dir=${spec%%/SPEC.md}; name=$(sed -n 's/^name: //p' "$spec" | head -1)
  if test -f "$dir/undeclared.txt"; then printf '%%s' "$name" > .fireweed/recovery-seen; printf '%%s: undeclared import\n' "$name"; fi
done
if test -f .fireweed/recovery-seen && ! find packages -name undeclared.txt | grep -q .; then printf '%%s: lingering violation\n' "$(cat .fireweed/recovery-seen)"; fi
`))
	g.commit()
	r := checks.Doctor(context.Background(), g.request(true))
	if !hasDoctorDiagnostic(r, "C14", false, "recovery", "undeclared", "boundar") {
		t.Errorf("missing failed recovery diagnostic: %s", doctorText(r))
	}
}

func Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific(t *testing.T) {

 succeeded in 614ms:
func Test_CHK_063_CHK_064_CHK_065_CHK_066_CHK_067_AgentProbes(t *testing.T) {
	f := localDoctorFixture(t)
	f.commit()
	for _, d := range project.Check(f.root) {
		if !d.Warning {
			t.Fatalf("invalid agent fixture: %+v", d)
		}
	}
	r := checks.Doctor(context.Background(), f.request(false))
	log := doctorLog(f)
	if doctorErrors(r) || strings.Count(log, "worker:") != 1 || strings.Count(log, "audit:") != 1 || strings.Count(log, "wrap:") != 2 || !strings.Contains(log, "FIREWEED_DOCTOR_OK") {
		t.Errorf("worker/audit/wrap probes: %s; log=%s", doctorText(r), log)
	}
	workerPrompt, auditPrompt := false, false
	for _, line := range strings.Split(log, "\n") {
		if fields := strings.SplitN(line, ":", 3); len(fields) == 3 && strings.Contains(fields[2], "FIREWEED_DOCTOR_OK") {
			switch fields[0] {
			case "worker":
				workerPrompt = true
			case "audit":
				auditPrompt = true
			}
		}
	}
	if !workerPrompt {
		t.Errorf("worker did not receive its fixed-reply prompt: %s", log)
	}
	if !auditPrompt {
		t.Errorf("audit did not receive its fixed-reply prompt: %s", log)
	}
	for _, part := range []string{"/fireweed/worktrees/", "worker:", "audit:"} {
		if !strings.Contains(log, part) {
			t.Errorf("scratch/role marker %s absent: %s", part, log)
		}
	}

	g := localDoctorFixture(t)
	g.editConfig(func(config map[string]any) {
		delete(config["audit"].(map[string]any), "command")
		config["agent"].(map[string]any)["command"] = "bash .fireweed/agent.sh {prompt} {git_common_dir}"
	})
	g.commit()
	rr := checks.Doctor(context.Background(), g.request(false))
	if !hasDoctorDiagnostic(rr, "C15", true, "shared") || strings.Count(doctorLog(g), "worker:") != 2 {
		t.Errorf("audit fallback: %s; log=%s", doctorText(rr), doctorLog(g))
	}

	h := localDoctorFixture(t)
	h.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/bash\nfor i in $(seq 1 25); do printf 'line-%%02d\\n' \"$i\"; done\necho RATE LIMIT reached\nexit 9\n"))
	h.commit()
	bad := checks.Doctor(context.Background(), h.request(false))
	if !hasDoctorDiagnostic(bad, "C15", false, "usage limit", "rate limit", "line-25") {
		t.Errorf("usage-limit tail missing: %s", doctorText(bad))
	}
	if strings.Contains(doctorText(bad), "line-01") {
		t.Errorf("agent output was not trimmed to final 20 lines: %s", doctorText(bad))
	}
	outOfTail := localDoctorFixture(t)
	outOfTail.put(".fireweed/agent.sh", "#!/bin/bash\necho RATE LIMIT early\nfor i in $(seq 1 25); do printf 'later-%02d\\n' \"$i\"; done\nexit 9\n")
	outOfTail.commit()
	ordinary := checks.Doctor(context.Background(), outOfTail.request(false))
	if !hasDoctorDiagnostic(ordinary, "C15", false, "agent.sh", "later-25") {
		t.Errorf("ordinary out-of-tail failure absent: %s", doctorText(ordinary))
	}
	for _, d := range ordinary.Diagnostics {
		if d.Code == "C15" && !d.Warning && strings.Contains(strings.ToLower(d.Message), "usage limit") {
			t.Errorf("out-of-tail pattern classified as usage limit: %s", doctorText(ordinary))
		}
	}
	for _, tc := range []struct{ name, script string }{{"wrong reply", "#!/bin/bash\necho WRONG_REPLY\n"}, {"nonzero with reply", "#!/bin/bash\necho FIREWEED_DOCTOR_OK\nexit 7\n"}, {"ordinary failure", "#!/bin/bash\nfor i in $(seq 1 25); do printf 'ordinary-%02d\\n' \"$i\"; done\nexit 8\n"}} {
		t.Run(tc.name, func(t *testing.T) {
			q := localDoctorFixture(t)
			q.put(".fireweed/agent.sh", tc.script)
			q.commit()
			r := checks.Doctor(context.Background(), q.request(false))
			if !hasDoctorDiagnostic(r, "C15", false, "agent.sh") {
				t.Errorf("agent failure accepted: %s", doctorText(r))
			}
			if tc.name == "ordinary failure" && (!hasDoctorDiagnostic(r, "C15", false, "ordinary-25") || strings.Contains(doctorText(r), "ordinary-01")) {
				t.Errorf("ordinary failure tail incorrect: %s", doctorText(r))
			}
		})
	}
	wrongAudit := localDoctorFixture(t)
	wrongAudit.put(".fireweed/audit.sh", "#!/bin/bash\necho WRONG_REPLY\n")
	wrongAudit.commit()
	auditFailure := checks.Doctor(context.Background(), wrongAudit.request(false))
	if !hasDoctorDiagnostic(auditFailure, "C15", false, "audit.sh") || hasDoctorCheck(auditFailure, "audit", true) {
		t.Errorf("wrong audit reply accepted: %s", doctorText(auditFailure))
	}

	i := localDoctorFixture(t)
	i.put(".fireweed/audit.sh", "#!/bin/bash\nsleep 2\n")
	i.commit()
	req := i.request(false)
	req.Timeout = 100 * time.Millisecond
	start := time.Now()
	timed := checks.Doctor(context.Background(), req)
	if !hasDoctorDiagnostic(timed, "C15", false, "audit") || time.Since(start) > 4*time.Second {
		t.Errorf("audit timeout not bounded: %s; elapsed=%s", doctorText(timed), time.Since(start))
	}

	defaulted := localDoctorFixture(t)
	defaulted.put(".fireweed/agent.sh", "#!/bin/bash\nsleep 0.2\nprintf 'FIREWEED_DOCTOR_OK\\n'\n")
	defaulted.put(".fireweed/audit.sh", "#!/bin/bash\nsleep 0.2\nprintf 'FIREWEED_DOCTOR_OK\\n'\n")
	defaulted.commit()
	defaultRequest := defaulted.request(false)
	defaultRequest.Timeout = 0
	defaultReport := checks.Doctor(context.Background(), defaultRequest)
	if doctorErrors(defaultReport) || !hasDoctorCheck(defaultReport, "worker", true) || !hasDoctorCheck(defaultReport, "audit", true) {
		t.Errorf("unset timeout did not give both agent probes their default allowance: %s", doctorText(defaultReport))
	}

	j := localDoctorFixture(t)
	j.commit()
	skip := checks.Doctor(context.Background(), j.request(true))
	if strings.Contains(doctorLog(j), "worker:") || strings.Contains(doctorLog(j), "audit:") || !hasDoctorCheck(skip, "undeclared", true) || !strings.Contains(doctorLog(j), j.bin) {
		t.Errorf("NoAgent skipped hooks or ran agent: %s; log=%s", doctorText(skip), doctorLog(j))
	}
}

func Test_CHK_065_UnsetTimeoutBoundsAuditProbe(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/audit.sh", fmt.Sprintf("#!/bin/bash\nprintf 'audit:%%s:%%s\n' \"$PWD\" \"$1\" >> %q\nprintf 'audit-timeout-started\n'\nsleep 210\nprintf 'FIREWEED_DOCTOR_OK\n'\n", f.log))
	f.commit()
	req := f.request(false)
	req.Timeout = 0

	start := time.Now()
	report := checks.Doctor(context.Background(), req)
	elapsed := time.Since(start)
	log := doctorLog(f)
	if !hasDoctorCheck(report, "worker", true) || strings.Count(log, "worker:") != 1 || strings.Count(log, "audit:") != 1 {
		t.Errorf("agent probes did not run once with immediate worker reply: %s; log=%s", doctorText(report), log)
	}
	if !hasDoctorCheck(report, "audit", false) || !hasDoctorDiagnostic(report, "C15", false, "audit.sh", "audit-timeout-started") {
		t.Errorf("long audit probe lacked a C15 timeout with its command and output: %s", doctorText(report))
	}
	if elapsed < 165*time.Second || elapsed > 200*time.Second {
		t.Errorf("unset audit timeout took %s, want a bound near three minutes", elapsed)
	}
}

func Test_CHK_068_CHK_069_ManagedGuidanceAndPlaybookBase(t *testing.T) {
	f := localDoctorFixture(t)
	text, err := assets.Playbook("evals")
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256([]byte(text))
	base := hex.EncodeToString(sum[:])
	f.put(".fireweed/playbooks/evals.md", "---\nembedded_base: "+base+"\n---\n# Local override\n")
	f.commit()
	r := checks.Doctor(context.Background(), f.request(true))
	if hasDoctorDiagnostic(r, "C16", true) {
		t.Errorf("current guidance/base warned: %s", doctorText(r))
	}
	f.put("AGENTS.md", "Before.\n"+doctorGuidance+"After.\n")
	surrounded := checks.Doctor(context.Background(), f.request(true))
	if hasDoctorDiagnostic(surrounded, "C16", true) {
		t.Errorf("current managed guidance with surrounding text warned: %s", doctorText(surrounded))
	}
	f.put("AGENTS.md", "Before.\n<!-- fireweed:start -->\nold guidance\n<!-- fireweed:end -->\nAfter.\n")
	f.put(".fireweed/playbooks/evals.md", "---\nembedded_base: stale\n---\n# Local override\n")
	stale := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(stale, "C16", true, "AGENTS.md") || !hasDoctorDiagnostic(stale, "C16", true, "evals") {
		t.Errorf("stale guidance/base not warned: %s", doctorText(stale))
	}
	if err := os.Remove(filepath.Join(f.root, "AGENTS.md")); err != nil {
		t.Fatal(err)
	}
	f.put(".fireweed/playbooks/evals.md", "# No embedded base\n")
	missing := checks.Doctor(context.Background(), f.request(true))
	if !hasDoctorDiagnostic(missing, "C16", true, "AGENTS.md") || !hasDoctorDiagnostic(missing, "C16", true, "evals") {
		t.Errorf("missing guidance/base not warned: %s", doctorText(missing))
	}
}

func mustRead(t *testing.T, path string) []byte {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return b
}
func doctorBinary(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("caller")
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(file), "..", "..", ".."))
	bin := filepath.Join(t.TempDir(), "fireweed-new")

 succeeded in 615ms:
package evals_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

const historyConfig = `version: 1
base: main
modules: ["packages/*"]
protected: [locked/**]
module_defaults:
  evals: [evals/**]
  interface: [api.go]
  append_only: [ledger/**]
`

type historyFixture struct {
	t    *testing.T
	root string
	base string
}

func (f historyFixture) git(args ...string) string {
	f.t.Helper()
	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Args: append([]string{"git", "-C", f.root}, args...),
		Dir:  f.root,
		Env: map[string]string{
			"GIT_CONFIG_GLOBAL": "/dev/null", "GIT_CONFIG_NOSYSTEM": "1",
			"GIT_AUTHOR_NAME": "History Fixture", "GIT_AUTHOR_EMAIL": "fixture@example.invalid",
			"GIT_COMMITTER_NAME": "History Fixture", "GIT_COMMITTER_EMAIL": "fixture@example.invalid",
		},
	})
	if err != nil {
		f.t.Fatalf("git %v: %v: %s", args, err, result.Stderr)
	}
	return strings.TrimSpace(result.Stdout)
}

func (f historyFixture) put(path, body string) {
	f.t.Helper()
	full := filepath.Join(f.root, filepath.FromSlash(path))
	if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
		f.t.Fatal(err)
	}
	if err := os.WriteFile(full, []byte(body), 0644); err != nil {
		f.t.Fatal(err)
	}
}

func (f historyFixture) remove(path string) {
	f.t.Helper()
	if err := os.Remove(filepath.Join(f.root, filepath.FromSlash(path))); err != nil {
		f.t.Fatal(err)
	}
}

func (f historyFixture) move(oldPath, newPath string) {
	f.t.Helper()
	if err := os.MkdirAll(filepath.Dir(filepath.Join(f.root, filepath.FromSlash(newPath))), 0755); err != nil {
		f.t.Fatal(err)
	}
	f.git("mv", oldPath, newPath)
}

func (f historyFixture) commit(subject string) string {
	f.t.Helper()
	f.git("add", "-A")
	f.git("commit", "-qm", subject)
	return f.git("rev-parse", "HEAD")
}

func (f historyFixture) change(path, phase, set string) {
	f.t.Helper()
	body := fmt.Sprintf("---\nmodule: alpha\nphase: %s\n%s---\n\n## Motivation\nHistory fixture.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n\n## Notes\n", phase, set)
	if _, err := contracts.ParseDocument(body); err != nil {
		f.t.Fatalf("fixture change: %v", err)
	}
	f.put(path, body)
}

func newHistoryFixture(t *testing.T) historyFixture {
	t.Helper()
	f := historyFixture{t: t, root: t.TempDir()}
	f.git("init", "-q", "-b", "main")
	f.put("fireweed.yaml", historyConfig)
	f.put("packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha behavior.\n")
	f.put("packages/alpha/api.go", "package alpha\n")
	f.put("packages/alpha/evals/existing_test.go", "package evals\n")
	f.put("packages/alpha/ledger/entry.txt", "initial\n")
	f.put("packages/alpha/newledger/entry.txt", "initial\n")
	f.put("locked/config.txt", "initial\n")
	f.base = f.commit("base")
	if _, err := project.LoadAt(f.root, f.base); err != nil {
		t.Fatalf("recorded base project: %v", err)
	}
	return f
}

func (f historyFixture) report(revisionRange string) checks.HistoryReport {
	f.t.Helper()
	report, err := checks.History(f.root, revisionRange)
	if err != nil {
		f.t.Fatalf("History(%q): %v", revisionRange, err)
	}
	return report
}

type historyFinding struct{ code, path, commit string }

func expectHistoryFindings(t *testing.T, report checks.HistoryReport, want ...historyFinding) {
	t.Helper()
	if len(report.Diagnostics) != len(want) {
		t.Fatalf("diagnostics = %#v; want %d findings", report.Diagnostics, len(want))
	}
	used := make([]bool, len(report.Diagnostics))
	for _, expected := range want {
		found := false
		for i, d := range report.Diagnostics {
			if !used[i] && d.Code == expected.code && d.Path == expected.path && strings.Contains(d.Message, expected.commit) {
				used[i], found = true, true
				break
			}
		}
		if !found {
			t.Errorf("missing %s on %s at %s in %#v", expected.code, expected.path, expected.commit, report.Diagnostics)
		}
	}
}

func Test_CHK_070_HistoryExplicitRangeOrderAndFullOIDs(t *testing.T) {
	f := newHistoryFixture(t)
	f.put("notes/one.txt", "one")
	one := f.commit("first")
	f.put("notes/two.txt", "two")
	two := f.commit("second")
	got := f.report(f.base + ".." + two)
	if !reflect.DeepEqual(got.Commits, []string{one, two}) {
		t.Errorf("commits = %#v; want [%s %s]", got.Commits, one, two)
	}
	expectHistoryFindings(t, got)
}

func Test_CHK_070_HistoryBranchingRangeTopologicalOrder(t *testing.T) {
	f := newHistoryFixture(t)
	f.git("branch", "side")
	f.put("notes/main-one.txt", "main one\n")
	mainOne := f.commit("main first")
	f.put("notes/main-two.txt", "main two\n")
	mainTwo := f.commit("main second")
	f.git("checkout", "-q", "side")
	f.put("notes/side-one.txt", "side one\n")
	sideOne := f.commit("side first")
	f.put("notes/side-two.txt", "side two\n")
	sideTwo := f.commit("side second")
	f.git("checkout", "-q", "main")
	f.git("merge", "--no-ff", "-m", "merge side", "side")
	merge := f.git("rev-parse", "HEAD")

	got := f.report(f.base + ".." + merge)
	want := map[string]bool{mainOne: true, mainTwo: true, sideOne: true, sideTwo: true, merge: true}
	if len(got.Commits) != len(want) {
		t.Errorf("branching commits = %#v; want exactly five commits", got.Commits)
	}
	positions := make(map[string]int)
	for i, commit := range got.Commits {
		if !want[commit] {
			t.Errorf("unexpected commit %s in %#v", commit, got.Commits)
		}
		if _, duplicate := positions[commit]; duplicate {
			t.Errorf("duplicate commit %s in %#v", commit, got.Commits)
		}
		positions[commit] = i
	}
	for commit := range want {
		if _, present := positions[commit]; !present {
			t.Errorf("missing commit %s in %#v", commit, got.Commits)
		}
	}
	for _, edge := range [][2]string{{mainOne, mainTwo}, {sideOne, sideTwo}, {mainTwo, merge}, {sideTwo, merge}} {
		parent, hasParent := positions[edge[0]]
		child, hasChild := positions[edge[1]]
		if hasParent && hasChild && parent >= child {
			t.Errorf("parent %s must precede child %s in %#v", edge[0], edge[1], got.Commits)
		}
	}
	expectHistoryFindings(t, got)
}

func Test_CHK_070_HistoryDefaultRangeUsesLocalConfiguredBase(t *testing.T) {
	f := newHistoryFixture(t)
	f.put("fireweed.yaml", strings.Replace(historyConfig, "base: main", "base: review-base", 1))
	f.commit("configure review base")
	f.git("branch", "review-base")
	f.put("notes/common.txt", "common\n")
	common := f.commit("common after review base")
	f.git("branch", "feature")
	f.git("checkout", "-q", "feature")
	f.put("notes/one.txt", "one")
	one := f.commit("first")
	f.put("notes/two.txt", "two")
	two := f.commit("second")
	f.git("checkout", "-q", "review-base")
	f.put("notes/review-only.txt", "review only\n")
	reviewOnly := f.commit("review base advances independently")
	f.git("checkout", "-q", "feature")
	got := f.report("")
	if !reflect.DeepEqual(got.Commits, []string{common, one, two}) {
		t.Errorf("default commits = %#v; want [%s %s %s]", got.Commits, common, one, two)
	}
	for _, commit := range got.Commits {
		if commit == reviewOnly {
			t.Errorf("default commits include configured-base-only commit %s", reviewOnly)
		}
	}
}

func Test_CHK_071_CHK_072_HistoryPhaseContentsAndDeletionControl(t *testing.T) {
	f := newHistoryFixture(t)
	path := "changes/0001-work.md"
	f.change(path, "evals", "")
	f.commit("add queued change")
	f.change(path, "audit", "")
	f.put("locked/config.txt", "worker edit\n")
	worker := f.commit("ordinary subject")
	f.remove(path)
	f.put("locked/config.txt", "author edit\n")
	f.commit("fireweed impl 0001-work")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got, historyFinding{"H01", "locked/config.txt", worker})
}

func Test_CHK_072_CHK_073_CHK_077_HistoryPlaybookOverrideAndProjectedInterface(t *testing.T) {
	f := newHistoryFixture(t)
	pending := "changes/0001-work.md"
	archive := "changes/archive/0001-work.md"
	set := "set:\n  interface: [newguard/**]\n"
	f.change(pending, "impl", set)
	f.commit("ready")
	projected, err := project.Project(f.root, "")
	if err != nil {
		t.Fatalf("Project with interface set: %v", err)
	}
	if len(projected.Modules) != 1 || projected.Modules[0].Name != "alpha" || !reflect.DeepEqual(projected.Modules[0].Zones.Interface, []string{"newguard/**"}) {
		t.Fatalf("projected alpha interface = %#v; want [newguard/**]", projected.Modules)
	}
	f.move(pending, archive)
	f.change(archive, "applied", set)
	f.put("packages/alpha/newguard/api.go", "package newguard\n")
	f.put(".fireweed/playbooks/evals.md", "local override\n")
	applied := f.commit("phase assigned")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got,
		historyFinding{"H02", "packages/alpha/newguard/api.go", applied},
		historyFinding{"H01", ".fireweed/playbooks/evals.md", applied})
}

func Test_CHK_071_CHK_073_HistoryAppliedImplTouchesEvalsAndInterface(t *testing.T) {
	f := newHistoryFixture(t)
	pending := "changes/0001-work.md"
	archive := "changes/archive/0001-work.md"
	f.change(pending, "impl", "")
	f.commit("ready for implementation")
	f.move(pending, archive)
	f.change(archive, "applied", "")
	f.put("packages/alpha/evals/existing_test.go", "changed test\n")
	f.put("packages/alpha/api.go", "changed api\n")
	applied := f.commit("unrelated subject")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got,
		historyFinding{"H02", "packages/alpha/evals/existing_test.go", applied},
		historyFinding{"H02", "packages/alpha/api.go", applied})
}

func Test_CHK_074_HistoryAuthorAppendOnlyModifyDeleteRename(t *testing.T) {
	for _, action := range []string{"modify", "delete", "rename"} {
		t.Run(action, func(t *testing.T) {
			f := newHistoryFixture(t)
			path := "packages/alpha/ledger/entry.txt"
			switch action {
			case "modify":
				f.put(path, "modified\n")
			case "delete":
				f.remove(path)
			case "rename":
				f.move(path, "packages/alpha/ledger/renamed.txt")
			}
			commit := f.commit("human maintenance")
			got := f.report(f.base + "..HEAD")
			if len(got.Diagnostics) == 0 {
				t.Errorf("%s diagnostics = %#v; want H03", action, got.Diagnostics)
			}
			for _, d := range got.Diagnostics {
				if d.Code != "H03" || !strings.Contains(d.Message, commit) || (d.Path != path && d.Path != "packages/alpha/ledger/renamed.txt") {
					t.Errorf("%s unexpected diagnostic %#v", action, d)
				}
			}
		})
	}
}

func Test_CHK_074_HistoryMergeResolutionInMergeOnlyRange(t *testing.T) {
	f := newHistoryFixture(t)
	f.git("checkout", "-qb", "side")
	f.put("packages/alpha/ledger/entry.txt", "side\n")
	f.commit("side edits ledger")
	f.git("checkout", "-q", "main")
	f.put("packages/alpha/ledger/entry.txt", "main\n")
	f.commit("main edits ledger")
	result, err := gitstore.Run(context.Background(), gitstore.Request{Args: []string{"git", "-C", f.root, "merge", "--no-ff", "--no-commit", "side"}, Dir: f.root})
	if err == nil || !strings.Contains(result.Stdout+result.Stderr, "CONFLICT") {
		t.Fatalf("expected ordinary merge conflict, got %v: %s %s", err, result.Stdout, result.Stderr)
	}
	f.put("packages/alpha/ledger/entry.txt", "resolved uniquely\n")
	merge := f.commit("resolve conflict")
	got := f.report(merge + "^!")
	if !reflect.DeepEqual(got.Commits, []string{merge}) {
		t.Errorf("merge-only selection = %#v; want [%s]", got.Commits, merge)
	}
	expectHistoryFindings(t, got, historyFinding{"H03", "packages/alpha/ledger/entry.txt", merge})
}

func Test_CHK_075_HistoryArchiveModifyAndDeletionControl(t *testing.T) {
	f := newHistoryFixture(t)
	path := "changes/archive/0001-old.md"
	f.change(path, "applied", "")
	f.commit("archived baseline")
	f.put(path, "---\nmodule: alpha\nphase: applied\n---\n\nUpdated archive.\n")
	modified := f.commit("edit archive")
	f.remove(path)
	f.commit("delete archive")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got, historyFinding{"H04", path, modified})
}

func Test_CHK_076_HistoryMultipleSquashedChangesKeepH03H04(t *testing.T) {
	f := newHistoryFixture(t)
	archiveOld := "changes/archive/0000-old.md"
	f.change(archiveOld, "applied", "")
	archiveBaseline := f.commit("existing archive")
	for _, path := range []string{"changes/archive/0001-one.md", "changes/archive/0002-two.md"} {
		f.change(path, "applied", "")
	}
	f.put("locked/config.txt", "squash protection\n")
	f.put("packages/alpha/evals/existing_test.go", "squash eval\n")
	f.put("packages/alpha/ledger/entry.txt", "squash ledger\n")
	f.put(archiveOld, "---\nmodule: alpha\nphase: applied\n---\n\nChanged old archive.\n")
	squash := f.commit("squashed changes")
	got := f.report(archiveBaseline + "..HEAD")
	wantSquashed := map[checks.SquashedChange]bool{
		{Commit: squash, Change: "changes/archive/0001-one.md"}: true,
		{Commit: squash, Change: "changes/archive/0002-two.md"}: true,
	}
	if len(got.Squashed) != len(wantSquashed) {
		t.Errorf("squashed = %#v; want %#v", got.Squashed, wantSquashed)
	} else {
		for _, entry := range got.Squashed {
			if !wantSquashed[entry] {
				t.Errorf("unexpected squashed entry %#v; want %#v", entry, wantSquashed)
			}
			delete(wantSquashed, entry)
		}
		if len(wantSquashed) != 0 {
			t.Errorf("missing squashed entries %#v", wantSquashed)
		}
	}
	expectHistoryFindings(t, got,
		historyFinding{"H03", "packages/alpha/ledger/entry.txt", squash},
		historyFinding{"H04", archiveOld, squash})
}

func Test_CHK_077_HistoryUsesRecordedBeforeAfterZones(t *testing.T) {
	f := newHistoryFixture(t)
	pending := "changes/0001-work.md"
	f.change(pending, "impl", "")
	f.commit("ready")
	f.move(pending, "changes/archive/0001-work.md")
	f.change("changes/archive/0001-work.md", "applied", "")
	f.put("packages/alpha/evals/new_test.go", "new eval\n")
	applied := f.commit("apply")
	f.put("fireweed.yaml", strings.Replace(historyConfig, "evals: [evals/**]", "evals: [later/**]", 1))
	f.commit("later config")
	got := f.report(f.base + ".." + applied)
	expectHistoryFindings(t, got, historyFinding{"H02", "packages/alpha/evals/new_test.go", applied})
}

func Test_CHK_077_HistoryChecksBothRecordedZoneMaps(t *testing.T) {
	f := newHistoryFixture(t)
	f.put("fireweed.yaml", strings.Replace(historyConfig, "append_only: [ledger/**]", "append_only: [newledger/**]", 1))
	f.put("packages/alpha/ledger/entry.txt", "changed under before zone\n")
	f.put("packages/alpha/newledger/entry.txt", "changed under after zone\n")
	commit := f.commit("author changes zone and both existing files")
	got := f.report(f.base + "..HEAD")
	expectHistoryFindings(t, got,
		historyFinding{"H03", "packages/alpha/ledger/entry.txt", commit},
		historyFinding{"H03", "packages/alpha/newledger/entry.txt", commit})
}

func Test_CHK_078_HistoryRejectsUnreadableRangeAndRecordedDocuments(t *testing.T) {
	for _, tc := range []struct {
		name       string
		edit       func(historyFixture)
		rangeValue string
	}{
		{"invalid_range", nil, "missing-ref..HEAD"},
		{"missing_config", func(f historyFixture) { f.remove("fireweed.yaml") }, ""},
		{"malformed_config", func(f historyFixture) { f.put("fireweed.yaml", "modules: [broken\n") }, ""},
		{"missing_SPEC", func(f historyFixture) { f.remove("packages/alpha/SPEC.md") }, ""},
		{"malformed_change", func(f historyFixture) { f.put("changes/0001-work.md", "---\nmodule: [broken\n---\n") }, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := newHistoryFixture(t)
			if tc.edit != nil {
				tc.edit(f)
				f.commit("record invalid input")
			}
			revisionRange := tc.rangeValue
			if revisionRange == "" {
				revisionRange = f.base + "..HEAD"
			}
			if _, err := checks.History(f.root, revisionRange); err == nil {
				t.Errorf("History(%q) accepted %s", revisionRange, tc.name)
			}
		})
	}
}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562/packages/checks/internal/engine/doctor.go
diff --git a/packages/checks/internal/engine/doctor.go b/packages/checks/internal/engine/doctor.go
index 6d69904707b2fb0e65938a4b6a72db7f41fce294..fed7afde0de7f6e8ca57e17186a6eafdcc0e54ca
--- a/packages/checks/internal/engine/doctor.go
+++ b/packages/checks/internal/engine/doctor.go
@@ -247,11 +247,9 @@
 	switch name {
 	case "undeclared", "internal_eval":
 		ds, e := boundaryProbe(ctx, s, m, boundHook, bin, timeout)
-		if e == nil {
-			for _, d := range ds {
-				if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
-					pass = true
-				}
+		for _, d := range ds {
+			if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
+				pass = true
 			}
 		}
 		detail = fmt.Sprintf("%v %v", ds, e)

exec
/bin/zsh -lc 'gofmt -w packages/checks/internal/engine/*.go && git diff -- packages/checks/internal/engine/doctor.go && git status --short' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc 'go test ./packages/checks/evals' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 580ms:
?? DRILL_NOTES.md
?? packages/checks/internal/

 exited 1 in 579ms:
go: cannot find main module, but found .git/config in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
	to create a module there, run:
	go mod init

codex
I found a Doctor bug: the boundary hook’s expected violation returns an X01 diagnostic and a nonnil error, but the canary probe required a nil error, so it always marked detection as failed. I’ve corrected that check. I’m also tightening History’s use of recorded zone maps before verification.
apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562/packages/checks/internal/engine/history.go
diff --git a/packages/checks/internal/engine/doctor.go b/packages/checks/internal/engine/doctor.go
index 6d69904707b2fb0e65938a4b6a72db7f41fce294..fed7afde0de7f6e8ca57e17186a6eafdcc0e54ca
--- a/packages/checks/internal/engine/doctor.go
+++ b/packages/checks/internal/engine/doctor.go
@@ -247,11 +247,9 @@
 	switch name {
 	case "undeclared", "internal_eval":
 		ds, e := boundaryProbe(ctx, s, m, boundHook, bin, timeout)
-		if e == nil {
-			for _, d := range ds {
-				if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
-					pass = true
-				}
+		for _, d := range ds {
+			if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
+				pass = true
 			}
 		}
 		detail = fmt.Sprintf("%v %v", ds, e)
diff --git a/packages/checks/internal/engine/history.go b/packages/checks/internal/engine/history.go
index c24b4442fabf69509ab2a74ec3856f9e375d48d5..49f7a17daff01e815fd6ed4866666d36f2ad19e6
--- a/packages/checks/internal/engine/history.go
+++ b/packages/checks/internal/engine/history.go
@@ -210,11 +210,15 @@
 			if e != nil {
 				return rep, e
 			}
-			before = append(before, s)
 			m, e := historyChanges(root, parent)
 			if e != nil {
 				return rep, e
 			}
+			s, e = projectedHistory(root, parent, s, m)
+			if e != nil {
+				return rep, e
+			}
+			before = append(before, s)
 			beforeChanges = append(beforeChanges, m)
 		}
 		var diffs []hchange

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562/packages/checks/internal/engine/doctor.go
diff --git a/packages/checks/internal/engine/doctor.go b/packages/checks/internal/engine/doctor.go
index 6d69904707b2fb0e65938a4b6a72db7f41fce294..d25debb3484e1c8cae5f1bf90cea1de4484fbf00
--- a/packages/checks/internal/engine/doctor.go
+++ b/packages/checks/internal/engine/doctor.go
@@ -179,11 +179,6 @@
 		return nil, fmt.Errorf("canary canaries.yaml version %d", raw.Version)
 	}
 	out := map[string]map[string]string{"baseline": raw.Baseline, "undeclared": raw.Undeclared, "internal_eval": raw.InternalEval, "failing_eval": raw.FailingEval, "missing_id": raw.MissingID}
-	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id"} {
-		if len(out[name]) == 0 {
-			return nil, fmt.Errorf("canary %s missing", name)
-		}
-	}
 	return out, nil
 }
 func evalProbe(ctx context.Context, s project.Snapshot, m project.Module, hook any, bin string, timeout time.Duration) (model.Report, error) {
@@ -247,11 +242,9 @@
 	switch name {
 	case "undeclared", "internal_eval":
 		ds, e := boundaryProbe(ctx, s, m, boundHook, bin, timeout)
-		if e == nil {
-			for _, d := range ds {
-				if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
-					pass = true
-				}
+		for _, d := range ds {
+			if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
+				pass = true
 			}
 		}
 		detail = fmt.Sprintf("%v %v", ds, e)
@@ -429,6 +422,14 @@
 		doctorDiag(&r, "C14", e.Error(), false)
 		return r
 	}
+	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id"} {
+		if len(maps[name]) == 0 {
+			doctorDiag(&r, "C14", "canary "+name+" missing", false)
+		}
+	}
+	if len(maps["baseline"]) == 0 {
+		return r
+	}
 	modText := ""
 	if b, e := os.ReadFile(filepath.Join(scratch, "go.mod")); e == nil {
 		for _, l := range strings.Split(string(b), "\n") {
@@ -519,6 +520,9 @@
 			vars["other_dir"] = dirs[1]
 		}
 		for _, name := range []string{"undeclared", "internal_eval", "failing_eval", "missing_id"} {
+			if len(maps[name]) == 0 {
+				continue
+			}
 			if bh == nil && (name == "undeclared" || name == "internal_eval") {
 				continue
 			}
diff --git a/packages/checks/internal/engine/history.go b/packages/checks/internal/engine/history.go
index c24b4442fabf69509ab2a74ec3856f9e375d48d5..49f7a17daff01e815fd6ed4866666d36f2ad19e6
--- a/packages/checks/internal/engine/history.go
+++ b/packages/checks/internal/engine/history.go
@@ -210,11 +210,15 @@
 			if e != nil {
 				return rep, e
 			}
-			before = append(before, s)
 			m, e := historyChanges(root, parent)
 			if e != nil {
 				return rep, e
 			}
+			s, e = projectedHistory(root, parent, s, m)
+			if e != nil {
+				return rep, e
+			}
+			before = append(before, s)
 			beforeChanges = append(beforeChanges, m)
 		}
 		var diffs []hchange

exec
/bin/zsh -lc "sed -n '790,950p' packages/checks/evals/doctor_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc 'git status --short && find packages/checks/internal -type f -maxdepth 3 -print' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc "sed -n '300,370p' packages/checks/evals/doctor_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 569ms:
		remaining[d]++
	}
	for _, d := range want {
		if remaining[d] == 0 {
			t.Errorf("complete project diagnostic %+v lost: %s", d, doctorText(report))
		}
		remaining[d]--
	}
	if strings.Contains(doctorLog(f), "setup:") || strings.Contains(doctorLog(f), "eval:") || strings.Contains(doctorLog(f), "boundaries:") || strings.Contains(doctorLog(f), "worker:") || strings.Contains(doctorLog(f), "audit:") {
		t.Errorf("static error ran probes: %s", doctorLog(f))
	}
}

func Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_CHK_060_CHK_061_CHK_067_RealGoRecipe(t *testing.T) {
	f := newDoctorFixture(t)
	f.commit()
	f.bin = doctorBinary(t)
	f.editConfig(func(config map[string]any) {
		config["modules"] = []string{"packages/alpha"}
		config["hooks"].(map[string]any)["setup"] = fmt.Sprintf("printf 'setup:%%s\\n' \"$PWD\" >> %q; test -z \"$(git symbolic-ref -q HEAD)\" && test -f .fireweed/doctor-untracked && go mod download", f.log)
	})
	f.put(".fireweed/doctor-untracked", "ready\n")
	for _, d := range project.Check(f.root) {
		if !d.Warning {
			t.Fatalf("invalid Go doctor fixture: %+v", d)
		}
	}
	beforeStatus, beforeWorktrees := f.git("status", "--porcelain", "--untracked-files=all"), worktrees(f)
	report := checks.Doctor(context.Background(), f.request(true))
	if doctorErrors(report) {
		t.Errorf("real Go recipe doctor failed: %s", doctorText(report))
	}
	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id", "recovery"} {
		if !hasDoctorCheck(report, name, true) {
			t.Errorf("missing passing %s outcome: %s", name, doctorText(report))
		}
	}
	recoveries := 0
	for _, check := range report.Checks {
		if strings.Contains(strings.ToLower(check.Name), "recovery") && check.Passed {
			recoveries++
		}
	}
	if recoveries < 4 {
		t.Errorf("clean recovery not reported after each canary: %s", doctorText(report))
	}
	modules := map[string]bool{}
	for _, check := range report.Checks {
		if check.Module != "" && check.Module != "alpha" {
			modules[check.Module] = true
		}
	}
	if len(modules) < 2 {
		t.Errorf("two distinct throwaway modules absent: %s", doctorText(report))
	}
	if !strings.Contains(doctorLog(f), "setup:") || !strings.Contains(doctorLog(f), "/fireweed/worktrees/") {
		t.Errorf("setup did not run in scratch: %s", doctorLog(f))
	}
	if strings.Contains(doctorLog(f), "worker:") || strings.Contains(doctorLog(f), "audit:") {
		t.Errorf("NoAgent ran agents: %s", doctorLog(f))
	}
	if got := f.git("status", "--porcelain", "--untracked-files=all"); got != beforeStatus {
		t.Errorf("caller changed: before %q after %q", beforeStatus, got)
	}
	if got := worktrees(f); got != beforeWorktrees {
		t.Errorf("scratch registration leaked: before %q after %q", beforeWorktrees, got)
	}
	assertScratchRemoved(t, f)
}

func Test_CHK_053_NestedModuleTemplatesRenderForBothCanaries(t *testing.T) {

 succeeded in 573ms:
func Test_CHK_060_OverridesMustDetectTheirCanaries(t *testing.T) {
	for _, tc := range []struct {
		name, hook, command, script, canary string
	}{
		{"eval", "eval", "sh .fireweed/override-eval.sh {module} {module_dir} {report}", `#!/bin/sh
prefix=$(sed -n 's/^id_prefix: //p' "$2/SPEC.md" | head -1)
printf '<testsuite><testcase name="%s-001"/></testsuite>' "$prefix" > "$3"
`, "failing_eval"},
		{"boundary", "boundaries", "sh .fireweed/override-boundary.sh {module}", "#!/bin/sh\nexit 0\n", "internal_eval"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := localDoctorFixture(t)
			path := ".fireweed/override-" + tc.name + ".sh"
			f.put(path, tc.script)
			f.put("packages/alpha/SPEC.md", strings.Replace(string(mustRead(t, filepath.Join(f.root, "packages/alpha/SPEC.md"))), "depends_on: []", "depends_on: []\nhooks:\n  "+tc.hook+": "+tc.command, 1))
			f.commit()
			r := checks.Doctor(context.Background(), f.request(true))
			diagnosed := false
			for _, d := range r.Diagnostics {
				message := strings.ToLower(d.Message)
				boundaryHook := tc.name == "boundary" && !strings.Contains(message, "global") &&
					(strings.Contains(message, "alpha boundary hook") ||
						strings.Contains(message, "alpha boundaries hook") ||
						strings.Contains(message, "alpha boundaries override boundary hook") ||
						strings.Contains(message, "alpha boundaries override boundaries hook"))
				if d.Code == "C14" && !d.Warning && strings.Contains(message, tc.canary) &&
					(tc.name == "eval" && strings.Contains(message, "alpha "+tc.hook+" override") || boundaryHook) {
					diagnosed = true
				}
			}
			if !diagnosed {
				t.Errorf("broken %s override missed its %s canary without C14: %s", tc.name, tc.canary, doctorText(r))
			}
		})
	}
}

func Test_CHK_063_CHK_064_CHK_065_CHK_066_CHK_067_AgentProbes(t *testing.T) {
	f := localDoctorFixture(t)
	f.commit()
	for _, d := range project.Check(f.root) {
		if !d.Warning {
			t.Fatalf("invalid agent fixture: %+v", d)
		}
	}
	r := checks.Doctor(context.Background(), f.request(false))
	log := doctorLog(f)
	if doctorErrors(r) || strings.Count(log, "worker:") != 1 || strings.Count(log, "audit:") != 1 || strings.Count(log, "wrap:") != 2 || !strings.Contains(log, "FIREWEED_DOCTOR_OK") {
		t.Errorf("worker/audit/wrap probes: %s; log=%s", doctorText(r), log)
	}
	workerPrompt, auditPrompt := false, false
	for _, line := range strings.Split(log, "\n") {
		if fields := strings.SplitN(line, ":", 3); len(fields) == 3 && strings.Contains(fields[2], "FIREWEED_DOCTOR_OK") {
			switch fields[0] {
			case "worker":
				workerPrompt = true
			case "audit":
				auditPrompt = true
			}
		}
	}
	if !workerPrompt {
		t.Errorf("worker did not receive its fixed-reply prompt: %s", log)
	}
	if !auditPrompt {
		t.Errorf("audit did not receive its fixed-reply prompt: %s", log)
	}
	for _, part := range []string{"/fireweed/worktrees/", "worker:", "audit:"} {
		if !strings.Contains(log, part) {
			t.Errorf("scratch/role marker %s absent: %s", part, log)
		}
	}

	g := localDoctorFixture(t)
	g.editConfig(func(config map[string]any) {
		delete(config["audit"].(map[string]any), "command")
		config["agent"].(map[string]any)["command"] = "bash .fireweed/agent.sh {prompt} {git_common_dir}"
	})
	g.commit()
	rr := checks.Doctor(context.Background(), g.request(false))
	if !hasDoctorDiagnostic(rr, "C15", true, "shared") || strings.Count(doctorLog(g), "worker:") != 2 {
		t.Errorf("audit fallback: %s; log=%s", doctorText(rr), doctorLog(g))
	}

	h := localDoctorFixture(t)
	h.put(".fireweed/agent.sh", fmt.Sprintf("#!/bin/bash\nfor i in $(seq 1 25); do printf 'line-%%02d\\n' \"$i\"; done\necho RATE LIMIT reached\nexit 9\n"))
	h.commit()
	bad := checks.Doctor(context.Background(), h.request(false))
	if !hasDoctorDiagnostic(bad, "C15", false, "usage limit", "rate limit", "line-25") {
		t.Errorf("usage-limit tail missing: %s", doctorText(bad))
	}
	if strings.Contains(doctorText(bad), "line-01") {
		t.Errorf("agent output was not trimmed to final 20 lines: %s", doctorText(bad))
	}
	outOfTail := localDoctorFixture(t)
	outOfTail.put(".fireweed/agent.sh", "#!/bin/bash\necho RATE LIMIT early\nfor i in $(seq 1 25); do printf 'later-%02d\\n' \"$i\"; done\nexit 9\n")
	outOfTail.commit()
	ordinary := checks.Doctor(context.Background(), outOfTail.request(false))
	if !hasDoctorDiagnostic(ordinary, "C15", false, "agent.sh", "later-25") {
		t.Errorf("ordinary out-of-tail failure absent: %s", doctorText(ordinary))
	}
	for _, d := range ordinary.Diagnostics {
		if d.Code == "C15" && !d.Warning && strings.Contains(strings.ToLower(d.Message), "usage limit") {
			t.Errorf("out-of-tail pattern classified as usage limit: %s", doctorText(ordinary))
		}
	}
	for _, tc := range []struct{ name, script string }{{"wrong reply", "#!/bin/bash\necho WRONG_REPLY\n"}, {"nonzero with reply", "#!/bin/bash\necho FIREWEED_DOCTOR_OK\nexit 7\n"}, {"ordinary failure", "#!/bin/bash\nfor i in $(seq 1 25); do printf 'ordinary-%02d\\n' \"$i\"; done\nexit 8\n"}} {
		t.Run(tc.name, func(t *testing.T) {
			q := localDoctorFixture(t)
			q.put(".fireweed/agent.sh", tc.script)
			q.commit()
			r := checks.Doctor(context.Background(), q.request(false))
			if !hasDoctorDiagnostic(r, "C15", false, "agent.sh") {
				t.Errorf("agent failure accepted: %s", doctorText(r))
			}
			if tc.name == "ordinary failure" && (!hasDoctorDiagnostic(r, "C15", false, "ordinary-25") || strings.Contains(doctorText(r), "ordinary-01")) {
				t.Errorf("ordinary failure tail incorrect: %s", doctorText(r))
			}
		})
	}
	wrongAudit := localDoctorFixture(t)
	wrongAudit.put(".fireweed/audit.sh", "#!/bin/bash\necho WRONG_REPLY\n")
	wrongAudit.commit()
	auditFailure := checks.Doctor(context.Background(), wrongAudit.request(false))
	if !hasDoctorDiagnostic(auditFailure, "C15", false, "audit.sh") || hasDoctorCheck(auditFailure, "audit", true) {
		t.Errorf("wrong audit reply accepted: %s", doctorText(auditFailure))
	}

	i := localDoctorFixture(t)
	i.put(".fireweed/audit.sh", "#!/bin/bash\nsleep 2\n")
	i.commit()
	req := i.request(false)
	req.Timeout = 100 * time.Millisecond
	start := time.Now()
	timed := checks.Doctor(context.Background(), req)
	if !hasDoctorDiagnostic(timed, "C15", false, "audit") || time.Since(start) > 4*time.Second {
		t.Errorf("audit timeout not bounded: %s; elapsed=%s", doctorText(timed), time.Since(start))
	}

	defaulted := localDoctorFixture(t)
	defaulted.put(".fireweed/agent.sh", "#!/bin/bash\nsleep 0.2\nprintf 'FIREWEED_DOCTOR_OK\\n'\n")
	defaulted.put(".fireweed/audit.sh", "#!/bin/bash\nsleep 0.2\nprintf 'FIREWEED_DOCTOR_OK\\n'\n")
	defaulted.commit()
	defaultRequest := defaulted.request(false)
	defaultRequest.Timeout = 0
	defaultReport := checks.Doctor(context.Background(), defaultRequest)
	if doctorErrors(defaultReport) || !hasDoctorCheck(defaultReport, "worker", true) || !hasDoctorCheck(defaultReport, "audit", true) {
		t.Errorf("unset timeout did not give both agent probes their default allowance: %s", doctorText(defaultReport))
	}

	j := localDoctorFixture(t)
	j.commit()
	skip := checks.Doctor(context.Background(), j.request(true))
	if strings.Contains(doctorLog(j), "worker:") || strings.Contains(doctorLog(j), "audit:") || !hasDoctorCheck(skip, "undeclared", true) || !strings.Contains(doctorLog(j), j.bin) {
		t.Errorf("NoAgent skipped hooks or ran agent: %s; log=%s", doctorText(skip), doctorLog(j))
	}
}

func Test_CHK_065_UnsetTimeoutBoundsAuditProbe(t *testing.T) {
	f := localDoctorFixture(t)
	f.put(".fireweed/audit.sh", fmt.Sprintf("#!/bin/bash\nprintf 'audit:%%s:%%s\n' \"$PWD\" \"$1\" >> %q\nprintf 'audit-timeout-started\n'\nsleep 210\nprintf 'FIREWEED_DOCTOR_OK\n'\n", f.log))

 succeeded in 595ms:
?? DRILL_NOTES.md
?? packages/checks/internal/
packages/checks/internal/engine/zones_boundaries.go
packages/checks/internal/engine/core.go
packages/checks/internal/engine/history.go
packages/checks/internal/engine/doctor.go

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562/DRILL_NOTES.md
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
index 9aaeb601a62191f4c917ce09f650b743f9cd75b4..942ee041d90cf3f54006b8fd0ecade8327c8fbc2
--- a/DRILL_NOTES.md
+++ b/DRILL_NOTES.md
@@ -6,4 +6,5 @@
 
 ## Gaps
 
-No known behavioral gaps in the target checks module. The supplied workspace omits the Go manifest and dependency implementations, so the retained evals could not be run here. Native verification remains with the outside controller.
+- CHK-055 (checks): The supplied workspace has no `go.mod` or dependent internal implementations, so the retained Go evals cannot run here. The prior `PackageFailure`/coverage result cannot be reproduced or cleared in this input alone.
+- no spec (app, runner, workflow): The prior verification also reported `PackageFailure` for these modules. Their implementations are outside the authorized checks rebuild paths.
diff --git a/packages/checks/internal/engine/doctor.go b/packages/checks/internal/engine/doctor.go
index 6d69904707b2fb0e65938a4b6a72db7f41fce294..d25debb3484e1c8cae5f1bf90cea1de4484fbf00
--- a/packages/checks/internal/engine/doctor.go
+++ b/packages/checks/internal/engine/doctor.go
@@ -179,11 +179,6 @@
 		return nil, fmt.Errorf("canary canaries.yaml version %d", raw.Version)
 	}
 	out := map[string]map[string]string{"baseline": raw.Baseline, "undeclared": raw.Undeclared, "internal_eval": raw.InternalEval, "failing_eval": raw.FailingEval, "missing_id": raw.MissingID}
-	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id"} {
-		if len(out[name]) == 0 {
-			return nil, fmt.Errorf("canary %s missing", name)
-		}
-	}
 	return out, nil
 }
 func evalProbe(ctx context.Context, s project.Snapshot, m project.Module, hook any, bin string, timeout time.Duration) (model.Report, error) {
@@ -247,11 +242,9 @@
 	switch name {
 	case "undeclared", "internal_eval":
 		ds, e := boundaryProbe(ctx, s, m, boundHook, bin, timeout)
-		if e == nil {
-			for _, d := range ds {
-				if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
-					pass = true
-				}
+		for _, d := range ds {
+			if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
+				pass = true
 			}
 		}
 		detail = fmt.Sprintf("%v %v", ds, e)
@@ -429,6 +422,14 @@
 		doctorDiag(&r, "C14", e.Error(), false)
 		return r
 	}
+	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id"} {
+		if len(maps[name]) == 0 {
+			doctorDiag(&r, "C14", "canary "+name+" missing", false)
+		}
+	}
+	if len(maps["baseline"]) == 0 {
+		return r
+	}
 	modText := ""
 	if b, e := os.ReadFile(filepath.Join(scratch, "go.mod")); e == nil {
 		for _, l := range strings.Split(string(b), "\n") {
@@ -519,6 +520,9 @@
 			vars["other_dir"] = dirs[1]
 		}
 		for _, name := range []string{"undeclared", "internal_eval", "failing_eval", "missing_id"} {
+			if len(maps[name]) == 0 {
+				continue
+			}
 			if bh == nil && (name == "undeclared" || name == "internal_eval") {
 				continue
 			}
diff --git a/packages/checks/internal/engine/history.go b/packages/checks/internal/engine/history.go
index c24b4442fabf69509ab2a74ec3856f9e375d48d5..49f7a17daff01e815fd6ed4866666d36f2ad19e6
--- a/packages/checks/internal/engine/history.go
+++ b/packages/checks/internal/engine/history.go
@@ -210,11 +210,15 @@
 			if e != nil {
 				return rep, e
 			}
-			before = append(before, s)
 			m, e := historyChanges(root, parent)
 			if e != nil {
 				return rep, e
 			}
+			s, e = projectedHistory(root, parent, s, m)
+			if e != nil {
+				return rep, e
+			}
+			before = append(before, s)
 			beforeChanges = append(beforeChanges, m)
 		}
 		var diffs []hchange

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562/packages/checks/internal/engine/doctor.go
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
index 9aaeb601a62191f4c917ce09f650b743f9cd75b4..942ee041d90cf3f54006b8fd0ecade8327c8fbc2
--- a/DRILL_NOTES.md
+++ b/DRILL_NOTES.md
@@ -6,4 +6,5 @@
 
 ## Gaps
 
-No known behavioral gaps in the target checks module. The supplied workspace omits the Go manifest and dependency implementations, so the retained evals could not be run here. Native verification remains with the outside controller.
+- CHK-055 (checks): The supplied workspace has no `go.mod` or dependent internal implementations, so the retained Go evals cannot run here. The prior `PackageFailure`/coverage result cannot be reproduced or cleared in this input alone.
+- no spec (app, runner, workflow): The prior verification also reported `PackageFailure` for these modules. Their implementations are outside the authorized checks rebuild paths.
diff --git a/packages/checks/internal/engine/doctor.go b/packages/checks/internal/engine/doctor.go
index 6d69904707b2fb0e65938a4b6a72db7f41fce294..5ade63b0352aed46e9d26313fd65293e0bfd68f1
--- a/packages/checks/internal/engine/doctor.go
+++ b/packages/checks/internal/engine/doctor.go
@@ -179,11 +179,6 @@
 		return nil, fmt.Errorf("canary canaries.yaml version %d", raw.Version)
 	}
 	out := map[string]map[string]string{"baseline": raw.Baseline, "undeclared": raw.Undeclared, "internal_eval": raw.InternalEval, "failing_eval": raw.FailingEval, "missing_id": raw.MissingID}
-	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id"} {
-		if len(out[name]) == 0 {
-			return nil, fmt.Errorf("canary %s missing", name)
-		}
-	}
 	return out, nil
 }
 func evalProbe(ctx context.Context, s project.Snapshot, m project.Module, hook any, bin string, timeout time.Duration) (model.Report, error) {
@@ -247,11 +242,9 @@
 	switch name {
 	case "undeclared", "internal_eval":
 		ds, e := boundaryProbe(ctx, s, m, boundHook, bin, timeout)
-		if e == nil {
-			for _, d := range ds {
-				if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
-					pass = true
-				}
+		for _, d := range ds {
+			if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
+				pass = true
 			}
 		}
 		detail = fmt.Sprintf("%v %v", ds, e)
@@ -332,16 +325,25 @@
 		limit = 3 * time.Minute
 	}
 	out, stderr, exit, runErr := run(ctx, scratch, argv, nil, limit)
-	logPath := filepath.Join(req.Root, ".fireweed", "runs", runID, "doctor-"+role+"-1.log")
-	_ = os.MkdirAll(filepath.Dir(logPath), 0755)
-	_ = os.WriteFile(logPath, []byte(out+stderr), 0644)
+	logRoot, pathErr := filepath.Abs(req.Root)
+	if pathErr != nil {
+		logRoot = req.Root
+	}
+	logPath := filepath.Join(logRoot, ".fireweed", "runs", runID, "doctor-"+role+"-1.log")
+	logErr := os.MkdirAll(filepath.Dir(logPath), 0755)
+	if logErr == nil {
+		logErr = os.WriteFile(logPath, []byte(out+stderr), 0644)
+	}
 	detail := "log: " + logPath
 	combined := out + stderr
 	tail := tail20(combined)
-	passed := runErr == nil && exit == 0 && strings.Contains(out, "FIREWEED_DOCTOR_OK")
+	passed := logErr == nil && runErr == nil && exit == 0 && strings.Contains(out, "FIREWEED_DOCTOR_OK")
 	doctorCheck(r, role, "", passed, detail)
 	if !passed {
 		message := fmt.Sprintf("%s probe command %v failed: %v exit %d\n%s", role, argv, runErr, exit, tail)
+		if logErr != nil {
+			message += "\nprobe log: " + logErr.Error()
+		}
 		for _, pat := range patterns {
 			if pat != "" && strings.Contains(strings.ToLower(tail), strings.ToLower(pat)) {
 				message = "usage limit: " + message
@@ -429,6 +431,14 @@
 		doctorDiag(&r, "C14", e.Error(), false)
 		return r
 	}
+	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id"} {
+		if len(maps[name]) == 0 {
+			doctorDiag(&r, "C14", "canary "+name+" missing", false)
+		}
+	}
+	if len(maps["baseline"]) == 0 {
+		return r
+	}
 	modText := ""
 	if b, e := os.ReadFile(filepath.Join(scratch, "go.mod")); e == nil {
 		for _, l := range strings.Split(string(b), "\n") {
@@ -519,6 +529,9 @@
 			vars["other_dir"] = dirs[1]
 		}
 		for _, name := range []string{"undeclared", "internal_eval", "failing_eval", "missing_id"} {
+			if len(maps[name]) == 0 {
+				continue
+			}
 			if bh == nil && (name == "undeclared" || name == "internal_eval") {
 				continue
 			}
diff --git a/packages/checks/internal/engine/history.go b/packages/checks/internal/engine/history.go
index c24b4442fabf69509ab2a74ec3856f9e375d48d5..49f7a17daff01e815fd6ed4866666d36f2ad19e6
--- a/packages/checks/internal/engine/history.go
+++ b/packages/checks/internal/engine/history.go
@@ -210,11 +210,15 @@
 			if e != nil {
 				return rep, e
 			}
-			before = append(before, s)
 			m, e := historyChanges(root, parent)
 			if e != nil {
 				return rep, e
 			}
+			s, e = projectedHistory(root, parent, s, m)
+			if e != nil {
+				return rep, e
+			}
+			before = append(before, s)
 			beforeChanges = append(beforeChanges, m)
 		}
 		var diffs []hchange

codex
I found two more spec edge cases: Doctor’s fixed canary names could collide with real modules, and History could label a merge’s applied phase as squashed even when a parent had impl. I’m fixing those before the final static check.
apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562/packages/checks/internal/engine/doctor.go
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
index 9aaeb601a62191f4c917ce09f650b743f9cd75b4..942ee041d90cf3f54006b8fd0ecade8327c8fbc2
--- a/DRILL_NOTES.md
+++ b/DRILL_NOTES.md
@@ -6,4 +6,5 @@
 
 ## Gaps
 
-No known behavioral gaps in the target checks module. The supplied workspace omits the Go manifest and dependency implementations, so the retained evals could not be run here. Native verification remains with the outside controller.
+- CHK-055 (checks): The supplied workspace has no `go.mod` or dependent internal implementations, so the retained Go evals cannot run here. The prior `PackageFailure`/coverage result cannot be reproduced or cleared in this input alone.
+- no spec (app, runner, workflow): The prior verification also reported `PackageFailure` for these modules. Their implementations are outside the authorized checks rebuild paths.
diff --git a/packages/checks/internal/engine/doctor.go b/packages/checks/internal/engine/doctor.go
index 6d69904707b2fb0e65938a4b6a72db7f41fce294..235be7af645acc2c0467f5b84ce777ef7c175338
--- a/packages/checks/internal/engine/doctor.go
+++ b/packages/checks/internal/engine/doctor.go
@@ -179,11 +179,6 @@
 		return nil, fmt.Errorf("canary canaries.yaml version %d", raw.Version)
 	}
 	out := map[string]map[string]string{"baseline": raw.Baseline, "undeclared": raw.Undeclared, "internal_eval": raw.InternalEval, "failing_eval": raw.FailingEval, "missing_id": raw.MissingID}
-	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id"} {
-		if len(out[name]) == 0 {
-			return nil, fmt.Errorf("canary %s missing", name)
-		}
-	}
 	return out, nil
 }
 func evalProbe(ctx context.Context, s project.Snapshot, m project.Module, hook any, bin string, timeout time.Duration) (model.Report, error) {
@@ -247,11 +242,9 @@
 	switch name {
 	case "undeclared", "internal_eval":
 		ds, e := boundaryProbe(ctx, s, m, boundHook, bin, timeout)
-		if e == nil {
-			for _, d := range ds {
-				if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
-					pass = true
-				}
+		for _, d := range ds {
+			if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
+				pass = true
 			}
 		}
 		detail = fmt.Sprintf("%v %v", ds, e)
@@ -332,16 +325,25 @@
 		limit = 3 * time.Minute
 	}
 	out, stderr, exit, runErr := run(ctx, scratch, argv, nil, limit)
-	logPath := filepath.Join(req.Root, ".fireweed", "runs", runID, "doctor-"+role+"-1.log")
-	_ = os.MkdirAll(filepath.Dir(logPath), 0755)
-	_ = os.WriteFile(logPath, []byte(out+stderr), 0644)
+	logRoot, pathErr := filepath.Abs(req.Root)
+	if pathErr != nil {
+		logRoot = req.Root
+	}
+	logPath := filepath.Join(logRoot, ".fireweed", "runs", runID, "doctor-"+role+"-1.log")
+	logErr := os.MkdirAll(filepath.Dir(logPath), 0755)
+	if logErr == nil {
+		logErr = os.WriteFile(logPath, []byte(out+stderr), 0644)
+	}
 	detail := "log: " + logPath
 	combined := out + stderr
 	tail := tail20(combined)
-	passed := runErr == nil && exit == 0 && strings.Contains(out, "FIREWEED_DOCTOR_OK")
+	passed := logErr == nil && runErr == nil && exit == 0 && strings.Contains(out, "FIREWEED_DOCTOR_OK")
 	doctorCheck(r, role, "", passed, detail)
 	if !passed {
 		message := fmt.Sprintf("%s probe command %v failed: %v exit %d\n%s", role, argv, runErr, exit, tail)
+		if logErr != nil {
+			message += "\nprobe log: " + logErr.Error()
+		}
 		for _, pat := range patterns {
 			if pat != "" && strings.Contains(strings.ToLower(tail), strings.ToLower(pat)) {
 				message = "usage limit: " + message
@@ -429,6 +431,14 @@
 		doctorDiag(&r, "C14", e.Error(), false)
 		return r
 	}
+	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id"} {
+		if len(maps[name]) == 0 {
+			doctorDiag(&r, "C14", "canary "+name+" missing", false)
+		}
+	}
+	if len(maps["baseline"]) == 0 {
+		return r
+	}
 	modText := ""
 	if b, e := os.ReadFile(filepath.Join(scratch, "go.mod")); e == nil {
 		for _, l := range strings.Split(string(b), "\n") {
@@ -438,7 +448,8 @@
 			}
 		}
 	}
-	names := []string{"doctorprobeone", "doctorprobetwo"}
+	stamp := time.Now().UnixNano()
+	names := []string{fmt.Sprintf("doctorprobeone%d", stamp), fmt.Sprintf("doctorprobetwo%d", stamp)}
 	dirs := []string{"packages/" + names[0], "packages/" + names[1]}
 	var canMods []project.Module
 	for i, name := range names {
@@ -519,6 +530,9 @@
 			vars["other_dir"] = dirs[1]
 		}
 		for _, name := range []string{"undeclared", "internal_eval", "failing_eval", "missing_id"} {
+			if len(maps[name]) == 0 {
+				continue
+			}
 			if bh == nil && (name == "undeclared" || name == "internal_eval") {
 				continue
 			}
diff --git a/packages/checks/internal/engine/history.go b/packages/checks/internal/engine/history.go
index c24b4442fabf69509ab2a74ec3856f9e375d48d5..49f7a17daff01e815fd6ed4866666d36f2ad19e6
--- a/packages/checks/internal/engine/history.go
+++ b/packages/checks/internal/engine/history.go
@@ -210,11 +210,15 @@
 			if e != nil {
 				return rep, e
 			}
-			before = append(before, s)
 			m, e := historyChanges(root, parent)
 			if e != nil {
 				return rep, e
 			}
+			s, e = projectedHistory(root, parent, s, m)
+			if e != nil {
+				return rep, e
+			}
+			before = append(before, s)
 			beforeChanges = append(beforeChanges, m)
 		}
 		var diffs []hchange

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562/packages/checks/internal/engine/history.go
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
index 9aaeb601a62191f4c917ce09f650b743f9cd75b4..942ee041d90cf3f54006b8fd0ecade8327c8fbc2
--- a/DRILL_NOTES.md
+++ b/DRILL_NOTES.md
@@ -6,4 +6,5 @@
 
 ## Gaps
 
-No known behavioral gaps in the target checks module. The supplied workspace omits the Go manifest and dependency implementations, so the retained evals could not be run here. Native verification remains with the outside controller.
+- CHK-055 (checks): The supplied workspace has no `go.mod` or dependent internal implementations, so the retained Go evals cannot run here. The prior `PackageFailure`/coverage result cannot be reproduced or cleared in this input alone.
+- no spec (app, runner, workflow): The prior verification also reported `PackageFailure` for these modules. Their implementations are outside the authorized checks rebuild paths.
diff --git a/packages/checks/internal/engine/doctor.go b/packages/checks/internal/engine/doctor.go
index 6d69904707b2fb0e65938a4b6a72db7f41fce294..235be7af645acc2c0467f5b84ce777ef7c175338
--- a/packages/checks/internal/engine/doctor.go
+++ b/packages/checks/internal/engine/doctor.go
@@ -179,11 +179,6 @@
 		return nil, fmt.Errorf("canary canaries.yaml version %d", raw.Version)
 	}
 	out := map[string]map[string]string{"baseline": raw.Baseline, "undeclared": raw.Undeclared, "internal_eval": raw.InternalEval, "failing_eval": raw.FailingEval, "missing_id": raw.MissingID}
-	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id"} {
-		if len(out[name]) == 0 {
-			return nil, fmt.Errorf("canary %s missing", name)
-		}
-	}
 	return out, nil
 }
 func evalProbe(ctx context.Context, s project.Snapshot, m project.Module, hook any, bin string, timeout time.Duration) (model.Report, error) {
@@ -247,11 +242,9 @@
 	switch name {
 	case "undeclared", "internal_eval":
 		ds, e := boundaryProbe(ctx, s, m, boundHook, bin, timeout)
-		if e == nil {
-			for _, d := range ds {
-				if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
-					pass = true
-				}
+		for _, d := range ds {
+			if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
+				pass = true
 			}
 		}
 		detail = fmt.Sprintf("%v %v", ds, e)
@@ -332,16 +325,25 @@
 		limit = 3 * time.Minute
 	}
 	out, stderr, exit, runErr := run(ctx, scratch, argv, nil, limit)
-	logPath := filepath.Join(req.Root, ".fireweed", "runs", runID, "doctor-"+role+"-1.log")
-	_ = os.MkdirAll(filepath.Dir(logPath), 0755)
-	_ = os.WriteFile(logPath, []byte(out+stderr), 0644)
+	logRoot, pathErr := filepath.Abs(req.Root)
+	if pathErr != nil {
+		logRoot = req.Root
+	}
+	logPath := filepath.Join(logRoot, ".fireweed", "runs", runID, "doctor-"+role+"-1.log")
+	logErr := os.MkdirAll(filepath.Dir(logPath), 0755)
+	if logErr == nil {
+		logErr = os.WriteFile(logPath, []byte(out+stderr), 0644)
+	}
 	detail := "log: " + logPath
 	combined := out + stderr
 	tail := tail20(combined)
-	passed := runErr == nil && exit == 0 && strings.Contains(out, "FIREWEED_DOCTOR_OK")
+	passed := logErr == nil && runErr == nil && exit == 0 && strings.Contains(out, "FIREWEED_DOCTOR_OK")
 	doctorCheck(r, role, "", passed, detail)
 	if !passed {
 		message := fmt.Sprintf("%s probe command %v failed: %v exit %d\n%s", role, argv, runErr, exit, tail)
+		if logErr != nil {
+			message += "\nprobe log: " + logErr.Error()
+		}
 		for _, pat := range patterns {
 			if pat != "" && strings.Contains(strings.ToLower(tail), strings.ToLower(pat)) {
 				message = "usage limit: " + message
@@ -429,6 +431,14 @@
 		doctorDiag(&r, "C14", e.Error(), false)
 		return r
 	}
+	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id"} {
+		if len(maps[name]) == 0 {
+			doctorDiag(&r, "C14", "canary "+name+" missing", false)
+		}
+	}
+	if len(maps["baseline"]) == 0 {
+		return r
+	}
 	modText := ""
 	if b, e := os.ReadFile(filepath.Join(scratch, "go.mod")); e == nil {
 		for _, l := range strings.Split(string(b), "\n") {
@@ -438,7 +448,8 @@
 			}
 		}
 	}
-	names := []string{"doctorprobeone", "doctorprobetwo"}
+	stamp := time.Now().UnixNano()
+	names := []string{fmt.Sprintf("doctorprobeone%d", stamp), fmt.Sprintf("doctorprobetwo%d", stamp)}
 	dirs := []string{"packages/" + names[0], "packages/" + names[1]}
 	var canMods []project.Module
 	for i, name := range names {
@@ -519,6 +530,9 @@
 			vars["other_dir"] = dirs[1]
 		}
 		for _, name := range []string{"undeclared", "internal_eval", "failing_eval", "missing_id"} {
+			if len(maps[name]) == 0 {
+				continue
+			}
 			if bh == nil && (name == "undeclared" || name == "internal_eval") {
 				continue
 			}
diff --git a/packages/checks/internal/engine/history.go b/packages/checks/internal/engine/history.go
index c24b4442fabf69509ab2a74ec3856f9e375d48d5..7969fb59576a8b75f57690fb84132217366c1916
--- a/packages/checks/internal/engine/history.go
+++ b/packages/checks/internal/engine/history.go
@@ -210,11 +210,15 @@
 			if e != nil {
 				return rep, e
 			}
-			before = append(before, s)
 			m, e := historyChanges(root, parent)
 			if e != nil {
 				return rep, e
 			}
+			s, e = projectedHistory(root, parent, s, m)
+			if e != nil {
+				return rep, e
+			}
+			before = append(before, s)
 			beforeChanges = append(beforeChanges, m)
 		}
 		var diffs []hchange
@@ -265,10 +269,12 @@
 			}
 			oldPhase := ""
 			held := false
+			hasImplParent := false
 			for _, m := range beforeChanges {
 				if x, ok := m[p]; ok {
 					oldPhase = x
 					held = held || x == phase
+					hasImplParent = hasImplParent || x == "impl"
 				}
 				for old, x := range m {
 					if path.Base(old) == path.Base(p) {
@@ -276,6 +282,7 @@
 							oldPhase = x
 						}
 						held = held || x == phase
+						hasImplParent = hasImplParent || x == "impl"
 					}
 				}
 			}
@@ -284,9 +291,9 @@
 			}
 			worker = true
 			if phase == "applied" {
-				if oldPhase == "impl" && len(parents) == 1 {
+				if hasImplParent && len(parents) == 1 {
 					implApplied = true
-				} else {
+				} else if !hasImplParent {
 					rep.Squashed = append(rep.Squashed, model.SquashedChange{Commit: oid, Change: p})
 					squashed = true
 				}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562/packages/checks/internal/engine/history.go
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
index 9aaeb601a62191f4c917ce09f650b743f9cd75b4..942ee041d90cf3f54006b8fd0ecade8327c8fbc2
--- a/DRILL_NOTES.md
+++ b/DRILL_NOTES.md
@@ -6,4 +6,5 @@
 
 ## Gaps
 
-No known behavioral gaps in the target checks module. The supplied workspace omits the Go manifest and dependency implementations, so the retained evals could not be run here. Native verification remains with the outside controller.
+- CHK-055 (checks): The supplied workspace has no `go.mod` or dependent internal implementations, so the retained Go evals cannot run here. The prior `PackageFailure`/coverage result cannot be reproduced or cleared in this input alone.
+- no spec (app, runner, workflow): The prior verification also reported `PackageFailure` for these modules. Their implementations are outside the authorized checks rebuild paths.
diff --git a/packages/checks/internal/engine/doctor.go b/packages/checks/internal/engine/doctor.go
index 6d69904707b2fb0e65938a4b6a72db7f41fce294..235be7af645acc2c0467f5b84ce777ef7c175338
--- a/packages/checks/internal/engine/doctor.go
+++ b/packages/checks/internal/engine/doctor.go
@@ -179,11 +179,6 @@
 		return nil, fmt.Errorf("canary canaries.yaml version %d", raw.Version)
 	}
 	out := map[string]map[string]string{"baseline": raw.Baseline, "undeclared": raw.Undeclared, "internal_eval": raw.InternalEval, "failing_eval": raw.FailingEval, "missing_id": raw.MissingID}
-	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id"} {
-		if len(out[name]) == 0 {
-			return nil, fmt.Errorf("canary %s missing", name)
-		}
-	}
 	return out, nil
 }
 func evalProbe(ctx context.Context, s project.Snapshot, m project.Module, hook any, bin string, timeout time.Duration) (model.Report, error) {
@@ -247,11 +242,9 @@
 	switch name {
 	case "undeclared", "internal_eval":
 		ds, e := boundaryProbe(ctx, s, m, boundHook, bin, timeout)
-		if e == nil {
-			for _, d := range ds {
-				if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
-					pass = true
-				}
+		for _, d := range ds {
+			if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
+				pass = true
 			}
 		}
 		detail = fmt.Sprintf("%v %v", ds, e)
@@ -332,16 +325,25 @@
 		limit = 3 * time.Minute
 	}
 	out, stderr, exit, runErr := run(ctx, scratch, argv, nil, limit)
-	logPath := filepath.Join(req.Root, ".fireweed", "runs", runID, "doctor-"+role+"-1.log")
-	_ = os.MkdirAll(filepath.Dir(logPath), 0755)
-	_ = os.WriteFile(logPath, []byte(out+stderr), 0644)
+	logRoot, pathErr := filepath.Abs(req.Root)
+	if pathErr != nil {
+		logRoot = req.Root
+	}
+	logPath := filepath.Join(logRoot, ".fireweed", "runs", runID, "doctor-"+role+"-1.log")
+	logErr := os.MkdirAll(filepath.Dir(logPath), 0755)
+	if logErr == nil {
+		logErr = os.WriteFile(logPath, []byte(out+stderr), 0644)
+	}
 	detail := "log: " + logPath
 	combined := out + stderr
 	tail := tail20(combined)
-	passed := runErr == nil && exit == 0 && strings.Contains(out, "FIREWEED_DOCTOR_OK")
+	passed := logErr == nil && runErr == nil && exit == 0 && strings.Contains(out, "FIREWEED_DOCTOR_OK")
 	doctorCheck(r, role, "", passed, detail)
 	if !passed {
 		message := fmt.Sprintf("%s probe command %v failed: %v exit %d\n%s", role, argv, runErr, exit, tail)
+		if logErr != nil {
+			message += "\nprobe log: " + logErr.Error()
+		}
 		for _, pat := range patterns {
 			if pat != "" && strings.Contains(strings.ToLower(tail), strings.ToLower(pat)) {
 				message = "usage limit: " + message
@@ -429,6 +431,14 @@
 		doctorDiag(&r, "C14", e.Error(), false)
 		return r
 	}
+	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id"} {
+		if len(maps[name]) == 0 {
+			doctorDiag(&r, "C14", "canary "+name+" missing", false)
+		}
+	}
+	if len(maps["baseline"]) == 0 {
+		return r
+	}
 	modText := ""
 	if b, e := os.ReadFile(filepath.Join(scratch, "go.mod")); e == nil {
 		for _, l := range strings.Split(string(b), "\n") {
@@ -438,7 +448,8 @@
 			}
 		}
 	}
-	names := []string{"doctorprobeone", "doctorprobetwo"}
+	stamp := time.Now().UnixNano()
+	names := []string{fmt.Sprintf("doctorprobeone%d", stamp), fmt.Sprintf("doctorprobetwo%d", stamp)}
 	dirs := []string{"packages/" + names[0], "packages/" + names[1]}
 	var canMods []project.Module
 	for i, name := range names {
@@ -519,6 +530,9 @@
 			vars["other_dir"] = dirs[1]
 		}
 		for _, name := range []string{"undeclared", "internal_eval", "failing_eval", "missing_id"} {
+			if len(maps[name]) == 0 {
+				continue
+			}
 			if bh == nil && (name == "undeclared" || name == "internal_eval") {
 				continue
 			}
diff --git a/packages/checks/internal/engine/history.go b/packages/checks/internal/engine/history.go
index c24b4442fabf69509ab2a74ec3856f9e375d48d5..6b0c8482e418df9307ad6b003fa2e67443ad2bc6
--- a/packages/checks/internal/engine/history.go
+++ b/packages/checks/internal/engine/history.go
@@ -210,11 +210,15 @@
 			if e != nil {
 				return rep, e
 			}
-			before = append(before, s)
 			m, e := historyChanges(root, parent)
 			if e != nil {
 				return rep, e
 			}
+			s, e = projectedHistory(root, parent, s, m)
+			if e != nil {
+				return rep, e
+			}
+			before = append(before, s)
 			beforeChanges = append(beforeChanges, m)
 		}
 		var diffs []hchange
@@ -263,19 +267,17 @@
 			if phase == "" {
 				continue
 			}
-			oldPhase := ""
 			held := false
+			hasImplParent := false
 			for _, m := range beforeChanges {
 				if x, ok := m[p]; ok {
-					oldPhase = x
 					held = held || x == phase
+					hasImplParent = hasImplParent || x == "impl"
 				}
 				for old, x := range m {
 					if path.Base(old) == path.Base(p) {
-						if oldPhase == "" {
-							oldPhase = x
-						}
 						held = held || x == phase
+						hasImplParent = hasImplParent || x == "impl"
 					}
 				}
 			}
@@ -284,9 +286,9 @@
 			}
 			worker = true
 			if phase == "applied" {
-				if oldPhase == "impl" && len(parents) == 1 {
+				if hasImplParent && len(parents) == 1 {
 					implApplied = true
-				} else {
+				} else if !hasImplParent {
 					rep.Squashed = append(rep.Squashed, model.SquashedChange{Commit: oid, Change: p})
 					squashed = true
 				}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562/packages/checks/internal/engine/core.go
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
index 9aaeb601a62191f4c917ce09f650b743f9cd75b4..942ee041d90cf3f54006b8fd0ecade8327c8fbc2
--- a/DRILL_NOTES.md
+++ b/DRILL_NOTES.md
@@ -6,4 +6,5 @@
 
 ## Gaps
 
-No known behavioral gaps in the target checks module. The supplied workspace omits the Go manifest and dependency implementations, so the retained evals could not be run here. Native verification remains with the outside controller.
+- CHK-055 (checks): The supplied workspace has no `go.mod` or dependent internal implementations, so the retained Go evals cannot run here. The prior `PackageFailure`/coverage result cannot be reproduced or cleared in this input alone.
+- no spec (app, runner, workflow): The prior verification also reported `PackageFailure` for these modules. Their implementations are outside the authorized checks rebuild paths.
diff --git a/packages/checks/internal/engine/core.go b/packages/checks/internal/engine/core.go
index d7dd034c8404afe0ab861d7d8b90f032df207b71..f71fa3603ee3f48c2517380c534610bdb3032739
--- a/packages/checks/internal/engine/core.go
+++ b/packages/checks/internal/engine/core.go
@@ -160,7 +160,7 @@
 	}
 	known := map[string]bool{}
 	for _, s := range specs {
-		known[strings.ToUpper(s.ID)] = true
+		known[strings.ReplaceAll(strings.ReplaceAll(strings.ToUpper(s.ID), "-", ""), "_", "")] = true
 	}
 	var report model.Report
 	var walk func(junitNode) error
@@ -179,25 +179,26 @@
 					c.IDs = append(c.IDs, s.ID)
 				}
 			}
-			if len(c.IDs) == 0 {
-				for _, id := range unknownID.FindAllString(n.Name+" "+n.Classname, -1) {
-					normalized := strings.ToUpper(strings.ReplaceAll(id, "_", "-"))
-					if !known[normalized] {
-						candidate := strings.ContainsAny(id, "-_")
-						if !candidate {
-							for _, spec := range specs {
-								prefix := strings.SplitN(spec.ID, "-", 2)[0]
-								if strings.HasPrefix(strings.ToUpper(id), strings.ToUpper(prefix)) {
-									candidate = true
-									break
-								}
-							}
-						}
-						if candidate {
-							return fmt.Errorf("testcase %q references unknown spec ID %q", n.Name, id)
+			for _, id := range unknownID.FindAllString(n.Name+" "+n.Classname, -1) {
+				normalized := strings.ReplaceAll(strings.ReplaceAll(strings.ToUpper(id), "-", ""), "_", "")
+				if known[normalized] {
+					continue
+				}
+				candidate := strings.ContainsAny(id, "-_")
+				if !candidate {
+					for _, spec := range specs {
+						prefix := strings.SplitN(spec.ID, "-", 2)[0]
+						if strings.HasPrefix(strings.ToUpper(id), strings.ToUpper(prefix)) {
+							candidate = true
+							break
 						}
 					}
 				}
+				if candidate {
+					return fmt.Errorf("testcase %q references unknown spec ID %q", n.Name, id)
+				}
+			}
+			if len(c.IDs) == 0 {
 				return fmt.Errorf("testcase %q has no known spec ID", n.Name)
 			}
 			var failures []string
diff --git a/packages/checks/internal/engine/doctor.go b/packages/checks/internal/engine/doctor.go
index 6d69904707b2fb0e65938a4b6a72db7f41fce294..235be7af645acc2c0467f5b84ce777ef7c175338
--- a/packages/checks/internal/engine/doctor.go
+++ b/packages/checks/internal/engine/doctor.go
@@ -179,11 +179,6 @@
 		return nil, fmt.Errorf("canary canaries.yaml version %d", raw.Version)
 	}
 	out := map[string]map[string]string{"baseline": raw.Baseline, "undeclared": raw.Undeclared, "internal_eval": raw.InternalEval, "failing_eval": raw.FailingEval, "missing_id": raw.MissingID}
-	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id"} {
-		if len(out[name]) == 0 {
-			return nil, fmt.Errorf("canary %s missing", name)
-		}
-	}
 	return out, nil
 }
 func evalProbe(ctx context.Context, s project.Snapshot, m project.Module, hook any, bin string, timeout time.Duration) (model.Report, error) {
@@ -247,11 +242,9 @@
 	switch name {
 	case "undeclared", "internal_eval":
 		ds, e := boundaryProbe(ctx, s, m, boundHook, bin, timeout)
-		if e == nil {
-			for _, d := range ds {
-				if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
-					pass = true
-				}
+		for _, d := range ds {
+			if d.Code == "X01" && strings.HasPrefix(d.Message, m.Name+":") && (name != "internal_eval" || strings.Contains(strings.ToLower(d.Message), "interface")) {
+				pass = true
 			}
 		}
 		detail = fmt.Sprintf("%v %v", ds, e)
@@ -332,16 +325,25 @@
 		limit = 3 * time.Minute
 	}
 	out, stderr, exit, runErr := run(ctx, scratch, argv, nil, limit)
-	logPath := filepath.Join(req.Root, ".fireweed", "runs", runID, "doctor-"+role+"-1.log")
-	_ = os.MkdirAll(filepath.Dir(logPath), 0755)
-	_ = os.WriteFile(logPath, []byte(out+stderr), 0644)
+	logRoot, pathErr := filepath.Abs(req.Root)
+	if pathErr != nil {
+		logRoot = req.Root
+	}
+	logPath := filepath.Join(logRoot, ".fireweed", "runs", runID, "doctor-"+role+"-1.log")
+	logErr := os.MkdirAll(filepath.Dir(logPath), 0755)
+	if logErr == nil {
+		logErr = os.WriteFile(logPath, []byte(out+stderr), 0644)
+	}
 	detail := "log: " + logPath
 	combined := out + stderr
 	tail := tail20(combined)
-	passed := runErr == nil && exit == 0 && strings.Contains(out, "FIREWEED_DOCTOR_OK")
+	passed := logErr == nil && runErr == nil && exit == 0 && strings.Contains(out, "FIREWEED_DOCTOR_OK")
 	doctorCheck(r, role, "", passed, detail)
 	if !passed {
 		message := fmt.Sprintf("%s probe command %v failed: %v exit %d\n%s", role, argv, runErr, exit, tail)
+		if logErr != nil {
+			message += "\nprobe log: " + logErr.Error()
+		}
 		for _, pat := range patterns {
 			if pat != "" && strings.Contains(strings.ToLower(tail), strings.ToLower(pat)) {
 				message = "usage limit: " + message
@@ -429,6 +431,14 @@
 		doctorDiag(&r, "C14", e.Error(), false)
 		return r
 	}
+	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id"} {
+		if len(maps[name]) == 0 {
+			doctorDiag(&r, "C14", "canary "+name+" missing", false)
+		}
+	}
+	if len(maps["baseline"]) == 0 {
+		return r
+	}
 	modText := ""
 	if b, e := os.ReadFile(filepath.Join(scratch, "go.mod")); e == nil {
 		for _, l := range strings.Split(string(b), "\n") {
@@ -438,7 +448,8 @@
 			}
 		}
 	}
-	names := []string{"doctorprobeone", "doctorprobetwo"}
+	stamp := time.Now().UnixNano()
+	names := []string{fmt.Sprintf("doctorprobeone%d", stamp), fmt.Sprintf("doctorprobetwo%d", stamp)}
 	dirs := []string{"packages/" + names[0], "packages/" + names[1]}
 	var canMods []project.Module
 	for i, name := range names {
@@ -519,6 +530,9 @@
 			vars["other_dir"] = dirs[1]
 		}
 		for _, name := range []string{"undeclared", "internal_eval", "failing_eval", "missing_id"} {
+			if len(maps[name]) == 0 {
+				continue
+			}
 			if bh == nil && (name == "undeclared" || name == "internal_eval") {
 				continue
 			}
diff --git a/packages/checks/internal/engine/history.go b/packages/checks/internal/engine/history.go
index c24b4442fabf69509ab2a74ec3856f9e375d48d5..6b0c8482e418df9307ad6b003fa2e67443ad2bc6
--- a/packages/checks/internal/engine/history.go
+++ b/packages/checks/internal/engine/history.go
@@ -210,11 +210,15 @@
 			if e != nil {
 				return rep, e
 			}
-			before = append(before, s)
 			m, e := historyChanges(root, parent)
 			if e != nil {
 				return rep, e
 			}
+			s, e = projectedHistory(root, parent, s, m)
+			if e != nil {
+				return rep, e
+			}
+			before = append(before, s)
 			beforeChanges = append(beforeChanges, m)
 		}
 		var diffs []hchange
@@ -263,19 +267,17 @@
 			if phase == "" {
 				continue
 			}
-			oldPhase := ""
 			held := false
+			hasImplParent := false
 			for _, m := range beforeChanges {
 				if x, ok := m[p]; ok {
-					oldPhase = x
 					held = held || x == phase
+					hasImplParent = hasImplParent || x == "impl"
 				}
 				for old, x := range m {
 					if path.Base(old) == path.Base(p) {
-						if oldPhase == "" {
-							oldPhase = x
-						}
 						held = held || x == phase
+						hasImplParent = hasImplParent || x == "impl"
 					}
 				}
 			}
@@ -284,9 +286,9 @@
 			}
 			worker = true
 			if phase == "applied" {
-				if oldPhase == "impl" && len(parents) == 1 {
+				if hasImplParent && len(parents) == 1 {
 					implApplied = true
-				} else {
+				} else if !hasImplParent {
 					rep.Squashed = append(rep.Squashed, model.SquashedChange{Commit: oid, Change: p})
 					squashed = true
 				}

exec
/bin/zsh -lc 'gofmt -w packages/checks/internal/engine/*.go && gofmt -l packages/checks/internal/engine/*.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
exec
/bin/zsh -lc "git status --short; rg -n 'case \"undeclared\"|hasImplParent|canary .* missing|logRoot|names :=|known :=' packages/checks/internal/engine" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562
 succeeded in 547ms:
 succeeded in 565ms:
?? DRILL_NOTES.md
?? packages/checks/internal/
packages/checks/internal/engine/doctor.go:243:	case "undeclared", "internal_eval":
packages/checks/internal/engine/doctor.go:328:	logRoot, pathErr := filepath.Abs(req.Root)
packages/checks/internal/engine/doctor.go:330:		logRoot = req.Root
packages/checks/internal/engine/doctor.go:332:	logPath := filepath.Join(logRoot, ".fireweed", "runs", runID, "doctor-"+role+"-1.log")
packages/checks/internal/engine/doctor.go:436:			doctorDiag(&r, "C14", "canary "+name+" missing", false)
packages/checks/internal/engine/doctor.go:452:	names := []string{fmt.Sprintf("doctorprobeone%d", stamp), fmt.Sprintf("doctorprobetwo%d", stamp)}
packages/checks/internal/engine/history.go:271:			hasImplParent := false
packages/checks/internal/engine/history.go:275:					hasImplParent = hasImplParent || x == "impl"
packages/checks/internal/engine/history.go:280:						hasImplParent = hasImplParent || x == "impl"
packages/checks/internal/engine/history.go:289:				if hasImplParent && len(parents) == 1 {
packages/checks/internal/engine/history.go:291:				} else if !hasImplParent {
packages/checks/internal/engine/core.go:161:	known := map[string]bool{}

codex
The checks implementation is in the target disposable paths. I fixed Doctor’s boundary canary detection, made canary names unique, tightened JUnit unknown-ID handling, and corrected History’s phase and zone checks. I also updated [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-371735562/DRILL_NOTES.md).

`gofmt` passed. The retained Go evals could not run because this supplied workspace has no `go.mod` or dependent internal implementations, so the prior `PackageFailure` and coverage failures remain unverified. No commit was made.
tokens used
77,061


## Verification failures
checks eval: C11: testcase "PackageFailure" has no known spec ID
C10: no coverage for CHK-001
app eval: C11: testcase "PackageFailure" has no known spec ID
C10: no coverage for APP-001
runner eval: C11: testcase "PackageFailure" has no known spec ID
C10: no coverage for RUN-001
workflow eval: C11: testcase "PackageFailure" has no known spec ID
C10: no coverage for FLOW-001

## Preserved prior review and candidate validation

The complete prior failed report and reviewed recovery evidence remain committed
at `5fb2126147a64e28576d31d4f08d524ca08bed0e:.fireweed/drills/checks.md`.
The earlier failed rebuilt ref remains
`refs/fireweed/drills/checks/20261105T000910.693992000`
(`83c4f87e645303d5eec76b3a1f236c2885d31a46`). The following CK1 conclusion and
recovery validation checkpoint are copied from that committed evidence. They
are historical; the current additional-drill verdict and owner checkpoint follow.

### Original CK1 review conclusion

CK1's focused review of the existing checks spec, Decisions, public interfaces
and evals found no necessary product correction. The retained evals cover genuine
failed/errored fail-first evidence versus compile/collection/skip failures, phase
zones, normally exited boundary suppression versus launch/cancellation/timeout
errors, doctor canaries and recovery, and merge/squash history rules. No scoped
change slug or lifecycle archive was needed; no product implementation, spec,
Decision, interface or eval was changed on main.

### Audited recovery and CK2 validation checkpoint (historical)

The owner authorized retrying the existing workflow audit with the same model
once capacity returned, completing recovery validation and stopping before any
additional drill. Bounded pinned-controller doctor probes passed both worker and
auditor in separate ephemeral sessions with gpt-6-sol and high effort. The
original capacity-failure trace remains retained; no model substitution or
fallback occurred.

The pinned controller resumed `0171-boundary-cause-eval` at audit in its original
run `20261105T021206-ef462a22e9bb7d1fa7c39456`, preserving the already-passing evals
Provenance and audit_rounds 0. The separate restricted audit passed FLOW-062 with
no findings. The implementation worker edited only the change Notes and passed
full Verify in 762.523 seconds; controller-owned implementation verification
passed 642 cases and archived the change at `47f4a63a6d37781dc4c8f5e403c715a05816eb51`.
`0170-doctor-baseline-scope` remains archived at
`ed5b02a41de4bb2c00f295bae1832b44480202a2`, with its prior passing audit and full
verification. No replacement change/run or counter reset was performed.

Final source review confirms only the approved CHK-055 clarification and the
FLOW-062 unlaunchable-boundary assertion/public setup changed, alongside archived
changes and protected evidence. No runtime implementation, public interface,
hook, recipe, configuration or other eval assertion was edited. The failed
reconstruction was not imported. The original CK1 review and unaffected earlier
proofs remain retained.

A clean fresh checkout of `47f4a63a6d37781dc4c8f5e403c715a05816eb51` built an independent NEW 0.1.0 candidate
with Go 1.27.2, darwin/arm64 and vcs.modified=false. Its executable SHA-256 is
`59b60b9975f06b6fee0726f16b9690197f5057b793a1dea4d634af85c6ed6723`. It remained a product subject, separate from the unchanged controller
pin. All candidate CLI commands selected that executable explicitly; doctor
setup copied its invoking subject into scratch. No controller promotion occurred.

With GOCACHE=/tmp/fireweed-rebuild-cache and serial package execution through
GOFLAGS=-p=1, sequential validation passed without corrections or retries:

| Validation | Exit | Seconds |
| --- | --- | --- |
| Clean candidate build | 0 | 1.82 |
| go test ./... | 0 | 985.93 |
| go test -race ./... | 0 | 1076.85 |
| go vet ./... | 0 | 1.30 |
| Candidate check --json | 0 | 0.47 |
| Candidate doctor --no-agent --json | 0 | 10.22 |
| Candidate check-commits, explicit range | 0 | 1387.52 |
| Candidate ci, explicit range, hook doctor enabled | 0 | 2477.01 |

Both history commands used the same exact nonempty range, containing 1,145 commits:

```text
be35552b706aa3000ff8f11306d129ff1d701b85..47f4a63a6d37781dc4c8f5e403c715a05816eb51
```

Candidate check, doctor, check-commits and CI returned no diagnostics. Standalone
doctor passed all four canaries and recovery checks. CI enforced passing cases
and coverage for app, assets, checks, contracts, gitstore, project, runner and
workflow, plus boundaries and hook doctor; it did not run real agents. Applicable
valid caches and unaffected historical evidence were reused without omitting
required phase checks. Runtime validation was on darwin/arm64.

## Reviewed additional checks drill and owner checkpoint

The owner explicitly authorized one additional independent native checks drill
on the validated recovery source, followed by CK4 review if it passed. Exactly
one additional `drill checks --json` invocation ran under the unchanged pinned
controller. Its two configured attempts were exhausted; both prescribed real
workers returned exit 0, while final enforced verification failed. This is a
new separately authorized proof, not a restart or reset of the earlier failed
proof. No further drill, counter reset or fallback was performed.

The isolated caller started from the validated candidate source recorded above.
Protected setup commit `66a37629cedaf9fcdedc5b3462cc3887c5fabb81` changed only
SYSTEM.md reconstruction instructions and agent.wrap preflight gating. It
retained the original global product conventions without supplying old source
or proof hints. Module contracts, evals, interfaces, manifests, hooks and graph
were unchanged. Static check passed before invocation. The wrapper enforced the
exact prescribed Codex argv/model/high effort and put the input tree's own
.fireweed/bin first on PATH with the prescribed cache. The opaque invoker and
its caller/verifier copies matched the pinned executable throughout.

Before each real worker, the outside controller inspected the actual command,
prompt, input Git and separate verifier. Initial input contained exactly the
21 approved retained files: checks SPEC/Decisions/interfaces/evals, SYSTEM and
the declared contracts/gitstore/assets/project public and model interfaces.
It had one fresh input-only commit, no remotes, alternates or unreachable old
objects, and no old target code, dependency/dependent internals, manifests,
hooks, binaries, saved attempts or bootstrap. The verifier retained its full
manifests, protected hooks and real dependency/dependent implementations; seven
original target disposable paths had been removed before the first worker.
The retry retained only the six fresh Go files and notes from this same proof,
with the approved retained inputs and one input-only Git commit unchanged.

All 393 retained verifier files and all approved input bytes matched the frozen
setup before both workers and during verification. Native acceptance and retained
proof-tree comparisons confirmed those bytes remained unchanged. Every command,
complete shell script and incremental patch from both workers was reviewed.
The first trace had 45 commands; repeated cumulative patch snapshots were
deduplicated and formatting normalized without omitting semantic revisions.
The second trace had 29 commands and three small patches. Final file bytes were
matched against those reviewed patches. Reads stayed within supplied inputs and
fresh current-run code; writes stayed within target disposable paths and notes.
No worker invoked a lifecycle mutator, committed, queried a product binary as
an oracle, read old/deleted implementation or dependency internals, or invoked
another agent. Both separate ephemeral sessions used gpt-6-sol/high.

First verification exposed unsupported Git --path-format handling in the fresh
doctor scratch setup, plus a generic boundary error that lost its launch cause.
After review of all native feedback and renewed isolation checks, only the
second configured invocation was released. It corrected those two target paths
and notes, then returned. Full native own and all-transitive-dependent checking
still failed; no verifier gate or retained assertion was waived.

Final native results, with every current ID covered by non-skipped cases:

| Module | Passing cases | Failed cases | Skipped | Current IDs covered |
| --- | --- | --- | --- | --- |
| checks | 108 | 6 | 0 | 83/83 |
| app | 320 | 3 | 0 | 155/155 |
| runner | 159 | 0 | 0 | 67/67 |
| workflow | 159 | 1 | 0 | 132/132 |

Failed parent cases remain counted alongside their failed subcases. Coverage is
not passing proof: all ten failed enforced cases above remain acceptance blockers.
All native report fields, final cases and failure details were reviewed.

The unresolved behavior failures are the Go missing_id canary not being
recognized, proceeding to canaries after a failed global boundary baseline,
and workflow returning a doubled X02 prefix for the supplied launch failure.
The app doctor/CI-doctor failures also report the missed missing_id canary.
Two checks assertions additionally need owner review: CHK-057's substring check
matches a successful `recovery internal_eval` as a passed `internal_eval` probe,
although the actual canary failed and C14 was present; CHK-060's boundary subcase
searches for `boundary`, which does not match the emitted `boundaries hook`.
Those are report-selection/wording questions, not authority to weaken canary,
C14, override or cause-preservation assertions. No recovery was authored from
these findings. This failed reconstruction does not itself establish a necessary
main implementation change.

No contradicted Decisions or known behavioral gaps were reported by the workers.
The notes' local build limitation remains the intended input/verifier separation,
not a product gap or permission to obtain excluded inputs. The native verifier's
actual failures, rather than the workers' no-gap statement, determine acceptance.

The native retained ref is
`refs/fireweed/drills/checks/20261105T051932.421753000`
(`70e5ec6c117757b403e752a2fdfb1bf377b8fe74`). It was retained locally in the main
repository with prior proof/attempt refs unchanged. No reconstructed source was
imported onto main. Native cleanup removed input and verifier trees and their
registration. The task-owned invoking clone/gate were removed only after its
report and ref were retained. Main and validated candidate checkouts are clean;
no active run/shared claim remains. The pin, explicit recovery backup and
unrelated worktree are preserved. Live status showed an empty queue, unapproved
2 and only the pre-existing B02 warnings for CHK-055, APP-103 and APP-104. Existing
assertions, bounds and the original 2026-11-07 22:51:05 UTC (14:51:05 PST) deadline
remain unchanged. No push, publication, tag, promotion or regenerate ran.

CK3 remains failed and CK4's passing-proof/completion gate is open. This additional
authorization is consumed. Stop for owner agreement on the unresolved failures
and report-assertion questions before scoped audited recovery; any further real
drill requires separate explicit authorization. Do not reset attempts, broaden
scope or retry automatically. Prior CK1/CK2 and unaffected evidence remain valid.


Author static check passed without diagnostics. `git diff --check` reports only
whitespace within the native captured output and diff context above, rather
than authored review text. Those trace bytes are preserved verbatim.

## Reviewed recovery after the additional failed proof

The owner approved a checks eval_fix for CHK-057 and CHK-060 and a CHK-043
contract clarification, followed by clean candidate validation and a stop
before another drill. Review identified two separate reconstruction defects.
Doctor matched the wrong unmapped-ID wording for the Go canary, and its global
boundary baseline considered only the throwaway module, missing the real
module's violation and proceeding to planting. CHK-059 and the clarified
CHK-055 continue to enforce both behaviors. The reconstructed source stayed on
the retained proof ref.

- `0172-doctor-probe-assertions` applied at
  `35b8143`. It distinguishes the actual internal_eval probe from recovery.
  The first restricted audit's extra blank line outside Audit failed the
  import check. On owner authorization, the original run
  `20261105T215537-14c3ae392f33f060ba32aaf5` retried the independent audit with
  the same Codex invocation settings. A temporary launcher appended
  formatting-only guidance to the prompt. The native importer accepted the
  submission, the original audit log remained an exact prefix of the log, and
  the audit round count stayed zero.
- `0173-boundary-override-diagnostic` applied at
  `6df7f31`. Review found that 0172 still required the word override and
  rejected the second proof's alpha boundaries hook diagnostic. This narrow
  follow-up accepts both boundary hook wordings and preserves the error-level
  C14, module attribution and canary checks. The final predicate accepted the
  recorded and current diagnostics and rejected five controls covering the
  global hook, wrong module, wrong canary, warning severity and wrong code.
- `0174-boundary-hook-cause-contract` applied at
  `555ef08` as already_satisfied. Its existing public eval verifies the
  separate X02 diagnostic and the launch, cancellation and timeout causes in
  the returned error without an X02 prefix. FLOW-062 retains its exact cause
  and single-prefix assertion.

Each archived change records 114 passing checks evals, a separate restricted
audit with no findings and 756 passing implementation verification cases.
Only the checks spec, checks evals and archived changes differ from the source
before this recovery. The earlier CK1 review and unaffected evidence still
apply.

### CK2 candidate validation

Fresh clean source `555ef083c533aaa0c1be05af778a47157e03e448`, Go 1.27.2 on
darwin/arm64, embedded `vcs.modified=false`. The independently built candidate
SHA-256 is `0410472d48d722c425cf804be5e34219e09b777d83820cc8483a131a78e5b157`.
The pinned controller remained separate. The explicit nonempty history range
was `be35552b706aa3000ff8f11306d129ff1d701b85..555ef083c533aaa0c1be05af778a47157e03e448`
(1161 commits).

| Gate | Result | Seconds |
| --- | --- | ---: |
| go test ./... | Pass | 964.47 |
| go test -race ./... | Pass | 1136.92 |
| go vet ./... | Pass | 1.35 |
| Candidate check | Pass, no diagnostics | 0.45 |
| Candidate doctor --no-agent | Pass, all four canaries and recovery | 10.13 |
| Candidate check-commits, explicit range | Pass, no diagnostics | 1318.55 |
| Candidate ci, same range, hook doctor enabled | Pass, no diagnostics | 2329.27 |

The build exited successfully with a Go module stat-cache write warning from
the sandbox. Live status reported the existing CHK-055, APP-103 and APP-104
case-count warnings and a repeated-module warning for the three scoped checks
changes. Both checkouts were clean and all controller and candidate scratch
worktrees were cleaned up. No shared claims remained. Prior proof and attempt
refs and the unrelated worktree were verified against the pre-recovery state.
The controller pin, approval enforcement, configured bounds and original
deadline remained in place. The failed CK3 proof is still the latest native
result. Another real drill requires separate owner authorization and isolation
checks before CK4 can proceed.


## CK3 invocation stopped before worker release

The owner authorized one native invocation on validated recovery source
`555ef083c533aaa0c1be05af778a47157e03e448`. The clean isolated caller's
protected setup commit `f20db727d2aa58e11d66244cf57cbb2b7f397ca0` changed
only reconstruction instructions and the launch gate. The candidate identity
and controller pin matched CK2, and the caller's static check passed.

The first tool launch request timed out in automatic approval review before
creating a process. The approved retry started one native invocation. The gate
used the previously reviewed ten-minute manual inspection window. Both windows
expired before live input and verifier inspection completed. Neither release
marker was written, and the controller recorded two launcher exits. Neither
Codex worker ran. No reconstructed files, eval reports or rebuilt ref were
produced. The native controller removed the input and verifier. The prior
retained failed proofs remain intact. This is a launch stop and supplies no
new evidence about reconstruction behavior.

The approved invocation is spent. CK1 and CK2 evidence remains applicable.
CK3 has no passing proof and CK4 remains open. The controller pin, model and
effort configuration, configured attempts and audit bounds, prior proof and
attempt refs, unrelated worktree and original deadline were preserved.

Native stop report follows verbatim.

```text
---
date: 2026-11-06T04:37:33.857188Z
base_commit: f20db727d2aa58e11d66244cf57cbb2b7f397ca0
agent: codex
model: gpt-6-sol
result: fail
---

## Own evals

## Dependent evals

## Contradicted Decisions

## Gaps

## Agent observations

### Attempt 1 (exit 1)
Reconstruction held for outside-controller input/verifier inspection.

Preflight gate expired; real agent was not invoked


### Attempt 2 (exit 1)
Reconstruction held for outside-controller input/verifier inspection.

Preflight gate expired; real agent was not invoked


## Verification failures
reconstruction agent: exit status 1
```


## Reviewed fresh CK3 result and owner checkpoint

The owner authorized one new native checks drill on the CK2 recovery source
`555ef083c533aaa0c1be05af778a47157e03e448` after the earlier launch stop.
The isolated caller's only setup differences were reconstruction instructions
and the external launch gate. The earlier native report remains available in
commit `6b09bd5`. The current native report above is retained verbatim.

The first release was automatic only after live inspection succeeded. The
input contained 21 approved files in a single input-only Git commit, without
remotes, alternates or unreachable old objects. The detached verifier retained
396 frozen files, including real dependency and dependent implementations,
manifests and hooks, and excluded the seven original disposable paths. The
invoking executable and caller and verifier copies matched the pin. Worker
PATH began with the input tree's .fireweed/bin. Both invocations used the
prescribed gpt-6-sol/high Codex command in independent ephemeral sessions.

All 26 first-worker commands and 95 patch snapshots were reviewed. The 22
distinct snapshots were normalized for formatting, with 1974 lines of
incremental patch review. First-attempt output matched the reviewed bytes in
the input and verifier. The full native feedback was reviewed before the
second release. It reported missing real Go internal_eval and missing_id
canaries in checks and app, a runner cancellation timing failure, and the
workflow boundary-cause failure. The second live preflight confirmed the
same frozen retained inputs and verifier, and only the fresh first-worker
files and notes and native feedback were added to the worker's context.

All 19 second-worker commands and two distinct patches were reviewed. The
worker changed only doctor canary classification and notes. Both traces stayed
within supplied inputs and fresh output, with no access to prior code, binaries
or external workspaces and no edits to retained files or lifecycle mutations.
Both worker commands exited zero. The final reviewed files and retained
verifier bytes matched every blob in the native proof tree.

| Module | Passing | Failed | Skipped | Current IDs covered |
| --- | ---: | ---: | ---: | ---: |
| checks | 114 | 0 | 0 | 83/83 |
| app | 323 | 0 | 0 | 155/155 |
| runner | 159 | 0 | 0 | 67/67 |
| workflow | 159 | 1 | 0 | 132/132 |

FLOW-062 is the sole final enforced failure. The reconstructed Boundaries
returned the unprefixed launch error and a separate X02 diagnostic whose
message added hook context. Lifecycle Verify prepended the diagnostic
message and then appended the same launch detail again. The eval required
the original returned cause with one X02 prefix and rejected the duplicate
detail. This is the observed producer/consumer mismatch. The passing CHK-043
eval, all-module coverage and successful runner rerun within the second
configured attempt do not satisfy the remaining FLOW-062 gate.

The worker reported no contradicted Decisions. Its CHK-057 and CHK-059 notes
describe first-attempt misses that passed final verification. The absent
manifest is the intended worker input limitation. The runner cancellation
failure was observed on the first attempt and did not recur in the final
verification. The final workflow failure remains for owner review.

The native exit was 1, with the failed rebuilt ref
`refs/fireweed/drills/checks/20261106T072400.155233000` retained locally at
`6fe180c6de0851d6c5e531e68744043273e1e824`. Native input and verifier paths
and the linked-worktree registration were removed. Both main and the validated
candidate were clean before evidence authoring, with no live claim. The prior
proof and attempt refs and unrelated worktree match their original values.
The pin, model, counters, configured bounds and original deadline were
preserved. CK1 and CK2 evidence applies, and CK4 remains open. A further
invocation or product correction needs separate owner authorization.


## Reviewed FLOW-062 recovery and affected CK2 validation

The owner approved one audited workflow fix and affected CK2 validation,
followed by a stop before another drill. `0175-preserve-x02-boundary-cause`
applied at `0e2c84c638620df5a0393e95911eb7de9e00bce6` on held run
`20261106T085253-cf6c6656e37fecebf85adaeb`. The new public eval gets the actual
missing-hook cause from checks.Boundaries and requires Verify to return it
with one X02 prefix, while retaining the full diagnostics. The existing exact
unlaunchable-hook assertion and earlier evals are byte-unchanged. The
implementation uses the selected X02 diagnostic for the code and keeps the
wrapped checks error as the cause, with structured diagnostics returned
separately. Existing error precedence and X01 behavior are preserved.

The first eval worker attempted a prohibited phase advance. Its sandbox
denied the Git write and the worker restored the held change document. The
outside controller was stopped before it could commit the first draft's
hard-coded producer wording. HEAD, index and claim were checked, and the
first worker log was retained. A second prescribed gpt-6-sol/high worker
corrected the new assertion on the same held run within the configured eval
attempt bound. Pinned quick Verify passed and the outside owner advanced
using pinned Done. Its evals Provenance records the manual advancement and
161 cases, with the new case failing for the intended reason and the other
160 passing. The independent restricted audit passed with no findings and
zero recorded audit rounds. The temporary launcher added instructions on
Audit formatting and worker phase boundaries to the existing prompts, with
the configured Codex arguments and native enforcement intact. Full
implementation verification passed all 643 cases. The preserved original
worker log and the second worker log remain under the same run identity.

The fresh, independently built candidate had embedded source
`0e2c84c638620df5a0393e95911eb7de9e00bce6` and `vcs.modified=false`, built
with Go 1.27.2 on darwin/arm64. Its SHA-256 was
`cd55ff5ba0b8ab44b6098ee88e6dfc5bf3d63f939dd39023aaac0fafbe78c68e`.
The explicit history range was
`be35552b706aa3000ff8f11306d129ff1d701b85..0e2c84c638620df5a0393e95911eb7de9e00bce6`
(1169 commits). The separate pinned controller remained in place.

| Gate | Result | Seconds |
| --- | --- | ---: |
| go test ./... | Pass | 1048.72 |
| go test -race ./... | Pass on unchanged retry | 344.23 |
| go vet ./... | Pass | 1.30 |
| Candidate check | Pass, no diagnostics | 0.40 |
| Candidate doctor --no-agent | Pass, no diagnostics | 11.98 |
| Candidate Verify and Done through CLI fixtures | Pass | 5.41 |
| Candidate check-commits, explicit range | Pass, no diagnostics | 1377.58 |
| Candidate ci, same range, hook doctor enabled | Pass, no diagnostics | 2324.16 |

The initial race command exited 1 after 1143.60 seconds. The runner
RUN-007/RUN-009 cancellation subcase exceeded its prompt-stop timer, with no
data race reported. Runner and Git process code were unchanged. The same
full race command passed on the same clean candidate on one retry, with no
settings or bounds changed. The successful build emitted a sandbox warning
about a Go module stat-cache write.

The candidate CLI fixtures checked missing and unlaunchable hooks through
both Verify and Done. Each returned the original cause under X02 and retained
the structured diagnostics, with the fixture HEAD and held impl claim
unchanged on failure. Fixture claims were released after the checks. Main and
the candidate are clean, with no active claim or scratch worktree. The pin,
counters, prior proof and attempt refs, unrelated worktree and original
deadline were preserved. This validation does not advance CK3 or CK4. The
next independent drill needs separate owner authorization.


## Reviewed CK3 compilation stop and owner checkpoint

The owner authorized one native checks drill on validated recovery source
`0e2c84c638620df5a0393e95911eb7de9e00bce6`. The isolated caller's only setup
changes were the reconstruction instructions and rehearsed external gate.
The original CK1 review and affected CK2 validation still apply. The
previous native report remains in commit `08c58f2`. The current native report
above is retained verbatim.

The first gate passed live inspection before release. The worker input held
21 approved files in a single input-only Git commit, with no remotes,
alternates or reachable old history. The detached verifier retained 397
frozen files, including the real dependencies, dependents, manifests and
hooks, and excluded seven original disposable paths. The controller copies
matched the pin. Both invocations used the configured gpt-6-sol/high Codex
command in independent ephemeral sessions.

The first trace contained 51 commands, including the full shell heredocs
and incremental Python rewrites for four target Go files and notes. The
review covered each command and write script. Output hashes matched between
the input and verifier, with all retained bytes intact. The first native
feedback reported PackageFailure and missing coverage for all four modules,
with empty stdout/stderr. The full feedback was reviewed before release of
the second attempt. Its live preflight confirmed the frozen input plus
only the fresh first-worker output, the unchanged verifier and the same
pin. The second trace contained 29 commands and nine distinct cumulative
patch snapshots. The review covered the commands and each patch, and the
final patches applied in reverse to the retained proof recovered the
first-attempt Go file hashes. All retained proof blobs matched the verifier
inventory. Both worker exits were zero. Neither worker crossed the input
boundary or changed a retained file.

The second native verification again reported PackageFailure and missing
coverage for checks, app, runner and workflow. The native report contains
no eval cases, so none of the four modules established passing coverage.
The worker reported no contradicted Decisions. Its note about missing
manifests describes the intended input isolation. The dependent package
failures alone do not identify defects in their implementations.

A compile-only diagnostic on an exact temporary archive of the final
retained proof, with the prescribed cache and GOPROXY disabled, produced:

```text
packages/checks/internal/engine/doctor.go:413:72: undefined: reason
packages/checks/internal/engine/core.go:11:2: "path/filepath" imported and not used
```

These are reconstruction compilation defects in the disposable checks
output. The diagnostic changed no proof bytes and was separate from the
native verification. The FLOW-062 recovery remains validated on the
candidate and was not reached in this drill's verifier.

The native exit was 1. The failed rebuilt ref
`refs/fireweed/drills/checks/20261106T155422.780454000` is retained locally at
`314015015ba0b7ee96d97f010e78455b742c1360`. Both live input and verifier paths
and their registrations were removed. Main and the validated candidate were
clean before evidence authoring, with no active claim. The controller pin,
model, counters, configured bounds, original deadline, prior proof and
attempt refs and unrelated worktree were preserved. The authorized
invocation is spent and CK4 remains open.
