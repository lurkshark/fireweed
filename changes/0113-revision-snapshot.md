---
audit_pending:
    - PRJ-047
audit_rounds: 1
bases: {}
module: project
phase: audit
run: 20261030T015320-239465605405eafd1c3fdef6
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

### PRJ-047
LoadAt MUST return the selected Git revision's project configuration and modules, using its recorded fireweed.yaml module globs and module SPEC.md documents rather than working-tree contents.

### PRJ-048
LoadAt MUST apply the selected revision's module_defaults to omitted zone fields while preserving explicitly empty module zone lists.

### PRJ-049
LoadAt MUST set Snapshot.Root to the canonical repository root and Module.Dir to each module's repository-relative directory.

### PRJ-050
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

## Provenance

- evals 2026-10-30T02:09:46.890494Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=594 failing=8 passing=586
- audit 2026-10-30T02:14:26.061808Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0
- evals 2026-10-30T02:16:45.188257Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=594 failing=8 passing=586

## Audit

- PRJ-047: "LoadAt MUST return the selected Git revision's project configuration and modules" — Missing assertions that Snapshot.Config retains the recorded module_defaults and that a recorded SPEC.md outside the configured module globs is excluded from returned modules.
