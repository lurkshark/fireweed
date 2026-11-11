---
eval_fix:
    - RUN-036
    - RUN-037
module: runner
phase: audit
run: 20261111T033424-f986d59bb4e07731df6f0db4
---

## Motivation

`Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied` failed the first
full CI run on the controller candidate and passed in isolation. It gives a
10-second `--for` budget to a run whose first change must finish all of its
phases, with real Git and hook processes, while the whole suite runs in
parallel. Under load the first change does not always finish in time, so
the eval fails for reasons unrelated to RUN-036 or RUN-037. Full
verification gates every `done`, so a flaky eval can block an unrelated
change.

## Interface

Audited eval_fix only. Edit only the runner evals zone and this change.

## Added

## Modified

## Removed

## Decisions

## Notes

In packages/runner/evals/time_limit_test.go, take timing out of the
outcome. Make beta's evals-phase fake agent write its marker and then block
(for example `sleep 120`) so only the deadline can end it, and raise `For`
enough that alpha finishes comfortably under load (for example 30s). Keep
the elapsed-time bound relative to `For` with a fixed margin, and keep every
existing assertion: one applied alpha, held beta work, four invocations
ending at beta evals, and a time-stop error naming the limit.

Apply the same pattern to any other runner eval in that file whose success
depends on a step finishing inside a short real-time budget. Keep the
runner suite under the SYSTEM.md verify-time bound.

The cross-change fixture gives the run 30 seconds and allows four seconds
for cleanup. Alpha's impl agent spends six seconds so resetting the budget
for beta would exceed that margin. The interrupted fixtures write their
markers before blocking, and the usage-wait preflight has more time to reach
the agent while still requiring an early stop.

`fireweed verify time-budget-eval-margin` passed in 6m39s. The uncached
`go test -count=1 ./...` run also passed, with runner evals taking 6m31s.
## Provenance
- evals 2026-11-11T03:49:19.370208Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=166 failing=0 passing=166
