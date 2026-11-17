---
eval_fix:
    - RUN-072
module: runner
phase: audit
run: 20261117T105320-b1f1eb8ec47052681a31dcfb
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
## Provenance

- evals 2026-11-17T11:03:07.677894Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=217 failing=0 passing=217

