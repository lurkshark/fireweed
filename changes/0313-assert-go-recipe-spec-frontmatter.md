---
eval_fix:
    - AST-020
module: assets
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
