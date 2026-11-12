---
audit_rounds: 0
bases: {}
module: app
phase: applied
run: 20261029T204957-38077114aff7e59b34419d29
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

### APP-091
release <change> MUST discover the nearest project root and invoke
workflow.Release for the named change. On success it MUST print confirmation
with exit0, preserving work, runtime state and HEAD while dropping only the
owned claim.

### APP-092
release MUST return exit1 with the workflow cause code and corrective action
when claim release fails, including missing or foreign claims. It MUST return
exit2 with a corrective diagnostic for usage or project configuration errors.

### APP-093
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

## Provenance

- evals 2026-10-29T20:59:19.983187Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=203 failing=7 passing=196
- audit 2026-10-29T21:01:10.109033Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-29T21:09:10.527796Z agent=manual fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=203 failing=0 passing=203

## Audit

No findings
Checked IDs: APP-091, APP-092, APP-093
