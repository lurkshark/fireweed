---
bases:
    RUN-029: 5b18306a
module: runner
phase: applied
run: 20261111T072134-1273816dd7ab08e452ef8599
---

## Motivation

The design promises a stable code for every problem and output that agents
can read. RUN-029 rejects bad limit settings "with C01", but no runner error
carries a code that `contracts.ErrorCode` can read, so the app parses the
text instead (APP-078, APP-085). State RUN-029's code through
`contracts.ErrorCode` and apply SYSTEM.md's diagnostic-code convention to
runner's other coded errors.

## Interface

No signature changes. Errors that carry a code become (or wrap)
`contracts.CodedError`.

## Added

## Modified

### RUN-029
Next MUST read limit_wait and limit_budget as durations, defaulting to 30m and
4h when unset. It MUST reject malformed durations, a nonpositive wait or negative
budget with an error whose contracts ErrorCode is C01 before claiming a fresh
change or invoking an agent.

## Removed

## Decisions

## Notes

Apply the convention to the drill's configuration errors as well, and wrap
with `%w` so that a workflow or checks code stays readable through `Next`,
`Run` and `Drill`. Keep the existing message text until 0223.

Evals: assert `contracts.ErrorCode` for RUN-029's rejection and for one
`Next` failure that wraps a lifecycle code.

Fail-first: the focused runner evals found no structured C01 for malformed
wait or budget, zero or negative wait, or negative budget. The existing
evals-exhaustion scenario preserved the wrapped lifecycle code P2-01.
The usage-policy errors are in `packages/runner/internal/next/usage.go`
and the drill configuration errors are in `packages/runner/internal/drill/drill.go`.
## Provenance
- evals 2026-11-11T07:41:11.512385Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=167 failing=8 passing=159
- audit 2026-11-11T07:41:31.146594Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-11T08:06:39.327764Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=490 failing=0 passing=490
## Audit

No findings; checked RUN-029
