---
already_satisfied:
    - CHK-002
    - new:required-fail-first
    - new:required-passing-results
bases:
    CHK-002: a4cd59f6
module: checks
---



## Motivation

Continue the previously identified checks splits under SYSTEM's one-behavior
spec shape. CHK-002 currently bundles three public APIs. Give fail-first and
passing-result checks independent IDs and retain coverage under CHK-002,
preserving the design's Evals and Implementation step requirements.

## Interface



## Added

### new:required-fail-first
RequireFailFirst MUST find a failed or errored referencing case for every requested enforced ID except already_satisfied; passing, skipped, empty or unrelated failures do not qualify.

### new:required-passing-results
RequirePass MUST reject failed or errored cases referencing any non-quarantined spec.

## Modified

### CHK-002
RequireCoverage MUST cover each current spec except review or quarantine with a non-skipped referencing case.

## Removed



## Decisions



## Notes

Retarget existing junit_test.go tests only: coverage tests retain CHK-002,
the two FailFirst tests gain the fail-first ID instead, and PassRejectsNonQuarantinedFailures
uses the passing-results ID. Preserve all assertions and all 28 reported
checks cases, with no new tests, t.Run cases, fixture instances or helpers.
JUnit failure and error elements both become Result fail in ParseJUnit under
CHK-001; these APIs use that normalized report representation. Do not invent
another result-string value or new behavior. No product code or interface
changes. Implementation may edit only this change file.
