---
already_satisfied:
    - APP-003
    - APP-004
    - APP-005
    - APP-006
    - APP-007
    - APP-008
    - APP-013
    - APP-014
    - APP-015
    - APP-016
    - APP-017
    - APP-018
    - APP-019
    - APP-020
    - APP-021
    - APP-022
    - APP-023
    - APP-024
    - APP-025
    - APP-026
audit_rounds: 0
bases:
    APP-003: c2581669
    APP-004: cd043073
    APP-005: 5768a4f4
    APP-006: 0fe58c50
    APP-007: 9427709a
    APP-008: cb56cc70
module: app
phase: audit
run: 20261028T064628-e207a5b15498f58dcd06a2b2
---



## Motivation

Batch the identified standalone/authoring CLI bundles under SYSTEM's 30-item
spec-only cap. Separate command behavior from existing lookup, rejection and
JSON result contracts while preserving the design's CLI requirements. Twenty
affected items; no product behavior, interface, assertion or fixture changes.
Keep current diagnostic allocations, including modules' C01 promise, for later
normal D006 migration on existing paths. Do not implement unsupported commands.

## Interface



## Added

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

## Modified

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

## Removed



## Decisions



## Notes

Retarget existing app_test.go, project_commands_test.go and authoring_commands_test.go
labels only, preserving every assertion, literal test input and all 86 app cases.
InvalidRequests retains APP-003 and gains unknown-lookup ID on its existing
lookup cases; UnimplementedCommands and FailedJSON use their Added IDs instead
of APP-003. Keep the PATH convention probe unchanged. Modules tests retain
APP-004 only for graph assertions and gain discovery/JSON/status IDs as relevant;
ModulesConfigurationErrors uses result-status. NewChange, Stamp, Check and Show
success/error/JSON tests gain their matching Added IDs, retaining old IDs only
where narrowed behavior is still asserted. Preserve APP-001/002 and APP-009–012.
Do not mass-label all JSON tests with the common failure ID: its existing
FailedJSON assertion is retained, and command-specific format IDs own their
assertions. Keep helpers, fixture paths/strings, binary builds, hooks, code,
interfaces and diagnostic assertions unchanged. No new Test functions, t.Run
cases or fixture instances. Implementation may edit only this change file.

## Provenance

- evals 2026-10-28T06:50:09.147566Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=86 failing=0 passing=86
