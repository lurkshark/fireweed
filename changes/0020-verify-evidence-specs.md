---
already_satisfied:
    - new:approved-author-fields
    - new:approved-runtime-state
    - new:target-coverage
    - new:fail-first-evidence
bases: {}
module: workflow
---



## Motivation

Human reviews 3 and 4 require the remaining bundled specs to be split before
Milestone 2. Begin FLOW-002 with dedicated specs for existing verification gates.
Keep FLOW-002 intact until its other guarantees have dedicated IDs. Retarget
existing evals; do not add behavior or implementation.

## Interface



## Added

### new:approved-author-fields
Verify MUST reject an unapproved alteration of the active change's author fields before running a hook.

### new:approved-runtime-state
Verify MUST reject an unapproved alteration of the active change's recorded runtime state before running a hook.

### new:target-coverage
Whenever Verify runs the target module's evals it MUST require coverage of every enforced current spec.

### new:fail-first-evidence
During evals Verify MUST require genuine fail-first evidence for every enforced change ID that is not already_satisfied.

## Modified



## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Retarget the existing
metadata/state-before-hooks case to the first two IDs. Retarget the existing
coverage and per-ID fail-first cases to the last two IDs. Keep every assertion
and all existing FLOW-002 labels until the final narrowing step; one testcase
may carry multiple spec IDs. Keep 32 reported workflow cases. No interface or
implementation edits; impl may edit only this change file.
