---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
    - checks
    - workflow
id_prefix: RUN
interface:
    - api.go
    - model/**
name: runner
---

# runner

## Specs

### RUN-001
InvokeAgent MUST select agent.command for the worker role and audit.command for the audit role, falling back to agent.command when audit.command is unset.

### RUN-002
InvokeAgent MUST run the selected command from Snapshot.Root.

### RUN-003
InvokeAgent MUST expand {prompt} and {git_common_dir} once in command templates and wrapper argv. Commands MUST accept shell strings or argv, quoting shell substitutions and preserving argv values literally.

### RUN-004
InvokeAgent MUST prepend the configured agent.wrap argv prefix for both worker and audit commands.

### RUN-005
InvokeAgent MUST set the command's FIREWEED_BIN environment variable to the request Bin.

### RUN-006
InvokeAgent MUST retain the command's stdout, stderr and exit code in AgentResult.

### RUN-007
InvokeAgent MUST use a positive request Timeout in preference to the configured timeout, using the configured timeout when the request is unset. It MUST stop the command's process group when that bound expires or its context is cancelled.

### RUN-008
InvokeAgent MUST set UsageLimit true exactly when a nonzero command exit has a configured limit_patterns match in its final 20 output lines.

### RUN-009
InvokeAgent MUST return command-launch, timeout or cancellation errors without discarding captured output.

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
When invocation fails, or phase verification still fails after its attempt
allowance, Next MUST return an error and captured invocation results without
advancing that phase.

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

### RUN-021
Next MUST permit at most the configured attempts count of agent invocations
for a phase after verification failures, defaulting to two when unset. Each
phase MUST have its own attempt allowance.

### RUN-022
When phase verification fails and an attempt remains, Next MUST retain the
earlier work and retry that phase. The retry prompt MUST include the last
verification output.

## Decisions

[0075-bounded-agent-invocation] An unset configured timeout defaults to 90 minutes. limit_patterns are
case-insensitive literal substrings, using the final 20 lines of stdout followed
by stderr with terminal newline separators trimmed.
