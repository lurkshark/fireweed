---
eval_fix:
    - PRJ-011
    - PRJ-068
module: project
phase: applied
run: 20261118T034323-7f7543935eaf0763fe615274
---

## Motivation

The November 17 contracts drill produced a C05 diagnostic on the correct
module SPEC.md, but the project eval rejected "duplicate item heading".
PRJ-011 specifies C05 for duplicates, and PRJ-068 specifies the affected
document's path. These fields identify the failure and where to correct it.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

For the duplicate-module-ID fixture in
`Test_PRJ_015_PRJ_016_CheckDiagnostics`, assert a non-warning C05 diagnostic
with a nonempty message and `Diagnostic.Path` equal to
`packages/alpha/SPEC.md`. Cover PRJ-011 and PRJ-068 and retain the existing
assertions for the other fixtures.
## Provenance

- evals 2026-11-18T03:44:26.950446Z agent=codex fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=322 failing=0 passing=322
- audit 2026-11-18T03:44:43.95379Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-18T04:05:15.981084Z agent=codex fireweed=0.2.1 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=1299 failing=0 passing=1299

## Audit

No findings; checked PRJ-011, PRJ-068

