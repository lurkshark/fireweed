---
already_satisfied:
    - PRJ-024
bases:
    PRJ-024: 09d40177
module: project
phase: applied
run: 20261111T010825-6111527bc1bb48b7afeeb3ad
---

## Motivation

The design requires stable diagnostic codes, and APP-019 requires
`new-change` to report C13 for a duplicate slug. PRJ-024 requires the
rejection but names no code, so the project drill's first reconstruction
chose a different code and passed only after the retry showed it the app
failure. The code is part of the project boundary and belongs in its spec.

## Interface

No API changes. The current implementation already returns C13.

## Added

## Modified

### PRJ-024
NewChange MUST reject a slug already present in the pending queue or archive with diagnostic C13, without overwriting existing files.

## Removed

## Decisions

## Notes

The existing PRJ-024 evals should also check the C13 code through
`contracts.ErrorCode` for both the pending and the archived duplicate.
Do not assert the error's wording.

Fail-first: both PRJ-024 duplicate cases reject the slug and preserve the
existing file, but `contracts.ErrorCode` reports no code. The CLI supplies
C13 from the error text. The author should remove PRJ-024 from
`already_satisfied` so the controller requires this fail-first evidence.
The implementation can supply the code in
`packages/project/internal/fixture/fixture.go`.
## Provenance
- evals 2026-11-11T01:10:13.251641Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=134 failing=3 passing=131
- audit 2026-11-11T01:10:31.396721Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-11T01:44:04.695423Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=928 failing=0 passing=928
## Audit

No findings; checked PRJ-024
