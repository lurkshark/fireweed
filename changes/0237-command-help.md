---
module: app
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

### new:help-aliases
The CLI MUST return assets.Help for help, --help and -h, and assets.CommandHelp
for help <command>, <command> help, <command> --help and <command> -h, with
exit 0 and empty stderr for supported commands. These forms MUST work outside
a project and with an invalid fireweed.yaml, without requiring the command's
ordinary arguments or options.

### new:help-json
Every successful help form MUST accept --json and emit the schema-1 envelope
with string data identical to the corresponding plain output, an empty
diagnostics array and empty stderr.

### new:help-errors
An unknown command requested through a help form MUST return exit 2 with U02,
the rejected name and guidance to fireweed help. Extra arguments or unsupported
options in a help request MUST return exit 2 with corrective U01 guidance
showing a valid help form. JSON failures MUST use the existing schema-1 null
data and diagnostics envelope with empty stderr.

## Modified

## Removed

## Decisions

## Notes

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

