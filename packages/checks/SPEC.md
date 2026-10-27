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
ParseJUnit MUST accept testsuite/testsuites documents, including nested suites, and read testcase name, classname, failure/error, skipped and time, mapping IDs with contracts.MatchesID. It MUST reject malformed XML and testcases with missing/unknown IDs. It MUST classify failure/error as fail and skipped as skip, with skipped cases contributing no coverage.

### CHK-002
RequireCoverage MUST cover each current spec except review/quarantine with a non-skipped referencing case. RequireFailFirst MUST find a failed/errored case for every requested enforced ID except already_satisfied; passing, skipped, empty or unrelated failures do not qualify. RequirePass MUST reject failed/errored cases referencing any non-quarantined spec.

### CHK-003
Evaluate MUST run the module eval override or global eval hook, accepting shell strings or argv with one-pass module/module_dir/report placeholders (shell values quoted) and FIREWEED_BIN set to request Bin. It MUST parse the written JUnit and retain process output/exit. Launch/timeout, missing report and exit/report disagreement MUST fail; a nonzero exited hook with mapped failures is valid report evidence.

### CHK-004
Zones MUST restrict audit to the current change, reject protected/SPEC/foreign-module/foreign-change edits in every phase, and honor outside-module worker_writable only in evals/impl. Evals MUST allow target evals, declared interfaces, new disposable stubs, other module files and the current change. Impl MUST freeze evals/interfaces, allow disposable/preserved/other module files/current change, and permit append_only paths only when new.

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

## Decisions
