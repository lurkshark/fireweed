---
already_satisfied:
    - new:frozen-evals
    - new:frozen-interface
    - new:disposable-zone
    - new:preserve-zone
    - new:other-zone
bases: {}
module: checks
---



## Motivation

Continue the human-directed CHK-004 split with one impl-phase permission rule
per zone. The existing ImplFreezesEvalsAndInterfaces case already asserts all
five. Keep CHK-004 until the remaining common and append-only rules are split.

## Interface



## Added

### new:frozen-evals
During impl, Zones MUST reject edits in the target module's evals zone.

### new:frozen-interface
During impl, Zones MUST reject edits in the target module's interface zone.

### new:disposable-zone
During impl, Zones MUST allow edits in the target module's disposable zone.

### new:preserve-zone
During impl, Zones MUST allow edits in the target module's preserve zone.

### new:other-zone
During impl, Zones MUST allow edits in the target module's other zone.

## Modified



## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Keep 28 reported checks
cases and every assertion. Retarget ImplFreezesEvalsAndInterfaces to the Added
IDs; both its rejected paths and all allowed zone paths are already asserted.
Keep CHK-004 until final narrowing. No code, fixture, helper or interface edits
are needed. Impl may edit only this change file.
