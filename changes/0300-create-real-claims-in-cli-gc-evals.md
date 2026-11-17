---
eval_fix:
    - APP-150
    - APP-151
module: app
phase: audit
run: 20261117T015215-03856053a22f2e45c354da6b
---

## Motivation

APP-150 and APP-151 check CLI garbage collection through its plain and JSON
results. The existing fixture builds the live claims from handwritten
workflow blobs, including uncanonicalized temporary paths. This couples the
app eval to workflow storage and path handling. Real claims exercise the
normal caller path, while the workflow module can check explicit path aliases.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Create the primary and linked worktree's live claims through workflow.Start
or CLI start. Give each repo branch a distinct eligible queued change and
commit all fixture inputs before starting. Capture the repository state after
claim setup and before GC. Keep the stale claim, archived attempt, old and
recent drill refs, unrelated refs, nearest-project assertions and repeated
GC checks. Assert the exact sorted pruned refs and verify that both native
live claims survive. Keep the plain and JSON checks and existing spec ID
coverage. FLOW-128 covers the explicit symlink-alias case.
## Provenance

- evals 2026-11-17T02:03:04.562739Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=403 failing=0 passing=403

