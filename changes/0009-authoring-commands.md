---
audit_rounds: 0
bases: {}
module: app
phase: impl
run: 20261027T012941-7f5c120abfe217b330682184
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

### APP-006
stamp <change> MUST discover the root, invoke project.Stamp and print confirmation with exit 0, accepting the existing change aliases and leaving Added handles unassigned. It MUST return exit 2 for usage/configuration errors and exit 1 for stamping rejections, with code/message diagnostics. With --json it MUST return schema 1 with string confirmation data on success, or null data and diagnostics on failure.

### APP-007
check MUST invoke project.Check and exit 0 when no error diagnostics remain, or exit 1 with all diagnostics when static checks fail. Usage/configuration errors MUST return exit 2 with code/message diagnostics. With --json it MUST return schema 1 with string summary data on success, null data on failure and the reported diagnostics.

### APP-008
show <module> [ID] MUST return the current module document or selected item, using the whole-queue project projection with --projected. With --json it MUST return the same text as string data in schema 1 with empty diagnostics and exit 0. Unknown/ambiguous module or item, invalid arguments and configuration errors MUST return exit 2 with code/message diagnostics.

## Modified



## Removed



## Decisions



## Provenance

- evals 2026-10-27T01:37:29.370628Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=78 failing=27 passing=51
- audit 2026-10-27T01:40:36.206937Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings. Checked: APP-006, APP-007, APP-008.
