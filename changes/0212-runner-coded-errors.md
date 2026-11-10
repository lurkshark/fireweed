---
module: runner
---

## Motivation

The design promises a stable code for every problem and output that agents
can read. Runner assigns its own codes in a few places (RUN-029 and the
drill's configuration errors use C01; RUN-041 gains P01 in 0205) and
otherwise returns workflow and checks errors from `Next`, `Run` and
`Drill`. None of its own errors carry a code that `contracts.ErrorCode`
can read, so the app parses the text instead (APP-078, APP-085). Make the
code part of the error and keep wrapped codes readable through the runner.

## Interface

No signature changes. Errors that carry a code become (or wrap)
`contracts.CodedError`.

## Added

### new:coded-errors
Every error a public runner function returns for a failure this spec assigns a D006 code MUST carry that code so that contracts.ErrorCode reports it. An error that wraps another module's coded error, for a failure this spec assigns no code, MUST keep the wrapped code.

## Modified

## Removed

## Decisions

## Notes

Keep the existing message text. Wrap with `%w` so that a workflow or checks
code applied by 0210 and 0211 stays readable through `Next` and `Run`.

Evals: assert `contracts.ErrorCode` for RUN-029's C01 rejection and for one
`Next` failure that wraps a lifecycle code. Do not build new fixtures where
an existing test already reaches the failure.
