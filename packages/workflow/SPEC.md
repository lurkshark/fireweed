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
Start MUST atomically create the local claim for the selected queued change.

### FLOW-002
Verify MUST reject edits forbidden by the approved projected phase zones before running hooks, checking differences since the held phase base.

### FLOW-003
Done MUST require full verification to succeed before committing a phase transition.

### FLOW-004
Successful impl Done MUST apply Added/Modified/Removed, set values and Decisions (attributed to the slug), including Removed decisions, then archive with phase applied and commit together after static validation. Every phase commit MUST advance the claim's phase base, freezing prior eval/interface edits and releasing the exact held claim upon application. Failed verification MUST leave the phase unadvanced and create no commit.

### FLOW-005
Each successful Done MUST append provenance with phase, date, manual/configured agent, version, playbook hash and result counts, retaining reports/output under the run directory. Context MUST include the change, projected spec, interface contents, SYSTEM and matching playbook, listing other relevant files by path/size or including contents with full. State MUST expose the held projected snapshot so composing commands use the approved graph during hooks.

### FLOW-006
Start MUST include the approved projected project snapshot in its returned State.

### FLOW-007
Start MUST replace each local Added handle throughout the change with the next available spec ID for its module.

### FLOW-008
Start MUST assign a unique run identifier when starting an unstarted change.

### FLOW-009
Start MUST report P04 when neither the local base nor the shared base contains the change file.

### FLOW-010
Start MUST reject a worktree containing uncommitted changes.

### FLOW-011
Start MUST reject a requested change when an earlier unblocked queued change exists.

### FLOW-012
Start MUST return the canonical absolute worktree root in State.Snapshot.Root.

### FLOW-013
Start MUST set phase to evals when starting an unstarted spec change.

### FLOW-014
Start MUST leave HEAD unchanged.

### FLOW-015
Start MUST reject when the shared-base unapproved archive count reaches authoring.max_unapproved, whose default is 5.

### FLOW-016
Start MUST reject a competing worktree that attempts to claim the same change.

### FLOW-017
Current MUST resolve only a claim owned by the calling worktree.

### FLOW-018
Start MUST use the current HEAD as the initial phase base of its claim.

### FLOW-019
While a claim is held, Current MUST return the project snapshot approved at Start.

### FLOW-020
Verify MUST reject an unapproved alteration of the active change's author fields before running a hook.

### FLOW-021
Verify MUST reject an unapproved alteration of the active change's recorded runtime state before running a hook.

### FLOW-022
Whenever Verify runs the target module's evals it MUST require coverage of every enforced current spec.

### FLOW-023
During evals Verify MUST require genuine fail-first evidence for every enforced change ID that is not already_satisfied.

### FLOW-024
During impl Verify MUST reject a failing current-module testcase that references a non-quarantined spec.

### FLOW-025
Full impl Verify MUST use verify_dependents with a default of direct to select dependent modules.

### FLOW-026
Full impl Verify MUST reject a failed enforced testcase from a selected dependent module.

### FLOW-027
Full impl Verify MUST require boundary verification to succeed.

### FLOW-028
Verify MUST select hooks from the held approved projected project snapshot.

### FLOW-029
Quick Verify MUST skip dependent-module evals.

### FLOW-030
Quick Verify MUST skip boundary verification.

### FLOW-031
Verify MUST leave HEAD unchanged.

### FLOW-032
Done MUST commit a successfully verified evals phase and advance it to audit when audit is enabled.

### FLOW-033
Done MUST commit a successfully verified evals phase and advance it to impl when audit.enabled is false.

### FLOW-034
Audit Done MUST discard a finding whose spec ID is absent or whose quoted clause is not an exact substring of that spec.

### FLOW-035
On re-audit, Audit Done MUST discard findings for spec IDs that had no retained finding in the preceding audit round.

### FLOW-036
Audit Done MUST require every enforced ID in the change's Added, Modified or fixes entries to appear in a No findings submission's checked-ID list.

### FLOW-037
Retained testable audit findings within the round limit MUST return the change to evals.

### FLOW-038
An audit return to evals MUST increment and persist audit_rounds.

### FLOW-039
Audit Done MUST block the change when retained findings would exceed audit.max_rounds, whose default is 2.

### FLOW-040
Audit Done MUST block the change when a retained finding marks its quoted clause untestable.

### FLOW-041
Audit Done MUST require No findings with a complete checked-ID list before advancing to impl.

### FLOW-042
When audit blocks a change, Done MUST include the blocking findings in its Questions section.

### FLOW-043
Audit blocking MUST release the calling worktree's held change claim.

## Decisions

[0012-manual-lifecycle] Provenance lines begin `- <phase>` and use `playbook=<full SHA-256>` plus `tests=N failing=N passing=N` for the reports verified by that Done call.
