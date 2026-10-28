---
already_satisfied:
    - CHK-005
    - new:boundary-required-hook
    - new:boundary-optional-hook
bases:
    CHK-005: 26a5d821
module: checks
---



## Motivation

Finish the identified CHK-005 split under SYSTEM's one-behavior spec shape.
CHK-039–043 retain violations, opt-out warnings, suppression and failed-hook
rules. Give missing-hook requirements their own IDs and retain applicable hook
execution under CHK-005, preserving the design's Language support boundary
requirements. Existing public evals already assert all these behaviors.

## Interface



## Added

### new:boundary-required-hook
Boundaries MUST fail for a missing hook required by a module with disposable globs and no reasoned opt-out.

### new:boundary-optional-hook
Boundaries MUST warn without an error for a missing hook when no boundary enforcement is required.

## Modified

### CHK-005
Boundaries MUST run applicable global and per-module boundary hooks.

## Removed



## Decisions



## Notes

Keep CHK-005 on GlobalAndModuleBoundaryHooks, which checks both execution
markers, and preserve CHK-039 there. Remove CHK-005 from ExactBaselineOnly,
OptOutSuppressesOnlyItsModuleAndWarns, NormalNonzeroExitCanUseBaselineAndOptOut
and LaunchCancellationAndTimeoutStayErrors; their dedicated IDs remain.
MissingHookRequiredVersusOptional uses the two Added IDs instead of CHK-005,
retaining its disposable, nondisposable and reasoned-opt-out assertions.
Rename tests only. Keep all 28 checks cases and every assertion, fixture,
helper and other ID. No product code or interface edits. Implementation may
edit only this change file.
