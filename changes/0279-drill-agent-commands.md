---
bases:
    RUN-001: 7724c9f0
    RUN-056: f48368e1
module: runner
---

## Motivation

A project owner wants to compare reconstruction results from different harnesses
and models without editing the default worker command between drills. A drill
requires a clean checkout, so switching by editing the command also requires a
separate Git change.

## Interface

Add the optional string field `AgentAlias` to `runner.DrillRequest`. In
`fireweed.yaml`, `agent.commands` maps aliases to command templates.
`agent.default` names the default worker command. `agent.command` also
supports an inline default.

## Added

### new:drill-agent-resolution
Drill MUST resolve its selected command before creating scratch workspaces or running hooks or agents and propagate the resolution error.

### new:drill-agent-report-identity
For a named command selected explicitly or by default, Drill MUST derive Agent and Model from that command using the existing executable and model-flag conventions and write those values to the drill report.

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
