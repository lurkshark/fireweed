---
audit_rounds: 0
bases: {}
module: workflow
phase: audit
run: 20261030T141017-5d9748d32d502cd139e64639
---



## Motivation

The design's Cost is measured clause requires recorded JUnit time, cases per
spec, eval/disposable lines and B01–B03 warnings in verify and status. The
approved status slice first needs these values from actual Verify/Done calls,
then a workflow status model and immediate app wiring/NEW proof. Existing
verification reports contain JUnit data but no line counts or budget warnings.
Record a cost snapshot beside those reports; status must consume saved data.

## Interface

Add ModuleCost with Module string, Seconds float64, CasesPerSpec map[string]int,
EvalLines int and DisposableLines int. Add CostSnapshot with Run string,
Phase string, RecordedAt time.Time, Modules []ModuleCost and
Diagnostics []project.Diagnostic. Add Result.Costs []ModuleCost and public
Costs(root string) ([]CostSnapshot, error), with thin aliases/wrapper in api.go
and shared data in model/. The internal reader supplies the snapshots to the
next status change; it does not evaluate hooks.

## Added

### FLOW-097
Each Verify that evaluates modules, including through Done, MUST retain a cost snapshot under its existing run/phase directory and return those module costs in Result.Costs.
The snapshot MUST identify the run, phase and UTC recording time, replacing the previous cost snapshot for that run/phase.

### FLOW-098
Each evaluated module's cost MUST report its name, summed JUnit testcase seconds and number of matching testcases per spec ID, counting each testcase once per matching ID.
It MUST report eval and disposable line counts for regular files selected by the held approved projected zone map, counting an unterminated final line and excluding higher-priority zones.

### FLOW-099
Verify and Done MUST return warnings B01 for total evaluated-module seconds above budgets.verify_time (default15m), B02 for each spec's testcase count above budgets.cases_per_spec (default15), and B03 for each module's eval/disposable line ratio above budgets.eval_ratio (default4).
Warnings MUST identify measurement and limit without affecting verification success or failure.
Positive eval lines with zero disposable lines MUST exceed the ratio limit.

### FLOW-100
Costs MUST return the retained run/phase cost snapshots in recording-time order with their module values and budget diagnostics, or an empty slice when none exist.
The retained values MUST reflect the evaluated state even if source files or budget settings later change.

## Modified



## Removed



## Decisions

Cost snapshots use .fireweed/runs/<run>/<phase>/costs.json beside existing module
reports; repeated verification in that phase replaces its snapshot, as it does
the module reports. Each regular file's line count is newline count plus one
when nonempty content lacks a trailing newline.

## Rejected alternatives

Do not run hooks to collect status or create a second module graph/state store.

## Notes

Owner approved three changes: workflow costs, workflow status data, then app
status/actual NEW proof, reaching the fifth-application review; push only on
PASS and stop. Current applied count2. No phase exception or drill. Fixed
Sol/high workers and isolated assertion auditor, normal max2 findings rounds.
Workers never commit or advance phases; only controller owns workflow writes.
Preserve all existing behavior/assertions and held approved snapshots. Keep own
Verify below5m/full below15m; use existing deterministic fixture hooks, no real
agents/network/extra process observer. Add only focused cost cases. No private
tests, bootstrap source/tests, rebuild-v1, deleted/saved implementation oracles.
Missing B01–B03 explain entries are a known remaining integration obligation;
this single-module change cannot edit assets. Do not claim the full status
slice or fifth-review PASS while that obligation remains unresolved.

## Provenance

- evals 2026-10-30T14:25:35.399359Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=450 failing=6 passing=444
