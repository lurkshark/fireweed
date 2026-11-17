---
already_satisfied:
    - PRJ-014
bases:
    PRJ-014: 7e5a959c
module: project
phase: audit
run: 20261117T092954-2345e47df4982a54807522c6
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
## Provenance

- evals 2026-11-17T09:31:38.291589Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=322 failing=0 passing=322

