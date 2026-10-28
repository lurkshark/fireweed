---
audit_rounds: 0
bases: {}
module: runner
phase: impl
run: 20261028T232748-8a17fd5f3261c7dfe4264076
---



## Motivation

The design's Driving the state machine requires next to compose start, agents,
verify and done. The current runner exposes only agent invocation, so NEW cannot
complete even one queued change. This dependency adds the minimum successful
spec-change loop for a bounded fake-agent Go fixture; the next change returns
to app wiring and demonstrates the actual NEW next command. Retries, saved
attempts, projected dependency blocking and run limits follow that runnable path
as focused scenarios, rather than prerequisites for exposing it.

## Interface

Add Next(context.Context, NextRequest) (NextResult, error) to thin runner api.go.
Retain types in model/**: NextRequest has Root and Bin strings, AllowUnmerged
bool and Timeout time.Duration; NextResult has State workflow.State, Empty bool
and Steps []StepResult; StepResult has Phase and Role strings and Agent
AgentResult. Existing signatures remain. Keep behavior under internal/ and use
public dependency interfaces, without own-parent import cycles.

## Added

### RUN-010
Next MUST return Empty true without invoking agents when the queue is empty.

### RUN-011
For a nonempty queue, Next MUST select its first unblocked change and start it
through workflow.Start, propagating start failures without invoking an agent.

### RUN-012
Unless AllowUnmerged is true, Next MUST reject a selected change absent from
both the local and shared base before claiming it or assigning IDs.

### RUN-013
For each phase, Next MUST invoke the worker role for evals and impl, and the
audit role for audit, using the held snapshot and request Bin and Timeout.
Worker prompts MUST include full workflow context. Audit prompts MUST include
the change, projected spec and eval contents without implementation contents.

### RUN-014
After each successful agent invocation, Next MUST verify the current phase and
call workflow.Done only when verification passes, continuing until applied.
It MUST return the applied state and the ordered phase invocation results.

### RUN-015
When invocation or phase verification fails, Next MUST return an error and
captured invocation results without advancing that phase.

## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

This is the minimum runner dependency for the first NEW next lifecycle; do not
implement broad retry/recovery, run, CLI or whole-module completion here. Empty
queue and first unblocked selection use project.Queue; workflow.Start retains
its current eligibility/clean-tree/claim/approval-bound checks, including P07.
Use locally known base refs with the design's remote/default-base fallback.
Rejection without AllowUnmerged must precede workflow.Start side effects.
The later projected dependency-blocking scenario will extend selection; do not
claim it is complete here. An all-blocked queue returns an error without agents.

Compose public workflow.Context, Verify and Done with the existing invocation
primitive. A quick Verify may precede Done's required full checks. Audit context
is deliberately assembled from permitted inputs, not full context containing
implementation. Include the audit playbook and the exact held change path so
the auditor can submit Audit. Caller failures retain the existing held work for
manual recovery; retries/saved attempts and clean limit stops remain pending.
Do not add diagnostic allocations beyond preserving existing dependency causes
and a clear approval error. Coordinator advancement passes the configured agent
command as provenance, not manual when Next actually invokes it.

Public-interface evals use ordinary temporary repos, explicit Git identities,
local hooks and bounded deterministic shell fake agents; no bootstrap, auth,
network, implementation imports, process instrumentation or observer framework.
Prove runtime fail-first for every Added ID with compiling new stubs. Exercise
one complete spec-change lifecycle with distinct worker/auditor commands,
genuine eval fail-first, ordered phases, checked commits, SPEC/archive/provenance
and claim release, plus focused empty/approval/start/verification failure checks.
Reuse fixtures and representative assertions; own Verify stays under five minutes.
Do not multiply scenarios for unrelated features or rewrite existing evals.

The coordinator runs native start/context/done and owns all commits. Real CLI
workers use the fixed Sol/high profile; workers never call start/done/next/run,
commit or edit protected docs, phase/counters or other modules. Separate auditors
receive only an isolated snapshot of permitted inputs. Any new block or unresolved
finding follows normal audit/checkpoint bounds; no exception is authorized.

## Provenance

- evals 2026-10-28T23:34:54.251732Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=191 failing=12 passing=179
- audit 2026-10-28T23:36:43.906843Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings. Checked IDs: RUN-010, RUN-011, RUN-012, RUN-013, RUN-014, RUN-015.
The independent Sol/high auditor reviewed an isolated snapshot of permitted
inputs only; full command-log inspection confirms read-only supplied-file
access without implementation/history reads or test execution.
