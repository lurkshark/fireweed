---
already_satisfied:
    - FLOW-019
bases:
    FLOW-019: 845a3cd8
module: workflow
phase: audit
run: 20261117T101251-65d12dd990fc766c58f4e0b4
---

## Motivation

Users need to resume a held change with the configuration that approved it.
The workflow drill report in commit `f6f5997` shows both APP-164 resume cases
returning C01 after the reconstructed workflow looked up the held path through
the live queue. A malformed working file stopped `next`, and an incompatible
working version pin stopped `run` before either could invoke the approved agent.

## Interface

## Added

## Modified

### FLOW-019
For a resolvable queued change held by the calling worktree, Current MUST return the project snapshot approved at Start. When supplied the held change's repository-relative path, Current MUST do so even when the live project configuration is malformed or incompatible with the running Fireweed version. Approved MUST expose the same snapshot independently of live project and queue parsing.

## Removed

## Decisions

## Notes

Use the public Start fixture and keep its queued change document valid.
Replace only the working fireweed.yaml with malformed YAML or valid YAML
whose version pin excludes this binary. Call Current with the started
change's repository-relative path. Assert that the returned path and claim
identify the same held change and that its configuration matches the
Start-approved snapshot. Keep the existing Approved tests.

The existing implementation resolves that path and loads the snapshot from
the claim. A focused probe passed on main's implementation and reproduced
both C01 errors on the retained reconstruction `16c778e`.
## Provenance

- evals 2026-11-17T10:16:54.281895Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=218 failing=0 passing=218

