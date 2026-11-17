---
module: runner
eval_fix:
    - RUN-072
---

## Motivation

The RUN-072 eval must identify the warning that Next is required to retain.
In the workflow drill report in commit `f6f5997`, Start returned a P04
approval warning and a P05 base-tip warning, both naming
`refs/remotes/origin/main`. The eval selected the first diagnostic by ref
text alone and stopped before exercising Next.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

In `Test_RUN_072_RetainsStartWarning` and its shared matcher, identify the
base-tip diagnostic by P05 and the selected ref together. Keep the warning
flag check and the full diagnostic equality and uniqueness assertions across
the existing return paths. Keep the separate P04 retention and worker
recovery warning assertions. With this matcher, the focused eval and all its
subtests passed on both main's implementation and the retained reconstruction
`16c778e`.
