---
eval_fix:
    - RUN-022
    - RUN-024
    - RUN-043
module: runner
phase: audit
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
