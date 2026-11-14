---
module: checks
eval_fix:
    - CHK-049
    - CHK-050
---

## Motivation

Doctor's shared fixture pins the 0.1 release series. The 0.2 binary would exit at static validation before reaching the canaries and agent probes that the fixture is meant to exercise.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Use an exact compatible pin derived from assets.Version() in doctorConfig. Preserve all assertions and fixture behavior.
This is an eval-only correction before the AST-001 version bump.
