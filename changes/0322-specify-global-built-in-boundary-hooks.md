---
already_satisfied:
    - AST-073
module: assets
phase: impl
run: 20261118T215415-010a84df988d39e682800cf2
---

## Motivation

The November 18 assets drill rebuilt a Go boundaries hook that required a module
argument. The recipe eval supplied a module name when expanding the command, so own
evals passed and Doctor and CI failed. The design's language-support table specifies
a project-wide hook. TypeScript and Python recipe evals use the same
module-supplying fixture pattern.
[Drill reports](https://github.com/lurkshark/fireweed/pull/25)

## Interface

Specify the global hook contract on the Go, TypeScript with pnpm, and Python with uv
recipes returned by BuiltinRecipe.

## Added

### AST-073
Each built-in recipe's global boundaries command MUST run without a module selection
and check all modules returned by FIREWEED_BIN modules --json. It MUST attribute
each detected violation to its source module.

## Modified

## Removed

## Decisions

## Notes

Run the returned boundaries command with the same empty module and module directory
values as checks.Boundaries uses for a global hook. Use a clean fixture with two
graph modules, then place an undeclared import only in the second module and require
a violation line beginning with that module's name. Look for that line in the hook
output without assuming a summary's placement. Reuse the existing recipe fixtures
and offline tool setup. The current three built-in hooks already scan the full graph.

The recipe fixtures now invoke global boundaries hooks with empty module values.
`fireweed verify specify-global-built-in-boundary-hooks` passed in the evals phase.
All 157 assets testcases passed, including the three AST-073 recipe cases.
## Provenance

- evals 2026-11-18T21:58:01.475022Z agent=codex fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=157 failing=0 passing=157
- audit 2026-11-18T21:58:33.285275Z agent=manual fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

## Audit

No findings; checked AST-073

