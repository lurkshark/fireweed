---
module: app
---

## Motivation

People running drills need to choose a configured harness and model at
invocation time. Expose the alias supported by the preceding runner change.

## Interface

`fireweed drill <module> [--agent <alias>] [--json]`

## Added

### new:drill-agent-option
drill --agent <alias> MUST pass the supplied alias to runner.Drill as AgentAlias. A missing, empty or repeated --agent value MUST return exit 2 and corrective U01 before running hooks or agents.

### new:drill-agent-lookup-error
An unknown drill agent alias MUST return exit 2 with U02 naming the alias and pointing to the agent.commands mapping in the project configuration. In JSON mode, drill MUST emit this diagnostic in a schema 1 envelope with null data and empty stderr.

## Modified

## Removed

## Decisions
