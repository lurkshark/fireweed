---
already_satisfied:
    - RUN-027
    - new:audit-work-preservation
bases:
    RUN-013: a6739afb
    RUN-027: 84080c43
module: runner
---



## Motivation

Early core-loop cutover must preserve independent audits. The design's Audit
step says the auditor reads only the change, spec and evals; SYSTEM requires
restricted automatic auditor inputs. NEW currently narrows auditPrompt but
invokes the auditor from the held checkout, where implementation, tooling and
Git history remain available. Run audits in a disposable input workspace and
validate the returned edit before importing it. Prove this through NEW next
with an inspectable fake auditor, then use the same path for the real-worker
self-hosting gate in docs/rebuild/ACCEPTANCE.md.

## Interface

No public API changes. Next changes where it invokes the audit role and how it
imports that role's result; InvokeAgent retains its Snapshot.Root contract.
Worker roles keep the held root. Existing audit tests that assume the source
checkout as audit cwd must be updated to assert the restricted workspace.

## Added

### new:audit-input-workspace
Next MUST run every automatic audit invocation in a fresh temporary workspace containing only the current change, projected module SPEC.md and module eval files at project-relative paths, plus independent Git metadata.
It MUST omit the held checkout's other files, including implementation, interfaces, hooks and configuration.

### new:audit-git-isolation
The audit workspace's Git metadata MUST contain none of the held repository's history, refs or saved attempts.
Audit-agent commits within this workspace MUST remain absent from the held checkout's history and refs.

### new:audit-submission-scope
Next MUST accept only edits to the Audit section of the copied current change.
Edits elsewhere in that change, edits to copied specs or evals, or added or deleted files outside Git metadata MUST reject the submission without copying any workspace edits to the held checkout.

### new:audit-work-preservation
When audit workspace edits are permitted, Next MUST copy its Audit section back before retrying or returning, including agent failures and time or usage stops.
Copying audit work MUST not itself advance a phase or replace other held change contents.

### new:audit-workspace-cleanup
Next MUST remove every temporary audit workspace after invocation, including after invalid submissions, agent failures and time or usage stops.

## Modified

### RUN-013
For each phase, Next MUST invoke the worker role for evals and impl from the held checkout and the audit role from its restricted input workspace, using approved agent configuration and request Bin and Timeout.
Worker prompts MUST include full workflow context.
Audit prompts MUST include the change, projected spec, eval contents and audit playbook without implementation contents.

### RUN-027
When an evals or impl worker creates commits, Next MUST soft-reset the held checkout's HEAD to its phase base after invocation and before verification or returning an invocation failure, retaining the step's working-tree and index work.

## Removed



## Decisions



## Rejected alternatives

Prompt-only restrictions: the current prompt omits implementation, but leaves
the auditor in the checkout containing it.
A linked worktree: it exposes the source repository's Git objects and refs.

## Notes

RUN-027 is narrowed to preserve existing worker recovery, and permitted partial
audit-work retention already exists; those two IDs are already_satisfied. The
new workspace, input/history isolation, submission validation and cleanup still
require genuine fail-first evidence.

Keep this change within runner. A valid synthetic Git repository may be needed
for InvokeAgent's preflight; it must contain only audit inputs and its own
metadata. Configured agent settings, wrapping, Bin, timeout and usage handling
come from the approved held snapshot, without copying project tooling into the
audit workspace. This is input-workspace isolation, not a new OS sandbox.

Validate workspace file contents against the originally supplied inputs even
if the auditor commits its edits. Preserve only allowed Audit work on error or
limit exits; rejected edits must not reach the held checkout. Preserve existing
worker commit recovery and its RUN-028 warnings; retarget audit-commit fixtures
to the isolated history rule. Keep independent audit rounds, approval bounds,
phase-zone enforcement and invocation result capture enabled.

Use public APIs, temporary fixture repositories and bounded fake auditors to
inspect cwd, supplied files, source-history absence, accepted/rejected edits,
partial Audit retention and cleanup. Reuse representative cases and existing
fixtures; no process instrumentation, real-agent evals or bootstrap source.
