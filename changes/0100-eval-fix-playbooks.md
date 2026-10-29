---
audit_rounds: 0
bases: {}
module: assets
phase: impl
run: 20261029T181904-9eb88dd4f65008635fa9f934
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

### AST-029
The evals playbook MUST instruct eval_fix workers to correct only evals, retain
coverage of listed IDs and require all module evals to pass instead of fail first.

### AST-030
The audit playbook MUST instruct eval_fix auditors to check that corrected tests
still assert the listed specs.

### AST-031
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

## Provenance

- evals 2026-10-29T18:22:15.264704Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=569 failing=3 passing=566
- audit 2026-10-29T18:23:12.015658Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings
Checked IDs: AST-029, AST-030, AST-031
