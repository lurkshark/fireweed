---
module: workflow
eval_fix:
    - FLOW-006
    - FLOW-051
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
