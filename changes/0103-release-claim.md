---
bases: {}
module: workflow
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

### new:release-owned-claim
Release MUST drop only the named claim owned by the calling worktree while
preserving its working files, HEAD and recorded change state.

### new:release-ownership-rejection
Release MUST reject a missing claim or a claim owned by another worktree with
P02, leaving refs unchanged.

### new:release-selectors
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
