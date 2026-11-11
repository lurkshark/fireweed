---
blocked: impl verification attempts exhausted
module: project
phase: impl
run: 20261111T112501-4be838d1ab95fe6649a36d44
---

## Motivation

SYSTEM.md's diagnostic-code convention says error text does not repeat its
code once the app reads codes only through `contracts.ErrorCode`, which
0214 (app-reads-error-codes) does. Project's error text still begins with
the code, so the code appears twice and callers have two sources that can
disagree. Remove the labels.

## Interface

No signature changes. Error text no longer includes the code label.

## Added

## Modified

## Removed

## Decisions

## Notes

Maintenance: remove each leading `<code>: ` from project error text and keep
the code in `contracts.CodedError`. 0209 moved project's own evals to
`contracts.ErrorCode`, and 0218 moved workflow's, so all evals must pass
unchanged. Run after 0214.
