---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
    - checks
    - workflow
    - runner
id_prefix: APP
name: app
---

# app

## Specs

### APP-001
The CLI MUST run help, version, explain <code> and playbook <name> successfully
outside a project, returning the corresponding assets text on stdout with exit
0 and empty stderr. These commands MUST also work in a directory containing an
invalid fireweed.yaml.

### APP-002
Each information command MUST accept --json and emit one JSON object with schema
1, string data equal to its plain text result, and an empty diagnostics array.
It MUST return exit 0 and leave stderr empty.

### APP-003
Unknown commands/options, missing or extra arguments, and unknown lookup names
MUST return exit 2 with a nonempty corrective diagnostic. Commands listed by help
but still under construction MUST explicitly report that they are unimplemented
and return nonzero. Failed --json requests MUST use the schema-1 envelope with
null data and a diagnostic containing code and message.

### APP-004
The modules command MUST find the nearest ancestor fireweed.yaml from the working directory and report the loaded, validated module graph. With --json it MUST return schema 1, an array data containing each module's name, id_prefix, dir, depends_on and resolved zones, and an empty diagnostics array. Success and missing/invalid project configuration MUST return exits 0 and 2 respectively, with a corrective C01 diagnostic on failure.

### APP-005
The new-change <module> <slug> command MUST use the discovered project root to create the project API's queued skeleton and print its repository-relative path with exit 0. Missing/extra arguments, invalid module/slug and duplicate slugs MUST return exit 2 with a corrective diagnostic, preserving existing files. --json MUST use schema 1 with string path data on success, or null data and code/message diagnostics on failure.

### APP-006
stamp <change> MUST discover the root, invoke project.Stamp and print confirmation with exit 0, accepting the existing change aliases and leaving Added handles unassigned. It MUST return exit 2 for usage/configuration errors and exit 1 for stamping rejections, with code/message diagnostics. With --json it MUST return schema 1 with string confirmation data on success, or null data and diagnostics on failure.

### APP-007
check MUST invoke project.Check and exit 0 when no error diagnostics remain, or exit 1 with all diagnostics when static checks fail. Usage/configuration errors MUST return exit 2 with code/message diagnostics. With --json it MUST return schema 1 with string summary data on success, null data on failure and the reported diagnostics.

### APP-008
show <module> [ID] MUST return the current module document or selected item, using the whole-queue project projection with --projected. With --json it MUST return the same text as string data in schema 1 with empty diagnostics and exit 0. Unknown/ambiguous module or item, invalid arguments and configuration errors MUST return exit 2 with code/message diagnostics.

## Decisions
