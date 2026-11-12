---
bases:
    RUN-014: 7115e12f
module: runner
phase: impl
run: 20261112T014811-37a2c5be230bf289938126db
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

Drop the separate full Verify and call Done. Use the workflow prerequisite's
`workflow.IsVerificationFailure` predicate to count verification failures as
failed attempts and build retry feedback from Done's verification result.
Propagate other Done errors with their existing diagnostics. A quick Verify
(target module only) before Done is allowed.

Evals: count eval-hook invocations in a fixture project and require one
full verification per successful attempt, and a retry with the last
verification output when Done reports a dependent failure.

This replays the existing RUN-014 change after the workflow prerequisite.
Reuse the previously audited eval from commit `4ff5a9b` and verify it against
the current phase base before applying the implementation.

The controller is pinned, so the saving applies to queue runs only after
the controller is rebuilt from a main that includes this change.
## Provenance
- evals 2026-11-12T01:59:51.729717Z agent=manual fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=170 failing=3 passing=167
- audit 2026-11-12T02:00:12.49347Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked RUN-014
