---
eval_fix:
    - APP-156
    - APP-164
module: app
phase: impl
run: 20261116T212029-548f8e2cdd16f7e4db650262
---

## Motivation

APP-156 requires CLI help delegation and AST-008 leaves the help overview's
layout open. The retained assets build lists commands under headings with
bullets, which the application eval's fixed-width row parser did not accept.
APP-164 requires admission against the held approved configuration. Its
incompatible-claim fixture rewrites the claim blob as if configuration were
nested YAML text, which depends on workflow's private storage format.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

For APP-156, use an explicit inventory of supported commands to exercise
CLI help dispatch. For each command assert that CLI help returns the matching
assets page, exit 0 and empty stderr. Keep the overview delegation, help
forms, no-project and invalid-config assertions. AST-008 owns overview
completeness and grouping.

For APP-164, create the claim with an ordinary Start using a second local
binary. Use a fixture pin that accepts the alternate binary and excludes the
binary under test. Follow the AST-001 eval precedent of copying the current
build inputs to a temporary directory and building once with a different
embedded VERSION. Check the accepted claim
through workflow.Approved. Restore a compatible live pin before the rejection
commands so the assertion distinguishes approved admission from live
admission. Keep the modules, coverage, next and run rejections, correct cause
and guidance, plain/JSON envelopes, side-effect and hook/agent checks, and
compatible approved-snapshot cases. Keep the production VERSION and the
claim blob intact.
## Audit

No findings; checked APP-156, APP-164

## Provenance

- evals 2026-11-16T21:29:42.068548Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=403 failing=0 passing=403
- audit 2026-11-16T21:30:02.727607Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

