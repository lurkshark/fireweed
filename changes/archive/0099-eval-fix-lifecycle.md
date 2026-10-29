---
already_satisfied:
    - FLOW-023
audit_rounds: 0
bases:
    FLOW-023: 67dcdbbe
    FLOW-036: ba36463a
module: workflow
phase: applied
run: 20261029T175930-a47eee99acd8dcc2416ccf47
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

### FLOW-078
For an eval_fix change, Verify MUST permit evals-phase edits only in the target
module's evals zone and the current change file. During impl it MUST permit
edits only to that change file.

### FLOW-079
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

## Provenance

- evals 2026-10-29T18:04:53.124143Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=361 failing=6 passing=355
- audit 2026-10-29T18:06:40.411193Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-29T18:17:49.977369Z agent=manual fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=361 failing=0 passing=361

## Audit

No findings
Checked IDs: FLOW-023, FLOW-036, FLOW-078, FLOW-079
