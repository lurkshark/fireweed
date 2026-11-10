---
eval_fix:
    - FLOW-058
    - FLOW-062
module: workflow
---

## Motivation

FLOW-058 and FLOW-062 require workflow to keep the cause of a project or
checks failure. Their evals check that by matching the lower module's text,
including its leading code: `C08: missing base for`, `C10: no coverage for
ALP-001`, `C11: read JUnit report:` and `E01: failed case`. Neither the
code label nor the wording is in any spec. After 0209 and 0210 the codes are
readable through `contracts.ErrorCode`, and 0220 and 0221 remove the labels
from the text, which would break these evals for a reason unrelated to
workflow.

## Interface

Audited eval_fix only. Edit only the workflow evals zone and this change.

## Added

## Modified

## Removed

## Decisions

## Notes

In `Test_FLOW_058_FLOW_062_StaleLifecycleBaseKeepsCauseAndMissingBaseKeepsC08`
and `Test_FLOW_062_LifecyclePreservesEvalZoneAndBoundaryCauses`
(lifecycle_test.go), replace each prefix match on a project or checks error
with `contracts.ErrorCode` returning that code, plus a check for the
identifier the fixture broke (the spec ID, or the report path) where the
current assertion includes one. Do not assert the lower module's wording.

Leave the X02 handling in the FLOW-062 boundary tests, which already accept
an error with or without the label, and leave assertions on workflow's own
P codes, which 0222 changes.

Run after 0211 and before 0220.
