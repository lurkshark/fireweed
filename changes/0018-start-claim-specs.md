---
already_satisfied:
    - FLOW-014
    - FLOW-015
    - FLOW-016
    - FLOW-017
audit_rounds: 0
bases: {}
module: workflow
phase: audit
run: 20261027T193119-743f44df14db52ebe4d31933
---



## Motivation

Continue the review 3 FLOW-001 split with dedicated specs for four remaining
observable claim behaviors. Existing eval scenarios already assert these rules.
Retarget them without changing implementation or adding cases.

## Interface



## Added

### FLOW-014
Start MUST leave HEAD unchanged.

### FLOW-015
Start MUST reject when the shared-base unapproved archive count reaches authoring.max_unapproved, whose default is 5.

### FLOW-016
Start MUST reject a competing worktree that attempts to claim the same change.

### FLOW-017
Current MUST resolve only a claim owned by the calling worktree.

## Modified



## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Add the no-commit and
Current IDs to the existing combined start case. Retarget the existing bound
scenario to its dedicated ID. Add the competing-owner and Current IDs to the
existing linked-worktree case. Retain FLOW-001 until the final narrowing change.
Keep all assertions and the 32-case total. No implementation or interface edits;
the impl phase may edit only this change file.

## Provenance

- evals 2026-10-27T19:32:56.319674Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=118 failing=0 passing=118
