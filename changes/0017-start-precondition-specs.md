---
already_satisfied:
    - new:clean-start
    - new:queue-head
    - new:canonical-root
    - new:initial-evals-phase
bases: {}
module: workflow
---



## Motivation

Continue human review 3's FLOW-001 split by naming four existing start
preconditions or state properties independently. Keep the original until every
remaining guarantee has a dedicated ID. This is spec/eval retargeting only.

## Interface



## Added

### new:clean-start
Start MUST reject a worktree containing uncommitted changes.

### new:queue-head
Start MUST reject a requested change when an earlier unblocked queued change exists.

### new:canonical-root
Start MUST return the canonical absolute worktree root in State.Snapshot.Root.

### new:initial-evals-phase
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
