---
bases: {}
module: assets
---



## Motivation

The design requires stable diagnostic codes with readable causes and corrective
actions. Workflow0119/0120 now emits B01–B03 cost warnings and P08 repeated-module
warnings, but assets.Explain rejects all four. NEW explain must make those
messages actionable before the status caller is exposed. Owner approved the
scoped warning explanation change before the mandatory fifth review; P08 is the
same newly emitted status-warning companion, with no extra application or new
diagnostic allocation. Close this slice through actual NEW explain lookups.

## Interface

No API/model changes. Extend existing Explain lookups only.

## Added

### new:explain-eval-time-budget
Explain MUST describe B01 as a budget warning when summed JUnit testcase seconds for evaluated modules exceed budgets.verify_time, whose default is 15m.
It MUST advise reviewing expensive or redundant eval work while preserving required assertions.

### new:explain-testcase-budget
Explain MUST describe B02 as a budget warning when a spec's testcase count exceeds budgets.cases_per_spec, whose default is 15.
It MUST advise consolidating redundant cases while preserving required behavioral assertions.

### new:explain-eval-line-budget
Explain MUST describe B03 as a budget warning when eval lines divided by disposable lines exceed budgets.eval_ratio, whose default is 4, including positive eval lines with zero disposable lines.
It MUST advise simplifying redundant eval or fixture code while preserving required assertions.

### new:explain-repeated-module
Explain MUST describe P08 as a warning when at least three consecutive applied changes target the same module.
It MUST advise reviewing that sequence's scope with the project owner before continuing.

## Modified



## Removed



## Decisions



## Rejected alternatives

No unused diagnostic entries, code reallocations or broad explanation rewrite.

## Notes

Count4 before application,5 afterward. Owner approved revised status recovery:
this warning explanation change, mandatory separate Sol/high quality/goal review
and appropriate fresh checks, push only on PASS; then immediately app status
and actual NEW proof, stopping with the caller local. Stop on any block/finding.
No phase exception, drill, altered refs/bounds/counts or saved implementation.

Fresh fixed Sol/high workers and isolated restricted assertion auditor; normal
max2 audit rounds and re-audit only prior finding IDs. Workers never commit or
advance phases. Only public-interface evals in evals, only internal strings in
impl; no API/model, other-module, protected or unrelated code changes. Preserve
all existing assertions and explanations. Keep tests focused, own Verify under5m
and full under15m; no added instrumentation or raised timeouts. No bootstrap
source/tests, rebuild-v1, deleted/abandoned/saved implementation or old worker
logs as oracles. Actual NEW explain proof uses the same clean build as the
fifth-change checks; no extra lifecycle or drill.
