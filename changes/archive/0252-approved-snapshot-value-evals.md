---
eval_fix:
    - FLOW-019
    - FLOW-139
module: workflow
phase: applied
run: 20261113T181346-47231e30e8af21ec3057e102
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
## Audit

No findings; checked FLOW-019, FLOW-139

## Provenance

- evals 2026-11-13T18:18:23.601549Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=200 failing=0 passing=200
- audit 2026-11-13T18:18:50.902562Z agent=codex fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-13T18:34:06.843973Z agent=codex fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=739 failing=0 passing=739

