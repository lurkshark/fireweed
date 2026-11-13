---
bases: {}
eval_fix:
    - RUN-010
    - RUN-011
    - RUN-016
    - RUN-023
    - RUN-024
    - RUN-025
    - RUN-026
    - RUN-030
    - RUN-035
    - RUN-040
module: runner
phase: applied
run: 20261113T002202-afc2c1f1f78dc3296d2c2cec
---

## Motivation

The owner approved completing structured fail-first diagnostics before regeneration. SYSTEM.md makes diagnostic codes structured metadata, but runner evals still pin the dependency's `P2-01:` message prefix. Correct those assertions before checks removes that prefix, retaining the verification cause, retry feedback, blocking and stop assertions.

## Interface

Audited eval_fix only; no API, runtime or spec changes.

## Added



## Modified



## Removed

## Decisions

## Notes

Replace dependency-error prefix checks in next_test.go, run_test.go, usage_test.go and time_limit_test.go with contracts.ErrorCode plus the existing identifying cause and ref checks. Retry prompt assertions should require the fail-first cause and affected ID without requiring an error-label prefix; diagnostic output may still contain its separate code. Preserve all stop classification, work retention, queue and invocation assertions. All runner evals must pass against the current dependencies. Implementation edits only this change file.
Owner authorized these five regeneration prerequisites and their necessary module-local changes. Stop after this queue; no regeneration or drill is authorized. Follow SYSTEM.md and the controller context. Workers must not read docs/rebuild/PROGRESS.md or ACCEPTANCE.md. Preserve unrelated assertions and retained interfaces after evals. Use existing public-interface fixtures and focused assertions; do not redesign helpers or add a test matrix. Evals and implementation workers run fireweed context for this change before each step and fireweed verify for verification. Restricted auditors follow only the supplied audit playbook and inputs. Only the controller advances phases and commits.
## Provenance
- evals 2026-11-13T00:31:35.17234Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=170 failing=0 passing=170
- audit 2026-11-13T00:32:13.191304Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-13T00:44:27.012503Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=493 failing=0 passing=493
## Audit

No findings; checked RUN-010, RUN-011, RUN-016, RUN-023, RUN-024, RUN-025, RUN-026, RUN-030, RUN-035, RUN-040
