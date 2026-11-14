---
bases:
    RUN-021: 113a4a26
module: runner
phase: audit
run: 20261114T073249-d99cc6c13cfdd9da3ad488c6
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

### RUN-070
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

Evals verification passed with fail-first evidence for RUN-021 and RUN-070.
The current runner invokes another evals worker after the allowance is spent,
then reaches the audit-round limit without the evals attempt ref. This occurs
after two audit returns with the default allowance and on the first return
when an earlier verification retry spent the second attempt. The correction,
usage-limit retry and configured audit-round limit cases passed.
The attempt loop and blocking path are in the disposable
`packages/runner/internal/next/next.go`.

Drafted at the owner's request after the R01–R29 traceability review.
No lifecycle execution, agent invocation, drill or release is authorized by
this drafting request. Workers must obtain context before their step and use
fireweed verify; only the controller may advance or commit lifecycle work.
## Provenance

- evals 2026-11-14T07:49:14.005387Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=178 failing=4 passing=174

