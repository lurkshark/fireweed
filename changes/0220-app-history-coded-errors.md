---
module: app
---

## Motivation

0220 removes diagnostic prefixes from project error text. Its implementation
verification fails APP-135 because `show --history` still uses a message
prefix to choose the diagnostic code and exit status. Read the structured code
as required by the SYSTEM.md diagnostic-code convention.

## Interface

`show --history` preserves its APP-135 exit statuses, diagnostic codes,
messages, corrective guidance and plain and JSON output formats.

## Added

## Modified

## Removed

## Decisions

## Notes

Maintenance prerequisite for 0220-project-unlabeled-errors. Use
`contracts.ErrorCode` to classify errors returned by `project.SpecHistory`.
Strip one matching leading code from the message for compatibility with the
current project implementation. Preserve the P01 fallback for uncoded errors
and leave other commands' classification paths outside this change's scope.
Run the existing APP-135 evals. Verification of 0220 exercises the same evals
with the project errors carrying their codes only in metadata.
