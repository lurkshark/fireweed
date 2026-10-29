---
audit_rounds: 0
bases: {}
module: runner
phase: impl
run: 20261029T020109-5ec3f2b5c9478ce3dff394f1
---



## Motivation

The design's Driving the state machine says run keeps going until the queue is
empty or a limit is reached. NEW next now completes a real Go fixture lifecycle.
The next observable slice is NEW run over a small queue, --max termination,
later continuation and empty-queue success. Add only its minimum runner loop
now, then return promptly to app wiring and an actual NEW run fixture.

## Interface

Add Run(context.Context, RunRequest) (RunResult, error) to thin runner api.go.
Retain RunRequest and RunResult in model/**: RunRequest contains Root and Bin
strings, AllowUnmerged bool, Timeout time.Duration and Max int. RunResult has
Applied []NextResult, Empty bool and Limited bool. Existing interfaces remain.
Implementation belongs under internal/ without importing its public parent;
reuse the existing single-change execution path rather than duplicating it.

## Added

### RUN-016
When single-change executions succeed, Run without a positive Max MUST
continue until the queue is empty, returning each applied NextResult in queue
order.

### RUN-017
With a positive Max, Run MUST apply at most Max changes per invocation. A later
Run invocation MUST continue the remaining queue without reapplying archived
changes.

### RUN-018
On successful return, Run MUST set Empty when no queued changes remain and set
Limited exactly when a positive Max stopped it with queued work remaining.

### RUN-019
Run MUST supply its Root, Bin, AllowUnmerged and Timeout to single-change
execution for every change it processes.

### RUN-020
If single-change execution fails while the calling worktree still holds the
change, Run MUST return the error with its cause and already-applied results
retained, leaving later changes unstarted.

## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

This is the minimum dependency for the bounded successful queue CLI scenario;
app wiring immediately follows. Max <= 0 means unlimited at this API. An empty
queue returns zero Applied results, Empty true and Limited false without agents.
A positive Max that consumes the entire queue returns Empty true/Limited false.
Preserve the existing approval, claim, P07 and phase behavior through Next.
The first loop may stop on errors; do not claim design completion for approval
skipping, saved/block attempts, released-claim independent continuation, retries,
resume, agent-commit undo, time/usage stops, --for, other kinds or drills. Those
remain focused follow-on scenarios after NEW run is wired. No source cleanup
or whole-module completion is requested.

Public-interface evals use ordinary temporary Go repositories, explicit identity,
local hooks and bounded separate fake agents. Exercise a small queue across
independent modules, positive Max termination, later unlimited continuation,
queue order and empty behavior, plus forwarding and a held failure after an
applied change with later work still queued. Keep every earlier test/assertion.
Use public state, commits, archives and agent output rather than private imports,
process observers, assumed commit subjects or helper instrumentation. Validate
a representative fixture through the existing public Next primitive before
freezing the loop tests, so fixture protocol errors cannot hide behind the new
Run stub. Use the established separate-line No findings/Checked IDs format.
No extra permanent preflight framework or matrix is needed. Own Verify stays
under five minutes; required full Done remains authoritative.

Only declared retained API/types, evals and fresh compiling disposable stubs
are allowed in evals, with genuine runtime fail-first for each new ID. Evals,
interfaces and other modules freeze in impl. Workers never commit or call
controller start/done/next/run, edit phase/counters or protected docs. Coordinator
uses native phase advancement and separate isolated Sol/high auditors. No new
phase-write, bound, drill or push exception is authorized.

## Provenance

- evals 2026-10-29T02:10:13.23853Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=204 failing=7 passing=197
- audit 2026-10-29T02:13:54.40663Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings
Checked IDs: RUN-016, RUN-017, RUN-018, RUN-019, RUN-020

The separate Sol/high auditor used only permitted isolated snapshot inputs.
Complete command-log inspection confirms supplied-file reads only, without
implementation/history access, execution or writes.
