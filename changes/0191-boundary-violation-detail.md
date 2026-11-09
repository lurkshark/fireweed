---
already_satisfied:
    - FLOW-062
bases:
    FLOW-062: 7f7f65c1
module: workflow
phase: audit
run: 20261109T021210-8d01ef9c76f997c265674d0c
---

## Motivation

Users and composing commands need the concrete boundary violation when full
implementation verification fails. Two supervised workflow reconstructions
returned only an aggregate count for X01. The existing public eval requires the
violation detail, and the current runtime already includes it. Make that
obligation explicit while preserving the established X02 hook error behavior.

## Interface

The existing error code, error precedence, structured diagnostics and workflow
API remain the same. Strengthen the existing eval to establish the clarified
X01 detail on unchanged runtime.

## Added

## Modified

### FLOW-062
Lifecycle operations MUST preserve assigned static, eval, zone and boundary
causes when returning their existing failures, except for lifecycle stale-base
classification P05.
When full implementation verification returns an X01 boundary failure, the
error MUST include the selected unsuppressed violation's module-attributed
message.

## Removed

## Decisions

## Notes

In the existing public FLOW-062 lifecycle causes eval, keep the X01 prefix
assertion and require the fixture's complete module-attributed violation message
(alpha: forbidden dependency) in the returned error. Preserve the other cases,
including the separate assertions that X02 returns the hook failure cause with a
single prefix and retains structured diagnostics. Do not require exact X01
formatting, aggregate counts or new ordering. All workflow evals must pass
against unchanged runtime to establish already_satisfied. In impl only the
controller applies the clarified contract; no runtime changes are needed.

Use the pinned controller and configured worker PATH and GOCACHE. Read context
before each held step and run pinned verify --quick before returning. The
outside controller owns full verification, phase commits and archive. Workers
never commit, call lifecycle mutators, or edit protected guidance or prior
evidence. Keep the separate restricted audit and existing bounds.
## Provenance
- evals 2026-11-09T02:18:28.514347Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=178 failing=0 passing=178
