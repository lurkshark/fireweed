---
eval_fix:
    - CHK-038
module: checks
phase: applied
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

All three subtests passed in a focused run during implementation. The two
disagreement cases returned C11, and the matching nonzero exit returned a nil
error. Each report retained the hook exit and the mapped testcase.

Full `fireweed verify` passed in impl with app eval budget warnings.

## Provenance

- evals 2026-11-18T04:14:40.776105Z agent=codex fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=141 failing=0 passing=141
- audit 2026-11-18T04:15:03.433683Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-18T04:35:18.879346Z agent=codex fireweed=0.2.1 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=980 failing=0 passing=980

## Audit

No findings; checked CHK-038

