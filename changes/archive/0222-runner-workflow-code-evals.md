---
eval_fix:
    - RUN-011
    - RUN-012
module: runner
phase: applied
run: 20261111T183620-87ba2d6c8314ee9d90ae4223
---

## Motivation

0222 removes workflow's diagnostic labels from error text. Its runner
verification fails because the Next start-failure fixture looks for the P01
and P07 labels in propagated workflow error text. The assigned codes are
already available through `contracts.ErrorCode`.

## Interface

Audited runner eval correction. Existing Next behavior and coverage are preserved.

## Added

## Modified

## Removed

## Decisions

## Notes

In `Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents`, assert P01 through
`contracts.ErrorCode` for dirty start and P07 for the unapproved archive
limit. Preserve the failure conditions, no-agent checks, claim and approval
checks, and all other subtests. This change precedes 0222-workflow-unlabeled-errors.
## Provenance
- evals 2026-11-11T18:48:34.842995Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=167 failing=0 passing=167
- audit 2026-11-11T18:48:54.687472Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-11T19:04:43.696476Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=490 failing=0 passing=490
## Audit

No findings; checked RUN-011, RUN-012
