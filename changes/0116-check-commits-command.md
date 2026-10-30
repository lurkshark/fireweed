---
audit_rounds: 0
bases: {}
module: app
phase: impl
run: 20261030T050448-e34550cc4bb01812c7db3097
---



## Motivation

Design History checks and Merge strategies require check-commits to inspect
recorded phase commits and recognize squash merges. Public checks.History and
accurate H01–H04 explanations now exist, but NEW still rejects check-commits.
Wire the immediate caller and close the approved history slice with actual NEW
normal-history, squash and archive-violation results from one real Go lifecycle.
This is application5 after0112–0115; independent fifth-change quality/goal review
and fresh appropriate checks precede push only on PASS. Full acceptance stays open.

## Interface

No public API/model changes. Compose checks.History with nearest-project lookup
and existing CLI envelope/error helpers under app/internal/entry.

## Added

### APP-105
check-commits [range] MUST discover the nearest project root and call checks.History with the supplied range, or an empty range when omitted to use the configured local base.
Plain success MUST return exit0, report the number of checked commits, and list each squashed change with its full commit object ID and change path.

### APP-106
check-commits MUST return exit1 for history-rule violations and report every returned H01–H04 diagnostic with its affected path, full commit object ID and corrective action.

### APP-107
check-commits MUST return exit2 with corrective U01 for invalid arguments/options and the existing cause code for project configuration errors.
Operational history failures MUST return exit1 with their existing cause code, or P01 when uncoded, and a corrective action.

### APP-108
Successful check-commits --json MUST emit schema1 object data with commits in returned order and squashed entries containing commit and change, using empty arrays when appropriate and empty diagnostics.
Failures MUST emit null data with code/message diagnostics.
It MUST leave stderr empty.

## Modified



## Removed



## Decisions



## Notes

Owner explicitly approved the two consecutive history changes and fifth review.
Keep module ownership; no history-engine duplication or changes to other commands,
public models, protected files, frozen evals/interfaces during impl, or allocations.
Normal fixed real Sol/high workers and isolated audits; controller alone owns
claims/phases/commits/archive. Evals use public app.Run with representative local
Git fixtures; reuse detailed checks.History evals rather than duplicate its matrix.
Include ordinary authored queued documents with no phase and legitimate SPEC
application in the normal caller fixture, matching observed dependency fixes.
Keep own Verify below5m and full below15m with unchanged bounds, no calibration,
instrumentation or extra confidence probes. Scope audit to missing assertions,
max2 persisted findings round-trips; stop on any block. No bootstrap source/tests,
rebuild-v1, saved/deleted/abandoned implementation or old worker logs as oracles.
Close through clean-built NEW with one actual Go lifecycle/explicit nonempty
range, squash listing and H04 violation; then mandatory separate quality/goal
review/checks, push only PASS. No new drill authority or phase-write exception.

## Provenance

- evals 2026-10-30T05:14:08.486798Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=243 failing=6 passing=237
- audit 2026-10-30T05:15:20.208656Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings
Checked IDs: APP-105, APP-106, APP-107, APP-108
