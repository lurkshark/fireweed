---
already_satisfied:
    - CON-001
    - CON-002
    - CON-005
    - CON-006
    - new:invalid-frontmatter
    - new:render-document
    - new:replace-section
    - new:parse-items
    - new:invalid-item-headings
    - new:duplicate-items
    - new:render-item
    - new:resolve-zone
    - new:invalid-handles
bases:
    CON-001: "60569833"
    CON-002: "00932638"
    CON-005: 349314ec
    CON-006: 6739d280
module: contracts
---



## Motivation

Batch the previously identified contracts splits under the owner's approved
30-item spec-only cap and SYSTEM's atomic-spec rule. CON-001/002/005 bundle
independent public APIs; CON-006's generic item guarantees belong with the
item parser. Preserve the design's Markdown grammar, handle and zone-map
obligations while giving these behaviors focused IDs. All behavior already
exists. This change affects 13 spec items and changes no public interface or
product behavior.

## Interface



## Added

### new:invalid-frontmatter
ParseDocument MUST reject malformed, non-mapping or unterminated frontmatter.

### new:render-document
RenderDocument MUST produce a document whose parsed frontmatter and body equal its input.

### new:replace-section
ReplaceSection MUST replace a named level-two Markdown section without changing other sections, appending the section when absent.

### new:parse-items
ParseItems MUST return level-three items in source order with ID, optional bracketed tag and complete body, including mixed assigned-ID and handle items.

### new:invalid-item-headings
ParseItems MUST reject malformed item headings.

### new:duplicate-items
ParseItems MUST reject duplicate item IDs, including duplicate handles.

### new:render-item
RenderItem MUST reproduce an item's heading, tag and body.

### new:resolve-zone
ResolveZone MUST choose the first matching zone in order protected, SPEC.md, evals, interface, append_only, preserve, disposable, other, treating SPEC.md as the module's root spec path.

### new:invalid-handles
ParseItems MUST reject empty handles and handles containing characters other than lowercase letters, digits and hyphens.

## Modified

### CON-001
ParseDocument MUST decode leading --- delimited YAML mapping frontmatter and retain the remaining Markdown body.

### CON-002
Section MUST return the body and presence of a named level-two Markdown section.

### CON-005
MatchGlob MUST support slash-separated paths with * within one component and ** spanning zero or more components.

### CON-006
ParseItems MUST accept new:<slug> item headings and preserve the complete handle in Item.ID, where slug contains lowercase letters, digits and hyphens.

## Removed



## Decisions



## Notes

Retarget existing contracts_test.go tests only, preserving every assertion and
all 34 reported contracts cases. ParseDocument retains CON-001; invalid
frontmatter and document round-trip use their Added IDs. Sections retains
CON-002 and gains ReplaceSection; AppendSection uses ReplaceSection and CON-002.
ParseItems uses the new parser ID; RejectInvalidItems uses heading/duplicate
IDs with corresponding existing row labels. ParseHandlesInSourceOrder retains
CON-006 and gains parser ID for mixed ordering and tag/body retention.
RejectInvalidOrDuplicateHandles uses invalid-handle/duplicate IDs and row labels.
RenderItem and ResolveZone use their Added IDs; MatchGlob retains CON-005.
Keep CON-003/004 unchanged. Rename test functions/table labels only: do not
rewrite IDs or handle strings inside fixture text, input Items, function calls
or other test data. No new Test functions, t.Run cases, fixture instances,
helpers, implementation or interface edits. Implementation may edit only this
change file. The common parser/duplicate specs retain the guarantees formerly
repeated in CON-006; no obligation is removed.
