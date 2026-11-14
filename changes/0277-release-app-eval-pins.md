---
eval_fix:
    - APP-016
    - APP-074
    - APP-080
    - APP-136
    - APP-163
    - APP-164
module: app
phase: impl
run: 20261114T172312-49a8a7c79772e2b1e728724c
---

## Motivation

Shared app and binary lifecycle fixtures pin 0.1. The v0.2 binary would reject the fixtures before exercising their intended commands. Compatibility evals also need a version-appropriate supported pin while retaining deliberate incompatibility.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Make the common fixtureBinaryPin an exact compatible comparison derived from assets.Version(). Reuse it in appFixtureConfig and the next, run and drill CLI fixtures. Preserve the incompatible configuration substitutions, assertion strength and supported or incompatible historical and approved cases.
This is an eval-only correction before the AST-001 version bump.
## Audit

No findings; checked APP-016, APP-074, APP-080, APP-136, APP-163, APP-164

## Provenance

- evals 2026-11-14T17:35:07.05149Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=395 failing=0 passing=395
- audit 2026-11-14T17:35:29.068154Z agent=codex fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

