---
audit_rounds: 0
bases:
    RUN-010: 8c04a67a
    RUN-011: a39d0f24
    RUN-016: 2d6deeca
module: runner
phase: applied
run: 20261029T082745-ee6d49740ab2b5b8d2e2120f
---



## Motivation

The design's Blocking section says the queue does not stop: dependent work waits
while independent work runs. NEW next can already save/block exhausted work and
workflow now supplies projected eligibility, but runner still selects the first
merely unblocked entry and Run stops on every error. Return to that caller now.
Close this slice with actual NEW run: Alpha exhausts verification and is saved/
blocked, Beta depends on Alpha and waits, Gamma is independent and applies, and
the invocation exits 1 with the blocking cause/ref. This does not complete saved
attempt resume, agent-commit undo, time/usage limits or broader recovery.

## Interface

No public interface or model changes. Compose workflow.Eligible, existing Next
results and the existing lifecycle block result under runner's disposable code.
RunResult.Applied continues to contain only applied changes.

## Added

### RUN-025
When single-change execution returns an error after successfully saving and
blocking its change with a clean unclaimed checkout, Run MUST continue processing
eligible changes. Run MUST exclude that blocked change from Applied.

### RUN-026
If a change was blocked during Run, Run MUST return an error when eligible work
is exhausted or Max is reached. That error MUST retain a blocked change's
verification cause and saved attempt ref.

## Modified

### RUN-010
Next MUST return Empty true without invoking agents when workflow.Eligible
returns no changes, including an empty or entirely waiting queue.

### RUN-011
When workflow.Eligible returns changes, Next MUST select its first result and
start it through workflow.Start, propagating start failures without invoking
an agent.

### RUN-016
Run without a positive Max MUST continue single-change execution until no
eligible changes remain, returning applied NextResults in execution order.

## Removed



## Decisions



## Rejected alternatives



## Notes

Preserve current retry/save/block behavior, approval/P07, held-agent-failure
fail-stop, Max's applied-change bound, result forwarding, all assertions and
public signatures. RUN-018 retains its literal queued-work flags: no eligible
work can remain while blocked/waiting queue files still exist; do not report
those files as an empty queue. Do not broaden invocation-failure recovery,
usage/time waits, agent-auth handling, status views or from-attempt resume.

Use representative public runner fixtures and bounded fake agents/hooks. Cover
Next skipping a waiting head to start independent work, no eligible result with
no agent calls, clean saved block then independent application with dependent
work untouched, and blocking error/cause/ref retained at exhaustion and Max.
Keep fixture workers from mistaking committed blocked phase files for a held
change. Preserve existing assertions; adapt fixture setup only as needed.
Fail-first failures must reference only these changed/new IDs. No helper/
observer instrumentation, failure matrix or unsupported platforms. Own Verify
below five minutes and full below fifteen.

Only runner evals may change in evals; the existing public API already compiles
and needs no new stub. Real runner selection/continuation edits follow the
independent assertion audit. Workers never commit, run governing workflow
commands, change refs/counters/runtime, edit SPEC/protected files/other modules,
or read bootstrap source/tests, deleted code, saved implementations or rebuild-v1.
Use the fixed Sol/high workers and isolated assertion auditor; no drill or new
phase-write exception. Actual NEW Go-fixture proof follows application.

## Audit

No findings
Checked IDs: RUN-010, RUN-011, RUN-016, RUN-025, RUN-026

## Provenance

- evals 2026-10-29T08:40:13.802173Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=217 failing=5 passing=212
- audit 2026-10-29T08:42:28.824748Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-29T08:50:16.074592Z agent=manual fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=217 failing=0 passing=217
