---
already_satisfied:
    - FLOW-142
module: workflow
phase: audit
run: 20261114T022109-4d6a3a65325aa78d037f4c8b
---

## Motivation

The design's Starting a change and Git behavior / Base branch require a
warning when HEAD does not contain the locally resolved base tip (R06/R14).
Main 3b3b64d starts an approved change without any warning when origin/main
has a later commit absent from HEAD. This is distinct from a stale spec hash.

## Interface

No public signature or type changes.

## Added

### FLOW-141
On a successful Start or StartFromAttempt, when the base ref selected by gitstore.BaseRef is not an ancestor of the starting HEAD, the returned State.Diagnostics MUST contain a P05 warning naming that ref and explaining that the queue may be stale. This warning MUST NOT by itself reject or advance the change.

### FLOW-142
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

The owner approved marking FLOW-142 as already satisfied in d760e3f.
Controller verification of the resumed evals passed. All four FLOW-141
scenarios failed first because the P05 warning was missing. All four
FLOW-142 scenarios passed, as did the FLOW-058 stale spec-hash rejection.
The implementation belongs in the shared acquisition path in
`packages/workflow/internal/lifecycle/claim.go`.
## Provenance

- evals 2026-11-14T03:09:49.169683Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=211 failing=5 passing=206

