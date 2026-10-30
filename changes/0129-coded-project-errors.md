---
module: project
---

## Motivation

App assigns C08 to Stamp failures by matching ` bases must be a mapping`, U02 to
missing pending changes by matching `pending change ... not found`, and project's
own queue code classifies item errors by message prefix. The design requires
stable diagnostic codes; these depend on unspecified wording, as the contracts
drill showed. Return coded errors using the contracts CodedError from
0128-coded-parse-errors.

## Interface

No new public functions. Stamp and the functions that look up pending changes
return errors that carry a contracts CodedError in their chain.

## Added

### new:stamp-bases-code
Stamp MUST return an error whose contracts ErrorCode is C08 when the change's bases frontmatter is not a mapping.

### new:missing-change-code
Stamp MUST return an error whose contracts ErrorCode is U02, with the requested change as subject, when no pending change matches it.

### new:spec-parse-code
Load MUST return an error preserving the contracts ErrorCode C03 or C05 when a module SPEC.md fails to parse.

## Modified

## Removed

## Decisions

## Rejected alternatives

## Notes

Implementation should also replace the message-prefix check in
internal/queue/cause.go with contracts.ErrorCode.
