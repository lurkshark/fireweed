---
already_satisfied:
    - APP-003
    - APP-004
    - APP-005
    - APP-006
    - APP-007
    - APP-008
    - new:unknown-information-lookup
    - new:unsupported-command
    - new:failed-json-envelope
    - new:modules-discovery
    - new:modules-json
    - new:modules-result-status
    - new:new-change-errors
    - new:new-change-json
    - new:stamp-result-status
    - new:stamp-json
    - new:check-usage-config
    - new:check-json
    - new:show-json
    - new:show-errors
bases:
    APP-003: c2581669
    APP-004: cd043073
    APP-005: 5768a4f4
    APP-006: 0fe58c50
    APP-007: 9427709a
    APP-008: cb56cc70
module: app
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

### new:unknown-information-lookup
Unknown explain codes or playbook names MUST return exit 2 with a nonempty corrective diagnostic.

### new:unsupported-command
Commands listed by help but still under construction MUST explicitly report that they are unimplemented and return nonzero.

### new:failed-json-envelope
Failed --json requests MUST use the schema-1 envelope with null data and a diagnostic containing code and message.

### new:modules-discovery
modules MUST discover the nearest ancestor fireweed.yaml.

### new:modules-json
modules --json MUST return schema 1 with array data containing name, id_prefix, dir, depends_on and resolved zone arrays, plus empty diagnostics on success.

### new:modules-result-status
modules MUST return exit 0 on success or exit 2 for missing or invalid configuration, with a corrective C01 failure diagnostic.

### new:new-change-errors
new-change MUST return exit 2 with a corrective diagnostic for missing or extra arguments, invalid module or slug, or duplicate slugs, preserving existing files.

### new:new-change-json
new-change --json MUST use schema 1 with string path data on success, or null data and code/message diagnostics on failure.

### new:stamp-result-status
stamp MUST return exit 2 for usage or configuration errors and exit 1 for stamping rejections, with code/message diagnostics.

### new:stamp-json
stamp --json MUST return schema 1 with string confirmation data on success, or null data and diagnostics on failure.

### new:check-usage-config
check MUST return exit 2 for usage or configuration errors with code/message diagnostics.

### new:check-json
check --json MUST return schema 1 with string summary data on success, null data on failure and the reported diagnostics.

### new:show-json
show --json MUST return identical text as schema-1 string data with empty diagnostics and exit 0 on success.

### new:show-errors
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
