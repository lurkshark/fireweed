---
depends_on: []
id_prefix: CON
name: contracts
---

# contracts

## Specs

### CON-001
Test ID matching MUST ignore case, accept hyphen, underscore or no separator between prefix and digits, and search both testcase name and classname. A match MUST NOT be preceded by a letter or digit or followed by a digit.

### CON-002
A spec base hash MUST be the first eight lowercase hexadecimal characters of SHA-256 over its heading (including its tag) and body, normalizing CRLF to LF, trimming trailing whitespace on each line and trimming surrounding blank lines. Meaningful text and tag changes MUST change the hash.

### CON-003
Module spec parsing MUST require exactly one level-two Specs section, parse level-three PREFIX-NNN headings in order only inside that section, retain item bodies, and reject wrong-prefix or duplicate IDs, malformed headings, missing normative MUST, MUST NOT or SHOULD text, and empty bodies. Headings and apparent sections inside fenced code MUST remain body text.

### CON-004
Spec parsing MUST accept untagged items, review tags and quarantine tags with a nonempty reason, including review and quarantine together. It MUST reject unknown, repeated or malformed tags and quarantine without a reason. Section item parsing MUST accept unique lowercase-letter/digit/hyphen new: handles only when explicitly allowed and reject malformed or duplicate handles.

### CON-005
Glob matching MUST support literals, * and ? within a path segment, bracket classes/ranges/negation, and ** across segments, including zero directories for **/. It MUST match the complete relative path, preserve case and reject malformed classes, absolute paths, traversal components and backslashes. Empty paths and empty patterns MUST be rejected.

### CON-006
Glob intersection MUST determine whether two valid patterns can match a common path, without requiring that path to exist; it MUST account for wildcards, classes and zero-directory globstars, and return an error for either invalid pattern.

### CON-007
Zone classification MUST apply strict precedence: protected, SPEC.md, evals, interface, append-only, preserved, disposable, other module files. Protected paths MUST include root fireweed.yaml and SYSTEM.md, all .fireweed descendants, every AGENTS.md and CLAUDE.md, and configured protected globs, even outside the module. Files outside the module MUST otherwise be outside; module path containment MUST use segment boundaries. Invalid repository paths or globs MUST fail classification.

## Decisions

[spec-primitives]
Canonical rendering for bases is `### <ID><optional space and tag>\n<body>`, with no blank separator inserted. Tags are retained exactly after outer whitespace trimming; semantic text is never case-folded for hashing. Blank outer body lines are insignificant. IDs require a letter-led alphanumeric prefix and at least three decimal digits. Document item parsing ignores prose before its first item; nested level-four headings stay in the body. Invalid level-three headings fail validation rather than silently disappearing.

[zone-primitives]
Globs are case-sensitive with bracket negation `!` or `^`. Path components equal to . or .., repeated slashes and trailing slashes are invalid. A globstar matches slash only when it occupies a complete segment; adjacent ordinary stars remain segment wildcards. Zone fields are module-relative, configured protected globs repository-relative. Literal module directory itself is outside (zones classify files below it). All configured patterns are validated even when a stricter zone wins. The zone resolver classifies only; preserved interface and append-only regeneration behavior follows from strict precedence.
