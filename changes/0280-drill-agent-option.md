---
bases:
    APP-053: 9d4f8010
    APP-054: 6a5a89bc
module: app
---

## Motivation

People running drills need to choose a configured harness and model at
invocation time. Expose the alias supported by the preceding runner change
and have init generate the named-command configuration used by both roles.

## Interface

`fireweed drill <module> [--agent <alias>] [--json]`

## Added

### new:drill-agent-option
drill --agent <alias> MUST pass the supplied alias to runner.Drill as AgentAlias. A missing, empty or repeated --agent value MUST return exit 2 and corrective U01 before running hooks or agents.

### new:drill-agent-lookup-error
An unknown drill agent alias MUST return exit 2 with U02 naming the alias and pointing to the agent.commands mapping in the project configuration. In JSON mode, drill MUST emit this diagnostic in a schema 1 envelope with null data and empty stderr.

## Modified

### APP-053
init MUST store the worker command in agent.commands and set agent.default to its alias. It MUST keep codex and claude presets as argv lists and preserve a custom template containing {prompt} or {prompt_file} as a shell string.

### APP-054
init MUST enable audit by default. When --audit-agent is supplied, it MUST accept the same presets and custom forms as --agent and set audit.agent to the alias of that command in agent.commands. Otherwise it MUST omit the audit selector so audit inherits the worker default.

## Removed

## Decisions

## Notes

Update init evals to assert the named configuration. Convert the remaining
app fixtures to the final format and remove the inline fixture entries added
for the dependency transition. Preserve the assertions against the actual
invocations, including the approved audit configuration and agent sentinels.

After the queue has been applied, carry out this repository's configuration
cutover between runs. Rebuild the controller from clean committed main with the
alias support, then commit the alias-form configuration with the same Codex
worker and auditor commands and audit settings. Record the controller revision
using the procedure in SYSTEM.md.
