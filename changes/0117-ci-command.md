---
bases: {}
module: app
---



## Motivation

Design CI requires one deterministic check command: static/history, all-module
coverage, boundaries and doctor canaries, with no agent probes and an optional
--no-doctor. NEW still rejects ci; public dependency capabilities now exist.
Compose them directly in app and close with actual NEW Go coverage/history/
boundary/canary proof. This is the second consecutive app application after0116;
stop before a third unless separately authorized. Count0 after passing/pushed
0112–0116 gate; full ACCEPTANCE remains open.

## Interface

No API/model changes. Compose existing project/checks/assets interfaces in
app/internal/entry; preserve the existing command behavior and helpers.

## Added

### new:ci-static-history
ci [range] MUST discover the nearest project root and run static project checks plus history checks on the supplied range, or the configured local-base range when omitted.
Static or history errors MUST prevent a successful CI result.

### new:ci-coverage-boundaries
ci MUST run evals for every discovered module and require passing coverage of its current specs, regardless of verify_dependents.
The configured boundary checks MUST also pass for CI to succeed.

### new:ci-doctor
ci MUST run doctor canaries by default with agent probes disabled.
With --no-doctor it MUST skip the doctor canaries while retaining the other deterministic checks.

### new:ci-results
Successful ci MUST return exit0 with a confirmation.
Usage and project configuration errors MUST return exit2 with corrective guidance, using U01 for usage and the existing cause code for configuration.
Deterministic check failures MUST return exit1 retaining their cause codes and corrective guidance.

### new:ci-json
Successful ci --json MUST emit schema1 confirmation-string data and collected diagnostics, including warnings.
Failures MUST emit null data with code/message diagnostics.
It MUST leave stderr empty.

## Modified



## Removed



## Decisions



## Notes

No new checker library, public API/model, diagnostic allocation, agent execution,
recipe/backend changes or history duplication. Use existing Evaluate/RequirePass/
RequireCoverage/Boundaries/Doctor capabilities; CI always covers all modules and
uses Doctor NoAgent, preserving the invoking executable for fixture Go hooks.
Use representative cheap local Git/hook fixtures through public app.Run, keeping
all prior evals and helpers. Avoid duplicating backend matrices or expensive Go
canary tests; existing app own Verify about218s leaves limited room under5m.
Full remains under15m; no timeout increases, instrumentation, observer calibration,
extra confidence probes or native source/test oracles. No commits/phase writes by
workers, frozen evals/interfaces during impl, protected/other-module edits, new
drills or scope widening. Fixed real Sol/high workers, isolated bounded audits,
controller-owned claims/phases/commits/archive and normal stop rules apply.
Close through clean-built NEW with a disposable actual Go fixture, explicit
nonempty history range, all-module coverage/boundaries and doctor-canary evidence;
reuse valid completed evidence instead of repeating a lifecycle. Count1 after
application; then the third-consecutive-app checkpoint applies before more app.
