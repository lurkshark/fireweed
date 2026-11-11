---
eval_fix:
    - RUN-022
    - RUN-024
    - RUN-043
module: runner
phase: applied
run: 20261111T105608-3583ab308996d9fbe7ac9b35
---

## Motivation

RUN-022 and RUN-043 require retry prompts to carry the last verification
output, and RUN-024 requires the blocking error to keep the last
verification cause. Their evals check that through the checks module's
text `E01: failed case`, a code label and wording that no spec requires.
0221 removes the label from checks' error text, which would break these
evals for a reason unrelated to runner.

## Interface

Audited eval_fix only. Edit only the runner evals zone and this change.

## Added

## Modified

## Removed

## Decisions

## Notes

- `Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase`
  (next_test.go): check that each phase's retry prompt contains the failing
  fixture testcase's name instead of `E01: failed case`.
- `Test_RUN_023_RUN_024_ImplExhaustionSavesWorkAndBlocks`: check that
  `contracts.ErrorCode` on the returned error reports E01 and that its text
  names the failing testcase.
- `Test_RUN_043_RetryIncludesJUnitOnlyCasesFromOwnAndDependentModules`
  (retry_failure_details_test.go): keep the hook-marker check and replace
  `E01: failed case` with the failing testcase's name.

Run after 0212 and before 0221.
## Provenance
- evals 2026-11-11T11:09:17.124057Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=167 failing=0 passing=167
- audit 2026-11-11T11:09:35.690696Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-11T11:24:58.503473Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=490 failing=0 passing=490
## Audit

No findings; checked RUN-022, RUN-024, RUN-043
