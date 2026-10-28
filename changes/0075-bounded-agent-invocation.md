---
audit_pending:
    - RUN-003
    - RUN-008
audit_rounds: 1
bases: {}
module: runner
phase: evals
run: 20261028T211221-80fd9aefa1ac1fcec2faf6a6
set:
    interface:
        - api.go
        - model/**
---



## Motivation

The design's next/run loop launches a configured worker or auditor for each
workflow step, through wrap, with bounded execution and usage-limit detection.
Milestone 2's next observable step is a deterministic fake-agent queue run.
Provide its runner-owned invocation primitive now, then connect the queue loop
and CLI promptly after the five-change review. This is execution feature work,
not a drill or a process-observer framework.

## Interface

Add InvokeAgent(context.Context, AgentRequest) (AgentResult, error) to the thin
runner api.go facade, with alias types retained under model/**. AgentRequest
contains Snapshot project.Snapshot, Role string, Prompt string, Bin string and
Timeout time.Duration. Role is worker or audit. AgentResult contains Stdout and
Stderr strings, ExitCode int and UsageLimit bool. Keep execution implementation
under internal/ and shared data under model/, without importing the public
parent from its own implementation. Existing public signatures stay unchanged.
The projected set makes model/** retained interface before the first step.

## Added

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

## Modified



## Removed



## Decisions

An unset configured timeout defaults to 90 minutes. limit_patterns are
case-insensitive literal substrings, using the final 20 lines of stdout followed
by stderr with terminal newline separators trimmed.

## Rejected alternatives



## Notes

Read the design's Driving the state machine, Usage limits, Configuration and
Worktrees clauses and existing public process/project interfaces. Use
Snapshot.Config for agent/audit commands, wrap, timeout and limit_patterns, and
resolve git_common_dir from the ordinary fixture repository's actual common
Git directory. Reuse gitstore.Run for execution and process-group control rather
than duplicating native process machinery. Caller policy for queue selection,
approval, commits, attempts/saved refs and wait budgets belongs to the following
loop slice; do not implement it or the CLI in this change. This API returns
process errors; it adds no CLI diagnostic allocation.

Public-interface evals use ordinary temporary Git repos with explicit identities
and deterministic local shell scripts only, without bootstrap or authenticated
agents. Representative string/argv substitution, worker/audit/fallback, wrapper,
output, short configured/request timeouts, cancellation and last-20-line usage
checks cover the assigned behaviors. Reuse fixtures, avoid Cartesian matrices,
and keep own Verify under five minutes. An ordinary bounded fake descendant and
observable completion marker can exercise process-group cancellation; use no
interposers, process observers or platform branches. The 90-minute default is a
configuration choice, not a request for a 90-minute permanent test.

Only declared public interfaces/types, evals and new compiling disposable stubs
are allowed during evals. Added IDs need genuine runtime fail-first results,
not compilation or unrelated failures. Implementation freezes evals/interface;
workers never commit. Audit uses only the change, projected specs, evals,
retained interfaces and governing rules, without implementation/history reads;
re-audit only prior finding IDs and preserve prior passes and normal bounds.

## Provenance

- evals 2026-10-28T21:16:27.284693Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=177 failing=24 passing=153
- audit 2026-10-28T21:17:28.44516Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

- RUN-003: "Commands MUST accept shell strings or argv, quoting shell substitutions and preserving argv values literally." — No shell-string assertion checks that a `{git_common_dir}` containing spaces remains one argument.
- RUN-008: "InvokeAgent MUST set UsageLimit true exactly when a nonzero command exit has a configured limit_patterns match in its final 20 output lines." — No assertion checks a nonzero exit with a matching stdout line and no matching stderr line.
