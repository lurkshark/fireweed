---
eval_fix:
    - RUN-011
    - RUN-016
    - RUN-050
module: runner
phase: audit
run: 20261114T170042-73ffdf980e2d463efe8a127a
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
## Provenance

- evals 2026-11-14T17:10:54.747571Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=186 failing=0 passing=186

