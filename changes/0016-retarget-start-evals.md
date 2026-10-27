---
audit_rounds: 0
bases: {}
eval_fix:
    - FLOW-001
    - FLOW-006
    - FLOW-007
    - FLOW-008
    - FLOW-009
module: workflow
phase: impl
run: 20261027T192020-c8e969479fe05f41a6d19a48
---



## Motivation

Human review 3 directed the split to retarget existing tests rather than add
cases. Change 0015 instead separated assertions into extra Test functions and
duplicated fixtures. Correct that eval rewrite while preserving every assertion
and the coverage of the newly assigned spec IDs.

## Interface



## Added



## Modified



## Removed



## Decisions



## Audit

No findings
Checked IDs: FLOW-001, FLOW-006, FLOW-007, FLOW-008, FLOW-009

## Notes

Restore the pre-0015 test structure from
1fe462f:packages/workflow/evals/lifecycle_test.go as the reference for this eval
file only. Reading this current test revision is allowed; implementation history
is not an input. Keep all stronger assertions added by 0015, but merge them into
the existing cases and reuse the original fixture instances.

The original start case should carry FLOW-001, FLOW-006, FLOW-007 and FLOW-008 in
its testcase name. The original approval/bound case should carry FLOW-001 and
FLOW-009. Remove the redundant standalone snapshot/ID/run test functions and the
extra unapproved fixture; its P04 outcome is already asserted in the original
below-bound scenario. Restore 32 reported workflow cases with no lost assertions.
No new Test functions, t.Run cases, fixture instances, public interfaces or
implementation changes. The impl phase may edit only this change file.

## Provenance

- evals 2026-10-27T19:21:59.567362Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=118 failing=0 passing=118
- audit 2026-10-27T19:22:33.160006Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
