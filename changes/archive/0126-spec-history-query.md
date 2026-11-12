---
audit_rounds: 0
bases: {}
module: project
phase: applied
run: 20261030T214954-ce905353e72b33c260a183b8
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

### PRJ-051
For a nonempty ID, SpecHistory MUST return the target module's archived changes whose Added or Modified items name that ID, including tag-only modifications.
Pending changes and archives belonging to other modules or unrelated spec IDs MUST not appear in that selection.

### PRJ-052
Every history entry MUST identify its archive path and module, the matching IDs, its Motivation, added and removed Decisions, and its complete Provenance section.

### PRJ-053
SpecHistory MUST order entries by their latest valid implementation Provenance timestamp, newest first, with missing or invalid implementation dates after dated entries.
Equal-date or undated entries MUST use archive-path lexical order as a total-order fallback.

### PRJ-054
With an empty ID filter, SpecHistory MUST return each target-module archive containing Added or Modified spec edits once, with its unique affected IDs in lexical order.
A module with no matching archives MUST return an empty slice.

### PRJ-055
An unknown module or an ID absent from both current module Specs and matching archive edits MUST return U02 with the rejected name.
A current ID with no matching archive edits MUST return an empty slice.
An ID in matching archive edits MUST remain queryable when it no longer exists in current Specs.

### PRJ-056
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

## Provenance

- evals 2026-10-30T22:08:15.168784Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=681 failing=11 passing=670
- audit 2026-10-30T22:09:19.244928Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-30T22:20:16.386462Z agent=manual fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=681 failing=0 passing=681

## Audit

No findings
Checked IDs: PRJ-051, PRJ-052, PRJ-053, PRJ-054, PRJ-055, PRJ-056
