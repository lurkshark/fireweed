---
already_satisfied:
    - FLOW-136
bases:
    FLOW-136: e2faee8a
module: workflow
phase: applied
run: 20261117T103247-9d80efbe136e2353b6ed4bc4
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
## Audit

No findings; checked FLOW-136

## Provenance

- evals 2026-11-17T10:37:31.47855Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=218 failing=0 passing=218
- audit 2026-11-17T10:37:51.496296Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-17T10:53:17.68697Z agent=codex fireweed=0.2.1 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=838 failing=0 passing=838

