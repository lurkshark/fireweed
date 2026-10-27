---
already_satisfied:
    - FLOW-003
    - FLOW-041
    - FLOW-042
    - FLOW-043
audit_rounds: 0
bases:
    FLOW-003: 15bd8a0e
module: workflow
phase: impl
run: 20261027T213255-c54a80dd67be868a918b1001
---



## Motivation

Finish the review-directed FLOW-003 split. FLOW-032 through FLOW-040 now retain
its evals transitions, finding filters and audit state rules. Give its remaining
audit outcomes dedicated IDs and narrow FLOW-003 to the full-verification gate.
Existing behavior and fixtures suffice; no implementation change belongs here.

## Interface



## Added

### FLOW-041
Audit Done MUST require No findings with a complete checked-ID list before advancing to impl.

### FLOW-042
When audit blocks a change, Done MUST include the blocking findings in its Questions section.

### FLOW-043
Audit blocking MUST release the calling worktree's held change claim.

## Modified

### FLOW-003
Done MUST require full verification to succeed before committing a phase transition.

## Removed



## Decisions



## Audit

Checked IDs: FLOW-003, FLOW-041, FLOW-042, FLOW-043
No findings

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

## Provenance

- evals 2026-10-27T21:35:46.286599Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=118 failing=0 passing=118
- audit 2026-10-27T21:36:20.028732Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
