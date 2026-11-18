---
already_satisfied:
    - "new:retained-audit-finding-text"
module: workflow
---

## Motivation

The November 18 workflow reconstruction duplicated the separator in a validated
audit finding when rendering it for later phases. Runner evals found the altered
finding in the phase base, blocked change, saved attempt and returned state.
Reviewers and resumed workers need the retained finding the auditor actually
submitted.
[Drill reports](https://github.com/lurkshark/fireweed/pull/25)

## Interface

Specify the existing representation of findings that workflow.Done retains after
validation.

## Added

### new:retained-audit-finding-text
When audit Done retains a finding, it MUST preserve the finding line's text apart
from surrounding whitespace in the recorded Audit section. If the finding blocks the
change, its line in Questions MUST retain the same text.

## Modified

## Removed

## Decisions

## Notes

Use a valid finding with a quoted clause and an explanation separated by an em dash.
Check the retained line through the committed change after a return to evals and
through the blocked change's Questions on the blocking path. Continue asserting the
phase, round count and existing filtering rules. Main already stores the original
validated line.
