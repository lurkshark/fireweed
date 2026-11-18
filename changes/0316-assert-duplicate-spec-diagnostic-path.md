---
eval_fix:
    - PRJ-011
    - PRJ-068
module: project
phase: audit
run: 20261118T034323-7f7543935eaf0763fe615274
---

## Motivation

The November 17 contracts drill produced a C05 diagnostic on the correct
module SPEC.md, but the project eval rejected "duplicate item heading".
PRJ-011 specifies C05 for duplicates, and PRJ-068 specifies the affected
document's path. These fields identify the failure and where to correct it.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

For the duplicate-module-ID fixture in
`Test_PRJ_015_PRJ_016_CheckDiagnostics`, assert a non-warning C05 diagnostic
with a nonempty message and `Diagnostic.Path` equal to
`packages/alpha/SPEC.md`. Cover PRJ-011 and PRJ-068 and retain the existing
assertions for the other fixtures.
## Provenance

- evals 2026-11-18T03:44:26.950446Z agent=codex fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=322 failing=0 passing=322

