---
module: assets
phase: impl
run: 20261112T024709-8a7cd75303833a0d8b458750
---

## Motivation

0233 adds the design's spec-shape warnings C17, C18 and C19, which nothing
implements today. Every emitted code needs an `explain` entry with a
corrective action.

## Interface

None. Explain gains the C17, C18 and C19 entries.

## Added

### AST-060
Explain MUST describe C17 as a spec item over spec_limits.max_clauses normative keywords or max_words words, C18 as a sentence repeated in three or more spec items, and C19 as a module Decisions section over spec_limits.max_decision_words, each as a warning with a corrective action.

## Modified

## Removed

## Decisions

## Notes

Corrective actions, following the design: split or narrow an oversized item;
move a repeated sentence into SYSTEM.md as a convention; drop stale
decisions through Removed decisions or turn lasting ones into specs. Name
the defaults (3, 80 and 300). Run before 0233.

Fail-first verification passed with `fireweed verify spec-limit-explanations`.
The focused AST-060 eval run failed for C17, C18 and C19 because
`assets.Explain` returned an unknown diagnostic code for each.
The implementation can add the entries in
`packages/assets/internal/content/content.go`.
## Provenance
- evals 2026-11-12T02:50:33.803783Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=129 failing=3 passing=126
- audit 2026-11-12T02:51:12.668506Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked AST-060
