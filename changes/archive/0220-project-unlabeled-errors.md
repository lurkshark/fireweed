---
module: project
phase: applied
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
## Provenance
- impl 2026-11-11T16:00:24.545103Z agent=manual fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=940 failing=0 passing=940
