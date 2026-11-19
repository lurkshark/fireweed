---
already_satisfied:
    - CON-002
bases:
    CON-002: d56419d5
module: contracts
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
