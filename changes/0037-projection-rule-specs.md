---
already_satisfied:
    - PRJ-032
    - PRJ-033
    - PRJ-034
    - PRJ-035
    - PRJ-036
audit_rounds: 0
bases: {}
module: project
phase: audit
run: 20261028T051431-293cc6124adbf0841c2a126e
---



## Motivation

Begin PRJ-005's authorized split under SYSTEM's one-behavior spec shape.
Preserve the design's Change files projection and per-change handle rules,
giving selection and metadata validation independent IDs before narrowing
PRJ-005 in the next change. Existing public evals already assert these rules.

## Interface



## Added

### PRJ-032
Project MUST apply the queue through the named change inclusively, or the whole queue when through is empty.

### PRJ-033
Project MUST apply set values for depends_on, interface, evals, disposable, preserve, append_only, boundaries_opt_out and hooks to the projected module.

### PRJ-034
Project MUST reject identity or unknown fields in set.

### PRJ-035
Project MUST reject an invalid projected module graph.

### PRJ-036
Project MUST keep Added handles local to their declaring change, allowing different changes to reuse handle text.

## Modified



## Removed



## Decisions



## Notes

Retarget existing tests, retaining PRJ-005 until the following narrowing.
ProjectEditsAndLocalHandles gains the queue-prefix and local-handle IDs;
TargetAliasesAndAmbiguousSlug also asserts the prefix cut-off. SetFieldsAndRejections
gains settings and validation IDs, with existing identity_name/identity_prefix/
unknown_field rows labeled for field validation and invalid_graph labeled for
graph validation. Preserve every assertion and every other ID. Keep 53 project
cases with no new Test functions, t.Run cases, fixture instances, helpers,
product code or interface changes. Implementation may edit only this change file.

## Provenance

- evals 2026-10-28T05:15:06.447199Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=199 failing=0 passing=199
