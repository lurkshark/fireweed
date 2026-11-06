---
bases: {}
fixes:
    - FLOW-055
module: workflow
---

## Motivation

The owner authorized recovery after the independent checks drill reproduced a
FLOW-055 failure. With a held unsupported phase, Verify and Done can return
Z01 from the zone checker before they report the required P01. A protected
worktree edit reproduces this precedence problem with the current product
through public inputs.

## Interface

Preserve the existing Verify and Done interfaces and the FLOW-055 P01
diagnostic for unsupported phases.

## Added

## Modified

## Removed

## Decisions

## Notes

Add one public fail-first FLOW-055 eval using the existing lifecycle fixture.
Retain the existing unsupported-phase test unchanged. Create a fixture with a
held review phase, then write a file under a path configured as protected.
Require both Verify and Done to return the P01 unsupported-phase cause even
when that edit would otherwise produce Z01. Confirm the new eval fails on the
current product because of Z01. Keep the fixture HEAD and held phase unchanged
after rejection.

In workflow Verify, reject unsupported held phases before asking the checks
module for zone diagnostics, while preserving the existing claim and frozen
state validation. Done uses Verify and must return the same cause. Preserve the
normal audit, evals and impl paths and their enforcement. Keep the runtime
change confined to workflow.

Use the pinned controller's context and Verify before returning. The outside
controller runs required phase verification and owns lifecycle commands,
commits and archive. Workers never advance phases or edit run, phase, or
Provenance metadata. The auditor uses only its restricted workspace and edits
only the existing Audit body, preserving all bytes outside it. If coverage is
adequate, the Audit body is exactly `No findings; checked FLOW-055`.
