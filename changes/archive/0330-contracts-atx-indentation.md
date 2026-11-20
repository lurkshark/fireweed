---
already_satisfied:
    - CON-002
bases:
    CON-002: d56419d5
module: contracts
phase: applied
run: 20261119T222831-c3e628bd39ef3cc8cb6aae11
---

## Motivation

Project validation counts `Specs` headings using `contracts.Section`. The
rebuild missed valid headings with three leading spaces, so it rejected a
valid module and accepted a module with a duplicate `Specs` section. Spell
out the indentation rule so the contracts evals catch both kinds of error.

## Interface

## Added

## Modified

### CON-002
Section MUST return the body and presence of the first matching named level-two Markdown section. A section MUST end at the next level-one or level-two ATX heading outside fenced code. For section starts and boundaries, ATX headings can have zero to three leading spaces. Headings inside fenced code MUST NOT start or end a section.

## Removed

## Decisions

## Notes

Add CON-002 evals through the public `Section` API. Assert that named
level-two headings indented by one, two, or three spaces start a section
and return the expected body. Assert that indented level-one and level-two
headings end an open section and that a heading-like line with four leading
spaces remains in its body. The current implementation already follows this
indentation rule.
## Provenance

- evals 2026-11-19T22:29:30.917421Z agent=codex fireweed=0.2.4 playbook=5e05bacdae954379b468564c11429c5532016e26a7afbcd418f9bdb9441e151c tests=72 failing=0 passing=72
- audit 2026-11-19T22:29:45.81232Z agent=codex fireweed=0.2.4 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-19T22:59:37.782886Z agent=codex fireweed=0.2.4 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=1572 failing=0 passing=1572

## Audit

No findings; checked CON-002

