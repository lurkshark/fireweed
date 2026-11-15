---
eval_fix:
    - RUN-062
module: runner
phase: applied
run: 20261115T035529-2adc8c0ad73de5b5fd2321f7
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

Both focused RUN-062 evals passed with Apple Git 2.30.1 and Homebrew Git
2.56.1 selected through PATH. The existing assertions ran with both executables.
Full implementation-phase verification passed with Homebrew Git 2.56.1.

## Audit

No findings; checked RUN-062

## Provenance

- evals 2026-11-15T04:06:02.74736Z agent=codex fireweed=0.2.0 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=199 failing=0 passing=199
- audit 2026-11-15T04:06:25.290424Z agent=codex fireweed=0.2.0 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0
- impl 2026-11-15T04:18:45.117725Z agent=codex fireweed=0.2.0 playbook=5f856d6bcab308b9ff1a8a7c03b9d2547744112aba81dd30dea8f61ec8062b03 tests=602 failing=0 passing=602

