---
module: app
phase: applied
run: 20261112T201656-04580206e513d285b4be2d42
---

## Motivation

The owner's attempts to learn agents-md and doctor usage returned usage errors
instead of help. Expose the preceding assets change's command pages through
discoverable, conventional help forms, without requiring a working project or
executing the requested command.

## Interface

No public Go signature changes. Extend app.Run argument handling to dispatch
help through assets.Help and assets.CommandHelp before ordinary command
execution and project loading.

## Added

### APP-156
The CLI MUST return assets.Help for help, --help and -h, and assets.CommandHelp
for help <command>, <command> help, <command> --help and <command> -h, with
exit 0 and empty stderr for supported commands. These forms MUST work outside
a project and with an invalid fireweed.yaml, without requiring the command's
ordinary arguments or options.

### APP-157
Every successful help form MUST accept --json and emit the schema-1 envelope
with string data identical to the corresponding plain output, an empty
diagnostics array and empty stderr.

### APP-158
An unknown command requested through a help form MUST return exit 2 with U02,
the rejected name and guidance to fireweed help. Extra arguments or unsupported
options in a help request MUST return exit 2 with corrective U01 guidance
showing a valid help form. JSON failures MUST use the existing schema-1 null
data and diagnostics envelope with empty stderr.

## Modified

## Removed

## Decisions

## Notes

Use the controller's audit submission format: each finding is one line,
`- ID: "exact single-line substring of the spec" — missing assertion`.
A passing audit says `No findings` and includes `Checked IDs:` listing all
changed IDs. Re-audits inspect only pending IDs while retaining earlier
clearances. This avoids the formatting rejection recovered in the preceding
change; keep audit enabled and preserve all round and attempt limits.

Run after progressive-help-content. Help forms accept only their shown tokens
plus optional --json; ordinary execution flags or positional values do not
become valid just because --help appears. Reserve the exact two-token
<command> help form as a help request, matching the owner's doctor and
agents-md examples. Other ordinary argument parsing remains unchanged.

Route help before the special init/new-module paths as well as normal dispatch.
The overview itself advertises command help via assets.Help; do not duplicate
its command documentation in app. Preserve APP-001/APP-002 information-command
behavior and APP-003 malformed-input handling outside these new forms.

Use public app evals with representative no-argument commands, required-input
commands and mutating commands, plus compact coverage that every listed
command resolves. Include the owner's agents-md and doctor examples and
--help/-h on init and new-module. Verify plain/JSON equality and failures
outside a project and with invalid config. Follow SYSTEM's existing read-only
command convention: help must not execute hooks/agents or change files, index
or refs. Do not introduce per-command instrumentation or a Cartesian product
of aliases, commands and environments.

The public app evals failed first for APP-156, APP-157 and APP-158. Existing
dispatch rejects the new help forms with U01, including unknown command help
requests that require U02. The targeted run also passed APP-001 through APP-003.
The implementation can route help in `packages/app/internal/entry/entry.go`
before the special dispatch paths and reuse `cliResponse` and `cliFailure`
from `packages/app/internal/entry/lifecycle.go` for both formats.

## Provenance
- evals 2026-11-12T20:29:48.80213Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=359 failing=27 passing=332
- audit 2026-11-12T20:30:07.42306Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-12T20:41:15.210171Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=359 failing=0 passing=359
## Audit

No findings; checked APP-156, APP-157, APP-158
