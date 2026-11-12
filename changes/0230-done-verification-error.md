---
module: workflow
phase: impl
run: 20261112T001352-f31d70a3b6373e93f6711a3f
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

### FLOW-136
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
## Provenance
- evals 2026-11-12T00:23:23.262782Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=190 failing=2 passing=188
- audit 2026-11-12T00:23:39.936156Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked FLOW-136
