---
module: project
phase: impl
run: 20261114T042322-3bbd9232541dfb0b635d3ef7
---

## Motivation

R20 and the design distinguish file-format version from the binary version
pin. Main 3b3b64d accepts version: 999 during ordinary check. The owner
confirmed on 2026-11-13 that format 1 is the first format and has no predecessor;
current/previous reading and forward migration remain obligations when a
second format is introduced. Reject unsupported live formats now.

## Interface

No public signature or type changes.

## Added

### PRJ-085
CheckCompatibility MUST accept format version 1 and an omitted version as the initial format, and MUST reject any other present version value with a C01 CodedError identifying the unsupported value and supported format 1. Format acceptance MUST be independent of binary-range acceptance.

## Modified



## Removed

## Decisions

## Notes

Run after binary-version-compatibility, which introduces CheckCompatibility
and uses it from Load/Check. Exercise integer 1, omitted version, future and
zero versions, strings, null and malformed value types. Preserve migrate's
existing APP-153/155 requirement for an explicit format 1 and its no-op behavior;
do not invent format 0, rewrite files, bump the format, or claim a migration proof.
Historical LoadAt remains a raw recorded-snapshot reader. Future format work
must define an actual predecessor schema and a forward rewrite before R20's
future compatibility obligation can be exercised.
Omitted version retains the live loader's existing initial-format behavior;
init continues to write version: 1 explicitly.

Drafted at the owner's request after the R01–R29 traceability review.
No lifecycle execution, agent invocation, drill or release is authorized by
this drafting request. Workers must obtain context before their step and use
fireweed verify; only the controller may advance or commit lifecycle work.

Evals verification confirmed fail-first for PRJ-085. The current
CheckCompatibility accepts unsupported format values with or without a matching
binary pin, and live Load and Check accept `version: 999`. The recorded-format
LoadAt regression passes. The format check belongs in the implementation-writable
`packages/project/internal/fixture/compatibility.go`, shared by live Load and
Check.
## Audit

No findings; checked PRJ-085

## Provenance

- evals 2026-11-14T04:27:35.107164Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=278 failing=11 passing=267
- audit 2026-11-14T04:27:50.786491Z agent=codex fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

