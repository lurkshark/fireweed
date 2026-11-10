---
module: assets
---

## Motivation

0233 adds the design's spec-shape warnings C17, C18 and C19, which nothing
implements today. Every emitted code needs an `explain` entry with a
corrective action.

## Interface

None. Explain gains the C17, C18 and C19 entries.

## Added

### new:spec-limit-explanations
Explain MUST describe C17 as a spec item over spec_limits.max_clauses normative keywords or max_words words, C18 as a sentence repeated in three or more spec items, and C19 as a module Decisions section over spec_limits.max_decision_words, each as a warning with a corrective action.

## Modified

## Removed

## Decisions

## Notes

Corrective actions, following the design: split or narrow an oversized item;
move a repeated sentence into SYSTEM.md as a convention; drop stale
decisions through Removed decisions or turn lasting ones into specs. Name
the defaults (3, 80 and 300). Run before 0233.
