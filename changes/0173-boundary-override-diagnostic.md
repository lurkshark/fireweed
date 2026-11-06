---
bases: {}
eval_fix:
    - CHK-060
module: checks
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
