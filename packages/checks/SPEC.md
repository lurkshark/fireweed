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

## Decisions
