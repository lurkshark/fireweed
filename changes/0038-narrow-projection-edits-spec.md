---
already_satisfied:
    - PRJ-005
audit_rounds: 0
bases:
    PRJ-005: 5858aeac
module: project
phase: audit
run: 20261028T051636-4780cbd89dbfdcec5b3cd2b6
---



## Motivation

Finish PRJ-005's authorized split under SYSTEM's spec-shape rule. PRJ-032–036
now retain projection selection, settings, validation and local handle scope.
Narrow PRJ-005 to the design's Change files spec-edit application behavior,
retaining all existing behavior assertions.

## Interface



## Added



## Modified

### PRJ-005
Project MUST preserve complete Added and Modified item text and remove Removed items when applying pending spec edits.

## Removed



## Decisions



## Notes

Keep PRJ-005 on ProjectEditsAndLocalHandles, which already checks Added and
Modified text, tags and removal. Remove PRJ-005 from SetFieldsAndRejections;
PRJ-033–035 retain those assertions. Preserve every other ID and assertion,
including PRJ-032/036. Rename existing tests only, keeping all 53 project cases
and adding no functions, t.Run cases, fixture instances, helpers, product code
or interface changes. Implementation may edit only this change file.

## Provenance

- evals 2026-10-28T05:17:11.492029Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=199 failing=0 passing=199
