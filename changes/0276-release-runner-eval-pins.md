---
module: runner
eval_fix:
    - RUN-011
    - RUN-016
    - RUN-050
---

## Motivation

The Next, Run and Drill fixtures pin the 0.1 release series, so a later binary would reject them before exercising the runner. The v0.2 release requires these fixtures to continue reaching the same behavior.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

In next_test.go, run_test.go and drill_test.go, derive the compatible fixture pins from assets.Version(). Preserve all assertions, test names and explicit failure scenarios.
This is an eval-only correction before the AST-001 version bump.
