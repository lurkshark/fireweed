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
For a worker with a nonempty alias argument, ResolveAgentCommand MUST select the matching entry from the agent.commands mapping. It MUST return U02 naming an unlisted alias and C01 identifying a malformed mapping or invalid selected command.

### new:default-agent-alias
With no explicit worker alias, ResolveAgentCommand MUST select the command named by agent.default in agent.commands. A missing or invalid default reference MUST return C01 identifying the setting at fault.

### new:audit-agent-alias
For audit, ResolveAgentCommand MUST select the command named by audit.agent in agent.commands when that reference is configured. An invalid audit reference MUST return C01 identifying the setting at fault. An omitted audit reference MUST use the worker default and report that inheritance.

## Modified

## Removed

## Decisions

## Notes

A configured reference is a nonempty string identifying a nonempty shell
template or argv list with a nonempty executable. Test named defaults,
explicit worker selection, audit inheritance, an independent audit alias, and
missing or invalid references with local configuration fixtures.
