---
already_satisfied:
    - FLOW-010
    - FLOW-011
    - FLOW-012
    - FLOW-013
audit_rounds: 0
bases: {}
module: workflow
phase: impl
run: 20261027T192528-b8ecea655905b915695c36dc
---



## Motivation

Continue human review 3's FLOW-001 split by naming four existing start
preconditions or state properties independently. Keep the original until every
remaining guarantee has a dedicated ID. This is spec/eval retargeting only.

## Interface



## Added

### FLOW-010
Start MUST reject a worktree containing uncommitted changes.

### FLOW-011
Start MUST reject a requested change when an earlier unblocked queued change exists.

### FLOW-012
Start MUST return the canonical absolute worktree root in State.Snapshot.Root.

### FLOW-013
Start MUST set phase to evals when starting an unstarted spec change.

## Modified



## Removed



## Decisions



## Notes

Do not create Test functions, t.Run cases or fixture instances. Retarget the
existing clean/head test to the first two IDs and retain FLOW-001 while it still
covers the original. Add the root/phase IDs to the existing combined start test.
For its existing start call, root+"/." exercises a noncanonical input spelling
without a new fixture; keep the existing canonical-root and phase assertions.
Do not change implementation, interfaces or test scenario count (32 own cases).
The implementation phase may edit only this change file.

## Provenance

- evals 2026-10-27T19:27:06.115126Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=118 failing=0 passing=118
- audit 2026-10-27T19:27:26.428587Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

Checked IDs: FLOW-010, FLOW-011, FLOW-012, FLOW-013
No findings
