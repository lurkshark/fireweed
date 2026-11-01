---
bases: {}
module: project
phase: impl
run: 20261101T182723-767c7de65468c682ee0be1f7
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

### PRJ-060
StaleSpecs MUST return the target change's Modified and Removed items whose recorded bases differ from the corresponding preceding projected spec hashes, in edit order, with ID, RecordedHash and Current text.
Items without a bases entry MUST be omitted.

### PRJ-061
StaleSpecs MUST recover Before by projecting the spec immediately before the target change at the Git commit that last changed that change file's bases mapping.
It MUST set Recovered true only when the recovered item's normalized SpecHash equals RecordedHash.

### PRJ-062
When matching stamped history cannot be recovered, StaleSpecs MUST return the stale item's Current text and RecordedHash with Recovered false and empty Before.
Unavailable history MUST NOT prevent reporting other stale items.

### PRJ-063
StaleSpecs MUST leave repository files, index and refs unchanged.
It MUST return an empty slice when no recorded bases are stale.

### PRJ-064
StampMissing MUST add canonical preceding-projection hashes only for Modified or Removed IDs absent from bases, preserving every existing entry including stale hashes.
It MUST retain other frontmatter values and human section contents.

### PRJ-065
StampMissing MUST reject a nonmapping bases value or an edited ID absent from the preceding projection with C08 and preserve the target file on rejection.

### PRJ-066
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

Fail-first: `go test ./packages/project/evals -run 'Test_PRJ_0(60|61|62|63|64|65|66)' -count=1` failed for every added ID against the inert stubs. `fireweed verify` passed the evals phase. The implementation-writable replacement is `packages/project/internal/queue/stale.go` for both new operations; the public facade and model type are retained interfaces.

Use the unchanged configured Codex model/effort and normal fail-first, restricted
independent audit and frozen zones. Run context before each step and fireweed
verify for checks. Workers never commit or call done/advance phases; the owning
controller verifies, commits and archives. Keep behavioral choices in Decisions
only when absent from these specs; process notes belong here.

Recovery after the owner-approved verification-cost diagnosis: preserve the
existing implementation and all frozen interfaces/evals. Use `fireweed verify
--quick` for worker feedback; the owning controller's done step still runs full
dependent checks and boundaries. Do not repeat a full verify inside the worker,
commit, call done, change bounds, reset attempts or edit frozen files. Native
reports from the interrupted pass already show all project, app, checks and
runner cases passing; the complete controller verification remains required.

## Provenance
- evals 2026-11-01T18:31:05.565484Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=126 failing=17 passing=109
- audit 2026-11-01T18:31:54.971154Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked PRJ-060, PRJ-061, PRJ-062, PRJ-063, PRJ-064, PRJ-065, PRJ-066
