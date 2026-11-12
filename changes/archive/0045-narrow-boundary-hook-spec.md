---
already_satisfied:
    - CHK-005
    - CHK-044
    - CHK-045
audit_rounds: 0
bases:
    CHK-005: 26a5d821
module: checks
phase: applied
run: 20261028T055100-0b05225b487f7c0fcc71d7da
---



## Motivation

Finish the identified CHK-005 split under SYSTEM's one-behavior spec shape.
CHK-039–043 retain violations, opt-out warnings, suppression and failed-hook
rules. Give missing-hook requirements their own IDs and retain applicable hook
execution under CHK-005, preserving the design's Language support boundary
requirements. Existing public evals already assert all these behaviors.

## Interface



## Added

### CHK-044
Boundaries MUST fail for a missing hook required by a module with disposable globs and no reasoned opt-out.

### CHK-045
Boundaries MUST warn without an error for a missing hook when no boundary enforcement is required.

## Modified

### CHK-005
Boundaries MUST run applicable global and per-module boundary hooks.

## Removed



## Decisions



## Audit

No findings. Checked IDs: CHK-005, CHK-044, CHK-045.

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

The evals phase already retargeted the existing tests and preserved their
assertions. No implementation edit is needed for this split.

## Provenance

- evals 2026-10-28T05:51:43.093855Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=146 failing=0 passing=146
- audit 2026-10-28T05:52:19.553602Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-28T05:53:23.993346Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=146 failing=0 passing=146
