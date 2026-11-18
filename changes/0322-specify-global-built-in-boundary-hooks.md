---
already_satisfied:
    - "new:global-recipe-boundaries"
module: assets
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

### new:global-recipe-boundaries
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
