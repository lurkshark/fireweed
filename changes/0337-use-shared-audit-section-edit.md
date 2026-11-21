---
bases:
    RUN-047: d927095a
module: runner
phase: audit
run: 20261121T004139-c180dc2ec5057749af044a61
---

## Motivation

The separate auditor for changes 0332 and 0334 inserted the first Audit
section before Provenance with a leading blank separator. Next reported P09
and stopped even though the audit contained a valid result. Users need
automatic audit import to accept the same Audit edits that Verify and Done
accept, while keeping the copied workspace isolated.

## Interface

Use workflow's ApplyAuditEdit operation from the preceding workflow change
when Next checks and imports the auditor's edit to the copied change body.

## Added

## Modified

### RUN-047
Next MUST use workflow's Audit edit rule to accept an Audit section edit to
the copied current change, including adjoining blank separators. It MUST
reject edits to the change's frontmatter and edits, additions, or deletions
outside the change file and the workspace's Git metadata. A rejected import
MUST copy none of the workspace edits to the held checkout.

## Removed

## Decisions

## Notes

Extend the existing public Next fixture for first insertion before
Provenance. Add a blank separator immediately before the new Audit heading,
as in 0332 and 0334. Assert that Next imports the Audit section without
changing the protected body or frontmatter, including when an audit invocation
fails and Next later resumes. Retain the existing workspace rejection cases
and add an unrelated blank-line edit to the copied change. Assert P09 and
that no workspace edit reaches the held checkout.

Fail-first evidence for RUN-047 came from
`Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported`. Next returned P09
for the valid insertion and left the Audit section out of the held change and
the resumed auditor's input. All eight invalid-submission cases passed.
The implementation-writable import code is in
`packages/runner/internal/next/audit.go`.
## Provenance

- evals 2026-11-21T00:53:07.977628Z agent=codex fireweed=0.2.4 playbook=5e05bacdae954379b468564c11429c5532016e26a7afbcd418f9bdb9441e151c tests=222 failing=1 passing=221

