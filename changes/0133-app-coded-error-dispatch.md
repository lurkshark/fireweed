---
module: app
---

## Motivation

Maintenance after 0128–0132: the CLI layer chooses diagnostic codes, exit statuses
and advice by matching error text in entry.go (stamp C08, show C05/U02),
lifecycle.go (coverage C05), ci.go and check_commits.go (SPEC.md parse C03),
pr_body.go (endpoint configuration), next.go and run.go (approval) and
project.go (projectFailureCode). Replace each with contracts.ErrorCode,
runner.IsApprovalRequired and workflow.IsEndpointConfigError. Codes, exit
statuses and messages stay the same, and existing app evals guard them.

## Interface

No interface changes.

## Added

## Modified

## Removed

## Decisions

## Rejected alternatives

## Notes

Show's U02 for a duplicated handle uses the CodedError subject's `new:` prefix
instead of the message. The wrappedCause regex that reads `CODE:` prefixes is
out of scope for this change.
