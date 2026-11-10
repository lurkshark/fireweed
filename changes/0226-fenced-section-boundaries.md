---
bases:
    CON-002: 076c0dda
    CON-009: fc9dc150
module: contracts
---

## Motivation

CON-002 and CON-009 do not say where a level-two section ends or whether a
heading inside fenced code counts. The current code ends a section at any
line that looks like a level-two heading, including one inside a fenced
code block, while the contracts drill's rebuild skipped fenced code and
also stopped at level-one headings. Two builds from the same spec therefore
read a change file differently. A Notes section with a fenced Markdown
example containing `## ` would be cut short, and the rest of the example
would become a new section.

## Interface

None. Section and ReplaceSection boundaries are defined.

## Added

## Modified

### CON-002
Section MUST return the body and presence of the first matching named level-two Markdown section. A section MUST end at the next level-one or level-two ATX heading outside fenced code; headings inside fenced code MUST NOT start or end a section.

### CON-009
ReplaceSection MUST replace the first matching named level-two Markdown section, bounded as Section bounds it, without changing other sections, appending the section when absent.

## Removed

## Decisions

## Notes

Fenced code means CommonMark backtick or tilde fences. No change file,
archived change or SPEC.md in this repository has a heading inside fenced
code, so existing projections and bases are unchanged.

Evals: a Notes section containing a fenced block with a `## ` line must be
returned whole by Section and replaced whole by ReplaceSection, and a
level-one heading must end the preceding section.
