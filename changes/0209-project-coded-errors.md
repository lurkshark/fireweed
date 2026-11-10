---
bases:
    PRJ-055: c323db61
    PRJ-056: 1f45a62f
module: project
---

## Motivation

The design promises a stable code for every problem and output that agents
can read. Project states its codes two ways. PRJ-057, PRJ-058 and PRJ-059
require a code that `contracts.ErrorCode` reports, while PRJ-055 and PRJ-056
require the code only as the first characters of the error text. Most
project errors carry their code only in the text, so the app recovers it
with a regular expression or by running `project.Check` again (APP-037).
The app drill lists this as a gap, and the project drill's retry changed
error text to match it. SYSTEM.md's diagnostic-code convention makes the
code part of the error; state PRJ-055 and PRJ-056 that way.

## Interface

No signature changes. Errors that carry a code become (or wrap)
`contracts.CodedError`.

## Added

## Modified

### PRJ-055
An unknown module or an ID absent from both current module Specs and matching
archive edits MUST return an error whose contracts ErrorCode is U02 and whose
text names the rejected name.
A current ID with no matching archive edits MUST return an empty slice.
An ID in matching archive edits MUST remain queryable when it no longer exists
in current Specs.

### PRJ-056
Missing or malformed fireweed.yaml MUST produce an error whose contracts
ErrorCode is C01.
Malformed archived change documents or malformed Added/Modified item syntax
MUST produce an error whose contracts ErrorCode is C03 and whose text
identifies the archive path.

## Removed

## Decisions

## Rejected alternatives

A project spec item requiring every coded error to carry its code: the same
rule applies to every module, and copying it into each spec is what C18
warns about. It is a SYSTEM.md convention instead.

## Notes

Apply the SYSTEM.md diagnostic-code convention to every project error that
a spec assigns a code (PRJ-004, PRJ-024, PRJ-065, PRJ-066 and the Load and
Stamp failures), not only PRJ-055 and PRJ-056. Keep the existing message
text, including the leading code, until 0220.

Evals: in spec_history_test.go, replace the PRJ-055 and PRJ-056 prefix
checks (`U02: `, `C01: `, `C03: `) with `contracts.ErrorCode` checks, keeping
the checks for the rejected name and archive path.
