---
already_satisfied:
    - new:motivation
    - new:normative-text
    - new:spec-tags
    - new:duplicate-spec-ids
    - new:foreign-spec-ids
bases: {}
module: project
---



## Motivation

SYSTEM's spec-shape rule requires one observable behavior per item. Begin the
approved PRJ-007 split by giving its document validation rules separate IDs,
preserving the design's Specs and Change files obligations. Existing behavior
and assertions already satisfy these rules; keep PRJ-007 until the remaining
rules have IDs in the following change.

## Interface



## Added

### new:motivation
Check MUST reject a pending change without a nonempty Motivation section.

### new:normative-text
Check MUST reject a module spec or Added/Modified spec whose body contains no normative keyword.

### new:spec-tags
Check MUST reject spec tags other than review or quarantine with a nonempty reason.

### new:duplicate-spec-ids
Check MUST reject duplicate spec IDs within a module's Specs section.

### new:foreign-spec-ids
Check MUST reject module spec IDs and change edit IDs whose prefix belongs to another module.

## Modified



## Removed



## Decisions



## Notes

Retarget the existing CheckValidAndDiagnostics table cases to the corresponding
Added IDs, retaining PRJ-007 and all existing assertions, including its current
C01 assertion until that spec is narrowed in the next change. No new Test
functions, t.Run cases or fixture instances. Keep 53 reported project cases.
The existing missing_motivation, nonnormative_module/nonnormative_change,
invalid_review_tag/invalid_quarantine_tag, duplicate_module_id and
foreign_module_id/foreign_change_id cases provide the assertions. Only adjust
case names/ID labels as needed; retain every other test and behavior assertion.
No code, helper, fixture or interface changes are needed.
Implementation may edit only this change file.
