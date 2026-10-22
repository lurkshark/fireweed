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

## Decisions

[spec-primitives]
Canonical rendering for bases is `### <ID><optional space and tag>\n<body>`, with no blank separator inserted. Tags are retained exactly after outer whitespace trimming; semantic text is never case-folded for hashing. Blank outer body lines are insignificant. IDs require a letter-led alphanumeric prefix and at least three decimal digits. Document item parsing ignores prose before its first item; nested level-four headings stay in the body. Invalid level-three headings fail validation rather than silently disappearing.
