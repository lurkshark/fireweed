---
bases: {}
module: assets
---

## Motivation

ACCEPTANCE C4 selects clear Go recipe and module-template guidance for the
preceding-letter boundary in eval names. Agents can otherwise write
TestGRT001 and receive an unmapped-test diagnostic even though the name looks
like GRT-001. Ship a valid example and the explanation in both bundled guides.

## Interface

No public fields or signatures change. BuiltinRecipe(go) exposes the updated
instructions and module SPEC template through its existing outputs.

## Added

### new:go-eval-id-boundary
The Go recipe instructions and module SPEC template MUST explain that a letter immediately before a spec prefix prevents ID matching, showing Test_GRT_001_Name as a matching name for GRT-001 and TestGRT001 as a name that does not match.

## Modified

## Removed

## Decisions

## Notes

Preserve every existing assertion, recipe file, template substitution, hook,
canary, setup command, retained-facade/stub example and other instruction.
This change adds naming guidance; it does not change contracts.MatchesID,
test discovery, hook behavior, recipes or interfaces. Do not edit other
modules, protected project-owned copies, module SPEC/Decisions or bounds.
The outside author will copy only the approved guidance into the current
project-owned recipe/template between claims after this applies.

Use a small public BuiltinRecipe eval for both exposed guide locations.
Keep fail-first genuine, all earlier coverage intact and own Verify below
five minutes; reuse existing recipe fixtures rather than adding executions.

Workers run pinned fireweed context, verify with fireweed verify --quick,
poll until terminal and return. They never commit, call lifecycle mutators,
edit run/phase/Provenance or change frozen evals/interfaces. Only the outside
owning controller advances, commits, applies and archives with normal full
verification.

The separate restricted auditor edits only the existing Audit body, retaining
all other bytes. Findings use - ID: "exact clause" — missing assertion. Passing
audit writes No findings and Checked IDs. Re-audits cover only flagged IDs,
retaining prior clearances and all counters/bounds.

## Audit

