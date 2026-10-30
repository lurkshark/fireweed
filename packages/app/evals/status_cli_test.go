package evals_test

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func statusArchive(t *testing.T, root, name, module, date string) {
	t.Helper()
	playbook, err := assets.Playbook("impl")
	if err != nil {
		t.Fatal(err)
	}
	appPut(t, root, "changes/archive/"+name+".md", fmt.Sprintf(`---
module: %s
bases: {}
run: fixture-%s
phase: applied
---

## Motivation
Earlier implementation change.

## Interface

## Added

## Modified

## Removed

## Decisions

## Provenance
- impl %s agent=manual fireweed=%s playbook=%x tests=1 failing=0 passing=1
`, module, name, date, assets.Version(), sha256.Sum256([]byte(playbook))))
}

func statusRichFixture(t *testing.T) (string, workflow.StatusReport) {
	t.Helper()
	root := appProjectFixture(t)
	appPut(t, root, "fireweed.yaml", strings.Replace(appFixtureConfig, "  preserve: [data/**]\n", "  preserve: [data/**]\n  protected: [guard/**]\n  append_only: [ledger/**]\n", 1))
	appPut(t, root, "packages/alpha/SPEC.md", `---
name: alpha
id_prefix: ALP
depends_on: []
boundaries_opt_out: fixture exception for alpha
---
# Alpha

## Specs

### ALP-001 [quarantine: intermittent fixture]
Alpha MUST work.

### ALP-002
Alpha MUST also work.
`)
	appPut(t, root, "packages/beta/SPEC.md", `---
name: beta
id_prefix: BET
depends_on: [alpha]
---
# Beta

## Specs

### BET-001
Beta MUST work.
`)
	appPut(t, root, "packages/gamma/SPEC.md", `---
name: gamma
id_prefix: GAM
depends_on: []
---
# Gamma

## Specs

### GAM-001 [quarantine: unavailable fixture service]
Gamma MUST work.
`)
	cliGit(t, root, "init", "-q", "-b", "main")
	cliCommit(t, root, "approved base")
	base := cliGit(t, root, "rev-parse", "HEAD")
	cliGit(t, root, "checkout", "-qb", "feature")
	listChange(t, root, "0001-blocked.md", "alpha", "phase: impl\nblocked: fixture review required\naudit_rounds: 2\n")
	listChange(t, root, "0002-dependent.md", "beta", "")
	listChange(t, root, "0003-independent.md", "gamma", "phase: audit\naudit_rounds: 1\n")
	for _, a := range []struct{ name, module, date string }{
		{"0004-beta-old", "beta", "2026-06-11T12:00:00Z"},
		{"0005-alpha-old", "alpha", "2026-06-12T12:00:00Z"},
		{"0006-alpha-new", "alpha", "2026-06-13T12:00:00Z"},
		{"0007-alpha-newest", "alpha", "2026-06-14T12:00:00Z"},
	} {
		statusArchive(t, root, a.name, a.module, a.date)
	}
	appPut(t, root, ".fireweed/drills/alpha.md", fmt.Sprintf("---\ndate: 2026-06-10T12:00:00Z\nbase_commit: %s\nagent: fixture\nmodel: fixture\nresult: fail\n---\n", base))
	snapshot := workflow.CostSnapshot{
		Run: "fixture-measured-run", Phase: "impl", RecordedAt: time.Date(2026, 6, 15, 12, 0, 0, 0, time.UTC),
		Modules: []workflow.ModuleCost{
			{Module: "alpha", Seconds: 1.25, CasesPerSpec: map[string]int{"ALP-001": 3, "ALP-002": 7}, EvalLines: 31, DisposableLines: 47},
			{Module: "beta", Seconds: 2.75, CasesPerSpec: map[string]int{"BET-001": 11}, EvalLines: 59, DisposableLines: 83},
		},
		Diagnostics: []project.Diagnostic{
			{Code: "B03", Message: "alpha: fixture eval budget exceeded; use fewer cases", Path: "packages/alpha/evals/weight_test.go", Warning: true},
			{Code: "B01", Message: "beta fixture time budget exceeded; review cost", Path: "packages/beta/evals/weight_test.go", Warning: true},
		},
	}
	b, err := json.Marshal(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	appPut(t, root, ".fireweed/runs/fixture-measured-run/impl/costs.json", string(b))
	archivePath := filepath.Join(root, "changes", "archive", "0005-alpha-old.md")
	archive, err := os.ReadFile(archivePath)
	if err != nil {
		t.Fatal(err)
	}
	appPut(t, root, "changes/archive/0005-alpha-old.md", strings.Replace(string(archive), "phase: applied", "phase: queued", 1))
	cliCommit(t, root, "first archive phase")
	appPut(t, root, "changes/archive/0005-alpha-old.md", string(archive))
	cliCommit(t, root, "status fixture")
	report, err := workflow.Status(root)
	if err != nil {
		t.Fatalf("public Status fixture: %v", err)
	}
	if len(report.Modules) != 3 || len(report.Queue) != 3 || len(report.Quarantines) != 2 || len(report.BoundaryOptOuts) != 1 || len(report.Costs) != 2 || len(report.PhaseCommits) == 0 || report.Unapproved == 0 {
		t.Fatalf("rich public Status fixture incomplete: %+v", report)
	}
	for _, module := range report.Modules {
		for _, check := range []struct {
			patterns []string
			want     string
		}{{module.Zones.Preserve, "data/**"}, {module.Zones.Protected, "guard/**"}, {module.Zones.AppendOnly, "ledger/**"}} {
			found := false
			for _, pattern := range check.patterns {
				found = found || pattern == check.want
			}
			if !found {
				t.Fatalf("rich public Status fixture module %s lacks resolved zone %q", module.Name, check.want)
			}
		}
	}
	if report.PhaseCommits["0005-alpha-old"] == report.PhaseCommits["0004-beta-old"] {
		t.Fatalf("phase-count fixture needs distinct associations: %+v", report.PhaseCommits)
	}
	measured := map[string]float64{}
	for _, c := range report.Costs {
		measured[c.Module] = c.Seconds
	}
	if measured["alpha"] != 1.25 || measured["beta"] != 2.75 {
		t.Fatalf("cost fixture lost recorded measurements: %+v", report.Costs)
	}
	if report.Queue[1].WaitingOn == "" || report.Queue[1].WaitingRule == "" || report.DriftSinceDrill["alpha"] == 0 {
		t.Fatalf("rich public Status fixture lacks waiting or drift: %+v", report)
	}
	for _, code := range []string{"B01", "B03", "P08"} {
		found := false
		for _, d := range report.Diagnostics {
			found = found || d.Code == code && d.Warning
		}
		if !found {
			t.Fatalf("rich public Status fixture lacks %s warning: %+v", code, report.Diagnostics)
		}
	}
	return root, report
}

func statusData(t *testing.T, output string) (map[string]json.RawMessage, []map[string]json.RawMessage) {
	t.Helper()
	env := decodeEnvelope(t, output)
	var data map[string]json.RawMessage
	if err := json.Unmarshal(env["data"], &data); err != nil || data == nil {
		t.Fatalf("status data=%s; want object: %v", env["data"], err)
	}
	want := []string{"modules", "queue", "quarantines", "boundary_opt_outs", "drift_since_drill", "unapproved", "phase_commits", "metrics"}
	if len(data) != len(want) {
		t.Errorf("status fields=%v; want %v", reflect.ValueOf(data).MapKeys(), want)
	}
	for _, key := range want {
		if len(data[key]) == 0 || string(data[key]) == "null" {
			t.Errorf("status field %s=%s; want non-null", key, data[key])
		}
	}
	var diagnostics []map[string]json.RawMessage
	if err := json.Unmarshal(env["diagnostics"], &diagnostics); err != nil || diagnostics == nil {
		t.Fatalf("diagnostics=%s; want array: %v", env["diagnostics"], err)
	}
	return data, diagnostics
}

func statusRows(t *testing.T, raw json.RawMessage, keys ...string) []map[string]json.RawMessage {
	t.Helper()
	var rows []map[string]json.RawMessage
	if err := json.Unmarshal(raw, &rows); err != nil || rows == nil {
		t.Fatalf("rows=%s; want array: %v", raw, err)
	}
	for i, row := range rows {
		if len(row) != len(keys) {
			t.Errorf("row %d fields=%v; want %v", i, row, keys)
		}
		for _, key := range keys {
			if len(row[key]) == 0 || string(row[key]) == "null" {
				t.Errorf("row %d field %s=%s", i, key, row[key])
			}
		}
	}
	return rows
}

func statusEqualField(t *testing.T, row map[string]json.RawMessage, key string, want any) {
	t.Helper()
	if values, ok := want.([]string); ok && values == nil {
		want = []string{}
	}
	b, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}
	var gotValue, wantValue any
	if err := json.Unmarshal(row[key], &gotValue); err != nil {
		t.Errorf("%s=%s: %v", key, row[key], err)
		return
	}
	if err := json.Unmarshal(b, &wantValue); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(gotValue, wantValue) {
		t.Errorf("%s=%s; want %s", key, row[key], b)
	}
}

func Test_APP_118_APP_119_APP_120_APP_121_APP_122_APP_123_APP_124_APP_125_StatusRichReport(t *testing.T) {
	root, report := statusRichFixture(t)
	nested := appNestedDir(t, root)
	status, plain, stderr := listRun(t, nested, "status")
	if status != 0 {
		t.Fatalf("plain status exit=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	last := -1
	for _, entry := range report.Queue {
		pos := strings.Index(plain, entry.Path)
		if pos <= last {
			t.Errorf("queue path %q at %d after %d in %q", entry.Path, pos, last, plain)
		}
		last = pos
		for _, field := range []string{entry.Slug, entry.Path, entry.Module, entry.Phase, entry.Blocked, entry.WaitingOn, entry.WaitingRule, strconv.Itoa(entry.AuditRounds)} {
			if field != "" && !strings.Contains(plain, field) {
				t.Errorf("plain status missing queue field %q: %q", field, plain)
			}
		}
	}
	for _, module := range report.Modules {
		for _, field := range []string{module.Name, module.IDPrefix, module.Dir} {
			if !strings.Contains(plain, field) {
				t.Errorf("plain status missing module field %q", field)
			}
		}
		for _, dep := range module.DependsOn {
			if !strings.Contains(plain, dep) {
				t.Errorf("plain status missing dependency %q", dep)
			}
		}
		for _, zones := range [][]string{module.Zones.Interface, module.Zones.Evals, module.Zones.Disposable, module.Zones.Preserve, module.Zones.Protected, module.Zones.AppendOnly} {
			for _, zone := range zones {
				if !strings.Contains(plain, zone) {
					t.Errorf("plain status missing resolved zone %q", zone)
				}
			}
		}
	}
	for _, q := range report.Quarantines {
		for _, field := range []string{q.Module, q.ID, q.Reason} {
			if !strings.Contains(plain, field) {
				t.Errorf("plain status missing quarantine %q", field)
			}
		}
	}
	for _, o := range report.BoundaryOptOuts {
		if !strings.Contains(plain, o.Module) || !strings.Contains(plain, o.Reason) {
			t.Errorf("plain status missing opt-out %+v", o)
		}
	}
	if !strings.Contains(plain, strconv.Itoa(report.Unapproved)) {
		t.Errorf("plain status missing unapproved=%d", report.Unapproved)
	}
	for id, count := range report.DriftSinceDrill {
		if !strings.Contains(plain, id) || !strings.Contains(plain, strconv.Itoa(count)) {
			t.Errorf("plain status missing drift %s=%d", id, count)
		}
	}
	for id, count := range report.PhaseCommits {
		if !strings.Contains(plain, id) || !strings.Contains(plain, strconv.Itoa(count)) {
			t.Errorf("plain status missing phase commits %s=%d", id, count)
		}
	}
	for _, cost := range report.Costs {
		for _, field := range []string{cost.Module, strconv.FormatFloat(cost.Seconds, 'f', -1, 64), strconv.Itoa(cost.EvalLines), strconv.Itoa(cost.DisposableLines)} {
			if !strings.Contains(plain, field) {
				t.Errorf("plain status missing cost field %q", field)
			}
		}
		for id, count := range cost.CasesPerSpec {
			if !strings.Contains(plain, id) || !strings.Contains(plain, strconv.Itoa(count)) {
				t.Errorf("plain status missing cases %s=%d", id, count)
			}
		}
	}
	for _, d := range report.Diagnostics {
		if !strings.Contains(plain+stderr, d.Code) || !strings.Contains(plain+stderr, d.Message) {
			t.Errorf("plain status missing diagnostic %+v in stdout=%q stderr=%q", d, plain, stderr)
		}
	}
	status, output, stderr := listRun(t, nested, "status", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("JSON status exit=%d stdout=%q stderr=%q", status, output, stderr)
	}
	data, diagnostics := statusData(t, output)
	modules := statusRows(t, data["modules"], "name", "id_prefix", "dir", "depends_on", "zones")
	if len(modules) != len(report.Modules) {
		t.Fatalf("modules=%d; want %d", len(modules), len(report.Modules))
	}
	for i, module := range report.Modules {
		for key, want := range map[string]any{"name": module.Name, "id_prefix": module.IDPrefix, "dir": module.Dir, "depends_on": module.DependsOn} {
			statusEqualField(t, modules[i], key, want)
		}
		var zones map[string]json.RawMessage
		if err := json.Unmarshal(modules[i]["zones"], &zones); err != nil || len(zones) != 6 {
			t.Errorf("module %s zones=%s: %v", module.Name, modules[i]["zones"], err)
		}
		for key, want := range map[string]any{"interface": module.Zones.Interface, "evals": module.Zones.Evals, "disposable": module.Zones.Disposable, "preserve": module.Zones.Preserve, "protected": module.Zones.Protected, "append_only": module.Zones.AppendOnly} {
			statusEqualField(t, zones, key, want)
		}
	}
	queue := statusRows(t, data["queue"], "slug", "path", "module", "phase", "blocked", "audit_rounds", "waiting_on", "waiting_rule")
	if len(queue) != len(report.Queue) {
		t.Fatalf("queue=%d; want %d", len(queue), len(report.Queue))
	}
	for i, entry := range report.Queue {
		for key, want := range map[string]any{"slug": entry.Slug, "path": entry.Path, "module": entry.Module, "phase": entry.Phase, "blocked": entry.Blocked, "audit_rounds": entry.AuditRounds, "waiting_on": entry.WaitingOn, "waiting_rule": entry.WaitingRule} {
			statusEqualField(t, queue[i], key, want)
		}
	}
	quarantines := statusRows(t, data["quarantines"], "module", "id", "reason")
	if len(quarantines) != len(report.Quarantines) {
		t.Fatalf("quarantines=%d; want %d", len(quarantines), len(report.Quarantines))
	}
	for i, q := range report.Quarantines {
		for key, want := range map[string]any{"module": q.Module, "id": q.ID, "reason": q.Reason} {
			statusEqualField(t, quarantines[i], key, want)
		}
	}
	optOuts := statusRows(t, data["boundary_opt_outs"], "module", "reason")
	if len(optOuts) != len(report.BoundaryOptOuts) {
		t.Fatalf("opt-outs=%d; want %d", len(optOuts), len(report.BoundaryOptOuts))
	}
	for i, o := range report.BoundaryOptOuts {
		statusEqualField(t, optOuts[i], "module", o.Module)
		statusEqualField(t, optOuts[i], "reason", o.Reason)
	}
	metrics := statusRows(t, data["metrics"], "module", "seconds", "cases_per_spec", "eval_lines", "disposable_lines")
	if len(metrics) != len(report.Costs) {
		t.Fatalf("metrics=%d; want measured modules only %d", len(metrics), len(report.Costs))
	}
	for i, cost := range report.Costs {
		for key, want := range map[string]any{"module": cost.Module, "seconds": cost.Seconds, "cases_per_spec": cost.CasesPerSpec, "eval_lines": cost.EvalLines, "disposable_lines": cost.DisposableLines} {
			statusEqualField(t, metrics[i], key, want)
		}
	}
	for key, want := range map[string]any{"drift_since_drill": report.DriftSinceDrill, "unapproved": report.Unapproved, "phase_commits": report.PhaseCommits} {
		statusEqualField(t, data, key, want)
	}
	if len(diagnostics) != len(report.Diagnostics) {
		t.Fatalf("diagnostics=%d; want %d: %s", len(diagnostics), len(report.Diagnostics), output)
	}
	for i, d := range report.Diagnostics {
		for key, want := range map[string]any{"code": d.Code, "message": d.Message, "warning": d.Warning} {
			statusEqualField(t, diagnostics[i], key, want)
		}
		if _, present := diagnostics[i]["path"]; present {
			statusEqualField(t, diagnostics[i], "path", d.Path)
		} else if d.Path != "" {
			t.Errorf("diagnostic %s missing path %q", d.Code, d.Path)
		}
	}
}

func Test_APP_118_APP_119_APP_122_StatusNearestRootEmptyContainers(t *testing.T) {
	outer := appProjectFixture(t)
	cliGit(t, outer, "init", "-q", "-b", "main")
	cliCommit(t, outer, "outer project")
	listChange(t, outer, "0001-outer.md", "alpha", "")
	inner := filepath.Join(outer, "nested-project")
	appPut(t, inner, "fireweed.yaml", appFixtureConfig)
	appPut(t, inner, "packages/inner/SPEC.md", "---\nname: inner\nid_prefix: INN\ndepends_on: []\n---\n# Inner\n\n## Specs\n")
	cliGit(t, inner, "init", "-q", "-b", "main")
	cliCommit(t, inner, "inner project")
	report, err := workflow.Status(inner)
	if err != nil {
		t.Fatal(err)
	}
	deep := filepath.Join(inner, "packages", "inner", "deep")
	if err := os.MkdirAll(deep, 0755); err != nil {
		t.Fatal(err)
	}
	status, plain, stderr := listRun(t, deep, "status")
	if status != 0 || stderr != "" || !strings.Contains(strings.ToLower(plain), "empty") {
		t.Errorf("plain empty status exit=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	status, output, stderr := listRun(t, deep, "status", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("JSON empty status exit=%d stdout=%q stderr=%q", status, output, stderr)
	}
	data, diagnostics := statusData(t, output)
	if len(diagnostics) != len(report.Diagnostics) {
		t.Errorf("empty diagnostics=%d; want %d", len(diagnostics), len(report.Diagnostics))
	}
	for _, key := range []string{"queue", "quarantines", "boundary_opt_outs", "metrics"} {
		if string(data[key]) != "[]" {
			t.Errorf("empty %s=%s; want []", key, data[key])
		}
	}
	for _, key := range []string{"drift_since_drill", "phase_commits"} {
		if string(data[key]) == "null" || !strings.HasPrefix(string(data[key]), "{") {
			t.Errorf("empty %s=%s; want object", key, data[key])
		}
	}
	modules := statusRows(t, data["modules"], "name", "id_prefix", "dir", "depends_on", "zones")
	if len(modules) != 1 || string(modules[0]["name"]) != `"inner"` {
		t.Errorf("nearest project modules=%s", data["modules"])
	}
}

func Test_APP_126_StatusUsageConfigurationAndOperationalFailures(t *testing.T) {
	check := func(root string, args []string, wantExit int, code, cause string) {
		t.Helper()
		status, out, errOut := listRun(t, root, args...)
		if status != wantExit || !strings.Contains(out+errOut, code+":") || !strings.Contains(out+errOut, cause) || !corrective(out+errOut) {
			t.Errorf("plain %v exit=%d stdout=%q stderr=%q; want %d %s %q action", args, status, out, errOut, wantExit, code, cause)
		}
		status, out, errOut = listRun(t, root, append(append([]string{}, args...), "--json")...)
		if status != wantExit || errOut != "" {
			t.Errorf("JSON %v exit=%d stdout=%q stderr=%q", args, status, out, errOut)
		}
		message := assertFailedJSON(t, out, code)
		if !strings.Contains(message, cause) || !corrective(message) {
			t.Errorf("JSON %v message=%q; want %q and action", args, message, cause)
		}
	}
	root := appProjectFixture(t)
	check(root, []string{"status", "extra"}, 2, "U01", "extra")
	appPut(t, root, "fireweed.yaml", "modules: [broken\n")
	check(root, []string{"status"}, 2, "C01", "fireweed.yaml")
	appPut(t, root, "fireweed.yaml", appFixtureConfig)
	listChange(t, root, "0001-repeat.md", "alpha", "")
	listChange(t, root, "0002-repeat.md", "beta", "")
	cliGit(t, root, "init", "-q", "-b", "main")
	cliCommit(t, root, "duplicate queue")
	if _, err := workflow.Status(root); err == nil {
		t.Fatal("public Status accepted duplicate slugs")
	}
	check(root, []string{"status"}, 1, "C13", "repeat")
	baseRoot := appProjectFixture(t)
	appPut(t, baseRoot, "fireweed.yaml", strings.Replace(appFixtureConfig, "base: main", "base: unavailable-base", 1))
	cliGit(t, baseRoot, "init", "-q", "-b", "main")
	cliCommit(t, baseRoot, "unavailable base")
	if _, err := workflow.Status(baseRoot); err == nil {
		t.Fatal("public Status accepted unavailable base ref")
	}
	check(baseRoot, []string{"status"}, 1, "P01", "unavailable-base")
}
