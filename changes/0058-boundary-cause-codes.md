---
bases:
    CHK-039: e2f7fef5
    CHK-040: 4afd84bb
    CHK-043: 2055c878
    CHK-044: fee76f33
    CHK-045: 80ec1642
module: checks
---



## Motivation

D006 retires B05 and assigns boundary causes X01–X04. The design requires stable
codes, and Boundaries already emits these violations, hook failures and warnings.
Migrate these existing paths without expanding boundary behavior or adding
speculative failures. Warning severity remains separate from the code prefix.

## Interface

No interface change.

## Added



## Modified

### CHK-039
Boundaries MUST return errors with diagnostic X01 for unsuppressed module-attributed violations.

### CHK-040
Boundaries MUST return a visible X03 warning containing the reason for a reasoned boundaries_opt_out.

### CHK-043
Boundaries MUST report hook launch, cancellation or timeout failures as errors with diagnostic X02.

### CHK-044
Boundaries MUST fail with diagnostic X02 for a missing hook required by a module with disposable globs and no reasoned opt-out.

### CHK-045
Boundaries MUST warn with diagnostic X04 without an error for a missing hook when no boundary enforcement is required.

## Removed



## Decisions



## Notes

Add exact assigned-code and severity assertions to existing boundaries_test.go
fixtures for the changed IDs, retaining error, attribution, reason, suppression,
marker and timeout assertions. Reuse all 28 own cases and existing scenarios;
no new fixture/helper infrastructure or instrumentation. Keep genuine fail-first
coverage for each changed ID and preserve all other coverage. A small assertion
helper is allowed only if it reduces repetition of diagnostic checks.
Implementation changes codes on existing diagnostic paths and attaches X02 to
existing hook-launch/cancellation/timeout errors; no new error condition is
authorized. Keep all error returns, warnings, normal-exit/suppression semantics
and process control behavior. B05 is no longer emitted here. Do not modify
other modules or add unused codes. Controller owns specs, commits and archive.
