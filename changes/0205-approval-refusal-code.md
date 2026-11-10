---
bases:
    RUN-041: dcfedf6a
module: runner
---

## Motivation

The design requires a stable code for every problem. APP-075, APP-077 and
APP-082 require P01 when `next` or `run` refuses an unapproved change, but
RUN-041 only requires `IsApprovalRequired`. The runner drill's first
reconstruction put a different code in the refusal text, the app reported
it, and the drill passed only after the retry showed it the app failures.
Today the app gets P01 only by default, because the runner error carries no
code. Give the refusal its code at the runner boundary.

## Interface

The approval refusal returned by `Next` and `Run` becomes a coded error
readable through `contracts.ErrorCode`. No signature changes.

## Added

## Modified

### RUN-041
Next MUST return an error for which IsApprovalRequired reports true and contracts.ErrorCode returns P01 when the selected change file is on neither the local nor the shared base and AllowUnmerged is false.

## Removed

## Decisions

## Notes

Wrap the existing approval error in a `contracts.CodedError` with code P01
and the change path as subject, keeping `IsApprovalRequired` true through
the wrap. The app evals for APP-075, APP-077 and APP-082 must keep passing
unchanged.
