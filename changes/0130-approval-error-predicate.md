---
module: runner
---

## Motivation

App next and run choose their recovery advice by matching `requires approval` in
runner errors. The design gives unapproved changes their own handling (run skips
them without --allow-unmerged); callers need to recognise that refusal without
depending on message wording, which a regenerated runner may change.

## Interface

Add `IsApprovalRequired(err error) bool` to the runner public API, beside the
existing `IsUsageStop`.

Keep the public facade delegating to disposable implementation. During evals,
add the implementation stub in a new internal file so the impl step can replace
it without editing the frozen interface (the wiring repair from 0128).

## Added

### new:approval-required
Next MUST return an error for which IsApprovalRequired reports true when the selected change file is on neither the local nor the shared base and AllowUnmerged is false.

### new:approval-other-errors
IsApprovalRequired MUST report false for Next errors caused by anything other than a missing approval.

## Modified

## Removed

## Decisions

## Rejected alternatives
