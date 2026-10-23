---
bases:
    PRJ-016: 217b4a9d
fixes:
    - PRJ-016
module: project
phase: evals
run: 54064a814032fbec76c5c142
---

## Motivation

Public allocation currently materializes all three edit headings even when their typed lists are empty and their source headings are absent. This changes authored structure during initial allocation and resume, so a minimal approved change can appear unapproved when lifecycle code compares its started intent. Preserve absent empty sections while continuing to make typed edits authoritative. The public discovery matrix in D027 identifies the regression; implementation and permanent evals belong to this module.

## Interface

Keep the existing Allocate public API and models unchanged. Strengthen retained section preservation under PRJ-016 through public black-box regression evals. Other project functions retain their public contracts.

## Added

## Modified

### PRJ-016
Allocate MUST rewrite every occurrence of each declared handle in canonical change metadata, preamble and every body section, including Added/Modified text, Interface, Decisions, Notes and fenced examples. Overlapping handle names MUST NOT corrupt one another. already_satisfied and nested set hook references MUST become the same assigned IDs as their Added declarations. Returned typed fields and Document MUST agree; unrelated author text, section order, tags, real IDs, lifecycle state and metadata MUST retain their semantic values under the public Markdown normalization contract. Typed change fields MUST override stale retained metadata/edit sections. Returned Path, Slug and Number MUST retain the selected typed identity even when filename numbering differs from Number. On new allocation and valid resume, Added/Modified/Removed sections with nil or empty typed item lists and absent retained headings MUST remain absent. Existing empty edit headings MUST retain their names and positions. Nonempty typed edit lists MUST create or update their corresponding sections without introducing unrelated empty edit sections. All other retained heading presence and relative order MUST remain unchanged, and fenced heading text MUST NOT count as an existing edit section.

## Removed

## Decisions

Known edit headings are synthesized only when their typed lists have items. A retained edit heading remains present and in place even when its authoritative typed list is empty; stale retained items are cleared. A missing heading with nonempty typed items must still be synthesized. Preserve arbitrary human sections and their relative order, including empty sections and fenced examples. Both fresh allocation and valid started resume use the same canonical view. This repairs document preservation rather than weakening approval content checks or adding lifecycle behavior.

## Rejected alternatives

Appending every known heading changes source structure without authored intent. Removing every empty heading also changes explicitly authored templates. Ignoring typed items when their headings are missing breaks typed authority and document agreement. Special-casing workflow approval would hide the allocator regression and leave other consumers with changed documents.

## Provenance

- evals | 2026-10-23T05:36:23Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 439 | failing: 46 | passing: 393
- audit | 2026-10-23T05:44:17Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0

## Audit

- PRJ-016: "Existing empty edit headings MUST retain their names and positions." — all new actual empty edit headings are appended after every human section. The legacy rewrite case checks interleaving when all three authoritative edit lists are nonempty, but its empty-list/resume cases do not assert heading order. A transform that moves only authoritative-empty headings to the end still passes. Add fresh/resume cases with empty edit headings before and between human sections, including a mixture of retained empty headings and a different nonempty typed edit section, and compare the complete retained ordering.
- PRJ-016: "On new allocation and valid resume, Added/Modified/Removed sections with nil or empty typed item lists and absent retained headings MUST remain absent."; "Typed change fields MUST override stale retained metadata/edit sections." — stale retained item bodies are cleared only in an unstarted change in Test_PRJ016AllocationStaleBodiesBecomeRetainedEmptyHeadings. Resume cases currently begin with already-empty retained bodies. Extend stale-body clearing across valid started phases for each edit section and both nil/explicit-empty authoritative lists, retaining the heading and its position while clearing every stale item.
- PRJ-016: "Nonempty typed edit lists MUST create or update their corresponding sections without introducing unrelated empty edit sections."; "Typed change fields MUST override stale retained metadata/edit sections." — the new present=true update fixtures start with empty retained bodies, so they test filling a heading rather than replacing conflicting retained edits. Existing allocation_test.go covers a changed nonempty Added body, but not conflicting nonempty retained Modified/Removed lists. Add fresh/resume replacements where retained and typed IDs/bodies differ, requiring exactly the authoritative typed items and canonical Document agreement.
- PRJ-016: "Nonempty typed edit lists MUST create or update their corresponding sections without introducing unrelated empty edit sections."; "All other retained heading presence and relative order MUST remain unchanged" — the missing-heading fixture requires the new section to be appended at the very end, although the authored clause and Decisions do not specify a placement for newly created headings. Inserting a new section elsewhere while retaining all prior headings' relative order can satisfy the clause but fail this assertion. Either explicitly pin the append-at-end convention in the contract or relax the missing-heading assertion to require the correct new section plus unchanged retained heading presence/relative order.
