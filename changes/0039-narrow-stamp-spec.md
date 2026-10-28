---
already_satisfied:
    - PRJ-006
    - new:stamp-absent-targets
    - new:stamp-preserve-author-content
    - new:stamp-added-only
bases:
    PRJ-006: 51f3a69b
module: project
---



## Motivation

Complete the authorized remaining project split under SYSTEM's one-behavior
spec shape. Separate PRJ-006's write operation from its rejection and content
preservation rules, retaining the design's Change files canonical bases and
late ID assignment. Existing public evals already assert these behaviors.

## Interface



## Added

### new:stamp-absent-targets
Stamp MUST reject Modified or Removed IDs absent from the projection preceding the target change.

### new:stamp-preserve-author-content
Stamp MUST retain other frontmatter values and human section contents.

### new:stamp-added-only
Stamp MUST succeed on an Added-only change without assigning spec IDs.

## Modified

### PRJ-006
Stamp MUST write each Modified or Removed item's canonical SpecHash from the projection preceding the target change into its bases.

## Removed



## Decisions



## Notes

Retarget existing Stamp evals, preserving every assertion and all 53 project
cases. StampPrecedingProjectionAndPreservation retains PRJ-006 and gains the
content-preservation ID. StampAddedOnlyAndAbsentID gains the Added-only and
absent-target IDs instead of PRJ-006, retaining both Modified and Removed
rejection assertions and its successful Added-only assertion. No new Test
functions, t.Run cases, fixture instances, helpers, product code or interface
changes. Keep every unrelated ID. Implementation may edit only this change file.
