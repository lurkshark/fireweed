---
already_satisfied:
    - FLOW-001
    - FLOW-018
    - FLOW-019
audit_rounds: 0
bases:
    FLOW-001: 9eea47d3
module: workflow
phase: impl
run: 20261027T193554-6de57db51b466fbd3ba1a666
---



## Motivation

Finish the review 3 FLOW-001 migration. Its other start behaviors now have
dedicated specs from 0015, 0017 and 0018. Preserve the remaining base/snapshot
guarantees under their own IDs and narrow FLOW-001 to claim creation. All behavior
already works; retarget existing evals without implementation changes.

## Interface



## Added

### FLOW-018
Start MUST use the current HEAD as the initial phase base of its claim.

### FLOW-019
While a claim is held, Current MUST return the project snapshot approved at Start.

## Modified

### FLOW-001
Start MUST atomically create the local claim for the selected queued change.

## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Keep the 32-case total.
Use the existing combined start case for both Added IDs. Its phase-base equality
assertion already covers the first. Retain the snapshot assertions; verify the
existing Current result contains the approved projected snapshot, not a newly
read live projection. If needed, reuse that same fixture to edit the existing
change's Added body before its existing Current call and compare the returned
snapshot with the approved body. Do not add a second fixture or test case.

Keep FLOW-001 only on existing assertions that prove claim creation/atomicity;
other assertions now carry their dedicated IDs. Preserve all prior checks.
No implementation or interface edits; impl may edit only this change file.

## Provenance

- evals 2026-10-27T19:38:26.905521Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=118 failing=0 passing=118
- audit 2026-10-27T19:38:57.838374Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

Checked IDs: FLOW-001, FLOW-018, FLOW-019
No findings
