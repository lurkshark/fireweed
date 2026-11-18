---
eval_fix:
    - "PRJ-088"
module: project
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
inheritance is false. Use the same fixture with audit.agent omitted to confirm the
invalid worker reference then matters. Retain the existing invalid-audit-reference
and worker inheritance assertions.
