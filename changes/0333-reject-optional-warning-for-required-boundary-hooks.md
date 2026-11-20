---
eval_fix:
    - CHK-044
    - CHK-045
module: checks
---

## Motivation

Callers need the boundary diagnostics to tell them whether a missing hook
requires action. The checks drill reported both X02 for missing required hooks
and X04 saying no enforcement was required. The checks eval accepted that
combination and the dependent FLOW-062 eval caught it (local report `7592e03`).

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

In `Test_CHK_044_CHK_045_MissingHookRequiredVersusOptional`, assert that the
missing-required-hook case returns an error and X02 and contains no diagnostic
with code X04. Keep the positive X04 assertions for the cases with no
disposable globs and with a reasoned opt-out.
