---
bases: {}
module: checks
phase: audit
run: 20261103T092544-fc7c518c8bca8418b2986aee
---

## Motivation

ACCEPTANCE C3 adds {prompt_file} to configured worker and audit commands.
The design's Doctor agent probes execute those same templates independently
of runner, so they must accept this transport too. Keep existing fixed-reply
checks, fallback, diagnostics, logs, cleanup and {prompt} behavior intact.

## Interface

No public fields or signatures change. Doctor expands configured command
templates and scopes any prompt file to one executed probe.

## Added

### CHK-082
Doctor MUST expand {prompt}, {prompt_file} and {git_common_dir} once in selected worker/audit command templates and wrapper argv, quoting shell substitutions and preserving argv values literally.

### CHK-083
Doctor MUST create one temporary file containing the exact probe prompt only when a selected command or wrapper template uses {prompt_file}, and expand that placeholder to its absolute path. The file MUST remain available through execution and be removed afterward, including after launch failure, nonzero exit, timeout or cancellation.

## Modified

## Removed

## Decisions

## Notes

Preserve every existing assertion, command/fallback selection, fixed-reply
outcome, stream/log retention, diagnostic and scratch-cleanliness check.
No interface work is needed. Do not change runner, app, module SPEC/Decisions,
protected infrastructure or existing timeouts. This is C3 transport support,
not CK review or additional hardening.

Use public API evals and bounded bash fake agents. Reuse representative
fixtures/assertions and keep own Verify below five minutes. For new cleanup
checks, use a readiness signal for cancellation and a modest bounded timeout
that permits the fake script to start; avoid depending on subsecond process
startup. Preserve existing timeouts and all assertions. Do not duplicate
expensive Doctor executions or add instrumentation. Fail-first requires real
failed cases for both Added handles, not compile/collection failures.

Workers run pinned fireweed context before work, verify with fireweed verify
--quick, poll until terminal and return. They never commit, call lifecycle
mutators, edit run/phase/Provenance or modify frozen evals/interfaces. Only the
outside owning controller advances, commits, applies or archives, with normal
full verification.

The separate restricted auditor edits only the existing Audit body, retaining
all other bytes. Findings use - ID: "exact clause" — missing assertion. Passing
audit writes No findings and Checked IDs. Re-audits cover only previously
flagged IDs, preserving earlier clearances and all counters/bounds.

## Audit


## Provenance
- evals 2026-11-03T09:52:03.045Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=114 failing=7 passing=107
