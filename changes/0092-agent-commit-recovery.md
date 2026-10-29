---
audit_rounds: 0
bases: {}
module: runner
phase: audit
run: 20261029T100929-08538f37f6e35ed6a14eea8a
---



## Motivation

The design's Driving the state machine section says that if an agent creates
commits, fireweed undoes them with a soft reset and logs a warning. NEW next
currently leaves that HEAD movement to fail lifecycle verification. Recover
those commits before checking the step, retaining the work. Close through an
actual NEW Go next fixture whose bounded fake worker/auditor commit during
phases, showing normal checked phase commits/archive and visible warnings.
Do not expand this slice to usage/time waits, branch maintenance or status.

## Interface

No facade or model changes. Use existing State.Diagnostics in NextResult for
warnings; app next and run already report those diagnostics. Keep captured
AgentResult stdout/stderr/exit and ordered Steps intact.

## Added

### RUN-027
When a phase agent creates commits, Next MUST soft-reset HEAD to the held phase
base after invocation and before verification or returning an invocation failure,
retaining the step's working-tree and index work.

### RUN-028
For each such recovery, Next MUST return a P05 warning in State.Diagnostics
identifying the phase and that work was retained. It MUST preserve these warnings
across phase advancement and on an invocation-failure return.

## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

P05 is the existing D006 HEAD/phase-base mismatch cause; severity belongs to the
diagnostic. The message must make the automatic soft-reset recovery clear.
Do not allocate an unused code or rewrite unrelated explanations. Existing
assets explain(P05) covers a base mismatching intended state and restoring that
state through the workflow. Approval/P07, claims, assigned IDs/run/phase,
verification retries/blocking and held invocation-failure behavior stay intact.
Never falsify captured agent output to insert warnings.

Use representative public runner fixtures with explicit Git identity and
bounded fake agents/hooks. One complete lifecycle can exercise worker and audit
commits, including multiple commits in one invocation, proving reset to the held
base rather than just the last parent. Assert retained staged/tracked/untracked
step work, successful normal verification/archive, no agent commits among the
branch's phase commits, and phase-tagged P05 warnings after advancement. Include
an agent that commits then exits nonzero, proving reset/work/warning and captured
failure survive while the phase and claim remain held. Reuse existing assertions
for no-commit paths; no failure matrix, observers, unsupported cases or new
history/ref maintenance API. Own Verify below five minutes, full below fifteen.

Only runner evals change in evals; the public API already compiles, with no new
stub or interface. Real undo/warning code follows isolated assertion audit.
Workers never commit or invoke governing lifecycle commands in this checkout,
change its refs/counters/runtime, or edit SPEC/protected docs/other modules.
Deliberate fixture-agent commits are test inputs required by this design clause;
fixture-local public workflows/Git are allowed. No bootstrap source/tests,
deleted code, saved rebuild implementation or rebuild-v1 as oracle. No drill,
phase-write exception or general workflow-file reorganization.

After application rebuild NEW and close the declared actual CLI proof before
starting the next recovery mechanism. This is the fourth application since the
passing checkpoint if successful; the standing fifth-change review still applies.

## Provenance

- evals 2026-10-29T10:19:32.422623Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=231 failing=2 passing=229
