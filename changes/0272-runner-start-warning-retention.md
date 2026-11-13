---
module: runner
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

### new:start-warning-retention
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
