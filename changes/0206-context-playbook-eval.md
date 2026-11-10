---
module: workflow
eval_fix:
    - FLOW-005
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
