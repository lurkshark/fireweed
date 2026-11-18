---
eval_fix:
    - CHK-038
module: checks
phase: audit
run: 20261118T040518-2742408c54da06a6e657502c
---

## Motivation

CHK-038 distinguishes a nonzero hook exit with valid failing test evidence
from an exit status that disagrees with the JUnit report. The eval pins
"disagrees with JUnit failures" for the disagreement cases. Callers can
observe the distinction through the diagnostic code and returned report.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Keep all three fixtures in `Test_CHK_038_ExitAndReportAgreement`. Assert C11
through `contracts.ErrorCode` for a zero exit with a mapped failure and for a
nonzero exit with only a passing case. Assert the reported hook exit and the
parsed testcase identity and outcome in each case. For a nonzero exit with
a mapped failure, require a nil error and the same report evidence.
## Provenance

- evals 2026-11-18T04:14:40.776105Z agent=codex fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=141 failing=0 passing=141

