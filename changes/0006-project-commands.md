---
audit_rounds: 0
bases: {}
module: app
phase: impl
run: 20261027T005400-980c8460d09504c82dc3fc70
---



## Motivation

Milestone 1 requires promptly wiring working module operations through the NEW
CLI. Expose the already-tested project loader and authoring API so a disposable
fixture can discover modules and create its first change before lifecycle work.

## Interface

Keep the public Run entry point unchanged. Wire project.Load, ValidateGraph and
NewChange in app/internal. No root glue, dependency graph or other-module edits.

## Added

### APP-004
The modules command MUST find the nearest ancestor fireweed.yaml from the working directory and report the loaded, validated module graph. With --json it MUST return schema 1, an array data containing each module's name, id_prefix, dir, depends_on and resolved zones, and an empty diagnostics array. Success and missing/invalid project configuration MUST return exits 0 and 2 respectively, with a corrective C01 diagnostic on failure.

### APP-005
The new-change <module> <slug> command MUST use the discovered project root to create the project API's queued skeleton and print its repository-relative path with exit 0. Missing/extra arguments, invalid module/slug and duplicate slugs MUST return exit 2 with a corrective diagnostic, preserving existing files. --json MUST use schema 1 with string path data on success, or null data and code/message diagnostics on failure.

## Modified



## Removed



## Decisions



## Notes

App evals exercise the built NEW executable from nested temporary projects. APP-004 covers graph output, nearest-ancestor discovery and C01 configuration failures. APP-005 covers queued path/content, both formats, argument/module/slug errors and pending/archived duplicate preservation within the 15-case limit. The existing single PATH recorder now includes modules.

## Provenance

- evals 2026-10-27T00:57:53.605156Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=51 failing=12 passing=39
- audit 2026-10-27T00:59:19.484511Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings.

Checked: APP-004, APP-005.
