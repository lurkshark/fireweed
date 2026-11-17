---
already_satisfied:
    - FLOW-128
bases:
    FLOW-128: eca2e233
module: workflow
phase: impl
run: 20261117T013233-31284770071b3070b34201e5
---

## Motivation

GC's safety guarantee is to retain claims for live registered worktrees.
The retained workflow build compared recorded paths with canonical paths and
deleted two live claims when macOS temporary paths used a symlink alias.
Native Start records canonical ownership. GC already also accepts resolvable
aliases, which protects existing records under different path spellings.

## Interface

## Added

## Modified

### FLOW-128
GC MUST prune refs/fireweed/claims entries whose recorded worktree is missing
or is no longer registered with the same Git common directory, while
retaining claims for live registered worktrees. If the recorded path is a
symlink alias that resolves to such a worktree, GC MUST retain its claim.

## Removed

## Decisions

## Notes

Extend the public GC fixture with a symlink to a live registered worktree
and a recorded claim using the alias. Create the alias explicitly so the
case runs on macOS and Linux independently of the OS temporary directory.
Assert the claim survives and stays out of the pruned list. Retain existing
coverage for missing, unregistered and foreign worktrees and canonical
live claims. A focused rerun of the app GC case passed on the retained
reconstruction when the temporary directory had a canonical path.
## Audit

No findings; checked FLOW-128

## Provenance

- evals 2026-11-17T01:36:23.19437Z agent=codex fireweed=0.2.1 playbook=6a9c2d506363e3539cc66efd37bc6a06701249541dc4e4f01ed3aa6454bbb864 tests=215 failing=0 passing=215
- audit 2026-11-17T01:36:43.667481Z agent=codex fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0

