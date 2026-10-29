---
bases: {}
module: app
---



## Motivation

Design CLI Lifecycle and local refs require release <change> to drop a local
claim by hand. FLOW-084–086 now provide the checked public Release API, but NEW
still rejects release as unimplemented. Return immediately to its app caller
after the approved correction-gate PASS. Close this slice with an actual NEW
Go fixture: start, retain dirty worker content and runtime state, release the
owned claim, then report a repeated missing-claim failure without other changes.

## Interface

No public API/model changes. Compose workflow.Release using existing root
lookup, CLI diagnostics and envelope helpers.

## Added

### new:release-command
release <change> MUST discover the nearest project root and invoke
workflow.Release for the named change. On success it MUST print confirmation
with exit0, preserving work, runtime state and HEAD while dropping only the
owned claim.

### new:release-failure
release MUST return exit1 with the workflow cause code and corrective action
when claim release fails, including missing or foreign claims. It MUST return
exit2 with a corrective diagnostic for usage or project configuration errors.

### new:release-json
release --json MUST emit schema1 with confirmation string data and empty
diagnostics on success, or null data with code/message diagnostics on failure.
It MUST leave stderr empty.

## Modified



## Removed



## Decisions



## Notes

Reuse public lifecycle fixtures and backend Release behavior. Cover nested root
lookup/alias, dirty work/runtime retention, representative missing/foreign claim
rejection and existing argument validation, plus plain/JSON shapes. No private
claim serialization or release reimplementation, hooks/agents/Verify invocation,
new options, APIs, other commands, other modules, protected files or ref tricks.
Preserve existing assertions. Evals freeze before impl. Real Sol/high workers
never commit or advance phases; coordinator owns full verification and separate
restricted audit/application. Own under5m/full under15m; do not raise bounds.
This is the release slice's immediate caller; no third release behavior change.
First application after the passing correction gate; normal checkpoint bounds.
