---
already_satisfied:
    - RUN-078
module: runner
phase: audit
run: 20261117T111540-1388407b00abb23fae7b12e6
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

### RUN-078
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
## Provenance

- evals 2026-11-17T11:25:42.985011Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=217 failing=0 passing=217

