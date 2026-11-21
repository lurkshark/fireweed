---
already_satisfied:
    - FLOW-002
bases:
    FLOW-002: 595d44bc
module: workflow
phase: applied
run: 20261120T215924-f9264a60be65231dcb9d6500
---

## Motivation

When a worker edits protected `fireweed.yaml` during a held step, Z01 tells
them to restore the protected file. The workflow drill returned live-config
C01 first for malformed YAML in both `verify` and `done` even though the
approved configuration passed admission. The dependent APP-170 eval caught
the diagnostic precedence difference (local report `9a94f39`).

## Interface

## Added

## Modified

### FLOW-002
Verify MUST reject edits forbidden by the approved projected phase zones before
running hooks, checking differences since the held phase base. For a held
step, Verify and Done MUST return an error whose contracts.ErrorCode is Z01
for an edit to protected fireweed.yaml even when loading the edited live file
would return C01.

## Removed

## Decisions

## Notes

Use the public workflow fixture to start a change, then replace the working
`fireweed.yaml` with malformed YAML. Call Verify and Done in owned fixtures
and assert Z01 and the `fireweed.yaml` diagnostic path. Keep the existing
APP-170 coverage for CLI admission against an incompatible approved pin.
A focused workflow API probe returned Z01 with the diagnostic path on the
drill base (`c97329b`) and C01 on the retained reconstruction (`79148d3`)
for both operations.
## Provenance

- evals 2026-11-20T22:03:42.879344Z agent=codex fireweed=0.2.4 playbook=5e05bacdae954379b468564c11429c5532016e26a7afbcd418f9bdb9441e151c tests=221 failing=0 passing=221
- audit 2026-11-20T22:04:00.325985Z agent=codex fireweed=0.2.4 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-20T22:20:02.198018Z agent=codex fireweed=0.2.4 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=856 failing=0 passing=856

## Audit

No findings; checked FLOW-002

