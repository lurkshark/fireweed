---
bases:
    RUN-014: 7115e12f
blocked: impl verification attempts exhausted
module: runner
phase: impl
run: 20261111T232906-bd2a07c548b6842a8b999677
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

Evals fail-first: `Test_RUN_014_NextRunsFullVerificationOncePerAttempt`
observed dependent eval attempts `[1 1]` on a clean pass (expected `[1]`)
and `[1 2 2]` after a dependent failure (expected `[1 2]`). Both scenarios
reached `applied`, and the retry received the dependent hook output.
The runner change belongs in `packages/runner/internal/next/run.go`.

## Questions

May workflow expose a structured verification-failure marker from `Done` in a
separate workflow change? `Done` can return the same zone or validation codes
from verification and from checks around the phase transition. Its result
carries no failure stage, so runner cannot reliably tell which failures should
consume an attempt.
## Provenance
- evals 2026-11-11T23:50:24.912475Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=170 failing=3 passing=167
- audit 2026-11-11T23:50:44.907686Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked RUN-014
