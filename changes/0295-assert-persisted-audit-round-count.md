---
module: runner
eval_fix: [RUN-021]
---

## Motivation

RUN-021 governs the runner's attempt allowances. FLOW-038 requires the audit
round count to persist, and FLOW-039 requires blocking at the limit. The
retained workflow build stopped with P06, released the claim, retained the
finding and wrote YAML audit_rounds: 1. The runner eval failed because its
direct comparison of an untyped frontmatter value to Go int rejected the same
returned count represented as float64(1).

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Keep the expected evals/audit invocation order, P06, phase, released claim
and exact finding. Assert the returned round count as a number with value one
independent of concrete Go representation and also parse the pending document
to confirm the persisted integer count. Keep RUN-021 coverage. Leave the
public workflow claim format and the audit and eval attempt limits intact.
