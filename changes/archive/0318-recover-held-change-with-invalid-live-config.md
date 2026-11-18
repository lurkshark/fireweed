---
module: app
phase: applied
run: 20261118T201143-271a012fbbe75dba4a86f4ba
---

## Motivation

A malformed live fireweed.yaml blocked the CLI from showing held context or
releasing a claim, while the workflow APIs could do both. Users need those recovery
operations after a configuration edit goes wrong. APP-011 and APP-091 describe
operations on the held change.

## Interface

Keep the existing context and release arguments, output, and ownership rules.

## Added

### APP-168
After normal project discovery, context for a selected change owned by the caller's
worktree MUST return the held context even when live fireweed.yaml is malformed or
its binary pin excludes the running binary.

### APP-169
After normal project discovery, release for a selected claim owned by the caller's
worktree MUST release that claim even when live fireweed.yaml is malformed or its
binary pin excludes the running binary.

## Modified

## Removed

## Decisions

## Notes

Use the existing nested CLI fixture and start an owned change with compatible
configuration. Edit the live file after Start. Exercise malformed YAML and an
excluding live pin against context and release. Check that context returns the held
change and projected spec. Verify that release drops only the owned claim and
preserves HEAD, the index and edits. Give release an unrelated change selector
while the live configuration is malformed and confirm the caller's claim survives.
Reuse the existing selector and ownership assertions. Keep fresh start and other
live-config consumers on their existing admission paths.

Fail-first: `go test ./packages/app/evals -run '^Test_APP_16[89]_' -count=1 -v`
reproduced C01 with exit 2 for both commands with malformed YAML and with an
excluding binary pin. `fireweed verify recover-held-change-with-invalid-live-config`
passed in the evals phase. The implementation-writable admission path for both
commands is `packages/app/internal/entry/lifecycle.go`.
## Provenance

- evals 2026-11-18T20:21:42.83617Z agent=codex fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=409 failing=6 passing=403
- audit 2026-11-18T20:22:50.401113Z agent=manual fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-18T20:31:25.721926Z agent=codex fireweed=0.2.1 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=409 failing=0 passing=409

## Audit

No findings; checked APP-168, APP-169

