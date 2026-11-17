---
already_satisfied:
    - GIT-005
bases:
    GIT-005: 7cca5bc0
module: gitstore
phase: audit
run: 20261117T085953-bce7e6eae1f23398f110b2c8
---

## Motivation

Fireweed must recognize a worker timeout so it can stop reconstruction. The
fresh gitstore drill report in commit `028baa6` records a reconstruction that
passed its own timeout eval with a plain error. In the dependent evals, checks
lost the deadline cause in its diagnostics and runner treated the timeout as a
retryable failure. Callers need a timeout cause they can match.

## Interface

## Added

## Modified

### GIT-005
Run MUST use Args as the complete argv, with its first element naming the executable, and launch it with the requested directory, supplied environment and stdin. It MUST return captured stdout, stderr and exit code in Result, with process error text excluding the captured streams. Run MUST return an error on nonzero exit, timeout or context cancellation, terminating the launched process group on timeout or cancellation and returning an error matching `context.DeadlineExceeded` via `errors.Is` when `Request.Timeout` expires.

## Removed

## Decisions

## Notes

Extend the timeout subcase of `Test_GIT_005_RunCaptureExitTimeoutAndGroupCancel`
to assert `errors.Is(err, context.DeadlineExceeded)` when `Request.Timeout`
expires. Preserve the existing process-group and captured-output assertions.
Main already returns the deadline cause because `Run` applies `Request.Timeout`
through a context deadline and returns the context error on expiry.
## Provenance

- evals 2026-11-17T09:00:50.796597Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=14 failing=0 passing=14

