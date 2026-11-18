---
eval_fix:
    - CHK-048
module: checks
phase: impl
run: 20261118T031317-af5b1bc6d2cb4e8c913eb3e9
---

## Motivation

The November 17 contracts drill returned C05 for a duplicate module item
and identified the offending ID, but the checks eval rejected the wording
"duplicate item heading". CHK-048 distinguishes malformed syntax from
duplicate identities by diagnostic code. The eval currently also requires
specific wording for both outcomes.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Keep the malformed-heading and duplicate-item fixtures in
`Test_CHK_048_EvaluateModuleItemCauses`. Assert C03 and C05 respectively
through `contracts.ErrorCode`, including that the code is present.
CON-017 covers the duplicate's structured subject at the parser boundary.
## Provenance

- evals 2026-11-18T03:22:32.52775Z agent=codex fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=138 failing=0 passing=138
- audit 2026-11-18T03:23:19.394917Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

## Audit

No findings; checked CHK-048

