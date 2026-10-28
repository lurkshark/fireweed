---
bases:
    PRJ-008: 4efdf265
    PRJ-009: 3ab806fe
    PRJ-010: ac9cb4e7
    PRJ-011: 074dd58a
    PRJ-012: ad7c76da
module: project
---



## Motivation

The design requires stable diagnostic causes for agents and explain. D006 assigns
C07 to missing Motivation, C04 to normative text, C06 to spec tags and C05 to ID
ownership/uniqueness. Existing Check paths still label each C01. Give those
existing rejections their assigned codes without adding errors or features.

## Interface

No public interface change.

## Added



## Modified

### PRJ-008
Check MUST reject a pending change without a nonempty Motivation section with diagnostic C07.

### PRJ-009
Check MUST reject a module spec or Added/Modified spec whose body contains no normative keyword with diagnostic C04.

### PRJ-010
Check MUST reject spec tags other than review or quarantine with a nonempty reason with diagnostic C06.

### PRJ-011
Check MUST reject duplicate spec IDs within a module's Specs section with diagnostic C05.

### PRJ-012
Check MUST reject module spec IDs and change edit IDs whose prefix belongs to another module with diagnostic C05.

## Removed



## Decisions



## Notes

Use existing CheckValidAndDiagnostics table fixtures and its follow-up Modified
normative/tag assertions to assert the exact assigned code for these IDs. Keep
all valid and invalid inputs, reason and nonempty code/message assertions and ID
coverage. Reuse the 53 cases; no new fixture, test case or helper is needed.
Unchanged IDs retain their checks. The exact cause assertion must match the
rejection being tested, not an unrelated diagnostic from the same fixture.
These changed IDs must genuinely fail before implementation; no already_satisfied.
Implementation migrates only the existing error branches that supply these
causes. Keep messages, existing rejection behavior and other cause labels intact.
Do not add unsupported features or speculative error paths. Evals and interfaces
freeze for implementation; controller owns specs, commits and archive.
