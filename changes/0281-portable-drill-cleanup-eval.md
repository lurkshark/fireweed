---
eval_fix:
    - RUN-062
module: runner
---

## Motivation

The teardown of `Test_RUN_062_CleanupFailureCannotPass` fails on stock Apple
Git 2.30.1 because it runs `git worktree list --porcelain -z`. This prevents
the eval from verifying the expected cleanup error on a supported macOS host.
Change 0254 established that the fixture must handle both a registered worktree
and an unregistered directory left after cleanup fails.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Limit the correction to the runner eval teardown. Restore fixture permissions
and distinguish the registered scratch worktree from an unregistered leftover
using Git operations available on the stock host. Remove only the fixture's
scratch path and report unexpected teardown errors. Preserve the injected
cleanup failure and assertions that reconstruction reached the fixture, the
scratch directory remains after failed removal, and Drill returns an error
with a nonpassing result.

Apply this eval fix before `git-compatibility-gc-abandon`. Run Fireweed phase
verification with a current Git while the separate workflow product failures
are pending. Also run the focused RUN-062 evals with Apple Git 2.30.1 and
confirm the same assertions pass with both executables. Select Git through PATH
and keep assertions independent of Git version strings.
