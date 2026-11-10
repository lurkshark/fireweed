---
module: checks
---

## Motivation

SYSTEM.md's diagnostic-code convention says error text does not repeat its
code once the app reads codes only through `contracts.ErrorCode`, which
0214 (app-reads-error-codes) does. Checks' error text still begins with
the code, so the code appears twice and callers have two sources that can
disagree. Remove the labels.

## Interface

No signature changes. Error text no longer includes the code label.

## Added

## Modified

## Removed

## Decisions

## Notes

Maintenance: remove each leading `<code>: ` from checks error text and keep
the code in `contracts.CodedError`. 0210 moved checks' own evals to
`contracts.ErrorCode`, and 0218 and 0219 moved workflow's and runner's, so
all evals must pass unchanged. Run after 0214.
