---
eval_fix:
    - AST-001
module: assets
phase: audit
run: 20261114T232425-eff71f3e2fc8ef0592585742
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
## Provenance

- evals 2026-11-14T23:25:03.449493Z agent=manual fireweed=0.2.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=149 failing=0 passing=149

