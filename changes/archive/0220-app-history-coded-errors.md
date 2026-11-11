---
module: app
phase: applied
run: 20261111T154247-2e003823692d61e93d2ae25c
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
## Provenance
- impl 2026-11-11T15:53:21.408227Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=323 failing=0 passing=323
