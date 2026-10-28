---
already_satisfied:
    - new:boundary-violations
    - new:boundary-opt-out-warning
    - new:boundary-baseline
    - new:boundary-opt-out-suppression
    - new:boundary-failed-hooks
bases: {}
module: checks
---



## Motivation

Begin the identified CHK-005 split under SYSTEM's one-behavior spec shape.
Retain the design's Language support boundary suppression and hook-completion
rules under focused IDs. Keep CHK-005 until its remaining hook execution and
missing-hook rules are separated in the next change. Existing behavior and
public eval assertions already cover these requirements.

## Interface



## Added

### new:boundary-violations
Boundaries MUST return errors for unsuppressed module-attributed violations.

### new:boundary-opt-out-warning
Boundaries MUST return a visible warning containing the reason for a reasoned boundaries_opt_out.

### new:boundary-baseline
Boundaries MUST suppress only exact baseline lines from normally exited hooks.

### new:boundary-opt-out-suppression
Boundaries MUST suppress only the opted-out module's violation lines through a reasoned opt-out after a hook exits normally.

### new:boundary-failed-hooks
Boundaries MUST report hook launch, cancellation or timeout failures as errors.

## Modified



## Removed



## Decisions



## Notes

Retarget existing boundaries_test.go evals while retaining CHK-005. Add violation
ID to GlobalAndModuleBoundaryHooks; warning/opt-out IDs to
OptOutSuppressesOnlyItsModuleAndWarns; baseline ID to ExactBaselineOnly;
baseline/opt-out IDs to NormalNonzeroExitCanUseBaselineAndOptOut; and failed-hook
ID to LaunchCancellationAndTimeoutStayErrors. Preserve every assertion and
all 28 checks cases. No new Test functions, t.Run cases, fixture instances,
helpers, product code or interfaces. Implementation may edit only this change file.
