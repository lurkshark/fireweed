---
audit_pending:
    - APP-095
audit_rounds: 1
bases: {}
module: app
phase: evals
run: 20261029T220850-60043a84e502eabe52e89654
---



## Motivation

Design Lifecycle CLI requires commits <change> to list a change's commits on
the current branch and support later abandon. FLOW-087/088 now provide the
public dependency, but NEW still rejects commits as unimplemented. Return to
its immediate app caller. Close this slice with an actual NEW Go next lifecycle,
then plain/JSON commits matching its three phase object IDs without read-side
mutation or extra agent calls, plus an empty pending history and unknown name.

## Interface

No API/model changes. Compose the existing workflow.Commits with root lookup,
CLI diagnostics and envelope helpers; keep selection/history in workflow.

## Added

### APP-094
commits <change> MUST discover the nearest project root and invoke
workflow.Commits. Plain success MUST print the returned object IDs one per line
in returned order with exit0, returning empty output when the list is empty.

### APP-095
commits MUST return exit2 with corrective U01 for usage errors, U02 for selector
errors, and the existing cause code for project configuration errors. Other
workflow failures MUST return exit1 retaining their cause code and a corrective
action.

### APP-096
commits --json MUST emit schema1 with an array of returned object IDs, including
an empty array for no commits, and empty diagnostics on success. On failure it
MUST emit null data with code/message diagnostics. It MUST leave stderr empty.

## Modified



## Removed



## Decisions



## Notes

Reuse public app/lifecycle fixtures, explicit temporary Git identities and
workflow dependency. Cover nested root/representative pending and archive
selectors, returned order/plain/JSON, empty list and representative usage,
unknown selector/configuration and operational failures. Do not repeat backend
merge/history matrices or reimplement selection. Preserve existing assertions.
No new options, models, APIs, other commands/modules, hooks/agents, protected
edits, reorganization or instrumentation. Evals freeze before impl. Workers
never commit or advance phases; coordinator owns full native verification,
restricted independent audit/archive and the actual NEW proof. Fixed Sol/high;
own under5m/full under15m, no raised bounds. This is the commits slice's second
change and the third application since passing checkpoint; no third behavior
continuation is authorized by this change.

## Provenance

- evals 2026-10-29T22:22:11.348073Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=209 failing=5 passing=204
- audit 2026-10-29T22:25:13.269472Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

- APP-095: "workflow failures MUST return exit1 retaining their cause code and a corrective" — missing explicit operational cause-code assertion.
