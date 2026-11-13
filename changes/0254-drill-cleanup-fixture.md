---
module: runner
eval_fix: [RUN-062]
---

## Motivation

During regenerate-runner, Test_RUN_062_CleanupFailureCannotPass failed in
deferred fixture teardown because Git had dropped the worktree registration
although removing its directory failed. RUN-062 requires a cleanup error
instead of a passing drill, not surviving Git registration. The owner approved
repairing teardown so this eval measures that contract.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Audited eval_fix only. Edit only the runner evals zone and this change.

Correct teardown in Test_RUN_062_CleanupFailureCannotPass in
packages/runner/evals/drill_test.go. Restore fixture permissions and clean up
whether the scratch directory remains registered as a worktree or only as a
leftover directory. Use ordinary Git/filesystem operations on fixture paths;
do not instrument processes or require a particular Git removal order.

Preserve the fault injection and assertions that reconstruction was reached,
the fixture actually prevented directory removal, the returned error is
nonnil and Result is not pass. Preserve the other RUN-062 success/failure
cleanup scenarios. Do not suppress genuine teardown errors unrelated to the
already-removed registration.

Do not change implementation in this eval_fix. The later
simplify-drill-cleanup maintenance change separately assesses the permission
precheck that was introduced for this fixture.

Drafted and queued only. No lifecycle execution is authorized by the drafting
request.
