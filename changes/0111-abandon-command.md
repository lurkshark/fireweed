---
audit_pending:
    - APP-097
audit_rounds: 1
bases: {}
module: app
phase: impl
run: 20261030T001135-22be6a3923abe0974c22c059
---



## Motivation

The design's Abandoning a change section requires abandon <change>. FLOW-089–096
now provide the verified dependency, but NEW still rejects abandon as unfinished.
Return immediately to its app caller. Close the slice with actual NEW in a Go
fixture: bounded invalid implementation exhausts verification and saves a real
attempt, abandon restores phase work/deletes the pending file in one commit,
retains the attempt and frees the unapplied ID for subsequent Start. This is
application5 since passing review; require quality/goal review before more work.

## Interface

No public API/model changes. Compose workflow.Abandon with existing project root,
diagnostic and envelope helpers. Keep selection/reversal/claim behavior in workflow.

## Added

### APP-097
abandon <change> MUST discover the nearest project root and invoke workflow.Abandon for the named change.
Plain success MUST print confirmation naming the change and returned full commit object ID with exit0.

### APP-098
abandon MUST return exit2 with corrective U01 for usage errors, U02 for selector errors, and the existing cause code for project configuration errors.
Other workflow failures MUST return exit1 retaining their cause code and a corrective action.

### APP-099
abandon --json MUST emit schema1 with confirmation string data and empty diagnostics on success.
On failure it MUST emit null data with code/message diagnostics.
It MUST leave stderr empty.

## Modified



## Removed



## Decisions



## Notes

Owner explicitly approved native0109 recovery, fresh dependency, consecutive
app caller and fifth-change review. Canceled0109 is not an application. Reuse
current public app fixtures with temporary Git identities. Assert nested root,
plain/JSON confirmation/full OID and representative dirty P01, foreign P02,
applied P01, unknown U02, usage U01/configuration failures; reuse backend evals
rather than duplicating reversal/attempt/history matrices. Keep existing
assertions and interfaces unchanged. No other command/module, protected edit,
process primitive, instrumentation or reorganization. Do not read canceled0109
private implementation/worker logs, saved code, bootstrap source/tests or rebuild-v1.
Real Sol/high eval worker, separate restricted auditor, frozen evals/interface
then implementation and full native verification. Workers never commit or
advance phases. Own under5m/full under15m, no raised bounds. Actual NEW proof
and mandatory separate fifth quality/goal review/checks precede further product
changes; push and continue only on PASS under standing owner authorization.

## Provenance

- evals 2026-10-30T00:19:29.697432Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=217 failing=7 passing=210
- audit 2026-10-30T00:21:07.660677Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- evals 2026-10-30T00:26:54.145893Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=217 failing=7 passing=210
- audit 2026-10-30T00:28:45.164969Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings
Checked IDs: APP-097
