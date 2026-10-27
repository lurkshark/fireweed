---
already_satisfied:
    - FLOW-020
    - FLOW-021
    - FLOW-022
    - FLOW-023
audit_pending:
    - FLOW-022
audit_rounds: 1
bases: {}
module: workflow
phase: evals
run: 20261027T195208-ee4bc74d8bf371d523e12f77
---



## Motivation

Human reviews 3 and 4 require the remaining bundled specs to be split before
Milestone 2. Begin FLOW-002 with dedicated specs for existing verification gates.
Keep FLOW-002 intact until its other guarantees have dedicated IDs. Retarget
existing evals; do not add behavior or implementation.

## Interface



## Added

### FLOW-020
Verify MUST reject an unapproved alteration of the active change's author fields before running a hook.

### FLOW-021
Verify MUST reject an unapproved alteration of the active change's recorded runtime state before running a hook.

### FLOW-022
Whenever Verify runs the target module's evals it MUST require coverage of every enforced current spec.

### FLOW-023
During evals Verify MUST require genuine fail-first evidence for every enforced change ID that is not already_satisfied.

## Modified



## Removed



## Decisions



## Notes

No new Test functions, t.Run cases or fixture instances. Retarget the existing
metadata/state-before-hooks case to the first two IDs. Retarget the existing
coverage and per-ID fail-first cases to the last two IDs. Keep every assertion
and all existing FLOW-002 labels until the final narrowing step; one testcase
may carry multiple spec IDs. Keep 32 reported workflow cases. No interface or
implementation edits; impl may edit only this change file.

## Provenance

- evals 2026-10-27T19:53:55.61447Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=118 failing=0 passing=118
- audit 2026-10-27T19:54:27.5353Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

- FLOW-022: "Whenever Verify runs the target module's evals it MUST require coverage of every enforced current spec." — missing assertion.
