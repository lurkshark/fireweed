---
audit_rounds: 0
bases: {}
fixes:
    - AST-024
    - AST-026
module: assets
phase: audit
run: 20261028T101517-773eae3fd9c0ff3e5c9aecd8
---



## Motivation

D006 and the design require accurate causes and corrective actions for emitted
codes. NEW already demonstrated C02 for duplicate graph names, but Explain C02
says the cause is an unknown dependency or cycle. Existing P02 paths also report
missing or unreadable/undecodable claims, but its text only advises waiting for
a holder, which does not repair a missing or invalid claim. Correct these
specific inaccuracies without adding codes, features or validation conditions.

## Interface

No public interface change.

## Added



## Modified



## Removed



## Decisions



## Notes

Strengthen existing C02 and P02 explanation cases using requirePattern, without
new cases/fixtures/helpers. Retain every existing cause/action and unknown-code
assertion. C02 must also describe empty or duplicate module names/prefixes and
repairing those identities, alongside invalid dependencies/cycles. P02 must
include unreadable/undecodable or invalid claim data, recovery/review for it,
starting an eligible change when no claim exists, and coordinating ownership
when a claim is held elsewhere. The added assertions must genuinely fail before
implementation for AST-024 and AST-026.

Change only the C02/P02 text and, within the same lifecycle explanation group,
make P01 wording apply to the already demonstrated new-change operational
precondition as well as lifecycle conditions. Preserve existing message meanings
and actions that remain valid. P02 advice must not promise an unimplemented
command or recommend overwriting another owner's claim; missing/invalid claims
need the appropriate start or approved human recovery, while ownership conflicts
need coordination. Do not add unused codes/tests, matrices, new checks or source
changes outside explanation text. Controller owns phase commits and archive.

## Provenance

- evals 2026-10-28T10:18:09.074095Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=331 failing=5 passing=326
