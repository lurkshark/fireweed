---
module: app
eval_fix: [APP-156, APP-164]
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
