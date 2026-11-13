---
bases:
    CHK-025: 5d1f3fad
module: checks
phase: impl
run: 20261113T004454-f253704169dbf8f3da43df2c
---

## Motivation

The owner approved removing the last fail-first message-prefix exception before regeneration. CHK-025 still requires a P2-01-prefixed error while SYSTEM.md requires codes in contracts.CodedError and unlabeled cause text. Return the missing-evidence code through ErrorCode so callers need no text compatibility path.

## Interface

No signatures change. RequireFailFirst returns a contracts.CodedError for missing genuine evidence.

## Added



## Modified

### CHK-025
RequireFailFirst MUST find a failed or errored referencing case for every requested enforced ID except already_satisfied; passing, skipped, empty or unrelated failures do not qualify. Missing genuine fail-first evidence MUST return an error whose contracts.ErrorCode is P2-01 and whose subject identifies the uncovered spec ID.


## Removed

## Decisions

## Notes

Update the existing fail-first assertion helper to require the structured code and affected ID. Preserve all fail-first acceptance/rejection cases. Remove the leading label from the cause under SYSTEM.md. The preceding runner eval_fix removes downstream prefix assumptions; workflow already wraps the failure as P2-01 and must retain its cause. No weakening of fail-first or unrelated diagnostic changes.
Owner authorized these five regeneration prerequisites and their necessary module-local changes. Stop after this queue; no regeneration or drill is authorized. Follow SYSTEM.md and the controller context. Workers must not read docs/rebuild/PROGRESS.md or ACCEPTANCE.md. Preserve unrelated assertions and retained interfaces after evals. Use existing public-interface fixtures and focused assertions; do not redesign helpers or add a test matrix. Evals and implementation workers run fireweed context for this change before each step and fireweed verify for verification. Restricted auditors follow only the supplied audit playbook and inputs. Only the controller advances phases and commits.

Eval verification passed with genuine CHK-025 fail-first evidence. The current error exposes neither a structured code nor a subject and still prefixes its cause with `P2-01:`. The implementation change belongs in `RequireFailFirst` in `packages/checks/internal/engine/junit.go`.
## Provenance
- evals 2026-11-13T00:54:12.088617Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=130 failing=3 passing=127
- audit 2026-11-13T00:54:30.029329Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked CHK-025
