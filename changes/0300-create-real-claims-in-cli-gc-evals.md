---
module: app
eval_fix: [APP-150, APP-151]
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
