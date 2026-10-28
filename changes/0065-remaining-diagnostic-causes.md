---
already_satisfied:
    - new:dependency-causes
bases:
    APP-018: 809a11b0
    APP-019: 756f1953
module: app
---



## Motivation

D006 and the design require stable cause-specific diagnostics. The remaining
CLI_AUTHOR/STAMP/SHOW/COVERAGE/CONFIG/LIFECYCLE labels hide existing causes, and
project loading still labels graph, document and metadata failures C01. Complete
this existing app classification in one cohesive batch, retaining messages,
exits, JSON and rejection/success behavior. APP-018's old blanket C01 promise is
replaced by its assigned causes, with the behavior assertions retained.

## Interface

No public interface or signature changes.

## Added

### new:static-causes
Existing static failures exposed by the CLI MUST retain their D006 cause: C01 configuration, C02 graph, C03 syntax, C04 normative text, C05 identities, C06 tags, C07 change structure, C08 bases, C09 handles, C12 metadata or C13 queue identity.

### new:lookup-causes
CLI lookup rejections for unknown or ambiguous requested names MUST use U02, retaining C13 for queue identity ambiguity.

### new:operational-causes
Existing operational precondition failures without a more specific D006 cause MUST report P01.

### new:dependency-causes
For workflow or checks failures carrying an assigned code, the CLI MUST preserve that code and reported warnings.

## Modified

### APP-018
modules MUST return exit 0 on success or exit 2 for missing or invalid project inputs, with a corrective failure diagnostic: C01 for configuration, C02 for graph, C03 for document syntax or C12 for metadata errors.

### APP-019
new-change MUST return exit 2 with corrective diagnostics U01 for invalid arguments or slug, U02 for an unknown module and C13 for duplicate slugs, preserving existing files.

## Removed



## Decisions



## Notes

Reuse existing module, authoring, lookup, coverage and lifecycle error fixtures
for exact code assertions; preserve statuses, cause/correction, JSON envelope,
warning, file-preservation and success assertions and all ID coverage. Do not
build an every-command/error matrix. Add only a representative public-input case
for a named existing cause lacking a fixture. For P01, an existing new-change
output path blocked by a directory exercises its current create-file error:
queue scanning ignores directories, so this is an operational failure rather
than a duplicate change identity. Do not add a new rejection. The dependency
code forwarding contract is legitimately already satisfied by the coded library
errors; retain exact representative assertions. All other changed IDs need
genuine fail-first evidence. Use existing assertion/fixture helpers.

Migrate every remaining emitted CLI_* label on existing paths. Keep usage U01
and lookup U02; distinguish invalid slug, duplicate queue identity, missing
edit target, malformed document/item, invalid set/metadata and graph failures.
Keep existing dependency/static causes and warnings, including lifecycle P05,
P2-01, Z01 and X codes. Use P01 only for otherwise uncoded operational errors.
Do not select an unrelated project.Check diagnostic merely because it appears
first; classifications must match the actual returned cause. Minimal existing
call-site classification or matching public diagnostics is allowed, without
adding public interfaces or a general error framework. Preserve original cause
text and corrective actions; avoid duplicate identical prefixes in plain output.
All JSON/status rules remain unchanged. Do not implement unsupported commands,
add validations/speculative failures, change behavior or edit another module.
Controller owns spec application, commits and archive.
