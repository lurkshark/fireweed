---
bases:
    AST-001: 032ffc8f
module: assets
phase: audit
run: 20261114T174256-d44f5cbeb1a34dc325317a7f
---

## Motivation

Prepare Fireweed v0.2.0 from current origin/main. The release workflow checks
the built binary's version against the release tag, and source installs and
new-project pins use the same public version.

## Interface

## Added

## Modified

### AST-001
Version MUST return 0.2.0.

## Removed

## Decisions

## Notes

Update the direct AST-001 version assertion for fail-first evidence. Keep the
top-level help version consistent with Version, using the same version
constant. The project file format remains 1. The audited fixture corrections
precede this change so their existing assertions run on both release series.
Repository version pins and release documentation are separate author edits.

`Test_AST_001_Version` failed against the current implementation with `0.1.0`
instead of `0.2.0`. `fireweed verify release-v0-2-0` passed in evals.
The implementation step can update `packages/assets/internal/content/content.go`,
which supplies both `Version()` and the top-level help heading.
## Provenance

- evals 2026-11-14T17:45:33.795698Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=149 failing=1 passing=148

