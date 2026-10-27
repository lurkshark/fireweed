---
already_satisfied:
    - CHK-016
    - CHK-017
    - CHK-018
    - CHK-019
    - CHK-020
audit_rounds: 0
bases: {}
module: checks
phase: audit
run: 20261027T214319-fdb7540438b4a3a52824d47b
---



## Motivation

Continue the human-directed CHK-004 split with one impl-phase permission rule
per zone. The existing ImplFreezesEvalsAndInterfaces case already asserts all
five. Keep CHK-004 until the remaining common and append-only rules are split.

## Interface



## Added

### CHK-016
During impl, Zones MUST reject edits in the target module's evals zone.

### CHK-017
During impl, Zones MUST reject edits in the target module's interface zone.

### CHK-018
During impl, Zones MUST allow edits in the target module's disposable zone.

### CHK-019
During impl, Zones MUST allow edits in the target module's preserve zone.

### CHK-020
During impl, Zones MUST allow edits in the target module's other zone.

## Modified



## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Keep 28 reported checks
cases and every assertion. Retarget ImplFreezesEvalsAndInterfaces to the Added
IDs; both its rejected paths and all allowed zone paths are already asserted.
Keep CHK-004 until final narrowing. No code, fixture, helper or interface edits
are needed. Impl may edit only this change file.

## Provenance

- evals 2026-10-27T21:44:00.285336Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=146 failing=0 passing=146
