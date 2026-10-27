---
already_satisfied:
    - CHK-006
    - CHK-007
    - CHK-008
    - CHK-009
    - CHK-010
audit_rounds: 0
bases: {}
module: checks
phase: audit
run: 20261027T213810-39ce0814030218a923f53a6f
---



## Motivation

Begin the human-directed CHK-004 split with independent path prohibitions and
the outside-module exception. Each rule already works. Reuse existing zone
evals and keep CHK-004 until its remaining phase permissions have dedicated IDs.

## Interface



## Added

### CHK-006
Zones MUST reject edits to protected paths in every phase.

### CHK-007
Zones MUST reject edits to module SPEC.md files in every phase.

### CHK-008
Zones MUST reject edits within another module in every phase.

### CHK-009
Zones MUST reject edits to change files other than the current change in every phase.

### CHK-010
During evals and impl, Zones MUST allow edits outside modules only when the path matches worker_writable, except for the current change file.

## Modified



## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Keep 28 reported checks
cases and every assertion. Retarget the existing every-phase prohibitions test
and WorkerWritableOnlyEvalsAndImpl test. The current-change permission is
preserved in CHK-004 and will receive its own ID before narrowing. Keep CHK-004
labels until then. No code, fixture, helper or interface edits are needed.
Impl may edit only this change file.

## Provenance

- evals 2026-10-27T21:38:52.231631Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=146 failing=0 passing=146
