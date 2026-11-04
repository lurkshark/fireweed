---
bases:
    FLOW-059: c604037a
    FLOW-061: bd22d0d8
module: workflow
---

## Motivation

ACCEPTANCE C4 requires unknown lifecycle change names/filenames to receive
lookup guidance instead of Start's queue-head rejection. It also requires
audit-phase Verify to validate submissions like Done, so an incomplete audit
fails immediately with P09. These are the two workflow-owned C4 fixes.

## Interface

No public fields or signatures change. Retain accepted selector forms,
existing implicit selection, cause classification and lifecycle controls.

## Added

## Modified

### FLOW-059
During audit, Verify MUST apply Done's audit-submission checks. Verify and Done MUST prefix malformed or incomplete submission rejections with P09.

### FLOW-061
Start MUST distinguish an explicit unknown change selector from queue eligibility. Lifecycle operations MUST return U02 for unknown explicit change selectors, naming the rejected value and listing valid slugs for the operation.

## Removed

## Decisions

## Notes

Preserve every existing assertion, accepted selector form, implicit selection,
dirty/static/ownership/stale precondition and cause code. Known but ineligible
Start targets retain P01. Guidance uses each operation's existing pending or
archive lookup set; it does not replace eligibility checks. Cover applicable
selector APIs, including StartFromAttempt, Current/Context, Verify/Done,
Release, Commits and Abandon. No new aliases or public interfaces are needed.

Audit Verify validates only; it does not call Done or advance, normalize/write
submission state, mutate counters/claims or commit. Keep current finding
filtering, complete checked-ID requirements, re-audit scope, round limits and
Done transitions intact. Reuse the existing submission rules rather than
inventing a second format or relaxing validation. Existing audit-zone evals
whose incidental fixture Audit body is incomplete may supply a valid complete
submission during this evals phase; preserve their zone and state assertions.
Do not edit other modules, module SPEC/Decisions, protected setup or bounds.

Use representative public API evals with genuine fail-first for both Modified
IDs. Preserve all earlier coverage. Keep own Verify below five minutes and
avoid an exhaustive selector/form/precondition cross product or duplicated
expensive lifecycle fixtures.

Workers run pinned fireweed context, verify with fireweed verify --quick,
poll until terminal and return. They never commit, call lifecycle mutators in
the held checkout, edit run/phase/Provenance or alter frozen evals/interfaces.
Fixture lifecycle calls in public-interface evals remain necessary tests.
Only the outside owning controller advances, commits, applies and archives,
with normal full verification.

The separate restricted auditor edits only the existing Audit body, retaining
all other bytes. Findings use - ID: "exact clause" — missing assertion. Passing
audit writes No findings and Checked IDs. Re-audits cover only flagged IDs,
retaining prior clearances and all counters/bounds.

## Audit

