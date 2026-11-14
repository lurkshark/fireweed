---
module: project
phase: impl
run: 20261114T010921-44364fbcb197697a04f1101e
---

## Motivation

The design's Change files / Body sections states that Removed contains
headings only (R05). Main 3b3b64d accepts and stamps a Removed item with body
text, which is then ignored by projection. Reject that misleading input.

## Interface

No public signature or type changes.

## Added

### PRJ-078
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

The targeted PRJ-078 run failed on all four prose fixtures because Check
returned no diagnostics. The empty and whitespace-separated fixtures passed,
as did the C03, C05 and C08 diagnostic cases. The PRJ-072 warning assertions
still pass with a heading-only removal in their valid fixture.
The implementation can add the Removed section check in the disposable
`packages/project/internal/queue/validate.go` file.
## Audit

No findings; checked PRJ-078

## Provenance

- evals 2026-11-14T01:13:13.439281Z agent=codex fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=192 failing=5 passing=187
- audit 2026-11-14T01:13:31.733902Z agent=codex fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

