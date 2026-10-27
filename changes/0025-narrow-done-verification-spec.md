---
already_satisfied:
    - FLOW-003
    - new:passing-audit
    - new:audit-questions
    - new:audit-release
bases:
    FLOW-003: 15bd8a0e
module: workflow
---



## Motivation

Finish the review-directed FLOW-003 split. FLOW-032 through FLOW-040 now retain
its evals transitions, finding filters and audit state rules. Give its remaining
audit outcomes dedicated IDs and narrow FLOW-003 to the full-verification gate.
Existing behavior and fixtures suffice; no implementation change belongs here.

## Interface



## Added

### new:passing-audit
Audit Done MUST require No findings with a complete checked-ID list before advancing to impl.

### new:audit-questions
When audit blocks a change, Done MUST include the blocking findings in its Questions section.

### new:audit-release
Audit blocking MUST release the calling worktree's held change claim.

## Modified

### FLOW-003
Done MUST require full verification to succeed before committing a phase transition.

## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Keep 32 reported workflow
cases and every assertion. Retarget the existing passing/invalid audit case and
the existing audit-bound/untestable cases. Preserve all previously assigned IDs;
keep FLOW-003 only on verification-gate assertions.

For the full-versus-quick gate, reuse existing failing dependent/boundary fixtures
in ImplFreezesEvalsAndRunsDependentsAndBoundaries: after full Verify rejects, call
Done on that same held fixture and require rejection with HEAD/phase unchanged.
Do not add scenarios. Its existing quick Verify already succeeds. The existing
FailedDoneKeepsPhaseAndHead fixture also asserts full-verification rejection.
No interface or implementation edits; impl may edit only this change file.
