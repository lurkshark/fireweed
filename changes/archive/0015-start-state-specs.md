---
already_satisfied:
    - FLOW-006
    - FLOW-007
    - FLOW-008
    - FLOW-009
audit_rounds: 0
bases: {}
module: workflow
phase: applied
run: 20261027T191418-bdf562561a2be996d1440ce8
---



## Motivation

Human review 3 requires one observable behavior per spec. Begin splitting
FLOW-001 without dropping its existing guarantees: establish dedicated IDs for
four working behaviors now, then narrow the original through subsequent changes.
Retarget the existing evals and reuse their fixtures. No implementation change
or new feature belongs in this change.

## Interface



## Added

### FLOW-006
Start MUST include the approved projected project snapshot in its returned State.

### FLOW-007
Start MUST replace each local Added handle throughout the change with the next available spec ID for its module.

### FLOW-008
Start MUST assign a unique run identifier when starting an unstarted change.

### FLOW-009
Start MUST report P04 when neither the local base nor the shared base contains the change file.

## Modified



## Removed



## Decisions



## Notes

This is the additive stage of the FLOW-001 split. A later normal change will
narrow FLOW-001 once every remaining behavior has a dedicated spec. Keep the
existing test scenarios; rename or retarget assertions to the assigned IDs.
Implementation is already satisfied and must remain unchanged.

## Provenance

- evals 2026-10-27T19:16:49.307196Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=122 failing=0 passing=122
- audit 2026-10-27T19:17:14.764776Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-27T19:18:53.193972Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=122 failing=0 passing=122

## Audit

Checked IDs: FLOW-006, FLOW-007, FLOW-008, FLOW-009
No findings
