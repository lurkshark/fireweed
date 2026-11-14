---
module: runner
phase: impl
run: 20261114T085216-0b63d7b9c092d3f941ffc1de
---

## Motivation

The design's Starting a change and Git behavior / Base branch require the
stale-base warning to reach users of automatic execution (R06/R12/R14).
Draft 0262 supplies it from workflow.Start. Main 3b3b64d's Next replaces that
initial State after Done and retains only separately collected worker-recovery
warnings, so successful next/run would lose the new warning. Preserve start
warnings through the runner instead of widening the workflow change.

## Interface

No public signature or type changes.

## Added

### RUN-072
Next MUST retain warnings returned by workflow.Start in State.Diagnostics on subsequent returns, including successful application, usage stops, invocation failures and verification exhaustion. Retained diagnostics MUST preserve code, message, path and warning status without being duplicated by phase advancement.

## Modified

## Removed

## Decisions

## Notes

Run after 0262 provides the base-tip warning. Use bounded public Next fixtures
with a locally resolved base tip absent from HEAD, and fake workers/auditors.
Exercise completion, verification exhaustion after a phase transition, and
representative invocation-failure/usage-stop returns. Assert the warning
alongside the existing run, applied or blocked state and retained attempt ref
as applicable. Reuse RUN-028's worker-recovery warning scenario to ensure both
warnings survive; do not conflate distinct P05 messages or add global deduping.

Confirm public Run's existing applied-result aggregation retains Next's
warning. Existing app next and successful run output already consume these
diagnostics; keep a final acceptance binary next/run warning check open until
that composition is exercised. Do not edit app's frozen evals or interfaces
from this runner change or widen it to run-level failure rendering.

Do not persist warnings across invocations, recreate Start warnings for an
already-held step, change RUN-028's recovery behavior, or change attempts,
audit limits, approval checks or continuation rules. Coordinate 0268's P05
explanation before treating the automatic warning path as acceptance-complete.

Draft only. No lifecycle execution, agent invocation or drill is authorized.
Workers obtain context before their step and use fireweed verify; the
controller alone advances and commits lifecycle work.

Controller verification passed in the evals phase with 180 passing cases
and six expected failures (the RUN-072 test and its five subtests). All
five scenarios lost the stale-base P05 warning after phase advancement.
The applied Next result still reported both worker-recovery warnings.
Run also dropped the unmerged P04 warning from its applied result.

The implementation belongs in the disposable
`packages/runner/internal/next/next.go`, which replaces the returned state
after Done and Block. Run already appends the full Next result to Applied.
Changes 0262 and 0268 are applied. The binary next/run acceptance check
remains open until the runner fix is applied.
## Audit

No findings; checked RUN-072

## Provenance

- evals 2026-11-14T09:04:32.327331Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=186 failing=6 passing=180
- audit 2026-11-14T09:04:57.715069Z agent=codex fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

