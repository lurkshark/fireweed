---
module: assets
---

## Motivation

0225 adds the design's warning for a SPEC.md outside the configured module
globs and gives it the unallocated code C20. Every emitted code needs an
`explain` entry with a corrective action.

## Interface

None. Explain gains the C20 entry.

## Added

### new:c20-explanation
Explain MUST describe C20 as a warning for a SPEC.md file outside the configured modules globs and advise adding a matching glob or moving or renaming the file.

## Modified

## Removed

## Decisions

## Notes

Run before 0225.
