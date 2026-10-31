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
Next MUST return Empty true without invoking agents when workflow.Eligible
returns no changes, including an empty or entirely waiting queue.

### RUN-011
When workflow.Eligible returns changes, Next MUST select its first result and
start it through workflow.Start, propagating start failures without invoking
an agent.

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
For a non-usage-limit agent invocation failure, Next MUST return an error and
captured invocation results without advancing the phase or releasing its held
claim.

### RUN-016
Run without a positive Max MUST continue single-change execution until no
eligible changes remain or its For bound stops it, returning applied NextResults
in execution order.

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

### RUN-023
For exhausted phase verification with the worktree's approved claim and an
unchanged HEAD, Next MUST save and block the step through the lifecycle,
leaving a clean unclaimed checkout and a state-only block commit.

### RUN-024
After that blocking succeeds, Next MUST return the blocked lifecycle State and
ordered captured invocations without changing the run, assigned IDs or current
phase. Its error MUST retain the last verification cause and identify the saved
attempt ref.

### RUN-025
When single-change execution returns an error after successfully saving and
blocking its change with a clean unclaimed checkout, Run MUST continue processing
eligible changes. Run MUST exclude that blocked change from Applied.

### RUN-026
If a change was blocked during Run, Run MUST return an error when eligible work
is exhausted or Max is reached. That error MUST retain a blocked change's
verification cause and saved attempt ref.

### RUN-027
When a phase agent creates commits, Next MUST soft-reset HEAD to the held phase
base after invocation and before verification or returning an invocation failure,
retaining the step's working-tree and index work.

### RUN-028
For each such recovery, Next MUST return a P05 warning in State.Diagnostics
identifying the phase and that work was retained. It MUST preserve these warnings
across phase advancement and on an invocation-failure return.

### RUN-029
Next MUST read limit_wait and limit_budget as durations, defaulting to 30m and
4h when unset. It MUST reject malformed durations, a nonpositive wait or negative
budget with C01 before claiming a fresh change or invoking an agent.

### RUN-030
Subject to the usage budget and Run's For bound, an AgentResult with UsageLimit
true MUST make Next wait limit_wait and repeat the same phase attempt without
spending a failed attempt. It MUST retain earlier step work and captured
invocations, keeping existing verification feedback.

### RUN-031
When another usage wait would exceed limit_budget, Next MUST return a clean
usage-stop error identifying completed wait, next wait and budget. It MUST retain
the held State, work and ordered Steps without advancing, blocking, committing
or releasing the claim.

### RUN-032
Run MUST account for total usage wait across phases and changes in one
invocation. On a clean usage stop it MUST retain already-applied results and
propagate the stop classification.

### RUN-033
When the calling worktree already holds a queued step, Next MUST resume that
owned State using its approved snapshot without a fresh Start, retaining its
run, phase, assigned IDs and work. Existing approval and unapproved-work bounds
MUST remain enforced for starting another change.

### RUN-034
IsUsageStop MUST recognize clean usage-stop errors through wrapping and return
false for nil or ordinary failures.

### RUN-035
If Run blocked a change before a usage stop, it MUST return failure with the
blocking cause/ref and stop reason retained. IsUsageStop MUST return false for
that combined failure so the caller cannot report it as a clean stop.

### RUN-036
Run MUST treat For zero as unbounded and positive For as one elapsed-time
budget for the whole invocation, including agent and verification time across
phases and changes.

### RUN-037
When its For bound stops execution, Run MUST return a nonnil clean time-stop
error naming the limit, with already-applied results retained. A step interrupted
during agent invocation or verification MUST keep its work, claim, run, phase
and assigned IDs without advancing, blocking, committing or releasing that step.

### RUN-038
When a usage wait fits its usage budget but would pass Run's For deadline,
execution MUST return a clean time stop before waiting. It MUST preserve the
held step's work and identity.

### RUN-039
IsTimeStop MUST recognize clean time-stop errors through wrapping. It MUST
return false for nil, usage stops, ordinary failures and parent-context
cancellation.

### RUN-040
If Run blocked a change before a time stop, it MUST return ordinary failure
retaining the blocking cause/ref and time-stop reason. IsTimeStop MUST return
false for that combined failure.

### RUN-041
Next MUST return an error for which IsApprovalRequired reports true when the selected change file is on neither the local nor the shared base and AllowUnmerged is false.

### RUN-042
IsApprovalRequired MUST report false for Next errors caused by anything other than a missing approval.

## Decisions

[0075-bounded-agent-invocation] An unset configured timeout defaults to 90 minutes. limit_patterns are
case-insensitive literal substrings, using the final 20 lines of stdout followed
by stderr with terminal newline separators trimmed.
