---
already_satisfied:
    - FLOW-023
bases:
    FLOW-023: 67dcdbbe
    FLOW-036: ba36463a
module: workflow
---



## Motivation

The design's eval_fix kind repairs wrong frozen evals without changing correct
specs or implementation. NEW currently treats it as ordinary evals: it permits
interface/stub edits, accepts failing evals, and omits listed IDs from the audit
requirement. Implement that missing lifecycle path. Close the slice through
actual NEW next in a Go fixture: correct a wrong existing test, require passing
coverage, separately audit its listed ID, then archive without changing code or
contracts. Normal fixes/spec changes retain genuine fail-first enforcement.

## Interface

No public API/model changes. Compose existing workflow Start/Verify/Done,
public checks and approved projected zones. No new checks request flags are
needed. Worker guidance updates and regenerate semantics are separate work.

## Added

### new:eval-fix-zones
For an eval_fix change, Verify MUST permit evals-phase edits only in the target
module's evals zone and the current change file. During impl it MUST permit
edits only to that change file.

### new:eval-fix-passing-coverage
During evals for an eval_fix change, Verify MUST require all target-module
evals to pass rather than fail first, with every listed eval_fix ID covered by
at least one referencing testcase.

## Modified

### FLOW-023
During evals Verify for spec and fix changes MUST require genuine fail-first
evidence for every enforced change ID that is not already_satisfied.

### FLOW-036
Audit Done MUST require every enforced ID in the change's Added, Modified,
fixes or eval_fix entries to appear in a No findings submission's checked-ID list.

## Removed



## Decisions



## Rejected alternatives



## Notes

FLOW-023 is a qualification of its existing spec/fix obligation, already covered
by retained fail-first assertions; do not weaken those tests. Use valid
eval_fix-only documents with a list of existing IDs, no spec edits or fixes.
Keep normal coverage, audit rounds, approval/claims, impl dependents/boundaries,
archival and runtime validation intact. Reuse existing public fixture helpers
but override their default Modified content before committing the fixture.
Require real behavioral failure evidence for the new restrictions and audit
inclusion, not a compiler/helper failure. A representative public lifecycle can
correct an eval, pass evals/audit/impl and archive with unchanged implementation,
SPEC and interface; use stable canonical fixture setup or semantic comparisons
where formatting is not a requirement. Negative cases should distinguish
failing evals, missing listed coverage, forbidden interface/impl edits and an
incomplete audit checked-ID list. Assert forbidden edits are rejected before a
hook and failed transitions leave HEAD unchanged. No matrix or more cases just
for confidence. Existing helpers/assertions remain unchanged.

Evals/interfaces freeze before impl. No new diagnostic allocation, generic kind
framework, API, static grammar rewrite, code organization, other-module edits,
process instrumentation, protected infrastructure, drill or phase exception.
Use existing cause codes for existing rejection classes. Own check under5m,
full under15m. Workers never commit or advance workflow. Coordinator owns
native phases, isolated audits, clean NEW build and actual Go fixture.
