---
bases:
    RUN-029: 5b18306a
module: runner
---

## Motivation

The design promises a stable code for every problem and output that agents
can read. RUN-029 rejects bad limit settings "with C01", but no runner error
carries a code that `contracts.ErrorCode` can read, so the app parses the
text instead (APP-078, APP-085). State RUN-029's code through
`contracts.ErrorCode` and apply SYSTEM.md's diagnostic-code convention to
runner's other coded errors.

## Interface

No signature changes. Errors that carry a code become (or wrap)
`contracts.CodedError`.

## Added

## Modified

### RUN-029
Next MUST read limit_wait and limit_budget as durations, defaulting to 30m and
4h when unset. It MUST reject malformed durations, a nonpositive wait or negative
budget with an error whose contracts ErrorCode is C01 before claiming a fresh
change or invoking an agent.

## Removed

## Decisions

## Notes

Apply the convention to the drill's configuration errors as well, and wrap
with `%w` so that a workflow or checks code stays readable through `Next`,
`Run` and `Drill`. Keep the existing message text until 0223.

Evals: assert `contracts.ErrorCode` for RUN-029's rejection and for one
`Next` failure that wraps a lifecycle code.
