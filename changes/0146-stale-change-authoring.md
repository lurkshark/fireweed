---
bases: {}
module: project
---

## Motivation

The design's Stale changes and Fixing a stale change clauses require authors to
see what changed beneath a stamped edit and repair missing stamps without
silently resolving stale conflicts. Supply the project queries and narrow
stamping operation needed by show --stale and check --fix before their CLI
integration. Existing Stamp and ordinary Check behavior remains unchanged.

## Interface

Add StaleSpecs(root, change string) ([]StaleSpec, error) and
StampMissing(root, change string) error to api.go. Retain StaleSpec in model/**
and export its alias through api.go: ID, RecordedHash, Before and Current are
strings; Recovered is bool. Before and Current contain rendered complete spec
items, including headings and tags; Before is empty when Recovered is false.
Use thin public facades and implementation-ready inert internal stubs in evals.

## Added

### new:stale-selection
StaleSpecs MUST return the target change's Modified and Removed items whose recorded bases differ from the corresponding preceding projected spec hashes, in edit order, with ID, RecordedHash and Current text.
Items without a bases entry MUST be omitted.

### new:stamped-history
StaleSpecs MUST recover Before by projecting the spec immediately before the target change at the Git commit that last changed that change file's bases mapping.
It MUST set Recovered true only when the recovered item's normalized SpecHash equals RecordedHash.

### new:history-fallback
When matching stamped history cannot be recovered, StaleSpecs MUST return the stale item's Current text and RecordedHash with Recovered false and empty Before.
Unavailable history MUST NOT prevent reporting other stale items.

### new:stale-read
StaleSpecs MUST leave repository files, index and refs unchanged.
It MUST return an empty slice when no recorded bases are stale.

### new:missing-stamps
StampMissing MUST add canonical preceding-projection hashes only for Modified or Removed IDs absent from bases, preserving every existing entry including stale hashes.
It MUST retain other frontmatter values and human section contents.

### new:missing-errors
StampMissing MUST reject a nonmapping bases value or an edited ID absent from the preceding projection with C08 and preserve the target file on rejection.

### new:stale-selectors
StaleSpecs and StampMissing MUST accept a unique pending slug, queued basename or repository-relative queued path.
An unknown selector MUST return U02 naming the requested change; an ambiguous slug MUST return C13.

## Modified

## Removed

## Decisions

## Notes

Preserve all existing cases, assertions and ID coverage. Evals use public APIs,
temporary repositories with explicit Git identity and representative queue
fixtures. Historical tests must distinguish the last bases-changing commit
from a later unrelated file edit, exercise preceding queued projection and
normalized hashes including tags, and cover unavailable and mismatched history.
Exercise mixed missing/valid/stale bases, Modified and Removed, retained extra
entries, aliases, rejection effects and unchanged ordinary Check/Stamp behavior.
No process instrumentation, bootstrap comparison or real drill belongs in evals.

Use the unchanged configured Codex model/effort and normal fail-first, restricted
independent audit and frozen zones. Run context before each step and fireweed
verify for checks. Workers never commit or call done/advance phases; the owning
controller verifies, commits and archives. Keep behavioral choices in Decisions
only when absent from these specs; process notes belong here.

