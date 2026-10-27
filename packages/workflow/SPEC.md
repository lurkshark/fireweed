---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
    - checks
id_prefix: FLOW
interface:
    - api.go
    - model/**
name: workflow
---

# workflow

## Specs

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

## Decisions

[0012-manual-lifecycle] Provenance lines begin `- <phase>` and use `playbook=<full SHA-256>` plus `tests=N failing=N passing=N` for the reports verified by that Done call.
