---
already_satisfied:
    - GIT-005
bases:
    GIT-005: 9f878fe1
module: gitstore
---

## Motivation

The design's doctor agent probe shows the last output lines of a failed
command. CHK-065 sets a 20-line limit and CHK-081 provides a full log. The
retained gitstore build concatenated full stdout and stderr into the error.
Doctor displayed the error with its own tail, exposing lines outside the 20-line
window. Callers need separate captured streams to set
their own output limits.

## Interface

## Added

## Modified

### GIT-005
Run MUST launch the complete argv in Args, whose first element names the executable, without prepending any command. It MUST use the requested directory and supplied environment/stdin and return captured stdout, stderr and exit code in Result, with any process error's text excluding the captured streams. Run MUST return an error on nonzero exit, timeout or context cancellation, terminating the launched process group on timeout or cancellation.

## Removed

## Decisions

## Notes

Keep the existing argv, stream, exit, timeout and process-group assertions.
Use distinctive output markers to assert that Run returns both streams in
Result and omits them from the error text. Current main already separates
them. This clarification chooses the separate-result reading of GIT-005's
previous phrase "returning an error with that output". If the owner intends an
output-bearing Run error, the alternative is for checks to construct bounded
probe diagnostics independently of that error's text.
