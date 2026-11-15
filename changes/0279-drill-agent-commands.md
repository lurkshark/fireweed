---
bases:
    RUN-001: 7724c9f0
    RUN-056: f48368e1
module: runner
phase: audit
run: 20261115T020452-a5072c6fd3fc5168fdf45533
---

## Motivation

A project owner wants to compare reconstruction results from different harnesses
and models without editing the default worker command between drills. A drill
requires a clean checkout, so switching by editing the command also requires a
separate Git change.

## Interface

Add the optional string field `AgentAlias` to `runner.DrillRequest`. In
`fireweed.yaml`, `agent.commands` maps aliases to command templates.
`agent.default` names the default worker command.

## Added

### RUN-073
Drill MUST resolve its selected command before creating scratch workspaces or running hooks or agents and propagate the resolution error.

### RUN-074
Drill MUST derive Agent and Model from the selected command using the existing executable and model-flag conventions and write those values to the drill report.

## Modified

### RUN-001
InvokeAgent MUST select the command for the requested role from Snapshot.Config using project.ResolveAgentCommand and propagate any resolution error.

### RUN-056
Drill MUST use the configured worker default or the command selected by AgentAlias. Every invocation MUST use the configured wrapper and timeout through InvokeAgent and retain its AgentResult in Steps in order. Non-usage failures and verification retries MUST share the configured attempts bound, retain rebuilt work and include the preceding reconstruction or own-eval failure in retry prompts.

## Removed

## Decisions

## Notes

Use bounded local fake agents to exercise the worker and audit aliases,
the drill default, explicit selection across retries, and the executable
and model in the result and report. Give each command distinct metadata so
the fixture can detect a report that identifies the wrong command. Use the
snapshot already supplied to each invocation as the configuration source.
Move the runner eval fixtures to named commands and retain their existing
behavior assertions.

Fail-first against the current implementation: RUN-001 rejected named worker
and audit commands with `empty command argv` and lost the resolver's C01 error.
RUN-056 could not invoke the default or explicit command. RUN-073 created the
scratch worktree and ran both setup hooks before returning an invocation error
instead of the resolver's U02 or C01. RUN-074 wrote `unknown` for both fields
in the result and report for both selections.

`runner.DrillRequest.AgentAlias` is the declared interface addition. The
remaining work fits in the disposable files `internal/invoke/agent.go` and
`internal/invoke/command.go` for selection and executable naming, and
`internal/drill/drill.go` and `internal/drill/report.go` for drill preflight,
retry selection, and report metadata.
## Provenance

- evals 2026-11-15T02:11:13.701057Z agent=codex fireweed=0.2.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=199 failing=174 passing=25

