---
bases:
    CHK-043: 82e0e810
module: checks
phase: applied
run: 20261110T001313-e0ad9b8a10464ed606bb9d69
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

Fail-first on the evals base: the required global hook returned no diagnostic
or error for exit 3. The silent module override's exit 7 was ignored beside a
global X01 violation. The existing launch, cancellation and timeout eval
passed, as did the nonzero hook with all attributed lines suppressed.

The implementation can be changed in `packages/checks/internal/engine/boundaries.go`,
which is writable during impl.
## Provenance
- evals 2026-11-10T00:26:44.876737Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=121 failing=2 passing=119
- audit 2026-11-10T00:26:59.668313Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-10T01:00:35.907148Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=787 failing=0 passing=787
## Audit

No findings; checked CHK-043
