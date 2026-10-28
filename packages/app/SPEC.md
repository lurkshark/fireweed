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
Unknown commands or options and missing or extra arguments MUST return exit 2 with a nonempty corrective diagnostic.

### APP-004
modules MUST report the validated graph, using the held approved projected snapshot while this worktree owns a change.

### APP-005
new-change <module> <slug> MUST use the discovered project root to create the project API's queued skeleton and print its repository-relative path with exit 0.

### APP-006
stamp <change> MUST discover the root, invoke project.Stamp and print confirmation with exit 0, accepting the existing change aliases and leaving Added handles unassigned.

### APP-007
check MUST invoke project.Check and exit 0 when no error diagnostics remain, or exit 1 with all diagnostics when static checks fail.

### APP-008
show <module> [ID] MUST return the current module document or selected item, using the whole-queue project projection with --projected.

### APP-009
start <change> MUST discover the root, invoke workflow.Start and print first-step workflow.Context with exit 0, assigning IDs, run and phase without committing.

### APP-010
verify [change] MUST invoke workflow with the actual invoking executable as Bin, returning 0 for success, 1 for failed checks or blocked state and 2 for usage or configuration errors; plain success names the resulting phase.

### APP-011
context <change> [--full] MUST return the held workflow.Context text, using compact or full behavior as requested, with exit 0 and empty diagnostics.

### APP-012
coverage <module> MUST evaluate current specs, or the selected projection with --with-change <change>, using approved metadata while held and setting FIREWEED_BIN to the actual executable.

### APP-013
Unknown explain codes or playbook names MUST return exit 2 with a nonempty corrective diagnostic.

### APP-014
Commands listed by help but still under construction MUST explicitly report that they are unimplemented and return nonzero.

### APP-015
Failed --json requests MUST use the schema-1 envelope with null data and a diagnostic containing code and message.

### APP-016
modules MUST discover the nearest ancestor fireweed.yaml.

### APP-017
modules --json MUST return schema 1 with array data containing name, id_prefix, dir, depends_on and resolved zone arrays, plus empty diagnostics on success.

### APP-018
modules MUST return exit 0 on success or exit 2 for missing or invalid configuration, with a corrective C01 failure diagnostic.

### APP-019
new-change MUST return exit 2 with a corrective diagnostic for missing or extra arguments, invalid module or slug, or duplicate slugs, preserving existing files.

### APP-020
new-change --json MUST use schema 1 with string path data on success, or null data and code/message diagnostics on failure.

### APP-021
stamp MUST return exit 2 for usage or configuration errors and exit 1 for stamping rejections, with code/message diagnostics.

### APP-022
stamp --json MUST return schema 1 with string confirmation data on success, or null data and diagnostics on failure.

### APP-023
check MUST return exit 2 for usage or configuration errors with code/message diagnostics.

### APP-024
check --json MUST return schema 1 with string summary data on success, null data on failure and the reported diagnostics.

### APP-025
show --json MUST return identical text as schema-1 string data with empty diagnostics and exit 0 on success.

### APP-026
show MUST return exit 2 with code/message diagnostics for unknown or ambiguous module or item, invalid arguments or configuration errors.

### APP-027
start MUST preserve reported warnings and return exit 1 for lifecycle rejection or exit 2 for usage or configuration errors, with corrective code/message diagnostics.

### APP-028
start --json MUST return schema 1 with context string data on success or null data on failure.

### APP-029
done [change] MUST invoke full workflow.Done with the actual invoking executable as Bin, returning 0 for success, 1 for failed checks or blocked state and 2 for usage or configuration errors; plain success names the resulting phase.

### APP-030
verify --quick MUST use quick checks without committing.

### APP-031
verify and done --json MUST return schema 1 with result data on success or null on failure and reported code/message diagnostics.

### APP-032
context --json MUST return identical workflow.Context text as schema-1 string data.

### APP-033
context MUST return exit 2 for usage or configuration errors and exit 1 for held-state lookup failures, with corrective code/message diagnostics and null JSON data.

### APP-034
coverage MUST map results to IDs and require coverage and pass, returning exit 0 on success or exit 1 on failed checks.

### APP-035
coverage MUST return exit 2 with corrective diagnostics for usage, configuration or unknown module/change errors.

### APP-036
coverage --json MUST return schema 1 with report data on success, null on failure and code/message diagnostics.

## Decisions
