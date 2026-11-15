---
bases:
    CHK-063: d58f1f3a
    CHK-064: 5165fb12
module: checks
phase: impl
run: 20261115T023040-7d6eb1d324f8a001fd2a1e78
---

## Motivation

Projects using named commands need doctor to probe the commands that lifecycle
workers and auditors will actually use.

## Interface

## Added

### CHK-085
Doctor MUST report an agent selection error as a failed C15 probe identifying the setting at fault and skip invocation for that unresolved role.

## Modified

### CHK-063
Doctor MUST resolve the worker and audit commands from the project configuration and run each resolved command once in scratch through agent.wrap with a fixed-reply prompt, requiring exit 0 and the requested reply for success.

### CHK-064
When the audit probe inherits the worker default, Doctor MUST probe that command for audit and emit a C15 warning identifying the shared-command blind spot.

## Removed

## Decisions

## Notes

Move the doctor eval fixtures to named commands and retain their existing
probe, wrapper, prompt, timeout, and log assertions.

Fail-first verification passed for CHK-063, CHK-064, and CHK-085. The
current Doctor does not invoke the named commands or identify
`agent.default` and `audit.agent` in its C15 failures. All three behaviors
can be implemented in `packages/checks/internal/engine/doctor.go`.
## Audit

No findings; checked CHK-063, CHK-064, CHK-085

## Provenance

- evals 2026-11-15T02:39:54.496149Z agent=codex fireweed=0.2.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=135 failing=23 passing=112
- audit 2026-11-15T02:40:21.790724Z agent=codex fireweed=0.2.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

