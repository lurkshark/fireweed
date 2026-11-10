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
require the code only as the first characters of the error text, and the
other coded items say neither. Most project errors carry their code only in
the text, so the app recovers it with a regular expression or by running
`project.Check` again (APP-037). The app drill lists this as a gap, and the
project drill's retry changed error text to match it. Make the code part of
the error, as CON-016 already allows.

## Interface

No signature changes. Errors that carry a code become (or wrap)
`contracts.CodedError`.

## Added

### new:coded-errors
Every error a public project function returns for a failure this spec assigns a D006 code MUST carry that code so that contracts.ErrorCode reports it. An error that wraps another module's coded error, for a failure this spec assigns no code, MUST keep the wrapped code.

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

Removing the code from every message now: the app still reads some codes
from text until 0214 applies, so the prefix would disappear before its
replacement exists. Once the app reads codes only through
`contracts.ErrorCode`, the text prefix is presentation and a later change
can drop it.

## Notes

Keep the existing message text, including any leading code, so the app's
current output does not change. The prefix is simply no longer the contract.

Evals: extend the existing tests for each coded error path (PRJ-004,
PRJ-024, PRJ-055, PRJ-056, PRJ-065, PRJ-066 and the Load and Stamp failures)
to assert `contracts.ErrorCode`. Do not build new fixtures where an existing
test already reaches the failure. Check returns Diagnostics, which already
carry codes, and is not affected.
