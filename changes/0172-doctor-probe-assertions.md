---
bases: {}
eval_fix:
    - CHK-057
    - CHK-060
module: checks
phase: audit
run: 20261105T215537-14c3ae392f33f060ba32aaf5
---

## Motivation

The owner approved correcting two report assertions exposed by the second
independent checks drill. The design's Language support section requires doctor
to prove the hook catches its canary and passes after removal. A successful
recovery does not show the planted canary was detected. Either singular or
plural boundary hook wording can identify the affected override.

## Interface

No public signature, spec, runtime implementation, hook or configuration edits.

## Added

## Modified

## Removed

## Decisions

## Notes

In Test_CHK_057_InternalEvalRequiresInterfaceViolation, distinguish the actual
internal_eval probe from recovery when checking for a passed probe. Both bare
and variant-suffixed probe names are used by conforming reports. Keep the
error-level C14 check and the independent undeclared-canary assertion.

In Test_CHK_060_OverridesMustDetectTheirCanaries, accept "boundary hook" and
"boundaries hook" for the boundary override. Require an error-level C14 naming
the affected override and its canary. Preserve the eval override subcase and
all cases and ID coverage. Limit edits to the necessary assertions and their
local support in the checks doctor eval file.

Follow the pinned controller's eval_fix lifecycle. Run fireweed context before
each worker step and the applicable fireweed verify before returning. Use the
separate restricted auditor. Workers never commit, advance or edit run/phase
or Provenance. Preserve prior evidence, counters, audit bounds and the pin.
No drill is part of this change.
## Provenance
- evals 2026-11-05T22:06:47.32438Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=114 failing=0 passing=114
