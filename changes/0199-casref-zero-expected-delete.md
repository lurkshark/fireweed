---
fixes:
  - GIT-010
module: gitstore
---

## Motivation

The owner approved fixing the GIT-010 deletion mismatch found in the gitstore
drill. After 0198, a public API probe confirmed that an empty expected value
preserves the ref, but the equivalent explicit all-zero SHA-1 value still
deletes it. Both absence representations must preserve an existing ref.

## Interface

Keep the public `CASRef` signature. A delete that supplies the all-zero
SHA-1 value as its expected object uses the same absence check as an empty
expected value.

## Added

## Modified

## Removed

## Decisions

## Notes

Extend the existing GIT-010 public regression to exercise the explicit zero
object ID with a separate fixture or subcase. Assert the mismatch error and
that the ref still points to the original object. Preserve the absent-ref
no-op, matching nonzero deletion and the passing empty-expected regression.

Reuse the atomic absence check from 0198 for the explicit zero value.
Keep changes within the existing gitstore scope. Workers can use the pinned
`fireweed verify --quick` for iteration. The controller runs full verification.
