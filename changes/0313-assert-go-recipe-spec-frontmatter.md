---
eval_fix:
    - AST-020
module: assets
phase: impl
run: 20261118T021932-98ece86ba0ecf89cc59eb058
---

## Motivation

AST-020 requires the Go recipe's module template to support identity
substitutions. The current eval searches the rendered `SPEC.md` for identity
text anywhere in the file. Checking the frontmatter ties the assertion to
the metadata Fireweed reads and accepts equivalent YAML formatting.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

In `Test_AST_007_AST_020_AST_021_AST_022_AST_023_TemplatesAndCanaries`, parse
the rendered `packages/beta/SPEC.md` with `contracts.ParseDocument`. Assert
that parsing succeeds and frontmatter contains name `beta` and ID prefix
`BET`. Keep the template substitution, compilation, and canary checks.
## Provenance

- evals 2026-11-18T02:21:16.099695Z agent=codex fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=153 failing=0 passing=153
- audit 2026-11-18T02:22:01.433276Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

## Audit

No findings; checked AST-020

