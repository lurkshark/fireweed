---
eval_fix:
    - FLOW-056
module: workflow
phase: audit
run: 20261116T210217-37ddb831905583c27fde9bc7
---

## Motivation

Two worktrees contending for a change need a recognizable P02 cause and an
intact winning claim, as specified by FLOW-016 and FLOW-056. The workflow eval
also required the command name update-ref in the returned error. GIT-010
allows an explanatory CAS mismatch error, which the retained gitstore build
returned without that command name.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Keep the P02 assertions for unclaimed Current, contended Start and foreign
Current. Keep the assertions distinguishing a claim already held from a claim
owned by another worktree. Remove the dependency on the Git subcommand name
in the contention message and use the public claim state or the original ref
value to verify the winning claim survived the second Start. Preserve
FLOW-056's existing testcase and coverage.
## Provenance

- evals 2026-11-16T21:06:07.361967Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=211 failing=0 passing=211

