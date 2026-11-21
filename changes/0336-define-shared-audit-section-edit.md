---
bases:
    FLOW-145: 889731d8
module: workflow
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

### new:apply-audit-edit
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
