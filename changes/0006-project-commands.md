---
bases: {}
module: app
---



## Motivation

Milestone 1 requires promptly wiring working module operations through the NEW
CLI. Expose the already-tested project loader and authoring API so a disposable
fixture can discover modules and create its first change before lifecycle work.

## Interface

Keep the public Run entry point unchanged. Wire project.Load, ValidateGraph and
NewChange in app/internal. No root glue, dependency graph or other-module edits.

## Added

### new:modules-command
The modules command MUST find the nearest ancestor fireweed.yaml from the working directory and report the loaded, validated module graph. With --json it MUST return schema 1, an array data containing each module's name, id_prefix, dir, depends_on and resolved zones, and an empty diagnostics array. Success and missing/invalid project configuration MUST return exits 0 and 2 respectively, with a corrective C01 diagnostic on failure.

### new:author-command
The new-change <module> <slug> command MUST use the discovered project root to create the project API's queued skeleton and print its repository-relative path with exit 0. Missing/extra arguments, invalid module/slug and duplicate slugs MUST return exit 2 with a corrective diagnostic, preserving existing files. --json MUST use schema 1 with string path data on success, or null data and code/message diagnostics on failure.

## Modified



## Removed



## Decisions
