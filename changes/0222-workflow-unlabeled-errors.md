---
bases:
    FLOW-133: 7de84872
module: workflow
phase: impl
run: 20261111T171536-ca2a2d5d9f81416ba56ea512
---

## Motivation

SYSTEM.md's diagnostic-code convention says error text does not repeat its
code once the app reads codes only through `contracts.ErrorCode`, which
0214 (app-reads-error-codes) does. Workflow's error text still begins with
the code, and FLOW-133 fixes the whole text including the label. Remove the
labels and restate FLOW-133.

## Interface

No signature changes. Error text no longer includes the code label.

## Added

## Modified

### FLOW-133
For a valid queue, Current(root, "") MUST return an error whose contracts ErrorCode is P02 and whose text is `worktree holds no change` when the calling worktree owns none of its change claims.

## Removed

## Decisions

## Notes

Remove each leading `<code>: ` from workflow error text and keep the code in
`contracts.CodedError`. Update workflow's own evals that still match a
leading P code (0211 moved the modified items' evals to
`contracts.ErrorCode`). Run after 0220 and 0221, so that wrapped project and
checks causes are already unlabeled.

`fireweed verify workflow-unlabeled-errors` accepted fail-first evidence.
All three FLOW-133 scenarios returned structured P02 and failed on the
leading label in the message. The workflow report recorded 170 passing
and 18 failing cases, including parent cases, with failures on the old
message labels.

The implementation-writable code is in
`packages/workflow/internal/lifecycle/lifecycle.go`, including the
coded-error helper, diagnostic conversion and boundary wrapping. CHK-025
still requires the checks cause to start with `P2-01:`, and runner evals
expect that cause text. The workflow eval checks its structured P2-01
and the fail-first message.
## Provenance
- evals 2026-11-11T17:27:45.768577Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=188 failing=18 passing=170
- audit 2026-11-11T17:28:02.38974Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked FLOW-133
