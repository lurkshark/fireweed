---
module: project
bases:
    PRJ-014: 7e5a959c
already_satisfied:
    - PRJ-014
---

## Motivation

Authors can refer to proposed specs in Motivation and Interface. The project
drill report in commit `422411b` shows check accepting a change whose
Motivation referred to an undeclared handle, causing the APP-037 handle case
to fail.

## Interface

project.Check reports code C09 for undeclared handle references in a pending
change document.

## Added

## Modified

### PRJ-014
Check MUST report a non-warning C09 diagnostic naming any handle referenced in a pending change document that is absent from that change's Added entries. This rule includes references in Motivation and Interface.

## Removed

## Decisions

## Notes

Add a public Check fixture for each prose section with an undeclared handle
appearing only in that section. Assert a non-warning C09 diagnostic naming
the handle. Include an otherwise valid change whose prose references a handle
declared in its own Added entries and assert that Check accepts it. Keep the
existing Modified-item case that uses a handle declared by a different change.
Current Check scans the complete rendered change document.
