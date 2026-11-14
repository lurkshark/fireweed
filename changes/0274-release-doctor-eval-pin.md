---
eval_fix:
    - CHK-049
    - CHK-050
module: checks
phase: audit
run: 20261114T161212-45fe08f74d5fe5a47b140635
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
## Provenance

- evals 2026-11-14T16:21:06.493552Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=130 failing=0 passing=130

