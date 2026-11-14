---
bases:
    AST-001: bda96ef8
module: assets
---

## Motivation

A release-number bump should be an edit to protected build metadata. The owner
asked to remove the recurring spec and eval edits while keeping this build at
0.2.0.

## Interface

## Added

## Modified

### AST-001
Version MUST return the stable MAJOR.MINOR.PATCH release version recorded in
`packages/assets/internal/content/VERSION` at build time.

## Removed

## Decisions

## Notes

Use an alternate valid VERSION value in a temporary build to demonstrate
fail-first behavior and keep the checked-in release value at 0.2.0. Preserve
the version heading in top-level help and keep it backed by the same value.
Keep the existing tag check and repository compatibility range.
