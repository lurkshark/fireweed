---
already_satisfied:
    - new:apply-metadata
    - new:apply-decisions
    - new:advance-phase-base
    - new:release-applied-claim
    - new:failed-done-state
bases: {}
module: workflow
---



## Motivation

Begin the previously identified FLOW-004 split under SYSTEM's one-behavior
spec shape. Preserve the design's Implementation step and claim lifecycle,
giving metadata, Decisions and state transitions focused IDs before narrowing
FLOW-004 in the next change. Existing behavior already satisfies these rules.

## Interface



## Added

### new:apply-metadata
Successful impl Done MUST apply the change's set values to module metadata.

### new:apply-decisions
Successful impl Done MUST apply Decisions attributed to the change slug and remove entries quoted in Removed decisions.

### new:advance-phase-base
Each successful Done MUST advance the phase base to its new commit.

### new:release-applied-claim
Successful impl Done MUST release only the exact held claim upon application.

### new:failed-done-state
Failed Done verification MUST leave the phase unadvanced and create no commit.

## Modified



## Removed



## Decisions



## Notes

Retain FLOW-004 until the next narrowing. Retarget existing lifecycle tests:
ImplDoneAppliesSpecsSetAndDecisionsAndArchives gains metadata/Decisions/claim
IDs; EvalsDoneUsesDefaultAuditOrDisabledImpl gains phase-base; and
FailedDoneKeepsPhaseAndHead gains failed-state. Preserve every assertion and
all 32 workflow cases. Reuse existing Done results in successful phase
transition tests to assert returned State.PhaseBase equals the new HEAD where
needed, including audit and application, without new fixtures or cases.
Existing FLOW-002 and CHK-016/017 retain the prior eval/interface freeze checks.
No new Test functions, t.Run cases, fixture instances, helper changes, product
code or interface edits. Implementation may edit only this change file.
