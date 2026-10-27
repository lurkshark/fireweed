---
depends_on: []
id_prefix: CON
interface:
    - api.go
    - model/**
name: contracts
---

# contracts

## Specs

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

### CON-006
ParseItems MUST accept new:<slug> item headings and preserve the complete handle in Item.ID, where slug contains lowercase letters, digits and hyphens. It MUST reject empty handles and handles containing other characters. Assigned-ID and handle items MUST retain their source order in mixed input, with the existing tag/body preservation and duplicate-ID checks applying to handles too.

## Decisions
