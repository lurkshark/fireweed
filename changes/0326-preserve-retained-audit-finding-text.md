---
already_satisfied:
    - FLOW-146
module: workflow
phase: impl
run: 20261118T232253-9f0607fb6aa31868e916fe83
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

### FLOW-146
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
## Provenance

- evals 2026-11-18T23:27:12.13211Z agent=codex fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=218 failing=0 passing=218
- audit 2026-11-18T23:27:29.822983Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

## Audit

No findings; checked FLOW-146

