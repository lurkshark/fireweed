---
module: project
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

### new:current-project-format
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
