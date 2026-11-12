---
module: workflow
---

## Motivation

0230 delegates full verification to Done so the runner can verify each
successful attempt once. Done can currently return the same diagnostic code
from verification and from phase-transition checks. RUN-021 retries failed
verification, so the runner needs the origin of the error as well as its code.

## Interface

Add `workflow.IsVerificationFailure(err error) bool` to identify an error
from Done's full verification step, including when the caller wraps that error.

## Added

### new:done-verification-error
Done MUST mark errors returned by its full verification step so callers can distinguish them from precondition and phase-transition errors. Marking MUST preserve the original error text, cause and `contracts.ErrorCode`, along with the verification result. `IsVerificationFailure` MUST recognize the marker through error wrapping.

## Modified

## Removed

## Decisions

## Notes

Use an error wrapper with `Unwrap` and a public predicate, following the
endpoint configuration error pattern. Apply the marker only when the full
Verify call inside Done returns an error. Keep verification reports,
diagnostics, and direct Verify behavior intact.

Evals: exercise the public predicate on Done errors from failed full
verification and from checks outside verification. Cover a failure during
verification with its original cause and diagnostic, a wrapped marked error,
and a phase-transition rejection with a code that verification can also
return. An eval-phase interface stub can return false until implementation.
This workflow change precedes 0230-single-phase-verification.
