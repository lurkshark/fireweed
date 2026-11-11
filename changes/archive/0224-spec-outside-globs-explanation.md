---
module: assets
phase: applied
run: 20261111T193939-f48e0648f90b9af622e22113
---

## Motivation

0225 adds the design's warning for a SPEC.md outside the configured module
globs and gives it the unallocated code C20. Every emitted code needs an
`explain` entry with a corrective action.

## Interface

None. Explain gains the C20 entry.

## Added

### AST-059
Explain MUST describe C20 as a warning for a SPEC.md file outside the configured modules globs and advise adding a matching glob or moving or renaming the file.

## Modified

## Removed

## Decisions

## Notes

Run before 0225.

Fail-first verification passed for the evals phase. The AST-059 eval failed
because `assets.Explain("C20")` returned `unknown diagnostic code "C20"`.
Add the C20 explanation in `packages/assets/internal/content/content.go` during
implementation.
## Provenance
- evals 2026-11-11T19:42:49.780695Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=126 failing=1 passing=125
- audit 2026-11-11T19:43:07.714329Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-11T20:10:19.283479Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=1066 failing=0 passing=1066
## Audit

No findings; checked AST-059
