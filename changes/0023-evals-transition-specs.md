---
already_satisfied:
    - new:evals-to-audit
    - new:evals-to-impl
    - new:valid-finding
    - new:convergent-finding
bases: {}
module: workflow
---



## Motivation

Begin the review-directed FLOW-003 split with independent evals transitions and
audit-finding validation rules. All already work. Retarget existing evals and keep
FLOW-003 intact until its remaining audit behaviors have dedicated specs.

## Interface



## Added

### new:evals-to-audit
Done MUST commit a successfully verified evals phase and advance it to audit when audit is enabled.

### new:evals-to-impl
Done MUST commit a successfully verified evals phase and advance it to impl when audit.enabled is false.

### new:valid-finding
Audit Done MUST discard a finding whose spec ID is absent or whose quoted clause is not an exact substring of that spec.

### new:convergent-finding
On re-audit, Audit Done MUST discard findings for spec IDs that had no retained finding in the preceding audit round.

## Modified



## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Keep the 32-case total
and every assertion. Retarget the existing default-audit/disabled-audit transition
case, the exact-clause/invalid-ID case, and the existing two-round convergence
case to these Added IDs. Preserve FLOW-003 labels until final narrowing. Do not
add edge cases or alter fixtures for behavior already asserted.
No interface or implementation edits; impl may edit only this change file.
