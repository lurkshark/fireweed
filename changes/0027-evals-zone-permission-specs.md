---
already_satisfied:
    - CHK-011
    - CHK-012
    - CHK-013
    - CHK-014
    - CHK-015
audit_rounds: 0
bases: {}
module: checks
phase: audit
run: 20261027T214135-75fb195e1e81a18a7d747c10
---



## Motivation

Continue the human-directed CHK-004 split with one evals-phase permission rule
per zone. The current EvalsAllowedZonesAndNewStubs case already asserts them all.
Keep CHK-004 until the impl and remaining common permissions are split too.

## Interface



## Added

### CHK-011
During evals, Zones MUST allow edits in the target module's evals zone.

### CHK-012
During evals, Zones MUST allow edits in the target module's interface zone only when InterfaceDeclared is true.

### CHK-013
During evals, Zones MUST allow edits in the target module's disposable zone only for files absent from the phase base.

### CHK-014
During evals, Zones MUST allow edits in the target module's other zone.

### CHK-015
During evals, Zones MUST reject edits in the target module's preserve zone.

## Modified



## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Keep 28 reported checks
cases and every assertion. Retarget EvalsAllowedZonesAndNewStubs to the Added
IDs; it already checks the allowed paths and rejects an undeclared interface,
an existing disposable file and a preserved file. Preserve CHK-004 until final
narrowing. No code, fixture, helper or interface edits are needed.
Impl may edit only this change file.

## Provenance

- evals 2026-10-27T21:42:04.141765Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=146 failing=0 passing=146
