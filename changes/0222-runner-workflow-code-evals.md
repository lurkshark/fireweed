---
eval_fix:
    - RUN-011
    - RUN-012
module: runner
---

## Motivation

0222 removes workflow's diagnostic labels from error text. Its runner
verification fails because the Next start-failure fixture looks for the P01
and P07 labels in propagated workflow error text. The assigned codes are
already available through `contracts.ErrorCode`.

## Interface

Audited runner eval correction. Existing Next behavior and coverage are preserved.

## Added

## Modified

## Removed

## Decisions

## Notes

In `Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents`, assert P01 through
`contracts.ErrorCode` for dirty start and P07 for the unapproved archive
limit. Preserve the failure conditions, no-agent checks, claim and approval
checks, and all other subtests. This change precedes 0222-workflow-unlabeled-errors.
