---
fixes:
    - FLOW-089
    - FLOW-128
module: workflow
phase: applied
run: 20261115T041910-31d4d7a888bda651e1d4b566
---

## Motivation

On stock Apple Git 2.30.1, `fireweed gc` fails to list worktrees because
`git worktree list --porcelain -z` is unsupported. `fireweed abandon` fails
after a phase commit because that Git rejects `git apply --cached --3way`.
The existing workflow and CLI evals reproduce both failures on a supported
macOS host.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Apply after `portable-drill-cleanup-eval`. Keep the existing public contracts
and use real temporary Git repositories for the regressions.

Extend the FLOW-089 coverage with a later commit that changes a distinct line
within the reverse patch's context. Abandon must undo the phase edit, keep
the later edit and remove the pending change in one commit. Retain existing
claim and attempt assertions. A scratch checkout and isolated index with
`git apply --index --3way --reverse` preserved the later edit on Apple Git
2.30.1 and Git 2.56.1. Keep reverse application isolated from the live
checkout and preserve the existing publish and rollback behavior.

Extend the FLOW-128 coverage with a live linked worktree whose path contains
a newline and an existing unregistered directory at the path prefix before
that newline. GC must keep the live claim and prune the unregistered claim.
Match worktree registration against the same Git common directory without
depending on `worktree list -z` on older Git. Legacy porcelain output can
contain a literal newline inside a path. Preserve the rest of GC's pruning
and repository-state guarantees.

Use Apple Git 2.30.1 on PATH for fail-first and passing Fireweed verification.
Rerun the focused workflow and CLI evals with a current Git. Use the same
behavior assertions across versions without version-string assertions, Git
shims or skipped cases.
## Audit

No findings; checked FLOW-089, FLOW-128

## Provenance

- evals 2026-11-15T04:25:31.039477Z agent=codex fireweed=0.2.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=211 failing=8 passing=203
- audit 2026-11-15T04:25:54.38572Z agent=codex fireweed=0.2.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-15T04:54:35.277328Z agent=codex fireweed=0.2.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=813 failing=0 passing=813

