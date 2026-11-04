---
bases: {}
module: assets
phase: audit
run: 20261104T115242-ba3a71fd409548c330669535
---

## Motivation

ACCEPTANCE C5 requires the gc age setting and its default to be documented
through the owning module. The backend and CLI have applied; assets owns
embedded Help, which users can read outside a project.

## Interface

No public signatures or recipe payloads change. Extend assets.Help text only.

## Added

### AST-048
Help MUST document that gc prunes stale claims, attempt refs for archived changes and drill refs older than gc.drill_max_age. It MUST identify gc.drill_max_age as a positive Go duration in fireweed.yaml with default 720h (30 days).

## Modified

## Removed

## Decisions

## Notes

Fail-first: `go test ./packages/assets/evals -run '^Test_AST_048_GCHelp$' -count=1` failed against the pre-implementation Help text on all five gc documentation assertions.
Implementation-writable file for the missing Help behavior: `packages/assets/internal/content/content.go`.

Keep every existing command listing and assertion. This closes C5; migrate,
additional recipes and regeneration remain later ordered stages. Do not edit
other modules, project configuration, recipe payloads, protected setup, module
SPEC/Decisions, bounds or timeouts. Use a compact public Help assertion with
genuine fail-first; do not add expensive fixture runs for text documentation.

Workers run pinned fireweed context, verify with fireweed verify --quick,
poll until terminal and return. They never commit, call lifecycle mutators in
the held checkout, edit run/phase/Provenance or alter frozen evals/interfaces.
Only the outside owning controller advances, commits, applies and archives
with normal full verification.

The separate restricted auditor edits only the existing Audit body. Findings
use - ID: "exact clause" — missing assertion. Passing audit writes No findings
and Checked IDs. Re-audits cover only flagged IDs, retaining prior clearances
and all counters/bounds.

## Provenance
- evals 2026-11-04T11:55:53.763419Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=86 failing=1 passing=85
