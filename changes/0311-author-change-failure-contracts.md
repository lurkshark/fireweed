---
module: assets
phase: impl
run: 20261117T235631-5c9552d53a272e15f63ca764
---

## Motivation

The design places observable behavior at module interfaces and uses dependent
evals to reveal missing contracts during reconstruction. Recent drills passed
their own evals but broke downstream behavior when callers lost the information
they used to choose a next action after failure. Authors need a prompt to capture
those distinctions where the failure originates so they can be covered by the
producing module's evals.

## Interface

Add a paragraph to the built-in `Playbook("author-change")` output, after the
guidance on specifying observable behavior at the module interface.

## Added

### AST-072
The author-change playbook MUST direct authors to specify, in terms testable
through the producing interface, how callers or users can distinguish failure
outcomes that require different actions. It MUST direct authors to specify what
changes or remains available after failure when that affects the next action.
It MUST advise requiring exact wording only when the interface promises it.

## Modified

## Removed

## Decisions

## Notes

Suggested playbook paragraph

> When a caller or user must respond differently to two failure outcomes,
> specify how they can tell those outcomes apart at the interface that produces
> them. State what remains available or changes after failure when it affects
> the next action. Make the distinction testable through that interface.
> Require exact wording only when the interface promises it.

Check the guidance through `Playbook("author-change")`. Assert the meaning of
the guidance so equivalent phrasing can pass.

`Test_AST_072_AuthorChangeFailureGuidance` failed against the current playbook
because the failure guidance paragraph is absent. The controller's
`fireweed verify author-change-failure-contracts` passed the evals phase checks.

Implement the paragraph in `packages/assets/internal/content/content.go`, the
disposable file that defines the built-in author-change playbook.
## Provenance

- evals 2026-11-18T00:00:30.09084Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=153 failing=1 passing=152
- audit 2026-11-18T00:03:00.110265Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

## Audit

No findings; checked AST-072

