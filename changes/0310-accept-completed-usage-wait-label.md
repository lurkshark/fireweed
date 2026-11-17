---
module: app
eval_fix:
    - APP-076
    - APP-079
    - APP-083
    - APP-086
---

## Motivation

The fresh runner drill report in commit `b25dc33` shows plain and JSON
usage stops returning the expected completed wait, next wait and budget.
The app evals rejected `completed 0s` and `completed 5ms` because the shared
matcher requires `completed wait` or a form of `waited`. The app specs
require the reason to identify each duration and allow equivalent labels.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

In `cliStopReason` in `packages/app/evals/usage_cli_test.go`, accept
`completed` as a completed-wait label in the usage-stop reason. Keep
accepting `completed wait` and `waited`, including `waited for`. Require a
label and the expected value independently for completed wait, next wait
and budget. Check each matching value against the expectation for its role.

Verify that the observed zero and nonzero forms pass. Use distinct expected
durations to confirm that a missing label, a wrong value or values swapped
between roles still fail. Keep the existing plain and JSON assertions for
exit status, resume guidance, retained work and applied changes. Preserve
the approved-configuration assertions in the APP-164 cases that share the
matcher.
