---
already_satisfied:
    - PRJ-007
    - new:duplicate-edits
    - new:undeclared-handles
    - new:stamped-bases
    - new:absent-edit-targets
bases:
    PRJ-007: f92e1669
module: project
---



## Motivation

Complete the approved PRJ-007 split under SYSTEM's spec-shape rule. PRJ-008–012
retain its document rules; give the remaining change validation rules their own
IDs and narrow PRJ-007 to diagnostic results. Remove its blanket C01 allocation
so D006 can later assign stable cause-specific codes through audited eval_fix
and normal changes. Preserve the design's Change files validation obligations;
all current behavior already satisfies these specs.

## Interface



## Added

### new:duplicate-edits
Check MUST reject a change that uses the same edit identity more than once across Added, Modified and Removed.

### new:undeclared-handles
Check MUST reject a pending change that references a handle absent from that change's Added entries.

### new:stamped-bases
Check MUST reject Modified or Removed entries whose bases hash is missing or differs from the preceding projected spec's canonical hash.

### new:absent-edit-targets
Check MUST reject Modified or Removed IDs absent from the projection preceding the change.

## Modified

### PRJ-007
Check MUST return no error diagnostics for valid module specs and changes. Each failure diagnostic MUST have a nonempty code and message.

## Removed



## Decisions



## Notes

Reuse the existing CheckValidAndDiagnostics table. Retarget duplicate_edit_identity,
undeclared_handle, missing_base/stale_base and absent_modified/absent_removed to
the corresponding Added IDs. Retain PRJ-007 on the valid/error diagnostic
contract and retain every PRJ-008–012 label and all assertions from 0030.
No new Test functions, t.Run cases or fixture instances; keep 53 project cases.
Keep the existing C01 assertions for now: their correction belongs to a later
eval_fix, and assigned D006 codes belong to later normal changes. Do not change
product code, helpers, fixtures or interfaces.
Implementation may edit only this change file.
