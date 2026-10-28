---
already_satisfied:
    - APP-009
    - APP-010
    - APP-011
    - APP-012
    - new:start-result-status
    - new:start-json
    - new:done-command
    - new:verify-quick
    - new:verify-done-json
    - new:context-json
    - new:context-errors
    - new:coverage-results
    - new:coverage-input-errors
    - new:coverage-json
bases:
    APP-009: 26729c7a
    APP-010: 8d083138
    APP-011: 4e8cec43
    APP-012: b9303680
module: app
---



## Motivation

Finish the identified lifecycle CLI bundles under SYSTEM's 30-item spec-only
cap. Separate verify and done, mode/result contracts and JSON/error reporting
while retaining the design's lifecycle, context and coverage obligations.
Fourteen affected items; all behavior and public interfaces already exist.
Preserve every assertion and fixture. No diagnostic allocation, unused error
path, unimplemented command or product behavior changes are authorized.

## Interface



## Added

### new:start-result-status
start MUST preserve reported warnings and return exit 1 for lifecycle rejection or exit 2 for usage or configuration errors, with corrective code/message diagnostics.

### new:start-json
start --json MUST return schema 1 with context string data on success or null data on failure.

### new:done-command
done [change] MUST invoke full workflow.Done with the actual invoking executable as Bin, returning 0 for success, 1 for failed checks or blocked state and 2 for usage or configuration errors; plain success names the resulting phase.

### new:verify-quick
verify --quick MUST use quick checks without committing.

### new:verify-done-json
verify and done --json MUST return schema 1 with result data on success or null on failure and reported code/message diagnostics.

### new:context-json
context --json MUST return identical workflow.Context text as schema-1 string data.

### new:context-errors
context MUST return exit 2 for usage or configuration errors and exit 1 for held-state lookup failures, with corrective code/message diagnostics and null JSON data.

### new:coverage-results
coverage MUST map results to IDs and require coverage and pass, returning exit 0 on success or exit 1 on failed checks.

### new:coverage-input-errors
coverage MUST return exit 2 with corrective diagnostics for usage, configuration or unknown module/change errors.

### new:coverage-json
coverage --json MUST return schema 1 with report data on success, null on failure and code/message diagnostics.

## Modified

### APP-009
start <change> MUST discover the root, invoke workflow.Start and print first-step workflow.Context with exit 0, assigning IDs, run and phase without committing.

### APP-010
verify [change] MUST invoke workflow with the actual invoking executable as Bin, returning 0 for success, 1 for failed checks or blocked state and 2 for usage or configuration errors; plain success names the resulting phase.

### APP-011
context <change> [--full] MUST return the held workflow.Context text, using compact or full behavior as requested, with exit 0 and empty diagnostics.

### APP-012
coverage <module> MUST evaluate current specs, or the selected projection with --with-change <change>, using approved metadata while held and setting FIREWEED_BIN to the actual executable.

## Removed



## Decisions



## Notes

Retarget existing lifecycle_cli_test.go labels only. Start context/assignment
and warning tests retain APP-009 and gain status/JSON IDs as asserted. Context
compact/full/errors retains APP-011 and gains Context JSON/error IDs. Coverage
current/projection/errors retains APP-012 and gains results/input/JSON IDs.
VerifyDoneAuditAndFullChecks and VerifyRejectsFalseFailFirstAndFrozenEdits keep
APP-010 where verify is asserted and gain Done/quick/JSON IDs as relevant.
DoneReportsBlockedAudit uses Done/JSON IDs instead of APP-010. Preserve all
explicit executable-Bin, quick/no-commit, full-Done, blocked-state, result-data,
warning, error, metadata and coverage assertions. Keep every unrelated ID.
No new Test functions, t.Run cases or fixture instances; keep 86 app cases.
Do not change helpers, command invocations, source fixture strings, hook code,
binary builds, product code, interfaces or diagnostic assertions. Implementation
may edit only this change file.
