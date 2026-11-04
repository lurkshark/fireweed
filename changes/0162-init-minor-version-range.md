---
bases:
    APP-041: "96246406"
module: app
---

## Motivation

ACCEPTANCE C4 selects a compatible minor-version window for new projects.
Init currently pins =X.Y.Z, unnecessarily excluding later patch builds.
Write the selected >=MAJOR.MINOR, <MAJOR.(MINOR+1) range based on assets.Version.

## Interface

No public fields or signatures change. Only the generated fireweed.yaml
binary-range value changes; its format version and other setup stay intact.

## Added

## Modified

### APP-041
init MUST create fireweed.yaml in the current directory with format version 1, a binary range >=MAJOR.MINOR, <MAJOR.(MINOR+1) derived from assets.Version, base main, remote origin and modules [packages/*].

## Removed

## Decisions

## Notes

Preserve every existing assertion and init behavior, including current-version
acceptance, all other config fields, recipes, presets/custom prompt forms,
audit fallback, diagnostics, non-overwrite rules and non-execution of hooks or
agents. Do not edit other modules, root config, module SPEC/Decisions, protected
setup, bounds or timeouts. This finishes C4; gc/migrate and recipes follow later.

Use a representative public CLI init eval for the exact minor window and its
patch/next-minor semantics. Reuse the existing version-range helper and init
fixtures without weakening header/version checks or duplicating expensive
lifecycle runs. Fail-first must be a genuine APP-041 assertion failure. Keep
own Verify below five minutes and preserve all earlier coverage.

Workers run pinned fireweed context, verify with fireweed verify --quick,
poll until terminal and return. They never commit, call lifecycle mutators,
edit run/phase/Provenance or alter frozen evals/interfaces. Only the outside
owning controller advances, commits, applies and archives with normal full
verification.

The separate restricted auditor edits only the existing Audit body, retaining
all other bytes. Findings use - ID: "exact clause" — missing assertion. Passing
audit writes No findings and Checked IDs. Re-audits cover only flagged IDs,
retaining prior clearances and all counters/bounds.

## Audit

