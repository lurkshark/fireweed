---
already_satisfied:
    - FLOW-004
    - FLOW-049
    - FLOW-050
audit_rounds: 0
bases:
    FLOW-004: "103006e6"
module: workflow
phase: impl
run: 20261028T060306-e0fe15b43c2f8421a76d5286
---



## Motivation

Finish the identified FLOW-004 split under SYSTEM's one-behavior spec shape.
FLOW-044–048 retain metadata, Decisions, phase-base, claim-release and failed
verification behavior. Separate the design's static validation and archive
transaction, leaving spec application under FLOW-004. Existing implementation
already satisfies all these obligations.

## Interface



## Added

### FLOW-049
Impl Done MUST require the applied project to pass static validation before committing.

### FLOW-050
Successful impl Done MUST move the active change to the archive with phase applied and commit it together with the application.

## Modified

### FLOW-004
Successful impl Done MUST apply Added, Modified and Removed spec edits to the module's SPEC.md.

## Removed



## Decisions



## Notes

Keep FLOW-004 on ImplDoneAppliesSpecsSetAndDecisionsAndArchives, adding the
archive-commit ID. Reuse that fixture to assert the archived change and applied
SPEC.md are present together at the new HEAD, using existing public Git APIs;
retain the existing working-tree, phase and claim assertions.
FailedDoneKeepsPhaseAndHead uses static-validation ID instead of FLOW-004,
retaining FLOW-003/024/048 and its injected invalid-spec failure assertions.
Keep all other IDs and every assertion. No new Test functions, t.Run cases,
fixture instances or helpers; retain 32 workflow cases. No product code or
interface changes. Implementation may edit only this change file.

## Provenance

- evals 2026-10-28T06:06:20.196059Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=118 failing=0 passing=118
- audit 2026-10-28T06:06:59.956597Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

Checked IDs: FLOW-004, FLOW-049, FLOW-050
No findings
