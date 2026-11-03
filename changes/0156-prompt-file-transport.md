---
bases:
    RUN-003: 7d803e4b
module: runner
phase: audit
run: 20261103T075827-096b27941f3bb219cec10904
---

## Motivation

ACCEPTANCE C3 explicitly selects {prompt_file} as an additional transport for
agent.command and audit.command, retaining {prompt} behavior. Passing the
prompt through a temporary file supports commands that should not receive
large prompt arguments. This runner portion covers workers, audits, fallback
and wrappers; doctor expansion and init custom-command acceptance follow in
their owning modules before C3 is complete.

## Interface

No public fields or function signatures need change. Expand {prompt_file} in
selected command templates and wrapper argv. The file contains the exact
request Prompt and its lifetime is scoped to the invocation. Retain all existing
result fields, logging, command selection and execution behavior.

## Added

### RUN-067
InvokeAgent MUST create one temporary file containing the exact request Prompt only when {prompt_file} occurs in the selected command or wrapper templates, and expand that placeholder to its absolute path. The file MUST remain available through command execution and be removed afterward, including after launch failure, nonzero exit, timeout or cancellation.

## Modified

### RUN-003
InvokeAgent MUST expand {prompt}, {prompt_file} and {git_common_dir} once in command templates and wrapper argv. Commands MUST accept shell strings or argv, quoting shell substitutions and preserving argv values literally.

## Removed

## Decisions

## Notes

Preserve every existing assertion, {prompt} substitution semantics, streams,
exit/usage classification, logging, timeouts, retries and isolation controls.
Only selected commands/wrappers determine whether a prompt file is needed.
The same file serves all {prompt_file} occurrences for one invocation. Do not
introduce prompt-file I/O for invocations that do not use the placeholder.
Use an external temporary-file location so audit/reconstruction input workspace
contents retain their existing restrictions; do not whitelist extra input files.

Use public API evals with bounded fake agents and representative assertions.
New prompt-bearing fixture scripts use bash or stdin, preserving C1. Keep own
verification below five minutes by reusing existing fixtures/assertions where
appropriate; avoid duplicated expensive lifecycle executions. Do not weaken
coverage, change bounds or add process instrumentation. Fail-first needs real
failed testcases for the Added and Modified IDs, not collection/compile errors.

Evals/impl workers run pinned fireweed context, verify with fireweed verify
--quick, poll live handles until terminal and return. Only the outside owning
controller advances phases, commits, applies or archives and performs normal
full verification. Never edit run/phase/Provenance or frozen evals/interfaces.

The separate restricted auditor edits only the existing Audit body, retaining
all other bytes. Findings use - ID: "exact clause" — missing assertion. Passing
audit writes No findings and Checked IDs. Re-audits inspect only previously
flagged IDs; retain earlier clearances without repeating their review.

## Audit

## Provenance
- evals 2026-11-03T08:27:40.623917Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=158 failing=14 passing=144
