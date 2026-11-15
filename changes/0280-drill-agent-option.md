---
bases:
    APP-053: 9d4f8010
    APP-054: 6a5a89bc
module: app
phase: impl
run: 20261115T030854-df5790b964e017e06e81986a
---

## Motivation

People running drills need to choose a configured harness and model at
invocation time. Expose the alias supported by the preceding runner change
and have init generate the named-command configuration used by both roles.

## Interface

`fireweed drill <module> [--agent <alias>] [--json]`

## Added

### APP-166
drill --agent <alias> MUST pass the supplied alias to runner.Drill as AgentAlias. A missing, empty or repeated --agent value MUST return exit 2 and corrective U01 before running hooks or agents.

### APP-167
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

The controller's `fireweed verify drill-agent-option` passed the evals phase.
The app report recorded 394 passing cases and 9 expected failures, including
the parent and subcases for audit command forms. APP-166 and APP-167 fail
because drill rejects `--agent`. APP-053 and APP-054 fail because the generated
configuration lacks `agent.commands`, `agent.default`, and the explicit
`audit.agent` selector.

Implementation can update option parsing in `packages/app/internal/entry/entry.go`,
drill dispatch and diagnostics in `packages/app/internal/entry/lifecycle.go`, and
generated configuration in `packages/app/internal/entry/setup.go`.
## Audit

No findings; checked APP-053, APP-054, APP-166, APP-167

## Provenance

- evals 2026-11-15T03:19:28.971846Z agent=codex fireweed=0.2.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=403 failing=9 passing=394
- audit 2026-11-15T03:19:45.577594Z agent=codex fireweed=0.2.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

