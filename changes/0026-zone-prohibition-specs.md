---
already_satisfied:
    - new:protected-edits
    - new:spec-edits
    - new:foreign-module-edits
    - new:foreign-change-edits
    - new:outside-worker-writable
bases: {}
module: checks
---



## Motivation

Begin the human-directed CHK-004 split with independent path prohibitions and
the outside-module exception. Each rule already works. Reuse existing zone
evals and keep CHK-004 until its remaining phase permissions have dedicated IDs.

## Interface



## Added

### new:protected-edits
Zones MUST reject edits to protected paths in every phase.

### new:spec-edits
Zones MUST reject edits to module SPEC.md files in every phase.

### new:foreign-module-edits
Zones MUST reject edits within another module in every phase.

### new:foreign-change-edits
Zones MUST reject edits to change files other than the current change in every phase.

### new:outside-worker-writable
During evals and impl, Zones MUST allow edits outside modules only when the path matches worker_writable, except for the current change file.

## Modified



## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Keep 28 reported checks
cases and every assertion. Retarget the existing every-phase prohibitions test
and WorkerWritableOnlyEvalsAndImpl test. The current-change permission is
preserved in CHK-004 and will receive its own ID before narrowing. Keep CHK-004
labels until then. No code, fixture, helper or interface edits are needed.
Impl may edit only this change file.
