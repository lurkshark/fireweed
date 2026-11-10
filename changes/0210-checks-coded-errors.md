---
module: checks
---

## Motivation

The design promises a stable code for every problem and output that agents
can read. Checks requires most of its error codes only as message prefixes
(CHK-002, CHK-026 to CHK-028, CHK-036 to CHK-038, CHK-046 to CHK-048), so
callers must parse text to learn the code. Workflow and the app pass these
errors on, and the app recovers codes with a regular expression that also
matches spec IDs (APP-078, APP-085). Make the code part of the error.

## Interface

No signature changes. Errors that carry a code become (or wrap)
`contracts.CodedError`. Diagnostics are unchanged.

## Added

### new:coded-errors
Every error a public checks function returns for a failure this spec assigns a D006 code MUST carry that code so that contracts.ErrorCode reports it. An error that wraps another module's coded error, for a failure this spec assigns no code, MUST keep the wrapped code.

## Modified

## Removed

## Decisions

## Rejected alternatives

Rewording each prefix item to the ErrorCode form: it would rewrite a dozen
items and their evals for no behavior beyond the added item. The prefix
items stay true, and a later change can relax them once the app reads codes
only through `contracts.ErrorCode`.

## Notes

Keep the existing message text. CHK-043's rule that the boundary error
carries no leading X02 still holds: the X02 code travels in the Diagnostic,
and the error may also carry it through `contracts.CodedError`.

Evals: extend the existing tests for each coded error to assert
`contracts.ErrorCode` alongside the prefix. Do not build new fixtures where
an existing test already reaches the failure.
