---
bases:
    RUN-021: 113a4a26
module: runner
---

## Motivation

The design's Audit step says a return to evals counts against the evals
step's attempts (R08/R12). In main 3b3b64d, Next starts attempt at 1 on every
phase re-entry, so every audit return receives another full evals allowance.
Bound unattended retries across audit returns within one Next invocation while
retaining the separate persisted audit.max_rounds enforcement.

## Interface

No public signature or type changes.

## Added

### new:audit-return-exhaustion
When audit returns to evals after that allowance is exhausted, Next MUST save and block the held evals step through workflow.Block before invoking another worker. Its failure MUST identify the exhausted evals allowance and retained attempt ref, leaving the existing run, assigned IDs and audit findings recoverable.

## Modified

### RUN-021
Next MUST limit each phase's verification-attempt sequence to the configured attempts count, defaulting to two when unset. Evals sequences separated by audit findings within one Next invocation MUST share one allowance, excluding usage-limit retries. Audit and impl sequences MUST retain separate allowances.

## Removed

## Decisions

## Notes

Use bounded fake workers and a restricted fake auditor returning valid exact
clause findings. Cover an audit correction within allowance, exhaustion on a
later return, an earlier verification retry that consumes the same allowance,
and a usage-limit retry that does not. Reuse the existing block/ref/independent
continuation assertions. Keep audit.max_rounds as a separate upper bound that
also applies to manual Done; whichever limit is reached first stops automation.
Do not add cross-invocation attempt persistence or change resumability. RUN-022's
initial-context prefix remains scoped to one phase attempt sequence: re-entry
after audit still gets fresh context including findings. No timer changes,
new controller state or altered audit restrictions are authorized.

Drafted at the owner's request after the R01–R29 traceability review.
No lifecycle execution, agent invocation, drill or release is authorized by
this drafting request. Workers must obtain context before their step and use
fireweed verify; only the controller may advance or commit lifecycle work.
