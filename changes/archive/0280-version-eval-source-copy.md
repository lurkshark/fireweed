---
eval_fix:
    - AST-001
module: assets
phase: applied
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
## Audit

No findings; checked AST-001

## Provenance

- evals 2026-11-14T23:25:03.449493Z agent=manual fireweed=0.2.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=149 failing=0 passing=149
- audit 2026-11-14T23:25:39.292521Z agent=codex fireweed=0.2.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-14T23:27:29.855076Z agent=manual fireweed=0.2.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=1349 failing=0 passing=1349

