---
already_satisfied:
    - new:checked-ids
    - new:findings-to-evals
    - new:audit-round-count
    - new:audit-round-limit
    - new:untestable-block
bases: {}
module: workflow
---



## Motivation

Continue the review-directed FLOW-003 split with independent audit state rules.
Existing evals already assert each behavior. Keep FLOW-003 until its remaining
passing-audit, Questions and claim-release behavior can be retained separately.
No new feature or implementation change belongs here.

## Interface



## Added

### new:checked-ids
Audit Done MUST require every enforced ID in the change's Added, Modified or fixes entries to appear in a No findings submission's checked-ID list.

### new:findings-to-evals
Retained testable audit findings within the round limit MUST return the change to evals.

### new:audit-round-count
An audit return to evals MUST increment and persist audit_rounds.

### new:audit-round-limit
Audit Done MUST block the change when retained findings would exceed audit.max_rounds, whose default is 2.

### new:untestable-block
Audit Done MUST block the change when a retained finding marks its quoted clause untestable.

## Modified



## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Keep 32 reported workflow
cases and every existing assertion. Retarget the existing checked-No-findings
case, the two-round/default and one-round/configured bound case, and the existing
untestable-finding case. Preserve FLOW-003 labels until final narrowing. These
cases already check phase, persistent count, block commits, Questions and claim
release; do not add scenarios or alter fixtures for asserted behavior.
No interface or implementation edits; impl may edit only this change file.
