---
module: runner
already_satisfied:
    - new:drill-eval-error-code
---

## Motivation

Callers need the stable E01 diagnostic assigned by D006 when a required drill
testcase fails. In the fresh runner drill report in commit `b25dc33`, both
APP-141 fixtures received P01. The reconstructed runner returned the
checks error through a joined error that `contracts.ErrorCode` could not read.

## Interface

The existing `runner.Drill` return error exposes the code and spec ID through
`contracts.ErrorCode`.

## Added

### new:drill-eval-error-code
When the primary error from Drill is a failed enforced testcase in the target or a dependent module, `contracts.ErrorCode` on the returned error MUST report E01 and the failed spec ID supplied by `checks.RequirePass`.

## Modified

## Removed

## Decisions

## Notes

Use the public `runner.Drill` fixture with a valid eval report and matching
hook exit status. Exercise target-only and dependent-only enforced failures.
Give the failing target and dependent distinct fixture spec IDs. In both
cases assert Result is `fail` and `contracts.ErrorCode(err)` returns E01 and
the failing spec ID. Keep the existing report and retry assertions for both
paths.

The current runner preserves the primary code and subject through eval
error wrapping and aggregation, so the added spec is already satisfied.
