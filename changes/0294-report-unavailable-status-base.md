---
already_satisfied:
    - FLOW-110
bases:
    FLOW-110: aefba67c
module: workflow
phase: audit
run: 20261116T231512-351afb748a2afa7645895f8a
---

## Motivation

Status reports unapproved changes against the resolved local base. The design's
Git behavior uses the known remote tip with a local fallback, and APP-126
requires a corrective P01 when the base is unavailable. The retained workflow
build swallowed base resolution failure and returned a successful report
with a zero count. Callers need to distinguish unavailable data from a
genuine zero.

## Interface

## Added

## Modified

### FLOW-110
Status MUST count applied changes archived on HEAD whose change file is absent
from the locally known configured remote/base ref, falling back to the local
base when the remote ref is absent. The base and remote MUST use the existing
configured/default base-ref resolution. If neither ref exists, Status MUST
return an error identifying the unavailable base.

## Removed

## Decisions

## Notes

Add a public Status case with a configured base missing in both locations.
Assert failure naming the base, even with an empty archive. Keep the existing
remote/local and archive-identity coverage. Current Status already propagates
the resolver error. APP-126 owns the CLI P01 mapping.
## Provenance

- evals 2026-11-16T23:19:00.571363Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=215 failing=0 passing=215

