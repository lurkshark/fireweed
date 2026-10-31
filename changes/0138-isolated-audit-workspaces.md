---
already_satisfied:
    - RUN-027
    - RUN-048
audit_pending:
    - RUN-013
    - RUN-048
audit_rounds: 1
bases:
    RUN-013: a6739afb
    RUN-027: 84080c43
module: runner
phase: evals
run: 20261031T085104-0a74a70a942f28f9dee76ea3
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

### RUN-045
Next MUST run every automatic audit invocation in a fresh temporary workspace containing only the current change, projected module SPEC.md and module eval files at project-relative paths, plus independent Git metadata.
It MUST omit the held checkout's other files, including implementation, interfaces, hooks and configuration.

### RUN-046
The audit workspace's Git metadata MUST contain none of the held repository's history, refs or saved attempts.
Audit-agent commits within this workspace MUST remain absent from the held checkout's history and refs.

### RUN-047
Next MUST accept only edits to the Audit section of the copied current change.
Edits elsewhere in that change, edits to copied specs or evals, or added or deleted files outside Git metadata MUST reject the submission without copying any workspace edits to the held checkout.

### RUN-048
When audit workspace edits are permitted, Next MUST copy its Audit section back before retrying or returning, including agent failures and time or usage stops.
Copying audit work MUST not itself advance a phase or replace other held change contents.

### RUN-049
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

The owner authorized a fresh independent audit on 2026-10-31 after incomplete
checked-ID submissions exhausted the prior audit attempts. Check all seven
changed IDs: RUN-013, RUN-027, RUN-045, RUN-046, RUN-047, RUN-048 and RUN-049.
already_satisfied exempts fail-first evidence, not audit review. After checking
their assertions, a passing Audit response lists all seven in Checked IDs;
otherwise report clause findings normally. Existing bounds remain unchanged.

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

## Provenance

- evals 2026-10-31T09:09:45.579501Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=378 failing=14 passing=364
- audit 2026-10-31T11:22:12.762795Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.1 playbook=5cb9f9dac24b0b6bb1b970edd4fab4406d2b4d0435ddefc3ef59a2cfdd267f18 tests=0 failing=0 passing=0

## Audit

- RUN-013: "For each phase, Next MUST invoke the worker role for evals and impl from the held checkout and the audit role from its restricted input workspace, using approved agent configuration and request Bin and Timeout." — No assertion distinguishes approved agent configuration from later unapproved changes to the held checkout's agent settings.
- RUN-048: "When audit workspace edits are permitted, Next MUST copy its Audit section back before retrying or returning, including agent failures and time or usage stops." — No assertion checks that a permitted Audit edit reaches the held change before an audit verification retry.
