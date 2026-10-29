---
audit_rounds: 0
bases: {}
module: workflow
phase: impl
run: 20261029T193658-a91309a40464314bde9455ee
---



## Motivation

The design's release <change> drops the local claim so its calling worktree can
stop owning a change without discarding work. NEW currently rejects release.
Add only the workflow operation needed by that caller. This is the fifth
application since review: after the required checkpoint, wire app immediately
and close the slice through an actual NEW start/release fixture showing retained
work/runtime state, no commit and no held claim. No other management commands.

## Interface

Add the thin public workflow.Release(root, change string) error facade, forwarding
to disposable lifecycle code. It accepts the same unambiguous path, filename or
slug selectors as Current. No model types or other public API changes. A new
compiling private stub is permitted during evals, with actual behavior in impl.

## Added

### FLOW-084
Release MUST drop only the named claim owned by the calling worktree while
preserving its working files, HEAD and recorded change state.

### FLOW-085
Release MUST reject a missing claim or a claim owned by another worktree with
P02, leaving refs unchanged.

### FLOW-086
Release MUST accept the same unambiguous change path, filename and slug
selectors as Current.

## Modified



## Removed



## Decisions



## Rejected alternatives



## Notes

Use public workflow/Git interfaces and representative temporary repositories.
For owned release retain dirty worker work and assigned run/phase/IDs in the
change, leave HEAD and work bytes/status intact and confirm the claim is gone.
For foreign ownership use an ordinary linked worktree and show the owner's
claim remains identical; missing-claim rejection must not create or move refs.
Test the three documented selector forms without adding an input matrix.
Keep existing helpers/assertions unchanged. Do not parse or construct private
claim serialization in evals; use public Start/Current and repository refs.

Implementation should use the existing canonical root/claim facilities and
exact compare-and-swap release, without verification, hooks, agents, phase
transitions, resets or commits. Preserve other worktrees and unrelated refs.
No new diagnostics (P02/U02 already describe ownership/lookup), generic lookup
framework, app wiring, abandon/commits/gc behavior, reorganization, other-module
or protected edits, instrumentation or drill. Evals/interfaces freeze in impl.
Workers never commit or advance workflow; coordinator owns native checks,
isolated audit/archive and the mandatory fifth-change review. Actual NEW release
proof follows the app caller after that checkpoint. Own under5m/full under15m;
do not raise bounds. The broader goal remains incomplete.

## Provenance

- evals 2026-10-29T19:51:18.209266Z agent=manual fireweed=0.1.0 playbook=b26da634732c7f1c7dc600115582a63cb8f82e1fd2bd1d522d608fe837a04ba7 tests=374 failing=6 passing=368
- audit 2026-10-29T19:52:54.322254Z agent=manual fireweed=0.1.0 playbook=0fb044cfaca1b9a844a25b4c86223177e52f2f35fb149eddf044e83822665c88 tests=0 failing=0 passing=0

## Audit

No findings
Checked IDs: FLOW-084, FLOW-085, FLOW-086
