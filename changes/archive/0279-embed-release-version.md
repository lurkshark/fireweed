---
bases:
    AST-001: bda96ef8
module: assets
phase: applied
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
- audit 2026-11-14T22:46:43.308527Z agent=codex fireweed=0.2.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-14T23:24:06.847735Z agent=manual fireweed=0.2.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=1349 failing=0 passing=1349

## Audit

No findings; checked AST-001

