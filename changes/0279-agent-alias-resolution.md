---
module: project
---

## Motivation

Project owners want to define each agent command once, select the worker default
by name, and optionally refer to a named command for the auditor. The runner
and doctor need the same selection rules.

## Interface

`agent.commands` maps aliases to shell templates or argv lists.
`agent.default` names the default worker command. `audit.agent` optionally
names the audit command from the same map.

Add `project.ResolveAgentCommand(config map[string]any, role, alias string)
(command any, usesWorkerDefault bool, err error)`. Pass an empty alias for
ordinary role selection. The boolean reports whether the auditor inherited
the worker default.

## Added

### new:explicit-agent-alias
For a worker with a nonempty alias argument, ResolveAgentCommand MUST select the matching entry in agent.commands. It MUST return U02 naming an unlisted alias and C01 identifying a malformed mapping or invalid selected command.

### new:default-agent-alias
With no explicit worker alias, ResolveAgentCommand MUST select agent.default from agent.commands when configured, otherwise use the inline worker template from agent.command. An invalid configured reference or simultaneous agent.default and agent.command values MUST return C01 identifying the configuration at fault.

### new:audit-agent-alias
For audit, ResolveAgentCommand MUST select audit.agent from agent.commands or use audit.command when supplied inline. When neither is configured, it MUST select the worker default and report that inheritance. An invalid configured reference or simultaneous audit.agent and audit.command values MUST return C01 identifying the configuration at fault.

## Modified

## Removed

## Decisions

## Notes

A configured reference is a nonempty string identifying a nonempty shell
template or argv list with a nonempty executable. Test legacy inline commands,
named defaults, explicit worker selection, audit inheritance, an independent
audit alias, and invalid or conflicting settings with local configuration
fixtures.
