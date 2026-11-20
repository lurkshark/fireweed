---
eval_fix:
    - CHK-044
    - CHK-045
module: checks
phase: impl
run: 20261120T211958-bec677b122f4e1e75663cd15
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
## Provenance

- evals 2026-11-20T21:29:12.318159Z agent=codex fireweed=0.2.4 playbook=5e05bacdae954379b468564c11429c5532016e26a7afbcd418f9bdb9441e151c tests=142 failing=0 passing=142
- audit 2026-11-20T21:29:29.180606Z agent=codex fireweed=0.2.4 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

## Audit

No findings; checked CHK-044, CHK-045

