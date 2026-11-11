---
module: runner
phase: applied
run: 20261111T191822-42289c58ed1e75f1545afe7f
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
## Provenance
- impl 2026-11-11T19:39:35.702962Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=490 failing=0 passing=490
