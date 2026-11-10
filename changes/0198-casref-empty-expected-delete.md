---
fixes:
    - GIT-010
module: gitstore
phase: impl
run: 20261110T045913-75966edcce447fcbf1ecc222
---

## Motivation

GIT-010 requires `CASRef` to delete only when the current ref matches the
expected value. The gitstore drill found that an empty expected value, which
means the ref should be absent, can instead delete an existing ref and report
success. A local public API probe reproduced the data loss. The owner
approved a fix before the assets drill.

## Interface

Keep the public `CASRef` signature. An empty-expected delete of an existing
ref returns an error and leaves its object ID intact.

## Added

## Modified

## Removed

## Decisions

## Notes

Use an independent temporary Git repository and a public GIT-010 eval.
Create a ref, call `CASRef` with empty expected and replacement values, and
assert an error and the original ref object ID. Confirm this case fails on
the current implementation. Preserve the existing absent-ref no-op,
matching deletion, and create/update mismatch behavior.

Keep the absence comparison and delete atomic through Git. Limit runtime
edits to gitstore. Leave the other drill findings for owner review.

`Test_GIT_010_DeleteWithEmptyExpectedRef` failed on the evals base. The call
returned success for an existing ref, and `Ref` reported that the ref was gone.
The absent-ref no-op passed.

The implementation can be changed in `packages/gitstore/internal/local/local.go`,
which is writable during impl.
## Provenance
- evals 2026-11-10T05:00:54.992058Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=11 failing=1 passing=10
- audit 2026-11-10T05:01:17.410766Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
## Audit

No findings; checked GIT-001, GIT-002, GIT-003, GIT-004, GIT-005, GIT-006, GIT-007, GIT-008, GIT-009, GIT-010
