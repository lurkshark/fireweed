---
bases:
    AST-003: 56ae01a3
module: assets
---



## Motivation

The design requires every emitted diagnostic to have useful explain output.
D006 migration now emits static, check and lifecycle causes that Explain still
rejects, while C07/U02 text is narrower than the current change-target and lookup
errors. Complete explanations for actually emitted codes, retaining every
existing cause and corrective action without unused allocations or features.

## Interface

No public interface change.

## Added

### new:static-explanations
Explain MUST describe the cause and a corrective action for C02 (invalid module graph), C03 (unreadable or malformed documents/items), C12 (invalid metadata or zones) and C13 (queue identity errors).

### new:eval-explanations
Explain MUST describe the cause and a corrective action for C10 (missing enforced coverage), C11 (eval-hook/JUnit contract failures) and E01 (failed enforced testcases).

### new:lifecycle-explanations
Explain MUST describe the cause and a corrective action for P01 (preconditions), P02 (claims), P03 (unapproved state alteration), P04 (approval warning), P05 (stale bases), P06 (blocked change), P07 (unapproved-work bound), P09 (audit submissions) and Z01 (forbidden phase-zone edits).

## Modified

### AST-003
Explain MUST describe the cause and corrective action for C01 (invalid configuration), C04 (non-normative text), C05 (foreign or duplicate spec/edit identities), C06 (invalid tags), C07 (missing Motivation or invalid change module target), C08 (missing/stale edit bases or absent targets), C09 (undeclared handles), X01 (boundary violations), X02 (missing/incomplete boundary hooks), X03 (reasoned boundary opt-out), X04 (unnecessary boundary enforcement), U01 (invalid arguments/options or unsupported commands), U02 (unknown or ambiguous requested names) and P2-01 (missing genuine fail-first evidence). Explain MUST reject unknown codes.

## Removed



## Decisions



## Notes

Keep all existing explanation and unknown-code assertions. Strengthen the C07
case for missing/invalid change module targets, and U02 for requested module,
item/change and information lookup names, retaining Motivation and information
lookup coverage. For each newly supported code use one representative cause
and corrective-action assertion in the existing public explanation test style,
reusing requirePattern; no fixtures, helper infrastructure or redundant matrices.
The new code groups and clarified AST-003 behavior must genuinely fail before
implementation. Keep each code-family case count modest.

Explain only the sixteen newly supported emitted codes plus the existing text
clarifications. C10 needs a mapped non-skipped case, not necessarily a passing
one. C11 includes hook launch/setup/report/mapping/exit disagreement. C13 does
not prohibit tied numeric queue positions: explain duplicate slugs or unparseable
numbers, with an unused slug/valid numeric prefix as the repair. C07 concerns
the pending change's module field; U02 concerns a requested lookup, including
ambiguous projected local handles. Keep P04 explicitly a warning, and P07 the
shared-base bound resolved through review/synchronization. P05 repairs stale
bases or restores intended state through the approved workflow; P06 requires
review of Questions; P09 needs No findings and a complete Checked IDs list.
P02 actions must not rely on an unimplemented command. Z01 restores forbidden
edits or moves work to the appropriate authorized phase/zone. Do not recommend
bypassing enforcement, weakening evals or changing claims to evade bounds.
No explanations/tests for unused B/H allocations, P08, C17–C19, reserved numbers
or unsupported features. Recipes, playbooks, public interfaces and other
behavior remain unchanged. Controller owns specs, commits and archive.
