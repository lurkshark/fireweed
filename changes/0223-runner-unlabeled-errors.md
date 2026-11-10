---
module: runner
---

## Motivation

SYSTEM.md's diagnostic-code convention says error text does not repeat its
code once the app reads codes only through `contracts.ErrorCode`, which
0214 (app-reads-error-codes) does. Runner's own errors, such as RUN-029's
C01, still begin with the code. Remove the labels so that, with 0220 to
0222, no fireweed error states its code twice.

## Interface

No signature changes. Error text no longer includes the code label.

## Added

## Modified

## Removed

## Decisions

## Notes

Maintenance: remove each leading `<code>: ` from runner error text and keep
the code in `contracts.CodedError`. This covers errors runner returns, not
retry prompts or drill reports, which quote verification output. 0219 moved
runner's evals off checks' labels, so all evals must pass unchanged. Run
after 0222.
