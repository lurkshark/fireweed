---
eval_fix:
    - APP-076
    - APP-079
    - APP-083
    - APP-086
module: app
phase: applied
run: 20261117T120108-8182c16b11014258a023822d
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

A temporary matcher probe passed all four distinct plain and JSON reasons from
the runner drill, covering `completed 0s` and `completed 5ms`. With distinct
expected durations of 5ms, 11ms and 17ms, the matcher accepted each supported
completed-wait label and rejected each missing role label, an incorrect value,
swapped values and conflicting values for one role.

## Audit

No findings; checked APP-076, APP-079, APP-083, APP-086

## Provenance

- evals 2026-11-17T12:09:19.75668Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=403 failing=0 passing=403
- audit 2026-11-17T12:09:37.580535Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-17T12:17:45.350906Z agent=codex fireweed=0.2.1 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=403 failing=0 passing=403

