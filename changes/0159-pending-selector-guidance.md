---
bases:
    PRJ-004: 5cdd4002
    PRJ-066: 44fc667c
module: project
---

## Motivation

ACCEPTANCE C4 requires unknown change names and filenames to identify valid
slugs instead of leaving the user without a usable selector. Project owns the
pending-selector resolver used by projection, stamping and stale-change APIs;
workflow's lifecycle selectors follow in their owning C4 change.

## Interface

No public fields or signatures change. Retain existing aliases, U02 cause and
subject data, static/ambiguity diagnostics and successful operation behavior.

## Added

## Modified

### PRJ-004
Project and Stamp MUST resolve a change argument from its unique slug, queued basename or repository-relative queued path, rejecting missing or ambiguous targets. An unknown target MUST return U02 naming the rejected selector and listing valid pending change slugs.

### PRJ-066
StaleSpecs and StampMissing MUST accept a unique pending slug, queued basename or repository-relative queued path.
An unknown selector MUST return U02 naming the requested change and listing valid pending change slugs; an ambiguous slug MUST return C13.

## Removed

## Decisions

## Notes

Preserve every existing assertion, accepted selector form, queue ordering,
projection/stamping effect, static error cause, ambiguity rejection and the
empty-through meaning of Project. Guidance lists pending slugs, not archive
entries. Empty queues should give usable empty-set guidance. No lifecycle,
app, assets, module SPEC/Decisions or protected-infrastructure edits belong
in this project change; no public interface work is needed.

Use representative public API evals for both Modified IDs and all four
selector APIs without an exhaustive cross product. Keep existing coverage
and own Verify below five minutes, aiming for at most fifteen cases per ID.
Fail-first requires real failed assertions for the missing valid-slug lists.

Workers run pinned fireweed context, verify with fireweed verify --quick,
poll until terminal and return. They never commit, call lifecycle mutators,
edit run/phase/Provenance or alter frozen evals/interfaces. Only the outside
owning controller advances, commits, applies and archives with full normal
verification.

The separate restricted auditor edits only the existing Audit body, retaining
all other bytes. Findings use - ID: "exact clause" — missing assertion. Passing
audit writes No findings and Checked IDs. Re-audits cover only flagged IDs,
retaining previous clearances and every counter/bound.

## Audit

