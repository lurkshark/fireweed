---
already_satisfied:
    - new:queue-order
    - new:queue-payload
    - new:queue-malformed
    - new:queue-module
    - new:queue-slugs
bases: {}
module: project
---



## Motivation

Continue the authorized PRJ-004 split under SYSTEM's one-observable-behavior
spec shape, retaining the design's Change files queue and document rules.
Give Queue's existing behaviors dedicated IDs before narrowing PRJ-004 to
change-argument resolution in the next change. Existing assertions already
cover these rules, so implementation stays unchanged.

## Interface



## Added

### new:queue-order
Queue MUST return pending NNNN-slug.md changes in numeric-then-slug order, allowing tied numbers and excluding archives.

### new:queue-payload
Queue MUST return each pending change's module, frontmatter and parsed Added, Modified and Removed items.

### new:queue-malformed
Queue MUST reject a malformed pending change document.

### new:queue-module
Queue MUST reject a pending change naming an unknown module.

### new:queue-slugs
Queue MUST reject pending changes that share a slug.

## Modified



## Removed



## Decisions



## Notes

Retarget existing queue evals while preserving PRJ-004 for this change.
QueueOrderContentAndArchive gains the order and payload IDs;
QueueRejectsInvalidDocuments gains the malformed/module IDs with corresponding
row labels; TargetAliasesAndAmbiguousSlug gains the slug ID on its existing
ambiguous-queue assertion. Retain every assertion and all 53 project cases.
No new Test functions, t.Run cases, fixture instances, helpers, product code
or interface edits. The next change narrows PRJ-004 only after these rules
have IDs. Implementation may edit only this change file.
