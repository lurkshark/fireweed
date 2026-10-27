---
already_satisfied:
    - new:evals-zone
    - new:declared-interface
    - new:new-disposable
    - new:other-zone
    - new:preserve-zone
bases: {}
module: checks
---



## Motivation

Continue the human-directed CHK-004 split with one evals-phase permission rule
per zone. The current EvalsAllowedZonesAndNewStubs case already asserts them all.
Keep CHK-004 until the impl and remaining common permissions are split too.

## Interface



## Added

### new:evals-zone
During evals, Zones MUST allow edits in the target module's evals zone.

### new:declared-interface
During evals, Zones MUST allow edits in the target module's interface zone only when InterfaceDeclared is true.

### new:new-disposable
During evals, Zones MUST allow edits in the target module's disposable zone only for files absent from the phase base.

### new:other-zone
During evals, Zones MUST allow edits in the target module's other zone.

### new:preserve-zone
During evals, Zones MUST reject edits in the target module's preserve zone.

## Modified



## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Keep 28 reported checks
cases and every assertion. Retarget EvalsAllowedZonesAndNewStubs to the Added
IDs; it already checks the allowed paths and rejects an undeclared interface,
an existing disposable file and a preserved file. Preserve CHK-004 until final
narrowing. No code, fixture, helper or interface edits are needed.
Impl may edit only this change file.
