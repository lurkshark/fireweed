---
bases:
    RUN-014: 7115e12f
module: runner
---

## Motivation

RUN-014 makes Next verify each successful agent attempt and then call
`workflow.Done`, and FLOW-003 requires Done to run full verification before it
commits. So every successful phase runs the full verification twice
(`run.go` calls `workflow.Verify` and then `workflow.Done`). SYSTEM.md
allows up to 35 minutes for full implementation verification, so an impl
phase can spend that long again for no new information. The runner drill
lists this as a gap.

## Interface

None. Next's verification sequence changes.

## Added

## Modified

### RUN-014
After each successful agent invocation, Next MUST run full verification of that attempt at most once and call workflow.Done, continuing until applied. A verification failure reported by Done MUST count as a failed attempt under RUN-021 and RUN-022. Next MUST return the applied state and the ordered phase invocation results.

## Removed

## Decisions

## Notes

The simplest shape is to drop the separate full Verify and call Done,
retrying on its verification failure; a quick Verify (target module only)
before Done is allowed. If Done's result cannot distinguish a verification
failure from other Done errors, block with a Question rather than parsing
error text.

Evals: count eval-hook invocations in a fixture project and require one
full verification per successful attempt, and a retry with the last
verification output when Done reports a dependent failure.

The controller is pinned, so the saving applies to queue runs only after
the controller is rebuilt from a main that includes this change.
