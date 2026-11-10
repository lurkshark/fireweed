---
fixes:
  - GIT-010
module: gitstore
---

## Motivation

GIT-010 requires `CASRef` to delete only when the current ref matches the
expected value. The gitstore drill found that an empty expected value, which
means the ref should be absent, can instead delete an existing ref and report
success. A local public API probe reproduced the data loss. The owner
approved a fix before the assets drill.

## Interface

Keep the public `CASRef` signature. An empty-expected delete of an existing
ref returns an error and leaves its object ID intact.

## Added

## Modified

## Removed

## Decisions

## Notes

Use an independent temporary Git repository and a public GIT-010 eval.
Create a ref, call `CASRef` with empty expected and replacement values, and
assert an error and the original ref object ID. Confirm this case fails on
the current implementation. Preserve the existing absent-ref no-op,
matching deletion, and create/update mismatch behavior.

Keep the absence comparison and delete atomic through Git. Limit runtime
edits to gitstore. Leave the other drill findings for owner review.
