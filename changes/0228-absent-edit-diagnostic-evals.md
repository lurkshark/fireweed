---
eval_fix:
    - PRJ-016
module: project
---

## Motivation

PRJ-016 requires Check to report C08 for a Modified or Removed ID absent
from the preceding projection. Its eval requires two C08 diagnostics for
each case, one with the generic text `ID ALP-999 absent from preceding
projection` and one beginning `Modified ID` or `Removed ID`. The spec leaves
the number and wording open, and the project drill lists this as a gap.

## Interface

Audited eval_fix only. Edit only the project evals zone and this change.

## Added

## Modified

## Removed

## Decisions

## Notes

In `Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics` (queue_test.go), require
for each PRJ-016 case at least one non-warning C08 diagnostic whose message
names ALP-999. Leave the PRJ-014 and PRJ-015 cases as they are.
