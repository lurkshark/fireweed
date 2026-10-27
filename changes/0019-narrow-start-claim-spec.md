---
already_satisfied:
    - FLOW-001
    - new:initial-phase-base
    - new:retained-snapshot
bases:
    FLOW-001: 9eea47d3
module: workflow
---



## Motivation

Finish the review 3 FLOW-001 migration. Its other start behaviors now have
dedicated specs from 0015, 0017 and 0018. Preserve the remaining base/snapshot
guarantees under their own IDs and narrow FLOW-001 to claim creation. All behavior
already works; retarget existing evals without implementation changes.

## Interface



## Added

### new:initial-phase-base
Start MUST use the current HEAD as the initial phase base of its claim.

### new:retained-snapshot
While a claim is held, Current MUST return the project snapshot approved at Start.

## Modified

### FLOW-001
Start MUST atomically create the local claim for the selected queued change.

## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Keep the 32-case total.
Use the existing combined start case for both Added IDs. Its phase-base equality
assertion already covers the first. Retain the snapshot assertions; verify the
existing Current result contains the approved projected snapshot, not a newly
read live projection. If needed, reuse that same fixture to edit the existing
change's Added body before its existing Current call and compare the returned
snapshot with the approved body. Do not add a second fixture or test case.

Keep FLOW-001 only on existing assertions that prove claim creation/atomicity;
other assertions now carry their dedicated IDs. Preserve all prior checks.
No implementation or interface edits; impl may edit only this change file.
