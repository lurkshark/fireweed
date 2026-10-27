---
already_satisfied:
    - new:no-start-commit
    - new:shared-work-bound
    - new:competing-owner
    - new:current-owner
bases: {}
module: workflow
---



## Motivation

Continue the review 3 FLOW-001 split with dedicated specs for four remaining
observable claim behaviors. Existing eval scenarios already assert these rules.
Retarget them without changing implementation or adding cases.

## Interface



## Added

### new:no-start-commit
Start MUST leave HEAD unchanged.

### new:shared-work-bound
Start MUST reject when the shared-base unapproved archive count reaches authoring.max_unapproved, whose default is 5.

### new:competing-owner
Start MUST reject a competing worktree that attempts to claim the same change.

### new:current-owner
Current MUST resolve only a claim owned by the calling worktree.

## Modified



## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Add the no-commit and
Current IDs to the existing combined start case. Retarget the existing bound
scenario to its dedicated ID. Add the competing-owner and Current IDs to the
existing linked-worktree case. Retain FLOW-001 until the final narrowing change.
Keep all assertions and the 32-case total. No implementation or interface edits;
the impl phase may edit only this change file.
