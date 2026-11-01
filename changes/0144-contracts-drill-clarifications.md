---
already_satisfied:
    - CON-002
    - CON-009
bases:
    CON-002: f7885c49
    CON-008: 71986f82
    CON-009: 23ed3cc6
    CON-010: be0602a2
module: contracts
---

## Motivation

The independent NEW contracts drill passed its own and all transitive dependent
evals but reported three ambiguities preventing DR4 closure. The owner explicitly
authorized scoped clarifications using the reported interpretations: reject nil
frontmatter while limiting round-trip type identity to YAML representations,
use uppercase alphanumeric assigned-ID prefixes and a decimal suffix, and select
the first matching section when names repeat. The design's Regeneration section
requires reviewed gaps to become owner-agreed changes; this is that follow-up,
not permission for new features or a wholesale main implementation replacement.

A public check of current main accepts nil frontmatter and lowercase ID prefixes,
and rejects digit-leading prefixes. These supply genuine fail-first cases.
Section and ReplaceSection already select the first duplicate; preserve that
behavior and all existing coverage using legitimate already_satisfied entries.

## Interface

No public signature, type, module metadata or configuration changes.

## Added

## Modified

### CON-002
Section MUST return the body and presence of the first matching named level-two Markdown section.

### CON-008
For non-nil frontmatter maps in ParseDocument's YAML scalar, sequence and mapping representations, RenderDocument MUST produce a document whose parsed frontmatter and body equal its input. It MUST reject nil frontmatter maps; other YAML-serializable Go values are not required to retain their concrete types.

### CON-009
ReplaceSection MUST replace the first matching named level-two Markdown section without changing other sections, appending the section when absent.

### CON-010
ParseItems MUST return level-three items in source order with ID, optional bracketed tag and complete body, including mixed assigned-ID and handle items. An assigned ID MUST consist of a nonempty uppercase ASCII alphanumeric prefix, one hyphen and a nonempty ASCII decimal suffix.

## Removed

## Decisions

## Notes

The owner approved these interpretations after reviewing the real drill's gap
prose. Use the prose here, not the report's Agent observations, retained rebuilt
code, old implementation/history, saved attempts or bootstrap as an implementation
oracle. No real drill runs in evals. This ordinary scoped change does not import
the drill implementation or authorize another module's drill.

Preserve all existing contracts cases, assertions and ID coverage. Add focused
public assertions for nil-frontmatter rejection, representative non-nil YAML
scalar/sequence/mapping round trips with exact body retention, assigned-ID
acceptance (including digit-leading prefixes) and malformed lowercase/empty/
punctuation forms, plus duplicate-section selection and replacement with other
sections byte-unchanged. Reuse the existing fixtures; no speculative edge-case
matrix. A non-nil empty mapping remains supported. Arbitrary Go concrete type
identity is outside the clarified round-trip promise, not a new rejection rule.

CON-008 and CON-010 require genuine failing behavioral cases before implementation.
CON-002 and CON-009 are already satisfied by the public duplicate-section check;
their strengthened regression assertions should pass. Do not manufacture failure
by stubbing or regressing correct implementation. Keep normal separate restricted
audit, frozen evals/interfaces, approval and all configured bounds. Workers read
context before each step and never commit or call done manually; the controller
owns verification and phase commits. Use focused public tests for iteration;
automatic next/done still run the unchanged full phase checks.

## Audit

Pending independent audit.
