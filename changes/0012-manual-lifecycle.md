---
audit_pending:
    - FLOW-005
audit_rounds: 2
bases: {}
module: workflow
phase: impl
run: 20261027T064001-843425dd0d42609ccb1f39a7
set:
    interface:
        - api.go
        - model/**
---



## Motivation

The design's manual lifecycle finish line needs real claims, phase checks,
audit convergence, apply/archive and inspectable evidence. Connect the existing
project/Git/checks primitives now; unattended retries, all change kinds and
attempt recovery remain later work.

## Interface

Retain data-only State with Change project.Change, Snapshot project.Snapshot,
Run/Phase/Head/PhaseBase/ClaimOID strings and Diagnostics []project.Diagnostic.
Request holds Root/Change/Bin/Agent strings, Quick bool and Timeout time.Duration.
Result holds State, Reports map[string]checks.Report and Diagnostics.
Thin functions: Start(root,change string) (State,error),
Current(root,change string) (State,error),
Verify(ctx context.Context,request Request) (Result,error),
Done(ctx context.Context,request Request) (Result,error), and
Context(root,change string,full bool) (string,error).
Behavior lives under internal; dependency types use public imports.

## Added

### FLOW-001
Start MUST atomically claim a clean eligible queued change for its canonical worktree, retain the approved projected snapshot/base, assign next free module IDs throughout the change, and write unique run plus evals without committing. It MUST warn when approval exists on neither local nor shared base and refuse when shared-base unapproved archives reach max_unapproved (default 5). Claims and Current MUST reject competing owners and resolve only this worktree's held change.

### FLOW-002
Verify MUST reject altered approved author metadata/state and forbidden phase zones before hooks, using the held snapshot and diff since phase base. Evals MUST require coverage and fail-first for enforced Added/Modified/fixes IDs except already_satisfied. Impl MUST require current coverage/pass and configured dependent passes (direct default, all/none supported), plus boundaries; quick skips dependents/boundaries, and Verify never commits.

### FLOW-003
Done MUST fully verify and commit evals into audit by default, or impl when audit is disabled. Audit MUST validate exact ID/clause findings, discard invalid findings and new IDs on re-audit, and require checked IDs with No findings to pass. Valid findings MUST increment persistent audit_rounds and return to evals; exceeding max_rounds (default 2), or untestable findings, blocks with Questions and releases the claim.

### FLOW-004
Successful impl Done MUST apply Added/Modified/Removed, set values and Decisions (attributed to the slug), including Removed decisions, then archive with phase applied and commit together after static validation. Every phase commit MUST advance the claim's phase base, freezing prior eval/interface edits and releasing the exact held claim upon application. Failed verification MUST leave the phase unadvanced and create no commit.

### FLOW-005
Each successful Done MUST append provenance with phase, date, manual/configured agent, version, playbook hash and result counts, retaining reports/output under the run directory. Context MUST include the change, projected spec, interface contents, SYSTEM and matching playbook, listing other relevant files by path/size or including contents with full. State MUST expose the held projected snapshot so composing commands use the approved graph during hooks.

## Modified



## Removed



## Decisions

Provenance lines begin `- <phase>` and use `playbook=<full SHA-256>` plus `tests=N failing=N passing=N` for the reports verified by that Done call.

## Provenance

- evals 2026-10-27T06:51:15.53473Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=108 failing=28 passing=80
- audit 2026-10-27T13:47:58.117575Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- evals 2026-10-27T13:53:38.862633Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=108 failing=28 passing=80
- audit 2026-10-27T13:55:48.99758Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- evals 2026-10-27T13:56:45.520373Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=108 failing=28 passing=80
- audit 2026-10-27T13:58:51.538243Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings; checked FLOW-005
