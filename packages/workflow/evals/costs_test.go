package evals_test

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func costAt(t *testing.T, root, run, phase string) (workflow.CostSnapshot, bool) {
	t.Helper()
	path := filepath.Join(root, ".fireweed", "runs", run, phase, "costs.json")
	b, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("cost snapshot %s: %v", path, err)
		return workflow.CostSnapshot{}, false
	}
	var snapshot workflow.CostSnapshot
	if err := json.Unmarshal(b, &snapshot); err != nil {
		t.Errorf("decode cost snapshot %s: %v", path, err)
		return workflow.CostSnapshot{}, false
	}
	return snapshot, true
}

func costFor(costs []workflow.ModuleCost, name string) (workflow.ModuleCost, bool) {
	for _, cost := range costs {
		if cost.Module == name {
			return cost, true
		}
	}
	return workflow.ModuleCost{}, false
}

func warningFor(ds []project.Diagnostic, code string) (project.Diagnostic, bool) {
	for _, d := range ds {
		if d.Code == code {
			return d, true
		}
	}
	return project.Diagnostic{}, false
}

func requireCostWarning(t *testing.T, ds []project.Diagnostic, code string, terms ...string) {
	t.Helper()
	d, ok := warningFor(ds, code)
	if !ok || !d.Warning {
		t.Errorf("%s warning absent from %#v", code, ds)
		return
	}
	for _, term := range terms {
		if !strings.Contains(d.Message, term) {
			t.Errorf("%s warning %q does not identify %q", code, d.Message, term)
		}
	}
}

func Test_FLOW_097_VerifyAndDoneReplaceRunPhaseCostSnapshot(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	write(t, root, ".fireweed/eval.sh", `#!/bin/sh
if [ "$1" = alpha ]; then
  if [ -f packages/alpha/internal/ready ]; then failure=''; code=0
  else failure='<failure/>'; code=1; fi
  if [ -f packages/alpha/evals/second ]; then seconds=0.75
  else seconds=0.25; fi
  printf '<testsuite><testcase name="ALP-001" time="%s">%s</testcase><testcase name="ALP-002" time="0.1"/></testsuite>' "$seconds" "$failure" > "$2"
  exit "$code"
fi
printf '<testsuite><testcase name="BET-001" time="0.2"/></testsuite>' > "$2"
`)
	commit(t, root, "deterministic cost report")
	state := start(t, root)
	before := time.Now().UTC()
	first, err := verify(root, true)
	if err != nil {
		t.Fatalf("first Verify: %v", err)
	}
	if cost, ok := costFor(first.Costs, "alpha"); !ok || math.Abs(cost.Seconds-0.35) > 1e-9 {
		t.Errorf("FLOW-097 Verify costs = %#v; want alpha at 0.35 seconds", first.Costs)
	}
	savedFirst, okFirst := costAt(t, root, state.Run, "evals")
	if okFirst {
		if savedFirst.Run != state.Run || savedFirst.Phase != "evals" || savedFirst.RecordedAt.IsZero() || savedFirst.RecordedAt.Location() != time.UTC || savedFirst.RecordedAt.Before(before) || savedFirst.RecordedAt.After(time.Now().UTC()) || !reflect.DeepEqual(savedFirst.Modules, first.Costs) {
			t.Errorf("FLOW-097 saved evals identity/time/modules = %#v; returned costs %#v", savedFirst, first.Costs)
		}
	}
	write(t, root, "packages/alpha/evals/second", "second\n")
	second, err := verify(root, true)
	if err != nil {
		t.Fatalf("second Verify: %v", err)
	}
	if cost, ok := costFor(second.Costs, "alpha"); !ok || math.Abs(cost.Seconds-0.85) > 1e-9 {
		t.Errorf("FLOW-097 replacement Verify costs = %#v; want alpha at 0.85 seconds", second.Costs)
	}
	savedSecond, okSecond := costAt(t, root, state.Run, "evals")
	if okSecond && (savedSecond.Run != state.Run || savedSecond.Phase != "evals" || !reflect.DeepEqual(savedSecond.Modules, second.Costs) || okFirst && savedSecond.RecordedAt.Before(savedFirst.RecordedAt)) {
		t.Errorf("FLOW-097 replacement snapshot = %#v; first %#v, returned %#v", savedSecond, savedFirst, second.Costs)
	}
	retained, err := workflow.Costs(root)
	if err != nil || len(retained) != 1 || retained[0].Phase != "evals" || !reflect.DeepEqual(retained[0].Modules, second.Costs) {
		t.Errorf("FLOW-097 repeated Verify retained snapshots = %#v, %v; want one replaced evals snapshot", retained, err)
	}
	evalsDone, err := done(root)
	if err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	if cost, ok := costFor(evalsDone.Costs, "alpha"); !ok || math.Abs(cost.Seconds-0.85) > 1e-9 {
		t.Errorf("FLOW-097 evals Done costs = %#v", evalsDone.Costs)
	}
	if saved, ok := costAt(t, root, state.Run, "evals"); ok && !reflect.DeepEqual(saved.Modules, evalsDone.Costs) {
		t.Errorf("FLOW-097 evals Done snapshot modules = %#v; result %#v", saved.Modules, evalsDone.Costs)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	implDone, err := done(root)
	if err != nil {
		t.Fatalf("impl Done: %v", err)
	}
	if saved, ok := costAt(t, root, state.Run, "impl"); ok && (saved.Run != state.Run || saved.Phase != "impl" || saved.RecordedAt.Location() != time.UTC || !reflect.DeepEqual(saved.Modules, implDone.Costs)) {
		t.Errorf("FLOW-097 impl Done snapshot = %#v; result costs %#v", saved, implDone.Costs)
	}
	if _, ok := costFor(implDone.Costs, "beta"); !ok {
		t.Errorf("FLOW-097 full impl Done omitted evaluated dependent: %#v", implDone.Costs)
	}
}

func Test_FLOW_098_ModuleJUnitCostsAndHeldZoneLineCounts(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\n", "set:\n  evals: [evals/**, internal/overlap.txt]\n", "", "", "")
	path := filepath.Join(root, "packages/alpha/SPEC.md")
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, "packages/alpha/SPEC.md", strings.Replace(string(b), "depends_on: []\n", "depends_on: []\nprotected: [internal/protected.txt]\npreserve: [internal/preserve.txt]\n", 1))
	write(t, root, "packages/alpha/evals/one.txt", "a\nb")
	write(t, root, "packages/alpha/internal/overlap.txt", "x\n")
	write(t, root, "packages/alpha/internal/code.txt", "p\nq")
	write(t, root, "packages/alpha/internal/preserve.txt", "r\n")
	write(t, root, "packages/alpha/internal/protected.txt", "s\n")
	if err := os.Symlink("code.txt", filepath.Join(root, "packages/alpha/internal/link.txt")); err != nil {
		t.Fatal(err)
	}
	write(t, root, "packages/beta/evals/one.txt", "z")
	write(t, root, "packages/beta/internal/code.txt", "i\n")
	write(t, root, ".fireweed/eval.sh", `#!/bin/sh
if [ "$1" = alpha ]; then
  if [ -f packages/alpha/internal/ready ]; then failure=''; code=0
  else failure='<failure/>'; code=1; fi
  printf '<testsuite><testcase name="ALP-001 ALP-002" classname="ALP-001" time="1.25">%s</testcase><testcase name="ALP-001 second" time="0.5"/></testsuite>' "$failure" > "$2"
  exit "$code"
fi
printf '<testsuite><testcase name="BET-001" time="2"/></testsuite>' > "$2"
`)
	commit(t, root, "held zone and multi-ID cost fixture")
	start(t, root)
	if _, err := done(root); err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	result, err := verify(root, false)
	if err != nil {
		t.Fatalf("full impl Verify: %v", err)
	}
	if cases := result.Reports["alpha"].Cases; len(cases) != 2 || !reflect.DeepEqual(cases[0].IDs, []string{"ALP-001", "ALP-002"}) || !reflect.DeepEqual(cases[1].IDs, []string{"ALP-001"}) || result.Reports["alpha"].Seconds != 1.75 || result.Reports["beta"].Seconds != 2 {
		t.Fatalf("cost fixture reports have unexpected cases or time: %#v", result.Reports)
	}
	if len(result.Costs) != 2 {
		t.Errorf("FLOW-098 costs for evaluated modules = %#v; want alpha and beta", result.Costs)
	}
	alpha, ok := costFor(result.Costs, "alpha")
	if !ok || alpha.Seconds != 1.75 || !reflect.DeepEqual(alpha.CasesPerSpec, map[string]int{"ALP-001": 2, "ALP-002": 1}) || alpha.EvalLines != 3 || alpha.DisposableLines != 3 {
		t.Errorf("FLOW-098 alpha cost = %#v, present %v; want seconds 1.75, counts 2/1, eval 3, disposable 3", alpha, ok)
	}
	beta, ok := costFor(result.Costs, "beta")
	if !ok || beta.Seconds != 2 || !reflect.DeepEqual(beta.CasesPerSpec, map[string]int{"BET-001": 1}) || beta.EvalLines != 1 || beta.DisposableLines != 1 {
		t.Errorf("FLOW-098 beta cost = %#v, present %v", beta, ok)
	}
}

func Test_FLOW_099_CustomWarningsSurviveSuccessAndFailure(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\nbudgets:\n  verify_time: 3s\n  cases_per_spec: 1\n  eval_ratio: 1\n", "", "", "", "")
	write(t, root, "packages/alpha/evals/weight.txt", "a\nb\n")
	write(t, root, ".fireweed/beta-fail", "yes\n")
	write(t, root, ".fireweed/eval.sh", `#!/bin/sh
if [ "$1" = alpha ]; then
  if [ -f packages/alpha/internal/ready ]; then failure=''; code=0
  else failure='<failure/>'; code=1; fi
  printf '<testsuite><testcase name="ALP-001 first" time="1">%s</testcase><testcase name="ALP-001 second" time="1"/><testcase name="ALP-002" time="1"/></testsuite>' "$failure" > "$2"
  exit "$code"
fi
printf '<testsuite><testcase name="BET-001" time="1"><failure/></testcase></testsuite>' > "$2"
exit 1
`)
	commit(t, root, "custom budget cost fixture")
	start(t, root)
	verified, err := verify(root, true)
	if err != nil {
		t.Fatalf("warning made evals Verify fail: %v", err)
	}
	if d, ok := warningFor(verified.Diagnostics, "B01"); ok {
		t.Errorf("FLOW-099 B01 at custom 3s threshold = %#v", d)
	}
	requireCostWarning(t, verified.Diagnostics, "B02", "ALP-001", "2", "1")
	requireCostWarning(t, verified.Diagnostics, "B03", "alpha", "1")
	if cost, ok := costFor(verified.Costs, "alpha"); !ok || cost.EvalLines != 2 || cost.DisposableLines != 0 {
		t.Errorf("FLOW-099 zero-disposable measurement = %#v", verified.Costs)
	}
	evalsDone, err := done(root)
	if err != nil {
		t.Fatalf("warning made evals Done fail: %v", err)
	}
	if d, ok := warningFor(evalsDone.Diagnostics, "B01"); ok {
		t.Errorf("FLOW-099 evals Done B01 at custom 3s threshold = %#v", d)
	}
	for _, code := range []string{"B02", "B03"} {
		requireCostWarning(t, evalsDone.Diagnostics, code)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	before := head(t, root)
	failed, err := done(root)
	if err == nil {
		t.Error("fixture expected dependent testcase failure from impl Done")
	}
	if head(t, root) != before {
		t.Error("failed impl Done committed")
	}
	if len(failed.Reports["beta"].Cases) != 1 {
		t.Errorf("failure fixture did not evaluate beta: %#v", failed.Reports)
	}
	for _, code := range []string{"B01", "B02", "B03"} {
		requireCostWarning(t, failed.Diagnostics, code)
	}
	requireCostWarning(t, failed.Diagnostics, "B01", "4", "3")
	if snapshot, ok := costAt(t, root, failed.State.Run, "impl"); ok {
		if !reflect.DeepEqual(snapshot.Modules, failed.Costs) {
			t.Errorf("FLOW-099 failed Done snapshot costs = %#v; result %#v", snapshot.Modules, failed.Costs)
		}
		for _, code := range []string{"B01", "B02", "B03"} {
			retained, found := warningFor(snapshot.Diagnostics, code)
			returned, present := warningFor(failed.Diagnostics, code)
			if !found || !present || retained != returned {
				t.Errorf("FLOW-099 %s failed Done retained warning = %#v; result %#v", code, retained, returned)
			}
		}
	}
}

func Test_FLOW_099_DefaultBudgetBoundaries(t *testing.T) {
	for _, tc := range []struct {
		name  string
		cases int
		secs  int
		lines int
		warn  bool
	}{{"equal", 15, 900, 4, false}, {"above", 16, 901, 5, true}} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t, "", "", "", "", "")
			write(t, root, "packages/alpha/evals/weight.txt", strings.Repeat("line\n", tc.lines))
			write(t, root, "packages/alpha/internal/code.txt", "line\n")
			var cases strings.Builder
			for i := 0; i < tc.cases; i++ {
				seconds, failure := 0, ""
				if i == 0 {
					seconds, failure = tc.secs, "<failure/>"
				}
				fmt.Fprintf(&cases, `<testcase name="ALP-001 case %d" time="%d">%s</testcase>`, i, seconds, failure)
			}
			cases.WriteString(`<testcase name="ALP-002" time="0"/>`)
			write(t, root, ".fireweed/eval.sh", "#!/bin/sh\ncat > \"$2\" <<'XML'\n<testsuite>"+cases.String()+"</testsuite>\nXML\nexit 1\n")
			commit(t, root, "default budget boundary fixture")
			start(t, root)
			result, err := verify(root, true)
			if err != nil {
				t.Fatalf("boundary Verify: %v", err)
			}
			if report := result.Reports["alpha"]; len(report.Cases) != tc.cases+1 || report.Seconds != float64(tc.secs) {
				t.Fatalf("budget fixture JUnit = %#v; want %d cases and %ds", report, tc.cases+1, tc.secs)
			}
			for _, code := range []string{"B01", "B02", "B03"} {
				d, ok := warningFor(result.Diagnostics, code)
				if ok != tc.warn || ok && !d.Warning {
					t.Errorf("FLOW-099 %s at defaults (secs %d, cases %d, ratio %d) = %#v; want warning %v", code, tc.secs, tc.cases, tc.lines, result.Diagnostics, tc.warn)
				}
			}
			if tc.warn {
				requireCostWarning(t, result.Diagnostics, "B01")
				if d, _ := warningFor(result.Diagnostics, "B01"); !strings.Contains(d.Message, "901") && !strings.Contains(d.Message, "15m1s") {
					t.Errorf("FLOW-099 B01 warning omits measured 901s: %q", d.Message)
				}
				if d, _ := warningFor(result.Diagnostics, "B01"); !strings.Contains(d.Message, "900") && !strings.Contains(d.Message, "15m") {
					t.Errorf("FLOW-099 B01 warning omits default 15m limit: %q", d.Message)
				}
				requireCostWarning(t, result.Diagnostics, "B02", "ALP-001", "16", "15")
				requireCostWarning(t, result.Diagnostics, "B03", "alpha", "5", "4")
				completed, err := done(root)
				if err != nil {
					t.Fatalf("default warnings made evals Done fail: %v", err)
				}
				for _, code := range []string{"B01", "B02", "B03"} {
					requireCostWarning(t, completed.Diagnostics, code)
				}
			}
		})
	}
}

func Test_FLOW_100_CostsEmptyOrderedAndRetainedAfterEdits(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\nbudgets:\n  eval_ratio: 1\n", "", "", "", "")
	initial, err := workflow.Costs(root)
	if err != nil || len(initial) != 0 {
		t.Errorf("FLOW-100 empty Costs = %#v, %v", initial, err)
	}
	write(t, root, "packages/alpha/evals/weight.txt", "one\ntwo\n")
	commit(t, root, "measured eval lines")
	state := start(t, root)
	evalsResult, err := done(root)
	if err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	implResult, err := done(root)
	if err != nil {
		t.Fatalf("impl Done: %v", err)
	}
	snapshots, err := workflow.Costs(root)
	if err != nil {
		t.Fatalf("Costs: %v", err)
	}
	if len(snapshots) != 2 {
		t.Errorf("FLOW-100 retained snapshots = %#v; want evals then impl", snapshots)
	} else {
		if snapshots[0].Run != state.Run || snapshots[0].Phase != "evals" || snapshots[1].Run != state.Run || snapshots[1].Phase != "impl" || snapshots[0].RecordedAt.After(snapshots[1].RecordedAt) {
			t.Errorf("FLOW-100 recording order/identity = %#v", snapshots)
		}
		if !reflect.DeepEqual(snapshots[0].Modules, evalsResult.Costs) || !reflect.DeepEqual(snapshots[1].Modules, implResult.Costs) {
			t.Errorf("FLOW-100 retained module values = %#v; Done costs %#v, %#v", snapshots, evalsResult.Costs, implResult.Costs)
		}
		for _, snapshot := range snapshots {
			requireCostWarning(t, snapshot.Diagnostics, "B03")
		}
	}
	write(t, root, "packages/alpha/evals/weight.txt", "changed\n")
	write(t, root, "packages/alpha/internal/ready", "changed\ntoo\n")
	config, err := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, "fireweed.yaml", strings.Replace(string(config), "eval_ratio: 1", "eval_ratio: 100", 1))
	after, err := workflow.Costs(root)
	if err != nil || !reflect.DeepEqual(after, snapshots) {
		t.Errorf("FLOW-100 Costs changed after source/settings edits: before %#v, after %#v, err %v", snapshots, after, err)
	}
}
