---
audit_rounds: 0
bases: {}
module: checks
phase: applied
run: 20261027T062042-cf45c8a72e4c0031e0c14be7
set:
    interface:
        - api.go
        - model/**
---



## Motivation

The design's Evals/Implementation steps and Language support require JUnit
coverage, genuine fail-first evidence, approved hook execution and phase zones.
Connect these minimum checks to the manual slice next; doctor, history, budgets
and reporting beyond these results remain later work.

## Interface

Expose retained data-only Case (Name, Classname, IDs, Result, Seconds), Report
(Cases, Seconds, HookExit, Stdout, Stderr), HookRequest (Snapshot, Module, Bin,
Timeout), BoundaryRequest (Snapshot, Bin, Timeout) and ZoneRequest (Snapshot,
Module, Phase, ChangePath, Base, InterfaceDeclared, Changes). Use project.Snapshot,
project.Diagnostic, gitstore.Change and contracts.Item through public imports;
strings/slices/time.Duration are ordinary Go data. Case Result is pass/fail/skip.
Thin API: ParseJUnit(data []byte,specs []contracts.Item) (Report,error),
RequireCoverage(report Report,specs []contracts.Item) error,
RequireFailFirst(report Report,ids,already []string) error,
RequirePass(report Report,specs []contracts.Item) error,
Evaluate(ctx context.Context,request HookRequest) (Report,error),
Boundaries(ctx context.Context,request BoundaryRequest) ([]project.Diagnostic,error),
Zones(request ZoneRequest) []project.Diagnostic. Behavior belongs under internal/.

## Added

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

## Modified



## Removed



## Decisions



## Notes

Evals-phase quick verify: 28 compiling cases; 26 fail first and 2 pass.
CHK-001 4/4 fail, CHK-002 3/5 fail, CHK-003 6/6 fail,
CHK-004 7/7 fail, CHK-005 6/6 fail. The only diagnostic is B03
(589 eval lines / 19 disposable stub lines), a warning.

## Provenance

- evals 2026-10-27T06:29:48.598308Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=106 failing=25 passing=81
- audit 2026-10-27T06:31:23.289119Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- impl 2026-10-27T06:38:27.010488Z agent=manual fireweed=0.1.0 playbook=07554a3249d8fca0ba70df7504135f297491465caa63458734a4b8351734d0e5 tests=106 failing=0 passing=106

## Audit

No findings. Checked CHK-001, CHK-002, CHK-003, CHK-004, CHK-005.
