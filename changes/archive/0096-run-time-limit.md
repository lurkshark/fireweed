---
already_satisfied:
    - RUN-016
    - RUN-030
audit_rounds: 0
bases:
    RUN-016: e626633e
    RUN-030: 1dde6570
module: runner
phase: applied
run: 20261029T155505-139e6ccaaac6b169ce975070
---



## Motivation

The design exposes run --for6h and requires a clean exit3 when the total time
limit expires or another usage wait would pass it. Usage-budget stops now work
through NEW, but Run has no elapsed-time bound. Add the minimal runner deadline
control/classification, then return immediately to app's --for parser and clean
output. Close the slice with actual NEW stopping a bounded Go fixture agent
with work/run/IDs retained, followed by later owned completion. This is total
invocation time, distinct from the existing per-agent Timeout and wait budget.

## Interface

Add For time.Duration to model.RunRequest (the existing public alias exposes it)
and IsTimeStop(err error) bool to the thin runner facade. For zero means no
elapsed-time bound; a positive value is the command's total duration. CLI parsing
will reject invalid/nonpositive explicit --for values in the following app
change; no new negative-request diagnostic path is required in this dependency.
No NextRequest field, other signature/model field or public clock/controller API.
All timing/error representation/control belongs in disposable runner code.
A clean time stop is a nonnil error, like the existing usage stop, so an old
caller cannot report interrupted work as successful before app integration.

## Added

### RUN-036
Run MUST treat For zero as unbounded and positive For as one elapsed-time
budget for the whole invocation, including agent and verification time across
phases and changes.

### RUN-037
When its For bound stops execution, Run MUST return a nonnil clean time-stop
error naming the limit, with already-applied results retained. A step interrupted
during agent invocation or verification MUST keep its work, claim, run, phase
and assigned IDs without advancing, blocking, committing or releasing that step.

### RUN-038
When a usage wait fits its usage budget but would pass Run's For deadline,
execution MUST return a clean time stop before waiting. It MUST preserve the
held step's work and identity.

### RUN-039
IsTimeStop MUST recognize clean time-stop errors through wrapping. It MUST
return false for nil, usage stops, ordinary failures and parent-context
cancellation.

### RUN-040
If Run blocked a change before a time stop, it MUST return ordinary failure
retaining the blocking cause/ref and time-stop reason. IsTimeStop MUST return
false for that combined failure.

## Modified

### RUN-016
Run without a positive Max MUST continue single-change execution until no
eligible changes remain or its For bound stops it, returning applied NextResults
in execution order.

### RUN-030
Subject to the usage budget and Run's For bound, an AgentResult with UsageLimit
true MUST make Next wait limit_wait and repeat the same phase attempt without
spending a failed attempt. It MUST retain earlier step work and captured
invocations, keeping existing verification feedback.

## Removed



## Decisions



## Notes

RUN-016/030 only qualify preserved ordinary execution/usage behavior and reuse
existing assertions as already_satisfied; new time requirements need genuine
fail-first. For zero leaves every current path unchanged. Keep Max, approval/
P07, owned approved-snapshot resume, usage wait accounting, agent-commit recovery,
verification retries, audit bounds and earlier-block failure semantics intact.
Distinguish this invocation's own For deadline from a parent context cancelling;
the latter remains its actual ordinary cause, not a clean limit result. Usage
budget overflow retains its existing usage-stop behavior; the new pre-wait time
case applies when that wait otherwise fits the usage budget. Cooperative agent/
hook cancellation uses existing process-group control; no native code, process
instrumentation or extra public control framework. Do not widen atomic lifecycle
commit behavior or add special probes at commit boundaries.

Evals phase may add only the declared model/facade fields, a new compiling false
classifier stub and runner evals. Preserve existing tests/assertions/helpers.
Use representative bounded public Run fixtures and real timers, with generous
short bounds. Cover a single total budget across phases/changes and earlier
Applied retention; interrupted agent and verification work; usage preflight
that stops before a full wait; wrapped/nil/usage/ordinary classifier cases and
parent cancellation; and prior-block precedence. Reuse/consolidate fixtures,
keep cases below15/spec and own Verify below5m/full below15m. Ordinary local
fixture hooks may emit deterministic JUnit to keep elapsed-time cases stable;
the later actual NEW proof will use genuine Go hooks. No fake clocks, elaborate
observer calibration, timing matrix or tests merely to increase confidence.

Workers never commit or run governing workflow commands; no protected docs,
SPEC/other-module/undeclared interface changes, author/runtime/phase/Audit/
Provenance/ref/claim/counter/bound edits, delegation or drill. No bootstrap
source/tests, rebuild-v1, deleted/abandoned/saved implementation or old worker
logs as oracles. No CLI --for/parser/status/new-module or unused diagnostics in
this dependency. After archive, wire app immediately and close actual NEW proof.
Two applied changes since review precede this one; normal checkpoints continue.

## Provenance

- evals 2026-10-29T16:27:41.954923Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=262 failing=5 passing=257
- audit 2026-10-29T16:36:40.288023Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-29T16:49:15.609604Z agent=manual fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=262 failing=0 passing=262

## Audit

No findings
Checked IDs: RUN-016, RUN-030, RUN-036, RUN-037, RUN-038, RUN-039, RUN-040
