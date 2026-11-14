---
bases:
    AST-001: bda96ef8
module: assets
phase: audit
run: 20261114T224408-20a0d920b11800befbe81f5a
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

The alternate build reports 0.2.0 with a 0.2.1 VERSION before implementation.
The implementation step can update the disposable `internal/content/content.go`
and leave the protected VERSION file and retained facade in place.
## Provenance

- evals 2026-11-14T22:46:03.700257Z agent=manual fireweed=0.2.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=149 failing=1 passing=148

