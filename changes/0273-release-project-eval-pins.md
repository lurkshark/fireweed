---
module: project
eval_fix:
    - PRJ-001
    - PRJ-084
---

## Motivation

The shared project fixture in project_test.go pins the 0.1 release series. A version bump would make ordinary project fixtures fail compatibility before they can exercise project behavior.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Use an exact compatible pin derived from assets.Version() in the shared fixture. Preserve explicit incompatible-pin tests and all existing assertions.
This is an eval-only correction before the AST-001 version bump.
