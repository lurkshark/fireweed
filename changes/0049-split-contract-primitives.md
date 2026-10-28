---
already_satisfied:
    - CON-001
    - CON-002
    - CON-005
    - CON-006
    - CON-007
    - CON-008
    - CON-009
    - CON-010
    - CON-011
    - CON-012
    - CON-013
    - CON-014
    - CON-015
audit_rounds: 0
bases:
    CON-001: "60569833"
    CON-002: "00932638"
    CON-005: 349314ec
    CON-006: 6739d280
module: contracts
phase: impl
run: 20261028T062314-ffc823125b2ebb9b8c75c473
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

### CON-007
ParseDocument MUST reject malformed, non-mapping or unterminated frontmatter.

### CON-008
RenderDocument MUST produce a document whose parsed frontmatter and body equal its input.

### CON-009
ReplaceSection MUST replace a named level-two Markdown section without changing other sections, appending the section when absent.

### CON-010
ParseItems MUST return level-three items in source order with ID, optional bracketed tag and complete body, including mixed assigned-ID and handle items.

### CON-011
ParseItems MUST reject malformed item headings.

### CON-012
ParseItems MUST reject duplicate item IDs, including duplicate handles.

### CON-013
RenderItem MUST reproduce an item's heading, tag and body.

### CON-014
ResolveZone MUST choose the first matching zone in order protected, SPEC.md, evals, interface, append_only, preserve, disposable, other, treating SPEC.md as the module's root spec path.

### CON-015
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

## Provenance

- evals 2026-10-28T06:24:20.408504Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=261 failing=0 passing=261
- audit 2026-10-28T06:25:05.683247Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings.

Checked IDs: CON-001, CON-002, CON-003, CON-004, CON-005, CON-006, CON-007, CON-008, CON-009, CON-010, CON-011, CON-012, CON-013, CON-014, CON-015.
