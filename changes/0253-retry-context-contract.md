---
already_satisfied:
    - RUN-022
bases:
    RUN-022: cb605598
module: runner
phase: audit
run: 20261113T183408-96f52bf7680db1757b88265d
---

## Motivation

During regenerate-runner, refreshing workflow context between retries failed
an eval that required the original prompt as a prefix. The design's Retries
rule adds the last verification output to the prompt, but RUN-022 did not
explicitly choose whether worker context is frozen or refreshed. The owner
approved specifying the implemented policy: preserve initial full context
within one Next invocation's phase attempt sequence and append the latest
verification feedback.

## Interface

## Added

## Modified

### RUN-022
When phase verification fails and an attempt remains, Next MUST retain the earlier work and retry that phase. The retry prompt MUST include the last verification output. Within one Next invocation and one phase attempt sequence, worker verification-retry prompts MUST retain that sequence's initial full workflow context as their prefix.

## Removed

## Decisions

## Notes

No signature or type changes. Clarify worker verification retries within
one Next invocation and one phase attempt sequence.

The current implementation captures full worker context before each phase
attempt loop within a Next invocation and appends verification feedback
inside that loop; this narrower RUN-022 clarification is already_satisfied.

Use the existing retry fixtures and original-prefix assertion in
packages/runner/evals/retry_failure_details_test.go to cover the clarified
clause, attributing coverage to RUN-022 as appropriate. Preserve the
RUN-043/044 failed-case contents, ordering and Unicode budget assertions,
earlier work, per-phase attempt bounds and last-verification feedback.

A later Next invocation resuming a held phase, or re-entry to a phase after
audit, may obtain fresh context. This change does not require persisting
prompts across invocations or phase re-entry. Keep audit restricted-context
rules, phase transitions, usage-limit retry semantics and drill retry
behavior unchanged.

Drafted and queued only. No lifecycle execution is authorized by the drafting
request.
## Provenance

- evals 2026-11-13T18:42:47.314825Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=170 failing=0 passing=170

