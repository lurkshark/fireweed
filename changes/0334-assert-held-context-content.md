---
eval_fix:
    - APP-168
module: app
phase: audit
run: 20261120T214929-cb5917c650aa8e8aec7bca59
---

## Motivation

Users need `context` to recover the held change and its projected requirements
after breaking their live configuration. In the workflow drill, both were
present and stable, but the APP-168 eval reported them missing when the
reconstructed renderer used different Markdown heading levels (local report
`9a94f39`). APP-168 and FLOW-005 specify the recovery and the content.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

In `Test_APP_168_ContextRecoversWithInvalidLiveConfiguration`, keep both live
configuration cases, the plain and full JSON calls, and the comparison with
the held context captured before the edit. Check for the complete held change
in the returned text. Exclude that matched block when checking for the
assigned spec ID and its normative text, since the held change already
contains the same requirement. Match the content without requiring a
particular Markdown heading level.
## Provenance

- evals 2026-11-20T21:57:34.333427Z agent=codex fireweed=0.2.4 playbook=5e05bacdae954379b468564c11429c5532016e26a7afbcd418f9bdb9441e151c tests=415 failing=0 passing=415

