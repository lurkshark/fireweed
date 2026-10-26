---
audit_rounds: 0
bases: {}
module: contracts
phase: impl
run: 20261026T203821-10f7bcd023c7b0ff09a196c4
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

### CON-001
ParseDocument MUST decode leading --- delimited YAML mapping frontmatter and retain the remaining Markdown body. It MUST reject malformed, non-mapping or unterminated frontmatter. RenderDocument MUST produce a document whose parsed frontmatter and body equal its input.

### CON-002
Section and ReplaceSection MUST read or replace a named level-two Markdown section without changing other sections, appending an absent section. ParseItems MUST return level-three items in source order with ID, optional bracketed tag and complete body, rejecting duplicate IDs and malformed item headings. RenderItem MUST reproduce an item's heading, tag and body.

### CON-003
SpecHash MUST return the first eight lowercase hex characters of SHA-256 over the rendered item, including its ID and tag, after normalizing CRLF to LF, removing spaces/tabs at line ends and canonicalizing trailing blank lines to one final newline.

### CON-004
MatchesID MUST find a spec ID in either testcase name or classname, ignoring case and treating hyphen, underscore and an omitted separator equally. It MUST reject a match preceded by a letter/digit or followed by a digit.

### CON-005
MatchGlob MUST support slash-separated paths with * within one component and ** spanning zero or more components. ResolveZone MUST choose the first matching zone in order protected, SPEC.md, evals, interface, append_only, preserve, disposable, other, treating SPEC.md as the module's root spec path.

## Modified



## Removed



## Decisions



## Provenance

- evals 2026-10-26T20:46:13.979686Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=61 failing=18 passing=43
- audit 2026-10-26T20:47:35.761624Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings.

Checked: CON-001, CON-002, CON-003, CON-004, CON-005.
