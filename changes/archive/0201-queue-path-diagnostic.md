---
already_satisfied:
    - PRJ-069
module: project
phase: applied
run: 20261110T161922-8520833f1e629a2ca962e094
---

## Motivation

The app `list` contract preserves the cause code for queue failures. Its
APP-116/APP-117 eval requires C03 when `changes` is a regular file. The
project drill returned C13 for that input because the project contract left
the classification implicit. The owner approved recording the existing
C03 behavior at the project boundary before renewing the drill.

## Interface

Use the existing `project.Queue` and `project.Check` signatures and results.
The current implementation already satisfies the requirement.

## Added

### PRJ-069
When `changes` is a regular file, Queue MUST return an error identifying `changes`. Check MUST report a non-warning C03 diagnostic identifying `changes`.

## Modified

## Removed

## Decisions

## Notes

Add one project-level public eval using the assigned ID and a valid fixture
with a regular file at `changes`. Assert that `Queue` rejects the path with
an error identifying `changes` and `Check` reports a non-warning C03
diagnostic identifying the same path.
Preserve the existing app eval and the C13 behavior for ambiguous slugs.
The new case must pass against the existing implementation. Keep product
code, interfaces, and other project contracts unchanged.

Use the pinned controller's `fireweed verify --quick` for worker iteration.
The controller runs full verification before applying the change.
## Provenance
- evals 2026-11-10T16:21:00.431922Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=134 failing=0 passing=134
- audit 2026-11-10T16:42:53.701887Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c model_reasoning_effort="medium" -s workspace-write {prompt} fireweed=0.1.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-10T16:44:09.250682Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=925 failing=0 passing=925
## Audit

No findings; checked PRJ-069
