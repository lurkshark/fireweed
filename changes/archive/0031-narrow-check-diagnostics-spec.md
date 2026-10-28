---
already_satisfied:
    - PRJ-007
    - PRJ-013
    - PRJ-014
    - PRJ-015
    - PRJ-016
audit_pending:
    - PRJ-007
    - PRJ-015
audit_rounds: 1
bases:
    PRJ-007: f92e1669
module: project
phase: applied
run: 20261028T034109-159648d6b99fd8c3dd90ba77
---



## Motivation

Complete the approved PRJ-007 split under SYSTEM's spec-shape rule. PRJ-008–012
retain its document rules; give the remaining change validation rules their own
IDs and narrow PRJ-007 to diagnostic results. Remove its blanket C01 allocation
so D006 can later assign stable cause-specific codes through audited eval_fix
and normal changes. Preserve the design's Change files validation obligations;
all current behavior already satisfies these specs.

## Interface



## Added

### PRJ-013
Check MUST reject a change that uses the same edit identity more than once across Added, Modified and Removed.

### PRJ-014
Check MUST reject a pending change that references a handle absent from that change's Added entries.

### PRJ-015
Check MUST reject Modified or Removed entries whose bases hash is missing or differs from the preceding projected spec's canonical hash.

### PRJ-016
Check MUST reject Modified or Removed IDs absent from the projection preceding the change.

## Modified

### PRJ-007
Check MUST return no error diagnostics for valid module specs and changes. Each failure diagnostic MUST have a nonempty code and message.

## Removed



## Decisions



## Notes

Reuse the existing CheckValidAndDiagnostics table. Retarget duplicate_edit_identity,
undeclared_handle, missing_base/stale_base and absent_modified/absent_removed to
the corresponding Added IDs. Retain PRJ-007 on the valid/error diagnostic
contract and retain every PRJ-008–012 label and all assertions from 0030.
No new Test functions, t.Run cases or fixture instances; keep 53 project cases.
Keep the existing C01 assertions for now: their correction belongs to a later
eval_fix, and assigned D006 codes belong to later normal changes. Do not change
product code, helpers, fixtures or interfaces.
Implementation may edit only this change file.

## Provenance

- evals 2026-10-28T03:41:54.076842Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=199 failing=0 passing=199
- audit 2026-10-28T03:42:35.060369Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- evals 2026-10-28T03:43:34.735574Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=199 failing=0 passing=199
- audit 2026-10-28T03:44:03.610531Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-28T03:44:35.924916Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=199 failing=0 passing=199

## Audit

No findings. Checked PRJ-007 and PRJ-015.
