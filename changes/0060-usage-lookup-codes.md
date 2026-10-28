---
bases:
    APP-003: c93927b8
    APP-013: dd1c50dd
    APP-014: 4a097bdf
module: app
---



## Motivation

The design requires stable diagnostic codes. D006 replaces legacy CLI_* labels
with U01 for invalid arguments/options and U02 for unknown lookup names. Existing
CLI usage, lookup and explicit unsupported-command rejections still emit legacy
labels. Correct those paths without implementing any unsupported command.

## Interface

No public interface change.

## Added



## Modified

### APP-003
Unknown commands or options and missing or extra arguments MUST return exit 2 with a nonempty corrective U01 diagnostic.

### APP-013
Unknown explain codes or playbook names MUST return exit 2 with a nonempty corrective U02 diagnostic.

### APP-014
Commands listed by help but still under construction MUST explicitly report that they are unimplemented with diagnostic U01 and return nonzero.

## Removed



## Decisions



## Notes

Reuse InvalidRequests, UnimplementedCommands and existing usage-error fixtures.
Add exact U01/U02 assertions for these existing plain and JSON cases while
retaining status, cause/argument, corrective-action, envelope and unimplemented
notice checks. No fixture/helper infrastructure or new scenarios are needed.
Reporting labels may change for correct ID attribution; retain all coverage.
Changed IDs must genuinely fail before implementation. Migrate existing
CLI_USAGE, CLI_OPTION and CLI_COMMAND paths to U01, including lifecycle arity
and misplaced-flag paths. Migrate CLI_LOOKUP to U02 and CLI_UNIMPLEMENTED to U01.
Leave CLI_AUTHOR/STAMP/SHOW/COVERAGE/CONFIG/LIFECYCLE and configuration causes for
subsequent scoped changes. Preserve every existing rejection and success path,
message, exit and JSON shape. Do not implement unsupported commands or add
speculative errors. Controller owns specs, commits and archive.
