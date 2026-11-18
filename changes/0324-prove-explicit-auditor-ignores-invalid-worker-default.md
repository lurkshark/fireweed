---
eval_fix:
    - PRJ-088
module: project
phase: audit
run: 20261118T224115-d3eabc7ca8a540d264da7d02
---

## Motivation

PRJ-088 gives explicit audit.agent selection its own authority. The November 18
project reconstruction consulted the worker default first, so a broken worker
default suppressed a valid auditor in Doctor. The current owning eval uses a valid
worker default alongside the explicit auditor.
[Drill reports](https://github.com/lurkshark/fireweed/pull/25)

## Interface

## Added

## Modified

## Removed

## Decisions

## Notes

Extend Test_PRJ_088_ResolveAuditAliasOrInheritWorkerDefault with a listed explicit
audit command and an unlisted agent.default. Assert the selected command and that
inheritance is false. With audit.agent omitted in the same fixture, require C01
identifying agent.default. Retain the existing invalid-audit-reference and worker
inheritance assertions.
## Provenance

- evals 2026-11-18T22:42:05.696818Z agent=codex fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=326 failing=0 passing=326

