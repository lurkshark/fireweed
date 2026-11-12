---
module: app
phase: applied
run: 20261111T165401-8f19a5fc134b476361e012aa
---

## Motivation

0221 fails app verification when checks errors carry C10, C11 and E01 only
in metadata. Several CLI commands still infer the code from the error text
despite the migration specified by 0214 and the SYSTEM.md diagnostic-code
convention.

## Interface

The CLI preserves existing exit statuses, assigned codes, corrective messages,
warning diagnostics, artifact references and plain and JSON output formats.

## Added

## Modified

## Removed

## Decisions

## Notes

Maintenance prerequisite for 0221-checks-unlabeled-errors. Complete the
implementation of 0214. Use `contracts.ErrorCode` to classify returned errors
throughout the app, including the common formatter for lifecycle, coverage and
CI, the wrapped runner failures for drill, next and run, and project and
new-change failures. Keep each command's existing fallback for uncoded errors.
Use explicit structured defaults for configuration or spec failures where the
CLI assigns the code itself. Remove message-based classification and the
`project.Check` rerun that was used to infer a code. Message cleanup may remove
one leading label only when it matches the assigned code.

Preserve primary blocked causes when run also returns a time or usage stop.
Verify the existing app evals, including APP-040, APP-078, APP-084–APP-086,
APP-110, APP-112, APP-113 and APP-141. The follow-on verification of 0221
exercises the same cases against the unlabeled checks errors.
## Provenance
- impl 2026-11-11T17:06:17.962877Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=323 failing=0 passing=323
