---
module: project
---

## Motivation

The design's Change files / Body sections states that Removed contains
headings only (R05). Main 3b3b64d accepts and stamps a Removed item with body
text, which is then ignored by projection. Reject that misleading input.

## Interface

No public signature or type changes.

## Added

### new:removed-headings-only
Check MUST reject non-whitespace content anywhere in a pending change's Removed section other than valid assigned-ID headings, using non-warning C07 and identifying the change and offending item when present. An empty Removed section or valid headings separated only by whitespace MUST remain valid when their targets and bases are valid.

## Modified



## Removed

## Decisions

## Notes

Use existing projection/stamp fixtures. Cover an item body, prose before the
first heading, a prose-only section, an empty section, and multiple valid
headings separated by blank lines. Inspect the full section because parsed
item bodies omit leading prose. Keep absent/stale target diagnostics C08
and existing malformed-heading/identity codes intact; C07 concerns otherwise
ignored non-heading content. Do not alter the shared item parser
or discard authored prose as a repair; report it so the author can move it
into Motivation or Notes. The assets follow-up describes this C07 cause.

Drafted at the owner's request after the R01–R29 traceability review.
No lifecycle execution, agent invocation, drill or release is authorized by
this drafting request. Workers must obtain context before their step and use
fireweed verify; only the controller may advance or commit lifecycle work.
