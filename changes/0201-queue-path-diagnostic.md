---
already_satisfied:
    - PRJ-069
blocked: audit verification attempts exhausted
module: project
phase: audit
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
When `changes` is a regular file, Queue MUST return an error with code C03 and Check MUST return a non-warning diagnostic with code C03. Both MUST identify `changes`.

## Modified

## Removed

## Decisions

## Notes

Add one project-level public eval using the assigned ID and a valid fixture
with a regular file at `changes`. Assert the C03 cause from `Queue` and a
non-warning C03 diagnostic from `Check`, both identifying the queue path.
Preserve the existing app eval and the C13 behavior for ambiguous slugs.
The new case must pass against the existing implementation. Keep product
code, interfaces, and other project contracts unchanged.

Use the pinned controller's `fireweed verify --quick` for worker iteration.
The controller runs full verification before applying the change.
## Provenance
- evals 2026-11-10T16:21:00.431922Z agent=codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c model_reasoning_effort="high" -s workspace-write {prompt} fireweed=0.1.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=134 failing=0 passing=134
