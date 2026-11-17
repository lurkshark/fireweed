---
already_satisfied:
    - FLOW-136
bases:
    FLOW-136: e2faee8a
module: workflow
---

## Motivation

An auditor needs a chance to correct a failed verification while attempts
remain, as required by the design's Retries rule and RUN-022. The workflow
drill report in commit `f6f5997` shows the RUN-048 fixture retaining its
permitted Audit edit and then stopping after Done returned an unmarked audit
zone violation. Next could not recognize the verification failure or retry.

## Interface

## Added

## Modified

### FLOW-136
Done MUST mark errors returned by its full verification step, including audit-phase zone violations, so callers can distinguish them from precondition and phase-transition errors. Marking MUST preserve the original error text, cause and `contracts.ErrorCode`, along with the verification result. `IsVerificationFailure` MUST recognize the marker through error wrapping.

## Removed

## Decisions

## Notes

Advance a public lifecycle fixture to audit, supply a valid Audit section and
add one file outside the audit-writable change document before verification.
Check that Verify rejects the file with its zone diagnostic. Assert that Done
returns the same verification result and error text, code and subject, and
that IsVerificationFailure recognizes the Done error through wrapping.
Assert the audit phase, claim and HEAD remain in place.

The existing implementation marks the audit verification error and preserves
its diagnostics. A focused probe passed on main's implementation and found
the missing marker and diagnostics on the retained reconstruction `16c778e`.
