---
bases:
    CHK-043: 82e0e810
module: checks
---

## Motivation

The design requires the boundaries hook to pass during full implementation
verification. The workflow drill reported that a required hook which prints
an unattributed startup error and exits 3 can pass `checks.Boundaries`.
A local public API probe reproduced the false pass. The owner approved a
scoped fix so Verify and CI cannot treat this hook failure as a clean result.

## Interface

Use the existing `checks.Boundaries` return values. Report the hook failure
with X02 and an error that identifies the nonzero exit status.

## Added

## Modified

### CHK-043
Boundaries MUST report hook launch, cancellation or timeout failures as errors with a diagnostic whose Code is X02. The returned error MUST preserve the hook failure cause without a leading X02 prefix.
If a hook exits nonzero without emitting any module-attributed violation, Boundaries MUST return an X02 diagnostic and an error identifying the exit status.

## Removed

## Decisions

## Notes

Use bounded public API evals with deterministic fixture hooks. Prove genuine
fail-first for CHK-043 with a required global hook that prints only an
unattributed error to stderr and exits 3. Cover a silent nonzero module
override as well. Assert the non-warning X02 diagnostic and the returned error
with its exit status. Keep the existing launch, cancellation and timeout eval.

Classify each hook's output independently. Recognize module attribution before
applying baseline and reasoned opt-out suppression. Keep X01 for unsuppressed
attributed violations and keep the existing pass when all attributed lines are
suppressed. Preserve zero-exit handling, the existing error prefix convention,
and the public API. Limit runtime edits to checks and leave raw boundary output
retention outside this change.
