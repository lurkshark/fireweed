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
For each phase, Next MUST invoke the worker role for evals and impl from the held checkout and the audit role from its restricted input workspace, using approved agent configuration and request Bin and Timeout.
Worker prompts MUST include full workflow context.
Audit prompts MUST include the change, projected spec, eval contents and audit playbook without implementation contents.

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
When an evals or impl worker creates commits, Next MUST soft-reset the held checkout's HEAD to its phase base after invocation and before verification or returning an invocation failure, retaining the step's working-tree and index work.

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

### RUN-043
On a verification retry, Next MUST append a detail section identifying the module, classname, testcase name and Case.Failure for failed cases from every reported module, including dependents, subject to RUN-044.
It MUST select cases by lexical module name and then report order, excluding passing and skipped cases.

### RUN-044
The failed-case detail section MUST fit within 16,384 Unicode code points, including labels and truncation markers.
Each Case.Failure excerpt MUST contain at most its first 4,096 code points.
Shortened diagnostic text or omitted cases MUST be marked.

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

### RUN-050
Drill MUST reject an unknown module, dirty invoking checkout or any live shared change claim before setup or agent invocation, leaving the invoking checkout and refs unchanged.

### RUN-051
Drill MUST create a detached scratch worktree at the invoking HEAD under the common Git directory's fireweed/worktrees directory and run the configured setup hook there with FIREWEED_BIN set to request Bin. Setup failure MUST prevent agent invocation and return its captured failure.

### RUN-052
Drill MUST delete only target-module files classified disposable by the approved zone precedence before reconstruction. It MUST retain every other target-module file and every other module's files in the verification worktree.

### RUN-053
Drill MUST invoke reconstruction in a separate input workspace containing the target's retained files and only interface files from its transitive dependencies, plus SYSTEM.md and fresh independent Git metadata. It MUST exclude original implementation, other-module internals, source history, saved attempts and binaries from those inputs.

### RUN-054
The reconstruction prompt MUST include the target spec and Decisions, identify its retained interfaces and evals, and require rebuilding only disposable paths without reading outside supplied inputs. It MUST request separate Contradicted Decisions and Gaps sections in DRILL_NOTES.md, each gap naming the nearest spec ID or "no spec" and a target module.

### RUN-055
Drill MUST accept reconstruction edits only to target disposable paths and DRILL_NOTES.md, rejecting retained-file changes, deletions or other additions before importing code. It MUST import accepted rebuilt disposable files into the verification worktree without changing retained files or the invoking implementation.

### RUN-056
Drill MUST use the configured worker command, wrapper and timeout through InvokeAgent, retaining ordered AgentResults in Steps. Non-usage failures and verification retries MUST share the configured attempts bound, retain rebuilt work and include the preceding failure in retry prompts.

### RUN-057
Drill MUST treat usage-limit responses as the same attempt, applying configured limit_wait and limit_budget without spending a failed attempt. An exhausted usage budget MUST return an IsUsageStop error with Result stopped; cancellation or timeout MUST return an error and stop reconstruction.

### RUN-058
After accepted reconstruction, Drill MUST run the target module's eval hook and require passing enforced cases and coverage through checks. It MUST retain the complete returned own report in Own, including failed testcase details when verification fails.

### RUN-059
Drill MUST evaluate every direct and transitive dependent in lexical module-name order regardless of verify_dependents, retaining each report in Dependents. Every dependent MUST satisfy passing enforced cases and coverage for a passing drill, even when own or another dependent's evals fail.

### RUN-060
After reconstruction has been invoked, Drill MUST write .fireweed/drills/<module>.md with date, base_commit, agent, model and result frontmatter and Own evals, Dependent evals, Contradicted Decisions and Gaps sections. Eval entries MUST identify module, testcase, spec IDs and result; agent observations and verification failures MUST remain visible, with each gap identifying a nearest spec ID or "no spec" and target module.

### RUN-061
After accepted reconstruction, Drill MUST retain the rebuilt verification tree at refs/fireweed/drills/<module>/<date>, returning that ref without moving the invoking HEAD. A passing result MUST require accepted reconstruction and passing own and all dependent evals; failure MUST return a nonnil error with Result fail.

### RUN-062
Drill MUST remove its verification worktree and input workspace on success, agent or verification failure, invalid edits, usage stop and cancellation. Cleanup failure MUST be returned as an error rather than reported as a passing drill.

### RUN-063
InvokeAgent MUST persist the invocation's complete stdout and stderr, including partial output on command failure, timeout or cancellation, at .fireweed/runs/<run>/<phase>-<attempt>.log under its log project root. Repeated invocations at the same path MUST append without discarding earlier records. AgentResult.LogPath MUST identify that file by an absolute path.

### RUN-064
Next MUST assign each invocation the held change's run, current phase and one-based phase attempt, writing its log under the invoking project's .fireweed/runs directory. Audit logs MUST stay in that invoking project, outside the restricted audit input workspace.

### RUN-065
Drill MUST allocate one log run identity per invocation and reuse it across reconstruction retries and usage waits. Its logs MUST use phase drill, the current one-based reconstruction attempt and the invoking project's root, outside scratch and reconstruction input workspaces.

### RUN-066
InvokeAgent MUST accept a log project root, run, phase and positive one-based attempt in AgentRequest. Unset fields MUST default to Snapshot.Root, a fresh run identifier, request Role and attempt 1, respectively.

## Decisions

[0075-bounded-agent-invocation] An unset configured timeout defaults to 90 minutes. limit_patterns are
case-insensitive literal substrings, using the final 20 lines of stdout followed
by stderr with terminal newline separators trimmed.

- agent-invocation-logs: Agent logs append length-labeled stdout and stderr records after each invocation, preserving both streams' bytes even without final newlines. A log write error is returned alongside any agent error while the captured result remains available.
