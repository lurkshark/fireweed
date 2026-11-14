---
eval_fix:
    - RUN-011
    - RUN-016
    - RUN-050
module: runner
phase: applied
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
## Audit

No findings; checked RUN-011, RUN-016, RUN-050

## Provenance

- evals 2026-11-14T17:10:54.747571Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=186 failing=0 passing=186
- audit 2026-11-14T17:11:23.178148Z agent=codex fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-14T17:23:10.565549Z agent=codex fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=581 failing=0 passing=581

