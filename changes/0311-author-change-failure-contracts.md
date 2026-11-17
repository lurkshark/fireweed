---
module: assets
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

### new:author-failure-contracts
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
