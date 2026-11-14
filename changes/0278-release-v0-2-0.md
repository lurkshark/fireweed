---
bases:
    AST-001: 032ffc8f
module: assets
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
