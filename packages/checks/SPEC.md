---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
id_prefix: CHK
interface:
    - api.go
    - model/**
name: checks
---

# checks

## Specs

### CHK-001
ParseJUnit MUST accept testsuite or testsuites documents, including nested suites, returning testcase name, classname and time.

### CHK-002
RequireCoverage MUST cover each current spec except review or quarantine with a non-skipped referencing case.

### CHK-003
Evaluate MUST run the module eval override or global eval hook, accepting shell strings or argv with one-pass module/module_dir/report placeholders (shell values quoted) and FIREWEED_BIN set to request Bin. It MUST parse the written JUnit and retain process output/exit. Launch/timeout, missing report and exit/report disagreement MUST fail; a nonzero exited hook with mapped failures is valid report evidence.

### CHK-004
During audit, Zones MUST allow edits only to the current change file.

### CHK-005
Boundaries MUST run applicable global/module hooks, returning errors for unsuppressed module-attributed violations and visible warnings for reasoned opt-outs. It MUST suppress only exact baseline lines or opted-out module lines from normally exited hooks; launch/cancel/timeout failures remain errors. A missing hook MUST fail for unopted disposable modules and warn when no boundary enforcement is required.

### CHK-006
Zones MUST reject edits to protected paths in every phase.

### CHK-007
Zones MUST reject edits to module SPEC.md files in every phase.

### CHK-008
Zones MUST reject edits within another module in every phase.

### CHK-009
Zones MUST reject edits to change files other than the current change in every phase.

### CHK-010
During evals and impl, Zones MUST allow edits outside modules only when the path matches worker_writable, except for the current change file.

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

### CHK-016
During impl, Zones MUST reject edits in the target module's evals zone.

### CHK-017
During impl, Zones MUST reject edits in the target module's interface zone.

### CHK-018
During impl, Zones MUST allow edits in the target module's disposable zone.

### CHK-019
During impl, Zones MUST allow edits in the target module's preserve zone.

### CHK-020
During impl, Zones MUST allow edits in the target module's other zone.

### CHK-021
During evals and impl, Zones MUST allow edits to the current change file.

### CHK-022
During impl, Zones MUST allow only newly added files in the target module's append_only zone.

### CHK-023
During evals, Zones MUST reject edits in the target module's append_only zone.

### CHK-024
Zones MUST check both the source and destination of a rename against phase permissions.

### CHK-025
RequireFailFirst MUST find a failed or errored referencing case for every requested enforced ID except already_satisfied; passing, skipped, empty or unrelated failures do not qualify.

### CHK-026
RequirePass MUST reject failed or errored cases referencing any non-quarantined spec.

### CHK-027
ParseJUnit MUST reject malformed XML.

### CHK-028
ParseJUnit MUST reject testcases with missing or unknown spec IDs.

### CHK-029
ParseJUnit MUST map testcase name and classname to current spec IDs using contracts.MatchesID.

### CHK-030
ParseJUnit MUST classify failure and error elements as fail, and skipped elements as skip.

## Decisions
