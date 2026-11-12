---
eval_fix:
    - APP-091
    - APP-092
    - APP-093
    - APP-130
    - APP-131
module: app
phase: applied
run: 20261111T181540-8d74109a616aa89607d012cd
---

## Motivation

0222 removes workflow's diagnostic labels from error text. Its app verification
fails because the release and PR body fixtures inspect text to identify the
P02 and P01 errors returned by workflow. The assigned codes are already
available through `contracts.ErrorCode`.

## Interface

Audited app eval correction. Existing CLI behavior and coverage are preserved.

## Added

## Modified

## Removed

## Decisions

## Notes

In the release fixture, assert P02 through `contracts.ErrorCode` for the
`workflow.Current` call after release. In the PR body fixture, assert P01
through `contracts.ErrorCode` for the unavailable revision range and retain
the rejected revision in the diagnostic cause. Keep all CLI exit-status,
plain and JSON diagnostics, claim, ref, and dirty-work checks. This change
precedes 0222-workflow-unlabeled-errors.
## Provenance
- evals 2026-11-11T18:26:07.665717Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=323 failing=0 passing=323
- audit 2026-11-11T18:26:29.433145Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-11T18:36:16.98751Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=323 failing=0 passing=323
## Audit

No findings; checked APP-091, APP-092, APP-093, APP-130, APP-131
