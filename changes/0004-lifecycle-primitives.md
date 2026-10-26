---
bases: {}
module: contracts
set:
    interface:
        - api.go
        - model/**
---



## Motivation

The design's Specs, Change files and zone-map sections supply the grammar and
checks needed for the NEW binary's first manual lifecycle. Provide just the
shared document, item, ID/hash and zone operations used by that fixture.

## Interface

Expose small data-only Document (Frontmatter map[string]any, Body string), Item
(ID, Tag, Body strings) and Zones (Protected, Evals, Interface, AppendOnly,
Preserve, Disposable []string) types from retained model/ through public aliases.
Thin api.go wrappers expose ParseDocument/RenderDocument, Section/ReplaceSection,
ParseItems/RenderItem, SpecHash, MatchesID, MatchGlob and ResolveZone. Define
signatures in evals phase. Keep all parsing, rendering and lookup behavior under
internal/, importing model rather than the public parent to avoid cycles.

## Added

### new:documents
ParseDocument MUST decode leading --- delimited YAML mapping frontmatter and retain the remaining Markdown body. It MUST reject malformed, non-mapping or unterminated frontmatter. RenderDocument MUST produce a document whose parsed frontmatter and body equal its input.

### new:sections-items
Section and ReplaceSection MUST read or replace a named level-two Markdown section without changing other sections, appending an absent section. ParseItems MUST return level-three items in source order with ID, optional bracketed tag and complete body, rejecting duplicate IDs and malformed item headings. RenderItem MUST reproduce an item's heading, tag and body.

### new:spec-hashes
SpecHash MUST return the first eight lowercase hex characters of SHA-256 over the rendered item, including its ID and tag, after normalizing CRLF to LF, removing spaces/tabs at line ends and canonicalizing trailing blank lines to one final newline.

### new:id-matching
MatchesID MUST find a spec ID in either testcase name or classname, ignoring case and treating hyphen, underscore and an omitted separator equally. It MUST reject a match preceded by a letter/digit or followed by a digit.

### new:zones
MatchGlob MUST support slash-separated paths with * within one component and ** spanning zero or more components. ResolveZone MUST choose the first matching zone in order protected, SPEC.md, evals, interface, append_only, preserve, disposable, other, treating SPEC.md as the module's root spec path.

## Modified



## Removed



## Decisions
