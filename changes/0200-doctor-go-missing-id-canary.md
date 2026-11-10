---
fixes:
  - CHK-059
module: checks
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
