---
already_satisfied:
    - CHK-004
    - CHK-021
    - CHK-022
    - CHK-023
    - CHK-024
audit_rounds: 0
bases:
    CHK-004: 5783c034
module: checks
phase: applied
run: 20261027T214522-7082d82fbdbfff2790673ab4
---



## Motivation

Finish the human-directed CHK-004 split. CHK-006 through CHK-020 now retain its
path prohibitions and phase-specific zone permissions. Give its remaining rules
separate IDs and narrow CHK-004 to audit's current-file restriction. All behavior
already works and all assertions already exist; implementation stays unchanged.

## Interface



## Added

### CHK-021
During evals and impl, Zones MUST allow edits to the current change file.

### CHK-022
During impl, Zones MUST allow only newly added files in the target module's append_only zone.

### CHK-023
During evals, Zones MUST reject edits in the target module's append_only zone.

### CHK-024
Zones MUST check both the source and destination of a rename against phase permissions.

## Modified

### CHK-004
During audit, Zones MUST allow edits only to the current change file.

## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Keep 28 reported checks
cases and every assertion. Keep CHK-004 on AuditOnlyCurrentChange; remove it from
unrelated cases now covered by dedicated IDs. Retarget the existing current-file
assertions in EvalsAllowedZonesAndNewStubs and ImplFreezesEvalsAndInterfaces,
ImplAppendOnlyNewVersusExisting, and RenamesCheckBothPaths to the Added IDs.
Keep the checks for both directions of foreign-module and append-only renames.
No code, fixture, helper or interface edits are needed.
Impl may edit only this change file.

## Provenance

- evals 2026-10-27T21:46:00.054279Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=146 failing=0 passing=146
- audit 2026-10-27T21:46:27.146734Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-27T21:46:56.746765Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=146 failing=0 passing=146

## Audit

No findings. Checked IDs: CHK-004, CHK-021, CHK-022, CHK-023, CHK-024.
