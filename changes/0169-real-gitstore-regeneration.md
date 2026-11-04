---
bases: {}
module: gitstore
regenerate: true
---

## Motivation

ACCEPTANCE C8 and the design's Regeneration section require one real replacement
of gitstore's disposable implementation through the normal maintenance lifecycle.
Rebuild from retained specs, Decisions, interfaces and evals, preserving every
assertion and normal verification gate.

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

This is the milestone's single real regenerate change, not a drill. The owning
pinned controller deletes disposable files and alone applies or archives work.
Reconstruction inputs are retained gitstore contracts/evals and declared
contracts dependency public interfaces only. Old implementation, history,
binaries, saved attempts, bootstrap and dependency internals are excluded.
Verification may execute actual dependencies opaquely; they are not model inputs.

