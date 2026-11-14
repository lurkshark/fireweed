---
eval_fix:
    - PRJ-001
    - PRJ-084
module: project
phase: impl
run: 20261114T153519-f07e7cb5a0b03fdcc8941f56
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
## Audit

No findings; checked PRJ-001, PRJ-084

## Provenance

- evals 2026-11-14T15:36:13.705225Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=278 failing=0 passing=278
- audit 2026-11-14T15:36:37.207307Z agent=codex fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

