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
Successful impl Done MUST apply Added, Modified and Removed spec edits to the module's SPEC.md.

### FLOW-005
Context MUST include the change, projected spec, interface contents, SYSTEM and matching playbook. It MUST list other relevant files by path and size, or include their contents when full is requested.

### FLOW-006
Start MUST include the approved projected project snapshot in its returned State.

### FLOW-007
Start MUST replace each local Added handle throughout the change with the next available spec ID for its module.

### FLOW-008
Start MUST assign a unique run identifier when starting an unstarted change.

### FLOW-009
Start MUST report P04 as a warning when neither the local base nor the shared base contains the change file.

### FLOW-010
Start MUST reject a worktree containing uncommitted changes.

### FLOW-011
Start MUST reject a requested change unless it is the first change returned by
Eligible. With no requested target, Start MUST start that first eligible
change. When Eligible returns no change, Start MUST reject with P01.

### FLOW-012
Start MUST return the canonical absolute worktree root in State.Snapshot.Root.

### FLOW-013
Start MUST set phase to evals when starting an unstarted spec change.

### FLOW-014
Start MUST leave HEAD unchanged.

### FLOW-015
Start MUST reject with P07 when the shared-base unapproved archive count reaches authoring.max_unapproved, whose default is 5.

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

### FLOW-044
Successful impl Done MUST apply the change's set values to module metadata.

### FLOW-045
Successful impl Done MUST apply Decisions attributed to the change slug and remove entries quoted in Removed decisions.

### FLOW-046
Each successful Done MUST advance the phase base to its new commit.

### FLOW-047
Successful impl Done MUST release only the exact held claim upon application.

### FLOW-048
Failed Done verification MUST leave the phase unadvanced and create no commit.

### FLOW-049
Impl Done MUST require the applied project to pass static validation before committing.

### FLOW-050
Successful impl Done MUST move the active change to the archive with phase applied and commit it together with the application.

### FLOW-051
Each successful Done MUST append provenance containing phase, date, manual or configured agent, version, playbook hash and result counts.

### FLOW-052
Successful Done MUST retain verification reports and process output under the run directory.

### FLOW-053
State MUST expose the held projected snapshot so composing commands use the approved graph during hooks.

### FLOW-054
Start MUST prefix dirty-worktree and queue-eligibility rejections with P01.

### FLOW-055
Verify and Done MUST prefix unsupported-phase rejections with P01.

### FLOW-056
Start and Current MUST prefix existing claim ownership or contention failures with P02.

### FLOW-057
Verify and Done MUST prefix unapproved author, runtime or audit-state alteration errors with P03 before running hooks.

### FLOW-058
Start, Verify and Done MUST prefix stale-base rejections with P05, retaining the underlying cause.

### FLOW-059
Audit Done MUST prefix malformed or incomplete No findings submission errors with P09.

### FLOW-060
Done MUST prefix its existing blocked-change error with P06 while preserving the committed blocking state and claim release.

### FLOW-061
Current and Context MUST prefix unknown requested change lookup errors with U02.

### FLOW-062
Lifecycle operations MUST preserve assigned static, eval, zone and boundary causes when returning their existing failures, except for lifecycle stale-base classification P05.

### FLOW-063
Block MUST require the calling worktree's current claim with unchanged approved
metadata and phase. It MUST reject a changed HEAD rather than blocking a stale
phase base.

### FLOW-064
Block MUST save the held step's tracked and untracked work as a commit at the
next unused positive refs/fireweed/attempts/<slug>/<n>, preserving prior attempt
refs. BlockResult MUST identify that ref and its commit OID.

### FLOW-065
After saving the attempt, Block MUST restore the step's tracked files, untracked
files and index to the phase base, removing the step's edits from the checkout.

### FLOW-066
Block MUST commit only the held change's assigned IDs, run, current phase and
requested blocked reason over the phase base. The block commit MUST preserve
the approved author and already-committed audit content and exclude step work.

### FLOW-067
On successful blocking, Block MUST release the current claim and return the
blocked State with no held ClaimOID. The resulting working tree MUST be clean.

### FLOW-068
Eligible MUST return unblocked queued changes not held back by earlier blocked
changes, in queue order. Empty or entirely waiting queues MUST return an empty
result without an error.

### FLOW-069
Eligible MUST omit a later change targeting the same module as an earlier
blocked change.

### FLOW-070
Eligible MUST omit a later change whose module depends directly or transitively
on an earlier blocked change's module.

### FLOW-071
When a later change has a nonempty Interface section, Eligible MUST omit it if
an earlier blocked change's module depends directly or transitively on its
module. With an empty Interface, that dependency change MUST remain eligible
unless another blocking rule applies.

### FLOW-072
For each candidate, Eligible MUST apply the blocking rules to the graph
projected through that candidate, including preceding queued set edits and its
own set edits.

## Decisions

[0012-manual-lifecycle] Provenance lines begin `- <phase>` and use `playbook=<full SHA-256>` plus `tests=N failing=N passing=N` for the reports verified by that Done call.
