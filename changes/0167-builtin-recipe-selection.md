---
already_satisfied:
    - APP-045
bases:
    APP-045: d29266b3
module: app
---

## Motivation

ACCEPTANCE C7 selects additional embedded recipes, each proved by fresh
init --recipe <name> and doctor. Init currently recognizes only the built-in
go name. Generalize selection through the owning assets.BuiltinRecipe API
before the assets recipe change, preserving Go, none and local behavior.

## Interface

No public signatures or dependency behavior change. Keep initialization's
shared copy/render/configuration path and broaden only built-in selection.

## Added

## Modified

### APP-045
init --recipe <name> MUST accept any currently available recipe returned by assets.BuiltinRecipe(name) and copy its Files into their project-relative destinations without changing their contents.

## Removed

## Decisions

## Notes

APP-045 is legitimately already_satisfied at this stage: assets' current public
AST-004 contract provides only go and rejects unknown names. The existing
public GoInit eval mapped to APP-045 compares every returned Files entry byte
for byte after init. Reuse those assertions and all existing none/local/unknown
recipe, diagnostic, non-overwrite, configuration and non-execution assertions.
Do not fabricate a second built-in, modify assets or mock its implementation
to manufacture fail-first. No new recipe is claimed available by this change.

Implementation should ask assets.BuiltinRecipe for a non-none built-in name
and retain existing local-directory and unknown-name handling when it is not
available. Keep common file copying, template rendering and defaults intact.
The subsequent owning assets change must add the two real bundles and prove
each with doctor --no-agent in its own fresh init --recipe <name> project;
neither proof is waived or replaced by this preparatory app change.

C7 prerequisites were checked: existing pnpm 10.15.1 runs with that project
packageManager pin without downloads, and uv is available. Do not install or
replace tools. Keep those native recipe proofs, canaries, graph-derived
boundaries and scratch invoker-copy/PATH requirements in the recipe change.

Preserve every assertion and public interface. Do not edit other modules,
root config, module SPEC/Decisions, protected setup, bounds or timeouts. Keep
own Verify below five minutes and full impl Verify below twenty-five.
Workers run pinned fireweed context, verify with fireweed verify --quick,
poll until terminal and return. They never commit, invoke lifecycle mutators
in the held checkout, edit run/phase/Provenance or alter frozen evals/interfaces.
Only the outside owning controller advances, commits, applies and archives
with normal full verification.

The separate restricted auditor edits only the existing Audit body. Findings
use - ID: "exact clause" — missing assertion. Passing audit writes No findings
and Checked IDs. Re-audits cover only flagged IDs, retaining prior clearances
and all counters/bounds.

