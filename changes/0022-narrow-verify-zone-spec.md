---
already_satisfied:
    - FLOW-002
    - new:approved-hooks
    - new:quick-dependents
    - new:quick-boundaries
    - new:verify-no-commit
bases:
    FLOW-002: a6ff549a
module: workflow
---



## Motivation

Finish the review-directed FLOW-002 split. Its coverage, fail-first, approved-state
and full-verification gates now have dedicated IDs. Preserve its remaining
independent behaviors and narrow FLOW-002 to its phase-zone check. All behavior
already works; reuse and retarget existing evals without implementation changes.

## Interface



## Added

### new:approved-hooks
Verify MUST select hooks from the held approved projected project snapshot.

### new:quick-dependents
Quick Verify MUST skip dependent-module evals.

### new:quick-boundaries
Quick Verify MUST skip boundary verification.

### new:verify-no-commit
Verify MUST leave HEAD unchanged.

## Modified

### FLOW-002
Verify MUST reject edits forbidden by the approved projected phase zones before running hooks, checking differences since the held phase base.

## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Keep 32 reported workflow
cases and every existing assertion. Retarget FLOW-002 labels on assertions that
now belong to FLOW-020 through FLOW-027 or the Added IDs; retain FLOW-002 on zone
checks. Existing evals and impl quick/full scenarios already exercise both skips
and HEAD stability. The frozen-eval scenario checks differences since an advanced
phase base.

For approved hook selection, reuse the existing impl-apply fixture with its set
metadata: author a set.hooks.eval override pointing to the existing eval script,
and point the global eval hook to a simple failing script before Start. Keep
its other set/application assertions. No new fixture or separate test case.
No interface or implementation edits; impl may edit only this change file.
