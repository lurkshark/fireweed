---
eval_fix:
    - AST-001
module: assets
---

## Motivation

The release version eval must work with the Go 1.23 toolchain supported by
go.mod. Its build overlay does not replace embedded data on that toolchain.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Retain the public Version comparison with the protected VERSION metadata and
the build-time assertion with a different value. Build from a temporary source
copy with that different value. Leave the release number and implementation
as applied in 0279.
