---
eval_fix:
    - APP-091
    - APP-092
    - APP-093
    - APP-130
    - APP-131
module: app
---

## Motivation

0222 removes workflow's diagnostic labels from error text. Its app verification
fails because the release and PR body fixtures inspect text to identify the
P02 and P01 errors returned by workflow. The assigned codes are already
available through `contracts.ErrorCode`.

## Interface

Audited app eval correction. Existing CLI behavior and coverage are preserved.

## Added

## Modified

## Removed

## Decisions

## Notes

In the release fixture, assert P02 through `contracts.ErrorCode` for the
`workflow.Current` call after release. In the PR body fixture, assert P01
through `contracts.ErrorCode` for the unavailable revision range and retain
the rejected revision in the diagnostic cause. Keep all CLI exit-status,
plain and JSON diagnostics, claim, ref, and dirty-work checks. This change
precedes 0222-workflow-unlabeled-errors.
