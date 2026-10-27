---
already_satisfied:
    - new:held-snapshot
    - new:assigned-ids
    - new:run-id
    - new:approval-warning
bases: {}
module: workflow
---



## Motivation

Human review 3 requires one observable behavior per spec. Begin splitting
FLOW-001 without dropping its existing guarantees: establish dedicated IDs for
four working behaviors now, then narrow the original through subsequent changes.
Retarget the existing evals and reuse their fixtures. No implementation change
or new feature belongs in this change.

## Interface



## Added

### new:held-snapshot
Start MUST include the approved projected project snapshot in its returned State.

### new:assigned-ids
Start MUST replace each local Added handle throughout the change with the next available spec ID for its module.

### new:run-id
Start MUST assign a unique run identifier when starting an unstarted change.

### new:approval-warning
Start MUST report P04 when neither the local base nor the shared base contains the change file.

## Modified



## Removed



## Decisions



## Notes

This is the additive stage of the FLOW-001 split. A later normal change will
narrow FLOW-001 once every remaining behavior has a dedicated spec. Keep the
existing test scenarios; rename or retarget assertions to the assigned IDs.
Implementation is already satisfied and must remain unchanged.
