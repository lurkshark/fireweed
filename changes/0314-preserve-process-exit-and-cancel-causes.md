---
already_satisfied:
    - new:run-process-error-causes
module: gitstore
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

### new:run-process-error-causes
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
