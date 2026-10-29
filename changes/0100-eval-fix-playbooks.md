---
bases: {}
module: assets
---



## Motivation

The design makes the same embedded playbooks guide interactive and unattended
steps, and eval_fix repairs evals while keeping implementation unchanged. NEW
now enforces that kind, but its impl playbook still directs the worker to edit
disposable code. Add the corresponding conditional instructions to the three
step playbooks. Close this small caller-facing slice by building NEW and reading
the evals/audit/impl playbooks through its actual CLI; reuse the passing0099 Go
eval-fix lifecycle proof rather than repeat it.

## Interface

No public API changes. Extend the existing Playbook(name) text for evals, audit
and impl with conditional eval_fix rules; preserve ordinary spec/fix guidance,
audit scope/convergence/example and worker no-commit instructions.

## Added

### new:eval-fix-evals-guidance
The evals playbook MUST instruct eval_fix workers to correct only evals, retain
coverage of listed IDs and require all module evals to pass instead of fail first.

### new:eval-fix-audit-guidance
The audit playbook MUST instruct eval_fix auditors to check that corrected tests
still assert the listed specs.

### new:eval-fix-impl-guidance
The impl playbook MUST instruct eval_fix workers to edit only the current change
file, leaving implementation unchanged.

## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

Use public assets.Playbook evals for observable returned instructions. Keep
requirements conditional on eval_fix; do not remove the normal behavior's
fail-first/implementation guidance. Preserve existing assertions and helpers,
including the passing Audit format example and restricted re-audit rules. No
new API, hook/recipe/version/diagnostic behavior, prose reorganization, matrix,
instrumentation, other-module/protected edits or drill. One focused assertion
per returned playbook can cover this slice. Evals/interfaces freeze in impl;
workers never commit or advance phases. Coordinator owns native verification,
isolated audit, clean NEW build and actual CLI guidance check. Own under5m,
full under15m; do not raise bounds. This is the second consecutive eval-fix
behavior change and completes its guidance follow-up; no third continuation
is authorized by this change.
