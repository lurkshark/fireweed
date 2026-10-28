---
already_satisfied:
    - PRJ-005
bases:
    PRJ-005: 5858aeac
module: project
---



## Motivation

Finish PRJ-005's authorized split under SYSTEM's spec-shape rule. PRJ-032–036
now retain projection selection, settings, validation and local handle scope.
Narrow PRJ-005 to the design's Change files spec-edit application behavior,
retaining all existing behavior assertions.

## Interface



## Added



## Modified

### PRJ-005
Project MUST preserve complete Added and Modified item text and remove Removed items when applying pending spec edits.

## Removed



## Decisions



## Notes

Keep PRJ-005 on ProjectEditsAndLocalHandles, which already checks Added and
Modified text, tags and removal. Remove PRJ-005 from SetFieldsAndRejections;
PRJ-033–035 retain those assertions. Preserve every other ID and assertion,
including PRJ-032/036. Rename existing tests only, keeping all 53 project cases
and adding no functions, t.Run cases, fixture instances, helpers, product code
or interface changes. Implementation may edit only this change file.
