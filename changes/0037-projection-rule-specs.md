---
already_satisfied:
    - new:projection-prefix
    - new:projection-settings
    - new:projection-field-validation
    - new:projection-graph-validation
    - new:projection-local-handles
bases: {}
module: project
---



## Motivation

Begin PRJ-005's authorized split under SYSTEM's one-behavior spec shape.
Preserve the design's Change files projection and per-change handle rules,
giving selection and metadata validation independent IDs before narrowing
PRJ-005 in the next change. Existing public evals already assert these rules.

## Interface



## Added

### new:projection-prefix
Project MUST apply the queue through the named change inclusively, or the whole queue when through is empty.

### new:projection-settings
Project MUST apply set values for depends_on, interface, evals, disposable, preserve, append_only, boundaries_opt_out and hooks to the projected module.

### new:projection-field-validation
Project MUST reject identity or unknown fields in set.

### new:projection-graph-validation
Project MUST reject an invalid projected module graph.

### new:projection-local-handles
Project MUST keep Added handles local to their declaring change, allowing different changes to reuse handle text.

## Modified



## Removed



## Decisions



## Notes

Retarget existing tests, retaining PRJ-005 until the following narrowing.
ProjectEditsAndLocalHandles gains the queue-prefix and local-handle IDs;
TargetAliasesAndAmbiguousSlug also asserts the prefix cut-off. SetFieldsAndRejections
gains settings and validation IDs, with existing identity_name/identity_prefix/
unknown_field rows labeled for field validation and invalid_graph labeled for
graph validation. Preserve every assertion and every other ID. Keep 53 project
cases with no new Test functions, t.Run cases, fixture instances, helpers,
product code or interface changes. Implementation may edit only this change file.
