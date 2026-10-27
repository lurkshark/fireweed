---
bases: {}
eval_fix:
    - FLOW-001
    - FLOW-006
    - FLOW-007
    - FLOW-008
    - FLOW-009
module: workflow
---



## Motivation

Human review 3 directed the split to retarget existing tests rather than add
cases. Change 0015 instead separated assertions into extra Test functions and
duplicated fixtures. Correct that eval rewrite while preserving every assertion
and the coverage of the newly assigned spec IDs.

## Interface



## Added



## Modified



## Removed



## Decisions



## Notes

Restore the pre-0015 test structure from
1fe462f:packages/workflow/evals/lifecycle_test.go as the reference for this eval
file only. Reading this current test revision is allowed; implementation history
is not an input. Keep all stronger assertions added by 0015, but merge them into
the existing cases and reuse the original fixture instances.

The original start case should carry FLOW-001, FLOW-006, FLOW-007 and FLOW-008 in
its testcase name. The original approval/bound case should carry FLOW-001 and
FLOW-009. Remove the redundant standalone snapshot/ID/run test functions and the
extra unapproved fixture; its P04 outcome is already asserted in the original
below-bound scenario. Restore 32 reported workflow cases with no lost assertions.
No new Test functions, t.Run cases, fixture instances, public interfaces or
implementation changes. The impl phase may edit only this change file.
