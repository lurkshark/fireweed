---
module: assets
---

## Motivation

Someone using `fireweed help drill` needs to find the agent selector and learn
where to define alternate commands. The doctor help should explain which
commands its probes exercise.

## Interface

## Added

### new:drill-agent-help
CommandHelp("drill") MUST document --agent <alias>, identify agent.commands as the alias mapping and agent.default or agent.command as the default, and show an example selecting a named command.

### new:doctor-agent-help
CommandHelp("doctor") MUST describe the worker default and explain that audit.agent selects a named command and audit.command supplies an inline audit command.

## Modified

## Removed

## Decisions

## Notes

Use aliases in examples that describe the configured harness, model, and
effort. Show `agent.default: codex-gpt-6-1-sol-high` and
`audit.agent: codex-gpt-6-astra-medium`. For a one-off drill, show
`fireweed drill contracts --agent opencode-deepseek-v4-1-flash-high`.
The alias is a lookup name. The command template supplies the model and effort.
