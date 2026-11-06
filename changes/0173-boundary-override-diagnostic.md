---
bases: {}
eval_fix:
    - CHK-060
module: checks
phase: impl
run: 20261106T001159-5ce3e44967e4517f77a749dc
---

## Motivation

Complete the owner's approved CHK-060 diagnostic wording correction. The
Language support section of the design requires doctor to prove each hook and
name a hook that misses its canary. The second checks drill reported an
error-level C14 for `alpha boundaries hook missed canary internal_eval`. The
preceding audited correction still requires an additional literal `override`,
which rejects that diagnostic even though it names the module, hook and canary.

## Interface

The edit is limited to the existing checks doctor eval for the boundary override.

## Added

## Modified

## Removed

## Decisions

## Notes

Correct only the boundary subcase in
Test_CHK_060_OverridesMustDetectTheirCanaries. Keep the eval override subcase
and the other evals intact. Keep an error-level C14, the alpha module, the
boundary hook label and internal_eval in the same diagnostic. Accept either
`boundary hook` or `boundaries hook`, with or without the word `override`.

Review the actual predicate against these messages. Each must satisfy it:

- `alpha boundaries hook missed canary internal_eval`
- `internal_eval alpha boundaries override boundary hook: missed internal_eval canary`

A diagnostic for a global hook, another module, another canary, a warning, or a
different code cannot satisfy it. Keep case count and spec ID coverage. Verify
through the pinned controller and the separate restricted audit. The prior
run, its rejected audit submission, native proof attempts and all bounds remain
part of the record.

The worker runs pinned fireweed verify --quick before returning. The owning
controller runs full phase verification and advances the lifecycle.
The auditor edits only the existing Audit body. Preserve every byte outside it.

## Audit

No findings; checked CHK-060
## Provenance
- evals 2026-11-06T00:22:48.494475Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=114 failing=0 passing=114
- audit 2026-11-06T00:23:34.632819Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
