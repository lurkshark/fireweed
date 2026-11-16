---
module: workflow
fixes: [FLOW-019]
---

## Motivation

FLOW-019 requires Current and Approved to return the Start-approved snapshot.
A focused eval on current main showed that Start accepts audit.max_rounds:
1.0 as float64(1), then both claim reads return int(1). Workflow only treats
an integer max_rounds as a configured round limit, so claim reload changes
the setting's interpretation. The workflow drill exposed the reverse error
when an approved YAML integer version: 1 reloaded as float64(1) and blocked
valid CLI operations under APP-164 and PRJ-085.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Add a public fail-first test that starts a change with audit.max_rounds
written as 1.0 and compares the value and numeric kind in Start, Current and
Approved. Assert that the accepted value is preserved across claim reload.
Cover integral audit.max_rounds and version: 1 as regressions, including
successful CheckCompatibility after reload. Compare the values directly
before JSON encoding. Preserve the existing approved graph, canonical-root
and malformed live-input assertions.

Keep configuration acceptance and the meaning of audit.max_rounds at Start.
Preserve the ability to read existing valid claims. Leave the encoding
choice to implementation.
