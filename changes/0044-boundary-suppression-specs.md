---
already_satisfied:
    - CHK-039
    - CHK-040
    - CHK-041
    - CHK-042
    - CHK-043
audit_rounds: 0
bases: {}
module: checks
phase: audit
run: 20261028T054107-b7e31d3bfecf376475df71dc
---



## Motivation

Begin the identified CHK-005 split under SYSTEM's one-behavior spec shape.
Retain the design's Language support boundary suppression and hook-completion
rules under focused IDs. Keep CHK-005 until its remaining hook execution and
missing-hook rules are separated in the next change. Existing behavior and
public eval assertions already cover these requirements.

## Interface



## Added

### CHK-039
Boundaries MUST return errors for unsuppressed module-attributed violations.

### CHK-040
Boundaries MUST return a visible warning containing the reason for a reasoned boundaries_opt_out.

### CHK-041
Boundaries MUST suppress only exact baseline lines from normally exited hooks.

### CHK-042
Boundaries MUST suppress only the opted-out module's violation lines through a reasoned opt-out after a hook exits normally.

### CHK-043
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

## Provenance

- evals 2026-10-28T05:41:57.09782Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=146 failing=0 passing=146
