---
already_satisfied:
    - PRJ-004
bases:
    PRJ-004: 93a0113f
module: project
---



## Motivation

Finish PRJ-004's authorized split under SYSTEM's spec-shape rule. PRJ-027–031
now retain Queue's ordering, payload and rejection rules. Narrow PRJ-004 to
the design's CLI change-argument resolution behavior used by Project and
Stamp, retaining all existing assertions and behavior.

## Interface



## Added



## Modified

### PRJ-004
Project and Stamp MUST resolve a change argument from its unique slug, queued basename or repository-relative queued path, rejecting missing or ambiguous targets.

## Removed



## Decisions



## Notes

Retain PRJ-004 on TargetAliasesAndAmbiguousSlug, which already asserts all three
aliases and missing/ambiguous targets through Project and Stamp. Remove
PRJ-004 from QueueOrderContentAndArchive and QueueRejectsInvalidDocuments;
their PRJ-027–030 IDs now own those assertions. Preserve PRJ-031 on the
ambiguous-slug case and every other assertion/ID. Keep all 53 project cases.
Only rename existing tests/labels; no new Test functions, t.Run cases,
fixtures, helpers, product code or interface edits. Implementation may edit
only this change file.
