---
module: workflow
eval_fix: [FLOW-019, FLOW-139]
---

## Motivation

During regenerate-workflow, the approved-snapshot eval failed solely because
empty resolved zone slices serialized as [] rather than null. FLOW-019 and
FLOW-139 require the approved snapshot and its independence from changed,
malformed or missing live inputs; they do not specify this representation
distinction. The owner approved narrowing the comparison to contractual
snapshot contents.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Audited eval_fix only. Edit only the workflow evals zone and this change.

In Test_FLOW_019_FLOW_139_ApprovedSurvivesLiveEdits in
packages/workflow/evals/approved_test.go, compare snapshot contents while
treating nil and empty resolved zone slices equivalently. Limit
normalization to those resolved zone collections.

Keep exact checks of all nonempty zone values, configuration, module
identity/dependencies/documents, canonical root, change slug/path and found
status. Keep all changed, malformed and missing live-input scenarios.

Do not normalize raw configuration or document frontmatter: PRJ-017's
omitted-versus-explicit-empty inheritance distinction must remain observable.
Do not change claim serialization or workflow implementation in this eval_fix.

Drafted and queued only. No lifecycle execution is authorized by the drafting
request.
