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
In audit phase, Verify and Done MUST treat blank separator lines adjoining an
otherwise permitted Audit edit as formatting when comparing the submission with
the approved change.

## Modified

## Removed

## Decisions

## Notes

Use the existing public audit lifecycle fixture. Append the first valid Audit
section after Provenance with a blank separator. Verify the submission before Done
and check the resulting phase. In the same fixture, change the adjacent approved
Provenance line as a negative control and require P03 before advancing. Keep
existing validity and zone assertions. The current workflow accepts the valid form.
