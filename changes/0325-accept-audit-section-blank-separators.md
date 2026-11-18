---
already_satisfied:
    - "new:audit-section-separators"
module: workflow
---

## Motivation

The November 18 workflow drill rejected valid audit edits as unauthorized because
appending an Audit section added a blank separator after approved content. Its
retries often led to a second Audit heading. Treating surrounding separator lines as
formatting restored 49 dependent cases. Auditors need to submit the allowed section
without touching approved author content.
[Drill reports](https://github.com/lurkshark/fireweed/pull/25)

## Interface

Specify the allowed blank separators when Verify or Done checks an Audit edit
against the approved change.

## Added

### new:audit-section-separators
In audit phase, Verify and Done MUST accept a valid Audit submission whose only
edits outside Audit are blank lines separating it from an adjacent section. Changes
to approved author content or runtime fields MUST still be rejected.

## Modified

## Removed

## Decisions

## Notes

Use the existing public audit lifecycle fixture. Append the first valid Audit
section after Provenance with a blank separator, then verify and advance through
Done. Exercise a substantive change to the adjacent approved content as a negative
control for P03. Scope the new assertion to separators adjacent to Audit. The
current workflow accepts the valid form.
