---
eval_fix:
    - PRJ-011
    - PRJ-068
module: project
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
