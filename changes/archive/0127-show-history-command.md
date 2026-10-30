---
already_satisfied:
    - APP-008
    - APP-025
audit_rounds: 0
bases:
    APP-008: a90726f5
    APP-025: cb2c30f2
module: app
phase: applied
run: 20261030T231146-e5304d598617e8a0a182fbc5
---



## Motivation

The design Specs/Spec history and CLI inventory require show <module> [ID]
--history to explain specs from archived changes. The project query now works;
wire it immediately through NEW with plain/schema1 JSON and actual proof, then
stop with this caller local at application1 after the passing0122–0126 review.
The owner explicitly approved this continuation and fifth-review/push sequence.

## Interface

No public interface or dependency changes. Add only app internal option parsing,
a thin project.SpecHistory caller and rendering. Preserve current/projected show
behavior. Qualify APP-008 and APP-025 to their existing non-history scope, reusing
all existing assertions as already_satisfied. Evals may extend only the existing
APP-001 blind PATH-recorder loop with history plain/JSON calls, preserving its
assertions and adding no second convention observer. No eval-phase behavior stub
is needed: --history is currently rejected as unsupported.

## Added

### APP-132
show <module> [ID] --history MUST discover the nearest project root, call project.SpecHistory with the optional ID filter and return its entries in order with exit0.
Plain output MUST include each archive path, module, matching IDs, implementation date, Motivation, added and removed Decisions and complete Provenance.
Empty history MUST print an explicit no-history message with exit0.

### APP-133
Successful show --history --json MUST emit schema1 with ordered array data, empty diagnostics and empty stderr.
Each row MUST contain path, module, ids, motivation, decisions, removed_decisions, provenance and implementation_date, the latter an RFC3339 timestamp or an empty string when undated.
Empty history MUST use an empty array.

### APP-134
--history MUST be accepted only with show and MUST reject combination with --projected, using exit2 and a corrective U01 diagnostic for either misuse.

### APP-135
History MUST return exit2 for lookup failures with U02 naming the rejected module or ID, or for missing or malformed project configuration with C01.
Other project.SpecHistory failures MUST return exit1 retaining their cause code and corrective guidance.
Failure output MUST contain code/message diagnostics: JSON uses null data and empty stderr; plain output uses empty stdout and stderr diagnostics.

## Modified

### APP-008
Without --history, show <module> [ID] MUST return the current module document or selected item, using the whole-queue project projection with --projected.

### APP-025
Without --history, show --json MUST return identical text as schema-1 string data with empty diagnostics and exit 0 on success.

## Removed



## Decisions

History JSON uses lowercase keys and preserves project query section strings;
plain output labels each section and identifies undated entries explicitly.

## Notes

Owner-approved immediate caller after fifth gate PASS and verified origin push
5610268. Independent exact Sol/high workers and isolated restricted assertion
auditor; max2 persisted rounds/re-audit only prior finding IDs. Own Verify <5m,
full <15m; representative fixtures, no extra instrumentation. Frozen contracts,
no protected/other-module edits, no worker commits/phase/claim/ref writes.
After normal archive, clean-build NEW and prove actual history output in one
disposable authored-archive fixture; record results and stop local at count1.
Owner-approved native cancellation/re-authoring corrected only clause wording;
the exact runtime-document restoration exception is consumed. No further phase
exception, second push, drill, GC or dependency expansion is authorized.
Stop after this caller completes, as the owner explicitly reaffirmed.

## Provenance

- evals 2026-10-30T23:17:29.438907Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=282 failing=18 passing=264
- audit 2026-10-30T23:18:54.542019Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-30T23:25:31.690301Z agent=manual fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=282 failing=0 passing=282

## Audit

No findings
Checked IDs: APP-132, APP-133, APP-134, APP-135, APP-008, APP-025
