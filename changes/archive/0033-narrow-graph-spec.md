---
already_satisfied:
    - PRJ-002
    - PRJ-019
    - PRJ-020
    - PRJ-021
    - PRJ-022
audit_rounds: 0
bases:
    PRJ-002: 42ff7021
module: project
phase: applied
run: 20261028T035342-54c189bac9cfd060909cf9a1
---



## Motivation

Continue the explicitly authorized project splits under SYSTEM's one-behavior
spec-shape rule. Separate PRJ-002's graph constraints from its validation result
contract, retaining the design's Modules and zone map invariants. All existing
Graph table assertions and behavior already satisfy the split.

## Interface



## Added

### PRJ-019
ValidateGraph MUST reject module names that are empty or duplicated in the graph.

### PRJ-020
ValidateGraph MUST reject module prefixes that are empty or duplicated in the graph.

### PRJ-021
ValidateGraph MUST reject a dependency that does not name a different existing module.

### PRJ-022
ValidateGraph MUST reject a graph containing a dependency cycle.

## Modified

### PRJ-002
ValidateGraph MUST return nil for a valid graph and an explanatory error identifying the rejection cause for an invalid graph.

## Removed



## Decisions



## Notes

Retarget the existing Graph table without adding cases. Keep PRJ-002 on the
validation-result/error assertions, and label existing rows with their Added
IDs: empty_name/duplicate_name, empty_prefix/duplicate_prefix,
missing_dependency/self_dependency, and cycle respectively. Preserve the
acyclic acceptance and every explanatory-error assertion. Keep the unchanged
53 reported project cases; no new fixtures, Test functions, t.Run cases or
helpers. No product code or interface changes are needed. Implementation may
edit only this change file. Retain every other test and ID.

Implementation review found the Graph rows already retargeted to PRJ-019–022,
with PRJ-002 retained on the result and explanatory-error assertions. No product
change was needed. Quick verification passed all 53 project cases.

## Provenance

- evals 2026-10-28T03:54:38.367009Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=199 failing=0 passing=199
- audit 2026-10-28T03:55:07.674195Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-28T03:55:38.296605Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=199 failing=0 passing=199

## Audit

No findings. Checked PRJ-002, PRJ-019, PRJ-020, PRJ-021, and PRJ-022.
