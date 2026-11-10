---
fixes:
    - CHK-059
module: checks
phase: audit
run: 20261110T072428-7b8e707c0c34060a699246fc
---

## Motivation

Doctor must detect the supplied missing-ID canary and reject unrelated
failures. In the assets drill, Go ran the valid `TestMissingIDCanary` test
and the eval checker rejected it with C11. Doctor still reported a missed
canary because its matcher expected an underscore after `Test`. This caused
four app and checks eval failures. The owner approved correcting the Go
matching before the assets proof is renewed.

## Interface

Keep the existing public Doctor interface. A Go missing-ID canary can use a
test name without an underscore after `Test`. Detection still requires the
unmapped testcase rejection to identify the supplied canary.

## Added

## Modified

## Removed

## Decisions

## Notes

Prove genuine fail-first through public `checks.Doctor` using a supplied Go
canary named `TestMissingIDCanary`. Assert the passing missing-ID outcome
and recovery. Include an unrelated unmapped testcase that shares the
canary's name as a prefix and require Doctor to reject it as evidence.

Match the complete declared Go test name, preserving its normal Go subtest
association. Keep non-Go canary matching, recipe files, existing tests and
the C11 unmapped-case requirement intact. Limit runtime edits to checks.
Use the pinned controller's `fireweed verify --quick` for worker iteration.
The controller runs full verification before applying the change.

The focused `Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary` run failed
first on the declared Go test and its subtest. The real Go hook ran
`TestMissingIDCanary` successfully and the eval checker returned
`C11: testcase "TestMissingIDCanary" has no known spec ID`, but Doctor
reported a miss. The shared-prefix rejection and missing-ID recovery
assertions passed. The implementation belongs in the existing disposable
`packages/checks/internal/engine/doctor.go`.

The pinned controller accepted `fireweed verify --quick` for the evals
phase. The 125-case checks report had failures only in the new regression
and its two positive cases. Verification warned that the existing 16
CHK-055 cases exceed the configured limit of 15.
## Provenance
- evals 2026-11-10T07:38:54.652412Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=125 failing=3 passing=122
