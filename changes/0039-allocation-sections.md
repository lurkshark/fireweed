---
bases:
    PRJ-016: 217b4a9d
fixes:
    - PRJ-016
module: project
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
