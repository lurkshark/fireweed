---
already_satisfied:
    - PRJ-001
    - PRJ-017
    - PRJ-018
audit_rounds: 0
bases:
    PRJ-001: 47df5144
module: project
phase: audit
run: 20261028T035113-10f721a50c6546f26eb894c1
---



## Motivation

The approved project spec-split slice applies SYSTEM's one-behavior spec shape
to PRJ-001. Preserve the design's Modules and zone map defaults and
Configuration requirements while separating discovery, zone default resolution
and document rejection. Existing public evals already assert these behaviors.

## Interface



## Added

### PRJ-017
Load MUST inherit an omitted module zone field from module_defaults while retaining an explicitly empty module zone list as empty.

### PRJ-018
Load MUST return an error for missing or malformed config or module documents.

## Modified

### PRJ-001
Load MUST read fireweed.yaml and module SPEC.md files only in configured modules globs, returning module names, prefixes, dependencies, documents and zones.

## Removed



## Decisions



## Notes

Retarget existing tests only; retain every assertion and 53 reported project
cases. LoadConfiguredModulesAndDefaults retains PRJ-001 and gains the defaults
ID. LoadInvalidDocuments and its existing four cases use the invalid-documents
ID rather than PRJ-001. No new Test functions, t.Run cases, fixture instances,
helpers, product code or interface edits are needed. Implementation may edit
only this change file. Keep all other tests and IDs unchanged.

## Provenance

- evals 2026-10-28T03:51:58.868099Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=199 failing=0 passing=199
