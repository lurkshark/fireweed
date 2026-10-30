---
bases: {}
module: project
---



## Motivation

The design Specs/Spec history clause requires show <module> <ID> --history to
explain a spec through the archived changes that added, modified or retagged it,
newest first with Motivation, Decisions and provenance. No public project query
supplies this view yet. Build only that query, complete the mandatory fifth-change
review/checks and push on PASS, then immediately wire app/actual NEW proof and
stop with that caller local, as explicitly approved by the owner.

## Interface

Add SpecHistory(root, module, id string) ([]SpecHistoryEntry, error) through a
thin api.go facade and model.SpecHistoryEntry. Fields: Path, Module string;
IDs []string; Motivation, Decisions, RemovedDecisions, Provenance string;
ImplementationDate time.Time. Behavior belongs under internal/. Empty id is the
module-wide view needed by show <module> [ID] --history. No dependency changes.
An evals-only new compiling stub may return a plain not-implemented error.

## Added

### new:history-selection
For a nonempty ID, SpecHistory MUST return the target module's archived changes whose Added or Modified items name that ID, including tag-only modifications.
Pending changes and archives belonging to other modules or unrelated spec IDs MUST not appear in that selection.

### new:history-entries
Every history entry MUST identify its archive path and module, the matching IDs, its Motivation, added and removed Decisions, and its complete Provenance section.

### new:history-date-order
SpecHistory MUST order entries by their latest valid implementation Provenance timestamp, newest first, with missing or invalid implementation dates after dated entries.
Equal-date or undated entries MUST use archive-path lexical order as a total-order fallback.

### new:history-module-view
With an empty ID filter, SpecHistory MUST return each target-module archive containing Added or Modified spec edits once, with its unique affected IDs in lexical order.
A module with no matching archives MUST return an empty slice.

### new:history-lookup
An unknown module or an ID absent from both current module Specs and matching archive edits MUST return U02 with the rejected name.
A current ID with no matching archive edits MUST return an empty slice.
An ID in matching archive edits MUST remain queryable when it no longer exists in current Specs.

### new:history-input-errors
Missing or malformed fireweed.yaml MUST produce a C01 error.
Malformed archived change documents or malformed Added/Modified item syntax MUST produce a C03 error identifying the archive path.

## Modified



## Removed



## Decisions

History reads the project's current archive files, as the archive-backed show
view; it does not infer chronology from commit ancestry or filenames. Entry
metadata are Markdown section contents without their level-two heading; a
nonempty ID filter returns that matching ID in each entry's IDs field.

## Notes

Owner-approved two-change continuation: this is count4 before/count5 after.
Fresh fixed Sol/high eval/impl workers, isolated restricted assertion auditor,
normal max2 rounds/re-audit only prior findings. Own Verify under5m/full under15m;
representative public fixtures only, no instrumentation/calibration or convention
observer. Workers never commit/alter phase/claims/refs; controller owns native
commits/application, with write-capable Done. No bootstrap source/tests, old
history/implementation oracle, protected/other-module edits or dependency changes.

After native archive, stop authoring for the fifth gate over0122–0126: separate
Sol/high quality/goal subagent review, fresh appropriate build/vet/race/static
and explicit nonempty controller history extension79f7113..reviewed source.
Any finding/block stops for owner. Push only on PASS and verify remote OID; review
resets count, not push. Then immediately app show --history/NEW proof and stop
with caller local at count1. No phase exception, drill, GC or further push.
