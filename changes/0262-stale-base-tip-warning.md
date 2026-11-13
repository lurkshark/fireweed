---
module: workflow
---

## Motivation

The design's Starting a change and Git behavior / Base branch require a
warning when HEAD does not contain the locally resolved base tip (R06/R14).
Main 3b3b64d starts an approved change without any warning when origin/main
has a later commit absent from HEAD. This is distinct from a stale spec hash.

## Interface

No public signature or type changes.

## Added

### new:base-tip-warning
On a successful Start or StartFromAttempt, when the base ref selected by gitstore.BaseRef is not an ancestor of the starting HEAD, the returned State.Diagnostics MUST contain a P05 warning naming that ref and explaining that the queue may be stale. This warning MUST NOT by itself reject or advance the change.

### new:contained-base-tip
Start and StartFromAttempt MUST omit the base-tip P05 warning when the selected base tip is equal to or an ancestor of the starting HEAD.

## Modified



## Removed

## Decisions

## Notes

Use local refs only. Exercise behind/diverged HEAD, equal/ahead HEAD, a
locally known remote ref and the local-base fallback. Preserve P04 approval,
P07 limits, claims, assigned IDs, resume state and FLOW-058 stale-spec-hash
rejection. A base-tip warning must not turn that existing error into a warning.
Use public lifecycle fixtures, including one resumed phase. APP-027 already
carries Start diagnostics, but Next currently loses acquisition warnings when
it replaces the state after Done. The separate runner follow-up 0272 retains
those warnings during automatic execution; complete it before claiming the
automatic path is covered. Keep this change within workflow; do not edit
runner or app evals or interfaces here. Coordinate the P05 explanation with
0268 before treating the warning as acceptance-complete.

Drafted at the owner's request after the R01–R29 traceability review.
No lifecycle execution, agent invocation, drill or release is authorized by
this drafting request. Workers must obtain context before their step and use
fireweed verify; only the controller may advance or commit lifecycle work.
