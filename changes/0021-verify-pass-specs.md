---
already_satisfied:
    - new:current-module-pass
    - new:dependent-selection
    - new:dependent-pass
    - new:boundary-pass
bases: {}
module: workflow
---



## Motivation

Continue the review-directed FLOW-002 split with independent full-verification
rules. Each already works. Reuse existing eval scenarios and keep FLOW-002 until
the last narrowing step. No new feature or implementation change belongs here.

## Interface



## Added

### new:current-module-pass
During impl Verify MUST reject a failing current-module testcase that references a non-quarantined spec.

### new:dependent-selection
Full impl Verify MUST use verify_dependents with a default of direct to select dependent modules.

### new:dependent-pass
Full impl Verify MUST reject a failed enforced testcase from a selected dependent module.

### new:boundary-pass
Full impl Verify MUST require boundary verification to succeed.

## Modified



## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Retarget the existing
impl dependent/boundary case to the corresponding IDs. The existing failed-Done
fixture has a failing current-module eval; reuse that fixture for the own-pass
assertion, adding a direct Verify call there only if needed to prove the gate.
Keep every existing assertion and all FLOW-002 labels until final narrowing.
Keep 32 reported workflow cases. No interface or implementation edits; impl may
edit only this change file.
