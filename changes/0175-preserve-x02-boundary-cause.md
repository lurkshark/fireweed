---
bases: {}
fixes:
    - FLOW-062
module: workflow
phase: impl
run: 20261106T085253-cf6c6656e37fecebf85adaeb
---

## Motivation

The owner approved this FLOW-062 recovery after the independent checks drill
exposed a duplicated boundary hook launch cause in Verify. D006 assigns X02
to a required hook that cannot complete. FLOW-062 requires preserving the
assigned cause. A missing required boundary hook also exposes the composition
problem on the current product through public inputs.

## Interface

Use the existing public checks and workflow interfaces. The returned workflow
failure preserves the X02 cause, and the result still carries the full
structured boundary diagnostics.

## Added

## Modified

## Removed

## Decisions

## Notes

Add one public fail-first FLOW-062 eval using the lifecycle fixture with audit
disabled and its required boundary hook removed before Start. Advance to impl
and satisfy the implementation eval. Obtain the unprefixed cause and diagnostics
from public checks.Boundaries and verify the diagnostics name the affected
modules. Require full workflow Verify to return that cause with exactly one
X02 prefix and retain the original structured diagnostics. Confirm the new
case fails on the current implementation because Verify also includes the
first diagnostic's text. Keep the existing exact unlaunchable-hook assertion
and every other FLOW-062 case intact.

In implementation, use the selected hard X02 diagnostic for its code and
preserve the returned checks error as the sole cause. Add one prefix only
when needed. Retain structured diagnostics and error wrapping. Preserve
existing error precedence, X01 details, warning behavior and all other gates.
Limit runtime changes to workflow lifecycle composition.

Use the pinned controller's context and Verify before each worker returns.
The worker can use verify --quick and the owning controller runs the required
full checks and advances phases. Use an independent restricted auditor. The
auditor edits only the existing Audit body and preserves every byte outside it.
Workers never commit or edit lifecycle metadata. Preserve the pin, counters,
proof refs and deadline. Stop after affected clean candidate validation.

## Audit

No findings; checked FLOW-062
## Provenance
- evals 2026-11-06T09:09:48.285053Z agent=manual fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=161 failing=1 passing=160
- audit 2026-11-06T09:10:21.616203Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
