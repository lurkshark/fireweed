---
bases: {}
module: assets
phase: applied
run: 20261104T095615-46a54aa85931068e8a8dedf7
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

### AST-047
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

Fail-first: `go test ./packages/assets/evals -run '^Test_AST_047_GoRecipeEvalNameBoundary$' -count=1` failed for both the recipe instructions and module SPEC template against the current bundle. The implementation-writable replacements are `packages/assets/internal/recipe/files/go.md` and `packages/assets/internal/recipe/files/SPEC.md`.

## Audit

No findings; checked AST-047
## Provenance
- evals 2026-11-04T10:00:03.624915Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=85 failing=3 passing=82
- audit 2026-11-04T10:00:31.74772Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-04T10:11:03.315785Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=947 failing=0 passing=947
