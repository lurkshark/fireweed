---
module: checks
phase: applied
run: 20261111T160130-7263a1ff1c5883e76786f87a
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

`fireweed verify checks-unlabeled-errors` passed checks, runner and workflow
evals. App verification failed because `coverage` and `ci` still use
`lifecycleError` to read a text prefix, and `drill`, `next` and `run` still
use `wrappedCause`. Those commands report P01 for unlabeled C10, C11 and
E01 errors. The app must read the structured codes before this change can
pass full verification. CHK-025 still requires the P2-01 text prefix.
## Provenance
- impl 2026-11-11T17:11:56.953863Z agent=manual fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=806 failing=0 passing=806
