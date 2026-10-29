---
already_satisfied:
    - RUN-015
audit_pending:
    - RUN-033
audit_rounds: 2
bases:
    RUN-015: 63014c25
module: runner
phase: evals
run: 20261029T104429-263c67507b7c55993d7c1bdb
---



## Motivation

The design's Usage limits section says a usage-limit exit waits and repeats the
same attempt, rather than failing it, and stops cleanly when total wait would
pass limit_budget. InvokeAgent already classifies UsageLimit, but Next treats it
as ordinary failure. Add bounded wait/control and owned-step resume, then return
to app's exit-3 mapping and actual NEW Go proof: a limit followed by success,
then budget stop with work retained and a later invocation completing the same
run. This library dependency reaches the standing fifth-change checkpoint;
review/PASS/push interrupts the slice, then the app caller follows immediately.

## Interface

Add IsUsageStop(err error) bool to the thin runner facade. It recognizes a clean
usage-budget stop for the app's later exit-3 mapping. Existing signatures and
model fields stay unchanged. Error representation, control policy, shared wait
accounting and classifier behavior belong in disposable runner code. A clean
stop is returned as a non-nil error with State/Steps retained, so old callers do
not falsely report an unapplied change as applied while integration is pending.

## Added

### RUN-029
Next MUST read limit_wait and limit_budget as durations, defaulting to 30m and
4h when unset. It MUST reject malformed durations, a nonpositive wait or negative
budget with C01 before claiming a fresh change or invoking an agent.

### RUN-030
For an AgentResult with UsageLimit true, Next MUST wait limit_wait and repeat
the same phase attempt without spending a failed attempt. It MUST retain earlier
step work and captured invocations, keeping existing verification feedback.

### RUN-031
When another usage wait would exceed limit_budget, Next MUST return a clean
usage-stop error identifying completed wait, next wait and budget. It MUST retain
the held State, work and ordered Steps without advancing, blocking, committing
or releasing the claim.

### RUN-032
Run MUST account for total usage wait across phases and changes in one
invocation. On a clean usage stop it MUST retain already-applied results and
propagate the stop classification.

### RUN-033
When the calling worktree already holds a queued step, Next MUST resume that
owned State using its approved snapshot without a fresh Start, retaining its
run, phase, assigned IDs and work. Existing approval and unapproved-work bounds
MUST remain enforced for starting another change.

### RUN-034
IsUsageStop MUST recognize clean usage-stop errors through wrapping and return
false for nil or ordinary failures.

### RUN-035
If Run blocked a change before a usage stop, it MUST return failure with the
blocking cause/ref and stop reason retained. IsUsageStop MUST return false for
that combined failure so the caller cannot report it as a clean stop.

## Modified

### RUN-015
For a non-usage-limit agent invocation failure, Next MUST return an error and
captured invocation results without advancing the phase or releasing its held
claim.

## Removed



## Decisions



## Rejected alternatives



## Notes

RUN-015 only clarifies unchanged non-limit failure behavior and reuses existing
assertions as already_satisfied. Usage handling uses the existing UsageLimit
classification; do not change InvokeAgent's final-20-line matching. Keep bounded
verification attempts, P05 commit recovery/warnings, approval/P07, audit rounds,
Max, saved blocking and ordinary held-error fail-stop intact. Usage waits do not
manufacture verification failures or change their feedback. Wait accounting is
per command invocation, shared by Run's Next executions, not reset each phase.

Resume the calling worktree's existing claim after a clean stop, rather than
rejecting its recoverable dirty work. Do not steal a foreign claim or accept
unapproved runtime/author changes through new snapshot/configuration lookup.
Use current approved snapshot while held. Waits must remain cancellable; a
context cancellation returns its actual cause and retains recoverable state.
Do not implement --for, CLI exit 3, non-limit agent retries, new-module/status,
new model-control fields or general process/Git maintenance in this dependency.

Use representative public runner fixtures, short configured waits and bounded
fake agents. Cover successful same-attempt retry with attempts=1 and retained
partial work/feedback; budget stop/default policy values without long waits;
shared Run budget with earlier results retained; same owned run/phase/work on
later invocation; wrapped/nil/ordinary classifier cases; earlier block failure
precedence; and invalid policy before mutation/agent calls. Default durations
can be observed in an immediate budget-stop explanation, not by waiting hours.
Preserve every existing applicable assertion; no timing failure matrix, fake
clock, observer instrumentation or speculative unsupported cases. Own Verify
below five minutes, full below fifteen; keep cases representative/below15/spec.

Only the facade addition/new compiling classifier stub and runner evals in evals;
real control changes follow isolated assertion audit. Workers never commit,
edit SPEC/protected docs/other modules, change governing refs/counters/runtime,
or run governing workflow commands. Fixture-local public workflows/Git allowed
with explicit identities. No bootstrap source/tests, deleted code, saved rebuild
implementation, old worker logs or rebuild-v1 as oracle. No drill or phase-write
exception. Stop at the fifth-change quality checkpoint after application; do not
start the app caller before its PASS and authorized push.

## Audit

- RUN-033: "run, phase, assigned IDs and work." — no assertion verifies that the saved partial step work remains after the owned-step resume completes.

## Provenance

- evals 2026-10-29T10:58:09.536757Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=246 failing=15 passing=231
- audit 2026-10-29T11:01:11.642301Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- evals 2026-10-29T11:08:09.568902Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=246 failing=15 passing=231
- audit 2026-10-29T11:11:25.646866Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
