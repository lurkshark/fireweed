---
bases: {}
module: project
---



## Motivation

Design History checks require inspecting protected/eval/interface/append-only
zones recorded around each selected commit, not today's mutable worktree. The
next observable command is NEW check-commits on an explicit nonempty fixture
range. Project currently exposes only working-tree Load; add the minimal public
revision snapshot dependency so checks can inspect recorded config/modules using
its own Git reader, without duplicating project parsing or making temporary
checkouts. Then immediately build checks history, its H01–H04 explanations and
app caller/actual NEW proof. Do not expand the rest of project first.

## Interface

Add thin public LoadAt(root, revision string) (Snapshot, error), using the
existing Snapshot/Module types and allowed gitstore dependency. Behavior lives
under project/internal; no new models, dependencies or process primitives.

## Added

### new:recorded-project
LoadAt MUST return the selected Git revision's project configuration and modules, using its recorded fireweed.yaml module globs and module SPEC.md documents rather than working-tree contents.

### new:recorded-defaults
LoadAt MUST apply the selected revision's module_defaults to omitted zone fields while preserving explicitly empty module zone lists.

### new:revision-paths
LoadAt MUST set Snapshot.Root to the canonical repository root and Module.Dir to each module's repository-relative directory.

### new:revision-errors
LoadAt MUST return an error when the revision cannot be resolved, its recorded fireweed.yaml is missing or malformed, or a matched recorded module directory lacks a readable SPEC.md.

## Modified



## Removed



## Decisions



## Notes

Use public project/gitstore/contracts interfaces and ordinary temporary Git repos
with explicit identities. One small historical/current fixture can distinguish
recorded config, module paths, defaults and explicit empty zones from dirty
working-tree edits/deletions; representative invalid revision/config/module-doc
inputs assert errors. No history checker, new command, model, graph rewrite,
background process, temporary checkout/ref or new dependency in this change.
Keep current Load behavior/assertions intact, share internal parsing as useful
without broad reorganization. Do not add aliasing/purity/syscall specs or per-spec
instrumentation. Workers never commit, change phase/claims or protected files.
Evals/interface plus a compiling stub first, freeze, separate restricted fixed
Sol/high audit, then implementation/full native verification. Keep own under5m,
full under15m and cases/spec near15, unchanged bounds. No bootstrap source/tests,
rebuild-v1, deleted/abandoned/saved implementation or old worker-log oracles.
Applied count1 after passing/pushed gate; this dependency chain returns to the
check-commits CLI proof as soon as needed interfaces work. No new drill authority.
