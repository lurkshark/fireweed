---
module: contracts
phase: impl
run: f511aa9df35ecc447d267333
---

## Motivation

The new executable needs language-independent spec parsing, stable bases and precise test-to-spec matching. These primitives establish the retained document contract without relying on bootstrap internals.

## Interface

Expose Spec {ID, Tag, Body string}; ParseSpecs(document, prefix string) ([]Spec, error); ParseItems(section, prefix string, allowHandles bool) ([]Spec, error); SpecHash(Spec) string; and MatchesID(name, classname, id string) bool. SpecHash hashes the complete rendered heading and body. ParseSpecs reads only one required Specs section; ParseItems reads a section body for future change parsing.

## Added

### CON-001
Test ID matching MUST ignore case, accept hyphen, underscore or no separator between prefix and digits, and search both testcase name and classname. A match MUST NOT be preceded by a letter or digit or followed by a digit.

### CON-002
A spec base hash MUST be the first eight lowercase hexadecimal characters of SHA-256 over its heading (including its tag) and body, normalizing CRLF to LF, trimming trailing whitespace on each line and trimming surrounding blank lines. Meaningful text and tag changes MUST change the hash.

### CON-003
Module spec parsing MUST require exactly one level-two Specs section, parse level-three PREFIX-NNN headings in order only inside that section, retain item bodies, and reject wrong-prefix or duplicate IDs, malformed headings, missing normative MUST, MUST NOT or SHOULD text, and empty bodies. Headings and apparent sections inside fenced code MUST remain body text.

### CON-004
Spec parsing MUST accept untagged items, review tags and quarantine tags with a nonempty reason, including review and quarantine together. It MUST reject unknown, repeated or malformed tags and quarantine without a reason. Section item parsing MUST accept unique lowercase-letter/digit/hyphen new: handles only when explicitly allowed and reject malformed or duplicate handles.

## Modified

## Removed

## Decisions

Canonical rendering for bases is `### <ID><optional space and tag>\n<body>`, with no blank separator inserted. Tags are retained exactly after outer whitespace trimming; semantic text is never case-folded for hashing. Blank outer body lines are insignificant. IDs require a letter-led alphanumeric prefix and at least three decimal digits. Document item parsing ignores prose before its first item; nested level-four headings stay in the body. Invalid level-three headings fail validation rather than silently disappearing.

## Rejected alternatives

Regex-only whole-document parsing would treat fenced examples as live headings. Hashing only body text would silently miss quarantine and review edits.

## Provenance

- evals | 2026-10-22T06:23:30Z | agent: manual | fireweed: 0.4.0 | playbook: 5867363302d4c874074520977b1627d26991d95ee2fbacb9ff63e827661edd0b | tests: 4 | failing: 4 | passing: 0
- audit | 2026-10-22T06:24:46Z | agent: manual | fireweed: 0.4.0 | playbook: a9c982d24b476621316049b527e6c7bbc319f06b12268af150fa7d8c6df92230 | tests: 0 | failing: 0 | passing: 0

## Audit

No findings. Independently checked CON-001, CON-002, CON-003 and CON-004 against all corresponding assertions. Auditor: separate context /root/audit_spec_primitives, restricted to change, module spec, public interface and evals; no implementation or history access.
