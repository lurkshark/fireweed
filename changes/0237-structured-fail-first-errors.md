---
bases:
    CHK-025: 5d1f3fad
module: checks
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
