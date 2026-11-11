---
eval_fix:
    - FLOW-005
module: workflow
phase: applied
run: 20261111T030507-4fdc9908fee978c87d994c52
---

## Motivation

FLOW-005 requires Context to include "the matching playbook". Its eval
checks this by looking for the literal heading `# Write evals`, which is
assets' wording, not anything workflow promises. When the assets drill
rebuilt the playbooks with a different heading, this workflow eval failed
even though Context still included the matching playbook. The eval should
compare against what assets publishes.

## Interface

Audited eval_fix only. Edit only the workflow evals zone and this change.

## Added

## Modified

## Removed

## Decisions

## Notes

In `Test_FLOW_005_FLOW_051_FLOW_052_FLOW_053_ContextAndProvenanceAndRetainedReports`
(packages/workflow/evals/lifecycle_test.go), replace the `"# Write evals"`
expectation with the full text returned by `assets.Playbook("evals")`, which
the same file already uses for provenance hashes. Keep every other
expectation unchanged.
## Provenance
- evals 2026-11-11T03:11:00.990913Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=181 failing=0 passing=181
- audit 2026-11-11T03:11:25.59818Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-11T03:34:20.702523Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=670 failing=0 passing=670
## Audit

No findings; checked FLOW-005
