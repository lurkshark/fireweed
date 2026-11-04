---
bases:
    APP-041: "96246406"
module: app
phase: impl
run: 20261104T101354-2c33ac1a4ecbf37807fa3201
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

No findings; checked APP-041
## Provenance
- evals 2026-11-04T10:27:59.744809Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=306 failing=1 passing=305
- audit 2026-11-04T10:28:30.466184Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
