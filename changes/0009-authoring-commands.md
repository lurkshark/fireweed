---
bases: {}
module: app
---



## Motivation

Immediately connect the working project queue/projection/stamp/check APIs to the
NEW CLI. This lets the disposable manual slice stamp and check its Added change
and inspect the projected spec before Git claims and lifecycle are implemented.

## Interface

Keep public Run unchanged. Compose public project/contracts APIs under
app/internal; add no root glue, dependency or other-module changes. Existing
information/module/new-change behavior remains supported.

## Added

### new:stamp-command
stamp <change> MUST discover the root, invoke project.Stamp and print confirmation with exit 0, accepting the existing change aliases and leaving Added handles unassigned. It MUST return exit 2 for usage/configuration errors and exit 1 for stamping rejections, with code/message diagnostics. With --json it MUST return schema 1 with string confirmation data on success, or null data and diagnostics on failure.

### new:check-command
check MUST invoke project.Check and exit 0 when no error diagnostics remain, or exit 1 with all diagnostics when static checks fail. Usage/configuration errors MUST return exit 2 with code/message diagnostics. With --json it MUST return schema 1 with string summary data on success, null data on failure and the reported diagnostics.

### new:show-command
show <module> [ID] MUST return the current module document or selected item, using the whole-queue project projection with --projected. With --json it MUST return the same text as string data in schema 1 with empty diagnostics and exit 0. Unknown/ambiguous module or item, invalid arguments and configuration errors MUST return exit 2 with code/message diagnostics.

## Modified



## Removed



## Decisions
