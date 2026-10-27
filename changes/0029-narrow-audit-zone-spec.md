---
already_satisfied:
    - CHK-004
    - new:current-change
    - new:impl-append-only
    - new:evals-append-only
    - new:rename-paths
bases:
    CHK-004: 5783c034
module: checks
---



## Motivation

Finish the human-directed CHK-004 split. CHK-006 through CHK-020 now retain its
path prohibitions and phase-specific zone permissions. Give its remaining rules
separate IDs and narrow CHK-004 to audit's current-file restriction. All behavior
already works and all assertions already exist; implementation stays unchanged.

## Interface



## Added

### new:current-change
During evals and impl, Zones MUST allow edits to the current change file.

### new:impl-append-only
During impl, Zones MUST allow only newly added files in the target module's append_only zone.

### new:evals-append-only
During evals, Zones MUST reject edits in the target module's append_only zone.

### new:rename-paths
Zones MUST check both the source and destination of a rename against phase permissions.

## Modified

### CHK-004
During audit, Zones MUST allow edits only to the current change file.

## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Keep 28 reported checks
cases and every assertion. Keep CHK-004 on AuditOnlyCurrentChange; remove it from
unrelated cases now covered by dedicated IDs. Retarget the existing current-file
assertions in EvalsAllowedZonesAndNewStubs and ImplFreezesEvalsAndInterfaces,
ImplAppendOnlyNewVersusExisting, and RenamesCheckBothPaths to the Added IDs.
Keep the checks for both directions of foreign-module and append-only renames.
No code, fixture, helper or interface edits are needed.
Impl may edit only this change file.
