---
module: app
eval_fix:
    - APP-016
    - APP-074
    - APP-080
    - APP-136
    - APP-163
    - APP-164
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
