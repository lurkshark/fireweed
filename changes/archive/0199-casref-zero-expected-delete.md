---
fixes:
    - GIT-010
module: gitstore
phase: applied
run: 20261110T053651-4ac1f514567e5a5195c1008e
---

## Motivation

The owner approved fixing the GIT-010 deletion mismatch found in the gitstore
drill. After 0198, a public API probe confirmed that an empty expected value
preserves the ref, but the equivalent explicit all-zero SHA-1 value still
deletes it. Both absence representations must preserve an existing ref.

## Interface

Keep the public `CASRef` signature. A delete that supplies the all-zero
SHA-1 value as its expected object uses the same absence check as an empty
expected value.

## Added

## Modified

## Removed

## Decisions

## Notes

Extend the existing GIT-010 public regression to exercise the explicit zero
object ID with a separate fixture or subcase. Assert the mismatch error and
that the ref still points to the original object. Preserve the absent-ref
no-op, matching nonzero deletion and the passing empty-expected regression.

Reuse the atomic absence check from 0198 for the explicit zero value.
Keep changes within the existing gitstore scope. Workers can use the pinned
`fireweed verify --quick` for iteration. The controller runs full verification.

`Test_GIT_010_DeleteWithAbsentExpectedRef/ZeroSHA1` failed on the evals base.
Deleting the existing ref returned success, and `Ref` reported it absent.
Both absent-ref no-ops and the empty-expected mismatch passed, as did the
matching nonzero deletion.

The existing atomic absence check is in
`packages/gitstore/internal/local/local.go`, which is writable during impl.
## Provenance
- evals 2026-11-10T05:38:20.739701Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=13 failing=2 passing=11
- audit 2026-11-10T05:38:37.155503Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-10T05:53:19.463804Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=933 failing=0 passing=933
## Audit

No findings; checked GIT-010
