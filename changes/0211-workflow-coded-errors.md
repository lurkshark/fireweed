---
module: workflow
---

## Motivation

The design promises a stable code for every problem and output that agents
can read. Workflow requires its lifecycle codes only as message prefixes
(FLOW-054 to FLOW-060, FLOW-125, FLOW-133 and others), and no workflow
error carries a code that `contracts.ErrorCode` can read. The runner passes
these errors on, and the app recovers the code from the text with a regular
expression that takes the last code-shaped token, which can be a wrapped
cause or a spec ID rather than the lifecycle code (APP-078, APP-085). Make
the code part of the error.

## Interface

No signature changes. Errors that carry a code become (or wrap)
`contracts.CodedError`. Diagnostics are unchanged.

## Added

### new:coded-errors
Every error a public workflow function returns for a failure this spec assigns a D006 code MUST carry that code so that contracts.ErrorCode reports it. An error that wraps another module's coded error, for a failure this spec assigns no code, MUST keep the wrapped code.

## Modified

## Removed

## Decisions

## Rejected alternatives

Rewording each prefix item to the ErrorCode form: it would rewrite many
items and their evals for no behavior beyond the added item. The prefix
items stay true, and a later change can relax them once the app reads codes
only through `contracts.ErrorCode`.

## Notes

Keep the existing message text, including FLOW-133's exact text. Where
FLOW-058 and FLOW-062 classify a stale-base failure as P05 while retaining
the underlying cause, P05 is the outer code and the cause stays in the
chain.

Evals: extend the existing tests for each coded error to assert
`contracts.ErrorCode` alongside the prefix, including one P05 stale-base
case whose cause has its own code. Do not build new fixtures where an
existing test already reaches the failure.
