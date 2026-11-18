---
eval_fix:
    - CHK-048
module: checks
---

## Motivation

The November 17 contracts drill returned C05 for a duplicate module item
and identified the offending ID, but the checks eval rejected the wording
"duplicate item heading". CHK-048 distinguishes malformed syntax from
duplicate identities by diagnostic code. The eval currently also requires
specific wording for both outcomes.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Keep the malformed-heading and duplicate-item fixtures in
`Test_CHK_048_EvaluateModuleItemCauses`. Assert C03 and C05 respectively
through `contracts.ErrorCode`, including that the code is present.
CON-017 covers the duplicate's structured subject at the parser boundary.
