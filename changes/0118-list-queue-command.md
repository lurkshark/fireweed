---
audit_rounds: 0
bases: {}
module: app
phase: audit
run: 20261030T063511-92982ff89d0d2ea4ed17a021
---



## Motivation

The design CLI reading surface includes the queue, blocked changes and audit
rounds. NEW still rejects list although project.Queue provides ordered change
documents. Expose that immediate caller and close with an actual NEW queued,
held and empty-queue read proof. Owner approved the scoped recovery in
/tmp/fireweed-0118-audit-checkpoint.md: native abandon of the blocked attempt,
remove only its unassertable P01 conditional, and carry the plain empty-reason
finding into fresh evals/audit/implementation/verification. Count remains1 before
application and becomes2 afterward; stop then. No new push or phase exception.

## Interface

No public API/model changes. Use existing project.Queue and nearest validated
project lookup with internal CLI result types and existing envelope/error helpers.

## Added

### APP-114
list MUST discover the nearest project root and list project.Queue entries in returned order.
Plain success MUST return exit0 and identify every queued change by slug, path and module, with empty output for an empty queue.

### APP-115
Each listed entry MUST report its phase, using queued when the phase is absent.
It MUST report the blocked reason and audit rounds, using an empty reason and zero rounds when absent.

### APP-116
Usage and project configuration errors MUST return exit2 with corrective guidance, using U01 for usage and the existing cause code for configuration.
Queue failures MUST return exit1 with their existing cause code and corrective guidance.

### APP-117
Successful list --json MUST emit schema1 ordered-array data with slug, path, module, phase, blocked and audit_rounds fields, including an empty array for an empty queue, and empty diagnostics.
Failures MUST emit null data with code/message diagnostics.
It MUST leave stderr empty.

## Modified



## Removed



## Decisions



## Notes

Owner: "Approve scoped recovery". Original0118-list-command was natively blocked
at audit with APP-115 plain empty-reason and APP-116 unassertable P01-conditional
Questions; normal native abandon78705e1 removed its unapplied phase work. This
corrected contract resolves only the owner-reviewed generic fallback requirement,
retaining all observable Queue cause/exit/action requirements. Fresh worker uses
preserved PUBLIC change/evals only, not any saved implementation (none exists).
Restore every existing assertion, including C03 regular-file error/canonical path
and C13 ambiguity, and add the missing plain empty-reason assertion. Keep the two
list calls in the existing single read-command convention eval, with no new
observer, calibration, process instrumentation, unsupported cases or matrix.
No API/model, protected, other-module, status/PR/library, diagnostic allocation or
new error path. No claim/phase/ref writes or commits by workers; controller owns
all native phases/commits/application/archive. Fresh fixed Sol/high eval worker,
restricted independent audit and fresh impl worker; normal max2 findings rounds.
Own Verify previously about252s, must remain under5m; full under15m, no raised
bounds or extra confidence probes. No bootstrap source/tests, rebuild-v1, old
worker logs, deleted/abandoned/saved implementation oracles. Actual NEW read proof
uses public author/start/release in a disposable fixture; no repeated lifecycle
or drill. Count stays1 until a real application, then2; stop afterward under the
one-change consecutive-app scope. No further recovery on block without owner.

## Provenance

- evals 2026-10-30T06:41:28.560667Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=250 failing=3 passing=247
