---
module: assets
phase: impl
run: 20261115T032919-a0110b235a8184999dd461dd
---

## Motivation

Someone using `fireweed help drill` needs to find the agent selector and learn
where to define alternate commands. The doctor help should explain which
commands its probes exercise.

## Interface

## Added

### AST-069
CommandHelp("drill") MUST document --agent <alias>, identify agent.commands as the alias mapping and agent.default as the default selector, and show an example selecting a named command.

### AST-070
CommandHelp("doctor") MUST identify agent.default and audit.agent as the configured probe selectors and explain that audit inherits the worker default when its selector is omitted.

## Modified

## Removed

## Decisions

## Notes

Use aliases in examples that describe the configured harness, model, and
effort. Show `agent.default: codex-gpt-6-1-sol-high` and
`audit.agent: codex-gpt-6-astra-medium`. For a one-off drill, show
`fireweed drill contracts --agent opencode-deepseek-v4-1-flash-high`.
The alias is a lookup name. The command template supplies the model and effort.

The focused assets evals failed on the missing drill selector and example
for AST-069 and the missing probe selectors and audit fallback for AST-070.
`fireweed verify drill-agent-help` passed in evals. Both help entries can be
updated in the disposable `packages/assets/internal/content/command_help.go`.
## Audit

No findings; checked AST-069, AST-070

## Provenance

- evals 2026-11-15T03:33:54.23129Z agent=codex fireweed=0.2.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=151 failing=2 passing=149
- audit 2026-11-15T03:34:16.187128Z agent=codex fireweed=0.2.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

