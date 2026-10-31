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
ParseDocument MUST decode leading --- delimited YAML mapping frontmatter and retain the remaining Markdown body.

### CON-002
Section MUST return the body and presence of a named level-two Markdown section.

### CON-003
SpecHash MUST return the first eight lowercase hex characters of SHA-256 over the rendered item, including its ID and tag, after normalizing CRLF to LF, removing spaces/tabs at line ends and canonicalizing trailing blank lines to one final newline.

### CON-004
MatchesID MUST find a spec ID in either testcase name or classname, ignoring case and treating hyphen, underscore and an omitted separator equally. It MUST reject a match preceded by a letter/digit or followed by a digit.

### CON-005
MatchGlob MUST support slash-separated paths with * within one component and ** spanning zero or more components.

### CON-006
ParseItems MUST accept new:<slug> item headings and preserve the complete handle in Item.ID, where slug contains lowercase letters, digits and hyphens.

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

### CON-016
ErrorCode MUST return the code and subject of the first CodedError in an error chain, including one wrapped with %w, and MUST report not ok when the chain has none.

### CON-017
ParseItems MUST return a CodedError with code C05 and the duplicated ID or handle as subject when a Specs section repeats an item heading.

### CON-018
ParseItems MUST return a CodedError with code C03 when an item heading is malformed.

### CON-019
ParseDocument MUST return a CodedError with code C03 when frontmatter is missing, unterminated, invalid YAML or not a mapping.

## Decisions

[0128-coded-parse-errors] CodedError exposes an optional Cause error. Error returns the cause's text when present and the code otherwise; Unwrap returns Cause.
