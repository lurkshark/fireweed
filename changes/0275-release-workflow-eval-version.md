---
eval_fix:
    - FLOW-006
    - FLOW-051
module: workflow
phase: impl
run: 20261114T164043-0e74efa5c66485cb358a1c6b
---

## Motivation

The shared lifecycle fixture and one provenance assertion pin 0.1. The v0.2 release needs lifecycle evals to exercise claiming and recording the running version with a compatible fixture.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Derive the shared lifecycle fixture pin and the provenance version assertion from assets.Version(). Keep the per-phase exact version assertion and every other provenance and lifecycle assertion.
This is an eval-only correction before the AST-001 version bump.
## Provenance

- evals 2026-11-14T16:44:35.643685Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=211 failing=0 passing=211
- audit 2026-11-14T16:45:44.940534Z agent=codex fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

## Audit

No findings; checked FLOW-006, FLOW-051

