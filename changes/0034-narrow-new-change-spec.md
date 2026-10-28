---
already_satisfied:
    - PRJ-003
    - new:change-numbering
    - new:change-slug-collisions
    - new:change-module-lookup
    - new:change-slug-validation
bases:
    PRJ-003: c0a8d930
module: project
---



## Motivation

Continue the owner-authorized project spec split under SYSTEM's spec-shape
rule. Retain the design's Change files ordering and authoring behavior while
separating PRJ-003's queued skeleton from numbering and argument/collision
checks. Every rule is already implemented and asserted by the public evals.

## Interface



## Added

### new:change-numbering
NewChange MUST use the next number after the largest pending or archived change number.

### new:change-slug-collisions
NewChange MUST reject a slug already present in the pending queue or archive without overwriting existing files.

### new:change-module-lookup
NewChange MUST reject an unknown target module.

### new:change-slug-validation
NewChange MUST reject an empty slug or a slug containing characters other than lowercase letters, digits and hyphens.

## Modified

### PRJ-003
NewChange MUST create a queued Markdown skeleton for an existing module with Motivation, Interface, Added, Modified, Removed and Decisions sections, containing no assigned spec IDs.

## Removed



## Decisions



## Notes

Retarget existing evals only, preserving every assertion and all 53 reported
project cases. NewChangeSkeletonAndNumbering keeps PRJ-003 and gains the
numbering ID. NewChangeRejectsDuplicatesWithoutOverwrite uses the collision
ID. NewChangeInvalidArguments covers module lookup and slug validation with
its existing unknown_module and uppercase/underscore/slash/empty_slug rows;
label these with the corresponding Added IDs, keeping both relevant IDs on
the table parent as necessary for coverage. No new Test functions, t.Run cases,
fixture instances, helpers, product code or interface changes are needed.
Implementation may edit only this change file. Preserve every unrelated ID.
