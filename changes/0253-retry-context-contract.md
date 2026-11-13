---
already_satisfied:
    - RUN-022
bases:
    RUN-022: cb605598
module: runner
---

## Motivation

During regenerate-runner, refreshing workflow context between retries failed
an eval that required the original prompt as a prefix. The design's Retries
rule adds the last verification output to the prompt, but RUN-022 did not
explicitly choose whether worker context is frozen or refreshed. The owner
approved specifying the implemented policy: preserve the initial full context
within a phase and append the latest verification feedback.

## Interface

No signature or type changes. Clarify the context delivered to worker retries
within one phase.

## Added

## Modified

### RUN-022
When phase verification fails and an attempt remains, Next MUST retain the earlier work and retry that phase. The retry prompt MUST include the last verification output. Within a phase, worker retry prompts MUST retain the full workflow context used for that phase's first worker invocation as their prefix.

## Removed

## Decisions

## Notes

The current implementation captures full worker context once per phase,
before the attempt loop, and appends verification feedback; RUN-022 is
already_satisfied.

Use the existing retry fixtures and original-prefix assertion in
packages/runner/evals/retry_failure_details_test.go to cover the clarified
clause, attributing coverage to RUN-022 as appropriate. Preserve the
RUN-043/044 failed-case contents, ordering and Unicode budget assertions,
earlier work, per-phase attempt bounds and last-verification feedback.

This clarification concerns worker verification retries within a phase.
Keep audit restricted-context rules, phase transitions, usage-limit retry
semantics and drill retry behavior unchanged.

Drafted and queued only. No lifecycle execution is authorized by the drafting
request.
