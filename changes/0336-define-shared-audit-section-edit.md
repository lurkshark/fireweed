---
bases:
    FLOW-145: 889731d8
module: workflow
phase: impl
run: 20261121T001337-9c2d72f129f4ea534ca5854d
---

## Motivation

Fireweed rejected valid automatic audits for changes 0332 and 0334 when the
auditor inserted a Markdown blank separator before the first Audit heading,
immediately before Provenance. The 0334 audit passed Verify and Done after
manual resubmission. Workflow currently ignores blank lines throughout the
protected body when comparing audit submissions. Audit callers need one rule
that accepts separators at the Audit boundary and protects content elsewhere.

## Interface

Expose `ApplyAuditEdit(original, submitted string) (string, error)` from
workflow. Both inputs are Markdown change bodies. The result is the body to
use when importing the edit. Verify and Done use the same rule to check the
approved body.

## Added

### FLOW-147
ApplyAuditEdit MUST accept an Audit section edit and any changes to blank
separator lines immediately adjoining that edit. It MUST return the body
with the submitted Audit edit applied to the original, preserving the
original content and separators outside Audit. It MUST reject changes
elsewhere in the body and submissions with multiple Audit sections.

## Modified

### FLOW-145
In audit phase, Verify and Done MUST compare the submitted body with the
approved body under the ApplyAuditEdit rule. A change to protected body
content, including a blank line away from the Audit boundary, MUST fail with
P03.

## Removed

## Decisions

## Notes

Exercise first insertion before Provenance and replacement of an existing
Audit section with an adjacent separator change. Assert that the shared
operation imports the Audit edit and preserves the original protected text.
Add a negative case for an unrelated blank-line edit inside protected
content. Exercise acceptance and P03 rejection through the public Verify and
Done paths, preserving the existing assertions about phase and HEAD.

Extend the existing FLOW-075 saved-audit-attempt fixture to save a valid
submission with the separator layout. After StartFromAttempt, assert that
the restored Audit still passes Verify and that the current approved author
fields, run, phase, and Provenance are retained. Retain the existing
restoration contract.

Fail-first: the ApplyAuditEdit stub rejects valid insertions and replacements
(FLOW-147). Verify and Done accept an extra blank line inside Motivation, and
Done advances to impl (FLOW-145). The first Audit insertion before Provenance
and saved-attempt restoration both pass against the current lifecycle.

The ApplyAuditEdit stub is in `internal/lifecycle/audit_edit.go` under the
workflow module. The approval check is in `internal/lifecycle/claim.go`.
## Audit

No findings; checked FLOW-145, FLOW-147

## Provenance

- evals 2026-11-21T00:19:24.331633Z agent=codex fireweed=0.2.4 playbook=5e05bacdae954379b468564c11429c5532016e26a7afbcd418f9bdb9441e151c tests=229 failing=4 passing=225
- audit 2026-11-21T00:19:41.607921Z agent=codex fireweed=0.2.4 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

