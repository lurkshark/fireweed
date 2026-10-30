---
bases: {}
module: app
---



## Motivation

The design CLI table requires pr-body <range> to provide the review description.
Workflow0123 now supplies endpoint review data/Markdown/B04, and0124 explains B04.
Return immediately to this caller: expose that report through NEW, prove the
actual plain/JSON command in a disposable fixture, then stop with this caller
local at count4 under the approved three-change scope. No push or drill.

## Interface

No public API/model changes. Compose workflow.PRBody and the existing CLI
schema1/diagnostic/error helpers. Do not recompute review context or metrics.
Project discovery must not validate report data using a later live config in
place of the requested committed endpoint. Preserve coded report causes.

## Added

### new:pr-body-command
pr-body MUST discover the nearest project root and expose workflow.PRBody for the supplied explicit range.
Requested endpoint data MUST remain authoritative when later HEAD or working-tree metadata differs.

### new:pr-body-plain
Plain pr-body MUST print the report Body and show every returned diagnostic code and message.
Successful reports MUST return exit0, including empty selections and warning-only reports.

### new:pr-body-json
Successful pr-body --json MUST emit schema1 with string data equal to the report Body and retained diagnostic codes, messages, warning status and path when present.
It MUST leave stderr empty.

### new:pr-body-errors
Missing or extra ranges and invalid options MUST return exit2 with U01 and corrective guidance.
Project discovery or endpoint configuration-loading failures MUST return exit2 with the existing cause code or C01 and corrective guidance.
Unresolvable ranges MUST return exit1 with P01, the range resolution cause and corrective guidance.

### new:pr-body-json-errors
JSON pr-body failures MUST emit null data with cause-code/message diagnostics and empty stderr.

## Modified



## Removed



## Decisions



## Notes

Owner-approved final change in this slice, count3 before/count4 after. Fresh fixed
Sol/high eval/impl workers, isolated restricted audit, normal native checks and
unchanged max2 audit bounds. Own Verify under5m/full under15m. Workers never
commit/modify phase/claims/refs; controller owns application/archive.

Retarget only the obsolete APP-014 pr-body-under-construction examples to another
still-pending command, preserving every assertion. Extend only the existing
APP-001 convention test's blind Git-recorder loop with pr-body plain/JSON and a
valid empty HEAD..HEAD range, preserving its original zero-tool assertions and
hook/agent/fetch/push checks. Do NOT assert new command success in APP-001 before
implementation; dedicated new cases require success. No new observer/calibration.

Use representative public fixtures and existing output helpers. No API/model,
other-module/protected edits, extra feature, speculative error paths or history
oracle. Preserve all existing evals. Actual NEW proof includes endpoint selection,
review content, warning paths/flags, ordinary/exact-threshold/root-protected
artifact exclusions and invalid range. Source data are authored fixture inputs,
not a repeated real lifecycle or drill. Stop after proof, clean/unclaimed/local;
no push, extra application or phase exception.
