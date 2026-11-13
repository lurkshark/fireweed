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
