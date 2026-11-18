---
already_satisfied:
    - GIT-011
module: gitstore
phase: applied
run: 20261118T025100-2e450daf8cf26faf3e45d81e
---

## Motivation

The November 17 gitstore drill passed its own evals while replacing the typed
process-exit cause with formatted text. Checks uses the cause to decide whether
to read a hook's JUnit report, and workflow uses its exit status to distinguish
a stale base-tip warning from a failed ancestry check. Runner also recognizes
context cancellation to stop a drill promptly. These distinctions
belong in the public Run contract and its evals.

## Interface

Specify how callers distinguish the existing `gitstore.Run` failure outcomes
using the returned error and `Result.ExitCode`.

## Added

### GIT-011
When a started process exits normally with a nonzero status, Run MUST return
an error matching `*exec.ExitError` via `errors.As`, with the exit status
matching `Result.ExitCode`. An error from a failure to launch MUST NOT match
`*exec.ExitError`. When caller cancellation stops a running process, Run MUST
return an error matching `context.Canceled` via `errors.Is`.

## Modified

## Removed

## Decisions

## Notes

Use the existing exit-7 fixture to assert the typed cause and check its
status against the result. Use a missing executable at a fixture-local
absolute path for launch failure. Extend the synchronized cancellation case
so cancellation occurs while the process is still running. Retain the
existing deadline, captured-output, and process-group checks. The current
implementation returns the process wait error, launch error, or context
error for these respective outcomes.
## Provenance

- evals 2026-11-18T02:52:18.941639Z agent=codex fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=15 failing=0 passing=15
- audit 2026-11-18T02:53:03.089531Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-18T03:13:14.788232Z agent=codex fireweed=0.2.1 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=1314 failing=0 passing=1314

## Audit

No findings; checked GIT-011

