---
already_satisfied:
    - FLOW-004
    - new:apply-static-validation
    - new:archive-application-commit
bases:
    FLOW-004: "103006e6"
module: workflow
---



## Motivation

Finish the identified FLOW-004 split under SYSTEM's one-behavior spec shape.
FLOW-044–048 retain metadata, Decisions, phase-base, claim-release and failed
verification behavior. Separate the design's static validation and archive
transaction, leaving spec application under FLOW-004. Existing implementation
already satisfies all these obligations.

## Interface



## Added

### new:apply-static-validation
Impl Done MUST require the applied project to pass static validation before committing.

### new:archive-application-commit
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
