---
bases: {}
module: app
---



## Motivation

The design CLI reading surface includes the queue, blocked changes and audit
rounds. NEW still rejects list although project.Queue already provides ordered
change documents. Expose that immediate caller and close with actual NEW queued,
held and empty-queue reads. Owner explicitly approved the one-change list
continuation in /tmp/fireweed-list-continuation-checkpoint.md; this third
consecutive app application becomes count2 since the last passing review.
Stop afterward; status, PR, drills and other remaining surface stay out of scope.

## Interface

No public API/model changes. Use existing project.Queue and nearest validated
project lookup with internal CLI result types and existing envelope/error helpers.

## Added

### new:list-command
list MUST discover the nearest project root and list project.Queue entries in returned order.
Plain success MUST return exit0 and identify every queued change by slug, path and module, with empty output for an empty queue.

### new:list-state
Each listed entry MUST report its phase, using queued when the phase is absent.
It MUST report the blocked reason and audit rounds, using an empty reason and zero rounds when absent.

### new:list-errors
Usage and project configuration errors MUST return exit2 with corrective guidance, using U01 for usage and the existing cause code for configuration.
Queue failures MUST return exit1 with their existing cause code, or P01 when uncoded, and corrective guidance.

### new:list-json
Successful list --json MUST emit schema1 ordered-array data with slug, path, module, phase, blocked and audit_rounds fields, including an empty array for an empty queue, and empty diagnostics.
Failures MUST emit null data with code/message diagnostics.
It MUST leave stderr empty.

## Modified



## Removed



## Decisions



## Notes

Owner: "Approve list continuation". This single change is the scoped consecutive
app exception; count1 before/count2 after, then stop again. No push/phase-write
exception/new drill or status/PR/library expansion. Reuse existing public Queue,
nearest root and diagnostics; preserve other command behavior and all old evals.
Use a few representative public app.Run cases for ordered queued/held/blocked
state, empty queue, nearest competing root, errors and JSON. Extend only the
existing single app-level PATH-shim read-convention eval with list invocations;
no new convention specs, observers, calibration or process-instrumentation tests.
Fixed real Sol/high workers and isolated restricted assertion auditor, max2
persisted findings rounds, frozen inputs during impl. Workers never commit or
advance phases; controller owns claims/phases/commits/apply/archive. Own Verify
currently about242s, must remain under5m; full under15m, no raised bounds or
extra confidence probes/matrices. No bootstrap source/tests, rebuild-v1, old
worker logs or deleted/abandoned/saved implementation oracles. Actual NEW read
proof may use public author/start/release in a disposable fixture; reuse prior
lifecycle evidence, perform no repeat lifecycle or drill. Stop on any block.
