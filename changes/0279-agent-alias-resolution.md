---
module: project
phase: impl
run: 20261115T011617-65309e5591fc11bf209bcb93
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

### PRJ-086
For a worker with a nonempty alias argument, ResolveAgentCommand MUST select the matching entry from the agent.commands mapping. It MUST return U02 naming an unlisted alias and C01 identifying a malformed mapping or invalid selected command.

### PRJ-087
With no explicit worker alias, ResolveAgentCommand MUST select the command named by agent.default in agent.commands. A missing or invalid default reference MUST return C01 identifying the setting at fault.

### PRJ-088
For audit, ResolveAgentCommand MUST select the command named by audit.agent in agent.commands when that reference is configured. An invalid audit reference MUST return C01 identifying the setting at fault. An omitted audit reference MUST use the worker default and report that inheritance.

## Modified

## Removed

## Decisions

## Notes

A configured reference is a nonempty string identifying a nonempty shell
template or argv list with a nonempty executable. Test named defaults,
explicit worker selection, audit inheritance, an independent audit alias, and
missing or invalid references with local configuration fixtures.

The new project evals compiled and failed on their resolver assertions for
PRJ-086, PRJ-087 and PRJ-088. `fireweed verify agent-alias-resolution` passed
the evals gate. The public facade delegates to the disposable stub at
`packages/project/internal/fixture/agent_command.go`, where the implementation
phase can add command resolution.
## Audit

No findings; checked PRJ-086, PRJ-087, PRJ-088

## Provenance

- evals 2026-11-15T01:19:24.737368Z agent=codex fireweed=0.2.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=305 failing=27 passing=278
- audit 2026-11-15T01:19:42.411889Z agent=codex fireweed=0.2.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

